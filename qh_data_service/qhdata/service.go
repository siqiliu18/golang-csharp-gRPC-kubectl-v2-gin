package qhdata

// Service is service
type Service struct {
	Name string
}

// NewService is new service
func NewService() (Service, error) {
	svc := Service{
		Name: "qh_data_service",
	}
	return svc, nil
}
