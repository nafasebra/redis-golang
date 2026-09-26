package store

type Storer interface {
	Get(key string) (string, error)
	Set(key string, value string) error
	keys() []string
	Delete(key string)
	Len() int
}
