package cdn

type AccountingDataDetail struct {
	BillingRegion *string `json:",omitempty"`
	Metrics       []MetricTimeStampValue
	Name          *string `json:",omitempty"`
}

type AddCdnCertInfo struct {
	Desc *string `json:",omitempty"`
}

type AddCdnCertificateRequest struct {
	CertInfo    *AddCdnCertInfo `json:",omitempty"`
	Certificate Certificate
	Source      *string `json:",omitempty"`
}

type AddCdnCertificateResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           AddCdnCertificateResult
}

type AddCdnCertificateResult struct {
	AddCdnCertificateResult *string `json:",omitempty"`
}

type AddCdnDomainRequest struct {
	AreaAccessRule     *AreaAccessRule `json:",omitempty"`
	BandwidthLimit     *BandwidthLimit `json:",omitempty"`
	Cache              []CacheControlRule
	CacheKey           []CacheKeyGenerationRule
	Compression        *Compression     `json:",omitempty"`
	CustomErrorPage    *CustomErrorPage `json:",omitempty"`
	Domain             string
	DownloadSpeedLimit *DownloadSpeedLimit `json:",omitempty"`
	FollowRedirect     *bool               `json:",omitempty"`
	HTTPS              *HTTPS              `json:",omitempty"`
	HttpForcedRedirect *HttpForcedRedirect `json:",omitempty"`
	IPv6               *IPv6               `json:",omitempty"`
	IpAccessRule       *IpAccessRule       `json:",omitempty"`
	IpFreqLimit        *IpFreqLimit        `json:",omitempty"`
	IpSpeedLimit       *IpSpeedLimit       `json:",omitempty"`
	MethodDeniedRule   *MethodDeniedRule   `json:",omitempty"`
	NegativeCache      []NegativeCache
	Origin             []OriginRule
	OriginAccessRule   *OriginAccessRule `json:",omitempty"`
	OriginArg          []OriginArgRule
	OriginHost         *string `json:",omitempty"`
	OriginProtocol     string
	OriginRange        *bool               `json:",omitempty"`
	OriginSni          *OriginSni          `json:",omitempty"`
	Project            *string             `json:",omitempty"`
	Quic               *Quic               `json:",omitempty"`
	RedirectionRewrite *RedirectionRewrite `json:",omitempty"`
	RefererAccessRule  *RefererAccessRule  `json:",omitempty"`
	RemoteAuth         *RemoteAuth         `json:",omitempty"`
	RequestHeader      []RequestHeaderRule
	ResponseHeader     []ResponseHeaderRule
	ServiceRegion      *string              `json:",omitempty"`
	ServiceType        *string              `json:",omitempty"`
	SignedUrlAuth      *SignedUrlAuth       `json:",omitempty"`
	UaAccessRule       *UserAgentAccessRule `json:",omitempty"`
	VideoDrag          *VideoDrag           `json:",omitempty"`
}

type AddCdnDomainResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}

type AddResourceTagsRequest struct {
	ResourceTags []ResourceTag
	Resources    []string
}

type AddResourceTagsResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}

type AreaAccessRule struct {
	Area     []string
	RuleType *string `json:",omitempty"`
	Switch   *bool   `json:",omitempty"`
}

type AreaAccessRuleUpdate struct {
	Area     []string
	RuleType *string `json:",omitempty"`
	Switch   *bool   `json:",omitempty"`
}

type AuthCacheAction struct {
	Action     string
	CacheKey   []string
	IgnoreCase bool
	Ttl        int64
}

type AuthCacheActionUpdate struct {
	Action     *string `json:",omitempty"`
	CacheKey   []string
	IgnoreCase *bool  `json:",omitempty"`
	Ttl        *int64 `json:",omitempty"`
}

type AuthModeConfig struct {
	BackupRemoteAddr *string `json:",omitempty"`
	MasterRemoteAddr string
	PathType         string
	PathValue        *string `json:",omitempty"`
	RequestMethod    string
}

type AuthModeConfigUpdate struct {
	BackupRemoteAddr *string `json:",omitempty"`
	MasterRemoteAddr *string `json:",omitempty"`
	PathType         *string `json:",omitempty"`
	PathValue        *string `json:",omitempty"`
	RequestMethod    *string `json:",omitempty"`
}

type AuthRequestHeaderRule struct {
	RequestHeaderComponents *RequestHeaderComponent `json:",omitempty"`
	RequestHeaderInstances  []RequestHeaderInstance
}

type AuthRequestHeaderRuleUpdate struct {
	RequestHeaderComponents *RequestHeaderComponentUpdate `json:",omitempty"`
	RequestHeaderInstances  []RequestHeaderInstanceUpdate
}

type AuthResponseConfig struct {
	CacheAction      *AuthCacheAction  `json:",omitempty"`
	ResponseAction   *ResponseAction   `json:",omitempty"`
	StatusCodeAction *StatusCodeAction `json:",omitempty"`
	TimeOutAction    *TimeOutAction    `json:",omitempty"`
}

type AuthResponseConfigUpdate struct {
	CacheAction      *AuthCacheActionUpdate  `json:",omitempty"`
	ResponseAction   *ResponseActionUpdate   `json:",omitempty"`
	StatusCodeAction *StatusCodeActionUpdate `json:",omitempty"`
	TimeOutAction    *TimeOutActionUpdate    `json:",omitempty"`
}

type BandwidthLimit struct {
	BandwidthLimitRule *BandwidthLimitRule `json:",omitempty"`
	Switch             bool
}

type BandwidthLimitAction struct {
	BandwidthThreshold int64
	LimitType          string
	SpeedLimitRate     *int64 `json:",omitempty"`
}

type BandwidthLimitActionUpdate struct {
	BandwidthThreshold *int64  `json:",omitempty"`
	LimitType          *string `json:",omitempty"`
	SpeedLimitRate     *int64  `json:",omitempty"`
}

type BandwidthLimitRule struct {
	BandwidthLimitAction *BandwidthLimitAction `json:",omitempty"`
	Condition            *Condition            `json:",omitempty"`
}

type BandwidthLimitRuleUpdate struct {
	BandwidthLimitAction *BandwidthLimitActionUpdate `json:",omitempty"`
	Condition            *ConditionUpdate            `json:",omitempty"`
}

type BandwidthLimitUpdate struct {
	BandwidthLimitRule *BandwidthLimitRuleUpdate `json:",omitempty"`
	Switch             *bool                     `json:",omitempty"`
}

type BatchDeployCertRequest struct {
	CertId string
	Domain string
}

type BatchDeployCertResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           BatchDeployCertResult
}

type BatchDeployCertResult struct {
	DeployResult []DomainCertDeployStatus
}

type CacheAction struct {
	Action     string
	IgnoreCase bool
	Ttl        int64
}

type CacheActionUpdate struct {
	Action     *string `json:",omitempty"`
	IgnoreCase *bool   `json:",omitempty"`
	Ttl        *int64  `json:",omitempty"`
}

type CacheControlRule struct {
	CacheAction CacheAction
	Condition   Condition
}

type CacheControlRuleUpdate struct {
	CacheAction *CacheActionUpdate `json:",omitempty"`
	Condition   *ConditionUpdate   `json:",omitempty"`
}

type CacheKeyAction struct {
	CacheKeyComponents []CacheKeyComponent
}

type CacheKeyActionUpdate struct {
	CacheKeyComponents []CacheKeyComponentUpdate
}

