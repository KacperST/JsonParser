package models

import (
	"encoding/json"
	"parser.com/root/utils"
)

type Statement struct {
	Sid       *string                         `json:"Sid"`
	Effect    string                         `json:"Effect"`
	Principal *map[string][]string           `json:"Principal"`
	Action    interface{}                    `json:"Action"`
	Resource  interface{}                    `json:"Resource"`
	Condition *map[string]map[string]string  `json:"Condition"`
}

// UnmarshalJSON is a method that will be called when you use json.Unmarshal for Statement
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
	obj, err := utils.CheckJSONKeys(data, []string{"Effect", "Action", "Resource"}, []string{"Sid", "Principal", "Condition"})
	if err != nil {
		return err
	}
	if obj["Sid"] != nil {
		if err := utils.CheckString(obj["Sid"]); err != nil {
			return err
		}
	}
	if err := utils.CheckEffect(obj["Effect"]); err != nil {
		return err
	}
	if obj["Principal"] != nil {
		if err := utils.CheckPrincipal(obj["Principal"]); err != nil {
			return err
		}
	}
	if err := utils.CheckListOrString(obj["Action"]); err != nil {
		return err
	}

	if err := utils.CheckResource(obj["Resource"]); err != nil {
		return err
	}
	if obj["Condition"] != nil {
		if err := utils.CheckCondition(obj["Condition"]); err != nil {
			return err
		}
	}


	return nil
}
