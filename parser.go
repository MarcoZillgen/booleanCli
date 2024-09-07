package main

type ASTNode struct {
	Left  *ASTNode
	Right *ASTNode
	Value string
}

func newASTNode(left *ASTNode, right *ASTNode, value string) *ASTNode {
	return &ASTNode{
		Left:  left,
		Right: right,
		Value: value,
	}
}

func parser(tokens *[]Token) *ASTNode {
	if len(*tokens) == 0 {
		return nil
	}

	return parseEquivalency(tokens)
}

func parseEquivalency(tokens *[]Token) *ASTNode {
	left := parseImplies(tokens)

	for len(*tokens) != 0 && (*tokens)[0].Type == Equivalency {
		*tokens = (*tokens)[1:]

		right := parseImplies(tokens)
		left = newASTNode(left, right, "<->")
	}

	return left
}

func parseImplies(tokens *[]Token) *ASTNode {
	left := parseOr(tokens)

	for len(*tokens) != 0 && (*tokens)[0].Type == Implies {
		*tokens = (*tokens)[1:]

		right := parseOr(tokens)
		left = newASTNode(left, right, "->")
	}

	return left
}

func parseOr(tokens *[]Token) *ASTNode {
	left := parseXor(tokens)

	for len((*tokens)) != 0 && (*tokens)[0].Type == Or {
		(*tokens) = (*tokens)[1:]

		right := parseXor(tokens)
		left = newASTNode(left, right, "|")
	}

	return left
}

func parseXor(tokens *[]Token) *ASTNode {
	left := parseAnd(tokens)

	for len((*tokens)) != 0 && (*tokens)[0].Type == Xor {
		(*tokens) = (*tokens)[1:]

		right := parseAnd(tokens)
		left = newASTNode(left, right, "^")
	}

	return left
}

func parseAnd(tokens *[]Token) *ASTNode {
	left := parseNand(tokens)

	for len((*tokens)) != 0 && (*tokens)[0].Type == And {
		(*tokens) = (*tokens)[1:]

		right := parseNand(tokens)
		left = newASTNode(left, right, "&")
	}

	return left
}

func parseNand(tokens *[]Token) *ASTNode {
	left := parseNor(tokens)

	for len((*tokens)) != 0 && (*tokens)[0].Type == Nand {
		(*tokens) = (*tokens)[1:]

		right := parseNor(tokens)
		left = newASTNode(left, right, "!&")
	}

	return left
}

func parseNor(tokens *[]Token) *ASTNode {
	left := parseNot(tokens)

	for len((*tokens)) != 0 && (*tokens)[0].Type == Nor {
		(*tokens) = (*tokens)[1:]

		right := parseNot(tokens)
		left = newASTNode(left, right, "!|")
	}

	return left
}

func parseNot(tokens *[]Token) *ASTNode {
	left := parseAtom(tokens)

	for len((*tokens)) != 0 && (*tokens)[0].Type == Not {
		(*tokens) = (*tokens)[1:]

		right := parseAtom(tokens)
		left = newASTNode(nil, right, "!")
	}

	return left
}

func parseAtom(tokens *[]Token) *ASTNode {
	if len((*tokens)) == 0 {
		return nil
	}

	if (*tokens)[0].Type == Identifier {
		value := (*tokens)[0].Value
		(*tokens) = (*tokens)[1:]
		return newASTNode(nil, nil, value)
	}

	if (*tokens)[0].Type == LeftParen {
		(*tokens) = (*tokens)[1:]

		node := parser(tokens)

		if len((*tokens)) == 0 || (*tokens)[0].Type != RightParen {
			return nil
		}

		(*tokens) = (*tokens)[1:]
		return node
	}

	return nil
}
