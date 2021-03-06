package parsers

import (
	"github.com/deepvalue-network/software/pangolin/domain/lexers"
	lexer_parser "github.com/deepvalue-network/software/pangolin/domain/lexers/parser"
)

type parserBuilder struct {
	lexerBuilder                     lexers.Builder
	lexerParserApplication           lexer_parser.Application
	lexerParserBuilder               lexer_parser.Builder
	programBuilder                   ProgramBuilder
	testableBuilder                  TestableBuilder
	executableBuilder                ExecutableBuilder
	scopesBuilder                    ScopesBuilder
	scopeBuilder                     ScopeBuilder
	commandBuilder                   CommandBuilder
	languageCommandBuilder           LanguageCommandBuilder
	scriptCommandBuilder             ScriptCommandBuilder
	headCommandBuilder               HeadCommandBuilder
	mainCommandBuilder               MainCommandBuilder
	mainCommandInstructionBuilder    MainCommandInstructionBuilder
	testCommandBuilder               TestCommandBuilder
	testCommandInstructionBuilder    TestCommandInstructionBuilder
	labelCommandBuilder              LabelCommandBuilder
	labelCommandInstructionBuilder   LabelCommandInstructionBuilder
	languageApplicationBuilder       LanguageApplicationBuilder
	languageMainSectionBuilder       LanguageMainSectionBuilder
	languageTestSectionBuilder       LanguageTestSectionBuilder
	languageTestDeclarationBuilder   LanguageTestDeclarationBuilder
	languageTestInstructionBuilder   LanguageTestInstructionBuilder
	languageLabelSectionBuilder      LanguageLabelSectionBuilder
	languageLabelDeclarationBuilder  LanguageLabelDeclarationBuilder
	languageLabelInstructionBuilder  LanguageLabelInstructionBuilder
	languageInstructionBuilder       LanguageInstructionBuilder
	languageInstructionCommonBuilder LanguageInstructionCommonBuilder
	languageDefinitionBuilder        LanguageDefinitionBuilder
	languageValueBuilder             LanguageValueBuilder
	scriptBuilder                    ScriptBuilder
	scriptValueBuilder               ScriptValueBuilder
	scriptTestsBuilder               ScriptTestsBuilder
	scriptTestBuilder                ScriptTestBuilder
	patternMatchBuilder              PatternMatchBuilder
	patternLabelsBuilder             PatternLabelsBuilder
	relativePathsBuilder             RelativePathsBuilder
	relativePathBuilder              RelativePathBuilder
	folderSectionBuilder             FolderSectionBuilder
	folderNameBuilder                FolderNameBuilder
	applicationBuilder               ApplicationBuilder
	testSectionBuilder               TestSectionBuilder
	testDeclarationBuilder           TestDeclarationBuilder
	testInstructionBuilder           TestInstructionBuilder
	assertBuilder                    AssertBuilder
	readFileBuilder                  ReadFileBuilder
	headSectionBuilder               HeadSectionBuilder
	headValueBuilder                 HeadValueBuilder
	loadSingleBuilder                LoadSingleBuilder
	importSingleBuilder              ImportSingleBuilder
	labelSectionBuilder              LabelSectionBuilder
	labelDeclarationBuilder          LabelDeclarationBuilder
	labelInstructionBuilder          LabelInstructionBuilder
	mainSectionBuilder               MainSectionBuilder
	instructionBuilder               InstructionBuilder
	registryBuilder                  RegistryBuilder
	fetchRegistryBuilder             FetchRegistryBuilder
	unregisterBuilder                UnregisterBuilder
	registerBuilder                  RegisterBuilder
	specificTokenCodeBuilder         SpecificTokenCodeBuilder
	tokenSectionBuilder              TokenSectionBuilder
	codeMatchBuilder                 CodeMatchBuilder
	tokenBuilder                     TokenBuilder
	variableBuilder                  VariableBuilder
	concatenationBuilder             ConcatenationBuilder
	declarationBuilder               DeclarationBuilder
	assignmentBuilder                AssignmentBuilder
	valueRepresentationBuilder       ValueRepresentationBuilder
	valueBuilder                     ValueBuilder
	numericValueBuilder              NumericValueBuilder
	typeBuilder                      TypeBuilder
	operationBuilder                 OperationBuilder
	arythmeticBuilder                ArythmeticBuilder
	relationalBuilder                RelationalBuilder
	logicalBuilder                   LogicalBuilder
	standardOperationBuilder         StandardOperationBuilder
	remainingOperationBuilder        RemainingOperationBuilder
	printBuilder                     PrintBuilder
	jumpBuilder                      JumpBuilder
	matchBuilder                     MatchBuilder
	exitBuilder                      ExitBuilder
	callBuilder                      CallBuilder
	moduleBuilder                    ModuleBuilder
	switchBuilder                    SwitchBuilder
	saveBuilder                      SaveBuilder
	stackFrameBuilder                StackFrameBuilder
	indexBuilder                     IndexBuilder
	skipBuilder                      SkipBuilder
	intPointerBuilder                IntPointerBuilder
	lexerAdapter                     lexers.Adapter
}

