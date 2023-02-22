//go:build no_runtime_type_checking

// Grafana construct for cdk8s.
package cdk8sgrafana

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_Dashboard) validateAddPluginsParameters(plugins *[]*GrafanaPlugin) error {
	return nil
}

func validateDashboard_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewDashboardParameters(scope constructs.Construct, id *string, props *DashboardProps) error {
	return nil
}

