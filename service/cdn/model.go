package cdn

type AccountingData struct {
	Metric string
	Values []DataPoint
}

type AccountingDataDetail struct {
	BillingRegion string
	Metrics       []AccountingData
	Name          string
}

type AccountingSummary struct {
	BillingCode string
	Name        string
	TimeStamp   int64
	Value       float64
	Values      []DataPoint
}

type AddCdnCertInfo struct {
	Desc *string `json:",omitempty"`
}

type AddCdnCertificateRequest struct {
	CertInfo      *AddCdnCertInfo `json:",omitempty"`
	Certificate   Certificate
	CloseSigCheck *bool   `json:",omitempty"`
	Source        *string `json:",omitempty"`
}

type AddCdnCertificateResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           string
}

type AddCdnDomainRequest struct {
	AreaAccessRule      *AreaAccessRule `json:",omitempty"`
	BandwidthLimit      *BandwidthLimit `json:",omitempty"`
	BrowserCache        []BrowserCacheControlRule
	Cache               []CacheControlRule
	CacheHost           *CacheHost `json:",omitempty"`
	CacheKey            []CacheKeyRule
	Compression         *Compression         `json:",omitempty"`
	CustomErrorPage     *CustomErrorPage     `json:",omitempty"`
	CustomizeAccessRule *CustomizeAccessRule `json:",omitempty"`
	Domain              string
	DownloadSpeedLimit  *DownloadSpeedLimit `json:",omitempty"`
	FollowRedirect      *bool               `json:",omitempty"`
	HTTPS               *HTTPS              `json:",omitempty"`
	HeaderLogging       *HeaderLog          `json:",omitempty"`
	HttpForcedRedirect  *HttpForcedRedirect `json:",omitempty"`
	IPv6                *IPv6               `json:",omitempty"`
	IpAccessRule        *IpAccessRule       `json:",omitempty"`
	IpFreqLimit         *IpFreqLimit        `json:",omitempty"`
	IpSpeedLimit        *IpSpeedLimit       `json:",omitempty"`
	MethodDeniedRule    *MethodDeniedRule   `json:",omitempty"`
	NegativeCache       []NegativeCache
	Origin              []OriginRule
	OriginAccessRule    *OriginAccessRule `json:",omitempty"`
	OriginArg           []OriginArgRule
	OriginHost          *string `json:",omitempty"`
	OriginIPv6          *string `json:",omitempty"`
	OriginProtocol      string
	OriginRange         *bool               `json:",omitempty"`
	OriginRewrite       *OriginRewrite      `json:",omitempty"`
	OriginSni           *OriginSni          `json:",omitempty"`
	PageOptimization    *PageOptimization   `json:",omitempty"`
	Project             *string             `json:",omitempty"`
	Quic                *Quic               `json:",omitempty"`
	RedirectionRewrite  *RedirectionRewrite `json:",omitempty"`
	RefererAccessRule   *RefererAccessRule  `json:",omitempty"`
	RemoteAuth          *RemoteAuth         `json:",omitempty"`
	RequestBlockRule    *RequestBlockRule   `json:",omitempty"`
	RequestHeader       []RequestHeaderRule
	ResourceTags        []ResourceTag
	ResponseHeader      []ResponseHeaderRule
	ServiceRegion       *string              `json:",omitempty"`
	ServiceType         *string              `json:",omitempty"`
	SignedUrlAuth       *SignedUrlAuth       `json:",omitempty"`
	Sparrow             *Sparrow             `json:",omitempty"`
	Timeout             *Timeout             `json:",omitempty"`
	UaAccessRule        *UserAgentAccessRule `json:",omitempty"`
	UrlNormalize        *URLNormalize        `json:",omitempty"`
	VideoDrag           *VideoDrag           `json:",omitempty"`
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

type AddTemplateDomainRequest struct {
	CertId            *string `json:",omitempty"`
	CipherTemplateId  *string `json:",omitempty"`
	Domain            string
	HTTPSSwitch       *string `json:",omitempty"`
	Project           *string `json:",omitempty"`
	ServiceRegion     *string `json:",omitempty"`
	ServiceTemplateId string
}

type AddTemplateDomainResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}

type AreaAccessRule struct {
	Area     []string
	RuleType *string `json:",omitempty"`
	Switch   *bool   `json:",omitempty"`
}

type AuthCacheAction struct {
	Action     *string `json:",omitempty"`
	CacheKey   []string
	IgnoreCase *bool  `json:",omitempty"`
	Ttl        *int64 `json:",omitempty"`
}

type AuthModeConfig struct {
	BackupRemoteAddr *string `json:",omitempty"`
	MasterRemoteAddr *string `json:",omitempty"`
	PathType         *string `json:",omitempty"`
	PathValue        *string `json:",omitempty"`
	RequestMethod    *string `json:",omitempty"`
}

type AuthRequestHeaderRule struct {
	RequestHeaderComponents *RequestHeaderComponent `json:",omitempty"`
	RequestHeaderInstances  []RequestHeaderInstance
	RequestHost             *string `json:",omitempty"`
}

type AuthResponseConfig struct {
	CacheAction      *AuthCacheAction   `json:",omitempty"`
	ResponseAction   *ResponseAction    `json:",omitempty"`
	StatusCodeAction *StatusCodeAction  `json:",omitempty"`
	TimeOutAction    *AuthTimeoutAction `json:",omitempty"`
}

type AuthTimeoutAction struct {
	Action *string `json:",omitempty"`
	Time   *int64  `json:",omitempty"`
}

type BandwidthLimit struct {
	BandwidthLimitRule *BandwidthLimitRule `json:",omitempty"`
	Switch             *bool               `json:",omitempty"`
}

type BandwidthLimitAction struct {
	BandwidthThreshold *int64  `json:",omitempty"`
	LimitType          *string `json:",omitempty"`
	SpeedLimitRate     *int64  `json:",omitempty"`
	SpeedLimitRateMax  *int64  `json:",omitempty"`
}

type BandwidthLimitRule struct {
	BandwidthLimitAction *BandwidthLimitAction `json:",omitempty"`
	Condition            *Condition            `json:",omitempty"`
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
	DeployResult []CertDeployStatus
}

type BlockAction struct {
	Action      *string `json:",omitempty"`
	ErrorPage   *string `json:",omitempty"`
	RedirectUrl *string `json:",omitempty"`
	StatusCode  *string `json:",omitempty"`
}

type BlockRule struct {
	BlockAction *BlockAction `json:",omitempty"`
	Condition   *Condition   `json:",omitempty"`
	RuleName    *string      `json:",omitempty"`
}

type BlockTaskInfo struct {
	BlockReason string
	CreateTime  int64
	Status      string
	TaskID      string
	TaskType    string
	Url         string
}

type BoundDomain struct {
	BoundTime int64
	Domain    string
}

type BrowserCacheControlRule struct {
	CacheAction *CacheAction `json:",omitempty"`
	Condition   *Condition   `json:",omitempty"`
}

type CacheAction struct {
	Action        *string `json:",omitempty"`
	DefaultPolicy *string `json:",omitempty"`
	IgnoreCase    *bool   `json:",omitempty"`
	Ttl           *int64  `json:",omitempty"`
}

type CacheControlRule struct {
	CacheAction *CacheAction `json:",omitempty"`
	Condition   *Condition   `json:",omitempty"`
}

type CacheHost struct {
	CacheHostRule []CacheHostRule
	Switch        *bool `json:",omitempty"`
}

type CacheHostAction struct {
	CacheHost *string `json:",omitempty"`
}

