package run

import (
	"github.com/Mona-bele/action-integration-terraform/pkg/model"
	"github.com/Mona-bele/action-integration-terraform/pkg/utils"
)

type ISetting interface {
	Setting() (*model.SettingAPI, error)
}

type Setting struct {
	config *utils.ConfigEnv
}

func (s Setting) Setting() (*model.SettingAPI, error) {
	st, err := model.NewSettingAPI(s.config.TFAPITOKEN, s.config.TFORGANIZATION, s.config.TFWORKSPACE, s.config.TFRUNTYPE)
	if err != nil {
		return nil, err
	}

	return st, nil
}

func NewSetting(config *utils.ConfigEnv) ISetting {
	return &Setting{
		config: config,
	}
}