type CacheKeyComponent struct {
	Action     string
	IgnoreCase bool
	Object     string
	Subobject  string
}

type CacheKeyComponentUpdate struct {
	Action     *string `json:",omitempty"`
	IgnoreCase *bool   `json:",omitempty"`
	Object     *string `json:",omitempty"`
	Subobject  *string `json:",omitempty"`
}

type CacheKeyGenerationRule struct {
	CacheKeyAction CacheKeyAction
	Condition      Condition
}

type CacheKeyGenerationRuleUpdate struct {
	CacheKeyAction *CacheKeyActionUpdate `json:",omitempty"`
	Condition      *ConditionUpdate      `json:",omitempty"`
}

type CertInfo struct {
	CertId           *string      `json:",omitempty"`
	CertName         *string      `json:",omitempty"`
	Certificate      *Certificate `json:",omitempty"`
	ConfiguredDomain *string      `json:",omitempty"`
	Desc             *string      `json:",omitempty"`
	DnsName          *string      `json:",omitempty"`
	EffectiveTime    *int64       `json:",omitempty"`
	ExpireTime       *int64       `json:",omitempty"`
	Source           *string      `json:",omitempty"`
	Status           *string      `json:",omitempty"`
}

type CertInfoUpdate struct {
	CertId           *string            `json:",omitempty"`
	CertName         *string            `json:",omitempty"`
	Certificate      *CertificateUpdate `json:",omitempty"`
	ConfiguredDomain *string            `json:",omitempty"`
	Desc             *string            `json:",omitempty"`
	DnsName          *string            `json:",omitempty"`
	EffectiveTime    *int64             `json:",omitempty"`
	ExpireTime       *int64             `json:",omitempty"`
	Source           *string            `json:",omitempty"`
	Status           *string            `json:",omitempty"`
}

type Certificate struct {
	Certificate string
	PrivateKey  string
}

type CertificateUpdate struct {
	Certificate *string `json:",omitempty"`
	PrivateKey  *string `json:",omitempty"`
}

type Compression struct {
	CompressionRules []CompressionRule
	Switch           bool
}

type CompressionAction struct {
	CompressionFormat *string `json:",omitempty"`
	CompressionTarget *string `json:",omitempty"`
	CompressionType   []string
}

type CompressionActionUpdate struct {
	CompressionFormat *string `json:",omitempty"`
	CompressionTarget *string `json:",omitempty"`
	CompressionType   []string
}

type CompressionRule struct {
	CompressionAction CompressionAction
	Condition         *Condition `json:",omitempty"`
}

type CompressionRuleUpdate struct {
	CompressionAction *CompressionActionUpdate `json:",omitempty"`
	Condition         *ConditionUpdate         `json:",omitempty"`
}

type CompressionUpdate struct {
	CompressionRules []CompressionRuleUpdate
	Switch           *bool `json:",omitempty"`
}

type Condition struct {
	ConditionRule []ConditionRule
	Connective    string
}

type ConditionRule struct {
	Object   string
	Operator string
	Type     string
	Value    string
}

type ConditionRuleUpdate struct {
	Object   *string `json:",omitempty"`
	Operator *string `json:",omitempty"`
	Type     *string `json:",omitempty"`
	Value    *string `json:",omitempty"`
}

type ConditionUpdate struct {
	ConditionRule []ConditionRuleUpdate
	Connective    *string `json:",omitempty"`
}

type ContentTask struct {
	CreateTime *int64  `json:",omitempty"`
	Process    *string `json:",omitempty"`
	Remark     *string `json:",omitempty"`
	Status     *string `json:",omitempty"`
	TaskID     *string `json:",omitempty"`
	TaskType   *string `json:",omitempty"`
	Url        *string `json:",omitempty"`
}

type CustomErrorPage struct {
	ErrorPageRule []ErrorPageRule
	Switch        bool
}

type CustomErrorPageUpdate struct {
	ErrorPageRule []ErrorPageRuleUpdate
	Switch        *bool `json:",omitempty"`
}

type DeleteCdnDomainRequest struct {
	Domain string
}

type DeleteCdnDomainResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}

type DeleteResourceTagsRequest struct {
	ResourceTags []ResourceTag
	Resources    []string
}

type DeleteResourceTagsResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}

type DescribeAccountingDataRequest struct {
	Aggregate        *string `json:",omitempty"`
	BillingRegion    *string `json:",omitempty"`
	Domain           *string `json:",omitempty"`
	EndTime          int64
	Interval         *int64 `json:",omitempty"`
	IsWildcardDomain *bool  `json:",omitempty"`
	Metric           string
	Protocol         *string `json:",omitempty"`
	StartTime        int64
}

type DescribeAccountingDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeAccountingDataResult
}

type DescribeAccountingDataResult struct {
	Resources []AccountingDataDetail
}

type DescribeCdnAccessLogRequest struct {
	Domain        string
	EndTime       int64
	FileInterval  *int64  `json:",omitempty"`
	PageNum       *int64  `json:",omitempty"`
	PageSize      *int64  `json:",omitempty"`
	ServiceRegion *string `json:",omitempty"`
	StartTime     int64
}

type DescribeCdnAccessLogResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeCdnAccessLogResult
}

type DescribeCdnAccessLogResult struct {
	Domain           *string `json:",omitempty"`
	DomainLogDetails []DomainLogDetail
	PageNum          *int64 `json:",omitempty"`
	PageSize         *int64 `json:",omitempty"`
	TotalCount       *int64 `json:",omitempty"`
}

type DescribeCdnConfigRequest struct {
	Domain string
}

type DescribeCdnConfigResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeCdnConfigResult
}

type DescribeCdnConfigResult struct {
	DomainConfig     *DomainVolcanoDetail `json:",omitempty"`
	ModuleLockConfig *ModuleLockConfig    `json:",omitempty"`
}

type DescribeCdnDataDetailRequest struct {
	Domain    string
	EndTime   int64
	Interval  *string `json:",omitempty"`
	IpVersion *string `json:",omitempty"`
	Metric    string
	Protocol  *string `json:",omitempty"`
	StartTime int64
}

type DescribeCdnDataDetailResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeCdnDataDetailResult
}

type DescribeCdnDataDetailResult struct {
	DataDetails []NrtDataDetails
	Name        *string `json:",omitempty"`
}

type DescribeCdnDataRequest struct {
	Aggregate           *string `json:",omitempty"`
	Area                *string `json:",omitempty"`
	BillingRegion       *string `json:",omitempty"`
	DisaggregateMetrics *string `json:",omitempty"`
	Domain              *string `json:",omitempty"`
	EndTime             int64
	Interval            *string `json:",omitempty"`
	IpVersion           *string `json:",omitempty"`
	IsWildcardDomain    *bool   `json:",omitempty"`
	Isp                 *string `json:",omitempty"`
	Metric              string
	Protocol            *string `json:",omitempty"`
	Region              *string `json:",omitempty"`
	StartTime           int64
}

type DescribeCdnDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeCdnDataResult
}

type DescribeCdnDataResult struct {
	Resources []NrtDataResource
}

type DescribeCdnDomainTopDataRequest struct {
	Area      *string `json:",omitempty"`
	Domain    *string `json:",omitempty"`
	EndTime   int64
	Item      string
	Metric    string
	StartTime int64
}

type DescribeCdnDomainTopDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeCdnDomainTopDataResult
}

type DescribeCdnDomainTopDataResult struct {
	Domain         *string `json:",omitempty"`
	TopDataDetails []TopDataDetail
}

type DescribeCdnOriginDataRequest struct {
	Aggregate           *string `json:",omitempty"`
	BillingRegion       *string `json:",omitempty"`
	DisaggregateMetrics *string `json:",omitempty"`
	Domain              *string `json:",omitempty"`
	EndTime             int64
	Interval            *string `json:",omitempty"`
	IsWildcardDomain    *bool   `json:",omitempty"`
	Metric              string
	StartTime           int64
}

type DescribeCdnOriginDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeCdnOriginDataResult
}

type DescribeCdnOriginDataResult struct {
	Resources []NrtDataResource
}

type DescribeCdnRegionAndIspRequest struct {
	Area *string `json:",omitempty"`
}

type DescribeCdnRegionAndIspResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeCdnRegionAndIspResult
}

type DescribeCdnRegionAndIspResult struct {
	Isps    []NamePair
	Regions []NamePair
}

type DescribeCdnServiceResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeCdnServiceResult
}

type DescribeCdnServiceResult struct {
	ServiceInfos []TopInstanceDetail
}

type DescribeCdnUpperIpRequest struct {
	Domain    string
	IpVersion *string `json:",omitempty"`
}

type DescribeCdnUpperIpResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeCdnUpperIpResult
}

type DescribeCdnUpperIpResult struct {
	CdnIpv4 []string
	CdnIpv6 []string
}

type DescribeCertConfigRequest struct {
	CertId *string `json:",omitempty"`
	Status *string `json:",omitempty"`
}

type DescribeCertConfigResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeCertConfigResult
}

type DescribeCertConfigResult struct {
	CertNotConfig       []DomainCertResult
	OtherCertConfig     []DomainCertResult
	SpecifiedCertConfig []DomainCertResult
}

type DescribeContentBlockTasksRequest struct {
	EndTime   *int64  `json:",omitempty"`
	PageNum   *int64  `json:",omitempty"`
	PageSize  *int64  `json:",omitempty"`
	StartTime *int64  `json:",omitempty"`
	Status    *string `json:",omitempty"`
	TaskID    *string `json:",omitempty"`
	TaskType  string
	URL       *string `json:",omitempty"`
}

type DescribeContentBlockTasksResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeContentBlockTasksResult
}

type DescribeContentBlockTasksResult struct {
	Data     []DescribeContentBlockTasksTaskInfo
	PageNum  *int64 `json:",omitempty"`
	PageSize *int64 `json:",omitempty"`
	Total    *int64 `json:",omitempty"`
}

type DescribeContentBlockTasksTaskInfo struct {
	CreateTime *int64  `json:",omitempty"`
	Status     *string `json:",omitempty"`
	TaskID     *string `json:",omitempty"`
	TaskType   *string `json:",omitempty"`
	Url        *string `json:",omitempty"`
}

type DescribeContentQuotaResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeContentQuotaResult
}

type DescribeContentQuotaResult struct {
	PreloadQuota     *int64 `json:",omitempty"`
	PreloadRemain    *int64 `json:",omitempty"`
	RefreshDirQuota  *int64 `json:",omitempty"`
	RefreshDirRemain *int64 `json:",omitempty"`
	RefreshQuota     *int64 `json:",omitempty"`
	RefreshRemain    *int64 `json:",omitempty"`
}

type DescribeContentTasksRequest struct {
	DomainName *string `json:",omitempty"`
	EndTime    *int64  `json:",omitempty"`
	PageNum    *int64  `json:",omitempty"`
	PageSize   *int64  `json:",omitempty"`
	Remark     *string `json:",omitempty"`
	StartTime  *int64  `json:",omitempty"`
	Status     *string `json:",omitempty"`
	TaskID     *string `json:",omitempty"`
	TaskType   string
	Url        *string `json:",omitempty"`
}

type DescribeContentTasksResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeContentTasksResult
}

type DescribeContentTasksResult struct {
	Data     []ContentTask
	PageNum  *int64 `json:",omitempty"`
	PageSize *int64 `json:",omitempty"`
	Total    *int64 `json:",omitempty"`
}

type DescribeEdgeNrtDataSummaryRequest struct {
	Aggregate           *string `json:",omitempty"`
	Area                *string `json:",omitempty"`
	BillingRegion       *string `json:",omitempty"`
	DisaggregateMetrics *string `json:",omitempty"`
	Domain              *string `json:",omitempty"`
	EndTime             int64
	Interval            *string `json:",omitempty"`
	IpVersion           *string `json:",omitempty"`
	Isp                 *string `json:",omitempty"`
	Metric              string
	Protocol            *string `json:",omitempty"`
	Region              *string `json:",omitempty"`
	StartTime           int64
}

type DescribeEdgeNrtDataSummaryResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeEdgeNrtDataSummaryResult
}

type DescribeEdgeNrtDataSummaryResult struct {
	Resources []NrtDataSummaryResource
}

type DescribeEdgeStatisticalDataRequest struct {
	Area      *string `json:",omitempty"`
	Domain    *string `json:",omitempty"`
	EndTime   int64
	Interval  *string `json:",omitempty"`
	IpVersion *string `json:",omitempty"`
	Metric    string
	Region    *string `json:",omitempty"`
	StartTime int64
}

type DescribeEdgeStatisticalDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeEdgeStatisticalDataResult
}

type DescribeEdgeStatisticalDataResult struct {
	Resources []EdgeStatisticalDataResource
}

type DescribeEdgeTopNrtDataRequest struct {
	Area          *string `json:",omitempty"`
	BillingRegion *string `json:",omitempty"`
	Domain        *string `json:",omitempty"`
	EndTime       int64
	Interval      *string `json:",omitempty"`
	Item          string
	Metric        string
	StartTime     int64
}

type DescribeEdgeTopNrtDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeEdgeTopNrtDataResult
}

type DescribeEdgeTopNrtDataResult struct {
	Item           *string `json:",omitempty"`
	Metric         *string `json:",omitempty"`
	Name           *string `json:",omitempty"`
	TopDataDetails []TopNrtDataDetail
}

type DescribeEdgeTopStatisticalDataRequest struct {
	Area      *string `json:",omitempty"`
	Domain    string
	EndTime   int64
	Item      string
	Metric    string
	StartTime int64
}

type DescribeEdgeTopStatisticalDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeEdgeTopStatisticalDataResult
}

type DescribeEdgeTopStatisticalDataResult struct {
	Item           *string `json:",omitempty"`
	Metric         *string `json:",omitempty"`
	Name           *string `json:",omitempty"`
	TopDataDetails []EdgeTopStatisticalDataDetail
}

type DescribeEdgeTopStatusCodeRequest struct {
	Area          *string `json:",omitempty"`
	BillingRegion *string `json:",omitempty"`
	Domain        *string `json:",omitempty"`
	EndTime       int64
	Item          string
	Metric        string
	StartTime     int64
}

type DescribeEdgeTopStatusCodeResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeEdgeTopStatusCodeResult
}

type DescribeEdgeTopStatusCodeResult struct {
	Item           *string `json:",omitempty"`
	Metric         *string `json:",omitempty"`
	Name           *string `json:",omitempty"`
	TopDataDetails []TopStatusCodeDetail
}

