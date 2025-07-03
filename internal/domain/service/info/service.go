package info

import (
	jsonInfo "Leech-ru/internal/adapters/repository/json/info"
)

// Repository описывает поведение хранилища info.json
type jsonInfoRepository interface {
	Read() (*jsonInfo.Info, error)
	Write(info *jsonInfo.Info) (*jsonInfo.Info, error)
}

type infoService struct {
	jsonInfoRepository jsonInfoRepository
	jsonInfoConfig     jsonInfoConfig
}

type jsonInfoConfig interface {
	PathToJsonFile() string
}

func NewService(jsonInfoConfig jsonInfoConfig) *infoService {
	return &infoService{
		jsonInfoRepository: jsonInfo.NewRepository(jsonInfoConfig.PathToJsonFile()),
		jsonInfoConfig:     jsonInfoConfig,
	}
}
