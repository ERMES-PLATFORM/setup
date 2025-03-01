package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/ermes-labs/api-go/infrastructure"
	log "github.com/my-ermes-labs/log"
)

func main() {
	// Check if the command-line argument is provided.
	if len(os.Args) < 2 {
		fmt.Println("Usage: deploy <input_file>")
		return
	}

	// The first command-line argument is the program path, the second is the
	// desired input file.
	filePath := os.Args[1]
	// Read JSON file.
	file, err := os.ReadFile(filePath)
	// Check if there was an error reading the file.
	if err != nil {
		fmt.Printf("Error reading JSON file: %s\n", err)
		return
	}

	// Unmarshal JSON file.
	infra, _, err := infrastructure.UnmarshalInfrastructure(file)
	// Check if there was an error unmarshaling the JSON file.
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %s\n", err)
		return
	}

	// Get all areas.
	areas := infra.Flatten()

	for _, area := range areas {

		// Marshal JSON node.
		jsonNode, err := infrastructure.MarshalNode(area.Node)
		if err != nil {
			fmt.Printf("Error marshaling JSON: %s\n", err)
			return
		}
		encodedJson := base64.StdEncoding.EncodeToString([]byte(jsonNode))

		areasString, err := json.Marshal(area.Areas)
		if err != nil {
			fmt.Printf("\nError marshaling areas: %s\n", err)
			return
		}
		if string(areasString) == "null" {
			areasString = []byte("[]")
			log.GeneralLog("\nEMPTY String Node Areas:")
		}
		log.GeneralLog("\nareasString62: " + string(areasString))

		encodedAreas := base64.StdEncoding.EncodeToString(areasString)

		cmd := exec.Command("ansible-playbook", "-i", "inventory.ini", "deploy.yml", "--extra-vars",
			fmt.Sprintf("target_node='%s' target_hosts='%s' target_areas='%s'", encodedJson, area.Host, encodedAreas))

		// Set environment variables and execute
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			fmt.Printf("Error running Ansible playbook: %s\n", err)
			return
		}
	}
}
