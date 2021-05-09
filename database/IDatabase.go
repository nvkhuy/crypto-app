package database

type IDatabase interface {
	Open(configUrl string)
	Close()
}
