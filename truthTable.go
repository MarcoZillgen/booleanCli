package main

import (
	"math"
)

type TruthTable struct {
	Variables map[string][]bool
	Result    []bool
}

func newTruthTable(vars []string) TruthTable {
	tt := TruthTable{
		Variables: make(map[string][]bool),
		Result:    []bool{},
	}

	for i, v := range vars {
		tt.Variables[v] = []bool{}
		falses := []bool{}
		trues := []bool{}
		for range int(math.Pow(2, float64(i))) {
			falses = append(falses, false)
			trues = append(trues, true)
		}

		for len(tt.Variables[v]) < int(math.Pow(2, float64(len(vars)))) {
			tt.Variables[v] = append(tt.Variables[v], falses...)
			tt.Variables[v] = append(tt.Variables[v], trues...)
		}
	}

	return tt
}

func getVariables(ast *ASTNode) []string {
	vars := []string{}

	if isAtom(ast.Value) {
		vars = append(vars, ast.Value)
	} else {
		if ast.Left != nil {
			vars = append(vars, getVariables(ast.Left)...)
		}
		if ast.Right != nil {
			vars = append(vars, getVariables(ast.Right)...)
		}
	}

	return vars
}

func solveAST(ast *ASTNode, vars map[string]bool) bool {
	switch ast.Value {
	case "&":
		return solveAST(ast.Left, vars) && solveAST(ast.Right, vars)
	case "!&":
		return !(solveAST(ast.Left, vars) && solveAST(ast.Right, vars))
	case "|":
		return solveAST(ast.Left, vars) || solveAST(ast.Right, vars)
	case "!|":
		return !(solveAST(ast.Left, vars) || solveAST(ast.Right, vars))
	case "^":
		return solveAST(ast.Left, vars) != solveAST(ast.Right, vars)
	case "!":
		return !solveAST(ast.Left, vars)
	case "->":
		if solveAST(ast.Left, vars) {
			return solveAST(ast.Right, vars)
		}
		return true
	case "<->":
		return solveAST(ast.Left, vars) == solveAST(ast.Right, vars)
	default:
		if isAtom(ast.Value) {
			return vars[ast.Value]
		}
	}

	return false
}

func solveTruthTable(ast *ASTNode) TruthTable {
	vars := getVariables(ast)

	tt := newTruthTable(vars)

	for tableRun := 0; tableRun < int(math.Pow(2, float64(len(vars)))); tableRun++ {
		currVars := make(map[string]bool)
		for _, v := range vars {
			currVars[v] = tt.Variables[v][tableRun]
		}

		tt.Result = append(tt.Result, solveAST(ast, currVars))
	}

	return tt
}
