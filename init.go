package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func main() {
	pn := flag.String("pn", "", "name of the problems")
	flag.Parse()

	err := os.Mkdir(*pn, os.ModeDir)
	if err != nil {
		log.Error(err)
	}

	var packageName string
	for i := 0; i < len(*pn); i++ {
		if isalnum((*pn)[i]) {
			packageName += string((*pn)[i])
		}
	}

	fileName := []string{"solution.go", "solution_test.go", "README.md"}
	for _, name := range fileName {
		file, err := os.Create(*pn + "/" + name)
		if err != nil {
			log.Error(err)
		}
		defer file.Close()

		switch filepath.Base(file.Name()) {
		case "solution.go":
			content := fmt.Sprintf("package %s\nfunc solution1(s string){}", packageName)
			file.WriteString(content)
		case "solution_test.go":
			content := fmt.Sprintf("package %s\nimport (\n\t\"testing\"\n)\nfunc TestSolution(t *testing.T){\n\ttestCase := []struct {\n\t\ts string\n\t\texpected bool}{}\n\tfor _, tc:= range testCase {}}", packageName)
			file.WriteString(content)
		}

	}

}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
