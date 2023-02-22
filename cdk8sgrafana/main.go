package cdk8sgrafana

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterEnum(
		"cdk8s-grafana.AccessType",
		reflect.TypeOf((*AccessType)(nil)).Elem(),
		map[string]interface{}{
			"PROXY": AccessType_PROXY,
			"DIRECT": AccessType_DIRECT,
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-grafana.Dashboard",
		reflect.TypeOf((*Dashboard)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addPlugins", GoMethod: "AddPlugins"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Dashboard{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-grafana.DashboardProps",
		reflect.TypeOf((*DashboardProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-grafana.DataSource",
		reflect.TypeOf((*DataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataSource{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-grafana.DataSourceProps",
		reflect.TypeOf((*DataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-grafana.Grafana",
		reflect.TypeOf((*Grafana)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDashboard", GoMethod: "AddDashboard"},
			_jsii_.MemberMethod{JsiiMethod: "addDataSource", GoMethod: "AddDataSource"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Grafana{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-grafana.GrafanaPlugin",
		reflect.TypeOf((*GrafanaPlugin)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-grafana.GrafanaProps",
		reflect.TypeOf((*GrafanaProps)(nil)).Elem(),
	)
}
