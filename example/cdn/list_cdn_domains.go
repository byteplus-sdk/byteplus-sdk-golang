package cdn

import (
	"encoding/json"
	"fmt"
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
	"testing"
)

func ListCdnDomains(t *testing.T) {
	resp, err := DefaultInstance.ListCdnDomains(&cdn.ListCdnDomainsRequest{})
	test, _ := json.Marshal(resp)
	fmt.Println(string(test))
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.Data)
}
