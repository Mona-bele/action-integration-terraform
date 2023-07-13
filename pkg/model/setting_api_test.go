package model

import (
	"testing"
)

func TestNewSettingAPI(t *testing.T) {
	sttApi, err := NewSettingAPI("token", "org", "workspace", "type_run")
	if err != nil {
		t.Errorf("NewSettingAPI() error = %v", err)
	}
	t.Log(sttApi)
}

func TestEmptySettingAPI(t *testing.T) {
	_, err := NewSettingAPI("", "", "", "")
	if err == nil {
		t.Errorf("NewSettingAPI() error = %v", err)
	}
}

func TestEmptyApiToken(t *testing.T) {
	_, err := NewSettingAPI("", "org", "workspace", "type_run")
	if err == nil {
		t.Errorf("NewSettingAPI() error = %v", err)
	}
}

func TestOrgEqualIsNotSet(t *testing.T) {
	_, err := NewSettingAPI("token", "", "workspace", "type_run")
	if err.Error() == "TF_API_ORG is not set" {
		t.Log("TF_API_ORG is not set")
	}
}
