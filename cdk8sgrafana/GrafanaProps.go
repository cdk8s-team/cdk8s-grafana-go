package cdk8sgrafana


type GrafanaProps struct {
	// Default admin password.
	// Default: "secret".
	//
	AdminPassword *string `field:"optional" json:"adminPassword" yaml:"adminPassword"`
	// Default admin username.
	// Default: "root".
	//
	AdminUser *string `field:"optional" json:"adminUser" yaml:"adminUser"`
	// Default data source - equivalent to calling `grafana.addDataSource`.
	// Default: - no data source added.
	//
	DefaultDataSource *DataSourceProps `field:"optional" json:"defaultDataSource" yaml:"defaultDataSource"`
	// Specify a custom image for Grafana.
	// Default: "public.ecr.aws/ubuntu/grafana:latest"
	//
	Image *string `field:"optional" json:"image" yaml:"image"`
	// Create an ingress to provide external access to the Grafana cluster.
	// Default: true.
	//
	Ingress *bool `field:"optional" json:"ingress" yaml:"ingress"`
	// Labels to apply to all Grafana resources.
	// Default: - { app: "grafana" }.
	//
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Namespace to apply to all Grafana resources.
	//
	// The Grafana Operator must be
	// installed in this namespace for resources to be recognized.
	// Default: - undefined (will be assigned to the 'default' namespace).
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Require login in order to view or manage dashboards.
	// Default: false.
	//
	RequireLogin *bool `field:"optional" json:"requireLogin" yaml:"requireLogin"`
	// Type of service to be created (NodePort, ClusterIP or LoadBalancer).
	// Default: ClusterIP.
	//
	ServiceType *string `field:"optional" json:"serviceType" yaml:"serviceType"`
}

