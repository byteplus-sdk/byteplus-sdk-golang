package cdn

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func CreateRuleEngineTemplate(t *testing.T) {
	resp, err := DefaultInstance.CreateRuleEngineTemplate(&cdn.CreateRuleEngineTemplateRequest{
		Title:   "example",
		Message: cdn.GetStrPtr("test"),
		Project: cdn.GetStrPtr("default"),
		Rule: cdn.GetStrPtr(`{"IfBlock": {
			"SubRules": [],
			"Actions": [{
				"Action": "request_header",
				"Stage": "client_request",
				"Phase": 1,
				"Module": 1,
				"Groups": [{
					"Dimension": "request_header",
					"GroupParameters": [{
						"Parameters": [{
							"Name": "action",
							"Values": ["set"]
						}, {
							"Name": "header_name",
							"Values": ["header123"]
						}, {
							"Name": "header_value",
							"Values": ["value456"]
						}]
					}]
				}]
			}],
			"Condition": {
				"IsGroup": true,
				"Connective": "and",
				"ConditionGroups": [{
					"IsGroup": true,
					"Connective": "and",
					"ConditionGroups": [{
						"IsGroup": false,
						"Condition": {
							"Object": "http_referer",
							"Operator": "equal",
							"Value": ["header123"],
							"IgnoreCase": true
						}
					}]
				}]
			}
		}
	}`),
	})
	test, _ := json.Marshal(resp)
	fmt.Println(string(test))
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Result)
}
