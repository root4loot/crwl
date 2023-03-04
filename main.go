package main

import (
	"encoding/json"
	"log"
	"os"

	crwl "github.com/root4loot/crwl/pkg"
)

var (
	file *os.File
)

func main() {
	args := parse()
	results := crwl.Go(*args)

	if *args.Outfile != "" {
		file, err := os.OpenFile(*args.Outfile,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		if *args.JSON {
			r, _ := json.Marshal(results)
			file.Write(r)
		} else {
			for _, res := range results {
				file.WriteString(res + "\n")
			}
		}
	}

	defer file.Close()

}
