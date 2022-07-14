package cdn

import (
	"fmt"
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
	"testing"
)

func DescribeCdnConfig(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnConfig(&cdn.DescribeCdnConfigRequest{
		Domain: exampleDomain,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp.Result.DomainConfig)
	domain := resp.Result.DomainConfig
	fmt.Printf("%+v\n", domain)
}
