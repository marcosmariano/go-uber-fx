package usecase

import (
	"healthchecker/adapters/logger"
	"healthchecker/adapters/service"
	"net/http"
)

type CheckComponentUseCase struct {
	logger  logger.Logger
	service service.ComponentService
}

func NewCheckComponentUseCase(l logger.Logger, s service.ComponentService) CheckComponentUseCase {
	return CheckComponentUseCase{
		logger:  l,
		service: s,
	}
}

func IsUrlHealth(url string) (bool, error) {
	resp, err := http.Get(url)
	defer http.NoBody.Close()
	if err != nil {
		return false, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return true, nil
	} else {
		return false, nil
	}
}

func (c CheckComponentUseCase) Check() {
	c.logger.Info("CheckComponentUseCase :: Checking components status")

	for _, cp := range c.service.GetComponents() {
		c.logger.Debug("Component Name: ", cp.Name)
		result, err := IsUrlHealth(cp.Url)

		if result {
			c.logger.Info("Component is health")
			if !cp.IsHealth {
				cp.IsHealth = true

				update_error := c.service.UpdateComponent(cp)
				if update_error != nil {
					c.logger.Error("Error when update component")
				}
			}
		} else {
			c.logger.Info("Component is not health", err)
			if cp.IsHealth {
				cp.IsHealth = false

				update_error := c.service.UpdateComponent(cp)
				if update_error != nil {
					c.logger.Error("Error when update component")
				}
			}
		}
	}
}
