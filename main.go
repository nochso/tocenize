package main

import (
	"flag"
	"fmt"
	"os"
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
	job := Job{}
	flag.IntVar(&job.MinDepth, "min", 1, "minimum depth")
	flag.IntVar(&job.MaxDepth, "max", 99, "maximum depth")
	flag.Parse()

	if flag.NArg() == 0 {
		exit("too few arguments", ExitUsageError)
	}

	for _, path := range flag.Args() {
		doc := NewDocument(path)
		toc := NewTOC(doc, job)
		doc.Update(toc, job)
		fmt.Println(toc)
	}
}

type Job struct {
	MinDepth int
	MaxDepth int
}

func exit(msg string, status int) {
	if status > 0 {
		fmt.Print("error: ")
	}
	fmt.Println(msg)
	fmt.Printf("exit code: %d\n", status)
	os.Exit(status)
}