type CacheHostRule struct {
	CacheHostAction *CacheHostAction `json:",omitempty"`
	Condition       *Condition       `json:",omitempty"`
}

type CacheKeyAction struct {
	CacheKeyComponents []CacheKeyComponent
}

type CacheKeyComponent struct {
	Action     *string `json:",omitempty"`
	IgnoreCase *bool   `json:",omitempty"`
	Object     *string `json:",omitempty"`
	Subobject  *string `json:",omitempty"`
}

type CacheKeyRule struct {
	CacheKeyAction *CacheKeyAction `json:",omitempty"`
	Condition      *Condition      `json:",omitempty"`
}

type CdnTemplate struct {
	Exception  bool
	TemplateId string
	Title      string
	Type       string
}

type CdnTemplateDomain struct {
	CertInfo      TemplateCertInfo
	Cname         string
	Domain        string
	HTTPSSwitch   string
	LockStatus    string
	Project       string
	Remark        string
	ServiceRegion string
	Status        string
	Templates     []CdnTemplate
	WAFStatus     string
}

type CertDeployStatus struct {
	Domain   string
	ErrorMsg string
	Status   string
}

type CertFingerprint struct {
	Sha1   string
	Sha256 string
}

type CertInfo struct {
	CertId        *string      `json:",omitempty"`
	CertName      *string      `json:",omitempty"`
	Certificate   *Certificate `json:",omitempty"`
	Desc          *string      `json:",omitempty"`
	EffectiveTime *int64       `json:",omitempty"`
	ExpireTime    *int64       `json:",omitempty"`
	Source        *string      `json:",omitempty"`
}

type Certificate struct {
	Certificate *string `json:",omitempty"`
	PrivateKey  *string `json:",omitempty"`
}

type CommonGlobalConfig struct {
	ConfigName *string `json:",omitempty"`
}

type CommonReferType struct {
	IgnoreCase   *bool `json:",omitempty"`
	IgnoreScheme *bool `json:",omitempty"`
	Referers     []string
}

type Compression struct {
	CompressionRules []CompressionRule
	Switch           *bool `json:",omitempty"`
}

type CompressionAction struct {
	CompressionFormat *string `json:",omitempty"`
	CompressionTarget *string `json:",omitempty"`
	CompressionType   []string
	MinFileSizeKB     *int64 `json:",omitempty"`
}

type CompressionRule struct {
	CompressionAction *CompressionAction `json:",omitempty"`
	Condition         *Condition         `json:",omitempty"`
}

type Condition struct {
	ConditionRule []ConditionRule
	Connective    *string `json:",omitempty"`
}

type ConditionRule struct {
	Name     *string `json:",omitempty"`
	Object   *string `json:",omitempty"`
	Operator *string `json:",omitempty"`
	Type     *string `json:",omitempty"`
	Value    *string `json:",omitempty"`
}

type ConfiguredDomain struct {
	Domain string
	Type   string
}

type ContentTask struct {
	Area          string
	CreateTime    int64
	Delete        bool
	Isp           string
	Layer         string
	Process       string
	RefreshPrefix bool
	Region        string
	Remark        string
	Status        string
	SubArea       string
	TaskID        string
	TaskType      string
	Url           string
}

type CreateCipherTemplateRequest struct {
	HTTPS              *HTTPSCommon        `json:",omitempty"`
	HttpForcedRedirect *HttpForcedRedirect `json:",omitempty"`
	Message            *string             `json:",omitempty"`
	Project            *string             `json:",omitempty"`
	Quic               *Quic               `json:",omitempty"`
	Title              string
}

type CreateCipherTemplateResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           CreateCipherTemplateResult
}

type CreateCipherTemplateResult struct {
	TemplateId string
}

type CreateServiceTemplateRequest struct {
	AreaAccessRule      *AreaAccessRule `json:",omitempty"`
	BandwidthLimit      *BandwidthLimit `json:",omitempty"`
	BrowserCache        []BrowserCacheControlRule
	Cache               []CacheControlRule
	CacheHost           *CacheHost `json:",omitempty"`
	CacheKey            []CacheKeyRule
	Compression         *Compression         `json:",omitempty"`
	CustomErrorPage     *CustomErrorPage     `json:",omitempty"`
	CustomizeAccessRule *CustomizeAccessRule `json:",omitempty"`
	DownloadSpeedLimit  *DownloadSpeedLimit  `json:",omitempty"`
	FollowRedirect      *bool                `json:",omitempty"`
	IPv6                *IPv6                `json:",omitempty"`
	IpAccessRule        *IpAccessRule        `json:",omitempty"`
	IpFreqLimit         *IpFreqLimit         `json:",omitempty"`
	IpSpeedLimit        *IpSpeedLimit        `json:",omitempty"`
	Message             *string              `json:",omitempty"`
	MethodDeniedRule    *MethodDeniedRule    `json:",omitempty"`
	NegativeCache       []NegativeCache
	Origin              []OriginRule
	OriginAccessRule    *OriginAccessRule `json:",omitempty"`
	OriginArg           []OriginArgRule
	OriginHost          *string `json:",omitempty"`
	OriginIPv6          *string `json:",omitempty"`
	OriginProtocol      string
	OriginRange         *bool               `json:",omitempty"`
	OriginRewrite       *OriginRewrite      `json:",omitempty"`
	OriginSni           *OriginSni          `json:",omitempty"`
	PageOptimization    *PageOptimization   `json:",omitempty"`
	Project             *string             `json:",omitempty"`
	RedirectionRewrite  *RedirectionRewrite `json:",omitempty"`
	RefererAccessRule   *RefererAccessRule  `json:",omitempty"`
	RemoteAuth          *RemoteAuth         `json:",omitempty"`
	RequestHeader       []RequestHeaderRule
	ResponseHeader      []ResponseHeaderRule
	SignedUrlAuth       *SignedUrlAuth `json:",omitempty"`
	Timeout             *Timeout       `json:",omitempty"`
	Title               string
	UaAccessRule        *UserAgentAccessRule `json:",omitempty"`
	VideoDrag           *VideoDrag           `json:",omitempty"`
}

type CreateServiceTemplateResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           CreateServiceTemplateResult
}

type CreateServiceTemplateResult struct {
	TemplateId string
}

type CustomErrorPage struct {
	ErrorPageRule []ErrorPageRule
	Switch        *bool `json:",omitempty"`
}

type CustomSignedUrlParam struct {
	ParamType     *string       `json:",omitempty"`
	RequestHeader *string       `json:",omitempty"`
	SupContent    *string       `json:",omitempty"`
	UriParamSup   *UriParamSup  `json:",omitempty"`
	UrlParam      *ParamCapRule `json:",omitempty"`
}

type CustomVariableInstance struct {
	Operator *string `json:",omitempty"`
	Type     *string `json:",omitempty"`
	Value    *string `json:",omitempty"`
}

type CustomVariableRules struct {
	CustomVariableInstances []CustomVariableInstance
}

type CustomizeAccessAction struct {
	AllowEmpty    *bool `json:",omitempty"`
	ListRules     []string
	RequestHeader *string `json:",omitempty"`
	RuleType      *string `json:",omitempty"`
}

type CustomizeAccessRule struct {
	CustomizeInstances []CustomizeInstance
	Switch             *bool `json:",omitempty"`
}

type CustomizeInstance struct {
	CustomizeRule *CustomizeRule `json:",omitempty"`
}

type CustomizeRule struct {
	AccessAction *CustomizeAccessAction `json:",omitempty"`
	Condition    *Condition             `json:",omitempty"`
}

type DataPoint struct {
	TimeStamp int64
	Value     float64
}

type DeleteCdnCertificateRequest struct {
	CertId string
}

type DeleteCdnCertificateResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
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

type DeleteTemplateRequest struct {
	TemplateId string
}

type DeleteTemplateResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}

type DescribeAccountingDataRequest struct {
	Aggregate        *string `json:",omitempty"`
	BillingRegion    *string `json:",omitempty"`
	Domain           *string `json:",omitempty"`
	EndTime          int64
	Interval         *int64 `json:",omitempty"`
	InverseDomain    *bool  `json:",omitempty"`
	IsWildcardDomain *bool  `json:",omitempty"`
	Metric           string
	Project          *string `json:",omitempty"`
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

type DescribeAccountingSummaryRequest struct {
	Aggregate     *string `json:",omitempty"`
	BillingCode   *string `json:",omitempty"`
	BillingRegion *string `json:",omitempty"`
	Domain        string
	EndTime       int64
	FreeTime      *string `json:",omitempty"`
	InverseDomain *bool   `json:",omitempty"`
	Project       *string `json:",omitempty"`
	StartTime     int64
	TimeZone      *string `json:",omitempty"`
}

type DescribeAccountingSummaryResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeAccountingSummaryResult
}

type DescribeAccountingSummaryResult struct {
	Resources []AccountingSummary
}

type DescribeCdnAccessLogRequest struct {
	Domain        string
	EndTime       int64
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
	Domain           string
	DomainLogDetails []DomainLogDetail
	PageNum          int64
	PageSize         int64
	TotalCount       int64
}

type DescribeCdnConfigRequest struct {
	Domain string
}

type DescribeCdnConfigResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeCdnConfigResult
}

type DescribeCdnConfigResult struct {
	DomainConfig  DomainConfig
	FeatureConfig FeatureConfig
}

type DescribeCdnDataDetailRequest struct {
	Domain    string
	EndTime   int64
	Interval  *string `json:",omitempty"`
	IpVersion *string `json:",omitempty"`
	Metric    string
	Protocol  *string `json:",omitempty"`
	StartTime int64
	TimeZone  *string `json:",omitempty"`
}

type DescribeCdnDataDetailResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeCdnDataDetailResult
}

type DescribeCdnDataDetailResult struct {
	DataDetails []NrtDataDetails
	Name        string
}

type DescribeCdnDataRequest struct {
	Aggregate           *string `json:",omitempty"`
	Area                *string `json:",omitempty"`
	BillingRegion       *string `json:",omitempty"`
	DisaggregateMetrics *string `json:",omitempty"`
	Domain              *string `json:",omitempty"`
	EndTime             int64
	Interval            *string `json:",omitempty"`
	InverseDomain       *bool   `json:",omitempty"`
	IpVersion           *string `json:",omitempty"`
	IsWildcardDomain    *bool   `json:",omitempty"`
	Isp                 *string `json:",omitempty"`
	Metric              string
	Project             *string `json:",omitempty"`
	Protocol            *string `json:",omitempty"`
	Region              *string `json:",omitempty"`
	StartTime           int64
	TimeZone            *string `json:",omitempty"`
}

type DescribeCdnDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeCdnDataResult
}

type DescribeCdnDataResult struct {
	Resources []NrtDataResource
}

type DescribeCdnOriginDataRequest struct {
	Aggregate           *string `json:",omitempty"`
	BillingRegion       *string `json:",omitempty"`
	DisaggregateMetrics *string `json:",omitempty"`
	Domain              *string `json:",omitempty"`
	EndTime             int64
	Interval            *string `json:",omitempty"`
	InverseDomain       *bool   `json:",omitempty"`
	IsWildcardDomain    *bool   `json:",omitempty"`
	Metric              string
	Project             *string `json:",omitempty"`
	StartTime           int64
	TimeZone            *string `json:",omitempty"`
}

type DescribeCdnOriginDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeCdnOriginDataResult
}

type DescribeCdnOriginDataResult struct {
	Resources []NrtDataResource
}

type DescribeCdnRegionAndIspRequest struct {
	Area    *string `json:",omitempty"`
	Feature *string `json:",omitempty"`
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
	CertNotConfig       []DomainStatus
	OtherCertConfig     []DomainCertStatus
	SpecifiedCertConfig []DomainCertStatus
}

type DescribeCipherTemplateRequest struct {
	TemplateId string
}

type DescribeCipherTemplateResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeCipherTemplateResult
}

type DescribeCipherTemplateResult struct {
	BoundDomains       []BoundDomain
	CreateTime         int64
	Exception          bool
	HTTPS              HTTPSCommon
	HttpForcedRedirect HttpForcedRedirect
	Message            string
	Project            string
	Quic               Quic
	Status             string
	TemplateId         string
	Title              string
	Type               string
	UpdateTime         int64
}

type DescribeContentBlockTasksRequest struct {
	DomainName *string `json:",omitempty"`
	EndTime    *int64  `json:",omitempty"`
	PageNum    *int64  `json:",omitempty"`
	PageSize   *int64  `json:",omitempty"`
	StartTime  *int64  `json:",omitempty"`
	Status     *string `json:",omitempty"`
	TaskID     *string `json:",omitempty"`
	TaskType   string
	URL        *string `json:",omitempty"`
}

type DescribeContentBlockTasksResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeContentBlockTasksResult
}

type DescribeContentBlockTasksResult struct {
	Data     []BlockTaskInfo
	PageNum  int64
	PageSize int64
	Total    int64
}

type DescribeContentQuotaResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeContentQuotaResult
}

type DescribeContentQuotaResult struct {
	BlockLimit         int64
	BlockQuota         int64
	BlockRemain        int64
	PreloadLimit       int64
	PreloadQuota       int64
	PreloadRemain      int64
	RefreshDirLimit    int64
	RefreshDirQuota    int64
	RefreshDirRemain   int64
	RefreshQuota       int64
	RefreshQuotaLimit  int64
	RefreshRegexLimit  int64
	RefreshRegexQuota  int64
	RefreshRegexRemain int64
	RefreshRemain      int64
	UnblockLimit       int64
	UnblockQuota       int64
	UnblockRemain      int64
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
	PageNum  int64
	PageSize int64
	Total    int64
}

type DescribeDistrictIspDataRequest struct {
	Aggregate *string `json:",omitempty"`
	Domain    string
	EndTime   int64
	Interval  *string `json:",omitempty"`
	IpVersion *string `json:",omitempty"`
	Metric    string
	Protocol  *string `json:",omitempty"`
	StartTime int64
	TimeZone  *string `json:",omitempty"`
}

type DescribeDistrictIspDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeDistrictIspDataResult
}

type DescribeDistrictIspDataResult struct {
	Resources []DomainNrtDetailData
}

type DescribeEdgeNrtDataSummaryRequest struct {
	Aggregate           *string `json:",omitempty"`
	Area                *string `json:",omitempty"`
	BillingRegion       *string `json:",omitempty"`
	DisaggregateMetrics *string `json:",omitempty"`
	Domain              *string `json:",omitempty"`
	EndTime             int64
	Interval            *string `json:",omitempty"`
	InverseDomain       *bool   `json:",omitempty"`
	IpVersion           *string `json:",omitempty"`
	Isp                 *string `json:",omitempty"`
	Metric              string
	Project             *string `json:",omitempty"`
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
	InverseDomain *bool   `json:",omitempty"`
	Item          string
	Metric        string
	Project       *string `json:",omitempty"`
	StartTime     int64
}

type DescribeEdgeTopNrtDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeEdgeTopNrtDataResult
}

type DescribeEdgeTopNrtDataResult struct {
	Item           string
	Metric         string
	Name           string
	TopDataDetails []TopNrtDataDetail
}

