package cdn

import (
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
	"testing"
)

func DeleteCdnDomain(t *testing.T) {
	resp, err := DefaultInstance.DeleteCdnDomain(&cdn.DeleteCdnDomainRequest{Domain: operateDomain})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
