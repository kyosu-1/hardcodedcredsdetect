package main

import (
	"github.com/kyosu-1/hardcodedcredsdetect"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(hardcodedcredsdetect.Analyzer) }