type DescribeEdgeTopStatisticalDataRequest struct {
	Area      *string `json:",omitempty"`
	Domain    string
	EndTime   int64
	Item      string
	Metric    string
	StartTime int64
	UaType    *string `json:",omitempty"`
}

type DescribeEdgeTopStatisticalDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeEdgeTopStatisticalDataResult
}

type DescribeEdgeTopStatisticalDataResult struct {
	Item           string
	Metric         string
	Name           string
	TopDataDetails []EdgeTopStatisticalDataDetail
}

type DescribeEdgeTopStatusCodeRequest struct {
	Area          *string `json:",omitempty"`
	BillingRegion *string `json:",omitempty"`
	Domain        *string `json:",omitempty"`
	EndTime       int64
	InverseDomain *bool `json:",omitempty"`
	Item          string
	Metric        string
	Project       *string `json:",omitempty"`
	StartTime     int64
}

type DescribeEdgeTopStatusCodeResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeEdgeTopStatusCodeResult
}

type DescribeEdgeTopStatusCodeResult struct {
	Item           string
	Metric         string
	Name           string
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
	CdnIp    bool
	IP       string
	ISP      string
	Location string
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
	InverseDomain       *bool   `json:",omitempty"`
	IpVersion           *string `json:",omitempty"`
	Isp                 *string `json:",omitempty"`
	Metric              string
	Project             *string `json:",omitempty"`
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
	InverseDomain *bool   `json:",omitempty"`
	Item          string
	Metric        string
	Project       *string `json:",omitempty"`
	StartTime     int64
}

type DescribeOriginTopNrtDataResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeOriginTopNrtDataResult
}

type DescribeOriginTopNrtDataResult struct {
	Item           string
	Metric         string
	Name           string
	TopDataDetails []TopNrtDataDetail
}

type DescribeOriginTopStatusCodeRequest struct {
	Area          *string `json:",omitempty"`
	BillingRegion *string `json:",omitempty"`
	Domain        *string `json:",omitempty"`
	EndTime       int64
	InverseDomain *bool `json:",omitempty"`
	Item          string
	Metric        string
	Project       *string `json:",omitempty"`
	StartTime     int64
}

type DescribeOriginTopStatusCodeResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeOriginTopStatusCodeResult
}

type DescribeOriginTopStatusCodeResult struct {
	Item           string
	Metric         string
	Name           string
	TopDataDetails []TopStatusCodeDetail
}

type DescribeServiceTemplateRequest struct {
	TemplateId string
}

type DescribeServiceTemplateResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeServiceTemplateResult
}

type DescribeServiceTemplateResult struct {
	AreaAccessRule      *AreaAccessRule `json:",omitempty"`
	BandwidthLimit      *BandwidthLimit `json:",omitempty"`
	BoundDomains        []BoundDomain
	BrowserCache        []BrowserCacheControlRule
	Cache               []CacheControlRule
	CacheHost           *CacheHost `json:",omitempty"`
	CacheKey            []CacheKeyRule
	Compression         *Compression         `json:",omitempty"`
	CreateTime          *int64               `json:",omitempty"`
	CustomErrorPage     *CustomErrorPage     `json:",omitempty"`
	CustomizeAccessRule *CustomizeAccessRule `json:",omitempty"`
	DownloadSpeedLimit  *DownloadSpeedLimit  `json:",omitempty"`
	Exception           *bool                `json:",omitempty"`
	FollowRedirect      *bool                `json:",omitempty"`
	IPv6                *IPv6                `json:",omitempty"`
	IpAccessRule        *IpAccessRule        `json:",omitempty"`
	IpFreqLimit         *IpFreqLimit         `json:",omitempty"`
	IpSpeedLimit        *IpSpeedLimit        `json:",omitempty"`
	Message             *string              `json:",omitempty"`
	MethodDeniedRule    *MethodDeniedRule    `json:",omitempty"`
	NegativeCache       []NegativeCache
	Origin              []OriginRule
	OriginAccessRule    *OriginAccessRule `json:",omitempty"`
	OriginArg           []OriginArgRule
	OriginHost          *string `json:",omitempty"`
	OriginIPv6          *string `json:",omitempty"`
	OriginProtocol      string
	OriginRange         *bool               `json:",omitempty"`
	OriginRewrite       *OriginRewrite      `json:",omitempty"`
	OriginSni           *OriginSni          `json:",omitempty"`
	PageOptimization    *PageOptimization   `json:",omitempty"`
	Project             *string             `json:",omitempty"`
	RedirectionRewrite  *RedirectionRewrite `json:",omitempty"`
	RefererAccessRule   *RefererAccessRule  `json:",omitempty"`
	RemoteAuth          *RemoteAuth         `json:",omitempty"`
	RequestHeader       []RequestHeaderRule
	ResponseHeader      []ResponseHeaderRule
	SignedUrlAuth       *SignedUrlAuth       `json:",omitempty"`
	Status              *string              `json:",omitempty"`
	TemplateId          *string              `json:",omitempty"`
	Timeout             *Timeout             `json:",omitempty"`
	Title               *string              `json:",omitempty"`
	Type                *string              `json:",omitempty"`
	UaAccessRule        *UserAgentAccessRule `json:",omitempty"`
	UpdateTime          *int64               `json:",omitempty"`
	VideoDrag           *VideoDrag           `json:",omitempty"`
}

type DescribeTemplateDomainsRequest struct {
	Filters  []Filter
	PageNum  *int64 `json:",omitempty"`
	PageSize *int64 `json:",omitempty"`
}

type DescribeTemplateDomainsResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeTemplateDomainsResult
}

type DescribeTemplateDomainsResult struct {
	Domains    []CdnTemplateDomain
	TotalCount int64
}

type DescribeTemplatesRequest struct {
	Filters  []Filter
	PageNum  *int64 `json:",omitempty"`
	PageSize *int64 `json:",omitempty"`
}

type DescribeTemplatesResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DescribeTemplatesResult
}

type DescribeTemplatesResult struct {
	Templates  []DomainTemplateInfo
	TotalCount int64
}

type DomainCertStatus struct {
	CerStatus  string
	Domain     string
	DomainLock DomainLock
	Status     string
	Type       string
}

