package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, "no command given. 'decode' or 'encode' is expected with exactly one argument.")
		os.Exit(1)
	}

	decCmd := flag.NewFlagSet("decode", flag.ExitOnError)
	encCmd := flag.NewFlagSet("encode", flag.ExitOnError)

	switch os.Args[1] {
	case "decode":
		if err := decCmd.Parse(os.Args[2:]); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(2)
		}

		if len(decCmd.Args()) == 0 {
			fmt.Fprint(os.Stderr, "no arguments given. expected exactly one.")
			os.Exit(1)
		}

		input := decCmd.Args()[0]

		out, err := url.QueryUnescape(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "bad input: %v", err)
			os.Exit(123)
		}
		fmt.Fprintf(os.Stdout, "input:\t | '%s'\ndecoded: | '%s'\n", input, out)
	case "encode":
		if err := encCmd.Parse(os.Args[2:]); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(2)
		}

		if len(encCmd.Args()) == 0 {
			fmt.Fprint(os.Stderr, "no arguments given. expected exactly one.")
			os.Exit(1)
		}

		input := encCmd.Args()[0]
		fmt.Fprintf(os.Stdout, "input:\t | '%s'\nencoded: | '%s'\n", input, url.QueryEscape(input))
	default:
		fmt.Fprintf(os.Stderr, "Unknown command. 'decode' or 'encode' is expected")
		os.Exit(2)
	}
}