func createParserBuilder(
	lexerParserApplication lexer_parser.Application,
	lexerParserBuilder lexer_parser.Builder,
	lexerBuilder lexers.Builder,
	programBuilder ProgramBuilder,
	testableBuilder TestableBuilder,
	executableBuilder ExecutableBuilder,
	scopesBuilder ScopesBuilder,
	scopeBuilder ScopeBuilder,
	commandBuilder CommandBuilder,
	languageCommandBuilder LanguageCommandBuilder,
	scriptCommandBuilder ScriptCommandBuilder,
	headCommandBuilder HeadCommandBuilder,
	mainCommandBuilder MainCommandBuilder,
	mainCommandInstructionBuilder MainCommandInstructionBuilder,
	testCommandBuilder TestCommandBuilder,
	testCommandInstructionBuilder TestCommandInstructionBuilder,
	labelCommandBuilder LabelCommandBuilder,
	labelCommandInstructionBuilder LabelCommandInstructionBuilder,
	languageApplicationBuilder LanguageApplicationBuilder,
	languageMainSectionBuilder LanguageMainSectionBuilder,
	languageTestSectionBuilder LanguageTestSectionBuilder,
	languageTestDeclarationBuilder LanguageTestDeclarationBuilder,
	languageTestInstructionBuilder LanguageTestInstructionBuilder,
	languageLabelSectionBuilder LanguageLabelSectionBuilder,
	languageLabelDeclarationBuilder LanguageLabelDeclarationBuilder,
	languageLabelInstructionBuilder LanguageLabelInstructionBuilder,
	languageInstructionBuilder LanguageInstructionBuilder,
	languageInstructionCommonBuilder LanguageInstructionCommonBuilder,
	languageDefinitionBuilder LanguageDefinitionBuilder,
	languageValueBuilder LanguageValueBuilder,
	scriptBuilder ScriptBuilder,
	scriptValueBuilder ScriptValueBuilder,
	scriptTestsBuilder ScriptTestsBuilder,
	scriptTestBuilder ScriptTestBuilder,
	patternMatchBuilder PatternMatchBuilder,
	patternLabelsBuilder PatternLabelsBuilder,
	relativePathsBuilder RelativePathsBuilder,
	relativePathBuilder RelativePathBuilder,
	folderSectionBuilder FolderSectionBuilder,
	folderNameBuilder FolderNameBuilder,
	applicationBuilder ApplicationBuilder,
	testSectionBuilder TestSectionBuilder,
	testDeclarationBuilder TestDeclarationBuilder,
	testInstructionBuilder TestInstructionBuilder,
	assertBuilder AssertBuilder,
	readFileBuilder ReadFileBuilder,
	headSectionBuilder HeadSectionBuilder,
	headValueBuilder HeadValueBuilder,
	loadSingleBuilder LoadSingleBuilder,
	importSingleBuilder ImportSingleBuilder,
	labelSectionBuilder LabelSectionBuilder,
	labelDeclarationBuilder LabelDeclarationBuilder,
	labelInstructionBuilder LabelInstructionBuilder,
	mainSectionBuilder MainSectionBuilder,
	instructionBuilder InstructionBuilder,
	registryBuilder RegistryBuilder,
	fetchRegistryBuilder FetchRegistryBuilder,
	unregisterBuilder UnregisterBuilder,
	registerBuilder RegisterBuilder,
	specificTokenCodeBuilder SpecificTokenCodeBuilder,
	tokenSectionBuilder TokenSectionBuilder,
	codeMatchBuilder CodeMatchBuilder,
	tokenBuilder TokenBuilder,
	variableBuilder VariableBuilder,
	concatenationBuilder ConcatenationBuilder,
	declarationBuilder DeclarationBuilder,
	assignmentBuilder AssignmentBuilder,
	valueRepresentationBuilder ValueRepresentationBuilder,
	valueBuilder ValueBuilder,
	numericValueBuilder NumericValueBuilder,
	typeBuilder TypeBuilder,
	operationBuilder OperationBuilder,
	arythmeticBuilder ArythmeticBuilder,
	relationalBuilder RelationalBuilder,
	logicalBuilder LogicalBuilder,
	standardOperationBuilder StandardOperationBuilder,
	remainingOperationBuilder RemainingOperationBuilder,
	printBuilder PrintBuilder,
	jumpBuilder JumpBuilder,
	matchBuilder MatchBuilder,
	exitBuilder ExitBuilder,
	callBuilder CallBuilder,
	moduleBuilder ModuleBuilder,
	switchBuilder SwitchBuilder,
	saveBuilder SaveBuilder,
	stackFrameBuilder StackFrameBuilder,
	indexBuilder IndexBuilder,
	skipBuilder SkipBuilder,
	intPointerBuilder IntPointerBuilder,
) ParserBuilder {
	out := parserBuilder{
		lexerParserApplication:           lexerParserApplication,
		lexerParserBuilder:               lexerParserBuilder,
		lexerBuilder:                     lexerBuilder,
		programBuilder:                   programBuilder,
		testableBuilder:                  testableBuilder,
		executableBuilder:                executableBuilder,
		scopesBuilder:                    scopesBuilder,
		scopeBuilder:                     scopeBuilder,
		commandBuilder:                   commandBuilder,
		languageCommandBuilder:           languageCommandBuilder,
		scriptCommandBuilder:             scriptCommandBuilder,
		headCommandBuilder:               headCommandBuilder,
		mainCommandBuilder:               mainCommandBuilder,
		mainCommandInstructionBuilder:    mainCommandInstructionBuilder,
		testCommandBuilder:               testCommandBuilder,
		testCommandInstructionBuilder:    testCommandInstructionBuilder,
		labelCommandBuilder:              labelCommandBuilder,
		labelCommandInstructionBuilder:   labelCommandInstructionBuilder,
		languageApplicationBuilder:       languageApplicationBuilder,
		languageMainSectionBuilder:       languageMainSectionBuilder,
		languageTestSectionBuilder:       languageTestSectionBuilder,
		languageTestDeclarationBuilder:   languageTestDeclarationBuilder,
		languageTestInstructionBuilder:   languageTestInstructionBuilder,
		languageLabelSectionBuilder:      languageLabelSectionBuilder,
		languageLabelDeclarationBuilder:  languageLabelDeclarationBuilder,
		languageLabelInstructionBuilder:  languageLabelInstructionBuilder,
		languageInstructionBuilder:       languageInstructionBuilder,
		languageInstructionCommonBuilder: languageInstructionCommonBuilder,
		languageDefinitionBuilder:        languageDefinitionBuilder,
		languageValueBuilder:             languageValueBuilder,
		scriptBuilder:                    scriptBuilder,
		scriptValueBuilder:               scriptValueBuilder,
		scriptTestsBuilder:               scriptTestsBuilder,
		scriptTestBuilder:                scriptTestBuilder,
		patternMatchBuilder:              patternMatchBuilder,
		patternLabelsBuilder:             patternLabelsBuilder,
		relativePathsBuilder:             relativePathsBuilder,
		relativePathBuilder:              relativePathBuilder,
		folderSectionBuilder:             folderSectionBuilder,
		folderNameBuilder:                folderNameBuilder,
		applicationBuilder:               applicationBuilder,
		testSectionBuilder:               testSectionBuilder,
		testDeclarationBuilder:           testDeclarationBuilder,
		testInstructionBuilder:           testInstructionBuilder,
		assertBuilder:                    assertBuilder,
		readFileBuilder:                  readFileBuilder,
		headSectionBuilder:               headSectionBuilder,
		headValueBuilder:                 headValueBuilder,
		loadSingleBuilder:                loadSingleBuilder,
		importSingleBuilder:              importSingleBuilder,
		labelSectionBuilder:              labelSectionBuilder,
		labelDeclarationBuilder:          labelDeclarationBuilder,
		labelInstructionBuilder:          labelInstructionBuilder,
		mainSectionBuilder:               mainSectionBuilder,
		instructionBuilder:               instructionBuilder,
		registryBuilder:                  registryBuilder,
		fetchRegistryBuilder:             fetchRegistryBuilder,
		unregisterBuilder:                unregisterBuilder,
		registerBuilder:                  registerBuilder,
		specificTokenCodeBuilder:         specificTokenCodeBuilder,
		tokenSectionBuilder:              tokenSectionBuilder,
		codeMatchBuilder:                 codeMatchBuilder,
		tokenBuilder:                     tokenBuilder,
		variableBuilder:                  variableBuilder,
		concatenationBuilder:             concatenationBuilder,
		declarationBuilder:               declarationBuilder,
		assignmentBuilder:                assignmentBuilder,
		valueRepresentationBuilder:       valueRepresentationBuilder,
		valueBuilder:                     valueBuilder,
		numericValueBuilder:              numericValueBuilder,
		typeBuilder:                      typeBuilder,
		operationBuilder:                 operationBuilder,
		arythmeticBuilder:                arythmeticBuilder,
		relationalBuilder:                relationalBuilder,
		logicalBuilder:                   logicalBuilder,
		standardOperationBuilder:         standardOperationBuilder,
		remainingOperationBuilder:        remainingOperationBuilder,
		printBuilder:                     printBuilder,
		jumpBuilder:                      jumpBuilder,
		matchBuilder:                     matchBuilder,
		exitBuilder:                      exitBuilder,
		callBuilder:                      callBuilder,
		moduleBuilder:                    moduleBuilder,
		switchBuilder:                    switchBuilder,
		saveBuilder:                      saveBuilder,
		stackFrameBuilder:                stackFrameBuilder,
		indexBuilder:                     indexBuilder,
		skipBuilder:                      skipBuilder,
		intPointerBuilder:                intPointerBuilder,
		lexerAdapter:                     nil,
	}

	return &out
}

