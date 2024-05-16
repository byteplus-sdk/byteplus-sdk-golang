package cdn

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func AddTemplateDomain(t *testing.T) {
	resp, err := DefaultInstance.AddTemplateDomain(&cdn.AddTemplateDomainRequest{
		ServiceRegion:     cdn.GetStrPtr("global"),
		ServiceTemplateId: "tpl-example",
		CertId:            cdn.GetStrPtr("cert-example"),
		CipherTemplateId:  cdn.GetStrPtr("tpl-example"),
		Domain:            "example.com",
		HTTPSSwitch:       cdn.GetStrPtr("on"),
	})
	test, _ := json.Marshal(resp)
	fmt.Println(string(test))
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
}
