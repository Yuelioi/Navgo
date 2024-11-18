package db

type Manager interface {
	Read() error
	Build() error
	Reduce() error
}
