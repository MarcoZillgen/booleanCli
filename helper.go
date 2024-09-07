package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	RedColor    = "\033[1;31m"
	YellowColor = "\033[1;33m"
	ResetColor  = "\033[0m"
)

func getUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\nEnter a logical expression:")
	fmt.Print("> ")
	scanner.Scan()
	userInput := scanner.Text()
	return userInput
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isAtom(s string) bool {
	return isAlpha(s[0])
}

func printWarn(format string, a ...interface{}) {
	fmt.Printf(YellowColor+format+ResetColor, a...)
}

func printError(format string, a ...interface{}) {
	fmt.Printf(RedColor+format+ResetColor, a...)
}

func printTokens(tokens []Token) {
	for _, token := range tokens {
		fmt.Println(token)
	}
}

func printAST(node *ASTNode, level int) {
	if node == nil {
		return
	}

	fmt.Println(strings.Repeat("-", level), node.Value)
	printAST(node.Left, level+1)
	printAST(node.Right, level+1)
}

func printTruthTable(tt TruthTable) {
	// Print header
	for k := range tt.Variables {
		fmt.Printf("%-18s", k)
	}
	fmt.Println("Result\n")

	// Print table
	for i := 0; i < len(tt.Result); i++ {
		for _, v := range tt.Variables {
			fmt.Printf("%-18t", v[i])
		}
		fmt.Println(tt.Result[i])
	}
}
