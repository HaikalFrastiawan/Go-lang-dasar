//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func InitializedService(IsError bool) (*SimpleService, error) {
	wire.Build(NewSimpleRepository, NewSimpleService) 
	return nil, nil
}

func InitializedDatabaseRepository() *DatabaseRepository{
	wire.Build(
		NewDatabaseMongoDB,
		NewDatabasePostgreSQL,
		NewDatabaseRepository,
	)
	return nil
}