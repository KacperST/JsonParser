package models_test

import (
	"encoding/json"
	"testing"
	"parser.com/root/models"
)

func TestPositivePolicyDocument(t *testing.T) {
	proper_policy_document := read_json("jsons/policy_documents/proper_policy_document.json")
	var policy_document models.PolicyDocument
	err := json.Unmarshal(proper_policy_document, &policy_document)
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func TestNegattiveAdditionalsKeys(t *testing.T) {
	policy_document_with_additional_keys := read_json("jsons/policy_documents/policy_document_with_additional_keys.json")
	var policy_document models.PolicyDocument
	err := json.Unmarshal(policy_document_with_additional_keys, &policy_document)
	if err == nil {
		t.Errorf("error: %v", err)
	}
}
func TestNegativeLackingKeys(t *testing.T) {
	policy_document_lacking_keys := read_json("jsons/policy_documents/policy_document_lacking_keys.json")
	var policy_document models.PolicyDocument
	err := json.Unmarshal(policy_document_lacking_keys, &policy_document)
	if err == nil {
		t.Errorf("error: %v", err)
	}
}
func TestNegativeVersionIsEmptyString(t *testing.T) {
	policy_document_version_not_sting := read_json("jsons/policy_documents/policy_document_version_empty_string.json")
	var policy_document models.PolicyDocument
	err := json.Unmarshal(policy_document_version_not_sting, &policy_document)
	if err == nil {
		t.Errorf("error: %v", err)
	}
}
func TestNegativeVersionIsNotAString(t *testing.T) {
	policy_document_version_not_string := read_json("jsons/policy_documents/policy_document_version_not_string.json")
	var policy_document models.PolicyDocument
	err := json.Unmarshal(policy_document_version_not_string, &policy_document)
	if err == nil {
		t.Errorf("error: %v", err)
	}
}
