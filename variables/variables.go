package variables

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

var CC_COUNT uint16 = 1

var (
	CHANNEL_NAME      string
	OUTPUT_PACKAGE    string
	CHAINCODE_PATH    string
	LABEL             string
	NAME              string
	TAR_PATH          string
	CC_PACKAGE_ID     string
	SEQUENCE          string
	INIT_FUN          string
	FUN               string
	TEST_NETWORK_PATH string
)

func SetPackageID(cmd *exec.Cmd) {
	b, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	sout := strings.Split(string(b), "\n")
	if len(sout) == 1 {
		log.Fatal("error while install chaincode")
	}
	for i, s := range sout {
		if i == 0 {
			continue
		}
		packLine := strings.Split(s, ":")
		if len(packLine) < 3 {
			continue
		}
		if strings.TrimSpace(packLine[3]) == LABEL {
			sep := strings.Split(s, ",")
			cc_package_id := strings.Join(strings.Split(sep[0], ":")[1:3], ":")
			CC_PACKAGE_ID = strings.TrimSpace(cc_package_id)
		}
	}
}
