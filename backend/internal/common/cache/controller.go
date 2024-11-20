package cache

type Controller interface {
	ID() string
	Add(id string, item any) error
	Remove(id string) error
	Exists(id string) bool

	Preload() error
	AfterModify() error
}
