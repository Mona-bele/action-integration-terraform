package run

import (
	"context"
	"errors"
	"github.com/Mona-bele/action-integration-terraform/internal/utils"
	"log"
	"strings"
)

type ActionRun string

const (
	VariableCreate ActionRun = "variable_create" // action type variable create in terraform cloud
	VariableDelete ActionRun = "variable_delete" // action type variable delete in terraform cloud
)

type IRunBuilder interface {
	GetRunBuilder(ctx context.Context) // variableType string
}

func GetBuilder(builderType string) (IRunBuilder, error) {

	config, err := utils.LoadEnv(".")
	if err != nil {
		return nil, err
	}

	if strings.HasPrefix(builderType, "variable") == true {
		return NewVariableRunBuilder(config, builderType), nil
	}

	return nil, errors.New("invalid builder type")
}

type VariableRunBuilder struct {
	config       *utils.ConfigEnv
	variableType string
}

func NewVariableRunBuilder(config *utils.ConfigEnv, variableType string) IRunBuilder {
	return &VariableRunBuilder{
		config:       config,
		variableType: variableType,
	}
}

func (v VariableRunBuilder) initialVariableRunBuilder() (IVariable, error) {

	setting := NewSetting(v.config)

	client := NewNetHttpAdapter(v.config)

	workspace := NewWorkspaceRun(client, setting)

	vRun := NewVariableRun(client, v.config, workspace)

	return vRun, nil
}

func (v VariableRunBuilder) GetRunBuilder(ctx context.Context) {

	vR, err := v.initialVariableRunBuilder()
	if err != nil {
		log.Println(err)
	}

	switch v.variableType {

	case "variable_create":

		err = vR.Create(ctx)
		if err != nil {
			log.Println(err)
		}

	case "variable_delete":

		err = vR.Delete(ctx)
		if err != nil {
			log.Println(err)
		}

	default:
		log.Println("No variable type found")
	}

}
