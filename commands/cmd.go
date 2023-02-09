package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/amidgo/gofab/variables"
)

func Launch() *exec.Cmd {
	return exec.Command(variables.NETWORK_SCRIPT, "up", "createChannel", "-c", variables.CHANNEL_NAME)
}
func Stop() *exec.Cmd {
	return exec.Command(variables.NETWORK_SCRIPT, "down")
}
func Build() *exec.Cmd {
	return exec.Command("peer", "lifecycle", "chaincode", "package", variables.OUTPUT_PACKAGE, "--path", variables.CHAINCODE_PATH, "--lang", "golang", "--label", variables.LABEL)
}
func OrgInstall() *exec.Cmd {
	return exec.Command("peer", "lifecycle", "chaincode", "install", variables.TAR_PATH)
}
func CheckInstalled() *exec.Cmd {
	return exec.Command("peer", "lifecycle", "chaincode", "queryinstalled")
}
func OrgApprove() *exec.Cmd {
	pwd := os.Getenv("PWD")
	return exec.Command("peer", "lifecycle", "chaincode", "approveformyorg", "-o", "localhost:7050", "--ordererTLSHostnameOverride", "orderer.example.com", "--channelID", variables.CHANNEL_NAME, "--name", variables.NAME, "--version", variables.SEQUENCE+".0", "--package-id", variables.CC_PACKAGE_ID, "--sequence", variables.SEQUENCE, "--tls", "--cafile", path.Join(pwd, "organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"))
}
func Commit() *exec.Cmd {
	pwd := os.Getenv("PWD")
	return exec.Command("peer", "lifecycle", "chaincode", "commit", "-o", "localhost:7050", "--ordererTLSHostnameOverride", "orderer.example.com", "--channelID", variables.CHANNEL_NAME, "--name", variables.NAME, "--version", variables.SEQUENCE+".0", "--sequence", variables.SEQUENCE, "--tls", "--cafile", path.Join(pwd, "organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"), "--peerAddresses", "localhost:7051", "--tlsRootCertFiles", path.Join(pwd, "organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt"), "--peerAddresses", "localhost:9051", "--tlsRootCertFiles", path.Join(pwd, "organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt"))
}

func InvokeChaincode() *exec.Cmd {
	pwd := os.Getenv("PWD")
	return exec.Command("peer", "chaincode", "invoke", "-o", "localhost:7050", "--ordererTLSHostnameOverride", "orderer.example.com", "--tls", "--cafile", path.Join(pwd, "organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"), "-C", variables.CHANNEL_NAME, "-n", variables.NAME, "--peerAddresses", "localhost:7051", "--tlsRootCertFiles", path.Join(pwd, "organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt"), "--peerAddresses", "localhost:9051", "--tlsRootCertFiles", path.Join(pwd, "organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt"), "-c", fmt.Sprintf(`{"function":"%s","Args":[]}`, variables.INIT_FUN))
}
