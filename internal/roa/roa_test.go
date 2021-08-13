package roa

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cloudflare/cfrpki/validator/lib"
)

func TestROALoad(t *testing.T) {
	roaFiles, err := filepath.Glob("../../test/roa/*.roa")
	if len(roaFiles) < 1 {
		t.Fatalf("no ROA files found")
	}
	if err != nil {
		t.Error(err)
	}
	for _, roaFile := range roaFiles {
		roaBytes, err := os.ReadFile(roaFile)
		if err != nil {
			t.Error(err)
		}
		decoder := &librpki.DecoderConfig{ValidateStrict: false}
		if _, err := Load(decoder, roaBytes, "plain"); err != nil {
			t.Error(err)
		}
	}
}
