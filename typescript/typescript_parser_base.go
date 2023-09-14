package typescript

import (
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type TypeScriptParserBase struct {
	*antlr.BaseParser
}

func (p *TypeScriptParserBase) P(str string) bool {
	return p.Prev(str)
}

func (p *TypeScriptParserBase) Prev(str string) bool {
	return p.GetTokenStream().LT(-1).GetText() == str
}

func (p *TypeScriptParserBase) N(str string) bool {
	return p.Next(str)
}

func (p *TypeScriptParserBase) Next(str string) bool {
	return p.GetTokenStream().LT(1).GetText() == str
}

func (p *TypeScriptParserBase) NotLineTerminator() bool {
	return !p.Here(TypeScriptParserLineTerminator)
}

func (p *TypeScriptParserBase) NotOpenBraceAndNotFunction() bool {
	nextTokenType := p.GetTokenStream().LT(1).GetTokenType()
	return nextTokenType != TypeScriptParserOpenBrace && nextTokenType != TypeScriptParserFunction_
}

func (p *TypeScriptParserBase) CloseBrace() bool {
	return p.GetTokenStream().LT(1).GetTokenType() == TypeScriptParserCloseBrace
}

func (p *TypeScriptParserBase) Here(tokenType int) bool {
	possibleIndexEosToken := p.GetCurrentToken().GetTokenIndex() - 1
	ahead := p.GetTokenStream().Get(possibleIndexEosToken)
	return ahead.GetChannel() == antlr.LexerHidden && ahead.GetTokenType() == tokenType
}

func (p *TypeScriptParserBase) LineTerminatorAhead() bool {
	possibleIndexEosToken := p.GetCurrentToken().GetTokenIndex() - 1
	ahead := p.GetTokenStream().Get(possibleIndexEosToken)
	if ahead.GetChannel() != antlr.LexerHidden {
		return false
	}

	if ahead.GetTokenType() == TypeScriptParserLineTerminator {
		return true
	}

	if ahead.GetTokenType() == TypeScriptParserWhiteSpaces {
		possibleIndexEosToken = p.GetCurrentToken().GetTokenIndex() - 2
		ahead = p.GetTokenStream().Get(possibleIndexEosToken)
	}

	text := ahead.GetText()
	tokenType := ahead.GetTokenType()

	return (tokenType == TypeScriptParserMultiLineComment && (strings.Contains(text, "\r") || strings.Contains(text, "\n"))) ||
		(tokenType == TypeScriptParserLineTerminator)
}
