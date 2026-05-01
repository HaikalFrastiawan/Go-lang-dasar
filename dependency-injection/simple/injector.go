//go:build wireinject
// +build wireinject

package simple


func InitializeService() *SimpleService {
	wire.Build(
		NewSimpleRepository,NewSimpleService,
	)
	return nil
}


