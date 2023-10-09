package billing

import (
	"github.com/byteplus-sdk/byteplus-sdk-golang/base"
	"net/url"
)

type BillDetail struct {
	BillPeriod             string `json:"BillPeriod"`
	ExpenseDate            string `json:"ExpenseDate"`
	PayerID                string `json:"PayerID"`
	PayerUserName          string `json:"PayerUserName"`
	PayerCustomerName      string `json:"PayerCustomerName"`
	SellerID               string `json:"SellerID"`
	SellerUserName         string `json:"SellerUserName"`
	SellerCustomerName     string `json:"SellerCustomerName"`
	OwnerID                string `json:"OwnerID"`
	OwnerUserName          string `json:"OwnerUserName"`
	OwnerCustomerName      string `json:"OwnerCustomerName"`
	BusinessMode           string `json:"BusinessMode"`
	Product                string `json:"Product"`
	ProductZh              string `json:"ProductZh"`
	BillingMode            string `json:"BillingMode"`
	ExpenseBeginTime       string `json:"ExpenseBeginTime"`
	ExpenseEndTime         string `json:"ExpenseEndTime"`
	UseDuration            string `json:"UseDuration"`
	UseDurationUnit        string `json:"UseDurationUnit"`
	TradeTime              string `json:"TradeTime"`
	BillID                 string `json:"BillID"`
	BillCategory           string `json:"BillCategory"`
	InstanceNo             string `json:"InstanceNo"`
	InstanceName           string `json:"InstanceName"`
	ConfigName             string `json:"ConfigName"`
	Element                string `json:"Element"`
	Region                 string `json:"Region"`
	Zone                   string `json:"Zone"`
	Factor                 string `json:"Factor"`
	ExpandField            string `json:"ExpandField"`
	Price                  string `json:"Price"`
	PriceUnit              string `json:"PriceUnit"`
	Count                  string `json:"Count"`
	Unit                   string `json:"Unit"`
	DeductionCount         string `json:"DeductionCount"`
	OriginalBillAmount     string `json:"OriginalBillAmount"`
	PreferentialBillAmount string `json:"PreferentialBillAmount"`
	DiscountBillAmount     string `json:"DiscountBillAmount"`
	CouponAmount           string `json:"CouponAmount"`
	PaidAmount             string `json:"PaidAmount"`
	UnpaidAmount           string `json:"UnpaidAmount"`
	Currency               string `json:"Currency"`
	SettlementType         string `json:"SettlementType"`
	Project                string `json:"Project"`
	Tag                    string `json:"Tag"`
	SellingMode            string `json:"SellingMode"`
	SolutionZh             string `json:"SolutionZh"`
	ReservationInstance    string `json:"ReservationInstance"`
	BillDetailId           string `json:"BillDetailId"`
	ElementCode            string `json:"ElementCode"`
	RegionCode             string `json:"RegionCode"`
	ZoneCode               string `json:"ZoneCode"`
	FactorCode             string `json:"FactorCode"`
	ConfigurationCode      string `json:"ConfigurationCode"`
	DeductionUseDuration   string `json:"DeductionUseDuration"`
	Tax                    string `json:"Tax"`
	PosttaxAmount          string `json:"PosttaxAmount"`
	PretaxAmount           string `json:"PretaxAmount"`
	CountryRegion          string `json:"CountryRegion"`
}

type BillDetailList struct {
	List   []*BillDetail `json:"List"`
	Total  int           `json:"Total"`
	Limit  int           `json:"Limit"`
	Offset int           `json:"Offset"`
}

type ListBillDetailRequest struct {
	Offset        int64  `json:"Offset"`
	Limit         int64  `json:"Limit"`
	BillPeriod    string `json:"BillPeriod"`
	GroupTerm     int64  `json:"GroupTerm"`
	GroupPeriod   int64  `json:"GroupPeriod"`
	Product       string `json:"Product,omitempty"` // 当该字段的值为该字段类型的零值时，忽略该字段
	BillingMode   string `json:"BillingMode,omitempty"`
	BillCategory  string `json:"BillCategory,omitempty"`
	InstanceNo    string `json:"InstanceNo,omitempty"`
	IgnoreZero    int64  `json:"IgnoreZero"`
	NeedRecordNum int64  `json:"NeedRecordNum"`
}

