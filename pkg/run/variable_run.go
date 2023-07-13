package run

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Mona-bele/action-integration-terraform/internal/utils"
	"github.com/Mona-bele/action-integration-terraform/pkg/model"
	"log"
)

type IVariable interface {
	Create(ctx context.Context) error
	Read(ctx context.Context) (*model.WorkspaceVariables, error)
	Update(ctx context.Context) error
	Delete(ctx context.Context) error
}

type VariableRun struct {
	Config      *utils.ConfigEnv
	HttpClient  HttpClient
	IWorkspaces IWorkspaces
}

func (v VariableRun) verifyWorkspaceID(ctx context.Context) (string, error) {
	WorkspaceID, err := v.IWorkspaces.ReturnWorkspaceID(ctx)
	if err != nil {
		return "", err
	}
	return WorkspaceID, nil
}

func (v VariableRun) Read(ctx context.Context) (*model.WorkspaceVariables, error) {
	var data *model.WorkspaceVariablesList

	WorkspaceID, err := v.verifyWorkspaceID(ctx)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("/api/v2/workspaces/%s/vars", WorkspaceID)

	body, err := v.HttpClient.NewRequest("GET", endpoint, nil)

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	for _, variable := range data.Data {
		if variable.Attributes.Key == v.Config.VariableKey {
			return &model.WorkspaceVariables{
				Data: variable,
			}, nil
		}
	}

	return nil, nil
}

func (v VariableRun) Create(ctx context.Context) error {

	attributes := model.Attributes{
		Key:         v.Config.VariableKey,
		Value:       v.Config.VariableValue,
		Description: v.Config.VariableDescription,
		Category:    v.Config.VariableCategory,
		Hcl:         v.Config.VariableHcl,
		Sensitive:   v.Config.VariableSensitive,
	}

	vw, err := model.NewWorkspaceVariables(v.Config.VariableType, attributes)
	if err != nil {
		return err
	}

	WorkspaceID, err := v.verifyWorkspaceID(ctx)
	if err != nil {
		return err
	}

	endpoint := fmt.Sprintf("/api/v2/workspaces/%s/vars", WorkspaceID)

	body, err := json.Marshal(vw)
	if err != nil {
		return err
	}

	_, err = v.HttpClient.NewRequest("POST", endpoint, bytes.NewReader(body))
	if err != nil {
		return err
	}

	return nil
}

func (v VariableRun) Update(ctx context.Context) error {
	readKey, err := v.Read(ctx)
	if err != nil {
		return err
	}

	if readKey == nil {
		return fmt.Errorf("variable %s not found", v.Config.VariableKey)
	}

	wV := &model.WorkspaceVariables{
		Data: &model.Data{
			Type: v.Config.VariableType,
			Attributes: model.Attributes{
				Key:         v.Config.VariableKey,
				Value:       v.Config.VariableValue,
				Description: v.Config.VariableDescription,
				Category:    v.Config.VariableCategory,
				Hcl:         v.Config.VariableHcl,
				Sensitive:   v.Config.VariableSensitive,
			},
		},
	}

	err = readKey.Update(wV)
	if err != nil {
		return err
	}

	WorkspaceID, err := v.verifyWorkspaceID(ctx)
	if err != nil {
		return err
	}

	endpoint := fmt.Sprintf("/api/v2/workspaces/%s/vars/%s", WorkspaceID, readKey.Data.ID)

	body, err := json.Marshal(readKey)
	if err != nil {
		return err
	}

	_, err = v.HttpClient.NewRequest("PATCH", endpoint, bytes.NewReader(body))
	if err != nil {
		return err
	}

	return nil
}

func (v VariableRun) Delete(ctx context.Context) error {
	readKey, err := v.Read(ctx)
	if err != nil {
		return err
	}

	if readKey == nil {
		return fmt.Errorf("variable %s not found", v.Config.VariableKey)
	}

	WorkspaceID, err := v.verifyWorkspaceID(ctx)
	if err != nil {
		return err
	}

	endpoint := fmt.Sprintf("/api/v2/workspaces/%s/vars/%s", WorkspaceID, readKey.Data.ID)

	_, err = v.HttpClient.NewRequest("DELETE", endpoint, nil)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func NewVariableRun(client HttpClient, config *utils.ConfigEnv, workspaces IWorkspaces) IVariable {
	return &VariableRun{
		HttpClient:  client,
		Config:      config,
		IWorkspaces: workspaces,
	}
}
