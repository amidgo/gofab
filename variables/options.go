package variables

type Option string

const (
	Launch               Option = "launch"
	Stop                 Option = "stop"
	Deploy               Option = "deploy"
	Build                Option = "build"
	InstallIntoFirstOrg  Option = "first-install"
	InstallIntoSecondOrg Option = "second-install"
	ApproveForFirstOrg   Option = "first-approve"
	ApproveForSecondOrg  Option = "second-approve"
	UpdateChaincode      Option = "update"
	DeployExample        Option = "deploy-example"
	UpdateExample        Option = "update-example"
)
