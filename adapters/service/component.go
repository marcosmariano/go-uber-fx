package service

import (
	"healthchecker/adapters/logger"
	"healthchecker/adapters/repository"
	"healthchecker/domain/model"
)

type ComponentService struct {
	logger     logger.Logger
	repository repository.ComponentRepository
}

func NewComponentService(l logger.Logger, r repository.ComponentRepository) ComponentService {
	return ComponentService{
		logger:     l,
		repository: r,
	}
}

func (s ComponentService) SaveComponent(component model.Component) error {
	//return s.r.Save(&component)
	return nil
}

func (s ComponentService) UpdateComponent(component model.Component) error {
	s.logger.Debug("ComponentService :: updating component")
	//return s.r.Save(&component)
	return nil
}

func (s ComponentService) GetComponents() []model.Component {
	s.logger.Debug("ComponentService :: getting components")
	list := []model.Component{{ID: 1, Name: "EKS", Url: "https://eks.teste.com.br", Retry: 2, IsHealth: true}}
	return list
}
