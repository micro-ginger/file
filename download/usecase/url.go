package usecase

import "net/url"

func (uc *useCase) getUrlIdPrefix() string {
	path, err := url.JoinPath(uc.config.BaseUrl, "k")
	if err != nil {
		return ""
	}
	return path
}

func (uc *useCase) GetAbsUrlById(id string) string {
	path, err := url.JoinPath(uc.getUrlIdPrefix(), id)
	if err != nil {
		return ""
	}
	return path
}
