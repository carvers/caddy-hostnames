package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mholt/caddy/caddyfile"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./caddy-hostnames [path to Caddyfile]")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error opening Caddyfile %q: %s\n", os.Args[1], err)
		os.Exit(1)
	}
	defer f.Close()
	serverBlocks, err := caddyfile.Parse("/usr/local/etc/caddy/Caddyfile", f, nil)
	if err != nil {
		fmt.Printf("Error parsing Caddyfile %q: %s\n", os.Args[1], err)
		f.Close()
		os.Exit(1)
	}
	for _, block := range serverBlocks {
		for _, key := range block.Keys {
			if strings.Contains(key, ":") {
				continue
			}
			fmt.Println(key)
		}
	}
}
