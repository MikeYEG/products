package linkers

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/labels"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/tests"
)

type applicationBuilder struct {
	name    string
	version string
	ins     instructions.Instructions
	tests   tests.Tests
	lbls    labels.Labels
	imports []External
}

func createApplicationBuilder() ApplicationBuilder {
	out := applicationBuilder{
		name:    "",
		version: "",
		ins:     nil,
		tests:   nil,
		lbls:    nil,
		imports: nil,
	}

	return &out
}

// Create initializes the builder
func (app *applicationBuilder) Create() ApplicationBuilder {
	return createApplicationBuilder()
}

// WithName adds a name to the builder
func (app *applicationBuilder) WithName(name string) ApplicationBuilder {
	app.name = name
	return app
}

// WithInstructions add instructions to the builder
func (app *applicationBuilder) WithInstructions(ins instructions.Instructions) ApplicationBuilder {
	app.ins = ins
	return app
}

// WithTests add tests to the builder
func (app *applicationBuilder) WithTests(tests tests.Tests) ApplicationBuilder {
	app.tests = tests
	return app
}

// WithLabels add labels to the builder
func (app *applicationBuilder) WithLabels(lbls labels.Labels) ApplicationBuilder {
	app.lbls = lbls
	return app
}

// WithImports add external imports to the builder
func (app *applicationBuilder) WithImports(imports []External) ApplicationBuilder {
	app.imports = imports
	return app
}

// WithVersion adds a version to the builder
func (app *applicationBuilder) WithVersion(version string) ApplicationBuilder {
	app.version = version
	return app
}

// Now builds a new Application instance
func (app *applicationBuilder) Now() (Application, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Application instance")
	}

	if app.version == "" {
		return nil, errors.New("the version is mandatory in order to build an Application instance")
	}

	if app.ins == nil {
		return nil, errors.New("the instructions are mandatory in order to build an Application instance")
	}

	if app.tests == nil {
		return nil, errors.New("the tests are mandatory in order to build an Application instance")
	}

	if app.lbls == nil {
		return nil, errors.New("the labels are mandatory in order to build an Application instance")
	}

	if app.imports != nil {
		mp := map[string]Executable{}
		for _, oneImport := range app.imports {
			mp[oneImport.Name()] = oneImport.Executable()
		}

		return createApplicationWithImports(
			app.name,
			app.version,
			app.ins,
			app.tests,
			app.lbls,
			app.imports,
			mp,
		), nil
	}

	return createApplication(
		app.name,
		app.version,
		app.ins,
		app.tests,
		app.lbls,
	), nil
}
