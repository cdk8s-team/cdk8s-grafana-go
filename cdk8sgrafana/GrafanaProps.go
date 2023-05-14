package cdk8sgrafana


type GrafanaProps struct {
	// Default admin password.
	AdminPassword *string `field:"optional" json:"adminPassword" yaml:"adminPassword"`
	// Default admin username.
	AdminUser *string `field:"optional" json:"adminUser" yaml:"adminUser"`
	// Default data source - equivalent to calling `grafana.addDataSource`.
	DefaultDataSource *DataSourceProps `field:"optional" json:"defaultDataSource" yaml:"defaultDataSource"`
	// Specify a custom image for Grafana.
	Image *string `field:"optional" json:"image" yaml:"image"`
	// Create an ingress to provide external access to the Grafana cluster.
	Ingress *bool `field:"optional" json:"ingress" yaml:"ingress"`
	// Labels to apply to all Grafana resources.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Namespace to apply to all Grafana resources.
	//
	// The Grafana Operator must be
	// installed in this namespace for resources to be recognized.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Require login in order to view or manage dashboards.
	RequireLogin *bool `field:"optional" json:"requireLogin" yaml:"requireLogin"`
	// Type of service to be created (NodePort, ClusterIP or LoadBalancer).
	ServiceType *string `field:"optional" json:"serviceType" yaml:"serviceType"`
}

