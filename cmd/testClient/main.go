package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/icecrime/docker-api/api"
	"github.com/icecrime/docker-api/client"
)

func printResult(v interface{}, err error) {
	if err == nil {
		if b, err := json.MarshalIndent(v, " ", " "); err == nil {
			fmt.Printf("%s\n", string(b))
		} else {
			fmt.Printf("%v\n", v)
		}
	} else {
		fmt.Printf("[ERROR] %v\n", err)
	}
}

func main() {
	h := &http.Client{}
	c := client.New(h, "http://localhost:8080")

	testConfig := map[string]interface{}{
		"Image":    "busybox",
		"Hostname": "PoCHostname",
	}

	fmt.Printf("> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		switch strings.ToLower(strings.TrimSpace(scanner.Text())) {
		case "list":
			printResult(c.List(&api.ListContainersParams{}))
		case "ping":
			printResult(c.Ping())
		case "version":
			printResult(c.Version())
		case "create":
			printResult(c.Create(testConfig))
		case "start":
			printResult(c.Start("testContainerID"))
		default:
			fmt.Printf("Unknown command %q\n", scanner.Text())
		}
		fmt.Printf("> ")
	}
}
