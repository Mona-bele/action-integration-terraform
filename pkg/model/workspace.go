package model

type Workspaces struct {
	Data DataWorkspaces `json:"data"`
}

type DataWorkspaces struct {
	ID         string               `json:"id"`
	Type       string               `json:"type"`
	Attributes AttributesWorkspaces `json:"attributes"`
}

type AttributesWorkspaces struct {
	Name string `json:"name"`
}
