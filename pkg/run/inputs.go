package run

type Inputs struct {
	TfApiToken     string
	TfWorkspace    string
	TfOrganization string
	TfRunType      string

	VariableType        string
	VariableKey         string
	VariableValue       string
	VariableCategory    string
	VariableSensitive   bool
	VariableHcl         bool
	VariableDescription string

	DeleteVariable string
	CreateVariable string
}
