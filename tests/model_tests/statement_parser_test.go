package models_test

import (
	"encoding/json"
	"os"
	"testing"
	"parser.com/root/models"
)

func read_json(filename string) []byte {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil
	}
	return fileBytes
}

func TestPositiveStatementParser(t *testing.T) {
	proper_statement := read_json("jsons/statements/proper_statement.json")
	var statement models.Statement
	err := json.Unmarshal(proper_statement, &statement)
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func TestPositiveStatementNoSid(t *testing.T) {
	statement_no_sid := read_json("jsons/statements/proper_statement_no_sid.json")
	var statement models.Statement
	err := json.Unmarshal(statement_no_sid, &statement)
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func TestPositiveStatementNoPrincipal(t *testing.T) {
	statement_no_principal := read_json("jsons/statements/proper_statement_no_principal.json")
	var statement models.Statement
	err := json.Unmarshal(statement_no_principal, &statement)
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func TestPositiveStatementNoCondition(t *testing.T) {
	statement_no_condition := read_json("jsons/statements/proper_statement_no_condition.json")
	var statement models.Statement
	err := json.Unmarshal(statement_no_condition, &statement)
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func TestNegativeStatementWithAdditionalKeys(t *testing.T) {
	statement_with_additional_field := read_json("jsons/statements/statement_with_additional_keys.json")
	var statement models.Statement
	err := json.Unmarshal(statement_with_additional_field, &statement)
	if err == nil {
		t.Errorf("error: %v", err)
	}
}

func TestNegativeStatementSidIsEmptyString(t *testing.T) {
	statement_sid_not_string := read_json("jsons/statements/statement_sid_empty_string.json")
	var statement models.Statement
	err := json.Unmarshal(statement_sid_not_string, &statement)
	if err == nil {
		t.Errorf("error: %v", err)
	}
}

func TestNegativeStatementEffectIsNotString(t *testing.T) {
	statement_effect_not_string := read_json("jsons/statements/statement_effect_not_string.json")
	var statement models.Statement
	err := json.Unmarshal(statement_effect_not_string, &statement)
	if err == nil {
		t.Errorf("error: %v", err)
	}
}

func TestNegativeStatementEffectIsNotAllowOrDeny(t *testing.T) {
	statement_effect_not_allow_or_deny := read_json("jsons/statements/statement_effect_not_allow_or_deny.json")
	var statement models.Statement
	err := json.Unmarshal(statement_effect_not_allow_or_deny, &statement)
	if err == nil {
		t.Errorf("error: %v", err)
	}
}

func TestNegativeStatementPrincipalValueIsNotListOrString(t *testing.T) {
	statement_principal_value_not_list := read_json("jsons/statements/statement_principal_value_not_list.json")
	var statement models.Statement
	err := json.Unmarshal(statement_principal_value_not_list, &statement)
	if err == nil {
		t.Errorf("error: %v", err)
	}
}
func TestNegativeStatementResouceIsAsterisk(t *testing.T) {
	statement_resource_asterisk := read_json("jsons/statements/statement_resource_asterisk.json")
	var statement models.Statement
	err := json.Unmarshal(statement_resource_asterisk, &statement)
	if err == nil {
		t.Errorf("error: %v", err)
	}
}

func TestNegativeStatementResourceValueIsNotStringOrList(t *testing.T) {
	statement_resource_not_string_or_list := read_json("jsons/statements/statement_resource_not_string_or_list.json")
	var statement models.Statement
	err := json.Unmarshal(statement_resource_not_string_or_list, &statement)
	if err == nil {
		t.Errorf("error: %v", err)
	}
}

func TestNegativeStatementConditionValueIsNotMap(t *testing.T) {
	statement_condition_value_not_map := read_json("jsons/statements/statement_condition_value_not_map.json")
	var statement models.Statement
	err := json.Unmarshal(statement_condition_value_not_map, &statement)
	if err == nil {
		t.Errorf("error: %v", err)
	}
}

func TestNegativeActionIsEmptyString(t *testing.T) {
	statement_action_empty_string := read_json("jsons/statements/statement_action_empty_string.json")
	var statement models.Statement
	err := json.Unmarshal(statement_action_empty_string, &statement)
	if err == nil {
		t.Errorf("error: %v", err)
	}
}