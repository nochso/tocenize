package tocenize

import (
	"flag"
	"testing"

	"github.com/nochso/golden"
)

var update = flag.Bool("update", false, "update golden files")

func TestDocument_Update(t *testing.T) {
	job := Job{
		MinDepth: 1,
		MaxDepth: 99,
	}
	golden.TestDir(t, "test-fixtures", func(tc golden.Case) {
		doc, err := NewDocFromPath(tc.In.Path)
		if err != nil {
			tc.T.Fatal(err)
		}
		toc := NewTOC(doc, job)
		newDoc, err := doc.Update(toc, job.ExistingOnly)
		if err != nil {
			tc.T.Fatal(err)
		}
		if *update {
			tc.Out.Update([]byte(newDoc.String()))
		}
		tc.Diff(newDoc.String())
	})
}
