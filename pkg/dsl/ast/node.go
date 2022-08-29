package ast

import (
	"bytes"
	"strings"

	"github.com/Permify/permify/pkg/dsl/token"
	"github.com/Permify/permify/pkg/tuple"
)

// Node -
type Node interface {
	TokenLiteral() string
	String() string
}

// Expression -
type Expression interface {
	Node
	expressionNode()
	IsInfix() bool
	Type() string
}

// Statement -
type Statement interface {
	Node
	statementNode()
}

// Schema -
type Schema struct {
	Statements []Statement
}

// Validate -
func (sch Schema) Validate() (err error) {
	for _, st := range sch.Statements {
		name := st.(*EntityStatement).Name.Literal
		if name == tuple.USER {
			return nil
		}
	}
	return UserEntityRequiredErr
}

// EntityStatement -
type EntityStatement struct {
	Token              token.Token // token.ENTITY
	Name               token.Token // token.IDENT
	RelationStatements []Statement
	ActionStatements   []Statement
	Option             token.Token // token.OPTION
}

// statementNode -
func (ls *EntityStatement) statementNode() {}

// TokenLiteral -
func (ls *EntityStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// String -
func (ls *EntityStatement) String() string {
	var sb strings.Builder
	sb.WriteString("entity")
	sb.WriteString(" ")
	sb.WriteString(ls.Name.Literal)
	sb.WriteString(" {")
	sb.WriteString("\n")

	for _, rs := range ls.RelationStatements {
		sb.WriteString(rs.String())
		sb.WriteString("\n")
	}

	sb.WriteString("\n")

	for _, rs := range ls.ActionStatements {
		sb.WriteString(rs.String())
		sb.WriteString("\n")
	}

	sb.WriteString("}")
	sb.WriteString(" ")

	if ls.Option.Literal != "" {
		sb.WriteString("`")
		sb.WriteString(ls.Option.Literal)
		sb.WriteString("`")
	}

	sb.WriteString("\n")
	return sb.String()
}

// RelationStatement -
type RelationStatement struct {
	Token         token.Token // token.RELATION
	Name          token.Token // token.IDENT
	RelationTypes []Statement
	Option        token.Token // token.OPTION
}

// statementNode -
func (ls *RelationStatement) statementNode() {}

// TokenLiteral -
func (ls *RelationStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// String -
func (ls *RelationStatement) String() string {
	var sb strings.Builder
	sb.WriteString("\trelation")
	sb.WriteString(" ")
	sb.WriteString(ls.Name.Literal)
	sb.WriteString(" ")

	for _, rs := range ls.RelationTypes {
		sb.WriteString(rs.String())
		sb.WriteString(" ")
	}

	sb.WriteString(" ")

	if ls.Option.Literal != "" {
		sb.WriteString("`")
		sb.WriteString(ls.Option.Literal)
		sb.WriteString("`")
	}

	return sb.String()
}

// RelationTypeStatement -
type RelationTypeStatement struct {
	Sign  token.Token // token.sign
	Token token.Token // token.IDENT
}

// statementNode -
func (ls *RelationTypeStatement) statementNode() {}

// TokenLiteral -
func (ls *RelationTypeStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *RelationTypeStatement) String() string {
	var sb strings.Builder
	sb.WriteString(ls.Sign.Literal)
	sb.WriteString(ls.Token.Literal)
	return sb.String()
}

// Identifier -
type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

// statementNode -
func (ls *Identifier) expressionNode() {}

// TokenLiteral -
func (ls *Identifier) TokenLiteral() string {
	return ls.Token.Literal
}

// String -
func (ls *Identifier) String() string {
	return ls.Value
}

// IsInfix -
func (ls *Identifier) IsInfix() bool {
	return false
}

// Type -
func (ls *Identifier) Type() string {
	return "identifier"
}

// ActionStatement -
type ActionStatement struct {
	Token               token.Token // token.ACTION
	Name                token.Token // token.IDENT
	ExpressionStatement Statement
}

// statementNode -
func (ls *ActionStatement) statementNode() {}

// TokenLiteral -
func (ls *ActionStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// String -
func (ls *ActionStatement) String() string {
	var sb strings.Builder
	sb.WriteString("\t" + ls.TokenLiteral() + " ")
	sb.WriteString(ls.Name.Literal)
	sb.WriteString(" = ")
	if ls.ExpressionStatement != nil {
		sb.WriteString(ls.ExpressionStatement.String())
	}
	return sb.String()
}

// ExpressionStatement struct
type ExpressionStatement struct {
	Expression Expression
}

// statementNode function on ExpressionStatement
func (es *ExpressionStatement) statementNode() {}

// TokenLiteral function on ExpressionStatement
func (es *ExpressionStatement) TokenLiteral() string {
	return "start"
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// InfixExpression -
type InfixExpression struct {
	Token    token.Token // The operator token, e.g. and, or
	Left     Expression
	Operator string
	Right    Expression
}

// expressionNode -
func (ie *InfixExpression) expressionNode() {}

// TokenLiteral -
func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

// String -
func (ie *InfixExpression) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(ie.Left.String())
	sb.WriteString(" " + ie.Operator)
	sb.WriteString(" ")
	sb.WriteString(ie.Right.String())
	sb.WriteString(")")
	return sb.String()
}

// IsInfix -
func (ie *InfixExpression) IsInfix() bool {
	return true
}

// Type -
func (ie *InfixExpression) Type() string {
	return "inflix"
}

// PrefixExpression -
type PrefixExpression struct {
	Token    token.Token // not
	Operator string
	Value    string
}

// TokenLiteral -
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }

// String -
func (pe *PrefixExpression) String() string {
	var sb bytes.Buffer
	sb.WriteString(pe.Operator)
	sb.WriteString(" ")
	sb.WriteString(pe.Value)
	return sb.String()
}

// expressionNode -
func (pe *PrefixExpression) expressionNode() {}

// IsInfix -
func (pe *PrefixExpression) IsInfix() bool {
	return false
}

// Type -
func (pe *PrefixExpression) Type() string {
	return "prefix"
}
