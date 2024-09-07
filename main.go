package main

func main() {
	for {
		userInput := getUserInput()

		tokens := lexer(userInput)
		if tokens == nil {
			continue
		}
		// printTokens(tokens)

		ast := parser(&tokens)
		if ast == nil {
			continue
		}
		printAST(ast, 0)

		truthTable := solveTruthTable(ast)
		printTruthTable(truthTable)
	}
}
