package cdk8sgrafana

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-grafana-go/cdk8sgrafana/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-grafana-go/cdk8sgrafana/internal"
)

// A Grafana dashboard.
// See: https://grafana.com/docs/grafana/latest/http_api/dashboard/
//
type Dashboard interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// Adds one or more plugins.
	AddPlugins(plugins ...*GrafanaPlugin)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Dashboard
type jsiiProxy_Dashboard struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_Dashboard) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewDashboard(scope constructs.Construct, id *string, props *DashboardProps) Dashboard {
	_init_.Initialize()

	if err := validateNewDashboardParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Dashboard{}

	_jsii_.Create(
		"cdk8s-grafana.Dashboard",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDashboard_Override(d Dashboard, scope constructs.Construct, id *string, props *DashboardProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-grafana.Dashboard",
		[]interface{}{scope, id, props},
		d,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Dashboard_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDashboard_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-grafana.Dashboard",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Dashboard) AddPlugins(plugins ...*GrafanaPlugin) {
	if err := d.validateAddPluginsParameters(&plugins); err != nil {
		panic(err)
	}
	args := []interface{}{}
	for _, a := range plugins {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		d,
		"addPlugins",
		args,
	)
}

func (d *jsiiProxy_Dashboard) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

