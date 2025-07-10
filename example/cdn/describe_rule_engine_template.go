package cdn

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func DescribeRuleEngineTemplate(t *testing.T) {
	resp, err := DefaultInstance.DescribeRuleEngineTemplate(&cdn.DescribeRuleEngineTemplateRequest{
		TemplateId:      "tpl-xxxxxxx",
		TemplateVersion: cdn.GetStrPtr("draft"),
	})
	test, _ := json.Marshal(resp)
	fmt.Println(string(test))
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Result)
}
