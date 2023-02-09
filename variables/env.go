package variables

import (
	"log"
	"os"
	"path"
)

func BasicEnv() []string {
	pwd := os.Getenv("PWD")
	pth := os.Getenv("PATH")
	if err := os.Setenv("PATH", path.Join(pwd, BINARY_PATH)+":"+pth); err != nil {
		log.Fatal(err)
	}
	env := os.Environ()
	basicVars := []string{
		`CORE_PEER_TLS_ENABLED=` + CORE_PEER_TLS_ENABLED,
		`FABRIC_CFG_PATH=` + path.Join(pwd, FABRIC_CFG_PATH),
	}
	env = append(env, basicVars...)
	return env
}

func FirstOrgEnv() []string {
	env := BasicEnv()
	pwd := os.Getenv("PWD")
	firstOrgEnv := []string{
		`CORE_PEER_LOCALMSPID=` + CORE_PEER_LOCALMSPID_ORG1,
		`CORE_PEER_TLS_ROOTCERT_FILE=` + path.Join(pwd, CORE_PEER_TLS_ROOTCERT_FILE_ORG1),
		`CORE_PEER_MSPCONFIGPATH=` + path.Join(pwd, CORE_PEER_MSPCONFIGPATH_ORG1),
		`CORE_PEER_ADDRESS=` + CORE_PEER_ADDRESS_ORG1,
	}
	env = append(env, firstOrgEnv...)
	return env
}

func SecondOrgEnv() []string {
	env := BasicEnv()
	pwd := os.Getenv("PWD")
	secondOrgEnv := []string{
		`CORE_PEER_LOCALMSPID=` + CORE_PEER_LOCALMSPID_ORG2,
		`CORE_PEER_TLS_ROOTCERT_FILE=` + path.Join(pwd, CORE_PEER_TLS_ROOTCERT_FILE_ORG2),
		`CORE_PEER_MSPCONFIGPATH=` + path.Join(pwd, CORE_PEER_MSPCONFIGPATH_ORG2),
		`CORE_PEER_ADDRESS=` + CORE_PEER_ADDRESS_ORG2,
	}
	env = append(env, secondOrgEnv...)
	return env
}
