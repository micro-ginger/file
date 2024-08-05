package usecase

import "github.com/micro-ginger/file/properties/domain/properties"

type useCase struct {
	properties *properties.Properties
}

func New() properties.UseCase {
	uc := &useCase{
		properties: new(properties.Properties),
	}
	return uc
}

func (uc *useCase) SetKeyBaseUrl(baseUrl string) {
	uc.properties.KeyBaseUrl = baseUrl
}

func (uc *useCase) GetProperties() *properties.Properties {
	return uc.properties
}
