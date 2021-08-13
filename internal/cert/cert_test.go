package cert

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCertLoad(t *testing.T) {
	certFiles, err := filepath.Glob("../../test/*/*.cer")

	if err != nil {
		t.Error(err)
	}
	for _, certFile := range certFiles {
		certBytes, err := os.ReadFile(certFile)
		if err != nil {
			t.Error(err)
		}
		if _, err := Load(certBytes, "plain"); err != nil {
			t.Error(err)
		}
	}
}
