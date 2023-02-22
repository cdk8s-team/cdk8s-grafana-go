//go:build no_runtime_type_checking

// Grafana construct for cdk8s.
package cdk8sgrafana

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_Grafana) validateAddDashboardParameters(id *string, props *DashboardProps) error {
	return nil
}

func (g *jsiiProxy_Grafana) validateAddDataSourceParameters(id *string, props *DataSourceProps) error {
	return nil
}

func validateGrafana_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewGrafanaParameters(scope constructs.Construct, id *string, props *GrafanaProps) error {
	return nil
}