type DescribeIPInfoRequest struct {
	IP string
}

type DescribeIPInfoResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeIPInfoResult
}

type DescribeIPInfoResult struct {
	CdnIp    *bool   `json:",omitempty"`
	IP       *string `json:",omitempty"`
	ISP      *string `json:",omitempty"`
	Location *string `json:",omitempty"`
}

type DescribeIPListInfoRequest struct {
	IpList string
}

type DescribeIPListInfoResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           []IPInfo
}

type DescribeOriginNrtDataSummaryRequest struct {
	Aggregate           *string `json:",omitempty"`
	Area                *string `json:",omitempty"`
	BillingRegion       *string `json:",omitempty"`
	DisaggregateMetrics *string `json:",omitempty"`
	Domain              *string `json:",omitempty"`
	EndTime             int64
	Interval            *string `json:",omitempty"`
	IpVersion           *string `json:",omitempty"`
	Isp                 *string `json:",omitempty"`
	Metric              string
	Protocol            *string `json:",omitempty"`
	Region              *string `json:",omitempty"`
	StartTime           int64
}

type DescribeOriginNrtDataSummaryResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeOriginNrtDataSummaryResult
}

type DescribeOriginNrtDataSummaryResult struct {
	Resources []NrtDataSummaryResource
}

type DescribeOriginTopNrtDataRequest struct {
	Area          *string `json:",omitempty"`
	BillingRegion *string `json:",omitempty"`
	Domain        *string `json:",omitempty"`
	EndTime       int64
	Interval      *string `json:",omitempty"`
	Item          string
	Metric        string
	StartTime     int64
}

type DescribeOriginTopNrtDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeOriginTopNrtDataResult
}

type DescribeOriginTopNrtDataResult struct {
	Item           *string `json:",omitempty"`
	Metric         *string `json:",omitempty"`
	Name           *string `json:",omitempty"`
	TopDataDetails []TopNrtDataDetail
}

type DescribeOriginTopStatusCodeRequest struct {
	Area          *string `json:",omitempty"`
	BillingRegion *string `json:",omitempty"`
	Domain        *string `json:",omitempty"`
	EndTime       int64
	Item          string
	Metric        string
	StartTime     int64
}

type DescribeOriginTopStatusCodeResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeOriginTopStatusCodeResult
}

type DescribeOriginTopStatusCodeResult struct {
	Item           *string `json:",omitempty"`
	Metric         *string `json:",omitempty"`
	Name           *string `json:",omitempty"`
	TopDataDetails []TopStatusCodeDetail
}

type DomainCertDeployStatus struct {
	Domain   *string `json:",omitempty"`
	ErrorMsg *string `json:",omitempty"`
	Status   *string `json:",omitempty"`
}

type DomainCertResult struct {
	CerStatus *string `json:",omitempty"`
	Domain    *string `json:",omitempty"`
	Status    *string `json:",omitempty"`
}

type DomainLogDetail struct {
	EndTime   *int64  `json:",omitempty"`
	LogName   *string `json:",omitempty"`
	LogPath   *string `json:",omitempty"`
	LogSize   *int64  `json:",omitempty"`
	StartTime *int64  `json:",omitempty"`
}

type DomainVolcanoDetail struct {
	AreaAccessRule     *AreaAccessRule `json:",omitempty"`
	BandwidthLimit     *BandwidthLimit `json:",omitempty"`
	Cache              []CacheControlRule
	CacheKey           []CacheKeyGenerationRule
	Cname              *string             `json:",omitempty"`
	Compression        *Compression        `json:",omitempty"`
	CreateTime         *int64              `json:",omitempty"`
	CustomErrorPage    *CustomErrorPage    `json:",omitempty"`
	Domain             *string             `json:",omitempty"`
	DownloadSpeedLimit *DownloadSpeedLimit `json:",omitempty"`
	FollowRedirect     *bool               `json:",omitempty"`
	HTTPS              *HTTPS              `json:",omitempty"`
	HttpForcedRedirect *HttpForcedRedirect `json:",omitempty"`
	IPv6               *IPv6               `json:",omitempty"`
	IpAccessRule       *IpAccessRule       `json:",omitempty"`
	IpFreqLimit        *IpFreqLimit        `json:",omitempty"`
	IpSpeedLimit       *IpSpeedLimit       `json:",omitempty"`
	MethodDeniedRule   *MethodDeniedRule   `json:",omitempty"`
	NegativeCache      []NegativeCache
	Origin             []OriginRule
	OriginAccessRule   *OriginAccessRule `json:",omitempty"`
	OriginArg          []OriginArgRule
	OriginHost         *string `json:",omitempty"`
	OriginProtocol     string
	OriginRange        *bool               `json:",omitempty"`
	OriginSni          *OriginSni          `json:",omitempty"`
	Project            *string             `json:",omitempty"`
	Quic               *Quic               `json:",omitempty"`
	RedirectionRewrite *RedirectionRewrite `json:",omitempty"`
	RefererAccessRule  *RefererAccessRule  `json:",omitempty"`
	RemoteAuth         *RemoteAuth         `json:",omitempty"`
	RequestHeader      []RequestHeaderRule
	ResponseHeader     []ResponseHeaderRule
	ServiceRegion      *string              `json:",omitempty"`
	ServiceType        *string              `json:",omitempty"`
	SignedUrlAuth      *SignedUrlAuth       `json:",omitempty"`
	Status             *string              `json:",omitempty"`
	UaAccessRule       *UserAgentAccessRule `json:",omitempty"`
	UpdateTime         *int64               `json:",omitempty"`
	VideoDrag          *VideoDrag           `json:",omitempty"`
}

type DownloadSpeedLimit struct {
	DownloadSpeedLimitRules []DownloadSpeedLimitRule
	Switch                  bool
}

type DownloadSpeedLimitAction struct {
	SpeedLimitRate      int64
	SpeedLimitRateAfter *int64          `json:",omitempty"`
	SpeedLimitTime      *SpeedLimitTime `json:",omitempty"`
}

type DownloadSpeedLimitActionUpdate struct {
	SpeedLimitRate      *int64                `json:",omitempty"`
	SpeedLimitRateAfter *int64                `json:",omitempty"`
	SpeedLimitTime      *SpeedLimitTimeUpdate `json:",omitempty"`
}

type DownloadSpeedLimitRule struct {
	Condition                *Condition `json:",omitempty"`
	DownloadSpeedLimitAction DownloadSpeedLimitAction
}

type DownloadSpeedLimitRuleUpdate struct {
	Condition                *ConditionUpdate                `json:",omitempty"`
	DownloadSpeedLimitAction *DownloadSpeedLimitActionUpdate `json:",omitempty"`
}

type DownloadSpeedLimitUpdate struct {
	DownloadSpeedLimitRules []DownloadSpeedLimitRuleUpdate
	Switch                  *bool `json:",omitempty"`
}

type EdgeStatisticalDataResource struct {
	Metrics []MetricTimeStampValue
	Name    *string `json:",omitempty"`
}

type EdgeTopStatisticalDataDetail struct {
	ItemKey   *string  `json:",omitempty"`
	ItemKeyCN *string  `json:",omitempty"`
	Value     *float64 `json:",omitempty"`
}

type ErrorObj struct {
	CodeN   int64
	Code    string
	Message string
}

