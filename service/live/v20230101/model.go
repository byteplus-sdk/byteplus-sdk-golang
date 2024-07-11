package live_v20230101

type BindCertBody struct {

	// REQUIRED; 需要绑定的 HTTPS 证书的证书链 ID，可以通过查询证书列表 [https://www.volcengine.com/docs/6469/1126822]接口获取。
	ChainID string `json:"ChainID"`

	// REQUIRED; 填写需要配置 HTTPS 证书的域名。 您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看需要绑定证书的域名。
	Domain string `json:"Domain"`

	// 证书域名。
	CertDomain *string `json:"CertDomain,omitempty"`

	// 是否启用 HTTPS 协议，默认值为 false，取值及含义如下所示。
	// * false：关闭；
	// * true：启用。
	HTTPS *bool `json:"HTTPS,omitempty"`

	// 最大支持的TLS版本，不填默认不校验，可选值为：TLSv1.0、TLSv1.1、TLSv1.2、TLSv1.3
	MaxTLSVersion *string `json:"MaxTLSVersion,omitempty"`

	// 最小支持的TLS版本，不填默认为TLSv1.2，可选值为：TLSv1.0、TLSv1.1、TLSv1.2、TLSv1.3
	MinTLSVersion *string `json:"MinTLSVersion,omitempty"`

	// 是否是客户自定义的证书链，如果是则跳过证书合法性校验。不填默认为false。
	UserDefinedChain *bool `json:"UserDefinedChain,omitempty"`

	// 视频直播服务的配置空间，由 1 到 60 位数字、字母、下划线及"-"和"."组成 :::tip 与 Domain 二选一。
	Vhost *string `json:"Vhost,omitempty"`
}

type BindCertRes struct {

	// REQUIRED
	ResponseMetadata BindCertResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type BindCertResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                            `json:"Version"`
	Error   *BindCertResResponseMetadataError `json:"Error,omitempty"`
}

type BindCertResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type BindEncryptDRMBody struct {

	// REQUIRED; 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 是否开启当前 DRM 加密配置，取值及含义如下所示。
	// * true：开启；
	// * false：关闭。
	Enable bool `json:"Enable"`

	// REQUIRED; 待加密的转码流对应的转码流后缀配置。您可以调用查询转码配置列表 [https://www.volcengine.com/docs/6469/1126853]接口或在视频直播控制台的转码配置 [https://console.volcengine.com/live/main/application/transcode]页面，查看转码配置的转码流后缀。
	EncryptTranscodeSuffix []string `json:"EncryptTranscodeSuffix"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`
}

type BindEncryptDRMRes struct {

	// REQUIRED
	ResponseMetadata BindEncryptDRMResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; Anything
	Result interface{} `json:"Result"`
}

type BindEncryptDRMResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

// Components1404CjzSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesHlsparamPropertiesTosparam
// - TOS 存储相关配置
// 说明
// TOSParam和VODParam配置且配置其中一个。
type Components1404CjzSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesHlsparamPropertiesTosparam struct {
	Bucket string `json:"Bucket"`

	Enable bool `json:"Enable"`

	ExactObject string `json:"ExactObject"`

	StorageDir string `json:"StorageDir"`
}

type Components17Ohct5SchemasDescribeliveasrdurationdataresPropertiesResultPropertiesAsrdurationdetaildataItemsPropertiesAsrdurationdataItems struct {
	Duration float32 `json:"Duration"`

	TimeStamp string `json:"TimeStamp"`
}

// Components1Hkcrc4SchemasListvhostsnapshotpresetresPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetPropertiesCallbackdetail
// - 回调信息。
type Components1Hkcrc4SchemasListvhostsnapshotpresetresPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetPropertiesCallbackdetail struct {
	URL string `json:"URL"`

	CallbackType *string `json:"CallbackType,omitempty"`
}

// Components1Via6UrSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesMp4ParamPropertiesTosparam
// - TOS 存储相关配置
// 说明
// TOSParam和VODParam配置且配置其中一个。
type Components1Via6UrSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesMp4ParamPropertiesTosparam struct {
	Bucket string `json:"Bucket"`

	Enable bool `json:"Enable"`

	ExactObject string `json:"ExactObject"`

	StorageDir string `json:"StorageDir"`
}

type Components1Wv3ClqSchemasUpdatetranscodepresetbodyPropertiesTranscodestructPropertiesAbtestAdditionalproperties struct {
	Label *int32 `json:"Label,omitempty"`
	Ratio *int32 `json:"Ratio,omitempty"`
}

// Components44Na0KSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesFlvparam
// - 录制为 FLV 格式时的录制参数。
type Components44Na0KSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesFlvparam struct {
	Duration *int32 `json:"Duration,omitempty"`

	Enable *bool `json:"Enable,omitempty"`

	RealtimeRecordDuration *int32 `json:"RealtimeRecordDuration,omitempty"`

	Splice *int32 `json:"Splice,omitempty"`

	TOSParam *ComponentsBbqv7RSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesFlvparamPropertiesTosparam `json:"TOSParam,omitempty"`

	VODParam *ComponentsKovkk9SchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesFlvparamPropertiesVodparam `json:"VODParam,omitempty"`
}

// ComponentsAoysk3SchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesHlsparam
// - 录制为 HLS 格式时的录制参数。
type ComponentsAoysk3SchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesHlsparam struct {
	Duration *int32 `json:"Duration,omitempty"`

	Enable *bool `json:"Enable,omitempty"`

	RealtimeRecordDuration *int32 `json:"RealtimeRecordDuration,omitempty"`

	Splice *int32 `json:"Splice,omitempty"`

	TOSParam *Components1404CjzSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesHlsparamPropertiesTosparam `json:"TOSParam,omitempty"`

	VODParam *ComponentsS0Ofr3SchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesHlsparamPropertiesVodparam `json:"VODParam,omitempty"`
}

// ComponentsBbqv7RSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesFlvparamPropertiesTosparam
// - TOS 存储相关配置。
type ComponentsBbqv7RSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesFlvparamPropertiesTosparam struct {
	Bucket string `json:"Bucket"`

	Enable bool `json:"Enable"`

	ExactObject string `json:"ExactObject"`

	StorageDir string `json:"StorageDir"`
}

// ComponentsFuamuzSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfig
// - 录制模板详细配置。
type ComponentsFuamuzSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfig struct {
	FlvParam *Components44Na0KSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesFlvparam `json:"FlvParam,omitempty"`

	HlsParam *ComponentsAoysk3SchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesHlsparam `json:"HlsParam,omitempty"`

	Mp4Param *ComponentsKqy98ZSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesMp4Param `json:"Mp4Param,omitempty"`

	OriginRecord *int32 `json:"OriginRecord,omitempty"`

	SliceDuration *int32 `json:"SliceDuration,omitempty"`

	TranscodeRecord *int32 `json:"TranscodeRecord,omitempty"`

	TranscodeSuffixList []*string `json:"TranscodeSuffixList,omitempty"`
}

// ComponentsKovkk9SchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesFlvparamPropertiesVodparam
// - VOD 存储相关配置。
type ComponentsKovkk9SchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesFlvparamPropertiesVodparam struct {
	ClassificationID *int32 `json:"ClassificationID,omitempty"`

	Enable *bool `json:"Enable,omitempty"`

	ExactObject *string `json:"ExactObject,omitempty"`

	StorageClass *int32 `json:"StorageClass,omitempty"`

	VodNamespace *string `json:"VodNamespace,omitempty"`

	WorkflowID *string `json:"WorkflowID,omitempty"`
}

// ComponentsKqy98ZSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesMp4Param
// - 录制为 HLS 格式时的录制参数。
type ComponentsKqy98ZSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesMp4Param struct {
	Duration *int32 `json:"Duration,omitempty"`

	Enable *bool `json:"Enable,omitempty"`

	RealtimeRecordDuration *int32 `json:"RealtimeRecordDuration,omitempty"`

	Splice *int32 `json:"Splice,omitempty"`

	TOSParam *Components1Via6UrSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesMp4ParamPropertiesTosparam `json:"TOSParam,omitempty"`

	VODParam *ComponentsQms0JiSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesMp4ParamPropertiesVodparam `json:"VODParam,omitempty"`
}

// ComponentsQms0JiSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesMp4ParamPropertiesVodparam
// - VOD 存储相关配置
// 说明
// TOSParam和VODParam配置且配置其中一个。
type ComponentsQms0JiSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesMp4ParamPropertiesVodparam struct {
	ClassificationID *int32 `json:"ClassificationID,omitempty"`

	Enable *bool `json:"Enable,omitempty"`

	ExactObject *string `json:"ExactObject,omitempty"`

	StorageClass *int32 `json:"StorageClass,omitempty"`

	VodNamespace *string `json:"VodNamespace,omitempty"`

	WorkflowID *string `json:"WorkflowID,omitempty"`
}

// ComponentsS0Ofr3SchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesHlsparamPropertiesVodparam
// - VOD 存储相关配置
// 说明
// TOSParam和VODParam配置且配置其中一个。
type ComponentsS0Ofr3SchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesHlsparamPropertiesVodparam struct {
	ClassificationID *int32 `json:"ClassificationID,omitempty"`

	Enable *bool `json:"Enable,omitempty"`

	ExactObject *string `json:"ExactObject,omitempty"`

	StorageClass *int32 `json:"StorageClass,omitempty"`

	VodNamespace *string `json:"VodNamespace,omitempty"`

	WorkflowID *string `json:"WorkflowID,omitempty"`
}

type ComponentsXsjbgcSchemasCreatetranscodepresetbodyPropertiesTranscodestructPropertiesAbtestAdditionalproperties struct {
	Label *int32 `json:"Label,omitempty"`
	Ratio *int32 `json:"Ratio,omitempty"`
}

type CreateCertBody struct {

	// REQUIRED; 证书信息。
	Rsa CreateCertBodyRsa `json:"Rsa"`

	// REQUIRED; 证书用途，默认为 https，取值及含义如下所示。
	// * https：用于 HTTPS 加密；
	// * sign：用于签名加密。
	UseWay string `json:"UseWay"`

	// 证书名称。
	CertName *string `json:"CertName,omitempty"`

	// 证书链 ID，用于标识整个证书链，包括叶子证书（服务器证书）、中间证书（中间 CA 证书）以及根证书（根 CA 证书）。
	ChainID *string `json:"ChainID,omitempty"`

	// 项目名称，默认值为 default，您可以登录访问控制 [https://console.volcengine.com/iam/resourcemanage/project]获取项目名称。
	ProjectName *string `json:"ProjectName,omitempty"`
}

// CreateCertBodyRsa - 证书信息。
type CreateCertBodyRsa struct {

	// REQUIRED; 证书私钥。
	Prikey string `json:"Prikey"`

	// REQUIRED; 证书公钥。
	Pubkey string `json:"Pubkey"`
}

type CreateCertRes struct {

	// REQUIRED
	ResponseMetadata CreateCertResResponseMetadata `json:"ResponseMetadata"`
	Result           *CreateCertResResult          `json:"Result,omitempty"`
}

type CreateCertResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                              `json:"Version"`
	Error   *CreateCertResResponseMetadataError `json:"Error,omitempty"`
}

type CreateCertResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type CreateCertResResult struct {
	AccountID *string `json:"AccountID,omitempty"`

	// 证书链 ID。
	ChainID *string `json:"ChainID,omitempty"`

	// 使用该证书的域名。
	Domain *string `json:"Domain,omitempty"`

	// 证书用途，取值及含义如下所示。
	// * https：用于 HTTPS 加密；
	// * sign：用于签名加密。
	UseWay *string `json:"UseWay,omitempty"`
}

type CreateDomainV2Body struct {

	// REQUIRED; 待添加到视频直播服务进行加速的域名列表信息。 :::tip 一个域名空间下最多包含 10 个域名。 :::
	Domains []CreateDomainV2BodyDomainsItem `json:"Domains"`

	// REQUIRED; 域名空间，是一组关联域名的集合，由字母（A - Z、a -z）、数字（0 - 9）和连字符（-） 组成。您可以自定义新的域名空间或调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口获取已有域名使用的域名空间。
	Vhost string `json:"Vhost"`

	// 为域名空间设置所属项目，默认为空表示归属到 default 项目， 新建域名空间时需要为域名空间绑定项目，您可以在访问控制-项目列表 [https://console.volcengine.com/iam/resourcemanage/project]查看已有项目并对项目进行管理。
	// 项目是火山引擎提供的一种资源管理方式，您可以对不同业务或项目使用的云资源进行分组管理，以实现基于项目的账单查看、子账号授权等功能。
	// * 新建域名空间时，需为域名空间设置所属项目。每个域名空间只能属于一个项目，选择已有的域名空间时，项目不可配置；
	// * 使用基于项目的账单查看需提前开通分账服务，请前往账单管理-分账账单 [https://console.volcengine.com/finance/bill/split-bill/]进行服务开通；
	// * 使用基于项目的子账号授权请参见使用 IAM 用户进行项目权限划分 [https://www.volcengine.com/docs/6469/1166573]。
	ProjectName *string `json:"ProjectName,omitempty"`

	// 为域名空间设置资源标签。您可以通过资源标签从不同维度对云资源进行分类和聚合管理，如资源分账等场景。 :::tip 如需使用标签进行资源分账，可以在可以在账单管理-费用标签 [https://console.volcengine.com/finance/bill/tag/]处管理启用标签，将对应标签运用到账单明细等数据中。
	// :::
	Tags []*CreateDomainV2BodyTagsItem `json:"Tags,omitempty"`

	// 是否进行域名归属校验，不填默认需要校验
	VerifyCheck *bool `json:"VerifyCheck,omitempty"`
}

type CreateDomainV2BodyDomainsItem struct {

	// REQUIRED; 域名名称，域名由字母（A - Z、a -z）、数字（0 - 9）和连字符（-） 组成，长度为 1 到 60 个字符。
	DomainName string `json:"DomainName"`

	// REQUIRED; 国内可传入：
	// * cn 国内
	// * cn-oversea 海外及港澳台
	// * cn-global 全球 byteplus可传入：
	// * cn 中国
	// * oversea 非中国区
	// * global 全球
	Region string `json:"Region"`

	// REQUIRED; 域名类型，取值及含义如下所示。
	// * push：推流域名；
	// * pull-flv：拉流域名。
	Type string `json:"Type"`

	// HTTPS 证书链 ID，默认为空表示不为域名绑定 HTTPS 证书。您可以调用 ListCertV2 [https://www.volcengine.com/docs/6469/1126823] 接口或在视频直播控制台的证书管理 [https://console.volcengine.com/live/main/certificate]页面，获取待绑定的证书链
	// ID。
	ChainID *string `json:"ChainID,omitempty"`
}

type CreateDomainV2BodyTagsItem struct {

	// REQUIRED; 标签类型，支持以下取值。
	// * System：系统内置标签；
	// * Custom：自定义标签。
	Category string `json:"Category"`

	// REQUIRED; 标签 Key 值。
	Key string `json:"Key"`

	// REQUIRED; 标签 Value 值。
	Value string `json:"Value"`
}

type CreateDomainV2Res struct {

	// REQUIRED
	ResponseMetadata CreateDomainV2ResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type CreateDomainV2ResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                  `json:"Version"`
	Error   *CreateDomainV2ResResponseMetadataError `json:"Error,omitempty"`
}

type CreateDomainV2ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type CreatePullToPushTaskBody struct {

	// REQUIRED; 任务的结束时间，Unix 时间戳，单位为秒。 :::tip 拉流转推任务持续时间最长为 7 天。 :::
	EndTime int32 `json:"EndTime"`

	// REQUIRED; 设置点播视频转推至第三方推流域名时是否使用推流优先级参数，缺省情况下表示不使用此参数，支持的取值及含义如下。
	// * true：使用；
	//
	//
	// * false：不使用。 :::tip
	//
	//
	// * 使用点播视频转推直播实现视频循环播放（轮播）时，支持使用带有推流优先级参数的推流地址进行推流，如在第一个点播视频的推流地址后添加 pri=10、在第二个点播视频的推流地址后添加 pri=11，可达到使用推流优先级高的流替换推流优先级低的流的目的。相比不使用推流优先级参数时可实现更平滑的轮播视频切换。
	//
	//
	// * 推流至非第三方域名时，默认支持使用带有推流优先级参数的推流地址。
	//
	//
	// * 推流至第三方域名时，如需使用推流优先级参数实现新流替换旧流时，需在创建拉流转推时为推流域名开启推流优先级参数配置开关。 :::
	PushPriority bool `json:"PushPriority"`

	// REQUIRED; 任务的开始时间，Unix 时间戳，单位为秒。 :::tip 拉流转推任务持续时间最长为 7 天。 :::
	StartTime int32 `json:"StartTime"`

	// REQUIRED; 拉流来源类型，支持的取值及含义如下。
	// * 0：直播源；
	// * 1：点播视频。
	Type int32 `json:"Type"`

	// 推流应用名称，推流地址（DstAddr）为空时必传；反之，则该参数不生效。
	App *string `json:"App,omitempty"`

	// 接收拉流转推任务状态回调的地址，最大长度为 2000 个字符，默认为空。
	CallbackURL *string `json:"CallbackURL,omitempty"`

	// 续播策略，续播策略指转推点播视频进行直播时出现断流并恢复后，如何继续播放的策略，拉流来源类型为点播视频（Type 为 1）时参数生效，支持的取值及含义如下。
	// * 0：从断流处续播（默认值）；
	// * 1：从断流处+自然流逝时长处续播。
	ContinueStrategy *int32 `json:"ContinueStrategy,omitempty"`

	// 点播视频文件循环播放模式，当拉流来源类型为点播视频时为必选参数，参数取值及含义如下所示。
	// * -1：无限次循环，至任务结束；
	// * 0：有限次循环，循环次数以 PlayTimes 取值为准；
	// * >0：有限次循环，循环次数以 CycleMode 取值为准。
	CycleMode *int32 `json:"CycleMode,omitempty"`

	// 推流域名，推流地址（DstAddr）为空时必传；反之，则该参数不生效。
	Domain *string `json:"Domain,omitempty"`

	// 推流地址，即直播源或点播视频转推的目标地址。
	DstAddr *string `json:"DstAddr,omitempty"`

	// 点播文件启播时间偏移值，单位为秒，仅当点播视频播放地址列表（SrcAddrS）只有一个地址，且未配置 Offsets 时生效，缺省情况下为空表示不进行偏移。
	Offset *float32 `json:"Offset,omitempty"`

	// 点播文件启播时间偏移值，单位为秒，数量与拉流地址列表中地址数量相等，缺省情况下为空表示不进行偏移。拉流来源类型为点播视频时，参数生效。
	OffsetS []*float32 `json:"OffsetS,omitempty"`

	// 点播视频文件循环播放次数，当 CycleMode 取值为 0 时，PlayTimes 取值将作为循环播放次数。 :::tip PlayTimes 为冗余参数，您可以将 PlayTimes 置 0 后直接使用 CycleMode 指定点播视频文件循环播放次数。
	// :::
	PlayTimes *int32 `json:"PlayTimes,omitempty"`

	// 是否开启点播预热，开启点播预热后，系统会自动将点播视频文件缓存到 CDN 节点上，当用户请求直播时，可以直播从 CDN 节点获取视频，从而提高直播流畅度。拉流来源类型为点播视频时，参数生效。
	// * 0：不开启；
	// * 1：开启（默认值）。
	PreDownload *int32 `json:"PreDownload,omitempty"`

	// 直播源的拉流地址，拉流来源类型为直播源时，为必传参数，最大长度为 1000 个字符。
	SrcAddr *string `json:"SrcAddr,omitempty"`

	// 点播视频播放地址列表，拉流来源类型为点播视频时，为必传参数，最多支持传入 30 个点播视频播放地址，每个地址最大长度为 1000 个字符。
	SrcAddrS []*string `json:"SrcAddrS,omitempty"`

	// 推流的流名称，推流地址（DstAddr）为空时必传；反之，则该参数不生效。
	Stream *string `json:"Stream,omitempty"`

	// 拉流转推任务的名称，默认为空表示不配置任务名称。支持由中文、大小写字母（A - Z、a - z）和数字（0 - 9）组成，长度为 1 到 20 各字符。
	Title *string `json:"Title,omitempty"`

	// 为拉流转推视频添加的水印配置信息。
	Watermark *CreatePullToPushTaskBodyWatermark `json:"Watermark,omitempty"`
}

// CreatePullToPushTaskBodyWatermark - 为拉流转推视频添加的水印配置信息。
type CreatePullToPushTaskBodyWatermark struct {

	// REQUIRED; 水印图片字符串，图片最大 2MB，最小 100Bytes，最大分辨率为 1080×1080。图片 Data URL 格式为：data:image/<mediatype>;base64,<data>。
	// * mediatype：图片类型，支持 png、jpg、jpeg 格式；
	// * data：base64 编码的图片字符串。
	// 例如，data:image/png;base64,iVBORw0KGg****mCC
	Picture string `json:"Picture"`

	// REQUIRED; 水印宽度占直播原始画面宽度百分比，支持精度为小数点后两位。
	Ratio float32 `json:"Ratio"`

	// REQUIRED; 水平偏移，表示水印左侧边与转码流画面左侧边之间的距离，使用相对比率，取值范围为 [0,1)。
	RelativePosX float32 `json:"RelativePosX"`

	// REQUIRED; 垂直偏移，表示水印顶部边与转码流画面顶部边之间的距离，使用相对比率，取值范围为 [0,1)。
	RelativePosY float32 `json:"RelativePosY"`
}

type CreatePullToPushTaskRes struct {

	// REQUIRED
	ResponseMetadata CreatePullToPushTaskResResponseMetadata `json:"ResponseMetadata"`
	Result           *CreatePullToPushTaskResResult          `json:"Result,omitempty"`
}

type CreatePullToPushTaskResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                        `json:"Version"`
	Error   *CreatePullToPushTaskResResponseMetadataError `json:"Error,omitempty"`
}

type CreatePullToPushTaskResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type CreatePullToPushTaskResResult struct {

	// REQUIRED; 任务 ID，任务的唯一标识。
	TaskID string `json:"TaskId"`
}

type CreateRecordPresetV2Body struct {

	// REQUIRED; 直播流录制配置的详细配置。
	RecordPresetConfig CreateRecordPresetV2BodyRecordPresetConfig `json:"RecordPresetConfig"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，查看需要录制的直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 应用名称，取值与直播流地址的 AppName 字段取值相同，支持填写星号（*）或由 1 到 30 位数字（0 - 9）、大写小字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，默认为空。 :::tip
	// * App 取值为空时，Stream 取值也需为空，表示录制配置为 Vhost 级别的全局配置。
	// * App 取值不为空时，Stream 取值含义请参见 Stream 参数说明。 :::
	App *string `json:"App,omitempty"`

	// 流名称，取值与直播流地址的 StreamName 字段取值相同，支持填写星号（*）或由 1 到 100 位数字（0 - 9）、字母、下划线（_）、短横线（-）和句点（.）组成。
	// :::tip
	// * App 取值不为空、Stream 取值为空时，表示录制配置为 Vhost + App 级别的配置。
	// * App 取值不为空、Stream 取值不为空时，表示录制为 Vhost + App + Stream 的配置。 :::
	Stream *string `json:"Stream,omitempty"`
}

// CreateRecordPresetV2BodyRecordPresetConfig - 直播流录制配置的详细配置。
type CreateRecordPresetV2BodyRecordPresetConfig struct {

	// 录制为 FLV 格式时的录制参数。 :::tip 您需至少配置一种录制格式，即 FlvParam、HlsParam、Mp4Param 至少开启一个。 :::
	FlvParam *CreateRecordPresetV2BodyRecordPresetConfigFlvParam `json:"FlvParam,omitempty"`

	// 录制为 HLS 合适时的录制参数。 :::tip 您需至少配置一种录制格式，即 FlvParam、HlsParam、Mp4Param 至少开启一个。 :::
	HlsParam *CreateRecordPresetV2BodyRecordPresetConfigHlsParam `json:"HlsParam,omitempty"`

	// 录制为 MP4 格式时的录制参数。 :::tip 您需至少配置一种录制格式，即 FlvParam、HlsParam、Mp4Param 至少开启一个。 :::
	Mp4Param *CreateRecordPresetV2BodyRecordPresetConfigMp4Param `json:"Mp4Param,omitempty"`

	// 是否源流录制，默认值为 0，支持的取值及含义如下所示。
	// * 0：不录制；
	// * 1：录制。
	// :::tip 转码流和源流需至少选一个进行录制，即是否录制转码流（TranscodeRecord）和是否录制源流（OriginRecord）的取值至少一个不为 0。 :::
	OriginRecord *int32 `json:"OriginRecord,omitempty"`

	// 录制为 HLS 格式时，单个 TS 切片时长，单位为秒，默认值为 10，取值范围为 [5,30]。
	SliceDuration *int32 `json:"SliceDuration,omitempty"`

	// 是否录制转码流，默认值为 0，支持的取值及含义如下所示。
	// * 0：不录制；
	// * 1：录制全部转码流；
	// * 2：录制指定转码流，即通过转码后缀列表 TranscodeSuffixList匹配转码流进行录制，如果转码流后缀列表为空仍表示录制全部转码流。
	// :::tip 转码流和源流需至少选一个进行录制，即是否录制转码流（TranscodeRecord）和是否录制源流（OriginRecord）的取值至少一个不为 0。 :::
	TranscodeRecord *int32 `json:"TranscodeRecord,omitempty"`

	// 转码流后缀列表，转码流录制配置为根据转码流列表匹配（TranscodeRecord 取值为 2）时生效，TranscodeSuffixList 默认配置为空，效果等同于录制全部转码流。
	TranscodeSuffixList []*string `json:"TranscodeSuffixList,omitempty"`
}

// CreateRecordPresetV2BodyRecordPresetConfigFlvParam - 录制为 FLV 格式时的录制参数。 :::tip 您需至少配置一种录制格式，即 FlvParam、HlsParam、Mp4Param
// 至少开启一个。 :::
type CreateRecordPresetV2BodyRecordPresetConfigFlvParam struct {

	// 实时录制场景下，断流等待时长，单位为秒，默认值为 180，取值范围为 [0,3600]。如果实际断流时间小于断流等待时长，录制任务不会停止；如果实际断流时间大于断流等待时长，录制任务会停止，断流恢复后重新开始一个新的录制任务。
	ContinueDuration *int32 `json:"ContinueDuration,omitempty"`

	// 断流录制场景下，单文件录制时长，单位为秒，默认值为 7200，取值范围为 -1 和 [300,86400]。
	// * 取值为 -1 时，表示不限制录制时长，录制结束后生成一个完整的录制文件。
	// * 取值为 [300,86400] 之间的值时，表示根据设置的录制文件时分段长生成录制文件，完成录制后一起上传。
	// :::tip 断流录制场景仅在录制格式为 HLS 时生效，且断流录制和实时录制为二选一配置。 :::
	Duration *int32 `json:"Duration,omitempty"`

	// 当前格式的录制是否开启，默认值为 false，支持的取值及含义如下所示。
	// * false：不开启；
	// * true：开启。
	Enable *bool `json:"Enable,omitempty"`

	// 实时录制场景下，单文件录制时长，单位为秒，默认值为 1800，取值范围为 [300,21600]。录制时间到达设置的单文件录制时长时，会立即生成录制文件实时上传存储。
	RealtimeRecordDuration *int32 `json:"RealtimeRecordDuration,omitempty"`

	// 断流录制场景下，断流拼接时长，单位为秒，默认值为 0，支持的取值及含义如下所示。
	// * -1：一直拼接，表示每次断流都不会影响录制任务，录制完成后生成一个完整的录制文件；
	// * 0：不拼接，表示每次断流结束录制任务生成一个录制文件，断流恢复重新开始一个新的录制任务；
	// * 大于 0：拼接容错时间，表示如果断流时间小于拼接容错时间时，则录制任务不会停止，不会生成新的录制文件；如果断流时间大于拼接容错时间，则录制任务停止，断流恢复后重新开始一个新的录制任务。
	// :::tip 断流录制场景仅在录制格式为 HLS 时生效，且断流录制和实时录制为二选一配置。 :::
	Splice *int32 `json:"Splice,omitempty"`

	// TOS 存储相关配置。 :::tip 录制文件只能选择一个位置进行存储，即 TOSParam 和 VODParam 配置且配置其中一个。 :::
	TOSParam *CreateRecordPresetV2BodyRecordPresetConfigFlvParamTOSParam `json:"TOSParam,omitempty"`

	// VOD 存储相关配置。 :::tip 录制文件只能选择一个位置进行存储，即 TOSParam 和 VODParam 配置且配置其中一个。 :::
	VODParam *CreateRecordPresetV2BodyRecordPresetConfigFlvParamVODParam `json:"VODParam,omitempty"`
}

// CreateRecordPresetV2BodyRecordPresetConfigFlvParamTOSParam - TOS 存储相关配置。 :::tip 录制文件只能选择一个位置进行存储，即 TOSParam 和 VODParam
// 配置且配置其中一个。 :::
type CreateRecordPresetV2BodyRecordPresetConfigFlvParamTOSParam struct {

	// TOS 存储对应的 Bucket。例如，存储位置为 live-test-tos-example/live/liveapp 时，Bucket 取值为 live-test-tos-example。 :::tip 如果使用 TOS 存储，即 TOSParam
	// 中 Enable 取值为 true 时，Bucket 为必填。 :::
	Bucket *string `json:"Bucket,omitempty"`

	// 是否使用 TOS 存储，默认为 false，取值及含义如下所示。
	// * false：不使用；
	// * true：使用。
	Enable *bool `json:"Enable,omitempty"`

	// 录制文件的存储规则，最大长度为 200 个字符，支持以record/{PubDomain}/{App}/{Stream}/{StartTime}_{EndTime} 样式设置存储规则，支持输入字母（A - Z、a - z）、数字（0 -
	// 9）、短横线（-）、叹号（!）、下划线（_）、句点（.）、星号（*）及占位符。
	// 存储规则设置注意事项如下。
	// * 目录层级至少包含2级及以上，如live/{App}/{Stream}。
	// * record 为自定义字段；
	// * {PubDomain} 取值为当前配置的 vhost 值；
	// * {App} 取值为当前配置的 AppName 值；
	// * {Stream} 取值为当前配置的 StreamName 值；
	// * {StartTime} 取值为录制的开始时间戳；
	// * {EndTime} 取值为录制的结束时间戳。
	ExactObject *string `json:"ExactObject,omitempty"`

	// TOS 存储对应 Bucket 下的存储目录，默认为空。例如，存储位置为 live-test-tos-example/live/liveapp 时，StorageDir 取值为 live/liveapp。
	StorageDir *string `json:"StorageDir,omitempty"`
}

// CreateRecordPresetV2BodyRecordPresetConfigFlvParamVODParam - VOD 存储相关配置。 :::tip 录制文件只能选择一个位置进行存储，即 TOSParam 和 VODParam
// 配置且配置其中一个。 :::
type CreateRecordPresetV2BodyRecordPresetConfigFlvParamVODParam struct {

	// 直播录制文件存储到点播时的视频分类 ID，您可以通过视频点播的ListVideoClassifications [https://www.volcengine.com/docs/4/101661]接口查询视频分类 ID 等信息，默认为空。
	ClassificationID *int32 `json:"ClassificationID,omitempty"`

	// 是否使用 VOD 存储，默认为 false，支持的取值及含义如下所示。
	// * false：不使用；
	// * true：使用。
	Enable *bool `json:"Enable,omitempty"`

	// 录制文件的存储规则，最大长度为 200 个字符，支持以record/{PubDomain}/{App}/{Stream}/{StartTime}_{EndTime} 样式设置存储规则，支持输入字母（A - Z、a - z）、数字（0 -
	// 9）、短横线（-）、叹号（!）、下划线（_）、句点（.）、星号（*）及占位符。
	// 存储规则设置注意事项如下。
	// * 目录层级至少包含2级及以上，如live/{App}/{Stream}。
	// * record 为自定义字段；
	// * {PubDomain} 取值为当前配置的 vhost 值；
	// * {App} 取值为当前配置的 AppName 值；
	// * {Stream} 取值为当前配置的 StreamName 值；
	// * {StartTime} 取值为录制的开始时间戳；
	// * {EndTime} 取值为录制的结束时间戳。
	ExactObject *string `json:"ExactObject,omitempty"`

	// 直播录制文件存储到点播时的存储类型，存储类型介绍请参考媒资存储管理 [https://www.volcengine.com/docs/4/73629#媒资存储]。默认值为 1，支持的取值及含义如下所示。
	// * 1：标准存储；
	// * 2：归档存储。
	StorageClass *int32 `json:"StorageClass,omitempty"`

	// 视频点播（VOD）空间名称。可登录视频点播控制台 [https://console.volcengine.com/vod/]查询。 :::tip 如果使用 VOD 存储，即 VODParam 中 Enable 取值为 true 时，VodNamespace
	// 为必填。 :::
	VodNamespace *string `json:"VodNamespace,omitempty"`

	// 视频点播工作流模板 ID，对于存储在点播的录制文件，会使用该工作流模版对录制的视频进行处理，可登录视频点播控制台 [https://console.volcengine.com/vod/]获取工作流模板 ID，默认为空。
	WorkflowID *string `json:"WorkflowID,omitempty"`
}

// CreateRecordPresetV2BodyRecordPresetConfigHlsParam - 录制为 HLS 合适时的录制参数。 :::tip 您需至少配置一种录制格式，即 FlvParam、HlsParam、Mp4Param
// 至少开启一个。 :::
type CreateRecordPresetV2BodyRecordPresetConfigHlsParam struct {

	// 断流等待时长，取值范围[0,3600]
	ContinueDuration *int32 `json:"ContinueDuration,omitempty"`

	// 断流录制单文件录制时长，单位为秒，默认值为 7200，取值范围为 -1，[300,86400]，-1 表示一直录制，目前只对 HLS生效.
	Duration *int32 `json:"Duration,omitempty"`

	// 当前格式的录制是否开启，默认 false，取值及含义如下所示。
	// * false：不开启；
	// * true：开启。
	Enable *bool `json:"Enable,omitempty"`

	// 实时录制文件时长，单位为 s，取值范围为 [300,21600]
	RealtimeRecordDuration *int32 `json:"RealtimeRecordDuration,omitempty"`

	// 断流拼接间隔时长，对实时录制无效，单位为 s，默认值为 0。支持的取值如下所示。
	// * -1：一直拼接；
	// * 0：不拼接；
	// * 大于 0：断流拼接时间间隔，对 HLS 录制生效。
	Splice *int32 `json:"Splice,omitempty"`

	// TOS 存储相关配置 :::tipTOSParam和VODParam配置且配置其中一个。 :::
	TOSParam *CreateRecordPresetV2BodyRecordPresetConfigHlsParamTOSParam `json:"TOSParam,omitempty"`

	// VOD 存储相关配置 :::tipTOSParam和VODParam配置且配置其中一个。 :::
	VODParam *CreateRecordPresetV2BodyRecordPresetConfigHlsParamVODParam `json:"VODParam,omitempty"`
}

// CreateRecordPresetV2BodyRecordPresetConfigHlsParamTOSParam - TOS 存储相关配置 :::tipTOSParam和VODParam配置且配置其中一个。 :::
type CreateRecordPresetV2BodyRecordPresetConfigHlsParamTOSParam struct {

	// TOS 存储空间，一般使用 CDN 对应的 Bucket :::tip 如果 TOSParam 中的 Enable 取值为 true，则 Bucket 必填。 :::
	Bucket *string `json:"Bucket,omitempty"`

	// 是否使用 TOS 存储，默认为 false，取值及含义如下所示。
	// * false：不使用；
	// * true：使用。
	Enable *bool `json:"Enable,omitempty"`

	// 录制文件的存储位置。存储路径为record/{PubDomain}/{App}/{Stream}/{StartTime}_{EndTime}
	ExactObject *string `json:"ExactObject,omitempty"`

	// TOS 存储目录，默认为空
	StorageDir *string `json:"StorageDir,omitempty"`
}

// CreateRecordPresetV2BodyRecordPresetConfigHlsParamVODParam - VOD 存储相关配置 :::tipTOSParam和VODParam配置且配置其中一个。 :::
type CreateRecordPresetV2BodyRecordPresetConfigHlsParamVODParam struct {

	// 直播录制文件存储到点播时的视频分类 ID，您可以通过视频点播的ListVideoClassifications [https://www.volcengine.com/docs/4/101661]接口查询视频分类 ID 等信息。
	ClassificationID *int32 `json:"ClassificationID,omitempty"`

	// 是否使用 VOD 存储，默认为 false，取值及含义如下所示。
	// * false：不使用；
	// * true：使用。
	Enable *bool `json:"Enable,omitempty"`

	// 录制文件的存储位置，最大长度为 200 个字符。默认的存储位置为record/{PubDomain}/{App}/{Stream}/{StartTime}_{EndTime}，参数格式要求如下所示。
	// * 支持删除固定路径，如 {App}/{Stream}；
	// * 不支持以正斜线（/）或者反斜线（\）开头；
	// * 不支持 “//” 和 “/./” 等字符串；
	// * 不支持 \b、\t、\n、\v、\f、\r 等字符；
	// * 不支持 “..” 作为文件名；
	// * 目录层级至少包含 2 级及以上，如live/{App}/{Stream}。
	ExactObject *string `json:"ExactObject,omitempty"`

	// 直播录制文件存储到点播时的存储类型。默认值为 1，支持的取值及含义如下所示。
	// * 1：标准存储；
	// * 2：归档存储。
	StorageClass *int32 `json:"StorageClass,omitempty"`

	// 视频点播（VOD）空间名称。可登录视频点播控制台 [https://console.volcengine.com/vod/]查询 :::tip 如果 VODParam 中的 Enable 取值为 true，则 VodNamespace 必填。
	// :::
	VodNamespace *string `json:"VodNamespace,omitempty"`

	// 工作流模版 ID，对于存储在点播的录制文件，会使用该工作流模版对视频进行处理。可登录视频点播控制台 [https://console.volcengine.com/vod/]获取 ID
	WorkflowID *string `json:"WorkflowID,omitempty"`
}

// CreateRecordPresetV2BodyRecordPresetConfigMp4Param - 录制为 MP4 格式时的录制参数。 :::tip 您需至少配置一种录制格式，即 FlvParam、HlsParam、Mp4Param
// 至少开启一个。 :::
type CreateRecordPresetV2BodyRecordPresetConfigMp4Param struct {

	// 断流等待时长，取值范围[0,3600]
	ContinueDuration *int32 `json:"ContinueDuration,omitempty"`

	// 断流录制单文件录制时长，单位为 s，默认值为 7200，取值范围为 -1，[300,86400]，-1表示一直录制，目前只对HLS生效
	Duration *int32 `json:"Duration,omitempty"`

	// 当前格式的录制是否开启，默认 false，取值及含义如下所示。
	// * false：不开启；
	// * true：开启。
	Enable *bool `json:"Enable,omitempty"`

	// 实时录制文件时长，单位为 s，取值范围为 [300,21600]
	RealtimeRecordDuration *int32 `json:"RealtimeRecordDuration,omitempty"`

	// 断流拼接间隔时长，对实时录制无效，单位为 s，默认值为 0。支持的取值如下所示。
	// * -1：一直拼接；
	// * 0：不拼接；
	// * 大于 0：断流拼接时间间隔，对 HLS 录制生效。
	Splice *int32 `json:"Splice,omitempty"`

	// TOS 存储相关配置 :::tipTOSParam和VODParam配置且配置其中一个。 :::
	TOSParam *CreateRecordPresetV2BodyRecordPresetConfigMp4ParamTOSParam `json:"TOSParam,omitempty"`

	// VOD 存储相关配置 :::tipTOSParam和VODParam配置且配置其中一个。 :::
	VODParam *CreateRecordPresetV2BodyRecordPresetConfigMp4ParamVODParam `json:"VODParam,omitempty"`
}

// CreateRecordPresetV2BodyRecordPresetConfigMp4ParamTOSParam - TOS 存储相关配置 :::tipTOSParam和VODParam配置且配置其中一个。 :::
type CreateRecordPresetV2BodyRecordPresetConfigMp4ParamTOSParam struct {

	// TOS 存储空间，一般使用 CDN 对应的 Bucket :::tip 如果 TOSParam 中的 Enable 取值为 true，则 Bucket 必填。 :::
	Bucket *string `json:"Bucket,omitempty"`

	// 是否使用 TOS 存储，默认为 false，取值及含义如下所示。
	// * false：不使用；
	// * true：使用。
	Enable *bool `json:"Enable,omitempty"`

	// 录制文件的存储位置。存储路径为record/{PubDomain}/{App}/{Stream}/{StartTime}_{EndTime}
	ExactObject *string `json:"ExactObject,omitempty"`

	// TOS 存储目录，默认为空
	StorageDir *string `json:"StorageDir,omitempty"`
}

// CreateRecordPresetV2BodyRecordPresetConfigMp4ParamVODParam - VOD 存储相关配置 :::tipTOSParam和VODParam配置且配置其中一个。 :::
type CreateRecordPresetV2BodyRecordPresetConfigMp4ParamVODParam struct {

	// 直播录制文件存储到点播时的视频分类 ID，您可以通过视频点播的ListVideoClassifications [https://www.volcengine.com/docs/4/101661]接口查询视频分类 ID 等信息。
	ClassificationID *int32 `json:"ClassificationID,omitempty"`

	// 是否使用 VOD 存储，默认为 false，取值及含义如下所示。
	// * false：不使用；
	// * true：使用。
	Enable *bool `json:"Enable,omitempty"`

	// 录制文件的存储位置，最大长度为 200 个字符。默认的存储位置为record/{PubDomain}/{App}/{Stream}/{StartTime}_{EndTime}，参数格式要求如下所示。
	// * 支持删除固定路径，如 {App}/{Stream}；
	// * 不支持以正斜线（/）或者反斜线（\）开头；
	// * 不支持 “//” 和 “/./” 等字符串；
	// * 不支持 \b、\t、\n、\v、\f、\r 等字符；
	// * 不支持 “..” 作为文件名；
	// * 目录层级至少包含 2 级及以上，如live/{App}/{Stream}。
	ExactObject *string `json:"ExactObject,omitempty"`

	// 直播录制文件存储到点播时的存储类型。默认值为 1，支持的取值及含义如下所示。
	// * 1：标准存储；
	// * 2：归档存储。
	StorageClass *int32 `json:"StorageClass,omitempty"`

	// 视频点播（VOD）空间名称。可登录视频点播控制台 [https://console.volcengine.com/vod/]查询 :::tip 如果 VODParam 中的 Enable 取值为 true，则 VodNamespace 必填。
	// :::
	VodNamespace *string `json:"VodNamespace,omitempty"`

	// 工作流模版 ID，对于存储在点播的录制文件，会使用该工作流模版对视频进行处理。可登录视频点播控制台 [https://console.volcengine.com/vod/]获取 ID
	WorkflowID *string `json:"WorkflowID,omitempty"`
}

type CreateRecordPresetV2Res struct {

	// REQUIRED
	ResponseMetadata CreateRecordPresetV2ResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type CreateRecordPresetV2ResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                        `json:"Version"`
	Error   *CreateRecordPresetV2ResResponseMetadataError `json:"Error,omitempty"`
}

type CreateRecordPresetV2ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type CreateRelaySourceV4Body struct {

	// REQUIRED; 应用名称，即直播流地址的AppName字段取值，支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 拉流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 回源地址列表，支持输入 RTMP、FLV、HLS 和 SRT 协议的直播推流地址。 :::tip
	// * 当源站使用了非默认端口时，支持在回源地址中以域名:端口或IP:端口的形式配置端口。
	// * 最多支持添加 10 个回源地址，回源失败时，将按照您添加的地址顺序轮循尝试。 :::
	SrcAddrS []string `json:"SrcAddrS"`

	// REQUIRED; 流名称，即直播流地址的StreamName字段取值，支持由大小写字母（A - Z、a - z）、数字（0 - 9）、字母、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	Stream string `json:"Stream"`

	// 回源结束时间，Unix 时间戳，单位为秒。 :::tip
	// * 回源开始到结束最大时间跨度为 7 天；
	// * 开始时间与结束时间同时缺省，表示永久回源。 :::
	EndTime *int32 `json:"EndTime,omitempty"`

	// 回源开始时间，Unix 时间戳，单位为秒。 :::tip
	// * 回源开始到结束最大时间跨度为 7 天；
	// * 开始时间与结束时间同时缺省，表示永久回源。 :::
	StartTime *int32 `json:"StartTime,omitempty"`
}

type CreateRelaySourceV4Res struct {

	// REQUIRED
	ResponseMetadata CreateRelaySourceV4ResResponseMetadata `json:"ResponseMetadata"`
	Result           *CreateRelaySourceV4ResResult          `json:"Result,omitempty"`
}

type CreateRelaySourceV4ResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                       `json:"Version"`
	Error   *CreateRelaySourceV4ResResponseMetadataError `json:"Error,omitempty"`
}

type CreateRelaySourceV4ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type CreateRelaySourceV4ResResult struct {

	// REQUIRED; 固定回源配置的 ID。
	TaskID string `json:"TaskId"`
}

type CreateSnapshotPresetBody struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 截图间隔时间，单位为 s，默认值为 10，取值范围为正整数。
	Interval int32 `json:"Interval"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// ToS 存储的 Bucket。 :::tipBucket 与 ServiceID 传且仅传一个。 :::
	Bucket *string `json:"Bucket,omitempty"`

	// 回调详情。
	CallbackDetailList []*CreateSnapshotPresetBodyCallbackDetailListItem `json:"CallbackDetailList,omitempty"`

	// 存储方式为覆盖截图时的存储规则，支持以 {Domain}/{App}/{Stream} 样式设置存储规则，支持输入字母、数字、"-"、"!"、"_"、"."、"*"及占位符。
	OverwriteObject *string `json:"OverwriteObject,omitempty"`

	// veImageX 的服务 ID。 :::tipBucket 与 ServiceID 传且仅传一个。 :::
	ServiceID *string `json:"ServiceID,omitempty"`

	// 截图格式。默认值为 jpeg，支持如下取值。
	// * jpeg
	// * jpg
	SnapshotFormat *string `json:"SnapshotFormat,omitempty"`

	// 存储方式为实时存储时的存储规则，支持以 {Domain}/{App}/{Stream}/{UnixTimestamp} 样式设置存储规则，支持输入字母、数字、"-"、"!"、"_"、"."、"*"及占位符。
	SnapshotObject *string `json:"SnapshotObject,omitempty"`

	// 截图模版状态状态。默认开启。
	// * 1：开启。
	// * 0：关闭。
	Status *int32 `json:"Status,omitempty"`

	// ToS 存储目录，不传为空。
	StorageDir *string `json:"StorageDir,omitempty"`
}

type CreateSnapshotPresetBodyCallbackDetailListItem struct {

	// 回调类型，默认值为 http。
	CallbackType *string `json:"CallbackType,omitempty"`

	// 回调地址。
	URL *string `json:"URL,omitempty"`
}

type CreateSnapshotPresetRes struct {

	// REQUIRED
	ResponseMetadata CreateSnapshotPresetResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type CreateSnapshotPresetResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                        `json:"Version"`
	Error   *CreateSnapshotPresetResResponseMetadataError `json:"Error,omitempty"`
}

type CreateSnapshotPresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type CreateTimeShiftPresetV3Body struct {

	// REQUIRED; 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 最大时移时长，即允许用户回看的最长时间，单位为秒，支持的取值如下所示。
	// * 86400：1 天；
	// * 259200：3 天；
	// * 604800：7 天；
	// * 1296000：15 天。
	MaxShiftTime int32 `json:"MaxShiftTime"`

	// REQUIRED; 时移拉流域名。 :::tip 录制到 TOS 时，需配置直播流对应的拉流域名。录制到 VOD 时，需配置 VOD 的空间对应的分发域名。 :::
	PullDomain string `json:"PullDomain"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看需要时移的直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 用于多码率时移的参数，为json字符串
	MasterFormat *string `json:"MasterFormat,omitempty"`

	// 0表示不需要 1表示需要
	NeedTranscode *int32 `json:"NeedTranscode,omitempty"`

	// 流名称，取值与直播流地址中 StreamName 字段取值相同，默认为空表示当前应用下的所有流都进行时移。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100
	// 个字符。 :::tip 流名称不为空时，表示只为此条流开启时移。通过流名称配置时移时，同一个 App 最多可指定 20 路流开启时移。 :::
	Stream *string `json:"Stream,omitempty"`

	// 时移类型。支持的取值如下所示。
	// * 0：录制时移，即时移复用录制配置，需提前创建对应级别的录制配置；
	// * 1：独立时移，即时移不复用录制配置。
	TimeShiftType *int32 `json:"TimeShiftType,omitempty"`
}

type CreateTimeShiftPresetV3Res struct {

	// REQUIRED
	ResponseMetadata CreateTimeShiftPresetV3ResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type CreateTimeShiftPresetV3ResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                           `json:"Version"`
	Error   *CreateTimeShiftPresetV3ResResponseMetadataError `json:"Error,omitempty"`
}

type CreateTimeShiftPresetV3ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type CreateTranscodePresetBody struct {

	// REQUIRED; 应用名称，取值与直播流地址的 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 转码后缀，支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）和短横线（-）组成，长度为 1 到 10 个字符。
	// 转码后缀通常以流名称后缀的形式来使用，常见的标识有 _sd、_hd、_uhd，例如，当转码配置的标识为 _hd 时，拉取转码流时转码流的流名名称为 源流的流名称_hd。
	SuffixName string `json:"SuffixName"`

	// REQUIRED; 视频编码格式，支持的取值及含义如下所示。
	// * h264：使用 H.264 视频编码格式；
	// * h265：使用 H.265 视频编码格式；
	// * h266：使用 H.266 视频编码格式；
	// * copy：不进行视频转码，所有视频编码参数不生效，视频编码参数包括视频帧率（FPS）、视频码率（VideoBitrate）、分辨率设置（As、Width、Height、ShortSide、LongSide）、GOP 和 BFrames
	// 等。
	Vcodec string `json:"Vcodec"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，查看需要转码的直播流使用的域名所属的域名空间。
	Vhost     string  `json:"Vhost"`
	ALayout   *string `json:"ALayout,omitempty"`
	AProfile  *string `json:"AProfile,omitempty"`
	AR        *int32  `json:"AR,omitempty"`
	AbrMode   *int32  `json:"AbrMode,omitempty"`
	AccountID *string `json:"AccountID,omitempty"`

	// 音频编码格式，默认值为 aac，支持的取值及含义如下所示。
	// * aac：使用 AAC 音频编码格式；
	// * opus：使用 Opus 音频编码格式。
	// * copy：不进行音频转码，所有音频编码参数不生效，音频编码参数包括音频码率（AudioBitrate）等。
	Acodec         *string `json:"Acodec,omitempty"`
	AdvancedParam  *string `json:"AdvancedParam,omitempty"`
	AllowAudioCopy *int32  `json:"AllowAudioCopy,omitempty"`
	AllowVideoCopy *int32  `json:"AllowVideoCopy,omitempty"`
	An             *int32  `json:"An,omitempty"`

	// 视频分辨率自适应模式开关，默认值为 0。支持的取值及含义如下。
	// * 0：关闭视频分辨率自适应；
	// * 1：开启视频分辨率自适应。 :::tip
	// * 关闭视频分辨率自适应模式（As 取值为 0）时，转码配置的视频分辨率取视频宽度（Width）和视频高度（Height）的值对转码视频进行拉伸；
	// * 开启视频分辨率自适应模式（As 取值为 1）时，转码配置的视频分辨率按照短边长度（ShortSide）、长边长度（LongSide）、视频宽度（Width）、视频高度（Height）的优先级取值，另一边等比缩放。 :::
	As *string `json:"As,omitempty"`

	// 音频码率，单位为 kbps，默认值为 128，取值范围为 [0,1000]；取值为 0 时，表示与源流的音频码率相同。
	AudioBitrate *int32 `json:"AudioBitrate,omitempty"`
	AutoTransAb  *int32 `json:"AutoTransAb,omitempty"`
	AutoTransAl  *int32 `json:"AutoTransAl,omitempty"`
	AutoTransAr  *int32 `json:"AutoTransAr,omitempty"`

	// 是否开启转码视频分辨率不超过源流分辨率，默认值为 1 表示开启。开启后，当源流分辨率低于转码配置分辨率时（即源流宽低于转码配置宽且源流高低于转码配置高时），将按源流视频分辨率进行转码。
	// * 0：关闭；
	// * 1：开启。
	AutoTransResolution *int32 `json:"AutoTransResolution,omitempty"`

	// 是否开启转码视频码率不超过源流码率，默认值为 1 表示开启。开启后，当源流码率低于转码配置码率时，将按照源流视频码率进行转码。
	// * 0：关闭；
	// * 1：开启。
	AutoTransVb *int32 `json:"AutoTransVb,omitempty"`

	// 是否开启转码视频帧率不超过源流帧率，默认值为 1 表示开启。开启后，当源流帧率低于转码配置帧率时，将按照源流视频帧率进行转码。
	// * 0：关闭；
	// * 1：开启。
	AutoTransVr *int32 `json:"AutoTransVr,omitempty"`
	BCM         *int32 `json:"BCM,omitempty"`

	// 转码输出视频中 2 个参考帧之间的最大 B 帧数量，默认值为 3，取值为 0 时表示去除 B 帧。
	// 最大 B 帧数量的取值范围根据视频编码格式（Vcodec）的不同有所差异，取值范围如下所示。
	// * 视频编码格式为 H.264 （Vcodec 取值为 h264）时取值范围为 [0,7]；
	// * 视频编码格式为 H.265 或 H.266 （Vcodec 取值为 h265 或 h266）时取值范围为 [0,3]、7、15。
	BFrames  *int32  `json:"BFrames,omitempty"`
	Describe *string `json:"Describe,omitempty"`

	// 动态范围，画质增强类型生效，ParamType=hvq时必填，并且h264只支持SDR
	// * SDR：输出为SDR
	// * HDR：输出为HDR
	DynamicRange *string `json:"DynamicRange,omitempty"`

	// 是否开启智能插帧，只对画质增强类型生效，不填默认为不开启
	// * 0：不开启
	// * 1：开启
	FISwitch *int32 `json:"FISwitch,omitempty"`

	// 视频帧率，单位为 fps，默认值为 25，取值为 0 时表示与源流视频帧率相同。
	// 视频帧率的取值范围根据视频编码格式（Vcodec）的不同有所差异，视频码率的取值范围如下所示。
	// * 视频编码格式为 H.264 或 H.265 （Vcodec 取值为 h264 或 h265）时，视频帧率取值范围为 [0,60]；
	// * 视频编码格式为 H.266 （Vcodec 取值为 h266）时，视频帧率取值范围为 [0,35]。
	FPS *int32 `json:"FPS,omitempty"`

	// IDR 帧之间的最大间隔时间，单位为秒，默认值为 4，取值范围为 [1,20]。
	GOP    *int32 `json:"GOP,omitempty"`
	GopMin *int32 `json:"GopMin,omitempty"`
	HVSPre *bool  `json:"HVSPre,omitempty"`

	// 视频高度，默认值为 0。
	// 视频高度的取值范围根据视频编码格式（Vcodec）的不同所有差异，视频高度取值如下所示。
	// * 视频编码格式为 H.264 或 H.265 （Vcodec 取值为 h264 或 h265）时，取值范围为 [150,1920]；
	// * 视频编码格式为 H.266 （Vcodec 取值为 h266）时，不支持设置 Width 和 Height。
	// :::tip
	// * 当关闭视频分辨率自适应（As 取值为 0）时，转码分辨率将取 Width 和 Height 的值对转码视频进行拉伸；
	// * 当关闭视频分辨率自适应（As 取值为 0）时，Width 和 Height 任一取值为 0 时，转码视频将保持源流尺寸。 :::
	Height *int32 `json:"Height,omitempty"`

	// 长边长度，默认值为 0。配置不同的转码类型（Roi）和视频编码方式（Vcodec）时，短边长度的取值范围存在如下。
	// * 转码类型为标准转码（Roi 取值为 false）时： * 视频编码方式为 H.264 （Vcodec 取值为 h264）时取值范围为 0 和 [150,4096]；
	// * 视频编码方式为 H.265 （Vcodec 取值为 h265）时取值范围为 0 和 [150,7680]；
	// * 视频编码方式为 H.266 （Vcodec 取值为 h266）时取值范围为 0 和 [150,1280]。
	//
	//
	// * 转码类型为极智超清转码（Roi 取值为 true）时： * 视频编码方式为 H.264 或 H.265 （Vcodec 取值为 h264 或 h265）时取值范围为 0 和 [150,1920]。
	//
	//
	// :::tip
	// * 当开启视频分辨率自适应模式时（As 取值为 1）时，参数生效，反之则不生效。
	// * 当开启视频分辨率自适应模式时（As 取值为 1）时，如果 LongSide 、 ShortSide 、Width 、Height 同时取 0，表示保持源流尺寸。 :::
	LongSide    *int32  `json:"LongSide,omitempty"`
	LookAhead   *int32  `json:"LookAhead,omitempty"`
	Modifier    *string `json:"Modifier,omitempty"`
	NvBf        *int32  `json:"NvBf,omitempty"`
	NvCodec     *string `json:"NvCodec,omitempty"`
	NvGop       *int32  `json:"NvGop,omitempty"`
	NvHVSPre    *bool   `json:"NvHVSPre,omitempty"`
	NvLookahead *int32  `json:"NvLookahead,omitempty"`
	NvPercent   *int32  `json:"NvPercent,omitempty"`
	NvPreset    *string `json:"NvPreset,omitempty"`
	NvPriority  *int32  `json:"NvPriority,omitempty"`
	NvProfile   *string `json:"NvProfile,omitempty"`
	NvRefs      *int32  `json:"NvRefs,omitempty"`
	NvTempAQ    *int32  `json:"NvTempAQ,omitempty"`
	Ocr         *bool   `json:"Ocr,omitempty"`

	// 转码模板参数的类型
	// * hvq：表示使用画质增强
	// 选择画质增强时，转码类型默认为极智超清转码（Roi 默认值为 true）。
	// 选择画质增强时，支持使用 shortside 来设置分辨率。
	// * ParamType 取 hvq 时： * 视频编码方式为 H.264 （Vcodec 取值为 h264）时，shortside 取值范围为 0 和 [150,1280]；
	// * 视频编码方式为 H.265 （Vcodec取值为h265）是，shortside 取值范围为 0 和 [150,1280]；
	ParamType    *string `json:"ParamType,omitempty"`
	Preset       *string `json:"Preset,omitempty"`
	PresetKind   *int32  `json:"PresetKind,omitempty"`
	PresetType   *int32  `json:"PresetType,omitempty"`
	Qp           *int32  `json:"Qp,omitempty"`
	RegionConfig *string `json:"RegionConfig,omitempty"`
	Revision     *string `json:"Revision,omitempty"`

	// 转码类型是否为极智超清转码，默认值为 false，取值及含义如下。
	// * true：极智超清转码；
	// * false：标准转码。
	// :::tip 视频编码格式为 H.266 （Vcodec取值为h266）时，转码类型不支持极智超清转码。 :::
	Roi  *bool `json:"Roi,omitempty"`
	SITI *bool `json:"SITI,omitempty"`

	// 使用场景，画质增强时生效，不填不生效
	// * football：足球场景
	SceneType *string `json:"SceneType,omitempty"`

	// 短边长度，默认值为 0。配置不同的转码类型（Roi）和视频编码方式（Vcodec）时，短边长度的取值范围存在如下。
	// * 转码类型为标准转码（Roi 取值为 false）时： * 视频编码方式为 H.264 （Vcodec 取值为 h264）时取值范围为 0 和 [150,2160]；
	// * 视频编码方式为 H.265 （Vcodec 取值为 h265）时取值范围为 0 和 [150,4096]；
	// * 视频编码方式为 H.266 （Vcodec 取值为 h266）时取值范围为 0 和 [150,720]。
	//
	//
	// * 转码类型为极智超清转码（Roi 取值为 true）时： * 视频编码方式为 H.264 或 H.265 （Vcodec 取值为 h264 或 h265）时取值范围为 0 和 [150,1920]。 :::tip
	//
	//
	// * 当开启视频分辨率自适应模式（As 取值为 1）时，参数生效，反之则不生效。
	// * 当开启视频分辨率自适应模式（As 取值为 1）时，如果 LongSide 、 ShortSide 、Width 、Height 同时取 0，表示保持源流尺寸。 :::
	ShortSide *int32 `json:"ShortSide,omitempty"`
	Status    *int32 `json:"Status,omitempty"`

	// 转码停止时长，支持触发方式为拉流转码（TransType 取值为 Pull）时设置，表示断开拉流后转码停止的时长，单位为秒，取值范围为 -1 和 [0,300]，-1 表示不停止转码，默认值为 60。
	StopInterval *int32 `json:"StopInterval,omitempty"`
	Threads      *int32 `json:"Threads,omitempty"`

	// 转码触发方式，默认值为 Pull，支持的取值及含义如下。
	// * Push：推流转码，直播推流后会自动启动转码任务，生成转码流；
	// * Pull：拉流转码，直播推流后，需要主动播放转码流才会启动转码任务，生成转码流。
	TransType       *string                                   `json:"TransType,omitempty"`
	TranscodeStruct *CreateTranscodePresetBodyTranscodeStruct `json:"TranscodeStruct,omitempty"`
	VBRatio         *int32                                    `json:"VBRatio,omitempty"`
	VBVBufSize      *int32                                    `json:"VBVBufSize,omitempty"`
	VBVMaxRate      *int32                                    `json:"VBVMaxRate,omitempty"`
	VLevel          *string                                   `json:"VLevel,omitempty"`
	VPreset         *string                                   `json:"VPreset,omitempty"`
	VProfile        *string                                   `json:"VProfile,omitempty"`
	VR              *int32                                    `json:"VRVr,omitempty"`
	VRBBframes      *int32                                    `json:"VRBBframes,omitempty"`
	VRBHeightNum    *int32                                    `json:"VRBHeightNum,omitempty"`
	VRBPreset       *string                                   `json:"VRBPreset,omitempty"`
	VRBProfile      *string                                   `json:"VRBProfile,omitempty"`
	VRBSuffix       *string                                   `json:"VRBSuffix,omitempty"`
	VRBVb           *int32                                    `json:"VRBVb,omitempty"`
	VRBWidthNum     *int32                                    `json:"VRBWidthNum,omitempty"`
	VRGop           *int32                                    `json:"VRGop,omitempty"`
	VRGopDen        *int32                                    `json:"VRGopDen,omitempty"`
	VRHvspre        *string                                   `json:"VRHvspre,omitempty"`
	VRProjection    *string                                   `json:"VRProjection,omitempty"`
	VRRoi           *string                                   `json:"VRRoi,omitempty"`
	VRTBframes      *int32                                    `json:"VRTBframes,omitempty"`
	VRTPreset       *string                                   `json:"VRTPreset,omitempty"`
	VRTProfile      *string                                   `json:"VRTProfile,omitempty"`
	VRTSuffix       *string                                   `json:"VRTSuffix,omitempty"`
	VRTVb           *int32                                    `json:"VRTVb,omitempty"`
	VRTileMod       *int32                                    `json:"VRTileMod,omitempty"`
	VRateCtrl       *string                                   `json:"VRateCtrl,omitempty"`
	VbThreshold     *string                                   `json:"VbThreshold,omitempty"`
	Vclass          *bool                                     `json:"Vclass,omitempty"`

	// 视频码率，单位为 bps，默认值为 1000000；取值为 0 时，表示与源流的视频码率相同。
	// 视频码率的取值范围根据视频编码格式（Vcodec）的不同有所差异，视频码率的取值范围如下所示。
	// * 视频编码格式为 H.264 或 H.265 （Vcodec 取值为 h264 或 h265）时，视频码率取值范围为 [0,30000000]；
	// * 视频编码格式为 H.266 （Vcodec 取值为 h266）时，视频码率取值范围为 [0,6000000]。
	VideoBitrate *int32  `json:"VideoBitrate,omitempty"`
	Vn           *int32  `json:"Vn,omitempty"`
	Watermark    *string `json:"Watermark,omitempty"`

	// 视频宽度，单位为 px，默认值为 0。
	// 视频宽度的取值范围根据视频编码格式（Vcodec）的不同所有差异，视频宽度取值如下所示。
	// * 视频编码格式为 H.264 或 H.265 （Vcodec 取值为 h264 或 h265）时，取值范围为 [150,1920]；
	// * 视频编码格式为 H.266 （Vcodec 取值为 h266）时，不支持设置 Width 和 Height。
	// :::tip
	// * 当关闭视频分辨率自适应（As 取值为 0）时，转码分辨率将取 Width 和 Height 的值对转码视频进行拉伸；
	// * 当关闭视频分辨率自适应（As 取值为 0）时，Width 和 Height 任一取值为 0 时，转码视频将保持源流尺寸。 :::
	Width *int32 `json:"Width,omitempty"`
}

type CreateTranscodePresetBodyTranscodeStruct struct {

	// Dictionary of
	ABTest       map[string]*ComponentsXsjbgcSchemasCreatetranscodepresetbodyPropertiesTranscodestructPropertiesAbtestAdditionalproperties `json:"ABTest,omitempty"`
	Codec        *string                                                                                                                   `json:"Codec,omitempty"`
	PresetName   *string                                                                                                                   `json:"PresetName,omitempty"`
	StopInterval *int32                                                                                                                    `json:"StopInterval,omitempty"`
	Suffix       *string                                                                                                                   `json:"Suffix,omitempty"`
	Type         *string                                                                                                                   `json:"Type,omitempty"`
}

type CreateTranscodePresetRes struct {

	// REQUIRED
	ResponseMetadata CreateTranscodePresetResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type CreateTranscodePresetResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version   string                                         `json:"Version"`
	Error     *CreateTranscodePresetResResponseMetadataError `json:"Error,omitempty"`
	RequestID *string                                        `json:"RequestID,omitempty"`
}

type CreateTranscodePresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type CreateWatermarkPresetBody struct {

	// REQUIRED; 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 水平偏移，表示水印左侧边与转码流画面左侧边之间的距离，使用相对比率，取值范围为 [0,1]。
	PosX float32 `json:"PosX"`

	// REQUIRED; 垂直偏移，表示水印顶部边与转码流画面顶部边之间的距离，使用相对比率，取值范围为 [0,1]。
	PosY float32 `json:"PosY"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 需要添加水印的直播画面方向，支持 2 种取值。
	// * vertical：竖屏；
	// * horizontal：横屏。 :::tip 该参数属于历史版本参数，预计将于未来移除。建议使用预览背景高度（PreviewHeight）、预览背景宽度（PreviewWidth）参数代替。 :::
	Orientation *string `json:"Orientation,omitempty"`

	// 水印图片编码字符串，图片最大 2MB，最小 100Bytes，最大分辨率为 1080×1080。图片 Data URL 格式为：data:[<mediatype>];[base64],<data>。
	// * mediatype：图片类型，支持 png、jpg、jpeg 格式；
	// * data：base64 编码的图片字符串。
	Picture *string `json:"Picture,omitempty"`

	// 水印图片对应的 HTTP 地址。与水印图片编码字符串字段二选一传入，同时传入时，以水印图片编码字符串参数为准。
	PictureURL *string `json:"PictureUrl,omitempty"`

	// 水印图片预览背景高度，单位为 px。
	PreviewHeight *float32 `json:"PreviewHeight,omitempty"`

	// 水印图片预览背景宽度，单位为 px。
	PreviewWidth *float32 `json:"PreviewWidth,omitempty"`

	// 水印相对高度，水印高度占直播转码流画面高度的比例，取值范围为 [0,1]，水印宽度会随高度等比缩放。与水印相对宽度字段冲突，请选择其中一个传参。
	RelativeHeight *float32 `json:"RelativeHeight,omitempty"`

	// 水印相对宽度，水印宽度占直播转码流画面宽度的比例，取值范围为 [0,1]，水印高度会随宽度等比缩放。与水印相对高度字段冲突，请选择其中一个传参。
	RelativeWidth *float32 `json:"RelativeWidth,omitempty"`

	// 流名称。
	Stream *string `json:"Stream,omitempty"`
}

type CreateWatermarkPresetRes struct {

	// REQUIRED
	ResponseMetadata CreateWatermarkPresetResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type CreateWatermarkPresetResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                         `json:"Version"`
	Error   *CreateWatermarkPresetResResponseMetadataError `json:"Error,omitempty"`
}

type CreateWatermarkPresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DeleteCMAFConfigBody struct {

	// REQUIRED
	Vhost string  `json:"Vhost"`
	App   *string `json:"App,omitempty"`
}

type DeleteCMAFConfigRes struct {

	// REQUIRED
	ResponseMetadata DeleteCMAFConfigResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteCMAFConfigResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type DeleteCallbackBody struct {

	// 应用名称，与创建回调时传的值一致。您可以调用 DescribeCallback [https://www.volcengine.com/docs/6469/1126931] 接口查看待删除回调配置的 App 取值。
	App *string `json:"App,omitempty"`

	// 推流域名，与创建回调时传的值一致。您可以调用 DescribeCallback [https://www.volcengine.com/docs/6469/1126931] 接口查看待删除回调配置的 Domain 取值。
	Domain *string `json:"Domain,omitempty"`

	// 消息类型，与创建回调时传的值一致。您可以调用 DescribeCallback [https://www.volcengine.com/docs/6469/1126931] 接口查看待删除回调配置的 MessageType 取值。
	// * push：推流开始回调；
	// * push_end：推流结束回调；
	// * snapshot：截图回调；
	// * record：录制任务状态回调；
	// * audit_snapshot：截图审核结果回调。
	MessageType *string `json:"MessageType,omitempty"`

	// 域名空间，与创建回调时传的值一致。您可以调用 DescribeCallback [https://www.volcengine.com/docs/6469/1126931] 接口查看待删除回调配置的 Vhost 取值。
	Vhost *string `json:"Vhost,omitempty"`
}

type DeleteCallbackRes struct {

	// REQUIRED
	ResponseMetadata DeleteCallbackResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteCallbackResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                  `json:"Version"`
	Error   *DeleteCallbackResResponseMetadataError `json:"Error,omitempty"`
}

type DeleteCallbackResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DeleteCertBody struct {

	// REQUIRED; 待删除的 HTTPS 证书的证书链 ID，可以通过查询证书列表 [https://www.volcengine.com/docs/6469/1126822]接口获取。
	ChainID string `json:"ChainID"`
}

type DeleteCertRes struct {

	// REQUIRED
	ResponseMetadata DeleteCertResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type DeleteCertResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                              `json:"Version"`
	Error   *DeleteCertResResponseMetadataError `json:"Error,omitempty"`
}

type DeleteCertResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DeleteClusterRateLimitBody struct {

	// REQUIRED
	Vhost  string  `json:"Vhost"`
	App    *string `json:"App,omitempty"`
	Domain *string `json:"Domain,omitempty"`
}

type DeleteClusterRateLimitRes struct {

	// REQUIRED
	ResponseMetadata DeleteClusterRateLimitResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteClusterRateLimitResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestId为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type DeleteDomainBody struct {

	// REQUIRED; 待删除域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看需要待删除域名的信息。
	Domain string `json:"Domain"`
}

type DeleteDomainRes struct {

	// REQUIRED
	ResponseMetadata DeleteDomainResResponseMetadata `json:"ResponseMetadata"`
}

type DeleteDomainResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                `json:"Version"`
	Error   *DeleteDomainResResponseMetadataError `json:"Error,omitempty"`
}

type DeleteDomainResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DeleteFormatAccessRuleBody struct {

	// REQUIRED
	Domain string `json:"Domain"`

	// REQUIRED
	Vhost string `json:"Vhost"`
}

type DeleteFormatAccessRuleRes struct {

	// REQUIRED
	ResponseMetadata DeleteFormatAccessRuleResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteFormatAccessRuleResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestId为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type DeleteHLSConfigBody struct {

	// REQUIRED; 域名空间
	Vhost string `json:"Vhost"`

	// 应用名称
	App *string `json:"App,omitempty"`

	// 服务类型
	ServiceType *string `json:"ServiceType,omitempty"`
}

type DeleteHLSConfigRes struct {

	// REQUIRED
	ResponseMetadata DeleteHLSConfigResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type DeleteHLSConfigResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type DeleteHTTPHeaderConfigBody struct {

	// REQUIRED; HTTP Header 类型，您可以调用 DescribeHTTPHeaderConfig [https://www.volcengine.com/docs/6469/1232744] 接口查看 HTTP Header
	// 配置的 Phase 取值。
	Phase int32 `json:"Phase"`

	// REQUIRED; 域名空间，您可以调用 DescribeHTTPHeaderConfig [https://www.volcengine.com/docs/6469/1232744] 接口查看 HTTP Header 配置的 Vhost
	// 取值。
	Vhost string `json:"Vhost"`

	// 拉流域名，您可以调用 DescribeHTTPHeaderConfig [https://www.volcengine.com/docs/6469/1232744] 接口查看 HTTP Header 配置的 Domain 取值。
	Domain *string `json:"Domain,omitempty"`
}

type DeleteHTTPHeaderConfigRes struct {

	// REQUIRED
	ResponseMetadata DeleteHTTPHeaderConfigResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteHTTPHeaderConfigResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type DeleteIPAccessRuleBody struct {

	// REQUIRED; 推流域名或拉流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取需要删除
	// IP 访问限制的域名。
	Domain string `json:"Domain"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取需要删除
	// IP 访问限制的域名所属的域名空间。
	Vhost string `json:"Vhost"`
}

type DeleteIPAccessRuleRes struct {

	// REQUIRED
	ResponseMetadata DeleteIPAccessRuleResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteIPAccessRuleResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type DeleteLatencyConfigBody struct {

	// REQUIRED
	Domain string `json:"Domain"`
}

type DeleteLatencyConfigRes struct {

	// REQUIRED
	ResponseMetadata DeleteLatencyConfigResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteLatencyConfigResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type DeletePullToPushTaskBody struct {

	// REQUIRED; 任务 ID，任务的唯一标识，您可以通过获取拉流转推任务列表 [https://www.volcengine.com/docs/6469/1126896]接口获取。
	TaskID string `json:"TaskId"`
}

type DeletePullToPushTaskRes struct {

	// REQUIRED
	ResponseMetadata DeletePullToPushTaskResResponseMetadata `json:"ResponseMetadata"`
}

type DeletePullToPushTaskResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                        `json:"Version"`
	Error   *DeletePullToPushTaskResResponseMetadataError `json:"Error,omitempty"`
}

type DeletePullToPushTaskResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DeleteRecordPresetBody struct {

	// REQUIRED; 录制配置的名称。可调用 ListVhostRecordPresetV2 [https://www.volcengine.com/docs/6469/1126858] 接口查看待删除录制配置的名称。
	Preset string `json:"Preset"`

	// 应用名称，您可以调用ListVhostRecordPresetV2 [https://www.volcengine.com/docs/6469/1126858]接口查看待删除的录制配置 App 取值。
	App *string `json:"App,omitempty"`

	// 域名空间。您可以调用 ListVhostRecordPresetV2 [https://www.volcengine.com/docs/6469/1126858] 接口查看待删除录制配置的 Vhost 取值。
	Vhost *string `json:"Vhost,omitempty"`
}

type DeleteRecordPresetRes struct {

	// REQUIRED
	ResponseMetadata DeleteRecordPresetResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type DeleteRecordPresetResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                      `json:"Version"`
	Error   *DeleteRecordPresetResResponseMetadataError `json:"Error,omitempty"`
}

type DeleteRecordPresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DeleteRefererBody struct {

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取需要配置
	// Referer 的拉流域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 应用名称，取值与直播流地址中 AppName 字段取值相同，默认为空，表示所有应用名称。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。 :::tip
	// * 如创建时传了 App，删除时需要传该参数；
	// * 如创建时未传 App，删除时不传该参数。 :::
	App *string `json:"App,omitempty"`

	// 拉流域名。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取需要配置
	// Referer 的拉流域名。 :::tip
	// * 如创建时传了 Domain，删除时需要传该参数；
	// * 如创建时未传 Domain，删除时不传该参数。 :::
	Domain *string `json:"Domain,omitempty"`
}

type DeleteRefererRes struct {

	// REQUIRED
	ResponseMetadata DeleteRefererResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteRefererResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                 `json:"Version"`
	Error   *DeleteRefererResResponseMetadataError `json:"Error,omitempty"`
}

type DeleteRefererResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DeleteRegionAccessRuleBody struct {

	// REQUIRED
	Domain string `json:"Domain"`

	// REQUIRED
	Vhost string  `json:"Vhost"`
	App   *string `json:"App,omitempty"`
}

type DeleteRegionAccessRuleRes struct {

	// REQUIRED
	ResponseMetadata DeleteRegionAccessRuleResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteRegionAccessRuleResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type DeleteRelaySourceRewriteBody struct {

	// 需要设置黑白名单的拉流域名。域名请在工信部完成备案。
	Domain *string `json:"Domain,omitempty"`

	// 域名空间名称
	Vhost *string `json:"Vhost,omitempty"`
}

type DeleteRelaySourceRewriteRes struct {

	// REQUIRED
	ResponseMetadata DeleteRelaySourceRewriteResResponseMetadata `json:"ResponseMetadata"`
	Result           *DeleteRelaySourceRewriteResResult          `json:"Result,omitempty"`
}

type DeleteRelaySourceRewriteResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type DeleteRelaySourceRewriteResResult struct {

	// REQUIRED; 异动列表
	Items []interface{} `json:"Items"`
}

type DeleteRelaySourceV3Body struct {

	// REQUIRED; 直播流使用的域名所属的域名空间，您可以调用DescribeRelaySourceV3 [https://www.volcengine.com/docs/6469/1126874]接口获取待删除配置的 Vhost 取值。
	Vhost string `json:"Vhost"`

	// 应用名称，您可以调用DescribeRelaySourceV3 [https://www.volcengine.com/docs/6469/1126874]接口获取待删除配置的 App 取值。
	App *string `json:"App,omitempty"`

	// 回源组名称。
	Group *string `json:"Group,omitempty"`
}

type DeleteRelaySourceV3Res struct {

	// REQUIRED
	ResponseMetadata DeleteRelaySourceV3ResResponseMetadata `json:"ResponseMetadata"`
}

type DeleteRelaySourceV3ResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                       `json:"Version"`
	Error   *DeleteRelaySourceV3ResResponseMetadataError `json:"Error,omitempty"`
}

type DeleteRelaySourceV3ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DeleteRelaySourceV4Body struct {

	// REQUIRED; 应用名称，您可以调用ListRelaySourceV4 [https://www.volcengine.com/docs/6469/1126878]接口，获取待删除固定回源配置的 App 取值。
	App string `json:"App"`

	// REQUIRED; 拉流域名，您可以调用ListRelaySourceV4 [https://www.volcengine.com/docs/6469/1126878]接口，获取待删除固定回源配置的 Domain 取值。
	Domain string `json:"Domain"`

	// REQUIRED; 流名称，您可以调用ListRelaySourceV4 [https://www.volcengine.com/docs/6469/1126878]接口，获取待删除固定回源配置的 Stream 取值。
	Stream string `json:"Stream"`
}

type DeleteRelaySourceV4Res struct {

	// REQUIRED
	ResponseMetadata DeleteRelaySourceV4ResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteRelaySourceV4ResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                       `json:"Version"`
	Error   *DeleteRelaySourceV4ResResponseMetadataError `json:"Error,omitempty"`
}

type DeleteRelaySourceV4ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DeleteSnapshotPresetBody struct {

	// REQUIRED; 截图配置的名称，您可以调用 ListVhostSnapshotPresetV2 [https://www.volcengine.com/docs/6469/1208858] 接口获取，取值与 Name 字段取值相同。
	Preset string `json:"Preset"`

	// 应用名称，您可以调用ListVhostSnapshotPresetV2 [https://www.volcengine.com/docs/6469/1208858]接口，获取待更新截图配置的 App 取值。
	App *string `json:"App,omitempty"`

	// 域名空间，您可以调用 ListVhostSnapshotPresetV2 [https://www.volcengine.com/docs/6469/1208858] 接口，获取待删除截图配置的 Vhost 取值。
	Vhost *string `json:"Vhost,omitempty"`
}

type DeleteSnapshotPresetRes struct {

	// REQUIRED
	ResponseMetadata DeleteSnapshotPresetResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type DeleteSnapshotPresetResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                        `json:"Version"`
	Error   *DeleteSnapshotPresetResResponseMetadataError `json:"Error,omitempty"`
}

type DeleteSnapshotPresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DeleteTimeShiftPresetV3Body struct {

	// REQUIRED; 应用名称，您可以调用ListTimeShiftPresetV2 [https://www.volcengine.com/docs/6469/1126883]接口，获取待删除时移配置的App取值。
	App string `json:"App"`

	// REQUIRED; 域名空间名称，您可以调用ListTimeShiftPresetV2 [https://www.volcengine.com/docs/6469/1126883]接口，获取待删除时移配置的Vhost取值。
	Vhost string `json:"Vhost"`

	// 流名
	Stream *string `json:"Stream,omitempty"`
}

type DeleteTimeShiftPresetV3Res struct {

	// REQUIRED
	ResponseMetadata DeleteTimeShiftPresetV3ResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteTimeShiftPresetV3ResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                           `json:"Version"`
	Error   *DeleteTimeShiftPresetV3ResResponseMetadataError `json:"Error,omitempty"`
}

type DeleteTimeShiftPresetV3ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DeleteTranscodePresetBody struct {

	// REQUIRED; 应用名称，您可以调用 ListVhostTransCodePreset [https://www.volcengine.com/docs/6469/1126853] 接口查看待删除转码配置的 App 取值。
	App string `json:"App"`

	// REQUIRED; 转码配置名称，您可以调用 ListVhostTransCodePreset [https://www.volcengine.com/docs/6469/1126853] 接口查看待删除转码配置的 Preset 取值。
	Preset string `json:"Preset"`

	// REQUIRED; 域名空间，您可以调用 ListVhostTransCodePreset [https://www.volcengine.com/docs/6469/1126853] 接口查看待删除转码配置的 Vhost 取值。
	Vhost string `json:"Vhost"`
}

type DeleteTranscodePresetRes struct {

	// REQUIRED
	ResponseMetadata DeleteTranscodePresetResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type DeleteTranscodePresetResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version   string                                         `json:"Version"`
	Error     *DeleteTranscodePresetResResponseMetadataError `json:"Error,omitempty"`
	RequestID *string                                        `json:"RequestID,omitempty"`
}

type DeleteTranscodePresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DeleteUserAgentAccessRuleBody struct {

	// REQUIRED
	Vhost  string  `json:"Vhost"`
	Domain *string `json:"Domain,omitempty"`
}

type DeleteUserAgentAccessRuleRes struct {

	// REQUIRED
	ResponseMetadata DeleteUserAgentAccessRuleResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteUserAgentAccessRuleResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestId为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type DeleteWatermarkPresetBody struct {

	// REQUIRED; 应用名称，您可以调用ListVhostWatermarkPreset [https://www.volcengine.com/docs/6469/1126889]接口，查看待删除水印配置的 App 取值。
	App string `json:"App"`

	// REQUIRED; 域名空间，您可以调用 ListVhostWatermarkPreset [https://www.volcengine.com/docs/6469/1126889] 接口，查看待删除水印配置的 Vhost 取值。
	Vhost string `json:"Vhost"`

	// 直播地址流名。可选。
	Stream *string `json:"Stream,omitempty"`
}

type DeleteWatermarkPresetRes struct {

	// REQUIRED
	ResponseMetadata DeleteWatermarkPresetResResponseMetadata `json:"ResponseMetadata"`
}

type DeleteWatermarkPresetResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                         `json:"Version"`
	Error   *DeleteWatermarkPresetResResponseMetadataError `json:"Error,omitempty"`
}

type DeleteWatermarkPresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeAuthBody struct {

	// REQUIRED; 鉴权场景类型，取值及含义如下所示。
	// * push：推流鉴权；
	// * pull：拉流鉴权。
	SceneType string `json:"SceneType"`

	// 应用名称，取值与直播流地址中 AppName 字段取值相同，默认为空，表示所有应用名称。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App *string `json:"App,omitempty"`

	// 直播流使用的域名。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看待配置鉴权的推拉流域名。
	// :::tip 参数 Domain 和 Vhost 传且仅传一个。
	// :::
	Domain *string `json:"Domain,omitempty"`

	// 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看需要查询的直播流使用的域名所属的域名空间。
	// :::tip 参数
	// Domain 和 Vhost 传且仅传一个。 :::
	Vhost *string `json:"Vhost,omitempty"`
}

type DescribeAuthRes struct {

	// REQUIRED
	ResponseMetadata DescribeAuthResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeAuthResResult          `json:"Result,omitempty"`
}

type DescribeAuthResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                `json:"Version"`
	Error   *DescribeAuthResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeAuthResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeAuthResResult struct {

	// 推/拉流鉴权列表。
	AuthList []*DescribeAuthResResultAuthListItem `json:"AuthList,omitempty"`
}

type DescribeAuthResResultAuthListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 是否开启 URL 地址鉴权，取值及含义如下所示。
	// * false：关闭；
	// * true：开启。
	AuthStatus bool `json:"AuthStatus"`

	// REQUIRED; 推/拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 鉴权场景类型，取值及含义如下所示。
	// * push：推流鉴权；
	// * pull：拉流鉴权。
	SceneType string `json:"SceneType"`

	// REQUIRED; 鉴权生效时长，单位为秒。
	ValidDuration int32 `json:"ValidDuration"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 鉴权详情。
	AuthDetailList []*DescribeAuthResResultAuthListPropertiesItemsItem `json:"AuthDetailList,omitempty"`
}

// DescribeAuthResResultAuthListPropertiesItemsItem - 鉴权详情。
type DescribeAuthResResultAuthListPropertiesItemsItem struct {

	// 自定义推拉流地址中，鉴权参数volcSecret和volcTime的名称。
	AuthField map[string]*string `json:"AuthField,omitempty"`

	// 旁路鉴权时，授权服务器的地址
	BypassAuthURL *string `json:"BypassAuthURL,omitempty"`

	// 接收旁路鉴权失败消息的回调地址
	BypassFailCallbackURL *string `json:"BypassFailCallbackURL,omitempty"`

	// 生成加密字符串使用的加密字段。
	EncryptField []*string `json:"EncryptField,omitempty"`

	// 对称加密算法。
	EncryptionAlgorithm *string `json:"EncryptionAlgorithm,omitempty"`

	// 旁路鉴权重试时长，单位为 s
	RetryInternalSecond *int32 `json:"RetryInternalSecond,omitempty"`

	// 旁路鉴权重试次数
	RetryTimes *int32 `json:"RetryTimes,omitempty"`

	// 自定义鉴权密钥。
	SecretKey *string `json:"SecretKey,omitempty"`

	// 旁路鉴权超时时长，单位为 s
	TimeoutSecond *int32 `json:"TimeoutSecond,omitempty"`
}

type DescribeCDNSnapshotHistoryBody struct {

	// REQUIRED; 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。 :::tip
	// * 当您查询指定截图任务详情时，DateFrom 应设置为推流开始时间之前的任意时间。
	// * 查询的最大时间跨度为 7 天。 :::
	DateFrom string `json:"DateFrom"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	DateTo string `json:"DateTo"`

	// REQUIRED; 流名称，取值与直播流地址中 StreamName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	Stream string `json:"Stream"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 查询数据的页码，默认为 1，表示查询第一页的数据。
	PageNum *int32 `json:"PageNum,omitempty"`

	// 每页显示的数据条数，默认为 10，最大值为 1000。
	PageSize *int32 `json:"PageSize,omitempty"`

	// 截图文件保存位置，取值及含义如下所示。
	// * tos：（默认值）TOS 对象存储服务；
	// * imageX：veImageX 图片服务。
	Type *string `json:"Type,omitempty"`
}

type DescribeCDNSnapshotHistoryRes struct {

	// REQUIRED
	ResponseMetadata DescribeCDNSnapshotHistoryResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeCDNSnapshotHistoryResResult          `json:"Result,omitempty"`
}

type DescribeCDNSnapshotHistoryResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                              `json:"Version"`
	Error   *DescribeCDNSnapshotHistoryResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeCDNSnapshotHistoryResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeCDNSnapshotHistoryResResult struct {

	// REQUIRED; 分页信息。
	Pagination DescribeCDNSnapshotHistoryResResultPagination `json:"Pagination"`

	// 截图文件信息。
	Data []*DescribeCDNSnapshotHistoryResResultDataItem `json:"Data,omitempty"`
}

type DescribeCDNSnapshotHistoryResResultDataItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 截图高度，单位为 px。
	Height int32 `json:"Height"`

	// REQUIRED
	ID int32 `json:"ID"`

	// REQUIRED; 截图文件保存的路径。
	Path string `json:"Path"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 截图时间戳，精度为毫秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 域名空间。
	Vhost string `json:"Vhost"`

	// REQUIRED; 截图宽度，单位为 px。
	Width int32 `json:"Width"`
}

// DescribeCDNSnapshotHistoryResResultPagination - 分页信息。
type DescribeCDNSnapshotHistoryResResultPagination struct {

	// REQUIRED; 查询数据的页码。
	PageCur int32 `json:"PageCur"`

	// REQUIRED; 每页显示的数据量条数。
	PageSize int32 `json:"PageSize"`

	// REQUIRED; 查询结果的数据总页数。
	PageTotal int32 `json:"PageTotal"`

	// REQUIRED; 查询结果的数据总条数。
	TotalCount int32 `json:"TotalCount"`
}

type DescribeCMAFConfigBody struct {

	// REQUIRED
	Vhost string  `json:"Vhost"`
	App   *string `json:"App,omitempty"`
}

type DescribeCMAFConfigRes struct {

	// REQUIRED
	ResponseMetadata DescribeCMAFConfigResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeCMAFConfigResResult `json:"Result"`
}

type DescribeCMAFConfigResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`
}

type DescribeCMAFConfigResResult struct {
	CMAFConfigList []*DescribeCMAFConfigResResultCMAFConfigListItem `json:"CMAFConfigList,omitempty"`
}

type DescribeCMAFConfigResResultCMAFConfigListItem struct {
	App               *string  `json:"App,omitempty"`
	DefaultLatency    *int32   `json:"DefaultLatency,omitempty"`
	DisableLowLatency *bool    `json:"DisableLowLatency,omitempty"`
	Interval          *float32 `json:"Interval,omitempty"`
	PlaylistLength    *int32   `json:"PlaylistLength,omitempty"`
	Vhost             *string  `json:"Vhost,omitempty"`
}

type DescribeCallbackBody struct {

	// domain, app二选一必传
	App *string `json:"App,omitempty"`

	// domain, app二选一必传
	Domain *string `json:"Domain,omitempty"`

	// 回调类型。默认为空，表示查询全部回调类型，取值及含义如下所示。
	// * push：推流开始回调；
	// * push_end：推流结束回调；
	// * snapshot：截图回调；
	// * record：录制回调；
	// * audit_snapshot：截图审核回调。
	MessageType *string `json:"MessageType,omitempty"`

	// 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看需要查询的直播流使用的域名所属的域名空间。
	// :::tipVhost和
	// Domain传且仅传一个。 :::
	Vhost *string `json:"Vhost,omitempty"`
}

type DescribeCallbackRes struct {

	// REQUIRED
	ResponseMetadata DescribeCallbackResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeCallbackResResult          `json:"Result,omitempty"`
}

type DescribeCallbackResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                    `json:"Version"`
	Error   *DescribeCallbackResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeCallbackResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeCallbackResResult struct {

	// 回调列表。
	CallbackList []*DescribeCallbackResResultCallbackListItem `json:"CallbackList,omitempty"`
}

type DescribeCallbackResResultCallbackListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 回调消息发送是否开启鉴权，默认为false，取值及含义如下所示。
	// * false：不开启；
	// * true：开启。
	AuthEnable bool `json:"AuthEnable"`

	// REQUIRED
	AuthField DescribeCallbackResResultCallbackListItemAuthField `json:"AuthField"`

	// REQUIRED; 回调消息发送鉴权密钥，开启回调消息鉴权时生效。
	AuthKeyPrimary string `json:"AuthKeyPrimary"`

	// REQUIRED; 回调创建时间
	CreateTime string `json:"CreateTime"`

	// REQUIRED; 格式为rfc3339，时区为utc的回调创建时间，
	CreateTimeUTC string `json:"CreateTimeUTC"`

	// REQUIRED; 回调的消息类型，取值及含义如下所示。
	// * push：推流开始回调；
	// * push_end：推流结束回调；
	// * snapshot：截图回调；
	// * record：录制回调；
	// * audit_snapshot：截图审核回调。
	MessageType string `json:"MessageType"`

	// REQUIRED; 是否开启转码流回调，默认为 0。取值及含义如下所示。
	// * 0：不开启；
	// * 1：开启。
	TranscodeCallback int32 `json:"TranscodeCallback"`

	// REQUIRED; 域名空间。
	Vhost         string  `json:"Vhost"`
	AuthKeySecond *string `json:"AuthKeySecond,omitempty"`

	// 回调数据列表。
	CallbackDetailList  []*DescribeCallbackResResultCallbackListPropertiesItemsItem `json:"CallbackDetailList,omitempty"`
	CallbackField       []*string                                                   `json:"CallbackField,omitempty"`
	Domain              *string                                                     `json:"Domain,omitempty"`
	EncryptField        []*string                                                   `json:"EncryptField,omitempty"`
	EncryptionAlgorithm *string                                                     `json:"EncryptionAlgorithm,omitempty"`
	HTTPMethod          *string                                                     `json:"HttpMethod,omitempty"`
	NotUseVhost         *bool                                                       `json:"NotUseVhost,omitempty"`
	RetryInternalSecond *int32                                                      `json:"RetryInternalSecond,omitempty"`
	RetryTimes          *int32                                                      `json:"RetryTimes,omitempty"`
	SecHandlerType      *string                                                     `json:"SecHandlerType,omitempty"`

	// 任务状态回调开关
	TaskStatusCallback *int32 `json:"TaskStatusCallback,omitempty"`
	TimeoutSecond      *int32 `json:"TimeoutSecond,omitempty"`
	ValidDuration      *int32 `json:"ValidDuration,omitempty"`
}

type DescribeCallbackResResultCallbackListItemAuthField struct {

	// REQUIRED
	AuthKeyPrimary string `json:"AuthKeyPrimary"`

	// REQUIRED
	AuthKeySecond string `json:"AuthKeySecond"`

	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]*string
}

type DescribeCallbackResResultCallbackListPropertiesItemsItem struct {

	// REQUIRED; 回调类型，返回 HTTP，表示可以使用 HTTP 和 HTTPS 地址接收回调消息。
	CallbackType string `json:"CallbackType"`

	// REQUIRED; 回调消息接收地址。
	URL string `json:"URL"`
}

type DescribeCertDRMQuery struct {

	// REQUIRED; 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App" query:"App"`

	// REQUIRED; 域名空间，即直播流地址的域名（Domain）所属的域名空间（Vhost）。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost" query:"Vhost"`
}

type DescribeCertDRMRes struct {

	// REQUIRED
	ResponseMetadata DescribeCertDRMResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DescribeCertDRMResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type DescribeCertDetailSecretV2Body struct {

	// 账号ID
	AccountID *string `json:"AccountID,omitempty"`

	// 证书实例 ID，您可以通过ListCertV2 [https://www.volcengine.com/docs/6469/81242]接口获取证书示例 ID。 :::tip 参数ChainID与CertID传且仅传一个。 :::
	CertID *string `json:"CertID,omitempty"`

	// 证书链 ID，您可以通过ListcCertV2 [https://www.volcengine.com/docs/6469/81242]接口获取 证书链 ID。 :::tip 参数ChainID与CertID传且仅传一个。 :::
	ChainID *string `json:"ChainID,omitempty"`
}

type DescribeCertDetailSecretV2Res struct {

	// REQUIRED
	ResponseMetadata DescribeCertDetailSecretV2ResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeCertDetailSecretV2ResResult `json:"Result,omitempty"`
}

type DescribeCertDetailSecretV2ResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                              `json:"Version"`
	Error   *DescribeCertDetailSecretV2ResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeCertDetailSecretV2ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

// DescribeCertDetailSecretV2ResResult - 视请求的接口而定
type DescribeCertDetailSecretV2ResResult struct {

	// REQUIRED; 与证书绑定的域名列表。
	CertDomainList []string `json:"CertDomainList"`

	// REQUIRED; 证书名称。
	CertName string `json:"CertName"`

	// REQUIRED; 证书链 ID。
	ChainID string `json:"ChainID"`

	// REQUIRED; 证书的过期时间，RFC3339 格式的 UTC 时间，精度为秒。
	NotAfter string `json:"NotAfter"`

	// REQUIRED; 证书的生效日期，RFC3339 格式的 UTC 时间，精度为秒。
	NotBefore string `json:"NotBefore"`

	// REQUIRED; 证书状态，取值及含义如下所示。
	// * OK：正常；
	// * Expire：过期；
	// * 30days：有效期剩余 30 天；
	// * 15days：有效期剩余 15 天；
	// * 7days：有效期剩余 7 天；
	// * 1days：有效期剩余 1 天。
	Status string `json:"Status"`

	// 加密算法
	EncryptionAlgorithm *string `json:"EncryptionAlgorithm,omitempty"`

	// 证书指纹（SHA1）
	FingerprintSHA1 *string `json:"FingerprintSHA1,omitempty"`

	// 证书指纹（SHA256）
	FingerprintSHA256 *string `json:"FingerprintSHA256,omitempty"`

	// 签发者信息
	Issuer *string `json:"Issuer,omitempty"`

	// openssl解析结果
	OpenSSLFormat *string `json:"OpenSSLFormat,omitempty"`

	// 证书详细信息。
	SSL *DescribeCertDetailSecretV2ResResultSSL `json:"SSL,omitempty"`

	// 签名算法
	SignatureAlgorithm *string `json:"SignatureAlgorithm,omitempty"`
}

// DescribeCertDetailSecretV2ResResultSSL - 证书详细信息。
type DescribeCertDetailSecretV2ResResultSSL struct {

	// REQUIRED; 证书链，包括叶子证书（服务器证书）、中间证书（中间 CA 证书）以及根证书（根 CA 证书）。证书链中的证书使用 PEM 编码格式。
	Chain []string `json:"Chain"`

	// REQUIRED; 密钥类型，默认为rsa。
	KeyType string `json:"KeyType"`

	// REQUIRED; 证书私钥
	PrivateKey string `json:"PrivateKey"`

	// 证书链解析后的证书链简短信息。
	ChainBriefInfo []*DescribeCertDetailSecretV2ResResultSSLChainBriefInfoItem `json:"ChainBriefInfo,omitempty"`
}

type DescribeCertDetailSecretV2ResResultSSLChainBriefInfoItem struct {

	// 加密算法
	EncryptionAlgorithm *string `json:"EncryptionAlgorithm,omitempty"`

	// 证书指纹（SHA1）
	FingerprintSHA1 *string `json:"FingerprintSHA1,omitempty"`

	// 证书指纹（SHA256）
	FingerprintSHA256 *string `json:"FingerprintSHA256,omitempty"`

	// 签发者信息
	Issuer *string `json:"Issuer,omitempty"`

	// 签名算法
	SignatureAlgorithm *string `json:"SignatureAlgorithm,omitempty"`

	// 主题信息
	Subject *string `json:"Subject,omitempty"`
}

type DescribeClosedStreamInfoByPageQuery struct {

	// REQUIRED; 查询的起始时间，RFC3339 格式的时间戳，精度为秒。筛选直播流结束时间符合查询条件的历史流。
	EndTimeFrom string `json:"EndTimeFrom" query:"EndTimeFrom"`

	// REQUIRED; 查询的结束时间，RFC3339 格式表示的时间戳，精度为秒。筛选直播流结束时间符合查询条件的历史流。
	EndTimeTo string `json:"EndTimeTo" query:"EndTimeTo"`

	// REQUIRED; 查询数据的页码，取值范围为正整数。
	PageNum int32 `json:"PageNum" query:"PageNum"`

	// REQUIRED; 每页显示的数据条数，取值范围为 [1,1000]。
	PageSize int32 `json:"PageSize" query:"PageSize"`

	// 应用名称，取值与直播流地址中 AppName 字段取值相同，默认为空，表示查询所有应用名称。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App *string `json:"App,omitempty" query:"App"`

	// 直播流使用的域名，默认为空，表示查询所有当前域名空间下的历史直播流。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，查看需要查询的历史直播流使用的域名。
	Domain *string `json:"Domain,omitempty" query:"Domain"`

	// 想要查询的目标信息，使用英文逗号作为分隔符“,”，例如，bitrate,framerate。缺省情况下表示 bitrate,framerate。支持如下取值。 all：所有信息；onlineuser：在线人数；bandwidth：带宽信息;bitrate：码率信息；framerate：帧率信息；.
	InfoType *string `json:"InfoType,omitempty" query:"InfoType"`

	// 使用流名称进行查询的方式，默认值为 strict，支持的取值即含义如下所示。
	// * fuzzy：模糊匹配；
	// * strict：精准匹配。
	QueryType *string `json:"QueryType,omitempty" query:"QueryType"`

	// 排列方式，根据直播流结束时间排序，默认值为 desc，支持的取值及含义如下所示。
	// * asc：从时间最远到最近排序；
	// * desc：从时间最近到最远排序。
	Sort *string `json:"Sort,omitempty" query:"Sort"`

	// 历史直播流的来源类型，默认为空，表示查询所有来源类型，支持的取值及含义如下所示。
	// * push：直推流；
	// * relay：回源流。
	SourceType *string `json:"SourceType,omitempty" query:"SourceType"`

	// 流名称，取值与直播流地址中 StreamName 字段取值相同，默认为空表示查询所有流名称。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	Stream *string `json:"Stream,omitempty" query:"Stream"`

	// 流类型，缺省情况下表示全选。支持如下取值。Origin：原始流；trans：转码流。.
	StreamType *string `json:"StreamType,omitempty" query:"StreamType"`

	// 域名空间，即直播流地址的域名所属的域名空间，默认为空，表示查询所有域名空间下的历史直播流。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]
	// 页面，查看需要查询的历史直播流使用的域名所属的域名空间。
	Vhost *string `json:"Vhost,omitempty" query:"Vhost"`
}

type DescribeClosedStreamInfoByPageRes struct {

	// REQUIRED
	ResponseMetadata DescribeClosedStreamInfoByPageResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeClosedStreamInfoByPageResResult          `json:"Result,omitempty"`
}

type DescribeClosedStreamInfoByPageResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                                  `json:"Version"`
	Error   *DescribeClosedStreamInfoByPageResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeClosedStreamInfoByPageResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeClosedStreamInfoByPageResResult struct {

	// REQUIRED; 查询结果中历史流的数量。
	RoughCount int32 `json:"RoughCount"`

	// 历史直播流信息列表。
	StreamInfoList []*DescribeClosedStreamInfoByPageResResultStreamInfoListItem `json:"StreamInfoList,omitempty"`
}

type DescribeClosedStreamInfoByPageResResultStreamInfoListItem struct {

	// REQUIRED; 历史直播流使用的应用名称。
	App string `json:"App"`

	// REQUIRED; 历史直播流使用的域名。
	Domain string `json:"Domain"`

	// REQUIRED; 直播流的结束时间，RFC3339 格式的 UTC 时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 历史直播流的来源类型，取值及含义如下所示。
	// * push：直推流；
	// * relay：回源流。
	SourceType string `json:"SourceType"`

	// REQUIRED; 直播流的开始时间，RFC3339 格式的 UTC 时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 历史直播流使用的流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 历史直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`
}

type DescribeClusterRateLimitBody struct {

	// REQUIRED
	Vhost  string  `json:"Vhost"`
	App    *string `json:"App,omitempty"`
	Domain *string `json:"Domain,omitempty"`
}

type DescribeClusterRateLimitRes struct {

	// REQUIRED
	ResponseMetadata DescribeClusterRateLimitResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeClusterRateLimitResResult `json:"Result"`
}

type DescribeClusterRateLimitResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestID"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`
}

type DescribeClusterRateLimitResResult struct {

	// REQUIRED
	ClusterRateLimitItemList []DescribeClusterRateLimitResResultClusterRateLimitItemListItem `json:"ClusterRateLimitItemList"`
}

type DescribeClusterRateLimitResResultClusterRateLimitItemListItem struct {
	AggregationPeriod *int32  `json:"AggregationPeriod,omitempty"`
	App               *string `json:"App,omitempty"`
	Domain            *string `json:"Domain,omitempty"`
	Limit             *int32  `json:"Limit,omitempty"`
	Param             *string `json:"Param,omitempty"`
	RejectCode        *int32  `json:"RejectCode,omitempty"`
	RejectDuration    *int32  `json:"RejectDuration,omitempty"`
	Status            *int32  `json:"Status,omitempty"`
	Type              *string `json:"Type,omitempty"`
	Vhost             *string `json:"Vhost,omitempty"`
}

type DescribeDomainBody struct {

	// REQUIRED; 待查询域名信息的域名列表。
	DomainList []string `json:"DomainList"`
}

type DescribeDomainRes struct {

	// REQUIRED
	ResponseMetadata DescribeDomainResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeDomainResResult          `json:"Result,omitempty"`
}

type DescribeDomainResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                  `json:"Version"`
	Error   *DescribeDomainResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeDomainResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeDomainResResult struct {

	// 域名详细信息列表。
	DomainList []*DescribeDomainResResultDomainListItem `json:"DomainList,omitempty"`
}

type DescribeDomainResResultDomainListItem struct {

	// REQUIRED; CNAME 信息。
	CNAME string `json:"CNAME"`

	// REQUIRED; 绑定的 HTTPS 证书支持的泛域名。
	CertDomain string `json:"CertDomain"`

	// REQUIRED; 绑定的证书名称。
	CertName string `json:"CertName"`

	// REQUIRED; 绑定的 HTTPS 证书的证书链 ID 信息。
	ChainID string `json:"ChainID"`

	// REQUIRED; CNAME 状态。
	// * 0：未配置 CNAME；
	// * 1：已配置 CNAME。
	CnameCheck int32 `json:"CnameCheck"`

	// REQUIRED; 域名添加时间，RFC3339 格式的 UTC 时间戳，精度为秒。
	CreateTime string `json:"CreateTime"`

	// REQUIRED; 推/拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 域名是否可用的状态。
	// * 0：正常，域名为可用状态；
	// * 1：配置中，域名为可用状态；
	// * 2：不可用，域名为其他的不可用状态。
	DomainCheck int32 `json:"DomainCheck"`

	// REQUIRED; ICP 备案校验是否通过，是否过期信息。
	// * 1：备案正常，未过期；
	// * 2：查存不到备案信息。
	ICPCheck int32 `json:"ICPCheck"`

	// REQUIRED; 当前域名所属的域名空间下的推流域名。
	PushDomain string `json:"PushDomain"`

	// REQUIRED; 域名加速区域，包含以下类型。
	// * cn：中国内地；
	// * cn-global：全球加速；
	// * cn-oversea：海外及港澳台。
	Region string `json:"Region"`

	// REQUIRED; 域名状态。状态说明如下所示。
	// * 0：正常；
	// * 1：审核中；
	// * 2：禁用，禁止使用，此时域名加速不生效；
	// * 3：删除；
	// * 4：审核被驳回，审核不通过，需要重新创建并审核；
	// * 5：欠费关停。
	Status int32 `json:"Status"`

	// REQUIRED; 域名类型，包含两种类型。
	// * push：推流域名；
	// * pull-flv：拉流域名，包含 RTMP、FLV、HLS 格式。
	Type string `json:"Type"`

	// REQUIRED; 域名空间。
	Vhost string `json:"Vhost"`
}

type DescribeEncryptDRMRes struct {

	// REQUIRED
	ResponseMetadata DescribeEncryptDRMResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeEncryptDRMResResult `json:"Result"`
}

type DescribeEncryptDRMResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`
}

type DescribeEncryptDRMResResult struct {

	// REQUIRED; DRM 加密配置列表。
	DRMItem DescribeEncryptDRMResResultDRMItem `json:"DRMItem"`
}

// DescribeEncryptDRMResResultDRMItem - DRM 加密配置列表。
type DescribeEncryptDRMResResultDRMItem struct {

	// REQUIRED; DRM 证书管理平台 API 访问密钥。
	APIKey string `json:"APIKey"`

	// REQUIRED; 申请 FairPlay 证书过程中 Apple 返回的 ASk（Application Secret Key）字符串。
	ApplicationSecretKey string `json:"ApplicationSecretKey"`

	// REQUIRED; FairPlay 证书文件的名称。
	CertificateFileName string `json:"CertificateFileName"`

	// REQUIRED; 自定义的 FairPlay 证书名称。
	CertificateName string `json:"CertificateName"`

	// REQUIRED; 申请 FairPlay 证书时创建的私钥文件密钥。
	PrivateKey string `json:"PrivateKey"`

	// REQUIRED; 申请 FairPlay 证书时创建的私钥文件名称。
	PrivateKeyFileName string `json:"PrivateKeyFileName"`
}

type DescribeForbiddenStreamInfoByPageQuery struct {

	// REQUIRED; 查询数据的页码，取值范围为正整数。
	PageNum int32 `json:"PageNum" query:"PageNum"`

	// REQUIRED; 每页显示的数据条数，取值范围为 [1,1000]。
	PageSize int32 `json:"PageSize" query:"PageSize"`

	// 应用名称，取值与禁推直播流时设置的应用名称相同，默认为空，表示查询当前域名空间下所有的禁推流。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App *string `json:"App,omitempty" query:"App"`

	// 直播流使用的域名，取值与禁推直播流时设置的应用名称相同，默认为空，表示查询所有当前域名空间下的禁推直播流。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]
	// 接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]
	// 页面，查看需要查询的禁推直播流使用的域名。
	Domain *string `json:"Domain,omitempty" query:"Domain"`

	// 指定是否模糊匹配流名称。缺省情况为精准匹配，支持的取值及含义如下所示。
	// * fuzzy：模糊匹配；
	// * strict：精准匹配。
	QueryType *string `json:"QueryType,omitempty" query:"QueryType"`

	// 排列方式，根据推流结束时间排序，默认值为 desc，支持的取值及含义如下所示。
	// * asc：从时间最远到最近排序；
	// * desc：从时间最近到最远排序。
	Sort *string `json:"Sort,omitempty" query:"Sort"`

	// 流名称，取值与禁推直播流时设置的流名称相同，默认为空，表示查询所有流名称。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	Stream *string `json:"Stream,omitempty" query:"Stream"`

	// 域名空间，取值与禁推直播流时设置的域名空间相同，默认为空，表示查询所有域名空间下的禁推流。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]
	// 页面，查看需要查询的禁推流使用的域名所属的域名空间。
	Vhost *string `json:"Vhost,omitempty" query:"Vhost"`
}

type DescribeForbiddenStreamInfoByPageRes struct {

	// REQUIRED
	ResponseMetadata DescribeForbiddenStreamInfoByPageResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeForbiddenStreamInfoByPageResResult          `json:"Result,omitempty"`
}

type DescribeForbiddenStreamInfoByPageResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                                     `json:"Version"`
	Error   *DescribeForbiddenStreamInfoByPageResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeForbiddenStreamInfoByPageResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeForbiddenStreamInfoByPageResResult struct {

	// REQUIRED; 查询结果中禁推流数量。
	RoughCount int32 `json:"RoughCount"`

	// 禁推流的信息列表。
	StreamInfoList []*DescribeForbiddenStreamInfoByPageResResultStreamInfoListItem `json:"StreamInfoList,omitempty"`
}

type DescribeForbiddenStreamInfoByPageResResultStreamInfoListItem struct {

	// REQUIRED; 禁推流的应用名称。
	App string `json:"App"`

	// REQUIRED; 禁推流被禁推的开始时间，RFC3339 格式的 UTC 时间戳，精度为秒。
	CreateTime string `json:"CreateTime"`

	// REQUIRED; 禁推流的域名。
	Domain string `json:"Domain"`

	// REQUIRED; 禁推流结束禁推的时间，RFC3339 格式的 UTC 时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 当前流的禁推配置是否启用。
	// * true：启用；
	// * false：禁用。
	Status bool `json:"Status"`

	// REQUIRED; 禁推流的流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 禁推流的域名空间。
	Vhost string `json:"Vhost"`
}

type DescribeFormatAccessRuleBody struct {

	// REQUIRED
	Domain string `json:"Domain"`

	// REQUIRED
	Vhost string `json:"Vhost"`
}

type DescribeFormatAccessRuleRes struct {

	// REQUIRED
	ResponseMetadata DescribeFormatAccessRuleResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeFormatAccessRuleResResult `json:"Result"`
}

type DescribeFormatAccessRuleResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`
}

type DescribeFormatAccessRuleResResult struct {

	// REQUIRED
	AccessRuleLists []DescribeFormatAccessRuleResResultAccessRuleListsItem `json:"AccessRuleLists"`
}

type DescribeFormatAccessRuleResResultAccessRuleListsItem struct {

	// REQUIRED
	Domain string `json:"Domain"`

	// REQUIRED
	FormatAccessRule DescribeFormatAccessRuleResResultAccessRuleListsItemFormatAccessRule `json:"FormatAccessRule"`

	// REQUIRED
	Vhost string `json:"Vhost"`
}

type DescribeFormatAccessRuleResResultAccessRuleListsItemFormatAccessRule struct {

	// REQUIRED
	Enable string `json:"Enable"`

	// REQUIRED
	FormatList []string `json:"FormatList"`

	// REQUIRED
	Type string `json:"Type"`
}

type DescribeHLSConfigBody struct {

	// REQUIRED
	Vhost string `json:"Vhost"`
}

type DescribeHLSConfigRes struct {

	// REQUIRED
	ResponseMetadata DescribeHLSConfigResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeHLSConfigResResult `json:"Result,omitempty"`
}

type DescribeHLSConfigResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

// DescribeHLSConfigResResult - 视请求的接口而定
type DescribeHLSConfigResResult struct {

	// REQUIRED
	HLSConfigList []DescribeHLSConfigResResultHLSConfigListItem `json:"HLSConfigList"`
}

type DescribeHLSConfigResResultHLSConfigListItem struct {

	// REQUIRED
	CreateTime string `json:"CreateTime"`

	// REQUIRED
	FirstPlaylistLength int32 `json:"FirstPlaylistLength"`

	// REQUIRED
	Interval float32 `json:"Interval"`

	// REQUIRED
	PartTargetDuration float32 `json:"PartTargetDuration"`

	// REQUIRED
	PlaylistLength int32 `json:"PlaylistLength"`

	// REQUIRED
	UpdateTime string `json:"UpdateTime"`

	// REQUIRED
	Vhost string `json:"Vhost"`
}

type DescribeHTTPHeaderConfigBody struct {

	// REQUIRED; HTTP Header 的类型，支持的取值及含义如下所示。
	// * 0：请求响应头；
	// * 1：回源请求头。
	Phase int32 `json:"Phase"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看需要域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 拉流域名。默认为空，表示查询 Vhost 下的全部拉流域名的 HTTP Header 配置。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，查看待查询的拉流域名。
	Domain *string `json:"Domain,omitempty"`
}

type DescribeHTTPHeaderConfigRes struct {

	// REQUIRED
	ResponseMetadata DescribeHTTPHeaderConfigResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 视请求的接口而定
	Result DescribeHTTPHeaderConfigResResult `json:"Result"`
}

type DescribeHTTPHeaderConfigResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

// DescribeHTTPHeaderConfigResResult - 视请求的接口而定
type DescribeHTTPHeaderConfigResResult struct {

	// REQUIRED; HTTP Header 配置信息。
	HeaderConfigList []DescribeHTTPHeaderConfigResResultHeaderConfigListItem `json:"HeaderConfigList"`
}

type DescribeHTTPHeaderConfigResResultHeaderConfigListItem struct {

	// REQUIRED; 是否保留原 Header 配置，取值及含义如下所示。
	// * 0：保留；
	// * 1：不保留。
	BlockOriginal int32 `json:"BlockOriginal"`

	// REQUIRED; 拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 配置是否启用，取值及含义如下所示。
	// * true：启用；
	// * false：禁用。
	Enable bool `json:"Enable"`

	// REQUIRED; 域名的 HTTP Header 具体字段配置。
	HeaderDetailList []DescribeHTTPHeaderConfigResResultHeaderConfigListPropertiesItemsItem `json:"HeaderDetailList"`

	// REQUIRED; 域名空间。
	Vhost string `json:"Vhost"`
}

type DescribeHTTPHeaderConfigResResultHeaderConfigListPropertiesItemsItem struct {

	// REQUIRED; Header 配置中字段 Value 值的类型，取值及含义如下所示。
	// * 0：常量；
	// * 1：变量。
	HeaderFieldType int32 `json:"HeaderFieldType"`

	// REQUIRED; Header 配置中字段的 Key 值。
	HeaderKey string `json:"HeaderKey"`

	// REQUIRED; Header 配置中字段的 Value 值。
	HeaderValue string `json:"HeaderValue"`
}

type DescribeIPAccessRuleBody struct {

	// REQUIRED; 推流域名或拉流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取需要查询
	// IP 访问限制的域名。
	Domain string `json:"Domain"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取需要配置
	// IP 访问限制的域名所属的域名空间。
	Vhost string `json:"Vhost"`
}

type DescribeIPAccessRuleRes struct {

	// REQUIRED
	ResponseMetadata DescribeIPAccessRuleResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeIPAccessRuleResResult `json:"Result"`
}

type DescribeIPAccessRuleResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Error string `json:"Error"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`
}

type DescribeIPAccessRuleResResult struct {

	// REQUIRED; IP 访问限制规则列表。
	AccessRuleLists []DescribeIPAccessRuleResResultAccessRuleListsItem `json:"AccessRuleLists"`
}

type DescribeIPAccessRuleResResultAccessRuleListsItem struct {

	// 推/拉流域名。
	Domain *string `json:"Domain,omitempty"`

	// IP 访问限制规则。
	IPAccessRule *DescribeIPAccessRuleResResultAccessRuleListsItemIPAccessRule `json:"IPAccessRule,omitempty"`

	// 域名空间名称。
	Vhost *string `json:"Vhost,omitempty"`
}

// DescribeIPAccessRuleResResultAccessRuleListsItemIPAccessRule - IP 访问限制规则。
type DescribeIPAccessRuleResResultAccessRuleListsItemIPAccessRule struct {

	// REQUIRED; 是否开启当前限制，取值及含义如下所示。
	// * true: 开启；
	// * false: 关闭。
	Enable bool `json:"Enable"`

	// REQUIRED; 名单中的 IP 信息。
	IPList []string `json:"IPList"`

	// REQUIRED; IP 访问限制的类型，取值及含义如下所示。
	// * allow: 白名单；
	// * deny: 黑名单。
	Type string `json:"Type"`
}

type DescribeIPInfoBody struct {

	// REQUIRED; 待查询的 IP 地址列表。支持 IPv4 和 IPv6 地址，一次最多查询 50 个 IP 地址。
	IPs []string `json:"Ips"`
}

type DescribeIPInfoRes struct {

	// REQUIRED
	ResponseMetadata DescribeIPInfoResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeIPInfoResResult          `json:"Result,omitempty"`
}

type DescribeIPInfoResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                  `json:"Version"`
	Error   *DescribeIPInfoResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeIPInfoResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeIPInfoResResult struct {

	// REQUIRED; IP 详情列表。
	List []DescribeIPInfoResResultListItem `json:"List"`
}

type DescribeIPInfoResResultListItem struct {

	// REQUIRED; CDN 节点所属城市信息。非归属火山引擎的 CDN 节点 IP 时，返回“-”。
	City string `json:"City"`

	// REQUIRED; IP 地址。
	IP string `json:"Ip"`

	// REQUIRED; CDN 节点所属网络运营商。非归属火山引擎 CDN 节点的 IP 时，返回”-“。您可以通过DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974]接口查看运营商标识符对应的运营商名称。
	Isp string `json:"Isp"`

	// REQUIRED; 当前 IP 地址是否是归属于火山引擎的 CDN 节点 IP。
	// * true：属于；
	// * false：不属于。
	LiveCdnIP bool `json:"LiveCdnIp"`

	// REQUIRED; CDN 节点所属国家或地区信息。非归属火山引擎的 CDN 节点 IP 时，返回“-”。
	Location string `json:"Location"`

	// REQUIRED; CDN 节点所属省份信息。非归属火山引擎 CDN 节点的 IP 时，返回“-”。
	Province string `json:"Province"`
}

type DescribeLatencyConfigBody struct {
	Domain *string `json:"Domain,omitempty"`
	Vhost  *string `json:"Vhost,omitempty"`
}

type DescribeLatencyConfigRes struct {

	// REQUIRED
	ResponseMetadata DescribeLatencyConfigResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeLatencyConfigResResult `json:"Result,omitempty"`
}

type DescribeLatencyConfigResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

// DescribeLatencyConfigResResult - 视请求的接口而定
type DescribeLatencyConfigResResult struct {

	// REQUIRED
	LatencyConfigList []DescribeLatencyConfigResResultLatencyConfigListItem `json:"LatencyConfigList"`
}

type DescribeLatencyConfigResResultLatencyConfigListItem struct {

	// REQUIRED
	Domain string `json:"Domain"`

	// REQUIRED; 单位ms
	GopCacheSize string `json:"GopCacheSize"`

	// REQUIRED
	Vhost string `json:"Vhost"`
}

type DescribeLicenseDRMQuery struct {

	// REQUIRED; 应用名称，取值与直播流地址的 AppName 字段相同，由大写小字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App" query:"App"`

	// REQUIRED; DRM 加密的类型，取值及含义如下所示。
	// * fp：FairPlay 加密；
	// * wv：Widevine 加密。
	DRMType string `json:"DRMType" query:"DRMType"`

	// REQUIRED; 拉流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的拉流域名。
	Domain string `json:"Domain" query:"Domain"`

	// REQUIRED; 流名称，取值与直播流地址中 StreamName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	StreamName string `json:"StreamName" query:"StreamName"`

	// REQUIRED; 拉取加密流时使用的拉流域名所在的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的拉流域名所属的域名空间。
	Vhost string `json:"Vhost" query:"Vhost"`
}

type DescribeLicenseDRMRes struct {

	// REQUIRED
	ResponseMetadata DescribeLicenseDRMResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DescribeLicenseDRMResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type DescribeLiveASRDurationDataBody struct {

	// REQUIRED; The end time of your query's range (UTC time in RFC 3339 format with second precision).
	EndTime string `json:"EndTime"`

	// REQUIRED; The start time of your query's range (UTC time in RFC 3339 format with second precision).
	StartTime string `json:"StartTime"`

	// AI字幕原始语言，枚举值：ZH(中文)，EN(英文), JA(日文), KO(韩文)
	ASRSourceType *string `json:"ASRSourceType,omitempty"`

	// AI字幕转换的语言类型，ASRTargetTypeList和ASRSourceType必须同时指定， 枚举值：ZH(中文)，EN(英文), JA(日文), KO(韩文)
	ASRTargetTypeList []*string `json:"ASRTargetTypeList,omitempty"`

	// The granularity of data aggregation, measured in seconds, with the following supported options:
	// * 300 (default): 5 minutes. When aggregated in 5-minute intervals, the maximum time span for a single query is 31 days,
	// and for historical queries, the maximum time range is 366 days.
	// * 3600: 1 hour. When aggregated in 1-hour intervals, the maximum time span for a single query is 93 days, and for historical
	// queries, the maximum time range is 366 days.
	// * 86400: 1 day. When aggregated in 1-day intervals, the maximum time span for a single query is 93 days, and for historical
	// queries, the maximum time range is 366 days.
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用列表，缺省表示该用户所有的app
	AppList []*string `json:"AppList,omitempty"`

	// 域名列表，缺省情况表示该用户的所有域名
	DomainList []*string `json:"DomainList,omitempty"`
}

type DescribeLiveASRDurationDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveASRDurationDataResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeLiveASRDurationDataResResult `json:"Result,omitempty"`
}

type DescribeLiveASRDurationDataResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

// DescribeLiveASRDurationDataResResult - 视请求的接口而定
type DescribeLiveASRDurationDataResResult struct {

	// REQUIRED; 每个时间点对应的转换时长
	ASRDurationData []DescribeLiveASRDurationDataResResultASRDurationDataItem `json:"ASRDurationData"`

	// REQUIRED; 每个字幕转换类型在每个时间点的时长
	ASRDurationDetailData []DescribeLiveASRDurationDataResResultASRDurationDetailDataItem `json:"ASRDurationDetailData"`

	// REQUIRED; Data granularity, measured in seconds.
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; AI字幕总时长总量，unit is minute
	Duration float32 `json:"Duration"`

	// REQUIRED; The end time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	EndTime string `json:"EndTime"`

	// REQUIRED; The start time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	StartTime string `json:"StartTime"`

	// AI字幕原始语言
	ASRSourceType *string `json:"ASRSourceType,omitempty"`

	// AI字幕转换的语言类型
	ASRTargetTypeList []*string `json:"ASRTargetTypeList,omitempty"`

	// 应用列表，缺省表示该用户所有的app
	AppList []*string `json:"AppList,omitempty"`

	// 域名列表，缺省情况表示该用户的所有域名
	DomainList []*string `json:"DomainList,omitempty"`
}

type DescribeLiveASRDurationDataResResultASRDurationDataItem struct {

	// REQUIRED; AI字幕时长，单位为分钟
	Duration float32 `json:"Duration"`

	// REQUIRED; The start time of each time granularity when data is aggregated by time granularity, formatted in RFC3339 UTC
	// time with second precision.
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLiveASRDurationDataResResultASRDurationDetailDataItem struct {

	// REQUIRED; 每个时间点对应的时长
	ASRDurationData []Components17Ohct5SchemasDescribeliveasrdurationdataresPropertiesResultPropertiesAsrdurationdetaildataItemsPropertiesAsrdurationdataItems `json:"ASRDurationData"`

	// REQUIRED; AI字幕原始语言
	ASRSourceType string `json:"ASRSourceType"`

	// REQUIRED; AI字幕转换的语言类型
	ASRTargetTypeList []string `json:"ASRTargetTypeList"`
}

type DescribeLiveBandwidthDataBody struct {

	// REQUIRED; The end time of your query's range (UTC time in RFC 3339 format with second precision).
	EndTime string `json:"EndTime"`

	// REQUIRED; The start time of your query's range (UTC time in RFC 3339 format with second precision).
	StartTime string `json:"StartTime"`

	// The granularity of data aggregation, measured in seconds, with the following supported options:
	// * 300 (default): 5 minutes. When aggregated in 5-minute intervals, the maximum time span for a single query is 31 days,
	// and for historical queries, the maximum time range is 366 days.
	// * 3600: 1 hour. When aggregated in 1-hour intervals, the maximum time span for a single query is 93 days, and for historical
	// queries, the maximum time range is 366 days.
	// * 86400: 1 day. When aggregated in 1-day intervals, the maximum time span for a single query is 93 days, and for historical
	// queries, the maximum time range is 366 days.
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 数据拆分的维度，默认为空表示不按维度进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * ISP：运营商；
	// * Protocol：推拉流协议。 :::tip 配置数据拆分的维度时，对应的维度参数传入多个值时才会返回按此维度拆分的数据。例如，配置按 Domain 进行数据拆分时， DomainList 传入多个 Domain 值时，才会返回按 Domain
	// 拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，默认为空，表示查询您视频直播产品下所有域名的带宽用量数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// Identifiers of operators providing network access services. By default, all operators are indicated. Supported operators
	// are as follows.
	// * unicom: China Unicom;
	// * railcom: China Railway Telecom;
	// * telecom: China Telecom;
	// * mobile: China Mobile;
	// * cernet: China Education and Research Network (CERNET);
	// * tianwei: China Tianwei;
	// * alibaba: Alibaba Group;
	// * tencent: Tencent Holdings;
	// * drpeng: Dr. Peng Telecom & Media Group;
	// * btvn: China Broadcasting Network;
	// * huashu: Huashu Media;
	// * other: Denotes other/unspecified options.
	// If you need to obtain the identifiers of various operators, you can call the DescribeLiveISPData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveispdata].
	ISPList []*string `json:"ISPList,omitempty"`

	// byteplus比火山多了CMAF协议
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// List of regions to which CDN node IPs belong, by default indicating all regions. :::tipRegionList and UserRegionList cannot
	// be used together in the same request. :::
	RegionList []*DescribeLiveBandwidthDataBodyRegionListItem `json:"RegionList,omitempty"`

	// List of regions to which client IPs belong, by default indicating all regions.
	// :::tipRegionList and UserRegionList cannot be used together in a single request. :::
	UserRegionList []*DescribeLiveBandwidthDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveBandwidthDataBodyRegionListItem struct {

	// The identifier for the major region in the regional information can be obtained by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata].
	Area *string `json:"Area,omitempty"`

	// The country identifier in the regional information can be obtained by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata].
	// When filtering by
	// country, both 'Area' and 'Country' need to be passed in simultaneously.
	Country *string `json:"Country,omitempty"`

	// The province identifier in the regional information is currently not supported for countries or regions outside mainland
	// China, Hong Kong, Macao, and Taiwan. You can obtain the identifier information
	// by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata]. When
	// filtering by province, you need to simultaneously pass in Area, Country, and
	// Province
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveBandwidthDataBodyUserRegionListItem struct {

	// The identifier for the major region in the regional information can be obtained by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata].
	Area *string `json:"Area,omitempty"`

	// The country identifier in the regional information can be obtained by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata].
	// When filtering by
	// country, both 'Area' and 'Country' need to be passed in simultaneously.
	Country *string `json:"Country,omitempty"`

	// The province identifier in the regional information is currently not supported for countries or regions outside mainland
	// China, Hong Kong, Macao, and Taiwan. You can obtain the identifier information
	// by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata]. When
	// filtering by province, you need to simultaneously pass in Area, Country, and
	// Province
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveBandwidthDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveBandwidthDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveBandwidthDataResResult `json:"Result"`
}

type DescribeLiveBandwidthDataResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`
}

type DescribeLiveBandwidthDataResResult struct {

	// REQUIRED; Data granularity, measured in seconds.
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 所有时间粒度的数据。
	BandwidthDataList []DescribeLiveBandwidthDataResResultBandwidthDataListItem `json:"BandwidthDataList"`

	// REQUIRED; The end time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询时间范围内的下行峰值带宽，单位为 Mbps。
	PeakDownBandwidth float32 `json:"PeakDownBandwidth"`

	// REQUIRED; 查询时间范围内的上行峰值带宽，单位为 Mbps。
	PeakUpBandwidth float32 `json:"PeakUpBandwidth"`

	// REQUIRED; The start time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	StartTime string `json:"StartTime"`

	// 按维度拆分后的数据。 :::tip 当配置了数据拆分的维度时，对应的维度参数传入多个值才会返回按维度拆分的数据。 :::
	BandwidthDetailDataList []*DescribeLiveBandwidthDataResResultBandwidthDetailDataListItem `json:"BandwidthDetailDataList,omitempty"`

	// 数据拆分的维度，维度说明如下所示。
	// * Domain：域名；
	// * ISP：运营商；
	// * Protocol：推拉流协议。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// Identifiers of operators providing network access services. By default, all operators are indicated. Supported operators
	// are as follows.
	// * unicom: China Unicom;
	// * railcom: China Railway Telecom;
	// * telecom: China Telecom;
	// * mobile: China Mobile;
	// * cernet: China Broadcasting Network;
	// * tianwei: China Tianwei;
	// * alibaba: Alibaba Group;
	// * tencent: Tencent Holdings;
	// * drpeng: Dr. Peng Telecom & Media Group;
	// * btvn: Broadcasting Television Network (BTVN);
	// * huashu: Huashu Media;
	// * other: Denotes other/unspecified options.
	ISPList []*string `json:"ISPList,omitempty"`

	// byteplus比火山多了CMAF协议
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// List of regions to which CDN node IPs belong.
	RegionList []*DescribeLiveBandwidthDataResResultRegionListItem `json:"RegionList,omitempty"`

	// List of regions to which client IPs belong.
	UserRegionList []*DescribeLiveBandwidthDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveBandwidthDataResResultBandwidthDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的下行峰值带宽，单位为 Mbps。
	DownBandwidth float32 `json:"DownBandwidth"`

	// REQUIRED; The start time of each time granularity when data is aggregated by time granularity, formatted in RFC3339 UTC
	// time with second precision.
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内的上行峰值带宽，单位为 Mbps。
	UpBandwidth float32 `json:"UpBandwidth"`
}

type DescribeLiveBandwidthDataResResultBandwidthDetailDataListItem struct {

	// REQUIRED; 按维度进行数据拆分后，当前维度下所有时间粒度的数据。
	BandwidthDataList []DescribeLiveBandwidthDataResResultBandwidthDetailDataListPropertiesItemsItem `json:"BandwidthDataList"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的下行峰值带宽，单位为 Mbps。
	PeakDownBandwidth float32 `json:"PeakDownBandwidth"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的上行峰值带宽，单位为 Mbps。
	PeakUpBandwidth float32 `json:"PeakUpBandwidth"`

	// 按域名维度进行数据拆分时的域名信息。
	Domain *string `json:"Domain,omitempty"`

	// 按运营商维度进行数据拆分时的运营商信息。
	ISP *string `json:"ISP,omitempty"`

	// 按推拉流协议维度进行数据拆分时的协议信息。
	Protocol *string `json:"Protocol,omitempty"`
}

type DescribeLiveBandwidthDataResResultBandwidthDetailDataListPropertiesItemsItem struct {

	// REQUIRED; 下行带宽，单位为 Mbps
	DownBandwidth float32 `json:"DownBandwidth"`

	// REQUIRED; The start time of each time granularity when data is aggregated by time granularity, formatted in RFC3339 UTC
	// time with second precision.
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 上行带宽，单位为 Mbps
	UpBandwidth float32 `json:"UpBandwidth"`
}

type DescribeLiveBandwidthDataResResultRegionListItem struct {

	// The regional identifier in regional information.
	Area *string `json:"Area,omitempty"`

	// The country identifier in the regional information.
	Country *string `json:"Country,omitempty"`

	// The province identifier in the regional information.
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveBandwidthDataResResultUserRegionListItem struct {

	// The regional identifier in regional information.
	Area *string `json:"Area,omitempty"`

	// The country identifier in the regional information.
	Country *string `json:"Country,omitempty"`

	// The province identifier in the regional information.
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveISPDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveISPDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveISPDataResResult `json:"Result"`
}

type DescribeLiveISPDataResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`
}

type DescribeLiveISPDataResResult struct {

	// REQUIRED; 运营商信息，视频直播提供的网络运营商标识，支持的运营商如下所示。
	// * unicom：联通；
	// * railcom：铁通；
	// * telecom：电信；
	// * mobile：移动；
	// * cernet：教育网；
	// * tianwei：天威；
	// * alibaba：阿里巴巴；
	// * tencent：腾讯；
	// * drpeng：鹏博士；
	// * btvn：广电；
	// * huashu：华数；
	// * other：其他。
	ISPList []DescribeLiveISPDataResResultISPListItem `json:"ISPList"`
}

type DescribeLiveISPDataResResultISPListItem struct {

	// REQUIRED; 运营商标识符。
	Code string `json:"Code"`

	// REQUIRED; 运营商名称。
	Name string `json:"Name"`
}

type DescribeLiveLogDataBody struct {

	// REQUIRED; The end time of your query's range (UTC time in RFC 3339 format with second precision).
	EndTime string `json:"EndTime"`

	// REQUIRED; The start time of your query's range (UTC time in RFC 3339 format with second precision).
	StartTime string `json:"StartTime"`

	// REQUIRED; 日志类型，支持的类型如下所示。
	// * pull：拉流日志；
	// * push：推流日志；
	// * source：回源日志；
	// * relay：拉流转推日志。
	Type string `json:"Type"`

	// 域名列表，默认为空，表示查询您视频直播产品下所有域名的日志文件信息。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的域名。 :::tip
	// 日志类型为拉流转推日志（Type 取值为 relay）时，该参数无效。 :::
	DomainList []*string `json:"DomainList,omitempty"`

	// 查询数据的页码，默认为 1，表示查询第一页的数据。
	PageNum *int32 `json:"PageNum,omitempty"`

	// 每页显示的数据条数，默认为 20，最大值为 1000。
	PageSize *int32 `json:"PageSize,omitempty"`
}

type DescribeLiveLogDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveLogDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveLogDataResResult `json:"Result"`
}

type DescribeLiveLogDataResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`
}

type DescribeLiveLogDataResResult struct {

	// REQUIRED; The end time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	EndTime string `json:"EndTime"`

	// REQUIRED; 日志文件的信息列表。
	LogInfoList []DescribeLiveLogDataResResultLogInfoListItem `json:"LogInfoList"`

	// REQUIRED; 数据分页信息。
	Pagination DescribeLiveLogDataResResultPagination `json:"Pagination"`

	// REQUIRED; The start time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	StartTime string `json:"StartTime"`

	// REQUIRED; 日志类型，类型说明如下所示。
	// * pull：拉流日志
	// * push：推流日志
	// * source：回源日志
	// * relay：拉流转推日志
	Type string `json:"Type"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`
}

type DescribeLiveLogDataResResultLogInfoListItem struct {

	// REQUIRED; 日志文件对应的小时区间，RFC3339 格式的时间戳，精度为秒。
	DateTime string `json:"DateTime"`

	// REQUIRED; 日志文件下载链接。
	DownloadURL string `json:"DownloadUrl"`

	// REQUIRED; 日志文件名称，日志文件命名规则如下。
	// * 与域名相关时：加速域名年月日时间开始时间结束文件拆分序号。例如，push.example.com_2023_08_11_000000_010000_0.gz；
	// * 与域名无关时：年月日时间开始时间结束_文件拆分序号。例如，2023_08_11_000000_010000_0.gz； :::tip 如果某个小时内，当前事件产生的日志大于 150 万条，则会生成为多个日志文件，用文件名最后的序号标注日志文件顺序，例如，2023_08_11_000000_010000_0.gz、2023_08_11_000000_010000_1.gz。
	// :::
	LogName string `json:"LogName"`

	// REQUIRED; 日志文件大小，单位为 byte。
	LogSize int32 `json:"LogSize"`

	// 域名。 :::tip 查询拉流转推日志（Type 取值为 relay）时，该字段为空。 :::
	Domain *string `json:"Domain,omitempty"`
}

// DescribeLiveLogDataResResultPagination - 数据分页信息。
type DescribeLiveLogDataResResultPagination struct {

	// REQUIRED; 当前所在分页的页码。
	PageNum int32 `json:"PageNum"`

	// REQUIRED; 每页显示的数据条数。
	PageSize int32 `json:"PageSize"`

	// REQUIRED; 查询结果的数据总条数。
	TotalCount int32 `json:"TotalCount"`
}

type DescribeLiveMetricBandwidthDataBody struct {

	// REQUIRED; The end time of your query's range (UTC time in RFC 3339 format with second precision).
	EndTime string `json:"EndTime"`

	// REQUIRED; The start time of your query's range (UTC time in RFC 3339 format with second precision).
	StartTime string `json:"StartTime"`

	// The granularity of data aggregation, measured in seconds, with the following supported options:
	// * 60: 1 minute. When aggregated every 1 minute, the maximum time span for a single query is 24 hours, and the historical
	// query time range is 366 days;
	// * 300: (default) 5 minutes. When aggregated every 5 minutes, the maximum time span for a single query is 31 days, and the
	// historical query time range is 366 days;
	// * 3600: 1 hour. When aggregated every 1 hour, the maximum time span for a single query is 93 days, and the historical query
	// time range is 366 days.
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// The Application Name must match the value of the AppName field in the live stream URL. It can include uppercase letters
	// (A-Z), lowercase letters (a-z), numbers (0-9), underscores (_), hyphens (-), and
	// periods (.), with a length ranging from 1 to 30 characters.
	// :::tip When querying stream granularity data, both the App and Stream parameters are required. :::
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，默认为空表示不按维度进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * Protocol：推拉流协议；
	// * ISP：运营商。
	// :::tip 配置数据拆分的维度时，对应的维度参数传入多个值时才会返回按此维度拆分的数据。例如，配置按 Domain 进行数据拆分时， DomainList 传入多个 Domain 值时，才会返回按 Domain 拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，默认为空，表示查询您视频直播产品下所有域名的带宽监控数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// Identifiers of operators providing network access services. By default, all operators are indicated. Supported operators
	// are as follows.
	// * unicom: China Unicom;
	// * railcom: China Railway Telecom;
	// * telecom: China Telecom;
	// * mobile: China Mobile;
	// * cernet: China Education and Research Network (CERNET);
	// * tianwei: China Tianwei;
	// * alibaba: Alibaba Group;
	// * tencent: Tencent Holdings;
	// * drpeng: Dr. Peng Telecom & Media Group;
	// * btvn: China Broadcasting Network;
	// * huashu: Huashu Media;
	// * other: Denotes other/unspecified options.
	// If you need to obtain the identifiers of various operators, you can call the DescribeLiveISPData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveispdata].
	ISPList []*string `json:"ISPList,omitempty"`

	// byteplus比火山多了CMAF协议
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// List of regions to which CDN node IPs belong, by default indicating all regions. :::tipRegionList and UserRegionList cannot
	// be used together in the same request. :::
	RegionList []*DescribeLiveMetricBandwidthDataBodyRegionListItem `json:"RegionList,omitempty"`

	// The Application Name must match the value of the AppName field in the live stream URL. It can include uppercase letters
	// (A-Z), lowercase letters (a-z), numbers (0-9), underscores (_), hyphens (-), and
	// periods (.), with a length ranging from 1 to 30 characters.
	// :::tip When querying stream granularity data, both the App and Stream parameters are required. :::
	Stream *string `json:"Stream,omitempty"`

	// List of regions to which client IPs belong, by default indicating all regions.
	// :::tipRegionList and UserRegionList cannot be used together in a single request. :::
	UserRegionList []*DescribeLiveMetricBandwidthDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveMetricBandwidthDataBodyRegionListItem struct {

	// The identifier for the major region in the regional information can be obtained by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata].
	Area *string `json:"Area,omitempty"`

	// The country identifier in the regional information can be obtained by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata].
	// When filtering by
	// country, both 'Area' and 'Country' need to be passed in simultaneously.
	Country *string `json:"Country,omitempty"`

	// The province identifier in the regional information is currently not supported for countries or regions outside mainland
	// China, Hong Kong, Macao, and Taiwan. You can obtain the identifier information
	// by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata]. When
	// filtering by province, you need to simultaneously pass in Area, Country, and
	// Province
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveMetricBandwidthDataBodyUserRegionListItem struct {

	// The identifier for the major region in the regional information can be obtained by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata].
	Area *string `json:"Area,omitempty"`

	// The country identifier in the regional information can be obtained by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata].
	// When filtering by
	// country, both 'Area' and 'Country' need to be passed in simultaneously.
	Country *string `json:"Country,omitempty"`

	// The province identifier in the regional information is currently not supported for countries or regions outside mainland
	// China, Hong Kong, Macao, and Taiwan. You can obtain the identifier information
	// by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata]. When
	// filtering by province, you need to simultaneously pass in Area, Country, and
	// Province
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveMetricBandwidthDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveMetricBandwidthDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveMetricBandwidthDataResResult `json:"Result"`
}

type DescribeLiveMetricBandwidthDataResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`
}

type DescribeLiveMetricBandwidthDataResResult struct {

	// REQUIRED; Data granularity, measured in seconds.
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 所有时间粒度的数据。
	BandwidthDataList []DescribeLiveMetricBandwidthDataResResultBandwidthDataListItem `json:"BandwidthDataList"`

	// REQUIRED; The end time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询时间范围内的下行峰值，单位为 Mbps。
	PeakDownBandwidth float32 `json:"PeakDownBandwidth"`

	// REQUIRED; 查询时间范围内的上行峰值，单位为 Mbps。
	PeakUpBandwidth float32 `json:"PeakUpBandwidth"`

	// REQUIRED; The start time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	StartTime string `json:"StartTime"`

	// The application name when querying stream granularity data.
	App *string `json:"App,omitempty"`

	// 按维度拆分后的数据。
	BandwidthDetailDataList []*DescribeLiveMetricBandwidthDataResResultBandwidthDetailDataListItem `json:"BandwidthDetailDataList,omitempty"`

	// 数据拆分的维度，维度说明如下所示。
	// * Domain：域名；
	// * Protocol：推拉流协议；
	// * ISP：运营商。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// Identifiers of operators providing network access services. By default, all operators are indicated. Supported operators
	// are as follows.
	// * unicom: China Unicom;
	// * railcom: China Railway Telecom;
	// * telecom: China Telecom;
	// * mobile: China Mobile;
	// * cernet: China Broadcasting Network;
	// * tianwei: China Tianwei;
	// * alibaba: Alibaba Group;
	// * tencent: Tencent Holdings;
	// * drpeng: Dr. Peng Telecom & Media Group;
	// * btvn: Broadcasting Television Network (BTVN);
	// * huashu: Huashu Media;
	// * other: Denotes other/unspecified options.
	ISPList []*string `json:"ISPList,omitempty"`

	// byteplus比火山多了CMAF协议
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// List of regions to which CDN node IPs belong.
	RegionList []*DescribeLiveMetricBandwidthDataResResultRegionListItem `json:"RegionList,omitempty"`

	// The stream name when querying stream granularity data.
	Stream *string `json:"Stream,omitempty"`

	// List of regions to which client IPs belong.
	UserRegionList []*DescribeLiveMetricBandwidthDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveMetricBandwidthDataResResultBandwidthDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的下行峰值带宽，单位为 Mbps。
	DownBandwidth float32 `json:"DownBandwidth"`

	// REQUIRED; The start time of each time granularity when data is aggregated by time granularity, formatted in RFC3339 UTC
	// time with second precision.
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内的上行峰值带宽，单位为 Mbps。
	UpBandwidth float32 `json:"UpBandwidth"`
}

type DescribeLiveMetricBandwidthDataResResultBandwidthDetailDataListItem struct {

	// REQUIRED; 按维度进行数据拆分后，当前维度下所有时间粒度的数据。
	BandwidthDataList []DescribeLiveMetricBandwidthDataResResultBandwidthDetailDataListPropertiesItemsItem `json:"BandwidthDataList"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的下行峰值带宽，单位为 Mbps。
	PeakDownBandwidth float32 `json:"PeakDownBandwidth"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的上行峰值带宽，单位为 Mbps。
	PeakUpBandwidth float32 `json:"PeakUpBandwidth"`

	// 按域名维度进行数据拆分时的域名信息。
	Domain *string `json:"Domain,omitempty"`

	// 按运营商维度进行数据拆分时的运营商信息。
	ISP *string `json:"ISP,omitempty"`

	// 按推拉流协议维度进行数据拆分时的协议信息。
	Protocol *string `json:"Protocol,omitempty"`
}

type DescribeLiveMetricBandwidthDataResResultBandwidthDetailDataListPropertiesItemsItem struct {

	// REQUIRED; 下行带宽，单位为 Mbps
	DownBandwidth float32 `json:"DownBandwidth"`

	// REQUIRED; The start time of each time granularity when data is aggregated by time granularity, formatted in RFC3339 UTC
	// time with second precision.
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 上行带宽，单位为 Mbps
	UpBandwidth float32 `json:"UpBandwidth"`
}

type DescribeLiveMetricBandwidthDataResResultRegionListItem struct {

	// The regional identifier in regional information.
	Area *string `json:"Area,omitempty"`

	// The country identifier in the regional information.
	Country *string `json:"Country,omitempty"`

	// The province identifier in the regional information.
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveMetricBandwidthDataResResultUserRegionListItem struct {

	// The regional identifier in regional information.
	Area *string `json:"Area,omitempty"`

	// The country identifier in the regional information.
	Country *string `json:"Country,omitempty"`

	// The province identifier in the regional information.
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveMetricTrafficDataBody struct {

	// REQUIRED; The end time of your query's range (UTC time in RFC 3339 format with second precision).
	EndTime string `json:"EndTime"`

	// REQUIRED; The start time of your query's range (UTC time in RFC 3339 format with second precision).
	StartTime string `json:"StartTime"`

	// The granularity of data aggregation, measured in seconds, with the following supported options:
	// * 60: 1 minute. When aggregated every 1 minute, the maximum time span for a single query is 24 hours, and the historical
	// query time range is 366 days;
	// * 300: (default) 5 minutes. When aggregated every 5 minutes, the maximum time span for a single query is 31 days, and the
	// historical query time range is 366 days;
	// * 3600: 1 hour. When aggregated every 1 hour, the maximum time span for a single query is 93 days, and the historical query
	// time range is 366 days.
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// The Application Name must match the value of the AppName field in the live stream URL. It can include uppercase letters
	// (A-Z), lowercase letters (a-z), numbers (0-9), underscores (_), hyphens (-), and
	// periods (.), with a length ranging from 1 to 30 characters.
	// :::tip When querying stream granularity data, both the App and Stream parameters are required. :::
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，默认为空表示不按维度进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * Protocol：推拉流协议；
	// * ISP：运营商。
	// :::tip 配置数据拆分的维度时，对应的维度参数传入多个值时才会返回按此维度拆分的数据。例如，配置按 Domain 进行数据拆分时， DomainList 传入多个 Domain 值时，才会返回按 Domain 拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，默认为空，表示查询您视频直播产品下所有域名的流量监控数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// Identifiers of operators providing network access services. By default, all operators are indicated. Supported operators
	// are as follows.
	// * unicom: China Unicom;
	// * railcom: China Railway Telecom;
	// * telecom: China Telecom;
	// * mobile: China Mobile;
	// * cernet: China Education and Research Network (CERNET);
	// * tianwei: China Tianwei;
	// * alibaba: Alibaba Group;
	// * tencent: Tencent Holdings;
	// * drpeng: Dr. Peng Telecom & Media Group;
	// * btvn: China Broadcasting Network;
	// * huashu: Huashu Media;
	// * other: Denotes other/unspecified options.
	// If you need to obtain the identifiers of various operators, you can call the DescribeLiveISPData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveispdata].
	ISPList []*string `json:"ISPList,omitempty"`

	// byteplus比火山多了CMAF协议
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// List of regions to which CDN node IPs belong, by default indicating all regions. :::tipRegionList and UserRegionList cannot
	// be used together in the same request. :::
	RegionList []*DescribeLiveMetricTrafficDataBodyRegionListItem `json:"RegionList,omitempty"`

	// 指定查询的流量数据为闲时或忙时，缺省情况下为查询全部数据，支持的取值如下。
	// * busy：忙时；
	// * free：闲时。
	Stage *string `json:"Stage,omitempty"`

	// The Application Name must match the value of the AppName field in the live stream URL. It can include uppercase letters
	// (A-Z), lowercase letters (a-z), numbers (0-9), underscores (_), hyphens (-), and
	// periods (.), with a length ranging from 1 to 30 characters.
	// :::tip When querying stream granularity data, both the App and Stream parameters are required. :::
	Stream *string `json:"Stream,omitempty"`

	// List of regions to which client IPs belong, by default indicating all regions.
	// :::tipRegionList and UserRegionList cannot be used together in a single request. :::
	UserRegionList []*DescribeLiveMetricTrafficDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveMetricTrafficDataBodyRegionListItem struct {

	// The identifier for the major region in the regional information can be obtained by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata].
	Area *string `json:"Area,omitempty"`

	// The country identifier in the regional information can be obtained by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata].
	// When filtering by
	// country, both 'Area' and 'Country' need to be passed in simultaneously.
	Country *string `json:"Country,omitempty"`

	// The province identifier in the regional information is currently not supported for countries or regions outside mainland
	// China, Hong Kong, Macao, and Taiwan. You can obtain the identifier information
	// by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata]. When
	// filtering by province, you need to simultaneously pass in Area, Country, and
	// Province
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveMetricTrafficDataBodyUserRegionListItem struct {

	// The identifier for the major region in the regional information can be obtained by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata].
	Area *string `json:"Area,omitempty"`

	// The country identifier in the regional information can be obtained by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata].
	// When filtering by
	// country, both 'Area' and 'Country' need to be passed in simultaneously.
	Country *string `json:"Country,omitempty"`

	// The province identifier in the regional information is currently not supported for countries or regions outside mainland
	// China, Hong Kong, Macao, and Taiwan. You can obtain the identifier information
	// by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata]. When
	// filtering by province, you need to simultaneously pass in Area, Country, and
	// Province
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveMetricTrafficDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveMetricTrafficDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveMetricTrafficDataResResult `json:"Result"`
}

type DescribeLiveMetricTrafficDataResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`
}

type DescribeLiveMetricTrafficDataResResult struct {

	// REQUIRED; Data granularity, measured in seconds.
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; The end time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	EndTime string `json:"EndTime"`

	// REQUIRED; The start time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	StartTime string `json:"StartTime"`

	// REQUIRED; 查询时间范围内的下行总流量，单位为 GB。
	TotalDownTraffic float32 `json:"TotalDownTraffic"`

	// REQUIRED; 查询时间范围内的上行总流量，单位为 GB。
	TotalUpTraffic float32 `json:"TotalUpTraffic"`

	// REQUIRED; 所有时间粒度的数据。
	TrafficDataList []DescribeLiveMetricTrafficDataResResultTrafficDataListItem `json:"TrafficDataList"`

	// The application name when querying stream granularity data.
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，维度说明如下所示。
	// * Domain：域名；
	// * Protocol：推拉流协议；
	// * ISP：运营商。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// Identifiers of operators providing network access services. By default, all operators are indicated. Supported operators
	// are as follows.
	// * unicom: China Unicom;
	// * railcom: China Railway Telecom;
	// * telecom: China Telecom;
	// * mobile: China Mobile;
	// * cernet: China Broadcasting Network;
	// * tianwei: China Tianwei;
	// * alibaba: Alibaba Group;
	// * tencent: Tencent Holdings;
	// * drpeng: Dr. Peng Telecom & Media Group;
	// * btvn: Broadcasting Television Network (BTVN);
	// * huashu: Huashu Media;
	// * other: Denotes other/unspecified options.
	ISPList []*string `json:"ISPList,omitempty"`

	// byteplus比火山多了CMAF协议
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// List of regions to which CDN node IPs belong.
	RegionList []*DescribeLiveMetricTrafficDataResResultRegionListItem `json:"RegionList,omitempty"`

	// 流量数据为闲时或忙时，取值说明如下。
	// * busy：忙时；
	// * free：闲时。
	Stage *string `json:"Stage,omitempty"`

	// The stream name when querying stream granularity data.
	Stream *string `json:"Stream,omitempty"`

	// 按维度拆分后的数据。 :::tip 配置数据拆分维度时，对应的维度参数传入多个值时会返回按维度进行拆分的数据；对应的维度只传入一个值时不返回按此维度进行拆分的数据。 :::
	TrafficDetailDataList []*DescribeLiveMetricTrafficDataResResultTrafficDetailDataListItem `json:"TrafficDetailDataList,omitempty"`

	// List of regions to which client IPs belong.
	UserRegionList []*DescribeLiveMetricTrafficDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveMetricTrafficDataResResultRegionListItem struct {

	// The regional identifier in regional information.
	Area *string `json:"Area,omitempty"`

	// The country identifier in the regional information.
	Country *string `json:"Country,omitempty"`

	// The province identifier in the regional information.
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveMetricTrafficDataResResultTrafficDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内产生的总下行流量，单位 GB。
	DownTraffic float32 `json:"DownTraffic"`

	// REQUIRED; The start time of each time granularity when data is aggregated by time granularity, formatted in RFC3339 UTC
	// time with second precision.
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内产生的总上行流量，单位 GB。
	UpTraffic float32 `json:"UpTraffic"`
}

type DescribeLiveMetricTrafficDataResResultTrafficDetailDataListItem struct {

	// REQUIRED; 按维度进行数据拆分后，当前维度的下行总流量，单位为 GB。
	TotalDownTraffic float32 `json:"TotalDownTraffic"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的上行总流量，单位为 GB。
	TotalUpTraffic float32 `json:"TotalUpTraffic"`

	// REQUIRED; 按维度进行数据拆分后，当前维度下所有时间粒度的数据。
	TrafficDataList []DescribeLiveMetricTrafficDataResResultTrafficDetailDataListPropertiesItemsItem `json:"TrafficDataList"`

	// 按域名维度进行数据拆分时的域名信息。
	Domain *string `json:"Domain,omitempty"`

	// 按运营商维度进行数据拆分时的运营商信息。
	ISP *string `json:"ISP,omitempty"`

	// 按推拉流协议维度进行数据拆分时的协议信息。
	Protocol *string `json:"Protocol,omitempty"`
}

type DescribeLiveMetricTrafficDataResResultTrafficDetailDataListPropertiesItemsItem struct {

	// REQUIRED; 下行流量，单位 GB
	DownTraffic float32 `json:"DownTraffic"`

	// REQUIRED; The start time of each time granularity when data is aggregated by time granularity, formatted in RFC3339 UTC
	// time with second precision.
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 上行流量，单位 GB
	UpTraffic float32 `json:"UpTraffic"`
}

type DescribeLiveMetricTrafficDataResResultUserRegionListItem struct {

	// The regional identifier in regional information.
	Area *string `json:"Area,omitempty"`

	// The country identifier in the regional information.
	Country *string `json:"Country,omitempty"`

	// The province identifier in the regional information.
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveP95PeakBandwidthDataBody struct {

	// REQUIRED; The end time of your query's range (UTC time in RFC 3339 format with second precision).
	EndTime string `json:"EndTime"`

	// REQUIRED; The start time of your query's range (UTC time in RFC 3339 format with second precision).
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，当前接口默认且仅支持按 300 秒进行数据拆分。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 域名列表，默认为空，表示查询您视频直播产品下所有域名的 95 峰值带宽用量数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// Push-pull streaming protocol. By default, all protocol types are indicated. Supported protocols are as follows.
	// * HTTP-FLV: A push-pull streaming protocol based on the HTTP protocol, using the FLV format for video transmission.
	// * HTTP-HLS: A push-pull streaming protocol based on the HTTP protocol, using the TS format for video transmission.
	// * RTMP: Real-Time Messaging Protocol for real-time message transmission.
	// * RTM: Real-Time Media protocol for ultra-low latency live streaming.
	// * SRT: Secure Reliable Transport protocol for secure and reliable streaming.
	// * QUIC: Quick UDP Internet Connections, a new low-latency internet transmission protocol based on UDP.
	// * CMAF: Common Media Application Format, a versatile streaming protocol.
	// :::tip If querying the QUIC protocol, other protocols cannot be queried simultaneously. :::
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// List of regions to which CDN node IPs belong, by default indicating all regions. :::tipRegionList and UserRegionList cannot
	// be used together in the same request. :::
	RegionList []*DescribeLiveP95PeakBandwidthDataBodyRegionListItem `json:"RegionList,omitempty"`

	// List of regions to which client IPs belong, by default indicating all regions.
	// :::tipRegionList and UserRegionList cannot be used together in a single request. :::
	UserRegionList []*DescribeLiveP95PeakBandwidthDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveP95PeakBandwidthDataBodyRegionListItem struct {

	// The identifier for the major region in the regional information can be obtained by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata].
	Area *string `json:"Area,omitempty"`

	// The country identifier in the regional information can be obtained by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata].
	// When filtering by
	// country, both 'Area' and 'Country' need to be passed in simultaneously.
	Country *string `json:"Country,omitempty"`

	// The province identifier in the regional information is currently not supported for countries or regions outside mainland
	// China, Hong Kong, Macao, and Taiwan. You can obtain the identifier information
	// by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata]. When
	// filtering by province, you need to simultaneously pass in Area, Country, and
	// Province
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveP95PeakBandwidthDataBodyUserRegionListItem struct {

	// The identifier for the major region in the regional information can be obtained by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata].
	Area *string `json:"Area,omitempty"`

	// The country identifier in the regional information can be obtained by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata].
	// When filtering by
	// country, both 'Area' and 'Country' need to be passed in simultaneously.
	Country *string `json:"Country,omitempty"`

	// The province identifier in the regional information is currently not supported for countries or regions outside mainland
	// China, Hong Kong, Macao, and Taiwan. You can obtain the identifier information
	// by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata]. When
	// filtering by province, you need to simultaneously pass in Area, Country, and
	// Province
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveP95PeakBandwidthDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveP95PeakBandwidthDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveP95PeakBandwidthDataResResult `json:"Result"`
}

type DescribeLiveP95PeakBandwidthDataResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`
}

type DescribeLiveP95PeakBandwidthDataResResult struct {

	// REQUIRED; Data granularity, measured in seconds.
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; The end time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	EndTime string `json:"EndTime"`

	// REQUIRED; 时间范围内的上下行 95 峰值带宽总和。 :::tip 如果请求时，Regionlist中传入多个 region，则返回这些 region 的上下行带宽 95 峰值总和。 :::
	P95PeakBandwidth float32 `json:"P95PeakBandwidth"`

	// REQUIRED; 95 峰值带宽的时间戳，RFC3339 格式的时间戳，精度为秒。
	P95PeakTimestamp string `json:"P95PeakTimestamp"`

	// REQUIRED; The start time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	StartTime string `json:"StartTime"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// Push-pull streaming protocol. By default, all protocol types are indicated. Supported protocols are as follows.
	// * HTTP-FLV: A push-pull streaming protocol based on the HTTP protocol, using the FLV format for video transmission.
	// * HTTP-HLS: A push-pull streaming protocol based on the HTTP protocol, using the TS format for video transmission.
	// * RTMP: Real-Time Messaging Protocol for real-time message transmission.
	// * RTM: Real-Time Media protocol for ultra-low latency live streaming.
	// * SRT: Secure Reliable Transport protocol for secure and reliable streaming.
	// * QUIC: Quick UDP Internet Connections, a new low-latency internet transmission protocol based on UDP.
	// * CMAF: Common Media Application Format, a versatile streaming protocol.
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// List of regions to which CDN node IPs belong.
	RegionList []*DescribeLiveP95PeakBandwidthDataResResultRegionListItem `json:"RegionList,omitempty"`

	// List of regions to which client IPs belong.
	UserRegionList []*DescribeLiveP95PeakBandwidthDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveP95PeakBandwidthDataResResultRegionListItem struct {

	// The regional identifier in regional information.
	Area *string `json:"Area,omitempty"`

	// The country identifier in the regional information.
	Country *string `json:"Country,omitempty"`

	// The province identifier in the regional information.
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveP95PeakBandwidthDataResResultUserRegionListItem struct {

	// The regional identifier in regional information.
	Area *string `json:"Area,omitempty"`

	// The country identifier in the regional information.
	Country *string `json:"Country,omitempty"`

	// The province identifier in the regional information.
	Province *string `json:"Province,omitempty"`
}

type DescribeLivePlayStatusCodeDataBody struct {

	// REQUIRED; The end time of your query's range (UTC time in RFC 3339 format with second precision).
	EndTime string `json:"EndTime"`

	// REQUIRED; The start time of your query's range (UTC time in RFC 3339 format with second precision).
	StartTime string `json:"StartTime"`

	// The granularity of data aggregation, measured in seconds, with the following supported options:
	// * 60: 1 minute. When aggregated every 1 minute, the maximum time span for a single query is 24 hours, and the historical
	// query time range is 366 days;
	// * 300: (default) 5 minutes. When aggregated every 5 minutes, the maximum time span for a single query is 31 days, and the
	// historical query time range is 366 days;
	// * 3600: 1 hour. When aggregated every 1 hour, the maximum time span for a single query is 93 days, and the historical query
	// time range is 366 days.
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 数据拆分的维度，默认为空表示不按维度进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * ISP：运营商。
	// :::tip 配置数据拆分的维度时，对应的维度参数传入多个值时才会返回按此维度拆分的数据。例如，配置按 Domain 进行数据拆分时， DomainList 传入多个 Domain 值时，才会返回按 Domain 拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，默认为空时表示查询所有域名下产生的请求状态码占比数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取待查询请求状态码占比数据的域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// Identifiers of operators providing network access services. By default, all operators are indicated. Supported operators
	// are as follows.
	// * unicom: China Unicom;
	// * railcom: China Railway Telecom;
	// * telecom: China Telecom;
	// * mobile: China Mobile;
	// * cernet: China Education and Research Network (CERNET);
	// * tianwei: China Tianwei;
	// * alibaba: Alibaba Group;
	// * tencent: Tencent Holdings;
	// * drpeng: Dr. Peng Telecom & Media Group;
	// * btvn: China Broadcasting Network;
	// * huashu: Huashu Media;
	// * other: Denotes other/unspecified options.
	// If you need to obtain the identifiers of various operators, you can call the DescribeLiveISPData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveispdata].
	ISPList []*string `json:"ISPList,omitempty"`

	// List of regions to which CDN node IPs belong, by default indicating all regions. :::tipRegionList and UserRegionList cannot
	// be used together in the same request. :::
	RegionList []*DescribeLivePlayStatusCodeDataBodyRegionListItem `json:"RegionList,omitempty"`

	// 请求类型，取值及含义如下所示。
	// * Access：（默认值）推流请求和拉流请求；
	// * Source：回源请求。
	Type *string `json:"Type,omitempty"`

	// List of regions to which client IPs belong, by default indicating all regions.
	// :::tipRegionList and UserRegionList cannot be used together in a single request. :::
	UserRegionList []*DescribeLivePlayStatusCodeDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLivePlayStatusCodeDataBodyRegionListItem struct {

	// The identifier for the major region in the regional information can be obtained by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata].
	Area *string `json:"Area,omitempty"`

	// The country identifier in the regional information can be obtained by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata].
	// When filtering by
	// country, both 'Area' and 'Country' need to be passed in simultaneously.
	Country *string `json:"Country,omitempty"`

	// The province identifier in the regional information is currently not supported for countries or regions outside mainland
	// China, Hong Kong, Macao, and Taiwan. You can obtain the identifier information
	// by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata]. When
	// filtering by province, you need to simultaneously pass in Area, Country, and
	// Province
	Province *string `json:"Province,omitempty"`
}

type DescribeLivePlayStatusCodeDataBodyUserRegionListItem struct {

	// The identifier for the major region in the regional information can be obtained by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata].
	Area *string `json:"Area,omitempty"`

	// The country identifier in the regional information can be obtained by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata].
	// When filtering by
	// country, both 'Area' and 'Country' need to be passed in simultaneously.
	Country *string `json:"Country,omitempty"`

	// The province identifier in the regional information is currently not supported for countries or regions outside mainland
	// China, Hong Kong, Macao, and Taiwan. You can obtain the identifier information
	// by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata]. When
	// filtering by province, you need to simultaneously pass in Area, Country, and
	// Province
	Province *string `json:"Province,omitempty"`
}

type DescribeLivePlayStatusCodeDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLivePlayStatusCodeDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLivePlayStatusCodeDataResResult `json:"Result"`
}

type DescribeLivePlayStatusCodeDataResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`
}

type DescribeLivePlayStatusCodeDataResResult struct {

	// REQUIRED; Data granularity, measured in seconds.
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; The end time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	EndTime string `json:"EndTime"`

	// REQUIRED; The start time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	StartTime string `json:"StartTime"`

	// REQUIRED; 所有时间粒度的数据。
	StatusDataList []DescribeLivePlayStatusCodeDataResResultStatusDataListItem `json:"StatusDataList"`

	// REQUIRED; 当前查询条件下的状态码占比数据。
	StatusSummaryDataList []DescribeLivePlayStatusCodeDataResResultStatusSummaryDataListItem `json:"StatusSummaryDataList"`

	// REQUIRED; 请求类型，取值及含义如下所示。
	// * Access：推流请求和拉流请求；
	// * Source：回源请求。
	Type string `json:"Type"`

	// 数据拆分的维度，维度说明如下所示。
	// * Domain：域名；
	// * ISP：运营商。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// Identifiers of operators providing network access services. By default, all operators are indicated. Supported operators
	// are as follows.
	// * unicom: China Unicom;
	// * railcom: China Railway Telecom;
	// * telecom: China Telecom;
	// * mobile: China Mobile;
	// * cernet: China Broadcasting Network;
	// * tianwei: China Tianwei;
	// * alibaba: Alibaba Group;
	// * tencent: Tencent Holdings;
	// * drpeng: Dr. Peng Telecom & Media Group;
	// * btvn: Broadcasting Television Network (BTVN);
	// * huashu: Huashu Media;
	// * other: Denotes other/unspecified options.
	ISPList []*string `json:"ISPList,omitempty"`

	// List of regions to which CDN node IPs belong.
	RegionList []*DescribeLivePlayStatusCodeDataResResultRegionListItem `json:"RegionList,omitempty"`

	// 按维度拆分后的数据。
	StatusDetailDataList []*DescribeLivePlayStatusCodeDataResResultStatusDetailDataListItem `json:"StatusDetailDataList,omitempty"`

	// List of regions to which client IPs belong.
	UserRegionList []*DescribeLivePlayStatusCodeDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLivePlayStatusCodeDataResResultRegionListItem struct {

	// The regional identifier in regional information.
	Area *string `json:"Area,omitempty"`

	// The country identifier in the regional information.
	Country *string `json:"Country,omitempty"`

	// The province identifier in the regional information.
	Province *string `json:"Province,omitempty"`
}

type DescribeLivePlayStatusCodeDataResResultStatusDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的状态码详细数据。
	StatusSummaryDataList []DescribeLivePlayStatusCodeDataResResultStatusDataListPropertiesItemsItem `json:"StatusSummaryDataList"`

	// REQUIRED; The start time of each time granularity when data is aggregated by time granularity, formatted in RFC3339 UTC
	// time with second precision.
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLivePlayStatusCodeDataResResultStatusDataListPropertiesItemsItem struct {

	// 占比。
	Percent *float32 `json:"Percent,omitempty"`

	// 状态码。
	StatusCode *int32 `json:"StatusCode,omitempty"`

	// 出现次数。
	Value *int32 `json:"Value,omitempty"`
}

type DescribeLivePlayStatusCodeDataResResultStatusDetailDataListItem struct {

	// 拉流域名。
	Domain *string `json:"Domain,omitempty"`

	// 运营商。
	ISP *string `json:"ISP,omitempty"`

	// 每个时间点的粒度数据。
	StatusDataList []*DescribeLivePlayStatusCodeDataResResultStatusDetailDataListPropertiesItemsItem `json:"StatusDataList,omitempty"`
}

type DescribeLivePlayStatusCodeDataResResultStatusDetailDataListPropertiesItemsItem struct {

	// REQUIRED; 按状态码区分的数据列表。
	StatusSummaryDataList []DescribeLivePlayStatusCodeDataResResultStatusDetailDataListPropertiesItemsStatusSummaryDataListItem `json:"StatusSummaryDataList"`

	// REQUIRED; The start time of each time granularity when data is aggregated by time granularity, formatted in RFC3339 UTC
	// time with second precision.
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLivePlayStatusCodeDataResResultStatusDetailDataListPropertiesItemsStatusSummaryDataListItem struct {

	// 占比
	Percent *float32 `json:"Percent,omitempty"`

	// 状态码
	StatusCode *int32 `json:"StatusCode,omitempty"`

	// 出现次数
	Value *int32 `json:"Value,omitempty"`
}

type DescribeLivePlayStatusCodeDataResResultStatusSummaryDataListItem struct {

	// 当前状态码出现次数在总状态码次数中的占比。
	Percent *float32 `json:"Percent,omitempty"`

	// 请求的状态码。
	StatusCode *int32 `json:"StatusCode,omitempty"`

	// 当前状态码出现的次数。
	Value *int32 `json:"Value,omitempty"`
}

type DescribeLivePlayStatusCodeDataResResultUserRegionListItem struct {

	// The regional identifier in regional information.
	Area *string `json:"Area,omitempty"`

	// The country identifier in the regional information.
	Country *string `json:"Country,omitempty"`

	// The province identifier in the regional information.
	Province *string `json:"Province,omitempty"`
}

type DescribeLivePullToPushBandwidthDataBody struct {

	// REQUIRED; The end time of your query's range (UTC time in RFC 3339 format with second precision).
	EndTime string `json:"EndTime"`

	// REQUIRED; The start time of your query's range (UTC time in RFC 3339 format with second precision).
	StartTime string `json:"StartTime"`

	// The granularity of data aggregation, measured in seconds, with the following supported options:
	// * 300 (default): 5 minutes. When aggregated in 5-minute intervals, the maximum time span for a single query is 31 days,
	// and for historical queries, the maximum time range is 366 days.
	// * 3600: 1 hour. When aggregated in 1-hour intervals, the maximum time span for a single query is 93 days, and for historical
	// queries, the maximum time range is 366 days.
	// * 86400: 1 day. When aggregated in 1-day intervals, the maximum time span for a single query is 93 days, and for historical
	// queries, the maximum time range is 366 days.
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 数据拆分的维度，默认为空表示不按维度进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * DstAddrType：推流地址类型。 :::tip 配置数据拆分的维度时，对应的维度参数传入多个值时才会返回按此维度拆分的数据。例如，配置按 Domain 进行数据拆分时， DomainList 传入多个 Domain 值时，才会返回按
	// Domain 拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，默认为空，表示查询您视频直播产品下所有域名下的带宽用量数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的域名。
	DomainList []*string `json:"DomainList,omitempty"`
}

type DescribeLivePullToPushBandwidthDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLivePullToPushBandwidthDataResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeLivePullToPushBandwidthDataResResult          `json:"Result,omitempty"`
}

type DescribeLivePullToPushBandwidthDataResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                                       `json:"Version"`
	Error   *DescribeLivePullToPushBandwidthDataResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeLivePullToPushBandwidthDataResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeLivePullToPushBandwidthDataResResult struct {

	// REQUIRED; Data granularity, measured in seconds.
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 所有时间粒度的数据。
	BandwidthDataList []DescribeLivePullToPushBandwidthDataResResultBandwidthDataListItem `json:"BandwidthDataList"`

	// REQUIRED; The end time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	EndTime string `json:"EndTime"`

	// REQUIRED; 当前查询条件下的拉流转推峰值带宽，单位为 Mbps。
	PeakUpBandwidth float32 `json:"PeakUpBandwidth"`

	// REQUIRED; The start time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	StartTime string `json:"StartTime"`

	// 按维度拆分后的数据。 :::tip 当配置了数据拆分的维度时，对应的维度参数传入多个值才会返回按维度拆分的数据。 :::
	BandwidthDetailDataList []*DescribeLivePullToPushBandwidthDataResResultBandwidthDetailDataListItem `json:"BandwidthDetailDataList,omitempty"`

	// 数据拆分的维度，维度说明如下。
	// * Domain：域名；
	// * DstAddrType：推流地址类型。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`
}

type DescribeLivePullToPushBandwidthDataResResultBandwidthDataListItem struct {

	// REQUIRED; The start time of each time granularity when data is aggregated by time granularity, formatted in RFC3339 UTC
	// time with second precision.
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内的拉流转推峰值带宽，单位为 Mbps。
	UpBandwidth float32 `json:"UpBandwidth"`
}

type DescribeLivePullToPushBandwidthDataResResultBandwidthDetailDataListItem struct {

	// REQUIRED; 按维度进行数据拆分后，当前维度下所有时间粒度的数据。
	BandwidthDataList []DescribeLivePullToPushBandwidthDataResResultBandwidthDetailDataListPropertiesItemsItem `json:"BandwidthDataList"`

	// REQUIRED; 查询时间范围内的维度下的拉流转推峰值带宽，单位为 Mbps。
	PeakUpBandwidth float32 `json:"PeakUpBandwidth"`

	// 按域名维度进行数据拆分时的域名信息。
	Domain *string `json:"Domain,omitempty"`

	// 按推流地址类型维度进行数据拆分时的地址类型信息。
	DstAddrType *string `json:"DstAddrType,omitempty"`
}

type DescribeLivePullToPushBandwidthDataResResultBandwidthDetailDataListPropertiesItemsItem struct {

	// REQUIRED; The start time of each time granularity when data is aggregated by time granularity, formatted in RFC3339 UTC
	// time with second precision.
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 转推带宽，单位为 Mbps
	UpBandwidth float32 `json:"UpBandwidth"`
}

type DescribeLivePullToPushDataBody struct {

	// REQUIRED; The end time of your query's range (UTC time in RFC 3339 format with second precision).
	EndTime string `json:"EndTime"`

	// REQUIRED; The start time of your query's range (UTC time in RFC 3339 format with second precision).
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 60：1 分钟。时间粒度为 1 分钟时，单次查询最大时间跨度为 1 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询时间跨度为 93 天，历史查询时间范围为 366 天；
	// * 86400：（默认值）1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// The Application Name must match the value of the AppName field in the live stream URL. It can include uppercase letters
	// (A-Z), lowercase letters (a-z), numbers (0-9), underscores (_), hyphens (-), and
	// periods (.), with a length ranging from 1 to 30 characters.
	// :::tip When querying stream granularity data, both the App and Stream parameters are required. :::
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，默认为空表示不按维度进行数据拆分，当前接口仅支持填写 Domain 表示按查询的域名为维度进行数据拆分。 :::tip 配置数据拆分的维度时，对应的维度参数传入多个值时才会返回按此维度拆分的数据。例如，配置按 Domain
	// 进行数据拆分时， DomainList 传入多个 Domain 值时，才会返回按 Domain 拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// The Stream Name must correspond to the value of the StreamName field in the live stream URL. It can include uppercase and
	// lowercase letters (A-Z, a-z), numbers (0-9), underscores (_), hyphens (-), and
	// periods (.), with a length ranging from 1 to 100 characters.
	// :::tip When querying stream granularity data, both the App and Stream parameters must be provided. :::
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLivePullToPushDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLivePullToPushDataResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeLivePullToPushDataResResult          `json:"Result,omitempty"`
}

type DescribeLivePullToPushDataResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                              `json:"Version"`
	Error   *DescribeLivePullToPushDataResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeLivePullToPushDataResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeLivePullToPushDataResResult struct {

	// REQUIRED; Data granularity, measured in seconds.
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; The end time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	EndTime string `json:"EndTime"`

	// REQUIRED; 所有时间粒度的数据。
	PullToPushDataList []DescribeLivePullToPushDataResResultPullToPushDataListItem `json:"PullToPushDataList"`

	// REQUIRED; The start time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	StartTime string `json:"StartTime"`

	// REQUIRED; 当前查询条件下的拉流转推总时长，单位为分钟。
	TotalDuration float32 `json:"TotalDuration"`

	// The application name when querying stream granularity data.
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，当前接口仅支持按 Domain 即域名维度进行数据拆分。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 按维度拆分后的数据。
	PullToPushDetailDataList []*DescribeLivePullToPushDataResResultPullToPushDetailDataListItem `json:"PullToPushDetailDataList,omitempty"`

	// The stream name when querying stream granularity data.
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLivePullToPushDataResResultPullToPushDataListItem struct {

	// REQUIRED; The start time of each time granularity when data is aggregated by time granularity, formatted in RFC3339 UTC
	// time with second precision.
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内的拉流转推总时长，单位为分钟。
	Value float32 `json:"Value"`
}

type DescribeLivePullToPushDataResResultPullToPushDetailDataListItem struct {

	// REQUIRED; 按维度进行数据拆分后，当前维度下所有时间粒度的数据。
	PullToPushDataList []DescribeLivePullToPushDataResResultPullToPushDetailDataListPropertiesItemsItem `json:"PullToPushDataList"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的拉流转推总时长，单位分钟。
	TotalDuration float32 `json:"TotalDuration"`

	// 按域名维度进行数据拆分时的域名信息。
	Domain *string `json:"Domain,omitempty"`
}

type DescribeLivePullToPushDataResResultPullToPushDetailDataListPropertiesItemsItem struct {

	// REQUIRED; The start time of each time granularity when data is aggregated by time granularity, formatted in RFC3339 UTC
	// time with second precision.
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 该时间片内的拉流转推总时长，单位分钟，保留小数点后 2 位
	Value float32 `json:"Value"`
}

type DescribeLivePushStreamInfoDataBody struct {

	// REQUIRED; 查询的结束时间。只能查询93d以内的数据
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的起始时间。
	StartTime string `json:"StartTime"`

	// 应用名称
	App *string `json:"App,omitempty"`

	// 域名列表，缺省情况表示该用户的所有域名
	DomainList []*string `json:"DomainList,omitempty"`

	// 分页页数，默认1
	PageNum *int32 `json:"PageNum,omitempty"`

	// 每页大小， 默认20
	PageSize *int32 `json:"PageSize,omitempty"`

	// 直播流名称
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLivePushStreamInfoDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLivePushStreamInfoDataResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeLivePushStreamInfoDataResResult `json:"Result,omitempty"`
}

type DescribeLivePushStreamInfoDataResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

// DescribeLivePushStreamInfoDataResResult - 视请求的接口而定
type DescribeLivePushStreamInfoDataResResult struct {

	// REQUIRED; 结束时间。格式rfc3339
	EndTime string `json:"EndTime"`

	// REQUIRED; 分页信息
	Pagination DescribeLivePushStreamInfoDataResResultPagination `json:"Pagination"`

	// REQUIRED; 按照搜索过滤字段和时间粒度聚合的数据
	PushStreamInfoDataList []DescribeLivePushStreamInfoDataResResultPushStreamInfoDataListItem `json:"PushStreamInfoDataList"`

	// REQUIRED; 起始时间。格式rfc3339
	StartTime string `json:"StartTime"`

	// 应用名称
	App *string `json:"App,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 直播流名称
	Stream *string `json:"Stream,omitempty"`
}

// DescribeLivePushStreamInfoDataResResultPagination - 分页信息
type DescribeLivePushStreamInfoDataResResultPagination struct {

	// REQUIRED; 当前页数
	PageCur int32 `json:"PageCur"`

	// REQUIRED; 每页大小
	PageSize int32 `json:"PageSize"`

	// REQUIRED; 总共推流个数
	TotalCount int32 `json:"TotalCount"`
}

type DescribeLivePushStreamInfoDataResResultPushStreamInfoDataListItem struct {

	// REQUIRED; 应用名称
	App string `json:"App"`

	// REQUIRED; 推流开始到结束的时长，单位s
	Duration int32 `json:"Duration"`

	// REQUIRED; 推流结束时间，格式rfc3339
	EndTime string `json:"EndTime"`

	// REQUIRED; 显示推流客户端的IP地址，如没有IP信息，返回空
	IP string `json:"IP"`

	// REQUIRED; 推流开始时间，格式rfc3339
	StartTime string `json:"StartTime"`

	// REQUIRED; 直播流名称
	Stream string `json:"Stream"`

	// REQUIRED; 推流断开原因
	StreamBreakReason string `json:"StreamBreakReason"`
}

type DescribeLivePushStreamMetricsBody struct {

	// REQUIRED; Application name, which corresponds to the value of theAppNamefield in the live stream address. It can consist
	// of uppercase and lowercase letters (A-Z, a-z), numbers (0-9), underscores (_), hyphens
	// (-), and periods (.), with a length of 1 to 30 characters.
	App string `json:"App"`

	// REQUIRED; 推流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的推流域名。
	Domain string `json:"Domain"`

	// REQUIRED; The end time of your query's range (UTC time in RFC 3339 format with second precision).
	EndTime string `json:"EndTime"`

	// REQUIRED; The start time of your query's range (UTC time in RFC 3339 format with second precision).
	StartTime string `json:"StartTime"`

	// REQUIRED; Stream name, which corresponds to the value of the StreamName field in the live stream address. It can consist
	// of uppercase and lowercase letters (A-Z, a-z), numbers (0-9), underscores (_), hyphens
	// (-), and periods (.), with a length of 1 to 100 characters.
	Stream string `json:"Stream"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 5：5 秒；
	// * 30：（默认值）30 秒。
	Aggregation *int32 `json:"Aggregation,omitempty"`
}

type DescribeLivePushStreamMetricsRes struct {

	// REQUIRED
	ResponseMetadata DescribeLivePushStreamMetricsResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeLivePushStreamMetricsResResult          `json:"Result,omitempty"`
}

type DescribeLivePushStreamMetricsResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                                 `json:"Version"`
	Error   *DescribeLivePushStreamMetricsResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeLivePushStreamMetricsResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeLivePushStreamMetricsResResult struct {

	// Data granularity, measured in seconds.
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 推流域名。
	Domain *string `json:"Domain,omitempty"`

	// The end time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	EndTime *string `json:"EndTime,omitempty"`

	// 所有时间粒度的数据。
	MetricList []*DescribeLivePushStreamMetricsResResultMetricListItem `json:"MetricList,omitempty"`

	// The start time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	StartTime *string `json:"StartTime,omitempty"`

	// 流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLivePushStreamMetricsResResultMetricListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的音频码率最大值，单位为 kbps。
	AudioBitrate float32 `json:"AudioBitrate"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻音频帧显示时间戳差值的最大值，单位为毫秒。
	AudioFrameGap int32 `json:"AudioFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内的音频帧率最大值，单位为 fps。
	AudioFramerate float32 `json:"AudioFramerate"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个音频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	AudioPts int32 `json:"AudioPts"`

	// REQUIRED; 当前数据聚合时间粒度内的视频码率最大值，单位为 kbps。
	Bitrate float32 `json:"Bitrate"`

	// REQUIRED; 当前数据聚合时间粒度内的视频帧率最大值，单位为 fps。
	Framerate float32 `json:"Framerate"`

	// REQUIRED; 当前数据聚合时间粒度内，所有音视频帧显示时间戳差值的最大值，即所有 AudioPts 与 VideoPts 差值的最大值，单位为毫秒。
	PtsDelta int32 `json:"PtsDelta"`

	// REQUIRED; The start time of each time granularity when data is aggregated by time granularity, formatted in RFC3339 UTC
	// time with second precision.
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻视频帧显示时间戳差值的最大值，单位为毫秒。
	VideoFrameGap int32 `json:"VideoFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个视频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	VideoPts int32 `json:"VideoPts"`
}

type DescribeLiveRecordDataBody struct {

	// REQUIRED; The end time of your query's range (UTC time in RFC 3339 format with second precision).
	EndTime string `json:"EndTime"`

	// REQUIRED; The start time of your query's range (UTC time in RFC 3339 format with second precision).
	StartTime string `json:"StartTime"`

	// The granularity of data aggregation, measured in seconds, with the following supported options:
	// * 300 (default): 5 minutes. When aggregated in 5-minute intervals, the maximum time span for a single query is 31 days,
	// and for historical queries, the maximum time range is 366 days.
	// * 3600: 1 hour. When aggregated in 1-hour intervals, the maximum time span for a single query is 93 days, and for historical
	// queries, the maximum time range is 366 days.
	// * 86400: 1 day. When aggregated in 1-day intervals, the maximum time span for a single query is 93 days, and for historical
	// queries, the maximum time range is 366 days.
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// The Application Name must match the value of the AppName field in the live stream URL. It can include uppercase letters
	// (A-Z), lowercase letters (a-z), numbers (0-9), underscores (_), hyphens (-), and
	// periods (.), with a length ranging from 1 to 30 characters.
	// :::tip When querying stream granularity data, both the App and Stream parameters are required. :::
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，默认为空表示不按维度进行数据拆分，当前接口仅支持填写 Domain 表示按查询的域名为维度进行数据拆分。 :::tip 配置数据拆分的维度时，对应的维度参数传入多个值时才会返回按此维度拆分的数据。例如，配置按 Domain
	// 进行数据拆分时， DomainList 传入多个 Domain 值时，才会返回按 Domain 拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，默认为空，表示查询您视频直播产品下所有域名的录制用量数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// The Stream Name must correspond to the value of the StreamName field in the live stream URL. It can include uppercase and
	// lowercase letters (A-Z, a-z), numbers (0-9), underscores (_), hyphens (-), and
	// periods (.), with a length ranging from 1 to 100 characters.
	// :::tip When querying stream granularity data, both the App and Stream parameters must be provided. :::
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveRecordDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveRecordDataResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeLiveRecordDataResResult          `json:"Result,omitempty"`
}

type DescribeLiveRecordDataResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                          `json:"Version"`
	Error   *DescribeLiveRecordDataResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeLiveRecordDataResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeLiveRecordDataResResult struct {

	// REQUIRED; Data granularity, measured in seconds.
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; The end time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	EndTime string `json:"EndTime"`

	// REQUIRED; 所有时间粒度的数据。
	RecordDataList []DescribeLiveRecordDataResResultRecordDataListItem `json:"RecordDataList"`

	// REQUIRED; 当前查询条件下的录制并发路数最大值。
	RecordPeak int32 `json:"RecordPeak"`

	// REQUIRED; The start time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	StartTime string `json:"StartTime"`

	// The application name when querying stream granularity data.
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，当前接口仅支持按 Domain 即域名维度进行数据拆分。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 按维度拆分后的数据。
	RecordDetailDataList []*DescribeLiveRecordDataResResultRecordDetailDataListItem `json:"RecordDetailDataList,omitempty"`

	// The stream name when querying stream granularity data.
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveRecordDataResResultRecordDataListItem struct {

	// REQUIRED; The start time of each time granularity when data is aggregated by time granularity, formatted in RFC3339 UTC
	// time with second precision.
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内录制并发路数最大值。
	Value int32 `json:"Value"`
}

type DescribeLiveRecordDataResResultRecordDetailDataListItem struct {

	// REQUIRED; 按维度进行数据拆分后，当前维度下所有时间粒度的数据。
	RecordDataList []DescribeLiveRecordDataResResultRecordDetailDataListPropertiesItemsItem `json:"RecordDataList"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的录制并发路数最大值。
	RecordPeak int32 `json:"RecordPeak"`

	// 按域名维度进行数据拆分时的域名信息。
	Domain *string `json:"Domain,omitempty"`
}

type DescribeLiveRecordDataResResultRecordDetailDataListPropertiesItemsItem struct {

	// REQUIRED; The start time of each time granularity when data is aggregated by time granularity, formatted in RFC3339 UTC
	// time with second precision.
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 录制峰值
	Value int32 `json:"Value"`
}

type DescribeLiveRegionDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveRegionDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveRegionDataResResult `json:"Result"`
}

type DescribeLiveRegionDataResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`
}

type DescribeLiveRegionDataResResult struct {

	// REQUIRED; 区域信息。
	Areas []DescribeLiveRegionDataResResultAreasItem `json:"Areas"`
}

type DescribeLiveRegionDataResResultAreasItem struct {

	// REQUIRED; 大区标识符。
	Code string `json:"Code"`

	// REQUIRED; 国家信息。
	Countries []DescribeLiveRegionDataResResultAreasPropertiesItemsItem `json:"Countries"`

	// REQUIRED; 大区名称。
	Name string `json:"Name"`
}

type DescribeLiveRegionDataResResultAreasPropertiesItemsItem struct {

	// REQUIRED; 国家标识符。
	Code string `json:"Code"`

	// REQUIRED; 国家名称。
	Name string `json:"Name"`

	// REQUIRED; 省份信息，国外暂不支持该参数。
	Provinces []DescribeLiveRegionDataResResultAreasPropertiesItemsProvincesItem `json:"Provinces"`
}

type DescribeLiveRegionDataResResultAreasPropertiesItemsProvincesItem struct {

	// REQUIRED; 省份标识符。
	Code string `json:"Code"`

	// REQUIRED; 省份名称。
	Name string `json:"Name"`
}

type DescribeLiveSnapshotDataBody struct {

	// REQUIRED; The end time of your query's range (UTC time in RFC 3339 format with second precision).
	EndTime string `json:"EndTime"`

	// REQUIRED; The start time of your query's range (UTC time in RFC 3339 format with second precision).
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 300：5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天；
	// * 86400：（默认值）1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// The Application Name must match the value of the AppName field in the live stream URL. It can include uppercase letters
	// (A-Z), lowercase letters (a-z), numbers (0-9), underscores (_), hyphens (-), and
	// periods (.), with a length ranging from 1 to 30 characters.
	// :::tip When querying stream granularity data, both the App and Stream parameters are required. :::
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，默认为空表示不按维度进行数据拆分，当前接口仅支持填写 Domain 表示按查询的域名为维度进行数据拆分。
	// :::tip 配置数据拆分的维度时，对应的维度参数传入多个值时才会返回按此维度拆分的数据。例如，配置按 Domain 进行数据拆分时， DomainList 传入多个 Domain 值时，才会返回按 Domain 拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，默认为空，表示查询您视频直播产品下所有域名的截图用量数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// The Stream Name must correspond to the value of the StreamName field in the live stream URL. It can include uppercase and
	// lowercase letters (A-Z, a-z), numbers (0-9), underscores (_), hyphens (-), and
	// periods (.), with a length ranging from 1 to 100 characters.
	// :::tip When querying stream granularity data, both the App and Stream parameters must be provided. :::
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveSnapshotDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveSnapshotDataResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeLiveSnapshotDataResResult          `json:"Result,omitempty"`
}

type DescribeLiveSnapshotDataResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                            `json:"Version"`
	Error   *DescribeLiveSnapshotDataResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeLiveSnapshotDataResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeLiveSnapshotDataResResult struct {

	// REQUIRED; Data granularity, measured in seconds.
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; The end time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	EndTime string `json:"EndTime"`

	// REQUIRED; 所有时间粒度的数据。
	SnapshotDataList []DescribeLiveSnapshotDataResResultSnapshotDataListItem `json:"SnapshotDataList"`

	// REQUIRED; The start time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	StartTime string `json:"StartTime"`

	// REQUIRED; 当前查询条件下的截图总张数。
	Total int32 `json:"Total"`

	// The application name when querying stream granularity data.
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，当前接口仅支持按 Domain 即域名维度进行数据拆分。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 按维度拆分后的数据。
	SnapshotDetailData []*DescribeLiveSnapshotDataResResultSnapshotDetailDataItem `json:"SnapshotDetailData,omitempty"`

	// The stream name when querying stream granularity data.
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveSnapshotDataResResultSnapshotDataListItem struct {

	// REQUIRED; The start time of each time granularity when data is aggregated by time granularity, formatted in RFC3339 UTC
	// time with second precision.
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内的截图总张数。
	Value int32 `json:"Value"`
}

type DescribeLiveSnapshotDataResResultSnapshotDetailDataItem struct {

	// REQUIRED; 按域名维度进行数据拆分时的域名信息。
	Domain string `json:"Domain"`

	// REQUIRED; 按维度进行数据拆分后，当前维度下所有时间粒度的数据。
	SnapshotDataList []DescribeLiveSnapshotDataResResultSnapshotDetailDataPropertiesItemsItem `json:"SnapshotDataList"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的截图总张数。
	Total int32 `json:"Total"`
}

type DescribeLiveSnapshotDataResResultSnapshotDetailDataPropertiesItemsItem struct {

	// REQUIRED; The start time of each time granularity when data is aggregated by time granularity, formatted in RFC3339 UTC
	// time with second precision.
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 截图总张数
	Value int32 `json:"Value"`
}

type DescribeLiveSourceBandwidthDataBody struct {

	// REQUIRED; The end time of your query's range (UTC time in RFC 3339 format with second precision).
	EndTime string `json:"EndTime"`

	// REQUIRED; The start time of your query's range (UTC time in RFC 3339 format with second precision).
	StartTime string `json:"StartTime"`

	// The granularity of data aggregation, measured in seconds, with the following supported options:
	// * 60: 1 minute. When aggregated every 1 minute, the maximum time span for a single query is 24 hours, and the historical
	// query time range is 366 days;
	// * 300: (default) 5 minutes. When aggregated every 5 minutes, the maximum time span for a single query is 31 days, and the
	// historical query time range is 366 days;
	// * 3600: 1 hour. When aggregated every 1 hour, the maximum time span for a single query is 93 days, and the historical query
	// time range is 366 days.
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// The Application Name must match the value of the AppName field in the live stream URL. It can include uppercase letters
	// (A-Z), lowercase letters (a-z), numbers (0-9), underscores (_), hyphens (-), and
	// periods (.), with a length ranging from 1 to 30 characters.
	// :::tip When querying data at the stream granularity, you must specify the Domain, App, and Stream parameters simultaneously.
	// :::
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，默认为空表示不按维度进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * ISP：运营商。
	// :::tip 配置数据拆分的维度时，对应的维度参数传入多个值时才会返回按此维度拆分的数据。例如，配置按 Domain 进行数据拆分时， DomainList 传入多个 Domain 值时，才会返回按 Domain 拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 拉流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的拉流域名。
	// :::tip 查询流粒度的回源带宽监控数据时，需同时指定 Domain 、App
	// 和 Stream 来指定回源流。 :::
	Domain *string `json:"Domain,omitempty"`

	// 拉流域名列表，默认为空，表示查询所有域名的回源带宽监控数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的拉流域名。
	// :::tipDomainList
	// 和 Domain 传且仅传一个。 :::
	DomainList []*string `json:"DomainList,omitempty"`

	// Identifiers of operators providing network access services. By default, all operators are indicated. Supported operators
	// are as follows.
	// * unicom: China Unicom;
	// * railcom: China Railway Telecom;
	// * telecom: China Telecom;
	// * mobile: China Mobile;
	// * cernet: China Education and Research Network (CERNET);
	// * tianwei: China Tianwei;
	// * alibaba: Alibaba Group;
	// * tencent: Tencent Holdings;
	// * drpeng: Dr. Peng Telecom & Media Group;
	// * btvn: China Broadcasting Network;
	// * huashu: Huashu Media;
	// * other: Denotes other/unspecified options.
	// If you need to obtain the identifiers of various operators, you can call the DescribeLiveISPData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveispdata].
	ISPList []*string `json:"ISPList,omitempty"`

	// The Stream Name is required when querying stream granularity data. It supports uppercase and lowercase letters (A-Z, a-z),
	// underscores (_), hyphens (-), and periods (.), with a length of 1 to 100
	// characters.
	// :::tip When querying stream granularity data, you must specify the Domain, App, and Stream parameters simultaneously. :::
	Stream *string `json:"Stream,omitempty"`

	// List of regions to which client IPs belong, by default indicating all regions.
	UserRegionList []*DescribeLiveSourceBandwidthDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveSourceBandwidthDataBodyUserRegionListItem struct {

	// The identifier for the major region in the regional information can be obtained by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata].
	Area *string `json:"Area,omitempty"`

	// The country identifier in the regional information can be obtained by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata].
	// When filtering by
	// country, both 'Area' and 'Country' need to be passed in simultaneously.
	Country *string `json:"Country,omitempty"`

	// The province identifier in the regional information is currently not supported for countries or regions outside mainland
	// China, Hong Kong, Macao, and Taiwan. You can obtain the identifier information
	// by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata]. When
	// filtering by province, you need to simultaneously pass in Area, Country, and
	// Province
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveSourceBandwidthDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveSourceBandwidthDataResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeLiveSourceBandwidthDataResResult          `json:"Result,omitempty"`
}

type DescribeLiveSourceBandwidthDataResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                                   `json:"Version"`
	Error   *DescribeLiveSourceBandwidthDataResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeLiveSourceBandwidthDataResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeLiveSourceBandwidthDataResResult struct {

	// REQUIRED; 所有时间粒度的数据。
	BandwidthDataList []DescribeLiveSourceBandwidthDataResResultBandwidthDataListItem `json:"BandwidthDataList"`

	// Data granularity, measured in seconds.
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// The application name when querying stream granularity data.
	App *string `json:"App,omitempty"`

	// 按维度拆分后的数据。
	BandwidthDetailDataList []*DescribeLiveSourceBandwidthDataResResultBandwidthDetailDataListItem `json:"BandwidthDetailDataList,omitempty"`

	// 数据拆分的维度，维度说明如下所示。
	// * Domain：域名；
	// * ISP：运营商。
	DetailField []*string `json:"DetailField,omitempty"`

	// 查询流粒度数据时的域名。
	Domain *string `json:"Domain,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// The end time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	EndTime *string `json:"EndTime,omitempty"`

	// Identifiers of operators providing network access services. By default, all operators are indicated. Supported operators
	// are as follows.
	// * unicom: China Unicom;
	// * railcom: China Railway Telecom;
	// * telecom: China Telecom;
	// * mobile: China Mobile;
	// * cernet: China Broadcasting Network;
	// * tianwei: China Tianwei;
	// * alibaba: Alibaba Group;
	// * tencent: Tencent Holdings;
	// * drpeng: Dr. Peng Telecom & Media Group;
	// * btvn: Broadcasting Television Network (BTVN);
	// * huashu: Huashu Media;
	// * other: Denotes other/unspecified options.
	ISPList []*string `json:"ISPList,omitempty"`

	// 查询时间范围内的回源峰值带宽，单位为 Mbps。
	PeakBandwidth *float32 `json:"PeakBandwidth,omitempty"`

	// The start time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	StartTime *string `json:"StartTime,omitempty"`

	// The stream name when querying stream granularity data.
	Stream *string `json:"Stream,omitempty"`

	// List of regions to which client IPs belong.
	UserRegionList []*DescribeLiveSourceBandwidthDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveSourceBandwidthDataResResultBandwidthDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的回源峰值带宽，单位为 Mbps。
	Bandwidth float32 `json:"Bandwidth"`

	// REQUIRED; The start time of each time granularity when data is aggregated by time granularity, formatted in RFC3339 UTC
	// time with second precision.
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLiveSourceBandwidthDataResResultBandwidthDetailDataListItem struct {

	// 按维度进行数据拆分后，当前维度下所有时间粒度的数据。
	BandwidthDataList []*DescribeLiveSourceBandwidthDataResResultBandwidthDetailDataListPropertiesItemsItem `json:"BandwidthDataList,omitempty"`

	// 按域名维度进行数据拆分时的域名信息。
	Domain *string `json:"Domain,omitempty"`

	// 按运营商维度进行数据拆分时的运营商信息。
	ISP *string `json:"ISP,omitempty"`

	// 按维度进行数据拆分后，当前维度的回源峰值带宽，单位为 Mbps。
	PeakBandwidth *float32 `json:"PeakBandwidth,omitempty"`
}

type DescribeLiveSourceBandwidthDataResResultBandwidthDetailDataListPropertiesItemsItem struct {

	// 时间片内回源带宽峰值，单位 Mbps
	Bandwidth *float32 `json:"Bandwidth,omitempty"`

	// The start time of each time granularity when data is aggregated by time granularity, formatted in RFC3339 UTC time with
	// second precision.
	TimeStamp *string `json:"TimeStamp,omitempty"`
}

type DescribeLiveSourceBandwidthDataResResultUserRegionListItem struct {

	// The regional identifier in regional information.
	Area *string `json:"Area,omitempty"`

	// The country identifier in the regional information.
	Country *string `json:"Country,omitempty"`

	// The province identifier in the regional information.
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveSourceStreamMetricsBody struct {

	// REQUIRED; Application name, which corresponds to the value of theAppNamefield in the live stream address. It can consist
	// of uppercase and lowercase letters (A-Z, a-z), numbers (0-9), underscores (_), hyphens
	// (-), and periods (.), with a length of 1 to 30 characters.
	App string `json:"App"`

	// REQUIRED; 拉流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看回源流使用的拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; The end time of your query's range (UTC time in RFC 3339 format with second precision).
	EndTime string `json:"EndTime"`

	// REQUIRED; The start time of your query's range (UTC time in RFC 3339 format with second precision).
	StartTime string `json:"StartTime"`

	// REQUIRED; Stream name, which corresponds to the value of the StreamName field in the live stream address. It can consist
	// of uppercase and lowercase letters (A-Z, a-z), numbers (0-9), underscores (_), hyphens
	// (-), and periods (.), with a length of 1 to 100 characters.
	Stream string `json:"Stream"`

	// 数据聚合的时间粒度，单位为秒，当前接口默认且仅支持按 30 秒进行数据聚合。
	Aggregation *int32 `json:"Aggregation,omitempty"`
}

type DescribeLiveSourceStreamMetricsRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveSourceStreamMetricsResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveSourceStreamMetricsResResult `json:"Result"`
}

type DescribeLiveSourceStreamMetricsResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`
}

type DescribeLiveSourceStreamMetricsResResult struct {

	// REQUIRED; Data granularity, measured in seconds.
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; The end time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	EndTime string `json:"EndTime"`

	// REQUIRED; 所有时间粒度的数据。
	MetricList []DescribeLiveSourceStreamMetricsResResultMetricListItem `json:"MetricList"`

	// REQUIRED; The start time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	StartTime string `json:"StartTime"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`
}

type DescribeLiveSourceStreamMetricsResResultMetricListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的音频码率最大值，单位为 kbps。
	AudioBitrate float32 `json:"AudioBitrate"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻音频帧显示时间戳差值的最大值，单位为毫秒。
	AudioFrameGap int32 `json:"AudioFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内的音频帧率最大值，单位为 fps。
	AudioFramerate float32 `json:"AudioFramerate"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个音频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	AudioPts int32 `json:"AudioPts"`

	// REQUIRED; 当前数据聚合时间粒度内的视频码率最大值，单位为 kbps。
	Bitrate float32 `json:"Bitrate"`

	// REQUIRED; 当前数据聚合时间粒度内的视频帧率最大值，单位为 fps
	Framerate float32 `json:"Framerate"`

	// REQUIRED; 当前数据聚合时间粒度内，所有音视频帧显示时间戳差值的最大值，即所有 AudioPts 与 VideoPts 差值的最大值，单位为毫秒。
	PtsDelta int32 `json:"PtsDelta"`

	// REQUIRED; The start time of each time granularity when data is aggregated by time granularity, formatted in RFC3339 UTC
	// time with second precision.
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻视频帧显示时间戳差值的最大值，单位为毫秒。
	VideoFrameGap int32 `json:"VideoFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个视频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	VideoPts int32 `json:"VideoPts"`
}

type DescribeLiveSourceTrafficDataBody struct {

	// REQUIRED; The end time of your query's range (UTC time in RFC 3339 format with second precision).
	EndTime string `json:"EndTime"`

	// REQUIRED; The start time of your query's range (UTC time in RFC 3339 format with second precision).
	StartTime string `json:"StartTime"`

	// The granularity of data aggregation, measured in seconds, with the following supported options:
	// * 60: 1 minute. When aggregated every 1 minute, the maximum time span for a single query is 24 hours, and the historical
	// query time range is 366 days;
	// * 300: (default) 5 minutes. When aggregated every 5 minutes, the maximum time span for a single query is 31 days, and the
	// historical query time range is 366 days;
	// * 3600: 1 hour. When aggregated every 1 hour, the maximum time span for a single query is 93 days, and the historical query
	// time range is 366 days.
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// The Application Name must match the value of the AppName field in the live stream URL. It can include uppercase letters
	// (A-Z), lowercase letters (a-z), numbers (0-9), underscores (_), hyphens (-), and
	// periods (.), with a length ranging from 1 to 30 characters.
	// :::tip When querying data at the stream granularity, you must specify the Domain, App, and Stream parameters simultaneously.
	// :::
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，默认为空表示按维度进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * ISP：运营商。
	// :::tip 配置数据拆分的维度时，对应的维度参数传入多个值时才会返回按此维度拆分的数据。例如，配置按 Domain 进行数据拆分时， DomainList 传入多个 Domain 值时，才会返回按 Domain 拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 拉流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的拉流域名。
	// :::tip 查询流粒度的回源流量监控数据时，需同时指定 Domain 、App
	// 和 Stream 来指定回源流。 :::
	Domain *string `json:"Domain,omitempty"`

	// 拉流域名列表，默认为空，表示查询所有域名的回源流量监控数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的拉流域名。
	// :::tipDomainList 和 Domain 传且仅传一个。 :::
	DomainList []*string `json:"DomainList,omitempty"`

	// Identifiers of operators providing network access services. By default, all operators are indicated. Supported operators
	// are as follows.
	// * unicom: China Unicom;
	// * railcom: China Railway Telecom;
	// * telecom: China Telecom;
	// * mobile: China Mobile;
	// * cernet: China Education and Research Network (CERNET);
	// * tianwei: China Tianwei;
	// * alibaba: Alibaba Group;
	// * tencent: Tencent Holdings;
	// * drpeng: Dr. Peng Telecom & Media Group;
	// * btvn: China Broadcasting Network;
	// * huashu: Huashu Media;
	// * other: Denotes other/unspecified options.
	// If you need to obtain the identifiers of various operators, you can call the DescribeLiveISPData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveispdata].
	ISPList []*string `json:"ISPList,omitempty"`

	// The Stream Name is required when querying stream granularity data. It supports uppercase and lowercase letters (A-Z, a-z),
	// underscores (_), hyphens (-), and periods (.), with a length of 1 to 100
	// characters.
	// :::tip When querying stream granularity data, you must specify the Domain, App, and Stream parameters simultaneously. :::
	Stream *string `json:"Stream,omitempty"`

	// List of regions to which client IPs belong, by default indicating all regions.
	UserRegionList []*DescribeLiveSourceTrafficDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveSourceTrafficDataBodyUserRegionListItem struct {

	// The identifier for the major region in the regional information can be obtained by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata].
	Area *string `json:"Area,omitempty"`

	// The country identifier in the regional information can be obtained by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata].
	// When filtering by
	// country, both 'Area' and 'Country' need to be passed in simultaneously.
	Country *string `json:"Country,omitempty"`

	// The province identifier in the regional information is currently not supported for countries or regions outside mainland
	// China, Hong Kong, Macao, and Taiwan. You can obtain the identifier information
	// by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata]. When
	// filtering by province, you need to simultaneously pass in Area, Country, and
	// Province
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveSourceTrafficDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveSourceTrafficDataResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeLiveSourceTrafficDataResResult          `json:"Result,omitempty"`
}

type DescribeLiveSourceTrafficDataResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                                 `json:"Version"`
	Error   *DescribeLiveSourceTrafficDataResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeLiveSourceTrafficDataResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeLiveSourceTrafficDataResResult struct {

	// REQUIRED; 查询时间范围内的回源总流量，单位为 GB。
	TotalTraffic float32 `json:"TotalTraffic"`

	// REQUIRED; 所有时间粒度的数据。
	TrafficDataList []DescribeLiveSourceTrafficDataResResultTrafficDataListItem `json:"TrafficDataList"`

	// Data granularity, measured in seconds.
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// The application name when querying stream granularity data.
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，维度说明如下所示。
	// * Domain：域名；
	// * ISP：运营商。
	DetailField []*string `json:"DetailField,omitempty"`

	// 查询流粒度数据时的域名。
	Domain *string `json:"Domain,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// The end time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	EndTime *string `json:"EndTime,omitempty"`

	// Identifiers of operators providing network access services. By default, all operators are indicated. Supported operators
	// are as follows.
	// * unicom: China Unicom;
	// * railcom: China Railway Telecom;
	// * telecom: China Telecom;
	// * mobile: China Mobile;
	// * cernet: China Broadcasting Network;
	// * tianwei: China Tianwei;
	// * alibaba: Alibaba Group;
	// * tencent: Tencent Holdings;
	// * drpeng: Dr. Peng Telecom & Media Group;
	// * btvn: Broadcasting Television Network (BTVN);
	// * huashu: Huashu Media;
	// * other: Denotes other/unspecified options.
	ISPList []*string `json:"ISPList,omitempty"`

	// The start time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	StartTime *string `json:"StartTime,omitempty"`

	// The stream name when querying stream granularity data.
	Stream *string `json:"Stream,omitempty"`

	// 按维度拆分后的数据。
	TrafficDetailDataList []*DescribeLiveSourceTrafficDataResResultTrafficDetailDataListItem `json:"TrafficDetailDataList,omitempty"`

	// List of regions to which client IPs belong.
	UserRegionList []*DescribeLiveSourceTrafficDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveSourceTrafficDataResResultTrafficDataListItem struct {

	// REQUIRED; The start time of each time granularity when data is aggregated by time granularity, formatted in RFC3339 UTC
	// time with second precision.
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内产生的回源流量，单位 GB。
	Traffic float32 `json:"Traffic"`
}

type DescribeLiveSourceTrafficDataResResultTrafficDetailDataListItem struct {

	// REQUIRED; 按维度进行数据拆分后，当前维度的回源总流量，单位为 GB。
	TotalTraffic float32 `json:"TotalTraffic"`

	// REQUIRED; 按维度进行数据拆分后，当前维度下所有时间粒度的数据。
	TrafficDataList []DescribeLiveSourceTrafficDataResResultTrafficDetailDataListPropertiesItemsItem `json:"TrafficDataList"`

	// 按域名维度进行数据拆分时的域名信息。
	Domain *string `json:"Domain,omitempty"`

	// 按运营商维度进行数据拆分时的运营商信息。
	ISP *string `json:"ISP,omitempty"`
}

type DescribeLiveSourceTrafficDataResResultTrafficDetailDataListPropertiesItemsItem struct {

	// REQUIRED; The start time of each time granularity when data is aggregated by time granularity, formatted in RFC3339 UTC
	// time with second precision.
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 回源流量，单位 GB
	Traffic float32 `json:"Traffic"`
}

type DescribeLiveSourceTrafficDataResResultUserRegionListItem struct {

	// The regional identifier in regional information.
	Area *string `json:"Area,omitempty"`

	// The country identifier in the regional information.
	Country *string `json:"Country,omitempty"`

	// The province identifier in the regional information.
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveStorageSpaceDataBody struct {

	// REQUIRED; The start time of your query's range (UTC time in RFC 3339 format with second precision).
	EndTime string `json:"EndTime"`

	// REQUIRED; The end time of your query's range (UTC time in RFC 3339 format with second precision).
	StartTime string `json:"StartTime"`

	// 时间粒度，单位为 s，支持配置为 86400，单次查询时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 桶名称
	Buckets []*string `json:"Buckets,omitempty"`
}

type DescribeLiveStorageSpaceDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveStorageSpaceDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveStorageSpaceDataResResult `json:"Result"`
}

type DescribeLiveStorageSpaceDataResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`
}

type DescribeLiveStorageSpaceDataResResult struct {

	// REQUIRED; Data granularity, measured in seconds.
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; The end time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	EndTime string `json:"EndTime"`

	// REQUIRED; The start time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	StartTime string `json:"StartTime"`

	// REQUIRED; 时移存储数据。
	StorageDataList []DescribeLiveStorageSpaceDataResResultStorageDataListItem `json:"StorageDataList"`

	// 域名空间列表。
	Buckets []*string `json:"Buckets,omitempty"`
}

type DescribeLiveStorageSpaceDataResResultStorageDataListItem struct {

	// REQUIRED; 存储用量，单位为 GB。
	Storage float32 `json:"Storage"`

	// REQUIRED; 时间点，默认为每日的结束时间。例如，返回 2022-02-16T00:00:00+08:00，表示取该时刻的存储用量作为 2 月 16 日的计费用量。
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLiveStreamInfoByPageQuery struct {

	// REQUIRED; 查询数据的页码，取值为正整数。
	PageNum int32 `json:"PageNum" query:"PageNum"`

	// REQUIRED; 每页显示的数据条数，取值范围为 [1,1000]。
	PageSize int32 `json:"PageSize" query:"PageSize"`

	// 应用名称，取值与直播流地址中 AppName 字段取值相同，默认为空，表示查询所有应用名称。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App *string `json:"App,omitempty" query:"App"`

	// 直播流使用的域名，默认为空，表示查询所有当前域名空间下的在线流。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名。
	Domain *string `json:"Domain,omitempty" query:"Domain"`

	// 想要查询的目标信息，使用英文逗号作为分隔符“,”，例如，bitrate,framerate。缺省情况下表示 bitrate,framerate。支持如下取值。 all：所有信息；onlineuser：在线人数；bandwidth：带宽信息;bitrate：码率信息；framerate：帧率信息；.
	InfoType *string `json:"InfoType,omitempty" query:"InfoType"`

	// 使用流名称进行查询的方式，默认值为 strict，支持的取值即含义如下所示。
	// * fuzzy：模糊匹配；
	// * strict：精准匹配。
	QueryType *string `json:"QueryType,omitempty" query:"QueryType"`

	// 在线流的来源类型，默认为空，表示查询所有来源类型，支持的取值即含义如下所示。
	// * push：直推流；
	// * relay：回源流。
	SourceType *string `json:"SourceType,omitempty" query:"SourceType"`

	// 流名称，取值与直播流地址中 StreamName 字段取值相同，默认为空表示查询所有流名称。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	Stream *string `json:"Stream,omitempty" query:"Stream"`

	// 在线流的流类型，默认为空，表示查询所有类型，支持的取值即含义如下所示。
	// * origin：原始流；
	// * trans：转码流。
	StreamType *string `json:"StreamType,omitempty" query:"StreamType"`

	// 域名空间，即直播流地址的域名所属的域名空间，默认为空，表示查询所有域名空间下的在线流。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]
	// 页面，查看需要查询的直播流使用的域名所属的域名空间。
	Vhost *string `json:"Vhost,omitempty" query:"Vhost"`
}

type DescribeLiveStreamInfoByPageRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveStreamInfoByPageResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeLiveStreamInfoByPageResResult          `json:"Result,omitempty"`
}

type DescribeLiveStreamInfoByPageResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                                `json:"Version"`
	Error   *DescribeLiveStreamInfoByPageResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeLiveStreamInfoByPageResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeLiveStreamInfoByPageResResult struct {

	// REQUIRED; 查询结果中在线流的数量。
	RoughCount int32 `json:"RoughCount"`

	// 在线流信息列表。
	StreamInfoList []*DescribeLiveStreamInfoByPageResResultStreamInfoListItem `json:"StreamInfoList,omitempty"`
}

type DescribeLiveStreamInfoByPageResResultStreamInfoListItem struct {

	// REQUIRED; 在线流使用的应用名称。
	App string `json:"App"`

	// REQUIRED; 带宽
	BandWidth string `json:"BandWidth"`

	// REQUIRED; 码率
	Bitrate string `json:"Bitrate"`

	// REQUIRED; 在线流使用的域名。
	Domain string `json:"Domain"`

	// REQUIRED; 帧率
	Framerate string `json:"Framerate"`

	// REQUIRED; 在线流的 ID。
	ID int64 `json:"ID"`

	// REQUIRED; 在线人数
	OnlineUser string `json:"OnlineUser"`

	// REQUIRED; 预览地址
	PreviewURL string `json:"PreviewURL"`

	// REQUIRED; 在线流的开始时间，RFC3339 格式的 UTC 时间戳，精度为毫秒。
	SessionStartTime string `json:"SessionStartTime"`

	// REQUIRED; 在线流的来源类型，取值及含义如下所示。
	// * push：直推流；
	// * relay：回源流。
	SourceType string `json:"SourceType"`

	// REQUIRED; 在线流使用的流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 在线流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`
}

type DescribeLiveStreamSessionDataBody struct {

	// REQUIRED; The end time of your query's range (UTC time in RFC 3339 format with second precision).
	EndTime string `json:"EndTime"`

	// REQUIRED; The start time of your query's range (UTC time in RFC 3339 format with second precision).
	StartTime string `json:"StartTime"`

	// The granularity of data aggregation, measured in seconds, with the following supported options:
	// * 60: 1 minute. When aggregated every 1 minute, the maximum time span for a single query is 24 hours, and the historical
	// query time range is 366 days;
	// * 300: (default) 5 minutes. When aggregated every 5 minutes, the maximum time span for a single query is 31 days, and the
	// historical query time range is 366 days;
	// * 3600: 1 hour. When aggregated every 1 hour, the maximum time span for a single query is 93 days, and the historical query
	// time range is 366 days.
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// The Application Name must match the value of the AppName field in the live stream URL. It can include uppercase letters
	// (A-Z), lowercase letters (a-z), numbers (0-9), underscores (_), hyphens (-), and
	// periods (.), with a length ranging from 1 to 30 characters.
	// :::tip When querying data at the stream granularity, you must specify the Domain, App, and Stream parameters simultaneously.
	// :::
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，默认为空表示不按维度进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * ISP：运营商；
	// * Protocol：推拉流协议；
	// * Referer：请求的 Referer 信息。
	// :::tip 配置数据拆分的维度时，对应的维度参数传入多个值时才会返回按此维度拆分的数据。例如，配置按 Domain 进行数据拆分时， DomainList 传入多个 Domain 值时，才会返回按 Domain 拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 拉流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的拉流域名。
	// :::tip 查询流粒度的请求数和在线人数数据时，需同时指定 Domain 、
	// App 和 Stream 来指定直播流。 :::
	Domain *string `json:"Domain,omitempty"`

	// 拉流域名列表，默认为空，表示查询所有域名的请求数和在线人数。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的拉流域名。
	// :::tipDomainList 和 Domain 传且仅传一个。 :::
	DomainList []*string `json:"DomainList,omitempty"`

	// Identifiers of operators providing network access services. By default, all operators are indicated. Supported operators
	// are as follows.
	// * unicom: China Unicom;
	// * railcom: China Railway Telecom;
	// * telecom: China Telecom;
	// * mobile: China Mobile;
	// * cernet: China Education and Research Network (CERNET);
	// * tianwei: China Tianwei;
	// * alibaba: Alibaba Group;
	// * tencent: Tencent Holdings;
	// * drpeng: Dr. Peng Telecom & Media Group;
	// * btvn: China Broadcasting Network;
	// * huashu: Huashu Media;
	// * other: Denotes other/unspecified options.
	// If you need to obtain the identifiers of various operators, you can call the DescribeLiveISPData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveispdata].
	ISPList []*string `json:"ISPList,omitempty"`

	// Push-pull streaming protocol. By default, all protocol types are indicated. Supported protocols are as follows.
	// * HTTP-FLV: A push-pull streaming protocol based on the HTTP protocol, using the FLV format for video transmission.
	// * HTTP-HLS: A push-pull streaming protocol based on the HTTP protocol, using the TS format for video transmission.
	// * RTMP: Real-Time Messaging Protocol for real-time message transmission.
	// * RTM: Real-Time Media protocol for ultra-low latency live streaming.
	// * SRT: Secure Reliable Transport protocol for secure and reliable streaming.
	// * QUIC: Quick UDP Internet Connections, a new low-latency internet transmission protocol based on UDP.
	// * CMAF: Common Media Application Format, a versatile streaming protocol.
	// :::tip If querying the QUIC protocol, other protocols cannot be queried simultaneously. :::
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// 指定拉流请求的 Referer 信息，默认为空，表示不对拉流请求的 Referer 字段进行校验。
	RefererList []*string `json:"RefererList,omitempty"`

	// List of regions to which CDN node IPs belong, by default indicating all regions.
	RegionList []*DescribeLiveStreamSessionDataBodyRegionListItem `json:"RegionList,omitempty"`

	// The Stream Name is required when querying stream granularity data. It supports uppercase and lowercase letters (A-Z, a-z),
	// underscores (_), hyphens (-), and periods (.), with a length of 1 to 100
	// characters.
	// :::tip When querying stream granularity data, you must specify the Domain, App, and Stream parameters simultaneously. :::
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveStreamSessionDataBodyRegionListItem struct {

	// The identifier for the major region in the regional information can be obtained by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata].
	Area *string `json:"Area,omitempty"`

	// The country identifier in the regional information can be obtained by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata].
	// When filtering by
	// country, both 'Area' and 'Country' need to be passed in simultaneously.
	Country *string `json:"Country,omitempty"`

	// The province identifier in the regional information is currently not supported for countries or regions outside mainland
	// China, Hong Kong, Macao, and Taiwan. You can obtain the identifier information
	// by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata]. When
	// filtering by province, you need to simultaneously pass in Area, Country, and
	// Province
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveStreamSessionDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveStreamSessionDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveStreamSessionDataResResult `json:"Result"`
}

type DescribeLiveStreamSessionDataResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`
}

type DescribeLiveStreamSessionDataResResult struct {

	// REQUIRED; Data granularity, measured in seconds.
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; The end time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询时间范围内的在线人数峰值。
	PeakOnlineUser int32 `json:"PeakOnlineUser"`

	// REQUIRED; 所有时间粒度的数据。
	SessionDataList []DescribeLiveStreamSessionDataResResultSessionDataListItem `json:"SessionDataList"`

	// REQUIRED; The start time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	StartTime string `json:"StartTime"`

	// REQUIRED; 查询时间范围内的请求数。
	TotalRequest int32 `json:"TotalRequest"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，维度说明如下所示。
	// * Domain：域名；
	// * ISP：运营商；
	// * Protocol：推拉流协议；
	// * Referer：请求的 Referer 信息。
	DetailField []*string `json:"DetailField,omitempty"`

	// 拉流域名。
	Domain *string `json:"Domain,omitempty"`

	// List of domains.
	DomainList []*string `json:"DomainList,omitempty"`

	// Identifiers of operators providing network access services. By default, all operators are indicated. Supported operators
	// are as follows.
	// * unicom: China Unicom;
	// * railcom: China Railway Telecom;
	// * telecom: China Telecom;
	// * mobile: China Mobile;
	// * cernet: China Broadcasting Network;
	// * tianwei: China Tianwei;
	// * alibaba: Alibaba Group;
	// * tencent: Tencent Holdings;
	// * drpeng: Dr. Peng Telecom & Media Group;
	// * btvn: Broadcasting Television Network (BTVN);
	// * huashu: Huashu Media;
	// * other: Denotes other/unspecified options.
	ISPList []*string `json:"ISPList,omitempty"`

	// Push-pull streaming protocol. By default, all protocol types are indicated. Supported protocols are as follows.
	// * HTTP-FLV: A push-pull streaming protocol based on the HTTP protocol, using the FLV format for video transmission.
	// * HTTP-HLS: A push-pull streaming protocol based on the HTTP protocol, using the TS format for video transmission.
	// * RTMP: Real-Time Messaging Protocol for real-time message transmission.
	// * RTM: Real-Time Media protocol for ultra-low latency live streaming.
	// * SRT: Secure Reliable Transport protocol for secure and reliable streaming.
	// * QUIC: Quick UDP Internet Connections, a new low-latency internet transmission protocol based on UDP.
	// * CMAF: Common Media Application Format, a versatile streaming protocol.
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// 拉流请求的 Referer 信息。
	RefererList []*string `json:"RefererList,omitempty"`

	// List of regions to which CDN node IPs belong.
	RegionList []*DescribeLiveStreamSessionDataResResultRegionListItem `json:"RegionList,omitempty"`

	// 按维度拆分的数据。
	SessionDetailDataList []*DescribeLiveStreamSessionDataResResultSessionDetailDataListItem `json:"SessionDetailDataList,omitempty"`

	// 流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveStreamSessionDataResResultRegionListItem struct {

	// The regional identifier in regional information.
	Area *string `json:"Area,omitempty"`

	// The country identifier in the regional information.
	Country *string `json:"Country,omitempty"`

	// The province identifier in the regional information.
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveStreamSessionDataResResultSessionDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的在线人数最大值。
	OnlineUser int32 `json:"OnlineUser"`

	// REQUIRED; 当前数据聚合时间粒度内的请求数。
	Request int32 `json:"Request"`

	// REQUIRED; The start time of each time granularity when data is aggregated by time granularity, formatted in RFC3339 UTC
	// time with second precision.
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLiveStreamSessionDataResResultSessionDetailDataListItem struct {

	// REQUIRED; 按维度进行数据拆分后，当前维度的在线人数峰值。
	PeakOnlineUser int32 `json:"PeakOnlineUser"`

	// REQUIRED; 按维度进行数据拆分后，当前维度下所有时间粒度的数据。
	SessionDataList []DescribeLiveStreamSessionDataResResultSessionDetailDataListPropertiesItemsItem `json:"SessionDataList"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的请求数。
	TotalRequest int32 `json:"TotalRequest"`

	// 按域名维度进行数据拆分时的域名信息。
	Domain *string `json:"Domain,omitempty"`

	// 按运营商维度进行数据拆分时的运营商信息。
	ISP *string `json:"ISP,omitempty"`

	// 按推拉流协议维度进行数据拆分时的协议信息。
	Protocol *string `json:"Protocol,omitempty"`

	// 按请求的 Referer 信息进行数据拆分时的 Referer 信息。
	Referer *string `json:"Referer,omitempty"`
}

type DescribeLiveStreamSessionDataResResultSessionDetailDataListPropertiesItemsItem struct {

	// REQUIRED; 在线人数
	OnlineUser int32 `json:"OnlineUser"`

	// REQUIRED; 请求数
	Request int32 `json:"Request"`

	// REQUIRED; The start time of each time granularity when data is aggregated by time granularity, formatted in RFC3339 UTC
	// time with second precision.
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLiveStreamStateQuery struct {

	// REQUIRED; 应用名称，取值与直播流地址的 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App" query:"App"`

	// REQUIRED; 流名称，取值与直播流地址的 StreamName 字段取值相同。支持由大小写字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	Stream string `json:"Stream" query:"Stream"`

	// 填写直播流使用的域名，默认为空，表示查询所有直推流和回源流的状态和类型。 您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，查看需要查询的直播流使用的域名。
	// :::tipVhost 和 Domain 传且仅传一个。 :::
	Domain *string `json:"Domain,omitempty" query:"Domain"`

	// 域名空间，即直播流地址的域名（Domain）所属的域名空间（Vhost）。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]
	// 页面，查看需要查询的直播流使用的域名所属的域名空间。 :::tipVhost 和 Domain 传且仅传一个。 :::
	Vhost *string `json:"Vhost,omitempty" query:"Vhost"`
}

type DescribeLiveStreamStateRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveStreamStateResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeLiveStreamStateResResult          `json:"Result,omitempty"`
}

type DescribeLiveStreamStateResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                           `json:"Version"`
	Error   *DescribeLiveStreamStateResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeLiveStreamStateResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeLiveStreamStateResResult struct {

	// REQUIRED; 直播流状态，取值及含义如下所示。
	// * online：在线流；
	// * offline：历史流；
	// * forbidden：禁推流。
	StreamState string `json:"stream_state"`

	// REQUIRED; 直播流类型，取值及含义如下所示。
	// * push：直推流；
	// * pull：回源流。
	Type string `json:"type"`
}

type DescribeLiveTrafficDataBody struct {

	// REQUIRED; The end time of your query's range (UTC time in RFC 3339 format with second precision).
	EndTime string `json:"EndTime"`

	// REQUIRED; The start time of your query's range (UTC time in RFC 3339 format with second precision).
	StartTime string `json:"StartTime"`

	// The granularity of data aggregation, measured in seconds, with the following supported options:
	// * 300 (default): 5 minutes. When aggregated in 5-minute intervals, the maximum time span for a single query is 31 days,
	// and for historical queries, the maximum time range is 366 days.
	// * 3600: 1 hour. When aggregated in 1-hour intervals, the maximum time span for a single query is 93 days, and for historical
	// queries, the maximum time range is 366 days.
	// * 86400: 1 day. When aggregated in 1-day intervals, the maximum time span for a single query is 93 days, and for historical
	// queries, the maximum time range is 366 days.
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 数据拆分的维度，默认为空表示不按维度进行数据拆分，支持的维度如下。
	// * Domain：域名；
	// * ISP：运营商；
	// * Protocol：推拉流协议。 :::tip 配置数据拆分的维度时，对应的维度参数传入多个值时才会返回按此维度拆分的数据。例如，配置按 Domain 进行数据拆分时， DomainList 传入多个 Domain 值时，才会返回按 Domain
	// 拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，默认为空，表示查询您视频直播产品下所有域名的流量用量数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// Identifiers of operators providing network access services. By default, all operators are indicated. Supported operators
	// are as follows.
	// * unicom: China Unicom;
	// * railcom: China Railway Telecom;
	// * telecom: China Telecom;
	// * mobile: China Mobile;
	// * cernet: China Education and Research Network (CERNET);
	// * tianwei: China Tianwei;
	// * alibaba: Alibaba Group;
	// * tencent: Tencent Holdings;
	// * drpeng: Dr. Peng Telecom & Media Group;
	// * btvn: China Broadcasting Network;
	// * huashu: Huashu Media;
	// * other: Denotes other/unspecified options.
	// If you need to obtain the identifiers of various operators, you can call the DescribeLiveISPData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveispdata].
	ISPList []*string `json:"ISPList,omitempty"`

	// byteplus比火山多了CMAF协议
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// List of regions to which CDN node IPs belong, by default indicating all regions. :::tipRegionList and UserRegionList cannot
	// be used together in the same request. :::
	RegionList []*DescribeLiveTrafficDataBodyRegionListItem `json:"RegionList,omitempty"`

	// 指定查询的流量数据为闲时或忙时，缺省情况下为查询全部数据，支持的取值如下。
	// * busy：忙时；
	// * free：闲时。
	Stage *string `json:"Stage,omitempty"`

	// List of regions to which client IPs belong, by default indicating all regions.
	// :::tipRegionList and UserRegionList cannot be used together in a single request. :::
	UserRegionList []*DescribeLiveTrafficDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveTrafficDataBodyRegionListItem struct {

	// The identifier for the major region in the regional information can be obtained by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata].
	Area *string `json:"Area,omitempty"`

	// The country identifier in the regional information can be obtained by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata].
	// When filtering by
	// country, both 'Area' and 'Country' need to be passed in simultaneously.
	Country *string `json:"Country,omitempty"`

	// The province identifier in the regional information is currently not supported for countries or regions outside mainland
	// China, Hong Kong, Macao, and Taiwan. You can obtain the identifier information
	// by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata]. When
	// filtering by province, you need to simultaneously pass in Area, Country, and
	// Province
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveTrafficDataBodyUserRegionListItem struct {

	// The identifier for the major region in the regional information can be obtained by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata].
	Area *string `json:"Area,omitempty"`

	// The country identifier in the regional information can be obtained by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata].
	// When filtering by
	// country, both 'Area' and 'Country' need to be passed in simultaneously.
	Country *string `json:"Country,omitempty"`

	// The province identifier in the regional information is currently not supported for countries or regions outside mainland
	// China, Hong Kong, Macao, and Taiwan. You can obtain the identifier information
	// by calling DescribeLiveRegionData [https://docs.byteplus.com/en/docs/byteplus-media-live/describeliveregiondata]. When
	// filtering by province, you need to simultaneously pass in Area, Country, and
	// Province
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveTrafficDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveTrafficDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveTrafficDataResResult `json:"Result"`
}

type DescribeLiveTrafficDataResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`
}

type DescribeLiveTrafficDataResResult struct {

	// REQUIRED; Data granularity, measured in seconds.
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; The end time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	EndTime string `json:"EndTime"`

	// REQUIRED; The start time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	StartTime string `json:"StartTime"`

	// REQUIRED; 查询时间范围内的下行总流量，单位为 GB。
	TotalDownTraffic float32 `json:"TotalDownTraffic"`

	// REQUIRED; 查询时间范围内的上行总流量，单位为 GB。
	TotalUpTraffic float32 `json:"TotalUpTraffic"`

	// REQUIRED; 所有时间粒度的数据。
	TrafficDataList []DescribeLiveTrafficDataResResultTrafficDataListItem `json:"TrafficDataList"`

	// 数据拆分的维度，维度说明如下所示。
	// * Domain：域名；
	// * ISP：运营商；
	// * Protocol：推拉流协议。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// Identifiers of operators providing network access services. By default, all operators are indicated. Supported operators
	// are as follows.
	// * unicom: China Unicom;
	// * railcom: China Railway Telecom;
	// * telecom: China Telecom;
	// * mobile: China Mobile;
	// * cernet: China Broadcasting Network;
	// * tianwei: China Tianwei;
	// * alibaba: Alibaba Group;
	// * tencent: Tencent Holdings;
	// * drpeng: Dr. Peng Telecom & Media Group;
	// * btvn: Broadcasting Television Network (BTVN);
	// * huashu: Huashu Media;
	// * other: Denotes other/unspecified options.
	ISPList []*string `json:"ISPList,omitempty"`

	// byteplus比火山多了CMAF协议
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// List of regions to which CDN node IPs belong.
	RegionList []*DescribeLiveTrafficDataResResultRegionListItem `json:"RegionList,omitempty"`

	// 流量数据为闲时或忙时，取值说明如下。
	// * busy：忙时；
	// * free：闲时。
	Stage *string `json:"Stage,omitempty"`

	// 按维度拆分后的数据。 :::tip 当配置了数据拆分的维度时，对应的维度参数传入多个值才会返回按维度拆分的数据。 :::
	TrafficDetailDataList []*DescribeLiveTrafficDataResResultTrafficDetailDataListItem `json:"TrafficDetailDataList,omitempty"`

	// List of regions to which client IPs belong.
	UserRegionList []*DescribeLiveTrafficDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveTrafficDataResResultRegionListItem struct {

	// The regional identifier in regional information.
	Area *string `json:"Area,omitempty"`

	// The country identifier in the regional information.
	Country *string `json:"Country,omitempty"`

	// The province identifier in the regional information.
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveTrafficDataResResultTrafficDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内产生的总下行流量，单位 GB。
	DownTraffic float32 `json:"DownTraffic"`

	// REQUIRED; The start time of each time granularity when data is aggregated by time granularity, formatted in RFC3339 UTC
	// time with second precision.
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内产生的总上行流量，单位 GB。
	UpTraffic float32 `json:"UpTraffic"`
}

type DescribeLiveTrafficDataResResultTrafficDetailDataListItem struct {

	// REQUIRED; 按维度进行数据拆分后，当前维度的下行总流量，单位为 GB。
	TotalDownTraffic float32 `json:"TotalDownTraffic"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的上行总流量，单位为 GB。
	TotalUpTraffic float32 `json:"TotalUpTraffic"`

	// REQUIRED; 按维度进行数据拆分后，当前维度下所有时间粒度的数据。
	TrafficDataList []DescribeLiveTrafficDataResResultTrafficDetailDataListPropertiesItemsItem `json:"TrafficDataList"`

	// 按域名维度进行数据拆分时的域名信息。
	Domain *string `json:"Domain,omitempty"`

	// 按运营商维度进行数据拆分时的运营商信息。
	ISP *string `json:"ISP,omitempty"`

	// 按推拉流协议维度进行数据拆分时的协议信息。
	Protocol *string `json:"Protocol,omitempty"`
}

type DescribeLiveTrafficDataResResultTrafficDetailDataListPropertiesItemsItem struct {

	// REQUIRED; 下行流量，单位 GB
	DownTraffic float32 `json:"DownTraffic"`

	// REQUIRED; The start time of each time granularity when data is aggregated by time granularity, formatted in RFC3339 UTC
	// time with second precision.
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 上行流量，单位 GB
	UpTraffic float32 `json:"UpTraffic"`
}

type DescribeLiveTrafficDataResResultUserRegionListItem struct {

	// The regional identifier in regional information.
	Area *string `json:"Area,omitempty"`

	// The country identifier in the regional information.
	Country *string `json:"Country,omitempty"`

	// The province identifier in the regional information.
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveTranscodeDataBody struct {

	// REQUIRED; The end time of your query's range (UTC time in RFC 3339 format with second precision).
	EndTime string `json:"EndTime"`

	// REQUIRED; The start time of your query's range (UTC time in RFC 3339 format with second precision).
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，当前接口默认且仅支持按 86400 秒进行数据聚合。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// The Application Name must match the value of the AppName field in the live stream URL. It can include uppercase letters
	// (A-Z), lowercase letters (a-z), numbers (0-9), underscores (_), hyphens (-), and
	// periods (.), with a length ranging from 1 to 30 characters.
	// :::tip When querying stream granularity data, both the App and Stream parameters are required. :::
	App *string `json:"App,omitempty"`

	// 域名列表，默认为空，表示查询您视频直播产品下所有域名的转码用量数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 分辨率。- 480P：640 × 480； - 720P：1280 × 720； - 1080P：1920 × 1088； - 2K：2560 × 1440； - 4K：4096 × 2160；- 8K：大于4K； - 0P：纯音频流；
	Resolution []*string `json:"Resolution,omitempty"`

	// The Stream Name must correspond to the value of the StreamName field in the live stream URL. It can include uppercase and
	// lowercase letters (A-Z, a-z), numbers (0-9), underscores (_), hyphens (-), and
	// periods (.), with a length ranging from 1 to 100 characters.
	// :::tip When querying stream granularity data, both the App and Stream parameters must be provided. :::
	Stream *string `json:"Stream,omitempty"`

	// BytePlus 不支持 ByteQE（画质增强）
	TranscodeType []*string `json:"TranscodeType,omitempty"`
}

type DescribeLiveTranscodeDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveTranscodeDataResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeLiveTranscodeDataResResult          `json:"Result,omitempty"`
}

type DescribeLiveTranscodeDataResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                             `json:"Version"`
	Error   *DescribeLiveTranscodeDataResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeLiveTranscodeDataResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeLiveTranscodeDataResResult struct {

	// REQUIRED; Data granularity, measured in seconds.
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询时间范围内的转码总时长，单位为分钟。
	Duration float32 `json:"Duration"`

	// REQUIRED; The end time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	EndTime string `json:"EndTime"`

	// REQUIRED; The start time of the query in UTC time, formatted according to RFC3339, with precision to the second.
	StartTime string `json:"StartTime"`

	// REQUIRED; 所有时间粒度的数据。
	TranscodeDataList []DescribeLiveTranscodeDataResResultTranscodeDataListItem `json:"TranscodeDataList"`

	// The application name when querying stream granularity data.
	App *string `json:"App,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 分辨率。- 480P：640 × 480； - 720P：1280 × 720； - 1080P：1920 × 1088； - 2K：2560 × 1440； - 4K：4096 × 2160；- 8K：大于4K； - 0P：纯音频流；
	Resolution []*string `json:"Resolution,omitempty"`

	// The stream name when querying stream granularity data.
	Stream *string `json:"Stream,omitempty"`

	// 视频编码格式，支持的取值和含义如下所示。- NormalH264：H.264 标准转码； - NormalH265：H.265 标准转码； - NormalH266：H.266 标准转码； - ByteHDH264：H.264 极智超清；
	// - ByteHDH265：H.265 极智超清； - ByteHDH266：H.266 极智超清；- ByteQE：画质增强；- Audio：纯音频流；
	TranscodeType []*string `json:"TranscodeType,omitempty"`
}

type DescribeLiveTranscodeDataResResultTranscodeDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的转码时长，单位为分钟。
	Duration float32 `json:"Duration"`

	// REQUIRED; 分辨率。- 480P：640 × 480； - 720P：1280 × 720； - 1080P：1920 × 1088； - 2K：2560 × 1440； - 4K：4096 × 2160；- 8K：大于4K； -
	// 0P：纯音频流；
	Resolution string `json:"Resolution"`

	// REQUIRED; The start time of each time granularity when data is aggregated by time granularity, formatted in RFC3339 UTC
	// time with second precision.
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 视频编码格式，支持的取值和含义如下所示。- NormalH264：H.264 标准转码； - NormalH265：H.265 标准转码； - NormalH266：H.266 标准转码； - ByteHDH264：H.264
	// 极智超清； - ByteHDH265：H.265 极智超清； - ByteHDH266：H.266 极智超清；- ByteQE：画质增强；- Audio：纯音频流；
	TranscodeType string `json:"TranscodeType"`
}

type DescribeRecordTaskFileHistoryBody struct {

	// REQUIRED; 开始录制时间，RFC3339 格式的时间戳，精度为秒。当您查询指定录制任务详情时，DateFrom 应设置为开始时间之前的任意时间。
	DateFrom string `json:"DateFrom"`

	// REQUIRED; 结束录制时间，RFC3339 格式的时间戳，精度为秒。结束时间需晚于 DateFrom，且与 DateFrom 间隔不超过 7 天。
	DateTo string `json:"DateTo"`

	// REQUIRED; 查询数据的页码，默认为 1，表示查询第一页的数据，取值范围为正整数。
	PageNum int32 `json:"PageNum"`

	// REQUIRED; 每页显示的数据条数，取值范围为正整数。
	PageSize int32 `json:"PageSize"`

	// 应用名称，取值与直播流地址的 AppName 字段取值相同，默认为空表示查询 vhost 下的所有录制历史。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30
	// 个字符。
	App *string `json:"App,omitempty"`

	// 流名称，取值与直播流地址的 StreamName 字段取值相同，默认为空表示查询 App 下的所有录制历史。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100
	// 个字符。 :::tip 如果指定 Stream，必须同时指定 App 的值。 :::
	Stream *string `json:"Stream,omitempty"`

	// 录制文件保存位置，支持的取值及含义如下所示。
	// * tos：存储到 TOS（默认值）；
	// * vod：存储到 VOD。
	Type *string `json:"Type,omitempty"`

	// 域名空间，即直播流地址的域名所属的域名空间，默认为空表示查询所有录制历史。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名所属的域名空间。
	Vhost *string `json:"Vhost,omitempty"`
}

type DescribeRecordTaskFileHistoryRes struct {

	// REQUIRED
	ResponseMetadata DescribeRecordTaskFileHistoryResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeRecordTaskFileHistoryResResult          `json:"Result,omitempty"`
}

type DescribeRecordTaskFileHistoryResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                                 `json:"Version"`
	Error   *DescribeRecordTaskFileHistoryResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeRecordTaskFileHistoryResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeRecordTaskFileHistoryResResult struct {

	// REQUIRED; 录制文件列表。
	Data []DescribeRecordTaskFileHistoryResResultDataItem `json:"Data"`

	// REQUIRED; 查询结果的分页信息。
	Pagination DescribeRecordTaskFileHistoryResResultPagination `json:"Pagination"`
}

type DescribeRecordTaskFileHistoryResResultDataItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 存储位置为 TOS 时的 Bucket。
	Bucket string `json:"Bucket"`

	// REQUIRED; 录制时长。
	Duration string `json:"Duration"`

	// REQUIRED; 结束录制时间。
	EndTime string `json:"EndTime"`

	// REQUIRED; 结束录制时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTimeUTC string `json:"EndTimeUTC"`

	// REQUIRED; 录制文件的文件名。
	FileName string `json:"FileName"`

	// REQUIRED; 录制文件存储格式。
	Format string `json:"Format"`

	// REQUIRED; 存储位置为 TOS 时，在 Bucket 中的存储路径。
	Path string `json:"Path"`

	// REQUIRED; 开始录制时间。
	StartTime string `json:"StartTime"`

	// REQUIRED; 开始录制时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTimeUTC string `json:"StartTimeUTC"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 域名空间。
	Vhost string `json:"Vhost"`

	// REQUIRED; 录制文件保存在 VOD 时，录制文件的 ID。
	Vid string `json:"Vid"`
}

// DescribeRecordTaskFileHistoryResResultPagination - 查询结果的分页信息。
type DescribeRecordTaskFileHistoryResResultPagination struct {

	// REQUIRED; 当前所在分页的页码。
	PageCur int32 `json:"PageCur"`

	// REQUIRED; 每页显示的数据条数。
	PageSize int32 `json:"PageSize"`

	// REQUIRED; 查询结果的数据总页数。
	PageTotal int32 `json:"PageTotal"`

	// REQUIRED; 查询结果的数据总条数。
	TotalCount int32 `json:"TotalCount"`
}

type DescribeRefererBody struct {

	// 应用名称，取值与直播流地址中 AppName 字段取值相同，默认为空，表示所有应用名称。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。 :::tip
	// 参数 Domain 和 App 至少传一个。 :::
	App *string `json:"App,omitempty"`

	// 拉流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的拉流域名。
	// :::tip
	// * 参数 Domain 和 Vhost 传且仅传一个。
	// * 参数 Domain 和 App 至少传一个。 :::
	Domain *string `json:"Domain,omitempty"`

	// 域名空间，即直播流地址的域名（Domain）所属的域名空间（Vhost）。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，查看拉流域名所属的域名空间。 :::tip
	// 参数 Domain 和 Vhost 传且仅传一个。 :::
	Vhost *string `json:"Vhost,omitempty"`
}

type DescribeRefererRes struct {

	// REQUIRED
	ResponseMetadata DescribeRefererResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeRefererResResult          `json:"Result,omitempty"`
}

type DescribeRefererResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                   `json:"Version"`
	Error   *DescribeRefererResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeRefererResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeRefererResResult struct {

	// Referer 防盗链信息列表。
	RefererList []*DescribeRefererResResultRefererListItem `json:"RefererList,omitempty"`
}

type DescribeRefererResResultRefererListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; Referer 防盗链详情列表。
	RefererInfoList []DescribeRefererResResultRefererListPropertiesItemsItem `json:"RefererInfoList"`

	// REQUIRED; 域名空间。
	Vhost string `json:"Vhost"`
}

type DescribeRefererResResultRefererListPropertiesItemsItem struct {

	// REQUIRED; 用于标识 referer 防盗链的关键词，返回值为 referer。
	Key string `json:"Key"`

	// REQUIRED; 优先级，当前默认返回值为 0。当多域名返回值一致时，按照域名输入顺序区分，越早加入列表的域名优先级越高。
	Priority int32 `json:"Priority"`

	// REQUIRED; referer 防盗链黑白名单类型，取值即含义如下所示。
	// * deny：黑名单；
	// * allow：白名单。
	Type string `json:"Type"`

	// REQUIRED; Referer 字段规则，即设置的黑名单或白名单的域名。
	Value string `json:"Value"`
}

type DescribeRegionAccessRuleBody struct {

	// REQUIRED
	Domain string `json:"Domain"`

	// REQUIRED
	Vhost string `json:"Vhost"`
}

type DescribeRegionAccessRuleRes struct {

	// REQUIRED
	ResponseMetadata DescribeRegionAccessRuleResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeRegionAccessRuleResResult `json:"Result,omitempty"`
}

type DescribeRegionAccessRuleResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

// DescribeRegionAccessRuleResResult - 视请求的接口而定
type DescribeRegionAccessRuleResResult struct {

	// REQUIRED
	AccessRuleLists []DescribeRegionAccessRuleResResultAccessRuleListsItem `json:"AccessRuleLists"`
}

type DescribeRegionAccessRuleResResultAccessRuleListsItem struct {

	// REQUIRED
	Domain string `json:"Domain"`

	// REQUIRED
	RegionAccessRule DescribeRegionAccessRuleResResultAccessRuleListsItemRegionAccessRule `json:"RegionAccessRule"`

	// REQUIRED
	Vhost string  `json:"Vhost"`
	App   *string `json:"App,omitempty"`
}

type DescribeRegionAccessRuleResResultAccessRuleListsItemRegionAccessRule struct {

	// REQUIRED
	Enable string `json:"Enable"`

	// REQUIRED
	Type         string    `json:"Type"`
	CountryList  []*string `json:"CountryList,omitempty"`
	ProvinceList []*string `json:"ProvinceList,omitempty"`
}

type DescribeRelaySourceRewriteBody struct {

	// REQUIRED; 域名空间名称
	Vhost string `json:"Vhost"`

	// 需要设置黑白名单的拉流域名。域名请在工信部完成备案。
	Domain *string `json:"Domain,omitempty"`
}

type DescribeRelaySourceRewriteRes struct {
	ResponseMetadata *DescribeRelaySourceRewriteResResponseMetadata `json:"ResponseMetadata,omitempty"`
	Result           *DescribeRelaySourceRewriteResResult           `json:"Result,omitempty"`
}

type DescribeRelaySourceRewriteResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`
}

type DescribeRelaySourceRewriteResResult struct {

	// 回源改写列表
	RelaySourceRewriteList *DescribeRelaySourceRewriteResResultRelaySourceRewriteList `json:"RelaySourceRewriteList,omitempty"`
}

// DescribeRelaySourceRewriteResResultRelaySourceRewriteList - 回源改写列表
type DescribeRelaySourceRewriteResResultRelaySourceRewriteList struct {

	// 需要设置黑白名单的拉流域名。域名请在工信部完成备案。
	Domain *string `json:"Domain,omitempty"`

	// 改写规则
	RewriteRule *DescribeRelaySourceRewriteResResultRelaySourceRewriteListRewriteRule `json:"RewriteRule,omitempty"`

	// 域名空间名称
	Vhost *string `json:"Vhost,omitempty"`
}

// DescribeRelaySourceRewriteResResultRelaySourceRewriteListRewriteRule - 改写规则
type DescribeRelaySourceRewriteResResultRelaySourceRewriteListRewriteRule struct {

	// 功能开关。- true: 开 - false: 关
	Enable *bool `json:"Enable,omitempty"`

	// 改写规则列表
	RewriteRuleList []*DescribeRelaySourceRewriteResResultRelaySourceRewriteListRewriteRuleListItem `json:"RewriteRuleList,omitempty"`
}

type DescribeRelaySourceRewriteResResultRelaySourceRewriteListRewriteRuleListItem struct {

	// 改写后地址是否包含原始地址的param参数
	IncludeParams *bool `json:"IncludeParams,omitempty"`

	// 原始path
	OriginPath *string `json:"OriginPath,omitempty"`

	// 改写后目标path
	TargetPath *string `json:"TargetPath,omitempty"`
}

type DescribeRelaySourceV3Body struct {

	// REQUIRED; 直播流使用的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名。所属的域名空间。
	Vhost string `json:"Vhost"`

	// 应用名称，即直播流地址的AppName字段取值，默认为空，表示查询当前域名空间下所有播放触发回源配置。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App *string `json:"App,omitempty"`

	// 回源组名称。
	Group *string `json:"Group,omitempty"`
}

type DescribeRelaySourceV3Res struct {

	// REQUIRED
	ResponseMetadata DescribeRelaySourceV3ResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeRelaySourceV3ResResult          `json:"Result,omitempty"`
}

type DescribeRelaySourceV3ResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                         `json:"Version"`
	Error   *DescribeRelaySourceV3ResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeRelaySourceV3ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeRelaySourceV3ResResult struct {

	// 回源配置列表。
	RelaySourceConfigList []*DescribeRelaySourceV3ResResultRelaySourceConfigListItem `json:"RelaySourceConfigList,omitempty"`
}

type DescribeRelaySourceV3ResResultRelaySourceConfigListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 回源组配置详情。
	GroupDetails []DescribeRelaySourceV3ResResultRelaySourceConfigListPropertiesItemsItem `json:"GroupDetails"`

	// REQUIRED; 域名空间。
	Vhost string `json:"Vhost"`

	// 组的重试次数
	RetryTimes *int64 `json:"RetryTimes,omitempty"`
}

type DescribeRelaySourceV3ResResultRelaySourceConfigListPropertiesItemsItem struct {

	// REQUIRED; 回源组名称。
	Group string `json:"Group"`

	// REQUIRED; 回源服务器配置列表。
	Servers []DescribeRelaySourceV3ResResultRelaySourceConfigListPropertiesItemsServersItem `json:"Servers"`
}

type DescribeRelaySourceV3ResResultRelaySourceConfigListPropertiesItemsServersItem struct {

	// REQUIRED; 回源地址。
	RelaySourceDomain string `json:"RelaySourceDomain"`

	// REQUIRED; 自定义回源参数。
	RelaySourceParams map[string]string `json:"RelaySourceParams"`

	// REQUIRED; 回源协议。
	RelaySourceProtocol string `json:"RelaySourceProtocol"`
}

type DescribeUserAgentAccessRuleBody struct {

	// REQUIRED
	Vhost  string  `json:"Vhost"`
	Domain *string `json:"Domain,omitempty"`
}

type DescribeUserAgentAccessRuleRes struct {

	// REQUIRED
	ResponseMetadata DescribeUserAgentAccessRuleResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeUserAgentAccessRuleResResult `json:"Result"`
}

type DescribeUserAgentAccessRuleResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`
}

type DescribeUserAgentAccessRuleResResult struct {

	// REQUIRED
	UserAgentList []DescribeUserAgentAccessRuleResResultUserAgentListItem `json:"UserAgentList"`
}

type DescribeUserAgentAccessRuleResResultUserAgentListItem struct {
	Domain       *string                                                            `json:"Domain,omitempty"`
	UaAccessRule *DescribeUserAgentAccessRuleResResultUserAgentListItemUaAccessRule `json:"UaAccessRule,omitempty"`
	Vhost        *string                                                            `json:"Vhost,omitempty"`
}

type DescribeUserAgentAccessRuleResResultUserAgentListItemUaAccessRule struct {

	// REQUIRED
	AllowEmpty bool `json:"AllowEmpty"`

	// REQUIRED
	Enable bool `json:"Enable"`

	// REQUIRED
	Type string `json:"Type"`

	// REQUIRED
	UserAgent []string `json:"UserAgent"`
}

type DisableDomainBody struct {

	// REQUIRED; 待禁用域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看需要待禁用域名的信息。
	Domain string `json:"Domain"`
}

type DisableDomainRes struct {

	// REQUIRED
	ResponseMetadata DisableDomainResResponseMetadata `json:"ResponseMetadata"`
}

type DisableDomainResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                 `json:"Version"`
	Error   *DisableDomainResResponseMetadataError `json:"Error,omitempty"`
}

type DisableDomainResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type EnableDomainBody struct {

	// REQUIRED; 待启用域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看需要待启用域名的信息。
	Domain string `json:"Domain"`
}

type EnableDomainRes struct {

	// REQUIRED
	ResponseMetadata EnableDomainResResponseMetadata `json:"ResponseMetadata"`
}

type EnableDomainResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                `json:"Version"`
	Error   *EnableDomainResResponseMetadataError `json:"Error,omitempty"`
}

type EnableDomainResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type EnableHTTPHeaderConfigBody struct {

	// REQUIRED; 启用或禁用配置，取值及含义如下所示。
	// * true：启用；
	// * false：禁用。
	Enable bool `json:"Enable"`

	// REQUIRED; HTTP Header 类型，您可以调用 DescribeHTTPHeaderConfig [https://www.volcengine.com/docs/6469/1232744] 接口查看 HTTP Header
	// 配置的 Phase 取值。
	Phase int32 `json:"Phase"`

	// REQUIRED; 域名空间，您可以调用 DescribeHTTPHeaderConfig [https://www.volcengine.com/docs/6469/1232744] 接口查看 HTTP Header 配置的 Vhost
	// 取值。
	Vhost string `json:"Vhost"`

	// 拉流域名，您可以调用 DescribeHTTPHeaderConfig [https://www.volcengine.com/docs/6469/1232744] 接口查看 HTTP Header 配置的 Domain 取值。
	Domain *string `json:"Domain,omitempty"`
}

type EnableHTTPHeaderConfigRes struct {

	// REQUIRED
	ResponseMetadata EnableHTTPHeaderConfigResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type EnableHTTPHeaderConfigResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type ForbidStreamBody struct {

	// REQUIRED; 应用名称，取值与直播流地址的 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 流名称，取值与直播流地址的 StreamName 字段取值相同。支持由大小写字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	Stream string `json:"Stream"`

	// 直播流使用的域名，您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看待禁推的直播流使用的域名。
	// :::tip 参数 Domain 和 Vhost
	// 传且仅传一个。 :::
	Domain *string `json:"Domain,omitempty"`

	// 禁推的结束时间，RFC3339 格式的 UTC 时间，精度为毫秒，禁推有效期最长为 90 天，默认为当前时间加 90 天。
	EndTime *string `json:"EndTime,omitempty"`

	// 域名空间，即直播流地址的域名（Domain）所属的域名空间（Vhost）。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]
	// 页面，查看待禁推的直播流使用的域名所属的域名空间。 :::tip 参数 Domain 和 Vhost 传且仅传一个。 :::
	Vhost *string `json:"Vhost,omitempty"`
}

type ForbidStreamRes struct {

	// REQUIRED
	ResponseMetadata ForbidStreamResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type ForbidStreamResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                `json:"Version"`
	Error   *ForbidStreamResResponseMetadataError `json:"Error,omitempty"`
}

type ForbidStreamResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type GeneratePlayURLBody struct {

	// REQUIRED; 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 拉流域名。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 流名称，取值与直播流地址中 StreamName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	Stream string `json:"Stream"`

	// 拉流地址的过期时间，RFC3339 格式的 UTC 时间，精度为秒，过期后需要重新生成。缺省情况下表示当前时间往后的 7 天。 :::tip 如果同时设置 ValidDuration 和 ExpiredTime，以 ExpiredTime
	// 的时间为准。 :::
	ExpiredTime *string `json:"ExpiredTime,omitempty"`

	// 转码流后缀，默认为空，表示生成源流的拉流地址。配置不为空时表示生成转码流的拉流地址，可通过调用 ListVhostTransCodePreset [https://www.volcengine.com/docs/6469/1126853]
	// 接口查询对应流的转码流后缀。
	Suffix *string `json:"Suffix,omitempty"`

	// CDN 类型，默认值为 fcdn，支持如下取值。
	// * fcdn：火山引擎流媒体直播 CDN；
	// * 3rd：第三方 CDN。
	Type *string `json:"Type,omitempty"`

	// 拉流地址的有效时长，单位为秒，超过有效时长后需要重新生成。缺省情况下表示 7 天，取值范围为正整数。 :::tip 如果同时设置 ValidDuration 和 ExpiredTime，以 ExpiredTime 的时间为准。 :::
	ValidDuration *int32 `json:"ValidDuration,omitempty"`

	// 域名空间名称
	Vhost *string `json:"Vhost,omitempty"`
}

type GeneratePlayURLRes struct {

	// REQUIRED
	ResponseMetadata GeneratePlayURLResResponseMetadata `json:"ResponseMetadata"`
	Result           *GeneratePlayURLResResult          `json:"Result,omitempty"`
}

type GeneratePlayURLResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                   `json:"Version"`
	Error   *GeneratePlayURLResResponseMetadataError `json:"Error,omitempty"`
}

type GeneratePlayURLResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type GeneratePlayURLResResult struct {

	// REQUIRED; 拉流地址信息列表。
	URLList []GeneratePlayURLResResultURLListItem `json:"URLList"`
}

type GeneratePlayURLResResultURLListItem struct {

	// REQUIRED; CDN 类型。
	// * fcdn：火山引擎流媒体直播 CDN；
	// * 3rd：第三方 CDN。
	CDN string `json:"CDN"`

	// REQUIRED; 协议类型，包括 hls、flv、rtmp、udp 和 cmaf。
	Protocol string `json:"Protocol"`

	// REQUIRED; 地址类型，取值及含义如下所示。
	// * pull：拉流地址；
	// * 3rd_play(relay_source)：第三方回源地址，当配置了回源且 CDN 类型为第三方 CDN 时返回；
	// * 3rd_play(relay_sink)：第三方转推地址，当配置了拉流转推且 CDN 类型为第三方 CDN 时返回。
	Type string `json:"Type"`

	// REQUIRED; 生成的拉流地址。
	URL string `json:"URL"`
}

type GeneratePushURLBody struct {

	// REQUIRED; 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 流名称，取值与直播流地址中 StreamName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	Stream string `json:"Stream"`

	// REQUIRED; 域名空间，即推流域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看推流域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 推流域名，默认为空，表示生成域名空间下所有推流域名的推流地址。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看需要的推流域名。
	Domain *string `json:"Domain,omitempty"`

	// 推流地址的过期时间，RFC3339 格式的时间字符串，精度为秒，过期后需要重新生成。缺省情况下表示当前时间往后的 7 天。 :::tip 如果同时设置 ValidDuration 和 ExpiredTime，以 ExpiredTime 的时间为准。
	// :::
	ExpiredTime *string `json:"ExpiredTime,omitempty"`

	// 推流地址的有效时长，单位为秒，超过有效时长后需要重新生成。缺省情况下表示 7 天，取值范围为正整数。 :::tip 如果同时设置 ValidDuration 和 ExpiredTime，以 ExpiredTime 的时间为准。 :::
	ValidDuration *int32 `json:"ValidDuration,omitempty"`
}

type GeneratePushURLRes struct {

	// REQUIRED
	ResponseMetadata GeneratePushURLResResponseMetadata `json:"ResponseMetadata"`
	Result           *GeneratePushURLResResult          `json:"Result,omitempty"`
}

type GeneratePushURLResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                   `json:"Version"`
	Error   *GeneratePushURLResResponseMetadataError `json:"Error,omitempty"`
}

type GeneratePushURLResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type GeneratePushURLResResult struct {

	// REQUIRED; RTMP 推流地址。
	PushURLList []string `json:"PushURLList"`

	// REQUIRED; 推流地址详情。
	PushURLListDetail []GeneratePushURLResResultPushURLListDetailItem `json:"PushURLListDetail"`

	// REQUIRED; RTM 推流地址。
	RtmURLList []string `json:"RtmURLList"`

	// REQUIRED; RTMP over SRT 推流地址。
	RtmpOverSrtURLList []string `json:"RtmpOverSrtURLList"`

	// REQUIRED; TS over SRT 推流地址。
	TsOverSrtURLList []string `json:"TsOverSrtURLList"`

	// REQUIRED; 网络传输推流地址。
	WebTransportURLList []string `json:"WebTransportURLList"`
}

type GeneratePushURLResResultPushURLListDetailItem struct {

	// REQUIRED; OBS 推流地址，例如，rtmp://push.example.com/live/。
	DomainApp string `json:"DomainApp"`

	// REQUIRED; OBS 推流名称，例如，streamname1?volcTime=1675652376&volcSecret=c57d247c2f19b395b6ec9b182******7。
	StreamSign string `json:"StreamSign"`

	// REQUIRED; 推流地址。
	URL string `json:"URL"`
}

type KillStreamBody struct {

	// 直播流使用的应用名称。
	App *string `json:"App,omitempty"`

	// 推流域名。 参数 Domain 和 Vhost传且仅传一个。
	Domain *string `json:"Domain,omitempty"`

	// 禁推的结束时间，禁推有效期最长为 90 天，默认为当前时间加 90 天
	EndTime *string `json:"EndTime,omitempty"`

	// 直播流使用的流名称。
	Stream *string `json:"Stream,omitempty"`

	// 域名空间，您可以调用 DescribeLiveStreamInfoByPage [https://www.volcengine.com/docs/6469/1126841] 接口，查看待断开的在线流的信息，包括 Vhost、APP 和 Stream。
	Vhost *string `json:"Vhost,omitempty"`
}

type KillStreamRes struct {

	// REQUIRED
	ResponseMetadata KillStreamResResponseMetadata `json:"ResponseMetadata"`
}

type KillStreamResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                              `json:"Version"`
	Error   *KillStreamResResponseMetadataError `json:"Error,omitempty"`
}

type KillStreamResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type ListBindEncryptDRMBody struct {

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 应用名称，取值与直播流地址中 AppName 字段取值相同，默认为空，表示查询符合域名空间取值的所有的 DRM 加密配置。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为
	// 1 到 30 个字符。
	App *string `json:"App,omitempty"`
}

type ListBindEncryptDRMRes struct {

	// REQUIRED
	ResponseMetadata ListBindEncryptDRMResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 视请求的接口而定
	Result ListBindEncryptDRMResResult `json:"Result"`
}

type ListBindEncryptDRMResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

// ListBindEncryptDRMResResult - 视请求的接口而定
type ListBindEncryptDRMResResult struct {

	// DRM 加密配置列表。
	DRMBindingList []*ListBindEncryptDRMResResultDRMBindingListItem `json:"DRMBindingList,omitempty"`
}

type ListBindEncryptDRMResResultDRMBindingListItem struct {

	// REQUIRED; 账号
	AccountID string `json:"AccountID"`

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 当前 DRM 配置是否开启，取值及含义如下所示。
	// * true：开启；
	// * false：关闭。
	Enable bool `json:"Enable"`

	// REQUIRED; 进行 DRM 加密的转码流对应的转码流后缀配置。
	EncryptTranscodeSuffix []string `json:"EncryptTranscodeSuffix"`

	// REQUIRED; 域名空间。
	Vhost string `json:"Vhost"`
}

type ListCertV2Body struct {

	// 火山引擎账号 ID
	AccountID *string `json:"AccountID,omitempty"`

	// 证书是否启用，默认值为 true，支持的取值及含义如下所示。
	// * true：启用证书；
	// * false：禁用证书。
	Available *bool `json:"Available,omitempty"`

	// 证书名称，支持输入证书名称中的关键字，进行模糊查询.
	CertName *string `json:"CertName,omitempty"`

	// 域名，查询该域名对应的证书，支持精确查询。默认为空，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看域名信息。
	Domain *string `json:"Domain,omitempty"`

	// 只有填了Available，这个字段才生效。
	Expiring *bool `json:"Expiring,omitempty"`
}

type ListCertV2Res struct {

	// REQUIRED
	ResponseMetadata ListCertV2ResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListCertV2ResResult          `json:"Result,omitempty"`
}

type ListCertV2ResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                              `json:"Version"`
	Error   *ListCertV2ResResponseMetadataError `json:"Error,omitempty"`
}

type ListCertV2ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type ListCertV2ResResult struct {

	// 证书列表。
	CertList []*ListCertV2ResResultCertListItem `json:"CertList,omitempty"`
}

type ListCertV2ResResultCertListItem struct {

	// REQUIRED; 与证书绑定的域名列表。
	CertDomainList []string `json:"CertDomainList"`

	// REQUIRED; 证书实例 ID。
	CertID string `json:"CertID"`

	// REQUIRED; 证书名称。
	CertName string `json:"CertName"`

	// REQUIRED; 证书链 ID。
	ChainID string `json:"ChainID"`

	// REQUIRED; 火山引擎证书中心证书链 ID。
	ChainIDVolc string `json:"ChainIDVolc"`

	// REQUIRED; 证书的过期时间，RFC3339 格式的 UTC 时间，精度为秒。
	NotAfter string `json:"NotAfter"`

	// REQUIRED; 证书的生效日期，RFC3339 格式的 UTC 时间，精度为秒。
	NotBefore string `json:"NotBefore"`

	// REQUIRED; 项目名称。
	ProjectName string `json:"ProjectName"`

	// REQUIRED; 证书状态，由证书管理平台返回，支持的取值如下所示。
	// * OK：正常；
	// * Expire：过期；
	// * 30days：有效期剩余 30 天；
	// * 15days：有效期剩余 15 天；
	// * 7days：有效期剩余 7 天；
	// * 1days：有效期剩余 1 天。
	Status string `json:"Status"`
}

type ListCommonTransPresetDetailBody struct {

	// 模板名称列表，缺省情况下，表示查询所有系统内置转码档位。
	PresetList []*string `json:"PresetList,omitempty"`
}

type ListCommonTransPresetDetailRes struct {

	// REQUIRED
	ResponseMetadata ListCommonTransPresetDetailResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListCommonTransPresetDetailResResult          `json:"Result,omitempty"`
}

type ListCommonTransPresetDetailResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version   string                                               `json:"Version"`
	Error     *ListCommonTransPresetDetailResResponseMetadataError `json:"Error,omitempty"`
	RequestID *string                                              `json:"RequestID,omitempty"`
}

type ListCommonTransPresetDetailResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type ListCommonTransPresetDetailResResult struct {

	// 极智超清转码配置。
	NarrowBandHDPresetDetail []*ListCommonTransPresetDetailResResultNarrowBandHDPresetDetailItem `json:"NarrowBandHDPresetDetail,omitempty"`

	// 标准转码配置。
	StandardPresetDetail []*ListCommonTransPresetDetailResResultStandardPresetDetailItem `json:"StandardPresetDetail,omitempty"`
}

type ListCommonTransPresetDetailResResultNarrowBandHDPresetDetailItem struct {
	ALayout   *string `json:"ALayout,omitempty"`
	AProfile  *string `json:"AProfile,omitempty"`
	AR        *int32  `json:"AR,omitempty"`
	AbrMode   *int32  `json:"AbrMode,omitempty"`
	AccountID *string `json:"AccountID,omitempty"`

	// 音频编码格式，支持的取值及含义如下。
	// * aac：使用 AAC 编码格式；
	// * opus：使用 Opus 编码格式；
	// * copy：不进行转码，所有音频编码参数不生效。
	Acodec         *string `json:"Acodec,omitempty"`
	AdvancedParam  *string `json:"AdvancedParam,omitempty"`
	AllowAudioCopy *int32  `json:"AllowAudioCopy,omitempty"`
	AllowVideoCopy *int32  `json:"AllowVideoCopy,omitempty"`
	An             *int32  `json:"An,omitempty"`

	// 视频分辨率自适应模式开关。支持的取值及含义如下。
	// * 0：关闭视频分辨率自适应；
	// * 1：开启视频分辨率自适应。 :::tip
	// * 关闭视频分辨率自适应模式（As 取值为 0）时，转码配置的视频分辨率取视频宽度（Width）和视频高度（Height）的值对转码视频进行拉伸；
	// * 开启视频分辨率自适应模式（As 取值为 1）时，转码配置的视频分辨率按照短边长度（ShortSide）、长边长度（LongSide）、视频宽度（Width）、视频高度（Height）的优先级取值，另一边等比缩放。 :::
	As *string `json:"As,omitempty"`

	// 音频码率，单位为 kbps。
	AudioBitrate        *int32 `json:"AudioBitrate,omitempty"`
	AutoTransAb         *int32 `json:"AutoTransAb,omitempty"`
	AutoTransAl         *int32 `json:"AutoTransAl,omitempty"`
	AutoTransAr         *int32 `json:"AutoTransAr,omitempty"`
	AutoTransResolution *int32 `json:"AutoTransResolution,omitempty"`
	AutoTransVb         *int32 `json:"AutoTransVb,omitempty"`
	AutoTransVr         *int32 `json:"AutoTransVr,omitempty"`
	BCM                 *int32 `json:"BCM,omitempty"`

	// 转码输出视频中 2 个参考帧之间的最大 B 帧数量，默认值为 3，取值为 0 时表示去除 B 帧。
	// 最大 B 帧数量的取值范围根据视频编码格式（Vcodec）的不同有所差异，取值范围如下所示。
	// * 视频编码格式为 H.264 （Vcodec 取值为 h264）时取值范围为 [0,7]；
	// * 视频编码格式为 H.265 或 H.266 （Vcodec 取值为 h265 或 h266）时取值范围为 [0,3]、7、15。
	BFrames  *int32  `json:"BFrames,omitempty"`
	Describe *string `json:"Describe,omitempty"`

	// 帧率，单位为 fps。帧率越大，画面越流畅。
	FPS *int32 `json:"FPS,omitempty"`

	// IDR 帧之间的最大间隔，单位为秒。
	GOP    *int32 `json:"GOP,omitempty"`
	GopMin *int32 `json:"GopMin,omitempty"`
	HVSPre *bool  `json:"HVSPre,omitempty"`

	// 视频高度。 :::tip
	// * 当关闭视频分辨率自适应（As 取值为 0）时，转码分辨率将取 Width 和 Height 的值对转码视频进行拉伸；
	// * 当关闭视频分辨率自适应（As 取值为 0）时，Width 和 Height 任一取值为 0 时，转码视频将保持源流尺寸；
	// * 编码格式为 H.266 时，不支持设置 Width 和 Height，请使用自适应配置。 :::
	Height *int32 `json:"Height,omitempty"`

	// 长边长度。 :::tip
	// * 当 As 的取值为 1 即开启宽高自适应时，参数生效，反之则不生效。
	// * 当 As 的取值为 1 时，如果 LongSide 、 ShortSide 、Width 、Height 同时取 0，表示保持源流尺寸。 :::
	LongSide     *int32  `json:"LongSide,omitempty"`
	LookAhead    *int32  `json:"LookAhead,omitempty"`
	Modifier     *string `json:"Modifier,omitempty"`
	NvBf         *int32  `json:"NvBf,omitempty"`
	NvCodec      *string `json:"NvCodec,omitempty"`
	NvGop        *int32  `json:"NvGop,omitempty"`
	NvHVSPre     *bool   `json:"NvHVSPre,omitempty"`
	NvLookahead  *int32  `json:"NvLookahead,omitempty"`
	NvPercent    *int32  `json:"NvPercent,omitempty"`
	NvPreset     *string `json:"NvPreset,omitempty"`
	NvPriority   *int32  `json:"NvPriority,omitempty"`
	NvProfile    *string `json:"NvProfile,omitempty"`
	NvRefs       *int32  `json:"NvRefs,omitempty"`
	NvTempAQ     *int32  `json:"NvTempAQ,omitempty"`
	Ocr          *bool   `json:"Ocr,omitempty"`
	Preset       *string `json:"Preset,omitempty"`
	PresetKind   *int32  `json:"PresetKind,omitempty"`
	PresetType   *int32  `json:"PresetType,omitempty"`
	Qp           *int32  `json:"Qp,omitempty"`
	RegionConfig *string `json:"RegionConfig,omitempty"`
	Revision     *string `json:"Revision,omitempty"`

	// 转码类型是否为极智超清转码，默认值为 false，取值及含义如下。
	// * true：极智超清转码；
	// * false：标准转码。
	// :::tip 视频编码格式为 H.266 （Vcodec取值为h266）时，转码类型不支持极智超清转码。 :::
	Roi  *bool `json:"Roi,omitempty"`
	SITI *bool `json:"SITI,omitempty"`

	// 短边长度。 :::tip
	// * 当 As 的取值为 1 即开启宽高自适应时，参数生效，反之则不生效。
	// * 当 As 的取值为 1 时，如果 LongSide 、 ShortSide 、Width 、Height 同时取 0，表示保持源流尺寸。 :::
	ShortSide    *int32 `json:"ShortSide,omitempty"`
	Status       *int32 `json:"Status,omitempty"`
	StopInterval *int32 `json:"StopInterval,omitempty"`

	// 转码后缀，支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）和短横线（-）组成，长度为 1 到 10 个字符。
	// 转码后缀通常以流名称后缀的形式来使用，常见的标识有 _sd、_hd、_uhd，例如，当转码配置的标识为 _hd 时，拉取转码流时转码流的流名名称为 源流的流名称_hd。
	SuffixName   *string `json:"SuffixName,omitempty"`
	Threads      *int32  `json:"Threads,omitempty"`
	VBRatio      *int32  `json:"VBRatio,omitempty"`
	VBVBufSize   *int32  `json:"VBVBufSize,omitempty"`
	VBVMaxRate   *int32  `json:"VBVMaxRate,omitempty"`
	VLevel       *string `json:"VLevel,omitempty"`
	VPreset      *string `json:"VPreset,omitempty"`
	VProfile     *string `json:"VProfile,omitempty"`
	VR           *int32  `json:"VRVr,omitempty"`
	VRBBframes   *int32  `json:"VRBBframes,omitempty"`
	VRBHeightNum *int32  `json:"VRBHeightNum,omitempty"`
	VRBPreset    *string `json:"VRBPreset,omitempty"`
	VRBProfile   *string `json:"VRBProfile,omitempty"`
	VRBSuffix    *string `json:"VRBSuffix,omitempty"`
	VRBVb        *int32  `json:"VRBVb,omitempty"`
	VRBWidthNum  *int32  `json:"VRBWidthNum,omitempty"`
	VRGop        *int32  `json:"VRGop,omitempty"`
	VRGopDen     *int32  `json:"VRGopDen,omitempty"`
	VRHvspre     *string `json:"VRHvspre,omitempty"`
	VRProjection *string `json:"VRProjection,omitempty"`
	VRRoi        *string `json:"VRRoi,omitempty"`
	VRTBframes   *int32  `json:"VRTBframes,omitempty"`
	VRTPreset    *string `json:"VRTPreset,omitempty"`
	VRTProfile   *string `json:"VRTProfile,omitempty"`
	VRTSuffix    *string `json:"VRTSuffix,omitempty"`
	VRTVb        *int32  `json:"VRTVb,omitempty"`
	VRTileMod    *int32  `json:"VRTileMod,omitempty"`
	VRateCtrl    *string `json:"VRateCtrl,omitempty"`
	VbThreshold  *string `json:"VbThreshold,omitempty"`
	Vclass       *bool   `json:"Vclass,omitempty"`

	// 视频编码格式。
	// * h264：使用 H.264 编码格式；
	// * h265：使用 H.265 编码格式；
	// * copy：不进行转码，所有视频编码参数不生效。
	Vcodec *string `json:"Vcodec,omitempty"`

	// 视频码率，单位为 kbps。
	VideoBitrate *int32  `json:"VideoBitrate,omitempty"`
	Vn           *int32  `json:"Vn,omitempty"`
	Watermark    *string `json:"Watermark,omitempty"`

	// 视频宽度。 :::tip
	// * 当关闭视频分辨率自适应（As 取值为 0）时，转码分辨率将取 Width 和 Height 的值对转码视频进行拉伸；
	// * 当关闭视频分辨率自适应（As 取值为 0）时，Width 和 Height 任一取值为 0 时，转码视频将保持源流尺寸；
	// * 编码格式为 H.266 时，不支持设置 Width 和 Height，请使用自适应配置。 :::
	Width *int32 `json:"Width,omitempty"`
}

type ListCommonTransPresetDetailResResultStandardPresetDetailItem struct {
	ALayout   *string `json:"ALayout,omitempty"`
	AProfile  *string `json:"AProfile,omitempty"`
	AR        *int32  `json:"AR,omitempty"`
	AbrMode   *int32  `json:"AbrMode,omitempty"`
	AccountID *string `json:"AccountID,omitempty"`

	// 音频编码格式。包括以下 3 种类型。
	// * aac：使用 aac 编码格式；
	// * copy：不进行转码，所有音频编码参数不生效；
	// * opus：使用 opus 编码格式。
	Acodec         *string `json:"Acodec,omitempty"`
	AdvancedParam  *string `json:"AdvancedParam,omitempty"`
	AllowAudioCopy *int32  `json:"AllowAudioCopy,omitempty"`
	AllowVideoCopy *int32  `json:"AllowVideoCopy,omitempty"`
	An             *int32  `json:"An,omitempty"`

	// 宽高自适应模式开关。
	// * 0：关闭宽高自适应，按照Width和Height的取值进行拉伸；
	// * 1：开启宽高自适应，按照ShortSide或LongSide等比缩放。
	As *string `json:"As,omitempty"`

	// 音频码率，单位为 kbps
	AudioBitrate        *int32 `json:"AudioBitrate,omitempty"`
	AutoTransAb         *int32 `json:"AutoTransAb,omitempty"`
	AutoTransAl         *int32 `json:"AutoTransAl,omitempty"`
	AutoTransAr         *int32 `json:"AutoTransAr,omitempty"`
	AutoTransResolution *int32 `json:"AutoTransResolution,omitempty"`
	AutoTransVb         *int32 `json:"AutoTransVb,omitempty"`
	AutoTransVr         *int32 `json:"AutoTransVr,omitempty"`
	BCM                 *int32 `json:"BCM,omitempty"`

	// 2 个参考帧之间的最大 B 帧数。BFrames取 0 时，表示去 B 帧。
	BFrames  *int32  `json:"BFrames,omitempty"`
	Describe *string `json:"Describe,omitempty"`

	// 帧率，单位为 fps。帧率越大，画面越流畅
	FPS *int32 `json:"FPS,omitempty"`

	// IDR 帧之间的最大间隔，单位为。
	GOP    *int32 `json:"GOP,omitempty"`
	GopMin *int32 `json:"GopMin,omitempty"`
	HVSPre *bool  `json:"HVSPre,omitempty"`

	// 视频高度。 :::tip 当 As 的取值为 0 时，Width 和 Height 中任意参数取 0，表示保持源流尺寸。 :::
	Height *int32 `json:"Height,omitempty"`

	// 长边长度。 :::tip 当 As 的取值为 1 时，如果 LongSide 和 ShortSide 都取 0，表示保持源流尺寸。 :::
	LongSide     *int32  `json:"LongSide,omitempty"`
	LookAhead    *int32  `json:"LookAhead,omitempty"`
	Modifier     *string `json:"Modifier,omitempty"`
	NvBf         *int32  `json:"NvBf,omitempty"`
	NvCodec      *string `json:"NvCodec,omitempty"`
	NvGop        *int32  `json:"NvGop,omitempty"`
	NvHVSPre     *bool   `json:"NvHVSPre,omitempty"`
	NvLookahead  *int32  `json:"NvLookahead,omitempty"`
	NvPercent    *int32  `json:"NvPercent,omitempty"`
	NvPreset     *string `json:"NvPreset,omitempty"`
	NvPriority   *int32  `json:"NvPriority,omitempty"`
	NvProfile    *string `json:"NvProfile,omitempty"`
	NvRefs       *int32  `json:"NvRefs,omitempty"`
	NvTempAQ     *int32  `json:"NvTempAQ,omitempty"`
	Ocr          *bool   `json:"Ocr,omitempty"`
	Preset       *string `json:"Preset,omitempty"`
	PresetKind   *int32  `json:"PresetKind,omitempty"`
	PresetType   *int32  `json:"PresetType,omitempty"`
	Qp           *int32  `json:"Qp,omitempty"`
	RegionConfig *string `json:"RegionConfig,omitempty"`
	Revision     *string `json:"Revision,omitempty"`

	// 是否极智超清转码。
	// * true：极智超清；
	// * false：标准转码。
	Roi  *bool `json:"Roi,omitempty"`
	SITI *bool `json:"SITI,omitempty"`

	// 短边长度。 :::tip 当 As 的取值为 1 时，如果 LongSide 和 ShortSide 都取 0，表示保持源流尺寸。 :::
	ShortSide    *int32 `json:"ShortSide,omitempty"`
	Status       *int32 `json:"Status,omitempty"`
	StopInterval *int32 `json:"StopInterval,omitempty"`

	// 转码流后缀名
	SuffixName   *string `json:"SuffixName,omitempty"`
	Threads      *int32  `json:"Threads,omitempty"`
	VBRatio      *int32  `json:"VBRatio,omitempty"`
	VBVBufSize   *int32  `json:"VBVBufSize,omitempty"`
	VBVMaxRate   *int32  `json:"VBVMaxRate,omitempty"`
	VLevel       *string `json:"VLevel,omitempty"`
	VPreset      *string `json:"VPreset,omitempty"`
	VProfile     *string `json:"VProfile,omitempty"`
	VR           *int32  `json:"VRVr,omitempty"`
	VRBBframes   *int32  `json:"VRBBframes,omitempty"`
	VRBHeightNum *int32  `json:"VRBHeightNum,omitempty"`
	VRBPreset    *string `json:"VRBPreset,omitempty"`
	VRBProfile   *string `json:"VRBProfile,omitempty"`
	VRBSuffix    *string `json:"VRBSuffix,omitempty"`
	VRBVb        *int32  `json:"VRBVb,omitempty"`
	VRBWidthNum  *int32  `json:"VRBWidthNum,omitempty"`
	VRGop        *int32  `json:"VRGop,omitempty"`
	VRGopDen     *int32  `json:"VRGopDen,omitempty"`
	VRHvspre     *string `json:"VRHvspre,omitempty"`
	VRProjection *string `json:"VRProjection,omitempty"`
	VRRoi        *string `json:"VRRoi,omitempty"`
	VRTBframes   *int32  `json:"VRTBframes,omitempty"`
	VRTPreset    *string `json:"VRTPreset,omitempty"`
	VRTProfile   *string `json:"VRTProfile,omitempty"`
	VRTSuffix    *string `json:"VRTSuffix,omitempty"`
	VRTVb        *int32  `json:"VRTVb,omitempty"`
	VRTileMod    *int32  `json:"VRTileMod,omitempty"`
	VRateCtrl    *string `json:"VRateCtrl,omitempty"`
	VbThreshold  *string `json:"VbThreshold,omitempty"`
	Vclass       *bool   `json:"Vclass,omitempty"`

	// 视频编码格式。
	// * h264：使用 H.264 编码格式；
	// * h265：使用 H.265 编码格式；
	// * copy：不进行转码，所有视频编码参数不生效。
	Vcodec *string `json:"Vcodec,omitempty"`

	// 视频码率，单位为 kbps
	VideoBitrate *int32  `json:"VideoBitrate,omitempty"`
	Vn           *int32  `json:"Vn,omitempty"`
	Watermark    *string `json:"Watermark,omitempty"`

	// 视频宽度。 :::tip 当 As 的取值为 0 时，如果 Width 和 Height 中任意参数取 0，表示保持源流尺寸。 :::
	Width *int32 `json:"Width,omitempty"`
}

type ListDomainDetailBody struct {

	// REQUIRED; 查询数据的页码，取值为 1 时表示查询第一页的数据，取值范围为 [1,1000]。
	PageNum int32 `json:"PageNum"`

	// REQUIRED; 每页显示的数据条数，取值为 10 时表示每页展示 10 条域名信息，取值范围为 [1, 1000]。
	PageSize int32 `json:"PageSize"`

	// 域名名称列表，缺省情况下表示全部。
	DomainNameList []*string `json:"DomainNameList,omitempty"`

	// 域名加速区域列表，缺省情况下表示查看全部。支持的取值如下所示。
	// * cn：中国内地；
	// * cn-global：全球加速；
	// * cn-oversea：海外及港澳台。
	DomainRegionList []*string `json:"DomainRegionList,omitempty"`

	// 域名状态列表，缺省情况下表示查询全部状态的域名。支持的取值如下所示。
	// * 0：正常；
	// * 1：审核中；
	// * 2：禁用，禁止使用，此时 domain 不生效；
	// * 3：删除；
	// * 4：审核被驳回。审核不通过，需要重新创建并审核；
	// * 5：欠费关停。
	DomainStatusList []*int32 `json:"DomainStatusList,omitempty"`

	// 域名类型列表，缺省情况下表示全部类型的域名。支持的取值如下所示。
	// * push：推流域名；
	// * pull-flv：拉流域名。
	DomainTypeList []*string `json:"DomainTypeList,omitempty"`

	// 域名空间列表，缺省情况下表示查询全部域名空间下的域名。
	VhostList []*string `json:"VhostList,omitempty"`
}

type ListDomainDetailRes struct {

	// REQUIRED
	ResponseMetadata ListDomainDetailResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListDomainDetailResResult          `json:"Result,omitempty"`
}

type ListDomainDetailResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                    `json:"Version"`
	Error   *ListDomainDetailResResponseMetadataError `json:"Error,omitempty"`
}

type ListDomainDetailResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type ListDomainDetailResResult struct {

	// REQUIRED; 总记录数。
	Total int32 `json:"Total"`

	// 域名详细信息列表。
	DomainList []*ListDomainDetailResResultDomainListItem `json:"DomainList,omitempty"`
}

type ListDomainDetailResResultDomainListItem struct {

	// REQUIRED; CNAME 信息。
	CNAME string `json:"CNAME"`

	// REQUIRED; 绑定的 HTTPS 证书支持的泛域名。
	CertDomain string `json:"CertDomain"`

	// REQUIRED; 绑定的 HTTPS 证书的证书链 ID 信息。
	ChainID string `json:"ChainID"`

	// REQUIRED; CNAME 状态，取值及含义如下所示。
	// * 0：未配置 CNAME；
	// * 1：已配置 CNAME。
	CnameCheck int32 `json:"CnameCheck"`

	// REQUIRED; 域名添加时间，RFC3339 格式的 UTC 时间戳，精度为秒。
	CreateTime string `json:"CreateTime"`

	// REQUIRED; 推/拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 域名是否可用的状态，取值及含义如下所示。
	// * 0：正常，域名为可用状态；
	// * 1：配置中，域名为可用状态；
	// * 2：不可用，域名为其他的不可用状态。
	DomainCheck int32 `json:"DomainCheck"`

	// REQUIRED; ICP 备案校验是否通过，是否过期信息。
	// * 1：备案正常，未过期；
	// * 2：查存不到备案信息。
	ICPCheck int32 `json:"ICPCheck"`

	// REQUIRED; 域名空间所属的项目名称。
	ProjectName string `json:"ProjectName"`

	// REQUIRED; 绑定的推流域名。
	PushDomain string `json:"PushDomain"`

	// REQUIRED; 域名加速区域，取值及含义如下所示。
	// * cn：中国大陆；
	// * cn-global：全球；
	// * cn-oversea：海外及港澳台。
	Region string `json:"Region"`

	// REQUIRED; 域名状态，取值及含义如下所示。
	// * 0：正常；
	// * 1：审核中；
	// * 2：禁用，禁止使用，此时 domain 不生效；
	// * 3：删除；
	// * 4：审核被驳回。审核不通过，需要重新创建并审核；
	// * 5：欠费关停。
	Status int32 `json:"Status"`

	// REQUIRED; 域名空间的标签信息。
	Tags []ListDomainDetailResResultDomainListPropertiesItemsItem `json:"Tags"`

	// REQUIRED; 域名类型，取值及含义如下所示。
	// * push：推流域名；
	// * pull-flv：拉流域名，包含 RTMP、FLV、HLS 格式。
	Type string `json:"Type"`

	// REQUIRED; 域名空间。
	Vhost string `json:"Vhost"`
}

type ListDomainDetailResResultDomainListPropertiesItemsItem struct {

	// REQUIRED; 标签类型，取值及含义如下所示。
	// * System：系统内置标签；
	// * Custom：自定义标签。
	Category string `json:"Category"`

	// REQUIRED; 标签 Key 值。
	Key string `json:"Key"`

	// REQUIRED; 标签 Value 值。
	Value string `json:"Value"`
}

type ListPullToPushTaskQuery struct {

	// 查询数据的页码，默认为 1，表示查询第一页的数据。
	Page *int32 `json:"Page,omitempty" query:"Page"`

	// 每页显示的数据条数，默认为 20，最大值为 500。
	Size *int32 `json:"Size,omitempty" query:"Size"`

	// 拉流转推任务的名称，不区分大小写，支持模糊查询。 例如，title取值为doc时，则返回任务名称为docspace、docs、DOC等 title 中包含doc关键词的所有任务列表。
	Title *string `json:"Title,omitempty" query:"Title"`
}

type ListPullToPushTaskRes struct {

	// REQUIRED
	ResponseMetadata ListPullToPushTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result ListPullToPushTaskResResult `json:"Result"`
}

type ListPullToPushTaskResResponseMetadata struct {
	Action    *string                                     `json:"Action,omitempty"`
	Error     *ListPullToPushTaskResResponseMetadataError `json:"Error,omitempty"`
	Region    *string                                     `json:"Region,omitempty"`
	RequestID *string                                     `json:"RequestId,omitempty"`
	Service   *string                                     `json:"Service,omitempty"`
	Version   *string                                     `json:"Version,omitempty"`
}

type ListPullToPushTaskResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type ListPullToPushTaskResResult struct {

	// 任务列表。
	List []*ListPullToPushTaskResResultListItem `json:"List,omitempty"`

	// 分页数量信息。
	Pagination *ListPullToPushTaskResResultPagination `json:"Pagination,omitempty"`
}

type ListPullToPushTaskResResultListItem struct {

	// 接收拉流转推任务状态回调的地址。
	CallbackURL *string `json:"CallbackURL,omitempty"`

	// 续播策略，续播策略指转推点播视频进行直播时出现断流并恢复后，如何继续播放的策略，拉流来源类型为点播视频时参数生效，支持的取值及含义如下。
	// * 0：从断流处续播（默认值）；
	// * 1：从断流处+自然流逝时长处续播。
	ContinueStrategy *int32 `json:"ContinueStrategy,omitempty"`

	// 点播视频文件循环播放模式，当拉流来源类型为点播视频时配置生效，参数取值及含义如下所示。
	// * -1：无限次循环，至任务结束；
	// * 0：有限次循环，循环次数以 PlayTimes 取值为准；
	// * >0：有限次循环，循环次数以 CycleMode 取值为准。
	CycleMode *int32 `json:"CycleMode,omitempty"`

	// 推流地址，即直播源或点播视频转推的目标地址。
	DstAddr *string `json:"DstAddr,omitempty"`

	// 推流地址类型。
	// * 1：非第三方，即推流地址域名已添加到视频直播。
	// * 2：第三方，即推流地址域名未添加到视频直播。
	DstAddrType *int32 `json:"DstAddrType,omitempty"`

	// 任务的结束时间，RFC3339 格式的 UTC 时间，单位为秒。
	EndTime *string `json:"EndTime,omitempty"`

	// 点播文件启播时间偏移值，单位为秒，数量与拉流地址列表中地址数量相等，缺省情况下为空表示不进行偏移。拉流来源类型为点播视频时，参数生效。
	OffsetS []*float32 `json:"OffsetS,omitempty"`

	// 点播视频文件循环播放次数，当 CycleMode 取值为 0 时，PlayTimes 取值将作为循环播放次数。
	PlayTimes *int32 `json:"PlayTimes,omitempty"`

	// 是否开启点播预热，开启点播预热后，系统会自动将点播视频文件缓存到 CDN 节点上，当用户请求直播时，可以直播从 CDN 节点获取视频，从而提高直播流畅度。拉流来源类型为点播视频时，参数生效。
	// * 0：不开启；
	// * 1：开启。
	PreDownload *int32 `json:"PreDownload,omitempty"`

	// 直播源的拉流地址，拉流来源类型为直播源时返回此值。
	SrcAddr *string `json:"SrcAddr,omitempty"`

	// 点播视频播放地址列表，拉流来源类型为点播视频时返回此值。
	SrcAddrS []*string `json:"SrcAddrS,omitempty"`

	// 任务的开始时间，RFC3339 格式的 UTC 时间，单位为秒。
	StartTime *string `json:"StartTime,omitempty"`

	// 拉流转推任务的状态，支持如下取值。
	// * 停用；
	// * 未开始；
	// * 生效中；
	// * 已结束。
	Status *string `json:"Status,omitempty"`

	// 任务 ID，任务的唯一标识。
	TaskID *string `json:"TaskId,omitempty"`

	// 拉流转推任务的名称。
	Title *string `json:"Title,omitempty"`

	// 拉流来源类型，支持的取值及含义如下。
	// * 0：直播源；
	// * 1：点播视频。
	Type *int32 `json:"Type,omitempty"`

	// 为拉流转推视频添加的水印配置信息。
	Watermark *ListPullToPushTaskResResultListItemWatermark `json:"Watermark,omitempty"`
}

// ListPullToPushTaskResResultListItemWatermark - 为拉流转推视频添加的水印配置信息。
type ListPullToPushTaskResResultListItemWatermark struct {

	// REQUIRED; 水印图片字符串，图片最大 2MB，最小 100Bytes，最大分辨率为 1080×1080。图片 Data URL 格式为：data:image/<mediatype>;base64,<data>。
	// * mediatype：图片类型，支持 png、jpg、jpeg 格式；
	// * data：base64 编码的图片字符串。
	// 例如，data:image/png;base64,iVBORw0KGg****mCC
	Picture string `json:"Picture"`

	// REQUIRED; 水印宽度占直播原始画面宽度百分比，支持精度为小数点后两位。
	Ratio float32 `json:"Ratio"`

	// REQUIRED; 水平偏移，表示水印左侧边与转码流画面左侧边之间的距离，使用相对比率，取值范围为 [0,1)。
	RelativePosX float32 `json:"RelativePosX"`

	// REQUIRED; 垂直偏移，表示水印顶部边与转码流画面顶部边之间的距离，使用相对比率，取值范围为 [0,1)。
	RelativePosY float32 `json:"RelativePosY"`
}

// ListPullToPushTaskResResultPagination - 分页数量信息。
type ListPullToPushTaskResResultPagination struct {

	// 当前任务所在分页。
	PageCur *int32 `json:"PageCur,omitempty"`

	// 每页显示的数据条数。
	PageSize *int32 `json:"PageSize,omitempty"`

	// 查询结果的数据总页数。
	PageTotal *int32 `json:"PageTotal,omitempty"`

	// 查询结果的数据总条数。
	TotalCount *int32 `json:"TotalCount,omitempty"`
}

type ListRelaySourceV4Body struct {

	// REQUIRED; 直播流使用的域名。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名。
	Domain string `json:"Domain"`

	// 查询数据的页码，默认为 1，表示查询第一页的数据。
	Page *int32 `json:"Page,omitempty"`

	// 每页显示的数据条数，默认为 20，最大值为 500。
	Size *int32 `json:"Size,omitempty"`
}

type ListRelaySourceV4Res struct {

	// REQUIRED
	ResponseMetadata ListRelaySourceV4ResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListRelaySourceV4ResResult          `json:"Result,omitempty"`
}

type ListRelaySourceV4ResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                     `json:"Version"`
	Error   *ListRelaySourceV4ResResponseMetadataError `json:"Error,omitempty"`
}

type ListRelaySourceV4ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type ListRelaySourceV4ResResult struct {

	// REQUIRED; 配置列表。
	List []ListRelaySourceV4ResResultListItem `json:"List"`

	// REQUIRED; 页码信息。
	Pagination ListRelaySourceV4ResResultPagination `json:"Pagination"`
}

type ListRelaySourceV4ResResultListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 直播流的使用的域名。
	Domain string `json:"Domain"`

	// REQUIRED; 回源结束时间，StartTime 和 EndTime 同时缺省的情况下，表示永久回源。
	EndTime int32 `json:"EndTime"`

	// REQUIRED; 回源地址列表。
	SrcAddrS []string `json:"SrcAddrS"`

	// REQUIRED; 回源开始时间，StartTime 和 EndTime 同时缺省的情况下，表示永久回源。
	StartTime int32 `json:"StartTime"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`
}

// ListRelaySourceV4ResResultPagination - 页码信息。
type ListRelaySourceV4ResResultPagination struct {

	// REQUIRED; 当前查询的页码。
	PageCur int32 `json:"PageCur"`

	// REQUIRED; 每页显示的数据条数。
	PageSize int32 `json:"PageSize"`

	// REQUIRED; 查询结果的数据总页数。
	PageTotal int32 `json:"PageTotal"`

	// REQUIRED; 查询结果的数据总条数。
	TotalCount int32 `json:"TotalCount"`
}

type ListTimeShiftPresetV2Body struct {

	// 时移类型，默认类型为 vod。
	// * vod：点播时移，表示查询时移录制存储在 VOD 中的时移配置；
	// * tos：直播时移，表示查询时移录制存储在 TOS 以及 fcdn-tos 中的时移配置。
	Type *string `json:"Type,omitempty"`

	// 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看需要时移的直播流使用的域名所属的域名空间。
	Vhost *string `json:"Vhost,omitempty"`
}

type ListTimeShiftPresetV2Res struct {

	// REQUIRED
	ResponseMetadata ListTimeShiftPresetV2ResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListTimeShiftPresetV2ResResult          `json:"Result,omitempty"`
}

type ListTimeShiftPresetV2ResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                         `json:"Version"`
	Error   *ListTimeShiftPresetV2ResResponseMetadataError `json:"Error,omitempty"`
}

type ListTimeShiftPresetV2ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type ListTimeShiftPresetV2ResResult struct {

	// 时移配置列表。
	List []*ListTimeShiftPresetV2ResResultListItem `json:"List,omitempty"`
}

type ListTimeShiftPresetV2ResResultListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; TOS 存储对应的 Bucket。
	Bucket string `json:"Bucket"`

	// REQUIRED; 拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 最大时移时长，即允许用户回看的最长时间，单位为秒。
	MaxShiftTime int32 `json:"MaxShiftTime"`

	// REQUIRED; 时移配置名称。
	Name string `json:"Name"`

	// REQUIRED; 直播时移配置启用状态。
	// * 0：配置中；
	// * 1：已启用。
	Status int32 `json:"Status"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 类型。默认类型为 vod。
	// * vod：录制类型为录制时移时，录制配置中存储位置为 VOD。
	// * tos：录制类型为录制时移时，录制配置中存储喂食为 TOS。
	// * fcdn-toS：独立时移。
	Type string `json:"Type"`

	// REQUIRED; 视频点播（VOD）空间名称。
	VODNamespace string `json:"VODNamespace"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`
}

type ListVhostRecordPresetV2Body struct {

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，查看需要录制的直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 直播录制的存储类型，默认值为 tos，支持的取值及含义如下所示。
	// * vod：录制文件存在 VOD；
	// * tos：录制文件存在 TOS。
	Type *string `json:"Type,omitempty"`
}

type ListVhostRecordPresetV2Res struct {

	// REQUIRED
	ResponseMetadata ListVhostRecordPresetV2ResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListVhostRecordPresetV2ResResult          `json:"Result,omitempty"`
}

type ListVhostRecordPresetV2ResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                           `json:"Version"`
	Error   *ListVhostRecordPresetV2ResResponseMetadataError `json:"Error,omitempty"`
}

type ListVhostRecordPresetV2ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type ListVhostRecordPresetV2ResResult struct {

	// REQUIRED; 录制配置列表。
	PresetList []ListVhostRecordPresetV2ResResultPresetListItem `json:"PresetList"`
}

type ListVhostRecordPresetV2ResResultPresetListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 录制参数详细信息。
	SlicePresetV2 *ListVhostRecordPresetV2ResResultPresetListItemSlicePresetV2 `json:"SlicePresetV2,omitempty"`
}

// ListVhostRecordPresetV2ResResultPresetListItemSlicePresetV2 - 录制参数详细信息。
type ListVhostRecordPresetV2ResResultPresetListItemSlicePresetV2 struct {

	// 录制配置 ID。
	ID *int32 `json:"ID,omitempty"`

	// 录制配置名称。
	Name *string `json:"Name,omitempty"`

	// 录制模板详细配置。
	RecordPresetConfig *ComponentsFuamuzSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfig `json:"RecordPresetConfig,omitempty"`
}

type ListVhostSnapshotPresetBody struct {

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 截图存储类型。
	// * tos；
	// * imageX。
	Type *string `json:"Type,omitempty"`
}

type ListVhostSnapshotPresetRes struct {

	// REQUIRED
	ResponseMetadata ListVhostSnapshotPresetResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListVhostSnapshotPresetResResult          `json:"Result,omitempty"`
}

type ListVhostSnapshotPresetResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                           `json:"Version"`
	Error   *ListVhostSnapshotPresetResResponseMetadataError `json:"Error,omitempty"`
}

type ListVhostSnapshotPresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type ListVhostSnapshotPresetResResult struct {

	// 模版列表。
	PresetList []*ListVhostSnapshotPresetResResultPresetListItem `json:"PresetList,omitempty"`
}

type ListVhostSnapshotPresetResResultPresetListItem struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 截图模板名称。
	SlicePreset *ListVhostSnapshotPresetResResultPresetListItemSlicePreset `json:"SlicePreset,omitempty"`
}

// ListVhostSnapshotPresetResResultPresetListItemSlicePreset - 截图模板名称。
type ListVhostSnapshotPresetResResultPresetListItemSlicePreset struct {

	// 截图在 ToS 中的存储位置。
	Bucket *string `json:"Bucket,omitempty"`

	// 回调信息。
	CallbackDetail *Components1Hkcrc4SchemasListvhostsnapshotpresetresPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetPropertiesCallbackdetail `json:"CallbackDetail,omitempty"`

	// 截图间隔时间。
	Interval *int32 `json:"Interval,omitempty"`

	// 截图模版名称。
	Preset *string `json:"Preset,omitempty"`

	// veImageX 的服务 ID。
	ServiceID *string `json:"ServiceID,omitempty"`

	// 截图模版状态。
	// * 1：开启
	// * 0：关闭
	Status *int32 `json:"Status,omitempty"`
}

type ListVhostTransCodePresetBody struct {

	// REQUIRED; 是否是hls abr 请求
	IsHlsAbr bool `json:"IsHlsAbr"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，查看需要录制的直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`
}

type ListVhostTransCodePresetRes struct {

	// REQUIRED
	ResponseMetadata ListVhostTransCodePresetResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListVhostTransCodePresetResResult          `json:"Result,omitempty"`
}

type ListVhostTransCodePresetResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version   string                                            `json:"Version"`
	Error     *ListVhostTransCodePresetResResponseMetadataError `json:"Error,omitempty"`
	RequestID *string                                           `json:"RequestID,omitempty"`
}

type ListVhostTransCodePresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type ListVhostTransCodePresetResResult struct {

	// REQUIRED; 全部转码配置列表。
	AllPresetList []ListVhostTransCodePresetResResultAllPresetListItem `json:"AllPresetList"`

	// REQUIRED; 使用内置参数的转码配置列表。
	CommonPresetList []ListVhostTransCodePresetResResultCommonPresetListItem `json:"CommonPresetList"`

	// REQUIRED; 使用自定义配置的转码配置列表。
	CustomizePresetList []ListVhostTransCodePresetResResultCustomizePresetListItem `json:"CustomizePresetList"`
}

type ListVhostTransCodePresetResResultAllPresetListItem struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 域名空间。
	Vhost string `json:"Vhost"`

	// 转码配置具体信息。
	TranscodePreset *ListVhostTransCodePresetResResultAllPresetListItemTranscodePreset `json:"TranscodePreset,omitempty"`
}

// ListVhostTransCodePresetResResultAllPresetListItemTranscodePreset - 转码配置具体信息。
type ListVhostTransCodePresetResResultAllPresetListItemTranscodePreset struct {

	// 音频编码格式。包括以下 3 种类型。
	// * aac：使用 aac 编码格式；
	// * copy：不进行转码，所有音频编码参数不生效；
	// * opus：使用 opus 编码格式。
	Acodec *string `json:"Acodec,omitempty"`

	// 宽高自适应模式开关。
	// * 0：关闭宽高自适应，按照Width和Height的取值进行拉伸；
	// * 1：开启宽高自适应，按照ShortSide或LongSide等比缩放。
	As *string `json:"As,omitempty"`

	// 音频码率，单位为 kbps。
	AudioBitrate *int32 `json:"AudioBitrate,omitempty"`

	// 是否开启转码不超过源流分辨率。开启后，当源流分辨率低于转码配置分辨率时(即源流宽低于转码配置宽且源流高低于转码配置高时)，将按源流视频分辨率进行转码，默认开启。
	// * 0：关闭
	// * 1：开启
	AutoTransResolution *int32 `json:"AutoTransResolution,omitempty"`

	// 是否开启不超过源流码率。开启后，当源流码率低于转码配置码率时，将按照源流视频码率进行转码，默认开启。
	// * 0：关闭
	// * 1：开启
	AutoTransVb *int32 `json:"AutoTransVb,omitempty"`

	// 是否开启不超过源流帧率。开启后，当源流帧率低于转码配置帧率时，将按照源流视频帧率进行转码，默认开启。
	// * 0：关闭
	// * 1：开启
	AutoTransVr *int32 `json:"AutoTransVr,omitempty"`

	// 2 个参考帧之间的最大 B 帧数。BFrames取 0 时，表示去 B 帧。
	BFrames *int32 `json:"BFrames,omitempty"`

	// 动态范围，画质增强类型生效
	// * SDR：输出为SDR
	// * HDR：输出为HDR
	DynamicRange *string `json:"DynamicRange,omitempty"`

	// 是否开启智能插帧，只对画质增强类型生效
	// * 0：不开启
	// * 1：开启
	FISwitch *int64 `json:"FISwitch,omitempty"`

	// 视频帧率，单位为 fps，帧率越大，画面越流畅。
	FPS *int32 `json:"FPS,omitempty"`

	// IDR 帧之间的最大间隔，单位为 s。
	GOP *int32 `json:"GOP,omitempty"`

	// 视频高度。
	Height *int32 `json:"Height,omitempty"`

	// 长边长度。 :::tip 当As的取值为 1 时，如果LongSide和ShortSide都取 0，表示保持源流尺寸。 :::
	LongSide *int32 `json:"LongSide,omitempty"`

	// 转码模板参数的类型
	// * va：表示使用画质增强
	ParamType *string `json:"ParamType,omitempty"`

	// 转码配置名称。
	Preset *string `json:"Preset,omitempty"`

	// 是否极智超清转码。
	// * true：极智超清；
	// * false：标准转码。
	Roi *bool `json:"Roi,omitempty"`

	// 使用场景，画质增强时生效 football：足球场景
	SceneType *string `json:"SceneType,omitempty"`

	// 短边长度。 :::tip 当As的取值为 1 时，如果LongSide和ShortSide都取 0，表示保持源流尺寸。 :::
	ShortSide *int32 `json:"ShortSide,omitempty"`

	// 转码停止时长，支持触发方式为拉流转码时设置，表示断开拉流后转码停止的时长，单位为 s，取值范围为 [0,300]，-1 表示不停止转码，默认值为 60。
	StopInterval *int32 `json:"StopInterval,omitempty"`

	// 转码流后缀名。
	SuffixName *string `json:"SuffixName,omitempty"`

	// 转码触发方式，默认为拉流转码，支持以下取值。
	// * Push：推流转码，直播推流后会自动启动转码任务，生成转码流；
	// * Pull：拉流转码，直播推流后，需要主动播放转码流才会启动转码任务，生成转码流。
	TransType *string `json:"TransType,omitempty"`

	// 视频编码格式。
	// * h264：使用 H.264 编码格式；
	// * h265：使用 H.265 编码格式；
	// * h266：使用 H.266 编码格式；
	// * copy：不进行转码，所有视频编码参数不生效。
	Vcodec *string `json:"Vcodec,omitempty"`

	// 视频码率，单位为 kbps。
	VideoBitrate *int32 `json:"VideoBitrate,omitempty"`

	// 视频宽度。
	Width *int32 `json:"Width,omitempty"`
}

type ListVhostTransCodePresetResResultCommonPresetListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 域名空间。
	Vhost string `json:"Vhost"`

	// 转码配置具体信息。
	TranscodePreset *ListVhostTransCodePresetResResultCommonPresetListItemTranscodePreset `json:"TranscodePreset,omitempty"`
}

// ListVhostTransCodePresetResResultCommonPresetListItemTranscodePreset - 转码配置具体信息。
type ListVhostTransCodePresetResResultCommonPresetListItemTranscodePreset struct {

	// 音频编码格式，默认值为 aac，支持的取值及含义如下所示。
	// * aac：使用 AAC 音频编码格式；
	// * opus：使用 Opus 音频编码格式。
	// * copy：不进行音频转码，所有音频编码参数不生效，音频编码参数包括音频码率（AudioBitrate）等。
	Acodec *string `json:"Acodec,omitempty"`

	// 视频分辨率自适应模式开关，默认值为 0。支持的取值及含义如下。
	// * 0：关闭视频分辨率自适应；
	// * 1：开启视频分辨率自适应。 :::tip
	// * 关闭视频分辨率自适应模式（As 取值为 0）时，转码配置的视频分辨率取视频宽度（Width）和视频高度（Height）的值对转码视频进行拉伸；
	// * 开启视频分辨率自适应模式（As 取值为 1）时，转码配置的视频分辨率按照短边长度（ShortSide）、长边长度（LongSide）、视频宽度（Width）、视频高度（Height）的优先级取值，另一边等比缩放。 :::
	As *string `json:"As,omitempty"`

	// 音频码率，单位为 kbps。
	AudioBitrate *int32 `json:"AudioBitrate,omitempty"`

	// 是否开启转码视频分辨率不超过源流分辨率，默认值为 1 表示开启。开启后，当源流分辨率低于转码配置分辨率时（即源流宽低于转码配置宽且源流高低于转码配置高时），将按源流视频分辨率进行转码。
	// * 0：关闭；
	// * 1：开启。
	AutoTransResolution *int32 `json:"AutoTransResolution,omitempty"`

	// 是否开启转码视频码率不超过源流码率，默认值为 1 表示开启。开启后，当源流码率低于转码配置码率时，将按照源流视频码率进行转码。
	// * 0：关闭；
	// * 1：开启。
	AutoTransVb *int32 `json:"AutoTransVb,omitempty"`

	// 是否开启转码视频帧率不超过源流帧率，默认值为 1 表示开启。开启后，当源流帧率低于转码配置帧率时，将按照源流视频帧率进行转码。
	// * 0：关闭；
	// * 1：开启。
	AutoTransVr *int32 `json:"AutoTransVr,omitempty"`

	// 转码输出视频中 2 个参考帧之间的最大 B 帧数量，取值为 0 时表示去除 B 帧。
	BFrames *int32 `json:"BFrames,omitempty"`

	// 视频帧率，单位为 fps，帧率越大，画面越流畅。
	FPS *int32 `json:"FPS,omitempty"`

	// IDR 帧之间的最大间隔，单位为秒。
	GOP *int32 `json:"GOP,omitempty"`

	// 视频高度。
	// :::tip
	// * 当关闭视频分辨率自适应（As 取值为 0）时，转码分辨率将取 Width 和 Height 的值对转码视频进行拉伸；
	// * 当关闭视频分辨率自适应（As 取值为 0）时，Width 和 Height 任一取值为 0 时，转码视频将保持源流尺寸。 :::
	Height *int32 `json:"Height,omitempty"`

	// 长边长度。
	// :::tip
	// * 当开启视频分辨率自适应模式时（As 取值为 1）时，参数生效，反之则不生效。
	// * 当开启视频分辨率自适应模式时（As 取值为 1）时，如果 LongSide 、 ShortSide 、Width 、Height 同时取 0，表示保持源流尺寸。 :::
	LongSide *int32 `json:"LongSide,omitempty"`

	// 转码配置名称。
	Preset *string `json:"Preset,omitempty"`

	// 转码类型是否为极智超清转码，默认值为 false，取值及含义如下。
	// * true：极智超清转码；
	// * false：标准转码。
	// :::tip 视频编码格式为 H.266 （Vcodec取值为h266）时，转码类型不支持极智超清转码。 :::
	Roi *bool `json:"Roi,omitempty"`

	// 短边长度。 :::tip
	// * 当 As 的取值为 1 即开启宽高自适应时，参数生效，反之则不生效。
	// * 当 As 的取值为 1 时，如果 LongSide 、 ShortSide 、Width 、Height 同时取 0，表示保持源流尺寸。 :::
	ShortSide *int32 `json:"ShortSide,omitempty"`

	// 转码停止时长，支持触发方式为拉流转码（TransType 取值为 Pull）时设置，表示断开拉流后转码停止的时长，单位为秒，取值范围为 -1 和 [0,300]，-1 表示不停止转码，默认值为 60。
	StopInterval *int32 `json:"StopInterval,omitempty"`

	// 转码流后缀名。
	SuffixName *string `json:"SuffixName,omitempty"`

	// 转码触发方式，支持的取值及含义如下。
	// * Push：推流转码，直播推流后会自动启动转码任务，生成转码流；
	// * Pull：拉流转码，直播推流后，需要主动播放转码流才会启动转码任务，生成转码流。
	TransType *string `json:"TransType,omitempty"`

	// 视频编码格式，支持的取值及含义如下所示。
	// * h264：使用 H.264 视频编码格式；
	// * h265：使用 H.265 视频编码格式；
	// * h266：使用 H.266 视频编码格式；
	// * copy：不进行视频转码，所有视频编码参数不生效，视频编码参数包括视频帧率（FPS）、视频码率（VideoBitrate）、分辨率设置（As、Width、Height、ShortSide、LongSide）、GOP 和 BFrames
	// 等。
	Vcodec *string `json:"Vcodec,omitempty"`

	// 视频码率，单位为 kbps。
	VideoBitrate *int32 `json:"VideoBitrate,omitempty"`

	// 视频宽度。 :::tip
	// * 当关闭视频分辨率自适应（As 取值为 0）时，转码分辨率将取 Width 和 Height 的值对转码视频进行拉伸；
	// * 当关闭视频分辨率自适应（As 取值为 0）时，Width 和 Height 任一取值为 0 时，转码视频将保持源流尺寸。 :::
	Width *int32 `json:"Width,omitempty"`
}

type ListVhostTransCodePresetResResultCustomizePresetListItem struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 域名空间。
	Vhost string `json:"Vhost"`

	// 转码配置具体信息。
	TranscodePreset *ListVhostTransCodePresetResResultCustomizePresetListItemTranscodePreset `json:"TranscodePreset,omitempty"`
}

// ListVhostTransCodePresetResResultCustomizePresetListItemTranscodePreset - 转码配置具体信息。
type ListVhostTransCodePresetResResultCustomizePresetListItemTranscodePreset struct {

	// 音频编码格式。包括以下 3 种类型。
	// * aac：使用 aac 编码格式；
	// * copy：不进行转码，所有音频编码参数不生效；
	// * opus：使用 opus 编码格式。
	Acodec *string `json:"Acodec,omitempty"`

	// 宽高自适应模式开关。
	// * 0：关闭宽高自适应，按照Width和Height的取值进行拉伸；
	// * 1：开启宽高自适应，按照ShortSide或LongSide等比缩放。
	As *string `json:"As,omitempty"`

	// 音频码率，单位为 kbps。
	AudioBitrate *int32 `json:"AudioBitrate,omitempty"`

	// 是否开启转码不超过源流分辨率。开启后，当源流分辨率低于转码配置分辨率时(即源流宽低于转码配置宽且源流高低于转码配置高时)，将按源流视频分辨率进行转码，默认开启。
	// * 0：关闭
	// * 1：开启
	AutoTransResolution *int32 `json:"AutoTransResolution,omitempty"`

	// 是否开启不超过源流码率。开启后，当源流码率低于转码配置码率时，将按照源流视频码率进行转码，默认开启。
	// * 0：关闭
	// * 1：开启
	AutoTransVb *int32 `json:"AutoTransVb,omitempty"`

	// 是否开启不超过源流帧率。开启后，当源流帧率低于转码配置帧率时，将按照源流视频帧率进行转码，默认开启。
	// * 0：关闭
	// * 1：开启
	AutoTransVr *int32 `json:"AutoTransVr,omitempty"`

	// 2 个参考帧之间的最大 B 帧数。BFrames取 0 时，表示去 B 帧。
	BFrames *int32 `json:"BFrames,omitempty"`

	// 动态范围，画质增强类型生效
	// * SDR：输出为SDR
	// * HDR：输出为HDR
	DynamicRange *string `json:"DynamicRange,omitempty"`

	// 是否开启智能插帧，只对画质增强类型生效
	// * 0：不开启
	// * 1：开启
	FISwitch *int64 `json:"FISwitch,omitempty"`

	// 视频帧率，单位为 fps，帧率越大，画面越流畅。
	FPS *int32 `json:"FPS,omitempty"`

	// IDR 帧之间的最大间隔，单位为 s。
	GOP *int32 `json:"GOP,omitempty"`

	// 视频高度。
	Height *int32 `json:"Height,omitempty"`

	// 长边长度。 :::tip 当As的取值为 1 时，如果LongSide和ShortSide都取 0，表示保持源流尺寸。 :::
	LongSide *int32 `json:"LongSide,omitempty"`

	// 转码模板参数的类型
	// * va：表示使用画质增强
	ParamType *string `json:"ParamType,omitempty"`

	// 转码配置名称。
	Preset *string `json:"Preset,omitempty"`

	// 是否极智超清转码。
	// * true：极智超清；
	// * false：标准转码。
	Roi *bool `json:"Roi,omitempty"`

	// 使用场景，画质增强时生效 football：足球场景
	SceneType *string `json:"SceneType,omitempty"`

	// 短边长度。 :::tip 当As的取值为 1 时，如果LongSide和ShortSide都取 0，表示保持源流尺寸。 :::
	ShortSide *int32 `json:"ShortSide,omitempty"`

	// 转码停止时长，支持触发方式为拉流转码时设置，表示断开拉流后转码停止的时长，单位为 s，取值范围为 [0,300]，-1 表示不停止转码，默认值为 60。
	StopInterval *int32 `json:"StopInterval,omitempty"`

	// 转码流后缀名。
	SuffixName *string `json:"SuffixName,omitempty"`
	Threads    *int32  `json:"Threads,omitempty"`

	// 转码触发方式，默认为拉流转码，支持以下取值。
	// * Push：推流转码，直播推流后会自动启动转码任务，生成转码流；
	// * Pull：拉流转码，直播推流后，需要主动播放转码流才会启动转码任务，生成转码流。
	TransType *string `json:"TransType,omitempty"`

	// 视频编码格式。
	// * h264：使用 H.264 编码格式；
	// * h265：使用 H.265 编码格式；
	// * h266：使用 H.266 编码格式；
	// * copy：不进行转码，所有视频编码参数不生效。
	Vcodec *string `json:"Vcodec,omitempty"`

	// 视频码率，单位为 kbps。
	VideoBitrate *int32 `json:"VideoBitrate,omitempty"`

	// 视频宽度。
	Width *int32 `json:"Width,omitempty"`
}

type ListVhostWatermarkPresetBody struct {

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`
}

type ListVhostWatermarkPresetRes struct {

	// REQUIRED
	ResponseMetadata ListVhostWatermarkPresetResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListVhostWatermarkPresetResResult          `json:"Result,omitempty"`
}

type ListVhostWatermarkPresetResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                            `json:"Version"`
	Error   *ListVhostWatermarkPresetResResponseMetadataError `json:"Error,omitempty"`
}

type ListVhostWatermarkPresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type ListVhostWatermarkPresetResResult struct {

	// 统计消息，提供查询成功和失败的数量。
	StaticsMsg *string `json:"StaticsMsg,omitempty"`

	// 获取配置失败的列表，返回获取失败的配置信息及获取失败的原因。
	WatermarkErrMsgList []*ListVhostWatermarkPresetResResultWatermarkErrMsgListItem `json:"WatermarkErrMsgList,omitempty"`

	// 水印配置列表。
	WatermarkPresetList []*ListVhostWatermarkPresetResResultWatermarkPresetListItem `json:"WatermarkPresetList,omitempty"`
}

type ListVhostWatermarkPresetResResultWatermarkErrMsgListItem struct {

	// 火山引擎账号 ID。
	AccountID *string `json:"AccountID,omitempty"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 获取水印模板失败的具体错误信息。
	ErrMsg *string `json:"ErrMsg,omitempty"`

	// 域名空间名称。
	Vhost *string `json:"Vhost,omitempty"`
}

type ListVhostWatermarkPresetResResultWatermarkPresetListItem struct {

	// 火山引擎账号 ID。
	AccountID *string `json:"AccountID,omitempty"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 水印模版 ID。
	ID *int32 `json:"ID,omitempty"`

	// 需要添加水印的直播画面方向。
	// * vertical：竖屏；
	// * horizontal：横屏。
	Orientation *string `json:"Orientation,omitempty"`

	// 水印图片编码字符串。
	Picture *string `json:"Picture,omitempty"`

	// 水印图片文件名。
	PictureKey *string `json:"PictureKey,omitempty"`

	// 水印图片对应的 HTTP 地址。与水印图片字符串字段二选一传入，同时传入时，以水印图片字符串参数为准。
	PictureURL *string `json:"PictureURL,omitempty"`

	// 水平偏移，表示水印左侧边与转码流画面左侧边之间的距离，使用相对比率，取值范围为 [0,1]。
	PosX *float32 `json:"PosX,omitempty"`

	// 垂直偏移，表示水印顶部边与转码流画面顶部边之间的距离，使用相对比率，取值范围为 [0,1]。
	PosY *float32 `json:"PosY,omitempty"`

	// 水印图片预览背景高度，单位为 px。
	PreviewHeight *float32 `json:"PreviewHeight,omitempty"`

	// 水印图片预览背景宽度，单位为 px。
	PreviewWidth *float32 `json:"PreviewWidth,omitempty"`

	// 水印相对高度，水印高度占直播转码流画面高度的比例，取值范围为 [0,1]，水印宽度会随高度等比缩放。
	RelativeHeight *float32 `json:"RelativeHeight,omitempty"`

	// 水印相对宽度，水印宽度占直播转码流画面宽度的比例，取值范围为 [0,1]，水印高度会随宽度等比缩放。
	RelativeWidth *float32 `json:"RelativeWidth,omitempty"`

	// 流名称。
	Stream *string `json:"Stream,omitempty"`

	// 域名空间。
	Vhost *string `json:"Vhost,omitempty"`
}

type ListWatermarkPresetBody struct {

	// REQUIRED; 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 直播流名。可选。
	Stream *string `json:"Stream,omitempty"`
}

type ListWatermarkPresetRes struct {

	// REQUIRED
	ResponseMetadata ListWatermarkPresetResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListWatermarkPresetResResult          `json:"Result,omitempty"`
}

type ListWatermarkPresetResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version   string                                       `json:"Version"`
	Error     *ListWatermarkPresetResResponseMetadataError `json:"Error,omitempty"`
	RequestID *string                                      `json:"RequestID,omitempty"`
}

type ListWatermarkPresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type ListWatermarkPresetResResult struct {

	// REQUIRED; 水印模板。
	Preset ListWatermarkPresetResResultPreset `json:"Preset"`
}

// ListWatermarkPresetResResultPreset - 水印模板。
type ListWatermarkPresetResResultPreset struct {

	// 火山引擎账号 ID。
	AccountID *string `json:"AccountID,omitempty"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 水印配置的 ID。
	ID *int32 `json:"ID,omitempty"`

	// 需要添加水印的直播画面方向。
	// * vertical：竖屏；
	// * horizontal：横屏。
	Orientation *string `json:"Orientation,omitempty"`

	// 水印图片编码字符串。
	Picture *string `json:"Picture,omitempty"`

	// 水印图片文件名。
	PictureKey *string `json:"PictureKey,omitempty"`

	// 水印图片对应的 HTTP 地址。与水印图片编码字符串字段二选一传入，同时传入时，以水印图片编码字符串参数为准。
	PictureURL *string `json:"PictureURL,omitempty"`

	// 水平偏移，表示水印左侧边与转码流画面左侧边之间的距离，使用相对比率，取值范围为 [0,1]。
	PosX *float32 `json:"PosX,omitempty"`

	// 垂直偏移，表示水印顶部边与转码流画面顶部边之间的距离，使用相对比率，取值范围为 [0,1]。
	PosY *float32 `json:"PosY,omitempty"`

	// 水印图片预览背景高度，单位为 px。
	PreviewHeight *float32 `json:"PreviewHeight,omitempty"`

	// 水印图片预览背景宽度，单位为 px。
	PreviewWidth *float32 `json:"PreviewWidth,omitempty"`

	// 水印相对高度，水印高度占直播转码流画面高度的比例，取值范围为 [0,1]，水印宽度会随高度等比缩放。
	RelativeHeight *float32 `json:"RelativeHeight,omitempty"`

	// 水印相对宽度，水印宽度占直播转码流画面宽度的比例，取值范围为 [0,1]，水印高度会随宽度等比缩放。
	RelativeWidth *float32 `json:"RelativeWidth,omitempty"`

	// 流名称。
	Stream *string `json:"Stream,omitempty"`

	// 域名空间。
	Vhost *string `json:"Vhost,omitempty"`
}

type RestartPullToPushTaskBody struct {

	// REQUIRED; 任务 ID，任务的唯一标识，您可以通过获取拉流转推任务列表 [https://www.volcengine.com/docs/6469/1126896]接口获取状态为停用的任务 ID。
	TaskID string `json:"TaskId"`
}

type RestartPullToPushTaskRes struct {

	// REQUIRED
	ResponseMetadata RestartPullToPushTaskResResponseMetadata `json:"ResponseMetadata"`
}

type RestartPullToPushTaskResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                         `json:"Version"`
	Error   *RestartPullToPushTaskResResponseMetadataError `json:"Error,omitempty"`
}

type RestartPullToPushTaskResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type ResumeStreamBody struct {

	// REQUIRED; 直播流使用的应用名称。
	App string `json:"App"`

	// REQUIRED; 直播流使用的流名称。
	Stream string `json:"Stream"`

	// 直播流使用的域名。 :::tip 参数 Domain 和 Vhost 传且仅传一个。 :::
	Domain *string `json:"Domain,omitempty"`

	// 域名空间。您可以调用 DescribeForbiddenStreamInfoByPage [https://www.volcengine.com/docs/6469/1126843] 接口，查看禁推直播流的信息，包括 Vhost、Domain、App
	// 和 Stream。 :::tip 参数 Domain 和 Vhost 传且仅传一个。 :::
	Vhost *string `json:"Vhost,omitempty"`
}

type ResumeStreamRes struct {

	// REQUIRED
	ResponseMetadata ResumeStreamResResponseMetadata `json:"ResponseMetadata"`
}

type ResumeStreamResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                `json:"Version"`
	Error   *ResumeStreamResResponseMetadataError `json:"Error,omitempty"`
}

type ResumeStreamResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type StopPullToPushTaskBody struct {

	// REQUIRED; 任务 ID，任务的唯一标识，您可以通过获取拉流转推任务列表 [https://www.volcengine.com/docs/6469/1126896]接口获取状态为未开始或生效中的任务 ID。
	TaskID string `json:"TaskId"`
}

type StopPullToPushTaskRes struct {

	// REQUIRED
	ResponseMetadata StopPullToPushTaskResResponseMetadata `json:"ResponseMetadata"`
}

type StopPullToPushTaskResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                      `json:"Version"`
	Error   *StopPullToPushTaskResResponseMetadataError `json:"Error,omitempty"`
}

type StopPullToPushTaskResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UnBindEncryptDRMBody struct {

	// REQUIRED; 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`
}

type UnBindEncryptDRMRes struct {

	// REQUIRED
	ResponseMetadata UnBindEncryptDRMResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UnBindEncryptDRMResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type UnbindCertBody struct {

	// REQUIRED; 填写需要解绑 HTTPS 证书的域名。 您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看需要解绑证书的域名。
	Domain string `json:"Domain"`

	// REQUIRED; 视频直播服务的配置空间，由 1 到 60 位数字、字母、下划线及"-"和"."组成。 :::tip 参数 Domain 和 Vhost 传且仅传一个。 :::
	Vhost string `json:"Vhost"`
}

type UnbindCertRes struct {

	// REQUIRED
	ResponseMetadata UnbindCertResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type UnbindCertResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                              `json:"Version"`
	Error   *UnbindCertResResponseMetadataError `json:"Error,omitempty"`
}

type UnbindCertResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateAuthKeyBody struct {

	// REQUIRED; 鉴权配置参数，包括鉴权密钥、鉴权字段、加密字符串生成算法等。
	AuthDetailList []UpdateAuthKeyBodyAuthDetailListItem `json:"AuthDetailList"`

	// REQUIRED; 鉴权场景类型，取值及含义如下所示。
	// * push：推流鉴权；
	// * pull：拉流鉴权。
	SceneType string `json:"SceneType"`

	// 应用名称，取值与直播流地址中 AppName 字段取值相同，默认为空，表示所有应用名称。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App *string `json:"App,omitempty"`

	// 直播流使用的域名。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名。
	// :::tip 参数 Domain 和 Vhost 传且仅传一个。 :::
	Domain *string `json:"Domain,omitempty"`

	// 是否开启 URL 地址鉴权，取值及含义如下所示。
	// * false：关闭（默认值）；
	// * true：开启。
	PushPullEnable *bool `json:"PushPullEnable,omitempty"`

	// 鉴权生效时长，单位为秒，默认值为 604800，取值范围为 [60,2592000]，超出生效时长后 URL 无法使用。
	ValidDuration *int32 `json:"ValidDuration,omitempty"`

	// 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看需要查询的直播流使用的域名所属的域名空间。
	// :::tip 参数
	// Domain 和 Vhost 传且仅传一个。 :::
	Vhost *string `json:"Vhost,omitempty"`
}

type UpdateAuthKeyBodyAuthDetailListItem struct {

	// REQUIRED; 推/拉流鉴权时必选
	EncryptionAlgorithm string `json:"EncryptionAlgorithm"`

	// REQUIRED; 推/拉流鉴权时必选
	SecretKey string `json:"SecretKey"`

	// 推/拉流鉴权时必选
	AuthField map[string]*string `json:"AuthField,omitempty"`

	// 推/拉流鉴权时必选
	EncryptField []*string `json:"EncryptField,omitempty"`
}

type UpdateAuthKeyRes struct {

	// REQUIRED
	ResponseMetadata UpdateAuthKeyResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type UpdateAuthKeyResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                 `json:"Version"`
	Error   *UpdateAuthKeyResResponseMetadataError `json:"Error,omitempty"`
}

type UpdateAuthKeyResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateCMAFConfigBody struct {

	// REQUIRED
	Vhost             string   `json:"Vhost"`
	App               *string  `json:"App,omitempty"`
	DefaultLatency    *float32 `json:"DefaultLatency,omitempty"`
	DisableLowLatency *bool    `json:"DisableLowLatency,omitempty"`
	Interval          *float32 `json:"Interval,omitempty"`
	PlaylistLength    *int32   `json:"PlaylistLength,omitempty"`
}

type UpdateCMAFConfigRes struct {

	// REQUIRED
	ResponseMetadata UpdateCMAFConfigResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateCMAFConfigResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type UpdateCallbackBody struct {

	// REQUIRED; 回调信息列表。
	CallbackDetailList []UpdateCallbackBodyCallbackDetailListItem `json:"CallbackDetailList"`

	// REQUIRED; 回调的消息类型，取值及含义如下所示。
	// * push：推流开始回调；
	// * push_end：推流结束回调；
	// * snapshot：截图回调；
	// * record：录制任务状态回调；
	// * audit_snapshot：截图审核结果回调。
	MessageType string `json:"MessageType"`

	// domain / app 二选一必传
	App *string `json:"App,omitempty"`

	// Dictionary of
	AppendField map[string]*string `json:"AppendField,omitempty"`

	// 回调消息发送是否开启鉴权，默认为 false，取值及含义如下所示。
	// * false：不开启；
	// * true：开启。
	AuthEnable *bool `json:"AuthEnable,omitempty"`

	// Dictionary of
	AuthField map[string]*string `json:"AuthField,omitempty"`

	// 回调消息发送鉴权密钥。 :::tip 如果 AuthEnable 为 true，则密钥必填。 :::
	AuthKeyPrimary *string   `json:"AuthKeyPrimary,omitempty"`
	AuthKeySecond  *string   `json:"AuthKeySecond,omitempty"`
	CallbackField  []*string `json:"CallbackField,omitempty"`

	// 直播流使用的推流域名。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名。
	// :::tipVhost和Domain传且仅传一个。 :::
	Domain              *string   `json:"Domain,omitempty"`
	EncryptField        []*string `json:"EncryptField,omitempty"`
	EncryptionAlgorithm *string   `json:"EncryptionAlgorithm,omitempty"`

	// 回调时的http方式，不填默认为post
	// * post: POST方式
	// * get: GET方式
	HTTPMethod *string `json:"HttpMethod,omitempty"`

	// 回调重试间隔，默认为3秒
	RetryInternalSecond *int32 `json:"RetryInternalSecond,omitempty"`

	// 回调重试次数，默认为3次
	RetryTimes *int32 `json:"RetryTimes,omitempty"`

	// 鉴权参数位置,该参数暂时没有做到通用，鉴权参数名称为 postauthparam 要经过低代码转换，不然不会生效。get 方式默认只能放在param里面，post 方式默认放在body，支持以下选择
	// * param: 鉴权放在链接里面
	// * body: 鉴权放在body里面
	SecHandlerType *string `json:"SecHandlerType,omitempty"`

	// 任务状态回调开关 0关闭 1开启
	TaskStatusCallback *int32 `json:"TaskStatusCallback,omitempty"`

	// 回调超时时间，默认为4秒
	TimeoutSecond *int32 `json:"TimeoutSecond,omitempty"`

	// 是否开启转码流回调，默认为 0。取值及含义如下所示。
	// * 0：false，不开启；
	// * 1：true，开启。
	// :::tip 回调类型为推流开始或推流结束时生效。 :::
	TranscodeCallback *int32 `json:"TranscodeCallback,omitempty"`
	ValidDuration     *int32 `json:"ValidDuration,omitempty"`

	// domain / app 二选一必传
	Vhost *string `json:"Vhost,omitempty"`
}

type UpdateCallbackBodyCallbackDetailListItem struct {

	// REQUIRED; 回调类型，返回 HTTP，表示可以使用 HTTP 和 HTTPS 地址接收回调消息。
	CallbackType string `json:"CallbackType"`

	// REQUIRED; 回调消息接收地址。
	URL string `json:"URL"`
}

type UpdateCallbackRes struct {

	// REQUIRED
	ResponseMetadata UpdateCallbackResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type UpdateCallbackResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                  `json:"Version"`
	Error   *UpdateCallbackResResponseMetadataError `json:"Error,omitempty"`
}

type UpdateCallbackResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateClusterRateLimitBody struct {

	// REQUIRED
	Limit int32 `json:"Limit"`

	// REQUIRED
	Type string `json:"Type"`

	// REQUIRED
	Vhost             string  `json:"Vhost"`
	AggregationPeriod *int32  `json:"AggregationPeriod,omitempty"`
	App               *string `json:"App,omitempty"`
	Domain            *string `json:"Domain,omitempty"`
	Param             *string `json:"Param,omitempty"`
	RejectCode        *int32  `json:"RejectCode,omitempty"`
	RejectDuration    *int32  `json:"RejectDuration,omitempty"`
	Status            *int32  `json:"Status,omitempty"`
}

type UpdateClusterRateLimitRes struct {

	// REQUIRED
	ResponseMetadata UpdateClusterRateLimitResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateClusterRateLimitResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestId为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type UpdateDomainVhostBody struct {

	// REQUIRED; 待修改所属域名空间的的拉流/推流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看域名信息。
	Domain string `json:"Domain"`

	// REQUIRED; 目的域名空间，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看域名所属的域名空间信息。
	Vhost string `json:"Vhost"`
}

type UpdateDomainVhostRes struct {

	// REQUIRED
	ResponseMetadata UpdateDomainVhostResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateDomainVhostResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                     `json:"Version"`
	Error   *UpdateDomainVhostResResponseMetadataError `json:"Error,omitempty"`
}

type UpdateDomainVhostResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateEncryptDRMBody struct {

	// DRM 证书管理平台 API 访问密钥，获取方法请参见最佳实践-直播 DRM 加密 [https://www.volcengine.com/docs/6469/1219836#在-intertrust-平台创建访问密钥]。
	APIKey *string `json:"APIKey,omitempty"`

	// 申请 FairPlay 证书过程中 Apple 返回的 ASk（Application Secret Key）字符串。
	ApplicationSecretKey *string `json:"ApplicationSecretKey,omitempty"`

	// FairPlay 证书文件内容。
	CertificateFile *string `json:"CertificateFile,omitempty"`

	// FairPlay 证书文件名称。
	CertificateFileName *string `json:"CertificateFileName,omitempty"`

	// 自定义 FairPlay 证书名称，支持由小写字母（a - z）、数字（0 - 9）和短横线（-）组成，最小长度为 2个字符，最大长度为 128 个字符。FairPlay 证书相关参数的获取方法请参见最佳实践-直播 DRM 加密 [https://www.volcengine.com/docs/6469/1219836#在-apple-官网获取-fairplay-证书]。
	CertificateName *string `json:"CertificateName,omitempty"`

	// 申请 FairPlay 证书时创建的私钥文件密钥。
	PrivateKey *string `json:"PrivateKey,omitempty"`

	// 申请 FairPlay 证书时创建的私钥文件内容。
	PrivateKeyFile *string `json:"PrivateKeyFile,omitempty"`

	// 申请 FairPlay 证书时创建的私钥文件名称。
	PrivateKeyFileName *string `json:"PrivateKeyFileName,omitempty"`
}

type UpdateEncryptDRMRes struct {

	// REQUIRED
	ResponseMetadata UpdateEncryptDRMResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 视请求的接口而定
	Result interface{} `json:"Result"`
}

type UpdateEncryptDRMResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type UpdateFormatAccessRuleBody struct {

	// REQUIRED
	Domain string `json:"Domain"`

	// REQUIRED
	FormatAccessRule UpdateFormatAccessRuleBodyFormatAccessRule `json:"FormatAccessRule"`

	// REQUIRED
	Vhost string `json:"Vhost"`
}

type UpdateFormatAccessRuleBodyFormatAccessRule struct {

	// REQUIRED
	Enable bool `json:"Enable"`

	// REQUIRED
	FormatList []string `json:"FormatList"`

	// REQUIRED
	Type string `json:"Type"`
}

type UpdateFormatAccessRuleRes struct {

	// REQUIRED
	ResponseMetadata UpdateFormatAccessRuleResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateFormatAccessRuleResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestId为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type UpdateHLSConfigBody struct {

	// REQUIRED; 服务类型
	ServiceType string `json:"ServiceType"`

	// REQUIRED; 域名空间
	Vhost string `json:"Vhost"`

	// 应用名称
	App *string `json:"App,omitempty"`

	// 永久存储ts，默认为true，也就是清零
	CleanUp *bool `json:"CleanUp,omitempty"`

	// json配置，通过json的方式添加时填写
	Config *string `json:"Config,omitempty"`

	// 时间戳置零,默认为false
	CopyTs *bool `json:"CopyTs,omitempty"`

	// 开启预取，默认false
	EnablePrefetch *bool `json:"EnablePrefetch,omitempty"`

	// 首个m3u8 ts的个数
	FirstPlaylistLength *int32 `json:"FirstPlaylistLength,omitempty"`

	// 时间戳gap，默认5s
	Gap      *int32   `json:"Gap,omitempty"`
	Interval *float32 `json:"Interval,omitempty"`

	// 切片最大帧数
	MaxFrame *int32 `json:"MaxFrame,omitempty"`

	// 切片最大大小，单位byte，默认 524288000
	MaxSize *int32 `json:"MaxSize,omitempty"`

	// 可选枚举值 "audio_only" "video_only "video_keyframe_only" "video_single_keyframe_only"
	PacketFilter *string `json:"PacketFilter,omitempty"`

	// ts存储位置
	Path *string `json:"Path,omitempty"`

	// m3u8的ts个数，默认3个
	PlaylistLength *int32 `json:"PlaylistLength,omitempty"`

	// 预取ts个数
	PrefetchNum *int32 `json:"PrefetchNum,omitempty"`

	// ts文件后缀
	Suffix *string `json:"Suffix,omitempty"`

	// ts缓存时间，单位s，默认60s
	TsExpiration *int32 `json:"TsExpiration,omitempty"`
}

type UpdateHLSConfigRes struct {

	// REQUIRED
	ResponseMetadata UpdateHLSConfigResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type UpdateHLSConfigResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type UpdateHTTPHeaderConfigBody struct {

	// REQUIRED; 配置完成后是否启用，取值及含义如下所示。
	// * true：启用；
	// * false：禁用。
	Enable bool `json:"Enable"`

	// REQUIRED; Header 具体字段配置。
	HeaderConfigList []UpdateHTTPHeaderConfigBodyHeaderConfigListItem `json:"HeaderConfigList"`

	// REQUIRED; HTTP Header 的类型，支持的取值及含义如下所示。
	// * 0：请求响应头；
	// * 1：回源请求头。
	Phase int32 `json:"Phase"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看需要域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 是否保留原 Header 配置，取值及含义如下所示。
	// * 0：保留（默认值）；
	// * 1：不保留。
	BlockOriginal *int32 `json:"BlockOriginal,omitempty"`

	// 拉流域名。默认为空，表示 Vhost 下的全部拉流域名。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看配置的拉流域名。
	Domain *string `json:"Domain,omitempty"`
}

type UpdateHTTPHeaderConfigBodyHeaderConfigListItem struct {

	// REQUIRED; Header 配置中字段 Value 值的类型，取值及含义如下所示。
	// * 0：常量；
	// * 1：变量。
	HeaderFieldType int32 `json:"HeaderFieldType"`

	// Header 配置中字段的 Key 值，最大长度为 1024 个字符，多个 Header 不可重名。
	HeaderKey *string `json:"HeaderKey,omitempty"`

	// Header 配置中字段的 Value 值，最大长度为 1024 个字符，支持使用常量和变量作为 Value 值。
	// HTTP Header 类型为回源请求头时，支持使用以下变量为 Value 赋值：
	// * ${domain}：客户端拉流请求中使用的域名。
	// * ${uri}：客户端拉流请求中不包括查询参数的路径。如果请求被重写，则表示重写后的路径。
	// * ${args}：客户端拉流请求中的查询参数。如果请求被重写，则表示重写后的参数。
	// * ${remote_addr}：发送拉流请求的客户端 IP 地址。
	// * ${server_addr}：响应客户端拉流请求的 CDN 节点 IP 地址。
	// HTTP Header 类型为请求响应头时，支持使用以下变量为 Value 赋值：
	// * ${upstream_host}：客户端拉流请求中使用的域名。
	// * ${upstream_uri}：客户端拉流请求中不包括查询参数的路径。如果请求被重写，则表示重写后的路径。
	// * ${upstream_args}：客户端拉流请求中的查询参数。如果请求被重写，则表示重写后的参数。
	// * ${remote_addr}：发送拉流请求的客户端 IP 地址。
	HeaderValue *string `json:"HeaderValue,omitempty"`
}

type UpdateHTTPHeaderConfigRes struct {

	// REQUIRED
	ResponseMetadata UpdateHTTPHeaderConfigResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateHTTPHeaderConfigResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type UpdateIPAccessRuleBody struct {

	// REQUIRED; 推流域名或拉流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取需要配置
	// IP 访问限制的域名。
	Domain string `json:"Domain"`

	// REQUIRED; IP 访问限制规则。
	IPAccessRule UpdateIPAccessRuleBodyIPAccessRule `json:"IPAccessRule"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取需要配置
	// IP 访问限制的域名所属的域名空间。
	Vhost string `json:"Vhost"`
}

// UpdateIPAccessRuleBodyIPAccessRule - IP 访问限制规则。
type UpdateIPAccessRuleBodyIPAccessRule struct {

	// REQUIRED; 是否开启当前限制，取值及含义如下所示。
	// * true: 开启；
	// * false: 关闭。
	Enable bool `json:"Enable"`

	// REQUIRED; 名单中的 IP 信息。
	IPList []string `json:"IPList"`

	// REQUIRED; IP 访问限制的类型，取值及含义如下所示。
	// * allow: 白名单；
	// * deny: 黑名单。
	Type string `json:"Type"`
}

type UpdateIPAccessRuleRes struct {

	// REQUIRED
	ResponseMetadata UpdateIPAccessRuleResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateIPAccessRuleResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type UpdateLatencyConfigBody struct {

	// REQUIRED; 单位毫秒，大于等于0
	GopCacheSize int32 `json:"GopCacheSize"`

	// 与Vhost二选一
	Domain *string `json:"Domain,omitempty"`
	Vhost  *string `json:"Vhost,omitempty"`
}

type UpdateLatencyConfigRes struct {

	// REQUIRED
	ResponseMetadata UpdateLatencyConfigResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateLatencyConfigResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type UpdatePullToPushTaskBody struct {

	// REQUIRED; 任务等结束时间，Unix 时间戳，单位为秒。 :::tip 拉流转推任务持续时间最长为 7 天。 :::
	EndTime int32 `json:"EndTime"`

	// REQUIRED; 设置点播视频转推至第三方推流域名时是否使用推流优先级参数，缺省情况下表示不使用此参数，支持的取值及含义如下。
	// * true：使用
	//
	//
	// * false：不使用
	//
	//
	// :::tip
	// * 使用点播视频转推直播实现视频循环播放（轮播）时，支持使用带有推流优先级参数的推流地址进行推流，如在第一个点播视频的推流地址后添加 pri=10、在第二个点播视频的推流地址后添加 pri=11，可达到使用推流优先级高的流替换推流优先级低的流的目的。相比不使用推流优先级参数时可实现更平滑的轮播视频切换。
	//
	//
	// * 推流至非第三方域名时，默认支持使用带有推流优先级参数的推流地址。
	//
	//
	// * 推流至第三方域名时，如需使用推流优先级参数实现新流替换旧流时，需在创建拉流转推时为推流域名开启推流优先级参数配置开关。 :::
	PushPriority bool `json:"PushPriority"`

	// REQUIRED; 任务的开始时间，Unix 时间戳，单位为秒。 :::tip 拉流转推任务持续时间最长为 7 天。 :::
	StartTime int32 `json:"StartTime"`

	// REQUIRED; 任务 ID，任务的唯一标识，您可以通过获取拉流转推任务列表 [https://www.volcengine.com/docs/6469/1126896]接口获取。
	TaskID string `json:"TaskId"`

	// REQUIRED; 拉流来源类型，支持的取值及含义如下。
	// * 0：直播源；
	// * 1：点播视频。
	Type int32 `json:"Type"`

	// 推流应用名称，推流地址（DstAddr）为空时必传；反之，则该参数不生效
	App *string `json:"App,omitempty"`

	// 接收拉流转推任务状态回调的地址，最大长度为 2000 个字符。
	CallbackURL *string `json:"CallbackURL,omitempty"`

	// 续播策略，续播策略指转推点播视频进行直播时出现断流并恢复后，如何继续播放的策略，拉流来源类型为点播视频（Type 为 1）时参数生效，支持的取值及含义如下。
	// * 0：从断流处续播（默认值）；
	// * 1：从断流处+自然流逝时长处续播。
	ContinueStrategy *int32 `json:"ContinueStrategy,omitempty"`

	// 点播视频文件循环播放模式，当拉流来源类型为点播视频时为必选参数，参数取值及含义如下所示。
	// * -1：无限次循环，至任务结束；
	// * 0：有限次循环，循环次数以 PlayTimes 取值为准；
	// * >0：有限次循环，循环次数以 CycleMode 取值为准。
	CycleMode *int32 `json:"CycleMode,omitempty"`

	// 推流域名，推流地址（DstAddr）为空时必传；反之，则该参数不生效
	Domain *string `json:"Domain,omitempty"`

	// 推流地址，即直播源或点播视频转推的目标地址。
	DstAddr *string `json:"DstAddr,omitempty"`

	// 点播文件启播时间偏移值，单位为秒，仅当点播视频播放地址列表（SrcAddrS）只有一个地址，且未配置 Offsets 时生效，缺省情况下表示不进行偏移。
	Offset *float32 `json:"Offset,omitempty"`

	// 点播文件启播时间偏移值，单位为秒，数量与拉流地址列表中地址数量相等，缺省情况下表示不进行偏移。 拉流来源类型为点播视频（Type 为 1）时，参数生效。
	OffsetS []*float32 `json:"OffsetS,omitempty"`

	// 点播视频文件循环播放次数，当 CycleMode 取值为 0 时，PlayTimes 取值将作为循环播放次数。 :::tip PlayTimes 为冗余参数，您可以将 PlayTimes 置 0 后直接使用 CycleMode 指定点播视频文件循环播放次数。
	// :::
	PlayTimes *int32 `json:"PlayTimes,omitempty"`

	// 是否开启点播预热，开启点播预热后，系统会自动将点播视频文件缓存到 CDN 节点上，当用户请求直播时，可以直播从 CDN 节点获取视频，从而提高直播流畅度。 拉流来源类型为点播视频（Type 为 1）时，参数生效。
	// * 0：不开启；
	// * 1：开启（默认值）。
	PreDownload *int32 `json:"PreDownload,omitempty"`

	// 直播源的拉流地址，拉流来源类型为直播源（Type 为 0）时，为必选参数，最大长度为 1000 个字符。
	SrcAddr *string `json:"SrcAddr,omitempty"`

	// 点播视频播放地址列表，拉流来源类型为点播视频（Type 为 1）时，为必选参数，最多支持传入 30 个点播视频播放地址，每个地址最大长度为 1000 个字符。
	SrcAddrS []*string `json:"SrcAddrS,omitempty"`

	// 推流的流名称，推流地址（DstAddr）为空时必传；反之，则该参数不生效
	Stream *string `json:"Stream,omitempty"`

	// 拉流转推任务的名称，默认为空表示不配置任务名称。支持由中文、大小写字母（A - Z、a - z）和数字（0 - 9）组成，长度为 1 到 20 各字符。
	Title *string `json:"Title,omitempty"`

	// 为拉流转推视频添加的水印配置信息。
	Watermark *UpdatePullToPushTaskBodyWatermark `json:"Watermark,omitempty"`
}

// UpdatePullToPushTaskBodyWatermark - 为拉流转推视频添加的水印配置信息。
type UpdatePullToPushTaskBodyWatermark struct {

	// REQUIRED; 水印图片字符串，图片最大 2MB，最小 100Bytes，最大分辨率为 1080×1080。图片 Data URL 格式为：data:image/<mediatype>;base64,<data>。
	// * mediatype：图片类型，支持 png、jpg、jpeg 格式；
	// * data：base64 编码的图片字符串。
	// 例如，data:image/png;base64,iVBORw0KGg****mCC
	Picture string `json:"Picture"`

	// REQUIRED; 水印宽度占直播原始画面宽度百分比，支持精度为小数点后两位。
	Ratio float32 `json:"Ratio"`

	// REQUIRED; 水平偏移，表示水印左侧边与转码流画面左侧边之间的距离，使用相对比率，取值范围为 [0,1)。
	RelativePosX float32 `json:"RelativePosX"`

	// REQUIRED; 垂直偏移，表示水印顶部边与转码流画面顶部边之间的距离，使用相对比率，取值范围为 [0,1)。
	RelativePosY float32 `json:"RelativePosY"`
}

type UpdatePullToPushTaskRes struct {

	// REQUIRED
	ResponseMetadata UpdatePullToPushTaskResResponseMetadata `json:"ResponseMetadata"`
}

type UpdatePullToPushTaskResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                        `json:"Version"`
	Error   *UpdatePullToPushTaskResResponseMetadataError `json:"Error,omitempty"`
}

type UpdatePullToPushTaskResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateRecordPresetV2Body struct {

	// REQUIRED; 录制配置的名称。您可以调用 ListVhostRecordPresetV2 [https://www.volcengine.com/docs/6469/1126858] 接口查看待更新录制配置的 Name 取值。
	Preset string `json:"Preset"`

	// REQUIRED; 域名空间。您可以调用 ListVhostRecordPresetV2 [https://www.volcengine.com/docs/6469/1126858] 接口查看待更新录制配置的 Vhost 取值。
	Vhost string `json:"Vhost"`

	// 应用名称，取值与直播流地址的 AppName 字段取值相同，用来指定待更新的录制配置，默认为空。您可以调用 ListVhostRecordPresetV2 [https://www.volcengine.com/docs/6469/1126858]
	// 接口查看待更新录制配置的 App 取值。
	App *string `json:"App,omitempty"`

	// 录制配置的详细参数配置。
	// :::tip 以下录制参数，未传入值时表示与更新前的配置相同。 :::
	RecordPresetConfig *UpdateRecordPresetV2BodyRecordPresetConfig `json:"RecordPresetConfig,omitempty"`

	// 流名称，取值与直播流地址的 StreamName 字段取值相应，用来指定待更新的录制配置，默认为空。您可以调用 ListVhostRecordPresetV2 [https://www.volcengine.com/docs/6469/1126858]
	// 接口查看待更新录制配置的 Stream 取值。
	Stream *string `json:"Stream,omitempty"`
}

// UpdateRecordPresetV2BodyRecordPresetConfig - 录制配置的详细参数配置。
// :::tip 以下录制参数，未传入值时表示与更新前的配置相同。 :::
type UpdateRecordPresetV2BodyRecordPresetConfig struct {

	// 录制为 FLV 格式时的录制参数。 :::tip 您需至少配置一个录制格式，即 FlvParam、HlsParam、Mp4Param 至少开启一个。 :::
	FlvParam *UpdateRecordPresetV2BodyRecordPresetConfigFlvParam `json:"FlvParam,omitempty"`

	// 录制为 HLS 格式时的录制参数。 :::tip 您需至少配置一个录制格式，即 FlvParam、HlsParam、Mp4Param 至少开启一个。 :::
	HlsParam *UpdateRecordPresetV2BodyRecordPresetConfigHlsParam `json:"HlsParam,omitempty"`

	// 录制为 MP4 格式时的录制参数。 :::tip 您需至少配置一个录制格式，即 FlvParam、HlsParam、Mp4Param 至少开启一个。 :::
	Mp4Param *UpdateRecordPresetV2BodyRecordPresetConfigMp4Param `json:"Mp4Param,omitempty"`

	// 是否录制源流，默认值为 0，支持的取值及含义如下所示。
	// * 0：不录制；
	// * 1：录制。
	// :::tip 转码流和源流需至少选一个进行录制，即是否录制转码流（TranscodeRecord）和是否录制源流（OriginRecord）的取值至少一个不为 0。 :::
	OriginRecord *int32 `json:"OriginRecord,omitempty"`

	// 录制为 HLS 格式时，单个 TS 切片时长，单位为秒，默认值为 10，取值范围为 [5,30]。
	SliceDuration *int32 `json:"SliceDuration,omitempty"`

	// 是否录制转码流，默认值为 0。支持的取值如下所示。
	// * 0：不录制；
	// * 1：录制全部转码流；
	// * 2：录制指定转码流，即通过转码后缀列表 TranscodeSuffixList匹配转码流进行录制，如果转码流后缀列表为空仍表示录制全部转码流。
	// :::tip 转码流和源流需至少选一个进行录制，即是否录制转码流（TranscodeRecord）和是否录制源流（OriginRecord）的取值至少一个不为 0。 :::
	TranscodeRecord *int32 `json:"TranscodeRecord,omitempty"`

	// 转码流后缀列表，转码流录制配置为根据转码流列表匹配（TranscodeRecord 取值为 2）时生效，TranscodeSuffixList 默认配置为空，效果等同于录制全部转码流。
	TranscodeSuffixList []*string `json:"TranscodeSuffixList,omitempty"`
}

// UpdateRecordPresetV2BodyRecordPresetConfigFlvParam - 录制为 FLV 格式时的录制参数。 :::tip 您需至少配置一个录制格式，即 FlvParam、HlsParam、Mp4Param
// 至少开启一个。 :::
type UpdateRecordPresetV2BodyRecordPresetConfigFlvParam struct {

	// 实时录制场景下，断流等待时长，单位为秒，默认值为 180，取值范围为 [0,3600]。如果实际断流时间小于断流等待时长，录制任务不会停止；如果实际断流时间大于断流等待时长，录制任务会停止，断流恢复后重新开始一个新的录制任务。
	ContinueDuration *int32 `json:"ContinueDuration,omitempty"`

	// 断流录制场景下，单文件录制时长，单位为秒，默认值为 7200，取值范围为 -1 和 [300,86400]。
	// * 取值为 -1 时，表示不限制录制时长，录制结束后生成一个完整的录制文件。
	// * 取值为 [300,86400] 之间的值时，表示根据设置的录制文件时长分段生成录制文件，完成录制后一起上传。
	// :::tip 断流录制场景仅在录制格式为 HLS 时生效，且断流录制和实时录制为二选一配置。 :::
	Duration *int32 `json:"Duration,omitempty"`

	// 当前格式的录制是否开启，默认 false，取值及含义如下所示。
	// * false：不开启；
	// * true：开启。
	Enable *bool `json:"Enable,omitempty"`

	// 实时录制场景下，单文件录制时长，单位为秒，默认值为 1800，取值范围为 [300,21600]。录制时间到达设置的单文件录制时长时，会立即生成录制文件实时上传存储。
	RealtimeRecordDuration *int32 `json:"RealtimeRecordDuration,omitempty"`

	// 断流录制场景下，断流拼接时长，单位为秒，默认值为 0，支持的取值及含义如下所示。
	// * -1：一直拼接，表示每次断流都不会影响录制任务，录制完成后生成一个完整的录制文件；
	// * 0：不拼接，表示每次断流结束录制任务生成一个录制文件，断流恢复重新开始一个新的录制任务；
	// * 大于 0：拼接容错时间，表示如果断流时间小于拼接容错时间时，则录制任务不会停止，不会生成新的录制文件；如果断流时间大于拼接容错时间，则录制任务停止，断流恢复后重新开始一个新的录制任务。
	// :::tip 断流录制场景仅在录制格式为 HLS 时生效，且断流录制和实时录制为二选一配置。 :::
	Splice *int32 `json:"Splice,omitempty"`

	// TOS 存储相关配置。 :::tip 录制文件只能选择一个位置进行存储，即 TOSParam 和 VODParam 配置且配置其中一个。 :::
	TOSParam *UpdateRecordPresetV2BodyRecordPresetConfigFlvParamTOSParam `json:"TOSParam,omitempty"`

	// VOD 存储相关配置。 :::tip 录制文件只能选择一个位置进行存储，即 TOSParam 和 VODParam 配置且配置其中一个。 :::
	VODParam *UpdateRecordPresetV2BodyRecordPresetConfigFlvParamVODParam `json:"VODParam,omitempty"`
}

// UpdateRecordPresetV2BodyRecordPresetConfigFlvParamTOSParam - TOS 存储相关配置。 :::tip 录制文件只能选择一个位置进行存储，即 TOSParam 和 VODParam
// 配置且配置其中一个。 :::
type UpdateRecordPresetV2BodyRecordPresetConfigFlvParamTOSParam struct {

	// TOS 存储对应的 Bucket。例如，存储位置为 live-test-tos-example/live/liveapp 时，Bucket 取值为 live-test-tos-example。 :::tip 如果使用 TOS 存储，即 TOSParam
	// 中 Enable 取值为 true 时，Bucket 为必填。 :::
	Bucket *string `json:"Bucket,omitempty"`

	// 是否使用 TOS 存储，默认为 false，取值及含义如下所示。
	// * false：不使用；
	// * true：使用。
	Enable *bool `json:"Enable,omitempty"`

	// 录制文件的存储规则，最大长度为 200 个字符，支持以record/{PubDomain}/{App}/{Stream}/{StartTime}_{EndTime} 样式设置存储规则，支持输入字母（A - Z、a - z）、数字（0 -
	// 9）、短横线（-）、叹号（!）、下划线（_）、句点（.）、星号（*）及占位符。
	// 存储规则设置注意事项如下。
	// * 目录层级至少包含2级及以上，如live/{App}/{Stream}。
	// * record 为自定义字段；
	// * {PubDomain} 取值为当前配置的 vhost 值；
	// * {App} 取值为当前配置的 AppName 值；
	// * {Stream} 取值为当前配置的 StreamName 值；
	// * {StartTime} 取值为录制的开始时间戳；
	// * {EndTime} 取值为录制的结束时间戳。
	ExactObject *string `json:"ExactObject,omitempty"`

	// TOS 存储对应 Bucket 下的存储目录，默认为空。例如，存储位置为 live-test-tos-example/live/liveapp 时，StorageDir 取值为 live/liveapp。
	StorageDir *string `json:"StorageDir,omitempty"`
}

// UpdateRecordPresetV2BodyRecordPresetConfigFlvParamVODParam - VOD 存储相关配置。 :::tip 录制文件只能选择一个位置进行存储，即 TOSParam 和 VODParam
// 配置且配置其中一个。 :::
type UpdateRecordPresetV2BodyRecordPresetConfigFlvParamVODParam struct {

	// 直播录制文件存储到点播时的视频分类 ID，您可以通过视频点播的ListVideoClassifications [https://www.volcengine.com/docs/4/101661]接口查询视频分类 ID 等信息，默认为空。
	ClassificationID *int32 `json:"ClassificationID,omitempty"`

	// 是否使用 VOD 存储，默认为 false，取值及含义如下所示。
	// * false：不使用；
	// * true：使用。
	Enable *bool `json:"Enable,omitempty"`

	// 录制文件的存储规则，最大长度为 200 个字符，支持以record/{PubDomain}/{App}/{Stream}/{StartTime}_{EndTime} 样式设置存储规则，支持输入字母（A - Z、a - z）、数字（0 -
	// 9）、短横线（-）、叹号（!）、下划线（_）、句点（.）、星号（*）及占位符。
	// 存储规则设置注意事项如下。
	// * 目录层级至少包含2级及以上，如live/{App}/{Stream}。
	// * record 为自定义字段；
	// * {PubDomain} 取值为当前配置的 vhost 值；
	// * {App} 取值为当前配置的 AppName 值；
	// * {Stream} 取值为当前配置的 StreamName 值；
	// * {StartTime} 取值为录制的开始时间戳；
	// * {EndTime} 取值为录制的结束时间戳。
	ExactObject *string `json:"ExactObject,omitempty"`

	// 直播录制文件存储到点播时的存储类型，存储类型介绍请参考媒资存储管理 [https://www.volcengine.com/docs/4/73629#媒资存储]。默认值为 1，支持的取值及含义如下所示。
	// * 1：标准存储；
	// * 2：归档存储。
	StorageClass *int32 `json:"StorageClass,omitempty"`

	// 视频点播（VOD）空间名称。可登录视频点播控制台 [https://console.volcengine.com/vod/]查询。 :::tip 如果使用 VOD 存储，即 VODParam 中 Enable 取值为 true 时，VodNamespace
	// 为必填。 :::
	VodNamespace *string `json:"VodNamespace,omitempty"`

	// 视频点播工作流模板 ID，对于存储在点播的录制文件，会使用该工作流模版对录制的视频进行处理，可登录视频点播控制台 [https://console.volcengine.com/vod/]获取工作流模板 ID，默认为空。
	WorkflowID *string `json:"WorkflowID,omitempty"`
}

// UpdateRecordPresetV2BodyRecordPresetConfigHlsParam - 录制为 HLS 格式时的录制参数。 :::tip 您需至少配置一个录制格式，即 FlvParam、HlsParam、Mp4Param
// 至少开启一个。 :::
type UpdateRecordPresetV2BodyRecordPresetConfigHlsParam struct {

	// 断流等待时长，取值范围[0, 3600]。
	ContinueDuration *int32 `json:"ContinueDuration,omitempty"`

	// 断流录制单文件录制时长，单位为 s，默认值为 7200，取值范围为 -1，[300,86400]，-1表示一直录制，目前只对 HLS 生效。
	Duration *int32 `json:"Duration,omitempty"`

	// 当前格式的录制是否开启，默认 false，取值及含义如下所示。
	// * false：不开启；
	// * true：开启。
	Enable *bool `json:"Enable,omitempty"`

	// 实时录制文件时长，单位为 s，取值范围为 [300,21600]。
	RealtimeRecordDuration *int32 `json:"RealtimeRecordDuration,omitempty"`

	// 断流拼接间隔时长，对实时录制无效，单位为 s，默认值为 0。支持的取值如下所示。
	// * -1：一直拼接；
	// * 0：不拼接；
	// * 大于 0：断流拼接时间间隔，对 HLS 录制生效。
	Splice *int32 `json:"Splice,omitempty"`

	// TOS 存储相关配置。 :::tipTOSParam和VODParam配置且配置其中一个。 :::
	TOSParam *UpdateRecordPresetV2BodyRecordPresetConfigHlsParamTOSParam `json:"TOSParam,omitempty"`

	// VOD 存储相关配置。 :::tipTOSParam和VODParam配置且配置其中一个。 :::
	VODParam *UpdateRecordPresetV2BodyRecordPresetConfigHlsParamVODParam `json:"VODParam,omitempty"`
}

// UpdateRecordPresetV2BodyRecordPresetConfigHlsParamTOSParam - TOS 存储相关配置。 :::tipTOSParam和VODParam配置且配置其中一个。 :::
type UpdateRecordPresetV2BodyRecordPresetConfigHlsParamTOSParam struct {

	// TOS 存储空间，一般使用 CDN 对应的 Bucket。 :::tip 如果 TOSParam 中的 Enable 取值为 true，则 Bucket 必填。 :::
	Bucket *string `json:"Bucket,omitempty"`

	// 是否使用 TOS 存储，默认为 false，取值及含义如下所示。
	// * false：不使用；
	// * true：使用。
	Enable *bool `json:"Enable,omitempty"`

	// 录制文件的存储位置。存储路径为record/{PubDomain}/{App}/{Stream}/{StartTime}_{EndTime}
	ExactObject *string `json:"ExactObject,omitempty"`

	// TOS 存储目录，默认为空。
	StorageDir *string `json:"StorageDir,omitempty"`
}

// UpdateRecordPresetV2BodyRecordPresetConfigHlsParamVODParam - VOD 存储相关配置。 :::tipTOSParam和VODParam配置且配置其中一个。 :::
type UpdateRecordPresetV2BodyRecordPresetConfigHlsParamVODParam struct {

	// 直播录制文件存储到点播时的视频分类 ID，您可以通过视频点播的ListVideoClassifications [https://www.volcengine.com/docs/4/101661]接口查询视频分类 ID 等信息。
	ClassificationID *int32 `json:"ClassificationID,omitempty"`

	// 是否使用 VOD 存储，默认为 false，取值及含义如下所示。
	// * false：不使用；
	// * true：使用。
	Enable *bool `json:"Enable,omitempty"`

	// 录制文件的存储位置，最大长度为 200 个字符。默认的存储位置为record/{PubDomain}/{App}/{Stream}/{StartTime}_{EndTime}，参数格式要求如下所示。
	// * 支持删除固定路径，如 {App}/{Stream}；
	// * 不支持以正斜线（/）或者反斜线（\）开头；
	// * 不支持 “//” 和 “/./” 等字符串；
	// * 不支持 \b、\t、\n、\v、\f、\r 等字符；
	// * 不支持 “..” 作为文件名；
	// * 目录层级至少包含 2 级及以上，如live/{App}/{Stream}。
	ExactObject *string `json:"ExactObject,omitempty"`

	// 直播录制文件存储到点播时的存储类型。默认值为 1，支持的取值及含义如下所示。
	// * 1：标准存储；
	// * 2：归档存储。
	StorageClass *int32 `json:"StorageClass,omitempty"`

	// 视频点播（VOD）空间名称。可登录视频点播控制台 [https://console.volcengine.com/vod/]查询。 :::tip 如果 VODParam 中的 Enable 取值为 true，则 VodNamespace
	// 必填。 :::
	VodNamespace *string `json:"VodNamespace,omitempty"`

	// 工作流模版 ID，对于存储在点播的录制文件，会使用该工作流模版对视频进行处理。可登录视频点播控制台 [https://console.volcengine.com/vod/]获取 ID。
	WorkflowID *string `json:"WorkflowID,omitempty"`
}

// UpdateRecordPresetV2BodyRecordPresetConfigMp4Param - 录制为 MP4 格式时的录制参数。 :::tip 您需至少配置一个录制格式，即 FlvParam、HlsParam、Mp4Param
// 至少开启一个。 :::
type UpdateRecordPresetV2BodyRecordPresetConfigMp4Param struct {

	// 断流等待时长，取值范围[0, 3600]。
	ContinueDuration *int32 `json:"ContinueDuration,omitempty"`

	// 断流录制单文件录制时长，单位为 s，默认值为 7200，取值范围为 -1，[300,86400]，-1表示一直录制，目前只对 HLS 生效。
	Duration *int32 `json:"Duration,omitempty"`

	// 当前格式的录制是否开启，默认 false，取值及含义如下所示。
	// * false：不开启；
	// * true：开启。
	Enable *bool `json:"Enable,omitempty"`

	// 实时录制文件时长，单位为 s，取值范围为 [300,21600]。
	RealtimeRecordDuration *int32 `json:"RealtimeRecordDuration,omitempty"`

	// 断流拼接间隔时长，对实时录制无效，单位为 s，默认值为 0。支持的取值如下所示。
	// * -1：一直拼接；
	// * 0：不拼接；
	// * 大于 0：断流拼接时间间隔，对 HLS 录制生效。
	Splice *int32 `json:"Splice,omitempty"`

	// TOS 存储相关配置。 :::tipTOSParam和VODParam配置且配置其中一个。 :::
	TOSParam *UpdateRecordPresetV2BodyRecordPresetConfigMp4ParamTOSParam `json:"TOSParam,omitempty"`

	// VOD 存储相关配置。 :::tipTOSParam和VODParam配置且配置其中一个。 :::
	VODParam *UpdateRecordPresetV2BodyRecordPresetConfigMp4ParamVODParam `json:"VODParam,omitempty"`
}

// UpdateRecordPresetV2BodyRecordPresetConfigMp4ParamTOSParam - TOS 存储相关配置。 :::tipTOSParam和VODParam配置且配置其中一个。 :::
type UpdateRecordPresetV2BodyRecordPresetConfigMp4ParamTOSParam struct {

	// TOS 存储空间，一般使用 CDN 对应的 Bucket。 :::tip 如果 TOSParam 中的 Enable 取值为 true，则 Bucket 必填。 :::
	Bucket *string `json:"Bucket,omitempty"`

	// 是否使用 TOS 存储，默认为 false，取值及含义如下所示。
	// * false：不使用；
	// * true：使用。
	Enable *bool `json:"Enable,omitempty"`

	// 录制文件的存储位置。存储路径为record/{PubDomain}/{App}/{Stream}/{StartTime}_{EndTime}
	ExactObject *string `json:"ExactObject,omitempty"`

	// TOS 存储目录，默认为空。
	StorageDir *string `json:"StorageDir,omitempty"`
}

// UpdateRecordPresetV2BodyRecordPresetConfigMp4ParamVODParam - VOD 存储相关配置。 :::tipTOSParam和VODParam配置且配置其中一个。 :::
type UpdateRecordPresetV2BodyRecordPresetConfigMp4ParamVODParam struct {

	// 直播录制文件存储到点播时的视频分类 ID，您可以通过视频点播的ListVideoClassifications [https://www.volcengine.com/docs/4/101661]接口查询视频分类 ID 等信息。
	ClassificationID *int32 `json:"ClassificationID,omitempty"`

	// 是否使用 VOD 存储，默认为 false，取值及含义如下所示。
	// * false：不使用；
	// * true：使用。
	Enable *bool `json:"Enable,omitempty"`

	// 录制文件的存储位置，最大长度为 200 个字符。默认的存储位置为record/{PubDomain}/{App}/{Stream}/{StartTime}_{EndTime}，参数格式要求如下所示。
	// * 支持删除固定路径，如 {App}/{Stream}；
	// * 不支持以正斜线（/）或者反斜线（\）开头；
	// * 不支持 “//” 和 “/./” 等字符串；
	// * 不支持 \b、\t、\n、\v、\f、\r 等字符；
	// * 不支持 “..” 作为文件名；
	// * 目录层级至少包含 2 级及以上，如live/{App}/{Stream}。
	ExactObject *string `json:"ExactObject,omitempty"`

	// 直播录制文件存储到点播时的存储类型。默认值为 1，支持的取值及含义如下所示。
	// * 1：标准存储；
	// * 2：归档存储。
	StorageClass *int32 `json:"StorageClass,omitempty"`

	// 视频点播（VOD）空间名称。可登录视频点播控制台 [https://console.volcengine.com/vod/]查询。 :::tip 如果 VODParam 中的 Enable 取值为 true，则 VodNamespace
	// 必填。 :::
	VodNamespace *string `json:"VodNamespace,omitempty"`

	// 工作流模版 ID，对于存储在点播的录制文件，会使用该工作流模版对视频进行处理。可登录视频点播控制台 [https://console.volcengine.com/vod/]获取 ID。
	WorkflowID *string `json:"WorkflowID,omitempty"`
}

type UpdateRecordPresetV2Res struct {

	// REQUIRED
	ResponseMetadata UpdateRecordPresetV2ResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateRecordPresetV2ResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                        `json:"Version"`
	Error   *UpdateRecordPresetV2ResResponseMetadataError `json:"Error,omitempty"`
}

type UpdateRecordPresetV2ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateRefererBody struct {

	// REQUIRED; Referer 防盗链规则列表。 :::tip
	// * 同一个 Vhost 下，默认支持配置不超过 100 个 Referer 规则，如需提升限额请创建工单 [https://console.volcengine.com/workorder/create?step=2&SubProductID=P00000076]获取技术支持；
	// * 单次请求最多支持配置 100 个 Referer 规则。 :::
	RefererInfoList []UpdateRefererBodyRefererInfoListItem `json:"RefererInfoList"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看需要查询的直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 应用名称，取值与直播流地址中 AppName 字段取值相同，默认为空，表示所有应用名称。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。 :::tip
	// 参数 Domain 和 App 传且仅传一个。 :::
	App *string `json:"App,omitempty"`

	// 拉流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的拉流域名。
	// :::tip 参数 Domain 和 App 传且仅传一个。 :::
	Domain *string `json:"Domain,omitempty"`
}

type UpdateRefererBodyRefererInfoListItem struct {

	// REQUIRED; 用于标识 referer 防盗链的关键词默认取值为 referer。
	Key string `json:"Key"`

	// REQUIRED; Referer 字段规则类型，取值即含义如下所示。
	// * deny：拒绝，即黑名单；
	// * allow：通过，即白名单。
	Type string `json:"Type"`

	// Referer 字段规则的匹配优先级，默认为 0，取值范围为 [0,100]，数值越大，优先级越高。如果优先级相同，则越早加入列表的域名优先级越高。
	Priority *int32 `json:"Priority,omitempty"`

	// Referer 字段规则，即设置的黑名单或白名单的域名，最大长度限制 300 个字符。
	Value *string `json:"Value,omitempty"`
}

type UpdateRefererRes struct {

	// REQUIRED
	ResponseMetadata UpdateRefererResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type UpdateRefererResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                 `json:"Version"`
	Error   *UpdateRefererResResponseMetadataError `json:"Error,omitempty"`
}

type UpdateRefererResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateRegionAccessRuleBody struct {

	// REQUIRED
	Domain string `json:"Domain"`

	// REQUIRED
	RegionAccessRule UpdateRegionAccessRuleBodyRegionAccessRule `json:"RegionAccessRule"`

	// REQUIRED
	Vhost string  `json:"Vhost"`
	App   *string `json:"App,omitempty"`
}

type UpdateRegionAccessRuleBodyRegionAccessRule struct {

	// REQUIRED
	Enable string `json:"Enable"`

	// REQUIRED
	Type         string    `json:"Type"`
	CountryList  []*string `json:"CountryList,omitempty"`
	ProvinceList []*string `json:"ProvinceList,omitempty"`
}

type UpdateRegionAccessRuleRes struct {

	// REQUIRED
	ResponseMetadata UpdateRegionAccessRuleResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateRegionAccessRuleResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type UpdateRelaySourceRewriteBody struct {

	// REQUIRED; 域名空间名称
	Vhost string `json:"Vhost"`

	// 需要设置黑白名单的拉流域名。域名请在工信部完成备案。
	Domain *string `json:"Domain,omitempty"`

	// 改写规则
	RewriteRule *UpdateRelaySourceRewriteBodyRewriteRule `json:"RewriteRule,omitempty"`
}

// UpdateRelaySourceRewriteBodyRewriteRule - 改写规则
type UpdateRelaySourceRewriteBodyRewriteRule struct {

	// REQUIRED; 功能开关。- true: 开 - false: 关
	Enable bool `json:"Enable"`

	// REQUIRED; 改写规则列表
	RewriteRuleList []UpdateRelaySourceRewriteBodyRewriteRuleListItem `json:"RewriteRuleList"`
}

type UpdateRelaySourceRewriteBodyRewriteRuleListItem struct {

	// 改写后地址是否包含原始地址的param参数
	IncludeParams *bool `json:"IncludeParams,omitempty"`

	// 原始path
	OriginPath *string `json:"OriginPath,omitempty"`

	// 改写后目标path
	TargetPath *string `json:"TargetPath,omitempty"`
}

type UpdateRelaySourceRewriteRes struct {

	// REQUIRED
	ResponseMetadata UpdateRelaySourceRewriteResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; Anything
	Result interface{} `json:"Result"`
}

type UpdateRelaySourceRewriteResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string `json:"Version"`
}

type UpdateRelaySourceV3Body struct {

	// REQUIRED; 回源组配置详情。
	GroupDetails []UpdateRelaySourceV3BodyGroupDetailsItem `json:"GroupDetails"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 应用名称，即直播流地址的AppName字段取值，默认为空，表示为当前域名空间的全局播放触发回源配置。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App *string `json:"App,omitempty"`

	// 生效类型（order/rand/hot），不填默认order
	LBType *string `json:"LBType,omitempty"`

	// 组的重试间隔/s，不填默认为1 s。
	RetryInterval *string `json:"RetryInterval,omitempty"`

	// 组的重试次数，不填默认等于group数-1
	RetryTimes *string `json:"RetryTimes,omitempty"`
}

type UpdateRelaySourceV3BodyGroupDetailsItem struct {

	// REQUIRED; 回源组名称。
	Group string `json:"Group"`

	// REQUIRED; 回源服务器配置列表。
	Servers    []UpdateRelaySourceV3BodyGroupDetailsPropertiesItemsItem `json:"Servers"`
	AuthParams *UpdateRelaySourceV3BodyGroupDetailsItemAuthParams       `json:"AuthParams,omitempty"`

	// 主还是备
	Backup *bool `json:"Backup,omitempty"`

	// CDN类型
	CDN *string `json:"CDN,omitempty"`

	// 应用层超时时间/s
	ConnTimeout *string `json:"ConnTimeout,omitempty"`

	// 传输层超时时间/s （不配置的话使用应用层时间）
	DialTimeout *string `json:"DialTimeout,omitempty"`

	// 是否禁用
	Disable *bool `json:"Disable,omitempty"`

	// server生效类型（order/rand/hot）
	LBType *string `json:"LBType,omitempty"`

	// 主secret key
	PrimarySK *string `json:"PrimarySK,omitempty"`

	// 不上火山
	PullAuth *bool `json:"PullAuth,omitempty"`

	// server重试间隔/s
	RetryInterval *string `json:"RetryInterval,omitempty"`

	// server重试次数
	RetryTimes *string `json:"RetryTimes,omitempty"`

	// 回源规则
	Rule *string `json:"Rule,omitempty"`

	// 副secret key
	SecondSK *string `json:"SecondSK,omitempty"`

	// Group 超时时间/s
	Timeout *string `json:"Timeout,omitempty"`

	// 鉴权有效时长
	ValidDuration *string `json:"ValidDuration,omitempty"`

	// 权重
	Weight *string `json:"Weight,omitempty"`
}

type UpdateRelaySourceV3BodyGroupDetailsItemAuthParams struct {

	// 鉴权参数名，如“sign”
	VolcSecret *string `json:"VolcSecret,omitempty"`

	// 有效期，如"expire"
	VolcTime *string `json:"VolcTime,omitempty"`
}

type UpdateRelaySourceV3BodyGroupDetailsItemServersItemOutboundConfig struct {

	// 代理配置列表，不传默认不使用代理
	ProxyConfigList []*UpdateRelaySourceV3BodyGroupDetailsPropertiesItemsOutboundConfigProxyConfigListItem `json:"ProxyConfigList,omitempty"`

	// 代理模式，0：固定模式，1: 解析模式，2：默认模式
	ProxyMode *string `json:"ProxyMode,omitempty"`
}

type UpdateRelaySourceV3BodyGroupDetailsPropertiesItemsItem struct {

	// REQUIRED; 直播源服务器的地址，支持填写回源服务的域名或 IP 地址。 :::tip
	// * 当源站使用了非默认端口时，支持在回源地址中以域名:端口或IP:端口的形式配置端口。
	// * 最多支持添加 10 个回源地址，回源失败时，将按照您添加的地址顺序轮循尝试。 :::
	RelaySourceDomain string `json:"RelaySourceDomain"`

	// REQUIRED; 回源协议，支持两种回源协议。
	// * rtmp：RTMP 回源协议；
	// * flv：FLV 回源协议。
	RelaySourceProtocol string `json:"RelaySourceProtocol"`

	// 回源Host
	Host           *string                                                           `json:"Host,omitempty"`
	OutboundConfig *UpdateRelaySourceV3BodyGroupDetailsItemServersItemOutboundConfig `json:"OutboundConfig,omitempty"`

	// 自定义回源参数，缺省情况下为空。格式为"Key":"Value"，例如，"domain":"live.push.net"。
	RelaySourceParams map[string]*string `json:"RelaySourceParams,omitempty"`

	// 权重
	Weight *string `json:"Weight,omitempty"`
}

type UpdateRelaySourceV3BodyGroupDetailsPropertiesItemsOutboundConfigProxyConfigListItem struct {

	// 集群
	Cluster *string `json:"Cluster,omitempty"`

	// 机房
	IDC *string `json:"IDC,omitempty"`

	// 运营商
	ISP *string `json:"ISP,omitempty"`

	// 代理列表
	ProxyList []*UpdateRelaySourceV3BodyGroupDetailsPropertiesItemsOutboundConfigProxyConfigListPropertiesItemsItem `json:"ProxyList,omitempty"`
}

type UpdateRelaySourceV3BodyGroupDetailsPropertiesItemsOutboundConfigProxyConfigListPropertiesItemsItem struct {

	// 代理地址
	URL *string `json:"URL,omitempty"`

	// 权重
	Weight *string `json:"Weight,omitempty"`
}

type UpdateRelaySourceV3Res struct {

	// REQUIRED
	ResponseMetadata UpdateRelaySourceV3ResResponseMetadata `json:"ResponseMetadata"`
}

type UpdateRelaySourceV3ResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                       `json:"Version"`
	Error   *UpdateRelaySourceV3ResResponseMetadataError `json:"Error,omitempty"`
}

type UpdateRelaySourceV3ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateRelaySourceV4Body struct {

	// REQUIRED; 应用名称，拉流域名，您可以调用ListRelaySourceV4 [https://www.volcengine.com/docs/6469/1126878]接口，获取待更新固定回源配置的 App 取值。
	App string `json:"App"`

	// REQUIRED; 拉流域名，您可以调用ListRelaySourceV4 [https://www.volcengine.com/docs/6469/1126878]接口，获取待更新固定回源配置的 Domain 取值。
	Domain string `json:"Domain"`

	// REQUIRED; 回源地址列表，支持 RTMP、FLV、HLS 和 SRT 回源协议。
	// :::tip
	// * 当源站使用了非默认端口时，支持在回源地址中以域名:端口或IP:端口的形式配置端口。
	// * 最多支持添加 10 个回源地址，回源失败时，将按照您添加的地址顺序轮循尝试。 :::
	SrcAddrS []string `json:"SrcAddrS"`

	// REQUIRED; 流名称，您可以调用ListRelaySourceV4 [https://www.volcengine.com/docs/6469/1126878]接口，获取待更新固定回源配置的 Domain 取值。
	Stream string `json:"Stream"`

	// 结束时间，Unix 时间戳，单位为秒。 :::tip
	// * 回源开始到结束最大时间跨度为 7 天；
	// * 开始时间与结束时间同时缺省，表示永久回源。 :::
	EndTime *int32 `json:"EndTime,omitempty"`

	// 开始时间，Unix 时间戳，单位为秒。 :::tip
	// * 回源开始到结束最大时间跨度为 7 天；
	// * 开始时间与结束时间同时缺省，表示永久回源。 :::
	StartTime *int32 `json:"StartTime,omitempty"`
}

type UpdateRelaySourceV4Res struct {

	// REQUIRED
	ResponseMetadata UpdateRelaySourceV4ResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateRelaySourceV4ResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                       `json:"Version"`
	Error   *UpdateRelaySourceV4ResResponseMetadataError `json:"Error,omitempty"`
}

type UpdateRelaySourceV4ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateSnapshotPresetBody struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 截图模板名称。
	Preset string `json:"Preset"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// ToS 的存储 Bucket。 :::tipBucket 与 ServiceID 传且仅传一个。 :::
	Bucket *string `json:"Bucket,omitempty"`

	// 回调详情。
	CallbackDetailList []*UpdateSnapshotPresetBodyCallbackDetailListItem `json:"CallbackDetailList,omitempty"`

	// 截图间隔时间，单位为 s，默认值为 10，取值范围为正整数
	Interval *int32 `json:"Interval,omitempty"`

	// 存储方式为覆盖截图时的存储规则，支持以 {Domain}/{App}/{Stream} 样式设置存储规则，支持输入字母、数字、"-"、"!"、"_"、"."、"*"及占位符。
	OverwriteObject *string `json:"OverwriteObject,omitempty"`

	// veImageX 的服务 ID。 :::tipBucket 与 ServiceID 传且仅传一个。 :::
	ServiceID *string `json:"ServiceID,omitempty"`

	// 截图格式。支持如下取值。- jpeg - jpg
	SnapshotFormat *string `json:"SnapshotFormat,omitempty"`

	// 存储方式为实时存储时的存储规则，支持以 {Domain}/{App}/{Stream}/{UnixTimestamp} 样式设置存储规则，支持输入字母、数字、"-"、"!"、"_"、"."、"*"及占位符。
	SnapshotObject *string `json:"SnapshotObject,omitempty"`

	// 截图模版状态。
	// * 1：开启
	// * 0：关闭
	Status *int32 `json:"Status,omitempty"`

	// ToS 的存储目录，不传为空。
	StorageDir *string `json:"StorageDir,omitempty"`
}

type UpdateSnapshotPresetBodyCallbackDetailListItem struct {

	// 回调类型，默认值为 http。
	CallbackType *string `json:"CallbackType,omitempty"`

	// 回调地址。
	URL *string `json:"URL,omitempty"`
}

type UpdateSnapshotPresetRes struct {

	// REQUIRED
	ResponseMetadata UpdateSnapshotPresetResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type UpdateSnapshotPresetResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                        `json:"Version"`
	Error   *UpdateSnapshotPresetResResponseMetadataError `json:"Error,omitempty"`
}

type UpdateSnapshotPresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateTimeShiftPresetV3Body struct {

	// REQUIRED; 应用名称，您可以调用ListTimeShiftPresetV2 [https://www.volcengine.com/docs/6469/1126883]接口，获取待更新时移配置的App取值。
	App string `json:"App"`

	// REQUIRED; 最大时移时长，即允许用户回看的最长时间，单位为秒，支持的取值如下所示。
	// * 86400：1 天；
	// * 259200：3 天；
	// * 604800：7 天；
	// * 1296000：15 天。
	MaxShiftTime int32 `json:"MaxShiftTime"`

	// REQUIRED; 域名空间名称，您可以调用ListTimeShiftPresetV2 [https://www.volcengine.com/docs/6469/1126883]接口，获取待更新时移配置的Vhost取值。
	Vhost string `json:"Vhost"`

	// 用于多码率时移的参数，为json字符串
	MasterFormat *string `json:"MasterFormat,omitempty"`

	// 0表示不需要 1表示需要
	NeedTranscode *int32 `json:"NeedTranscode,omitempty"`

	// nss配置
	NssConfig *string `json:"NssConfig,omitempty"`

	// 操作类型，不填默认更新关联，only_preset: 只更新模板配置，associate: 更新模板和关联
	OperationType *string `json:"OperationType,omitempty"`

	// 模板名称
	PresetName   *string `json:"PresetName,omitempty"`
	RecordObject *string `json:"RecordObject,omitempty"`
	Status       *int32  `json:"Status,omitempty"`

	// 开启时移的流名称，默认为空表示更新 App 级别的时移配置，不为空时表示更新 Stream 级别的时移配置。您可以调用ListTimeShiftPresetV2 [https://www.volcengine.com/docs/6469/1126883]接口，获取待更新时移配置的Stream取值并进行更新。
	Stream *string `json:"Stream,omitempty"`
}

type UpdateTimeShiftPresetV3Res struct {

	// REQUIRED
	ResponseMetadata UpdateTimeShiftPresetV3ResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateTimeShiftPresetV3ResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                           `json:"Version"`
	Error   *UpdateTimeShiftPresetV3ResResponseMetadataError `json:"Error,omitempty"`
}

type UpdateTimeShiftPresetV3ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateTranscodePresetBody struct {

	// REQUIRED; 转码配置的名称，您可以调用ListVhostTransCodePreset [https://www.volcengine.com/docs/6469/1126853]接口查看待更新转码配置的Preset取值。
	Preset string `json:"Preset"`

	// REQUIRED; 域名空间，您可以调用 ListVhostTransCodePreset [https://www.volcengine.com/docs/6469/1126853] 接口查看待更新转码配置的 Vhost 取值。
	Vhost     string  `json:"Vhost"`
	ALayout   *string `json:"ALayout,omitempty"`
	AProfile  *string `json:"AProfile,omitempty"`
	AR        *int32  `json:"AR,omitempty"`
	AbrMode   *int32  `json:"AbrMode,omitempty"`
	AccountID *string `json:"AccountID,omitempty"`

	// 音频编码格式，默认值为aac，支持的取值及含义如下所示。
	// * aac：使用 AAC 音频编码格式；
	// * opus：使用 Opus 音频编码格式。
	// * copy：不进行音频转码，所有音频编码参数不生效，音频编码参数包括音频码率（AudioBitrate）等。
	Acodec         *string `json:"Acodec,omitempty"`
	AdvancedParam  *string `json:"AdvancedParam,omitempty"`
	AllowAudioCopy *int32  `json:"AllowAudioCopy,omitempty"`
	AllowVideoCopy *int32  `json:"AllowVideoCopy,omitempty"`
	An             *int32  `json:"An,omitempty"`

	// 应用名称，取值与直播流地址的 AppName 字段取值相同，您可以调用 ListVhostTransCodePreset [https://www.volcengine.com/docs/6469/1126853] 接口查看待更新转码配置的
	// App 取值。
	App *string `json:"App,omitempty"`

	// 视频分辨率自适应模式开关，默认值为 0。支持的取值及含义如下。
	// * 0：关闭视频分辨率自适应；
	// * 1：开启视频分辨率自适应。 :::tip
	// * 关闭视频分辨率自适应模式（As 取值为 0）时，转码配置的视频分辨率取视频宽度（Width）和视频高度（Height）的值对转码视频进行拉伸；
	// * 开启视频分辨率自适应模式（As 取值为 1）时，转码配置的视频分辨率按照短边长度（ShortSide）、长边长度（LongSide）、视频宽度（Width）、视频高度（Height）的优先级取值，另一边等比缩放。 :::
	As *string `json:"As,omitempty"`

	// 音频码率，单位为 kbps，默认值为128，取值范围为 [0,1000]；取值为0时，表示与源流的音频码率相同。
	AudioBitrate *int32 `json:"AudioBitrate,omitempty"`
	AutoTransAb  *int32 `json:"AutoTransAb,omitempty"`
	AutoTransAl  *int32 `json:"AutoTransAl,omitempty"`
	AutoTransAr  *int32 `json:"AutoTransAr,omitempty"`

	// 是否开启转码视频分辨率不超过源流分辨率，默认值为 1 表示开启。开启后，当源流分辨率低于转码配置分辨率时（即源流宽低于转码配置宽且源流高低于转码配置高时），将按源流视频分辨率进行转码。
	// * 0：关闭；
	// * 1：开启。
	AutoTransResolution *int32 `json:"AutoTransResolution,omitempty"`

	// 是否开启转码视频码率不超过源流码率，默认值为 1 表示开启。开启后，当源流码率低于转码配置码率时，将按照源流视频码率进行转码。
	// * 0：关闭；
	// * 1：开启。
	AutoTransVb *int32 `json:"AutoTransVb,omitempty"`

	// 是否开启转码视频帧率不超过源流帧率，默认值为 1 表示开启。开启后，当源流帧率低于转码配置帧率时，将按照源流视频帧率进行转码。
	// * 0：关闭；
	// * 1：开启。
	AutoTransVr *int32 `json:"AutoTransVr,omitempty"`
	BCM         *int32 `json:"BCM,omitempty"`

	// 转码输出视频中 2 个参考帧之间的最大 B 帧数量，默认值为 3，取值为 0 时表示去除 B 帧。
	// 最大 B 帧数量的取值范围根据视频编码格式（Vcodec）的不同有所差异，取值范围如下所示。
	// * 视频编码格式为 H.264 （Vcodec 取值为 h264）时取值范围为 [0,7]；
	// * 视频编码格式为 H.265 或 H.266 （Vcodec 取值为 h265 或 h266）时取值范围为 [0,3]、7、15。
	BFrames  *int32  `json:"BFrames,omitempty"`
	Describe *string `json:"Describe,omitempty"`

	// 动态范围，画质增强类型生效
	// * SDR：输出为SDR
	// * HDR：输出为HDR
	DynamicRange *string `json:"DynamicRange,omitempty"`

	// 是否开启智能插帧，只对画质增强类型生效
	// * 0：不开启
	// * 1：开启
	FISwitch *int32 `json:"FISwitch,omitempty"`

	// 视频帧率，单位为 fps，默认值为 25，取值为 0 时表示与源流视频帧率相同。
	// 视频帧率的取值范围根据视频编码格式（Vcodec）的不同有所差异，视频码率的取值范围如下所示。
	// * 视频编码格式为 H.264 或 H.265 （Vcodec 取值为 h264 或 h265）时，视频帧率取值范围为 [0,60]；
	// * 视频编码格式为 H.266 （Vcodec 取值为 h266）时，视频帧率取值范围为 [0,35]。
	FPS *int32 `json:"FPS,omitempty"`

	// IDR 帧之间的最大间隔时间，单位为秒，默认值为 4，取值范围为 [1,20]。
	GOP    *int32 `json:"GOP,omitempty"`
	GopMin *int32 `json:"GopMin,omitempty"`
	HVSPre *bool  `json:"HVSPre,omitempty"`

	// 视频高度，默认值为 0。
	// 视频高度的取值范围根据视频编码格式（Vcodec）的不同所有差异，视频高度取值如下所示。
	// * 视频编码格式为 H.264 或 H.265 （Vcodec 取值为 h264 或 h265）时，取值范围为 [150,1920]；
	// * 视频编码格式为 H.266 （Vcodec 取值为 h266）时，不支持设置 Width 和 Height。
	// :::tip
	// * 当关闭视频分辨率自适应（As 取值为 0）时，转码分辨率将取 Width 和 Height 的值对转码视频进行拉伸；
	// * 当关闭视频分辨率自适应（As 取值为 0）时，Width 和 Height 任一取值为 0 时，转码视频将保持源流尺寸。 :::
	Height *int32 `json:"Height,omitempty"`

	// 长边长度，默认值为 0。配置不同的转码类型（Roi）和视频编码方式（Vcodec）时，短边长度的取值范围存在如下。
	// * 转码类型为标准转码（Roi 取值为 false）时： * 视频编码方式为 H.264 （Vcodec 取值为 h264）时取值范围为 0 和 [150,4096]；
	// * 视频编码方式为 H.265 （Vcodec 取值为 h265）时取值范围为 0 和 [150,7680]；
	// * 视频编码方式为 H.266 （Vcodec 取值为 h266）时取值范围为 0 和 [150,1280]。
	//
	//
	// * 转码类型为极智超清转码（Roi 取值为 true）时： * 视频编码方式为 H.264 或 H.265 （Vcodec 取值为 h264 或 h265）时取值范围为 0 和 [150,1920]。
	//
	//
	// :::tip
	// * 当开启视频分辨率自适应模式时（As 取值为 1）时，参数生效，反之则不生效。
	// * 当开启视频分辨率自适应模式时（As 取值为 1）时，如果 LongSide 、 ShortSide 、Width 、Height 同时取 0，表示保持源流尺寸。 :::
	LongSide    *int32  `json:"LongSide,omitempty"`
	LookAhead   *int32  `json:"LookAhead,omitempty"`
	Modifier    *string `json:"Modifier,omitempty"`
	NvBf        *int32  `json:"NvBf,omitempty"`
	NvCodec     *string `json:"NvCodec,omitempty"`
	NvGop       *int32  `json:"NvGop,omitempty"`
	NvHVSPre    *bool   `json:"NvHVSPre,omitempty"`
	NvLookahead *int32  `json:"NvLookahead,omitempty"`
	NvPercent   *int32  `json:"NvPercent,omitempty"`
	NvPreset    *string `json:"NvPreset,omitempty"`
	NvPriority  *int32  `json:"NvPriority,omitempty"`
	NvProfile   *string `json:"NvProfile,omitempty"`
	NvRefs      *int32  `json:"NvRefs,omitempty"`
	NvTempAQ    *int32  `json:"NvTempAQ,omitempty"`
	Ocr         *bool   `json:"Ocr,omitempty"`

	// 转码模板参数的类型
	// * hvq：表示使用画质增强
	// 选择画质增强时，支持使用 shortside 来设置分辨率。
	// * ParamType 取 hvq 时： * 视频编码方式为 H.264 （Vcodec 取值为 h264）时，shortside 取值范围为 0 和 [150,1280]；
	// * 视频编码方式为 H.265 （Vcodec取值为h265）是，shortside 取值范围为 0 和 [150,1280]；
	ParamType    *string `json:"ParamType,omitempty"`
	PresetKind   *int32  `json:"PresetKind,omitempty"`
	PresetType   *int32  `json:"PresetType,omitempty"`
	Qp           *int32  `json:"Qp,omitempty"`
	RegionConfig *string `json:"RegionConfig,omitempty"`
	Revision     *string `json:"Revision,omitempty"`

	// 转码类型是否为极智超清转码，默认值为 false，取值及含义如下。
	// * true：极智超清转码；
	// * false：标准转码。
	// :::tip 视频编码格式为 H.266 （Vcodec取值为h266）时，转码类型不支持极智超清转码。 :::
	Roi  *bool `json:"Roi,omitempty"`
	SITI *bool `json:"SITI,omitempty"`

	// 使用场景，画质增强时生效
	// football：足球场景
	SceneType *string `json:"SceneType,omitempty"`

	// 短边长度，默认值为 0。配置不同的转码类型（Roi）和视频编码方式（Vcodec）时，短边长度的取值范围存在如下。
	// * 转码类型为标准转码（Roi 取值为 false）时： * 视频编码方式为 H.264 （Vcodec 取值为 h264）时取值范围为 0 和 [150,2160]；
	// * 视频编码方式为 H.265 （Vcodec 取值为 h265）时取值范围为 0 和 [150,4096]；
	// * 视频编码方式为 H.266 （Vcodec 取值为 h266）时取值范围为 0 和 [150,720]。
	//
	//
	// * 转码类型为极智超清转码（Roi 取值为 true）时： * 视频编码方式为 H.264 或 H.265 （Vcodec 取值为 h264 或 h265）时取值范围为 0 和 [150,1920]。 :::tip
	//
	//
	// * 当开启视频分辨率自适应模式（As 取值为 1）时，参数生效，反之则不生效。
	// * 当开启视频分辨率自适应模式（As 取值为 1）时，如果 LongSide 、 ShortSide 、Width 、Height 同时取 0，表示保持源流尺寸。 :::
	ShortSide *int32 `json:"ShortSide,omitempty"`
	Status    *int32 `json:"Status,omitempty"`

	// 转码停止时长，支持触发方式为拉流转码（TransType 取值为 Pull）时设置，表示断开拉流后转码停止的时长，单位为秒，取值范围为 -1 和 [0,300]，-1 表示不停止转码，默认值为 60。
	StopInterval *int32 `json:"StopInterval,omitempty"`

	// 转码后缀，支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）和短横线（-）组成，长度为 1 到 10 个字符。
	// 转码后缀通常以流名称后缀的形式来使用，常见的标识有 _sd、_hd、_uhd，例如，当转码配置的标识为 _hd 时，拉取转码流时转码流的流名名称为 源流的流名称_hd。
	SuffixName *string `json:"SuffixName,omitempty"`
	Threads    *int32  `json:"Threads,omitempty"`

	// 转码触发方式，默认值为 Pull，支持的取值及含义如下。
	// * Push：推流转码，直播推流后会自动启动转码任务，生成转码流；
	// * Pull：拉流转码，直播推流后，需要主动播放转码流才会启动转码任务，生成转码流。
	TransType       *string                                   `json:"TransType,omitempty"`
	TranscodeStruct *UpdateTranscodePresetBodyTranscodeStruct `json:"TranscodeStruct,omitempty"`
	VBRatio         *int32                                    `json:"VBRatio,omitempty"`
	VBVBufSize      *int32                                    `json:"VBVBufSize,omitempty"`
	VBVMaxRate      *int32                                    `json:"VBVMaxRate,omitempty"`
	VLevel          *string                                   `json:"VLevel,omitempty"`
	VPreset         *string                                   `json:"VPreset,omitempty"`
	VProfile        *string                                   `json:"VProfile,omitempty"`
	VR              *int32                                    `json:"VRVr,omitempty"`
	VRBBframes      *int32                                    `json:"VRBBframes,omitempty"`
	VRBHeightNum    *int32                                    `json:"VRBHeightNum,omitempty"`
	VRBPreset       *string                                   `json:"VRBPreset,omitempty"`
	VRBProfile      *string                                   `json:"VRBProfile,omitempty"`
	VRBSuffix       *string                                   `json:"VRBSuffix,omitempty"`
	VRBVb           *int32                                    `json:"VRBVb,omitempty"`
	VRBWidthNum     *int32                                    `json:"VRBWidthNum,omitempty"`
	VRGop           *int32                                    `json:"VRGop,omitempty"`
	VRGopDen        *int32                                    `json:"VRGopDen,omitempty"`
	VRHvspre        *string                                   `json:"VRHvspre,omitempty"`
	VRProjection    *string                                   `json:"VRProjection,omitempty"`
	VRRoi           *string                                   `json:"VRRoi,omitempty"`
	VRTBframes      *int32                                    `json:"VRTBframes,omitempty"`
	VRTPreset       *string                                   `json:"VRTPreset,omitempty"`
	VRTProfile      *string                                   `json:"VRTProfile,omitempty"`
	VRTSuffix       *string                                   `json:"VRTSuffix,omitempty"`
	VRTVb           *int32                                    `json:"VRTVb,omitempty"`
	VRTileMod       *int32                                    `json:"VRTileMod,omitempty"`
	VRateCtrl       *string                                   `json:"VRateCtrl,omitempty"`
	VbThreshold     *string                                   `json:"VbThreshold,omitempty"`
	Vclass          *bool                                     `json:"Vclass,omitempty"`

	// 视频编码格式，支持的取值及含义如下所示。
	// * h264：使用 H.264 视频编码格式；
	// * h265：使用 H.265 视频编码格式；
	// * h266：使用 H.266 视频编码格式；
	// * copy：不进行视频转码，所有视频编码参数不生效，视频编码参数包括视频帧率（FPS）、视频码率（VideoBitrate）、分辨率设置（As、Width、Height、ShortSide、LongSide）、GOP 和 BFrames
	// 等。
	Vcodec *string `json:"Vcodec,omitempty"`

	// 视频码率，单位为 bps，默认值为 1000000；取值为 0 时，表示与源流的视频码率相同。
	// 视频码率的取值范围根据视频编码格式（Vcodec）的不同有所差异，视频码率的取值范围如下所示。
	// * 视频编码格式为 H.264 或 H.265 （Vcodec 取值为 h264 或 h265）时，视频码率取值范围为 [0,30000000]；
	// * 视频编码格式为 H.266 （Vcodec 取值为 h266）时，视频码率取值范围为 [0,6000000]。
	VideoBitrate *int32  `json:"VideoBitrate,omitempty"`
	Vn           *int32  `json:"Vn,omitempty"`
	Watermark    *string `json:"Watermark,omitempty"`

	// 视频宽度，单位为 px，默认值为 0。
	// 视频宽度的取值范围根据视频编码格式（Vcodec）的不同所有差异，视频宽度取值如下所示。
	// * 视频编码格式为 H.264 或 H.265 （Vcodec 取值为 h264 或 h265）时，取值范围为 [150,1920]；
	// * 视频编码格式为 H.266 （Vcodec 取值为 h266）时，不支持设置 Width 和 Height。
	// :::tip
	// * 当关闭视频分辨率自适应（As 取值为 0）时，转码分辨率将取 Width 和 Height 的值对转码视频进行拉伸；
	// * 当关闭视频分辨率自适应（As 取值为 0）时，Width 和 Height 任一取值为 0 时，转码视频将保持源流尺寸。 :::
	Width *int32 `json:"Width,omitempty"`
}

type UpdateTranscodePresetBodyTranscodeStruct struct {

	// Dictionary of
	ABTest       map[string]*Components1Wv3ClqSchemasUpdatetranscodepresetbodyPropertiesTranscodestructPropertiesAbtestAdditionalproperties `json:"ABTest,omitempty"`
	Codec        *string                                                                                                                    `json:"Codec,omitempty"`
	PresetName   *string                                                                                                                    `json:"PresetName,omitempty"`
	StopInterval *int32                                                                                                                     `json:"StopInterval,omitempty"`
	Suffix       *string                                                                                                                    `json:"Suffix,omitempty"`
	Type         *string                                                                                                                    `json:"Type,omitempty"`
}

type UpdateTranscodePresetRes struct {

	// REQUIRED
	ResponseMetadata UpdateTranscodePresetResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type UpdateTranscodePresetResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version   string                                         `json:"Version"`
	Error     *UpdateTranscodePresetResResponseMetadataError `json:"Error,omitempty"`
	RequestID *string                                        `json:"RequestID,omitempty"`
}

type UpdateTranscodePresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateUserAgentAccessRuleBody struct {

	// REQUIRED
	UaAccessRule UpdateUserAgentAccessRuleBodyUaAccessRule `json:"UaAccessRule"`

	// REQUIRED
	Vhost  string  `json:"Vhost"`
	Domain *string `json:"Domain,omitempty"`
}

type UpdateUserAgentAccessRuleBodyUaAccessRule struct {

	// REQUIRED
	AllowEmpty bool `json:"AllowEmpty"`

	// REQUIRED
	Enable bool `json:"Enable"`

	// REQUIRED
	Type string `json:"Type"`

	// REQUIRED
	UserAgent []string `json:"UserAgent"`
}

type UpdateUserAgentAccessRuleRes struct {

	// REQUIRED
	ResponseMetadata UpdateUserAgentAccessRuleResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateUserAgentAccessRuleResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestId为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string `json:"Version"`
}

type UpdateWatermarkPresetBody struct {

	// REQUIRED; 应用名称，您可以调用ListVhostWatermarkPreset [https://www.volcengine.com/docs/6469/1126889]接口，查看待更新水印配置的 App 取值。
	App string `json:"App"`

	// REQUIRED; 域名空间，您可以调用 ListVhostWatermarkPreset [https://www.volcengine.com/docs/6469/1126889] 接口，查看待更新水印配置的 Vhost 取值。
	Vhost string `json:"Vhost"`

	// 直播画面方向，支持 2 种取值。
	// * vertical：竖屏；
	// * horizontal：横屏。 :::tip 该参数属于历史版本参数，预计将于未来移除。建议使用预览背景高度（PreviewHeight）、预览背景宽度（PreviewWidth）参数代替。 :::
	Orientation *string `json:"Orientation,omitempty"`

	// 水印图片编码字符串，图片最大 2MB，最小 100Bytes，最大分辨率为 1080×1080。图片使用 data URI 协议，格式为：data:[<mediatype>];[base64],<data>。
	// * mediatype：图片类型，支持 png、jpg、jpeg 格式；
	// * data：base64 编码的图片字符串。
	// :::warning 如果水印图片不更新，请勿在更新配置时传入该参数，否则会造成水印无法显示。 :::
	Picture *string `json:"Picture,omitempty"`

	// 水印图片对应的 HTTP 地址。与水印图片编码字符串字段二选一传入。同时传入时，以水印图片编码字符串参数为准。 :::warning 如果水印图片不更新，请勿在更新配置时传入该参数，否则会造成水印无法显示。 :::
	PictureURL *string `json:"PictureUrl,omitempty"`

	// 水平偏移，表示水印左侧边与转码流画面左侧边之间的距离，使用相对比率，取值范围为 [0,1]。
	PosX *float32 `json:"PosX,omitempty"`

	// 垂直偏移，表示水印顶部边与转码流画面顶部边之间的距离，使用相对比率，取值范围为 [0,1]。
	PosY *float32 `json:"PosY,omitempty"`

	// 水印图片预览背景高度，单位为 px。
	PreviewHeight *float32 `json:"PreviewHeight,omitempty"`

	// 水印图片预览背景宽度，单位为 px。
	PreviewWidth *float32 `json:"PreviewWidth,omitempty"`

	// 水印相对高度，水印高度占直播转码流画面高度的比例，取值范围为 [0,1]，水印宽度会随高度等比缩放。与水印相对宽度字段冲突，请选择其中一个传参。
	RelativeHeight *float32 `json:"RelativeHeight,omitempty"`

	// 水印相对宽度，水印宽度占直播转码流画面宽度的比例，取值范围为 [0,1]，水印高度会随宽度等比缩放。与水印相对高度字段冲突，请选择其中一个传参。
	RelativeWidth *float32 `json:"RelativeWidth,omitempty"`

	// 直播地址流名。
	Stream *string `json:"Stream,omitempty"`
}

type UpdateWatermarkPresetRes struct {

	// REQUIRED
	ResponseMetadata UpdateWatermarkPresetResResponseMetadata `json:"ResponseMetadata"`
}

type UpdateWatermarkPresetResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                         `json:"Version"`
	Error   *UpdateWatermarkPresetResResponseMetadataError `json:"Error,omitempty"`
}

type UpdateWatermarkPresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}
type BindCert struct{}
type BindCertQuery struct{}
type BindEncryptDRM struct{}
type BindEncryptDRMQuery struct{}
type CreateCert struct{}
type CreateCertQuery struct{}
type CreateDomainV2 struct{}
type CreateDomainV2Query struct{}
type CreatePullToPushTask struct{}
type CreatePullToPushTaskQuery struct{}
type CreateRecordPresetV2 struct{}
type CreateRecordPresetV2Query struct{}
type CreateRelaySourceV4 struct{}
type CreateRelaySourceV4Query struct{}
type CreateSnapshotPreset struct{}
type CreateSnapshotPresetQuery struct{}
type CreateTimeShiftPresetV3 struct{}
type CreateTimeShiftPresetV3Query struct{}
type CreateTranscodePreset struct{}
type CreateTranscodePresetQuery struct{}
type CreateWatermarkPreset struct{}
type CreateWatermarkPresetQuery struct{}
type DeleteCMAFConfig struct{}
type DeleteCMAFConfigQuery struct{}
type DeleteCallback struct{}
type DeleteCallbackQuery struct{}
type DeleteCert struct{}
type DeleteCertQuery struct{}
type DeleteClusterRateLimit struct{}
type DeleteClusterRateLimitQuery struct{}
type DeleteDomain struct{}
type DeleteDomainQuery struct{}
type DeleteFormatAccessRule struct{}
type DeleteFormatAccessRuleQuery struct{}
type DeleteHLSConfig struct{}
type DeleteHLSConfigQuery struct{}
type DeleteHTTPHeaderConfig struct{}
type DeleteHTTPHeaderConfigQuery struct{}
type DeleteIPAccessRule struct{}
type DeleteIPAccessRuleQuery struct{}
type DeleteLatencyConfig struct{}
type DeleteLatencyConfigQuery struct{}
type DeletePullToPushTask struct{}
type DeletePullToPushTaskQuery struct{}
type DeleteRecordPreset struct{}
type DeleteRecordPresetQuery struct{}
type DeleteReferer struct{}
type DeleteRefererQuery struct{}
type DeleteRegionAccessRule struct{}
type DeleteRegionAccessRuleQuery struct{}
type DeleteRelaySourceRewrite struct{}
type DeleteRelaySourceRewriteQuery struct{}
type DeleteRelaySourceV3 struct{}
type DeleteRelaySourceV3Query struct{}
type DeleteRelaySourceV4 struct{}
type DeleteRelaySourceV4Query struct{}
type DeleteSnapshotPreset struct{}
type DeleteSnapshotPresetQuery struct{}
type DeleteTimeShiftPresetV3 struct{}
type DeleteTimeShiftPresetV3Query struct{}
type DeleteTranscodePreset struct{}
type DeleteTranscodePresetQuery struct{}
type DeleteUserAgentAccessRule struct{}
type DeleteUserAgentAccessRuleQuery struct{}
type DeleteWatermarkPreset struct{}
type DeleteWatermarkPresetQuery struct{}
type DescribeAuth struct{}
type DescribeAuthQuery struct{}
type DescribeCDNSnapshotHistory struct{}
type DescribeCDNSnapshotHistoryQuery struct{}
type DescribeCMAFConfig struct{}
type DescribeCMAFConfigQuery struct{}
type DescribeCallback struct{}
type DescribeCallbackQuery struct{}
type DescribeCertDRM struct{}
type DescribeCertDRMBody struct{}
type DescribeCertDetailSecretV2 struct{}
type DescribeCertDetailSecretV2Query struct{}
type DescribeClosedStreamInfoByPage struct{}
type DescribeClosedStreamInfoByPageBody struct{}
type DescribeClusterRateLimit struct{}
type DescribeClusterRateLimitQuery struct{}
type DescribeDomain struct{}
type DescribeDomainQuery struct{}
type DescribeEncryptDRM struct{}
type DescribeEncryptDRMBody struct{}
type DescribeEncryptDRMQuery struct{}
type DescribeForbiddenStreamInfoByPage struct{}
type DescribeForbiddenStreamInfoByPageBody struct{}
type DescribeFormatAccessRule struct{}
type DescribeFormatAccessRuleQuery struct{}
type DescribeHLSConfig struct{}
type DescribeHLSConfigQuery struct{}
type DescribeHTTPHeaderConfig struct{}
type DescribeHTTPHeaderConfigQuery struct{}
type DescribeIPAccessRule struct{}
type DescribeIPAccessRuleQuery struct{}
type DescribeIPInfo struct{}
type DescribeIPInfoQuery struct{}
type DescribeLatencyConfig struct{}
type DescribeLatencyConfigQuery struct{}
type DescribeLicenseDRM struct{}
type DescribeLicenseDRMBody struct{}
type DescribeLiveASRDurationData struct{}
type DescribeLiveASRDurationDataQuery struct{}
type DescribeLiveBandwidthData struct{}
type DescribeLiveBandwidthDataQuery struct{}
type DescribeLiveISPData struct{}
type DescribeLiveISPDataBody struct{}
type DescribeLiveISPDataQuery struct{}
type DescribeLiveLogData struct{}
type DescribeLiveLogDataQuery struct{}
type DescribeLiveMetricBandwidthData struct{}
type DescribeLiveMetricBandwidthDataQuery struct{}
type DescribeLiveMetricTrafficData struct{}
type DescribeLiveMetricTrafficDataQuery struct{}
type DescribeLiveP95PeakBandwidthData struct{}
type DescribeLiveP95PeakBandwidthDataQuery struct{}
type DescribeLivePlayStatusCodeData struct{}
type DescribeLivePlayStatusCodeDataQuery struct{}
type DescribeLivePullToPushBandwidthData struct{}
type DescribeLivePullToPushBandwidthDataQuery struct{}
type DescribeLivePullToPushData struct{}
type DescribeLivePullToPushDataQuery struct{}
type DescribeLivePushStreamInfoData struct{}
type DescribeLivePushStreamInfoDataQuery struct{}
type DescribeLivePushStreamMetrics struct{}
type DescribeLivePushStreamMetricsQuery struct{}
type DescribeLiveRecordData struct{}
type DescribeLiveRecordDataQuery struct{}
type DescribeLiveRegionData struct{}
type DescribeLiveRegionDataBody struct{}
type DescribeLiveRegionDataQuery struct{}
type DescribeLiveSnapshotData struct{}
type DescribeLiveSnapshotDataQuery struct{}
type DescribeLiveSourceBandwidthData struct{}
type DescribeLiveSourceBandwidthDataQuery struct{}
type DescribeLiveSourceStreamMetrics struct{}
type DescribeLiveSourceStreamMetricsQuery struct{}
type DescribeLiveSourceTrafficData struct{}
type DescribeLiveSourceTrafficDataQuery struct{}
type DescribeLiveStorageSpaceData struct{}
type DescribeLiveStorageSpaceDataQuery struct{}
type DescribeLiveStreamInfoByPage struct{}
type DescribeLiveStreamInfoByPageBody struct{}
type DescribeLiveStreamSessionData struct{}
type DescribeLiveStreamSessionDataQuery struct{}
type DescribeLiveStreamState struct{}
type DescribeLiveStreamStateBody struct{}
type DescribeLiveTrafficData struct{}
type DescribeLiveTrafficDataQuery struct{}
type DescribeLiveTranscodeData struct{}
type DescribeLiveTranscodeDataQuery struct{}
type DescribeRecordTaskFileHistory struct{}
type DescribeRecordTaskFileHistoryQuery struct{}
type DescribeReferer struct{}
type DescribeRefererQuery struct{}
type DescribeRegionAccessRule struct{}
type DescribeRegionAccessRuleQuery struct{}
type DescribeRelaySourceRewrite struct{}
type DescribeRelaySourceRewriteQuery struct{}
type DescribeRelaySourceV3 struct{}
type DescribeRelaySourceV3Query struct{}
type DescribeUserAgentAccessRule struct{}
type DescribeUserAgentAccessRuleQuery struct{}
type DisableDomain struct{}
type DisableDomainQuery struct{}
type EnableDomain struct{}
type EnableDomainQuery struct{}
type EnableHTTPHeaderConfig struct{}
type EnableHTTPHeaderConfigQuery struct{}
type ForbidStream struct{}
type ForbidStreamQuery struct{}
type GeneratePlayURL struct{}
type GeneratePlayURLQuery struct{}
type GeneratePushURL struct{}
type GeneratePushURLQuery struct{}
type KillStream struct{}
type KillStreamQuery struct{}
type ListBindEncryptDRM struct{}
type ListBindEncryptDRMQuery struct{}
type ListCertV2 struct{}
type ListCertV2Query struct{}
type ListCommonTransPresetDetail struct{}
type ListCommonTransPresetDetailQuery struct{}
type ListDomainDetail struct{}
type ListDomainDetailQuery struct{}
type ListPullToPushTask struct{}
type ListPullToPushTaskBody struct{}
type ListRelaySourceV4 struct{}
type ListRelaySourceV4Query struct{}
type ListTimeShiftPresetV2 struct{}
type ListTimeShiftPresetV2Query struct{}
type ListVhostRecordPresetV2 struct{}
type ListVhostRecordPresetV2Query struct{}
type ListVhostSnapshotPreset struct{}
type ListVhostSnapshotPresetQuery struct{}
type ListVhostTransCodePreset struct{}
type ListVhostTransCodePresetQuery struct{}
type ListVhostWatermarkPreset struct{}
type ListVhostWatermarkPresetQuery struct{}
type ListWatermarkPreset struct{}
type ListWatermarkPresetQuery struct{}
type RestartPullToPushTask struct{}
type RestartPullToPushTaskQuery struct{}
type ResumeStream struct{}
type ResumeStreamQuery struct{}
type StopPullToPushTask struct{}
type StopPullToPushTaskQuery struct{}
type UnBindEncryptDRM struct{}
type UnBindEncryptDRMQuery struct{}
type UnbindCert struct{}
type UnbindCertQuery struct{}
type UpdateAuthKey struct{}
type UpdateAuthKeyQuery struct{}
type UpdateCMAFConfig struct{}
type UpdateCMAFConfigQuery struct{}
type UpdateCallback struct{}
type UpdateCallbackQuery struct{}
type UpdateClusterRateLimit struct{}
type UpdateClusterRateLimitQuery struct{}
type UpdateDomainVhost struct{}
type UpdateDomainVhostQuery struct{}
type UpdateEncryptDRM struct{}
type UpdateEncryptDRMQuery struct{}
type UpdateFormatAccessRule struct{}
type UpdateFormatAccessRuleQuery struct{}
type UpdateHLSConfig struct{}
type UpdateHLSConfigQuery struct{}
type UpdateHTTPHeaderConfig struct{}
type UpdateHTTPHeaderConfigQuery struct{}
type UpdateIPAccessRule struct{}
type UpdateIPAccessRuleQuery struct{}
type UpdateLatencyConfig struct{}
type UpdateLatencyConfigQuery struct{}
type UpdatePullToPushTask struct{}
type UpdatePullToPushTaskQuery struct{}
type UpdateRecordPresetV2 struct{}
type UpdateRecordPresetV2Query struct{}
type UpdateReferer struct{}
type UpdateRefererQuery struct{}
type UpdateRegionAccessRule struct{}
type UpdateRegionAccessRuleQuery struct{}
type UpdateRelaySourceRewrite struct{}
type UpdateRelaySourceRewriteQuery struct{}
type UpdateRelaySourceV3 struct{}
type UpdateRelaySourceV3Query struct{}
type UpdateRelaySourceV4 struct{}
type UpdateRelaySourceV4Query struct{}
type UpdateSnapshotPreset struct{}
type UpdateSnapshotPresetQuery struct{}
type UpdateTimeShiftPresetV3 struct{}
type UpdateTimeShiftPresetV3Query struct{}
type UpdateTranscodePreset struct{}
type UpdateTranscodePresetQuery struct{}
type UpdateUserAgentAccessRule struct{}
type UpdateUserAgentAccessRuleQuery struct{}
type UpdateWatermarkPreset struct{}
type UpdateWatermarkPresetQuery struct{}
type BindCertReq struct {
	*BindCertQuery
	*BindCertBody
}
type BindEncryptDRMReq struct {
	*BindEncryptDRMQuery
	*BindEncryptDRMBody
}
type CreateCertReq struct {
	*CreateCertQuery
	*CreateCertBody
}
type CreateDomainV2Req struct {
	*CreateDomainV2Query
	*CreateDomainV2Body
}
type CreatePullToPushTaskReq struct {
	*CreatePullToPushTaskQuery
	*CreatePullToPushTaskBody
}
type CreateRecordPresetV2Req struct {
	*CreateRecordPresetV2Query
	*CreateRecordPresetV2Body
}
type CreateRelaySourceV4Req struct {
	*CreateRelaySourceV4Query
	*CreateRelaySourceV4Body
}
type CreateSnapshotPresetReq struct {
	*CreateSnapshotPresetQuery
	*CreateSnapshotPresetBody
}
type CreateTimeShiftPresetV3Req struct {
	*CreateTimeShiftPresetV3Query
	*CreateTimeShiftPresetV3Body
}
type CreateTranscodePresetReq struct {
	*CreateTranscodePresetQuery
	*CreateTranscodePresetBody
}
type CreateWatermarkPresetReq struct {
	*CreateWatermarkPresetQuery
	*CreateWatermarkPresetBody
}
type DeleteCMAFConfigReq struct {
	*DeleteCMAFConfigQuery
	*DeleteCMAFConfigBody
}
type DeleteCallbackReq struct {
	*DeleteCallbackQuery
	*DeleteCallbackBody
}
type DeleteCertReq struct {
	*DeleteCertQuery
	*DeleteCertBody
}
type DeleteClusterRateLimitReq struct {
	*DeleteClusterRateLimitQuery
	*DeleteClusterRateLimitBody
}
type DeleteDomainReq struct {
	*DeleteDomainQuery
	*DeleteDomainBody
}
type DeleteFormatAccessRuleReq struct {
	*DeleteFormatAccessRuleQuery
	*DeleteFormatAccessRuleBody
}
type DeleteHLSConfigReq struct {
	*DeleteHLSConfigQuery
	*DeleteHLSConfigBody
}
type DeleteHTTPHeaderConfigReq struct {
	*DeleteHTTPHeaderConfigQuery
	*DeleteHTTPHeaderConfigBody
}
type DeleteIPAccessRuleReq struct {
	*DeleteIPAccessRuleQuery
	*DeleteIPAccessRuleBody
}
type DeleteLatencyConfigReq struct {
	*DeleteLatencyConfigQuery
	*DeleteLatencyConfigBody
}
type DeletePullToPushTaskReq struct {
	*DeletePullToPushTaskQuery
	*DeletePullToPushTaskBody
}
type DeleteRecordPresetReq struct {
	*DeleteRecordPresetQuery
	*DeleteRecordPresetBody
}
type DeleteRefererReq struct {
	*DeleteRefererQuery
	*DeleteRefererBody
}
type DeleteRegionAccessRuleReq struct {
	*DeleteRegionAccessRuleQuery
	*DeleteRegionAccessRuleBody
}
type DeleteRelaySourceRewriteReq struct {
	*DeleteRelaySourceRewriteQuery
	*DeleteRelaySourceRewriteBody
}
type DeleteRelaySourceV3Req struct {
	*DeleteRelaySourceV3Query
	*DeleteRelaySourceV3Body
}
type DeleteRelaySourceV4Req struct {
	*DeleteRelaySourceV4Query
	*DeleteRelaySourceV4Body
}
type DeleteSnapshotPresetReq struct {
	*DeleteSnapshotPresetQuery
	*DeleteSnapshotPresetBody
}
type DeleteTimeShiftPresetV3Req struct {
	*DeleteTimeShiftPresetV3Query
	*DeleteTimeShiftPresetV3Body
}
type DeleteTranscodePresetReq struct {
	*DeleteTranscodePresetQuery
	*DeleteTranscodePresetBody
}
type DeleteUserAgentAccessRuleReq struct {
	*DeleteUserAgentAccessRuleQuery
	*DeleteUserAgentAccessRuleBody
}
type DeleteWatermarkPresetReq struct {
	*DeleteWatermarkPresetQuery
	*DeleteWatermarkPresetBody
}
type DescribeAuthReq struct {
	*DescribeAuthQuery
	*DescribeAuthBody
}
type DescribeCDNSnapshotHistoryReq struct {
	*DescribeCDNSnapshotHistoryQuery
	*DescribeCDNSnapshotHistoryBody
}
type DescribeCMAFConfigReq struct {
	*DescribeCMAFConfigQuery
	*DescribeCMAFConfigBody
}
type DescribeCallbackReq struct {
	*DescribeCallbackQuery
	*DescribeCallbackBody
}
type DescribeCertDRMReq struct {
	*DescribeCertDRMQuery
	*DescribeCertDRMBody
}
type DescribeCertDetailSecretV2Req struct {
	*DescribeCertDetailSecretV2Query
	*DescribeCertDetailSecretV2Body
}
type DescribeClosedStreamInfoByPageReq struct {
	*DescribeClosedStreamInfoByPageQuery
	*DescribeClosedStreamInfoByPageBody
}
type DescribeClusterRateLimitReq struct {
	*DescribeClusterRateLimitQuery
	*DescribeClusterRateLimitBody
}
type DescribeDomainReq struct {
	*DescribeDomainQuery
	*DescribeDomainBody
}
type DescribeEncryptDRMReq struct {
	*DescribeEncryptDRMQuery
	*DescribeEncryptDRMBody
}
type DescribeForbiddenStreamInfoByPageReq struct {
	*DescribeForbiddenStreamInfoByPageQuery
	*DescribeForbiddenStreamInfoByPageBody
}
type DescribeFormatAccessRuleReq struct {
	*DescribeFormatAccessRuleQuery
	*DescribeFormatAccessRuleBody
}
type DescribeHLSConfigReq struct {
	*DescribeHLSConfigQuery
	*DescribeHLSConfigBody
}
type DescribeHTTPHeaderConfigReq struct {
	*DescribeHTTPHeaderConfigQuery
	*DescribeHTTPHeaderConfigBody
}
type DescribeIPAccessRuleReq struct {
	*DescribeIPAccessRuleQuery
	*DescribeIPAccessRuleBody
}
type DescribeIPInfoReq struct {
	*DescribeIPInfoQuery
	*DescribeIPInfoBody
}
type DescribeLatencyConfigReq struct {
	*DescribeLatencyConfigQuery
	*DescribeLatencyConfigBody
}
type DescribeLicenseDRMReq struct {
	*DescribeLicenseDRMQuery
	*DescribeLicenseDRMBody
}
type DescribeLiveASRDurationDataReq struct {
	*DescribeLiveASRDurationDataQuery
	*DescribeLiveASRDurationDataBody
}
type DescribeLiveBandwidthDataReq struct {
	*DescribeLiveBandwidthDataQuery
	*DescribeLiveBandwidthDataBody
}
type DescribeLiveISPDataReq struct {
	*DescribeLiveISPDataQuery
	*DescribeLiveISPDataBody
}
type DescribeLiveLogDataReq struct {
	*DescribeLiveLogDataQuery
	*DescribeLiveLogDataBody
}
type DescribeLiveMetricBandwidthDataReq struct {
	*DescribeLiveMetricBandwidthDataQuery
	*DescribeLiveMetricBandwidthDataBody
}
type DescribeLiveMetricTrafficDataReq struct {
	*DescribeLiveMetricTrafficDataQuery
	*DescribeLiveMetricTrafficDataBody
}
type DescribeLiveP95PeakBandwidthDataReq struct {
	*DescribeLiveP95PeakBandwidthDataQuery
	*DescribeLiveP95PeakBandwidthDataBody
}
type DescribeLivePlayStatusCodeDataReq struct {
	*DescribeLivePlayStatusCodeDataQuery
	*DescribeLivePlayStatusCodeDataBody
}
type DescribeLivePullToPushBandwidthDataReq struct {
	*DescribeLivePullToPushBandwidthDataQuery
	*DescribeLivePullToPushBandwidthDataBody
}
type DescribeLivePullToPushDataReq struct {
	*DescribeLivePullToPushDataQuery
	*DescribeLivePullToPushDataBody
}
type DescribeLivePushStreamInfoDataReq struct {
	*DescribeLivePushStreamInfoDataQuery
	*DescribeLivePushStreamInfoDataBody
}
type DescribeLivePushStreamMetricsReq struct {
	*DescribeLivePushStreamMetricsQuery
	*DescribeLivePushStreamMetricsBody
}
type DescribeLiveRecordDataReq struct {
	*DescribeLiveRecordDataQuery
	*DescribeLiveRecordDataBody
}
type DescribeLiveRegionDataReq struct {
	*DescribeLiveRegionDataQuery
	*DescribeLiveRegionDataBody
}
type DescribeLiveSnapshotDataReq struct {
	*DescribeLiveSnapshotDataQuery
	*DescribeLiveSnapshotDataBody
}
type DescribeLiveSourceBandwidthDataReq struct {
	*DescribeLiveSourceBandwidthDataQuery
	*DescribeLiveSourceBandwidthDataBody
}
type DescribeLiveSourceStreamMetricsReq struct {
	*DescribeLiveSourceStreamMetricsQuery
	*DescribeLiveSourceStreamMetricsBody
}
type DescribeLiveSourceTrafficDataReq struct {
	*DescribeLiveSourceTrafficDataQuery
	*DescribeLiveSourceTrafficDataBody
}
type DescribeLiveStorageSpaceDataReq struct {
	*DescribeLiveStorageSpaceDataQuery
	*DescribeLiveStorageSpaceDataBody
}
type DescribeLiveStreamInfoByPageReq struct {
	*DescribeLiveStreamInfoByPageQuery
	*DescribeLiveStreamInfoByPageBody
}
type DescribeLiveStreamSessionDataReq struct {
	*DescribeLiveStreamSessionDataQuery
	*DescribeLiveStreamSessionDataBody
}
type DescribeLiveStreamStateReq struct {
	*DescribeLiveStreamStateQuery
	*DescribeLiveStreamStateBody
}
type DescribeLiveTrafficDataReq struct {
	*DescribeLiveTrafficDataQuery
	*DescribeLiveTrafficDataBody
}
type DescribeLiveTranscodeDataReq struct {
	*DescribeLiveTranscodeDataQuery
	*DescribeLiveTranscodeDataBody
}
type DescribeRecordTaskFileHistoryReq struct {
	*DescribeRecordTaskFileHistoryQuery
	*DescribeRecordTaskFileHistoryBody
}
type DescribeRefererReq struct {
	*DescribeRefererQuery
	*DescribeRefererBody
}
type DescribeRegionAccessRuleReq struct {
	*DescribeRegionAccessRuleQuery
	*DescribeRegionAccessRuleBody
}
type DescribeRelaySourceRewriteReq struct {
	*DescribeRelaySourceRewriteQuery
	*DescribeRelaySourceRewriteBody
}
type DescribeRelaySourceV3Req struct {
	*DescribeRelaySourceV3Query
	*DescribeRelaySourceV3Body
}
type DescribeUserAgentAccessRuleReq struct {
	*DescribeUserAgentAccessRuleQuery
	*DescribeUserAgentAccessRuleBody
}
type DisableDomainReq struct {
	*DisableDomainQuery
	*DisableDomainBody
}
type EnableDomainReq struct {
	*EnableDomainQuery
	*EnableDomainBody
}
type EnableHTTPHeaderConfigReq struct {
	*EnableHTTPHeaderConfigQuery
	*EnableHTTPHeaderConfigBody
}
type ForbidStreamReq struct {
	*ForbidStreamQuery
	*ForbidStreamBody
}
type GeneratePlayURLReq struct {
	*GeneratePlayURLQuery
	*GeneratePlayURLBody
}
type GeneratePushURLReq struct {
	*GeneratePushURLQuery
	*GeneratePushURLBody
}
type KillStreamReq struct {
	*KillStreamQuery
	*KillStreamBody
}
type ListBindEncryptDRMReq struct {
	*ListBindEncryptDRMQuery
	*ListBindEncryptDRMBody
}
type ListCertV2Req struct {
	*ListCertV2Query
	*ListCertV2Body
}
type ListCommonTransPresetDetailReq struct {
	*ListCommonTransPresetDetailQuery
	*ListCommonTransPresetDetailBody
}
type ListDomainDetailReq struct {
	*ListDomainDetailQuery
	*ListDomainDetailBody
}
type ListPullToPushTaskReq struct {
	*ListPullToPushTaskQuery
	*ListPullToPushTaskBody
}
type ListRelaySourceV4Req struct {
	*ListRelaySourceV4Query
	*ListRelaySourceV4Body
}
type ListTimeShiftPresetV2Req struct {
	*ListTimeShiftPresetV2Query
	*ListTimeShiftPresetV2Body
}
type ListVhostRecordPresetV2Req struct {
	*ListVhostRecordPresetV2Query
	*ListVhostRecordPresetV2Body
}
type ListVhostSnapshotPresetReq struct {
	*ListVhostSnapshotPresetQuery
	*ListVhostSnapshotPresetBody
}
type ListVhostTransCodePresetReq struct {
	*ListVhostTransCodePresetQuery
	*ListVhostTransCodePresetBody
}
type ListVhostWatermarkPresetReq struct {
	*ListVhostWatermarkPresetQuery
	*ListVhostWatermarkPresetBody
}
type ListWatermarkPresetReq struct {
	*ListWatermarkPresetQuery
	*ListWatermarkPresetBody
}
type RestartPullToPushTaskReq struct {
	*RestartPullToPushTaskQuery
	*RestartPullToPushTaskBody
}
type ResumeStreamReq struct {
	*ResumeStreamQuery
	*ResumeStreamBody
}
type StopPullToPushTaskReq struct {
	*StopPullToPushTaskQuery
	*StopPullToPushTaskBody
}
type UnBindEncryptDRMReq struct {
	*UnBindEncryptDRMQuery
	*UnBindEncryptDRMBody
}
type UnbindCertReq struct {
	*UnbindCertQuery
	*UnbindCertBody
}
type UpdateAuthKeyReq struct {
	*UpdateAuthKeyQuery
	*UpdateAuthKeyBody
}
type UpdateCMAFConfigReq struct {
	*UpdateCMAFConfigQuery
	*UpdateCMAFConfigBody
}
type UpdateCallbackReq struct {
	*UpdateCallbackQuery
	*UpdateCallbackBody
}
type UpdateClusterRateLimitReq struct {
	*UpdateClusterRateLimitQuery
	*UpdateClusterRateLimitBody
}
type UpdateDomainVhostReq struct {
	*UpdateDomainVhostQuery
	*UpdateDomainVhostBody
}
type UpdateEncryptDRMReq struct {
	*UpdateEncryptDRMQuery
	*UpdateEncryptDRMBody
}
type UpdateFormatAccessRuleReq struct {
	*UpdateFormatAccessRuleQuery
	*UpdateFormatAccessRuleBody
}
type UpdateHLSConfigReq struct {
	*UpdateHLSConfigQuery
	*UpdateHLSConfigBody
}
type UpdateHTTPHeaderConfigReq struct {
	*UpdateHTTPHeaderConfigQuery
	*UpdateHTTPHeaderConfigBody
}
type UpdateIPAccessRuleReq struct {
	*UpdateIPAccessRuleQuery
	*UpdateIPAccessRuleBody
}
type UpdateLatencyConfigReq struct {
	*UpdateLatencyConfigQuery
	*UpdateLatencyConfigBody
}
type UpdatePullToPushTaskReq struct {
	*UpdatePullToPushTaskQuery
	*UpdatePullToPushTaskBody
}
type UpdateRecordPresetV2Req struct {
	*UpdateRecordPresetV2Query
	*UpdateRecordPresetV2Body
}
type UpdateRefererReq struct {
	*UpdateRefererQuery
	*UpdateRefererBody
}
type UpdateRegionAccessRuleReq struct {
	*UpdateRegionAccessRuleQuery
	*UpdateRegionAccessRuleBody
}
type UpdateRelaySourceRewriteReq struct {
	*UpdateRelaySourceRewriteQuery
	*UpdateRelaySourceRewriteBody
}
type UpdateRelaySourceV3Req struct {
	*UpdateRelaySourceV3Query
	*UpdateRelaySourceV3Body
}
type UpdateRelaySourceV4Req struct {
	*UpdateRelaySourceV4Query
	*UpdateRelaySourceV4Body
}
type UpdateSnapshotPresetReq struct {
	*UpdateSnapshotPresetQuery
	*UpdateSnapshotPresetBody
}
type UpdateTimeShiftPresetV3Req struct {
	*UpdateTimeShiftPresetV3Query
	*UpdateTimeShiftPresetV3Body
}
type UpdateTranscodePresetReq struct {
	*UpdateTranscodePresetQuery
	*UpdateTranscodePresetBody
}
type UpdateUserAgentAccessRuleReq struct {
	*UpdateUserAgentAccessRuleQuery
	*UpdateUserAgentAccessRuleBody
}
type UpdateWatermarkPresetReq struct {
	*UpdateWatermarkPresetQuery
	*UpdateWatermarkPresetBody
}
