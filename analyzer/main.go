package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type PrcData struct {
	Name  string `json:"name"`
	Rules []int  `json:"rules"`
}

func main() {
	file := flag.String("d", "", "target file")
	flag.Parse()
	if *file == "" {
		flag.Usage()
		os.Exit(1)
	}

	raw, err := ioutil.ReadFile(*file)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var prcData []PrcData
	err = json.Unmarshal(raw, &prcData)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(len(prcData))

	ruleMap := make(map[string][]string)
	for _, prc := range prcData {
		ruleStr := strings.Trim(strings.Join(strings.Split(fmt.Sprint(prc.Rules), " "), "-"), "[]")
		if _, ok := ruleMap[ruleStr]; !ok {
			ruleMap[ruleStr] = []string{prc.Name}
		} else {
			ruleMap[ruleStr] = append(ruleMap[ruleStr], prc.Name)
		}
	}
	fmt.Println(len(ruleMap))

	for _, v := range ruleMap {
		if len(v) > 1 {
			fmt.Println(v)
			break
		}
	}
}
