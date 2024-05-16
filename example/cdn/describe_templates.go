package cdn

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func DescribeTemplates(t *testing.T) {
	resp, err := DefaultInstance.DescribeTemplates(&cdn.DescribeTemplatesRequest{
		PageNum:  cdn.GetInt64Ptr(1),
		PageSize: cdn.GetInt64Ptr(10),
		Filters: []cdn.Filter{
			{
				Fuzzy: cdn.GetBoolPtr(false),
				Name:  cdn.GetStrPtr("type"),
				Value: []string{"service"},
			},
		},
	})
	test, _ := json.Marshal(resp)
	fmt.Println(string(test))
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Result)
}
