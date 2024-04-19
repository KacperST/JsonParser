package utils

import (
	"encoding/json"
	"fmt"
	"regexp"
)

// ContainsKey checks if a key is present in a list of keys (for example in a keys of map)
// Parameters:
// - key: key to check
// - keys: list (map) of keys to check
// Returns:
// - bool: true if the key is present in the list, false otherwise
func ContainsKey(key string, keys []string) bool {
	for _, val := range keys {
		if val == key {
			return true
		}
	}
	return false
}

// CheckJSONKeys checks if the keys in a JSON object are present in a list of keys
// and if there are any unexpected or missing keys.
// Parameters:
// - data: JSON object
// - keys: list of keys that has to be present in the JSON object
// - optionalKeys: list of keys that are optional in the JSON object
// Returns:
// - mapped object: JSON object mapped to a map
// - error: error if there are any unexpected or missing keys
func CheckJSONKeys(data []byte, keys []string, optionalKeys []string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	keys_copy := make([]string, len(keys))
	copy(keys_copy, keys)
	keys_copy_set := make(map[string]bool)
	for _, key := range keys {
		keys_copy_set[key] = true
	}
	optional_keys_set := make(map[string]bool)
	for _, key := range optionalKeys {
		optional_keys_set[key] = true
	}
	for key := range m {
		if !ContainsKey(key, keys) && !ContainsKey(key, optionalKeys) {
			return nil, fmt.Errorf("unexpected key \"%s\" in PolicyDocument field", key)
		}
		delete(keys_copy_set, key)
	}
	for key := range keys_copy_set {
		if !optional_keys_set[key] {
			return nil, fmt.Errorf("missing key \"%s\" in PolicyDocument field", key)
		}
	}
	return m, nil
}

// CheckString checks if a value is a string and if it is not empty.
// Parameters:
// - str: value to check
// Returns:
// - error: nill or error if occures
func CheckString(str interface{}) error {
	if _, isString := str.(string); !isString {
		return fmt.Errorf("value %v is not a string", str)
	}
	if str == "" {
		return fmt.Errorf("string is empty")
	}
	return nil
}

// CheckListOrString checks if a value is a list of strings or a string.
// Some field in AWS::IAM::Role Policy might be a list of strings or a single string so this function checks both.
// Parameters:
// - list: value to check
// Returns:
// - error: nill or error if occures
func CheckListOrString(val interface{}) error {
	isListOrString := false
	if obj, isList := val.([]interface{}); isList {
		if len(obj) == 0 {
			return fmt.Errorf("list is empty")
		}
		for _, value := range obj {
			err := CheckString(value)
			if err != nil {
				return fmt.Errorf("list contains a value that is not a string")
			}
			
		}
		isListOrString = true
	}
	err := CheckString(val)
	if err == nil {
		isListOrString = true
	}
	
	if !isListOrString {
		return fmt.Errorf("value is not a list or string")
	}

	return nil
}

// CheckPolicyName checks if a value is a non-empty string that contains only valid characters and is not longer than 128 characters.
// Parameters:
// - policyName: value to check
// Returns:
// - error: nill or error if occures

func CheckPolicyName(policyName interface{}) error {
	if err := CheckString(policyName); err != nil {
		return err
	}
	strPolicyName := policyName.(string)
	if len(strPolicyName) > 128 {
		return fmt.Errorf("PolicyName is longer than 128 characters")
	}
	if match, _ := regexp.MatchString("[A-Za-z0-9_+=,.@-]+", strPolicyName); !match {
		return fmt.Errorf("PolicyName contains invalid characters")
	}
	return nil
}

// CheckEffect checks if a value is a string and if it is "Allow" or "Deny".
// Parameters:
// - effect: value to check
// Returns:
// - error: nill or error if occures
func CheckEffect(effect interface{}) error {
	if _, isString := effect.(string); !isString {
		return fmt.Errorf("effect is not a string")
	}
	strEffect := effect.(string)
	if strEffect != "Allow" && strEffect != "Deny" {
		return fmt.Errorf("effect is not \"Allow\" or \"Deny\"")
	}
	return nil
}

// CheckPrincipal checks if a value is a map of strings or lists of strings.
// Parameters:
// - principal: value to check
// Returns:
// - error: nill or error if occures
func CheckPrincipal(principal interface{}) error {
	if _, isMap := principal.(map[string]interface{}); !isMap {
		return fmt.Errorf("principal is not a map")
	}
	for key, value := range principal.(map[string]interface{}) {
		err := CheckListOrString(value)
		if err != nil {
			return fmt.Errorf("principal value for key \"%s\" is not a list or a string", key)
		}
	}

	return nil
}

// CheckResource checks if a value is a string or a list of strings and if it is not "*".
// Parameters:
// - resource: value to check
// Returns:
// - error: error if type is invalid or if the value is "*"
func CheckResource(resource interface{}) error {
	switch resources := resource.(type) {
	case string:
		if resources == "*" {
			return fmt.Errorf("resource is set to \"*\"")
		}
	case []interface{}:
		if err := CheckListOrString(resources); err != nil {
			return err
		}
	default:
		return fmt.Errorf("resource is not a string or list of strings")
	}
	return nil
}

// CheckCondition checks if a value is a map of maps of strings.
// Parameters:
// - condition: value to check
// Returns:
// - error: nill or error if occures
func CheckCondition(condition interface{}) error {
	if _, isMap := condition.(map[string]interface{}); !isMap {
		return fmt.Errorf("condition is not a map")
	}
	for key, value := range condition.(map[string]interface{}) {
		if _, isMap := value.(map[string]interface{}); !isMap {
			return fmt.Errorf("condition value for key \"%s\" is not a map", key)
		}
		for key2, value2 := range value.(map[string]interface{}) {
			err := CheckString(value2)
			if err != nil {
				return fmt.Errorf("condition value for key \"%s\" is not a string", key2)
			}
		}
	}
	return nil
}
