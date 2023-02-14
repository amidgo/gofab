package optionhandler

import (
	"log"

	"github.com/amidgo/gofab/commands"
	"github.com/amidgo/gofab/variables"
)

func update() {
	if variables.CHAINCODE_PATH == "" {
		log.Fatal("chaincode is empty")
	}
	variables.RunBasicCmd(commands.Build())
	variables.RunFirstOrgCmd(commands.OrgInstall())
	variables.SetPackageID(variables.FirstOrgCmd(commands.CheckInstalled()))
	variables.RunSecondOrgCmd(commands.OrgInstall())
	variables.RunFirstOrgCmd(commands.OrgApprove())
	variables.RunSecondOrgCmd(commands.OrgApprove())
	variables.RunFirstOrgCmd(commands.Commit())
	variables.RunFirstOrgCmd(commands.InvokeChaincode())
}