type ErrorPageAction struct {
	Action       *string `json:",omitempty"`
	RedirectCode string
	RedirectUrl  string
	StatusCode   string
}

type ErrorPageActionUpdate struct {
	Action       *string `json:",omitempty"`
	RedirectCode *string `json:",omitempty"`
	RedirectUrl  *string `json:",omitempty"`
	StatusCode   *string `json:",omitempty"`
}

type ErrorPageRule struct {
	ErrorPageAction ErrorPageAction
}

type ErrorPageRuleUpdate struct {
	ErrorPageAction *ErrorPageActionUpdate `json:",omitempty"`
}

type ForcedRedirect struct {
	EnableForcedRedirect bool
	StatusCode           string
}

type ForcedRedirectUpdate struct {
	EnableForcedRedirect *bool   `json:",omitempty"`
	StatusCode           *string `json:",omitempty"`
}

type HTTPS struct {
	CertInfo       *CertInfo       `json:",omitempty"`
	DisableHttp    *bool           `json:",omitempty"`
	ForcedRedirect *ForcedRedirect `json:",omitempty"`
	HTTP2          *bool           `json:",omitempty"`
	OCSP           *bool           `json:",omitempty"`
	Switch         bool
	TlsVersion     []string
}

type HTTPSUpdate struct {
	CertInfo       *CertInfoUpdate       `json:",omitempty"`
	DisableHttp    *bool                 `json:",omitempty"`
	ForcedRedirect *ForcedRedirectUpdate `json:",omitempty"`
	HTTP2          *bool                 `json:",omitempty"`
	OCSP           *bool                 `json:",omitempty"`
	Switch         *bool                 `json:",omitempty"`
	TlsVersion     []string
}

type HttpForcedRedirect struct {
	EnableForcedRedirect bool
	StatusCode           string
}

type HttpForcedRedirectUpdate struct {
	EnableForcedRedirect *bool   `json:",omitempty"`
	StatusCode           *string `json:",omitempty"`
}

type IPInfo struct {
	CdnIp    *bool   `json:",omitempty"`
	IP       *string `json:",omitempty"`
	ISP      *string `json:",omitempty"`
	Location *string `json:",omitempty"`
}

type IPv6 struct {
	Switch bool
}

type IPv6Update struct {
	Switch *bool `json:",omitempty"`
}

type IpAccessRule struct {
	Ip       []string
	RuleType string
	Switch   bool
}

type IpAccessRuleUpdate struct {
	Ip       []string
	RuleType *string `json:",omitempty"`
	Switch   *bool   `json:",omitempty"`
}

type IpFreqLimit struct {
	IpFreqLimitRules []IpFreqLimitRule
	Switch           bool
}

type IpFreqLimitAction struct {
	Action        string
	FreqLimitRate int64
	StatusCode    string
}

type IpFreqLimitActionUpdate struct {
	Action        *string `json:",omitempty"`
	FreqLimitRate *int64  `json:",omitempty"`
	StatusCode    *string `json:",omitempty"`
}

type IpFreqLimitRule struct {
	Condition         *Condition `json:",omitempty"`
	IpFreqLimitAction IpFreqLimitAction
}

type IpFreqLimitRuleUpdate struct {
	Condition         *ConditionUpdate         `json:",omitempty"`
	IpFreqLimitAction *IpFreqLimitActionUpdate `json:",omitempty"`
}

type IpFreqLimitUpdate struct {
	IpFreqLimitRules []IpFreqLimitRuleUpdate
	Switch           *bool `json:",omitempty"`
}

type IpSpeedLimit struct {
	IpSpeedLimitRules []IpSpeedLimitRule
	Switch            bool
}

type IpSpeedLimitAction struct {
	SpeedLimitRate *int64 `json:",omitempty"`
}

type IpSpeedLimitActionUpdate struct {
	SpeedLimitRate *int64 `json:",omitempty"`
}

type IpSpeedLimitRule struct {
	Condition          *Condition `json:",omitempty"`
	IpSpeedLimitAction IpSpeedLimitAction
}

type IpSpeedLimitRuleUpdate struct {
	Condition          *ConditionUpdate          `json:",omitempty"`
	IpSpeedLimitAction *IpSpeedLimitActionUpdate `json:",omitempty"`
}

type IpSpeedLimitUpdate struct {
	IpSpeedLimitRules []IpSpeedLimitRuleUpdate
	Switch            *bool `json:",omitempty"`
}

type ListCdnCertInfoRequest struct {
	CertId           *string `json:",omitempty"`
	ConfiguredDomain *string `json:",omitempty"`
	DnsName          *string `json:",omitempty"`
	PageNum          *int64  `json:",omitempty"`
	PageSize         *int64  `json:",omitempty"`
	Source           *string `json:",omitempty"`
	Status           *string `json:",omitempty"`
}

type ListCdnCertInfoResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           ListCdnCertInfoResult
}

type ListCdnCertInfoResult struct {
	CertInfo      []CertInfo
	ExpiringCount *int64 `json:",omitempty"`
	PageNum       *int64 `json:",omitempty"`
	PageSize      *int64 `json:",omitempty"`
	Total         *int64 `json:",omitempty"`
}

type ListCdnDomainDomain struct {
	BackupOrigin   []string
	Cname          *string `json:",omitempty"`
	CreateTime     *int64  `json:",omitempty"`
	Domain         *string `json:",omitempty"`
	HTTPS          *bool   `json:",omitempty"`
	IPv6           *bool   `json:",omitempty"`
	OriginProtocol *string `json:",omitempty"`
	PrimaryOrigin  []string
	ResourceTags   []ResourceTag
	ServiceRegion  *string `json:",omitempty"`
	ServiceType    *string `json:",omitempty"`
	Status         *string `json:",omitempty"`
	UpdateTime     *int64  `json:",omitempty"`
}

type ListCdnDomainsRequest struct {
	Domain         *string `json:",omitempty"`
	ExactMatch     *bool   `json:",omitempty"`
	HTTPS          *bool   `json:",omitempty"`
	IPv6           *bool   `json:",omitempty"`
	OriginProtocol *string `json:",omitempty"`
	PageNum        *int64  `json:",omitempty"`
	PageSize       *int64  `json:",omitempty"`
	PrimaryOrigin  *string `json:",omitempty"`
	Project        *string `json:",omitempty"`
	ResourceTags   []string
	ServiceRegion  *string `json:",omitempty"`
	ServiceType    *string `json:",omitempty"`
	Status         *string `json:",omitempty"`
	TagConnective  *string `json:",omitempty"`
}

type ListCdnDomainsResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           ListCdnDomainsResult
}

type ListCdnDomainsResult struct {
	Data     []ListCdnDomainDomain
	PageNum  *int64 `json:",omitempty"`
	PageSize *int64 `json:",omitempty"`
	Total    *int64 `json:",omitempty"`
}

type ListCertInfoRequest struct {
	Name          *string `json:",omitempty"`
	PageNum       *int64  `json:",omitempty"`
	PageSize      *int64  `json:",omitempty"`
	SetPagination *bool   `json:",omitempty"`
	Source        *string `json:",omitempty"`
	Status        *string `json:",omitempty"`
}

type ListCertInfoResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           ListCertInfoResult
}

type ListCertInfoResult struct {
	CertInfo      []CertInfo
	ExpiringCount *int64 `json:",omitempty"`
	PageNum       *int64 `json:",omitempty"`
	PageSize      *int64 `json:",omitempty"`
	Total         *int64 `json:",omitempty"`
}

type ListResourceTagsResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           ListResourceTagsResult
}

type ListResourceTagsResult struct {
	ResourceTags []ResourceTag
}

type MethodDeniedRule struct {
	Methods *string `json:",omitempty"`
	Switch  bool
}

type MethodDeniedRuleUpdate struct {
	Methods *string `json:",omitempty"`
	Switch  *bool   `json:",omitempty"`
}

type MetricTimeStampValue struct {
	Metric *string `json:",omitempty"`
	Values []TimeStampValue
}

type MetricTimestampValue struct {
	Metric *string `json:",omitempty"`
	Values []TimestampValue
}

type MetricValue struct {
	Metric *string  `json:",omitempty"`
	Value  *float64 `json:",omitempty"`
}

type ModuleLockConfig struct {
	CacheKeyLocked          *bool `json:",omitempty"`
	CacheLocked             *bool `json:",omitempty"`
	CompressionLocked       *bool `json:",omitempty"`
	IpAccessRuleLocked      *bool `json:",omitempty"`
	NegativeCacheLocked     *bool `json:",omitempty"`
	OriginLocked            *bool `json:",omitempty"`
	RefererAccessRuleLocked *bool `json:",omitempty"`
	RequestHeaderLocked     *bool `json:",omitempty"`
	ResponseHeaderLocked    *bool `json:",omitempty"`
	SignUrlAuthLocked       *bool `json:",omitempty"`
}

type NamePair struct {
	Code *string `json:",omitempty"`
	Name *string `json:",omitempty"`
}

type NegativeCache struct {
	Condition         *Condition `json:",omitempty"`
	NegativeCacheRule NegativeCacheAction
}

type NegativeCacheAction struct {
	Action     string
	StatusCode string
	Ttl        int64
}

type NegativeCacheActionUpdate struct {
	Action     *string `json:",omitempty"`
	StatusCode *string `json:",omitempty"`
	Ttl        *int64  `json:",omitempty"`
}

type NegativeCacheUpdate struct {
	Condition         *ConditionUpdate           `json:",omitempty"`
	NegativeCacheRule *NegativeCacheActionUpdate `json:",omitempty"`
}

type NrtDataDetails struct {
	Isp     *string `json:",omitempty"`
	Metrics []MetricTimestampValue
	Region  *string `json:",omitempty"`
}

type NrtDataResource struct {
	BillingRegion *string `json:",omitempty"`
	Metrics       []MetricTimestampValue
	Name          *string `json:",omitempty"`
}

type NrtDataSummaryResource struct {
	BillingRegion *string `json:",omitempty"`
	Metrics       []MetricValue
	Name          *string `json:",omitempty"`
}

type OriginAccessRule struct {
	AllowEmpty bool
	Origins    []string
	RuleType   string
	Switch     bool
}

type OriginAccessRuleUpdate struct {
	AllowEmpty *bool `json:",omitempty"`
	Origins    []string
	RuleType   *string `json:",omitempty"`
	Switch     *bool   `json:",omitempty"`
}

type OriginAction struct {
	OriginLines []OriginLine
}

type OriginActionUpdate struct {
	OriginLines []OriginLineUpdate
}

type OriginArgAction struct {
	OriginArgComponents []OriginArgComponents
}

type OriginArgActionUpdate struct {
	OriginArgComponents []OriginArgComponentsUpdate
}

type OriginArgComponents struct {
	Action    string
	Object    string
	Subobject string
}

type OriginArgComponentsUpdate struct {
	Action    *string `json:",omitempty"`
	Object    *string `json:",omitempty"`
	Subobject *string `json:",omitempty"`
}

type OriginArgRule struct {
	Condition       Condition
	OriginArgAction OriginArgAction
}

type OriginArgRuleUpdate struct {
	Condition       *ConditionUpdate       `json:",omitempty"`
	OriginArgAction *OriginArgActionUpdate `json:",omitempty"`
}

type OriginLine struct {
	Address             string
	HttpPort            string
	HttpsPort           string
	InstanceType        string
	OriginHost          *string `json:",omitempty"`
	OriginType          string
	PrivateBucketAccess *bool `json:",omitempty"`
	Weight              string
}

type OriginLineUpdate struct {
	Address             *string `json:",omitempty"`
	HttpPort            *string `json:",omitempty"`
	HttpsPort           *string `json:",omitempty"`
	InstanceType        *string `json:",omitempty"`
	OriginHost          *string `json:",omitempty"`
	OriginType          *string `json:",omitempty"`
	PrivateBucketAccess *bool   `json:",omitempty"`
	Weight              *string `json:",omitempty"`
}

type OriginRule struct {
	Condition    *Condition `json:",omitempty"`
	OriginAction OriginAction
}

type OriginRuleUpdate struct {
	Condition    *ConditionUpdate    `json:",omitempty"`
	OriginAction *OriginActionUpdate `json:",omitempty"`
}

type OriginSni struct {
	SniDomain *string `json:",omitempty"`
	Switch    *bool   `json:",omitempty"`
}

type OriginSniUpdate struct {
	SniDomain *string `json:",omitempty"`
	Switch    *bool   `json:",omitempty"`
}

type QueryStringComponents struct {
	Action *string `json:",omitempty"`
	Value  *string `json:",omitempty"`
}

type QueryStringComponentsUpdate struct {
	Action *string `json:",omitempty"`
	Value  *string `json:",omitempty"`
}

type QueryStringInstance struct {
	Action    *string `json:",omitempty"`
	Key       *string `json:",omitempty"`
	Value     *string `json:",omitempty"`
	ValueType *string `json:",omitempty"`
}

type QueryStringInstanceUpdate struct {
	Action    *string `json:",omitempty"`
	Key       *string `json:",omitempty"`
	Value     *string `json:",omitempty"`
	ValueType *string `json:",omitempty"`
}

type QueryStringRule struct {
	QueryStringComponents *QueryStringComponents `json:",omitempty"`
	QueryStringInstances  []QueryStringInstance
}

type QueryStringRuleUpdate struct {
	QueryStringComponents *QueryStringComponentsUpdate `json:",omitempty"`
	QueryStringInstances  []QueryStringInstanceUpdate
}

type Quic struct {
	Switch bool
}

type QuicUpdate struct {
	Switch *bool `json:",omitempty"`
}

type RedirectionAction struct {
	RedirectCode          string
	SourcePath            string
	TargetHost            string
	TargetPath            string
	TargetProtocol        string
	TargetQueryComponents TargetQueryComponents
}

type RedirectionActionUpdate struct {
	RedirectCode          *string                      `json:",omitempty"`
	SourcePath            *string                      `json:",omitempty"`
	TargetHost            *string                      `json:",omitempty"`
	TargetPath            *string                      `json:",omitempty"`
	TargetProtocol        *string                      `json:",omitempty"`
	TargetQueryComponents *TargetQueryComponentsUpdate `json:",omitempty"`
}

type RedirectionRewrite struct {
	RedirectionRule []RedirectionRule
	Switch          bool
}

type RedirectionRewriteUpdate struct {
	RedirectionRule []RedirectionRuleUpdate
	Switch          *bool `json:",omitempty"`
}

type RedirectionRule struct {
	RedirectionAction RedirectionAction
}

