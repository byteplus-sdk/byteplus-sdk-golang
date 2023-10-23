package billing

import (
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/billing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ListBillDetail(t *testing.T) {
	resp, status, err := DefaultInstance.ListBillDetail(&billing.ListBillDetailRequest{
		Offset:        0,
		Limit:         10,
		BillPeriod:    "2023-08",
		GroupTerm:     0,
		GroupPeriod:   2,
		Product:       "",
		BillingMode:   "",
		BillCategory:  "",
		InstanceNo:    "",
		IgnoreZero:    0,
		NeedRecordNum: 1,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result)
	t.Logf("status: %v, resp: %v", status, billing.ToJsonForLog(resp))
}

func TestListBillOverviewByProd(t *testing.T) {
	resp, status, err := DefaultInstance.ListBillOverviewByProd(&billing.ListBillOverviewByProdRequest{
		Offset:             0,
		Limit:              10,
		BillPeriod:         "2023-08",
		Product:            "",
		BillingMode:        "",
		BillCategoryParent: "",
		IgnoreZero:         0,
		NeedRecordNum:      1,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result)
	t.Logf("status: %v, resp: %v", status, billing.ToJsonForLog(resp))
}

func TestListBillOverviewByCategory(t *testing.T) {
	resp, status, err := DefaultInstance.ListBillOverviewByCategory(&billing.ListBillOverviewByCategoryRequest{
		Offset:             0,
		Limit:              10,
		BillPeriod:         "2023-08",
		BillingMode:        "",
		BillCategoryParent: "",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result)
	t.Logf("status: %v, resp: %v", status, billing.ToJsonForLog(resp))
}
