// Grafana construct for cdk8s.
package cdk8sgrafana

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-grafana-go/cdk8sgrafana/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-grafana-go/cdk8sgrafana/internal"
)

// A Grafana instance.
type Grafana interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// Creates a dashboard associated with a particular data source.
	//
	// By default,
	// labels are automatically added so that the data source is detected by
	// Grafana.
	AddDashboard(id *string, props *DashboardProps) Dashboard
	// Adds a data source.
	//
	// By default, labels are automatically added so that
	// the data source is detected by Grafana.
	AddDataSource(id *string, props *DataSourceProps) DataSource
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Grafana
type jsiiProxy_Grafana struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_Grafana) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewGrafana(scope constructs.Construct, id *string, props *GrafanaProps) Grafana {
	_init_.Initialize()

	if err := validateNewGrafanaParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Grafana{}

	_jsii_.Create(
		"cdk8s-grafana.Grafana",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewGrafana_Override(g Grafana, scope constructs.Construct, id *string, props *GrafanaProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-grafana.Grafana",
		[]interface{}{scope, id, props},
		g,
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
func Grafana_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGrafana_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-grafana.Grafana",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Grafana) AddDashboard(id *string, props *DashboardProps) Dashboard {
	if err := g.validateAddDashboardParameters(id, props); err != nil {
		panic(err)
	}
	var returns Dashboard

	_jsii_.Invoke(
		g,
		"addDashboard",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Grafana) AddDataSource(id *string, props *DataSourceProps) DataSource {
	if err := g.validateAddDataSourceParameters(id, props); err != nil {
		panic(err)
	}
	var returns DataSource

	_jsii_.Invoke(
		g,
		"addDataSource",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Grafana) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

