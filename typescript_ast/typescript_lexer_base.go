package typescript_ast

import "github.com/antlr4-go/antlr/v4"

type TypeScriptLexerBase struct {
	*antlr.BaseLexer

	scopeStrictModes []bool
	lastToken        antlr.Token

	useStrictDefault bool
	useStrictCurrent bool

	templateDepth int
	bracesDepth   int
}

func (t *TypeScriptLexerBase) GetUseStrictDefault() bool {
	return t.useStrictDefault
}

func (t *TypeScriptLexerBase) SetUseStrictDefault(value bool) {
	t.useStrictDefault = value
	t.useStrictCurrent = value
}

func (t *TypeScriptLexerBase) IsStartOfFile() bool {
	return t.lastToken == nil
}

func (t *TypeScriptLexerBase) IsStrictMode() bool {
	return t.useStrictCurrent
}

func (t *TypeScriptLexerBase) StartTemplateString() {
	t.bracesDepth = 0
}

func (t *TypeScriptLexerBase) IsInTemplateString() bool {
	return t.templateDepth > 0 && t.bracesDepth == 0
}

func (t *TypeScriptLexerBase) NextToken() antlr.Token {
	next := t.BaseLexer.NextToken().(antlr.Token)

	if next.GetChannel() == antlr.TokenDefaultChannel {
		t.lastToken = next
	}

	return next
}

func (t *TypeScriptLexerBase) ProcessOpenBrace() {
	t.bracesDepth++
	t.useStrictCurrent = (len(t.scopeStrictModes) > 0 && t.scopeStrictModes[len(t.scopeStrictModes)-1]) || t.useStrictDefault
	t.scopeStrictModes = append(t.scopeStrictModes, t.useStrictCurrent)
}

func (t *TypeScriptLexerBase) ProcessCloseBrace() {
	t.bracesDepth--
	if len(t.scopeStrictModes) > 0 {
		t.scopeStrictModes = t.scopeStrictModes[:len(t.scopeStrictModes)-1]
		t.useStrictCurrent = t.scopeStrictModes[len(t.scopeStrictModes)-1]
	} else {
		t.useStrictCurrent = t.useStrictDefault
	}
}

func (t *TypeScriptLexerBase) ProcessStringLiteral() {
	if t.lastToken == nil || t.lastToken.GetTokenType() == TypeScriptLexerOpenBrace {
		text := t.GetText()
		if text == "\"use strict\"" || text == "'use strict'" {
			if len(t.scopeStrictModes) > 0 {
				t.scopeStrictModes = t.scopeStrictModes[:len(t.scopeStrictModes)-1]
			}
			t.useStrictCurrent = true
			t.scopeStrictModes = append(t.scopeStrictModes, t.useStrictCurrent)
		}
	}
}

func (t *TypeScriptLexerBase) IncreaseTemplateDepth() {
	t.templateDepth++
}

func (t *TypeScriptLexerBase) DecreaseTemplateDepth() {
	t.templateDepth--
}

func (t *TypeScriptLexerBase) IsRegexPossible() bool {
	if t.lastToken == nil {
		return true
	}

	switch t.lastToken.GetTokenType() {
	case TypeScriptLexerIdentifier, TypeScriptLexerNullLiteral, TypeScriptLexerBooleanLiteral,
		TypeScriptLexerThis, TypeScriptLexerCloseBracket, TypeScriptLexerCloseParen,
		TypeScriptLexerOctalIntegerLiteral, TypeScriptLexerDecimalLiteral,
		TypeScriptLexerHexIntegerLiteral, TypeScriptLexerStringLiteral,
		TypeScriptLexerPlusPlus, TypeScriptLexerMinusMinus:
		return false
	default:
		return true
	}
}
