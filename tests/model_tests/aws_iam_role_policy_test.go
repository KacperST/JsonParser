package models_test

import (
	"encoding/json"
	"testing"
	"parser.com/root/models"
)


func TestPositiveAWSIAMRolePolicy(t *testing.T) {
	aws_iam_role_policy := read_json("jsons/aws_iam_role_policy/proper_aws_iam_role_policy.json")
	var proper_aws_iam_role_policy models.AWSIAMRolePolicy
	err := json.Unmarshal([]byte(aws_iam_role_policy), &proper_aws_iam_role_policy)
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func TestNegativeWrongKeys(t *testing.T) {
	aws_iam_role_policy := read_json("jsons/aws_iam_role_policy/policy_wrong_keys.json")
	var wrong_keys models.AWSIAMRolePolicy
	err := json.Unmarshal([]byte(aws_iam_role_policy), &wrong_keys)
	if err == nil {
		t.Errorf("error: %v", err)
	}
}

func TestNegativePolicyNameNotString(t* testing.T) {
	aws_iam_role_policy := read_json("jsons/aws_iam_role_policy/policy_name_not_string.json")
	var policy_name_not_string models.AWSIAMRolePolicy
	err := json.Unmarshal([]byte(aws_iam_role_policy), &policy_name_not_string)
	if err == nil {
		t.Errorf("error: %v", err)
	}
}

func TestNegativePolicyNameNotMatchesRegex(t* testing.T) {
	aws_iam_role_policy := read_json("jsons/aws_iam_role_policy/policy_name_not_matches_regex.json")
	var policy_name_not_matches_regex models.AWSIAMRolePolicy
	err := json.Unmarshal([]byte(aws_iam_role_policy), &policy_name_not_matches_regex)
	if err == nil {
		t.Errorf("error: %v", err)
	}
}
func TestNegativePolicyNameTooLong(t *testing.T) {
	aws_iam_role_policy := read_json("jsons/aws_iam_role_policy/policy_name_too_long.json")
	var policy_name_too_long models.AWSIAMRolePolicy
	err := json.Unmarshal([]byte(aws_iam_role_policy), &policy_name_too_long)
	if err == nil {
		t.Errorf("error: %v", err)
	}
}

