package tocenize

import (
	"bytes"
	"strings"
)

type TOC struct {
	Headings []Heading
	eol      string
}

func (t TOC) MinDepth() int {
	if len(t.Headings) == 0 {
		return 0
	}
	m := t.Headings[0].Depth
	for _, h := range t.Headings {
		if h.Depth < m {
			m = h.Depth
		}
	}
	return m
}

func (t TOC) String() string {
	minDepth := t.MinDepth()
	buf := &bytes.Buffer{}
	for i, h := range t.Headings {
		if i > 0 {
			buf.WriteString(t.eol)
		}
		buf.WriteString(h.string(minDepth))
	}
	return buf.String()
}

func NewTOC(doc Document, pj Job) TOC {
	toc := TOC{eol: doc.eol}
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
			toc.Headings = append(toc.Headings, *heading)
		}
		prevLine = l
	}
	return toc
}
