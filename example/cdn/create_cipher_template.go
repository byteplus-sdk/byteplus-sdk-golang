package cdn

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func CreateCipherTemplate(t *testing.T) {
	resp, err := DefaultInstance.CreateCipherTemplate(&cdn.CreateCipherTemplateRequest{
		Title:   "example",
		Message: cdn.GetStrPtr("example"),
		Quic: &cdn.Quic{
			Switch: cdn.GetBoolPtr(false),
		},
	})
	test, _ := json.Marshal(resp)
	fmt.Println(string(test))
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Result)
}
