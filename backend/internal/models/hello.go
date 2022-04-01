package models

type HelloService interface {
	Hello(name string) (string, error)
}
