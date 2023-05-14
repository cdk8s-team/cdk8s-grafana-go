package cdk8sgrafana


// Mode for accessing a data source.
// See: https://grafana.com/docs/grafana/latest/administration/provisioning/#example-data-source-config-file
//
type AccessType string

const (
	// Access via proxy.
	AccessType_PROXY AccessType = "PROXY"
	// Access directly (via server or browser in UI).
	AccessType_DIRECT AccessType = "DIRECT"
)

