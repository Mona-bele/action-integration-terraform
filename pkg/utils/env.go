package utils

import (
	"flag"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type ConfigEnv struct {
	TFAPITOKEN     string `json:"tf_api_token"`
	TFWORKSPACE    string `json:"tf_workspace"`
	TFORGANIZATION string `json:"tf_organization"`
	TFRUNTYPE      string `json:"tf_run_type"`

	VariableType        string `json:"variable_type"`
	VariableKey         string `json:"variable_key"`
	VariableValue       string `json:"variable_value"`
	VariableCategory    string `json:"variable_category"`
	VariableSensitive   bool   `json:"variable_sensitive"`
	VariableHcl         bool   `json:"variable_hcl"`
	VariableDescription string `json:"variable_description"`

	DeleteVariable string `json:"delete_variable"`
	CreateVariable string `json:"create_variable"`
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

var (
	variableType, variableKey, variableValue, variableDescription, variableCategory string
	variableHcl, variableSensitive                                                  bool
)

func ArgsInput() {
	flag.StringVar(&variableType, "variable_type", "", "variable type")
	flag.StringVar(&variableKey, "variable_key", "", "variable key")
	flag.StringVar(&variableValue, "variable_value", "", "variable value")
	flag.StringVar(&variableDescription, "variable_description", "", "variable description")
	flag.StringVar(&variableCategory, "variable_category", "", "variable category")
	flag.BoolVar(&variableHcl, "variable_hcl", false, "variable hcl")
	flag.BoolVar(&variableSensitive, "variable_sensitive", false, "variable sensitive")
	flag.Parse()

	// set env variable from flag value

	SetGithubEnv("variable_type", variableType)
	SetGithubEnv("variable_key", variableKey)
	SetGithubEnv("variable_value", variableValue)
	SetGithubEnv("variable_description", variableDescription)
	SetGithubEnv("variable_category", variableCategory)
	variableHclParse := strconv.FormatBool(variableHcl)
	SetGithubEnv("variable_hcl", variableHclParse)
	variableSensitiveParse := strconv.FormatBool(variableSensitive)
	SetGithubEnv("variable_sensitive", variableSensitiveParse)
}
