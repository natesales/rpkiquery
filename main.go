package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/natesales/rpkiquery/internal/cert"
	"github.com/natesales/rpkiquery/internal/roa"
	"github.com/natesales/rpkiquery/internal/tal"

	"github.com/cloudflare/cfrpki/validator/lib"
)

var (
	file           = flag.String("f", "", "input file")
	inputFormat    = flag.String("i", "auto", "input format (roa|tal|cert|auto)")
	outputFormat   = flag.String("o", "plain", "output format (plain|json)")
	validateStrict = flag.Bool("s", false, "use strict validation")
)

func main() {
	flag.Parse()

	if *file == "" {
		if len(os.Args) > 1 {
			file = &os.Args[len(os.Args)-1]
		} else {
			log.Fatal("input file must be defined")
		}
	}

	fileBytes, err := os.ReadFile(*file)
	if err != nil {
		log.Fatal(err)
	}
	decoder := &librpki.DecoderConfig{ValidateStrict: *validateStrict}

	var out string
	switch *inputFormat {
	case "roa":
		out, err = roa.Load(decoder, fileBytes, *outputFormat)
		if err != nil {
			log.Fatal(err)
		}
	case "tal":
		out, err = tal.Load(fileBytes, *outputFormat)
		if err != nil {
			log.Fatal(err)
		}
	case "cert":
		out, err = cert.Load(fileBytes, *outputFormat)
		if err != nil {
			log.Fatal(err)
		}
	case "auto":
		out = "unable to automatically detect format"
		out, err = roa.Load(decoder, fileBytes, *outputFormat)
		if err == nil {
			break
		}
		out, err = tal.Load(fileBytes, *outputFormat)
		if err == nil {
			break
		}
		out, err = cert.Load(fileBytes, *outputFormat)
		if err == nil {
			break
		}
	default:
		log.Fatal("input format must be one of tal|roa|cer|auto")
	}
	fmt.Print(out)
}
