package roa

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cloudflare/cfrpki/validator/lib"

	"github.com/natesales/rpkiquery/internal/util"
)

// Load parses a ROA and returns an output string
func Load(decoder *librpki.DecoderConfig, roaBytes []byte, format string) (string, error) {
	out := "Type: ROA\n"
	roa, err := decoder.DecodeROA(roaBytes)
	if err != nil {
		return "", fmt.Errorf("ROA decode: %s", err)
	}

	format = strings.ToLower(format)
	if format == "json" {
		jsonBytes, err := json.Marshal(roa)
		if err != nil {
			return "", fmt.Errorf("JSON decode: %s", err)
		}
		fmt.Println(string(jsonBytes))
	} else if format == "plain" {
		out += fmt.Sprintf("Subject key identifier: %s\n", util.FormatHex(roa.Certificate.Certificate.SubjectKeyId))
		out += fmt.Sprintf("Authority key identifier: %s\n", util.FormatHex(roa.Certificate.Certificate.AuthorityKeyId))
		out += fmt.Sprintf("Certificate validity: %v to %v\n", roa.Certificate.Certificate.NotBefore, roa.Certificate.Certificate.NotAfter)
		out += fmt.Sprintf("asID: %d\n", roa.ASN)
		for i, entry := range roa.Entries {
			out += fmt.Sprintf("    %d: %s (max %d)\n", i+1, entry.IPNet, entry.MaxLength)
		}
	}
	return out, nil
}
