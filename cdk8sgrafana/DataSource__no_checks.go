//go:build no_runtime_type_checking

package cdk8sgrafana

// Building without runtime type checking enabled, so all the below just return nil

func validateDataSource_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewDataSourceParameters(scope constructs.Construct, id *string, props *DataSourceProps) error {
	return nil
}