func (r *ListBillDetailRequest) ToQuery() url.Values {
	return ToUrlValues(r)
}

type ListBillDetailResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *BillDetailList `json:"Result,omitempty"`
}

type SplitBillDetail struct {
	BillPeriod             string `json:"BillPeriod"`
	ExpenseTime            string `json:"ExpenseTime"`
	PayerUserName          string `json:"PayerUserName"`
	SellerUserName         string `json:"SellerUserName"`
	OwnerUserName          string `json:"OwnerUserName"`
	Product                string `json:"Product"`
	ProductZh              string `json:"ProductZh"`
	BusinessMode           string `json:"BusinessMode"`
	BillingMode            string `json:"BillingMode"`
	UseDuration            string `json:"UseDuration"`
	UseDurationUnit        string `json:"UseDurationUnit"`
	TradeTime              string `json:"TradeTime"`
	BillID                 string `json:"BillID"`
	BillCategory           string `json:"BillCategory"`
	SettlementType         string `json:"SettlementType"`
	InstanceNo             string `json:"InstanceNo"`
	InstanceName           string `json:"InstanceName"`
	ConfigName             string `json:"ConfigName"`
	Element                string `json:"Element"`
	Region                 string `json:"Region"`
	Zone                   string `json:"Zone"`
	Factor                 string `json:"Factor"`
	ExpandField            string `json:"ExpandField"`
	SplitItemID            string `json:"SplitItemID"`
	SplitItemName          string `json:"SplitItemName"`
	Price                  string `json:"Price"`
	PriceUnit              string `json:"PriceUnit"`
	SplitItemAmount        string `json:"SplitItemAmount"`
	Unit                   string `json:"Unit"`
	SplitItemRatio         string `json:"SplitItemRatio"`
	DeductionCount         string `json:"DeductionCount"`
	SolutionZh             string `json:"SolutionZh"`
	OriginalBillAmount     string `json:"OriginalBillAmount"`
	PreferentialBillAmount string `json:"PreferentialBillAmount"`
	DiscountBillAmount     string `json:"DiscountBillAmount"`
	CouponDeductionAmount  string `json:"CouponDeductionAmount"`
	PaidAmount             string `json:"PaidAmount"`
	UnpaidAmount           string `json:"UnpaidAmount"`
	Currency               string `json:"Currency"`
	Project                string `json:"Project"`
	Tag                    string `json:"Tag"`
	SellingMode            string `json:"SellingMode"`
	SubjectName            string `json:"SubjectName"`
	ReservationInstance    string `json:"ReservationInstance"`
	SplitBillDetailId      string `json:"SplitBillDetailId"`
	ElementCode            string `json:"ElementCode"`
	RegionCode             string `json:"RegionCode"`
	ZoneCode               string `json:"ZoneCode"`
	FactorCode             string `json:"FactorCode"`
	ConfigurationCode      string `json:"ConfigurationCode"`
	Tax                    string `json:"Tax"`
	PosttaxAmount          string `json:"PosttaxAmount"`
	PretaxAmount           string `json:"PretaxAmount"`
	CountryRegion          string `json:"CountryRegion"`
}

type SplitBillDetailList struct {
	List   []*SplitBillDetail `json:"List"`
	Total  int                `json:"Total"`
	Limit  int                `json:"Limit"`
	Offset int                `json:"Offset"`
}

type ListSplitBillDetailRequest struct {
	Offset        int64  `json:"Offset"`
	Limit         int64  `json:"Limit"`
	BillPeriod    string `json:"BillPeriod"`
	GroupPeriod   int64  `json:"GroupPeriod"`
	Product       string `json:"Product,omitempty"`
	BillingMode   string `json:"BillingMode,omitempty"`
	BillCategory  string `json:"BillCategory,omitempty"`
	InstanceNo    string `json:"InstanceNo,omitempty"`
	SplitItemID   string `json:"SplitItemID,omitempty"`
	IgnoreZero    int64  `json:"IgnoreZero"`
	NeedRecordNum int64  `json:"NeedRecordNum"`
}

func (r *ListSplitBillDetailRequest) ToQuery() url.Values {
	return ToUrlValues(r)
}

type ListSplitBillDetailResp struct {
	ResponseMetadata *base.ResponseMetadata
	Result           *SplitBillDetailList `json:"Result,omitempty"`
}
