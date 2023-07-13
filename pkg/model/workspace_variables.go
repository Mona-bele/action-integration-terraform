package model

import "fmt"

type WorkspaceVariables struct {
	Data *Data `json:"data"`
}

type WorkspaceVariablesList struct {
	Data []*Data `json:"data"`
}

type Data struct {
	ID         string     `json:"id"`
	Type       string     `json:"type"`
	Attributes Attributes `json:"attributes"`
}

type Attributes struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Hcl         bool   `json:"hcl"`
	Sensitive   bool   `json:"sensitive"`
}

func NewWorkspaceVariables(typeVar string, attributes Attributes) (*WorkspaceVariables, error) {
	wv := &WorkspaceVariables{
		Data: &Data{
			Type:       typeVar,
			Attributes: attributes,
		},
	}

	if err := wv.invalid(); err != nil {
		return nil, err
	}

	return wv, nil
}

func (wv *WorkspaceVariables) invalid() error {
	if wv.Data.Type == "" {
		return fmt.Errorf("type is not set")
	}
	if wv.Data.Attributes.Key == "" {
		return fmt.Errorf("key is not set")
	}
	if wv.Data.Attributes.Value == "" {
		return fmt.Errorf("value is not set")
	}
	if wv.Data.Attributes.Category == "" {
		return fmt.Errorf("category is not set")
	}

	return nil
}

func (wv *WorkspaceVariables) Update(wV *WorkspaceVariables) error {
	wv.Data.Attributes.Key = wV.Data.Attributes.Key
	wv.Data.Attributes.Value = wV.Data.Attributes.Value
	wv.Data.Attributes.Description = wV.Data.Attributes.Description
	wv.Data.Attributes.Category = wV.Data.Attributes.Category
	wv.Data.Attributes.Hcl = wV.Data.Attributes.Hcl
	wv.Data.Attributes.Sensitive = wV.Data.Attributes.Sensitive

	if err := wv.invalid(); err != nil {
		return err
	}

	return nil
}
