package golang

import (
	"bytes"
	"flag"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/andyleap/gencode/schema"
	"github.com/kr/pretty"
)

var update = flag.Bool("update", false, "if true, update golden files in ./testdata directory")

func TestGolangBackend(t *testing.T) {
	dir, err := ioutil.TempDir("", "gencode")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(dir)

	for _, tc := range []string{
		"array.schema",
	} {
		inputF := filepath.Join("./testdata", tc)
		outputF := filepath.Join(dir, tc+".go")
		goldenF := inputF+".golden"

		in, err := ioutil.ReadFile(inputF)
		if err != nil {
			t.Fatalf("%v: Failed to read: %v", tc, err)
		}
		s, err := schema.ParseSchema(bytes.NewReader(in))
		if err != nil {
			t.Fatalf("%v: Failed schema.ParseSchema: %v", tc, err)
		}

		b := GolangBackend{Package: "array"}
		g, err := b.Generate(s)
		if err != nil {
			t.Fatalf("%v: Failed Generate: %v", tc, err)
		}
		out := []byte(g)
		if err = ioutil.WriteFile(outputF, out, 0777); err != nil {
			t.Fatalf("%v: Failed to write generated file: %v", tc, err)
		}
		if out, err := exec.Command("go", "build", outputF).CombinedOutput(); err != nil {
			t.Fatalf("%v: Failed to compile generated code (error: %v):\n%s", tc, err, out)
		}

		want, err := ioutil.ReadFile(goldenF)
		needUpdate := true
		if err != nil {
			t.Errorf("%v: Failed to read golden file: %v", tc, err)
		} else if diff := pretty.Diff(want, out); len(diff) != 0 {
			t.Errorf("%v: Diff(want, got) = %v", tc, strings.Join(diff, "\n"))
		} else {
			needUpdate = false
		}

		if needUpdate && *update {
			if err := ioutil.WriteFile(goldenF, out, 0777); err != nil {
				t.Errorf("%v: Failed to update golden file: %v", tc, err)
			}
		}
	}
}
