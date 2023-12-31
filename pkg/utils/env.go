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

	env.TFAPITOKEN = os.Getenv("INPUT_TF_API_TOKEN")
	env.TFWORKSPACE = os.Getenv("INPUT_TF_WORKSPACE")
	env.TFORGANIZATION = os.Getenv("INPUT_TF_ORGANIZATION")
	env.TFRUNTYPE = os.Getenv("INPUT_TF_RUN_TYPE")

	env.VariableType = os.Getenv("INPUT_VARIABLE_TYPE")
	env.VariableKey = os.Getenv("INPUT_TF_VARIABLE_KEY")
	env.VariableValue = os.Getenv("INPUT_TF_VARIABLE_VALUE")
	env.VariableDescription = os.Getenv("INPUT_TF_VARIABLE_DESCRIPTION")
	env.VariableCategory = os.Getenv("INPUT_TF_VARIABLE_CATEGORY")
	env.VariableSensitive, _ = strconv.ParseBool(os.Getenv("INPUT_TF_VARIABLE_SENSITIVE"))
	env.VariableHcl, _ = strconv.ParseBool(os.Getenv("INPUT_TF_VARIABLE_HCL"))

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