type DomainConfig struct {
	AreaAccessRule      *AreaAccessRule `json:",omitempty"`
	BackupCname         *string         `json:",omitempty"`
	BandwidthLimit      *BandwidthLimit `json:",omitempty"`
	BrowserCache        []BrowserCacheControlRule
	Cache               []CacheControlRule
	CacheHost           *CacheHost `json:",omitempty"`
	CacheKey            []CacheKeyRule
	Cname               *string              `json:",omitempty"`
	Compression         *Compression         `json:",omitempty"`
	CreateTime          *int64               `json:",omitempty"`
	CustomErrorPage     *CustomErrorPage     `json:",omitempty"`
	CustomizeAccessRule *CustomizeAccessRule `json:",omitempty"`
	Domain              *string              `json:",omitempty"`
	DownloadSpeedLimit  *DownloadSpeedLimit  `json:",omitempty"`
	FollowRedirect      *bool                `json:",omitempty"`
	HTTPS               *HTTPS               `json:",omitempty"`
	HeaderLogging       *HeaderLog           `json:",omitempty"`
	HttpForcedRedirect  *HttpForcedRedirect  `json:",omitempty"`
	IPv6                *IPv6                `json:",omitempty"`
	IpAccessRule        *IpAccessRule        `json:",omitempty"`
	IpFreqLimit         *IpFreqLimit         `json:",omitempty"`
	IpSpeedLimit        *IpSpeedLimit        `json:",omitempty"`
	LockStatus          *string              `json:",omitempty"`
	MethodDeniedRule    *MethodDeniedRule    `json:",omitempty"`
	NegativeCache       []NegativeCache
	Origin              []OriginRule
	OriginAccessRule    *OriginAccessRule `json:",omitempty"`
	OriginArg           []OriginArgRule
	OriginHost          *string `json:",omitempty"`
	OriginIPv6          *string `json:",omitempty"`
	OriginProtocol      string
	OriginRange         *bool               `json:",omitempty"`
	OriginRewrite       *OriginRewrite      `json:",omitempty"`
	OriginSni           *OriginSni          `json:",omitempty"`
	PageOptimization    *PageOptimization   `json:",omitempty"`
	Project             *string             `json:",omitempty"`
	Quic                *Quic               `json:",omitempty"`
	RedirectionRewrite  *RedirectionRewrite `json:",omitempty"`
	RefererAccessRule   *RefererAccessRule  `json:",omitempty"`
	RemoteAuth          *RemoteAuth         `json:",omitempty"`
	RequestBlockRule    *RequestBlockRule   `json:",omitempty"`
	RequestHeader       []RequestHeaderRule
	ResponseHeader      []ResponseHeaderRule
	ServiceRegion       *string `json:",omitempty"`
	ServiceType         string
	SignedUrlAuth       *SignedUrlAuth       `json:",omitempty"`
	Sparrow             *Sparrow             `json:",omitempty"`
	Status              *string              `json:",omitempty"`
	Timeout             *Timeout             `json:",omitempty"`
	UaAccessRule        *UserAgentAccessRule `json:",omitempty"`
	UpdateTime          *int64               `json:",omitempty"`
	UrlNormalize        *URLNormalize        `json:",omitempty"`
	VideoDrag           *VideoDrag           `json:",omitempty"`
}

type DomainLock struct {
	Remark string
	Status string
}

type DomainLogDetail struct {
	EndTime   int64
	LogName   string
	LogPath   string
	LogSize   int64
	StartTime int64
}

type DomainNrtDetailData struct {
	DataDetails []NrtDataDetails
	Name        string
}

type DomainStatus struct {
	Domain     string
	DomainLock DomainLock
	Status     string
	Type       string
}

type DomainSummary struct {
	BackupCname           string
	BackupOrigin          []string
	CacheShared           string
	CacheSharedTargetHost string
	Cname                 string
	ConfigStatus          string
	CreateTime            int64
	Domain                string
	DomainLock            DomainLock
	HTTPS                 bool
	IPv6                  bool
	IsConflictDomain      bool
	OriginProtocol        string
	PrimaryOrigin         []string
	Project               string
	ResourceTags          []ResourceTag
	ServiceRegion         string
	ServiceType           string
	SparrowList           []string
	Status                string
	UpdateTime            int64
	Waf                   bool
}

type DomainTemplateInfo struct {
	BoundDomains []BoundDomain
	CreateTime   int64
	Exception    bool
	Message      string
	Project      string
	Status       string
	TemplateId   string
	Title        string
	Type         string
	UpdateTime   int64
}

type DownloadSpeedLimit struct {
	DownloadSpeedLimitRules []DownloadSpeedLimitRule
	Switch                  *bool `json:",omitempty"`
}

type DownloadSpeedLimitAction struct {
	SpeedLimitRate      *int64          `json:",omitempty"`
	SpeedLimitRateAfter *int64          `json:",omitempty"`
	SpeedLimitTime      *SpeedLimitTime `json:",omitempty"`
}

type DownloadSpeedLimitRule struct {
	Condition                *Condition                `json:",omitempty"`
	DownloadSpeedLimitAction *DownloadSpeedLimitAction `json:",omitempty"`
}

type DuplicateTemplateRequest struct {
	Message            *string `json:",omitempty"`
	Project            *string `json:",omitempty"`
	ReferredTemplateId string
	Title              *string `json:",omitempty"`
}

type DuplicateTemplateResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           DuplicateTemplateResult
}

type DuplicateTemplateResult struct {
	TemplateId string
}

type EdgeStatisticalDataResource struct {
	Metrics []AccountingData
	Name    string
}

type EdgeTopStatisticalDataDetail struct {
	ItemKey   string
	ItemKeyCN string
	Value     float64
}

type ErrorObj struct {
	CodeN   int64
	Code    string
	Message string
}

type ErrorPageAction struct {
	Action       *string `json:",omitempty"`
	RedirectCode *string `json:",omitempty"`
	RedirectUrl  *string `json:",omitempty"`
	StatusCode   *string `json:",omitempty"`
}

type ErrorPageRule struct {
	ErrorPageAction *ErrorPageAction `json:",omitempty"`
}

type FeatureConfig struct {
	OriginV2 bool
}

type Filter struct {
	Fuzzy *bool   `json:",omitempty"`
	Name  *string `json:",omitempty"`
	Value []string
}

type ForcedRedirect struct {
	EnableForcedRedirect *bool   `json:",omitempty"`
	StatusCode           *string `json:",omitempty"`
}

type HTTPS struct {
	CertInfo       *CertInfo `json:",omitempty"`
	CertInfoList   []CertInfo
	DisableHttp    *bool           `json:",omitempty"`
	ForcedRedirect *ForcedRedirect `json:",omitempty"`
	HTTP2          *bool           `json:",omitempty"`
	Hsts           *Hsts           `json:",omitempty"`
	OCSP           *bool           `json:",omitempty"`
	Switch         *bool           `json:",omitempty"`
	TlsVersion     []string
}

type HTTPSCommon struct {
	DisableHttp    *bool           `json:",omitempty"`
	ForcedRedirect *ForcedRedirect `json:",omitempty"`
	HTTP2          *bool           `json:",omitempty"`
	Hsts           *Hsts           `json:",omitempty"`
	OCSP           *bool           `json:",omitempty"`
	TlsVersion     []string
}

type HeaderLog struct {
	HeaderLogging *string `json:",omitempty"`
	Switch        *bool   `json:",omitempty"`
}

type Hsts struct {
	Subdomain *string `json:",omitempty"`
	Switch    *bool   `json:",omitempty"`
	Ttl       *int64  `json:",omitempty"`
}

type HttpForcedRedirect struct {
	EnableForcedRedirect *bool   `json:",omitempty"`
	StatusCode           *string `json:",omitempty"`
}

type IPInfo struct {
	CdnIp    bool
	IP       string
	ISP      string
	Location string
}

type IPv6 struct {
	Switch *bool `json:",omitempty"`
}

type IpAccessRule struct {
	Ip           []string
	RuleType     *string             `json:",omitempty"`
	SharedConfig *CommonGlobalConfig `json:",omitempty"`
	Switch       *bool               `json:",omitempty"`
}

type IpFreqLimit struct {
	IpFreqLimitRules []IpFreqLimitRule
	Switch           *bool `json:",omitempty"`
}

type IpFreqLimitAction struct {
	Action        *string `json:",omitempty"`
	FreqLimitRate *int64  `json:",omitempty"`
	StatusCode    *string `json:",omitempty"`
}

type IpFreqLimitRule struct {
	Condition         *Condition         `json:",omitempty"`
	IpFreqLimitAction *IpFreqLimitAction `json:",omitempty"`
}

type IpSpeedLimit struct {
	IpSpeedLimitRules []IpSpeedLimitRule
	Switch            *bool `json:",omitempty"`
}

type IpSpeedLimitAction struct {
	SpeedLimitRate *int64 `json:",omitempty"`
}

