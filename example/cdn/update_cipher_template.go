package cdn

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func UpdateCipherTemplate(t *testing.T) {
	resp, err := DefaultInstance.UpdateCipherTemplate(&cdn.UpdateCipherTemplateRequest{
		Title:   cdn.GetStrPtr("example"),
		Message: cdn.GetStrPtr("example"),
		Quic: &cdn.Quic{
			Switch: cdn.GetBoolPtr(false),
		},
		HTTPS: &cdn.HTTPSCommon{
			HTTP2:      cdn.GetBoolPtr(true),
			OCSP:       cdn.GetBoolPtr(false),
			TlsVersion: []string{"tlsv1.1", "tlsv1.2"},
			Hsts: &cdn.Hsts{
				Switch: cdn.GetBoolPtr(false),
				Ttl:    cdn.GetInt64Ptr(0),
			},
			ForcedRedirect: &cdn.ForcedRedirect{
				EnableForcedRedirect: cdn.GetBoolPtr(false),
				StatusCode:           cdn.GetStrPtr(""),
			},
		},
		HttpForcedRedirect: &cdn.HttpForcedRedirect{
			EnableForcedRedirect: cdn.GetBoolPtr(false),
			StatusCode:           cdn.GetStrPtr(""),
		},
		TemplateId: "tpl-example",
	})
	test, _ := json.Marshal(resp)
	fmt.Println(string(test))
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.ResponseMetadata)
}
