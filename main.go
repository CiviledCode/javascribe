package main

import (
	"fmt"
	"io"
	"os"

	"github.com/civiledcode/gofast-testing/dfa"
	"github.com/civiledcode/gofast-testing/visitors"
	"github.com/t14raptor/go-fast/parser"
)

func main() {
	f, err := os.Open("./input.js")
	if err != nil {
		panic(err)
	}

	jsCode, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	f.Close()

	a, err := parser.ParseFile(string(jsCode))
	if err != nil {
		panic(err)
	}

	logVisitor := visitors.LogVisitor{}

	a.VisitWith(&logVisitor)

	//fmt.Println("\n\n\n\n")

	rdaCtx := dfa.CreateContextRDA(256)
	rdaCtx.Debug = true

	rdaCtx.Start(a)

	//fmt.Println("\n\n\n\n")

	for _, ud := range rdaCtx.UseDefs {
		fmt.Println("Use:", ud.Usage.Name)
		fmt.Println("Defs:", ud.Definitions)
		fmt.Println("")
	}
}
