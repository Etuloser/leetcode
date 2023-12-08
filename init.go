package main

import (
	"flag"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	pn := flag.String("pn", "", "name of the problems")
	flag.Parse()
	err := os.Mkdir(*pn, os.ModeDir)
	if err != nil {
		log.Error(err)
	}
	fileName := []string{"solution.go", "solution_test.go", "README.md"}
	for _, name := range fileName {
		file, err := os.Create(*pn + "/" + name)
		if err != nil {
			log.Error(err)
		}
		defer file.Close()
	}

}
