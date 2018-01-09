package parser

import (
	"os"
	"github.com/oprstchn/binlog-parser/database"
	"github.com/oprstchn/binlog-parser/parser/parser"
)

func ParseBinlog(binlogFilename string, tableMap database.TableMap, consumerChain ConsumerChain) error {
	if _, err := os.Stat(binlogFilename); os.IsNotExist(err) {
		return err
	}

	return parser.ParseBinlogToMessages(binlogFilename, tableMap, consumerChain.consumeMessage)
}
