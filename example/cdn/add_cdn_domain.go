package cdn

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func AddCdnDomain(t *testing.T) {
	resp, err := DefaultInstance.AddCdnDomain(&cdn.AddCdnDomainRequest{
		Domain:      operateDomain,
		ServiceType: cdn.GetStrPtr("web"),
		Origin: []cdn.OriginRule{
			{OriginAction: cdn.OriginAction{
				OriginLines: []cdn.OriginLine{
					{
						OriginType:   "primary",
						InstanceType: "domain",
						Address:      "origin.test.com",
						HttpPort:     "80",
						HttpsPort:    "80",
						Weight:       "100",
					},
				},
			}},
		},
		OriginProtocol: "http",
	})
	rsp, _ := json.Marshal(resp)
	fmt.Println(string(rsp))
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