type IpSpeedLimitRule struct {
	Condition          *Condition          `json:",omitempty"`
	IpSpeedLimitAction *IpSpeedLimitAction `json:",omitempty"`
}

type ListCdnCertInfoRequest struct {
	CertId           *string `json:",omitempty"`
	Configured       *bool   `json:",omitempty"`
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
	CertInfo      []ListCertInfo
	ExpiringCount int64
	PageNum       int64
	PageSize      int64
	Total         int64
}

type ListCdnDomainsRequest struct {
	Domain         *string `json:",omitempty"`
	ExactMatch     *bool   `json:",omitempty"`
	HTTPS          *bool   `json:",omitempty"`
	IPv6           *bool   `json:",omitempty"`
	NeedSparrows   *bool   `json:",omitempty"`
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
	Data     []DomainSummary
	PageNum  int64
	PageSize int64
	Total    int64
}

type ListCertInfo struct {
	CertFingerprint        CertFingerprint
	CertId                 string
	CertName               string
	ConfiguredDomain       string
	ConfiguredDomainDetail []ConfiguredDomain
	Desc                   string
	DnsName                string
	EffectiveTime          int64
	ExpireTime             int64
	Source                 string
	Status                 string
}

type ListCertInfoRequest struct {
	CertId           *string   `json:",omitempty"`
	ConfiguredDomain *string   `json:",omitempty"`
	FuzzyMatch       *bool     `json:",omitempty"`
	Name             *string   `json:",omitempty"`
	PageNum          *int64    `json:",omitempty"`
	PageSize         *int64    `json:",omitempty"`
	SetPagination    *bool     `json:",omitempty"`
	SortRule         *SortRule `json:",omitempty"`
	Source           *string   `json:",omitempty"`
	Status           *string   `json:",omitempty"`
}

type ListCertInfoResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           ListCertInfoResult
}

type ListCertInfoResult struct {
	CertInfo      []ListCertInfo
	ExpiringCount int64
	PageNum       int64
	PageSize      int64
	Total         int64
}

type ListResourceTagsResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           ListResourceTagsResult
}

type ListResourceTagsResult struct {
	ResourceTags []ResourceTag
}

type LockTemplateRequest struct {
	TemplateId string
}

type LockTemplateResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}

type MethodDeniedRule struct {
	Methods *string `json:",omitempty"`
	Switch  *bool   `json:",omitempty"`
}

type MetricTimestampValue struct {
	Metric string
	Values []TimestampValue
}

type MetricValue struct {
	Metric string
	Value  float64
}

type NamePair struct {
	Code string
	Name string
}

type NegativeCache struct {
	Condition         *Condition           `json:",omitempty"`
	NegativeCacheRule *NegativeCacheAction `json:",omitempty"`
}

type NegativeCacheAction struct {
	Action     *string `json:",omitempty"`
	IgnoreCase *bool   `json:",omitempty"`
	StatusCode *string `json:",omitempty"`
	Ttl        *int64  `json:",omitempty"`
}

type NrtDataDetails struct {
	Isp     string
	Metrics []MetricTimestampValue
	Region  string
}

type NrtDataResource struct {
	BillingRegion string
	Isp           string
	Metrics       []MetricTimestampValue
	Name          string
	Region        string
}

type NrtDataSummaryResource struct {
	BillingRegion string
	Metrics       []MetricValue
	Name          string
}

type OriginAccessRule struct {
	AllowEmpty *bool `json:",omitempty"`
	IgnoreCase *bool `json:",omitempty"`
	Origins    []string
	RuleType   *string `json:",omitempty"`
	Switch     *bool   `json:",omitempty"`
}

type OriginAction struct {
	OriginLines []OriginLine
}

type OriginArgAction struct {
	OriginArgComponents []OriginArgComponents
}

type OriginArgComponents struct {
	Action    *string `json:",omitempty"`
	Object    *string `json:",omitempty"`
	Subobject *string `json:",omitempty"`
}

type OriginArgRule struct {
	Condition       *Condition       `json:",omitempty"`
	OriginArgAction *OriginArgAction `json:",omitempty"`
}

type OriginLine struct {
	Address             *string            `json:",omitempty"`
	HttpPort            *string            `json:",omitempty"`
	HttpsPort           *string            `json:",omitempty"`
	InstanceType        *string            `json:",omitempty"`
	OriginHost          *string            `json:",omitempty"`
	OriginType          *string            `json:",omitempty"`
	PrivateBucketAccess *bool              `json:",omitempty"`
	PrivateBucketAuth   *PrivateBucketAuth `json:",omitempty"`
	Region              *string            `json:",omitempty"`
	Weight              *string            `json:",omitempty"`
}

type OriginRewrite struct {
	OriginRewriteRule []OriginRewriteRule
	Switch            *bool `json:",omitempty"`
}

type OriginRewriteAction struct {
	SourcePath *string `json:",omitempty"`
	TargetPath *string `json:",omitempty"`
}

type OriginRewriteRule struct {
	Condition           *Condition           `json:",omitempty"`
	OriginRewriteAction *OriginRewriteAction `json:",omitempty"`
}

type OriginRule struct {
	Condition    *Condition    `json:",omitempty"`
	OriginAction *OriginAction `json:",omitempty"`
}

type OriginSni struct {
	SniDomain *string `json:",omitempty"`
	Switch    *bool   `json:",omitempty"`
}

type PageOptimization struct {
	OptimizationType []string
	Switch           *bool `json:",omitempty"`
}

type ParamCapRule struct {
	CapMode   *string `json:",omitempty"`
	ParamName *string `json:",omitempty"`
	UriLevel  *int64  `json:",omitempty"`
}

type PreloadHeader struct {
	Key   *string `json:",omitempty"`
	Value *string `json:",omitempty"`
}

type PrivateBucketAuth struct {
	AuthType           *string             `json:",omitempty"`
	Switch             *bool               `json:",omitempty"`
	TosAuthInformation *TosAuthInformation `json:",omitempty"`
}

type QueryStringComponents struct {
	Action *string `json:",omitempty"`
	Value  *string `json:",omitempty"`
}

type QueryStringInstance struct {
	Action    *string `json:",omitempty"`
	Key       *string `json:",omitempty"`
	Value     *string `json:",omitempty"`
	ValueType *string `json:",omitempty"`
}

type QueryStringRule struct {
	QueryStringComponents *QueryStringComponents `json:",omitempty"`
	QueryStringInstances  []QueryStringInstance
}

type Quic struct {
	Switch *bool `json:",omitempty"`
}

type RedirectionAction struct {
	RedirectCode          *string                `json:",omitempty"`
	SourcePath            *string                `json:",omitempty"`
	TargetHost            *string                `json:",omitempty"`
	TargetPath            *string                `json:",omitempty"`
	TargetProtocol        *string                `json:",omitempty"`
	TargetQueryComponents *TargetQueryComponents `json:",omitempty"`
}

type RedirectionRewrite struct {
	RedirectionRule []RedirectionRule
	Switch          *bool `json:",omitempty"`
}

type RedirectionRule struct {
	RedirectionAction *RedirectionAction `json:",omitempty"`
}

type RefererAccessRule struct {
	AllowEmpty   *bool `json:",omitempty"`
	Referers     []string
	ReferersType *ReferersType       `json:",omitempty"`
	RuleType     *string             `json:",omitempty"`
	SharedConfig *CommonGlobalConfig `json:",omitempty"`
	Switch       *bool               `json:",omitempty"`
}

type RefererType struct {
	Referers []string
}

type ReferersType struct {
	CommonType  *CommonReferType `json:",omitempty"`
	RegularType *RefererType     `json:",omitempty"`
}

