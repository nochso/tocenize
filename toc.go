package main

import (
	"bytes"
	"strings"
)

type TOC []Heading

func (t TOC) MinDepth() int {
	if len(t) == 0 {
		return 0
	}
	m := t[0].Depth
	for _, h := range t {
		if h.Depth < m {
			m = h.Depth
		}
	}
	return m
}

func (t TOC) String() string {
	minDepth := t.MinDepth()
	buf := &bytes.Buffer{}
	for i, h := range t {
		if i > 0 {
			buf.WriteByte('\n')
		}
		buf.WriteString(h.string(minDepth))
	}
	return buf.String()
}

func NewTOC(doc Document, pj Job) TOC {
	toc := TOC{}
	prevLine := ""
	var heading *Heading
	for i, l := range doc.Lines {
		heading = nil
		if strings.HasPrefix(l, "#") {
			heading = NewHeadingATX(l, i)
		} else if prevLine != "" && l != "" && (strings.Trim(l, "=") == "" || strings.Trim(l, "-") == "") {
			heading = NewHeadingSE(prevLine, string(l[0]), i-1)
		}
		if heading != nil && heading.Depth <= pj.MaxDepth && heading.Depth >= pj.MinDepth {
			toc = append(toc, *heading)
		}
		prevLine = l
	}
	return toc
}
