package utilstests_test

import (
	"parser.com/root/utils"
	"strings"
	"testing"
)

func TestPositiveContainsKey(t *testing.T) {
	keys := []string{"key1", "key2", "key3"}
	if !utils.ContainsKey("key1", keys) {
		t.Errorf("error: Array %v should contain key %v", keys, "key1")
	}
}

func TestNegativeContainsKey(t *testing.T) {
	keys := []string{"key1", "key2", "key3"}
	if utils.ContainsKey("key4", keys) {
		t.Errorf("error: Array %v should not contain key %v", keys, "key4")
	}
}

func TestPositiveCheckJSONKeys(t *testing.T) {
	keys := []string{"key1", "key2", "key3"}
	data := []byte(`{"key1": "value1", "key2": "value2", "key3": "value3"}`)
	_, err := utils.CheckJSONKeys(data, keys, []string{})
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func TestNegativeCheckJSONKeys(t *testing.T) {
	keys := []string{"key1", "key2", "key3"}
	data := []byte(`{"key1": "value1", "key2": "value2", "key4": "value4"}`)
	_, err := utils.CheckJSONKeys(data, keys, []string{})
	if err == nil {
		t.Errorf("error: %v", err)
	}
}

func TestNegativeCheckJSONKeysWrongData(t *testing.T) {
	keys := []string{"key1", "key2", "key3"}
	data := []byte(`{1,2,3,4}`)
	_, err := utils.CheckJSONKeys(data, keys, []string{})
	if err == nil {
		t.Errorf("error: %v", err)
	}
}
	

func TestNegativeCheckJsonKeysMissingKeys(t *testing.T) {
	keys := []string{"key1", "key2", "key3"}
	data := []byte(`{"key1": "value1", "key2": "value2"}`)
	_, err := utils.CheckJSONKeys(data, keys, []string{})
	if err == nil {
		t.Errorf("error: %v", err)
	}
}

func TestPositiveCheckString(t *testing.T) {
	err := utils.CheckString("string")
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func TestNegativeCheckString(t *testing.T) {
	err := utils.CheckString(1)
	if err == nil {
		t.Errorf("error: %v", err)
	}
}
func TestNegativeCheckStringEmptyString(t *testing.T) {
	err := utils.CheckString("")
	if err == nil {
		t.Errorf("error: %v", err)
	}
}

func TestPositiveCheckListOrString(t *testing.T) {
	err := utils.CheckListOrString([]interface{}{"string1", "string2"})
	if err != nil {
		t.Errorf("error: %v", err)
	}
	err2 := utils.CheckListOrString("string")
	if err2 != nil {
		t.Errorf("error: %v", err2)
	}
}

func TestNegativeCheckListOrString(t *testing.T) {
	err := utils.CheckListOrString([]interface{}{1, 2})
	if err == nil {
		t.Errorf("error: %v", err)
	}
}

func TestNegativeCheckListOrStringEmptyList(t *testing.T) {
	err := utils.CheckListOrString([]interface{}{})
	if err == nil {
		t.Errorf("error: %v", err)
	}
}

func TestPositiveCheckPolicyName(t *testing.T) {
	err := utils.CheckPolicyName("policy_name")
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func TestNegativeCheckPolicyName(t *testing.T) {
	err := utils.CheckPolicyName(1)
	if err == nil {
		t.Errorf("error: %v", err)
	}
}

func TestNegativeCheckPolicyNameWrongRegex(t *testing.T) {
	err := utils.CheckPolicyName("!{}><?{})")
	if err == nil {
		t.Errorf("error: %v", err)
	}
}

func TestNegativeCheckPolicyNameTooLong(t *testing.T) {
	err := utils.CheckPolicyName(strings.Repeat("a", 129))
	if err == nil {
		t.Errorf("error: %v", err)
	}
}

func TestPositiveCheckEffect(t *testing.T) {
	err := utils.CheckEffect("Allow")
	if err != nil {
		t.Errorf("error: %v", err)
	}
	err2 := utils.CheckEffect("Deny")
	if err2 != nil {
		t.Errorf("error: %v", err2)
	}
}

func TestNegativeCheckEffect(t *testing.T) {
	err := utils.CheckEffect(1)
	if err == nil {
		t.Errorf("error: %v", err)
	}
}

func TestNegativeCheckEffectWrongStringValue(t *testing.T) {
	err := utils.CheckEffect("Wrong Value")
	if err == nil {
		t.Errorf("error: %v", err)
	}
}

func TestPositiveCheckPrincipal(t *testing.T) {
	prinicpal := map[string]interface{}{"key1": []interface{}{"value1", "value2"}}
	err := utils.CheckPrincipal(prinicpal)
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func TestNegativeCheckPrincipal(t *testing.T) {
	err := utils.CheckPrincipal(1)
	if err == nil {
		t.Errorf("error: %v", err)
	}
}
func TestNegativeCheckPrincipalWrongType(t *testing.T) {
	err := utils.CheckPrincipal(map[string]interface{}{"key1": 1})
	if err == nil {
		t.Errorf("error: %v", err)
	}
}

func TestPositiveCheckResource(t *testing.T) {
	err := utils.CheckResource("resource")
	if err != nil {
		t.Errorf("error: %v", err)
	}
	err2 := utils.CheckResource([]interface{}{"resource1", "resource2"})
	if err2 != nil {
		t.Errorf("error: %v", err2)
	}
}

func TestNegativeCheckResource(t *testing.T) {
	err := utils.CheckResource(1)
	if err == nil {
		t.Errorf("error: %v", err)
	}
	err2 := utils.CheckResource([]interface{}{1, 2})
	if err2 == nil {
		t.Errorf("error: %v", err2)
	}

}
func TestNegativeCheckResourceAsterisk(t *testing.T) {
	err := utils.CheckResource("*")
	if err == nil {
		t.Errorf("error: %v", err)
	}
}

func TestPositiveCheckCondition(t *testing.T) {
	condition := map[string]interface{}{"key1": map[string]interface{}{"key2": "value"}}
	err := utils.CheckCondition(condition)
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func TestNegativeCheckCondition(t *testing.T) {
	err := utils.CheckCondition(1)
	if err == nil {
		t.Errorf("error: %v", err)
	}
	err2 := utils.CheckCondition(map[string]interface{}{"key1": "value1"})
	if err2 == nil {
		t.Errorf("error: %v", err2)
	}
}
