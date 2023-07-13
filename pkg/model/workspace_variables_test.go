package model

import (
	"encoding/json"
	"testing"
)

func TestNewWorkspaceVariables(t *testing.T) {
	outputData := `{
			  "data": {
				"type":"vars",
				"attributes": {
				  "key":"some_key",
				  "value":"some_value",
				  "description":"some description",
				  "category":"terraform",
				  "hcl":false,
				  "sensitive":false
				}
			  }
			}`

	attributes := Attributes{
		Key:         "some_key",
		Value:       "some_value",
		Description: "some description",
		Category:    "terraform",
		Hcl:         false,
		Sensitive:   false,
	}

	outputDataWv := WorkspaceVariables{}

	wv, err := NewWorkspaceVariables("vars", attributes)
	if err != nil {
		t.Errorf("NewWorkspaceVariables() error = %v", err)
	}

	if wv.Data.Type != "vars" {
		t.Errorf("NewWorkspaceVariables() error = %v", err)
	}

	err = json.Unmarshal([]byte(outputData), &outputDataWv)
	if err != nil {
		t.Errorf("json.Unmarshal() error = %v", err)
	}

	if wv.Data.Type != outputDataWv.Data.Type {
		t.Errorf("Data_Input %v is not equal to Data_Output %v", wv.Data, outputDataWv)
	}
}

func TestListWorkspaceVariables(t *testing.T) {
	outputData := `{
			  "data": [
				{
				  "type":"vars",
				  "attributes": {
					"key":"some_key",
					"value":"some_value",
					"description":"some description",
					"category":"terraform",
					"hcl":false,
					"sensitive":false
				  }
				}
			  ]
			}`

	var outputDataWv WorkspaceVariablesList

	attributes := Attributes{
		Key:         "some_key",
		Value:       "some_value",
		Description: "some description",
		Category:    "terraform",
		Hcl:         false,
		Sensitive:   false,
	}

	wv, err := NewWorkspaceVariables("vars", attributes)
	if err != nil {
		t.Errorf("NewWorkspaceVariables() error = %v", err)
	}

	err = json.Unmarshal([]byte(outputData), &outputDataWv)
	if err != nil {
		t.Errorf("json.Unmarshal() error = %v", err)
	}

	if wv.Data.Type != outputDataWv.Data[0].Type {
		t.Errorf("Data_Input %v is not equal to Data_Output %v", wv.Data, outputDataWv.Data[0])
	}
}

func TestEmptyType(t *testing.T) {
	_, err := NewWorkspaceVariables("", Attributes{})
	if err == nil {
		t.Error("Type empty error")
	}

	if err.Error() != "type is not set" {
		t.Error("Type empty error")
	}
}
