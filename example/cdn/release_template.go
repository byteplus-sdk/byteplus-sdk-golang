package cdn

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func ReleaseTemplate(t *testing.T) {
	resp, err := DefaultInstance.ReleaseTemplate(&cdn.ReleaseTemplateRequest{
		TemplateId:      "tpl-xxxxxxx",
		Message:         cdn.GetStrPtr("test"),
		Env:             cdn.GetStrPtr("prod"),
		TemplateVersion: cdn.GetStrPtr("draft"),
	})
	test, _ := json.Marshal(resp)
	fmt.Println(string(test))
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
}
