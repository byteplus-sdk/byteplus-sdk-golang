package cdn

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func ListCertInfo(t *testing.T) {
	resp, err := DefaultInstance.ListCertInfo(&cdn.ListCertInfoRequest{
		Source: cdn.GetStrPtr("volc_cert_center"),
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	rsp, _ := json.Marshal(resp)
	fmt.Printf("%+v", rsp)
}
