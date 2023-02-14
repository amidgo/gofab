package main

import (
	"flag"

	"github.com/amidgo/gofab/optionhandler"
	"github.com/amidgo/gofab/variables"
)

func main() {
	var (
		option string
	)
	flag.StringVar(&option, "opt", string(variables.Launch), "type of function in hyperledger network")
	flag.StringVar(&variables.CHANNEL_NAME, "c", variables.DEFAULT_CHANNEL_NAME, "name of channel")
	flag.StringVar(&variables.OUTPUT_PACKAGE, "out", variables.DEFAULT_OUTPUT_PACKAGE, "output package of chaincode")
	flag.StringVar(&variables.TAR_PATH, "tar", variables.DEFAULT_OUTPUT_PACKAGE, "path to tar package")
	flag.StringVar(&variables.CHAINCODE_PATH, "chaincode", "", "path to chaincode package")
	flag.StringVar(&variables.NAME, "n", variables.DEFAULT_NAME, "name for chaincode package")
	flag.StringVar(&variables.SEQUENCE, "seq", variables.DEFAULT_SEQUENCE, "count of contracts")
	flag.StringVar(&variables.INIT_FUN, "init-fun", variables.DEFAULT_INIT_FUN, "init function of chaincode")
	flag.StringVar(&variables.FUN, "fun", "{}", "callable funtion of chaincode")
	flag.StringVar(&variables.TEST_NETWORK_PATH, "network-path", "", "path to test-network directory")
	flag.Parse()
	variables.LABEL = variables.NAME + "_" + variables.SEQUENCE
	optionhandler.Handle(option)
}