// Create initializes the builder
func (app *parserBuilder) Create() ParserBuilder {
	return createParserBuilder(
		app.lexerParserApplication,
		app.lexerParserBuilder,
		app.lexerBuilder,
		app.programBuilder,
		app.testableBuilder,
		app.executableBuilder,
		app.scopesBuilder,
		app.scopeBuilder,
		app.commandBuilder,
		app.languageCommandBuilder,
		app.scriptCommandBuilder,
		app.headCommandBuilder,
		app.mainCommandBuilder,
		app.mainCommandInstructionBuilder,
		app.testCommandBuilder,
		app.testCommandInstructionBuilder,
		app.labelCommandBuilder,
		app.labelCommandInstructionBuilder,
		app.languageApplicationBuilder,
		app.languageMainSectionBuilder,
		app.languageTestSectionBuilder,
		app.languageTestDeclarationBuilder,
		app.languageTestInstructionBuilder,
		app.languageLabelSectionBuilder,
		app.languageLabelDeclarationBuilder,
		app.languageLabelInstructionBuilder,
		app.languageInstructionBuilder,
		app.languageInstructionCommonBuilder,
		app.languageDefinitionBuilder,
		app.languageValueBuilder,
		app.scriptBuilder,
		app.scriptValueBuilder,
		app.scriptTestsBuilder,
		app.scriptTestBuilder,
		app.patternMatchBuilder,
		app.patternLabelsBuilder,
		app.relativePathsBuilder,
		app.relativePathBuilder,
		app.folderSectionBuilder,
		app.folderNameBuilder,
		app.applicationBuilder,
		app.testSectionBuilder,
		app.testDeclarationBuilder,
		app.testInstructionBuilder,
		app.assertBuilder,
		app.readFileBuilder,
		app.headSectionBuilder,
		app.headValueBuilder,
		app.loadSingleBuilder,
		app.importSingleBuilder,
		app.labelSectionBuilder,
		app.labelDeclarationBuilder,
		app.labelInstructionBuilder,
		app.mainSectionBuilder,
		app.instructionBuilder,
		app.registryBuilder,
		app.fetchRegistryBuilder,
		app.unregisterBuilder,
		app.registerBuilder,
		app.specificTokenCodeBuilder,
		app.tokenSectionBuilder,
		app.codeMatchBuilder,
		app.tokenBuilder,
		app.variableBuilder,
		app.concatenationBuilder,
		app.declarationBuilder,
		app.assignmentBuilder,
		app.valueRepresentationBuilder,
		app.valueBuilder,
		app.numericValueBuilder,
		app.typeBuilder,
		app.operationBuilder,
		app.arythmeticBuilder,
		app.relationalBuilder,
		app.logicalBuilder,
		app.standardOperationBuilder,
		app.remainingOperationBuilder,
		app.printBuilder,
		app.jumpBuilder,
		app.matchBuilder,
		app.exitBuilder,
		app.callBuilder,
		app.moduleBuilder,
		app.switchBuilder,
		app.saveBuilder,
		app.stackFrameBuilder,
		app.indexBuilder,
		app.skipBuilder,
		app.intPointerBuilder,
	)
}

