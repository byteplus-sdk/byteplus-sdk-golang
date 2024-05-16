package cdn

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func UpdateServiceTemplate(t *testing.T) {
	resp, err := DefaultInstance.UpdateServiceTemplate(&cdn.UpdateServiceTemplateRequest{
		Title:          cdn.GetStrPtr("example"),
		Message:        cdn.GetStrPtr("example"),
		OriginProtocol: "http",
		OriginHost:     cdn.GetStrPtr(""),
		Origin: []cdn.OriginRule{
			{
				OriginAction: &cdn.OriginAction{
					OriginLines: []cdn.OriginLine{
						{
							Address:      cdn.GetStrPtr("example.com"),
							HttpPort:     cdn.GetStrPtr("80"),
							HttpsPort:    cdn.GetStrPtr("443"),
							InstanceType: cdn.GetStrPtr("domain"),
							OriginHost:   cdn.GetStrPtr(""),
							OriginType:   cdn.GetStrPtr("primary"),
							Weight:       cdn.GetStrPtr("1"),
						},
					},
				},
			},
		},
		TemplateId: "tpl-example",
	})
	test, _ := json.Marshal(resp)
	fmt.Println(string(test))
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.ResponseMetadata)
}
