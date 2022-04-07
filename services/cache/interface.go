package services

type Cache interface {
	Get(key string) (value string, err error)
	Set(key string, value string) error
}
