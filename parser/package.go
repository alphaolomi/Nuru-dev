package parser

import (
	"github.com/AvicennaJr/Nuru/ast"
	"github.com/AvicennaJr/Nuru/token"
)

func (p *Parser) parsePackage() ast.Expression {
	expression := &ast.Package{Token: p.curToken}
	p.nextToken()
	expression.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	expression.Block = p.parseBlockStatement()
	return expression
}
