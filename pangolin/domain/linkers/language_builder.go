package linkers

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/definitions"
)

type languageBuilder struct {
	app     Application
	matches []definitions.PatternMatch
	paths   Paths
	root    string
}

func createLanguageBuilder() LanguageBuilder {
	out := languageBuilder{
		app:     nil,
		matches: nil,
		paths:   nil,
		root:    "",
	}

	return &out
}

// Create initializes the builder
func (app *languageBuilder) Create() LanguageBuilder {
	return createLanguageBuilder()
}

// WithApplication adds an application to the builder
func (app *languageBuilder) WithApplication(appli Application) LanguageBuilder {
	app.app = appli
	return app
}

// WithPatternMatches add pattern matches to the builder
func (app *languageBuilder) WithPatternMatches(matches []definitions.PatternMatch) LanguageBuilder {
	app.matches = matches
	return app
}

// WithPaths adds a paths to the builder
func (app *languageBuilder) WithPaths(paths Paths) LanguageBuilder {
	app.paths = paths
	return app
}

// WithRoot adds a root to the builder
func (app *languageBuilder) WithRoot(root string) LanguageBuilder {
	app.root = root
	return app
}

// Now builds a new Language instance
func (app *languageBuilder) Now() (Language, error) {
	if app.app == nil {
		return nil, errors.New("the application is mandatory in order to build a Language instance")
	}

	if app.paths == nil {
		return nil, errors.New("the Paths instance is mandatory in order to build a Language instance")
	}

	if app.root == "" {
		return nil, errors.New("the root pattern is mandatory in order to build a Language instance")
	}

	if app.matches != nil && len(app.matches) <= 0 {
		app.matches = nil
	}

	if app.matches == nil {
		return nil, errors.New("the []PatternMatch are mandatory in order to build a Language instance")
	}

	return createLanguage(app.app, app.matches, app.paths, app.root), nil
}
