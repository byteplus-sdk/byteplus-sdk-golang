package cdn

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func DescribeCipherTemplate(t *testing.T) {
	resp, err := DefaultInstance.DescribeCipherTemplate(&cdn.DescribeCipherTemplateRequest{
		TemplateId: "tpl-example",
	})
	test, _ := json.Marshal(resp)
	fmt.Println(string(test))
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Result)
}
