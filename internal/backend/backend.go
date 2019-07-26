package backend

type Store interface {
	Open() error
	Read() (map[string]interface{}, error)
	Close() error
}
