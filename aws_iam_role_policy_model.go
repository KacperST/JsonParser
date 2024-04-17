package main

import (
	"encoding/json"
	"fmt"
)

type AWSIAMRolePolicy struct {
	PolicyName string `json:"PolicyName"`
	PolicyDocument PolicyDocument `json:"PolicyDocument"`
}

type PolicyDocument struct {
	Version  string `json:"Version"`
	Statement []Statement `json:"Statement"`
}

type Statement struct {
	Sid *string `json:"Sid"`
	Effect string `json:"Effect"`
	Action interface{} `json:"Action"`
	Resource interface{} `json:"Resource"`
}

func containsKey(key string, keys []string) bool {
    for _, val := range keys {
        if val == key {
            return true
        }
    }
    return false
}

func UnmarshalJSONInterface(data []byte, keys []string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	for key := range m {
		if !containsKey(key, keys) {
			return nil, fmt.Errorf("Unexpected key \"%s\" in PolicyDocument field", key)
		}
	}
	return m, nil
}



func (a *AWSIAMRolePolicy) UnmarshalJSON(data []byte) error {
	type Alias AWSIAMRolePolicy
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if _, err := UnmarshalJSONInterface(data, []string{"PolicyName", "PolicyDocument"}); err != nil {
		return err
	}

	
	return nil
}

func (p *PolicyDocument) UnmarshalJSON(data []byte) error {
	type Alias PolicyDocument
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(p),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if _, err := UnmarshalJSONInterface(data, []string{"Version", "Statement"}); err != nil {
		return err
	}

	return nil
}

func (s *Statement) UnmarshalJSON(data []byte) error {
	type Alias Statement
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	object, err := UnmarshalJSONInterface(data, []string{"Sid", "Effect", "Action", "Resource"})
	if err != nil {
		return err
	}
	switch resources := object["Resource"].(type) {
		case string:
			if resources == "*" {
				return fmt.Errorf("Resource is set to \"*\"")
			}
		case []string:
			if len(resources) == 0 {
				return fmt.Errorf("Resource is an empty list")
			}
		default:
			return fmt.Errorf("Resource is not a string or list of strings")
	}

	return nil
}
