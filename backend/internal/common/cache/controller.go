package cache

type Controller interface {
	ID() string
	Add(id string, item any) error
	Get(id string) (any, error)
	Remove(id string) error
	Exists(id string) bool

	Preload(data any) error
	AfterModify(id string) error
}
