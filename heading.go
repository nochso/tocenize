package tocenize

import (
	"fmt"
	"regexp"
	"strings"
)

type Heading struct {
	Title         string
	Depth         int
	Index         int
	UniqueCounter int
}

func NewHeadingATX(line string, index int) *Heading {
	return &Heading{
		Title: strings.Trim(line, "# "),
		Depth: len(line) - len(strings.TrimLeft(line, "#")),
		Index: index,
	}
}

func NewHeadingSE(title string, sep string, index int) *Heading {
	depth := 1
	if sep[0] == '-' {
		depth = 2
	}
	return &Heading{
		Title: strings.TrimSpace(title),
		Depth: depth,
		Index: index,
	}
}

func (h Heading) String() string {
	return h.string(1)
}

func (h Heading) string(minDepth int) string {
	return fmt.Sprintf(
		"%s- [%s](%s)",
		strings.Repeat("\t", h.Depth-minDepth),
		h.Title,
		h.Anchor(),
	)
}

var rePunct = regexp.MustCompile(`([^\w -]+)`)

func (h Heading) Anchor() string {
	a := strings.ToLower(h.Title)
	a = rePunct.ReplaceAllString(a, "")
	a = strings.Replace(a, " ", "-", -1)
	if h.UniqueCounter > 0 {
		a = fmt.Sprintf("%s-%d", a, h.UniqueCounter)
	}
	return fmt.Sprintf("#%s", a)
}
