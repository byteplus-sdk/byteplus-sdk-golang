package cdn

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func UpdateTemplateDomain(t *testing.T) {
	resp, err := DefaultInstance.UpdateTemplateDomain(&cdn.UpdateTemplateDomainRequest{
		ServiceRegion:     cdn.GetStrPtr("global"),
		ServiceTemplateId: cdn.GetStrPtr("tpl-example"),
		CertId:            cdn.GetStrPtr("cert-example"),
		CipherTemplateId:  cdn.GetStrPtr("tpl-example"),
		Domains:           []string{"example.com"},
		HTTPSSwitch:       cdn.GetStrPtr("on"),
	})
	test, _ := json.Marshal(resp)
	fmt.Println(string(test))
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
}
