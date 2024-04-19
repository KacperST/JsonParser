package models

import (
	"encoding/json"
	"parser.com/root/utils"
)

type PolicyDocument struct {
	Version   string      `json:"Version"`
	Statement []Statement `json:"Statement"`
}

// UnmarshalJSON is a method that will be called when you use json.Unmarshal for PolicyDocument
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
	obj, err := utils.CheckJSONKeys(data, []string{"Version", "Statement"}, []string{})
	if err != nil {
		return err
	}
	if err := utils.CheckString(obj["Version"]); err != nil {
		return err
	}

	return nil
}
