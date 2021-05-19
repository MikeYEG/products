package parsers

type languageTestInstruction struct {
	lang LanguageInstructionCommon
	test TestInstruction
}

func createLanguageTestInstructionWithLanguage(
	lang LanguageInstructionCommon,
) LanguageTestInstruction {
	return createLanguageTestInstructionInternally(lang, nil)
}

func createLanguageTestInstructionWithTest(
	test TestInstruction,
) LanguageTestInstruction {
	return createLanguageTestInstructionInternally(nil, test)
}

func createLanguageTestInstructionInternally(
	lang LanguageInstructionCommon,
	test TestInstruction,
) LanguageTestInstruction {
	out := languageTestInstruction{
		lang: lang,
		test: test,
	}

	return &out
}

// IsLanguageInstruction returns true if there is a language instruction, false otherwise
func (obj *languageTestInstruction) IsLanguageInstruction() bool {
	return obj.lang != nil
}

// LanguageInstruction returns the language instruction, if any
func (obj *languageTestInstruction) LanguageInstruction() LanguageInstructionCommon {
	return obj.lang
}

// IsTestInstruction returns true if there is a test instruction, false otherwise
func (obj *languageTestInstruction) IsTestInstruction() bool {
	return obj.test != nil
}

// TestInstruction returns the test instruction, if any
func (obj *languageTestInstruction) TestInstruction() TestInstruction {
	return obj.test
}
