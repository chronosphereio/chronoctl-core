package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run your_program.go <file_path>")
		os.Exit(1)
	}

	filePath := os.Args[1]

	// Option 1: Read the entire file into memory (Good for smaller files)
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file using ioutil.ReadFile: %v", err)
		return
	}

	output, _, err := ConvertPrometheus(string(content), "", Opts{})

	if err != nil {
		fmt.Println(err)
		return
	}
	for _, m := range output.monitors {
		m.NotificationPolicySlug = "plaid-notification-policy"
		m.IntervalSecs = 60
		outYaml, err := json.Marshal(*m)
		if err != nil {
			fmt.Println("err", err)
			return
		}
		fmt.Println(string(outYaml))
	}

}
