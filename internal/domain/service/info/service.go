package info

import jsonInfo "Leech-ru/internal/adapters/repository/json/info"

// Repository описывает поведение хранилища info.json
type jsonInfoRepository interface {
	Read() (*jsonInfo.Info, error)
	Write(info *jsonInfo.Info) error
}

type infoService struct {
	jsonInfoRepository jsonInfoRepository
}

func NewService(jsonInfoRepository jsonInfoRepository) *infoService {
	return &infoService{
		jsonInfoRepository: jsonInfoRepository,
	}
}
