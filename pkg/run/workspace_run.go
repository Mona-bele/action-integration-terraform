package run

import (
	"context"
	"encoding/json"
	"github.com/Mona-bele/action-integration-terraform/pkg/model"
)

type IWorkspaces interface {
	ShowWorkspace(ctx context.Context) (*model.Workspaces, error)
	ReturnWorkspaceID(ctx context.Context) (string, error)
}

type WorkspaceRun struct {
	ISetting   ISetting
	HttpClient HttpClient
}

func (w *WorkspaceRun) ReturnWorkspaceID(ctx context.Context) (string, error) {
	data, err := w.ShowWorkspace(ctx)
	if err != nil {
		return "", err
	}
	return data.Data.ID, nil
}

func (w *WorkspaceRun) ShowWorkspace(ctx context.Context) (*model.Workspaces, error) {
	data := model.Workspaces{}

	setting, err := w.ISetting.Setting()
	if err != nil {
		return nil, err
	}

	endpoint := "/api/v2/organizations/" + setting.TfOrgID + "/workspaces/" + setting.TfWorkspaceName

	body, err := w.HttpClient.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func NewWorkspaceRun(client HttpClient, ISetting ISetting) IWorkspaces {
	return &WorkspaceRun{
		HttpClient: client,
		ISetting:   ISetting,
	}
}
