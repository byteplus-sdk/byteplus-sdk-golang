package live_v20230101

import (
	"net/http"
	"net/url"
	"time"

	common "github.com/byteplus-sdk/byteplus-sdk-golang/base"
)

const (
	ServiceName    = "live"
	DefaultTimeout = 10 * time.Second
)

var (
	ServiceInfoMap = map[string]common.ServiceInfo{
		"cn-north-1": {
			Timeout: DefaultTimeout,
			Scheme:  "https",
			Host:    "open.byteplusapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: common.Credentials{
				Region:  "cn-north-1",
				Service: ServiceName,
			},
		},
	}
	ApiListInfo = map[string]*common.ApiInfo{

		"DeleteTranscodePreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteTranscodePreset"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteCommonTransPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteCommonTransPreset"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteTranscodePresetPatchByAdmin": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteTranscodePresetPatchByAdmin"},
				"Version": []string{"2023-01-01"},
			},
		},
		"AddCommonTransPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AddCommonTransPreset"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateTranscodePreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateTranscodePreset"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListCommonTransPresetDetail": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListCommonTransPresetDetail"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeTranscodePresetDetail": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeTranscodePresetDetail"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListVhostTransCodePreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListVhostTransCodePreset"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CreateTranscodePreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateTranscodePreset"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CreateTranscodePresetPatchByAdmin": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateTranscodePresetPatchByAdmin"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteWatermarkPresetV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteWatermarkPresetV2"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateWatermarkPresetV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateWatermarkPresetV2"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeWatermarkPresetDetail": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeWatermarkPresetDetail"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CreateWatermarkPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateWatermarkPreset"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CreateWatermarkPresetV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateWatermarkPresetV2"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateWatermarkPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateWatermarkPreset"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteWatermarkPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteWatermarkPreset"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListWatermarkPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListWatermarkPreset"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListVhostWatermarkPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListVhostWatermarkPreset"},
				"Version": []string{"2023-01-01"},
			},
		},
		"StopPullRecordTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopPullRecordTask"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CreatePullRecordTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreatePullRecordTask"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteRecordHistory": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteRecordHistory"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteRecordPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteRecordPreset"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateRecordPresetV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateRecordPresetV2"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeRecordTaskFileHistory": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeRecordTaskFileHistory"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListVhostRecordPresetV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListVhostRecordPresetV2"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListPullRecordTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListPullRecordTask"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CreateRecordPresetV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateRecordPresetV2"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListVideoClassifications": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListVideoClassifications"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteSnapshotPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteSnapshotPreset"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateSnapshotPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateSnapshotPreset"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateSnapshotPresetV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateSnapshotPresetV2"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeCDNSnapshotHistory": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeCDNSnapshotHistory"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListVhostSnapshotPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListVhostSnapshotPreset"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListVhostSnapshotPresetV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListVhostSnapshotPresetV2"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CreateSnapshotPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateSnapshotPreset"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CreateSnapshotPresetV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateSnapshotPresetV2"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteTimeShiftPresetV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteTimeShiftPresetV2"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteTimeShiftPresetV3": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteTimeShiftPresetV3"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CreateTimeShiftPresetV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateTimeShiftPresetV2"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateTimeShiftPresetV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateTimeShiftPresetV2"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateTimeShiftPresetV3": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateTimeShiftPresetV3"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeTimeShiftPresetDetail": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeTimeShiftPresetDetail"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListTimeShiftPresetV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListTimeShiftPresetV2"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CreateTimeShiftPresetV3": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateTimeShiftPresetV3"},
				"Version": []string{"2023-01-01"},
			},
		},
		"GenerateTimeShiftPlayURL": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GenerateTimeShiftPlayURL"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListVhostDomainDetailByUserID": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListVhostDomainDetailByUserID"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateVhostTags": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateVhostTags"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListVhostDetail": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListVhostDetail"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListVhostDetailByAdmin": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListVhostDetailByAdmin"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeVhost": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeVhost"},
				"Version": []string{"2023-01-01"},
			},
		},
		"GetTags": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetTags"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListProjects": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListProjects"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteCallback": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteCallback"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeCallback": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeCallback"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateCallback": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateCallback"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteAuth": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteAuth"},
				"Version": []string{"2023-01-01"},
			},
		},
		"EnableAuth": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"EnableAuth"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeAuth": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeAuth"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DisableAuth": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DisableAuth"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteCert": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteCert"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateCert": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateCert"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeCertDetailSecretV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeCertDetailSecretV2"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListCertV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListCertV2"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeCertDetailV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeCertDetailV2"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CreateCert": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateCert"},
				"Version": []string{"2023-01-01"},
			},
		},
		"BindCert": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BindCert"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UnbindCert": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UnbindCert"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListObject": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListObject"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ManagerPullPushDomainBind": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ManagerPullPushDomainBind"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeCertDetailSecret": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeCertDetailSecret"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeDomainVerify": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeDomainVerify"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CreateVerifyContent": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateVerifyContent"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListVqosDimensionValues": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListVqosDimensionValues"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListCert": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListCert"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListCertBindInfo": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListCertBindInfo"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveFreeTimeInterval": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveFreeTimeInterval"},
				"Version": []string{"2023-01-01"},
			},
		},
		"VerifyDomainOwner": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"VerifyDomainOwner"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ValidateCert": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ValidateCert"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteDomain": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteDomain"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteDomainV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteDomainV2"},
				"Version": []string{"2023-01-01"},
			},
		},
		"EnableDomain": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"EnableDomain"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CreateDomainV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateDomainV2"},
				"Version": []string{"2023-01-01"},
			},
		},
		"RejectDomain": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RejectDomain"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateDomainVhost": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateDomainVhost"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateDomain": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateDomain"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeDomain": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeDomain"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListDomainDetail": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListDomainDetail"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CreateDomain": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateDomain"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DisableDomain": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DisableDomain"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CreateVQScoreTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateVQScoreTask"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeVQScoreTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeVQScoreTask"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListVQScoreTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListVQScoreTask"},
				"Version": []string{"2023-01-01"},
			},
		},
		"StopPullToPushTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopPullToPushTask"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CreatePullToPushTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreatePullToPushTask"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeletePullToPushTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeletePullToPushTask"},
				"Version": []string{"2023-01-01"},
			},
		},
		"RestartPullToPushTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RestartPullToPushTask"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdatePullToPushTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdatePullToPushTask"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListPullToPushTask": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListPullToPushTask"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteDenyConfigV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteDenyConfigV2"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeDenyConfigV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeDenyConfigV2"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteRelaySourceRewrite": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteRelaySourceRewrite"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteRelaySourceV4": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteRelaySourceV4"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteRelaySourceV3": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteRelaySourceV3"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateRelaySourceRewrite": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateRelaySourceRewrite"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateRelaySourceV4": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateRelaySourceV4"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeRelaySourceRewrite": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeRelaySourceRewrite"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListRelaySourceV4": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListRelaySourceV4"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeRelaySourceV3": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeRelaySourceV3"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CreateRelaySourceV4": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateRelaySourceV4"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateRelaySourceV3": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateRelaySourceV3"},
				"Version": []string{"2023-01-01"},
			},
		},
		"KillStream": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"KillStream"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeClosedStreamInfoByPage": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeClosedStreamInfoByPage"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveStreamInfoByPage": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveStreamInfoByPage"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveStreamState": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveStreamState"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeForbiddenStreamInfoByPage": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeForbiddenStreamInfoByPage"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ForbidStream": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ForbidStream"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ResumeStream": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ResumeStream"},
				"Version": []string{"2023-01-01"},
			},
		},
		"GeneratePlayURL": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GeneratePlayURL"},
				"Version": []string{"2023-01-01"},
			},
		},
		"GeneratePushURL": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GeneratePushURL"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateSDKLicense": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateSDKLicense"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CreateApp": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateApp"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteSDK": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteSDK"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateApp": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateApp"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateSDK": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateSDK"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeSDKDetail": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeSDKDetail"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeSDKParamsAvailable": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeSDKParamsAvailable"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeAppIDParamsAvailable": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeAppIDParamsAvailable"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CreateSDK": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateSDK"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListSDK": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListSDK"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListSDKAdmin": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListSDKAdmin"},
				"Version": []string{"2023-01-01"},
			},
		},
		"GetApps": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetApps"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateService": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateService"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListServices": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListServices"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeService": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeService"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateActivityBilling": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateActivityBilling"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateBilling": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateBilling"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListInstance": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListInstance"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeBillingForAdmin": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeBillingForAdmin"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeBilling": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeBilling"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeBillingMonthAvailable": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeBillingMonthAvailable"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListResourcePackage": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListResourcePackage"},
				"Version": []string{"2023-01-01"},
			},
		},
		"TerminateInstance": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"TerminateInstance"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteStreamQuotaConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteStreamQuotaConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateStreamQuotaConfigPatch": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateStreamQuotaConfigPatch"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeStreamQuotaConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeStreamQuotaConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateStreamQuotaConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateStreamQuotaConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListVqosMetricsDimensions": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListVqosMetricsDimensions"},
				"Version": []string{"2023-01-01"},
			},
		},
		"GetVqosRawData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetVqosRawData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"StopPullCDNSnapshotTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StopPullCDNSnapshotTask"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CreatePullCDNSnapshotTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreatePullCDNSnapshotTask"},
				"Version": []string{"2023-01-01"},
			},
		},
		"GetPullCDNSnapshotTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetPullCDNSnapshotTask"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListPullCDNSnapshotTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListPullCDNSnapshotTask"},
				"Version": []string{"2023-01-01"},
			},
		},
		"GetPullRecordTask": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetPullRecordTask"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteSnapshotAuditPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteSnapshotAuditPreset"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateSnapshotAuditPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateSnapshotAuditPreset"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeSnapshotAuditPresetDetail": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeSnapshotAuditPresetDetail"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListVhostSnapshotAuditPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListVhostSnapshotAuditPreset"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CreateSnapshotAuditPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateSnapshotAuditPreset"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeIpInfo": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeIpInfo"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveRegionData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveRegionData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveSourceStreamMetrics": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveSourceStreamMetrics"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLivePushStreamMetrics": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLivePushStreamMetrics"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLivePlayStatusCodeData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLivePlayStatusCodeData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveBatchSourceStreamMetrics": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveBatchSourceStreamMetrics"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveBatchSourceStreamAvgMetrics": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveBatchSourceStreamAvgMetrics"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveBatchOnlineStreamMetrics": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveBatchOnlineStreamMetrics"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveBatchPushStreamMetrics": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveBatchPushStreamMetrics"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveBatchPushStreamAvgMetrics": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveBatchPushStreamAvgMetrics"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveBatchStreamTranscodeData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveBatchStreamTranscodeData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveStreamCountData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveStreamCountData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLivePushStreamCountData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLivePushStreamCountData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveSourceBandwidthData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveSourceBandwidthData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveSourceTrafficData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveSourceTrafficData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveMetricBandwidthData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveMetricBandwidthData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveMetricTrafficData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveMetricTrafficData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveBatchStreamTrafficData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveBatchStreamTrafficData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveStreamSessionData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveStreamSessionData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveISPData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveISPData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveP95PeakBandwidthData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveP95PeakBandwidthData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveAuditData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveAuditData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLivePullToPushBandwidthData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLivePullToPushBandwidthData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLivePullToPushData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLivePullToPushData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveBandwidthData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveBandwidthData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveRecordData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveRecordData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveSnapshotData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveSnapshotData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveTrafficData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveTrafficData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveTranscodeData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveTranscodeData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveTimeShiftData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveTimeShiftData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeActionHistory": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeActionHistory"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListActionHistory": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListActionHistory"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteDenseSnapshotPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteDenseSnapshotPreset"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateDenseSnapshotPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateDenseSnapshotPreset"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListVhostDenseSnapshotPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListVhostDenseSnapshotPreset"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescDenseSnapshotPresetDetail": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescDenseSnapshotPresetDetail"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CreateDenseSnapshotPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateDenseSnapshotPreset"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveCustomizedLogData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveCustomizedLogData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveLogData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveLogData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"AssociatePreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AssociatePreset"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DisAssociatePreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DisAssociatePreset"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdatePresetAssociation": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdatePresetAssociation"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribePresetAssociation": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribePresetAssociation"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CreateTicket": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateTicket"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteReferer": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteReferer"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeDenyConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeDenyConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeReferer": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeReferer"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateDenyConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateDenyConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateDenyConfigV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateDenyConfigV2"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateReferer": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateReferer"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateAuthKey": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateAuthKey"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteRelaySink": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteRelaySink"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateRelaySink": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateRelaySink"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeRelaySink": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeRelaySink"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteHLSConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteHLSConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateHLSConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateHLSConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeHLSConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeHLSConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteHTTPHeaderConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteHTTPHeaderConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteHeaderConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteHeaderConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"EnableHTTPHeaderConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"EnableHTTPHeaderConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateHTTPHeaderConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateHTTPHeaderConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateHeaderConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateHeaderConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeHTTPHeaderConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeHTTPHeaderConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeHeaderConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeHeaderConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListHeaderEnum": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListHeaderEnum"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteNSSRewriteConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteNSSRewriteConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateNSSRewriteConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateNSSRewriteConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeNSSRewriteConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeNSSRewriteConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveActivityBandwidthData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveActivityBandwidthData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CreateLiveAccountFeeConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateLiveAccountFeeConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteLiveAccountFeeConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteLiveAccountFeeConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveAccountFeeConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveAccountFeeConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveStreamUsageData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveStreamUsageData"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveFeeConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveFeeConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLiveAccountFeeType": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveAccountFeeType"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateEncryptDRM": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateEncryptDRM"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeContentKey": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeContentKey"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeCertDRM": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeCertDRM"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLicenseDRM": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLicenseDRM"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeEncryptDRM": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeEncryptDRM"},
				"Version": []string{"2023-01-01"},
			},
		},
		"BindEncryptDRM": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BindEncryptDRM"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UnBindEncryptDRM": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UnBindEncryptDRM"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListBindEncryptDRM": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListBindEncryptDRM"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CreateCustomLogConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateCustomLogConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteCustomLogConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteCustomLogConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeCustomLogConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeCustomLogConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CheckCustomLogConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CheckCustomLogConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CreateTranscodePresetBatch": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateTranscodePresetBatch"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteTranscodePresetBatch": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteTranscodePresetBatch"},
				"Version": []string{"2023-01-01"},
			},
		},
		"AssociateRefConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"AssociateRefConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DisassociateRefConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DisassociateRefConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeRefConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeRefConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListReferenceNames": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListReferenceNames"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListReferenceTypes": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListReferenceTypes"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListReferenceInfo": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListReferenceInfo"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CreateAvSlicePreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateAvSlicePreset"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteAvSlicePreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteAvSlicePreset"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateAvSlicePreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateAvSlicePreset"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteIPAccessRule": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteIPAccessRule"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateIPAccessRule": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateIPAccessRule"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeIPAccessRule": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeIPAccessRule"},
				"Version": []string{"2023-01-01"},
			},
		},
		"CreateProxyConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateProxyConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteProxyConfigAssociation": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteProxyConfigAssociation"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteProxyConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteProxyConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateProxyConfigAssociation": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateProxyConfigAssociation"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateProxyConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateProxyConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeProxyConfigAssociation": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeProxyConfigAssociation"},
				"Version": []string{"2023-01-01"},
			},
		},
		"ListProxyConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ListProxyConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteCMAFConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteCMAFConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateCMAFConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateCMAFConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeCMAFConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeCMAFConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteLatencyConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteLatencyConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateLatencyConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateLatencyConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeLatencyConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLatencyConfig"},
				"Version": []string{"2023-01-01"},
			},
		},
	}
)
