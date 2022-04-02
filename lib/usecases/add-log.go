package usecases

import (
	"clean-architeture/lib/entities"
	"clean-architeture/lib/ports/repositories"
	usecase "clean-architeture/lib/ports/usecases"
)

type addLog interface {
	usecase.AddLogUsecases
}

type service struct {
	add repositories.AddLogRepository
}

func NewAddLog(add repositories.AddLogRepository) addLog {
	return &service{add}
}

func (s service) Add(data []byte) usecase.Response {
	log, err := entities.Create(data)
	if err != nil {
		return usecase.Response{"", err}
	}
	result := s.add.Add(log)
	return usecase.Response{result, nil}
}
