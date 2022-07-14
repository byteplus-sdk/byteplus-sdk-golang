package cdn

import (
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
	"testing"
)

func StopCdnDomain(t *testing.T) {
	resp, err := DefaultInstance.StopCdnDomain(&cdn.StopCdnDomainRequest{Domain: operateDomain})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
