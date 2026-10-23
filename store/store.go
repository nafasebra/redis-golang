package store

type Storer interface {
	Get(key string) (string, error)
	Set(key string, value string) error
	Keys() []string
	Delete(key string) error
	Len() int
}
