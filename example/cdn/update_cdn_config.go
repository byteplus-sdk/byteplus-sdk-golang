package cdn

import (
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func UpdateCdnConfig(t *testing.T) {
	resp, err := DefaultInstance.UpdateCdnConfig(&cdn.UpdateCdnConfigRequest{
		Domain: &exampleDomain,
		Origin: []cdn.OriginRule{
			{OriginAction: &cdn.OriginAction{
				OriginLines: []cdn.OriginLine{
					{
						OriginType:   cdn.GetStrPtr("primary"),
						InstanceType: cdn.GetStrPtr("ip"),
						Address:      cdn.GetStrPtr("1.1.1.1"),
						HttpPort:     cdn.GetStrPtr("80"),
						HttpsPort:    cdn.GetStrPtr("80"),
						Weight:       cdn.GetStrPtr("100"),
					},
				},
			}},
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
