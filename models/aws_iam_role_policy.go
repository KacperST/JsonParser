package models

import (
	"encoding/json"
	"parser.com/root/utils"
)

type AWSIAMRolePolicy struct {
	PolicyName     string         `json:"PolicyName"`
	PolicyDocument PolicyDocument `json:"PolicyDocument"`
}

// UnmarshalJSON is a method that will be called when you use json.Unmarshal for AWSIAMRolePolicy
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
	obj, err := utils.CheckJSONKeys(data, []string{"PolicyName", "PolicyDocument"}, []string{})
	if err != nil {
		return err
	}

	if err := utils.CheckPolicyName(obj["PolicyName"]); err != nil {
		return err
	}

	return nil
}
