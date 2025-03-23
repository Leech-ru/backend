package service_provider

type ServiceProvider struct {
	orderService orderService
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}
