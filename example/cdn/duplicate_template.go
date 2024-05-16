package cdn

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func DuplicateTemplate(t *testing.T) {
	resp, err := DefaultInstance.DuplicateTemplate(&cdn.DuplicateTemplateRequest{
		Title:              cdn.GetStrPtr("example"),
		Message:            cdn.GetStrPtr("example"),
		ReferredTemplateId: "tpl-example",
	})
	test, _ := json.Marshal(resp)
	fmt.Println(string(test))
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Result)
}
