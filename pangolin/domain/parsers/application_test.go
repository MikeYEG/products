package parsers

import (
	"testing"
)

func Test_application_simple_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("application", grammarFile)

	file := "./tests/codes/application/simple.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	app := ins.(Application)
	if app.HasLabel() {
		t.Errorf("the application was not expecting label")
		return
	}

	if app.HasTest() {
		t.Errorf("the application was not expecting test")
		return
	}
}

func Test_application_WithLabel_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("application", grammarFile)

	file := "./tests/codes/application/with_label.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	app := ins.(Application)
	if !app.HasLabel() {
		t.Errorf("the application was expecting label")
		return
	}

	if app.HasTest() {
		t.Errorf("the application was not expecting test")
		return
	}
}

func Test_application_WithLabel_withTest_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("application", grammarFile)

	file := "./tests/codes/application/with_label_with_test.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	app := ins.(Application)
	if !app.HasLabel() {
		t.Errorf("the application was expecting label")
		return
	}

	if !app.HasTest() {
		t.Errorf("the application was expecting test")
		return
	}
}

func Test_application_withTest_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("application", grammarFile)

	file := "./tests/codes/application/with_test.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	app := ins.(Application)
	if app.HasLabel() {
		t.Errorf("the application was not expecting label")
		return
	}

	if !app.HasTest() {
		t.Errorf("the application was expecting test")
		return
	}
}
