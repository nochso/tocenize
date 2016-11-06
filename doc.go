package tocenize

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/kylelemons/godebug/diff"
)

var Verbose = false

type Job struct {
	MinDepth int
	MaxDepth int
	Diff     bool
	Print    bool
	Update   bool
}

type Document struct {
	Path  string
	Lines []string
	eol   string
}

func NewDocument(path string) (Document, error) {
	doc := Document{
		Path: path,
	}
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return doc, err
	}
	doc.Lines, doc.eol = lines(string(b))
	return doc, nil
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

func (d Document) String() string {
	return strings.Join(d.Lines, d.eol)
}

func (d Document) Update(toc TOC, job Job) error {
	s, e := d.FindTOC()
	if s == -1 {
		s, e = d.SuggestTOC(toc)
	}
	nd := Document{Path: d.Path, eol: d.eol}
	if s > -1 && e > -1 {
		nd.Lines = make([]string, s)
		copy(nd.Lines, d.Lines[:s])
		for _, tocLine := range toc.Headings {
			nd.Lines = append(nd.Lines, tocLine.string(toc.MinDepth()))
		}
		nd.Lines = append(nd.Lines, d.Lines[e:]...)
	}
	if job.Diff {
		if Verbose {
			log.Println("Diff -old +new")
			fmt.Println()
		}
		fmt.Println(diff.Diff(d.String(), nd.String()))
		return nil
	}
	if job.Print {
		if Verbose {
			log.Println("printing full result")
			fmt.Println()
		}
		fmt.Println(nd.String())
		return nil
	}
	if job.Update {
		if Verbose {
			log.Println("updating file")
		}
		err := ioutil.WriteFile(d.Path, []byte(nd.String()), 0644)
		return err
	}
	return nil
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
	secCount := 0
	minIndex := 0
	for _, tocLine := range toc.Headings {
		if tocLine.Depth == toc.MinDepth() {
			minCount++
			if minCount > 1 {
				// too many root headings
				break
			}
			minIndex = tocLine.Index
		}
		if tocLine.Depth > toc.MinDepth() && minCount == 1 && secCount == 0 {
			secCount++
			minIndex = tocLine.Index
		}
	}
	if minCount == 1 {
		log.Printf("found end of root paragraph on line %d", minIndex+1)
		return minIndex, minIndex
	}
	if len(toc.Headings) > 0 {
		start = toc.Headings[0].Index
		if Verbose {
			log.Printf("chose line %d before first significant heading", start+1)
		}
		return start, start
	}
	// in doubt, insert at top
	if Verbose {
		log.Println("chose first line for new TOC (unable to find root paragraph or existing TOC)")
	}
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
			curEnd = i + 1
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
	if start > -1 && Verbose {
		log.Printf("found existing TOC on lines %d-%d", start+1, end+1)
	}
	return start, end
}

var reIsTOCLine = regexp.MustCompile(`^\s*- \[.*\]\(#.*\)\s*$`)
