package main

type TokenType int

const (
	And         TokenType = iota // &
	Nand                         // !&
	Or                           // |
	Xor                          // ^
	Nor                          // !|
	Not                          // !
	Implies                      // ->
	Equivalency                  // <->
	LeftParen                    // (
	RightParen                   // )
	Identifier                   // a-zA-Z
)

type Token struct {
	Type  TokenType
	Value string
}

func lexer(input string) []Token {
	tokens := []Token{}
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case ' ', '\t', '\n':
			continue
		case '&':
			tokens = append(tokens, Token{Type: And, Value: "&"})
		case '|':
			tokens = append(tokens, Token{Type: Or, Value: "|"})
		case '!':
			if i+1 < len(input) && input[i+1] == '&' {
				tokens = append(tokens, Token{Type: Nand, Value: "!&"})
				i++
			} else if i+1 < len(input) && input[i+1] == '|' {
				tokens = append(tokens, Token{Type: Nor, Value: "!|"})
				i++
			} else {
				tokens = append(tokens, Token{Type: Not, Value: "!"})
			}
		case '^':
			tokens = append(tokens, Token{Type: Xor, Value: "^"})
		case '<':
			if i+2 < len(input) && input[i+1] == '-' && input[i+2] == '>' {
				tokens = append(tokens, Token{Type: Equivalency, Value: "<->"})
				i += 2
			} else {
				printError("Error: invalid token [%s] at position %d\n", string(input[i]), i)
				return nil
			}
		case '-':
			if i+1 < len(input) && input[i+1] == '>' {
				tokens = append(tokens, Token{Type: Implies, Value: "->"})
				i++
			} else {
				printError("%sError: invalid token [%s] at position %d\n", RedColor, string(input[i]), i)
				return nil
			}
		case '(':
			tokens = append(tokens, Token{Type: LeftParen, Value: "("})
		case ')':
			tokens = append(tokens, Token{Type: RightParen, Value: ")"})
		default:
			if isAlpha(input[i]) {
				j := i
				for j < len(input) && isAlpha(input[j]) {
					j++
				}
				tokens = append(tokens, Token{Type: Identifier, Value: input[i:j]})
				i = j - 1
			} else {
				printError("Warning: invalid character [%s] at position %d\n", string(input[i]), i)
				return nil
			}
		}
	}

	return tokens
}