type RemoteAuth struct {
	RemoteAuthRules []RemoteAuthRule
	Switch          *bool `json:",omitempty"`
}

type RemoteAuthRule struct {
	Condition            *Condition            `json:",omitempty"`
	RemoteAuthRuleAction *RemoteAuthRuleAction `json:",omitempty"`
}

type RemoteAuthRuleAction struct {
	AuthModeConfig     *AuthModeConfig        `json:",omitempty"`
	AuthResponseConfig *AuthResponseConfig    `json:",omitempty"`
	QueryStringRules   *QueryStringRule       `json:",omitempty"`
	RequestBodyRules   *string                `json:",omitempty"`
	RequestHeaderRules *AuthRequestHeaderRule `json:",omitempty"`
}

type RequestBlockRule struct {
	BlockRule []BlockRule
	Switch    *bool `json:",omitempty"`
}

type RequestHeaderAction struct {
	RequestHeaderInstances []RequestHeaderInstance
}

type RequestHeaderComponent struct {
	Action *string `json:",omitempty"`
	Value  *string `json:",omitempty"`
}

type RequestHeaderInstance struct {
	Action    *string `json:",omitempty"`
	Key       *string `json:",omitempty"`
	Value     *string `json:",omitempty"`
	ValueType *string `json:",omitempty"`
}

type RequestHeaderRule struct {
	Condition           *Condition           `json:",omitempty"`
	RequestHeaderAction *RequestHeaderAction `json:",omitempty"`
}

type ResourceTag struct {
	Key   *string `json:",omitempty"`
	Value *string `json:",omitempty"`
}

type ResponseAction struct {
	StatusCode *string `json:",omitempty"`
}

type ResponseHeaderAction struct {
	ResponseHeaderInstances []ResponseHeaderInstance
}

type ResponseHeaderInstance struct {
	AccessOriginControl *bool   `json:",omitempty"`
	Action              *string `json:",omitempty"`
	Key                 *string `json:",omitempty"`
	Value               *string `json:",omitempty"`
	ValueType           *string `json:",omitempty"`
}

type ResponseHeaderRule struct {
	Condition            *Condition            `json:",omitempty"`
	ResponseHeaderAction *ResponseHeaderAction `json:",omitempty"`
}

type ResponseMetadata struct {
	RequestId string
	Service   *string   `json:",omitempty"`
	Region    *string   `json:",omitempty"`
	Action    *string   `json:",omitempty"`
	Version   *string   `json:",omitempty"`
	Error     *ErrorObj `json:",omitempty"`
}

type RewriteM3u8Rule struct {
	DeleteParam      *bool `json:",omitempty"`
	KeepM3u8Param    *bool `json:",omitempty"`
	TransferEncoding *bool `json:",omitempty"`
}

type SignedUrlAuth struct {
	SignedUrlAuthRules []SignedUrlAuthRule
	Switch             *bool `json:",omitempty"`
}

type SignedUrlAuthAction struct {
	BackupSecretKey     *string              `json:",omitempty"`
	CustomVariableRules *CustomVariableRules `json:",omitempty"`
	Duration            *int64               `json:",omitempty"`
	KeepOriginArg       *bool                `json:",omitempty"`
	MasterSecretKey     *string              `json:",omitempty"`
	MpdVarExpand        *bool                `json:",omitempty"`
	RewriteM3u8         *bool                `json:",omitempty"`
	RewriteM3u8Rule     *RewriteM3u8Rule     `json:",omitempty"`
	RewriteMpd          *bool                `json:",omitempty"`
	SignName            *string              `json:",omitempty"`
	SignatureRule       []string
	TimeFormat          *string `json:",omitempty"`
	TimeName            *string `json:",omitempty"`
	URLAuthType         *string `json:",omitempty"`
}

type SignedUrlAuthRule struct {
	Condition           *Condition           `json:",omitempty"`
	SignedUrlAuthAction *SignedUrlAuthAction `json:",omitempty"`
}

type SortRule struct {
	Asc     *bool   `json:",omitempty"`
	OrderBy *string `json:",omitempty"`
}

type Sparrow struct {
	SparrowRules []SparrowRule
	Switch       *bool `json:",omitempty"`
}

type SparrowAction struct {
	Action     *string `json:",omitempty"`
	IgnoreCase *bool   `json:",omitempty"`
	SparrowID  *string `json:",omitempty"`
}

type SparrowRule struct {
	Condition     *Condition     `json:",omitempty"`
	SparrowAction *SparrowAction `json:",omitempty"`
}

type SpeedLimitTime struct {
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
	Type *string `json:",omitempty"`
	Urls string
}

type SubmitBlockTaskResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           SubmitBlockTaskResult
}

type SubmitBlockTaskResult struct {
	TaskID string
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
	CommitNum *int64 `json:",omitempty"`
	TaskID    string
}

type SubmitRefreshTaskRequest struct {
	Delete *bool   `json:",omitempty"`
	Prefix *bool   `json:",omitempty"`
	Type   *string `json:",omitempty"`
	Urls   string
}

type SubmitRefreshTaskResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           SubmitRefreshTaskResult
}

type SubmitRefreshTaskResult struct {
	TaskID string
}

type SubmitUnblockTaskRequest struct {
	Type *string `json:",omitempty"`
	Urls string
}

type SubmitUnblockTaskResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
	Result           SubmitUnblockTaskResult
}

type SubmitUnblockTaskResult struct {
	TaskID string
}

type TargetQueryComponents struct {
	Action *string `json:",omitempty"`
	Value  *string `json:",omitempty"`
}

type TemplateCertInfo struct {
	CertId string
}

type Timeout struct {
	Switch       *bool `json:",omitempty"`
	TimeoutRules []TimeoutRule
}

type TimeoutAction struct {
	HttpTimeout *int64 `json:",omitempty"`
	TcpTimeout  *int64 `json:",omitempty"`
}

type TimeoutRule struct {
	Condition     *Condition     `json:",omitempty"`
	TimeoutAction *TimeoutAction `json:",omitempty"`
}

type TimestampValue struct {
	Timestamp int64
	Value     float64
}

type TopInstanceDetail struct {
	BeginTime        string
	BillingCode      string
	BillingCycle     string
	BillingData      string
	BillingDesc      string
	CreateTime       string
	FreeTimePeriods  []int64
	InstanceCategory string
	InstanceType     string
	MetricType       string
	ServiceRegion    string
	StartTime        string
	Status           string
}

type TopNrtDataDetail struct {
	Bandwidth                float64
	BandwidthPeakTime        int64
	BsBandwidth              float64
	BsBandwidthPeakTime      int64
	BsFlux                   float64
	BsFluxRatio              float64
	DynamicRequest           int64
	DynamicRequestRatio      float64
	Flux                     float64
	FluxRatio                float64
	InboundBandwidth         float64
	InboundBandwidthPeakTime int64
	InboundFlux              float64
	InboundFluxRatio         float64
	ItemKey                  string
	ItemKeyCN                string
	PV                       float64
	PVRatio                  float64
	Quic                     int64
	StaticRequest            int64
	StaticRequestRatio       float64
}

type TopStatusCodeDetail struct {
	Status2xx      float64 `json:"2xx"`
	Status2xxRatio float64 `json:"2xxRatio"`
	Status3xx      float64 `json:"3xx"`
	Status3xxRatio float64 `json:"3xxRatio"`
	Status4xx      float64 `json:"4xx"`
	Status4xxRatio float64 `json:"4xxRatio"`
	Status5xx      float64 `json:"5xx"`
	Status5xxRatio float64 `json:"5xxRatio"`
	ItemKey        string
}

