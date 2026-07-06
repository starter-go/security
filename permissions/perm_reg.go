package permissions

type Registration struct {
	DTO

	Priority int
}

type Registry interface {
	ListRegistrations() []*Registration
}
