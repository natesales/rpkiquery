package tal

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cloudflare/cfrpki/validator/lib"
)

// Load parses a TAL and returns an output string
func Load(talBytes []byte, format string) (string, error) {
	out := "Type: TAL\n"
	tal, err := librpki.DecodeTAL(talBytes)
	if err != nil {
		return "", fmt.Errorf("TAL decode: %s", err)
	}

	format = strings.ToLower(format)
	if format == "json" {
		jsonBytes, err := json.Marshal(tal)
		if err != nil {
			return "", fmt.Errorf("JSON decode: %s", err)
		}
		fmt.Println(string(jsonBytes))
	} else if format == "plain" {
		out += fmt.Sprintf("URI: %s\n", tal.URI)
		out += fmt.Sprintf("OID: %s\n", tal.OID)
		out += fmt.Sprintf("Algo: %s\n", tal.Algorithm.String())
	}
	return out, nil
}
