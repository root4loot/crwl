package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gookit/color"
	crwl "github.com/root4loot/crwl/pkg"
)

var version string

func parse() *crwl.Options {
	version = "v0.1.0"

	options := crwl.Options{
		Domain:    flag.String("domain", "", ""),
		Ext:       flag.String("ext", "", ""),
		Parallels: flag.Int("para", 10, ""),
		Depth:     flag.Int("depth", 999, ""),
		Delay:     flag.Int("delay", 0, ""),
		Delay2:    flag.Int("delay2", 0, ""),
		Whitelist: flag.String("whitelist", "", ""),
		UserAgent: flag.String("useragent", "", ""),
		Outfile:   flag.String("outfile", "", ""),
		Regexp:    flag.String("regex", "", ""),
		Async:     flag.Bool("async", false, ""),
		JSON:      flag.Bool("json", false, ""),
		Silent:    flag.Bool("silent", false, ""),
		Version:   flag.Bool("version", false, ""),
		Help:      flag.Bool("help", false, ""),
	}

	flag.Usage = func() {
		banner()
		usagex()
	}

	flag.Parse()
	// remove timestamp from exits
	log.SetFlags(0)

	if *options.Help {
		banner()
		usagex()
		os.Exit(0)
	}

	// print version
	if *options.Version {
		fmt.Println("crwl " + version)
		os.Exit(0)
	}

	if *options.Domain == "" {
		log.Fatalf("%s %s\n", color.FgRed.Text("[!]"), "Missing target domain (-domain)")
	}

	if *options.Silent && *options.Outfile == "" {
		log.Fatalf("%s %s\n", color.FgRed.Text("[!]"), "Please specify an -outfile when using --silent")
	}

	if *options.JSON && *options.Outfile == "" {
		log.Fatalf("%s %s\n", color.FgRed.Text("[!]"), "Missing -outfile for JSON write")
	}

	if *options.Ext != "" && *options.Regexp != "" {
		log.Fatalf("%s %s\n", color.FgRed.Text("[!]"), "-regex cannot be used with -ext")
	}

	if *options.Outfile != "" {
		// Attempt to create outfile
		_, err := os.Create(*options.Outfile)
		if err != nil {
			log.Fatalf("\n%s Could not write file %s. Bad permissions?", color.FgRed.Text("[!]"), *options.Outfile)
		}
	}

	return &options
}

func usagex() {
	fmt.Printf("\n%s\n\n", "Arguments:")
	fmt.Printf("    %s\t%s    %s\n", "-domain", "<string>", "Domain to crawl")
	fmt.Printf("    %s\t%s    %s\n", "-whitelist", "<string>", "Domains to be whitelisted (comma separated)")
	fmt.Printf("    %s\t%s    %s\n", "-ext", "<string>", "Filter results on URL file extensions (comma separated)")
	fmt.Printf("    %s\t%s    %s\n", "-outfile", "<string>", "File to write results")
	fmt.Printf("    %s\t%s    %s\n", "-regex", "<string>", "Regular expression to match against domains")
	fmt.Printf("    %s\t%s    %s\n", "-useragent", "<string>", "Set user-agent")
	fmt.Printf("    %s\t%s       %s\n", "-para", "<int>", "Max parallelism (Default: 10)")
	fmt.Printf("    %s\t%s       %s\n", "-depth", "<int>", "Maximum depth to crawl (Default: 999)")
	fmt.Printf("    %s\t%s       %s\n", "-delay", "<int>", "Seconds to wait before creating a new request to the matching domains")
	fmt.Printf("    %s\t%s       %s\n", "-delay2", "<int>", "Seconds to be randomized prior to each new request (Default: 0)")
	fmt.Printf("    %s\t\t%s    %s\n", "-json", "", "Output as JSON (-outfile)")
	fmt.Printf("    %s\t\t%s    %s\n", "-silent", "", "Suppress output from console")
	fmt.Printf("    %s\t\t%s    %s\n", "-async", "", "Enable asynchronous network requests (Default: ON)")
	fmt.Printf("    %s\t\t%s    %s\n", "-version", "", "Display version")
	fmt.Printf("    %s\t\t%s    %s\n", "-help", "", "Display help")
	fmt.Println("")
}

func banner() {
	banner := `
                              888 
 e88~~\ 888-~\ Y88b    e    / 888 
d888    888     Y88b  d8b  /  888 
8888    888      Y888/Y88b/   888 
Y888    888       Y8/  Y8/    888 
 "88__/ 888        Y    Y     888 
`

	fmt.Printf("%s\n", banner)
}