type RedirectionRuleUpdate struct {
	RedirectionAction *RedirectionActionUpdate `json:",omitempty"`
}

type RefererAccessRule struct {
	AllowEmpty bool
	Referers   []string
	RuleType   string
	Switch     bool
}

type RefererAccessRuleUpdate struct {
	AllowEmpty *bool `json:",omitempty"`
	Referers   []string
	RuleType   *string `json:",omitempty"`
	Switch     *bool   `json:",omitempty"`
}

type RemoteAuth struct {
	RemoteAuthRules []RemoteAuthRule
	Switch          bool
}

type RemoteAuthRule struct {
	Condition            Condition
	RemoteAuthRuleAction RemoteAuthRuleAction
}

type RemoteAuthRuleAction struct {
	AuthModeConfig     AuthModeConfig
	AuthResponseConfig AuthResponseConfig
	QueryStringRules   QueryStringRule
	RequestBodyRules   string
	RequestHeaderRules AuthRequestHeaderRule
}

type RemoteAuthRuleActionUpdate struct {
	AuthModeConfig     *AuthModeConfigUpdate        `json:",omitempty"`
	AuthResponseConfig *AuthResponseConfigUpdate    `json:",omitempty"`
	QueryStringRules   *QueryStringRuleUpdate       `json:",omitempty"`
	RequestBodyRules   *string                      `json:",omitempty"`
	RequestHeaderRules *AuthRequestHeaderRuleUpdate `json:",omitempty"`
}

type RemoteAuthRuleUpdate struct {
	Condition            *ConditionUpdate            `json:",omitempty"`
	RemoteAuthRuleAction *RemoteAuthRuleActionUpdate `json:",omitempty"`
}

type RemoteAuthUpdate struct {
	RemoteAuthRules []RemoteAuthRuleUpdate
	Switch          *bool `json:",omitempty"`
}

type RequestHeaderAction struct {
	RequestHeaderInstances []RequestHeaderInstance
}

type RequestHeaderActionUpdate struct {
	RequestHeaderInstances []RequestHeaderInstanceUpdate
}

type RequestHeaderComponent struct {
	Action string
	Value  *string `json:",omitempty"`
}

type RequestHeaderComponentUpdate struct {
	Action *string `json:",omitempty"`
	Value  *string `json:",omitempty"`
}

type RequestHeaderInstance struct {
	Action    string
	Key       string
	Value     *string `json:",omitempty"`
	ValueType string
}

type RequestHeaderInstanceUpdate struct {
	Action    *string `json:",omitempty"`
	Key       *string `json:",omitempty"`
	Value     *string `json:",omitempty"`
	ValueType *string `json:",omitempty"`
}

type RequestHeaderRule struct {
	Condition           *Condition `json:",omitempty"`
	RequestHeaderAction RequestHeaderAction
}

type RequestHeaderRuleUpdate struct {
	Condition           *ConditionUpdate           `json:",omitempty"`
	RequestHeaderAction *RequestHeaderActionUpdate `json:",omitempty"`
}

type ResourceTag struct {
	Key   *string `json:",omitempty"`
	Value *string `json:",omitempty"`
}

type ResponseAction struct {
	StatusCode string
}

type ResponseActionUpdate struct {
	StatusCode *string `json:",omitempty"`
}

type ResponseHeaderAction struct {
	ResponseHeaderInstances []ResponseHeaderInstance
}

type ResponseHeaderActionUpdate struct {
	ResponseHeaderInstances []ResponseHeaderInstanceUpdate
}

type ResponseHeaderInstance struct {
	Action    string
	Key       string
	Value     *string `json:",omitempty"`
	ValueType string
}

type ResponseHeaderInstanceUpdate struct {
	Action    *string `json:",omitempty"`
	Key       *string `json:",omitempty"`
	Value     *string `json:",omitempty"`
	ValueType *string `json:",omitempty"`
}

type ResponseHeaderRule struct {
	Condition            *Condition `json:",omitempty"`
	ResponseHeaderAction ResponseHeaderAction
}

type ResponseHeaderRuleUpdate struct {
	Condition            *ConditionUpdate            `json:",omitempty"`
	ResponseHeaderAction *ResponseHeaderActionUpdate `json:",omitempty"`
}

type ResponseMetadata struct {
	RequestId string
	Service   *string   `json:",omitempty"`
	Region    *string   `json:",omitempty"`
	Action    *string   `json:",omitempty"`
	Version   *string   `json:",omitempty"`
	Error     *ErrorObj `json:",omitempty"`
}

type SignedUrlAuth struct {
	SignedUrlAuthRules []SignedUrlAuthRule
	Switch             bool
}

type SignedUrlAuthAction struct {
	BackupSecretKey *string `json:",omitempty"`
	Duration        int64
	MasterSecretKey string
	RewriteM3u8     *bool   `json:",omitempty"`
	SignName        *string `json:",omitempty"`
	SignatureRule   []string
	TimeFormat      *string `json:",omitempty"`
	TimeName        *string `json:",omitempty"`
	URLAuthType     string
}

type SignedUrlAuthActionUpdate struct {
	BackupSecretKey *string `json:",omitempty"`
	Duration        *int64  `json:",omitempty"`
	MasterSecretKey *string `json:",omitempty"`
	RewriteM3u8     *bool   `json:",omitempty"`
	SignName        *string `json:",omitempty"`
	SignatureRule   []string
	TimeFormat      *string `json:",omitempty"`
	TimeName        *string `json:",omitempty"`
	URLAuthType     *string `json:",omitempty"`
}

type SignedUrlAuthRule struct {
	Condition           Condition
	SignedUrlAuthAction SignedUrlAuthAction
}

type SignedUrlAuthRuleUpdate struct {
	Condition           *ConditionUpdate           `json:",omitempty"`
	SignedUrlAuthAction *SignedUrlAuthActionUpdate `json:",omitempty"`
}

type SignedUrlAuthUpdate struct {
	SignedUrlAuthRules []SignedUrlAuthRuleUpdate
	Switch             *bool `json:",omitempty"`
}

type SpeedLimitTime struct {
	BeginTime *string `json:",omitempty"`
	DayWeek   string
	EndTime   *string `json:",omitempty"`
}

type SpeedLimitTimeUpdate struct {
	BeginTime *string `json:",omitempty"`
	DayWeek   *string `json:",omitempty"`
	EndTime   *string `json:",omitempty"`
}

type StartCdnDomainRequest struct {
	Domain string
}

type StartCdnDomainResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}

type StatusCodeAction struct {
	DefaultAction string
	FailCode      string
	SuccessCode   string
}

type StatusCodeActionUpdate struct {
	DefaultAction *string `json:",omitempty"`
	FailCode      *string `json:",omitempty"`
	SuccessCode   *string `json:",omitempty"`
}

type StopCdnDomainRequest struct {
	Domain string
}

type StopCdnDomainResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}

type SubmitBlockTaskRequest struct {
	Persistent *bool `json:",omitempty"`
	Urls       string
}

type SubmitBlockTaskResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           SubmitBlockTaskResult
}

type SubmitBlockTaskResult struct {
	TaskID *string `json:",omitempty"`
}

type SubmitPreloadTaskRequest struct {
	ConcurrentLimit *int64 `json:",omitempty"`
	Urls            string
}

type SubmitPreloadTaskResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           SubmitPreloadTaskResult
}

type SubmitPreloadTaskResult struct {
	TaskID string
}

