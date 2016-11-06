package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/k0kubun/pp"
)

type Document struct {
	Path  string
	Lines []string
	eol   string
}

func NewDocument(path string) Document {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		exit(err.Error(), ExitInputError)
	}
	doc := Document{
		Path: path,
		// Lines: strings.SplitAfter(string(b), "\n"),
	}
	doc.Lines, doc.eol = lines(string(b))
	return doc
}

func lines(s string) (lines []string, eol string) {
	lines = strings.Split(s, "\n")
	eol = "\n"
	for _, l := range lines {
		if len(l) > 0 && l[len(l)-1] == '\r' {
			eol = "\r\n"
			l = strings.TrimRight(l, "\r")
		}
	}
	return
}

func (d Document) Update(toc TOC, job Job) {
	s, e := d.FindTOC()
	if s == -1 {
		s, e = d.SuggestTOC(toc)
	}
	nd := Document{Path: d.Path, eol: d.eol}
	if s > -1 && e > -1 {
		nd.Lines = make([]string, s)
		copy(nd.Lines, d.Lines[:s])
		for _, tocLine := range toc {
			nd.Lines = append(nd.Lines, tocLine.string(toc.MinDepth()))
		}
		nd.Lines = append(nd.Lines, d.Lines[e:]...)
	}
	pp.Println(nd)
	fmt.Printf("%d %d\n", s, e)
}

// SuggestTOC looks for the first heading below a root heading.
// A root heading has minimum depth and the depth must only be used once.
// e.g.
//
// 	# Name
// 	--here--
// 	## A
// 	## B
func (d Document) SuggestTOC(toc TOC) (start, end int) {
	minCount := 0
	for _, tocLine := range toc {
		if tocLine.Depth == toc.MinDepth() {
			minCount++
			if minCount > 1 {
				// too many root headings
				break
			}
		}
		if tocLine.Depth > toc.MinDepth() {
			if minCount == 1 {
				fmt.Printf("found end of root paragraph on line %d for new TOC\n", tocLine.Index)
				return tocLine.Index, tocLine.Index
			}
			// ## appears before # which is odd
			break
		}
	}
	// in doubt, insert at top
	fmt.Println("chose first line for new TOC (unable to find root paragraph or existing TOC)")
	return 0, 0
}

func (d Document) FindTOC() (start, end int) {
	start = -1
	end = -1
	curStart := -1
	curEnd := -1
	isToc := false
	for i, l := range d.Lines {
		if reIsTOCLine.MatchString(l) {
			if !isToc {
				curStart = i
				isToc = true
			}
			curEnd = i
			continue
		}
		if isToc {
			isToc = false
			if curEnd-curStart > end-start {
				start = curStart
				end = curEnd
			}
		}
	}
	if start > -1 {
		fmt.Printf("found existing TOC on lines %d-%d\n", start+1, end+1)
	}
	return start, end
}

var reIsTOCLine = regexp.MustCompile(`^\s*- \[.*\]\(#.*\)\s*$`)
