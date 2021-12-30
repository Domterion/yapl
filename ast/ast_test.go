package ast

import (
	"testing"

	"github.com/domterion/yapl/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}

func TestProgram_TokenLiteral(t *testing.T) {
	type fields struct {
		Statements []Statement
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Program{
				Statements: tt.fields.Statements,
			}
			if got := p.TokenLiteral(); got != tt.want {
				t.Errorf("Program.TokenLiteral() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProgram_String(t *testing.T) {
	type fields struct {
		Statements []Statement
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Program{
				Statements: tt.fields.Statements,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("Program.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExpressionStatement_String(t *testing.T) {
	type fields struct {
		Token      token.Token
		Expression Expression
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			es := &ExpressionStatement{
				Token:      tt.fields.Token,
				Expression: tt.fields.Expression,
			}
			if got := es.String(); got != tt.want {
				t.Errorf("ExpressionStatement.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIdentifier_String(t *testing.T) {
	type fields struct {
		Token token.Token
		Value string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Identifier{
				Token: tt.fields.Token,
				Value: tt.fields.Value,
			}
			if got := i.String(); got != tt.want {
				t.Errorf("Identifier.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLetStatement_TokenLiteral(t *testing.T) {
	type fields struct {
		Token token.Token
		Name  *Identifier
		Value Expression
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ls := &LetStatement{
				Token: tt.fields.Token,
				Name:  tt.fields.Name,
				Value: tt.fields.Value,
			}
			if got := ls.TokenLiteral(); got != tt.want {
				t.Errorf("LetStatement.TokenLiteral() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLetStatement_String(t *testing.T) {
	type fields struct {
		Token token.Token
		Name  *Identifier
		Value Expression
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ls := &LetStatement{
				Token: tt.fields.Token,
				Name:  tt.fields.Name,
				Value: tt.fields.Value,
			}
			if got := ls.String(); got != tt.want {
				t.Errorf("LetStatement.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIdentifier_TokenLiteral(t *testing.T) {
	type fields struct {
		Token token.Token
		Value string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Identifier{
				Token: tt.fields.Token,
				Value: tt.fields.Value,
			}
			if got := i.TokenLiteral(); got != tt.want {
				t.Errorf("Identifier.TokenLiteral() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReturnStatement_TokenLiteral(t *testing.T) {
	type fields struct {
		Token       token.Token
		ReturnValue Expression
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := &ReturnStatement{
				Token:       tt.fields.Token,
				ReturnValue: tt.fields.ReturnValue,
			}
			if got := rs.TokenLiteral(); got != tt.want {
				t.Errorf("ReturnStatement.TokenLiteral() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReturnStatement_String(t *testing.T) {
	type fields struct {
		Token       token.Token
		ReturnValue Expression
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := &ReturnStatement{
				Token:       tt.fields.Token,
				ReturnValue: tt.fields.ReturnValue,
			}
			if got := rs.String(); got != tt.want {
				t.Errorf("ReturnStatement.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExpressionStatement_TokenLiteral(t *testing.T) {
	type fields struct {
		Token      token.Token
		Expression Expression
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			es := &ExpressionStatement{
				Token:      tt.fields.Token,
				Expression: tt.fields.Expression,
			}
			if got := es.TokenLiteral(); got != tt.want {
				t.Errorf("ExpressionStatement.TokenLiteral() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntegerLiteral_TokenLiteral(t *testing.T) {
	type fields struct {
		Token token.Token
		Value int64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			il := &IntegerLiteral{
				Token: tt.fields.Token,
				Value: tt.fields.Value,
			}
			if got := il.TokenLiteral(); got != tt.want {
				t.Errorf("IntegerLiteral.TokenLiteral() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntegerLiteral_String(t *testing.T) {
	type fields struct {
		Token token.Token
		Value int64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			il := &IntegerLiteral{
				Token: tt.fields.Token,
				Value: tt.fields.Value,
			}
			if got := il.String(); got != tt.want {
				t.Errorf("IntegerLiteral.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrefixExpression_TokenLiteral(t *testing.T) {
	type fields struct {
		Token    token.Token
		Operator string
		Right    Expression
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pe := &PrefixExpression{
				Token:    tt.fields.Token,
				Operator: tt.fields.Operator,
				Right:    tt.fields.Right,
			}
			if got := pe.TokenLiteral(); got != tt.want {
				t.Errorf("PrefixExpression.TokenLiteral() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrefixExpression_String(t *testing.T) {
	type fields struct {
		Token    token.Token
		Operator string
		Right    Expression
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pe := &PrefixExpression{
				Token:    tt.fields.Token,
				Operator: tt.fields.Operator,
				Right:    tt.fields.Right,
			}
			if got := pe.String(); got != tt.want {
				t.Errorf("PrefixExpression.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