type SubmitRefreshTaskRequest struct {
	Type *string `json:",omitempty"`
	Urls string
}

type SubmitRefreshTaskResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           SubmitRefreshTaskResult
}

type SubmitRefreshTaskResult struct {
	TaskID string
}

type SubmitUnblockTaskRequest struct {
	Persistent *bool `json:",omitempty"`
	Urls       string
}

type SubmitUnblockTaskResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           SubmitUnblockTaskResult
}

type SubmitUnblockTaskResult struct {
	TaskID *string `json:",omitempty"`
}

type TargetQueryComponents struct {
	Action string
	Value  string
}

type TargetQueryComponentsUpdate struct {
	Action *string `json:",omitempty"`
	Value  *string `json:",omitempty"`
}

type TimeOutAction struct {
	Action string
	Time   int64
}

type TimeOutActionUpdate struct {
	Action *string `json:",omitempty"`
	Time   *int64  `json:",omitempty"`
}

type TimeStampValue struct {
	TimeStamp *int64   `json:",omitempty"`
	Value     *float64 `json:",omitempty"`
}

type TimestampValue struct {
	Timestamp *int64   `json:",omitempty"`
	Value     *float64 `json:",omitempty"`
}

type TopDataDetail struct {
	Status2xx      *int64   `json:"2xx,omitempty"`
	Status2xxRatio *float64 `json:"2xxRatio,omitempty"`
	Status3xx      *int64   `json:"3xx,omitempty"`
	Status3xxRatio *float64 `json:"3xxRatio,omitempty"`
	Status4xx      *int64   `json:"4xx,omitempty"`
	Status4xxRatio *float64 `json:"4xxRatio,omitempty"`
	Status5xx      *int64   `json:"5xx,omitempty"`
	Status5xxRatio *float64 `json:"5xxRatio,omitempty"`
	Bandwidth      *float64 `json:",omitempty"`
	BandwidthRatio *float64 `json:",omitempty"`
	ClientIP       *int64   `json:",omitempty"`
	ClientIPRatio  *float64 `json:",omitempty"`
	Flux           *int64   `json:",omitempty"`
	FluxRatio      *float64 `json:",omitempty"`
	Item           *string  `json:",omitempty"`
	PV             *int64   `json:",omitempty"`
	PVRatio        *float64 `json:",omitempty"`
}

type TopInstanceDetail struct {
	BillingCode      *string `json:",omitempty"`
	BillingCycle     *string `json:",omitempty"`
	BillingData      *string `json:",omitempty"`
	BillingDesc      *string `json:",omitempty"`
	CreateTime       *string `json:",omitempty"`
	InstanceCategory *string `json:",omitempty"`
	InstanceType     *string `json:",omitempty"`
	MetricType       *string `json:",omitempty"`
	ServiceRegion    *string `json:",omitempty"`
	StartTime        *string `json:",omitempty"`
	Status           *string `json:",omitempty"`
}

type TopNrtDataDetail struct {
	Bandwidth         *float64 `json:",omitempty"`
	BandwidthPeakTime *int64   `json:",omitempty"`
	Flux              *float64 `json:",omitempty"`
	FluxRatio         *float64 `json:",omitempty"`
	ItemKey           *string  `json:",omitempty"`
	ItemKeyCN         *string  `json:",omitempty"`
	PV                *float64 `json:",omitempty"`
	PVRatio           *float64 `json:",omitempty"`
	Quic              *int64   `json:",omitempty"`
}

type TopStatusCodeDetail struct {
	Status2xx      *float64 `json:"2xx,omitempty"`
	Status2xxRatio *float64 `json:"2xxRatio,omitempty"`
	Status3xx      *float64 `json:"3xx,omitempty"`
	Status3xxRatio *float64 `json:"3xxRatio,omitempty"`
	Status4xx      *float64 `json:"4xx,omitempty"`
	Status4xxRatio *float64 `json:"4xxRatio,omitempty"`
	Status5xx      *float64 `json:"5xx,omitempty"`
	Status5xxRatio *float64 `json:"5xxRatio,omitempty"`
	ItemKey        *string  `json:",omitempty"`
}

type UpdateCdnConfigRequest struct {
	AreaAccessRule     *AreaAccessRuleUpdate `json:",omitempty"`
	BandwidthLimit     *BandwidthLimitUpdate `json:",omitempty"`
	Cache              []CacheControlRuleUpdate
	CacheKey           []CacheKeyGenerationRuleUpdate
	Compression        *CompressionUpdate        `json:",omitempty"`
	CustomErrorPage    *CustomErrorPageUpdate    `json:",omitempty"`
	Domain             *string                   `json:",omitempty"`
	DownloadSpeedLimit *DownloadSpeedLimitUpdate `json:",omitempty"`
	FollowRedirect     *bool                     `json:",omitempty"`
	HTTPS              *HTTPSUpdate              `json:",omitempty"`
	HttpForcedRedirect *HttpForcedRedirectUpdate `json:",omitempty"`
	IPv6               *IPv6Update               `json:",omitempty"`
	IpAccessRule       *IpAccessRuleUpdate       `json:",omitempty"`
	IpFreqLimit        *IpFreqLimitUpdate        `json:",omitempty"`
	IpSpeedLimit       *IpSpeedLimitUpdate       `json:",omitempty"`
	MethodDeniedRule   *MethodDeniedRuleUpdate   `json:",omitempty"`
	NegativeCache      []NegativeCacheUpdate
	Origin             []OriginRuleUpdate
	OriginAccessRule   *OriginAccessRuleUpdate `json:",omitempty"`
	OriginArg          []OriginArgRuleUpdate
	OriginHost         *string                   `json:",omitempty"`
	OriginProtocol     *string                   `json:",omitempty"`
	OriginRange        *bool                     `json:",omitempty"`
	OriginSni          *OriginSniUpdate          `json:",omitempty"`
	Quic               *QuicUpdate               `json:",omitempty"`
	RedirectionRewrite *RedirectionRewriteUpdate `json:",omitempty"`
	RefererAccessRule  *RefererAccessRuleUpdate  `json:",omitempty"`
	RemoteAuth         *RemoteAuthUpdate         `json:",omitempty"`
	RequestHeader      []RequestHeaderRuleUpdate
	ResponseHeader     []ResponseHeaderRuleUpdate
	ServiceRegion      *string                    `json:",omitempty"`
	ServiceType        *string                    `json:",omitempty"`
	SignedUrlAuth      *SignedUrlAuthUpdate       `json:",omitempty"`
	UaAccessRule       *UserAgentAccessRuleUpdate `json:",omitempty"`
	VideoDrag          *VideoDragUpdate           `json:",omitempty"`
}

type UpdateCdnConfigResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}

type UpdateResourceTagsRequest struct {
	ResourceTags []ResourceTag
	Resources    []string
}

type UpdateResourceTagsResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}

type UserAgentAccessRule struct {
	AllowEmpty bool
	RuleType   string
	Switch     bool
	UserAgent  []string
}

type UserAgentAccessRuleUpdate struct {
	AllowEmpty *bool   `json:",omitempty"`
	RuleType   *string `json:",omitempty"`
	Switch     *bool   `json:",omitempty"`
	UserAgent  []string
}

type VideoDrag struct {
	Switch bool
}

type VideoDragUpdate struct {
	Switch *bool `json:",omitempty"`
}
