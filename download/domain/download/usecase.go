package download

type UseCase interface {
	GetAbsUrlById(id string) string
}
