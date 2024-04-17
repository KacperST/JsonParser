package main

import (
	"encoding/json"
	"fmt"
	"os"
)
func read_json(filename string) ([]byte, error) {
	// Read the json file
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return fileBytes, nil
}

func parse_json(fileBytes []byte) (output bool, err error) {
	var awsIAMRolePolicy AWSIAMRolePolicy
	err = json.Unmarshal(fileBytes, &awsIAMRolePolicy)
	if err != nil {
		return false, err
	}
	fmt.Print(awsIAMRolePolicy)
	return true, nil
}


func main() {
	fileBytes, err := read_json("inputs/aws_iam_role_policy.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	output, err := parse_json(fileBytes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(output)
}
