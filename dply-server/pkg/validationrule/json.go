package validationrule

import (
	"encoding/json"
	"errors"

	"gopkg.in/validator.v2"
)

func RegisterJsonRule() {
	validator.SetValidationFunc("json", ruleJson)
}

func isJsonString(jsonStr string) bool {
	tmpJson := map[string]interface{}{}
	return json.Unmarshal([]byte(jsonStr), &tmpJson) == nil
}

func ruleJson(in interface{}, param string) (err error) {
	inStr := ""
	if v, ok := in.(string); ok {
		inStr = v
	} else {
		return errors.New("rule `json` can only be used by string/[]byte data type")
	}

	if !isJsonString(inStr) {
		return errors.New("invalid format. rule `json` is required")
	}

	return nil
}
