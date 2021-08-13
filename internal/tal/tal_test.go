package tal

import (
	"os"
	"path/filepath"
	"testing"
)

func TestTALLoad(t *testing.T) {
	talFiles, err := filepath.Glob("../../test/tal/*.tal")
	if len(talFiles) < 1 {
		t.Fatalf("no TAL files found")
	}
	if err != nil {
		t.Error(err)
	}
	for _, talFile := range talFiles {
		talBytes, err := os.ReadFile(talFile)
		if err != nil {
			t.Error(err)
		}
		if _, err := Load(talBytes, "plain"); err != nil {
			t.Error(err)
		}
	}
}
