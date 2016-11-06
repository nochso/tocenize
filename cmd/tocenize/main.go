package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nochso/tocenize"
)

const (
	_ = iota // ignore 0-1
	_
	ExitUsageError // 2 as used by flag
	ExitInputError
)

func main() {
	flag.Usage = func() {
		fmt.Println("tocenize [options] FILE...")
		fmt.Println()
		flag.PrintDefaults()
	}
	job := tocenize.Job{}
	flag.IntVar(&job.MinDepth, "min", 1, "minimum depth")
	flag.IntVar(&job.MaxDepth, "max", 99, "maximum depth")
	flag.BoolVar(&job.Diff, "d", false, "print diff")
	flag.BoolVar(&job.Print, "p", false, "print full output")
	flag.BoolVar(&job.Update, "u", true, "update existing file")
	flag.Parse()

	if flag.NArg() == 0 {
		exit("too few arguments", ExitUsageError)
	}

	for _, path := range flag.Args() {
		doc, err := tocenize.NewDocument(path)
		if err != nil {
			exit(err.Error(), ExitInputError)
		}
		toc := tocenize.NewTOC(doc, job)
		doc.Update(toc, job)
	}
}

func exit(msg string, status int) {
	if status > 0 {
		fmt.Print("error: ")
	}
	fmt.Println(msg)
	fmt.Printf("exit code: %d\n", status)
	os.Exit(status)
}
