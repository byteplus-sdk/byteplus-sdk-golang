package cdn

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func DescribeTemplateDomains(t *testing.T) {
	resp, err := DefaultInstance.DescribeTemplateDomains(&cdn.DescribeTemplateDomainsRequest{
		PageNum:  cdn.GetInt64Ptr(1),
		PageSize: cdn.GetInt64Ptr(10),
		Filters:  []cdn.Filter{},
	})
	test, _ := json.Marshal(resp)
	fmt.Println(string(test))
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Result.Domains)
}