// WithLexerAdapter adds a lexerAdapter to the builder
func (app *parserBuilder) WithLexerAdapter(lexerAdapter lexers.Adapter) ParserBuilder {
	app.lexerAdapter = lexerAdapter
	return app
}

// Now builds a new Parser instance
func (app *parserBuilder) Now() (Parser, error) {
	return createParser(
		app.lexerAdapter,
		app.lexerParserApplication,
		app.lexerParserBuilder,
		app.lexerBuilder,
		app.programBuilder,
		app.testableBuilder,
		app.executableBuilder,
		app.scopesBuilder,
		app.scopeBuilder,
		app.commandBuilder,
		app.languageCommandBuilder,
		app.scriptCommandBuilder,
		app.headCommandBuilder,
		app.mainCommandBuilder,
		app.mainCommandInstructionBuilder,
		app.testCommandBuilder,
		app.testCommandInstructionBuilder,
		app.labelCommandBuilder,
		app.labelCommandInstructionBuilder,
		app.languageApplicationBuilder,
		app.languageMainSectionBuilder,
		app.languageTestSectionBuilder,
		app.languageTestDeclarationBuilder,
		app.languageTestInstructionBuilder,
		app.languageLabelSectionBuilder,
		app.languageLabelDeclarationBuilder,
		app.languageLabelInstructionBuilder,
		app.languageInstructionBuilder,
		app.languageInstructionCommonBuilder,
		app.languageDefinitionBuilder,
		app.languageValueBuilder,
		app.scriptBuilder,
		app.scriptValueBuilder,
		app.scriptTestsBuilder,
		app.scriptTestBuilder,
		app.patternMatchBuilder,
		app.patternLabelsBuilder,
		app.relativePathsBuilder,
		app.relativePathBuilder,
		app.folderSectionBuilder,
		app.folderNameBuilder,
		app.applicationBuilder,
		app.testSectionBuilder,
		app.testDeclarationBuilder,
		app.testInstructionBuilder,
		app.assertBuilder,
		app.readFileBuilder,
		app.headSectionBuilder,
		app.headValueBuilder,
		app.loadSingleBuilder,
		app.importSingleBuilder,
		app.labelSectionBuilder,
		app.labelDeclarationBuilder,
		app.labelInstructionBuilder,
		app.mainSectionBuilder,
		app.instructionBuilder,
		app.registryBuilder,
		app.fetchRegistryBuilder,
		app.unregisterBuilder,
		app.registerBuilder,
		app.specificTokenCodeBuilder,
		app.tokenSectionBuilder,
		app.codeMatchBuilder,
		app.tokenBuilder,
		app.variableBuilder,
		app.concatenationBuilder,
		app.declarationBuilder,
		app.assignmentBuilder,
		app.valueRepresentationBuilder,
		app.valueBuilder,
		app.numericValueBuilder,
		app.typeBuilder,
		app.operationBuilder,
		app.arythmeticBuilder,
		app.relationalBuilder,
		app.logicalBuilder,
		app.standardOperationBuilder,
		app.remainingOperationBuilder,
		app.printBuilder,
		app.jumpBuilder,
		app.matchBuilder,
		app.exitBuilder,
		app.callBuilder,
		app.moduleBuilder,
		app.switchBuilder,
		app.saveBuilder,
		app.stackFrameBuilder,
		app.indexBuilder,
		app.skipBuilder,
		app.intPointerBuilder,
	)
}
