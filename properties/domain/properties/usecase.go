package properties

type UseCase interface {
	SetKeyBaseUrl(baseUrl string)
	GetProperties() *Properties
}
