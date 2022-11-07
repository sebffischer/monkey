package ast

import "monkey/token"

// This is a node in our AST. Each node must implement the function TokenLiteral()
type Node interface {
	TokenLiteral() string
}

// A statement is something that does not produce a value.
type Statement interface {
	Node
	statementNode() // This is only a dummy method and I am not exactly sure what it is used for
}

// An expression is something that produces a value.
type Expression interface {
	Node
	expressionNode() // This is only a dummy method and I am not exactly sure what it is used for
}

// A program is a collection of statements.
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (ls *ReturnStatement) statementNode() {}
func (ls *ReturnStatement) TokenLiteral() string {
  return ls.Token.Literal
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
