package model

type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}
