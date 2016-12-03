package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/kylelemons/godebug/diff"
	"github.com/nochso/tocenize"
)

var VERSION = "?"

func main() {
	log.SetFlags(0)
	flag.Usage = func() {
		fmt.Println("tocenize [options] FILE...")
		fmt.Println()
		flag.PrintDefaults()
	}
	job := tocenize.Job{}
	flag.IntVar(&job.MinDepth, "min", 1, "minimum depth")
	flag.IntVar(&job.MaxDepth, "max", 99, "maximum depth")
	flag.StringVar(&tocenize.Indent, "indent", "\t", "string used for nesting")
	flag.BoolVar(&job.Diff, "d", false, "print full diff to stdout")
	flag.BoolVar(&job.Print, "p", false, "print full result to stdout")
	flag.BoolVar(&job.ExistingOnly, "e", false, "update only existing TOC (no insert)")
	flag.BoolVar(&tocenize.Verbose, "v", false, "verbose output")
	showVersion := flag.Bool("V", false, "print version")
	flag.Parse()

	if *showVersion {
		fmt.Println(VERSION)
		os.Exit(0)
	}
	if flag.NArg() == 0 {
		log.Println("too few arguments")
		flag.Usage()
		os.Exit(2)
	}

	for _, arg := range flag.Args() {
		paths, err := filepath.Glob(arg)
		if err != nil {
			log.Println(err)
			continue
		}

		for _, path := range paths {
			updateFile(path, job)
		}
	}
}

func updateFile(path string, job tocenize.Job) {
	log.SetPrefix(path + ": ")
	doc, err := tocenize.NewDocument(path)
	if err != nil {
		log.Println(err)
		return
	}
	toc := tocenize.NewTOC(doc, job)
	newDoc, err := doc.Update(toc, job.ExistingOnly)
	if err != nil {
		log.Println(err)
		return
	}
	if job.Diff {
		vlog("Diff -old +new")
		log.Println()
		d := diff.Diff(doc.String(), newDoc.String())
		if d != "" {
			fmt.Println(d)
		}
		return
	}
	if job.Print {
		vlog("printing full result")
		fmt.Println(newDoc.String())
		return
	}
	vlog("updating file")
	err = ioutil.WriteFile(doc.Path, []byte(newDoc.String()), 0644)
	if err != nil {
		log.Println(err)
	}
}

func vlog(s string) {
	if tocenize.Verbose {
		log.Println(s)
	}
}
