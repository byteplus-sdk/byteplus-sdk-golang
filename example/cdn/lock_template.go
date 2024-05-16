package cdn

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func LockTemplate(t *testing.T) {
	resp, err := DefaultInstance.LockTemplate(&cdn.LockTemplateRequest{
		TemplateId: "tpl-example",
	})
	test, _ := json.Marshal(resp)
	fmt.Println(string(test))
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
}
