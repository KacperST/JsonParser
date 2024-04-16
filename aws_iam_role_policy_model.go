package main

type AWSIAMRolePolicy struct {
	// The name of the policy.
	PolicyName string `json:"PolicyName"`
	PolicyDocument PolicyDocument `json:"PolicyDocument"`
}

type PolicyDocument struct {
	// The policy document.
	Version  string `json:"Version"`
	Statement []Statement `json:"Statement"`
}

type Statement struct {
	// Sid is an optional identifier for the statement.
	Sid *string `json:"Sid"`
	// The effect of the policy.
	Effect string `json:"Effect"`
	// The actions in the policy.
	Action interface{} `json:"Action"`
	// The resources in the policy.
	Resource interface{} `json:"Resource"`
}
