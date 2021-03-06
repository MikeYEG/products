package languages

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/languages/definitions"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type adapter struct {
	patternMatchadapter definitions.PatternMatchAdapter
	builder             Builder
	valueBuilder        ValueBuilder
}

func createAdapter(
	patternMatchadapter definitions.PatternMatchAdapter,
	builder Builder,
	valueBuilder ValueBuilder,
) Adapter {
	out := adapter{
		patternMatchadapter: patternMatchadapter,
		builder:             builder,
		valueBuilder:        valueBuilder,
	}

	return &out
}

// ToLanguage converts a parsed language command to a language instance
func (app *adapter) ToLanguage(parsed parsers.LanguageCommand) (Language, error) {
	values := []Value{}
	parsedValues := parsed.Values()
	for _, oneParsedValue := range parsedValues {
		valueBuilder := app.valueBuilder.Create()
		if oneParsedValue.IsRoot() {
			root := oneParsedValue.Root()
			valueBuilder.WithRoot(root)
		}

		if oneParsedValue.IsTokens() {
			tokensPath := oneParsedValue.Tokens()
			valueBuilder.WithTokensPath(tokensPath)
		}

		if oneParsedValue.IsChannels() {
			channelsPath := oneParsedValue.Channels()
			valueBuilder.WithChannelsPath(channelsPath)
		}

		if oneParsedValue.IsRules() {
			rulesPath := oneParsedValue.Rules()
			valueBuilder.WithRulesPath(rulesPath)
		}

		if oneParsedValue.IsLogic() {
			logicsPath := oneParsedValue.Logic()
			valueBuilder.WithLogicsPath(logicsPath)
		}

		if oneParsedValue.IsInputVariable() {
			inputVariable := oneParsedValue.InputVariable()
			valueBuilder.WithInputVariable(inputVariable)
		}

		if oneParsedValue.IsExtends() {
			parsedExtendPaths := oneParsedValue.Extends()
			valueBuilder.WithExtends(parsedExtendPaths)
		}

		if oneParsedValue.IsPatternMatches() {
			parsedPatternMatches := oneParsedValue.PatternMatches()
			patternMatches, err := app.patternMatchadapter.ToPatternMatches(parsedPatternMatches)
			if err != nil {
				return nil, err
			}

			valueBuilder.WithPatternMatches(patternMatches)
		}

		value, err := valueBuilder.Now()
		if err != nil {
			return nil, err
		}

		values = append(values, value)
	}

	variable := parsed.Variable()
	return app.builder.Create().
		WithValues(values).
		WithVariable(variable).
		Now()
}
