package billing

import (
	"encoding/json"
)

func (b *Billing) ListBillDetail(req *ListBillDetailRequest) (*ListBillDetailResp, int, error) {
	output := new(ListBillDetailResp)
	respBody, status, err := b.Client.Query(ActionListBillDetail, req.ToQuery())
	if err != nil {
		return output, status, err
	}
	err = json.Unmarshal(respBody, output)
	return output, status, err
}

func (b *Billing) ListSplitBillDetail(req *ListSplitBillDetailRequest) (*ListSplitBillDetailResp, int, error) {
	output := new(ListSplitBillDetailResp)
	respBody, status, err := b.Client.Query(ActionListSplitBillDetail, req.ToQuery())
	if err != nil {
		return output, status, err
	}
	err = json.Unmarshal(respBody, output)
	return output, status, err
}
