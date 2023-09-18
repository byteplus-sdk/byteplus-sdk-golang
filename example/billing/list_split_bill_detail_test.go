package billing

import (
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/billing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ListSplitBillDetail(t *testing.T) {
	resp, status, err := DefaultInstance.ListSplitBillDetail(&billing.ListSplitBillDetailRequest{
		Offset:        0,
		Limit:         10,
		BillPeriod:    "2023-08",
		GroupPeriod:   2,
		Product:       "",
		BillingMode:   "",
		BillCategory:  "",
		InstanceNo:    "",
		SplitItemID:   "",
		IgnoreZero:    0,
		NeedRecordNum: 1,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result)
	t.Logf("status: %v, resp: %v", status, billing.ToJsonForLog(resp))
}
