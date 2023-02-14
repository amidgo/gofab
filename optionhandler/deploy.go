package optionhandler

import (
	"log"

	"github.com/amidgo/gofab/commands"
	"github.com/amidgo/gofab/variables"
)

func deploy() {
	if variables.CHAINCODE_PATH == "" {
		log.Fatal("chaincode paht in empty")
	}
	variables.RunBasicCmd(commands.RemoveDirectories())
	variables.RunBasicCmd(commands.Stop())
	variables.RunBasicCmd(commands.Launch())
	variables.RunBasicCmd(commands.Build())
	variables.RunFirstOrgCmd(commands.OrgInstall())
	variables.SetPackageID(variables.FirstOrgCmd(commands.CheckInstalled()))
	variables.RunSecondOrgCmd(commands.OrgInstall())
	variables.RunFirstOrgCmd(commands.OrgApprove())
	variables.RunSecondOrgCmd(commands.OrgApprove())
	variables.RunFirstOrgCmd(commands.Commit())
	variables.RunFirstOrgCmd(commands.InvokeChaincode())
	variables.RunBasicCmd(commands.EnrollAdmin())
	variables.RunBasicCmd(commands.RegisterAdmin())
}
