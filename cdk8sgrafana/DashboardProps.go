// Grafana construct for cdk8s.
package cdk8sgrafana

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type DashboardProps struct {
	// Title of the dashboard.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Specify a mapping from data source variables to data source names.
	//
	// This is only needed if you are importing an existing dashboard's JSON
	// and it specifies variables within an "__inputs" field.
	//
	// Example:
	//   { DS_PROMETHEUS: "my-prometheus-ds" }
	//
	DataSourceVariables *map[string]*string `field:"optional" json:"dataSourceVariables" yaml:"dataSourceVariables"`
	// Group dashboards into folders.
	Folder *string `field:"optional" json:"folder" yaml:"folder"`
	// All other dashboard customizations.
	// See: https://grafana.com/docs/grafana/latest/dashboards/json-model/
	//
	JsonModel *map[string]interface{} `field:"optional" json:"jsonModel" yaml:"jsonModel"`
	// Labels to apply to the kubernetes resource.
	//
	// When adding a dashboard to a Grafana instance using `grafana.addDashboard`,
	// labels provided to Grafana will be automatically applied. Otherwise,
	// labels must be added manually.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Namespace to apply to the kubernetes resource.
	//
	// When adding a dashboard to a Grafana instance using `grafana.addDashboard`,
	// the namespace will be automatically inherited.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specify plugins required by the dashboard.
	Plugins *[]*GrafanaPlugin `field:"optional" json:"plugins" yaml:"plugins"`
	// Auto-refresh interval.
	RefreshRate cdk8s.Duration `field:"optional" json:"refreshRate" yaml:"refreshRate"`
	// Time range for the dashboard, e.g. last 6 hours, last 7 days, etc.
	TimeRange cdk8s.Duration `field:"optional" json:"timeRange" yaml:"timeRange"`
}

