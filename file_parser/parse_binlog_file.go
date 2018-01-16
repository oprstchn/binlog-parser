package file_parser

import (
	"github.com/golang/glog"
	"github.com/oprstchn/binlog-parser/database"
	"github.com/oprstchn/binlog-parser/parser"
	"fmt"
	"encoding/json"
)

type binlogParseFunc func(string) error

func CreateBinlogParseFunc(dbDsn string, consumerChain parser.ConsumerChain) binlogParseFunc {
	return func(binlogFilename string) error {
		return parseBinlogFile(binlogFilename, dbDsn, consumerChain)
	}
}

func parseBinlogFile(binlogFilename, dbDsn string, consumerChain parser.ConsumerChain) error {
	glog.V(2).Infof("Parsing binlog file %s", binlogFilename)

	db, err := database.GetDatabaseInstance(dbDsn)

	if err != nil {
		return err
	}

	defer db.Close()

	tableMap := database.NewTableMap(db)
	jsonBytes, _ := json.Marshal(tableMap)
	fmt.Println("tableMap", string(jsonBytes))

	glog.V(2).Info("About to parse file ...")

	return parser.ParseBinlog(binlogFilename, tableMap, consumerChain)
}
