package tocenize

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/writeas/go-strip-markdown"
)

// Indent string used for nesting. See `-indent`.
var Indent = "\t"

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
		strings.Repeat(Indent, h.Depth-minDepth),
		h.LinkTitle(),
		h.Anchor(),
	)
}

var (
	reLink    = regexp.MustCompile(`\[(.*?)\][\[\(].*?[\]\)]`)
	reLinkRef = regexp.MustCompile(`\[(.*?)\]`)
	reImage   = regexp.MustCompile(`\!\[.*?\]\s?[\[\(].*?[\]\)]`)
)

// LinkTitle returns a cleaned up title for use in links.
// It removes Markdown images and links, keeping only the link text.
func (h Heading) LinkTitle() string {
	t := h.Title
	t = reLink.ReplaceAllString(t, "$1")    // remove links, keeping the text
	t = reImage.ReplaceAllString(t, "")     // remove images
	t = reLinkRef.ReplaceAllString(t, "$1") // remove reference [link]
	t = strings.TrimSpace(t)
	return t
}

// adaptation of the anchor-cleanup used by Github's pipeline
// https://github.com/jch/html-pipeline/blob/master/lib/html/pipeline/toc_filter.rb
// \p{Word} = Letter, Mark, Number and Connector_Punctuation
var rePunct = regexp.MustCompile(`([^\p{L}\p{M}\p{N}\p{Pc}\- ])`)

func (h Heading) Anchor() string {
	// Strip Markdown
	a := stripmd.Strip(h.Title)
	a = toLowerASCII(a)
	a = rePunct.ReplaceAllString(a, "")
	a = strings.Replace(a, " ", "-", -1)
	if h.UniqueCounter > 0 {
		a = fmt.Sprintf("%s-%d", a, h.UniqueCounter)
	}
	return fmt.Sprintf("#%s", a)
}

// toLowerASCII is like strings.ToLower but considers ASCII only.
// This should mirror Ruby's str.downcase(:ascii) which is used by Github's
// pipeline: https://github.com/jch/html-pipeline/blob/master/lib/html/pipeline/toc_filter.rb
func toLowerASCII(s string) string {
	b := make([]byte, len(s))
	for i := range b {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		}
		b[i] = c
	}
	return string(b)
}
