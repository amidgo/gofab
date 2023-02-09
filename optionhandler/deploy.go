package optionhandler

import (
	"fmt"
	"log"
	"strings"

	"github.com/amidgo/gofab/commands"
	"github.com/amidgo/gofab/variables"
)

func deploy() {
	if variables.CHAINCODE_PATH == "" {
		log.Fatal("chaincode paht in empty")
	}
	variables.RunBasicCmd(commands.Stop())
	variables.RunBasicCmd(commands.Launch())
	variables.RunBasicCmd(commands.Build())
	variables.RunFirstOrgCmd(commands.OrgInstall())
	cmd := variables.FirstOrgCmd(commands.CheckInstalled())
	b, err := cmd.Output()
	sout := strings.Split(string(b), ",")
	if len(sout) < 2 {
		fmt.Println(sout)
		fmt.Println(sout)
		log.Fatal("WRONG CMD OUT")
	}
	cc_pkg_id := sout[0]
	cc_pkg_id = strings.Join(strings.Split(cc_pkg_id, ":")[2:4], ":")
	variables.CC_PACKAGE_ID = strings.TrimSpace(cc_pkg_id)
	if err != nil {
		fmt.Println(err)
		return
	}
	variables.RunSecondOrgCmd(commands.OrgInstall())
	variables.RunFirstOrgCmd(commands.OrgApprove())
	variables.RunSecondOrgCmd(commands.OrgApprove())
	variables.RunFirstOrgCmd(commands.Commit())
	variables.RunBasicCmd(commands.InvokeChaincode())
	// variables.RunFirstOrgCmd(commands.InvokeChaincode())
	// fmt.Println(variables.BasicCMD(commands.InvokeChaincode()))
}
