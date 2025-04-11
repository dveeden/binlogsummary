package main

import (
	"flag"
	"fmt"
	"log/slog"

	"github.com/go-mysql-org/go-mysql/mysql"
	"github.com/go-mysql-org/go-mysql/replication"
	"github.com/google/uuid"
)

var gSet mysql.GTIDSet

func processEvent(e *replication.BinlogEvent) error {
	if g, ok := e.Event.(*replication.GTIDEvent); ok {
		u, _ := uuid.FromBytes(g.SID)
		gtid := fmt.Sprintf("%s:%d", u.String(), g.GNO)
		slog.Debug("processing gtid", "gtid", gtid)

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

	for _, file := range flag.Args() {
		slog.Info("parsing file", "filename", file)
		err := p.ParseFile(file, 4, processEvent)
		if err != nil {
			println(err.Error())
		}
		if gSet != nil {
			fmt.Println(gSet.String())
			gSet = nil
		}
	}
}
