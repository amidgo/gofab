package variables

// path constants for work
const (
	// path to fabric core.yaml package
	FABRIC_CFG_PATH = "/../config/"
	// path to fabric binary files
	BINARY_PATH = "../bin"
	// relative path in NETWORK_PATH directory to ./network.sh script
	NETWORK_SCRIPT = "./network.sh"
	// default output chaincode package
	DEFAULT_OUTPUT_PACKAGE = "contract.tar.gz"
	// default label
	DEFAULT_NAME = "contract"
	// default channel name
	DEFAULT_CHANNEL_NAME = "mychannel"
	// default sequence count
	DEFAULT_SEQUENCE = "1"
	// default init function in chaincode contract
	DEFAULT_INIT_FUN = "init"
)

const CORE_PEER_TLS_ENABLED = "true"

const (
	CORE_PEER_LOCALMSPID_ORG1        = "Org1MSP"
	CORE_PEER_TLS_ROOTCERT_FILE_ORG1 = "organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt"
	CORE_PEER_MSPCONFIGPATH_ORG1     = "organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp"
	CORE_PEER_ADDRESS_ORG1           = "localhost:7051"
)

// org2 export constants
const (
	CORE_PEER_LOCALMSPID_ORG2        = "Org2MSP"
	CORE_PEER_TLS_ROOTCERT_FILE_ORG2 = "organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt"
	CORE_PEER_MSPCONFIGPATH_ORG2     = "organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp"
	CORE_PEER_ADDRESS_ORG2           = "localhost:9051"
)

const (
	EXAMPLE_DEPLOY_COMMAND = "./gofab -opt deploy -c testchannel -out test.tar.gz -tar test.tar.gz -chaincode ../../../hyperledger/fabric-samples/asset-transfer-basic/chaincode-go/ -n test -seq 1 --init-fun initLedger"
	EXAMPLE_UPDATE_COMMAND = "./gofab -opt update -out test1.tar.gz -chaincode ../../../hyperledger/fabric-samples/asset-transfer-basic/chaincode-go/ -c testchannel -n test -seq 2 -tar test1.tar.gz"
)
