// Grafana construct for cdk8s.
package cdk8sgrafana


type GrafanaPlugin struct {
	// Name of the plugin, e.g. "grafana-piechart-panel".
	Name *string `field:"required" json:"name" yaml:"name"`
	// Version of the plugin, e.g. "1.3.6".
	Version *string `field:"required" json:"version" yaml:"version"`
}

