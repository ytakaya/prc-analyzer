package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ytakaya/prc-parser/parser"
	"github.com/ytakaya/prc-parser/visitor"
)

func dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}

type PrcData struct {
	Name  string `json:"name"`
	Rules []int  `json:"rules"`
}

func main() {
	dir := flag.String("d", "", "target dir")
	flag.Parse()
	if *dir == "" {
		flag.Usage()
		os.Exit(1)
	}

	paths := dirwalk(*dir)
	prcData := make([]PrcData, 0)
	for i, p := range paths {
		fmt.Printf("%s ** %d ** %s\n", time.Now().Format("15:04"), i, p)
		if strings.Contains(p, ".tbl") || strings.Contains(p, ".syn") {
			continue
		}

		prcVisitor := visitor.PRCVisitor{}

		input, _ := antlr.NewFileStream(p)
		lexer := parser.NewTSqlLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		p := parser.NewTSqlParser(stream)

		p.AddErrorListener(antlr.NewDefaultErrorListener())
		p.BuildParseTrees = true
		p.Tsql_file().Accept(&prcVisitor)

		for _, prcCtx := range prcVisitor.PrcCtx {
			prcName := prcCtx.GetProcName()
			if prcName == nil {
				continue
			}
			l := &visitor.PrcListener{}
			antlr.ParseTreeWalkerDefault.Walk(l, prcCtx)
			prcData = append(prcData, PrcData{prcName.GetText(), l.Rules})
		}
		fmt.Println("---")
		if i%100 == 0 {
			json, _ := json.Marshal(prcData)
			err := ioutil.WriteFile("./prc.json", json, 0777)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
	json, _ := json.Marshal(prcData)
	err := ioutil.WriteFile("./prc.json", json, 0777)
	if err != nil {
		fmt.Println(err.Error())
	}
}
