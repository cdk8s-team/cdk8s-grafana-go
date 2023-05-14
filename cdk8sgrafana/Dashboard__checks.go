//go:build !no_runtime_type_checking

package cdk8sgrafana

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (d *jsiiProxy_Dashboard) validateAddPluginsParameters(plugins *[]*GrafanaPlugin) error {
	for idx_92cd10, v := range *plugins {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter plugins[%#v]", idx_92cd10) }); err != nil {
			return err
		}
	}

	return nil
}

func validateDashboard_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewDashboardParameters(scope constructs.Construct, id *string, props *DashboardProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

