package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/go-mysql-org/go-mysql/mysql"
	"github.com/go-mysql-org/go-mysql/replication"
	"github.com/google/uuid"
)

var gSet mysql.GTIDSet

func processEvent(e *replication.BinlogEvent) error {
	if g, ok := e.Event.(*replication.GTIDEvent); ok {
		u, _ := uuid.FromBytes(g.SID)
		gtid := fmt.Sprintf("%s:%d", u.String(), g.GNO)

		var err error
		if gSet == nil || gSet.IsEmpty() {
			gSet, err = mysql.ParseMysqlGTIDSet(gtid)
		} else {
			err = gSet.Update(gtid)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	flag.Parse()

	p := replication.NewBinlogParser()

	gSets := make(map[string]mysql.GTIDSet, 0)

	for _, file := range flag.Args() {
		err := p.ParseFile(file, 4, processEvent)
		if err != nil {
			println(err.Error())
		}
		gSets[file] = gSet
		gSet = nil
	}

	binlogs := make([]string, len(gSets))
	for f := range gSets {
		binlogs = append(binlogs, f)
	}
	sort.Strings(binlogs)

	for _, f := range binlogs {
		if f == "" {
			continue
		}
		if gSets[f] == nil {
			fmt.Printf("%s\t%s\n", f, "<empty>")
		} else {
			fmt.Printf("%s\t%s\n", f, gSets[f].String())
		}
	}
}
