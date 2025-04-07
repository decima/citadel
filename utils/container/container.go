package container

import (
	"errors"
	"fmt"
)

type container struct {
	servicesCollection map[string]any
}

func Add[T any](k T) {
	serviceName := fmt.Sprintf("%T", *new(T))
	Container.servicesCollection[serviceName] = k
}

func Get[T any]() (*T, error) {
	serviceName := fmt.Sprintf("%T", *new(T))
	service, exists := Container.servicesCollection[serviceName]
	if !exists {
		return nil, errors.New("service not found")
	}
	return service.(*T), nil
}

func ShouldGet[T any]() T {
	serviceName := fmt.Sprintf("%T", *new(T))
	service := Container.servicesCollection[serviceName]
	return service.(T)
}

var Container = &container{
	servicesCollection: make(map[string]any),
}