type TosAuthInformation struct {
	AccessKeyId       *string `json:",omitempty"`
	AccessKeySecret   *string `json:",omitempty"`
	RoleAccountId     *string `json:",omitempty"`
	RoleName          *string `json:",omitempty"`
	RolePassAccountId *string `json:",omitempty"`
	RolePassName      *string `json:",omitempty"`
}

type URLNormalize struct {
	NormalizeObject []string
	Switch          *bool `json:",omitempty"`
}

type UpdateCdnConfigRequest struct {
	AreaAccessRule      *AreaAccessRule `json:",omitempty"`
	BandwidthLimit      *BandwidthLimit `json:",omitempty"`
	BrowserCache        []BrowserCacheControlRule
	Cache               []CacheControlRule
	CacheHost           *CacheHost `json:",omitempty"`
	CacheKey            []CacheKeyRule
	Compression         *Compression         `json:",omitempty"`
	CustomErrorPage     *CustomErrorPage     `json:",omitempty"`
	CustomizeAccessRule *CustomizeAccessRule `json:",omitempty"`
	Domain              *string              `json:",omitempty"`
	DownloadSpeedLimit  *DownloadSpeedLimit  `json:",omitempty"`
	FollowRedirect      *bool                `json:",omitempty"`
	HTTPS               *HTTPS               `json:",omitempty"`
	HeaderLogging       *HeaderLog           `json:",omitempty"`
	HttpForcedRedirect  *HttpForcedRedirect  `json:",omitempty"`
	IPv6                *IPv6                `json:",omitempty"`
	IpAccessRule        *IpAccessRule        `json:",omitempty"`
	IpFreqLimit         *IpFreqLimit         `json:",omitempty"`
	IpSpeedLimit        *IpSpeedLimit        `json:",omitempty"`
	MethodDeniedRule    *MethodDeniedRule    `json:",omitempty"`
	NegativeCache       []NegativeCache
	Origin              []OriginRule
	OriginAccessRule    *OriginAccessRule `json:",omitempty"`
	OriginArg           []OriginArgRule
	OriginHost          *string             `json:",omitempty"`
	OriginIPv6          *string             `json:",omitempty"`
	OriginProtocol      *string             `json:",omitempty"`
	OriginRange         *bool               `json:",omitempty"`
	OriginRewrite       *OriginRewrite      `json:",omitempty"`
	OriginSni           *OriginSni          `json:",omitempty"`
	PageOptimization    *PageOptimization   `json:",omitempty"`
	Quic                *Quic               `json:",omitempty"`
	RedirectionRewrite  *RedirectionRewrite `json:",omitempty"`
	RefererAccessRule   *RefererAccessRule  `json:",omitempty"`
	RemoteAuth          *RemoteAuth         `json:",omitempty"`
	RequestBlockRule    *RequestBlockRule   `json:",omitempty"`
	RequestHeader       []RequestHeaderRule
	ResponseHeader      []ResponseHeaderRule
	ServiceRegion       *string              `json:",omitempty"`
	ServiceType         *string              `json:",omitempty"`
	SignedUrlAuth       *SignedUrlAuth       `json:",omitempty"`
	Sparrow             *Sparrow             `json:",omitempty"`
	Timeout             *Timeout             `json:",omitempty"`
	UaAccessRule        *UserAgentAccessRule `json:",omitempty"`
	UrlNormalize        *URLNormalize        `json:",omitempty"`
	VideoDrag           *VideoDrag           `json:",omitempty"`
}

type UpdateCdnConfigResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}

type UpdateCipherTemplateRequest struct {
	HTTPS              *HTTPSCommon        `json:",omitempty"`
	HttpForcedRedirect *HttpForcedRedirect `json:",omitempty"`
	Message            *string             `json:",omitempty"`
	Quic               *Quic               `json:",omitempty"`
	TemplateId         string
	Title              *string `json:",omitempty"`
}

type UpdateCipherTemplateResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}

type UpdateResourceTagsRequest struct {
	ResourceTags []ResourceTag
	Resources    []string
}

type UpdateResourceTagsResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}

type UpdateServiceTemplateRequest struct {
	AreaAccessRule      *AreaAccessRule `json:",omitempty"`
	BandwidthLimit      *BandwidthLimit `json:",omitempty"`
	BrowserCache        []BrowserCacheControlRule
	Cache               []CacheControlRule
	CacheHost           *CacheHost `json:",omitempty"`
	CacheKey            []CacheKeyRule
	Compression         *Compression         `json:",omitempty"`
	CustomErrorPage     *CustomErrorPage     `json:",omitempty"`
	CustomizeAccessRule *CustomizeAccessRule `json:",omitempty"`
	DownloadSpeedLimit  *DownloadSpeedLimit  `json:",omitempty"`
	FollowRedirect      *bool                `json:",omitempty"`
	IPv6                *IPv6                `json:",omitempty"`
	IpAccessRule        *IpAccessRule        `json:",omitempty"`
	IpFreqLimit         *IpFreqLimit         `json:",omitempty"`
	IpSpeedLimit        *IpSpeedLimit        `json:",omitempty"`
	Message             *string              `json:",omitempty"`
	MethodDeniedRule    *MethodDeniedRule    `json:",omitempty"`
	NegativeCache       []NegativeCache
	Origin              []OriginRule
	OriginAccessRule    *OriginAccessRule `json:",omitempty"`
	OriginArg           []OriginArgRule
	OriginHost          *string `json:",omitempty"`
	OriginIPv6          *string `json:",omitempty"`
	OriginProtocol      string
	OriginRange         *bool               `json:",omitempty"`
	OriginRewrite       *OriginRewrite      `json:",omitempty"`
	OriginSni           *OriginSni          `json:",omitempty"`
	PageOptimization    *PageOptimization   `json:",omitempty"`
	RedirectionRewrite  *RedirectionRewrite `json:",omitempty"`
	RefererAccessRule   *RefererAccessRule  `json:",omitempty"`
	RemoteAuth          *RemoteAuth         `json:",omitempty"`
	RequestHeader       []RequestHeaderRule
	ResponseHeader      []ResponseHeaderRule
	SignedUrlAuth       *SignedUrlAuth `json:",omitempty"`
	TemplateId          string
	Timeout             *Timeout             `json:",omitempty"`
	Title               *string              `json:",omitempty"`
	UaAccessRule        *UserAgentAccessRule `json:",omitempty"`
	VideoDrag           *VideoDrag           `json:",omitempty"`
}

type UpdateServiceTemplateResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}

type UpdateTemplateDomainRequest struct {
	CertId            *string `json:",omitempty"`
	CipherTemplateId  *string `json:",omitempty"`
	Domains           []string
	HTTPSSwitch       *string `json:",omitempty"`
	ServiceRegion     *string `json:",omitempty"`
	ServiceTemplateId *string `json:",omitempty"`
}

type UpdateTemplateDomainResponse struct {
	ResponseMetadata *ResponseMetadata `json:",omitempty"`
}

type UriParamSup struct {
	JoinSymbol  *string `json:",omitempty"`
	SplitSymbol *string `json:",omitempty"`
	StartLevel  *int64  `json:",omitempty"`
	TermLevel   *int64  `json:",omitempty"`
}

type UserAgentAccessRule struct {
	AllowEmpty *bool   `json:",omitempty"`
	IgnoreCase *bool   `json:",omitempty"`
	RuleType   *string `json:",omitempty"`
	Switch     *bool   `json:",omitempty"`
	UserAgent  []string
}

type VideoDrag struct {
	Switch *bool `json:",omitempty"`
}
