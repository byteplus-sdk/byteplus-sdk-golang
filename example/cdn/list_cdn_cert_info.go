package cdn

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ListCdnCertInfo(t *testing.T) {
	resp, err := DefaultInstance.ListCdnCertInfo(nil)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", resp)
}
