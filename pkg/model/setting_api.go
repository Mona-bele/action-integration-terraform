package model

import "fmt"

type SettingAPI struct {
	TfApiToken      string
	TfOrgID         string
	TfWorkspaceName string
	TfRunType       string
}

func NewSettingAPI(apiToken, tfOrgID, tfWorkspaceName, tfRunType string) (*SettingAPI, error) {
	setting := &SettingAPI{
		TfApiToken:      apiToken,
		TfOrgID:         tfOrgID,
		TfWorkspaceName: tfWorkspaceName,
		TfRunType:       tfRunType,
	}

	if err := setting.invalid(); err != nil {
		return nil, err
	}

	return setting, nil
}

func (s *SettingAPI) invalid() error {
	if s.TfApiToken == "" {
		return fmt.Errorf("TF_API_TOKEN is not set")
	}
	if s.TfOrgID == "" {
		return fmt.Errorf("ID of organization is not set")
	}
	if s.TfWorkspaceName == "" {
		return fmt.Errorf("name of workspace is not set")
	}
	if s.TfRunType == "" {
		return fmt.Errorf("TF_RUN_TYPE is not set")
	}
	return nil
}

func TypeRun(runType string) string {
	switch runType {
	case "type_action":
		return "actions"
	default:
		return "runs"
	}
}
