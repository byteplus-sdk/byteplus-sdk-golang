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
		"CreateWatermarkPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateWatermarkPreset"},
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
		"CreateRecordPresetV2": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateRecordPresetV2"},
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
		"CreateSnapshotPreset": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateSnapshotPreset"},
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
		"UpdateTimeShiftPresetV3": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateTimeShiftPresetV3"},
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
		"DeleteCert": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteCert"},
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
		"DeleteDomain": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteDomain"},
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
		"UpdateDomainVhost": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateDomainVhost"},
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
		"DisableDomain": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DisableDomain"},
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
		"DescribeLiveStreamSessionData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveStreamSessionData"},
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
		"DescribeLivePushStreamInfoData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLivePushStreamInfoData"},
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
		"DescribeLiveASRDurationData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveASRDurationData"},
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
		"DescribeLiveStorageSpaceData": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLiveStorageSpaceData"},
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
		"DeleteReferer": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteReferer"},
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
		"DescribeAuth": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeAuth"},
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
		"EnableHTTPHeaderConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"EnableHTTPHeaderConfig"},
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
		"UpdateHTTPHeaderConfig": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateHTTPHeaderConfig"},
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
		"DescribeLicenseDRM": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeLicenseDRM"},
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
		"DeleteIPAccessRule": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteIPAccessRule"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteRegionAccessRule": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteRegionAccessRule"},
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
		"UpdateRegionAccessRule": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateRegionAccessRule"},
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
		"DescribeRegionAccessRule": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeRegionAccessRule"},
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
		"DeleteClusterRateLimit": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteClusterRateLimit"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeClusterRateLimit": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeClusterRateLimit"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateClusterRateLimit": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateClusterRateLimit"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteUserAgentAccessRule": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteUserAgentAccessRule"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeUserAgentAccessRule": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeUserAgentAccessRule"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateUserAgentAccessRule": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateUserAgentAccessRule"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DeleteFormatAccessRule": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DeleteFormatAccessRule"},
				"Version": []string{"2023-01-01"},
			},
		},
		"DescribeFormatAccessRule": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DescribeFormatAccessRule"},
				"Version": []string{"2023-01-01"},
			},
		},
		"UpdateFormatAccessRule": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UpdateFormatAccessRule"},
				"Version": []string{"2023-01-01"},
			},
		},
	}
)
