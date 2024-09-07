package main

func validateAST(ast *ASTNode) bool {
	if ast == nil {
		printError("Error: Invalid syntax\n")
		return false
	}

	if isAtom(ast.Value) {
		if ast.Left != nil || ast.Right != nil {
			printError("Error: Atom cannot have children\n")
			return false
		}
		if len(ast.Value) > 16 {
			printError("Error: Atom name too long "+ast.Value+"\n")
			return false
		}
	} else if ast.Value == "!" {
		if ast.Left != nil && ast.Right != nil {
			printError("Error: Unary operator cannot have two children\n")
			return false
		}
	} else {
		if ast.Left == nil || ast.Right == nil {
			printError("Error: Binary operator must have two children\n")
			return false
		}
	}

	return true
}
