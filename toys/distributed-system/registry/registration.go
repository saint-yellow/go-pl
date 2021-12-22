package registry

type Registration struct {
	ServiceName ServiceName
	ServiceURL string
	RequiredServices []ServiceName
	ServiceUpdateURL string
	HeartbeatURL string
}

type ServiceName string

const (
	LogService = ServiceName("Log")
	BusinessService = ServiceName("Business")
	PortalApplication = ServiceName("Portal")
)

type patchEntry struct {
	Name ServiceName
	URL string
}

type patch struct {
	Added []patchEntry
	Removed []patchEntry
}
