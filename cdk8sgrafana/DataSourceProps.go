// Grafana construct for cdk8s.
package cdk8sgrafana


type DataSourceProps struct {
	// Access type of the data source.
	Access AccessType `field:"required" json:"access" yaml:"access"`
	// Name of the data source.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Type of the data source.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Description of the data source.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Labels to apply to the kubernetes resource.
	//
	// When adding a data source to a Grafana instance using `grafana.addDataSource`,
	// labels provided to Grafana will be automatically applied. Otherwise,
	// labels must be added manually.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Namespace to apply to the kubernetes resource.
	//
	// When adding a data source to a Grafana instance using `grafana.addDataSource`,
	// the namespace will be automatically inherited.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// URL of the data source.
	//
	// Most resources besides the 'testdata' data source
	// type require this field in order to retrieve data.
	Url *string `field:"optional" json:"url" yaml:"url"`
}

