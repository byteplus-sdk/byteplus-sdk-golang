package cdn

import (
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
	"testing"
)

func StartCdnDomain(t *testing.T) {
	resp, err := DefaultInstance.StartCdnDomain(&cdn.StartCdnDomainRequest{Domain: operateDomain})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
