package cert

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cloudflare/cfrpki/validator/lib"

	"github.com/natesales/rpkiquery/internal/util"
)

// Load parses a certificate file and returns an output string
func Load(cerBytes []byte, format string) (string, error) {
	out := "Type: Certificate\n"
	cert, err := librpki.DecodeCertificate(cerBytes)
	if err != nil {
		return "", fmt.Errorf("certificate decode: %s", err)
	}

	format = strings.ToLower(format)
	if format == "json" {
		jsonBytes, err := json.Marshal(cert)
		if err != nil {
			return "", fmt.Errorf("certificate decode: %s", err)
		}
		fmt.Println(string(jsonBytes))
	} else if format == "plain" {
		out += fmt.Sprintf("Subject key identifier: %s\n", util.FormatHex(cert.Certificate.SubjectKeyId))
		out += fmt.Sprintf("Authority key identifier: %s\n", util.FormatHex(cert.Certificate.AuthorityKeyId))
		out += fmt.Sprintf("Certificate validity: %v to %v\n", cert.Certificate.NotBefore, cert.Certificate.NotAfter)
		out += fmt.Sprintf("SIAs:\n")
		for _, sia := range cert.SubjectInformationAccess {
			out += fmt.Sprintf("  - %s\n", sia.String())
		}
		out += fmt.Sprintf("IPs:\n")
		for _, ip := range cert.IPAddresses {
			out += fmt.Sprintf("  - %s\n", ip.String())
		}
		out += fmt.Sprintf("ASNs:\n")
		for _, asn := range append(cert.ASNums, cert.ASNRDI...) {
			out += fmt.Sprintf("  - %s\n", asn.String())
		}
	}
	return out, nil
}
