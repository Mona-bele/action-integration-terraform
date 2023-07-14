package utils

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type ConfigEnv struct {
	TFAPITOKEN     string `mapstructure:"tf_api_token"`
	TFWORKSPACE    string `mapstructure:"tf_workspace"`
	TFORGANIZATION string `mapstructure:"tf_organization"`
	TFRUNTYPE      string `mapstructure:"tf_run_type"`

	VariableType        string `mapstructure:"variable_type"`
	VariableKey         string `mapstructure:"variable_key"`
	VariableValue       string `mapstructure:"variable_value"`
	VariableCategory    string `mapstructure:"variable_category"`
	VariableSensitive   bool   `mapstructure:"variable_sensitive"`
	VariableHcl         bool   `mapstructure:"variable_hcl"`
	VariableDescription string `mapstructure:"variable_description"`

	DeleteVariable string `mapstructure:"delete_variable"`
	CreateVariable string `mapstructure:"create_variable"`
}

func LoadEnv(path string) (*ConfigEnv, error) {
	var env ConfigEnv

	_ = godotenv.Load()

	env.TFAPITOKEN = os.Getenv("tf_api_token")
	env.TFWORKSPACE = os.Getenv("tf_workspace")
	env.TFORGANIZATION = os.Getenv("tf_organization")
	env.TFRUNTYPE = os.Getenv("tf_run_type")

	env.VariableType = os.Getenv("variable_type")
	env.VariableKey = os.Getenv("variable_key")
	env.VariableValue = os.Getenv("variable_value")
	env.VariableDescription = os.Getenv("variable_description")
	env.VariableCategory = os.Getenv("variable_category")
	env.VariableSensitive, _ = strconv.ParseBool(os.Getenv("variable_sensitive"))
	env.VariableHcl, _ = strconv.ParseBool(os.Getenv("variable_hcl"))

	env.DeleteVariable = os.Getenv("delete_variable")
	env.CreateVariable = os.Getenv("create_variable")

	return &env, nil
}
