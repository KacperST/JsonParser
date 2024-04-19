package main

import (
	"encoding/json"
	"fmt"
	"os"
	"parser.com/root/models"
	"strconv"
	"strings"
)

// ReadJson reads the json file
// Parameters: filename string - the name of the file
// Returns: []byte, error - the data as bytes and an error if any

func ReadJson(filename string) ([]byte, error) {
	// Read the json file
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return fileBytes, nil
}

// ParseJson parses the json file
// Parameters: fileBytes []byte - the data as bytes
// Returns: bool, error - a boolean value and an error if any
func ParseJson(fileBytes []byte) (output bool, err error) {
	var awsIAMRolePolicy models.AWSIAMRolePolicy
	err = json.Unmarshal(fileBytes, &awsIAMRolePolicy)
	if err != nil {
		return false, err
	}
	return true, nil
}

// CheckJSONFormat checks if the json file is in the proper format
// Parameters: path string - the path of the file
// Returns: bool, error - a boolean value and an error if any
func CheckJSONFormat(path string) (bool, error) {
	fileBytes, err := ReadJson(path)
	if err != nil {
		return false, err
	}
	output, err := ParseJson(fileBytes)

	if err != nil {
		return false, err
	}
	return output, nil
}

func main() {
	isProperFormat, error := CheckJSONFormat("inputs/aws_iam_role_policy.json")
	if error != nil {
		fmt.Printf("%s -  The file is not in the proper format\nError: %v\n", strings.ToUpper(strconv.FormatBool(isProperFormat)), error)
	} else {
		fmt.Printf("%s -  The file is in the proper format.\n", strings.ToUpper(strconv.FormatBool(isProperFormat)))
	}
}
