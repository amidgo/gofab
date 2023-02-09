package optionhandler

import (
	"fmt"
	"log"

	"github.com/amidgo/gofab/commands"
	"github.com/amidgo/gofab/variables"
)

func Handle(option string) {
	var opt variables.Option = variables.Option(option)
	switch opt {
	case variables.Launch:
		variables.RunBasicCmd(commands.Stop())
		variables.RunBasicCmd(commands.Launch())
	case variables.Stop:
		variables.RunBasicCmd(commands.Stop())
	case variables.Build:
		if variables.CHAINCODE_PATH == "" {
			log.Fatal("chaincode path is empty")
		}
		variables.RunBasicCmd(commands.Build())
	case variables.InstallIntoFirstOrg:
		variables.RunFirstOrgCmd(commands.OrgInstall())
		variables.RunFirstOrgCmd(commands.CheckInstalled())
	case variables.InstallIntoSecondOrg:
		variables.RunSecondOrgCmd(commands.OrgInstall())
		variables.RunSecondOrgCmd(commands.CheckInstalled())
	case variables.ApproveForFirstOrg:
		variables.RunFirstOrgCmd(commands.OrgApprove())
	case variables.ApproveForSecondOrg:
		variables.RunSecondOrgCmd(commands.OrgApprove())
	case variables.Deploy:
		deploy()
	case variables.UpdateChaincode:
		update()
	case variables.UpdateExample:
		fmt.Println(variables.EXAMPLE_UPDATE_COMMAND)
	case variables.DeployExample:
		fmt.Println(variables.EXAMPLE_DEPLOY_COMMAND)
	default:
		log.Fatalf("wrong value, '%s' --opt flag doesn`t exist", opt)
	}
}
