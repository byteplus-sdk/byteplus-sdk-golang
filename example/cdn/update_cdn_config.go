package cdn

import (
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
	"testing"
)

func UpdateCdnConfig(t *testing.T) {
	resp, err := DefaultInstance.UpdateCdnConfig(&cdn.UpdateCdnConfigRequest{
		Domain: exampleDomain,
		Origin: []cdn.OriginRule{
			{OriginAction: cdn.OriginAction{
				OriginLines: []cdn.OriginLine{
					{
						OriginType:   "primary",
						InstanceType: "domain",
						Address:      "new-origin.com",
						HttpPort:     "80",
						HttpsPort:    "80",
						Weight:       "100",
					},
				},
			}},
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
