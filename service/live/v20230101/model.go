package live_v20230101

type AddCommonTransPresetBody struct {
	App        *string                                   `json:"App,omitempty"`
	PresetList []*AddCommonTransPresetBodyPresetListItem `json:"PresetList,omitempty"`
	Vhost      *string                                   `json:"Vhost,omitempty"`
}

type AddCommonTransPresetBodyPresetListItem struct {

	// Dictionary of
	ABTest       map[string]*Components1UawxzeSchemasAddcommontranspresetbodyPropertiesPresetlistItemsPropertiesAbtestAdditionalproperties `json:"ABTest,omitempty"`
	Codec        *string                                                                                                                   `json:"Codec,omitempty"`
	PresetName   *string                                                                                                                   `json:"PresetName,omitempty"`
	StopInterval *int32                                                                                                                    `json:"StopInterval,omitempty"`
	Suffix       *string                                                                                                                   `json:"Suffix,omitempty"`
	Type         *string                                                                                                                   `json:"Type,omitempty"`
}

type AddCommonTransPresetRes struct {

	// REQUIRED
	ResponseMetadata AddCommonTransPresetResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type AddCommonTransPresetResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version   string                                        `json:"Version"`
	Error     *AddCommonTransPresetResResponseMetadataError `json:"Error,omitempty"`
	RequestID *string                                       `json:"RequestID,omitempty"`
}

type AddCommonTransPresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type AssociatePresetBody struct {

	// REQUIRED; 模板名称
	PresetName string `json:"PresetName"`

	// REQUIRED; 模板类型， recor:录制 snapshot:密集抽帧 transcode:转码 avslice:音频切片 cdnsnapshot：截图 avextractor timeshift：时移 auditsnapshot：审核截图
	// data_migration watermark：水印
	PresetType string `json:"PresetType"`

	// REQUIRED; 域名空间名称
	Vhost string `json:"Vhost"`

	// 应用名称
	App           *string `json:"App,omitempty"`
	PresetNameOld *string `json:"PresetNameOld,omitempty"`

	// 录制配置
	RecordPresetParam *AssociatePresetBodyRecordPresetParam `json:"RecordPresetParam,omitempty"`

	// 录制类型：push, pull
	RecordType *string `json:"RecordType,omitempty"`

	// 流名
	Stream *string `json:"Stream,omitempty"`

	// 时移配置
	TimeShiftStruct *AssociatePresetBodyTimeShiftStruct `json:"TimeShiftStruct,omitempty"`
	TranscodeStruct *AssociatePresetBodyTranscodeStruct `json:"TranscodeStruct,omitempty"`
}

// AssociatePresetBodyRecordPresetParam - 录制配置
type AssociatePresetBodyRecordPresetParam struct {
	ContinueDuration *int32 `json:"ContinueDuration,omitempty"`

	// 源流录制，1表示录制
	OriginRecord           *int32  `json:"OriginRecord,omitempty"`
	OriginRegexp           *string `json:"OriginRegexp,omitempty"`
	RealtimeRecordDuration *int32  `json:"RealtimeRecordDuration,omitempty"`
	RelayEnable            *bool   `json:"RelayEnable,omitempty"`

	// 转码流录制，1表示录制,2录制全部
	TranscodeRecord *int32 `json:"TranscodeRecord,omitempty"`

	// 转码流录制后缀
	TranscodeSuffixList []*string `json:"TranscodeSuffixList,omitempty"`
}

// AssociatePresetBodyTimeShiftStruct - 时移配置
type AssociatePresetBodyTimeShiftStruct struct {

	// 是否需要转码流时移
	NeedTranscode *int32 `json:"NeedTranscode,omitempty"`

	// 时移的类型
	TimeShiftType *int32 `json:"TimeShiftType,omitempty"`
}

type AssociatePresetBodyTranscodeStruct struct {

	// Anything
	ABTest                interface{} `json:"ABTest,omitempty"`
	Codec                 *string     `json:"Codec,omitempty"`
	CreateTime            *string     `json:"CreateTime,omitempty"`
	PresetName            *string     `json:"PresetName,omitempty"`
	RelayDisableTranscode *bool       `json:"RelayDisableTranscode,omitempty"`
	StopInterval          *int32      `json:"StopInterval,omitempty"`
	Suffix                *string     `json:"Suffix,omitempty"`
	Type                  *string     `json:"Type,omitempty"`
	UpdateTime            *string     `json:"UpdateTime,omitempty"`
}

type AssociatePresetRes struct {

	// REQUIRED
	ResponseMetadata AssociatePresetResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type AssociatePresetResResponseMetadata struct {

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

type AssociateRefConfigBody struct {

	// REQUIRED; 引用名，若已存在相同类型的引用会直接覆盖
	Name string `json:"Name"`

	// 应用名称。 :::tip App 与 Domain 二选一填 :::
	App *string `json:"App,omitempty"`

	// 域名。 :::tip App 与 Domain 二选一填 :::
	Domain *string `json:"Domain,omitempty"`

	// 域名空间名称。
	Vhost *string `json:"Vhost,omitempty"`
}

type AssociateRefConfigRes struct {

	// REQUIRED
	ResponseMetadata AssociateRefConfigResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type AssociateRefConfigResResponseMetadata struct {

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

type BindCertBody struct {

	// REQUIRED; 需要绑定的 HTTPS 证书的证书链 ID，可以通过查询证书列表 [https://www.volcengine.com/docs/6469/1126822]接口获取。
	ChainID string `json:"ChainID"`

	// REQUIRED; 填写需要配置 HTTPS 证书的域名。 您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理 [https://console-stable.volcanicengine.com/live/main/domain/list]页面，查看需要绑定证书的域名。
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

	// REQUIRED; app
	App string `json:"App"`

	// REQUIRED; 是否开启drm配置
	Enable bool `json:"Enable"`

	// REQUIRED; 需要加密的转码模版后缀（源流默认加密）
	EncryptTranscodeList []string `json:"EncryptTranscodeList"`

	// REQUIRED; 域名空间
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

type CheckCustomLogConfigBody struct {

	// REQUIRED; 用户账号id
	AccountID string `json:"AccountId"`

	// REQUIRED; 用户账号名称
	AccountName string `json:"AccountName"`

	// REQUIRED; 日志获取接口名称
	ActionName string `json:"ActionName"`

	// REQUIRED; bmq集群名
	BmqCluster string `json:"BmqCluster"`

	// REQUIRED; tce集群名
	Cluster string `json:"Cluster"`

	// REQUIRED; 创建人
	Creator string `json:"Creator"`

	// REQUIRED; 下载时填的Type参数
	DownloadType string `json:"DownloadType"`

	// REQUIRED; 文件名字段名称
	FileNameFields CheckCustomLogConfigBodyFileNameFields `json:"FileNameFields"`

	// REQUIRED; 文件名pattern
	FileNamePattern string `json:"FileNamePattern"`

	// REQUIRED; 日志字段名称
	LogFields CheckCustomLogConfigBodyLogFields `json:"LogFields"`

	// REQUIRED; 日志pattern
	LogPattern string `json:"LogPattern"`

	// REQUIRED; 日志类型，如果是多个用逗号连接，全选可填*
	LogType string `json:"LogType"`

	// REQUIRED; 写入的topic
	Topic string `json:"Topic"`

	// REQUIRED; 特殊清洗状态，false:表示数仓单独任务进行清洗，true为通用清洗任务
	WashStatus bool `json:"WashStatus"`

	// 延迟时间，默认300s
	DelayTime *int32 `json:"DelayTime,omitempty"`

	// 填1或0，是否补空文件，默认为0
	EmptyFile *int32 `json:"EmptyFile,omitempty"`

	// 排除的账号ID
	ExcludedAccountIDs *string `json:"ExcludedAccountIds,omitempty"`

	// 如果没有Id，表示创建，带了Id表示更新
	ID *int32 `json:"Id,omitempty"`

	// 是否每个域名一个文件，默认为false
	SplitDomain *bool `json:"SplitDomain,omitempty"`

	// 默认false（前端默认填ture），同一个时间范围是否允许按照大小切割文件
	SplitFile *bool `json:"SplitFile,omitempty"`

	// 切割文件的行数，默��120w
	SplitLine *int32 `json:"SplitLine,omitempty"`

	// 切割文件的时间，单位秒，默认3600
	SplitTime *int32 `json:"SplitTime,omitempty"`

	// 默认0，状态，1：启动，0：禁止
	Status *int32 `json:"Status,omitempty"`
}

// CheckCustomLogConfigBodyFileNameFields - 文件名字段名称
type CheckCustomLogConfigBodyFileNameFields struct {

	// REQUIRED; 字段名称
	Key string `json:"Key"`

	// REQUIRED; 字段类型，不能为空
	Type string `json:"Type"`

	// 备注信息，没有可以为空
	FmtValue *string `json:"FmtValue,omitempty"`

	// 字段对应中文名
	KeyCn *string `json:"KeyCn,omitempty"`

	// 敏感词替换字符串，比如：ab,cd 表示用cd替换ab，如果有多组替换用分号连接
	Transform *int32 `json:"Transform,omitempty"`
}

// CheckCustomLogConfigBodyLogFields - 日志字段名称
type CheckCustomLogConfigBodyLogFields struct {

	// REQUIRED; 字段名称
	Key string `json:"Key"`

	// REQUIRED; 字段类型，不能为空
	Type string `json:"Type"`

	// 备注信息，没有可以为空
	FmtValue *string `json:"FmtValue,omitempty"`

	// 字段对应中文名
	KeyCn *string `json:"KeyCn,omitempty"`

	// 敏感词替换字符串，比如：ab,cd 表示用cd替换ab，如果有多组替换用分号连接
	Transform *int32 `json:"Transform,omitempty"`
}

type CheckCustomLogConfigRes struct {

	// REQUIRED
	ResponseMetadata CheckCustomLogConfigResResponseMetadata `json:"ResponseMetadata"`
	Result           *CheckCustomLogConfigResResult          `json:"Result,omitempty"`
}

type CheckCustomLogConfigResResponseMetadata struct {

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

type CheckCustomLogConfigResResult struct {

	// REQUIRED
	Message string `json:"Message"`

	// REQUIRED
	Status int32 `json:"Status"`
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

type Components17C6BtpSchemasDescribepresetassociationresPropertiesResultPropertiesListItemsPropertiesWatermarkpresetlistv2Items struct {
	CreateTime string `json:"CreateTime"`

	PresetName string `json:"PresetName"`

	UpdateTime string `json:"UpdateTime"`
}

type Components1Bmm523SchemasListvhostdetailbyadminresPropertiesResultPropertiesVhostlistItemsPropertiesDomainlistItems struct {
	CNAME      *string `json:"CNAME,omitempty"`
	CertDomain *string `json:"CertDomain,omitempty"`
	CertName   *string `json:"CertName,omitempty"`
	ChainID    *string `json:"ChainID,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty"`
	Domain     *string `json:"Domain,omitempty"`
	Priority   *int32  `json:"Priority,omitempty"`
	PushDomain *string `json:"PushDomain,omitempty"`
	Region     *string `json:"Region,omitempty"`
	Status     *int32  `json:"Status,omitempty"`
	Type       *string `json:"Type,omitempty"`
	Vhost      *string `json:"Vhost,omitempty"`
}

type Components1GzojhcSchemasListvhostsnapshotpresetresPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetPropertiesRecordtobItems struct {
	Duration     *int32  `json:"Duration,omitempty"`
	Format       *string `json:"Format,omitempty"`
	RecordObject *string `json:"RecordObject,omitempty"`
	Splice       *int32  `json:"Splice,omitempty"`
}

type Components1Je5O2CSchemasDescribepresetassociationresPropertiesResultPropertiesListItemsPropertiesSnapshotpresetlistv2Items struct {
	CreateTime string `json:"CreateTime"`

	PresetName string `json:"PresetName"`

	UpdateTime string `json:"UpdateTime"`
}

type Components1M64L84SchemasListvhostdetailresPropertiesResultPropertiesVhostlistItemsPropertiesTagsItems struct {
	Category string `json:"Category"`

	Key string `json:"Key"`

	Value string `json:"Value"`
}

type Components1Tzc8QlSchemasListvhostsnapshotpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesSnapshotpresetconfigPropertiesJpgparamPropertiesTosparam struct {
	ACL             *string `json:"ACL,omitempty"`
	AccessKey       *string `json:"AccessKey,omitempty"`
	Bucket          *string `json:"Bucket,omitempty"`
	Enable          *bool   `json:"Enable,omitempty"`
	ExactObject     *string `json:"ExactObject,omitempty"`
	OverwriteObject *string `json:"OverwriteObject,omitempty"`
	Region          *string `json:"Region,omitempty"`
	S3NetworkType   *string `json:"S3NetworkType,omitempty"`
	StorageDir      *string `json:"StorageDir,omitempty"`
	TosCluster      *string `json:"TosCluster,omitempty"`
	TosDC           *string `json:"TosDC,omitempty"`
	TosPSM          *string `json:"TosPSM,omitempty"`
}

type Components1UawxzeSchemasAddcommontranspresetbodyPropertiesPresetlistItemsPropertiesAbtestAdditionalproperties struct {
	Label *int32 `json:"Label,omitempty"`
	Ratio *int32 `json:"Ratio,omitempty"`
}

type Components1UxazjaSchemasListvhostsnapshotpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesSnapshotpresetconfigPropertiesJpgparamPropertiesImagexparam struct {
	ServiceID       string  `json:"ServiceID"`
	Enable          *bool   `json:"Enable,omitempty"`
	ExactObject     *string `json:"ExactObject,omitempty"`
	OverwriteObject *string `json:"OverwriteObject,omitempty"`
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

type Components1Yrp708SchemasDescribepresetassociationresPropertiesResultPropertiesListItemsPropertiesCdnsnapshotpresetlistv2Items struct {
	CreateTime string `json:"CreateTime"`

	PresetName string `json:"PresetName"`

	RecordType string `json:"RecordType"`

	UpdateTime string `json:"UpdateTime"`
}

type Components318E5PSchemasDescribepresetassociationresPropertiesResultPropertiesListItemsPropertiesSnapshotauditpresetlistv2Items struct {
	CreateTime string `json:"CreateTime"`

	PresetName string `json:"PresetName"`

	UpdateTime string `json:"UpdateTime"`
}

// Components44Na0KSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesFlvparam
// - FLV 录制参数，开启 FLV 录制时设置。
type Components44Na0KSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesFlvparam struct {
	Duration *int32 `json:"Duration,omitempty"`

	Enable *bool `json:"Enable,omitempty"`

	RealtimeRecordDuration *int32 `json:"RealtimeRecordDuration,omitempty"`

	Splice *int32 `json:"Splice,omitempty"`

	TOSParam *ComponentsBbqv7RSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesFlvparamPropertiesTosparam `json:"TOSParam,omitempty"`

	VODParam *ComponentsKovkk9SchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesFlvparamPropertiesVodparam `json:"VODParam,omitempty"`
}

type Components4Gwy08SchemasCreatetranscodepresetpatchbyadminbodyPropertiesPresetlistItemsPropertiesTranscodestructPropertiesAbtestAdditionalproperties struct {
	Label *int32 `json:"Label,omitempty"`
	Ratio *int32 `json:"Ratio,omitempty"`
}

type Components4Y1LroSchemasListvhostdetailresPropertiesResultPropertiesVhostlistItemsPropertiesDomainlistItems struct {
	CreateTime string `json:"CreateTime"`

	ProjectName string `json:"ProjectName"`

	Tags ListVhostDetailResResultVhostListItemDomainListItemTags `json:"Tags"`

	UpdateTime string  `json:"UpdateTime"`
	CNAME      *string `json:"CNAME,omitempty"`
	CertDomain *string `json:"CertDomain,omitempty"`
	CertName   *string `json:"CertName,omitempty"`
	ChainID    *string `json:"ChainID,omitempty"`
	Domain     *string `json:"Domain,omitempty"`
	Priority   *int32  `json:"Priority,omitempty"`
	PushDomain *string `json:"PushDomain,omitempty"`
	Region     *string `json:"Region,omitempty"`
	Status     *int32  `json:"Status,omitempty"`
	Type       *string `json:"Type,omitempty"`
	Vhost      *string `json:"Vhost,omitempty"`
}

type Components5Jn2JnSchemasDescribepresetassociationresPropertiesResultPropertiesListItemsPropertiesAvslicepresetlistv2Items struct {
	CreateTime string `json:"CreateTime"`

	PresetName string `json:"PresetName"`

	UpdateTime string `json:"UpdateTime"`
}

type Components7Eb4PfSchemasDescribepresetassociationresPropertiesResultPropertiesListItemsPropertiesAvextractorpresetlistv2Items struct {
	CompositionSubtitle string `json:"CompositionSubtitle"`

	CreateTime string `json:"CreateTime"`

	PresetName string `json:"PresetName"`

	TimeshiftSubtitle string `json:"TimeshiftSubtitle"`

	UpdateTime string `json:"UpdateTime"`
}

type ComponentsAer7PvSchemasDescribevhostresPropertiesResultPropertiesVhostlistItemsPropertiesDomainlistItems struct {
	CreateTime string `json:"CreateTime"`

	UpdateTime string `json:"UpdateTime"`

	CNAME *string `json:"CNAME,omitempty"`

	CertDomain *string `json:"CertDomain,omitempty"`

	ChainID *string `json:"ChainID,omitempty"`

	Domain *string `json:"Domain,omitempty"`

	Region *string `json:"Region,omitempty"`

	Type *string `json:"Type,omitempty"`
}

// ComponentsAoysk3SchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfigPropertiesHlsparam
// - HLS 录制参数，开启 HLS 录制时设置。
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

type ComponentsF9EcgzSchemasDescribepresetassociationresPropertiesResultPropertiesListItemsPropertiesDatamigrationpresetlistv2Items struct {
	CreateTime string `json:"CreateTime"`

	PresetName string `json:"PresetName"`

	UpdateTime string `json:"UpdateTime"`
}

type ComponentsFceumsSchemasListvqosmetricsdimensionsresPropertiesResultItemsPropertiesDimensionsItems struct {
	Alias string `json:"Alias"`

	Attribute string `json:"Attribute"`

	Name string `json:"Name"`
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

type ComponentsGhid1HSchemasDescribepresetassociationresPropertiesResultPropertiesListItemsPropertiesTimeshiftpresetlistv2Items struct {
	CreateTime string `json:"CreateTime"`

	NeedTranscode string `json:"NeedTranscode"`

	PresetName string `json:"PresetName"`

	TimeShiftType int32 `json:"TimeShiftType"`

	UpdateTime string `json:"UpdateTime"`
}

type ComponentsH8On9CSchemasDescribepresetassociationresPropertiesResultPropertiesListItemsPropertiesTranspresetlistItems struct {
	Codec string `json:"Codec"`

	CreateTime string `json:"CreateTime"`

	PresetName string `json:"PresetName"`

	StopInterval int32 `json:"StopInterval"`

	Suffix string `json:"Suffix"`

	Type string `json:"Type"`

	UpdateTime string `json:"UpdateTime"`
}

// ComponentsK46Cw0SchemasListvhostsnapshotpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesSnapshotpresetconfigPropertiesJpegparamPropertiesImagexparam
// - 截图存储到 veImageX 时的配置。 :::tip TOSParam 和 ImageXParam 配置且配置其中一个。 :::
type ComponentsK46Cw0SchemasListvhostsnapshotpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesSnapshotpresetconfigPropertiesJpegparamPropertiesImagexparam struct {
	Enable bool `json:"Enable"`

	ExactObject string `json:"ExactObject"`

	OverwriteObject string `json:"OverwriteObject"`

	ServiceID string `json:"ServiceID"`
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
// - MP4 录制参数，开启 MP4 录制时设置。
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

type ComponentsSgrw9KSchemasCreatetranscodepresetbatchbodyPropertiesPresetlistItemsPropertiesTranscodestructPropertiesAbtestAdditionalproperties struct {
	Label *int32 `json:"Label,omitempty"`
	Ratio *int32 `json:"Ratio,omitempty"`
}

// ComponentsSlabtaSchemasListvhostsnapshotpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesSnapshotpresetconfigPropertiesJpgparam
// - 截图格式为 JPG 时的截图参数。
type ComponentsSlabtaSchemasListvhostsnapshotpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesSnapshotpresetconfigPropertiesJpgparam struct {
	Enable      *bool                                                                                                                                                                                      `json:"Enable,omitempty"`
	ImageXParam *Components1UxazjaSchemasListvhostsnapshotpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesSnapshotpresetconfigPropertiesJpgparamPropertiesImagexparam `json:"ImageXParam,omitempty"`
	TOSParam    *Components1Tzc8QlSchemasListvhostsnapshotpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesSnapshotpresetconfigPropertiesJpgparamPropertiesTosparam    `json:"TOSParam,omitempty"`
}

type ComponentsXsjbgcSchemasCreatetranscodepresetbodyPropertiesTranscodestructPropertiesAbtestAdditionalproperties struct {
	Label *int32 `json:"Label,omitempty"`
	Ratio *int32 `json:"Ratio,omitempty"`
}

type CreateAppBody struct {

	// REQUIRED; app中文名称
	AppCnName string `json:"AppCnName"`

	// REQUIRED; app英文名称
	AppEnName string `json:"AppEnName"`

	// REQUIRED; app类型，WEB,APP,SERVICE,OTHERS
	AppType string `json:"AppType"`

	// REQUIRED; 项目名称
	Project string `json:"Project"`

	// REQUIRED; 地区，cn-north-1、as-singapore-1、us-east-1
	Region string `json:"Region"`

	// bundleID
	BundleID *string `json:"BundleID,omitempty"`

	// packageName
	PackageName *string `json:"PackageName,omitempty"`
}

type CreateAppRes struct {

	// REQUIRED
	ResponseMetadata CreateAppResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *CreateAppResResult `json:"Result,omitempty"`
}

type CreateAppResResponseMetadata struct {

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

// CreateAppResResult - 视请求的接口而定
type CreateAppResResult struct {

	// REQUIRED; appid
	AppID int32 `json:"AppID"`

	// REQUIRED; appKey
	AppKey string `json:"AppKey"`
}

type CreateAvSlicePresetBody struct {

	// REQUIRED
	Vhost       string  `json:"Vhost"`
	AccessKey   *string `json:"AccessKey,omitempty"`
	AccountID   *string `json:"AccountID,omitempty"`
	App         *string `json:"App,omitempty"`
	Bucket      *string `json:"Bucket,omitempty"`
	Callback    *string `json:"Callback,omitempty"`
	Description *string `json:"Description,omitempty"`
	NssConfig   *string `json:"NssConfig,omitempty"`
	Preset      *string `json:"Preset,omitempty"`
	Status      *int32  `json:"Status,omitempty"`
	Stream      *string `json:"Stream,omitempty"`
	TosCluster  *string `json:"TosCluster,omitempty"`
	TosDC       *string `json:"TosDC,omitempty"`
	TosPSM      *string `json:"TosPSM,omitempty"`
}

type CreateAvSlicePresetRes struct {

	// REQUIRED
	ResponseMetadata CreateAvSlicePresetResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type CreateAvSlicePresetResResponseMetadata struct {

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

type CreateCertBody struct {

	// REQUIRED; 密钥信息
	Rsa CreateCertBodyRsa `json:"Rsa"`

	// REQUIRED; 证书用途，默认为 https，支持的取值包括：
	// * https：https 认证；
	// * sign：签名校验。
	UseWay string `json:"UseWay"`

	// 证书名称
	CertName *string `json:"CertName,omitempty"`

	// 证书链 ID，用于标识整个证书链，包括叶子证书（服务器证书）、中间证书（中间 CA 证书）以及根证书（根 CA 证书）
	ChainID *string `json:"ChainID,omitempty"`

	// 是否是客户自定义的证书链，如果是则跳过证书合法性校验。不填默认为false。
	UserDefinedChain *bool `json:"UserDefinedChain,omitempty"`
}

// CreateCertBodyRsa - 密钥信息
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

	// 证书 ID。
	ChainID *string `json:"ChainID,omitempty"`

	// 使用该证书的域名。
	Domain *string `json:"Domain,omitempty"`

	// 证书用途，包括两种取值。
	// * https：HTTPS 认证；
	// * sign：签名校验。
	UseWay *string `json:"UseWay,omitempty"`
}

type CreateCustomLogConfigBody struct {

	// REQUIRED; 用户账号id
	AccountID string `json:"AccountId"`

	// REQUIRED; 用户账号名称
	AccountName string `json:"AccountName"`

	// REQUIRED; 日志获取接口名称
	ActionName string `json:"ActionName"`

	// REQUIRED; bmq集群名
	BmqCluster string `json:"BmqCluster"`

	// REQUIRED; tce集群名
	Cluster string `json:"Cluster"`

	// REQUIRED; 创建人
	Creator string `json:"Creator"`

	// REQUIRED; 下载时填的Type参数
	DownloadType string `json:"DownloadType"`

	// REQUIRED; 文件名字段名称
	FileNameFields CreateCustomLogConfigBodyFileNameFields `json:"FileNameFields"`

	// REQUIRED; 文件名pattern
	FileNamePattern string `json:"FileNamePattern"`

	// REQUIRED; 日志字段名称
	LogFields CreateCustomLogConfigBodyLogFields `json:"LogFields"`

	// REQUIRED; 日志pattern
	LogPattern string `json:"LogPattern"`

	// REQUIRED; 日志类型，如果是多个用逗号连接，全选可填*
	LogType string `json:"LogType"`

	// REQUIRED; 写入的topic
	Topic string `json:"Topic"`

	// REQUIRED; 特殊清洗状态，false:表示数仓单独任务进行清洗，true为通用清洗任务
	WashStatus bool `json:"WashStatus"`

	// 延迟时间，默认300s
	DelayTime *int32 `json:"DelayTime,omitempty"`

	// 填1或0，是否补空文件，默认为0
	EmptyFile *int32 `json:"EmptyFile,omitempty"`

	// 排除的账号ID
	ExcludedAccountIDs *string `json:"ExcludedAccountIds,omitempty"`

	// 如果没有Id，表示创建，带了Id表示更新
	ID *int32 `json:"Id,omitempty"`

	// 是否每个域名一个文件，默认为false
	SplitDomain *bool `json:"SplitDomain,omitempty"`

	// 默认false（前端默认填ture），同一个时间范围是否允许按照大小切割文件
	SplitFile *bool `json:"SplitFile,omitempty"`

	// 切割文件的行数，默认120w
	SplitLine *int32 `json:"SplitLine,omitempty"`

	// 切割文件的时间，单位秒，默认3600
	SplitTime *int32 `json:"SplitTime,omitempty"`

	// 默认0，状态，1：启动，0：禁止
	Status *int32 `json:"Status,omitempty"`
}

// CreateCustomLogConfigBodyFileNameFields - 文件名字段名称
type CreateCustomLogConfigBodyFileNameFields struct {

	// REQUIRED; 字段名称
	Key string `json:"Key"`

	// REQUIRED; 字段类型，不能为空
	Type string `json:"Type"`

	// 备注信息，没有可以为空
	FmtValue *string `json:"FmtValue,omitempty"`

	// 字段对应中文名
	KeyCn *string `json:"KeyCn,omitempty"`

	// 敏感词替换字符串，比如：ab,cd 表示用cd替换ab，如果有多组替换用分号连接
	Transform *int32 `json:"Transform,omitempty"`
}

// CreateCustomLogConfigBodyLogFields - 日志字段名称
type CreateCustomLogConfigBodyLogFields struct {

	// REQUIRED; 字段名称
	Key string `json:"Key"`

	// REQUIRED; 字段类型，不能为空
	Type string `json:"Type"`

	// 备注信息，没有可以为空
	FmtValue *string `json:"FmtValue,omitempty"`

	// 字段对应中文名
	KeyCn *string `json:"KeyCn,omitempty"`

	// 敏感词替换字符串，比如：ab,cd 表示用cd替换ab，如果有多组替换用分号连接
	Transform *int32 `json:"Transform,omitempty"`
}

type CreateCustomLogConfigRes struct {

	// REQUIRED
	ResponseMetadata CreateCustomLogConfigResResponseMetadata `json:"ResponseMetadata"`
	Result           *CreateCustomLogConfigResResult          `json:"Result,omitempty"`
}

type CreateCustomLogConfigResResponseMetadata struct {

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

type CreateCustomLogConfigResResult struct {

	// REQUIRED; 配置的Id
	ID string `json:"Id"`
}

type CreateDenseSnapshotPresetBody struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字幕、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 域名空间名称。
	Vhost     string  `json:"Vhost"`
	AccessKey *string `json:"AccessKey,omitempty"`
	AsLong    *int32  `json:"AsLong,omitempty"`
	AsShort   *int32  `json:"AsShort,omitempty"`

	// ToS 存储的 Bucket。 :::tipBucket 与 ServiceID 传且仅传一个。 :::
	Bucket *string `json:"Bucket,omitempty"`

	// 回调地址，支持 HTTP 和 HTTPS 的回调地址。如果同时使用 UpdateCallback 配置了回调地址，则此处回调地址配置优先级更高。
	CallbackURL *string `json:"CallbackUrl,omitempty"`
	Describe    *string `json:"Describe,omitempty"`
	Field36     *string `json:"Field36,omitempty"`
	Format      *string `json:"Format,omitempty"`
	Height      *int32  `json:"Height,omitempty"`

	// 截图间隔时间，单位为 s，默认为 10s，取值范围为正整数。
	Interval        *int32  `json:"Interval,omitempty"`
	IsTobSnapshot   *int32  `json:"IsTobSnapshot,omitempty"`
	KafkaCluster    *string `json:"KafkaCluster,omitempty"`
	KafkaTopic      *string `json:"KafkaTopic,omitempty"`
	Object          *string `json:"Object,omitempty"`
	OverwriteObject *string `json:"OverwriteObject,omitempty"`
	PlatformType    *string `json:"PlatformType,omitempty"`

	// 密集抽帧截图配置模板名称。
	Preset         *string `json:"Preset,omitempty"`
	Product        *string `json:"Product,omitempty"`
	Quality        *int32  `json:"Quality,omitempty"`
	Rate           *int32  `json:"Rate,omitempty"`
	Region         *string `json:"Region,omitempty"`
	S3NetworkType  *int32  `json:"S3NetworkType,omitempty"`
	SequenceObject *string `json:"SequenceObject,omitempty"`

	// veImageX 的服务 ID。 :::tipBucket 与 ServiceID 传且仅传一个。 :::
	ServiceID *string `json:"ServiceID,omitempty"`

	// 截图格式，支持 jpg 和 png，默认为 jpg。
	SnapshotFormat *string `json:"SnapshotFormat,omitempty"`

	// 存储规则。
	SnapshotObject *string `json:"SnapshotObject,omitempty"`

	// 密集截图配置模板的开启状态，默认为开启。
	// * 1：开启；
	// * 0：关闭。
	Status *int32 `json:"Status,omitempty"`

	// ToS 存储目录，不传为空。
	StorageDir      *string `json:"StorageDir,omitempty"`
	TosCluster      *string `json:"TosCluster,omitempty"`
	TosType         *int32  `json:"TosType,omitempty"`
	TranscodeSuffix *string `json:"TranscodeSuffix,omitempty"`
	Width           *int32  `json:"Width,omitempty"`
}

type CreateDenseSnapshotPresetRes struct {

	// REQUIRED
	ResponseMetadata CreateDenseSnapshotPresetResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result CreateDenseSnapshotPresetResResult `json:"Result"`
}

type CreateDenseSnapshotPresetResResponseMetadata struct {

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

type CreateDenseSnapshotPresetResResult struct {

	// REQUIRED; 模板名称。
	PresetName string `json:"PresetName"`
}

type CreateDomainBody struct {

	// REQUIRED; 推/拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 域名类型，包含两种类型。
	// * push：推流域名；
	// * pull-flv：拉流域名，包含 RTMP、FLV、HLS 格式。
	Type string `json:"Type"`

	// 区域，默认指为 cn，包含以下类型。
	// * cn：中国大陆；
	// * cn-global：全球；
	// * cn-oversea：海外及港澳台。
	Region *string `json:"Region,omitempty"`
}

type CreateDomainRes struct {

	// REQUIRED
	ResponseMetadata CreateDomainResResponseMetadata `json:"ResponseMetadata"`
}

type CreateDomainResResponseMetadata struct {

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
	Error   *CreateDomainResResponseMetadataError `json:"Error,omitempty"`
}

type CreateDomainResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type CreateDomainV2Body struct {

	// REQUIRED; 域名列表，总和最多十个。
	Domains []CreateDomainV2BodyDomainsItem `json:"Domains"`

	// REQUIRED; 区域，包含以下类型。
	// * cn：中国大陆；
	// * cn-global：全球；
	// * cn-oversea：海外及港澳台。
	Region string `json:"Region"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 项目名称，vhost 将归类这个项目里，仅在新创建 vhost 时需要设置。
	ProjectName *string `json:"ProjectName,omitempty"`

	// 标签列表，vhost 将归类这个 tag 里。
	Tags []*CreateDomainV2BodyTagsItem `json:"Tags,omitempty"`

	// 是否进行域名归属校验，不填默认需要校验
	VerifyCheck *bool `json:"VerifyCheck,omitempty"`
}

type CreateDomainV2BodyDomainsItem struct {

	// REQUIRED; 域名名称。
	DomainName string `json:"DomainName"`

	// REQUIRED; 域名类型，支持以下取值。
	// * push：推流域名
	// * pull-flv：拉流域名
	Type string `json:"Type"`

	// 证书 ID。
	ChainID *string `json:"ChainID,omitempty"`
}

type CreateDomainV2BodyTagsItem struct {

	// REQUIRED; 标签类型，支持以下取值。
	// * System：系统内置标签
	// * Custom：自定义标签
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

type CreateLiveAccountFeeConfigBody struct {

	// REQUIRED; 账号
	AccountID string `json:"AccountId"`

	// REQUIRED; 进制
	Base int32 `json:"Base"`

	// REQUIRED; 创建者
	Creator string `json:"Creator"`

	// REQUIRED; 上浮系数
	Factor float32 `json:"Factor"`

	// REQUIRED; 如果id是0，表示创建，否则表示更新
	ID int32 `json:"Id"`

	// REQUIRED; 是否开启闲忙时，True表示开启，false表示关闭
	StageEnable bool `json:"StageEnable"`

	// REQUIRED; 上浮系数生效时间
	StartTime string `json:"StartTime"`

	// 免流的类型
	FreeFeeList []*string `json:"FreeFeeList,omitempty"`
}

type CreateLiveAccountFeeConfigRes struct {

	// REQUIRED
	ResponseMetadata CreateLiveAccountFeeConfigResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *CreateLiveAccountFeeConfigResResult `json:"Result,omitempty"`
}

type CreateLiveAccountFeeConfigResResponseMetadata struct {

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

// CreateLiveAccountFeeConfigResResult - 视请求的接口而定
type CreateLiveAccountFeeConfigResResult struct {

	// REQUIRED; 账号id
	AccountID string `json:"AccountId"`

	// REQUIRED; 进制
	Base string `json:"Base"`

	// REQUIRED; 创建时间
	CreateTime string `json:"CreateTime"`

	// REQUIRED; 配置创建者
	Creator string `json:"Creator"`

	// REQUIRED; 上浮系数
	Factor string `json:"Factor"`

	// REQUIRED; 免流配置
	FreeFeeList []string `json:"FreeFeeList"`

	// REQUIRED; 配置id
	ID string `json:"Id"`

	// REQUIRED; 闲忙时开关
	StageEnable string `json:"StageEnable"`

	// REQUIRED; 闲忙时生效时间
	StageTime string `json:"StageTime"`

	// REQUIRED; 上浮系数生效时间
	StartTime string `json:"StartTime"`

	// REQUIRED; 更新时间
	UpdateTime string `json:"UpdateTime"`
}

type CreateProxyConfigBody struct {

	// REQUIRED; 生效类型，1：默认生效
	EffectType int32 `json:"EffectType"`

	// REQUIRED; 代理模式，0：固定模式，1：解析模式
	Mode int32 `json:"Mode"`

	// REQUIRED; 代理名称
	Name string `json:"Name"`

	// REQUIRED; 代理��表
	ProxyConfigList []CreateProxyConfigBodyProxyConfigListItem `json:"ProxyConfigList"`

	// 账号
	AccountID *string `json:"AccountID,omitempty"`

	// 是否与账号关联，associate:关联
	AssociateType *string `json:"AssociateType,omitempty"`

	// 配置级别，overall:全局，single:单客户
	ConfigLevel *string `json:"ConfigLevel,omitempty"`

	// 描述
	Description *string `json:"Description,omitempty"`
}

type CreateProxyConfigBodyProxyConfigListItem struct {

	// REQUIRED; 集群
	Cluster string `json:"Cluster"`

	// REQUIRED; 机房
	IDC string `json:"IDC"`

	// REQUIRED; 运营商
	ISP string `json:"ISP"`

	// REQUIRED; 地址列表
	ProxyList []CreateProxyConfigBodyProxyConfigListPropertiesItemsItem `json:"ProxyList"`
}

type CreateProxyConfigBodyProxyConfigListPropertiesItemsItem struct {

	// REQUIRED; 地址
	URL string `json:"URL"`

	// REQUIRED; 权重
	Weight int32 `json:"Weight"`
}

type CreateProxyConfigRes struct {

	// REQUIRED
	ResponseMetadata CreateProxyConfigResResponseMetadata `json:"ResponseMetadata"`
}

type CreateProxyConfigResResponseMetadata struct {

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

type CreatePullCDNSnapshotTaskBody struct {

	// REQUIRED; app名字
	App string `json:"App"`

	// REQUIRED; 域名
	Domain string `json:"Domain"`

	// REQUIRED; stream名称
	Stream string `json:"Stream"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty"`

	// 拉流地址
	StreamURL *string `json:"StreamURL,omitempty"`

	// 域名空间
	Vhost *string `json:"Vhost,omitempty"`
}

type CreatePullCDNSnapshotTaskRes struct {

	// REQUIRED
	ResponseMetadata CreatePullCDNSnapshotTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result CreatePullCDNSnapshotTaskResResult `json:"Result"`
}

type CreatePullCDNSnapshotTaskResResponseMetadata struct {

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

type CreatePullCDNSnapshotTaskResResult struct {

	// REQUIRED; 任务id
	TaskID string `json:"TaskId"`
}

type CreatePullRecordTaskBody struct {

	// REQUIRED; app名字
	App string `json:"App"`

	// REQUIRED; 域名
	Domain string `json:"Domain"`

	// REQUIRED; stream名称
	Stream string `json:"Stream"`

	// REQUIRED; 拉流地址
	StreamURL string `json:"StreamURL"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty"`

	// 域名空间
	Vhost *string `json:"Vhost,omitempty"`
}

type CreatePullRecordTaskRes struct {

	// REQUIRED
	ResponseMetadata CreatePullRecordTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 视请求的接口而定
	Result CreatePullRecordTaskResResult `json:"Result"`
}

type CreatePullRecordTaskResResponseMetadata struct {

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

// CreatePullRecordTaskResResult - 视请求的接口而定
type CreatePullRecordTaskResResult struct {

	// REQUIRED; 任务id
	TaskID string `json:"TaskId"`
}

type CreatePullToPushTaskBody struct {

	// REQUIRED; 任务的结束时间，Unix 时间戳，单位为秒。 :::tip 拉流转推任务持续时间最长为 7 天。 :::
	EndTime int32 `json:"EndTime"`

	// REQUIRED; 设置点播视频转推至第三方推流域名时是否使用推流优先级参数，缺省情况下表示不使用此参数，支持的取值及含义如下。
	// * true：使用
	//
	//
	// * false：不使用 :::tip
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

	// 续播策略，续播策略指转推点播视频进行直播时出现断流并恢复后，如何继续播放的策略，拉流���源类型为点播视频（Type 为 1）时参数生效，支持的取值及含义如下。
	// * 0：从断流处续播（默认值）；
	// * 1：从断流处+自然流逝时长处续播。
	ContinueStrategy *int32 `json:"ContinueStrategy,omitempty"`

	// 点播视频文件循环播放模式，当拉流来源类型为点播视频（Type 为 1）时为必选参数，参数取值及含义如下所示。
	// * -1：无限循环，至任务结束；
	// * 0：有限次循环，循环次数为 PlayTimes 取值为准。
	CycleMode *int32 `json:"CycleMode,omitempty"`

	// 推流域名，推流地址（DstAddr）为空时必传；反之，则该参数不生效。
	Domain *string `json:"Domain,omitempty"`

	// 推流地址，即直播源或点播视频转推的目标地址。
	DstAddr *string `json:"DstAddr,omitempty"`

	// 点播文件启播时间偏移值，单位为秒，仅当点播视频播放地址列表（SrcAddrS）只有一个地址，且未配置 Offsets 时生效，缺省情况下为空表示不进行偏移。
	Offset *float32 `json:"Offset,omitempty"`

	// 点播文件启播时间偏移值，单位为秒，数量与拉流地址列表中地址数量相等，缺省情况下为空表示不进行偏移。 拉流来源类型为点播视频（Type 为 1）时，参数生效。
	OffsetS []*float32 `json:"OffsetS,omitempty"`

	// 点播视频文件循环播放次数，当循环播放模式为有限次循环（CycleMode为0）时为必选参数。
	PlayTimes *int32 `json:"PlayTimes,omitempty"`

	// 是否开启点播预热，开启点播预热后，系统会自动将点播视频文件缓存到 CDN 节点上，当用户请求直播时，可以直播从 CDN 节点获取视频，从而提高直播流畅度。 拉流来源类型为点播视频（Type 为 1）时，参数生效。
	// * 0：不开启；
	// * 1：开启（默认值）。
	PreDownload *int32 `json:"PreDownload,omitempty"`

	// 直播源的拉流地址，拉流来源类型为直播源（Type 为 0）时，为必选参数，最大长度为 1000 个字符。
	SrcAddr *string `json:"SrcAddr,omitempty"`

	// 点播视频播放地址列表，拉流来源类型为点播视频（Type 为 1）时，为必选参数，最多支持传入 30 个点播视频播放地址，每个地址最大长度为 1000 个字符。
	SrcAddrS []*string `json:"SrcAddrS,omitempty"`

	// 推流的流名称，推流地址（DstAddr）为空时必传；反之，则该参数不生效。
	Stream *string `json:"Stream,omitempty"`

	// 拉流转推任务的名称，由 1 到 20 位中文、大小写字母和数字组成，默认为空，表示不配置任务名称。
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

	// REQUIRED; 应用名称，取值与直播流地址的 AppName 字段取值相同，支持填写星号（*）或由 1 到 30 位数字（0 - 9）、大写小字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成。
	// * 应用名称填写星号，即 App 取值为 *时， 流名称也需填写星号，此时表示录制配置为域名空间级别的配置，即直播流使用的域名属于此域名空间时，就会使用此配置进行录制。
	// * 应用名称填写非星号时，表示录制配置为域名空间 + 应用名称 + 流名称级别的配置，即直播流使用的域名属于此域名空间，且 AppName 和 StreamName 字段也同时与 App 和 Stream 的取值匹配时，就会使用此配置进行录制。
	// :::warning 当 App 取值为 * 时，Stream 取值必须为 *。 :::
	App string `json:"App"`

	// REQUIRED; 直播流录制配置的详细配置。
	RecordPresetConfig CreateRecordPresetV2BodyRecordPresetConfig `json:"RecordPresetConfig"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console-stable.volcanicengine.com/live/main/domain/list]页面，查看需要录制的直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 流名称，取值与直播流地址的 StreamName 字段取值相同，支持填写星号（*）或由 1 到 100 位数字（0 - 9）、字母、下划线（_）、短横线（-）和句点（.）组成。
	// * 流名称填写星号时，表示录制配置为应用名称级别的配置，即直播流使用的域名属于此域名空间，且 AppName 字段也与 App 取值同时匹配时，就会使用此配置进行录制。
	// * 流名称填写非星号时，表示录制配置为域名空间 + 应用名称 + 流名称级别的配置，即直播流使用的域名属于此域名空间，且 AppName 和 StreamName 字段也同时与 App 和 Stream 的取值匹配时，就会使用此配置进行录制。
	// :::warning 当 App 取值为 * 时，Stream 取值必须为 *。 :::
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

	// 是否源流录制，默认值为 0，支持的取值即含义如下所示。
	// * 0：不录制；
	// * 1：录制。
	// :::tip 转码流和源流需至少选一个进行录制，即是否录制转码流（TranscodeRecord）和是否录制源流（OriginRecord）的取值至少一个不为 0。 :::
	OriginRecord *int32 `json:"OriginRecord,omitempty"`

	// 录制为 HLS 格式时，单个 TS 切片时长，单位为秒，默认值为 10，取值范围为 [5,30]。
	SliceDuration *int32 `json:"SliceDuration,omitempty"`

	// 是否录制转码流，默认值为 0，支持的取值及含义如下所示。
	// * 0：不录制；
	// * 1：录制全部转码流；
	// * 2：通过转码流后缀列表，即TranscodeSuffixList 字段取值匹配转码流。
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
	// * 取值为 [300,86400] 之间的值时，表示根据设置的录制文件时长生成录制文件，完成录制后一起上传。
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
	TOSParam *CreateRecordPresetV2BodyRecordPresetConfigFlvParamTOSParam `json:"TOSParam,omitempty"`

	// VOD 存储相关配置。 :::tip 录制文件只能选择一个位置进行存储，即 TOSParam 和 VODParam 配置且配置其中一个。 :::
	VODParam *CreateRecordPresetV2BodyRecordPresetConfigFlvParamVODParam `json:"VODParam,omitempty"`
}

// CreateRecordPresetV2BodyRecordPresetConfigFlvParamTOSParam - TOS 存储相关配置。 :::tip 录制文件只能选择一个位置进行存储，即 TOSParam 和 VODParam
// 配置且配置其中一个。 :::
type CreateRecordPresetV2BodyRecordPresetConfigFlvParamTOSParam struct {

	// TOS 存储对应的 Bucket。例如，存储位置为 live-test-tos-example/live/liveapp 时，Bucket 取值为 live-test-tos-example。 :::tip 如果启用 TOSParam 配置（Enable
	// 取值为 true），则 Bucket 必填。 :::
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

	// 视频点播（VOD）空间名称。可登录视频点播控制台 [https://console.volcengine.com/vod/]查询。 :::tip 如果启用 VODParam 配置（Enable 取值为 true），则 VodNamespace
	// 必填。 :::
	VodNamespace *string `json:"VodNamespace,omitempty"`

	// 视频点播工作流模板 ID，对于存储在点播的录制文件，会使用该工作流模版对录制的视频进行处理，可登录视频点播控制台 [https://console.volcengine.com/vod/]获取工作流模板 ID，默认为空。
	WorkflowID *string `json:"WorkflowID,omitempty"`
}

// CreateRecordPresetV2BodyRecordPresetConfigHlsParam - 录制为 HLS 合适时的录制参数。 :::tip 您需至少配置一种录制格式，即 FlvParam、HlsParam、Mp4Param
// 至少开启一个。 :::
type CreateRecordPresetV2BodyRecordPresetConfigHlsParam struct {

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

	// REQUIRED; 直播流使用的域名。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console-stable.volcanicengine.com/live/main/domain/list]页面，查看直播流使用的域名。
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

type CreateSDKBody struct {

	// REQUIRED; 应用ID
	AppID string `json:"AppID"`

	// REQUIRED; 应用名称,长度小于129
	AppName string `json:"AppName"`

	// REQUIRED; 应用英文名称，长度小于31
	AppNameEN string `json:"AppNameEN"`

	// BundleID，和packageName二选一必填
	BundleID *string `json:"BundleID,omitempty"`

	// License 类型,0:无版本，1:基础版本，2：高级版本，3：试用版
	LicenseType *int32 `json:"LicenseType,omitempty"`

	// 流量包ID
	PackageID *string `json:"PackageID,omitempty"`

	// 包名，和bundleID二选一必填
	PackageName *string `json:"PackageName,omitempty"`

	// 应用类型，App, Web二选一，不填默认为App
	SDKType *string `json:"SDKType,omitempty"`

	// SDK版本，精简版：1，互动版：2，已经弃用
	SDKVersion *string `json:"SDKVersion,omitempty"`
}

type CreateSDKRes struct {

	// REQUIRED
	ResponseMetadata CreateSDKResResponseMetadata `json:"ResponseMetadata"`
}

type CreateSDKResResponseMetadata struct {
	Action    *string                            `json:"Action,omitempty"`
	Error     *CreateSDKResResponseMetadataError `json:"Error,omitempty"`
	Region    *string                            `json:"Region,omitempty"`
	RequestID *string                            `json:"RequestId,omitempty"`
	Service   *string                            `json:"Service,omitempty"`
	Version   *string                            `json:"Version,omitempty"`
}

type CreateSDKResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type CreateSnapshotAuditPresetBody struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 截图审核结果回调地址配置。
	CallbackDetailList []CreateSnapshotAuditPresetBodyCallbackDetailListItem `json:"CallbackDetailList"`

	// REQUIRED; 截图间隔时间，单位为秒，取值范围为 [0.1,10]，支持保留两位小数。
	Interval float32 `json:"Interval"`

	// REQUIRED; 存储策略，支持的取值及含义如下。
	// * 0：触发存储，只存储有风险图片；
	// * 1：全部存储，存储所有图片。
	StorageStrategy int32   `json:"StorageStrategy"`
	AshePresetName  *string `json:"AshePresetName,omitempty"`
	AuditType       *string `json:"AuditType,omitempty"`

	// TOS 存储对应的 Bucket。 例如，存储路径为 live-test-tos-example/live/liveapp 时，Bucket 取值为 live-test-tos-example。 :::tip 参数 Bucket 和 ServiceID
	// 传且仅传一个。 :::
	Bucket *string `json:"Bucket,omitempty"`

	// 截图审核配置的描述。
	Description *string `json:"Description,omitempty"`

	// 推流域名。 :::tip 参数 Domain 和 Vhost 传且仅传一个。 :::
	Domain *string `json:"Domain,omitempty"`

	// 审核标签，缺省情况下取值为 301、302、303、305 和 306，支持的取值及含义如下。
	// * 301：涉黄；
	// * 302：涉敏1；
	// * 303：涉敏2；
	// * 304：广告；
	// * 305：引人不适；
	// * 306：违禁；
	// * 307：二维码；
	// * 308：诈骗；
	// * 309：不良画面；
	// * 310：未成年相关；
	// * 320：文字违规。
	Label      []*string `json:"Label,omitempty"`
	PresetName *string   `json:"PresetName,omitempty"`

	// veImageX 的服务 ID。 :::tip 参数 Bucket 和 ServiceID 传且仅传一个。 :::
	ServiceID *string `json:"ServiceID,omitempty"`

	// 截图存储规则，支持以 {Domain}/{App}/{Stream}/{UnixTimestamp} 样式设置存储规则，支持输入字母、数字、-、!、_、.、* 及占位符，最大长度为 180 个字符，默认值为 {audit}/{PushDomain}/{App}/{Stream}/{UnixTimestamp}。
	SnapshotObject *string `json:"SnapshotObject,omitempty"`
	Status         *int32  `json:"Status,omitempty"`

	// ToS 存储对应的 bucket 下的存储目录，默认为空。 例如，存储位置为 live-test-tos-example/live/liveapp 时，StorageDir 取值为 live/liveapp。
	StorageDir *string `json:"StorageDir,omitempty"`

	// 域名空间名称。 :::tip 参数 Domain 和 Vhost 传且仅传一个。 :::
	Vhost *string `json:"Vhost,omitempty"`
}

type CreateSnapshotAuditPresetBodyCallbackDetailListItem struct {

	// REQUIRED; 回调地址的类型，当前仅支持 http。
	CallbackType string `json:"CallbackType"`

	// REQUIRED; 回调地址。
	URL string `json:"URL"`
}

type CreateSnapshotAuditPresetRes struct {

	// REQUIRED
	ResponseMetadata CreateSnapshotAuditPresetResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type CreateSnapshotAuditPresetResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                             `json:"Version"`
	Error   *CreateSnapshotAuditPresetResResponseMetadataError `json:"Error,omitempty"`
}

type CreateSnapshotAuditPresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type CreateSnapshotPresetBody struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 截图间隔时间，单位为 s，默认值为 10，取值范围为正整数。
	Interval int32 `json:"Interval"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// REQUIRED
	VodNamespace string  `json:"VodNamespace"`
	ACL          *string `json:"ACL,omitempty"`
	AccessKey    *string `json:"AccessKey,omitempty"`
	AccountID    *string `json:"AccountID,omitempty"`
	AsShort      *int32  `json:"AsShort,omitempty"`

	// ToS 存储的 Bucket。 :::tipBucket 与 ServiceID 传且仅传一个。 :::
	Bucket   *string `json:"Bucket,omitempty"`
	Callback *string `json:"Callback,omitempty"`

	// 回调详情。
	CallbackDetailList []*CreateSnapshotPresetBodyCallbackDetailListItem `json:"CallbackDetailList,omitempty"`
	Description        *string                                           `json:"Description,omitempty"`
	Duration           *int32                                            `json:"Duration,omitempty"`
	Format             []*string                                         `json:"Format,omitempty"`
	Height             *int32                                            `json:"Height,omitempty"`
	NssConfig          *string                                           `json:"NssConfig,omitempty"`

	// 存储方式为覆盖截图时的存储规则，支持以 {Domain}/{App}/{Stream} 样式设置存储规则，支持输入字母、数字、"-"、"!"、"_"、"."、"*"及占位符。
	OverwriteObject  *string                                  `json:"OverwriteObject,omitempty"`
	PlatformTypeList []*string                                `json:"PlatformTypeList,omitempty"`
	Preset           *string                                  `json:"Preset,omitempty"`
	PullDomain       *string                                  `json:"PullDomain,omitempty"`
	Quality          *int32                                   `json:"Quality,omitempty"`
	RecordConfig     *string                                  `json:"RecordConfig,omitempty"`
	RecordObject     *string                                  `json:"RecordObject,omitempty"`
	RecordTob        []*CreateSnapshotPresetBodyRecordTobItem `json:"RecordTob,omitempty"`
	RegionConfig     *string                                  `json:"RegionConfig,omitempty"`
	ReserveDays      *int32                                   `json:"ReserveDays,omitempty"`

	// veImageX 的服务 ID。 :::tipBucket 与 ServiceID 传且仅传一个。 :::
	ServiceID      *string `json:"ServiceID,omitempty"`
	SliceDuration  *int32  `json:"SliceDuration,omitempty"`
	SnapshotConfig *string `json:"SnapshotConfig,omitempty"`

	// 截图格式。默认值为 jpeg，支持如下取值。
	// * jpeg
	// * jpg
	SnapshotFormat *string `json:"SnapshotFormat,omitempty"`

	// 存储方式为实时存储时的存储规则，支持以 {Domain}/{App}/{Stream}/{UnixTimestamp} 样式设置存储规则，支持输入字母、数字、"-"、"!"、"_"、"."、"*"及占位符。
	SnapshotObject *string `json:"SnapshotObject,omitempty"`
	Splice         *int32  `json:"Splice,omitempty"`

	// 截图模版状态状态。默认开启。
	// * 1：开启。
	// * 0：关闭。
	Status *int32 `json:"Status,omitempty"`

	// ToS 存储目录，不传为空。
	StorageDir *string `json:"StorageDir,omitempty"`
	TosCluster *string `json:"TosCluster,omitempty"`
	TosDC      *string `json:"TosDC,omitempty"`
	TosPSM     *string `json:"TosPSM,omitempty"`
	Width      *int32  `json:"Width,omitempty"`
	WorkflowID *string `json:"WorkflowID,omitempty"`
}

type CreateSnapshotPresetBodyCallbackDetailListItem struct {

	// 回调类型，默认值为 http。
	CallbackType *string `json:"CallbackType,omitempty"`

	// 回调地址。
	URL *string `json:"URL,omitempty"`
}

type CreateSnapshotPresetBodyRecordTobItem struct {
	Duration     *int32  `json:"Duration,omitempty"`
	Format       *string `json:"Format,omitempty"`
	RecordObject *string `json:"RecordObject,omitempty"`
	Splice       *int32  `json:"Splice,omitempty"`
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

type CreateSnapshotPresetV2Body struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 截图配置的详细参数配置。
	SnapshotPresetConfig CreateSnapshotPresetV2BodySnapshotPresetConfig `json:"SnapshotPresetConfig"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 截图配置生效状态，默认为生效。
	// * 1：生效；
	// * 0：不生效。
	Status *int32 `json:"Status,omitempty"`
}

// CreateSnapshotPresetV2BodySnapshotPresetConfig - 截图配置的详细参数配置。
type CreateSnapshotPresetV2BodySnapshotPresetConfig struct {

	// 截图间隔时间，单位为秒，默认值为 10，取值范围为正整数。
	Interval *int32 `json:"Interval,omitempty"`

	// 图片格式为 JPEG 时的截图参数，开启 JPEG 截图时设置。 :::tip JPEG 截图和 JPG 截图必须开启且只能开启一个。 :::
	JPEGParam *CreateSnapshotPresetV2BodySnapshotPresetConfigJPEGParam `json:"JpegParam,omitempty"`

	// 截图格式为 JPG 时的截图参数，开启 JPG 截图时设置。 :::tip JPEG 截图和 JPG 截图必须开启且只能开启一个。 :::
	JpgParam *CreateSnapshotPresetV2BodySnapshotPresetConfigJpgParam `json:"JpgParam,omitempty"`
}

// CreateSnapshotPresetV2BodySnapshotPresetConfigJPEGParam - 图片格式为 JPEG 时的截图参数，开启 JPEG 截图时设置。 :::tip JPEG 截图和 JPG 截图必须开启且只能开启一个。
// :::
type CreateSnapshotPresetV2BodySnapshotPresetConfigJPEGParam struct {

	// 当前格式的截图是否开启，默认为 false，取值及含义如下所示。
	// * false：不开启；
	// * true：开启。
	Enable *bool `json:"Enable,omitempty"`

	// 截图存储到 veImageX 时的配置。 :::tip TOSParam 和 ImageXParam 配置且配置其中一个。 :::
	ImageXParam *CreateSnapshotPresetV2BodySnapshotPresetConfigJPEGParamImageXParam `json:"ImageXParam,omitempty"`

	// 截图存储到 TOS 时的配置。 :::tip TOSParam 和 ImageXParam 配置且配置其中一个。 :::
	TOSParam *CreateSnapshotPresetV2BodySnapshotPresetConfigJPEGParamTOSParam `json:"TOSParam,omitempty"`
}

// CreateSnapshotPresetV2BodySnapshotPresetConfigJPEGParamImageXParam - 截图存储到 veImageX 时的配置。 :::tip TOSParam 和 ImageXParam
// 配置且配置其中一个。 :::
type CreateSnapshotPresetV2BodySnapshotPresetConfigJPEGParamImageXParam struct {

	// 截图是否使用 veImageX 存储，默认为 false，取值及含义如下所示。
	// * false：不使用；
	// * true：使用。
	Enable *bool `json:"Enable,omitempty"`

	// 存储方式为实时截图时的存储规则，支持以 {Domain}/{App}/{Stream}/{UnixTimestamp} 样式设置存储规则，支持输入字母、数字、-、!、_、.、* 及占位符。 :::tip 参数 ExactObject 和
	// OverwriteObject 传且仅传一个。 :::
	ExactObject *string `json:"ExactObject,omitempty"`

	// 存储方式为覆盖截图时的存储规则，支持以 {Domain}/{App}/{Stream} 样式设置存储规则，支持输入字母、数字、-、!、_、.、* 及占位符。 :::tip 参数 ExactObject 和 OverwriteObject
	// 传且仅传一个。 :::
	OverwriteObject *string `json:"OverwriteObject,omitempty"`

	// 使用 veImageX 存储截图时，对应的 veImageX 的服务 ID。 :::tip 使用 veImageX 存储时 ServiceID 为必填项。 :::
	ServiceID *string `json:"ServiceID,omitempty"`
}

// CreateSnapshotPresetV2BodySnapshotPresetConfigJPEGParamTOSParam - 截图存储到 TOS 时的配置。 :::tip TOSParam 和 ImageXParam 配置且配置其中一个。
// :::
type CreateSnapshotPresetV2BodySnapshotPresetConfigJPEGParamTOSParam struct {

	// TOS 存储对应的 Bucket。 例如，存储路径为 live-test-tos-example/live/liveapp 时，Bucket 取值为 live-test-tos-example。 :::tip 使用 TOS 存储时 Bucket
	// 为必填项。 :::
	Bucket *string `json:"Bucket,omitempty"`

	// 截图是否使用 TOS 存储，默认为 false，取值及含义如下所示。
	// * false：不使用；
	// * true：使用。
	Enable *bool `json:"Enable,omitempty"`

	// 存储方式为实时截图时的存储规则，支持以 {Domain}/{App}/{Stream}/{UnixTimestamp} 样式设置存储规则，支持输入字母、数字、-、!、_、.、* 及占位符。 :::tip 参数 ExactObject 和
	// OverwriteObject 传且仅传一个。 :::
	ExactObject *string `json:"ExactObject,omitempty"`

	// 存储方式为覆盖截图时的存储规则，支持以 {Domain}/{App}/{Stream} 样式设置存储规则，支持输入字母、数字、-、!、_、.、* 及占位符。 :::tip 参数 ExactObject 和 OverwriteObject
	// 传且仅传一个。 :::
	OverwriteObject *string `json:"OverwriteObject,omitempty"`

	// ToS 存储对应的 bucket 下的存储目录，默认为空。 例如，存储位置为 live-test-tos-example/live/liveapp 时，StorageDir 取值为 live/liveapp。
	StorageDir *string `json:"StorageDir,omitempty"`
}

// CreateSnapshotPresetV2BodySnapshotPresetConfigJpgParam - 截图格式为 JPG 时的截图参数，开启 JPG 截图时设置。 :::tip JPEG 截图和 JPG 截图必须开启且只能开启一个。
// :::
type CreateSnapshotPresetV2BodySnapshotPresetConfigJpgParam struct {
	Enable      *bool                                                              `json:"Enable,omitempty"`
	ImageXParam *CreateSnapshotPresetV2BodySnapshotPresetConfigJpgParamImageXParam `json:"ImageXParam,omitempty"`
	TOSParam    *CreateSnapshotPresetV2BodySnapshotPresetConfigJpgParamTOSParam    `json:"TOSParam,omitempty"`
}

type CreateSnapshotPresetV2BodySnapshotPresetConfigJpgParamImageXParam struct {
	Enable          *bool   `json:"Enable,omitempty"`
	ExactObject     *string `json:"ExactObject,omitempty"`
	OverwriteObject *string `json:"OverwriteObject,omitempty"`
	ServiceID       *string `json:"ServiceID,omitempty"`
}

type CreateSnapshotPresetV2BodySnapshotPresetConfigJpgParamTOSParam struct {
	ACL             *string `json:"ACL,omitempty"`
	AccessKey       *string `json:"AccessKey,omitempty"`
	Bucket          *string `json:"Bucket,omitempty"`
	Enable          *bool   `json:"Enable,omitempty"`
	ExactObject     *string `json:"ExactObject,omitempty"`
	OverwriteObject *string `json:"OverwriteObject,omitempty"`
	Region          *string `json:"Region,omitempty"`
	S3NetworkType   *string `json:"S3NetworkType,omitempty"`
	StorageDir      *string `json:"StorageDir,omitempty"`
	TosCluster      *string `json:"TosCluster,omitempty"`
	TosDC           *string `json:"TosDC,omitempty"`
	TosPSM          *string `json:"TosPSM,omitempty"`
}

type CreateSnapshotPresetV2Res struct {

	// REQUIRED
	ResponseMetadata CreateSnapshotPresetV2ResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type CreateSnapshotPresetV2ResResponseMetadata struct {

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

type CreateTicketBody struct {

	// REQUIRED; 配置块信息
	ConfigList []CreateTicketBodyConfigListItem `json:"ConfigList"`

	// REQUIRED; 灰度分组参数
	GroupParam CreateTicketBodyGroupParam `json:"GroupParam"`

	// REQUIRED; 命名空间
	Namespace string `json:"Namespace"`

	// REQUIRED; 服务类型
	ServiceType string `json:"ServiceType"`

	// REQUIRED
	TicketType int32 `json:"TicketType"`

	// REQUIRED
	WorkFlowParams CreateTicketBodyWorkFlowParams `json:"WorkFlowParams"`
}

type CreateTicketBodyConfigListItem struct {

	// REQUIRED; 配置块名称
	TemplateName string `json:"TemplateName"`

	// REQUIRED; 域名空间
	Vhost string `json:"Vhost"`

	// 应用名称
	App *string `json:"App,omitempty"`

	// 域名
	Domain *string `json:"Domain,omitempty"`
}

// CreateTicketBodyGroupParam - 灰度分组参数
type CreateTicketBodyGroupParam struct {

	// REQUIRED; 分组数
	GroupNum int32 `json:"GroupNum"`

	// REQUIRED
	GroupStrategy int32 `json:"GroupStrategy"`
}

type CreateTicketBodyWorkFlowParams struct {

	// REQUIRED
	WaitTime int32 `json:"WaitTime"`
}

type CreateTicketRes struct {

	// REQUIRED
	ResponseMetadata CreateTicketResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type CreateTicketResResponseMetadata struct {

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

type CreateTimeShiftPresetV2Body struct {

	// REQUIRED
	MaxShiftTime int32 `json:"MaxShiftTime"`

	// REQUIRED
	PullDomain string `json:"PullDomain"`

	// REQUIRED
	Vhost        string  `json:"Vhost"`
	App          *string `json:"App,omitempty"`
	Bucket       *string `json:"Bucket,omitempty"`
	MasterFormat *string `json:"MasterFormat,omitempty"`
	Type         *string `json:"Type,omitempty"`
	VODNamespace *string `json:"VODNamespace,omitempty"`
}

type CreateTimeShiftPresetV2Res struct {

	// REQUIRED
	ResponseMetadata CreateTimeShiftPresetV2ResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type CreateTimeShiftPresetV2ResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version   string                                           `json:"Version"`
	Error     *CreateTimeShiftPresetV2ResResponseMetadataError `json:"Error,omitempty"`
	RequestID *string                                          `json:"RequestID,omitempty"`
}

type CreateTimeShiftPresetV2ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type CreateTimeShiftPresetV3Body struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 最大时移时长，即观看时移的最长时间，单位为 s。支持的取值如下所示。
	// * 86400
	// * 259200
	// * 604800
	// * 1296000
	MaxShiftTime int32 `json:"MaxShiftTime"`

	// REQUIRED; 时移拉流域名
	PullDomain string `json:"PullDomain"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 用于多码率时移的参数，为json字符串
	MasterFormat *string `json:"MasterFormat,omitempty"`

	// 0表示不需要 1表示需要
	NeedTranscode *int32 `json:"NeedTranscode,omitempty"`

	// 开启时移的流名称，同一个 App 最多可指定 20 路。
	Stream *string `json:"Stream,omitempty"`

	// 时移类型。支持的取值如下所示。
	// * 0：录制时移，即时移复用录制模板；
	// * 1：独立时移，即时移不复用录制模板。
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

type CreateTranscodePresetBatchBody struct {

	// REQUIRED
	PresetList []CreateTranscodePresetBatchBodyPresetListItem `json:"PresetList"`

	// REQUIRED; create associate hls-abr
	Type string `json:"Type"`
}

type CreateTranscodePresetBatchBodyPresetListItem struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 转码流后缀名。支持 10 个字符以内的大小写字母、下划线与中划线，常见后缀包括：sd、hd、_uhd 例如，配置的转码流后缀名为 _hd，则拉转码流时转码的流名为 stream-123456789_hd。
	SuffixName string `json:"SuffixName"`

	// REQUIRED; 视频编码格式。
	// * h264：使用 H.264 编码格式；
	// * h265：使用 H.265 编码格式；
	// * copy：不进行转码，所有视频编码参数不生效。
	Vcodec string `json:"Vcodec"`

	// REQUIRED; 域名空间名称。
	Vhost     string  `json:"Vhost"`
	ALayout   *string `json:"ALayout,omitempty"`
	AProfile  *string `json:"AProfile,omitempty"`
	AR        *int32  `json:"AR,omitempty"`
	AbrMode   *int32  `json:"AbrMode,omitempty"`
	AccountID *string `json:"AccountID,omitempty"`

	// 音频编码格式。默认格式为 acc，支持以下 3 种类型。
	// * aac：使用 aac 编码格式；
	// * copy：不进行转码，所有音频编码参数不生效；
	// * opus：使用 opus 编码格式。
	Acodec         *string `json:"Acodec,omitempty"`
	AdvancedParam  *string `json:"AdvancedParam,omitempty"`
	AllowAudioCopy *int32  `json:"AllowAudioCopy,omitempty"`
	AllowVideoCopy *int32  `json:"AllowVideoCopy,omitempty"`
	An             *int32  `json:"An,omitempty"`

	// 宽高自适应模式开关。默认值为 0。支持的取值包括。
	// * 0：关闭宽高自适应，按照 Width 和 Height 的取值进行拉伸；
	// * 1：开启宽高自适应，按照 ShortSide 或 LongSide 等比缩放。
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

	// 2 个参考帧之间的最大 B 帧数。不同编码格式的取值存在差异。
	// * H.264：取值范围为 [0,7]，默认值为 3；
	// * H.265：取值范围为 [0,1,2,3,7,15]，默认值为 3；
	// BFrames 取 0 时，表示去 B 帧。
	BFrames  *int32  `json:"BFrames,omitempty"`
	Describe *string `json:"Describe,omitempty"`

	// 视频帧率，单位为 fps，取值范围为 [0,60]，默认为 25fps。帧率越大，画面越流畅。
	FPS *int32 `json:"FPS,omitempty"`

	// IDR 帧之间的最大间隔，单位为 s，默认值为 0，取值范围为 [0,1000]
	GOP    *int32 `json:"GOP,omitempty"`
	GopMin *int32 `json:"GopMin,omitempty"`
	HVSPre *bool  `json:"HVSPre,omitempty"`

	// 视频高度。默认值为 0，取值范围为 [0,8192]。 :::tip
	// * 当 As 的取值为 0 时，参数生效；反之则不生效；
	// * 当 As 的取值为 0 时，如果 Width 和 Height 任意取值为 0，表示保持源流尺寸。 :::
	Height *int32 `json:"Height,omitempty"`

	// 长边长度。默认值为 0。
	// * Roi 取 false 时，取值范围为 [0,8192]；
	// * Roi 取 true 时，取值范围为 [0,1920]。 :::tip
	// * 当 As 的取值为 1 时，参数生效，反之则不生效。
	// * 当 As 的取值为 1 时，如果 LongSide 和 ShortSide 同时取 0，表示保持源流尺寸。
	// * 当 As 的取值为 1 时，如果同时配置 LongSide 和 ShortSide 的值，则按照 ShortSide 进行等比缩放。 :::
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

	// 是否极智超清转码。默认值为 false。
	// * true：极智超清；
	// * false：标准转码。
	Roi  *bool `json:"Roi,omitempty"`
	SITI *bool `json:"SITI,omitempty"`

	// 短边长度。默认值为 0。
	// * Roi 取 false 时，取值范围为 [0,4096]；
	// * Roi 取 true 时，取值范围为 [0,1080]。 :::tip
	// * 当 As 的取值为 1 时，参数生效，反之则不生效。
	// * 当 As 的取值为 1 时，如果 LongSide 和 ShortSide 同时取 0，表示保持源流尺寸。
	// * 当 As 的取值为 1 时，如果同时配置 LongSide 和 ShortSide 的值，则按照 ShortSide 进行等比缩放。 :::
	ShortSide       *int32                                                       `json:"ShortSide,omitempty"`
	Status          *int32                                                       `json:"Status,omitempty"`
	StopInterval    *int32                                                       `json:"StopInterval,omitempty"`
	Threads         *int32                                                       `json:"Threads,omitempty"`
	TranscodeStruct *CreateTranscodePresetBatchBodyPresetListItemTranscodeStruct `json:"TranscodeStruct,omitempty"`
	VBRatio         *int32                                                       `json:"VBRatio,omitempty"`
	VBVBufSize      *int32                                                       `json:"VBVBufSize,omitempty"`
	VBVMaxRate      *int32                                                       `json:"VBVMaxRate,omitempty"`
	VLevel          *string                                                      `json:"VLevel,omitempty"`
	VPreset         *string                                                      `json:"VPreset,omitempty"`
	VProfile        *string                                                      `json:"VProfile,omitempty"`
	VR              *int32                                                       `json:"VRVr,omitempty"`
	VRBBframes      *int32                                                       `json:"VRBBframes,omitempty"`
	VRBHeightNum    *int32                                                       `json:"VRBHeightNum,omitempty"`
	VRBPreset       *string                                                      `json:"VRBPreset,omitempty"`
	VRBProfile      *string                                                      `json:"VRBProfile,omitempty"`
	VRBSuffix       *string                                                      `json:"VRBSuffix,omitempty"`
	VRBVb           *int32                                                       `json:"VRBVb,omitempty"`
	VRBWidthNum     *int32                                                       `json:"VRBWidthNum,omitempty"`
	VRGop           *int32                                                       `json:"VRGop,omitempty"`
	VRGopDen        *int32                                                       `json:"VRGopDen,omitempty"`
	VRHvspre        *string                                                      `json:"VRHvspre,omitempty"`
	VRProjection    *string                                                      `json:"VRProjection,omitempty"`
	VRRoi           *string                                                      `json:"VRRoi,omitempty"`
	VRTBframes      *int32                                                       `json:"VRTBframes,omitempty"`
	VRTPreset       *string                                                      `json:"VRTPreset,omitempty"`
	VRTProfile      *string                                                      `json:"VRTProfile,omitempty"`
	VRTSuffix       *string                                                      `json:"VRTSuffix,omitempty"`
	VRTVb           *int32                                                       `json:"VRTVb,omitempty"`
	VRTileMod       *int32                                                       `json:"VRTileMod,omitempty"`
	VRateCtrl       *string                                                      `json:"VRateCtrl,omitempty"`
	VbThreshold     *string                                                      `json:"VbThreshold,omitempty"`
	Vclass          *bool                                                        `json:"Vclass,omitempty"`

	// 视频码率，单位为 bps，取值范围为 [0,30000000]；默认值为 1000000；取 0 时，表示使用源流码率。
	VideoBitrate *int32  `json:"VideoBitrate,omitempty"`
	Vn           *int32  `json:"Vn,omitempty"`
	Watermark    *string `json:"Watermark,omitempty"`

	// 视频宽度。默认值为 0，取值范围为 [0,8192]。 :::tip
	// * 当 As 的取值为 0 时，参数生效；反之则不生效；
	// * 当 As 的取值为 0 时，如果 Width 和 Height 任意取值为 0，表示保持源流尺寸。 :::
	Width *int32 `json:"Width,omitempty"`
}

type CreateTranscodePresetBatchBodyPresetListItemTranscodeStruct struct {

	// REQUIRED; 标记模版是否为hlsabr
	IsHlsAbr bool `json:"IsHlsAbr"`

	// Dictionary of
	ABTest       map[string]*ComponentsSgrw9KSchemasCreatetranscodepresetbatchbodyPropertiesPresetlistItemsPropertiesTranscodestructPropertiesAbtestAdditionalproperties `json:"ABTest,omitempty"`
	Codec        *string                                                                                                                                                 `json:"Codec,omitempty"`
	PresetName   *string                                                                                                                                                 `json:"PresetName,omitempty"`
	StopInterval *int32                                                                                                                                                  `json:"StopInterval,omitempty"`
	Suffix       *string                                                                                                                                                 `json:"Suffix,omitempty"`
	Type         *string                                                                                                                                                 `json:"Type,omitempty"`
}

type CreateTranscodePresetBatchRes struct {

	// REQUIRED
	ResponseMetadata CreateTranscodePresetBatchResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type CreateTranscodePresetBatchResResponseMetadata struct {

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

type CreateTranscodePresetBody struct {

	// REQUIRED; 应用名称，取值与直播流地址的 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 转码后缀，支持由大小写字母（A - Z、a - z）、下划线（_）和短横线（-）组成，长度为 1 到 10 个字符。
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
	// [https://console-stable.volcanicengine.com/live/main/domain/list]页面，查看需要转码的直播流使用的域名所属的域名空间。
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
	// * 视频编码格式为 H.265 （Vcodec 取值为 h265）时取值范围为 [0,3]、7、15；
	// * 视频编码格式为 H.266 （Vcodec 取值为 h266）时取值范围为 [0,3]、7、15。
	BFrames  *int32  `json:"BFrames,omitempty"`
	Describe *string `json:"Describe,omitempty"`

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

type CreateTranscodePresetPatchByAdminBody struct {

	// REQUIRED; 模板配置列表
	PresetList []CreateTranscodePresetPatchByAdminBodyPresetListItem `json:"PresetList"`

	// REQUIRED; 逻辑的的处理类型，create：单纯创建模板，不关联app/vhost，associate：创建模板的同时关联app/vhost
	Type string `json:"Type"`

	// 配置的类型，不填默认为全量配置
	// * simple：精简配置，后端会默认填充字段
	// * full：全量配置，后端不会做修改
	ConfigType *string `json:"ConfigType,omitempty"`
}

type CreateTranscodePresetPatchByAdminBodyPresetListItem struct {

	// REQUIRED; 转码流后缀名。支持 10 个字符以内的大小写字母、下划线与中划线，常见后缀包括：sd、hd、_uhd 例如，配置的转码流后缀名为 _hd，则拉转码流时转码的流名为 stream-123456789_hd。
	SuffixName string `json:"SuffixName"`

	// REQUIRED; 视频编码格式。
	// * h264：使用 H.264 编码格式；
	// * h265：使用 H.265 编码格式；
	// * copy：不进行转码，所有视频编码参数不生效。
	Vcodec    string  `json:"Vcodec"`
	ALayout   *string `json:"ALayout,omitempty"`
	AProfile  *string `json:"AProfile,omitempty"`
	AR        *int32  `json:"AR,omitempty"`
	AbrMode   *int32  `json:"AbrMode,omitempty"`
	AccountID *string `json:"AccountID,omitempty"`

	// 音频编码格式。默认格式为 acc，支持以下 3 种类型。
	// * aac：使用 aac 编码格式；
	// * copy：不进行转码，所有音频编码参数不生效；
	// * opus：使用 opus 编码格式。
	Acodec         *string `json:"Acodec,omitempty"`
	AdvancedParam  *string `json:"AdvancedParam,omitempty"`
	AllowAudioCopy *int32  `json:"AllowAudioCopy,omitempty"`
	AllowVideoCopy *int32  `json:"AllowVideoCopy,omitempty"`
	An             *int32  `json:"An,omitempty"`

	// 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App *string `json:"App,omitempty"`

	// 宽高自适应模式开关。默认值为 0。支持的取值包括。
	// * 0：关闭宽高自适应，按照 Width 和 Height 的取值进行拉伸；
	// * 1：开启宽高自适应，按照 ShortSide 或 LongSide 等比缩放。
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

	// 2 个参考帧之间的最大 B 帧数。不同编码格式的取值存在差异。
	// * H.264：取值范围为 [0,7]，默认值为 3；
	// * H.265：取值范围为 [0,1,2,3,7,15]，默认值为 3；
	// BFrames 取 0 时，表示去 B 帧。
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

	// 视频帧率，单位为 fps，取值范围为 [0,60]，默认为 25fps。帧率越大，画面越流畅。
	FPS *int32 `json:"FPS,omitempty"`

	// IDR 帧之间的最大间隔，单位为 s，默认值为 0，取值范围为 [0,1000]
	GOP    *int32 `json:"GOP,omitempty"`
	GopMin *int32 `json:"GopMin,omitempty"`
	HVSPre *bool  `json:"HVSPre,omitempty"`

	// 视频高度。默认值为 0，取值范围为 [0,8192]。 :::tip
	// * 当 As 的取值为 0 时，参数生效；反之则不生效；
	// * 当 As 的取值为 0 时，如果 Width 和 Height 任意取值为 0，表示保持源流尺寸。 :::
	Height *int32 `json:"Height,omitempty"`

	// 长边长度。默认值为 0。
	// * Roi 取 false 时，取值范围为 [0,8192]；
	// * Roi 取 true 时，取值范围为 [0,1920]。 :::tip
	// * 当 As 的取值为 1 时，参数生效，反之则不生效。
	// * 当 As 的取值为 1 时，如果 LongSide 和 ShortSide 同时取 0，表示保持源流尺寸。
	// * 当 As 的取值为 1 时，如果同时配置 LongSide 和 ShortSide 的值，则按照 ShortSide 进行等比缩放。 :::
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
	ParamType    *string `json:"ParamType,omitempty"`
	Preset       *string `json:"Preset,omitempty"`
	PresetKind   *int32  `json:"PresetKind,omitempty"`
	PresetType   *int32  `json:"PresetType,omitempty"`
	Qp           *int32  `json:"Qp,omitempty"`
	RegionConfig *string `json:"RegionConfig,omitempty"`
	Revision     *string `json:"Revision,omitempty"`

	// 是否极智超清转码。默认值为 false。
	// * true：极智超清；
	// * false：标准转码。
	Roi  *bool `json:"Roi,omitempty"`
	SITI *bool `json:"SITI,omitempty"`

	// 使用场景，画质增强时生效
	// * football：足球场景
	SceneType *string `json:"SceneType,omitempty"`

	// 短边长度。默认值为 0。
	// * Roi 取 false 时，取值范围为 [0,4096]；
	// * Roi 取 true 时，取值范围为 [0,1080]。 :::tip
	// * 当 As 的取值为 1 时，参数生效，反之则不生效。
	// * 当 As 的取值为 1 时，如果 LongSide 和 ShortSide 同时取 0，表示保持源流尺寸。
	// * 当 As 的取值为 1 时，如果同时配置 LongSide 和 ShortSide 的值，则按照 ShortSide 进行等比缩放。 :::
	ShortSide    *int32 `json:"ShortSide,omitempty"`
	Status       *int32 `json:"Status,omitempty"`
	StopInterval *int32 `json:"StopInterval,omitempty"`
	Threads      *int32 `json:"Threads,omitempty"`

	// 转码触发方式，默认为拉流转码，支持以下取值(给火山控制台使用)。
	// * Push：推流转码，直播推流后会自动启动转码任务，生成转码流；
	// * Pull：拉流转码，直播推流后，需要主动播放转码流才会启动转码任务，生成转码流。
	TransType       *string                                                             `json:"TransType,omitempty"`
	TranscodeStruct *CreateTranscodePresetPatchByAdminBodyPresetListItemTranscodeStruct `json:"TranscodeStruct,omitempty"`
	VBRatio         *int32                                                              `json:"VBRatio,omitempty"`
	VBVBufSize      *int32                                                              `json:"VBVBufSize,omitempty"`
	VBVMaxRate      *int32                                                              `json:"VBVMaxRate,omitempty"`
	VLevel          *string                                                             `json:"VLevel,omitempty"`
	VPreset         *string                                                             `json:"VPreset,omitempty"`
	VProfile        *string                                                             `json:"VProfile,omitempty"`
	VR              *int32                                                              `json:"VRVr,omitempty"`
	VRBBframes      *int32                                                              `json:"VRBBframes,omitempty"`
	VRBHeightNum    *int32                                                              `json:"VRBHeightNum,omitempty"`
	VRBPreset       *string                                                             `json:"VRBPreset,omitempty"`
	VRBProfile      *string                                                             `json:"VRBProfile,omitempty"`
	VRBSuffix       *string                                                             `json:"VRBSuffix,omitempty"`
	VRBVb           *int32                                                              `json:"VRBVb,omitempty"`
	VRBWidthNum     *int32                                                              `json:"VRBWidthNum,omitempty"`
	VRGop           *int32                                                              `json:"VRGop,omitempty"`
	VRGopDen        *int32                                                              `json:"VRGopDen,omitempty"`
	VRHvspre        *string                                                             `json:"VRHvspre,omitempty"`
	VRProjection    *string                                                             `json:"VRProjection,omitempty"`
	VRRoi           *string                                                             `json:"VRRoi,omitempty"`
	VRTBframes      *int32                                                              `json:"VRTBframes,omitempty"`
	VRTPreset       *string                                                             `json:"VRTPreset,omitempty"`
	VRTProfile      *string                                                             `json:"VRTProfile,omitempty"`
	VRTSuffix       *string                                                             `json:"VRTSuffix,omitempty"`
	VRTVb           *int32                                                              `json:"VRTVb,omitempty"`
	VRTileMod       *int32                                                              `json:"VRTileMod,omitempty"`
	VRateCtrl       *string                                                             `json:"VRateCtrl,omitempty"`
	VbThreshold     *string                                                             `json:"VbThreshold,omitempty"`
	Vclass          *bool                                                               `json:"Vclass,omitempty"`

	// 域名空间名称。
	Vhost *string `json:"Vhost,omitempty"`

	// 视频码率，单位为 bps，取值范围为 [0,30000000]；默认值为 1000000；取 0 时，表示使用源流码率。
	VideoBitrate *int32  `json:"VideoBitrate,omitempty"`
	Vn           *int32  `json:"Vn,omitempty"`
	Watermark    *string `json:"Watermark,omitempty"`

	// 视频宽度。默认值为 0，取值范围为 [0,8192]。 :::tip
	// * 当 As 的取值为 0 时，参数生效；反之则不生效；
	// * 当 As 的取值为 0 时，如果 Width 和 Height 任意取值为 0，表示保持源流尺寸。 :::
	Width *int32 `json:"Width,omitempty"`
}

type CreateTranscodePresetPatchByAdminBodyPresetListItemTranscodeStruct struct {

	// Dictionary of
	ABTest map[string]*Components4Gwy08SchemasCreatetranscodepresetpatchbyadminbodyPropertiesPresetlistItemsPropertiesTranscodestructPropertiesAbtestAdditionalproperties `json:"ABTest,omitempty"`
	Codec  *string                                                                                                                                                        `json:"Codec,omitempty"`

	// 模板名称
	PresetName *string `json:"PresetName,omitempty"`

	// 拉流转码多少秒停止
	StopInterval *int32 `json:"StopInterval,omitempty"`

	// 转码后缀
	Suffix *string `json:"Suffix,omitempty"`

	// 转码触发类型
	Type *string `json:"Type,omitempty"`
}

type CreateTranscodePresetPatchByAdminRes struct {

	// REQUIRED
	ResponseMetadata CreateTranscodePresetPatchByAdminResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type CreateTranscodePresetPatchByAdminResResponseMetadata struct {

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

type CreateVQScoreTaskBody struct {

	// 测评算法，支持vqscore分数计算
	Algorithm *string `json:"Algorithm,omitempty"`

	// 对比拉流地址。
	ContrastAddr *string `json:"ContrastAddr,omitempty"`

	// 测评运行时间
	// * 支持输入s整数
	// * 最大支持7*24小时的测评任务
	// * 最小支持1min的测评任务
	Duration *string `json:"Duration,omitempty"`

	// 抽帧间隔，目前只能密集抽帧模板控制，以模板为主。
	FrameInterval *string `json:"FrameInterval,omitempty"`

	// 主拉流地址。 支持输入FCDN拉流地址和第三方CDN拉流地址。
	MainAddr *string `json:"MainAddr,omitempty"`
}

type CreateVQScoreTaskRes struct {

	// REQUIRED
	ResponseMetadata CreateVQScoreTaskResResponseMetadata `json:"ResponseMetadata"`
	Result           *CreateVQScoreTaskResResult          `json:"Result,omitempty"`
}

type CreateVQScoreTaskResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RequestID string `json:"RequestId"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version string                                     `json:"Version"`
	Error   *CreateVQScoreTaskResResponseMetadataError `json:"Error,omitempty"`
}

type CreateVQScoreTaskResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type CreateVQScoreTaskResResult struct {
	ID *string `json:"ID,omitempty"`
}

type CreateVerifyContentBody struct {

	// REQUIRED; 推拉流域名
	Domain string `json:"Domain"`
}

type CreateVerifyContentRes struct {

	// REQUIRED
	ResponseMetadata CreateVerifyContentResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *CreateVerifyContentResResult `json:"Result,omitempty"`
}

type CreateVerifyContentResResponseMetadata struct {

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

// CreateVerifyContentResResult - 视请求的接口而定
type CreateVerifyContentResResult struct {

	// 校验内容记录值
	Content *string `json:"Content,omitempty"`

	// 主机记录
	SubDomain *string `json:"SubDomain,omitempty"`
}

type CreateWatermarkPresetBody struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 水平偏移，表示水印左侧边与转码流画面左侧边之间的距离，使用相对比率，取值范围为 [0,1]。
	PosX float32 `json:"PosX"`

	// REQUIRED; 垂直偏移，表示水印顶部边与转码流画面顶部边之间的距离，使用相对比率，取值范围为 [0,1]。
	PosY float32 `json:"PosY"`

	// REQUIRED; 域名空间名称，由 1 到 60 位数字、字母、下划线及"-"和"."组成。
	Vhost string `json:"Vhost"`

	// 需要添加水印的直播画面方向，支持 2 种取值。
	// * vertical：竖屏；
	// * horizontal：横屏。 :::tip 该参数属于历史版本参数，预计将于未来移除。建议使用预览背景高度（PreviewHeight）、预览背景宽度（PreviewWidth）参数代替。 :::
	Orientation *string `json:"Orientation,omitempty"`

	// 水印图片字符串，图片最大 2MB，最小 100Bytes，最大分辨率为 1080×1080。图片 Data URL 格式为：data:[<mediatype>];[base64],<data>。
	// * mediatype：图片类型，支持 png、jpg、jpeg 格式；
	// * data：base64 编码的图片字符串。
	Picture *string `json:"Picture,omitempty"`

	// 水印图片对应的 HTTP 地址。与水印图片字符串字段二选一传入，同时传入时，以水印图片字符串参数为准。
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

type CreateWatermarkPresetV2Body struct {

	// REQUIRED; 水印图片字符串，图片最大 2MB，最小 100Bytes，最大分辨率为 1080×1080。图片 Data URL 格式为：data:[<mediatype>];[base64],<data>。
	// * mediatype：图片类型，支持 png、jpg、jpeg 格式；
	// * data：base64 编码的图片字符串。
	// 例如，data:image/png;base64,iVBORw0KGg****mCC
	Picture string `json:"Picture"`

	// 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App *string `json:"App,omitempty"`

	// 需要添加水印的直播画面方向，支持 2 种取值。
	// * vertical：竖屏；
	// * horizontal：横屏。
	Orientation *string `json:"Orientation,omitempty"`
	PictureURL  *string `json:"PictureUrl,omitempty"`

	// 水平偏移，表示水印左侧边与转码流画面左侧边之间的距离，使用相对比率，取值范围为 [0,1]。
	PosX *float32 `json:"PosX,omitempty"`

	// 垂直偏移，表示水印顶部边与转码流画面顶部边之间的距离，使用相对比率，取值范围为 [0,1]。
	PosY           *float32 `json:"PosY,omitempty"`
	PresetName     *string  `json:"PresetName,omitempty"`
	PreviewHeight  *float32 `json:"PreviewHeight,omitempty"`
	PreviewWidth   *float32 `json:"PreviewWidth,omitempty"`
	RelativeHeight *float32 `json:"RelativeHeight,omitempty"`

	// 水印相对宽度，水印宽度占直播转码流画面宽度的比例，取值范围为 [0,1]，水印高度会随宽度等比缩放。
	RelativeWidth *float32 `json:"RelativeWidth,omitempty"`
	Scale         *float32 `json:"Scale,omitempty"`
	Stream        *string  `json:"Stream,omitempty"`

	// 域名空间名称。由 1 到 60 位数字、字母、下划线及"-"和"."组成。
	Vhost *string `json:"Vhost,omitempty"`
}

type CreateWatermarkPresetV2Res struct {

	// REQUIRED
	ResponseMetadata CreateWatermarkPresetV2ResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *CreateWatermarkPresetV2ResResult `json:"Result,omitempty"`
}

type CreateWatermarkPresetV2ResResponseMetadata struct {

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

// CreateWatermarkPresetV2ResResult - 视请求的接口而定
type CreateWatermarkPresetV2ResResult struct {

	// REQUIRED; 模板的ID
	ID int32 `json:"ID"`

	// REQUIRED; 模板名称
	PresetName string `json:"PresetName"`
}

type DeleteAuthBody struct {

	// REQUIRED; 鉴权场景，枚举：push， pull， bypass
	SceneType string `json:"SceneType"`

	// 应用名称
	App *string `json:"App,omitempty"`

	// 域名，vhost和domain二选一必填
	Domain *string `json:"Domain,omitempty"`

	// 域名空间，vhost和domain二选一必填
	Vhost *string `json:"Vhost,omitempty"`
}

type DeleteAuthRes struct {

	// REQUIRED
	ResponseMetadata DeleteAuthResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteAuthResResponseMetadata struct {

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

type DeleteAvSlicePresetBody struct {

	// REQUIRED
	Preset string  `json:"Preset"`
	App    *string `json:"App,omitempty"`
	Stream *string `json:"Stream,omitempty"`
	Vhost  *string `json:"Vhost,omitempty"`
}

type DeleteAvSlicePresetRes struct {

	// REQUIRED
	ResponseMetadata DeleteAvSlicePresetResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteAvSlicePresetResResponseMetadata struct {

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

	// 应用名称。缺省情况下表示删除 Vhost 下的所有回调配置。如果入参选择 Domain，则不可同时传 App。
	App *string `json:"App,omitempty"`

	// 推流域名。如创建回调 UpdateCallback [https://www.volcengine.com/docs/6469/78553] 时传了参数 Domain，删除时需要传 Domain。
	Domain *string `json:"Domain,omitempty"`

	// 消息类型。缺省情况下表示删除所有消息类型。包括以下类型。
	// * push：推流开始回调；
	// * push_end：推流结束回调；
	// * snapshot：截图回调；
	// * record：录制回调；
	// * audit_snapshot：截图审核回调。
	MessageType *string `json:"MessageType,omitempty"`

	// 域名空间名称。如创建回调 UpdateCallback [https://www.volcengine.com/docs/6469/78553] 时传了参数 Vhost，删除时需要传 Vhost。
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

	// REQUIRED; 待删除的证书链 ID。
	ChainID   string  `json:"ChainID"`
	AccountID *string `json:"AccountID,omitempty"`
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

type DeleteCommonTransPresetBody struct {

	// REQUIRED
	App string `json:"App"`

	// REQUIRED
	Preset string `json:"Preset"`

	// REQUIRED
	Vhost string `json:"Vhost"`
}

type DeleteCommonTransPresetRes struct {

	// REQUIRED
	ResponseMetadata DeleteCommonTransPresetResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type DeleteCommonTransPresetResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version   string                                           `json:"Version"`
	Error     *DeleteCommonTransPresetResResponseMetadataError `json:"Error,omitempty"`
	RequestID *string                                          `json:"RequestID,omitempty"`
}

type DeleteCommonTransPresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DeleteCustomLogConfigBody struct {

	// REQUIRED; 删除的配置Id
	ID string `json:"Id"`
}

type DeleteCustomLogConfigRes struct {

	// REQUIRED
	ResponseMetadata DeleteCustomLogConfigResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DeleteCustomLogConfigResResult `json:"Result,omitempty"`
}

type DeleteCustomLogConfigResResponseMetadata struct {

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

// DeleteCustomLogConfigResResult - 视请求的接口而定
type DeleteCustomLogConfigResResult struct {

	// REQUIRED; 删除的配置Id
	ID string `json:"Id"`
}

type DeleteDenseSnapshotPresetBody struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 密集抽帧截图配置模板名称。
	Preset string `json:"Preset"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`
}

type DeleteDenseSnapshotPresetRes struct {

	// REQUIRED
	ResponseMetadata DeleteDenseSnapshotPresetResResponseMetadata `json:"ResponseMetadata"`
}

type DeleteDenseSnapshotPresetResResponseMetadata struct {

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

type DeleteDenyConfigV2Body struct {

	// REQUIRED; 域名空间名称
	Vhost string `json:"Vhost"`

	// 推/拉流域名
	Domain *string `json:"Domain,omitempty"`
}

type DeleteDenyConfigV2Res struct {

	// REQUIRED
	ResponseMetadata DeleteDenyConfigV2ResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteDenyConfigV2ResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                      `json:"Version"`
	Error   *DeleteDenyConfigV2ResResponseMetadataError `json:"Error,omitempty"`
}

type DeleteDenyConfigV2ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DeleteDomainBody struct {

	// REQUIRED; 待删除域名。
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

type DeleteDomainV2Body struct {

	// REQUIRED; 域名列表
	Domains []string `json:"Domains"`
}

type DeleteDomainV2Res struct {

	// REQUIRED
	ResponseMetadata DeleteDomainV2ResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteDomainV2ResResponseMetadata struct {

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

	// REQUIRED; The type of HTTP header configurations you want to delete:
	// * 0: Response headers.
	// * 1: Request headers.
	Phase int32 `json:"Phase"`

	// REQUIRED; The domain name space.
	Vhost string `json:"Vhost"`

	// The domain name.
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

type DeleteHeaderConfigBody struct {

	// REQUIRED; 域名空间
	Vhost string `json:"Vhost"`

	// 应用名称
	App *string `json:"App,omitempty"`

	// header类型，支持hls,flv,dash
	HeaderType *string `json:"HeaderType,omitempty"`
}

type DeleteHeaderConfigRes struct {

	// REQUIRED
	ResponseMetadata DeleteHeaderConfigResResponseMetadata `json:"ResponseMetadata"`
}

type DeleteHeaderConfigResResponseMetadata struct {

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

	// REQUIRED; 推/拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 域名空间名称。
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

type DeleteLiveAccountFeeConfigBody struct {

	// REQUIRED
	AccountID string `json:"AccountId"`

	// REQUIRED; 如果id是0，表示创建，否则表示更新
	ID int32 `json:"Id"`
}

type DeleteLiveAccountFeeConfigRes struct {

	// REQUIRED
	ResponseMetadata DeleteLiveAccountFeeConfigResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DeleteLiveAccountFeeConfigResResult `json:"Result,omitempty"`
}

type DeleteLiveAccountFeeConfigResResponseMetadata struct {

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

// DeleteLiveAccountFeeConfigResResult - 视请求的接口而定
type DeleteLiveAccountFeeConfigResResult struct {

	// 配置id
	ID *int32 `json:"Id,omitempty"`
}

type DeleteNSSRewriteConfigBody struct {

	// REQUIRED; 域名空间
	Vhost string `json:"Vhost"`

	// 应用名称
	App *string `json:"App,omitempty"`

	// 服务类型
	ServiceType *string `json:"ServiceType,omitempty"`
}

type DeleteNSSRewriteConfigRes struct {

	// REQUIRED
	ResponseMetadata DeleteNSSRewriteConfigResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type DeleteNSSRewriteConfigResResponseMetadata struct {

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

type DeleteProxyConfigAssociationBody struct {

	// REQUIRED; 账号
	AccountID string `json:"AccountID"`

	// REQUIRED; 代理记录ID
	ProxyID int32 `json:"ProxyID"`
}

type DeleteProxyConfigAssociationRes struct {

	// REQUIRED
	ResponseMetadata DeleteProxyConfigAssociationResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteProxyConfigAssociationResResponseMetadata struct {

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

type DeleteProxyConfigBody struct {

	// REQUIRED; 记录ID
	ID string `json:"ID"`
}

type DeleteProxyConfigRes struct {

	// REQUIRED
	ResponseMetadata DeleteProxyConfigResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteProxyConfigResResponseMetadata struct {

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

type DeleteRecordHistoryBody struct {

	// REQUIRED; 域名空间
	Vhost string `json:"Vhost"`
}

type DeleteRecordHistoryRes struct {

	// REQUIRED
	ResponseMetadata DeleteRecordHistoryResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteRecordHistoryResResponseMetadata struct {

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
	Error   *DeleteRecordHistoryResResponseMetadataError `json:"Error,omitempty"`
}

type DeleteRecordHistoryResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DeleteRecordPresetBody struct {

	// REQUIRED; 模版名称。可调用 ListVhostRecordPresetV2 [https://www.volcengine.com/docs/6469/1126858] 接口，查询模版名称。
	Preset string `json:"Preset"`

	// 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App *string `json:"App,omitempty"`

	// 域名空间名称。
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

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 应用名称
	// * 如创建时传了 App，删除时需要传该参数；
	// * 如创建时未传 App，删除时不传该参数。
	App *string `json:"App,omitempty"`

	// 拉流域名。
	// * 如创建时传了 Domain，删除时需要传该参数；
	// * 如创建时未传 Domain，删除时不传该参数。
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

type DeleteRelaySinkBody struct {

	// REQUIRED
	Vhost string  `json:"Vhost"`
	App   *string `json:"App,omitempty"`
}

type DeleteRelaySinkRes struct {

	// REQUIRED
	ResponseMetadata DeleteRelaySinkResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteRelaySinkResResponseMetadata struct {

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

	// The domain name
	Domain *string `json:"Domain,omitempty"`

	// The domain name space
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

	// REQUIRED; A list of rewrite rules deleted
	Items []interface{} `json:"Items"`
}

type DeleteRelaySourceV3Body struct {

	// REQUIRED; 直播流使用的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console-stable.volcanicengine.com/live/main/domain/list]页面，查看直播流使用的域名。所属的域名空间。
	Vhost string `json:"Vhost"`

	// 应用名称，即直播流地址的AppName字段取值，默认为空，表示删除当前域名空间的全局播放触发回源配置。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
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

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 推流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 流名称，由 1 到 100 位数字、字母、下划线及"-"和"."组成。
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

type DeleteSDKBody struct {

	// REQUIRED; 账号
	AccountID string `json:"AccountID"`

	// REQUIRED; 库记录ID
	ID int32 `json:"ID"`
}

type DeleteSDKRes struct {
	ResponseMetadata *DeleteSDKResResponseMetadata `json:"ResponseMetadata,omitempty"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type DeleteSDKResResponseMetadata struct {
	Action *string `json:"Action,omitempty"`

	// Anything
	Error     interface{} `json:"Error,omitempty"`
	Region    *string     `json:"Region,omitempty"`
	RequestID *string     `json:"RequestID,omitempty"`
	Service   *string     `json:"Service,omitempty"`
	Version   *string     `json:"Version,omitempty"`
}

type DeleteSnapshotAuditPresetBody struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 截图审核配置的名称。
	PresetName string `json:"PresetName"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`
}

type DeleteSnapshotAuditPresetRes struct {

	// REQUIRED
	ResponseMetadata DeleteSnapshotAuditPresetResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteSnapshotAuditPresetResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                             `json:"Version"`
	Error   *DeleteSnapshotAuditPresetResResponseMetadataError `json:"Error,omitempty"`
}

type DeleteSnapshotAuditPresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DeleteSnapshotPresetBody struct {

	// REQUIRED; 截图配置名称。
	Preset string `json:"Preset"`

	// 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App *string `json:"App,omitempty"`

	// 域名空间名称。
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

type DeleteStreamQuotaConfigBody struct {

	// REQUIRED; 待删除限额配置的推流域名或拉流域名。
	Domain string `json:"Domain"`
}

type DeleteStreamQuotaConfigRes struct {

	// REQUIRED
	ResponseMetadata DeleteStreamQuotaConfigResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type DeleteStreamQuotaConfigResResponseMetadata struct {

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
	Error   *DeleteStreamQuotaConfigResResponseMetadataError `json:"Error,omitempty"`
}

type DeleteStreamQuotaConfigResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DeleteTimeShiftPresetV2Body struct {

	// REQUIRED
	Preset string `json:"Preset"`

	// REQUIRED
	Vhost string  `json:"Vhost"`
	App   *string `json:"App,omitempty"`
	Type  *string `json:"Type,omitempty"`
}

type DeleteTimeShiftPresetV2Res struct {

	// REQUIRED
	ResponseMetadata DeleteTimeShiftPresetV2ResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type DeleteTimeShiftPresetV2ResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version   string                                           `json:"Version"`
	Error     *DeleteTimeShiftPresetV2ResResponseMetadataError `json:"Error,omitempty"`
	RequestID *string                                          `json:"RequestID,omitempty"`
}

type DeleteTimeShiftPresetV2ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DeleteTimeShiftPresetV3Body struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 域名空间名称。
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

type DeleteTranscodePresetBatchBody struct {

	// REQUIRED; 删除模版的信息
	PresetList []DeleteTranscodePresetBatchBodyPresetListItem `json:"PresetList"`

	// REQUIRED; associate create hls-abr
	Type string `json:"Type"`
}

type DeleteTranscodePresetBatchBodyPresetListItem struct {

	// REQUIRED; 所属accountid
	AccountID string `json:"AccountID"`

	// REQUIRED; 解绑的app
	App string `json:"App"`

	// REQUIRED; 模版名
	Preset string `json:"Preset"`

	// REQUIRED; 解绑的stream
	Stream string `json:"Stream"`

	// REQUIRED; 解绑的vhost
	Vhost string `json:"Vhost"`
}

type DeleteTranscodePresetBatchRes struct {

	// REQUIRED
	ResponseMetadata DeleteTranscodePresetBatchResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type DeleteTranscodePresetBatchResResponseMetadata struct {

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

type DeleteTranscodePresetBody struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 转码配置名称。
	Preset string `json:"Preset"`

	// REQUIRED; 域名空间名称。
	Vhost     string  `json:"Vhost"`
	AccountID *string `json:"AccountID,omitempty"`
}

type DeleteTranscodePresetPatchByAdminBody struct {

	// REQUIRED
	PresetList []DeleteTranscodePresetPatchByAdminBodyPresetListItem `json:"PresetList"`

	// REQUIRED; 操作的类型，associate: 删除模板的同时取消关联，create: 只删除模板
	Type string `json:"Type"`
}

type DeleteTranscodePresetPatchByAdminBodyPresetListItem struct {

	// REQUIRED
	Preset    string  `json:"Preset"`
	AccountID *string `json:"AccountID,omitempty"`
	App       *string `json:"App,omitempty"`
	Stream    *string `json:"Stream,omitempty"`
	Vhost     *string `json:"Vhost,omitempty"`
}

type DeleteTranscodePresetPatchByAdminRes struct {

	// REQUIRED
	ResponseMetadata DeleteTranscodePresetPatchByAdminResResponseMetadata `json:"ResponseMetadata"`
	Result           *DeleteTranscodePresetPatchByAdminResResult          `json:"Result,omitempty"`
}

type DeleteTranscodePresetPatchByAdminResResponseMetadata struct {

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

type DeleteTranscodePresetPatchByAdminResResult struct {

	// REQUIRED
	ModuleDeployTasks []interface{} `json:"ModuleDeployTasks"`
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

type DeleteWatermarkPresetBody struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 域名空间名称，由 1 到 60 位数字、字母、下划线及"-"和"."组成。
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

type DeleteWatermarkPresetV2Body struct {

	// 模板ID
	ID *int32 `json:"ID,omitempty"`

	// 模板名称
	PresetName *string `json:"PresetName,omitempty"`
}

type DeleteWatermarkPresetV2Res struct {

	// REQUIRED
	ResponseMetadata DeleteWatermarkPresetV2ResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type DeleteWatermarkPresetV2ResResponseMetadata struct {

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

type DescDenseSnapshotPresetDetailBody struct {

	// REQUIRED
	PresetList []string `json:"PresetList"`
}

type DescDenseSnapshotPresetDetailRes struct {

	// REQUIRED
	PresetDetailList []DescDenseSnapshotPresetDetailResPresetDetailListItem `json:"PresetDetailList"`

	// REQUIRED
	ResponseMetadata DescDenseSnapshotPresetDetailResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DescDenseSnapshotPresetDetailResPresetDetailListItem struct {

	// REQUIRED
	AccessKey string `json:"AccessKey"`

	// REQUIRED
	AccountID string `json:"AccountID"`

	// REQUIRED
	AsLong int32 `json:"AsLong"`

	// REQUIRED
	AsShort int32 `json:"AsShort"`

	// REQUIRED
	Bucket string `json:"Bucket"`

	// REQUIRED
	CallBackURL string `json:"CallBackUrl"`

	// REQUIRED
	CreatedAt int32 `json:"CreatedAt"`

	// REQUIRED
	Describe string `json:"Describe"`

	// REQUIRED
	Format string `json:"Format"`

	// REQUIRED
	Height int32 `json:"Height"`

	// REQUIRED
	Interval float32 `json:"Interval"`

	// REQUIRED
	KafkaCluster string `json:"KafkaCluster"`

	// REQUIRED
	KafkaTopic string `json:"KafkaTopic"`

	// REQUIRED
	Object string `json:"Object"`

	// REQUIRED
	OverwriteObject string `json:"OverwriteObject"`

	// REQUIRED
	Preset string `json:"Preset"`

	// REQUIRED
	Product string `json:"Product"`

	// REQUIRED
	Quality int32 `json:"Quality"`

	// REQUIRED
	Rate int32 `json:"Rate"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	RegionConfig string `json:"RegionConfig"`

	// REQUIRED
	S3NetworkType int32 `json:"S3NetworkType"`

	// REQUIRED
	SequenceObject string `json:"SequenceObject"`

	// REQUIRED
	ServiceID string `json:"ServiceID"`

	// REQUIRED
	Status int32 `json:"Status"`

	// REQUIRED
	TosCluster string `json:"TosCluster"`

	// REQUIRED
	TosType int32 `json:"TosType"`

	// REQUIRED
	TranscodeSuffix string `json:"TranscodeSuffix"`

	// REQUIRED
	UpdatedAt int32 `json:"UpdatedAt"`

	// REQUIRED
	Width int32 `json:"Width"`
}

type DescDenseSnapshotPresetDetailResResponseMetadata struct {

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

type DescribeActionHistoryBody struct {

	// REQUIRED; 历史记录ID
	ID string `json:"ID"`
}

type DescribeActionHistoryRes struct {

	// REQUIRED
	ResponseMetadata DescribeActionHistoryResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeActionHistoryResResult `json:"Result,omitempty"`
}

type DescribeActionHistoryResResponseMetadata struct {

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

// DescribeActionHistoryResResult - 视请求的接口而定
type DescribeActionHistoryResResult struct {

	// REQUIRED; 操作接口名称
	Action string `json:"Action"`

	// REQUIRED; 变更状态
	ActionStatus string `json:"ActionStatus"`

	// REQUIRED; action的操作时间
	ActionTime string `json:"ActionTime"`

	// REQUIRED; 应用名称
	App string `json:"App"`

	// REQUIRED; 变更的内容
	Body string `json:"Body"`

	// REQUIRED; 配置项名称
	ConfigName string `json:"ConfigName"`

	// REQUIRED; 配置项名称英文
	ConfigNameEn string `json:"ConfigNameEn"`

	// REQUIRED; 配置平台
	ConfigPlatform string `json:"ConfigPlatform"`

	// REQUIRED; 域名
	Domain string `json:"Domain"`

	// REQUIRED; ID
	ID string `json:"ID"`

	// REQUIRED; 模板名称
	PresetName string `json:"PresetName"`

	// REQUIRED; 变更回复
	Response string `json:"Response"`

	// REQUIRED; 流名
	Stream string `json:"Stream"`

	// REQUIRED; 变更人
	UserID string `json:"UserID"`

	// REQUIRED; 域名空间
	Vhost string `json:"Vhost"`

	// 变更平台详情
	ApplicationInfo *DescribeActionHistoryResResultApplicationInfo `json:"ApplicationInfo,omitempty"`

	// config的配置详情
	ConfigInfo *DescribeActionHistoryResResultConfigInfo `json:"ConfigInfo,omitempty"`
}

// DescribeActionHistoryResResultApplicationInfo - 变更平台详情
type DescribeActionHistoryResResultApplicationInfo struct {

	// REQUIRED; 配置平台
	ApplicationURL string `json:"ApplicationURL"`

	// REQUIRED; 审批人
	ApproveUserID []string `json:"ApproveUserID"`

	// REQUIRED; 变更平台工单ID
	ID string `json:"ID"`

	// REQUIRED; 变更平台的状态
	Status string `json:"Status"`
}

// DescribeActionHistoryResResultConfigInfo - config的配置详情
type DescribeActionHistoryResResultConfigInfo struct {

	// 配置进度
	ConfigProgess *float32 `json:"ConfigProgess,omitempty"`

	// 配置机器数量
	Count *int32 `json:"Count,omitempty"`

	// 配置失败的数量
	FailureCount *int32 `json:"FailureCount,omitempty"`

	// 配置时间
	ProcessTime *string `json:"ProcessTime,omitempty"`
}

type DescribeAppIDParamsAvailableBody struct {

	// 中文名称
	AppCnName *string `json:"AppCnName,omitempty"`

	// 英文名称
	AppEnName *string `json:"AppEnName,omitempty"`

	// 应用ID
	AppID *int32 `json:"AppID,omitempty"`
}

type DescribeAppIDParamsAvailableRes struct {

	// REQUIRED
	ResponseMetadata DescribeAppIDParamsAvailableResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeAppIDParamsAvailableResResult `json:"Result,omitempty"`
}

type DescribeAppIDParamsAvailableResResponseMetadata struct {

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

// DescribeAppIDParamsAvailableResResult - 视请求的接口而定
type DescribeAppIDParamsAvailableResResult struct {

	// false:该名称不可用
	CheckAppCnName *bool `json:"CheckAppCnName,omitempty"`

	// false:该名称不可用
	CheckAppEnName *bool `json:"CheckAppEnName,omitempty"`
}

type DescribeAuthBody struct {

	// REQUIRED; 鉴权场景类型。
	// * push：推流鉴权；
	// * pull：拉流鉴权；
	SceneType string `json:"SceneType"`

	// 应用名称，默认为所有应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App *string `json:"App,omitempty"`

	// 推/拉流域名。 :::tip 参数 Domain 和 Vhost 传且仅传一个。 :::
	Domain *string `json:"Domain,omitempty"`

	// 域名空间名称。 :::tip 参数 Domain 和 Vhost 传且仅传一个。 :::
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

	// REQUIRED; 鉴权状态。
	// * false：关闭推拉流鉴权；
	// * true：开启推拉流鉴权。
	AuthStatus bool `json:"AuthStatus"`

	// REQUIRED; 推/拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 鉴权场景类型。
	// * push：推流鉴权；
	// * pull：拉流鉴权。
	SceneType string `json:"SceneType"`

	// REQUIRED; 有效时长，单位为 s。
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

	// 加密字段。
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

type DescribeBillingForAdminBody struct {

	// REQUIRED; 账号
	AccountID string `json:"AccountID"`
}

type DescribeBillingForAdminRes struct {

	// REQUIRED
	ResponseMetadata DescribeBillingForAdminResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeBillingForAdminResResult          `json:"Result,omitempty"`
}

type DescribeBillingForAdminResResponseMetadata struct {

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

type DescribeBillingForAdminResResult struct {

	// REQUIRED; 不填则更新为空
	ActivityBilling DescribeBillingForAdminResResultActivityBilling `json:"ActivityBilling"`

	// REQUIRED; 订单状态
	// * 0：正常
	// * 1：正在开通
	// * 2：没有开通
	BillingStatus int32 `json:"BillingStatus"`

	// REQUIRED; 标准直播计费项，支持以下取值
	// live-traffic: 日流量月结
	// live-day-bandwidth：带宽日峰值月结
	// live-month-bandwidth：带宽月95峰值月结
	// live-bandwidth-daily：直播日峰值带宽日结
	// live-traffic-daily：直播流量日结
	// live-bandwidth-95daily：直播日95带宽日结
	// live-month-bandwidth-average：按日峰值月平均计费
	// live-month-bandwidth-95average：按带宽日95峰月平均计费
	// live-month-bandwidth-inner：对内客户
	BillingType string `json:"BillingType"`

	// REQUIRED; 下个月生效的计费方式，取值与BillingType相同
	BillingTypeNextMonth string `json:"BillingTypeNextMonth"`

	// REQUIRED; 国内Quic直播计费项，不填跟随国内标准直播取值，月结时支持以下取值 live-day-bandwidth：带宽日峰值月结 live-month-bandwidth：带宽月95峰值月结 live-month-bandwidth-average：按日峰值月平均计费
	// live-month-bandwidth-95average：日95峰月平均计费
	BillingTypeQuic string `json:"BillingTypeQuic"`

	// REQUIRED; 国内低延迟直播计费项，不填跟随国内标准直播取值，月结时支持以下取值
	// live-day-bandwidth：带宽日峰值月结
	// live-month-bandwidth：带宽月95峰值月结
	// live-month-bandwidth-average：按日峰值月平均计费
	// live-month-bandwidth-95average：日95峰月平均计费
	BillingTypeRTM string `json:"BillingTypeRTM"`

	// REQUIRED; 自定义计费方式，入参为以为样式marshal后的json串：
	// {"key1":"value1","key2":"value2"}
	// key和value取值参考：【数据工程】FCDN控制台 v2.1.0 技术评审 [https://bytedance.feishu.cn/docx/Dqkvd8WAgogvjwxwlMpcW9HznIg]
	CustomBilling string `json:"CustomBilling"`

	// REQUIRED; 最新更新时间
	LastUpdateTime string `json:"LastUpdateTime"`

	// REQUIRED; 海外标准直播计费项，不填跟随国内标准直播取值，BillingType为日结方式时，该值必须与BillingType相同，如果为月结方式，则支持以下取值：
	// live-day-bandwidth：带宽日峰值月结
	// live-month-bandwidth：带宽月95峰值月结
	// live-month-bandwidth-average：按日峰值月平均计费
	// live-month-bandwidth-95average：日95峰月平均计费
	OverseaBillingType string `json:"OverseaBillingType"`

	// REQUIRED; 海外Quic直播计费项，不填跟随国内标准直播取值，月结时支持以下取值 live-day-bandwidth：带宽日峰值月结 live-month-bandwidth：带宽月95峰值月结 live-month-bandwidth-average：按日峰值月平均计费
	// live-month-bandwidth-95average：日95峰月平均计费
	OverseaBillingTypeQuic string `json:"OverseaBillingTypeQuic"`

	// REQUIRED; 海外低延迟直播计费项，不填跟随国内标准直播取值，月结时支持以下取值 live-day-bandwidth：带宽日峰值月结
	// live-month-bandwidth：带宽月95峰值月结
	// live-month-bandwidth-average：按日峰值月平均计费
	// live-month-bandwidth-95average：日95峰月平均计费
	OverseaBillingTypeRTM string `json:"OverseaBillingTypeRTM"`

	// REQUIRED; 海外标准直播计费方式，0：拆分大区计费，1：海外统一计费，默认为0
	OverseaChargeMode int32 `json:"OverseaChargeMode"`

	// REQUIRED; 海外Quic直播计费方式，0：拆分大区计费，1：海外统一计费，默认为0
	OverseaChargeModeQuic int32 `json:"OverseaChargeModeQuic"`

	// REQUIRED; 海外低延迟直播计费方式，0：拆分大区计费，1：海外统一计费，默认为0
	OverseaChargeModeRTM int32 `json:"OverseaChargeModeRTM"`

	// REQUIRED; 状态
	// * 0：正常
	// * 1：删除
	// * 2：人工开通审批中
	// * 3：试用
	// * 4：欠费关停
	Status int32 `json:"Status"`

	// REQUIRED; trade实例的状态，和status值可能不同，因为月结有可能被设置成不处理欠费状态和回收
	// * 0：正常
	// * 4：欠费关停
	// * 5：欠费回收
	TradeStatus int32 `json:"TradeStatus"`
}

// DescribeBillingForAdminResResultActivityBilling - 不填则更新为空
type DescribeBillingForAdminResResultActivityBilling struct {

	// REQUIRED; 活动条目列表
	Activity []DescribeBillingForAdminResResultActivityBillingActivityItem `json:"Activity"`

	// REQUIRED; 检测条件
	Detect DescribeBillingForAdminResResultActivityBillingDetect `json:"Detect"`

	// REQUIRED; 当前配置是否生效，1：生效，0：不生效
	Switch int32 `json:"Switch"`
}

type DescribeBillingForAdminResResultActivityBillingActivityItem struct {

	// REQUIRED; 日期
	Date string `json:"Date"`

	// REQUIRED; 条目列表
	FeeDetailList []DescribeBillingForAdminResResultActivityBillingActivityPropertiesItemsItem `json:"FeeDetailList"`
}

type DescribeBillingForAdminResResultActivityBillingActivityPropertiesItemsItem struct {

	// REQUIRED
	ProcDetailList []DescribeBillingForAdminResResultActivityBillingActivityPropertiesItemsProcDetailListItem `json:"ProcDetailList"`

	// REQUIRED; 协议
	Protocol string `json:"Protocol"`
}

type DescribeBillingForAdminResResultActivityBillingActivityPropertiesItemsProcDetailListItem struct {

	// REQUIRED; 区域
	AreaName string `json:"AreaName"`

	// REQUIRED; 带宽，单位Gbps
	Bandwidth float32 `json:"Bandwidth"`
}

// DescribeBillingForAdminResResultActivityBillingDetect - 检测条件
type DescribeBillingForAdminResResultActivityBillingDetect struct {

	// 突发增长量场景
	BandwidthCondition *DescribeBillingForAdminResResultActivityBillingDetectBandwidthCondition `json:"BandwidthCondition,omitempty"`

	// 日峰值带宽突发增长量
	BandwidthIncrCondition *DescribeBillingForAdminResResultActivityBillingDetectBandwidthIncrCondition `json:"BandwidthIncrCondition,omitempty"`

	// 请求数场景
	RequestBandwidthCondition *DescribeBillingForAdminResResultActivityBillingDetectRequestBandwidthCondition `json:"RequestBandwidthCondition,omitempty"`
}

// DescribeBillingForAdminResResultActivityBillingDetectBandwidthCondition - 突发增长量场景
type DescribeBillingForAdminResResultActivityBillingDetectBandwidthCondition struct {

	// REQUIRED; 增量数值超过 xx 的场景xx,单位Gbps
	BandwidthIncr float32 `json:"BandwidthIncr"`

	// REQUIRED; 突发增长量超过最近一个月日峰月均带宽值的x，增长倍数
	BandwidthIncrLoop float32 `json:"BandwidthIncrLoop"`

	// REQUIRED; 1：开启，0：关闭
	Switch int32 `json:"Switch"`
}

// DescribeBillingForAdminResResultActivityBillingDetectBandwidthIncrCondition - 日峰值带宽突发增长量
type DescribeBillingForAdminResResultActivityBillingDetectBandwidthIncrCondition struct {

	// REQUIRED; 日峰值带宽突发增长量大于 xx 的场景，增量带宽，单位Gbps
	BandwidthIncr float32 `json:"BandwidthIncr"`

	// REQUIRED; 1：开启，0：关闭
	Switch int32 `json:"Switch"`
}

// DescribeBillingForAdminResResultActivityBillingDetectRequestBandwidthCondition - 请求数场景
type DescribeBillingForAdminResResultActivityBillingDetectRequestBandwidthCondition struct {

	// REQUIRED; 日峰月均值不低于 xx 的场景,日峰值月平均带宽，单位Gbps
	Bandwidth float32 `json:"Bandwidth"`

	// REQUIRED; 请求数超过近一个月的日峰月均值的x倍，增加倍速
	RequestLoop float32 `json:"RequestLoop"`

	// REQUIRED; 1：开启，0：关闭
	Switch int32 `json:"Switch"`
}

type DescribeBillingMonthAvailableRes struct {

	// REQUIRED
	ResponseMetadata DescribeBillingMonthAvailableResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeBillingMonthAvailableResResult `json:"Result,omitempty"`
}

type DescribeBillingMonthAvailableResResponseMetadata struct {

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

// DescribeBillingMonthAvailableResResult - 视请求的接口而定
type DescribeBillingMonthAvailableResResult struct {

	// REQUIRED; true: 支持，false：不支持
	MonthAvailable bool `json:"MonthAvailable"`
}

type DescribeBillingRes struct {

	// REQUIRED
	ResponseMetadata DescribeBillingResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeBillingResResult `json:"Result"`
}

type DescribeBillingResResponseMetadata struct {
	Action    *string                                  `json:"Action,omitempty"`
	Error     *DescribeBillingResResponseMetadataError `json:"Error,omitempty"`
	Region    *string                                  `json:"Region,omitempty"`
	RequestID *string                                  `json:"RequestId,omitempty"`
	Service   *string                                  `json:"Service,omitempty"`
	Version   *string                                  `json:"Version,omitempty"`
}

type DescribeBillingResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type DescribeBillingResResult struct {

	// REQUIRED; 订单状态
	// * 0：正常
	// * 1：正在开通
	// * 2：没有开通
	BillingStatus int32 `json:"BillingStatus"`

	// REQUIRED; 标准直播计费项，支持以下取值
	// live-traffic: 日流量月结
	// live-day-bandwidth：带宽日峰值月结
	// live-month-bandwidth：带宽月95峰值月结
	// live-bandwidth-daily：直播日峰值带宽日结
	// live-traffic-daily：直播流量日结
	// live-bandwidth-95daily：直播日95带宽日结
	// live-month-bandwidth-average：按日峰值月平均计费
	// live-month-bandwidth-95average：按带宽日95峰月平均计费
	// live-month-bandwidth-inner：对内客户
	BillingType string `json:"BillingType"`

	// REQUIRED; 下个月生效的计费方式，取值与BillingType相同
	BillingTypeNextMonth string `json:"BillingTypeNextMonth"`

	// REQUIRED; 国内Quic直播计费项，不填跟随国内标准直播取值，月结时支持以下取值 live-day-bandwidth：带宽日峰值月结 live-month-bandwidth：带宽月95峰值月结 live-month-bandwidth-average：按日峰值月平均计费
	// live-month-bandwidth-95average：日95峰月平均计费
	BillingTypeQuic string `json:"BillingTypeQuic"`

	// REQUIRED; 国内低延迟直播计费项，不填跟随国内标准直播取值，月结时支持以下取值
	// live-day-bandwidth：带宽日峰值月结
	// live-month-bandwidth：带宽月95峰值月结
	// live-month-bandwidth-average：按日峰值月平均计费
	// live-month-bandwidth-95average：日95峰月平均计费
	BillingTypeRTM string `json:"BillingTypeRTM"`

	// REQUIRED; 自定义计费方式，入参为以为样式marshal后的json串：
	// {"key1":"value1","key2":"value2"}
	// key和value取值参考：【数据工程】FCDN控制台 v2.1.0 技术评审 [https://bytedance.feishu.cn/docx/Dqkvd8WAgogvjwxwlMpcW9HznIg]
	CustomBilling string `json:"CustomBilling"`

	// REQUIRED; 最新更新时间
	LastUpdateTime string `json:"LastUpdateTime"`

	// REQUIRED; 海外标准直播计费项，不填跟随国内标准直播取值，BillingType为日结方式时，该值必须与BillingType相同，如果为月结方式，则支持以下取值：
	// live-day-bandwidth：带宽日峰值月结
	// live-month-bandwidth：带宽月95峰值月结
	// live-month-bandwidth-average：按日峰值月平均计费
	// live-month-bandwidth-95average：日95峰月平均计费
	OverseaBillingType string `json:"OverseaBillingType"`

	// REQUIRED; 海外Quic直播计费项，不填跟随国内标准直播取值，月结时支持以下取值 live-day-bandwidth：带宽日峰值月结 live-month-bandwidth：带宽月95峰值月结 live-month-bandwidth-average：按日峰值月平均计费
	// live-month-bandwidth-95average：日95峰月平均计费
	OverseaBillingTypeQuic string `json:"OverseaBillingTypeQuic"`

	// REQUIRED; 海外低延迟直播计费项，不填跟随国内标准直播取值，月结时支持以下取值 live-day-bandwidth：带宽日峰值月结
	// live-month-bandwidth：带宽月95峰值月结
	// live-month-bandwidth-average：按日峰值月平均计费
	// live-month-bandwidth-95average：日95峰月平均计费
	OverseaBillingTypeRTM string `json:"OverseaBillingTypeRTM"`

	// REQUIRED; 海外标准直播计费方式，0：拆分大区计费，1：海外统一计费，默认为0
	OverseaChargeMode int32 `json:"OverseaChargeMode"`

	// REQUIRED; 海外Quic直播计费方式，0：拆分大区计费，1：海外统一计费，默认为0
	OverseaChargeModeQuic int32 `json:"OverseaChargeModeQuic"`

	// REQUIRED; 海外低延迟直播计费方式，0：拆分大区计费，1：海外统一计费，默认为0
	OverseaChargeModeRTM int32 `json:"OverseaChargeModeRTM"`

	// REQUIRED; 状态。
	// * 0：正常
	// * 1：删除
	// * 2：需要人工审批
	// * 3：试用
	// * 4：欠费关停
	Status int32 `json:"Status"`
}

type DescribeCDNSnapshotHistoryBody struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。 :::tip
	// * 当您查询指定截图任务详情时，DateFrom 应设置为推流开始时间之前的任意时间。
	// * 查询的最大时间跨度为 7 天。 :::
	DateFrom string `json:"DateFrom"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	DateTo string `json:"DateTo"`

	// REQUIRED; 流名称，由 1 到 100 位数字、字母、下划线及"-"和"."组成。
	Stream string `json:"Stream"`

	// REQUIRED; 域名空间名称，由 1 到 60 位数字、字母、下划线及"-"和"."组成。
	Vhost string `json:"Vhost"`

	// 查询数据的页码，默认为 1，表示查询第一页的数据。
	PageNum *int32 `json:"PageNum,omitempty"`

	// 每页��示的数据条数，默认为 10，最大值为 1000。
	PageSize *int32 `json:"PageSize,omitempty"`

	// 截图文件保存位置，默认取值为 tos。
	// * tos：TOS 对象存储服务；
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

	// REQUIRED; 截图高度。
	Height int32 `json:"Height"`

	// REQUIRED
	ID int32 `json:"ID"`

	// REQUIRED; 截图文件保存的路径。
	Path string `json:"Path"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 截图时间戳，精度为毫秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// REQUIRED; 截图宽度。
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

	// 消息类型，缺省情况下表示查询全部。包括以下类型。
	// * push：推流开始回调；
	// * push_end：推流结束回调；
	// * snapshot：截图回调；
	// * record：录制回调；
	// * audit_snapshot：截图审核回调。
	MessageType *string `json:"MessageType,omitempty"`

	// 域名空间名称，Vhost和Domain传且仅传一个。
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

	// REQUIRED; 开启鉴权。
	AuthEnable bool `json:"AuthEnable"`

	// REQUIRED
	AuthField DescribeCallbackResResultCallbackListItemAuthField `json:"AuthField"`

	// REQUIRED; 密钥。
	AuthKeyPrimary string `json:"AuthKeyPrimary"`

	// REQUIRED; 创建时间。
	CreateTime string `json:"CreateTime"`

	// REQUIRED; 消息类型。包括以下类型。
	// * push：推流开始回调；
	// * push_end：推流结束回调；
	// * snapshot：截图回调；
	// * record：录制回调；
	// * audit_snapshot：截图审核回调。
	MessageType string `json:"MessageType"`

	// REQUIRED; 是否开启转码流回调，默认为 0。取值及含义如下所示。
	// * 0：false，不开启；
	// * 1：true，开启。
	TranscodeCallback int32 `json:"TranscodeCallback"`

	// REQUIRED; 更新时间
	UpdateTime string `json:"UpdateTime"`

	// REQUIRED; 域名空间名称。
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

	// REQUIRED; 回调类型，返回 HTTP，表示可以使用 HTTP 和 HTTPS 接收回调事件。
	CallbackType string `json:"CallbackType"`

	// REQUIRED; 回调的 URL。
	URL string `json:"URL"`
}

type DescribeCertDRMQuery struct {

	// REQUIRED; app
	App string `json:"App" query:"App"`

	// REQUIRED; 域名空间
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

type DescribeCertDetailSecretBody struct {

	// REQUIRED; 证书 ID
	ChainID string `json:"ChainID"`
}

type DescribeCertDetailSecretRes struct {

	// REQUIRED
	ResponseMetadata DescribeCertDetailSecretResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeCertDetailSecretResResult          `json:"Result,omitempty"`
}

type DescribeCertDetailSecretResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version   string                                            `json:"Version"`
	Error     *DescribeCertDetailSecretResResponseMetadataError `json:"Error,omitempty"`
	RequestID *string                                           `json:"RequestID,omitempty"`
}

type DescribeCertDetailSecretResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeCertDetailSecretResResult struct {

	// 证书名称
	CertName *string `json:"CertName,omitempty"`

	// 证书 ID
	ChainID *string `json:"ChainID,omitempty"`

	// 与证书绑定的域名
	Domain *string `json:"Domain,omitempty"`

	// 证书详细信息
	Rsa *DescribeCertDetailSecretResResultRsa `json:"Rsa,omitempty"`

	// 证书状态
	Status *string `json:"Status,omitempty"`
	UseWay *string `json:"UseWay,omitempty"`
}

// DescribeCertDetailSecretResResultRsa - 证书详细信息
type DescribeCertDetailSecretResResultRsa struct {

	// 证书类型。
	CertType *string `json:"CertType,omitempty"`

	// 证书指纹，为唯一值。
	FingerPrint *string `json:"FingerPrint,omitempty"`

	// 证书过期时间。
	NotAfter *string `json:"NotAfter,omitempty"`

	// 证书生效时间。
	NotBefore *string `json:"NotBefore,omitempty"`
	PriKey    *string `json:"PriKey,omitempty"`
	PriName   *string `json:"PriName,omitempty"`

	// 公钥数据。
	PubKey *string `json:"PubKey,omitempty"`

	// 系统自动生成的公钥文本名称。
	PubName *string `json:"PubName,omitempty"`

	// 证书序列号，为唯一值。
	SerialNumber *string `json:"SerialNumber,omitempty"`
}

type DescribeCertDetailSecretV2Body struct {

	// 账号ID
	AccountID *string `json:"AccountID,omitempty"`

	// 证书实例 ID，可以通过查询证书列表 [https://www.volcengine.com/docs/6469/81242]接口获取。 :::tip 参数ChainID与CertID传且仅传一个。 :::
	CertID *string `json:"CertID,omitempty"`

	// 证书链 ID，可以通过查询证书列表 [https://www.volcengine.com/docs/6469/81242]接口获取。 :::tip 参数ChainID与CertID传且仅传一个。 :::
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

	// REQUIRED; 证书的过期时间，RFC3339 格式的 UTC 时间，精度为 s。
	NotAfter string `json:"NotAfter"`

	// REQUIRED; 证书的生效日期，RFC3339 格式的 UTC 时间，精度为 s。
	NotBefore string `json:"NotBefore"`

	// REQUIRED; 证书状态，取值与含义的对应关系如下所示。
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

type DescribeCertDetailV2Body struct {

	// 账号ID
	AccountID *string `json:"AccountID,omitempty"`

	// 证书ID，和ChainID二选一填
	CertID *string `json:"CertID,omitempty"`

	// 证书链ID，和CertID二选一填
	ChainID *string `json:"ChainID,omitempty"`
}

type DescribeCertDetailV2Res struct {

	// REQUIRED
	ResponseMetadata DescribeCertDetailV2ResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeCertDetailV2ResResult `json:"Result,omitempty"`
}

type DescribeCertDetailV2ResResponseMetadata struct {

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
	Error   *DescribeCertDetailV2ResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeCertDetailV2ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

// DescribeCertDetailV2ResResult - 视请求的接口而定
type DescribeCertDetailV2ResResult struct {

	// 证书包含的域名
	CertDomainList []*string `json:"CertDomainList,omitempty"`

	// 证书名称
	CertName *string `json:"CertName,omitempty"`

	// 证书链ID
	ChainID *string `json:"ChainID,omitempty"`

	// 加密算法
	EncryptionAlgorithm *string `json:"EncryptionAlgorithm,omitempty"`

	// 证书指纹（SHA1）
	FingerprintSHA1 *string `json:"FingerprintSHA1,omitempty"`

	// 证书指纹（SHA256）
	FingerprintSHA256 *string `json:"FingerprintSHA256,omitempty"`

	// 签发者信息
	Issuer *string `json:"Issuer,omitempty"`

	// 证书的过期时间，RFC3339 格式的 UTC 时间，精度为 s
	NotAfter *string `json:"NotAfter,omitempty"`

	// 证书的生效日期，RFC3339 格式的 UTC 时间，精度为 s
	NotBefore *string `json:"NotBefore,omitempty"`

	// openssl解析结果
	OpenSSLFormat *string `json:"OpenSSLFormat,omitempty"`

	// 证书详细信息
	SSL *DescribeCertDetailV2ResResultSSL `json:"SSL,omitempty"`

	// 签名算法
	SignatureAlgorithm *string `json:"SignatureAlgorithm,omitempty"`

	// 证书的状态过期时间，"OK", "Expire", "1days", "7days", "15days", "30days"
	Status *string `json:"Status,omitempty"`
}

// DescribeCertDetailV2ResResultSSL - 证书详细信息
type DescribeCertDetailV2ResResultSSL struct {

	// 证书链。从叶子证书开始，到根证书。PEM编码
	Chain []*string `json:"Chain,omitempty"`

	// 证书链解析后的证书链简短信息
	ChainBriefInfo []*DescribeCertDetailV2ResResultSSLChainBriefInfoItem `json:"ChainBriefInfo,omitempty"`

	// 密钥类型，默认rsa
	KeyType *string `json:"KeyType,omitempty"`

	// 证书私钥
	PrivateKey *string `json:"PrivateKey,omitempty"`
}

type DescribeCertDetailV2ResResultSSLChainBriefInfoItem struct {

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

	// REQUIRED; 查询的起始时间，RFC3339 格式的 UTC 时间戳，精度为秒。筛选直播流结束时间符合查询条件的历史流。
	EndTimeFrom string `json:"EndTimeFrom" query:"EndTimeFrom"`

	// REQUIRED; 查询的结束时间，RFC3339 格式表示的 UTC 时间戳，精度为秒。筛选直播流结束时间符合查询条件的历史流。
	EndTimeTo string `json:"EndTimeTo" query:"EndTimeTo"`

	// REQUIRED; 查询数据的页码，取值范围为正整数。
	PageNum int32 `json:"PageNum" query:"PageNum"`

	// REQUIRED; 每页显示的数据条数，取值范围为 [1,1000]。
	PageSize int32 `json:"PageSize" query:"PageSize"`

	// 应用名称，取值与直播流地址中 AppName 字段取值相同，默认为空，表示查询所有应用名称。支持由大小写字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App *string `json:"App,omitempty" query:"App"`

	// 直播流使用的域名，默认为空，表示查询所有当前域名空间（Vhost）下的历史直播流。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console-stable.volcanicengine.com/live/main/domain/list]
	// 页面，查看需要查询的历史直播流使用的域名。
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

	// 流名称，取值与直播流地址中 StreamName 字段取值相同，默认为空表示查询所有流名称。支持由大小写字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	Stream *string `json:"Stream,omitempty" query:"Stream"`

	// 流类型，缺省情况下表示全选。支持如下取值。Origin：原始流；trans：转码流。.
	StreamType *string `json:"StreamType,omitempty" query:"StreamType"`

	// 域名空间，即直播流地址的域名（Domain）所属的域名空间（Vhost），默认为空，表示查询所有域名空间（Vhost）下的历史直播流。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]
	// 接口或在视频直播控制台的域名管理
	// [https://console-stable.volcanicengine.com/live/main/domain/list]页面，查看需要查询的历史直播流使用的域名所属的域名空间。
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

	// REQUIRED; 直播流的结束时间。
	EndTime string `json:"EndTime"`

	// REQUIRED; 历史直播流的来源类型，取值及含义如下所示。
	// * push：直推流；
	// * relay：回源流。
	SourceType string `json:"SourceType"`

	// REQUIRED; 直播流的开始时间。
	StartTime string `json:"StartTime"`

	// REQUIRED; 历史直播流使用的流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 历史直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`
}

type DescribeContentKeyRes struct {

	// REQUIRED
	ResponseMetadata DescribeContentKeyResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DescribeContentKeyResResponseMetadata struct {

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

type DescribeCustomLogConfigRes struct {

	// REQUIRED
	ResponseMetadata DescribeCustomLogConfigResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeCustomLogConfigResResult          `json:"Result,omitempty"`
}

type DescribeCustomLogConfigResResponseMetadata struct {

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

type DescribeCustomLogConfigResResult struct {

	// REQUIRED; 用户账号id
	AccountID string `json:"AccountId"`

	// REQUIRED; 用户账号名称
	AccountName string `json:"AccountName"`

	// REQUIRED; 日志获取接口名称
	ActionName string `json:"ActionName"`

	// REQUIRED; bmq集群名
	BmqCluster string `json:"BmqCluster"`

	// REQUIRED; tce集群名
	Cluster string `json:"Cluster"`

	// REQUIRED; 创建人
	Creator string `json:"Creator"`

	// REQUIRED; 延迟时间，默认300s
	DelayTime int32 `json:"DelayTime"`

	// REQUIRED; 下载时填的Type参数
	DownloadType string `json:"DownloadType"`

	// REQUIRED; 填1或0，是否补空文件，默认为0
	EmptyFile int32 `json:"EmptyFile"`

	// REQUIRED; 排除的账号ID
	ExcludedAccountIDs string `json:"ExcludedAccountIds"`

	// REQUIRED; 文件名字段名称
	FileNameFields DescribeCustomLogConfigResResultFileNameFields `json:"FileNameFields"`

	// REQUIRED; 文件名pattern
	FileNamePattern string `json:"FileNamePattern"`

	// REQUIRED; 如果没有Id，表示创建，带了Id表示更新
	ID int32 `json:"Id"`

	// REQUIRED; 日志字段名称
	LogFields DescribeCustomLogConfigResResultLogFields `json:"LogFields"`

	// REQUIRED; 日志pattern
	LogPattern string `json:"LogPattern"`

	// REQUIRED; 日志类型，如果是多个用逗号连接，全选可填*
	LogType string `json:"LogType"`

	// REQUIRED; 是否每个域名一个文件，默认为false
	SplitDomain bool `json:"SplitDomain"`

	// REQUIRED; 默认false（前端默认填ture），同一个时间范围是否允许按照大小切割文件
	SplitFile bool `json:"SplitFile"`

	// REQUIRED; 切割文件的行数，默认120w
	SplitLine int32 `json:"SplitLine"`

	// REQUIRED; 切割文件的时间，单位秒，默认3600
	SplitTime int32 `json:"SplitTime"`

	// REQUIRED; 默认0，状态，1：启动，0：禁止
	Status int32 `json:"Status"`

	// REQUIRED; 写入的topic
	Topic string `json:"Topic"`

	// REQUIRED; 特殊清洗状态，false:表示数仓单独任务进行清洗，true为通用清洗任务
	WashStatus bool `json:"WashStatus"`
}

// DescribeCustomLogConfigResResultFileNameFields - 文件名字段名称
type DescribeCustomLogConfigResResultFileNameFields struct {

	// REQUIRED; 字段名称
	Key string `json:"Key"`

	// REQUIRED; 字段类型，不能为空
	Type string `json:"Type"`

	// 备注信息，没有可以为空
	FmtValue *string `json:"FmtValue,omitempty"`

	// 字段对应中文名
	KeyCn *string `json:"KeyCn,omitempty"`

	// 敏感词替换字符串，比如：ab,cd 表示用cd替换ab，如果有多组替换用分号连接
	Transform *int32 `json:"Transform,omitempty"`
}

// DescribeCustomLogConfigResResultLogFields - 日志字段名称
type DescribeCustomLogConfigResResultLogFields struct {

	// REQUIRED; 字段名称
	Key string `json:"Key"`

	// REQUIRED; 字段类型，不能为空
	Type string `json:"Type"`

	// 备注信息，没有可以为空
	FmtValue *string `json:"FmtValue,omitempty"`

	// 字段对应中文名
	KeyCn *string `json:"KeyCn,omitempty"`

	// 敏感词替换字符串，比如：ab,cd 表示用cd替换ab，如果有多组替换用分号连接
	Transform *int32 `json:"Transform,omitempty"`
}

type DescribeDenyConfigBody struct {

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// App 的名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。Domain 和 App 二选一填。
	App *string `json:"App,omitempty"`

	// 推/拉流域名。
	Domain *string `json:"Domain,omitempty"`
}

type DescribeDenyConfigRes struct {

	// REQUIRED
	ResponseMetadata DescribeDenyConfigResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeDenyConfigResResult          `json:"Result,omitempty"`
}

type DescribeDenyConfigResResponseMetadata struct {

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
	Error   *DescribeDenyConfigResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeDenyConfigResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeDenyConfigResResult struct {

	// 配置列表。
	DenyList []*DescribeDenyConfigResResultDenyListItem `json:"DenyList,omitempty"`
}

type DescribeDenyConfigResResultDenyListItem struct {

	// App的名称。
	App *string `json:"App,omitempty"`

	// 配置详情列表。
	DenyConfig []*DescribeDenyConfigResResultDenyListPropertiesItemsItem `json:"DenyConfig,omitempty"`

	// 推拉流域名。
	Domain *string `json:"Domain,omitempty"`

	// 域名空间名称。
	Vhost *string `json:"Vhost,omitempty"`
}

type DescribeDenyConfigResResultDenyListPropertiesItemsItem struct {

	// 白名单。
	AllowList []*string `json:"AllowList,omitempty"`

	// 城市
	City *string `json:"City,omitempty"`

	// 大洲
	Continent *string `json:"Continent,omitempty"`

	// 国家码
	Country *string `json:"Country,omitempty"`

	// 黑名单。
	DenyList []*string `json:"DenyList,omitempty"`

	// 格式类型，比如 HTTP、RTMP。
	FmtType []*string `json:"FmtType,omitempty"`

	// 运营商。
	ISP *string `json:"ISP,omitempty"`

	// 协议类型，比如 TCP、KCP、QUIC。
	ProType []*string `json:"ProType,omitempty"`

	// 区域
	Region *string `json:"Region,omitempty"`
}

type DescribeDenyConfigV2Body struct {

	// REQUIRED; 域名空间名称
	Vhost string `json:"Vhost"`

	// App名称，app和domain二选一填
	App *string `json:"App,omitempty"`

	// 推/拉流域名
	Domain *string `json:"Domain,omitempty"`

	// 服务类型，pull：拉流，push：推流
	ServiceType *string `json:"ServiceType,omitempty"`
}

type DescribeDenyConfigV2Res struct {

	// REQUIRED
	ResponseMetadata DescribeDenyConfigV2ResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeDenyConfigV2ResResult          `json:"Result,omitempty"`
}

type DescribeDenyConfigV2ResResponseMetadata struct {

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
	Error   *DescribeDenyConfigV2ResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeDenyConfigV2ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeDenyConfigV2ResResult struct {

	// 配置列表
	DenyList []*DescribeDenyConfigV2ResResultDenyListItem `json:"DenyList,omitempty"`
}

type DescribeDenyConfigV2ResResultDenyListItem struct {

	// REQUIRED; 创建时间
	CreateTime string `json:"CreateTime"`

	// REQUIRED; 推拉流域名
	Domain string `json:"Domain"`

	// REQUIRED; 服务类型
	ServiceType string `json:"ServiceType"`

	// REQUIRED; 更新时间
	UpdateTime string `json:"UpdateTime"`

	// REQUIRED; 域名空间名称
	Vhost string `json:"Vhost"`

	// App的名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App *string `json:"App,omitempty"`

	// 配置详情列表
	DenyConfigDetail []*DescribeDenyConfigV2ResResultDenyListPropertiesItemsItem `json:"DenyConfigDetail,omitempty"`
}

type DescribeDenyConfigV2ResResultDenyListPropertiesItemsItem struct {

	// REQUIRED; 黑/白名单 IP 列表。
	IPList []string `json:"IPList"`

	// REQUIRED; 传输协议
	ProType []string `json:"ProType"`

	// REQUIRED; 限制类型。
	// * allow：IP 白名单；
	// * deny：IP 黑名单。
	Type string `json:"Type"`

	// 城市限制
	City []*string `json:"City,omitempty"`

	// 国家限制，国家码
	Country []*string `json:"Country,omitempty"`

	// 拉流类型
	FmtType []*string `json:"FmtType,omitempty"`

	// 运营商限制
	ISP []*string `json:"ISP,omitempty"`

	// 省份限制
	Province []*string `json:"Province,omitempty"`

	// 大区限制
	Region []*string `json:"Region,omitempty"`

	// streams名称
	Streams []*string `json:"Streams,omitempty"`
}

type DescribeDomainBody struct {

	// REQUIRED; 域名列表。
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

	// REQUIRED; 所绑定证书支持的泛域名。
	CertDomain string `json:"CertDomain"`

	// REQUIRED; 绑定的证书名称。
	CertName string `json:"CertName"`

	// REQUIRED; 绑定的证书信息。
	ChainID string `json:"ChainID"`

	// REQUIRED; CNAME 状态。
	// * 0：未配置 CNAME；
	// * 1：已配置 CNAME。
	CnameCheck int32 `json:"CnameCheck"`

	// REQUIRED; 创建时间。
	CreateTime string `json:"CreateTime"`

	// REQUIRED; 推/拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 域名是否可用的状态。
	// * 0：正常，域名为可用状态；
	// * 1：配置中，域名为可用状态；
	// * 2：不可用，域名为其他的不可用状态。
	DomainCheck int32 `json:"DomainCheck"`

	// REQUIRED; ICP 备案校验是否通过，是否过期信息。
	ICPCheck int32 `json:"ICPCheck"`

	// REQUIRED; 绑定的推流域名。
	PushDomain string `json:"PushDomain"`

	// REQUIRED; 区域，包含以下类型。
	// * cn：中国大陆；
	// * cn-global：全球；
	// * cn-oversea：海外及港澳台。
	Region string `json:"Region"`

	// REQUIRED; 域名状态。状态说明如下所示。
	// * 0：正常；
	// * 1：审核中；
	// * 2：禁用，禁止使用，此时 domain 不生效；
	// * 3：删除；
	// * 4：审核被驳回。审核不通过，需要重新创建并审核；
	// * 5：欠费关停。
	Status int32 `json:"Status"`

	// REQUIRED; 域名类型，包含两种类型。
	// * push：推流域名；
	// * pull-flv：拉流域名，包含 RTMP、FLV、HLS 格式。
	Type string `json:"Type"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`
}

type DescribeDomainVerifyBody struct {

	// REQUIRED; 推拉流域名列表，最多十个
	Domains []string `json:"Domains"`
}

type DescribeDomainVerifyRes struct {

	// REQUIRED
	ResponseMetadata DescribeDomainVerifyResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeDomainVerifyResResult `json:"Result,omitempty"`
}

type DescribeDomainVerifyResResponseMetadata struct {

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

// DescribeDomainVerifyResResult - 视请求的接口而定
type DescribeDomainVerifyResResult struct {

	// 校验返回列表
	DomainList []*DescribeDomainVerifyResResultDomainListItem `json:"DomainList,omitempty"`
}

type DescribeDomainVerifyResResultDomainListItem struct {

	// REQUIRED; 推拉流域名
	Domains []string `json:"Domains"`

	// REQUIRED; 根域名
	RootDomain string `json:"RootDomain"`

	// REQUIRED; true: 已经校验
	Verify bool `json:"Verify"`
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

	// REQUIRED; drm配置
	DRMItem DescribeEncryptDRMResResultDRMItem `json:"DRMItem"`
}

// DescribeEncryptDRMResResultDRMItem - drm配置
type DescribeEncryptDRMResResultDRMItem struct {

	// REQUIRED; apikey
	APIKey string `json:"APIKey"`

	// REQUIRED
	ApplicationSecretKey string `json:"ApplicationSecretKey"`

	// REQUIRED; 证书文件名
	CertificateFileName string `json:"CertificateFileName"`

	// REQUIRED; 证书名称
	CertificateName string `json:"CertificateName"`

	// REQUIRED
	PrivateKey string `json:"PrivateKey"`

	// REQUIRED; 私钥文件名
	PrivateKeyFileName string `json:"PrivateKeyFileName"`
}

type DescribeForbiddenStreamInfoByPageQuery struct {

	// REQUIRED; 查询数据的页码，取值范围为正整数。
	PageNum int32 `json:"PageNum" query:"PageNum"`

	// REQUIRED; 每页显示的数据条数，取值范围为 [1,1000]。
	PageSize int32 `json:"PageSize" query:"PageSize"`

	// 应用名称，取值与禁推直播流时设置的应用名称相同，默认为空，表示查询当前域名空间（Vhost）下所有的禁推流。支持由大小写字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App *string `json:"App,omitempty" query:"App"`

	// 直播流使用的域名，取值与禁推直播流时设置的应用名称相同，默认为空，表示查询所有当前域名空间（Vhost）下的禁推直播流。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]
	// 接口或在视频直播控制台的域名管理
	// [https://console-stable.volcanicengine.com/live/main/domain/list]页面，查看需要查询的禁推直播流使用的域名。
	Domain *string `json:"Domain,omitempty" query:"Domain"`

	// 指定是否模糊匹配流名称。缺省情况为精准匹配，支持的取值及含义如下所示。
	// * fuzzy：模糊匹配；
	// * strict：精准匹配。
	QueryType *string `json:"QueryType,omitempty" query:"QueryType"`

	// 排列方式，根据推流结束时间排序，默认值为 desc，支持的取值及含义如下所示。
	// * asc：从时间最远到最近排序；
	// * desc：从时间最近到最远排序。
	Sort *string `json:"Sort,omitempty" query:"Sort"`

	// 流名称，取值与禁推直播流时设置的流名称相同，默认查询所有流名称，由 1 到 100 位数字、字母、下划线及"-"和"."组成，如果指定 Stream，必须同时指定 App 的值。
	Stream *string `json:"Stream,omitempty" query:"Stream"`

	// 域名空间，取值与禁推直播流时设置的域名空间相同，默认为空，表示查询所有域名空间（Vhost）下的禁推流。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台��域名管理
	// [https://console-stable.volcanicengine.com/live/main/domain/list]页面，查看需要查询的禁推流使用的域名所属的域名空间。
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

	// REQUIRED; 禁推流被禁推的开始时间。
	CreateTime string `json:"CreateTime"`

	// REQUIRED; 禁推流的域名。
	Domain string `json:"Domain"`

	// REQUIRED; 禁推流结束禁推的时间。
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

	// REQUIRED; The type of HTTP header configurations you want to query:
	// * 0: Response headers.
	// * 1: Request headers.
	Phase int32 `json:"Phase"`

	// REQUIRED; The domain name space.
	Vhost string `json:"Vhost"`

	// The domain name.
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

	// REQUIRED; A list of header configurations of the domain name.
	HeaderConfigList []DescribeHTTPHeaderConfigResResultHeaderConfigListItem `json:"HeaderConfigList"`
}

type DescribeHTTPHeaderConfigResResultHeaderConfigListItem struct {

	// REQUIRED; Whether the original headers is excluded.
	// * 0: Included.
	// * 1: Excluded.
	BlockOriginal int32 `json:"BlockOriginal"`

	// REQUIRED; The domain name.
	Domain string `json:"Domain"`

	// REQUIRED; Whether the configuration is enabled.
	// * true: Enabled.
	// * false: Disabled.
	Enable bool `json:"Enable"`

	// REQUIRED; A list of HTTP headers you want to query.
	HeaderDetailList []DescribeHTTPHeaderConfigResResultHeaderConfigListPropertiesItemsItem `json:"HeaderDetailList"`

	// REQUIRED; The domain name space.
	Vhost string `json:"Vhost"`
}

type DescribeHTTPHeaderConfigResResultHeaderConfigListPropertiesItemsItem struct {

	// REQUIRED; The type of the header value:
	// * 0: Constant
	// * 1: Variable
	HeaderFieldType int32 `json:"HeaderFieldType"`

	// REQUIRED; The header name.
	HeaderKey string `json:"HeaderKey"`

	// REQUIRED; The header value. The header value can be a constant or one of the following variables: For the header in a response,
	// the header value can be the following variables:
	// * ${domain}: The domain name in the client request. Example:example.com
	// * ${uri}: The path of the client request excluding the query parameters. If the client request is rewritten, this variable
	// represents the rewritten path. Example:/dir/sample.php
	// * ${args}: The query parameters in the client request. If the client request is rewritten, this variable represents the
	// rewritten parameters. Example:color=red&n=10
	// * ${remote_addr}: The IP address of the client sending the request. Example:10.10.10.10
	// * ${server_addr}: The IP address of the edge server responding to the client request. Example:10.10.10.10
	// For the header in a request, the header value can be the following variables:
	// * ${upstream_host}: The domain name in the origin-pull request. Example:example.com
	// * ${upstream_uri}: The path of the origin-pull request excluding the query parameters. If the request is rewritten, this
	// variable represents the rewritten path. Example:/dir/sample.php
	// * ${upstream_args}: The query parameters in the origin-pull request. If the request is rewritten, this variable represents
	// the rewritten parameters. Example:color=red&n=10
	HeaderValue string `json:"HeaderValue"`
}

type DescribeHeaderConfigBody struct {

	// REQUIRED
	Vhost      string  `json:"Vhost"`
	App        *string `json:"App,omitempty"`
	HeaderType *string `json:"HeaderType,omitempty"`
}

type DescribeHeaderConfigRes struct {

	// REQUIRED
	ResponseMetadata DescribeHeaderConfigResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeHeaderConfigResResult `json:"Result,omitempty"`
}

type DescribeHeaderConfigResResponseMetadata struct {

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

// DescribeHeaderConfigResResult - 视请求的接口而定
type DescribeHeaderConfigResResult struct {

	// REQUIRED
	HeaderConfigListV2 []DescribeHeaderConfigResResultHeaderConfigListV2Item `json:"HeaderConfigListV2"`
}

type DescribeHeaderConfigResResultHeaderConfigListV2Item struct {

	// REQUIRED
	App string `json:"App"`

	// REQUIRED; 创建时间
	CreateTime string `json:"CreateTime"`

	// REQUIRED
	HeaderDetailList []DescribeHeaderConfigResResultHeaderConfigListV2PropertiesItemsItem `json:"HeaderDetailList"`

	// REQUIRED
	RealJSON string `json:"RealJSON"`

	// REQUIRED; 更新时间
	UpdateTime string `json:"UpdateTime"`

	// REQUIRED
	Vhost string `json:"Vhost"`
}

type DescribeHeaderConfigResResultHeaderConfigListV2PropertiesItemsHeaderDetailListItem struct {

	// REQUIRED
	HeaderKey string `json:"HeaderKey"`

	// REQUIRED
	HeaderValue string `json:"HeaderValue"`
}

type DescribeHeaderConfigResResultHeaderConfigListV2PropertiesItemsItem struct {

	// REQUIRED
	HeaderDetailList []DescribeHeaderConfigResResultHeaderConfigListV2PropertiesItemsHeaderDetailListItem `json:"HeaderDetailList"`

	// REQUIRED
	HeaderType string `json:"HeaderType"`
}

type DescribeIPAccessRuleBody struct {

	// REQUIRED; 推/拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 域名空间名称。
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

	// REQUIRED; 是否开启当前限制。
	// * true: 开启；
	// * false: 关闭。
	Enable bool `json:"Enable"`

	// REQUIRED; 名单中的 IP 信息。
	IPList []string `json:"IPList"`

	// REQUIRED; IP 访问限制的类型，取值及含义如下。
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

	// REQUIRED; IP 所属地区。非归属火山引擎视频直播的 IP，返回“-”。
	City string `json:"City"`

	// REQUIRED; IP 地址
	IP string `json:"Ip"`

	// REQUIRED; IP 所属运营商。非归属火山引擎视频直播的 IP，返回”-“。 您可以通过DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974]接口查看运营商标识符对应的运营商名称。
	Isp string `json:"Isp"`

	// REQUIRED; 是否归属于火山引擎 CDN 节点。
	// * true：属于；
	// * false：不属于。
	LiveCdnIP bool `json:"LiveCdnIp"`

	// REQUIRED; IP 所属国家或地区。非归属火山引擎视频直播的 IP，返回“-”。
	Location string `json:"Location"`

	// REQUIRED; IP 所属省。非归属火山引擎视频直播的 IP，返回“-”。
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

	// REQUIRED; app
	App string `json:"App" query:"App"`

	// REQUIRED; DRM加密的类型枚举，可以取fp（代表fairplay）或wv（代表widevine）
	DRMType string `json:"DRMType" query:"DRMType"`

	// REQUIRED; 拉流域名
	Domain string `json:"Domain" query:"Domain"`

	// REQUIRED; 流名
	StreamName string `json:"StreamName" query:"StreamName"`

	// REQUIRED; 域名空间
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

type DescribeLiveAccountFeeConfigRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveAccountFeeConfigResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeLiveAccountFeeConfigResResult `json:"Result,omitempty"`
}

type DescribeLiveAccountFeeConfigResResponseMetadata struct {

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

// DescribeLiveAccountFeeConfigResResult - 视请求的接口而定
type DescribeLiveAccountFeeConfigResResult struct {

	// REQUIRED; 配置列表
	FeeConfigList []DescribeLiveAccountFeeConfigResResultFeeConfigListItem `json:"FeeConfigList"`

	// REQUIRED; 配置个数
	Total float32 `json:"Total"`
}

type DescribeLiveAccountFeeConfigResResultFeeConfigListItem struct {

	// REQUIRED; 账号id
	AccountID string `json:"AccountId"`

	// REQUIRED; 配置id
	ID int32 `json:"Id"`

	// 进制
	Base *int32 `json:"Base,omitempty"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty"`

	// 创建者
	Creator *string `json:"Creator,omitempty"`

	// 上浮系数
	Factor      *float32  `json:"Factor,omitempty"`
	FreeFeeList []*string `json:"FreeFeeList,omitempty"`

	// 是否开启闲忙时
	StageEnable *bool `json:"StageEnable,omitempty"`

	// 闲忙时生效时间
	StageTime *string `json:"StageTime,omitempty"`

	// 上浮系数生效时间
	StartTime *string `json:"StartTime,omitempty"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty"`
}

type DescribeLiveAccountFeeTypeBody struct {

	// REQUIRED; 账号
	AccountID string `json:"AccountId"`
}

type DescribeLiveAccountFeeTypeRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveAccountFeeTypeResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeLiveAccountFeeTypeResResult `json:"Result,omitempty"`
}

type DescribeLiveAccountFeeTypeResResponseMetadata struct {

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

// DescribeLiveAccountFeeTypeResResult - 视请求的接口而定
type DescribeLiveAccountFeeTypeResResult struct {

	// 计费类型，traffic流量计费，bandwidth带宽计费
	FeeType *string `json:"FeeType,omitempty"`
}

type DescribeLiveActivityBandwidthDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天；
	// * 86400：1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 域名列表，缺省情况表示当前用户的所有推拉流域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，缺省情况下表示所有运营商，支持的运营商如下所示。
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
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// 推拉流协议，缺省情况下表示所有协议类型，支持的协议如下所示。
	// * HTTP-FLV：基于 HTTP 协议的推拉流协议，使用 FLV 格式传输视频格式。
	// * HTTP-HLS：基于 HTTP 协议的推拉流协议，使用 TS 格式传输视频格式。
	// * RTMP：Real Time Message Protocol，实时信息传输协议。
	// * RTM：Real Time Media，超低延时直播协议。
	// * SRT：Secure Reliable Transport，安全可靠传输协议。
	// * QUIC：Quick UDP Internet Connections，一种基于 UDP 的全新的低延时互联网传输协议。
	// :::tip 如果查询推拉流协议为 QUIC，不能同时查询其他协议。 :::
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	RegionList []*DescribeLiveActivityBandwidthDataBodyRegionListItem `json:"RegionList,omitempty"`

	// 客户端 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	UserRegionList []*DescribeLiveActivityBandwidthDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveActivityBandwidthDataBodyRegionListItem struct {

	// 区域信息中的大区标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按国家筛选，需要同时传入Area和Country。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符，国外暂不支持该参数，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按省筛选，需要同时传入Area、Country和Province。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveActivityBandwidthDataBodyUserRegionListItem struct {

	// 大区，映射关系请参见区域映射
	Area *string `json:"Area,omitempty"`

	// 国家，映射关系请参见区域映射。如果按国家筛选，需要同时传入 Area 和 Country。
	Country *string `json:"Country,omitempty"`

	// 国内为省，国外暂不支持该参数，映射关系请参见区域映射。如果按省筛选，需要同时传入 Area、Country 和 Province。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveActivityBandwidthDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveActivityBandwidthDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveActivityBandwidthDataResResult `json:"Result"`
}

type DescribeLiveActivityBandwidthDataResResponseMetadata struct {

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

type DescribeLiveActivityBandwidthDataResResult struct {

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 300：5 分钟；
	// * 3600：1 小时；
	// * 86400：1 天。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 所有时间粒度的数据。
	BandwidthDataList []DescribeLiveActivityBandwidthDataResResultBandwidthDataListItem `json:"BandwidthDataList"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 当前查询条件下的峰值带宽，单位为 Mbps。
	PeakBandwidth int32 `json:"PeakBandwidth"`

	// REQUIRED; 峰值带宽的时间戳，RFC3339 格式的 UTC 时间，精度为秒。
	PeakTimestamp string `json:"PeakTimestamp"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，标识符与运营商的对应关系如下。
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
	ISPList []*string `json:"ISPList,omitempty"`

	// 推拉流协议，协议说明如下。
	// * HTTP-FLV：基于 HTTP 协议的推拉流协议，使用 FLV 格式传输视频格式。
	// * HTTP-HLS：基于 HTTP 协议的推拉流协议，使用 TS 格式传输视频格式。
	// * RTMP：Real Time Message Protocol，实时信息传输协议。
	// * RTM：Real Time Media，超低延时直播协议。
	// * SRT：Secure Reliable Transport，安全可靠传输协议。
	// * QUIC：Quick UDP Internet Connections，一种基于 UDP 的全新的低延时互联网传输协议。
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域列表。
	RegionList []*DescribeLiveActivityBandwidthDataResResultRegionListItem `json:"RegionList,omitempty"`

	// 客户端 IP 所属区域列表。
	UserRegionList []*DescribeLiveActivityBandwidthDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveActivityBandwidthDataResResultBandwidthDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的峰值带宽，单位为 Mbps。
	Bandwidth int32 `json:"Bandwidth"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLiveActivityBandwidthDataResResultRegionListItem struct {

	// 区域信息中的大区标识符。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveActivityBandwidthDataResResultUserRegionListItem struct {

	// 区域信息中的大区标识符。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveAuditDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 86400：（默认值）1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 数据拆分的维度，缺省情况下不进行数据拆分，支持的维度如下所示。
	// * Domain：域名。 :::tip 配置数据拆分维度时，对应的维度参数传入多个值时会返回按维度进行拆分的数据；对应的维度只传入一个值时不返回按维度进行拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，缺省情况表示当前用户的所有推拉流域名。
	DomainList []*string `json:"DomainList,omitempty"`
}

type DescribeLiveAuditDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveAuditDataResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeLiveAuditDataResResult          `json:"Result,omitempty"`
}

type DescribeLiveAuditDataResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                         `json:"Version"`
	Error   *DescribeLiveAuditDataResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeLiveAuditDataResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeLiveAuditDataResResult struct {

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 86400：1 天。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 所有时间粒度的数据。
	AuditDataList []DescribeLiveAuditDataResResultAuditDataListItem `json:"AuditDataList"`

	// REQUIRED; 按维度拆分后的数据。
	AuditDetailDataList []DescribeLiveAuditDataResResultAuditDetailDataListItem `json:"AuditDetailDataList"`

	// REQUIRED; 数据拆分的维度，维度说明如下。
	// * Domain：域名。
	DetailField []string `json:"DetailField"`

	// REQUIRED; 域名列表。
	DomainList []string `json:"DomainList"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 当前查询条件下的截图审核总张数。
	TotalCount int32 `json:"TotalCount"`
}

type DescribeLiveAuditDataResResultAuditDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的截图审核总张数。
	Count int32 `json:"Count"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLiveAuditDataResResultAuditDetailDataListItem struct {

	// REQUIRED; 按维度进行数据拆分后，当前维度下所有时间粒度的数据。
	AuditDataList []DescribeLiveAuditDataResResultAuditDetailDataListPropertiesItemsItem `json:"AuditDataList"`

	// REQUIRED; 按域名维度进行数据拆分时的域名信息。
	Domain string `json:"Domain"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的截图审核总张数。
	TotalCount int32 `json:"TotalCount"`
}

type DescribeLiveAuditDataResResultAuditDetailDataListPropertiesItemsItem struct {

	// REQUIRED; 截图审核总张数
	Count int32 `json:"Count"`

	// REQUIRED; 时间片起始时刻。RFC3339 格式的 UTC 时间，精度为 s，例如，2022-04-13T00:00:00+08:00
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLiveBandwidthDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询最大时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询最大时间范围为 366 天；
	// * 86400：1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询最大时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 数据拆分的维度，缺省情况下不进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * ISP：运营商；
	// * Protocol：推拉流协议。 :::tip 配置数据拆分维度时，对应的维度参数传入多个值时会返回按维度进行拆分的数据；对应的维度只传入一个值时不返回按维度进行拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，缺省情况表示当前用户的所有推拉流域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，缺省情况下表示所有运营商，支持的运营商如下所示。
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
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// 推拉流协议，缺省情况下表示所有协议类型，支持的协议如下所示。
	// * HTTP-FLV：基于 HTTP 协议的推拉流协议，使用 FLV 格式传输视频格式。
	// * HTTP-HLS：基于 HTTP 协议的推拉流协议，使用 TS 格式传输视频格式。
	// * RTMP：Real Time Message Protocol，实时信息传输协议。
	// * RTM：Real Time Media，超低延时直播协议。
	// * SRT：Secure Reliable Transport，安全可靠传输协议。
	// * QUIC：Quick UDP Internet Connections，一种基于 UDP 的全新的低延时互联网传输协议。
	// :::tip 如果查询推拉流协议为 QUIC，不能同时查询其他协议。 :::
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	RegionList []*DescribeLiveBandwidthDataBodyRegionListItem `json:"RegionList,omitempty"`

	// 客户端 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	UserRegionList []*DescribeLiveBandwidthDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveBandwidthDataBodyRegionListItem struct {

	// 区域信息中的大区标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按国家筛选，需要同时传入Area和Country。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符，国外暂不支持该参数，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按省筛选，需要同时传入Area、Country和Province。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveBandwidthDataBodyUserRegionListItem struct {

	// 大区，映射关系请参见区域映射
	Area *string `json:"Area,omitempty"`

	// 国家，映射关系请参见区域映射。如果按国家筛选，需要同时传入 Area 和 Country。
	Country *string `json:"Country,omitempty"`

	// 国内为省，国外暂不支持该参数，映射关系请参见区域映射。如果按省筛选，需要同时传入 Area、Country 和 Province。
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

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 300：5 分钟；
	// * 3600：1 小时；
	// * 86400：1 天。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 所有时间粒度的数据。
	BandwidthDataList []DescribeLiveBandwidthDataResResultBandwidthDataListItem `json:"BandwidthDataList"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询时间范围内的下行峰值带宽，单位为 Mbps。
	PeakDownBandwidth float32 `json:"PeakDownBandwidth"`

	// REQUIRED; 查询时间范围内的上行峰值带宽，单位为 Mbps。
	PeakUpBandwidth float32 `json:"PeakUpBandwidth"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 按维度拆分的数据。 :::tip 请求时，DomainList、ProtocolList和ISPList至少有一个参数传入了多个值时，会返回该参数；否则不返回该参数。
	// 优化：当配置了数据拆分的维度，且对应的维度参数传入多个值时会返回按维度拆分的数据。 :::
	BandwidthDetailDataList []*DescribeLiveBandwidthDataResResultBandwidthDetailDataListItem `json:"BandwidthDetailDataList,omitempty"`

	// 数据拆分的维度，维度说明如下所示。
	// * Domain：域名；
	// * ISP：运营商；
	// * Protocol：推拉流协议。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，标识符与运营商的对应关系如下。
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
	ISPList []*string `json:"ISPList,omitempty"`

	// 推拉流协议，协议说明如下。
	// * HTTP-FLV：基于 HTTP 协议的推拉流协议，使用 FLV 格式传输视频格式。
	// * HTTP-HLS：基于 HTTP 协议的推拉流协议，使用 TS 格式传输视频格式。
	// * RTMP：Real Time Message Protocol，实时信息传输协议。
	// * RTM：Real Time Media，超低延时直播协议。
	// * SRT：Secure Reliable Transport，安全可靠传输协议。
	// * QUIC：Quick UDP Internet Connections，一种基于 UDP 的全新的低延时互联网传输协议。
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域列表。
	RegionList []*DescribeLiveBandwidthDataResResultRegionListItem `json:"RegionList,omitempty"`

	// 客户端 IP 所属区域列表。
	UserRegionList []*DescribeLiveBandwidthDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveBandwidthDataResResultBandwidthDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的下行峰值带宽，单位为 Mbps。
	DownBandwidth float32 `json:"DownBandwidth"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
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

	// REQUIRED; 时间片起始时刻。RFC3339 格式的 UTC 时间，精度为 s，例如，2022-04-13T00:00:00+08:00
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 上行带宽，单位为 Mbps
	UpBandwidth float32 `json:"UpBandwidth"`
}

type DescribeLiveBandwidthDataResResultRegionListItem struct {

	// 区域信息中的大区标识符。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveBandwidthDataResResultUserRegionListItem struct {

	// 大区
	Area *string `json:"Area,omitempty"`

	// 国家
	Country *string `json:"Country,omitempty"`

	// 国内为省，国外暂不支持该参数
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveBatchOnlineStreamMetricsBody struct {

	// REQUIRED; 流类型，push:推流，relay: 回源流
	StreamType string `json:"StreamType"`

	// app
	App *string `json:"App,omitempty"`

	// 推流域名列表，缺省情况下表示所有域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 分页页码，默认是1，取值范围[1，10000]
	PageNum *int32 `json:"PageNum,omitempty"`

	// 每页的大小，默认100，取值范围[1, 1000]
	PageSize *int32 `json:"PageSize,omitempty"`

	// 流名
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveBatchOnlineStreamMetricsRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveBatchOnlineStreamMetricsResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveBatchOnlineStreamMetricsResResult `json:"Result"`
}

type DescribeLiveBatchOnlineStreamMetricsResResponseMetadata struct {

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

type DescribeLiveBatchOnlineStreamMetricsResResult struct {

	// REQUIRED; 查询结果的分页信息。
	Pagination DescribeLiveBatchOnlineStreamMetricsResResultPagination `json:"Pagination"`

	// REQUIRED; 按指定时间粒度聚合的监控数据。
	StreamMetricList []DescribeLiveBatchOnlineStreamMetricsResResultStreamMetricListItem `json:"StreamMetricList"`

	// REQUIRED; 流类型，push:推流，relay: 回源流
	StreamType string `json:"StreamType"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 推流域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 流名称。
	Stream *string `json:"Stream,omitempty"`
}

// DescribeLiveBatchOnlineStreamMetricsResResultPagination - 查询结果的分页信息。
type DescribeLiveBatchOnlineStreamMetricsResResultPagination struct {

	// REQUIRED; 当前所在分页的页码。
	PageNum int32 `json:"PageNum"`

	// REQUIRED; 每页显示的数据条数。
	PageSize int32 `json:"PageSize"`

	// REQUIRED; 查询结果的数据总条数。
	TotalCount int32 `json:"TotalCount"`
}

type DescribeLiveBatchOnlineStreamMetricsResResultStreamMetricListItem struct {

	// REQUIRED; 音频编码格式
	Acodec string `json:"Acodec"`

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 音频帧率，单位为 fps
	AudioFps float32 `json:"AudioFps"`

	// REQUIRED; 音频码率，单位为 kbps
	AudioRate float32 `json:"AudioRate"`

	// REQUIRED; 客户端ip
	ClientIP string `json:"ClientIp"`

	// REQUIRED; 推流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 分辨率
	Resolution string `json:"Resolution"`

	// REQUIRED; 服务器ip
	ServerIP string `json:"ServerIp"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 流开始时间，rfc3339格式
	StreamBeginTime string `json:"StreamBeginTime"`

	// REQUIRED; 编码格式
	Vcodec string `json:"Vcodec"`

	// REQUIRED; 视频码率，单位为 kbps
	VideoFps float32 `json:"VideoFps"`

	// REQUIRED; 视频帧率，单位为 fps
	VideoRate float32 `json:"VideoRate"`
}

type DescribeLiveBatchPushStreamAvgMetricsBody struct {

	// REQUIRED; 推流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。 :::tip 单次查询最大时间跨度为 1 天，历史查询最大时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 5：5 秒；
	// * 30：30 秒；
	// * 60：（默认值）1 分钟。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 流名称。 :::tip 使用 Stream 构造请求时，需同时定义 App 参数，不可缺省。 :::
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveBatchPushStreamAvgMetricsRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveBatchPushStreamAvgMetricsResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveBatchPushStreamAvgMetricsResResult `json:"Result"`
}

type DescribeLiveBatchPushStreamAvgMetricsResResponseMetadata struct {

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

type DescribeLiveBatchPushStreamAvgMetricsResResult struct {

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 5：5 秒；
	// * 30：30 秒；
	// * 60：1 分钟。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 推流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 流的信息，包含域名、应用名称、流名称和监控数据。
	StreamMetricList []DescribeLiveBatchPushStreamAvgMetricsResResultStreamMetricListItem `json:"StreamMetricList"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveBatchPushStreamAvgMetricsResResultStreamMetricListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 推流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 按指定时间粒度聚合的监控数据。
	MetricList []DescribeLiveBatchPushStreamAvgMetricsResResultStreamMetricListPropertiesItemsItem `json:"MetricList"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`
}

type DescribeLiveBatchPushStreamAvgMetricsResResultStreamMetricListPropertiesItemsItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的音频码率平均值，单位为 kbps。
	AudioBitrate float32 `json:"AudioBitrate"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻音频帧显示时间戳差值的平均值，单位为毫秒。
	AudioFrameGap int32 `json:"AudioFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内的音频帧率平均值，单位为 fps。
	AudioFramerate float32 `json:"AudioFramerate"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个音频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	AudioPts int32 `json:"AudioPts"`

	// REQUIRED; 当前数据聚合时间粒度内的视频码率平均值，单位为 kbps。
	Bitrate float32 `json:"Bitrate"`

	// REQUIRED; 当前数据聚合时间粒度内的视频帧率平均值，单位为 fps。
	Framerate float32 `json:"Framerate"`

	// REQUIRED; 当前数据聚合时间粒度内，所有音视频帧显示时间戳差值的平均值，即所有 AudioPts 与 VideoPts 差值的平均值，单位为毫秒。
	PtsDelta int32 `json:"PtsDelta"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间， RFC3339 格式的 UTC 时间，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻视频帧显示时间戳差值的平均值，单位为毫秒。
	VideoFrameGap int32 `json:"VideoFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个视频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	VideoPts int32 `json:"VideoPts"`
}

type DescribeLiveBatchPushStreamMetricsBody struct {

	// REQUIRED; 推流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	// :::tip 单次查询最大时间跨度为 1 天，历史查询最大时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 5：5 秒；
	// * 30：30 秒；
	// * 60：（默认值）1 分钟。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 流名称。 :::tip 使用 Stream 构造请求时，需同时定义 App 参数，不可缺省。 :::
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveBatchPushStreamMetricsRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveBatchPushStreamMetricsResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveBatchPushStreamMetricsResResult `json:"Result"`
}

type DescribeLiveBatchPushStreamMetricsResResponseMetadata struct {

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

type DescribeLiveBatchPushStreamMetricsResResult struct {

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 5：5 秒；
	// * 30：30 秒；
	// * 60：1 分钟。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 推流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; ��指定时间粒度聚合的监控数据。
	StreamMetricList []DescribeLiveBatchPushStreamMetricsResResultStreamMetricListItem `json:"StreamMetricList"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveBatchPushStreamMetricsResResultStreamMetricListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 推流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 按指定时间粒度聚合的监控数据。
	MetricList []DescribeLiveBatchPushStreamMetricsResResultStreamMetricListPropertiesItemsItem `json:"MetricList"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`
}

type DescribeLiveBatchPushStreamMetricsResResultStreamMetricListPropertiesItemsItem struct {

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

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻视频帧显示时间戳差值的最大值，单位为毫秒。
	VideoFrameGap int32 `json:"VideoFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个视频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	VideoPts int32 `json:"VideoPts"`
}

type DescribeLiveBatchSourceStreamAvgMetricsBody struct {

	// REQUIRED; 拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	// :::tip 单次查询最大时间跨度为 1 天，历史查询最大时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 30：30 秒；
	// * 60：（默认值）1 分钟。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveBatchSourceStreamAvgMetricsRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveBatchSourceStreamAvgMetricsResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveBatchSourceStreamAvgMetricsResResult `json:"Result"`
}

type DescribeLiveBatchSourceStreamAvgMetricsResResponseMetadata struct {

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

type DescribeLiveBatchSourceStreamAvgMetricsResResult struct {

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 30：30 秒；
	// * 60：1 分钟。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 流的信息，包含域名、应用名称、流名称和监控数据。
	StreamMetricList []DescribeLiveBatchSourceStreamAvgMetricsResResultStreamMetricListItem `json:"StreamMetricList"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveBatchSourceStreamAvgMetricsResResultStreamMetricListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 按指定时间粒度聚合的监控数据。
	MetricList []DescribeLiveBatchSourceStreamAvgMetricsResResultStreamMetricListPropertiesItemsItem `json:"MetricList"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`
}

type DescribeLiveBatchSourceStreamAvgMetricsResResultStreamMetricListPropertiesItemsItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的音频码率平均值，单位为 kbps。
	AudioBitrate float32 `json:"AudioBitrate"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻音频帧显示时间戳差值的平均值，单位为毫秒。
	AudioFrameGap int32 `json:"AudioFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内的音频帧率平均值，单位为 fps。
	AudioFramerate float32 `json:"AudioFramerate"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个音频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	AudioPts int32 `json:"AudioPts"`

	// REQUIRED; 当前数据聚合时间粒度内的视频码率平均值，单位为 kbps。
	Bitrate float32 `json:"Bitrate"`

	// REQUIRED; 当前数据聚合时间粒度内的视频帧率平均值，单位为 fps。
	Framerate float32 `json:"Framerate"`

	// REQUIRED; 当前数据聚合时间粒度内，所有音视频帧显示时间戳差值的平均值，即所有 AudioPts 与 VideoPts 差值的平均值，单位为毫秒。
	PtsDelta int32 `json:"PtsDelta"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻视频帧显示时间戳差值的平均值，单位为毫秒。
	VideoFrameGap int32 `json:"VideoFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个视频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	VideoPts int32 `json:"VideoPts"`
}

type DescribeLiveBatchSourceStreamMetricsBody struct {

	// REQUIRED; 拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	// :::tip 单次查询最大时间跨度为 1 天，历史查询最大时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 30：30 秒；
	// * 60：（默认值）1 分钟。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveBatchSourceStreamMetricsRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveBatchSourceStreamMetricsResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveBatchSourceStreamMetricsResResult `json:"Result"`
}

type DescribeLiveBatchSourceStreamMetricsResResponseMetadata struct {

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

type DescribeLiveBatchSourceStreamMetricsResResult struct {

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 30：30 秒；
	// * 60：1 分钟。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 流的监控数据。
	StreamMetricList []DescribeLiveBatchSourceStreamMetricsResResultStreamMetricListItem `json:"StreamMetricList"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveBatchSourceStreamMetricsResResultStreamMetricListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 按指定时间粒度聚合的监控数据。
	MetricList []DescribeLiveBatchSourceStreamMetricsResResultStreamMetricListPropertiesItemsItem `json:"MetricList"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`
}

type DescribeLiveBatchSourceStreamMetricsResResultStreamMetricListPropertiesItemsItem struct {

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

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻视频帧显示时间戳差值的最大值，单位为毫秒。
	VideoFrameGap int32 `json:"VideoFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个视频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	VideoPts int32 `json:"VideoPts"`
}

type DescribeLiveBatchStreamTrafficDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	// :::tip 查询历史数据的时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 域名列表，缺省情况下表示当前账号下的所有推拉流域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 查询数据的页码，默认值为 1，表示查询第一页的数据。
	PageNum *int32 `json:"PageNum,omitempty"`

	// 每页显示的数据条数，默认值为 1000，取值范围为 100～1000。
	PageSize *int32 `json:"PageSize,omitempty"`

	// 推拉流协议，缺省情况下表示所有协议类型，支持的协议如下所示。
	// * HTTP-FLV：基于 HTTP 协议的推拉流协议，使用 FLV 格式传输视频格式。
	// * HTTP-HLS：基于 HTTP 协议的推拉流协议，使用 TS 格式传输视频格式。
	// * RTMP：Real Time Message Protocol，实时信息传输协议。
	// * RTM：Real Time Media，超低延时直播协议。
	// * SRT：Secure Reliable Transport，安全可靠传输协议。
	// * QUIC：Quick UDP Internet Connections，一种基于 UDP 的全新的低延时互联网传输协议。
	// :::tip
	// * 如果查询推拉流协议为 QUIC，不能同时查询其他协议。
	// * 缺省情况下，查询的总流量数据为实际产生的上下行流量。
	// * 如果传入单个协议进行查询，并对各协议的流量求和，结果将大于实际总流量。
	ProtocolList []*string `json:"ProtocolList,omitempty"`
}

type DescribeLiveBatchStreamTrafficDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveBatchStreamTrafficDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveBatchStreamTrafficDataResResult `json:"Result"`
}

type DescribeLiveBatchStreamTrafficDataResResponseMetadata struct {

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

type DescribeLiveBatchStreamTrafficDataResResult struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 数据分页的信息。
	Pagination DescribeLiveBatchStreamTrafficDataResResultPagination `json:"Pagination"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 流维度的流量用量信息详情。
	StreamInfoList []DescribeLiveBatchStreamTrafficDataResResultStreamInfoListItem `json:"StreamInfoList"`

	// REQUIRED; 当前查询条件下，所有流的下行总流量，单位为 GB。
	TotalDownTraffic float32 `json:"TotalDownTraffic"`

	// REQUIRED; 当前查询条件下，所有流的上行总流量，单位为 GB。
	TotalUpTraffic float32 `json:"TotalUpTraffic"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 推拉流协议，协议说明如下。
	// * HTTP-FLV：基于 HTTP 协议的推拉流协议，使用 FLV 格式传输视频格式。
	// * HTTP-HLS：基于 HTTP 协议的推拉流协议，使用 TS 格式传输视频格式。
	// * RTMP：Real Time Message Protocol，实时信息传输协议。
	// * RTM：Real Time Media，超低延时直播协议。
	// * SRT：Secure Reliable Transport，安全可靠传输协议。
	// * QUIC：Quick UDP Internet Connections，一种基于 UDP 的全新的低延时互联网传输协议。
	ProtocolList []*string `json:"ProtocolList,omitempty"`
}

// DescribeLiveBatchStreamTrafficDataResResultPagination - 数据分页的信息。
type DescribeLiveBatchStreamTrafficDataResResultPagination struct {

	// REQUIRED; 当前所在分页的页码。
	PageNum int32 `json:"PageNum"`

	// REQUIRED; 每页显示的数据条数。
	PageSize int32 `json:"PageSize"`

	// REQUIRED; 查询结果的数据总条数。
	TotalCount int32 `json:"TotalCount"`
}

type DescribeLiveBatchStreamTrafficDataResResultStreamInfoListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 域名。
	Domain string `json:"Domain"`

	// REQUIRED; 当前流的下行流量，单位为 GB。
	DownTraffic float32 `json:"DownTraffic"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 当前流的上行流量，单位为 GB。
	UpTraffic float32 `json:"UpTraffic"`
}

type DescribeLiveBatchStreamTranscodeDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。 :::tip 查询历史数据的时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 域名列表，缺省情况下表示当前账号下的所有推拉流域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 查询数据的页码，默认值为 1，表示查询第一页的数据。
	PageNum *int32 `json:"PageNum,omitempty"`

	// 每页显示的数据条数，默认值为 1000，取值范围为 [100,1000]。
	PageSize *int32 `json:"PageSize,omitempty"`
}

type DescribeLiveBatchStreamTranscodeDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveBatchStreamTranscodeDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveBatchStreamTranscodeDataResResult `json:"Result"`
}

type DescribeLiveBatchStreamTranscodeDataResResponseMetadata struct {

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

type DescribeLiveBatchStreamTranscodeDataResResult struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 数据分页的信息。
	Pagination DescribeLiveBatchStreamTranscodeDataResResultPagination `json:"Pagination"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 流维度的转码用量信息详情。
	StreamInfoList []DescribeLiveBatchStreamTranscodeDataResResultStreamInfoListItem `json:"StreamInfoList"`

	// REQUIRED; 当前查询条件下，所有流的转码总时长，单位为分钟。
	TotalDuration float32 `json:"TotalDuration"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`
}

// DescribeLiveBatchStreamTranscodeDataResResultPagination - 数据分页的信息。
type DescribeLiveBatchStreamTranscodeDataResResultPagination struct {

	// REQUIRED; 当前所在分页的页码。
	PageNum int32 `json:"PageNum"`

	// REQUIRED; 每页展示的数据条数。
	PageSize int32 `json:"PageSize"`

	// REQUIRED; 查询结果的数据总条数。
	TotalCount int32 `json:"TotalCount"`
}

type DescribeLiveBatchStreamTranscodeDataResResultStreamInfoListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 当前流的转码码率。
	Coderate int32 `json:"Coderate"`

	// REQUIRED; 域名。
	Domain string `json:"Domain"`

	// REQUIRED; 当前流在查询时间内的转码总时长，单位为分钟。
	Duration float32 `json:"Duration"`

	// REQUIRED; 分辨率。- 480P：640 × 480； - 720P：1280 × 720； - 1080P：1920 × 1088； - 2K：2560 × 1440； - 4K：4096 × 2160；- 8K：大于4K； -
	// 0：纯音频流；
	Resolution string `json:"Resolution"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 视频编码格式，支持的取值和含义如下所示。- NormalH264：H.264 标准转码； - NormalH265：H.265 标准转码； - NormalH266：H.266 标准转码； - ByteHDH264：H.264
	// 极智超清； - ByteHDH265：H.265 极智超清； - ByteHDH266：H.266 极智超清；- ByteQE：画质增强；- Audio：纯音频流；
	VCodec string `json:"VCodec"`
}

type DescribeLiveCustomizedLogDataBody struct {

	// REQUIRED; The end time of the query, in UTC time in RFC 3339 format, with a precision of seconds.
	EndTime string `json:"EndTime"`

	// REQUIRED; The start time of the query, in UTC time in RFC 3339 format, with a precision of seconds. ::: tip Currently,
	// only query log data for the last 31 days is supported. :::
	StartTime string `json:"StartTime"`

	// REQUIRED; For log type, please contact technical support for parameter values.
	Type string `json:"Type"`

	// Domain name list, which by default represents all streaming and pulling domain names for the current user. ::: tips This
	// parameter is invalid when the log type is pull-stream-forward log (Type =
	// relay). :::
	DomainList []*string `json:"DomainList,omitempty"`

	// The page number of the query data, the default is 1, indicating the data on the first page of the query.
	PageNum *int32 `json:"PageNum,omitempty"`

	// The number of data bars displayed per page, the default is 20, and the maximum value is 1000.
	PageSize *int32 `json:"PageSize,omitempty"`
}

type DescribeLiveCustomizedLogDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveCustomizedLogDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveCustomizedLogDataResResult `json:"Result"`
}

type DescribeLiveCustomizedLogDataResResponseMetadata struct {

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

type DescribeLiveCustomizedLogDataResResult struct {

	// REQUIRED; List of domain names.
	DomainList []string `json:"DomainList"`

	// REQUIRED; The end time of the query, in UTC time in RFC 3339 format, with a precision of seconds.
	EndTime string `json:"EndTime"`

	// REQUIRED; List of information for the log file.
	LogInfoList []DescribeLiveCustomizedLogDataResResultLogInfoListItem `json:"LogInfoList"`

	// REQUIRED; Data paging information.
	Pagination DescribeLiveCustomizedLogDataResResultPagination `json:"Pagination"`

	// REQUIRED; The start time of the query, in UTC time in RFC 3339 format, with a precision of seconds.
	StartTime string `json:"StartTime"`

	// REQUIRED; The log type of the query.
	Type string `json:"Type"`
}

type DescribeLiveCustomizedLogDataResResultLogInfoListItem struct {

	// The hour interval corresponding to the log file, UTC time in RFC 3339 format, with a precision of seconds.
	DateTime *string `json:"DateTime,omitempty"`

	// Domain name. ::: tips This field is empty when querying the pull stream retweet log (Type = relay). :::
	Domain *string `json:"Domain,omitempty"`

	// Log file download link.
	DownloadURL *string `json:"DownloadUrl,omitempty"`

	// Log file name.
	LogName *string `json:"LogName,omitempty"`

	// Log file size in bytes.
	LogSize *int32 `json:"LogSize,omitempty"`
}

// DescribeLiveCustomizedLogDataResResultPagination - Data paging information.
type DescribeLiveCustomizedLogDataResResultPagination struct {

	// REQUIRED; The page number of the current pagination.
	PageNum int32 `json:"PageNum"`

	// REQUIRED; The number of data bars displayed per page.
	PageSize int32 `json:"PageSize"`

	// REQUIRED; The total number of pieces of data in the query result.
	TotalCount int32 `json:"TotalCount"`
}

type DescribeLiveFeeConfigRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveFeeConfigResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeLiveFeeConfigResResult `json:"Result,omitempty"`
}

type DescribeLiveFeeConfigResResponseMetadata struct {

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

// DescribeLiveFeeConfigResResult - 视请求的接口而定
type DescribeLiveFeeConfigResResult struct {

	// 账号id
	AccountID *string `json:"AccountId,omitempty"`

	// 进制
	Base *int32 `json:"Base,omitempty"`

	// 配置创建者
	Creator *string `json:"Creator,omitempty"`

	// 上浮系数
	Factor *float32 `json:"Factor,omitempty"`

	// 免流配置
	FreeFeeList []*string `json:"FreeFeeList,omitempty"`

	// 配置id
	ID *int32 `json:"Id,omitempty"`

	// 闲忙时开关
	StageEnable *string `json:"StageEnable,omitempty"`

	// 闲忙时生效时间
	StageTime *string `json:"StageTime,omitempty"`

	// 上浮系数生效时间
	StartTime *string `json:"StartTime,omitempty"`
}

type DescribeLiveFreeTimeIntervalRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveFreeTimeIntervalResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeLiveFreeTimeIntervalResResult `json:"Result,omitempty"`
}

type DescribeLiveFreeTimeIntervalResResponseMetadata struct {

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

// DescribeLiveFreeTimeIntervalResResult - 视请求的接口而定
type DescribeLiveFreeTimeIntervalResResult struct {

	// REQUIRED; 闲时时间段
	FreeTime string `json:"FreeTime"`
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

	// REQUIRED; The end time of the query, in UTC time in RFC 3339 format, with a precision of seconds.
	EndTime string `json:"EndTime"`

	// REQUIRED; The start time of the query, in UTC time in RFC 3339 format, with a precision of seconds. ::: tips Currently,
	// only query log data for the last 31 days is supported. :::
	StartTime string `json:"StartTime"`

	// REQUIRED; Log types, supported types are shown below.
	// * pull: pull stream log
	// * push: push stream log
	// * source: back to source log
	// * relay: pull stream retweet log
	Type string `json:"Type"`

	// Domain name list, which by default represents all streaming and pulling domain names for the current user. ::: tips This
	// parameter is invalid when the log type is pull-stream-forward log (Type =
	// relay). :::
	DomainList []*string `json:"DomainList,omitempty"`

	// The page number of the query data, the default is 1, indicating the data on the first page of the query.
	PageNum *int32 `json:"PageNum,omitempty"`

	// The number of data bars displayed per page, the default is 20, and the maximum value is 1000.
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

	// REQUIRED; The end time of the query, in UTC time in RFC 3339 format, with a precision of seconds.
	EndTime string `json:"EndTime"`

	// REQUIRED; List of information for the log file.
	LogInfoList []DescribeLiveLogDataResResultLogInfoListItem `json:"LogInfoList"`

	// REQUIRED; Data paging information.
	Pagination DescribeLiveLogDataResResultPagination `json:"Pagination"`

	// REQUIRED; The start time of the query, in UTC time in RFC 3339 format, with a precision of seconds.
	StartTime string `json:"StartTime"`

	// REQUIRED; Log type, type description is shown below.
	// * pull: pull stream log
	// * push: push stream log
	// * source: back to source log
	// * relay: pull stream retweet log
	Type string `json:"Type"`

	// List of domain names.
	DomainList []*string `json:"DomainList,omitempty"`
}

type DescribeLiveLogDataResResultLogInfoListItem struct {

	// REQUIRED; The hour interval corresponding to the log file, UTC time in RFC 3339 format, with a precision of seconds.
	DateTime string `json:"DateTime"`

	// REQUIRED; Log file download link.
	DownloadURL string `json:"DownloadUrl"`

	// REQUIRED; Log file names, log file naming conventions are as follows.
	// * In relation to domain names: Accelerated domain names start _ _ _ _ end _ file split serial number. For example, 'www.example.com202308110000000100000';
	// * When not related to the domain name: the file splitting serial number _ beginning _ end of time _ _ . For example, '202308110000000100000
	// .gz';
	// * If the current event generates more than 1.50 million logs within a certain hour, multiple log files will be generated,
	// and the order of log files will be marked with the serial number at the end
	// of the file name, for example, '202308110000000100000 .gz', '202308110000000100001 .gz'.
	LogName string `json:"LogName"`

	// REQUIRED; Log file size in bytes.
	LogSize int32 `json:"LogSize"`

	// Domain name. ::: tips This field is empty when querying the pull stream retweet log (Type = relay). :::
	Domain *string `json:"Domain,omitempty"`
}

// DescribeLiveLogDataResResultPagination - Data paging information.
type DescribeLiveLogDataResResultPagination struct {

	// REQUIRED; The page number of the current pagination.
	PageNum int32 `json:"PageNum"`

	// REQUIRED; The number of data bars displayed per page.
	PageSize int32 `json:"PageSize"`

	// REQUIRED; The total number of pieces of data in the query result.
	TotalCount int32 `json:"TotalCount"`
}

type DescribeLiveMetricBandwidthDataBody struct {

	// REQUIRED; The end time of the query, in UTC time in RFC 3339 format, with a precision of seconds.
	EndTime string `json:"EndTime"`

	// REQUIRED; The start time of the query, in UTC time in RFC 3339 format, with a precision of seconds.
	StartTime string `json:"StartTime"`

	// The time granularity of the aggregation, in seconds, is supported as follows.
	// * 60:1 minute. When the time granularity is 1 minute, the maximum time span for a single query is 24 hours, and the historical
	// query time range is 366 days;
	// * 300: (default) 5 minutes. When the time granularity is 5 minutes, the maximum time span for a single query is 31 days,
	// and the historical query time range is 366 days;
	// * 3600:1 hour. When the time granularity is 1 hour, the maximum time span for a single query is 93 days, and the time range
	// for historical queries is 366 days.
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// The name of the application when querying stream granular data. ::: tips When using'App 'to construct a request, you need
	// to define'Stream' parameters at the same time, which cannot be defaulted. :::
	App *string `json:"App,omitempty"`

	// Dimensions for data splitting. Data splitting is not performed by default. The supported dimensions are as follows.
	// * Domain: domain name;
	// * Protocol: up & down streaming protocol;
	// * IP: the IP address of the export extranet;
	// * ISP: Operator.
	// ::: tips When configuring a data split dimension, the corresponding dimension parameter will return the data split by dimension
	// when multiple values are passed; the corresponding dimension will not
	// return the data split by dimension when only one value is passed. :::
	DetailField []*string `json:"DetailField,omitempty"`

	// Domain name list, which by default represents all push-and-pull basin names for the current user.
	DomainList []*string `json:"DomainList,omitempty"`

	// The operator identifier that provides the network access service, which by default represents all operators. The supported
	// operators are shown below.
	// * unicom: unicom;
	// * railcom;
	// * telecom: telecommunications;
	// * mobile: mobile;
	// * cernet: education network;
	// * tianwei: tianwei;
	// * alibaba: alibaba;
	// * tenger: Tencent;
	// * drpeng: Dr. Peng;
	// * btvn: radio and television;
	// * huashu: wah number;
	// * other: other.
	// You can also obtain the operator's corresponding identifier through the DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974]
	// interface.
	ISPList []*string `json:"ISPList,omitempty"`

	// The up & down streaming protocol, representing all protocol types by default, supports the following protocols.
	// HTTP-FLV: Up & down streaming protocol based on HTTP protocol, using FLV format to transmit video formats. HTTP-HLS: Up
	// & down streaming protocol based on HTTP protocol, using TS format to transmit
	// video formats. RTMP: Real Time Message Protocol, Real Time Message Protocol. RTM: Real Time Media, ultra-low latency live
	// streaming protocol.
	// * SRT: Secure Reliable Transport, Secure Reliable Transport. QUIC: Quick UDP Internet Connections, a new low-latency Internet
	// transport protocol based on UDP.
	// ::: tips If the query up & down streaming protocol is QUIC, you cannot simultaneously query other protocols. :::
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// List of regions to which the CDN node IP belongs, representing all regions by default. ::: tips Parameters'RegionList 'and'UserRegionList'
	// do not support simultaneous passing in. :::
	RegionList []*DescribeLiveMetricBandwidthDataBodyRegionListItem `json:"RegionList,omitempty"`

	// The stream name parameter when querying stream granular data. ::: tips When using'Stream 'to construct a request, you need
	// to define'App' parameters at the same time, which cannot be defaulted. :::
	Stream *string `json:"Stream,omitempty"`

	// List of regions to which the client side IP belongs, representing all regions by default. ::: tips Parameters'RegionList
	// 'and'UserRegionList' do not support simultaneous passing in. :::
	UserRegionList []*DescribeLiveMetricBandwidthDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveMetricBandwidthDataBodyRegionListItem struct {

	// For the region identifier in the region information, see Query Region Identifier (https://www.volcengine.com/docs/6469/1133973).
	Area *string `json:"Area,omitempty"`

	// The country identifier in the region information, see Query Region Identifier [https://www.volcengine.com/docs/6469/1133973].
	// If filtering by country, you need to pass in both Area and Country.
	Country *string `json:"Country,omitempty"`

	// The province identifier in the regional information. This parameter is not supported abroad for the time being. Please
	// refer to Query Area Identifier [https://www.volcengine.com/docs/6469/1133973] for
	// how to obtain it. If filtering by province, you need to pass in'Area ',' Country 'and'Province' at the same time.
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveMetricBandwidthDataBodyUserRegionListItem struct {

	// Region, mapping relationship, see Region mapping
	Area *string `json:"Area,omitempty"`

	// Country, mapping relationship See Region mapping. If you want to filter by country, you need to pass in both Area and Country.
	Country *string `json:"Country,omitempty"`

	// Domestic is a province, and foreign countries do not support this parameter for the time being. For the mapping relationship,
	// please refer to Area Mapping. If filtering by province, you need to pass
	// in Area, Country and Province at the same time.
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

	// REQUIRED; The time granularity of the aggregation, in seconds.
	// * 60:1 minute;
	// * 300:5 minutes;
	// * 3600:1 hour.
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; All time-granular data.
	BandwidthDataList []DescribeLiveMetricBandwidthDataResResultBandwidthDataListItem `json:"BandwidthDataList"`

	// REQUIRED; The end time of the query, in UTC time in RFC 3339 format, with a precision of seconds.
	EndTime string `json:"EndTime"`

	// REQUIRED; The downlink peak in the query time range, expressed in Mbps.
	PeakDownBandwidth float32 `json:"PeakDownBandwidth"`

	// REQUIRED; The upstream peak in the query time range, expressed in Mbps.
	PeakUpBandwidth float32 `json:"PeakUpBandwidth"`

	// REQUIRED; The start time of the query, in UTC time in RFC 3339 format, with a precision of seconds.
	StartTime string `json:"StartTime"`

	// The name of the application when querying stream granular data.
	App *string `json:"App,omitempty"`

	// Data split by dimension.
	BandwidthDetailDataList []*DescribeLiveMetricBandwidthDataResResultBandwidthDetailDataListItem `json:"BandwidthDetailDataList,omitempty"`

	// The dimension of the data split, the dimension description is shown below.
	// * Domain: domain name;
	// * Protocol: up & down streaming protocol;
	// * IP: the IP address of the export extranet;
	// * ISP: Operator.
	DetailField []*string `json:"DetailField,omitempty"`

	// List of domain names.
	DomainList []*string `json:"DomainList,omitempty"`

	// The operator identifier that provides network access services, and the corresponding relationship between the identifier
	// and the operator is as follows.
	// * unicom: unicom;
	// * railcom;
	// * telecom: telecommunications;
	// * mobile: mobile;
	// * cernet: education network;
	// * tianwei: tianwei;
	// * alibaba: alibaba;
	// * tenger: Tencent;
	// * drpeng: Dr. Peng;
	// * btvn: radio and television;
	// * huashu: wah number;
	// * other: other.
	ISPList []*string `json:"ISPList,omitempty"`

	// Up & down streaming protocol, the protocol description is as follows. HTTP-FLV: Up & down streaming protocol based on HTTP
	// protocol, using FLV format to transmit video formats. HTTP-HLS: Up & down
	// streaming protocol based on HTTP protocol, using TS format to transmit video formats. RTMP: Real Time Message Protocol,
	// Real Time Message Protocol. RTM: Real Time Media, ultra-low latency live
	// streaming protocol.
	// * SRT: Secure Reliable Transport, Secure Reliable Transport. QUIC: Quick UDP Internet Connections, a new low-latency Internet
	// transport protocol based on UDP.
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN node IP region list.
	RegionList []*DescribeLiveMetricBandwidthDataResResultRegionListItem `json:"RegionList,omitempty"`

	// The name of the stream when querying stream granular data.
	Stream *string `json:"Stream,omitempty"`

	// List of client side IP regions.
	UserRegionList []*DescribeLiveMetricBandwidthDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveMetricBandwidthDataResResultBandwidthDataListItem struct {

	// REQUIRED; The downlink peak bandwidth within the current data aggregation time granularity, in Mbps.
	DownBandwidth float32 `json:"DownBandwidth"`

	// REQUIRED; When data is aggregated by time granularity, the start time of each time granularity, in UTC time in RFC 3339
	// format, with a precision of seconds.
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; The upstream peak bandwidth within the current data aggregation time granularity, in Mbps.
	UpBandwidth float32 `json:"UpBandwidth"`
}

type DescribeLiveMetricBandwidthDataResResultBandwidthDetailDataListItem struct {

	// REQUIRED; After splitting the data by dimension, the data of all time granularities in the current dimension.
	BandwidthDataList []DescribeLiveMetricBandwidthDataResResultBandwidthDetailDataListPropertiesItemsItem `json:"BandwidthDataList"`

	// REQUIRED; After splitting the data by dimension, the downlink peak bandwidth of the current dimension is expressed in Mbps.
	PeakDownBandwidth float32 `json:"PeakDownBandwidth"`

	// REQUIRED; After splitting the data by dimension, the upstream peak bandwidth of the current dimension is expressed in Mbps.
	PeakUpBandwidth float32 `json:"PeakUpBandwidth"`

	// Domain name information when data is split by domain name dimension.
	Domain *string `json:"Domain,omitempty"`

	// Operator information when data is split according to the operator dimension.
	ISP *string `json:"ISP,omitempty"`

	// Protocol information when splitting data by up & down streaming protocol dimension.
	Protocol *string `json:"Protocol,omitempty"`
}

type DescribeLiveMetricBandwidthDataResResultBandwidthDetailDataListPropertiesItemsItem struct {

	// REQUIRED; Downlink bandwidth in Mbps
	DownBandwidth float32 `json:"DownBandwidth"`

	// REQUIRED; Time slice start time. UTC time in RFC3339 format with precision s, for example, 2022-04-13T00:00:00 + 08:00
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; Upstream bandwidth in Mbps
	UpBandwidth float32 `json:"UpBandwidth"`
}

type DescribeLiveMetricBandwidthDataResResultRegionListItem struct {

	// The region identifier in the region information.
	Area *string `json:"Area,omitempty"`

	// The country identifier in the regional information.
	Country *string `json:"Country,omitempty"`

	// The identity identifier in the zone information.
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveMetricBandwidthDataResResultUserRegionListItem struct {
	Area     *string `json:"Area,omitempty"`
	Country  *string `json:"Country,omitempty"`
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveMetricTrafficDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 60：1 分钟。时间粒度为 1 分钟时，单次查询最大时间跨度为 24 小时，历史查询时间范围为 366 天；
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，缺省情况下不进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * Protocol：推拉流协议；
	// * ISP：运营商。
	// :::tip 配置数据拆分维度时，对应的维度参数需传入多个值时会返回按维度进行拆分的数据；对应的维度只传入一个值时不返回按维度进行拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，缺省情况表示当前用户的所有推拉流域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，缺省情况下表示所有运营商，支持的运营商如下所示。
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
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// 推拉流协议，缺省情况下表示所有协议类型，支持的协议如下所示。
	// * HTTP-FLV：基于 HTTP 协议的推拉流协议，使用 FLV 格式传输视频格式。
	// * HTTP-HLS：基于 HTTP 协议的推拉流协议，使用 TS 格式传输视频格式。
	// * RTMP：Real Time Message Protocol，实时信息传输协议。
	// * RTM：Real Time Media，超低延时直播协议。
	// * SRT：Secure Reliable Transport，安全可靠传输协议。
	// * QUIC：Quick UDP Internet Connections，一种基于 UDP 的全新的低延时互联网传输协议。
	// :::tip 如果查询推拉流协议为 QUIC，不能同时查询其他协议。 :::
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	RegionList []*DescribeLiveMetricTrafficDataBodyRegionListItem `json:"RegionList,omitempty"`

	// 指定查询的流量数据为闲时或忙时，缺省情况下为查询全部数据，支持的取值如下。
	// * busy：忙时；
	// * free：闲时。
	Stage *string `json:"Stage,omitempty"`

	// 流名称。 :::tip 使用 Stream 构造请求时，需同时定义 App 参数，不可缺省。 :::
	Stream *string `json:"Stream,omitempty"`

	// 客户端 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	UserRegionList []*DescribeLiveMetricTrafficDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveMetricTrafficDataBodyRegionListItem struct {

	// 区域信息中的大区标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按国家筛选，需要同时传入Area和Country。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符，国外暂不支持该参数，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按省筛选，需要同时传入Area、Country和Province。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveMetricTrafficDataBodyUserRegionListItem struct {

	// 大区，映射关系请参见区域映射
	Area *string `json:"Area,omitempty"`

	// 国家，映射关系请参见区域映射。如果按国家筛选，需要同时传入 Area 和 Country。
	Country *string `json:"Country,omitempty"`

	// 国内为省，国外暂不支持该参数，映射关系请参见区域映射。如果按省筛选，需要同时传入 Area、Country 和 Province。
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

	// REQUIRED; 聚合的时间粒度，单位为秒。
	// * 60：1 分钟；
	// * 300：5 分钟；
	// * 3600：1 小时。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 查询时间范围内的下行总流量，单位为 GB。
	TotalDownTraffic float32 `json:"TotalDownTraffic"`

	// REQUIRED; 查询时间范围内的上行总流量，单位为 GB。
	TotalUpTraffic float32 `json:"TotalUpTraffic"`

	// REQUIRED; 所有时间粒度的数据。
	TrafficDataList []DescribeLiveMetricTrafficDataResResultTrafficDataListItem `json:"TrafficDataList"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，维度说明如下所示。
	// * Domain：域名；
	// * Protocol：推拉流协议；
	// * ISP：运营商。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，标识符与运营商的对应关系如下。
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
	ISPList []*string `json:"ISPList,omitempty"`

	// 推拉流协议，协议说明如下。
	// * HTTP-FLV：基于 HTTP 协议的推拉流协议，使用 FLV 格式传输视频格式。
	// * HTTP-HLS：基于 HTTP 协议的推拉流协议，使用 TS 格式传输视频格式。
	// * RTMP：Real Time Message Protocol，实时信息传输协议。
	// * RTM：Real Time Media，超低延时直播协议。
	// * SRT：Secure Reliable Transport，安全可靠传输协议。
	// * QUIC：Quick UDP Internet Connections，一种基于 UDP 的全新的低延时互联网传输协议。
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域列表。
	RegionList []*DescribeLiveMetricTrafficDataResResultRegionListItem `json:"RegionList,omitempty"`

	// 流量数据为闲时或忙时，取值说明如下。
	// * busy：忙时；
	// * free：闲时。
	Stage *string `json:"Stage,omitempty"`

	// 流名称。
	Stream *string `json:"Stream,omitempty"`

	// 按维度拆分后的数据。 :::tip 请求时，DomainList、ProtocolList和ISPList至少有一个参数传入了多个值时，会返回该参数；否则不返回该参数。 :::
	TrafficDetailDataList []*DescribeLiveMetricTrafficDataResResultTrafficDetailDataListItem `json:"TrafficDetailDataList,omitempty"`

	// 客户端 IP 所属区域列表。
	UserRegionList []*DescribeLiveMetricTrafficDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveMetricTrafficDataResResultRegionListItem struct {

	// 区域信息中的大区标识符。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveMetricTrafficDataResResultTrafficDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内产生的总下行流量，单位 GB。
	DownTraffic float32 `json:"DownTraffic"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
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

	// REQUIRED; 时间片起始时刻。RFC3339 格式的 UTC 时间，精度为 s，例如，2022-04-13T00:00:00+08:00
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 上行流量，单位 GB
	UpTraffic float32 `json:"UpTraffic"`
}

type DescribeLiveMetricTrafficDataResResultUserRegionListItem struct {
	Area     *string `json:"Area,omitempty"`
	Country  *string `json:"Country,omitempty"`
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveP95PeakBandwidthDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。 :::tip 单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 300：（默认值）5 分钟。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 域名列表，缺省情况下表示当前用户的所有推拉流域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 推拉流协议，缺省情况下表示所有协议类型，支持的协议如下所示。
	// * HTTP-FLV：基于 HTTP 协议的推拉流协议，使用 FLV 格式传输视频格式。
	// * HTTP-HLS：基于 HTTP 协议的推拉流协议，使用 TS 格式传输视频格式。
	// * RTMP：Real Time Message Protocol，实时信息传输协议。
	// * RTM：Real Time Media，超低延时直播协议。
	// * SRT：Secure Reliable Transport，安全可靠传输协议。
	// * QUIC：Quick UDP Internet Connections，一种基于 UDP 的全新的低延时互联网传输协议。
	// :::tip 如果查询推拉流协议为 QUIC，不能同时查询其他协议。 :::
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	RegionList []*DescribeLiveP95PeakBandwidthDataBodyRegionListItem `json:"RegionList,omitempty"`

	// 客户端 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	UserRegionList []*DescribeLiveP95PeakBandwidthDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveP95PeakBandwidthDataBodyRegionListItem struct {

	// 区域信息中的大区标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按国家筛选，需要同时传入Area和Country。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按省筛选，需要同时传入Area、Country和Province。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveP95PeakBandwidthDataBodyUserRegionListItem struct {

	// 大区，映射关系请参见区域映射
	Area *string `json:"Area,omitempty"`

	// 国家，映射关系请参见区域映射。如果按国家筛选，需要同时传入 Area 和 Country。
	Country *string `json:"Country,omitempty"`

	// 国内为省，国外暂不支持该参数，映射关系请参见区域映射。如果按省筛选，需要同时传入 Area、Country 和 Province。
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

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 300：5 分钟。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 时间范围内的上下行 95 峰值带宽总和。 :::tip 如果请求时，Regionlist中传入多个 region，则返回这些 region 的上下行带宽 95 峰值总和。 :::
	P95PeakBandwidth float32 `json:"P95PeakBandwidth"`

	// REQUIRED; 95 峰值带宽的时间戳，RFC3339 格式的 UTC 时间，精度为秒。
	P95PeakTimestamp string `json:"P95PeakTimestamp"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 推拉流协议，协议说明如下。
	// * HTTP-FLV：基于 HTTP 协议的推拉流协议，使用 FLV 格式传输视频格式。
	// * HTTP-HLS：基于 HTTP 协议的推拉流协议，使用 TS 格式传输视频格式。
	// * RTMP：Real Time Message Protocol，实时信息传输协议。
	// * RTM：Real Time Media，超低延时直播协议。
	// * SRT：Secure Reliable Transport，安全可靠传输协议。
	// * QUIC：Quick UDP Internet Connections，一种基于 UDP 的全新的低延时互联网传输协议。
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域的列表。
	RegionList []*DescribeLiveP95PeakBandwidthDataResResultRegionListItem `json:"RegionList,omitempty"`

	// 客户端 IP 所属区域的列表。
	UserRegionList []*DescribeLiveP95PeakBandwidthDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveP95PeakBandwidthDataResResultRegionListItem struct {

	// 区域信息中的大区标识符。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveP95PeakBandwidthDataResResultUserRegionListItem struct {
	Area     *string `json:"Area,omitempty"`
	Country  *string `json:"Country,omitempty"`
	Province *string `json:"Province,omitempty"`
}

type DescribeLivePlayStatusCodeDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 60：（默认值）1 分钟。时间粒度为 1 分钟时，单次查询最大时间跨度为 24 小时，历史查询时间范围为 366 天；
	// * 300：5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 数据拆分的维度，缺省情况下不进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * ISP：运营商。
	// :::tip 配置数据拆分维度时，对应的维度参数需传入多个值时会返回按维度进行拆分的数据；对应的维度只传入一个值时不返回按维度进行拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，缺省情况下表示当前用户的所有推拉流域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，缺省情况下表示所有运营商，支持的运营商如下所示。
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
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// CDN 节点 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	RegionList []*DescribeLivePlayStatusCodeDataBodyRegionListItem `json:"RegionList,omitempty"`

	// 客户端 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	UserRegionList []*DescribeLivePlayStatusCodeDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLivePlayStatusCodeDataBodyRegionListItem struct {

	// 区域信息中的大区标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按国家筛选，需要同时传入Area和Country。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符，国外暂不支持该参数，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按省筛选，需要同时传入Area、Country和Province。
	Province *string `json:"Province,omitempty"`
}

type DescribeLivePlayStatusCodeDataBodyUserRegionListItem struct {

	// 大区，映射关系请参见区域映射
	Area *string `json:"Area,omitempty"`

	// 国家，映射关系请参见区域映射。如果按国家筛选，需要同时传入 Area 和 Country。
	Country *string `json:"Country,omitempty"`

	// 国内为省，国外暂不支持该参数，映射关系请参见区域映射。如果按省筛选，需要同时传入 Area、Country 和 Province。
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

	// REQUIRED; 聚合的时间粒度，单位为秒。
	// * 60：1 分钟；
	// * 300：5 分钟；
	// * 3600：1 小时。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 所有时间粒度的数据。
	StatusDataList []DescribeLivePlayStatusCodeDataResResultStatusDataListItem `json:"StatusDataList"`

	// REQUIRED; 当前查询条件下的状态码占比数据。
	StatusSummaryDataList []DescribeLivePlayStatusCodeDataResResultStatusSummaryDataListItem `json:"StatusSummaryDataList"`

	// 数据拆分的维度，维度说明如下所示。
	// * Domain：域名；
	// * ISP：运营商。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 运营商。
	ISPList []*string `json:"ISPList,omitempty"`

	// CDN 节点 IP 所属区域列表。
	RegionList []*DescribeLivePlayStatusCodeDataResResultRegionListItem `json:"RegionList,omitempty"`

	// 按维度拆分后的数据。
	StatusDetailDataList []*DescribeLivePlayStatusCodeDataResResultStatusDetailDataListItem `json:"StatusDetailDataList,omitempty"`

	// 客户端 IP 所属区域列表。
	UserRegionList []*DescribeLivePlayStatusCodeDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLivePlayStatusCodeDataResResultRegionListItem struct {

	// 区域信息中的大区标识符。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符。
	Province *string `json:"Province,omitempty"`
}

type DescribeLivePlayStatusCodeDataResResultStatusDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的状态码详细数据。
	StatusSummaryDataList []DescribeLivePlayStatusCodeDataResResultStatusDataListPropertiesItemsItem `json:"StatusSummaryDataList"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
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

	// 域名。
	Domain *string `json:"Domain,omitempty"`

	// 运营商。
	ISP *string `json:"ISP,omitempty"`

	// 每个时间点的粒度数据。
	StatusDataList []*DescribeLivePlayStatusCodeDataResResultStatusDetailDataListPropertiesItemsItem `json:"StatusDataList,omitempty"`
}

type DescribeLivePlayStatusCodeDataResResultStatusDetailDataListPropertiesItemsItem struct {

	// REQUIRED; 按状态码区分的数据列表。
	StatusSummaryDataList []DescribeLivePlayStatusCodeDataResResultStatusDetailDataListPropertiesItemsStatusSummaryDataListItem `json:"StatusSummaryDataList"`

	// REQUIRED; 时间片起始时刻。RFC3339 格式的 UTC 时间，精度为 s。
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
	Area     *string `json:"Area,omitempty"`
	Country  *string `json:"Country,omitempty"`
	Province *string `json:"Province,omitempty"`
}

type DescribeLivePullToPushBandwidthDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天；
	// * 86400：1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 数据拆分的维度，缺省情况下不进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * DstAddrType：推流地址类型。 :::tip 配置数据拆分维度时，对应的维度参数传入多个值时会返回按维度进行拆分的数据；对应的维度只传入一个值时不返回按维度进行拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 推流域名列表，缺省情况表示当前用户的所有推拉流域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 推流地址类型，可选值如下所示。
	// * Live：非第三方；
	// * Third：（默认值）第三方。
	DstAddrTypeList []*string `json:"DstAddrTypeList,omitempty"`
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

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 300：5 分钟；
	// * 3600：1 小时；
	// * 86400：1 天。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 所有时间粒度的数据。
	BandwidthDataList []DescribeLivePullToPushBandwidthDataResResultBandwidthDataListItem `json:"BandwidthDataList"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 当前查询条件下的拉流转推峰值带宽，单位为 Mbps。
	PeakUpBandwidth float32 `json:"PeakUpBandwidth"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 按维度拆分的数据。
	BandwidthDetailDataList []*DescribeLivePullToPushBandwidthDataResResultBandwidthDetailDataListItem `json:"BandwidthDetailDataList,omitempty"`

	// 数据拆分的维度，维度说明如下。
	// * Domain：域名；
	// * DstAddrType：推流地址类型。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 推流地址类型。
	// * Live：非第三方；
	// * Third：第三方。
	DstAddrTypeList []*string `json:"DstAddrTypeList,omitempty"`
}

type DescribeLivePullToPushBandwidthDataResResultBandwidthDataListItem struct {

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内的拉流转推峰值带宽，单位为 Mbps。
	UpBandwidth float32 `json:"UpBandwidth"`
}

type DescribeLivePullToPushBandwidthDataResResultBandwidthDetailDataListItem struct {

	// REQUIRED; 按维度进行数据拆分后，当前维度下所有时间粒度的数据。
	BandwidthDataList []DescribeLivePullToPushBandwidthDataResResultBandwidthDetailDataListPropertiesItemsItem `json:"BandwidthDataList"`

	// REQUIRED; 按维度进行数据拆分后，当前维度下的拉流转推峰值带宽，单位为 Mbps。
	PeakUpBandwidth float32 `json:"PeakUpBandwidth"`

	// 按域名维度进行数据拆分时的域名信息。
	Domain *string `json:"Domain,omitempty"`

	// 按推流地址类型维度进行数据拆分时的地址类型信息。
	// * Live：非第三方；
	// * Third：第三方。
	DstAddrType *string `json:"DstAddrType,omitempty"`
}

type DescribeLivePullToPushBandwidthDataResResultBandwidthDetailDataListPropertiesItemsItem struct {

	// REQUIRED; 时间片起始时刻。RFC3339 格式的 UTC 时间，精度为 s，例如，2022-04-13T00:00:00+08:00
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 转推带宽，单位为 Mbps
	UpBandwidth float32 `json:"UpBandwidth"`
}

type DescribeLivePullToPushDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 60：1 分钟。时间粒度为 1 分钟时，单次查询最大时间跨度为 1 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询时间跨度为 93 天，历史查询时间范围为 366 天；
	// * 86400：（默认值）1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 查询流粒度数据时的应用名称。 :::tip 使用 App 构造请求时，需同时定义 Stream 参数，不可缺省。 :::
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，缺省情况下不进行数据拆分，支持的维度如下所示。
	// * Domain：域名。 :::tip 配置数据拆分维度时，对应的维度参数传入多个值时会返回按维度进行拆分的数据；对应的维度只传入一个值时不返回按维度进行拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，缺省情况表示当前用户的所有推拉流域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 查询流粒度数据时的流名称。 :::tip 使用 Stream 构造请求时，需同时定义 App 参数，不可缺省。 :::
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

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 60：1 分钟；
	// * 3600：1 小时；
	// * 86400：1 天。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 所有时间粒度的数据。
	PullToPushDataList []DescribeLivePullToPushDataResResultPullToPushDataListItem `json:"PullToPushDataList"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 当前查询条件下的拉流转推总时长，单位为分钟。
	TotalDuration float32 `json:"TotalDuration"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，维度说明如下。
	// * Domain：域名。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 按维度拆分后的数据。
	PullToPushDetailDataList []*DescribeLivePullToPushDataResResultPullToPushDetailDataListItem `json:"PullToPushDetailDataList,omitempty"`

	// 流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLivePullToPushDataResResultPullToPushDataListItem struct {

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
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

	// REQUIRED; 时间片起始时刻。RFC3339 格式的 UTC 时间，精度为 s，例如，2022-04-13T00:00:00+08:00
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 该时间片内的拉流转推总时长，单位分钟，保留小数点后 2 位
	Value float32 `json:"Value"`
}

type DescribeLivePushStreamCountDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 60：1 分钟。时间粒度为 1 分钟时，单次查询最大时间跨度为 24 小时，历史查询时间范围为 366 天；
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天；
	// * 86400：1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 数据拆分的维度，缺省情况下不进行数据拆分，支持的维度如下所示。
	// * Domain：域名。
	// :::tip 配置数据拆分维度时，对应的维度参数传入多个值时会返回按维度进行拆分的数据；对应的维度只传入一个值时不返回按维度进行拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，缺省情况表示当前用户的所有推流域名和拉流域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，缺省情况下表示所有运营商，支持的运营商如下所示。
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
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// 客户端 IP 所属区域的列表，缺省情况下表示所有区域。
	UserRegionList []*DescribeLivePushStreamCountDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLivePushStreamCountDataBodyUserRegionListItem struct {

	// 区域信息中的大区标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按国家筛选，需要同时传入Area和Country。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符，国外暂不支持该参数，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按省筛选，需要同时传入Area、Country和Province。
	Province *string `json:"Province,omitempty"`
}

type DescribeLivePushStreamCountDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLivePushStreamCountDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLivePushStreamCountDataResResult `json:"Result"`
}

type DescribeLivePushStreamCountDataResResponseMetadata struct {

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

type DescribeLivePushStreamCountDataResResult struct {

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 60：1 分钟；
	// * 300：5 分钟；
	// * 3600：1 小时；
	// * 86400：1 天。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询时间范围内的推流数量最大值。
	PeakCount int32 `json:"PeakCount"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 所有时间粒度的数据。
	TotalStreamDataList []DescribeLivePushStreamCountDataResResultTotalStreamDataListItem `json:"TotalStreamDataList"`

	// 数据拆分的维度，维度说明如下所示。
	// * Domain：域名。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，标识符与运营商的对应关系如下。
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
	ISPList []*string `json:"ISPList,omitempty"`

	// 按维度拆分后的数据。
	StreamDetailDataList []*DescribeLivePushStreamCountDataResResultStreamDetailDataListItem `json:"StreamDetailDataList,omitempty"`

	// 客户端 IP 所属区域列表。
	UserRegionList []*DescribeLivePushStreamCountDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLivePushStreamCountDataResResultStreamDetailDataListItem struct {

	// REQUIRED; 按域名维度进行数据拆分时的域名信息。
	Domain string `json:"Domain"`

	// REQUIRED; 按维度进行数据拆分后，当前维度下的所有时间粒度数据。
	TotalStreamDataList []DescribeLivePushStreamCountDataResResultStreamDetailDataListPropertiesItemsItem `json:"TotalStreamDataList"`
}

type DescribeLivePushStreamCountDataResResultStreamDetailDataListPropertiesItemsItem struct {

	// REQUIRED
	PeakCount int32 `json:"PeakCount"`

	// REQUIRED
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLivePushStreamCountDataResResultTotalStreamDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的推流数量最大值。
	PeakCount int32 `json:"PeakCount"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLivePushStreamCountDataResResultUserRegionListItem struct {

	// 区域信息中的大区标识符。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符。
	Province *string `json:"Province,omitempty"`
}

type DescribeLivePushStreamMetricsBody struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 推流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	// :::tip 单次查询最大时间跨度为 1 天，历史查询最大时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// REQUIRED; 流名称。
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

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 5：5 秒；
	// * 30：30 秒。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 推流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 所有时间粒度的数据。
	MetricList []DescribeLivePushStreamMetricsResResultMetricListItem `json:"MetricList"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`
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

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻视频帧显示时间戳差值的最大值，单位为毫秒。
	VideoFrameGap int32 `json:"VideoFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个视频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	VideoPts int32 `json:"VideoPts"`
}

type DescribeLiveRecordDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询最大时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天；
	// * 86400：1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 查询流粒度数据时的应用名称。 :::tip 使用 App 构造请求时，需同时定义 Stream 参数，不可缺省。 :::
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，缺省情况下不进行数据拆分，支持的维度如下所示。
	// * Domain：域名。 :::tip 配置数据拆分维度时，对应的维度参数需传入多个值时会返回按维度进行拆分的数据；对应的维度只传入一个值时不返回按维度进行拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，缺省情况表示当前用户的所有推拉流域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 查询流粒度数据时的流名称， :::tip 使用 Stream 构造请求时，需同时定义 App 参数，不可缺省。 :::
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

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 300：5 分钟；
	// * 3600：1 小时；
	// * 86400：1 天。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 所有时间粒度的数据。
	RecordDataList []DescribeLiveRecordDataResResultRecordDataListItem `json:"RecordDataList"`

	// REQUIRED; 当前查询条件下的录制并发路数最大值。
	RecordPeak int32 `json:"RecordPeak"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 查询流粒度数据时的应用名称。
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，维度说明如下。
	// * Domain：域名。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 按维度拆分后的数据。
	RecordDetailDataList []*DescribeLiveRecordDataResResultRecordDetailDataListItem `json:"RecordDetailDataList,omitempty"`

	// 查询流粒度数据时的流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveRecordDataResResultRecordDataListItem struct {

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
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

	// REQUIRED; 时间
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 300：5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天；
	// * 86400：（默认值）1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 查询流粒度数据时的应用名称。 :::tip 使用 App 构造请求时，需要同时定义 Stream 参数，不可缺省。 :::
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，缺省情况下不进行数据拆分，支持的维度如下所示。
	// * Domain：域名。 :::tip 配置数据拆分维度时，对应的维度参数传入多个值时会返回按维度进行拆分的数据；对应的维度只传入一个值时不返回按维度进行拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，缺省情况表示当前用户的所有推拉流域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 查询流粒度数据时的流名称。 :::tip 使用 Stream 构造请求时，需要同时定义 App 参数，不可缺省。 :::
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

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 300：5 分钟；
	// * 3600：1 小时；
	// * 86400：1 天。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 所有时间粒度的数据。
	SnapshotDataList []DescribeLiveSnapshotDataResResultSnapshotDataListItem `json:"SnapshotDataList"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 当前查询条件下的截图总张数。
	Total int32 `json:"Total"`

	// 查询流粒度数据时的应用名称。
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，维度说明如下。
	// * Domain：域名。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 按维度拆分后的数据。
	SnapshotDetailData []*DescribeLiveSnapshotDataResResultSnapshotDetailDataItem `json:"SnapshotDetailData,omitempty"`

	// 查询流粒度数据时的流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveSnapshotDataResResultSnapshotDataListItem struct {

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
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

	// REQUIRED; 时间片起始时刻。RFC3339 格式的 UTC 时间，精度为 s；例如，2022-04-13T00:00:00+08:00
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 截图总张数
	Value int32 `json:"Value"`
}

type DescribeLiveSourceBandwidthDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 60：1 分钟。时间粒度为 1 分钟时，单次查询最大时间跨度为 24 小时，历史查询时间范围为 366 天；
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 查询流粒度数据时的应用名称。 :::tip 使用 App 构造请求时，需要同时定义 Domain 和 Stream 参数，不可缺省。 :::
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，缺省情况下不进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * ISP：运营商。
	// :::tip 配置数据拆分维度时，对应的维度参数传入多个值时会返回按维度进行拆分的数据；对应的维度只传入一个值时不返回按维度进行拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 查询流粒度数据时的域名，支持填写拉流域名。 :::tip 使用 Domain 构造请求时，需要同时定义 App 和 Stream 参数，不可缺省。 :::
	Domain *string `json:"Domain,omitempty"`

	// 拉流域名列表，缺省情况表示当前用户的所有推拉流域名。 :::tipDomainList 和 Domain 传且仅传一个。 :::
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，缺省情况下表示所有运营商，支持的运营商如下所示。
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
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// 查询流粒度数据时的流名称。 :::tip 使用 Stream 构造请求时，需要同时定义 Domain 和 App 参数，不可缺省。 :::
	Stream *string `json:"Stream,omitempty"`

	// 客户端 IP 所属区域的列表，缺省情况下表示所有区域。
	UserRegionList []*DescribeLiveSourceBandwidthDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveSourceBandwidthDataBodyUserRegionListItem struct {

	// 区域信息的大区标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。
	Area *string `json:"Area,omitempty"`

	// 区域信息的国家标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按国家筛选，需要同时传入 Area 和 Country。
	Country *string `json:"Country,omitempty"`

	// 区域信息的省份标识符，国外暂不支持该参数，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按省筛选，需要同时传入 Area、Country 和 Province。
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

	// REQUIRED; 聚合的时间粒度，单位为秒。
	// * 60：1 分钟；
	// * 300：5 分钟；
	// * 3600：1 小时。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询流粒度数据时的应用名称。
	App string `json:"App"`

	// REQUIRED; 所有时间粒度的数据。
	BandwidthDataList []DescribeLiveSourceBandwidthDataResResultBandwidthDataListItem `json:"BandwidthDataList"`

	// REQUIRED; 按维度拆分后的数据。
	BandwidthDetailDataList []DescribeLiveSourceBandwidthDataResResultBandwidthDetailDataListItem `json:"BandwidthDetailDataList"`

	// REQUIRED; 数据拆分的维度，维度说明如下所示。
	// * Domain：域名；
	// * ISP：运营商。
	DetailField []string `json:"DetailField"`

	// REQUIRED; 查询流粒度数据时的域名。
	Domain string `json:"Domain"`

	// REQUIRED; 域名列表。
	DomainList []string `json:"DomainList"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 提供网络接入服务的运营商标识符，标识符与运营商的对应关系如下。
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
	ISPList []string `json:"ISPList"`

	// REQUIRED; 查询时间范围内的回源峰值带宽，单位为 Mbps。
	PeakBandwidth float32 `json:"PeakBandwidth"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 查询流粒度数据时的流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 客户端 IP 所属区域列表。
	UserRegionList []DescribeLiveSourceBandwidthDataResResultUserRegionListItem `json:"UserRegionList"`
}

type DescribeLiveSourceBandwidthDataResResultBandwidthDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的回源峰值带宽，单位为 Mbps。
	Bandwidth float32 `json:"Bandwidth"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLiveSourceBandwidthDataResResultBandwidthDetailDataListItem struct {

	// REQUIRED; 按维度进行数据拆分后，当前维度下所有时间粒度的数据。
	BandwidthDataList []DescribeLiveSourceBandwidthDataResResultBandwidthDetailDataListPropertiesItemsItem `json:"BandwidthDataList"`

	// REQUIRED; 按域名维度进行数据拆分时的域名信息。
	Domain string `json:"Domain"`

	// REQUIRED; 按运营商维度进行数据拆分时的运营商信息。
	ISP string `json:"ISP"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的回源峰值带宽，单位为 Mbps。
	PeakBandwidth float32 `json:"PeakBandwidth"`
}

type DescribeLiveSourceBandwidthDataResResultBandwidthDetailDataListPropertiesItemsItem struct {

	// REQUIRED; 时间片内回源带宽峰值，单位 Mbps
	Bandwidth float32 `json:"Bandwidth"`

	// REQUIRED; 时间片起始时刻。RFC3339 格式的 UTC 时间，精度为 s，例如，2022-04-13T00:00:00+08:00
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLiveSourceBandwidthDataResResultUserRegionListItem struct {

	// REQUIRED; 区域信息中的大区标识符。
	Area string `json:"Area"`

	// REQUIRED; 区域信息中的国际标识符。
	Country string `json:"Country"`

	// REQUIRED; 区域信息中的省份标识符。
	Province string `json:"Province"`
}

type DescribeLiveSourceStreamMetricsBody struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	// :::tip 单次查询最大时间跨度为 1 天，历史查询最大时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 30：（默认值）30 秒。
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

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 30：30 秒。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 所有时间粒度的数据。
	MetricList []DescribeLiveSourceStreamMetricsResResultMetricListItem `json:"MetricList"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
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

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻视频帧显示时间戳差值的最大值，单位为毫秒。
	VideoFrameGap int32 `json:"VideoFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个视频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	VideoPts int32 `json:"VideoPts"`
}

type DescribeLiveSourceTrafficDataBody struct {

	// REQUIRED; The end time of the query, in UTC time in RFC 3339 format, with a precision of seconds.
	EndTime string `json:"EndTime"`

	// REQUIRED; The start time of the query, in UTC time in RFC 3339 format, with a precision of seconds.
	StartTime string `json:"StartTime"`

	// The time granularity of the data aggregation, in seconds, is supported as follows.
	// * 60:1 minute. When the time granularity is 1 minute, the maximum time span for a single query is 24 hours, and the historical
	// query time range is 366 days;
	// * 300: (default) 5 minutes. When the time granularity is 5 minutes, the maximum time span for a single query is 31 days,
	// and the historical query time range is 366 days;
	// * 3600:1 hour. When the time granularity is 1 day, the maximum time span for a single query is 93 days, and the historical
	// query time range is 366 days.
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// The name of the application when querying stream granular data. ::: tips When using'App 'to construct a request, you need
	// to define both'Domain' and'Stream 'parameters, which cannot be defaulted. :::
	App *string `json:"App,omitempty"`

	// Dimensions for data splitting. Data splitting is not performed by default. The supported dimensions are as follows.
	// * Domain: domain name;
	// * IP: the IP address of the export extranet;
	// * ISP: Operator.
	// ::: tips When configuring a data split dimension, the corresponding dimension parameter will return data split by dimension
	// when multiple values are passed; the corresponding dimension will not return
	// data split by dimension when only one value is passed. :::
	DetailField []*string `json:"DetailField,omitempty"`

	// The domain name for stream pulling. Use this parameter if you want to query the data for a specific stream. The Domain,
	// App, and Stream parameters must be specified at the same time.
	Domain *string `json:"Domain,omitempty"`

	// A list of the domain names you want to query. If unspecified, all the domain names under your account will be queried.Note:
	// Specify either DomainList or Domain, not both.
	DomainList []*string `json:"DomainList,omitempty"`

	// The operator identifier that provides the network access service, which by default represents all operators. The supported
	// operators are shown below.
	// * unicom: unicom;
	// * railcom;
	// * telecom: telecommunications;
	// * mobile: mobile;
	// * cernet: education network;
	// * tianwei: tianwei;
	// * alibaba: alibaba;
	// * tenger: Tencent;
	// * drpeng: Dr. Peng;
	// * btvn: radio and television;
	// * huashu: wah number;
	// * other: other.
	// You can also obtain the operator's corresponding identifier through the DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974]
	// interface.
	ISPList []*string `json:"ISPList,omitempty"`

	// The name of the stream when querying stream granular data. ::: tip When using'Stream 'to construct a request, you need
	// to define both'Domain' and'App 'parameters, which cannot be defaulted. :::
	Stream *string `json:"Stream,omitempty"`

	// List of regions to which the client side IP belongs, representing all regions by default.
	UserRegionList []*DescribeLiveSourceTrafficDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveSourceTrafficDataBodyUserRegionListItem struct {

	// For the region identifier in the region information, see Query Region Identifier (https://www.volcengine.com/docs/6469/1133973).
	Area *string `json:"Area,omitempty"`

	// The country identifier in the region information, see Query Region Identifier [https://www.volcengine.com/docs/6469/1133973].
	// If filtering by country, you need to pass in both Area and Country.
	Country *string `json:"Country,omitempty"`

	// The province identifier in the regional information. This parameter is not supported abroad for the time being. Please
	// refer to Query Area Identifier [https://www.volcengine.com/docs/6469/1133973] for
	// how to obtain it. If filtering by province, you need to pass in'Area ',' Country 'and'Province' at the same time.
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveSourceTrafficDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveSourceTrafficDataResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeLiveSourceTrafficDataResResult          `json:"Result,omitempty"`
}

type DescribeLiveSourceTrafficDataResResponseMetadata struct {

	// REQUIRED; The interface name of the request, which is a public parameter of the request.
	Action string `json:"Action"`

	// REQUIRED; The requested Region, for example: cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID is the unique identifier for each API request.
	RequestID string `json:"RequestId"`

	// REQUIRED; The requested service is a public parameter of the request.
	Service string `json:"Service"`

	// REQUIRED; The version number of the request, which is a public parameter of the request.
	Version string                                                 `json:"Version"`
	Error   *DescribeLiveSourceTrafficDataResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeLiveSourceTrafficDataResResponseMetadataError struct {

	// error code
	Code *string `json:"Code,omitempty"`

	// error message
	Message *string `json:"Message,omitempty"`
}

type DescribeLiveSourceTrafficDataResResult struct {

	// REQUIRED; The time granularity of the aggregation, in seconds.
	// * 60:1 minute;
	// * 300:5 minutes;
	// * 3600:1 hour.
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; The name of the application when querying stream granular data.
	App string `json:"App"`

	// REQUIRED; The dimension of the data split, the dimension description is shown below.
	// * Domain: domain name;
	// * IP: the IP address of the export extranet;
	// * ISP: Operator.
	DetailField []string `json:"DetailField"`

	// REQUIRED; The domain name when querying stream granular data.
	Domain string `json:"Domain"`

	// REQUIRED; List of domain names.
	DomainList []string `json:"DomainList"`

	// REQUIRED; The end time of the query, in UTC time in RFC 3339 format, with a precision of seconds.
	EndTime string `json:"EndTime"`

	// REQUIRED; The operator identifier that provides network access services, and the corresponding relationship between the
	// identifier and the operator is as follows.
	// * unicom: unicom;
	// * railcom;
	// * telecom: telecommunications;
	// * mobile: mobile;
	// * cernet: education network;
	// * tianwei: tianwei;
	// * alibaba: alibaba;
	// * tenger: Tencent;
	// * drpeng: Dr. Peng;
	// * btvn: radio and television;
	// * huashu: wah number;
	// * other: other.
	ISPList []string `json:"ISPList"`

	// REQUIRED; The start time of the query, in UTC time in RFC 3339 format, with a precision of seconds.
	StartTime string `json:"StartTime"`

	// REQUIRED; The name of the stream when querying stream granular data.
	Stream string `json:"Stream"`

	// REQUIRED; Query the total traffic back to the source within the time range, in GB.
	TotalTraffic float32 `json:"TotalTraffic"`

	// REQUIRED; All time-granular data.
	TrafficDataList []DescribeLiveSourceTrafficDataResResultTrafficDataListItem `json:"TrafficDataList"`

	// REQUIRED; Data split by dimension.
	TrafficDetailDataList []DescribeLiveSourceTrafficDataResResultTrafficDetailDataListItem `json:"TrafficDetailDataList"`

	// REQUIRED; List of client side IP regions.
	UserRegionList []DescribeLiveSourceTrafficDataResResultUserRegionListItem `json:"UserRegionList"`
}

type DescribeLiveSourceTrafficDataResResultTrafficDataListItem struct {

	// REQUIRED; When data is aggregated by time granularity, the start time of each time granularity, in UTC time in RFC 3339
	// format, with a precision of seconds.
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; Flow back to the source generated within the current data aggregation time granularity, in GB.
	Traffic float32 `json:"Traffic"`
}

type DescribeLiveSourceTrafficDataResResultTrafficDetailDataListItem struct {

	// REQUIRED; Domain name information when data is split by domain name dimension.
	Domain string `json:"Domain"`

	// REQUIRED; Operator information when data is split according to the operator dimension.
	ISP string `json:"ISP"`

	// REQUIRED; After splitting the data by dimension, the total traffic in the return source of the current dimension is in
	// GB.
	TotalTraffic float32 `json:"TotalTraffic"`

	// REQUIRED; After splitting the data by dimension, the data of all time granularities in the current dimension.
	TrafficDataList []DescribeLiveSourceTrafficDataResResultTrafficDetailDataListPropertiesItemsItem `json:"TrafficDataList"`
}

type DescribeLiveSourceTrafficDataResResultTrafficDetailDataListPropertiesItemsItem struct {

	// REQUIRED; Time slice start time. UTC time in RFC3339 format with precision s, for example, '2022-04-13T00:00:00 + 08:00'
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; Return traffic, in GB
	Traffic float32 `json:"Traffic"`
}

type DescribeLiveSourceTrafficDataResResultUserRegionListItem struct {

	// REQUIRED; The region identifier in the region information.
	Area string `json:"Area"`

	// REQUIRED; The country identifier in the regional information.
	Country string `json:"Country"`

	// REQUIRED; The province identifier in the regional information.
	Province string `json:"Province"`
}

type DescribeLiveStreamCountDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 60：1 分钟。时间粒度为 1 分钟时，单次查询最大时间跨度为 24 小时，历史查询时间范围为 366 天；
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天；
	// * 86400：1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 数据拆分的维度，缺省情况下不进行数据拆分，支持的维度如下所示。
	// * Domain：域名。
	// :::tip 配置数据拆分维度时，对应的维度参数传入多个值时会返回按维度进行拆分的数据；对应的维度只传入一个值时不返回按维度进行拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，缺省情况表示当前用户的所有推拉流域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 流类型，缺省情况下表示全部类型，支持的流类型取值如下。
	// * push：推流；
	// * relay-source：回源流；
	// * transcode：转码流。
	StreamType []*string `json:"StreamType,omitempty"`
}

type DescribeLiveStreamCountDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveStreamCountDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveStreamCountDataResResult `json:"Result"`
}

type DescribeLiveStreamCountDataResResponseMetadata struct {

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

type DescribeLiveStreamCountDataResResult struct {

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 60：1 分钟；
	// * 300：5 分钟；
	// * 3600：1 小时；
	// * 86400：1 天。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 当前查询条件下流数最大值。
	PeakCount int32 `json:"PeakCount"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 所有时间粒度的数据。
	TotalStreamDataList []DescribeLiveStreamCountDataResResultTotalStreamDataListItem `json:"TotalStreamDataList"`

	// 数据拆分的维度，维度说明如下所示。
	// * Domain：域名。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 按维度拆分后的数据。
	StreamDetailDataList []*DescribeLiveStreamCountDataResResultStreamDetailDataListItem `json:"StreamDetailDataList,omitempty"`

	// 流类型，流类型说明如下。
	// * push：拉流；
	// * relay-source：回源流；
	// * transcode：转码流。
	StreamType []*string `json:"StreamType,omitempty"`
}

type DescribeLiveStreamCountDataResResultStreamDetailDataListItem struct {

	// REQUIRED; 按域名维度进行数据拆分时的域名信息。
	Domain string `json:"Domain"`

	// REQUIRED; 按维度进行数据拆分后，当前维度下的所有时间粒度数据。
	TotalStreamDataList []DescribeLiveStreamCountDataResResultStreamDetailDataListPropertiesItemsItem `json:"TotalStreamDataList"`
}

type DescribeLiveStreamCountDataResResultStreamDetailDataListPropertiesItemsItem struct {

	// REQUIRED
	PeakCount int32 `json:"PeakCount"`

	// REQUIRED
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLiveStreamCountDataResResultTotalStreamDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的流数最大值。
	PeakCount int32 `json:"PeakCount"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLiveStreamInfoByPageQuery struct {

	// REQUIRED; 查询数据的页码，取值为正整数。
	PageNum int32 `json:"PageNum" query:"PageNum"`

	// REQUIRED; 每页显示的数据条数，取值范围为 [1,1000]。
	PageSize int32 `json:"PageSize" query:"PageSize"`

	// 应用名称，取值与直播流地址中 AppName 字段取值相同，默认为空，表示查询所有应用名称。支持由大小写字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App *string `json:"App,omitempty" query:"App"`

	// 直播流使用的域名，默认为空，表示查询所有当前域名空间（Vhost）下的在线流。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console-stable.volcanicengine.com/live/main/domain/list]
	// 页面，查看直播流使用的域名。
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

	// 流名称，取值与直播流地址中 StreamName 字段取值相同，默认为空表示查询所有流名称。支持由大小写字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	Stream *string `json:"Stream,omitempty" query:"Stream"`

	// 在线流的流类型，默认为空，表示查询所有类型，支持的取值即含义如下所示。
	// * origin：原始流；
	// * trans：转码流。
	StreamType *string `json:"StreamType,omitempty" query:"StreamType"`

	// 域名空间，即直播流地址的域名（Domain）所属的域名空间（Vhost），默认为空，表示查询所有域名空间（Vhost）下的在线流。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]
	// 接口或在视频直播控制台的域名管理
	// [https://console-stable.volcanicengine.com/live/main/domain/list]页面，查看需要查询的直播流使用的域名所属的域名空间。
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

	// REQUIRED; 在线流的开始时间。
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 60：1 分钟。时间粒度为 1 分钟时，单次查询最大时间跨度为 24 小时，历史查询时间范围为 366 天；
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 查询流粒度数据时的应用名。 :::tip 使用App构造请求时，需要同时定义Domain和Stream参数，不可缺省。 :::
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，缺省情况下不进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * ISP：运营商；
	// * Protocol：推拉流协议；
	// * Referer：请求的 Referer 信息。
	// :::tip 配置数据拆分维度时，对应的维度参数需传入多个值时会返回按维度进行拆分的数据；对应的维度只传入一个值时不返回按维度进行拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 查询流粒度数据时的域名参数。 :::tip 使用Domain构造请求时，需要同时定义App和Stream参数，不可缺省。 :::
	Domain *string `json:"Domain,omitempty"`

	// 域名列表，缺省情况表示该用户的所有推拉流域名。 :::tipDomainList和Domain传且仅传一个。 :::
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，缺省情况下表示所有运营商，支持的运营商如下所示。
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
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// 推拉流协议，缺省情况下表示所有协议类型，支持的协议如下所示。
	// * HTTP-FLV：基于 HTTP 协议的推拉流协议，使用 FLV 格式传输视频格式。
	// * HTTP-HLS：基于 HTTP 协议的推拉流协议，使用 TS 格式传输视频格式。
	// * RTMP：Real Time Message Protocol，实时信息传输协议。
	// * RTM：Real Time Media，超低延时直播协议。
	// * SRT：Secure Reliable Transport，安全可靠传输协议。
	// * QUIC：Quick UDP Internet Connections，一种基于 UDP 的全新的低延时互联网传输协议。
	// :::tip 如果查询推拉流协议为 QUIC，不能同时查询其他协议。 :::
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// 请求的 Referer 信息。
	RefererList []*string `json:"RefererList,omitempty"`

	// CDN 节点 IP 所属区域的列表，缺省情况下表示所有区域。
	RegionList []*DescribeLiveStreamSessionDataBodyRegionListItem `json:"RegionList,omitempty"`

	// 查询流粒度数据时的流名称。 :::tip 使用Stream构造请求时，需要同时定义Domain和App参数，不可缺省。 :::
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveStreamSessionDataBodyRegionListItem struct {

	// 区域信息中的大区标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按国家筛选，需要同时传入Area和Country。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符，国外暂不支持该参数，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按省筛选，需要同时传入Area、Country和Province。
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

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 60：1 分钟；
	// * 300：5 分钟；
	// * 3600：1 天。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询时间范围内的在线人数峰值。
	PeakOnlineUser int32 `json:"PeakOnlineUser"`

	// REQUIRED; 所有时间粒度的数据。
	SessionDataList []DescribeLiveStreamSessionDataResResultSessionDataListItem `json:"SessionDataList"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
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

	// 域名。
	Domain *string `json:"Domain,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，标识符与运营商的对应关系如下。
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
	ISPList []*string `json:"ISPList,omitempty"`

	// 推拉流协议，协议说明如下。
	// * HTTP-FLV：基于 HTTP 协议的推拉流协议，使用 FLV 格式传输视频格式。
	// * HTTP-HLS：基于 HTTP 协议的推拉流协议，使用 TS 格式传输视频格式。
	// * RTMP：Real Time Message Protocol，实时信息传输协议。
	// * RTM：Real Time Media，超低延时直播协议。
	// * SRT：Secure Reliable Transport，安全可靠传输协议。
	// * QUIC：Quick UDP Internet Connections，一种基于 UDP 的全新的低延时互联网传输协议。
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// 请求的 Referer 信息。
	RefererList []*string `json:"RefererList,omitempty"`

	// 区域列表。
	RegionList []*DescribeLiveStreamSessionDataResResultRegionListItem `json:"RegionList,omitempty"`

	// 按维度拆分的数据。
	SessionDetailDataList []*DescribeLiveStreamSessionDataResResultSessionDetailDataListItem `json:"SessionDetailDataList,omitempty"`

	// 流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveStreamSessionDataResResultRegionListItem struct {

	// 区域信息中的大区标识符。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveStreamSessionDataResResultSessionDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的在线人数最大值。
	OnlineUser int32 `json:"OnlineUser"`

	// REQUIRED; 当前数据聚合时间粒度内的请求数。
	Request int32 `json:"Request"`

	// REQUIRED; 数据按时间粒度聚合时，诶个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
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

	// REQUIRED; 时间片起始时刻。RFC3339 时间，例如，2022-04-13T00:00:00+08:00
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLiveStreamStateQuery struct {

	// REQUIRED; 应用名称，取值与直播流地址的 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App" query:"App"`

	// REQUIRED; 流名称，取值与直播流地址的 StreamName 字段取值相同。支持由大小写字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	Stream string `json:"Stream" query:"Stream"`

	// 填写直播流使用的域名，默认为空，表示查询所有直推流和回源流的状态和类型。 您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console-stable.volcanicengine.com/live/main/domain/list]
	// 页面，查看需要查询的直播流使用的域名。 :::tipVhost 和 Domain 传且仅传一个。 :::
	Domain *string `json:"Domain,omitempty" query:"Domain"`

	// 域名空间，即直播流地址的域名（Domain）所属的域名空间（Vhost）。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console-stable.volcanicengine.com/live/main/domain/list]
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

type DescribeLiveStreamUsageDataBody struct {

	// 域名。
	Domain *string `json:"Domain,omitempty"`

	// 查询时间，格式为 yyyy-mm-dd HH:MM。 例如：查询时间 2023-01-01 10:00 时表示查询时间范围为 10:00 到 10:01。 :::tip
	// * 未填写查询时间时，默认查询时间为当前时间减 5 分钟。
	// * 最长支持查询的历史时间范围为 31 天。 :::
	QueryTime *string `json:"QueryTime,omitempty"`

	// 流名称。
	StreamName *string `json:"StreamName,omitempty"`
}

type DescribeLiveStreamUsageDataRes struct {

	// REQUIRED; 响应数据。
	Response DescribeLiveStreamUsageDataResResponse `json:"Response"`

	// REQUIRED; 请求失败原因：
	// * 请求成功时 Result 为空
	// * 请求失败时 Result 展示失败的原因
	Result string `json:"Result"`

	// REQUIRED; 请求状态：
	// * 1：请求成功
	// * 0：请求失败
	Status int32 `json:"Status"`
}

// DescribeLiveStreamUsageDataResResponse - 响应数据。
type DescribeLiveStreamUsageDataResResponse struct {

	// REQUIRED; 详细数据信息。
	DataInfoList []DescribeLiveStreamUsageDataResResponseDataInfoListItem `json:"DataInfoList"`

	// REQUIRED; 查询时间。
	QueryTime string `json:"QueryTime"`

	// REQUIRED; 请求 ID。
	RequestID string `json:"RequestId"`
}

type DescribeLiveStreamUsageDataResResponseDataInfoListItem struct {

	// REQUIRED; 带宽，单位 kbps。
	Bandwidth int32 `json:"Bandwidth"`

	// REQUIRED; 域名
	Domain string `json:"Domain"`

	// REQUIRED; 在线人数。
	OnlineUser int32 `json:"OnlineUser"`

	// REQUIRED; 协议。
	Protocol string `json:"Protocol"`

	// REQUIRED; 请求数。
	Request int32 `json:"Request"`

	// REQUIRED; 流名称。
	StreamName string `json:"StreamName"`
}

type DescribeLiveTimeShiftDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。 :::tip 单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 86400：（默认值）1 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 域名空间列表，缺省情况表示查询当前用户的所有域名空间。
	Vhosts []*string `json:"Vhosts,omitempty"`
}

type DescribeLiveTimeShiftDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveTimeShiftDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveTimeShiftDataResResult `json:"Result"`
}

type DescribeLiveTimeShiftDataResResponseMetadata struct {

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

type DescribeLiveTimeShiftDataResResult struct {

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 86400：1 天。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 所有时间粒度的数据。
	TimeShiftDataList []DescribeLiveTimeShiftDataResResultTimeShiftDataListItem `json:"TimeShiftDataList"`

	// 域名空间列表。
	Vhosts []*string `json:"Vhosts,omitempty"`
}

type DescribeLiveTimeShiftDataResResultTimeShiftDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的时移文件存储用量，单位为 GB。
	Storage float32 `json:"Storage"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLiveTrafficDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 300：（默认值）5 分钟。聚合粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询最大时间范围为 366 天；
	// * 3600：1 小时。聚合粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询最大时间范围为 366 天；
	// * 86400：1 天。聚合粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询最大时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 数据拆分的维度，缺省情况下不进行数据拆分，支持的维度如下。
	// * Domain：域名；
	// * ISP：运营商；
	// * Protocol：推拉流协议。 :::tip 配置数据拆分维度时，对应的维度参数传入多个值时会返回按维度进行拆分的数据；对应的维度只传入一个值时不返回按维度进行拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，缺省情况表示当前用户的所有推拉流域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 提供网络接入服务的运营商标识符，缺省情况下表示所有运营商，支持的运营商如下所示。
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
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// 推拉流协议，缺省情况下表示所有协议类型，支持的协议如下所示。
	// * HTTP-FLV：基于 HTTP 协议的推拉流协议，使用 FLV 格式传输视频格式。
	// * HTTP-HLS：基于 HTTP 协议的推拉流协议，使用 TS 格式传输视频格式。
	// * RTMP：Real Time Message Protocol，实时信息传输协议。
	// * RTM：Real Time Media，超低延时直播协议。
	// * SRT：Secure Reliable Transport，安全可靠传输协议。
	// * QUIC：Quick UDP Internet Connections，一种基于 UDP 的全新的低延时互联网传输协议。
	// :::tip 如果查询推拉流协议为 QUIC，不能同时查询其他协议。 :::
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	RegionList []*DescribeLiveTrafficDataBodyRegionListItem `json:"RegionList,omitempty"`

	// 指定查询的流量数据为闲时或忙时，缺省情况下为查询全部数据，支持的取值如下。
	// * busy：忙时；
	// * free：闲时。
	Stage *string `json:"Stage,omitempty"`

	// 客户端 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	UserRegionList []*DescribeLiveTrafficDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveTrafficDataBodyRegionListItem struct {

	// 区域信息中的大区标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按国家筛选，需要同时传入Area和Country。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符，国外暂不支持该参数，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按省筛选，需要同时传入Area、Country和Province。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveTrafficDataBodyUserRegionListItem struct {

	// 大区，映射关系请参见区域映射 [https://www.volcengine.com/docs/6469/114196]
	Area *string `json:"Area,omitempty"`

	// 国家，映射关系请参见区域映射 [https://www.volcengine.com/docs/6469/114196]
	Country *string `json:"Country,omitempty"`

	// 国内为省，国外暂不支持该参数，映射关系请参见区域映射 [https://www.volcengine.com/docs/6469/114196]
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

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 300：5 分钟；
	// * 3600：1 小时；
	// * 86400：1 天。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
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

	// 提供网络接入服务的运营商标识符，标识符与运营商的对应关系如下。
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
	ISPList []*string `json:"ISPList,omitempty"`

	// 推拉流协议，协议说明如下。
	// * HTTP-FLV：基于 HTTP 协议的推拉流协议，使用 FLV 格式传输视频格式。
	// * HTTP-HLS：基于 HTTP 协议的推拉流协议，使用 TS 格式传输视频格式。
	// * RTMP：Real Time Message Protocol，实时信息传输协议。
	// * RTM：Real Time Media，超低延时直播协议。
	// * SRT：Secure Reliable Transport，安全可靠传输协议。
	// * QUIC：Quick UDP Internet Connections，一种基于 UDP 的全新的低延时互联网传输协议。
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域列表。
	RegionList []*DescribeLiveTrafficDataResResultRegionListItem `json:"RegionList,omitempty"`

	// 流量数据为闲时或忙时，取值说明如下。
	// * busy：忙时；
	// * free：闲时。
	Stage *string `json:"Stage,omitempty"`

	// 按维度拆分后的数据。 :::tip 请求时，DomainList、ProtocolList和ISPList至少有一个参数传入了多个值时，会返回该参数；否则不返回该参数。 :::
	TrafficDetailDataList []*DescribeLiveTrafficDataResResultTrafficDetailDataListItem `json:"TrafficDetailDataList,omitempty"`

	// 客户端 IP 所属区域列表。
	UserRegionList []*DescribeLiveTrafficDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveTrafficDataResResultRegionListItem struct {

	// 区域信息中的大区标识符。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveTrafficDataResResultTrafficDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内产生的总下行流量，单位 GB。
	DownTraffic float32 `json:"DownTraffic"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
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

	// REQUIRED; 时间片起始时刻。RFC3339 时间，例如，2022-04-13T00:00:00+08:00
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 上行流量，单位 GB
	UpTraffic float32 `json:"UpTraffic"`
}

type DescribeLiveTrafficDataResResultUserRegionListItem struct {

	// 大区
	Area *string `json:"Area,omitempty"`

	// 国家
	Country *string `json:"Country,omitempty"`

	// 国内为省，国外暂不支持该参数
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveTranscodeDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。 :::tip 单次查询最大时间跨度为 93 天，历史查询最大时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 86400：（默认值）1 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 查询流粒度数据时的应用名称。 :::tip 使用App构造请求时，需要同时定义Stream参数，不可缺省。 :::
	App *string `json:"App,omitempty"`

	// 域名列表，缺省情况表示当前用户的所有推拉流域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 查询流粒度数据时的流名称。 :::tip 使用Stream构造请求时，需要同时定义App参数，不可缺省。 :::
	Stream *string `json:"Stream,omitempty"`
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

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	// * 86400：1 天。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询时间范围内的转码总时长，单位为分钟。
	Duration float32 `json:"Duration"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的起始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 所有时间粒度的数据。
	TranscodeDataList []DescribeLiveTranscodeDataResResultTranscodeDataListItem `json:"TranscodeDataList"`

	// 查询流粒度数据时的应用名称。
	App *string `json:"App,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 查询流粒度数据时的流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveTranscodeDataResResultTranscodeDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的转码时长，单位为分钟。
	Duration float32 `json:"Duration"`

	// REQUIRED; 分辨率。- 480P：640 × 480； - 720P：1280 × 720； - 1080P：1920 × 1088； - 2K：2560 × 1440； - 4K：4096 × 2160；- 8K：大于4K； -
	// 0：纯音频流；
	Resolution string `json:"Resolution"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 视频编码格式，支持的取值和含义如下所示。- NormalH264：H.264 标准转码； - NormalH265：H.265 标准转码； - NormalH266：H.266 标准转码； - ByteHDH264：H.264
	// 极智超清； - ByteHDH265：H.265 极智超清； - ByteHDH266：H.266 极智超清；- ByteQE：画质增强；- Audio：纯音频流；
	TranscodeType string `json:"TranscodeType"`
}

type DescribeNSSRewriteConfigBody struct {

	// REQUIRED
	Vhost       string  `json:"Vhost"`
	App         *string `json:"App,omitempty"`
	ServiceType *string `json:"ServiceType,omitempty"`
}

type DescribeNSSRewriteConfigRes struct {

	// REQUIRED
	ResponseMetadata DescribeNSSRewriteConfigResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeNSSRewriteConfigResResult          `json:"Result,omitempty"`
}

type DescribeNSSRewriteConfigResResponseMetadata struct {

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

type DescribeNSSRewriteConfigResResult struct {

	// REQUIRED; @描述：已购资源包总数 @场景：公有云 @示例值：4
	ConfigList []DescribeNSSRewriteConfigResResultConfigListItem `json:"ConfigList"`
}

type DescribeNSSRewriteConfigResResultConfigListItem struct {

	// REQUIRED
	App string `json:"App"`

	// REQUIRED
	Config []string `json:"Config"`

	// REQUIRED
	CreateTime string `json:"CreateTime"`

	// REQUIRED
	DebugHeader string `json:"DebugHeader"`

	// REQUIRED
	Enable bool `json:"Enable"`

	// REQUIRED
	ServiceType string `json:"ServiceType"`

	// REQUIRED
	UpdateTime string `json:"UpdateTime"`

	// REQUIRED
	Vhost string `json:"Vhost"`
}

type DescribePresetAssociationBody struct {
	App    *string `json:"App,omitempty"`
	Domain *string `json:"Domain,omitempty"`
	Stream *string `json:"Stream,omitempty"`
	Vhost  *string `json:"Vhost,omitempty"`
}

type DescribePresetAssociationRes struct {

	// REQUIRED
	ResponseMetadata DescribePresetAssociationResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribePresetAssociationResResult `json:"Result,omitempty"`
}

type DescribePresetAssociationResResponseMetadata struct {

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

// DescribePresetAssociationResResult - 视请求的接口而定
type DescribePresetAssociationResResult struct {

	// REQUIRED
	List []DescribePresetAssociationResResultListItem `json:"List"`
}

type DescribePresetAssociationResResultListItem struct {

	// REQUIRED
	App string `json:"App"`

	// REQUIRED
	AvextractorPresetListV2 []Components7Eb4PfSchemasDescribepresetassociationresPropertiesResultPropertiesListItemsPropertiesAvextractorpresetlistv2Items `json:"AvextractorPresetListV2"`

	// REQUIRED
	AvslicePresetListV2 []Components5Jn2JnSchemasDescribepresetassociationresPropertiesResultPropertiesListItemsPropertiesAvslicepresetlistv2Items `json:"AvslicePresetListV2"`

	// REQUIRED
	CdnSnapshotPresetListV2 []Components1Yrp708SchemasDescribepresetassociationresPropertiesResultPropertiesListItemsPropertiesCdnsnapshotpresetlistv2Items `json:"CdnSnapshotPresetListV2"`

	// REQUIRED
	DataMigrationPresetListV2 []ComponentsF9EcgzSchemasDescribepresetassociationresPropertiesResultPropertiesListItemsPropertiesDatamigrationpresetlistv2Items `json:"DataMigrationPresetListV2"`

	// REQUIRED
	RecordPresetListV2 []DescribePresetAssociationResResultListPropertiesItemsItem `json:"RecordPresetListV2"`

	// REQUIRED
	SnapshotAuditPresetListV2 []Components318E5PSchemasDescribepresetassociationresPropertiesResultPropertiesListItemsPropertiesSnapshotauditpresetlistv2Items `json:"SnapshotAuditPresetListV2"`

	// REQUIRED
	SnapshotPresetListV2 []Components1Je5O2CSchemasDescribepresetassociationresPropertiesResultPropertiesListItemsPropertiesSnapshotpresetlistv2Items `json:"SnapshotPresetListV2"`

	// REQUIRED
	Stream string `json:"Stream"`

	// REQUIRED
	TimeshiftPresetListV2 []ComponentsGhid1HSchemasDescribepresetassociationresPropertiesResultPropertiesListItemsPropertiesTimeshiftpresetlistv2Items `json:"TimeshiftPresetListV2"`

	// REQUIRED; 转码配置列表
	TransPresetList []ComponentsH8On9CSchemasDescribepresetassociationresPropertiesResultPropertiesListItemsPropertiesTranspresetlistItems `json:"TransPresetList"`

	// REQUIRED
	Vhost string `json:"Vhost"`

	// REQUIRED
	WatermarkPresetListV2 []Components17C6BtpSchemasDescribepresetassociationresPropertiesResultPropertiesListItemsPropertiesWatermarkpresetlistv2Items `json:"WatermarkPresetListV2"`
}

type DescribePresetAssociationResResultListPropertiesItemsItem struct {

	// REQUIRED
	CreateTime string `json:"CreateTime"`

	// REQUIRED
	OriginRecord int32 `json:"OriginRecord"`

	// REQUIRED
	PresetName string `json:"PresetName"`

	// REQUIRED; 录制类型，pull,push
	RecordType string `json:"RecordType"`

	// REQUIRED
	RelayEnable bool `json:"RelayEnable"`

	// REQUIRED
	TranscodeRecord int32 `json:"TranscodeRecord"`

	// REQUIRED
	TranscodeSuffixList []string `json:"TranscodeSuffixList"`

	// REQUIRED
	UpdateTime string `json:"UpdateTime"`
}

type DescribeProxyConfigAssociationBody struct {

	// REQUIRED; 账号
	AccountID string `json:"AccountID"`
}

type DescribeProxyConfigAssociationRes struct {

	// REQUIRED
	ResponseMetadata DescribeProxyConfigAssociationResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeProxyConfigAssociationResResult `json:"Result,omitempty"`
}

type DescribeProxyConfigAssociationResResponseMetadata struct {

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

// DescribeProxyConfigAssociationResResult - 视请求的接口而定
type DescribeProxyConfigAssociationResResult struct {

	// REQUIRED; 关联列表
	List []DescribeProxyConfigAssociationResResultListItem `json:"List"`
}

type DescribeProxyConfigAssociationResResultListItem struct {

	// REQUIRED; 账号
	AccountID string `json:"AccountID"`

	// REQUIRED; 代理列表
	ProxyConfigList []DescribeProxyConfigAssociationResResultListPropertiesItemsItem `json:"ProxyConfigList"`
}

type DescribeProxyConfigAssociationResResultListPropertiesItemsItem struct {

	// REQUIRED
	CreateTime string `json:"CreateTime"`

	// REQUIRED; 代理记��值ID
	ID int32 `json:"ID"`

	// REQUIRED
	UpdateTime string `json:"UpdateTime"`
}

type DescribeRecordTaskFileHistoryBody struct {

	// REQUIRED; 开始录制时间，RFC3339 格式的 UTC 时间，精度为 s。当您查询指定录制任务详情时，DateFrom 应设置为开始时间之前的任意时间。
	DateFrom string `json:"DateFrom"`

	// REQUIRED; 结束录制时间，结束时间需晚于 DateFrom，且与 DateFrom 间隔不超过 7天，RFC3339 格式的 UTC 时间，精度为 s。
	DateTo string `json:"DateTo"`

	// REQUIRED; 分页查询页码。
	PageNum int32 `json:"PageNum"`

	// REQUIRED; 单个分页中，查询的结果数量。
	PageSize int32 `json:"PageSize"`

	// 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App *string `json:"App,omitempty"`

	// 流名称，默认查询所有流名称，由 1 到 100 位数字、字母、下划线及"-"和"."组成，如果指定 Stream，必须同时指定 App 的值。
	Stream *string `json:"Stream,omitempty"`

	// 录制文件保存位置。默认取值为 tos。
	// * tos
	// * vod
	Type *string `json:"Type,omitempty"`

	// 域名空间名称。由 1 到 60 位数字、字母、下划线及"-"和"."组成。
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

	// REQUIRED; 分页信息。
	Pagination DescribeRecordTaskFileHistoryResResultPagination `json:"Pagination"`
}

type DescribeRecordTaskFileHistoryResResultDataItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; ToS 存储空间。
	Bucket string `json:"Bucket"`

	// REQUIRED; 录制时长。
	Duration string `json:"Duration"`

	// REQUIRED; 结束录制时间。
	EndTime string `json:"EndTime"`

	// REQUIRED; 录制文件的文件名。
	FileName string `json:"FileName"`

	// REQUIRED; 录制文件存储格式。
	Format string `json:"Format"`

	// REQUIRED; ToS 中的保存路径。
	Path string `json:"Path"`

	// REQUIRED; 开始录制时间。
	StartTime string `json:"StartTime"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 域名空间名称。由 1 到 60 位数字、字母、下划线及"-"和"."组成。
	Vhost string `json:"Vhost"`

	// REQUIRED; 录制文件保存在 VoD 时，录制视频的 ID。
	Vid string `json:"Vid"`
}

// DescribeRecordTaskFileHistoryResResultPagination - 分页信息。
type DescribeRecordTaskFileHistoryResResultPagination struct {

	// REQUIRED; 当前页。
	PageCur int32 `json:"PageCur"`

	// REQUIRED; 当前页的大小。
	PageSize int32 `json:"PageSize"`

	// REQUIRED; 当前页的数据量。
	PageTotal int32 `json:"PageTotal"`

	// REQUIRED; 数据总量。
	TotalCount int32 `json:"TotalCount"`
}

type DescribeRefConfigBody struct {

	// 应用名称
	App *string `json:"App,omitempty"`

	// 域名
	Domain *string `json:"Domain,omitempty"`

	// 配置类型
	Type *string `json:"Type,omitempty"`

	// 域名空间
	Vhost *string `json:"Vhost,omitempty"`
}

type DescribeRefConfigRes struct {

	// REQUIRED
	ResponseMetadata DescribeRefConfigResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeRefConfigResResult          `json:"Result,omitempty"`
}

type DescribeRefConfigResResponseMetadata struct {

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

type DescribeRefConfigResResult struct {

	// REQUIRED; 配置列表
	AssociationList []DescribeRefConfigResResultAssociationListItem `json:"AssociationList"`
}

type DescribeRefConfigResResultAssociationListItem struct {

	// REQUIRED; 应用名称
	App string `json:"App"`

	// REQUIRED; 创建时间
	CreateTime string `json:"CreateTime"`

	// REQUIRED; 域名
	Domain string `json:"Domain"`

	// REQUIRED; 引用名
	Name string `json:"Name"`

	// REQUIRED; 配置类型
	Type string `json:"Type"`

	// REQUIRED; 更新时间
	UpdateTime string `json:"UpdateTime"`

	// REQUIRED; 域名空间
	Vhost string `json:"Vhost"`
}

type DescribeRefererBody struct {

	// 应用名称，默认为所有应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。 :::tip 参数 Domain 和 App 至少传一个。 :::
	App *string `json:"App,omitempty"`

	// 拉流域名。 :::tip
	// * 参数 Domain 和 Vhost 传且仅传一个。
	// * 参数 Domain 和 App 至少传一个。 :::
	Domain *string `json:"Domain,omitempty"`

	// 域名空间名称。 :::tip 参数 Domain 和 Vhost 传且仅传一个。 :::
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

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; Referer 防盗链详情列表。
	RefererInfoList []DescribeRefererResResultRefererListPropertiesItemsItem `json:"RefererInfoList"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`
}

type DescribeRefererResResultRefererListPropertiesItemsItem struct {

	// REQUIRED; 用于标识 referer 防盗链的关键词，返回值为 referer。
	Key string `json:"Key"`

	// REQUIRED; 优先级，当前默认返回值为 0。当多域名返回值一致时，按照域名输入顺序区分，越早加入列表的域名优先级越高。
	Priority int32 `json:"Priority"`

	// REQUIRED; 防盗链类型。
	// * deny：黑名单；
	// * allow：白名单。
	Type string `json:"Type"`

	// REQUIRED; 防盗链规则。
	Value string `json:"Value"`
}

type DescribeRelaySinkBody struct {

	// REQUIRED
	Vhost string  `json:"Vhost"`
	App   *string `json:"App,omitempty"`
	Group *string `json:"Group,omitempty"`
}

type DescribeRelaySinkRes struct {

	// REQUIRED
	ResponseMetadata DescribeRelaySinkResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeRelaySinkResResult `json:"Result,omitempty"`
}

type DescribeRelaySinkResResponseMetadata struct {

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

// DescribeRelaySinkResResult - 视请求的接口而定
type DescribeRelaySinkResResult struct {

	// REQUIRED
	RelaySinkList []DescribeRelaySinkResResultRelaySinkListItem `json:"RelaySinkList"`
}

type DescribeRelaySinkResResultRelaySinkListItem struct {

	// REQUIRED
	App string `json:"App"`

	// REQUIRED
	GroupList []DescribeRelaySinkResResultRelaySinkListPropertiesItemsItem `json:"GroupList"`

	// REQUIRED
	Vhost string `json:"Vhost"`
}

type DescribeRelaySinkResResultRelaySinkListPropertiesItemsFieldRelaySinkDetailListItem struct {

	// REQUIRED
	AK string `json:"AK"`

	// REQUIRED
	CDN string `json:"CDN"`

	// REQUIRED
	CreateTime string `json:"CreateTime"`

	// REQUIRED
	ID string `json:"ID"`

	// REQUIRED
	PullDomainList []DescribeRelaySinkResResultRelaySinkListPropertiesItemsFieldRelaySinkDetailListPropertiesItemsItem `json:"PullDomainList"`

	// REQUIRED
	PushAuth bool `json:"PushAuth"`

	// REQUIRED
	RelaySinkApp string `json:"RelaySinkApp"`

	// REQUIRED
	RelaySinkDomain string `json:"RelaySinkDomain"`

	// REQUIRED; Anything
	RelaySinkParams interface{} `json:"RelaySinkParams"`

	// REQUIRED
	SK string `json:"SK"`

	// REQUIRED
	Status int32 `json:"Status"`

	// REQUIRED
	UpdateTime string `json:"UpdateTime"`

	// REQUIRED
	ValidDuration int32 `json:"ValidDuration"`

	// REQUIRED
	Weight int32 `json:"Weight"`
}

type DescribeRelaySinkResResultRelaySinkListPropertiesItemsFieldRelaySinkDetailListPropertiesItemsItem struct {

	// REQUIRED
	Protocol string `json:"Protocol"`

	// REQUIRED
	PullDomain string `json:"PullDomain"`
}

type DescribeRelaySinkResResultRelaySinkListPropertiesItemsItem struct {

	// REQUIRED
	FieldRelaySinkDetailList []DescribeRelaySinkResResultRelaySinkListPropertiesItemsFieldRelaySinkDetailListItem `json:"FieldRelaySinkDetailList"`

	// REQUIRED
	Group string `json:"Group"`

	// REQUIRED
	IsThunderRelaySink bool `json:"IsThunderRelaySink"`

	// REQUIRED
	IsTranscodeRelaySink bool `json:"IsTranscodeRelaySink"`

	// REQUIRED
	PassRequestParam bool `json:"PassRequestParam"`
}

type DescribeRelaySourceRewriteBody struct {

	// REQUIRED; The domain name space
	Vhost string `json:"Vhost"`

	// The domain name
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

	// Origin address rewrite configurations
	RelaySourceRewriteList *DescribeRelaySourceRewriteResResultRelaySourceRewriteList `json:"RelaySourceRewriteList,omitempty"`
}

// DescribeRelaySourceRewriteResResultRelaySourceRewriteList - Origin address rewrite configurations
type DescribeRelaySourceRewriteResResultRelaySourceRewriteList struct {

	// The domain name space
	Domain *string `json:"Domain,omitempty"`

	// The origin address rewrite configurations
	RewriteRule *DescribeRelaySourceRewriteResResultRelaySourceRewriteListRewriteRule `json:"RewriteRule,omitempty"`

	// The domain name
	Vhost *string `json:"Vhost,omitempty"`
}

// DescribeRelaySourceRewriteResResultRelaySourceRewriteListRewriteRule - The origin address rewrite configurations
type DescribeRelaySourceRewriteResResultRelaySourceRewriteListRewriteRule struct {

	// Whether origin address rewrite is enabled
	Enable *bool `json:"Enable,omitempty"`

	// A list of rewrite rules
	RewriteRuleList []*DescribeRelaySourceRewriteResResultRelaySourceRewriteListRewriteRuleListItem `json:"RewriteRuleList,omitempty"`
}

type DescribeRelaySourceRewriteResResultRelaySourceRewriteListRewriteRuleListItem struct {

	// Whether the query parameters of the original path are included in the new path
	IncludeParams *bool `json:"IncludeParams,omitempty"`

	// The existing path
	OriginPath *string `json:"OriginPath,omitempty"`

	// The new path
	TargetPath *string `json:"TargetPath,omitempty"`
}

type DescribeRelaySourceV3Body struct {

	// REQUIRED; 直播流使用的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console-stable.volcanicengine.com/live/main/domain/list]页面，查看直播流使用的域名。所属的域名空间。
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

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 生效类型（order/rand/hot）
	LBType *string `json:"LBType,omitempty"`

	// 组的重试间隔/s
	RetryInterval *string `json:"RetryInterval,omitempty"`

	// 组的重试次数
	RetryTimes *string `json:"RetryTimes,omitempty"`
}

type DescribeRelaySourceV3ResResultRelaySourceConfigListItemGroupDetailsItemAuthParams struct {

	// 鉴权参数名，如“sign”
	VolcSecret *string `json:"VolcSecret,omitempty"`

	// 有效期，如"expire"
	VolcTime *string `json:"VolcTime,omitempty"`
}

type DescribeRelaySourceV3ResResultRelaySourceConfigListItemGroupDetailsItemServersItemOutboundConfig struct {

	// 代理配置列表，不传默认不使用代理
	ProxyConfigList []*DescribeRelaySourceV3ResResultRelaySourceConfigListPropertiesItemsServersPropertiesItemsItem `json:"ProxyConfigList,omitempty"`

	// 代理模式，0：固定模式，1: 回源模式
	ProxyMode *string `json:"ProxyMode,omitempty"`
}

type DescribeRelaySourceV3ResResultRelaySourceConfigListPropertiesItemsItem struct {

	// REQUIRED; 回源组名称。
	Group string `json:"Group"`

	// REQUIRED; 回源服务器配置列表。
	Servers    []DescribeRelaySourceV3ResResultRelaySourceConfigListPropertiesItemsServersItem    `json:"Servers"`
	AuthParams *DescribeRelaySourceV3ResResultRelaySourceConfigListItemGroupDetailsItemAuthParams `json:"AuthParams,omitempty"`

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

type DescribeRelaySourceV3ResResultRelaySourceConfigListPropertiesItemsServersItem struct {

	// REQUIRED; 回源地址。
	RelaySourceDomain string `json:"RelaySourceDomain"`

	// REQUIRED; 自定义回源参数。
	RelaySourceParams map[string]string `json:"RelaySourceParams"`

	// REQUIRED; 回源协议。
	RelaySourceProtocol string  `json:"RelaySourceProtocol"`
	CreateTime          *string `json:"CreateTime,omitempty"`

	// 回源Host
	Host           *string                                                                                           `json:"Host,omitempty"`
	OutboundConfig *DescribeRelaySourceV3ResResultRelaySourceConfigListItemGroupDetailsItemServersItemOutboundConfig `json:"OutboundConfig,omitempty"`
	UpdateTime     *string                                                                                           `json:"UpdateTime,omitempty"`

	// 权重
	Weight *string `json:"Weight,omitempty"`
}

type DescribeRelaySourceV3ResResultRelaySourceConfigListPropertiesItemsServersPropertiesItemsItem struct {

	// 集群
	Cluster *string `json:"Cluster,omitempty"`

	// 机房
	IDC *string `json:"IDC,omitempty"`

	// 运营商
	ISP *string `json:"ISP,omitempty"`

	// 代理列表
	ProxyList []*DescribeRelaySourceV3ResResultRelaySourceConfigListPropertiesItemsServersPropertiesItemsProxyListItem `json:"ProxyList,omitempty"`
}

type DescribeRelaySourceV3ResResultRelaySourceConfigListPropertiesItemsServersPropertiesItemsProxyListItem struct {

	// 代理地址
	URL *string `json:"URL,omitempty"`

	// 权重
	Weight *string `json:"Weight,omitempty"`
}

type DescribeSDKDetailBody struct {

	// REQUIRED
	ID int32 `json:"ID"`
}

type DescribeSDKDetailRes struct {

	// REQUIRED
	ResponseMetadata DescribeSDKDetailResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeSDKDetailResResult          `json:"Result,omitempty"`
}

type DescribeSDKDetailResResponseMetadata struct {

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

type DescribeSDKDetailResResult struct {

	// REQUIRED
	SDKDetail DescribeSDKDetailResResultSDKDetail `json:"SDKDetail"`
}

type DescribeSDKDetailResResultSDKDetail struct {

	// REQUIRED; 账号
	AccountID string `json:"AccountID"`

	// REQUIRED; 激活时间
	ActiveTime string `json:"ActiveTime"`

	// REQUIRED; 应用ID
	AppID int32 `json:"AppID"`

	// REQUIRED; 应用名称
	AppName string `json:"AppName"`

	// REQUIRED; 应用英文名称
	AppNameEn string `json:"AppNameEn"`

	// REQUIRED; 申请时间
	ApplyTime string `json:"ApplyTime"`

	// REQUIRED; BundleID
	BundleID string `json:"BundleID"`

	// REQUIRED; 创建时间
	CreateTime string `json:"CreateTime"`

	// REQUIRED; 过期时间
	ExpireTime string `json:"ExpireTime"`

	// REQUIRED; sdk记录ID
	ID int32 `json:"ID"`

	// REQUIRED; 证书ID
	LicenseID string `json:"LicenseID"`

	// REQUIRED; License类型，0：无版本，1：基础版本，2：高级版本，3：试用版
	LicenseType int32 `json:"LicenseType"`

	// REQUIRED; License下载地址
	LicenseURL string `json:"LicenseURL"`

	// REQUIRED; 操作时间
	OperateTime string `json:"OperateTime"`

	// REQUIRED; 操作者
	OperateUser string `json:"OperateUser"`

	// REQUIRED; 流量包ID
	PackageID string `json:"PackageID"`

	// REQUIRED; 包名
	PackageName string `json:"PackageName"`

	// REQUIRED; 应用类型，WEB, APP
	SDKType string `json:"SDKType"`

	// REQUIRED; //SDK版本，精简版：1、互动版：2
	SDKVersion int32 `json:"SDKVersion"`

	// REQUIRED; 购买方式，1：人工开通，2：线上购买
	SellType string `json:"SellType"`

	// REQUIRED; 状态，0：未激活，1：已激活，2：审核中，3：已过期
	Status int32 `json:"Status"`
}

type DescribeSDKParamsAvailableBody struct {

	// BundleID
	BundleID *string `json:"BundleID,omitempty"`

	// packageName
	PackageName *string `json:"PackageName,omitempty"`
}

type DescribeSDKParamsAvailableRes struct {

	// REQUIRED
	ResponseMetadata DescribeSDKParamsAvailableResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeSDKParamsAvailableResResult `json:"Result,omitempty"`
}

type DescribeSDKParamsAvailableResResponseMetadata struct {

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

// DescribeSDKParamsAvailableResResult - 视请求的接口而定
type DescribeSDKParamsAvailableResResult struct {

	// REQUIRED; false: 当前bundleID已存在,不可用
	CheckBundleID bool `json:"CheckBundleID"`

	// REQUIRED; false: 当前packageName已存在,不可用
	CheckPackageName bool `json:"CheckPackageName"`
}

type DescribeServiceRes struct {

	// REQUIRED
	ResponseMetadata DescribeServiceResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeServiceResResult          `json:"Result,omitempty"`
}

type DescribeServiceResResponseMetadata struct {

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
	Error   *DescribeServiceResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeServiceResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeServiceResResult struct {

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty"`

	// 当前账号资源限制
	LimitConfig *DescribeServiceResResultLimitConfig `json:"LimitConfig,omitempty"`

	// 上一个状态
	PreStatus *string `json:"PreStatus,omitempty"`

	// 1: 录制是否隐藏TOS 2: 截图是否隐藏TOS 3: 时移是否隐藏VOD 4: 云导播是否隐藏 5：海外加速计费是否隐藏 6：RTM单独加速计费是否隐藏 7：基础版License申请是否隐藏 8：高级版License申请是否隐藏 9：固定回源是否隐藏
	// 10: 月结欠费关停是否处理，1表示处理 11: IP限频是否隐藏 12：URL限频是否隐藏 13：URL参数限频是否隐藏
	// 14：IP访问相同URL限频是否隐藏 15: 活动带宽计费是否隐藏 16: 画质增强是否隐藏 17: Quic加速计费是否隐藏
	PresetConfigHide []*int32 `json:"PresetConfigHide,omitempty"`

	// 审核状态 审批状态
	// * 0: 正常
	// * 1: 未发起
	// * 2: 未审批
	// * 3: 审批未通过
	// * 4：试用
	// * 5：过期
	Status *string `json:"Status,omitempty"`

	// 过期时间
	TrailTime *string `json:"TrailTime,omitempty"`
}

// DescribeServiceResResultLimitConfig - 当前账号资源限制
type DescribeServiceResResultLimitConfig struct {

	// app数量
	AppLimit *int32 `json:"AppLimit,omitempty"`

	// vhost下domain的数量
	DomainLimit *int32 `json:"DomainLimit,omitempty"`

	// stream数量
	StreamLimit *int32 `json:"StreamLimit,omitempty"`

	// 账号下vhost的数量
	VhostLimit *int32 `json:"VhostLimit,omitempty"`
}

type DescribeSnapshotAuditPresetDetailBody struct {
	PresetList []*string `json:"PresetList,omitempty"`
}

type DescribeSnapshotAuditPresetDetailRes struct {

	// REQUIRED
	ResponseMetadata DescribeSnapshotAuditPresetDetailResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeSnapshotAuditPresetDetailResResult `json:"Result,omitempty"`
}

type DescribeSnapshotAuditPresetDetailResResponseMetadata struct {

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

// DescribeSnapshotAuditPresetDetailResResult - 视请求的接口而定
type DescribeSnapshotAuditPresetDetailResResult struct {

	// REQUIRED
	PresetDetailList []DescribeSnapshotAuditPresetDetailResResultPresetDetailListItem `json:"PresetDetailList"`
}

// DescribeSnapshotAuditPresetDetailResResultPresetDetailListItem - 审核模版详细信息。
type DescribeSnapshotAuditPresetDetailResResultPresetDetailListItem struct {

	// REQUIRED
	AccountID string `json:"AccountID"`

	// REQUIRED
	AshePresetName string `json:"AshePresetName"`

	// REQUIRED
	AuditType string `json:"AuditType"`

	// REQUIRED; ToS 存���空间 bucket。 :::tip 参数 Bucket 和 ServiceID 传且仅传一个。 :::
	Bucket string `json:"Bucket"`

	// REQUIRED; 审核结果回调配置。
	CallbackDetailList []DescribeSnapshotAuditPresetDetailResResultPresetDetailListPropertiesItemsItem `json:"CallbackDetailList"`

	// REQUIRED
	CreatedAt int32 `json:"CreatedAt"`

	// REQUIRED; 审核模板描述。
	Description string `json:"Description"`

	// REQUIRED; 审核标签名称，若为空，则默认请求所有基础模型。支持以下取值。
	// * 301：涉黄；
	// * 302：涉敏1；
	// * 303：涉敏2；
	// * 304：广告；
	// * 305：引人不适；
	// * 306：违禁；
	// * 307：二维码；
	// * 308：诈骗；
	// * 309：不良画面；
	// * 310：未成年相关；
	// * 320：文字违规。
	Label []string `json:"Label"`

	// REQUIRED; 审核模板名称。
	PresetName string `json:"PresetName"`

	// REQUIRED; veimageX 的服务 ID。 :::tip 参数 Bucket 和 ServiceID 传且仅传一个。 :::
	ServiceID string `json:"ServiceID"`

	// REQUIRED
	SnapshotConfig DescribeSnapshotAuditPresetDetailResResultPresetDetailListItemSnapshotConfig `json:"SnapshotConfig"`

	// REQUIRED; 存储策略。支持以下取值。
	// * 0：触发存储，只存储有风险图片；
	// * 1：全部存储，存储全部图片。
	StorageStrategy int32 `json:"StorageStrategy"`

	// REQUIRED
	UpdatedAt int32 `json:"UpdatedAt"`
}

type DescribeSnapshotAuditPresetDetailResResultPresetDetailListItemSnapshotConfig struct {

	// REQUIRED
	Format string `json:"Format"`

	// REQUIRED
	Interval float32 `json:"Interval"`

	// REQUIRED
	SnapshotObject string `json:"SnapshotObject"`

	// REQUIRED
	StorageDir string `json:"StorageDir"`
}

type DescribeSnapshotAuditPresetDetailResResultPresetDetailListPropertiesItemsItem struct {

	// REQUIRED; 回调的类型 HTTP。
	CallbackType string `json:"CallbackType"`

	// REQUIRED; 回调地址。
	URL string `json:"URL"`
}

type DescribeStreamQuotaConfigBody struct {

	// REQUIRED; 待查询限额配置的推流域名或拉流域名。
	Domain string `json:"Domain"`
}

type DescribeStreamQuotaConfigRes struct {

	// REQUIRED
	ResponseMetadata DescribeStreamQuotaConfigResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeStreamQuotaConfigResResult `json:"Result,omitempty"`
}

type DescribeStreamQuotaConfigResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                             `json:"Version"`
	Error   *DescribeStreamQuotaConfigResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeStreamQuotaConfigResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

// DescribeStreamQuotaConfigResResult - 视请求的接口而定
type DescribeStreamQuotaConfigResResult struct {

	// REQUIRED; 限额配置列表。
	QuotaList []DescribeStreamQuotaConfigResResultQuotaListItem `json:"QuotaList"`
}

type DescribeStreamQuotaConfigResResultQuotaListItem struct {

	// REQUIRED; 推流域名或拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 配置详情列表。
	QuotaDetailList []DescribeStreamQuotaConfigResResultQuotaListPropertiesItemsItem `json:"QuotaDetailList"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`
}

// DescribeStreamQuotaConfigResResultQuotaListItemQuotaDetailListItemBandwidthConfig - 推流域名的推流路数限额配置信息。 :::tipDomain 为拉流域名时返回。
// :::
type DescribeStreamQuotaConfigResResultQuotaListItemQuotaDetailListItemBandwidthConfig struct {

	// REQUIRED; 拉流带宽限额。
	Quota int32 `json:"Quota"`

	// REQUIRED; 拉流带宽限额的计量单位。
	QuotaUnit string `json:"QuotaUnit"`

	// 拉流带宽限额告警阈值，缺省情况表示未设置告警。
	AlarmThreshold *int32 `json:"AlarmThreshold,omitempty"`

	// 拉流带宽限额告警的计量单位，缺省情况表示不未设置告警。
	AlarmThresholdUnit *string `json:"AlarmThresholdUnit,omitempty"`
}

// DescribeStreamQuotaConfigResResultQuotaListItemQuotaDetailListItemStreamConfig - 推流域名的推流路数限额配置信息。 :::tipDomain 为推流域名时返回。
// :::
type DescribeStreamQuotaConfigResResultQuotaListItemQuotaDetailListItemStreamConfig struct {

	// REQUIRED; 推流路数限额。
	Quota int32 `json:"Quota"`

	// 推流路数限额告警阈值，缺省情况表示未设置告警。
	AlarmThreshold *int32 `json:"AlarmThreshold,omitempty"`
}

type DescribeStreamQuotaConfigResResultQuotaListPropertiesItemsItem struct {

	// 推流域名的推流路数限额配置信息。 :::tipDomain 为拉流域名时返回。 :::
	BandwidthConfig *DescribeStreamQuotaConfigResResultQuotaListItemQuotaDetailListItemBandwidthConfig `json:"BandwidthConfig,omitempty"`

	// 推流域名的推流路数限额配置信息。 :::tipDomain 为推流域名时返回。 :::
	StreamConfig *DescribeStreamQuotaConfigResResultQuotaListItemQuotaDetailListItemStreamConfig `json:"StreamConfig,omitempty"`
}

type DescribeTimeShiftPresetDetailBody struct {

	// REQUIRED; 模板名称列表
	PresetList []string `json:"PresetList"`
}

type DescribeTimeShiftPresetDetailRes struct {

	// REQUIRED
	ResponseMetadata DescribeTimeShiftPresetDetailResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeTimeShiftPresetDetailResResult `json:"Result,omitempty"`
}

type DescribeTimeShiftPresetDetailResResponseMetadata struct {

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

// DescribeTimeShiftPresetDetailResResult - 视请求的接口而定
type DescribeTimeShiftPresetDetailResResult struct {

	// REQUIRED
	PresetDetailList []DescribeTimeShiftPresetDetailResResultPresetDetailListItem `json:"PresetDetailList"`
}

type DescribeTimeShiftPresetDetailResResultPresetDetailListItem struct {

	// REQUIRED; 账号
	AccountID string `json:"AccountID"`

	// REQUIRED; tos bucket
	Bucket string `json:"Bucket"`

	// REQUIRED; 创建时间
	CreatedAt int32 `json:"CreatedAt"`

	// REQUIRED; 删除时间
	DeletedAt int32 `json:"DeletedAt"`

	// REQUIRED; 描述
	Description string `json:"Description"`

	// REQUIRED; 时移保存时间
	MaxShiftTime int32 `json:"MaxShiftTime"`

	// REQUIRED; nss配置
	NssConfig string `json:"NssConfig"`

	// REQUIRED; 模板名称
	Preset string `json:"Preset"`

	// REQUIRED; 时移的类型
	PresetType string `json:"PresetType"`

	// REQUIRED; 分发域名
	PullDomain string `json:"PullDomain"`

	// REQUIRED; 存储的集群
	RPCCluster string `json:"RPCCluster"`

	// REQUIRED; 上传路径
	RecordObject string `json:"RecordObject"`

	// REQUIRED; 状态
	Status int32 `json:"Status"`

	// REQUIRED; 更新时间
	UpdatedAt int32 `json:"UpdatedAt"`

	// REQUIRED; 点播空间名称
	VodNamespace string `json:"VodNamespace"`
}

type DescribeTranscodePresetDetailBody struct {

	// REQUIRED
	PresetList []string `json:"PresetList"`
}

type DescribeTranscodePresetDetailRes struct {

	// REQUIRED
	ResponseMetadata DescribeTranscodePresetDetailResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeTranscodePresetDetailResResult `json:"Result,omitempty"`
}

type DescribeTranscodePresetDetailResResponseMetadata struct {

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

// DescribeTranscodePresetDetailResResult - 视请求的接口而定
type DescribeTranscodePresetDetailResResult struct {

	// REQUIRED
	PresetDetailList DescribeTranscodePresetDetailResResultPresetDetailList `json:"PresetDetailList"`
}

type DescribeTranscodePresetDetailResResultPresetDetailList struct {

	// REQUIRED
	CreatedAt int32 `json:"CreatedAt"`

	// REQUIRED
	UpdatedAt int32   `json:"UpdatedAt"`
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

	// 音频码率，单位为 kbps。
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

	// IDR 帧之间的最大间隔，单位为 s
	GOP    *int32 `json:"GOP,omitempty"`
	GopMin *int32 `json:"GopMin,omitempty"`
	HVSPre *bool  `json:"HVSPre,omitempty"`

	// 视频高度。 :::tip 当As的取值为 0 时，如果Width和Height任意取值为 0，表示保持源流尺寸。 :::
	Height *int32 `json:"Height,omitempty"`

	// 长边长度。 :::tip 当As的取值为 1 时，如果LongSide和ShortSide都取 0，表示保持源流尺寸。 :::
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

	// 模板名称
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

	// 短边长度。 :::tip 当As的取值为 1 时，如果LongSide和ShortSide都取 0，表示保持源流尺寸。 :::
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

	// 视频宽度。 :::tip 当As的取值为 0 时，如果Width和Height任意取值为 0，表示保持源流尺寸。 :::
	Width *int32 `json:"Width,omitempty"`
}

type DescribeVQScoreTaskBody struct {

	// 测评任务ID
	ID *string `json:"ID,omitempty"`
}

type DescribeVQScoreTaskRes struct {

	// REQUIRED
	ResponseMetadata DescribeVQScoreTaskResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeVQScoreTaskResResult          `json:"Result,omitempty"`
}

type DescribeVQScoreTaskResResponseMetadata struct {

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
	Error   *DescribeVQScoreTaskResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeVQScoreTaskResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeVQScoreTaskResResult struct {

	// 拉流地址测评得分详细信息
	AddrScoreList []*DescribeVQScoreTaskResResultAddrScoreListItem `json:"AddrScoreList,omitempty"`

	// 对比拉流地址
	ContrastAddr *string `json:"ContrastAddr,omitempty"`

	// 计算对比拉流地址的画质平均得分
	ContrastAverageScore *float32 `json:"ContrastAverageScore,omitempty"`

	// 主评分与对比评分的差值，（主地址评分平均值-对比地址评分平均值）绝对值
	Difference *float32 `json:"Difference,omitempty"`

	// 主评分与对比评分的差值百分比，主评分与对比评分的差值/min(主地址评分平均值,对比地址评分平均值)*100
	DifferencePer *float32 `json:"DifferencePer,omitempty"`

	// 测评运行时间
	Duration *int32 `json:"Duration,omitempty"`

	// 主拉流地址
	MainAddr *string `json:"MainAddr,omitempty"`

	// 计算主拉流地址平均得分
	MainAverageScore *float32 `json:"MainAverageScore,omitempty"`

	// 计算取值点数
	TotalPointNum *int32 `json:"TotalPointNum,omitempty"`
}

type DescribeVQScoreTaskResResultAddrScoreListItem struct {

	// 拉流地址类型，1：主拉流地址，2：对比拉流地址
	AddrType *int32 `json:"AddrType,omitempty"`

	// 测评得分列表。
	ScoreInfoList []*DescribeVQScoreTaskResResultAddrScoreListPropertiesItemsItem `json:"ScoreInfoList,omitempty"`
}

type DescribeVQScoreTaskResResultAddrScoreListPropertiesItemsItem struct {

	// 测评取值时间点
	PointTime *string `json:"PointTime,omitempty"`

	// 测评得分
	Score *float32 `json:"Score,omitempty"`
}

type DescribeVhostBody struct {

	// REQUIRED; 域名空间名称列表，限制十
	VhostList []string `json:"VhostList"`
}

type DescribeVhostRes struct {

	// REQUIRED
	ResponseMetadata DescribeVhostResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeVhostResResult          `json:"Result,omitempty"`
}

type DescribeVhostResResponseMetadata struct {

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
	Error   *DescribeVhostResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeVhostResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeVhostResResult struct {

	// vhost详情列表
	VhostList []*DescribeVhostResResultVhostListItem `json:"VhostList,omitempty"`
}

type DescribeVhostResResultVhostListItem struct {

	// REQUIRED; 创建时间
	CreateTime string `json:"CreateTime"`

	// REQUIRED; 更新时间
	UpdateTime string `json:"UpdateTime"`

	// 账号
	AccountID *string `json:"AccountID,omitempty"`

	// domain详情列表
	DomainList []*ComponentsAer7PvSchemasDescribevhostresPropertiesResultPropertiesVhostlistItemsPropertiesDomainlistItems `json:"DomainList,omitempty"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty"`

	// 标签列表
	Tags []*DescribeVhostResResultVhostListPropertiesItemsItem `json:"Tags,omitempty"`

	// 域名空间名称
	Vhost *string `json:"Vhost,omitempty"`
}

type DescribeVhostResResultVhostListPropertiesItemsItem struct {

	// 标签类型，System, Custom
	Category *string `json:"Category,omitempty"`

	// 标签key
	Key *string `json:"Key,omitempty"`

	// 标签value
	Value *string `json:"Value,omitempty"`
}

type DescribeWatermarkPresetDetailBody struct {

	// REQUIRED
	PresetList []string `json:"PresetList"`
}

type DescribeWatermarkPresetDetailRes struct {

	// REQUIRED
	ResponseMetadata DescribeWatermarkPresetDetailResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeWatermarkPresetDetailResResult `json:"Result,omitempty"`
}

type DescribeWatermarkPresetDetailResResponseMetadata struct {

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

// DescribeWatermarkPresetDetailResResult - 视请求的接口而定
type DescribeWatermarkPresetDetailResResult struct {

	// REQUIRED
	PresetDetailList []DescribeWatermarkPresetDetailResResultPresetDetailListItem `json:"PresetDetailList"`
}

type DescribeWatermarkPresetDetailResResultPresetDetailListItem struct {

	// REQUIRED
	AccountID string `json:"AccountID"`

	// REQUIRED
	App string `json:"App"`

	// REQUIRED
	CreatedAt int32 `json:"CreatedAt"`

	// REQUIRED
	ID int32 `json:"ID"`

	// REQUIRED
	Name string `json:"Name"`

	// REQUIRED
	Orientation string `json:"Orientation"`

	// REQUIRED
	Picture string `json:"Picture"`

	// REQUIRED
	PictureKey string `json:"PictureKey"`

	// REQUIRED
	PictureURL string `json:"PictureURL"`

	// REQUIRED
	PosX float32 `json:"PosX"`

	// REQUIRED
	PosY float32 `json:"PosY"`

	// REQUIRED
	PreviewHeight string `json:"PreviewHeight"`

	// REQUIRED
	PreviewWidth string `json:"PreviewWidth"`

	// REQUIRED
	RelativeHeight string `json:"RelativeHeight"`

	// REQUIRED
	RelativeWidth string `json:"RelativeWidth"`

	// REQUIRED
	Stream string `json:"Stream"`

	// REQUIRED
	UpdatedAt int32 `json:"UpdatedAt"`

	// REQUIRED
	Vhost string `json:"Vhost"`
}

type DisAssociatePresetBody struct {

	// REQUIRED; 域名空间名称
	Vhost string `json:"Vhost"`

	// app名称
	App *string `json:"App,omitempty"`

	// 模板名称
	PresetName *string `json:"PresetName,omitempty"`

	// 模板类型
	PresetType *string `json:"PresetType,omitempty"`

	// 流名称
	Stream *string `json:"Stream,omitempty"`
}

type DisAssociatePresetRes struct {

	// REQUIRED
	ResponseMetadata DisAssociatePresetResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DisAssociatePresetResResponseMetadata struct {

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

type DisableAuthBody struct {

	// REQUIRED
	Domain string `json:"Domain"`

	// REQUIRED
	Vhost     string  `json:"Vhost"`
	App       *string `json:"App,omitempty"`
	SceneType *string `json:"SceneType,omitempty"`
}

type DisableAuthRes struct {

	// REQUIRED
	ResponseMetadata DisableAuthResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type DisableAuthResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version   string                               `json:"Version"`
	Error     *DisableAuthResResponseMetadataError `json:"Error,omitempty"`
	RequestID *string                              `json:"RequestID,omitempty"`
}

type DisableAuthResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DisableDomainBody struct {

	// REQUIRED; 待禁用域名。
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

type DisassociateRefConfigBody struct {

	// REQUIRED; 引用名，若不存在该引用绑定会直接报错
	Name string `json:"Name"`

	// 应用名称。 :::tip App 与 Domain 二选一填 :::
	App *string `json:"App,omitempty"`

	// 域名。 :::tip App 与 Domain 二选一填 :::
	Domain *string `json:"Domain,omitempty"`

	// 域名空间名称。
	Vhost *string `json:"Vhost,omitempty"`
}

type DisassociateRefConfigRes struct {

	// REQUIRED
	ResponseMetadata DisassociateRefConfigResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DisassociateRefConfigResResponseMetadata struct {

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

type EnableAuthBody struct {

	// REQUIRED
	Domain string `json:"Domain"`

	// REQUIRED
	Vhost     string  `json:"Vhost"`
	App       *string `json:"App,omitempty"`
	SceneType *string `json:"SceneType,omitempty"`
}

type EnableAuthRes struct {

	// REQUIRED
	ResponseMetadata EnableAuthResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type EnableAuthResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version   string                              `json:"Version"`
	Error     *EnableAuthResResponseMetadataError `json:"Error,omitempty"`
	RequestID *string                             `json:"RequestID,omitempty"`
}

type EnableAuthResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type EnableDomainBody struct {

	// REQUIRED; 待启用域名。
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

	// REQUIRED; Whether to enable the configuration.
	// * true: Enable
	// * false: Disable
	Enable bool `json:"Enable"`

	// REQUIRED; The type of HTTP header to be enabled or disabled:
	// * 0: In the response sent from an edge server to a client
	// * 1: In the request sent to a third-party origin server during an origin-pull task.
	Phase int32 `json:"Phase"`

	// REQUIRED; The domain name space.
	Vhost string `json:"Vhost"`

	// The domain name.
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

	// 直播流使用的域名，您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理 [https://console-stable.volcanicengine.com/live/main/domain/list]页面，查看待禁推的直播流使用的域名。
	// :::tip 参数 Domain 和
	// Vhost 传且仅传一个。 :::
	Domain *string `json:"Domain,omitempty"`

	// 禁推的结束时间，RFC3339 格式的 UTC 时间，精度为毫秒，禁推有效期最长为 90 天，默认为当前时间加 90 天。
	EndTime *string `json:"EndTime,omitempty"`

	// 域名空间，即直播流地址的域名（Domain）所属的域名空间（Vhost）。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console-stable.volcanicengine.com/live/main/domain/list]
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

	// REQUIRED; 应用名称，默认为所有应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// 过期时间，拉流地址的有效时间，过期后需要重新生成。RFC3339 格式的 UTC 时间，精度为秒，缺省情况下表示 7 天。 :::tip 如果同时设置 ValidDuration 和 ExpiredTime，以 ExpiredTime 的时间为准。
	// :::
	ExpiredTime *string `json:"ExpiredTime,omitempty"`

	// 转码流后缀，不传默认为空，可通过调用 ListVhostTransCodePreset [https://www.volcengine.com/docs/6469/1126853] 接口查询。
	Suffix *string `json:"Suffix,omitempty"`

	// CDN 类型，默认值为 fcdn，支持如下取值。
	// * fcdn：火山引擎流媒体直播 CDN；
	// * 3rd：第三方 CDN。
	Type *string `json:"Type,omitempty"`

	// 有效时长，拉流地址的有效时间，过期后需要重新生成。单位为秒，取值 ﹥0，缺省情况下表示 7 天。 :::tip 如果同时设置 ValidDuration 和 ExpiredTime，以 ExpiredTime 的时间为准。 :::
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

	// REQUIRED; 协议类型，包括 hls、flv 和 rtmp。
	Protocol string `json:"Protocol"`

	// REQUIRED; 地址类型，可能的值为：
	// * push：推流；
	// * pull：拉流；
	// * 3rdplay(relaysource)：第三方回源；
	// * 3rdplay(relaysink)：第三方转推。
	Type string `json:"Type"`

	// REQUIRED; 生成的拉流地址。
	URL string `json:"URL"`
}

type GeneratePushURLBody struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 推流域名名称，需要推流地址的域名，不填返回Vhost下所有推流域名生成的地址。
	Domain *string `json:"Domain,omitempty"`

	// 过期时间，推流地址的有效时间，过期后需要重新生成。RFC3339 格式的 UTC 时间，精度为秒，缺省情况下表示 7 天。 :::tip 如果同时设置 ValidDuration 和 ExpiredTime，以 ExpiredTime 的时间为准。
	// :::
	ExpiredTime *string `json:"ExpiredTime,omitempty"`

	// 有效时长，推流地址的有效时间，过期后需要重新生成。单位为秒，取值 ﹥0，缺省情况下表示 7 天。 :::tip 如果同时设置 ValidDuration 和 ExpiredTime，以 ExpiredTime 的时间为准。 :::
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

type GenerateTimeShiftPlayURLBody struct {

	// REQUIRED
	ExpireTime string `json:"ExpireTime"`

	// REQUIRED
	TimeShiftURL string `json:"TimeShiftURL"`

	// REQUIRED
	VodSpacename string `json:"VodSpacename"`
}

type GenerateTimeShiftPlayURLRes struct {

	// REQUIRED
	ResponseMetadata GenerateTimeShiftPlayURLResResponseMetadata `json:"ResponseMetadata"`
	Result           *GenerateTimeShiftPlayURLResResult          `json:"Result,omitempty"`
}

type GenerateTimeShiftPlayURLResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version   string                                            `json:"Version"`
	Error     *GenerateTimeShiftPlayURLResResponseMetadataError `json:"Error,omitempty"`
	RequestID *string                                           `json:"RequestID,omitempty"`
}

type GenerateTimeShiftPlayURLResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type GenerateTimeShiftPlayURLResResult struct {
	PlayURLList []*GenerateTimeShiftPlayURLResResultPlayURLListItem `json:"PlayURLList,omitempty"`
}

type GenerateTimeShiftPlayURLResResultPlayURLListItem struct {
	BackupPlayURL *string `json:"BackupPlayURL,omitempty"`
	MainPlayURL   *string `json:"MainPlayURL,omitempty"`
	URLExpire     *string `json:"URLExpire,omitempty"`
}

type GetAppsRes struct {

	// REQUIRED
	ResponseMetadata GetAppsResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetAppsResResult `json:"Result,omitempty"`
}

type GetAppsResResponseMetadata struct {

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

// GetAppsResResult - 视请求的接口而定
type GetAppsResResult struct {

	// REQUIRED; app列表
	Apps []GetAppsResResultAppsItem `json:"Apps"`
}

type GetAppsResResultAppsItem struct {

	// REQUIRED; app中文名称
	AppCnName string `json:"AppCnName"`

	// REQUIRED; app英文名称
	AppEnName string `json:"AppEnName"`

	// REQUIRED; APPID
	AppID int32 `json:"AppID"`

	// REQUIRED; bundleID
	BundleID string `json:"BundleID"`

	// REQUIRED; packageName
	PackageName string `json:"PackageName"`
}

type GetPullCDNSnapshotTaskBody struct {

	// REQUIRED; 直播截图任务 ID，任务的唯一标识。
	TaskID string `json:"TaskID"`
}

type GetPullCDNSnapshotTaskRes struct {

	// REQUIRED
	ResponseMetadata GetPullCDNSnapshotTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 视请求的接口而定
	Result GetPullCDNSnapshotTaskResResult `json:"Result"`
}

type GetPullCDNSnapshotTaskResResponseMetadata struct {

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

// GetPullCDNSnapshotTaskResResult - 视请求的接口而定
type GetPullCDNSnapshotTaskResResult struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 推流/拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 截图任务的结束时间。
	EndTime string `json:"EndTime"`

	// REQUIRED; 截图任务的开始时间。
	StartTime string `json:"StartTime"`

	// REQUIRED; 任务状态：
	// * 停用
	// * 未开始
	// * 生效中
	// * 已结束
	Status string `json:"Status"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 直播截图任务 ID。
	TaskID string `json:"TaskId"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`
}

type GetPullRecordTaskBody struct {

	// REQUIRED; 直播录制任务的 ID，任务的唯一标识。
	TaskID string `json:"TaskID"`
}

type GetPullRecordTaskRes struct {

	// REQUIRED
	ResponseMetadata GetPullRecordTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 视请求的接口而定
	Result GetPullRecordTaskResResult `json:"Result"`
}

type GetPullRecordTaskResResponseMetadata struct {

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

// GetPullRecordTaskResResult - 视请求的接口而定
type GetPullRecordTaskResResult struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 推流/拉流域名
	Domain string `json:"Domain"`

	// REQUIRED; 任务结束时间
	EndTime string `json:"EndTime"`

	// REQUIRED; 任务开始时间。
	StartTime string `json:"StartTime"`

	// REQUIRED; 任务状态，有以下几种状态：
	// * 停用
	// * 未开始
	// * 生效中
	// * 已结束
	Status string `json:"Status"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 直播录制任务的 ID。
	TaskID string `json:"TaskId"`

	// REQUIRED; 域名空间
	Vhost string `json:"Vhost"`
}

type GetTagsRes struct {

	// REQUIRED
	ResponseMetadata GetTagsResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetTagsResResult `json:"Result,omitempty"`
}

type GetTagsResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                           `json:"Version"`
	Error   *GetTagsResResponseMetadataError `json:"Error,omitempty"`
}

type GetTagsResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

// GetTagsResResult - 视请求的接口而定
type GetTagsResResult struct {

	// REQUIRED; 标签列表
	Tags []GetTagsResResultTagsItem `json:"Tags"`
}

type GetTagsResResultTagsItem struct {

	// REQUIRED; 标签类型：Custom, System
	Category string `json:"Category"`

	// REQUIRED; 标签Key
	TagKey string `json:"TagKey"`

	// REQUIRED; 标签value
	TagValue string `json:"TagValue"`
}

type GetVqosRawDataBody struct {

	// REQUIRED; 维度可为该service下任意一个维度
	Dimensions []string `json:"Dimensions"`

	// REQUIRED; 同上
	End int32 `json:"End"`

	// REQUIRED; 过滤条件
	Filter GetVqosRawDataBodyFilter `json:"Filter"`

	// REQUIRED; Limit取值上限为20000
	Limit int32 `json:"Limit"`

	// REQUIRED; line时表示图例数，point时表示数据点数
	LimitType string `json:"LimitType"`

	// REQUIRED; 指标可为该service下任意一个指标
	Metrics []string `json:"Metrics"`

	// REQUIRED; 默认0
	Offset int32 `json:"Offset"`

	// REQUIRED; UNIX时间戳，单位：s
	Start int32 `json:"Start"`

	// REQUIRED; 支持5m、1h、1d
	Window string `json:"Window"`
}

// GetVqosRawDataBodyFilter - 过滤条件
type GetVqosRawDataBodyFilter struct {

	// REQUIRED; Logic有值时，Filters必须有值
	Filters []GetVqosRawDataBodyFilterFiltersItem `json:"Filters"`

	// REQUIRED; 取值为空、"and"或者"or"
	Logic string `json:"Logic"`
}

type GetVqosRawDataBodyFilterFiltersItem struct {

	// 筛选项名称
	Field *string `json:"Field,omitempty"`

	// 操作符号
	Op *string `json:"Op,omitempty"`

	// Logic为空时生效，多值时使用英文逗号拼接
	Value *string `json:"Value,omitempty"`
}

type GetVqosRawDataQuery struct {

	// REQUIRED
	AppQueryType string `json:"AppQueryType" query:"AppQueryType"`

	// REQUIRED
	VqosService string `json:"VqosService" query:"VqosService"`
}

type GetVqosRawDataRes struct {

	// REQUIRED
	ResponseMetadata GetVqosRawDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result GetVqosRawDataResResult `json:"Result"`
}

type GetVqosRawDataResResponseMetadata struct {

	// REQUIRED
	RequestID string `json:"RequestId"`
}

type GetVqosRawDataResResult struct {

	// REQUIRED
	Columns []interface{} `json:"Columns"`

	// REQUIRED
	Data []interface{} `json:"Data"`

	// REQUIRED
	Limit int32 `json:"Limit"`

	// REQUIRED
	Offset int32 `json:"Offset"`

	// REQUIRED
	Total int32 `json:"Total"`

	// REQUIRED
	TotalPoint int32 `json:"TotalPoint"`
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

type ListActionHistoryBody struct {

	// REQUIRED; 页码
	PageNum int32 `json:"PageNum"`

	// REQUIRED; 每页数量
	PageSize int32 `json:"PageSize"`

	// 账号
	AccountID *string `json:"AccountID,omitempty"`

	// 应用名称
	App *string `json:"App,omitempty"`

	// 配置项名称
	ConfigName *string `json:"ConfigName,omitempty"`

	// 配置项名称英文
	ConfigNameEn *string `json:"ConfigNameEn,omitempty"`

	// 配置平台：Vadmin, 火山引擎控制台
	ConfigPlatform *string `json:"ConfigPlatform,omitempty"`

	// 域名
	Domain *string `json:"Domain,omitempty"`

	// 查询结束时间
	EndTime *int32 `json:"EndTime,omitempty"`

	// 模板名称
	PresetName *string `json:"PresetName,omitempty"`

	// 查询开始时间
	StartTime *int32 `json:"StartTime,omitempty"`

	// 操作人ID
	UserID *string `json:"UserID,omitempty"`

	// 域名空间
	Vhost *string `json:"Vhost,omitempty"`
}

type ListActionHistoryRes struct {

	// REQUIRED
	ResponseMetadata ListActionHistoryResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *ListActionHistoryResResult `json:"Result,omitempty"`
}

type ListActionHistoryResResponseMetadata struct {

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

// ListActionHistoryResResult - 视请求的接口而定
type ListActionHistoryResResult struct {

	// REQUIRED
	ActionHistoryList []ListActionHistoryResResultActionHistoryListItem `json:"ActionHistoryList"`

	// REQUIRED; 总共的数量
	Total int32 `json:"Total"`
}

type ListActionHistoryResResultActionHistoryListItem struct {

	// REQUIRED; 账号
	AccountID string `json:"AccountID"`

	// REQUIRED; 接口名称
	Action string `json:"Action"`

	// REQUIRED; 变更状态
	ActionStatus string `json:"ActionStatus"`

	// REQUIRED; 配置时间
	ActionTime string `json:"ActionTime"`

	// REQUIRED; 应用名称
	App string `json:"App"`

	// REQUIRED; 变更工单ID
	ApplicationID string `json:"ApplicationID"`

	// REQUIRED; 变更平台链接
	ApplicationURL string `json:"ApplicationURL"`

	// REQUIRED; 操作内容
	Body string `json:"Body"`

	// REQUIRED; 配置项名称
	ConfigName string `json:"ConfigName"`

	// REQUIRED; 配置项名称英文
	ConfigNameEn string `json:"ConfigNameEn"`

	// REQUIRED; 配置平台
	ConfigPlatform string `json:"ConfigPlatform"`

	// REQUIRED; 配置进度
	ConfigProgress float32 `json:"ConfigProgress"`

	// REQUIRED; 域名
	Domain string `json:"Domain"`

	// REQUIRED; 记录ID
	ID string `json:"ID"`

	// REQUIRED; 模板名称
	PresetName string `json:"PresetName"`

	// REQUIRED; 流名
	Stream string `json:"Stream"`

	// REQUIRED; 操作人
	UserID string `json:"UserID"`

	// REQUIRED; 域名空间
	Vhost string `json:"Vhost"`
}

type ListBindEncryptDRMBody struct {

	// REQUIRED; 域名空间
	Vhost string `json:"Vhost"`

	// app
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
	DRMBindingList []*ListBindEncryptDRMResResultDRMBindingListItem `json:"DRMBindingList,omitempty"`
}

type ListBindEncryptDRMResResultDRMBindingListItem struct {

	// REQUIRED; 账号
	AccountID string `json:"AccountID"`

	// REQUIRED; app
	App string `json:"App"`

	// REQUIRED; drm是否开启
	Enable bool `json:"Enable"`

	// REQUIRED; 进行drm加密的转码流的转码后缀
	EncryptTranscodeList []string `json:"EncryptTranscodeList"`

	// REQUIRED; 域名空间
	Vhost string `json:"Vhost"`
}

type ListCertBindInfoBody struct {

	// REQUIRED; 页码，x >=1
	PageNum int32 `json:"PageNum"`

	// REQUIRED; 一页数量，0 < x <= 1000
	PageSize int32 `json:"PageSize"`

	// 域名
	Domain *string `json:"Domain,omitempty"`

	// 域名类型，push，pull-flv
	DomainTypeList []*string `json:"DomainTypeList,omitempty"`

	// 过滤HTTPS,0:返回非HTTPS数据，1：返回HTTPS数据，2：返回所有数据，默认不填为1
	HTTPS *int32 `json:"HTTPS,omitempty"`

	// 域名空间
	Vhost *string `json:"Vhost,omitempty"`
}

type ListCertBindInfoRes struct {

	// REQUIRED
	ResponseMetadata ListCertBindInfoResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *ListCertBindInfoResResult `json:"Result,omitempty"`
}

type ListCertBindInfoResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                    `json:"Version"`
	Error   *ListCertBindInfoResResponseMetadataError `json:"Error,omitempty"`
}

type ListCertBindInfoResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

// ListCertBindInfoResResult - 视请求的接口而定
type ListCertBindInfoResResult struct {
	CertBindList []*ListCertBindInfoResResultCertBindListItem `json:"CertBindList,omitempty"`
}

type ListCertBindInfoResResultCertBindListItem struct {

	// 账号
	AccountID *string `json:"AccountID,omitempty"`

	// 证书域名列表
	CertDomainList []*string `json:"CertDomainList,omitempty"`

	// 证书实例ID
	CertID *string `json:"CertID,omitempty"`

	// 证书名称
	CertName *string `json:"CertName,omitempty"`

	// 证书链ID
	ChainID *string `json:"ChainID,omitempty"`

	// 域名绑定的域名
	Domain *string `json:"Domain,omitempty"`

	// 域名类型
	DomainType *string `json:"DomainType,omitempty"`

	// 是否https
	HTTPS *bool `json:"HTTPS,omitempty"`

	// 有效期
	NotAfter *string `json:"NotAfter,omitempty"`

	// 生效时间
	NotBefore *string `json:"NotBefore,omitempty"`

	// 域名空间
	Vhost *string `json:"Vhost,omitempty"`
}

type ListCertBody struct {

	// 火山引擎账号 ID
	AccountID *string `json:"AccountID,omitempty"`

	// 是否启用证书，默认值为 true，支持的取值包括：
	// * true：启用证书；
	// * false：禁用证书。
	Available *bool   `json:"Available,omitempty"`
	Domain    *string `json:"Domain,omitempty"`

	// 证书是否过期，默认值为 false，支持的取值包括：
	// * false：表示证书未过期；
	// * true：表示证书已过期。
	Expiring *bool `json:"Expiring,omitempty"`
}

type ListCertRes struct {

	// REQUIRED
	ResponseMetadata ListCertResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListCertResResult          `json:"Result,omitempty"`
}

type ListCertResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version   string                            `json:"Version"`
	Error     *ListCertResResponseMetadataError `json:"Error,omitempty"`
	RequestID *string                           `json:"RequestID,omitempty"`
}

type ListCertResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type ListCertResResult struct {

	// 证书列表
	CertList []*ListCertResResultCertListItem `json:"CertList,omitempty"`
}

type ListCertResResultCertListItem struct {
	AccountID *string `json:"AccountID,omitempty"`

	// 证书域名
	CertDomain *string `json:"CertDomain,omitempty"`

	// 证书名称
	CertName *string `json:"CertName,omitempty"`

	// 证书 ID
	ChainID    *string `json:"ChainID,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty"`

	// 证书的过期时间，RFC3339 格式的 UTC 时间，精度为 s
	NotAfter *string `json:"NotAfter,omitempty"`

	// 证书的生效日期，RFC3339 格式的 UTC 时间，精度为 s
	NotBefore *string `json:"NotBefore,omitempty"`

	// 证书状态，由证书管理平台返回，支持的取值如下所示。
	// * OK：正常；
	// * Expire：过期；
	// * 30days：有效期剩余 30 天；
	// * 15days：有效期剩余 15 天；
	// * 7days：有效期剩余 7 天；
	// * 1days：有效期剩余 1 天。
	Status *string `json:"Status,omitempty"`

	// 状态码，由证书管理平台返回，支持的取值与对应的 状态如下所示。
	// * 0：OK；
	// * 1：Expire；
	// * 7：30days；
	// * 10：15days；
	// * 8：7days；
	// * 9：1days。
	StatusCode *int32 `json:"StatusCode,omitempty"`
}

type ListCertV2Body struct {

	// 火山引擎账号 ID
	AccountID *string `json:"AccountID,omitempty"`

	// 是否启用证书，默认值为 true，支持的取值包括：
	// * true：启用证书；
	// * false：禁用证书。
	Available *bool `json:"Available,omitempty"`

	// 证书名称，支持输入证书名称中的关键字，进行模糊查询
	CertName *string `json:"CertName,omitempty"`

	// 域名，查询该域名对应的证书，支持精确查询
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

	// REQUIRED; 证书实例ID。
	CertID string `json:"CertID"`

	// REQUIRED; 证书名称。
	CertName string `json:"CertName"`

	// REQUIRED; 证书链 ID。
	ChainID string `json:"ChainID"`

	// REQUIRED; 火山证书中心证书链 ID。
	ChainIDVolc string `json:"ChainIDVolc"`

	// REQUIRED; 证书的过期时间，RFC3339 格式的 UTC 时间，精度为 s。
	NotAfter string `json:"NotAfter"`

	// REQUIRED; 证书的生效日期，RFC3339 格式的 UTC 时间，精度为 s。
	NotBefore string `json:"NotBefore"`

	// REQUIRED; 证书状态，由证书管理平台返回，支持的取值如下所示。
	// * OK：正常；
	// * Expire：过期；
	// * 30days：有效期剩余 30 天；
	// * 15days：有效期剩余 15 天；
	// * 7days：有效期剩余 7 天；
	// * 1days：有效期剩余 1 天。
	Status    string  `json:"Status"`
	AccountID *string `json:"AccountID,omitempty"`
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
	// * copy：不进行转��，所有音频编码参数不生效；
	// * opus：使用 Opus 编码格式。
	Acodec         *string `json:"Acodec,omitempty"`
	AdvancedParam  *string `json:"AdvancedParam,omitempty"`
	AllowAudioCopy *int32  `json:"AllowAudioCopy,omitempty"`
	AllowVideoCopy *int32  `json:"AllowVideoCopy,omitempty"`
	An             *int32  `json:"An,omitempty"`

	// 宽高自适应模式开关，支持的取值及含义如下。
	// * 0：关闭宽高自适应
	// * 1：开启宽高自适应 :::tip
	// * 关闭宽高自适应时，转码配置分辨率取 Width 和 Height 的值对转码视频进行拉伸；
	// * 开启宽高自适应时，转码配置分辨率按照 ShortSide 、 LongSide 、Width 、Height 的优先级取值，另一边等比缩放。 :::
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

	// 2 个参考帧之间的最大 B 帧数。BFrames取 0 时，表示去 B 帧。
	BFrames  *int32  `json:"BFrames,omitempty"`
	Describe *string `json:"Describe,omitempty"`

	// 帧率，单位为 fps。帧率越大，画面越流畅。
	FPS *int32 `json:"FPS,omitempty"`

	// IDR 帧之间的最大间隔，单位为秒。
	GOP    *int32 `json:"GOP,omitempty"`
	GopMin *int32 `json:"GopMin,omitempty"`
	HVSPre *bool  `json:"HVSPre,omitempty"`

	// 视频高度。 :::tip
	// * 当 As 的取值为 0 即关闭宽高自适应时，转码分辨率将取 Width 和 Height 的值对转码视频进行拉伸；
	// * Width 和 Height 任一配置为 0 时，转码视频将保持源流尺寸；
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

	// 是否极智超清转码，取值及含义如下。
	// * true：极智超清转码；
	// * false：标准转码。
	Roi  *bool `json:"Roi,omitempty"`
	SITI *bool `json:"SITI,omitempty"`

	// 短边长度。 :::tip
	// * 当 As 的取值为 1 即开启宽高自适应时，参数生效，反之则不生效。
	// * 当 As 的取值为 1 时，如果 LongSide 、 ShortSide 、Width 、Height 同时取 0，表示保持源流尺寸。 :::
	ShortSide    *int32 `json:"ShortSide,omitempty"`
	Status       *int32 `json:"Status,omitempty"`
	StopInterval *int32 `json:"StopInterval,omitempty"`

	// 转码流后缀名。
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

	// 视频宽度。 :::tip
	// * 当 As 的取值为 0 即关闭宽高自适应时，转码分辨率将取 Width 和 Height 的值对转码视频进行拉伸；
	// * Width 和 Height 任一配置为 0 时，转码视频将保持源流尺寸；
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

	// REQUIRED; 当前页码，取值范围为 [1,1000]。
	PageNum int32 `json:"PageNum"`

	// REQUIRED; 分页大小，取值范围为 [1, 1000]
	PageSize int32 `json:"PageSize"`

	// 域名名称列表，缺省情况下表示全部。
	DomainNameList []*string `json:"DomainNameList,omitempty"`

	// 域名区域列表。缺省情况下表示全部，区域类型支持以下取值。
	// * cn：中国大陆；
	// * cn-global：全球；
	// * cn-oversea：海外及港澳台。
	DomainRegionList []*string `json:"DomainRegionList,omitempty"`

	// 域名状态列表。缺省情况下表示全部。支持的取值如下所示。
	// * 0：正常；
	// * 1：审核中；
	// * 2：禁用，禁止使用，此时 domain 不生效；
	// * 3：删除；
	// * 4：审核被驳回。审核不通过，需要重新创建并审核；
	// * 5：欠费关停。
	DomainStatusList []*int32 `json:"DomainStatusList,omitempty"`

	// 域名类型列表。缺省情况下表示全部。域名类型支持以下取值。
	// * push：推流域名；
	// * pull-flv：拉流域名。
	DomainTypeList []*string `json:"DomainTypeList,omitempty"`

	// 域名空间列表，缺省情况下表示全部。
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

	// REQUIRED; 所绑定证书支持的泛域名。
	CertDomain string `json:"CertDomain"`

	// REQUIRED; 绑定的证书名称。
	CertName string `json:"CertName"`

	// REQUIRED; 绑定的证书信息。
	ChainID string `json:"ChainID"`

	// REQUIRED; CNAME 状态。
	// * 0：未配置 CNAME；
	// * 1：已配置 CNAME。
	CnameCheck int32 `json:"CnameCheck"`

	// REQUIRED; 创建时间。
	CreateTime string `json:"CreateTime"`

	// REQUIRED; 推/拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 域名是否可用的状态。
	// * 0：正常，域名为可用状态；
	// * 1：配置中，域名为可用状态；
	// * 2：不可用，域名为其他的不可用状态。
	DomainCheck int32 `json:"DomainCheck"`

	// REQUIRED; ICP 备案校验是否通过，是否过期信息。
	ICPCheck int32 `json:"ICPCheck"`

	// REQUIRED; 项目名称。
	ProjectName string `json:"ProjectName"`

	// REQUIRED; 绑定的推流域名。
	PushDomain string `json:"PushDomain"`

	// REQUIRED; 区域，包含以下类型。
	// * cn：中国大陆；
	// * cn-global：全球；
	// * cn-oversea：海外及港澳台。
	Region string `json:"Region"`

	// REQUIRED; 域名状态。状态说明如下所示。
	// * 0：正常；
	// * 1：审核中；
	// * 2：禁用，禁止使用，此时 domain 不生效；
	// * 3：删除；
	// * 4：审核被驳回。审核不通过，需要重新创建并审核；
	// * 5：欠费关停。
	Status int32 `json:"Status"`

	// REQUIRED; 标签信息。
	Tags []ListDomainDetailResResultDomainListPropertiesItemsItem `json:"Tags"`

	// REQUIRED; 域名类型，包含两种类型。
	// * push：推流域名；
	// * pull-flv：拉流域名，包含 RTMP、FLV、HLS 格式。
	Type string `json:"Type"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`
}

type ListDomainDetailResResultDomainListPropertiesItemsItem struct {

	// REQUIRED; 标签类型。
	// * System：系统内置标签
	// * Custom：自定义标签
	Category string `json:"Category"`

	// REQUIRED; 标签 Key 值。
	Key string `json:"Key"`

	// REQUIRED; 标签 Value 值。
	Value string `json:"Value"`
}

type ListHeaderEnumBody struct {

	// REQUIRED
	Phase int32 `json:"Phase"`
}

type ListHeaderEnumRes struct {

	// REQUIRED
	ResponseMetadata ListHeaderEnumResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *ListHeaderEnumResResult `json:"Result,omitempty"`
}

type ListHeaderEnumResResponseMetadata struct {

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

// ListHeaderEnumResResult - 视请求的接口而定
type ListHeaderEnumResResult struct {

	// REQUIRED; 常量列表
	ConstantList []ListHeaderEnumResResultConstantListItem `json:"ConstantList"`

	// REQUIRED; 变量列表
	VariableList []ListHeaderEnumResResultVariableListItem `json:"VariableList"`
}

type ListHeaderEnumResResultConstantListItem struct {

	// REQUIRED; 常量名
	Name string `json:"Name"`

	// REQUIRED; 提示词
	Prompt string `json:"Prompt"`
}

type ListHeaderEnumResResultVariableListItem struct {

	// REQUIRED
	Name string `json:"Name"`

	// REQUIRED
	Prompt string `json:"Prompt"`
}

type ListInstanceBody struct {

	// REQUIRED
	PageNum int32 `json:"PageNum"`

	// REQUIRED; 最大100
	PageSize int32 `json:"PageSize"`

	// REQUIRED; 以,分割
	Status string `json:"Status"`
}

type ListInstanceRes struct {

	// REQUIRED
	ResponseMetadata ListInstanceResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result ListInstanceResResult `json:"Result"`
}

type ListInstanceResResponseMetadata struct {

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

type ListInstanceResResult struct {

	// REQUIRED
	List []ListInstanceResResultListItem `json:"List"`

	// REQUIRED
	Total int32 `json:"Total"`
}

type ListInstanceResResultListItem struct {
	AccountID         *string `json:"AccountID,omitempty"`
	ConfigurationCode *string `json:"ConfigurationCode,omitempty"`
	ConfigurationName *string `json:"ConfigurationName,omitempty"`
	ExpireTime        *string `json:"ExpireTime,omitempty"`
	Number            *string `json:"Number,omitempty"`
	Product           *string `json:"Product,omitempty"`
	Status            *int32  `json:"Status,omitempty"`
	SubOrderNO        *string `json:"SubOrderNO,omitempty"`
	TerminateValid    *bool   `json:"TerminateValid,omitempty"`
	Type              *int32  `json:"Type,omitempty"`
}

type ListObjectBody struct {

	// REQUIRED
	BucketName string `json:"BucketName"`

	// REQUIRED
	PageNum int32 `json:"PageNum"`

	// REQUIRED
	PageSize int32   `json:"PageSize"`
	Prefix   *string `json:"Prefix,omitempty"`
}

type ListObjectRes struct {
	ResponseMetadata *ListObjectResResponseMetadata `json:"ResponseMetadata,omitempty"`
	Result           *ListObjectResResult           `json:"Result,omitempty"`
}

type ListObjectResResponseMetadata struct {
	Action *string `json:"Action,omitempty"`

	// Anything
	Error     interface{} `json:"Error,omitempty"`
	Region    *string     `json:"Region,omitempty"`
	RequestID *string     `json:"RequestID,omitempty"`
	Service   *string     `json:"Service,omitempty"`
	Version   *string     `json:"Version,omitempty"`
}

type ListObjectResResult struct {
	CurPage *int32    `json:"CurPage,omitempty"`
	Prefix  []*string `json:"Prefix,omitempty"`
	Total   *int32    `json:"Total,omitempty"`
}

type ListProjectsBody struct {

	// 每页数量限制，不填默认10
	Limit *int32 `json:"Limit,omitempty"`

	// 页码，不填默认0
	Offset *int32 `json:"Offset,omitempty"`
}

type ListProjectsRes struct {

	// REQUIRED
	ResponseMetadata ListProjectsResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *ListProjectsResResult `json:"Result,omitempty"`
}

type ListProjectsResResponseMetadata struct {

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
	Error   *ListProjectsResResponseMetadataError `json:"Error,omitempty"`
}

type ListProjectsResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

// ListProjectsResResult - 视请求的接口而定
type ListProjectsResResult struct {

	// REQUIRED; 每页数量
	Limit int32 `json:"Limit"`

	// REQUIRED; 页码
	Offset int32 `json:"Offset"`

	// REQUIRED; 项目列表
	Projects []ListProjectsResResultProjectsItem `json:"Projects"`

	// REQUIRED; 总数量
	Total int32 `json:"Total"`
}

type ListProjectsResResultProjectsItem struct {

	// REQUIRED; 账号
	AccountID string `json:"AccountID"`

	// REQUIRED; 创建时间
	CreateDate string `json:"CreateDate"`

	// REQUIRED; 项目展示名
	DisplayName string `json:"DisplayName"`

	// REQUIRED; 父项目名称，没有为空
	ParentProjectName string `json:"ParentProjectName"`

	// REQUIRED; 从根节点到当前节点的路径
	Path string `json:"Path"`

	// REQUIRED; 项目名称
	ProjectName string `json:"ProjectName"`

	// REQUIRED; 状态
	Status string `json:"Status"`

	// REQUIRED; 更新时间
	UpdateDate string `json:"UpdateDate"`
}

type ListProxyConfigBody struct {

	// REQUIRED; 页码
	PageNum int32 `json:"PageNum"`

	// REQUIRED; 页大小
	PageSize int32 `json:"PageSize"`

	// ID
	ID *int32 `json:"ID,omitempty"`

	// 代理模式，0：固定模式，1：解析模式
	Mode *int32 `json:"Mode,omitempty"`

	// 名称
	Name *string `json:"Name,omitempty"`
}

type ListProxyConfigRes struct {

	// REQUIRED
	ResponseMetadata ListProxyConfigResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *ListProxyConfigResResult `json:"Result,omitempty"`
}

type ListProxyConfigResResponseMetadata struct {

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

// ListProxyConfigResResult - 视请求的接口而定
type ListProxyConfigResResult struct {

	// REQUIRED
	List []ListProxyConfigResResultListItem `json:"List"`

	// REQUIRED; 总共数目
	Total int32 `json:"Total"`
}

type ListProxyConfigResResultListItem struct {

	// REQUIRED; 配置级别，overall:全局，single:单客户
	ConfigLevel string `json:"ConfigLevel"`

	// REQUIRED; 描述
	Description string `json:"Description"`

	// REQUIRED; 生效模式，1：默认生效
	EffectType int32 `json:"EffectType"`

	// REQUIRED; 记录ID
	ID string `json:"ID"`

	// REQUIRED; 代理模式，0：固定模式，1：解析模式
	Mode int32 `json:"Mode"`

	// REQUIRED; 名称
	Name string `json:"Name"`

	// REQUIRED
	ProxyConfigList []ListProxyConfigResResultListPropertiesItemsItem `json:"ProxyConfigList"`
}

type ListProxyConfigResResultListPropertiesItemsItem struct {

	// REQUIRED; 集群
	Cluster string `json:"Cluster"`

	// REQUIRED; 机房
	IDC string `json:"IDC"`

	// REQUIRED; 运营商
	ISP string `json:"ISP"`

	// REQUIRED
	ProxyList []ListProxyConfigResResultListPropertiesItemsProxyListItem `json:"ProxyList"`
}

type ListProxyConfigResResultListPropertiesItemsProxyListItem struct {

	// REQUIRED; 代理地址
	URL string `json:"URL"`

	// REQUIRED; 权重
	Weight int32 `json:"Weight"`
}

type ListPullCDNSnapshotTaskBody struct {

	// REQUIRED; 分页数
	PageNum string `json:"PageNum"`

	// REQUIRED; 分页的大小
	PageSize string `json:"PageSize"`
}

type ListPullCDNSnapshotTaskRes struct {

	// REQUIRED
	ResponseMetadata ListPullCDNSnapshotTaskResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *ListPullCDNSnapshotTaskResResult `json:"Result,omitempty"`
}

type ListPullCDNSnapshotTaskResResponseMetadata struct {

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

// ListPullCDNSnapshotTaskResResult - 视请求的接口而定
type ListPullCDNSnapshotTaskResResult struct {

	// REQUIRED; 直播截图列表记录。
	List []ListPullCDNSnapshotTaskResResultListItem `json:"List"`

	// REQUIRED; 分页信息。
	Pagination ListPullCDNSnapshotTaskResResultPagination `json:"Pagination"`
}

type ListPullCDNSnapshotTaskResResultListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 推流/拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 截图任务的结束时间。
	EndTime string `json:"EndTime"`

	// REQUIRED; 截图任务的开始时间。
	StartTime string `json:"StartTime"`

	// REQUIRED; 4种状态:停用，未开始，生效中，已结束
	Status string `json:"Status"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 任务ID，任务的唯一标识。
	TaskID string `json:"TaskId"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`
}

// ListPullCDNSnapshotTaskResResultPagination - 分页信息。
type ListPullCDNSnapshotTaskResResultPagination struct {

	// REQUIRED; 当前页。
	PageCur int32 `json:"PageCur"`

	// REQUIRED; 当前页的数据量。
	PageSize int32 `json:"PageSize"`

	// REQUIRED; 总页数。
	PageTotal int32 `json:"PageTotal"`

	// REQUIRED; 总数据量。
	TotalCount int32 `json:"TotalCount"`
}

type ListPullRecordTaskBody struct {

	// REQUIRED; 分页数
	PageNum int32 `json:"PageNum"`

	// REQUIRED; 分页的大小
	PageSize int32 `json:"PageSize"`
}

type ListPullRecordTaskRes struct {

	// REQUIRED
	ResponseMetadata ListPullRecordTaskResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *ListPullRecordTaskResResult `json:"Result,omitempty"`
}

type ListPullRecordTaskResResponseMetadata struct {

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

// ListPullRecordTaskResResult - 视请求的接口而定
type ListPullRecordTaskResResult struct {

	// REQUIRED; 直播录制列表记录。
	List []ListPullRecordTaskResResultListItem `json:"List"`

	// REQUIRED; 分页信息。
	Pagination ListPullRecordTaskResResultPagination `json:"Pagination"`
}

type ListPullRecordTaskResResultListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 推流域名或拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 录制的结束时间，RFC3339 格式表示的 UTC 时间戳，精度为 s。
	EndTime string `json:"EndTime"`

	// REQUIRED; 录制的开始时间，RFC3339 格式表示的 UTC 时间戳，精度为 s。
	StartTime string `json:"StartTime"`

	// REQUIRED; 4种状态: 停用，未开始，生效中，已结束
	Status string `json:"Status"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 任务 ID，任务的唯一标识。
	TaskID string `json:"TaskId"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`
}

// ListPullRecordTaskResResultPagination - 分页信息。
type ListPullRecordTaskResResultPagination struct {

	// REQUIRED; 目前页数
	PageCur int32 `json:"PageCur"`

	// REQUIRED; 目前分页页大小
	PageSize int32 `json:"PageSize"`

	// REQUIRED; 总共页数
	PageTotal int32 `json:"PageTotal"`

	// REQUIRED; 任务数量
	TotalCount int32 `json:"TotalCount"`
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

	// 续播策略，续播策略指转推点播视频进行直播时出现断流并恢复后，如何继续播放的策略，拉流来源类型为点播视频（Type 为 1）时参数生效，支持的取值及含义如下。
	// 0：从断流处续播（默认值）； 1：从断流处+自然流逝时长处续播。
	ContinueStrategy *int32 `json:"ContinueStrategy,omitempty"`

	// 点播视频文件循环播放模式，当拉流来源类型为点播视频（Type 为 1）时配置生效，参数取值及含义如下所示。
	// * -1：无限循环，至任务结束；
	// * 0：有限次循环，循环次数为 PlayTimes 取值为准。
	CycleMode *int32 `json:"CycleMode,omitempty"`

	// 推流地址，即直播源或点播视频转推的目标地址。
	DstAddr *string `json:"DstAddr,omitempty"`

	// 推流地址类型。
	// * 1：非第三方，即推流地址域名已添加到视频直播。
	// * 2：第三方，即推流地址域名未添加到视频直播。
	DstAddrType *int32 `json:"DstAddrType,omitempty"`

	// 任务的结束时间，RFC3339 格式的 UTC 时间，单位为秒。
	EndTime *string `json:"EndTime,omitempty"`

	// 点播文件启播时间偏移值，单位为秒，数量与拉流地址列表中地址数量相等，缺省情况下为空表示不进行偏移。 拉流来源类型为点播视频（Type 为 1）时，参数生效。
	OffsetS []*float32 `json:"OffsetS,omitempty"`

	// 点播视频文件循环播放次数，当循环播放模式为有限次循环（CycleMode为0）时配置生效。
	PlayTimes *int32 `json:"PlayTimes,omitempty"`

	// 是否开启点播预热，开启点播预热后，系统会自动将点播视频文件缓存到 CDN 节点上，当用户请求直播时，可以直播从 CDN 节点获取视频，从而提高直播流畅度。 拉流来源类型为点播视频（Type 为 1）时，参数生效。
	// * 0：不开启；
	// * 1：开启。
	PreDownload *int32 `json:"PreDownload,omitempty"`

	// 直播源的拉流地址，拉流来源类型为直播源（Type 为 0）时返回此值。
	SrcAddr *string `json:"SrcAddr,omitempty"`

	// 点播视频播放地址列表，拉流来源类型为点播视频（type 为 1）时返回此值。
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

type ListReferenceInfoBody struct {

	// 引用名称列表
	NameList []*string `json:"NameList,omitempty"`

	// 引用类型列表
	TypeList []*string `json:"TypeList,omitempty"`
}

type ListReferenceInfoRes struct {

	// REQUIRED
	ResponseMetadata ListReferenceInfoResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *ListReferenceInfoResResult `json:"Result,omitempty"`
}

type ListReferenceInfoResResponseMetadata struct {

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

// ListReferenceInfoResResult - 视请求的接口而定
type ListReferenceInfoResResult struct {

	// REQUIRED; 配置列表
	ReferenceInfos []ListReferenceInfoResResultReferenceInfosItem `json:"ReferenceInfos"`
}

type ListReferenceInfoResResultReferenceInfosItem struct {

	// REQUIRED; 引用名称
	Name string `json:"Name"`

	// REQUIRED; 服务类型
	ServiceType string `json:"ServiceType"`

	// REQUIRED; 配置块名称
	TemplateName string `json:"TemplateName"`

	// REQUIRED; 引用类型
	Type string `json:"Type"`

	// REQUIRED; 具体配置的值
	Value string `json:"Value"`
}

type ListReferenceNamesBody struct {

	// 引用类型列表
	TypeList []*string `json:"TypeList,omitempty"`
}

type ListReferenceNamesRes struct {

	// REQUIRED
	ResponseMetadata ListReferenceNamesResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListReferenceNamesResResult          `json:"Result,omitempty"`
}

type ListReferenceNamesResResponseMetadata struct {

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

type ListReferenceNamesResResult struct {

	// REQUIRED; 引用名称列表
	NameList []string `json:"NameList"`
}

type ListReferenceTypesRes struct {

	// REQUIRED
	ResponseMetadata ListReferenceTypesResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListReferenceTypesResResult          `json:"Result,omitempty"`
}

type ListReferenceTypesResResponseMetadata struct {

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

type ListReferenceTypesResResult struct {

	// REQUIRED; 引用类型列表
	TypeList []string `json:"TypeList"`
}

type ListRelaySourceV4Body struct {

	// REQUIRED; 直播流使用的域名。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console-stable.volcanicengine.com/live/main/domain/list]页面，查看直播流使用的域名。
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

type ListResourcePackageBody struct {

	// REQUIRED
	PageNum int32 `json:"PageNum"`

	// REQUIRED
	PageSize   int32  `json:"PageSize"`
	NeedFilter *int32 `json:"NeedFilter,omitempty"`

	// 1:运行中，vadmin查询可以填1
	Status *int32 `json:"Status,omitempty"`
}

type ListResourcePackageRes struct {
	ResponseMetadata *ListResourcePackageResResponseMetadata `json:"ResponseMetadata,omitempty"`
	Result           *ListResourcePackageResResult           `json:"Result,omitempty"`
}

type ListResourcePackageResResponseMetadata struct {
	Action *string `json:"Action,omitempty"`

	// Anything
	Error     interface{} `json:"Error,omitempty"`
	Region    *string     `json:"Region,omitempty"`
	RequestID *string     `json:"RequestID,omitempty"`
	Service   *string     `json:"Service,omitempty"`
	Version   *string     `json:"Version,omitempty"`
}

type ListResourcePackageResResult struct {

	// REQUIRED; 列表
	List []ListResourcePackageResResultListItem `json:"List"`

	// REQUIRED; 总数
	Total int32 `json:"Total"`
}

type ListResourcePackageResResultListItem struct {

	// REQUIRED; 账号
	AccountID string `json:"AccountID"`

	// REQUIRED; 资源包剩余容量
	AvailableAmount float32 `json:"AvailableAmount"`

	// REQUIRED; 是否已经绑定
	BindStatus int32 `json:"BindStatus"`

	// REQUIRED; 配置项
	ConfigurationCode string `json:"ConfigurationCode"`

	// REQUIRED; 过期时间
	ExpireTime string `json:"ExpireTime"`

	// REQUIRED; 1.基础版license、2.高级版license、3.试用版license、4.流量包
	LicenseSourceType string `json:"LicenseSourceType"`

	// REQUIRED; 资源包ID
	PackageID string `json:"PackageID"`

	// REQUIRED; 资源包名
	PackageName string `json:"PackageName"`

	// REQUIRED; 状态
	Status int32 `json:"Status"`

	// REQUIRED; 资源包总容量
	TotalAmount float32 `json:"TotalAmount"`

	// REQUIRED; 资源包单位
	Unit string `json:"Unit"`
}

type ListSDKAdminBody struct {

	// 应用ID
	AppID *string `json:"AppID,omitempty"`

	// 应用名称，支持模糊搜索
	AppName  *string `json:"AppName,omitempty"`
	BundleID *string `json:"BundleID,omitempty"`

	// sdk 记录ID
	ID          *string `json:"ID,omitempty"`
	LicenseID   *string `json:"LicenseID,omitempty"`
	PackageName *string `json:"PackageName,omitempty"`

	// 页码
	PageNum *string `json:"PageNum,omitempty"`

	// 每页大小
	PageSize *string `json:"PageSize,omitempty"`

	// SDK版本，精简版：1、互动版：2
	SDKVersion *string `json:"SDKVersion,omitempty"`

	// 购买方式
	SellType *string `json:"SellType,omitempty"`

	// 状态，0：未激活，1：激活，2：审批，3：过期，4：删除，5：试用过期，6：正式过期，7：试用激活，8：彻底删除
	Status []*int32 `json:"Status,omitempty"`
}

type ListSDKAdminRes struct {

	// REQUIRED
	ResponseMetadata ListSDKAdminResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *ListSDKAdminResResult `json:"Result,omitempty"`
}

type ListSDKAdminResResponseMetadata struct {

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

// ListSDKAdminResResult - 视请求的接口而定
type ListSDKAdminResResult struct {

	// REQUIRED; 列表
	List []ListSDKAdminResResultListItem `json:"List"`

	// REQUIRED; 总数
	Total int32 `json:"Total"`
}

type ListSDKAdminResResultListItem struct {

	// REQUIRED; 账号
	AccountID string `json:"AccountID"`

	// REQUIRED
	ActiveTime string `json:"ActiveTime"`

	// REQUIRED; appid
	AppID int32 `json:"AppID"`

	// REQUIRED; app名称
	AppName string `json:"AppName"`

	// REQUIRED
	ApplyTime string `json:"ApplyTime"`

	// REQUIRED; ios包名
	BundleID string `json:"BundleID"`

	// REQUIRED
	CreateTime string `json:"CreateTime"`

	// REQUIRED
	ExpireTime string `json:"ExpireTime"`

	// REQUIRED
	ID int32 `json:"ID"`

	// REQUIRED
	LicenseID string `json:"LicenseID"`

	// REQUIRED
	LicenseType int32 `json:"LicenseType"`

	// REQUIRED; license链接
	LicenseURL string `json:"LicenseURL"`

	// REQUIRED
	OperateTime string `json:"OperateTime"`

	// REQUIRED
	OperateUser string `json:"OperateUser"`

	// REQUIRED
	PackageID string `json:"PackageID"`

	// REQUIRED; 安卓包名
	PackageName string `json:"PackageName"`

	// REQUIRED; sdk版本
	SDKVersion int32 `json:"SDKVersion"`

	// REQUIRED
	SellType int32 `json:"SellType"`

	// REQUIRED; 状态
	Status int32 `json:"Status"`

	// REQUIRED; vadmin使用的显示
	StatusToShow int32 `json:"StatusToShow"`
}

type ListSDKBody struct {

	// REQUIRED; 页码
	PageNum int32 `json:"PageNum"`

	// REQUIRED; 每页大小
	PageSize int32 `json:"PageSize"`

	// 应用ID
	AppID    *int32  `json:"AppID,omitempty"`
	BundleID *string `json:"BundleID,omitempty"`

	// sdk 记录ID
	ID        *int32  `json:"ID,omitempty"`
	LicenseID *string `json:"LicenseID,omitempty"`

	// license类型：1：基础版，2：高级版，3：试用版，0：无版本
	LicenseType *int32  `json:"LicenseType,omitempty"`
	PackageName *string `json:"PackageName,omitempty"`

	// SDK版本，精简版：1、互动版：2
	SDKVersion *int32 `json:"SDKVersion,omitempty"`

	// 购买方式,1:线上，2：线下
	SellType *int32 `json:"SellType,omitempty"`

	// 状态
	Status []*int32 `json:"Status,omitempty"`
}

type ListSDKRes struct {

	// REQUIRED
	ResponseMetadata ListSDKResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListSDKResResult          `json:"Result,omitempty"`
}

type ListSDKResResponseMetadata struct {
	Action    *string                          `json:"Action,omitempty"`
	Error     *ListSDKResResponseMetadataError `json:"Error,omitempty"`
	Region    *string                          `json:"Region,omitempty"`
	RequestID *string                          `json:"RequestId,omitempty"`
	Service   *string                          `json:"Service,omitempty"`
	Version   *string                          `json:"Version,omitempty"`
}

type ListSDKResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type ListSDKResResult struct {

	// sdk详情列表
	List []*ListSDKResResultListItem `json:"List,omitempty"`

	// sdk总记录数
	Total *int32 `json:"Total,omitempty"`
}

type ListSDKResResultListItem struct {

	// 账号
	AccountID *string `json:"AccountID,omitempty"`

	// 激活时间
	ActiveTime *string `json:"ActiveTime,omitempty"`

	// 应用ID
	AppID *int32 `json:"AppID,omitempty"`

	// 应用名称
	AppName *string `json:"AppName,omitempty"`

	// 应用英文名称
	AppNameEn *string `json:"AppNameEn,omitempty"`

	// 申请时间
	ApplyTime *string `json:"ApplyTime,omitempty"`

	// BundleID
	BundleID *string `json:"BundleID,omitempty"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitempty"`

	// sdk记录ID
	ID *int32 `json:"ID,omitempty"`

	// 证书ID
	LicenseID *string `json:"LicenseID,omitempty"`

	// License类型，0：无版本，1：基础版本，2：高级版本，3：试用版
	LicenseType *int32 `json:"LicenseType,omitempty"`

	// License下载地址
	LicenseURL *string `json:"LicenseURL,omitempty"`

	// 操作时间
	OperateTime *string `json:"OperateTime,omitempty"`

	// 操作者
	OperateUser *string `json:"OperateUser,omitempty"`

	// 流量包ID
	PackageID *string `json:"PackageID,omitempty"`

	// 包名
	PackageName *string `json:"PackageName,omitempty"`

	// 应用类型，WEB, APP
	SDKType *string `json:"SDKType,omitempty"`

	// //SDK版本，精简版：1、互动版：2
	SDKVersion *int32 `json:"SDKVersion,omitempty"`

	// 购买方式，1：人工开通，2：线上购买
	SellType *int32 `json:"SellType,omitempty"`

	// 状态，0：未激活，1：已激活，2：审核中，3：已过期
	Status *int32 `json:"Status,omitempty"`
}

type ListServicesBody struct {

	// 用户账号
	AccountID *string `json:"AccountID,omitempty"`

	// 公司名称
	CompanyName *string `json:"CompanyName,omitempty"`

	// 联系人号码
	ContactNumber *string `json:"ContactNumber,omitempty"`

	// 联系人名称
	ContactPerson *string `json:"ContactPerson,omitempty"`

	// 用户创建的域名
	Domain *string `json:"Domain,omitempty"`

	// 页码大小 [1~1000]
	PageNum *string `json:"PageNum,omitempty"`

	// 分页大小 [1~1000]
	PageSize *string `json:"PageSize,omitempty"`
}

type ListServicesRes struct {

	// REQUIRED
	ResponseMetadata ListServicesResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result ListServicesResResult `json:"Result"`
}

type ListServicesResResponseMetadata struct {
	Action    *string                               `json:"Action,omitempty"`
	Error     *ListServicesResResponseMetadataError `json:"Error,omitempty"`
	Region    *string                               `json:"Region,omitempty"`
	RequestID *string                               `json:"RequestId,omitempty"`
	Service   *string                               `json:"Service,omitempty"`
	Version   *string                               `json:"Version,omitempty"`
}

type ListServicesResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type ListServicesResResult struct {

	// services列表
	List []*ListServicesResResultListItem `json:"List,omitempty"`

	// 总数
	Total *string `json:"Total,omitempty"`
}

type ListServicesResResultListItem struct {

	// 用户账号
	AccountID *string `json:"AccountID,omitempty"`

	// 申请功能
	ApplyService *string `json:"ApplyService,omitempty"`

	// 审批人ID
	ApproverID *string `json:"ApproverID,omitempty"`

	// 审批人姓名
	ApproverName *string `json:"ApproverName,omitempty"`

	// 订单审批状态
	BillingBillingStatus *string `json:"BillingBillingStatus,omitempty"`

	// 订单状态
	BillingStatus *string `json:"BillingStatus,omitempty"`

	// 计费方式
	BillingType *string `json:"BillingType,omitempty"`

	// 低延时直播计费方式
	BillingTypeRTM *string `json:"BillingTypeRTM,omitempty"`

	// 业务领域
	BusinessArea *string `json:"BusinessArea,omitempty"`

	// 业务场景
	BusinessScene *string `json:"BusinessScene,omitempty"`

	// 公司名称
	CompanyName *string `json:"CompanyName,omitempty"`

	// 联系人号码
	ContactNumber *string `json:"ContactNumber,omitempty"`

	// 联系人姓名
	ContactPerson *string `json:"ContactPerson,omitempty"`

	// 最新提交时间
	CreateTime  *string                                   `json:"CreateTime,omitempty"`
	LimitConfig *ListServicesResResultListItemLimitConfig `json:"LimitConfig,omitempty"`

	// 1: 录制是否隐藏TOS 2: 截图是否隐藏TOS 3: 时移是否隐藏VOD 4: 云导播是否隐藏 5：海外加速计费是否隐藏 6：RTM单独加速计费是否隐藏 7：基础版License申请是否隐藏 8：高级版License申请是否隐藏 9：固定回源是否隐藏
	// 10: 月结欠费关停是否处理，1表示处理 11: IP限频是否隐藏 12：URL限频是否隐藏 13：URL参数限频是否隐藏
	// 14：IP访问相同URL限频是否隐藏 15: 活动带宽计费是否隐藏 16: 画质增强是否隐藏 17: Quic加速计费是否隐藏
	PresetConfigHide []*int32 `json:"PresetConfigHide,omitempty"`

	// 服务状态 服务状态（仅在状态为审批通过后生效）
	// * 0: 正式
	// * 1: 试用
	// * 2: 过期状态
	// * -1 表示空
	ServiceStatus *string `json:"ServiceStatus,omitempty"`

	// 审批状态
	// * 0: 正常
	// * 1: 未发起
	// * 2: 未审批
	// * 3: 审批未通过
	// * 4：试用
	// * 5：过期
	Status *string `json:"Status,omitempty"`

	// 过期时间
	TrailTime *string `json:"TrailTime,omitempty"`

	// 处理时间
	UpdateTime *string `json:"UpdateTime,omitempty"`
}

type ListServicesResResultListItemLimitConfig struct {

	// 账号下app限制数目
	AppLimit *int32 `json:"AppLimit,omitempty"`

	// 账号vhost下的domain限制数目
	DomainLimit *int32 `json:"DomainLimit,omitempty"`

	// 账号下的vhost限制数目，-1表示无限制，0表示使用默认配置数目
	VhostLimit *int32 `json:"VhostLimit,omitempty"`
}

type ListTimeShiftPresetV2Body struct {

	// 时移类型。默认类型为 VoD。
	// * vod：点播时移，获取 vod 类型的时移模板配置信息；
	// * tos：直播时移，获取 tos 以及 fcdn-tos 类型的时移模板配置信息。
	Type *string `json:"Type,omitempty"`

	// 域名空间名称。
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

	// 模板列表。
	List []*ListTimeShiftPresetV2ResResultListItem `json:"List,omitempty"`
}

type ListTimeShiftPresetV2ResResultListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; ToS 存储目录。
	Bucket string `json:"Bucket"`

	// REQUIRED
	CreateTime string `json:"CreateTime"`

	// REQUIRED; 拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED
	MasterFormat string `json:"MasterFormat"`

	// REQUIRED; 最大时移时长，即观看时移的最长时间，单位为 s。
	MaxShiftTime int32 `json:"MaxShiftTime"`

	// REQUIRED; 模板名称。
	Name string `json:"Name"`

	// REQUIRED; 直播时移配置模版状态。
	// * 0：配置中；
	// * 1：已启用。
	Status int32 `json:"Status"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 类型。默认类型为 VoD。
	// * VoD
	// * ToS
	// * fcdn-ToS
	Type string `json:"Type"`

	// REQUIRED
	UpdateTime string `json:"UpdateTime"`

	// REQUIRED; 点播空间。
	VODNamespace string `json:"VODNamespace"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`
}

type ListVQScoreTaskBody struct {

	// 查询结束时间，UTC时间格式，支持查询最近30天内的数据
	EndTime *string `json:"EndTime,omitempty"`

	// 当前页码，取值范围为 [1,1000]。
	PageNum *string `json:"PageNum,omitempty"`

	// 分页大小，取值范围为 [1,1000]。
	PageSize *string `json:"PageSize,omitempty"`

	// 查询开始时间，UTC时间格式，支持查询最近30天内的数据
	StartTime *string `json:"StartTime,omitempty"`

	// 测评状态，0：全部，1：测试中，2：成功，3：失败
	Status *string `json:"Status,omitempty"`
}

type ListVQScoreTaskRes struct {

	// REQUIRED
	ResponseMetadata ListVQScoreTaskResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListVQScoreTaskResResult          `json:"Result,omitempty"`
}

type ListVQScoreTaskResResponseMetadata struct {

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
	Error   *ListVQScoreTaskResResponseMetadataError `json:"Error,omitempty"`
}

type ListVQScoreTaskResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type ListVQScoreTaskResResult struct {

	// 查询结束时间，UTC时间格式
	EndTime *string `json:"EndTime,omitempty"`

	// 查询开始时间，UTC时间格式
	StartTime *string `json:"StartTime,omitempty"`

	// 测评任务列表明细
	TaskList []*ListVQScoreTaskResResultTaskListItem `json:"TaskList,omitempty"`

	// 条目总数
	Total *int32 `json:"Total,omitempty"`
}

type ListVQScoreTaskResResultTaskListItem struct {

	// 账号
	AccountID *string `json:"AccountID,omitempty"`

	// 测评运行时间
	Duration *int32 `json:"Duration,omitempty"`

	// 测评任务ID
	ID *string `json:"ID,omitempty"`

	// 测评状态，1：测试中，2：成功，3：失败
	Status *int32 `json:"Status,omitempty"`
}

type ListVhostDenseSnapshotPresetBody struct {

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`
}

type ListVhostDenseSnapshotPresetRes struct {

	// REQUIRED
	ResponseMetadata ListVhostDenseSnapshotPresetResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 视请求的接口而定
	Result ListVhostDenseSnapshotPresetResResult `json:"Result"`
}

type ListVhostDenseSnapshotPresetResResponseMetadata struct {

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

// ListVhostDenseSnapshotPresetResResult - 视请求的接口而定
type ListVhostDenseSnapshotPresetResResult struct {

	// REQUIRED; 模板列表。
	PresetList []ListVhostDenseSnapshotPresetResResultPresetListItem `json:"PresetList"`
}

type ListVhostDenseSnapshotPresetResResultPresetListItem struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 密集抽帧截图模板配置信息。
	DenseSnapshotPreset ListVhostDenseSnapshotPresetResResultPresetListItemDenseSnapshotPreset `json:"DenseSnapshotPreset"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`
}

// ListVhostDenseSnapshotPresetResResultPresetListItemDenseSnapshotPreset - 密集抽帧截图模板配置信息。
type ListVhostDenseSnapshotPresetResResultPresetListItemDenseSnapshotPreset struct {

	// REQUIRED; 截图在 ToS 中的存储位置。
	Bucket string `json:"Bucket"`

	// REQUIRED; 回调地址。
	CallbackURL string `json:"CallbackUrl"`

	// REQUIRED; 截图间隔时间。
	Interval int32 `json:"Interval"`

	// REQUIRED; 密集抽帧截图配置模板名称。
	Preset string `json:"Preset"`

	// REQUIRED; veImageX 的服务 ID。
	ServiceID string `json:"ServiceID"`

	// REQUIRED; 密集抽帧截图配置模版的开启状态。
	// * 1：开启
	// * 0：关闭
	Status int32 `json:"Status"`
}

type ListVhostDetailBody struct {

	// REQUIRED
	PageNum int32 `json:"PageNum"`

	// REQUIRED
	PageSize         int32     `json:"PageSize"`
	AccountIDList    []*string `json:"AccountIDList,omitempty"`
	DomainRegionList []*string `json:"DomainRegionList,omitempty"`
	DomainStatusList []*int32  `json:"DomainStatusList,omitempty"`
	DomainTypeList   []*string `json:"DomainTypeList,omitempty"`
	VhostNameList    []*string `json:"VhostNameList,omitempty"`
	VhostStatusList  []*int32  `json:"VhostStatusList,omitempty"`
	VhostTypeList    []*string `json:"VhostTypeList,omitempty"`
}

type ListVhostDetailByAdminBody struct {

	// REQUIRED
	PageNum int32 `json:"PageNum"`

	// REQUIRED
	PageSize         int32     `json:"PageSize"`
	AccountIDList    []*string `json:"AccountIDList,omitempty"`
	DomainRegionList []*string `json:"DomainRegionList,omitempty"`
	DomainStatusList []*int32  `json:"DomainStatusList,omitempty"`
	DomainTypeList   []*string `json:"DomainTypeList,omitempty"`
	VhostNameList    []*string `json:"VhostNameList,omitempty"`
	VhostStatusList  []*int32  `json:"VhostStatusList,omitempty"`
	VhostTypeList    []*string `json:"VhostTypeList,omitempty"`
}

type ListVhostDetailByAdminRes struct {

	// REQUIRED
	ResponseMetadata ListVhostDetailByAdminResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListVhostDetailByAdminResResult          `json:"Result,omitempty"`
}

type ListVhostDetailByAdminResResponseMetadata struct {

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

type ListVhostDetailByAdminResResult struct {
	Total     *int32                                          `json:"Total,omitempty"`
	VhostList []*ListVhostDetailByAdminResResultVhostListItem `json:"VhostList,omitempty"`
}

type ListVhostDetailByAdminResResultVhostListItem struct {
	AccountID  *string                                                                                                               `json:"AccountID,omitempty"`
	AppList    []*string                                                                                                             `json:"AppList,omitempty"`
	DomainList []*Components1Bmm523SchemasListvhostdetailbyadminresPropertiesResultPropertiesVhostlistItemsPropertiesDomainlistItems `json:"DomainList,omitempty"`
	ID         *int32                                                                                                                `json:"ID,omitempty"`
	Priority   *int32                                                                                                                `json:"Priority,omitempty"`
	Status     *int32                                                                                                                `json:"Status,omitempty"`
	Type       *string                                                                                                               `json:"Type,omitempty"`
	Vhost      *string                                                                                                               `json:"Vhost,omitempty"`
}

type ListVhostDetailRes struct {

	// REQUIRED
	ResponseMetadata ListVhostDetailResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListVhostDetailResResult          `json:"Result,omitempty"`
}

type ListVhostDetailResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version   string                                   `json:"Version"`
	Error     *ListVhostDetailResResponseMetadataError `json:"Error,omitempty"`
	RequestID *string                                  `json:"RequestID,omitempty"`
}

type ListVhostDetailResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type ListVhostDetailResResult struct {

	// REQUIRED
	Total int32 `json:"Total"`

	// REQUIRED
	VhostList []ListVhostDetailResResultVhostListItem `json:"VhostList"`
}

type ListVhostDetailResResultVhostListItem struct {

	// REQUIRED
	AccountID string `json:"AccountID"`

	// REQUIRED
	AppList []string `json:"AppList"`

	// REQUIRED; 创建时间
	CreateTime string `json:"CreateTime"`

	// REQUIRED
	DomainList []Components4Y1LroSchemasListvhostdetailresPropertiesResultPropertiesVhostlistItemsPropertiesDomainlistItems `json:"DomainList"`

	// REQUIRED
	ID int32 `json:"ID"`

	// REQUIRED
	Priority int32 `json:"Priority"`

	// REQUIRED; 项目名称
	ProjectName string `json:"ProjectName"`

	// REQUIRED
	Status int32 `json:"Status"`

	// REQUIRED; 标签
	Tags []Components1M64L84SchemasListvhostdetailresPropertiesResultPropertiesVhostlistItemsPropertiesTagsItems `json:"Tags"`

	// REQUIRED; 更新时间
	UpdateTime string `json:"UpdateTime"`

	// REQUIRED
	Vhost string  `json:"Vhost"`
	Type  *string `json:"Type,omitempty"`
}

// ListVhostDetailResResultVhostListItemDomainListItemTags - 标签
type ListVhostDetailResResultVhostListItemDomainListItemTags struct {

	// REQUIRED
	Category string `json:"Category"`

	// REQUIRED
	Key string `json:"Key"`

	// REQUIRED
	Value string `json:"Value"`
}

type ListVhostDomainDetailByUserIDBody struct {

	// REQUIRED; 账号
	AccountID string `json:"AccountID"`

	// REQUIRED; 子账号userID
	UserID string `json:"UserID"`
}

type ListVhostDomainDetailByUserIDRes struct {

	// REQUIRED
	ResponseMetadata ListVhostDomainDetailByUserIDResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListVhostDomainDetailByUserIDResResult          `json:"Result,omitempty"`
}

type ListVhostDomainDetailByUserIDResResponseMetadata struct {

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

type ListVhostDomainDetailByUserIDResResult struct {
	DomainList []*ListVhostDomainDetailByUserIDResResultDomainListItem `json:"DomainList,omitempty"`
	VhostList  []*ListVhostDomainDetailByUserIDResResultVhostListItem  `json:"VhostList,omitempty"`
}

type ListVhostDomainDetailByUserIDResResultDomainListItem struct {
	Domain *string `json:"Domain,omitempty"`
	Status *int32  `json:"Status,omitempty"`
	Type   *string `json:"Type,omitempty"`
	Vhost  *string `json:"Vhost,omitempty"`
}

type ListVhostDomainDetailByUserIDResResultVhostListItem struct {
	AccountID *string `json:"AccountID,omitempty"`
	Status    *int32  `json:"Status,omitempty"`
	Type      *string `json:"Type,omitempty"`
	Vhost     *string `json:"Vhost,omitempty"`
}

type ListVhostRecordPresetV2Body struct {

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 直播录制的存储类型。默认值为 tos，支持以下取值。
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

	// REQUIRED; 录制模板列表。
	PresetList []ListVhostRecordPresetV2ResResultPresetListItem `json:"PresetList"`
}

type ListVhostRecordPresetV2ResResultPresetListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 模板详细信息。
	SlicePresetV2 *ListVhostRecordPresetV2ResResultPresetListItemSlicePresetV2 `json:"SlicePresetV2,omitempty"`
}

// ListVhostRecordPresetV2ResResultPresetListItemSlicePresetV2 - 模板详细信息。
type ListVhostRecordPresetV2ResResultPresetListItemSlicePresetV2 struct {

	// 模板 ID。
	ID *int32 `json:"ID,omitempty"`

	// 模板名称。
	Name *string `json:"Name,omitempty"`

	// 录制模板详细配置。
	RecordPresetConfig *ComponentsFuamuzSchemasListvhostrecordpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesRecordpresetconfig `json:"RecordPresetConfig,omitempty"`
}

type ListVhostSnapshotAuditPresetBody struct {

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`
}

type ListVhostSnapshotAuditPresetRes struct {

	// REQUIRED
	ResponseMetadata ListVhostSnapshotAuditPresetResResponseMetadata `json:"ResponseMetadata"`
	Result           *ListVhostSnapshotAuditPresetResResult          `json:"Result,omitempty"`
}

type ListVhostSnapshotAuditPresetResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                                `json:"Version"`
	Error   *ListVhostSnapshotAuditPresetResResponseMetadataError `json:"Error,omitempty"`
}

type ListVhostSnapshotAuditPresetResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type ListVhostSnapshotAuditPresetResResult struct {

	// REQUIRED; 截图审核配置列表。
	PresetList []ListVhostSnapshotAuditPresetResResultPresetListItem `json:"PresetList"`
}

type ListVhostSnapshotAuditPresetResResultPresetListItem struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 截图审核配置详细信息。
	AuditPreset ListVhostSnapshotAuditPresetResResultPresetListItemAuditPreset `json:"AuditPreset"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`
}

// ListVhostSnapshotAuditPresetResResultPresetListItemAuditPreset - 截图审核配置详细信息。
type ListVhostSnapshotAuditPresetResResultPresetListItemAuditPreset struct {

	// REQUIRED; ToS 存储对应的 Bucket。
	Bucket string `json:"Bucket"`

	// REQUIRED; 截图审核结果回调地址配置。
	CallbackDetailList []ListVhostSnapshotAuditPresetResResultPresetListPropertiesItemsItem `json:"CallbackDetailList"`

	// REQUIRED; 截图审核配置的描述。
	Description string `json:"Description"`

	// REQUIRED; 截图间隔时间，单位秒，取值范围为[0.1,10]，支持保留两位小数。
	Interval float32 `json:"Interval"`

	// REQUIRED; 审核标签名称，取值及含义如下。
	// * 301：涉黄；
	// * 302：涉敏1；
	// * 303：涉敏2；
	// * 304：广告；
	// * 305：引人不适；
	// * 306：违禁；
	// * 307：二维码；
	// * 308：诈骗；
	// * 309：不良画面；
	// * 310：未成年相关；
	// * 320：文字违规。
	Label []string `json:"Label"`

	// REQUIRED; 截图审核配置的名称。
	PresetName string `json:"PresetName"`

	// REQUIRED; veimageX 的服务 ID。 :::tip 参数 Bucket 和 ServiceID 传且仅传一个。 :::
	ServiceID string `json:"ServiceID"`

	// REQUIRED; 存储方式为实时存储时的存储规则，支持以 {Domain}/{App}/{Stream}/{UnixTimestamp} 样式设置存储规则，支持输入字母、数字、"-"、"!"、"_"、"."、"*"及占位符。
	SnapshotObject string `json:"SnapshotObject"`

	// REQUIRED; ToS 存储对应的 Bucket 下的存储目录。
	StorageDir string `json:"StorageDir"`

	// REQUIRED; 存储策略，取值及含义如下。
	// * 0：触发存储，只存储有风险图片；
	// * 1：全部存储，存储全部图片。
	StorageStrategy int32 `json:"StorageStrategy"`

	// REQUIRED; 配置信息的更新时间，RFC3339 格式的 UTC 时间，精度为秒。
	UpdateTime string `json:"UpdateTime"`
}

type ListVhostSnapshotAuditPresetResResultPresetListPropertiesItemsItem struct {

	// REQUIRED; 回调的类型 HTTP。
	CallbackType string `json:"CallbackType"`

	// REQUIRED; 回调地址。
	URL string `json:"URL"`
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
	AccessKey *string `json:"AccessKey,omitempty"`
	AccountID *string `json:"AccountID,omitempty"`
	AsShort   *int32  `json:"AsShort,omitempty"`

	// 截图在 ToS 中的存储位置。
	Bucket   *string `json:"Bucket,omitempty"`
	Callback *string `json:"Callback,omitempty"`

	// 回调信息。
	CallbackDetail []*ListVhostSnapshotPresetResResultPresetListPropertiesItemsItem `json:"CallbackDetail,omitempty"`
	CreatedAt      *string                                                          `json:"CreatedAt,omitempty"`
	Description    *string                                                          `json:"Description,omitempty"`
	Duration       *int32                                                           `json:"Duration,omitempty"`
	Format         []*string                                                        `json:"Format,omitempty"`
	Height         *int32                                                           `json:"Height,omitempty"`
	ID             *int32                                                           `json:"ID,omitempty"`

	// 截图间隔时间。
	Interval        *int32  `json:"Interval,omitempty"`
	NssConfig       *string `json:"NssConfig,omitempty"`
	OverwriteObject *string `json:"OverwriteObject,omitempty"`

	// 截图模版名称。
	Preset       *string                                                                                                                                     `json:"Preset,omitempty"`
	PullDomain   *string                                                                                                                                     `json:"PullDomain,omitempty"`
	Quality      *int32                                                                                                                                      `json:"Quality,omitempty"`
	RecordConfig *string                                                                                                                                     `json:"RecordConfig,omitempty"`
	RecordObject *string                                                                                                                                     `json:"RecordObject,omitempty"`
	RecordTob    []*Components1GzojhcSchemasListvhostsnapshotpresetresPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetPropertiesRecordtobItems `json:"RecordTob,omitempty"`
	RegionConfig *string                                                                                                                                     `json:"RegionConfig,omitempty"`
	ReserveDays  *int32                                                                                                                                      `json:"ReserveDays,omitempty"`

	// veImageX 的服务 ID。
	ServiceID      *string `json:"ServiceID,omitempty"`
	SliceDuration  *int32  `json:"SliceDuration,omitempty"`
	SnapshotConfig *string `json:"SnapshotConfig,omitempty"`
	SnapshotFormat *string `json:"SnapshotFormat,omitempty"`
	SnapshotObject *string `json:"SnapshotObject,omitempty"`

	// 截图模版状态。
	// * 1：开启
	// * 0：关闭
	Status       *int32  `json:"Status,omitempty"`
	StorageDir   *string `json:"StorageDir,omitempty"`
	TosCluster   *string `json:"TosCluster,omitempty"`
	TosDC        *string `json:"TosDC,omitempty"`
	TosPSM       *string `json:"TosPSM,omitempty"`
	VodNamespace *string `json:"VodNamespace,omitempty"`
	Width        *int32  `json:"Width,omitempty"`
	WorkflowID   *string `json:"WorkflowID,omitempty"`
}

// ListVhostSnapshotPresetResResultPresetListPropertiesItemsItem - 回调信息
type ListVhostSnapshotPresetResResultPresetListPropertiesItemsItem struct {

	// REQUIRED; 回调地址。
	URL string `json:"URL"`

	// 回调类型。
	// * http
	// * nsq
	// * kafka
	// * rpc
	CallbackType *string `json:"CallbackType,omitempty"`
}

type ListVhostSnapshotPresetV2Body struct {

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 截图配置中截图文件的存储位置，缺省情况下表示不对存储位置进行过滤，取值及含义如下所示。
	// * tos：TOS 对象存储服务；
	// * imageX：veImageX 图片服务。
	Type *string `json:"Type,omitempty"`
}

type ListVhostSnapshotPresetV2Res struct {

	// REQUIRED
	ResponseMetadata ListVhostSnapshotPresetV2ResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *ListVhostSnapshotPresetV2ResResult `json:"Result,omitempty"`
}

type ListVhostSnapshotPresetV2ResResponseMetadata struct {

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

// ListVhostSnapshotPresetV2ResResult - 视请求的接口而定
type ListVhostSnapshotPresetV2ResResult struct {

	// REQUIRED; 截图配置列表。
	PresetList []ListVhostSnapshotPresetV2ResResultPresetListItem `json:"PresetList"`
}

type ListVhostSnapshotPresetV2ResResultPresetListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 截图配置基础信息。
	SlicePresetV2 ListVhostSnapshotPresetV2ResResultPresetListItemSlicePresetV2 `json:"SlicePresetV2"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`
}

// ListVhostSnapshotPresetV2ResResultPresetListItemSlicePresetV2 - 截图配置基础信息。
type ListVhostSnapshotPresetV2ResResultPresetListItemSlicePresetV2 struct {

	// REQUIRED; 截图配置名称。
	Name string `json:"Name"`

	// REQUIRED; 截图配置详细信息。
	SnapshotPresetConfig ListVhostSnapshotPresetV2ResResultPresetListProperties `json:"SnapshotPresetConfig"`

	// REQUIRED; 截图配置生效状态。
	// * 1：生效；
	// * 0：不生效。
	Status int32 `json:"Status"`
}

// ListVhostSnapshotPresetV2ResResultPresetListProperties - 截图配置详细信息。
type ListVhostSnapshotPresetV2ResResultPresetListProperties struct {

	// REQUIRED; 截图间隔时间，单位为秒。
	Interval int32 `json:"Interval"`

	// 图片格式为 JPEG 时的截图参数。
	JPEGParam *ListVhostSnapshotPresetV2ResResultPresetListPropertiesPropertiesProperties `json:"JpegParam,omitempty"`

	// 截图格式为 JPG 时的截图参数。
	JpgParam *ComponentsSlabtaSchemasListvhostsnapshotpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesSnapshotpresetconfigPropertiesJpgparam `json:"JpgParam,omitempty"`
}

// ListVhostSnapshotPresetV2ResResultPresetListPropertiesPropertiesProperties - 图片格式为 JPEG 时的截图参数。
type ListVhostSnapshotPresetV2ResResultPresetListPropertiesPropertiesProperties struct {

	// REQUIRED; 当前格式的截图是否开启，默认为 false，取值及含义如下所示。
	// * false：不开启；
	// * true：开启。
	Enable bool `json:"Enable"`

	// 截图存储到 veImageX 时的配置。 :::tip TOSParam 和 ImageXParam 配置且配置其中一个。 :::
	ImageXParam *ComponentsK46Cw0SchemasListvhostsnapshotpresetv2ResPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetv2PropertiesSnapshotpresetconfigPropertiesJpegparamPropertiesImagexparam `json:"ImageXParam,omitempty"`

	// 截图存储到 TOS 时的配置。 :::tip TOSParam 和 ImageXParam 配置且配置其中一个。 :::
	TOSParam *ListVhostSnapshotPresetV2ResResultPresetListPropertiesPropertiesPropertiesProperties `json:"TOSParam,omitempty"`
}

// ListVhostSnapshotPresetV2ResResultPresetListPropertiesPropertiesPropertiesProperties - 截图存储到 TOS 时的配置。 :::tip TOSParam
// 和 ImageXParam 配置且配置其中一个。 :::
type ListVhostSnapshotPresetV2ResResultPresetListPropertiesPropertiesPropertiesProperties struct {

	// REQUIRED; TOS 存储对应的 Bucket。 例如，存储路径为 live-test-tos-example/live/liveapp 时，Bucket 取值为 live-test-tos-example。
	Bucket string `json:"Bucket"`

	// REQUIRED; 截图是否使用 TOS 存储，取值及含义如下所示。
	// * false：不使用；
	// * true：使用。
	Enable bool `json:"Enable"`

	// REQUIRED; 存储方式为实时截图时的存储规则，支持以 {Domain}/{App}/{Stream}/{UnixTimestamp} 样式设置存储规则，支持输入字母、数字、-、!、_、.、* 及占位符。
	ExactObject string `json:"ExactObject"`

	// REQUIRED; 存储方式为覆盖截图时的存储规则，支持以 {Domain}/{App}/{Stream} 样式设置存储规则，支持输入字母、数字、-、!、_、.、* 及占位符。
	OverwriteObject string `json:"OverwriteObject"`

	// REQUIRED; Bucket 目录。 例如，存储路径为 live-test-tos-example/live/liveapp 时，StorageDir 取值为 live/liveapp。
	StorageDir string `json:"StorageDir"`
}

type ListVhostTransCodePresetBody struct {

	// REQUIRED; 是否是hls abr 请求
	IsHlsAbr bool `json:"IsHlsAbr"`

	// REQUIRED; 域名空间。
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

	// REQUIRED; 应用名称
	App string `json:"App"`

	// REQUIRED; 域名空间
	Vhost string `json:"Vhost"`

	// 转码配置具体信息
	TranscodePreset *ListVhostTransCodePresetResResultAllPresetListItemTranscodePreset `json:"TranscodePreset,omitempty"`
}

// ListVhostTransCodePresetResResultAllPresetListItemTranscodePreset - 转码配置具体信息
type ListVhostTransCodePresetResResultAllPresetListItemTranscodePreset struct {
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

	// 音频码率，单位为 kbps。
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

	// IDR 帧之间的最大间隔，单位为 s
	GOP    *int32 `json:"GOP,omitempty"`
	GopMin *int32 `json:"GopMin,omitempty"`
	HVSPre *bool  `json:"HVSPre,omitempty"`

	// 视频高度。 :::tip 当As的取值为 0 时，如果Width和Height任意取值为 0，表示保持源流尺寸。 :::
	Height *int32 `json:"Height,omitempty"`

	// 长边长度。 :::tip 当As的取值为 1 时，如果LongSide和ShortSide都取 0，表示保持源流尺寸。 :::
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

	// 模板名称
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

	// 短边长度。 :::tip 当As的取值为 1 时，如果LongSide和ShortSide都取 0，表示保持源流尺寸。 :::
	ShortSide    *int32 `json:"ShortSide,omitempty"`
	Status       *int32 `json:"Status,omitempty"`
	StopInterval *int32 `json:"StopInterval,omitempty"`

	// 转码流后缀名
	SuffixName *string `json:"SuffixName,omitempty"`
	Threads    *int32  `json:"Threads,omitempty"`
	TransType  *string `json:"TransType,omitempty"`
	VBRatio    *int32  `json:"VBRatio,omitempty"`
	VBVBufSize *int32  `json:"VBVBufSize,omitempty"`
	VBVMaxRate *int32  `json:"VBVMaxRate,omitempty"`
	VLevel     *string `json:"VLevel,omitempty"`

	// 转码配置名称
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

	// 视频宽度。 :::tip 当As的取值为 0 时，如果Width和Height任意取值为 0，表示保持源流尺寸。 :::
	Width *int32 `json:"Width,omitempty"`
}

type ListVhostTransCodePresetResResultCommonPresetListItem struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 域名空间。
	Vhost string `json:"Vhost"`

	// 转码配置具体信息。
	TranscodePreset *ListVhostTransCodePresetResResultCommonPresetListItemTranscodePreset `json:"TranscodePreset,omitempty"`
}

// ListVhostTransCodePresetResResultCommonPresetListItemTranscodePreset - 转码配置具体信息。
type ListVhostTransCodePresetResResultCommonPresetListItemTranscodePreset struct {
	ALayout   *string `json:"ALayout,omitempty"`
	AProfile  *string `json:"AProfile,omitempty"`
	AR        *int32  `json:"AR,omitempty"`
	AbrMode   *int32  `json:"AbrMode,omitempty"`
	AccountID *string `json:"AccountID,omitempty"`

	// 音频编码格式，取值含义如下。
	// * aac：使用 AAC 编码格式；
	// * copy：不进行转码，所有音频编码参数不生效；
	// * opus：使用 Opus 编码格式。
	Acodec         *string `json:"Acodec,omitempty"`
	AdvancedParam  *string `json:"AdvancedParam,omitempty"`
	AllowAudioCopy *int32  `json:"AllowAudioCopy,omitempty"`
	AllowVideoCopy *int32  `json:"AllowVideoCopy,omitempty"`
	An             *int32  `json:"An,omitempty"`

	// 宽高自适应模式开关。
	// * 0：关闭宽高自适应；
	// * 1：开启宽高自适应。 :::tip
	// * 关闭宽高自适应时，转码配置分辨率取 Width 和 Height 的值对转码视频进行拉伸；
	// * 开启宽高自适应时，转码配置分辨率按照 ShortSide 、 LongSide 、Width 、Height 的优先级取值，另一边等比缩放。 :::
	As *string `json:"As,omitempty"`

	// 音频码率，单位为 kbps。
	AudioBitrate *int32 `json:"AudioBitrate,omitempty"`
	AutoTransAb  *int32 `json:"AutoTransAb,omitempty"`
	AutoTransAl  *int32 `json:"AutoTransAl,omitempty"`
	AutoTransAr  *int32 `json:"AutoTransAr,omitempty"`

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
	BCM         *int32 `json:"BCM,omitempty"`

	// 2 个参考帧之间的最大 B 帧数。取值为 0 时，表示去除 B 帧。
	BFrames  *int32  `json:"BFrames,omitempty"`
	Describe *string `json:"Describe,omitempty"`

	// 视频帧率，单位为 fps，帧率越大，画面越流畅。
	FPS *int32 `json:"FPS,omitempty"`

	// IDR 帧之间的最大间隔，单位为秒。
	GOP    *int32 `json:"GOP,omitempty"`
	GopMin *int32 `json:"GopMin,omitempty"`
	HVSPre *bool  `json:"HVSPre,omitempty"`

	// 视频高度。
	Height *int32 `json:"Height,omitempty"`

	// 长边长度。 :::tip
	// * 当 As 的取值为 1 即开启宽高自适应时，参数生效，反之则不生效。
	// * 当 As 的取值为 1 时，如果 LongSide 、 ShortSide 、Width 、Height 同时取 0，表示保持源流尺寸。 :::
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

	// 转码配置名称。
	Preset       *string `json:"Preset,omitempty"`
	PresetKind   *int32  `json:"PresetKind,omitempty"`
	PresetType   *int32  `json:"PresetType,omitempty"`
	Qp           *int32  `json:"Qp,omitempty"`
	RegionConfig *string `json:"RegionConfig,omitempty"`
	Revision     *string `json:"Revision,omitempty"`

	// 是否极智超清转码，取值及含义如下。
	// * true：极智超清转码；
	// * false：标准转码。
	Roi  *bool `json:"Roi,omitempty"`
	SITI *bool `json:"SITI,omitempty"`

	// 短边长度。 :::tip
	// * 当 As 的取值为 1 即开启宽高自适应时，参数生效，反之则不生效。
	// * 当 As 的取值为 1 时，如果 LongSide 、 ShortSide 、Width 、Height 同时取 0，表示保持源流尺寸。 :::
	ShortSide *int32 `json:"ShortSide,omitempty"`
	Status    *int32 `json:"Status,omitempty"`

	// 转码停止时长，支持触发方式为拉流转码时设置，表示断开拉流后转码停止的时长，单位为 s，取值范围为 -1 和 [0,300]，-1 表示不停止转码，默认值为 60。
	StopInterval *int32 `json:"StopInterval,omitempty"`

	// 转码流后缀名。
	SuffixName *string `json:"SuffixName,omitempty"`
	Threads    *int32  `json:"Threads,omitempty"`

	// 转码触发方式，取值及含义如下。
	// * Push：推流转码，直播推流后会自动启动转码任务，生成转码流；
	// * Pull：拉流转码，直播推流后，需要主动播放转码流才会启动转码任务，生成转码流。
	TransType    *string `json:"TransType,omitempty"`
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
	// * h266：使用 H.266 编码格式；
	// * copy：不进行转码，所有视频编码参数不生效。
	Vcodec *string `json:"Vcodec,omitempty"`

	// 视频码率，单位为 kbps。
	VideoBitrate *int32  `json:"VideoBitrate,omitempty"`
	Vn           *int32  `json:"Vn,omitempty"`
	Watermark    *string `json:"Watermark,omitempty"`

	// 视频宽度。
	Width *int32 `json:"Width,omitempty"`
}

type ListVhostTransCodePresetResResultCustomizePresetListItem struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成
	App string `json:"App"`

	// REQUIRED; 域名空间
	Vhost           string                                                                   `json:"Vhost"`
	TranscodePreset *ListVhostTransCodePresetResResultCustomizePresetListItemTranscodePreset `json:"TranscodePreset,omitempty"`
}

type ListVhostTransCodePresetResResultCustomizePresetListItemTranscodePreset struct {
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

	// 音频码率，单位为 kbps。
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

	// IDR 帧之间的最大间隔，单位为 s
	GOP    *int32 `json:"GOP,omitempty"`
	GopMin *int32 `json:"GopMin,omitempty"`
	HVSPre *bool  `json:"HVSPre,omitempty"`

	// 视频高度。 :::tip 当As的取值为 0 时，如果Width和Height任意取值为 0，表示保持源流尺寸。 :::
	Height *int32 `json:"Height,omitempty"`

	// 长边长度。 :::tip 当As的取值为 1 时，如果LongSide和ShortSide都取 0，表示保持源流尺寸。 :::
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

	// 模板名称
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

	// 短边长度。 :::tip 当As的取值为 1 时，如果LongSide和ShortSide都取 0，表示保持源流尺寸。 :::
	ShortSide    *int32 `json:"ShortSide,omitempty"`
	Status       *int32 `json:"Status,omitempty"`
	StopInterval *int32 `json:"StopInterval,omitempty"`

	// 转码流后缀名
	SuffixName   *string `json:"SuffixName,omitempty"`
	Threads      *int32  `json:"Threads,omitempty"`
	TransType    *string `json:"TransType,omitempty"`
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

	// 视频宽度。 :::tip 当As的取值为 0 时，如果Width和Height任意取值为 0，表示保持源流尺寸。 :::
	Width *int32 `json:"Width,omitempty"`
}

type ListVhostWatermarkPresetBody struct {

	// REQUIRED; 域名空间名称。由 1 到 60 位数字、字母、下划线及"-"和"."组成。
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

	// 获取模板失败的列表，返回获取失败的模版及获取失败的原因。
	WatermarkErrMsgList []*ListVhostWatermarkPresetResResultWatermarkErrMsgListItem `json:"WatermarkErrMsgList,omitempty"`

	// 水印模版列表。
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

	// 直播画面方向。
	// * vertical：竖屏；
	// * horizontal：横屏。
	Orientation *string `json:"Orientation,omitempty"`

	// 水印图片链接。
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

	// 域名空间名称。
	Vhost *string `json:"Vhost,omitempty"`
}

type ListVideoClassificationsBody struct {

	// REQUIRED; 空间名称
	SpaceName string `json:"SpaceName"`
}

type ListVideoClassificationsRes struct {

	// REQUIRED
	ResponseMetadata ListVideoClassificationsResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *ListVideoClassificationsResResult `json:"Result,omitempty"`
}

type ListVideoClassificationsResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                            `json:"Version"`
	Error   *ListVideoClassificationsResResponseMetadataError `json:"Error,omitempty"`
}

type ListVideoClassificationsResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

// ListVideoClassificationsResResult - 视请求的接口而定
type ListVideoClassificationsResResult struct {

	// REQUIRED; 分类列表
	ClassificationTrees []ListVideoClassificationsResResultClassificationTreesItem `json:"ClassificationTrees"`
}

type ListVideoClassificationsResResultClassificationTreesItem struct {

	// REQUIRED; 分类名称
	Classification string `json:"Classification"`

	// REQUIRED; 分类ID
	ClassificationID string `json:"ClassificationID"`

	// REQUIRED; 创建时间
	CreatedAt string `json:"CreatedAt"`

	// REQUIRED; 分类级别，1：一级，2：二级，3：三级
	Level int32 `json:"Level"`

	// REQUIRED; 父分类ID
	ParentClassificationID string `json:"ParentClassificationID"`

	// REQUIRED; 空间名
	SpaceName string `json:"SpaceName"`

	// REQUIRED; 与ClassificationTrees相同
	SubClassificationTrees []interface{} `json:"SubClassificationTrees"`
}

type ListVqosDimensionValuesBody struct {

	// REQUIRED
	Dimension string `json:"Dimension"`

	// REQUIRED
	End int32 `json:"End"`

	// REQUIRED
	Start int32 `json:"Start"`

	// REQUIRED
	Window string   `json:"Window"`
	Limit  *float32 `json:"Limit,omitempty"`
}

type ListVqosDimensionValuesQuery struct {

	// REQUIRED
	AppQueryType float32 `json:"AppQueryType" query:"AppQueryType"`

	// REQUIRED
	VqosService string `json:"VqosService" query:"VqosService"`
}

type ListVqosDimensionValuesRes struct {

	// REQUIRED
	ResponseMetadata ListVqosDimensionValuesResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result []ListVqosDimensionValuesResResultItem `json:"Result"`
}

type ListVqosDimensionValuesResResponseMetadata struct {

	// REQUIRED
	RequestID string `json:"RequestId"`
}

type ListVqosDimensionValuesResResultItem struct {

	// REQUIRED
	Alias string `json:"Alias"`

	// REQUIRED
	Count int32 `json:"Count"`

	// REQUIRED
	Name string `json:"Name"`
}

type ListVqosMetricsDimensionsQuery struct {

	// REQUIRED
	VqosService string `json:"VqosService" query:"VqosService"`
}

type ListVqosMetricsDimensionsRes struct {

	// REQUIRED
	ResponseMetadata ListVqosMetricsDimensionsResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result []ListVqosMetricsDimensionsResResultItem `json:"Result"`
}

type ListVqosMetricsDimensionsResResponseMetadata struct {

	// REQUIRED
	RequestID string `json:"RequestId"`
}

type ListVqosMetricsDimensionsResResultItem struct {
	Dimensions []*ComponentsFceumsSchemasListvqosmetricsdimensionsresPropertiesResultItemsPropertiesDimensionsItems `json:"Dimensions,omitempty"`
	Metrics    []*ListVqosMetricsDimensionsResResultPropertiesItemsItem                                             `json:"Metrics,omitempty"`
	Service    *string                                                                                              `json:"Service,omitempty"`
}

type ListVqosMetricsDimensionsResResultPropertiesItemsItem struct {

	// REQUIRED
	Alias string `json:"Alias"`

	// REQUIRED
	Attached string `json:"Attached"`

	// REQUIRED
	Attribute string `json:"Attribute"`

	// REQUIRED
	Desc string `json:"Desc"`

	// REQUIRED
	Name string `json:"Name"`

	// REQUIRED
	Type string `json:"Type"`
}

type ListWatermarkPresetBody struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 域名空间名称。由 1 到 60 位数字、字母、下划线及"-"和"."组成。
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

	// 水印模版 ID。
	ID *int32 `json:"ID,omitempty"`

	// 直播画面方向。
	// * vertical：竖屏；
	// * horizontal：横屏。
	Orientation *string `json:"Orientation,omitempty"`

	// 水印图片链接。
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

	// 域名空间名称。
	Vhost *string `json:"Vhost,omitempty"`
}

type ManagerPullPushDomainBindBody struct {

	// REQUIRED; 拉流域名。
	PullDomain string `json:"PullDomain"`

	// 需要绑定的推流域名。缺省情况下，表示解除推拉流域名的绑定关系。
	PushDomain *string `json:"PushDomain,omitempty"`
}

type ManagerPullPushDomainBindRes struct {

	// REQUIRED
	ResponseMetadata ManagerPullPushDomainBindResResponseMetadata `json:"ResponseMetadata"`
}

type ManagerPullPushDomainBindResResponseMetadata struct {

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
	Error   *ManagerPullPushDomainBindResResponseMetadataError `json:"Error,omitempty"`
}

type ManagerPullPushDomainBindResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type RejectDomainBody struct {

	// REQUIRED; 域名
	Domain string `json:"Domain"`
}

type RejectDomainRes struct {

	// REQUIRED
	ResponseMetadata RejectDomainResResponseMetadata `json:"ResponseMetadata"`
	Result           *RejectDomainResResult          `json:"Result,omitempty"`
}

type RejectDomainResResponseMetadata struct {

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

type RejectDomainResResult struct {

	// REQUIRED; VKE VMP 工作空间记录列表
	VkeVMPWorkspaceRecordList []interface{} `json:"VkeVMPWorkspaceRecordList"`
}

type RestartPullToPushTaskBody struct {

	// REQUIRED; 任务 ID，任务的唯一标识，您可以通过获取拉流转推任务列表 [https://www.volcengine.com/docs/6469/1126896]接口获取。
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

type StopPullCDNSnapshotTaskBody struct {

	// REQUIRED; 任务id
	TaskID string `json:"TaskId"`
}

type StopPullCDNSnapshotTaskRes struct {

	// REQUIRED
	ResponseMetadata StopPullCDNSnapshotTaskResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type StopPullCDNSnapshotTaskResResponseMetadata struct {

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

type StopPullRecordTaskBody struct {

	// REQUIRED; 停止任务的TaskId
	TaskID string `json:"TaskId"`
}

type StopPullRecordTaskRes struct {

	// REQUIRED
	ResponseMetadata StopPullRecordTaskResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type StopPullRecordTaskResResponseMetadata struct {

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

type StopPullToPushTaskBody struct {

	// REQUIRED; 任务 ID，任务的唯一标识，您可以通过获取拉流转推任务列表 [https://www.volcengine.com/docs/6469/1126896]接口获取。
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

type TerminateInstanceBody struct {

	// REQUIRED; 账号
	AccountID string `json:"AccountID"`

	// REQUIRED; 实例号
	InstanceNo string `json:"InstanceNo"`

	// REQUIRED; 产品
	Product string `json:"Product"`
}

type TerminateInstanceRes struct {

	// REQUIRED
	ResponseMetadata TerminateInstanceResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type TerminateInstanceResResponseMetadata struct {

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

type UnBindEncryptDRMBody struct {

	// REQUIRED; 域名空间
	Vhost string `json:"Vhost"`

	// app
	App *string `json:"App,omitempty"`
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

	// REQUIRED; 需要解绑证书的域名。
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

type UpdateActivityBillingBody struct {

	// 不填则更新为空
	ActivityBilling *UpdateActivityBillingBodyActivityBilling `json:"ActivityBilling,omitempty"`
}

// UpdateActivityBillingBodyActivityBilling - 不填则更新为空
type UpdateActivityBillingBodyActivityBilling struct {

	// REQUIRED; 活动条目列表
	Activity []UpdateActivityBillingBodyActivityBillingActivityItem `json:"Activity"`

	// REQUIRED; 检测条件
	Detect UpdateActivityBillingBodyActivityBillingDetect `json:"Detect"`

	// REQUIRED; 当前配置是否生效，1：生效，0：不生效
	Switch int32 `json:"Switch"`
}

type UpdateActivityBillingBodyActivityBillingActivityItem struct {

	// REQUIRED; 日期
	Date string `json:"Date"`

	// REQUIRED; 条目列表
	FeeDetailList []UpdateActivityBillingBodyActivityBillingActivityPropertiesItemsItem `json:"FeeDetailList"`
}

type UpdateActivityBillingBodyActivityBillingActivityPropertiesItemsItem struct {

	// REQUIRED
	ProcDetailList []UpdateActivityBillingBodyActivityBillingActivityPropertiesItemsProcDetailListItem `json:"ProcDetailList"`

	// REQUIRED; 协议
	Protocol string `json:"Protocol"`
}

type UpdateActivityBillingBodyActivityBillingActivityPropertiesItemsProcDetailListItem struct {

	// REQUIRED; 区域
	AreaName string `json:"AreaName"`

	// REQUIRED; 带宽，单位Gbps
	Bandwidth float32 `json:"Bandwidth"`
}

// UpdateActivityBillingBodyActivityBillingDetect - 检测条件
type UpdateActivityBillingBodyActivityBillingDetect struct {

	// 突发增长量场景
	BandwidthCondition *UpdateActivityBillingBodyActivityBillingDetectBandwidthCondition `json:"BandwidthCondition,omitempty"`

	// 日峰值带宽突发增长量
	BandwidthIncrCondition *UpdateActivityBillingBodyActivityBillingDetectBandwidthIncrCondition `json:"BandwidthIncrCondition,omitempty"`

	// 请求数场景
	RequestBandwidthCondition *UpdateActivityBillingBodyActivityBillingDetectRequestBandwidthCondition `json:"RequestBandwidthCondition,omitempty"`
}

// UpdateActivityBillingBodyActivityBillingDetectBandwidthCondition - 突发增长量场景
type UpdateActivityBillingBodyActivityBillingDetectBandwidthCondition struct {

	// REQUIRED; 增量数值超过 xx 的场景xx,单位Gbps
	BandwidthIncr float32 `json:"BandwidthIncr"`

	// REQUIRED; 突发增长量超过最近一个月日峰月均带宽值的x，增长倍数
	BandwidthIncrLoop float32 `json:"BandwidthIncrLoop"`

	// REQUIRED; 1：开启，0：关闭
	Switch int32 `json:"Switch"`
}

// UpdateActivityBillingBodyActivityBillingDetectBandwidthIncrCondition - 日峰值带宽突发增长量
type UpdateActivityBillingBodyActivityBillingDetectBandwidthIncrCondition struct {

	// REQUIRED; 日峰值带宽突发增长量大于 xx 的场景，增量带宽，单位Gbps
	BandwidthIncr float32 `json:"BandwidthIncr"`

	// REQUIRED; 1：开启，0：关闭
	Switch int32 `json:"Switch"`
}

// UpdateActivityBillingBodyActivityBillingDetectRequestBandwidthCondition - 请求数场景
type UpdateActivityBillingBodyActivityBillingDetectRequestBandwidthCondition struct {

	// REQUIRED; 日峰月均值不低于 xx 的场景,日峰值月平均带宽，单位Gbps
	Bandwidth float32 `json:"Bandwidth"`

	// REQUIRED; 请求数超过近一个月的日峰月均值的x倍，增加倍速
	RequestLoop float32 `json:"RequestLoop"`

	// REQUIRED; 1：开启，0：关闭
	Switch int32 `json:"Switch"`
}

type UpdateActivityBillingRes struct {

	// REQUIRED
	ResponseMetadata UpdateActivityBillingResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateActivityBillingResResponseMetadata struct {

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

type UpdateAppBody struct {

	// REQUIRED; appid
	AppID int32 `json:"AppID"`

	// 应用名称中文名
	AppCnName *string `json:"AppCnName,omitempty"`

	// 应用名称英文名
	AppEnName *string `json:"AppEnName,omitempty"`
}

type UpdateAppRes struct {

	// REQUIRED
	ResponseMetadata UpdateAppResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateAppResResponseMetadata struct {

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

type UpdateAuthKeyBody struct {

	// REQUIRED; 鉴权详情，数量阈值为 100。
	AuthDetailList []UpdateAuthKeyBodyAuthDetailListItem `json:"AuthDetailList"`

	// REQUIRED; 鉴权场景类型。
	// * push：推流鉴权；
	// * pull：拉流鉴权；
	SceneType string `json:"SceneType"`

	// 应用名称，默认为所有应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App *string `json:"App,omitempty"`

	// 推/拉流域名。 :::tip 参数 Domain 和 Vhost 传且仅传一个。 :::
	Domain *string `json:"Domain,omitempty"`

	// 鉴权状态。创建推拉流鉴权时，默认值为 false；更新推拉流鉴权时，缺省情况表示不修改推拉流鉴权状态。
	// * false：关闭推拉流鉴权；
	// * true：开启推拉流鉴权。
	PushPullEnable *bool `json:"PushPullEnable,omitempty"`

	// 有效时长，单位为 s，默认值为 604800，取值范围为 [60,2592000]。
	ValidDuration *int32 `json:"ValidDuration,omitempty"`

	// 域名空间名称。 :::tip 参数 Domain 和 Vhost 传且仅传一个。 :::
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

type UpdateAvSlicePresetBody struct {
	AccessKey   *string `json:"AccessKey,omitempty"`
	AccountID   *string `json:"AccountID,omitempty"`
	App         *string `json:"App,omitempty"`
	Bucket      *string `json:"Bucket,omitempty"`
	Callback    *string `json:"Callback,omitempty"`
	Description *string `json:"Description,omitempty"`
	NssConfig   *string `json:"NssConfig,omitempty"`
	Preset      *string `json:"Preset,omitempty"`
	Status      *int32  `json:"Status,omitempty"`
	Stream      *string `json:"Stream,omitempty"`
	TosCluster  *string `json:"TosCluster,omitempty"`
	TosDC       *string `json:"TosDC,omitempty"`
	TosPSM      *string `json:"TosPSM,omitempty"`
	Vhost       *string `json:"Vhost,omitempty"`
}

type UpdateAvSlicePresetRes struct {

	// REQUIRED
	ResponseMetadata UpdateAvSlicePresetResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateAvSlicePresetResResponseMetadata struct {

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

type UpdateBillingBody struct {

	// 标准直播计费项，支持以下取值 live-traffic: 日流量月结 live-day-bandwidth：带宽日峰值月结 live-month-bandwidth：带宽月95峰值月结 live-bandwidth-daily：直播日峰值带宽日结
	// live-traffic-daily：直播流量日结 live-bandwidth-95daily：直播日95带宽日结
	// live-month-bandwidth-average：按日峰值月平均计费 live-month-bandwidth-95average：日95峰月平均计费 live-month-bandwidth-inner：对内客户
	BillingType *string `json:"BillingType,omitempty"`

	// 低延迟直播计费项，不填跟随标准直播取值，支持以下取值 live-day-bandwidth：带宽日峰值月结 live-month-bandwidth：带宽月95峰值月结 live-month-bandwidth-average：按日峰值月平均计费
	// live-month-bandwidth-95average：日95峰月平均计费
	BillingTypeQuic *string `json:"BillingTypeQuic,omitempty"`

	// 低延迟直播计费项，不填跟随标准直播取值，支持以下取值 live-day-bandwidth：带宽日峰值月结 live-month-bandwidth：带宽月95峰值月结 live-month-bandwidth-average：按日峰值月平均计费
	// live-month-bandwidth-95average：日95峰月平均计费
	BillingTypeRTM *string `json:"BillingTypeRTM,omitempty"`

	// 自定义计费方式，入参为以为样式marshal后的json串： {"key1":"value1","key2":"value2"} key和value取值参考：数据工程 [https://bytedance.larkoffice.com/docx/Dqkvd8WAgogvjwxwlMpcW9HznIg]
	CustomBilling *string `json:"CustomBilling,omitempty"`

	// 海外标准直播计费项，不填跟随国内标准直播取值，BillingType为日结方式时，该值必须与BillingType相同，如果为月结方式，则支持以下取值： live-day-bandwidth：带宽日峰值月结 live-month-bandwidth：带宽月95峰值月结
	// live-month-bandwidth-average：按日峰值月平均计费
	// live-month-bandwidth-95average：日95峰月平均计费
	OverseaBillingType *string `json:"OverseaBillingType,omitempty"`

	// 海外Quic直播计费项，不填跟随国内标准直播取值，月结时支持以下取值 live-day-bandwidth：带宽日峰值月结 live-month-bandwidth：带宽月95峰值月结 live-month-bandwidth-average：按日峰值月平均计费
	// live-month-bandwidth-95average：日95峰月平均计费
	OverseaBillingTypeQuic *string `json:"OverseaBillingTypeQuic,omitempty"`

	// 海外低延迟直播计费项，不填跟随国内标准直播取值，月结时支持以下取值 live-day-bandwidth：带宽日峰值月结 live-month-bandwidth：带宽月95峰值月结 live-month-bandwidth-average：按日峰值月平均计费
	// live-month-bandwidth-95average：日95峰月平均计费
	OverseaBillingTypeRTM *string `json:"OverseaBillingTypeRTM,omitempty"`

	// 海外标准直播计费方式，0：拆分大区计费，1：海外统一���费，默认为0
	OverseaChargeMode *int32 `json:"OverseaChargeMode,omitempty"`

	// 海外Quic直播计费方式，0：拆分大区计费，1：海外统一计费，默认为0
	OverseaChargeModeQuic *int32 `json:"OverseaChargeModeQuic,omitempty"`

	// 海外低延迟直播计费方式，0：拆分大区计费，1：海外统一计费，默认为0
	OverseaChargeModeRTM *int32 `json:"OverseaChargeModeRTM,omitempty"`
}

type UpdateBillingRes struct {

	// REQUIRED
	ResponseMetadata UpdateBillingResResponseMetadata `json:"ResponseMetadata"`
}

type UpdateBillingResResponseMetadata struct {
	Action    *string                                `json:"Action,omitempty"`
	Error     *UpdateBillingResResponseMetadataError `json:"Error,omitempty"`
	Region    *string                                `json:"Region,omitempty"`
	RequestID *string                                `json:"RequestId,omitempty"`
	Service   *string                                `json:"Service,omitempty"`
	Version   *string                                `json:"Version,omitempty"`
}

type UpdateBillingResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
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

	// REQUIRED; 消息类型。包括以下类型。
	// * push：推流开始回调；
	// * push_end：推流结束回调；
	// * snapshot：截图回调；
	// * record：录制回调；
	// * audit_snapshot：截图审核回调。
	MessageType string `json:"MessageType"`

	// domain / app 二选一必传
	App *string `json:"App,omitempty"`

	// Dictionary of
	AppendField map[string]*string `json:"AppendField,omitempty"`

	// 是否开启鉴权，默认为 false。取值及含义如下所示。
	// * false：不开启；
	// * true：开启。
	AuthEnable *bool `json:"AuthEnable,omitempty"`

	// Dictionary of
	AuthField map[string]*string `json:"AuthField,omitempty"`

	// 密钥。 :::tip 如果 AuthEnable 为 true，则密钥必填。 :::
	AuthKeyPrimary *string   `json:"AuthKeyPrimary,omitempty"`
	AuthKeySecond  *string   `json:"AuthKeySecond,omitempty"`
	CallbackField  []*string `json:"CallbackField,omitempty"`

	// 推流域名。Vhost和Domain传且仅传一个。
	Domain              *string   `json:"Domain,omitempty"`
	EncryptField        []*string `json:"EncryptField,omitempty"`
	EncryptionAlgorithm *string   `json:"EncryptionAlgorithm,omitempty"`
	HTTPMethod          *string   `json:"HttpMethod,omitempty"`
	RetryInternalSecond *int32    `json:"RetryInternalSecond,omitempty"`
	RetryTimes          *int32    `json:"RetryTimes,omitempty"`
	SecHandlerType      *string   `json:"SecHandlerType,omitempty"`

	// 任务状态回调开关 0关闭 1开启
	TaskStatusCallback *int32 `json:"TaskStatusCallback,omitempty"`
	TimeoutSecond      *int32 `json:"TimeoutSecond,omitempty"`

	// 是否开启转码流回调，默认为 0。取值及含义如下所示。
	// * 0：false，不开启；
	// * 1：true，开启。
	TranscodeCallback *int32 `json:"TranscodeCallback,omitempty"`
	ValidDuration     *int32 `json:"ValidDuration,omitempty"`

	// domain / app 二选一必传
	Vhost *string `json:"Vhost,omitempty"`
}

type UpdateCallbackBodyCallbackDetailListItem struct {

	// REQUIRED; 回调类型，支持设置为 HTTP，表示可以使用 HTTP 和 HTTPS 接收回调事件。
	CallbackType string `json:"CallbackType"`

	// REQUIRED; 回调的 URL。
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

type UpdateCertBody struct {

	// REQUIRED; 更新后的证书名称
	CertName string `json:"CertName"`

	// REQUIRED; 需要更新证书名称的证书链 ID，可以通过查询证书列表 [https://www.volcengine.com/docs/6469/81242]接口获取
	ChainID string `json:"ChainID"`

	// REQUIRED; 证书用途，支持的取值包括：
	// * https：https 认证；
	// * sign：签名校验。
	UseWay    string             `json:"UseWay"`
	AccountID *string            `json:"AccountID,omitempty"`
	Rsa       *UpdateCertBodyRsa `json:"Rsa,omitempty"`
}

type UpdateCertBodyRsa struct {
	Prikey *string `json:"prikey,omitempty"`
	Pubkey *string `json:"pubkey,omitempty"`
}

type UpdateCertRes struct {

	// REQUIRED
	ResponseMetadata UpdateCertResResponseMetadata `json:"ResponseMetadata"`
	Result           *UpdateCertResResult          `json:"Result,omitempty"`
}

type UpdateCertResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version   string                              `json:"Version"`
	Error     *UpdateCertResResponseMetadataError `json:"Error,omitempty"`
	RequestID *string                             `json:"RequestID,omitempty"`
}

type UpdateCertResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateCertResResult struct {
	AccountID *string `json:"AccountID,omitempty"`

	// 证书 ID
	ChainID *string `json:"ChainID,omitempty"`

	// 与证书绑定的域名
	Domain *string `json:"Domain,omitempty"`

	// 证书用途，支持的取值包括：
	// * https：https 认证；
	// * sign：签名校验。
	UseWay *string `json:"UseWay,omitempty"`
}

type UpdateDenseSnapshotPresetBody struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 密集抽帧截图配置模板名称。
	Preset string `json:"Preset"`

	// REQUIRED; 域名空间名称。
	Vhost     string  `json:"Vhost"`
	AccessKey *string `json:"AccessKey,omitempty"`
	AsLong    *int32  `json:"AsLong,omitempty"`
	AsShort   *int32  `json:"AsShort,omitempty"`

	// ToS 的存储 Bucket。 :::tipBucket 与 ServiceID 传且仅传一个。 :::
	Bucket *string `json:"Bucket,omitempty"`

	// 回调地址，支持 HTTP 和 HTTPS 的回调地址。如果同时使用 UpdateCallback 配置了回调地址，则此处回调地址配置优先级更高。
	CallbackURL *string `json:"CallbackUrl,omitempty"`
	Describe    *string `json:"Describe,omitempty"`
	Field36     *string `json:"Field36,omitempty"`
	Format      *string `json:"Format,omitempty"`
	Height      *int32  `json:"Height,omitempty"`

	// 截图间隔时间，单位为 s，默认为 10s，取值范围为正整数。
	Interval        *int32  `json:"Interval,omitempty"`
	IsTobSnapshot   *int32  `json:"IsTobSnapshot,omitempty"`
	KafkaCluster    *string `json:"KafkaCluster,omitempty"`
	KafkaTopic      *string `json:"KafkaTopic,omitempty"`
	Object          *string `json:"Object,omitempty"`
	OverwriteObject *string `json:"OverwriteObject,omitempty"`
	PlatformType    *string `json:"PlatformType,omitempty"`
	Product         *string `json:"Product,omitempty"`
	Quality         *int32  `json:"Quality,omitempty"`
	Rate            *int32  `json:"Rate,omitempty"`
	Region          *string `json:"Region,omitempty"`
	S3NetworkType   *int32  `json:"S3NetworkType,omitempty"`
	SequenceObject  *string `json:"SequenceObject,omitempty"`

	// veImageX 的服务 ID。 :::tipBucket 与 ServiceID 传且仅传一个。 :::
	ServiceID *string `json:"ServiceID,omitempty"`

	// 截图格式，支持 jpg 和 png，默认为 jpg。
	SnapshotFormat *string `json:"SnapshotFormat,omitempty"`

	// 存储规则。
	SnapshotObject *string `json:"SnapshotObject,omitempty"`

	// 密集抽帧截图配置模版的开启状态。
	// * 1：开启
	// * 0：关闭
	Status *int32 `json:"Status,omitempty"`

	// ToS 的存储目录，不传为空。仅当传入了SnapshotObject时生效。
	StorageDir      *string `json:"StorageDir,omitempty"`
	TosCluster      *string `json:"TosCluster,omitempty"`
	TosType         *int32  `json:"TosType,omitempty"`
	TranscodeSuffix *string `json:"TranscodeSuffix,omitempty"`
	Width           *int32  `json:"Width,omitempty"`
}

type UpdateDenseSnapshotPresetRes struct {

	// REQUIRED
	ResponseMetadata UpdateDenseSnapshotPresetResResponseMetadata `json:"ResponseMetadata"`
}

type UpdateDenseSnapshotPresetResResponseMetadata struct {

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

type UpdateDenyConfigBody struct {

	// REQUIRED; 黑白名单配置列表。
	DenyConfigList []UpdateDenyConfigBodyDenyConfigListItem `json:"DenyConfigList"`

	// REQUIRED; 推/拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// App 的名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。Domain 和 App 二选一填。
	App *string `json:"App,omitempty"`
}

type UpdateDenyConfigBodyDenyConfigListItem struct {

	// 白名单。
	AllowList []*string `json:"AllowList,omitempty"`

	// 城市
	City *string `json:"City,omitempty"`

	// 大洲
	Continent *string `json:"Continent,omitempty"`

	// 国家码
	Country *string `json:"Country,omitempty"`

	// 黑名单。
	DenyList []*string `json:"DenyList,omitempty"`

	// 格式类型，比如 HTTP、RTMP。
	FmtType []*string `json:"FmtType,omitempty"`

	// 运营商
	ISP *string `json:"ISP,omitempty"`

	// 协议类型，比如 TCP、KCP、QUIC。
	ProType []*string `json:"ProType,omitempty"`

	// 区域
	Region *string `json:"Region,omitempty"`
}

type UpdateDenyConfigRes struct {

	// REQUIRED
	ResponseMetadata UpdateDenyConfigResResponseMetadata `json:"ResponseMetadata"`
}

type UpdateDenyConfigResResponseMetadata struct {

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
	Error   *UpdateDenyConfigResResponseMetadataError `json:"Error,omitempty"`
}

type UpdateDenyConfigResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateDenyConfigV2Body struct {

	// REQUIRED; 黑白名单配置详情
	DenyConfigList []UpdateDenyConfigV2BodyDenyConfigListItem `json:"DenyConfigList"`

	// REQUIRED; 需要设置 IP 黑白名单的推/拉流域名。域名请在工信部完成备案。
	Domain string `json:"Domain"`

	// REQUIRED; 域名空间名称
	Vhost string `json:"Vhost"`

	// App名称，app和domain二选一填
	App *string `json:"App,omitempty"`

	// 服务类型, pull: 拉流，push：推流
	ServiceType *string `json:"ServiceType,omitempty"`
}

// UpdateDenyConfigV2BodyDenyConfigListItem - 黑白名单配置详情
type UpdateDenyConfigV2BodyDenyConfigListItem struct {

	// REQUIRED; 限制类型。
	// * allow：IP 白名单；
	// * deny：IP 黑名单。
	Type string `json:"Type"`

	// 城市限制
	City []*string `json:"City,omitempty"`

	// 国家限制，国家码
	Country []*string `json:"Country,omitempty"`

	// 拉流类型
	FmtType []*string `json:"FmtType,omitempty"`

	// 黑/白名单 IP 列表，最大限制为 100 个。支持 CIDR（无类域间路由），例如，192.168.0.0 或 192.168.0.0/24。
	IPList []*string `json:"IPList,omitempty"`

	// 运营商限制
	ISP []*string `json:"ISP,omitempty"`

	// 传输协议
	ProType []*string `json:"ProType,omitempty"`

	// 省份限制
	Province []*string `json:"Province,omitempty"`

	// 大区限制
	Region []*string `json:"Region,omitempty"`

	// streams名称
	Streams []*string `json:"Streams,omitempty"`
}

type UpdateDenyConfigV2Res struct {

	// REQUIRED
	ResponseMetadata UpdateDenyConfigV2ResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                               `json:"Result,omitempty"`
}

type UpdateDenyConfigV2ResResponseMetadata struct {

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
	Error   *UpdateDenyConfigV2ResResponseMetadataError `json:"Error,omitempty"`
}

type UpdateDenyConfigV2ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateDomainBody struct {

	// 域名。一次只能提交一个域名。域名请在工信部完成备案。
	Domain *string `json:"Domain,omitempty"`

	// 区域，包含四种类型。cn：中国大陆；oversea：海外；cn-global：全球；cn-oversea：海外及港澳台
	Region *string `json:"Region,omitempty"`
}

type UpdateDomainRes struct {

	// REQUIRED
	ResponseMetadata UpdateDomainResResponseMetadata `json:"ResponseMetadata"`
}

type UpdateDomainResResponseMetadata struct {

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
	Error   *UpdateDomainResResponseMetadataError `json:"Error,omitempty"`
}

type UpdateDomainResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateDomainVhostBody struct {

	// REQUIRED; 需要切换的拉流/推流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 目的域名空间。
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

	// 标准DRM的ApiKey
	APIKey *string `json:"APIKey,omitempty"`

	// 向Apple申请的ask
	ApplicationSecretKey *string `json:"ApplicationSecretKey,omitempty"`

	// 证书文件内容
	CertificateFile *string `json:"CertificateFile,omitempty"`

	// 证书文件名
	CertificateFileName *string `json:"CertificateFileName,omitempty"`

	// 证书名称
	CertificateName *string `json:"CertificateName,omitempty"`

	// 私钥密码
	PrivateKey *string `json:"PrivateKey,omitempty"`

	// 私钥文件内容
	PrivateKeyFile *string `json:"PrivateKeyFile,omitempty"`

	// 私钥文件名
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

	// REQUIRED; Whether to enable the configuration for the domain.
	// * true: Enable
	// * false: Disable
	Enable bool `json:"Enable"`

	// REQUIRED; A list of HTTP headers to be added. You can add a maximum of 20 headers.
	HeaderConfigList []UpdateHTTPHeaderConfigBodyHeaderConfigListItem `json:"HeaderConfigList"`

	// REQUIRED; The type of HTTP header to be updated:
	// * 0: In the response sent from an edge server to a client
	// * 1: In the request sent to a third-party origin server during an origin-pull task.
	Phase int32 `json:"Phase"`

	// REQUIRED; The domain name space.
	Vhost string `json:"Vhost"`

	// Whether to remove the original headers.
	// * 0: Keep
	// * 1: Remove
	BlockOriginal *int32 `json:"BlockOriginal,omitempty"`

	// The domain name.
	Domain *string `json:"Domain,omitempty"`
}

type UpdateHTTPHeaderConfigBodyHeaderConfigListItem struct {

	// REQUIRED; The type of the header value:
	// * 0: Constant
	// * 1: Variable
	HeaderFieldType int32 `json:"HeaderFieldType"`

	// The header name which cannot exceed 1024 characters. And the header names must be distinct from each other.
	HeaderKey *string `json:"HeaderKey,omitempty"`

	// The header value. Specify a constant or a variable as the header value. For the
	// * ${domain}: The domain name in the client request. Example: example.com
	//
	//
	// * ${uri}: The path of the client request excluding the query parameters. If the client request is rewritten, this variable
	// represents the rewritten path. Example: /dir/sample.php
	//
	//
	// * ${args}: The query parameters in the client request. If the client request is rewritten, this variable represents the
	// rewritten parameters. Example: color=red&n=10
	//
	//
	// * ${remote_addr}: The IP address of the client sending the request. Example: 10.10.10.10
	//
	//
	// * ${server_addr}: The IP address of the edge server responding to the client request. Example: 10.10.10.10
	//
	//
	// * ${upstream_host}: The domain name in the origin-pull request. Example: example.com
	//
	//
	// * ${upstream_uri}: The path of the origin-pull request excluding the query parameters. If the request is rewritten, this
	// variable represents the rewritten path. Example: /dir/sample.php
	//
	//
	// * ${upstream_args}: The query parameters in the origin-pull request. If the request is rewritten, this variable represents
	// the rewritten parameters. Example: color=red&n=10
	HeaderValue *string `json:"HeaderValue,omitempty"`
}

type UpdateHTTPHeaderConfigRes struct {

	// REQUIRED
	ResponseMetadata UpdateHTTPHeaderConfigResResponseMetadata `json:"ResponseMetadata"`

	// Depending on the requested interface
	Result interface{} `json:"Result,omitempty"`
}

type UpdateHTTPHeaderConfigResResponseMetadata struct {

	// REQUIRED; The interface name of the request, which is a public parameter of the request.
	Action string `json:"Action"`

	// REQUIRED; The requested Region, for example: cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID is the unique identifier for each API request.
	RequestID string `json:"RequestId"`

	// REQUIRED; The requested service is a public parameter of the request.
	Service string `json:"Service"`

	// REQUIRED; The version number of the request, which is a public parameter of the request.
	Version string `json:"Version"`
}

type UpdateHeaderConfigBody struct {

	// REQUIRED; json配置，使用json更新时填写
	Config string `json:"Config"`

	// REQUIRED; 具体的header配置，目前生效
	HeaderConfigListV2 []UpdateHeaderConfigBodyHeaderConfigListV2Item `json:"HeaderConfigListV2"`

	// REQUIRED; 域名空间
	Vhost string `json:"Vhost"`

	// 应用名称
	App *string `json:"App,omitempty"`

	// 具体的header映射值，已经废弃
	HeaderDetailList []*UpdateHeaderConfigBodyHeaderDetailListItem `json:"HeaderDetailList,omitempty"`

	// 头部类型，已经废弃
	HeaderType *string `json:"HeaderType,omitempty"`
}

type UpdateHeaderConfigBodyHeaderConfigListV2Item struct {

	// REQUIRED; header配置映射
	HeaderDetailList []UpdateHeaderConfigBodyHeaderConfigListV2PropertiesItemsItem `json:"HeaderDetailList"`

	// REQUIRED; header的类型，hls,flv,dash
	HeaderType string `json:"HeaderType"`
}

type UpdateHeaderConfigBodyHeaderConfigListV2PropertiesItemsItem struct {
	HeaderKey   *string `json:"HeaderKey,omitempty"`
	HeaderValue *string `json:"HeaderValue,omitempty"`
}

type UpdateHeaderConfigBodyHeaderDetailListItem struct {

	// key
	HeaderKey *string `json:"HeaderKey,omitempty"`

	// value
	HeaderValue *string `json:"HeaderValue,omitempty"`
}

type UpdateHeaderConfigRes struct {

	// REQUIRED
	ResponseMetadata UpdateHeaderConfigResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type UpdateHeaderConfigResResponseMetadata struct {

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

	// REQUIRED; 推/拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; IP 访问限制规则。
	IPAccessRule UpdateIPAccessRuleBodyIPAccessRule `json:"IPAccessRule"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`
}

// UpdateIPAccessRuleBodyIPAccessRule - IP 访问限制规则。
type UpdateIPAccessRuleBodyIPAccessRule struct {

	// REQUIRED; 是否开启当前限制，取值及含义如下。
	// * true: 开启；
	// * false: 关闭。
	Enable bool `json:"Enable"`

	// REQUIRED; 名单中的 IP 信息。
	IPList []string `json:"IPList"`

	// REQUIRED; IP 访问限制的类型，取值及含义如下。
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

type UpdateNSSRewriteConfigBody struct {

	// REQUIRED; 具体的配置
	Config []string `json:"Config"`

	// REQUIRED; 是否开启
	Enable bool `json:"Enable"`

	// REQUIRED; 服务类型
	ServiceType string `json:"ServiceType"`

	// REQUIRED; 域名空间名称
	Vhost string `json:"Vhost"`

	// 应用名称
	App *string `json:"App,omitempty"`

	// debug
	DebugHeader *string `json:"DebugHeader,omitempty"`
}

type UpdateNSSRewriteConfigRes struct {

	// REQUIRED
	ResponseMetadata UpdateNSSRewriteConfigResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type UpdateNSSRewriteConfigResResponseMetadata struct {

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

type UpdatePresetAssociationBody struct {

	// REQUIRED; 模板名称
	PresetName string `json:"PresetName"`

	// REQUIRED; 模板类型， recor:录制 snapshot:密集抽帧 transcode:转码 avslice:音频切片 cdnsnapshot：截图 avextractor timeshift：时移 auditsnapshot：审核截图
	// data_migration watermark：水印
	PresetType string `json:"PresetType"`

	// REQUIRED; 域名空间名称
	Vhost string `json:"Vhost"`

	// 应用名称
	App *string `json:"App,omitempty"`

	// 旧的模板名
	PresetNameOld *string `json:"PresetNameOld,omitempty"`

	// 录制配置
	RecordParams *UpdatePresetAssociationBodyRecordParams `json:"RecordParams,omitempty"`

	// 录制类型：push, pull
	RecordType *string `json:"RecordType,omitempty"`

	// 流名
	Stream *string `json:"Stream,omitempty"`

	// 时移配置
	TimeShiftStruct *UpdatePresetAssociationBodyTimeShiftStruct `json:"TimeShiftStruct,omitempty"`
}

// UpdatePresetAssociationBodyRecordParams - 录制配置
type UpdatePresetAssociationBodyRecordParams struct {

	// REQUIRED; 默认开启转推
	RelayEnable bool `json:"RelayEnable"`

	// 源流录制，1表示录制
	OriginRecord *int32 `json:"OriginRecord,omitempty"`

	// 转码流录制，1表示录制
	TranscodeRecord *int32 `json:"TranscodeRecord,omitempty"`

	// 转码流录制后缀
	TranscodeSuffixList []*string `json:"TranscodeSuffixList,omitempty"`
}

// UpdatePresetAssociationBodyTimeShiftStruct - 时移配置
type UpdatePresetAssociationBodyTimeShiftStruct struct {

	// 是否需要转码流时移
	NeedTranscode *int32 `json:"NeedTranscode,omitempty"`

	// 时移的类型
	TimeShiftType *int32 `json:"TimeShiftType,omitempty"`
}

type UpdatePresetAssociationRes struct {

	// REQUIRED
	ResponseMetadata UpdatePresetAssociationResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdatePresetAssociationResResponseMetadata struct {

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

type UpdateProxyConfigAssociationBody struct {

	// REQUIRED; 账号
	AccountID string `json:"AccountID"`

	// REQUIRED; 代理记录值
	ProxyID int32 `json:"ProxyID"`
}

type UpdateProxyConfigAssociationRes struct {

	// REQUIRED
	ResponseMetadata UpdateProxyConfigAssociationResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateProxyConfigAssociationResResponseMetadata struct {

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

type UpdateProxyConfigBody struct {

	// REQUIRED; 代理记录ID
	ID int32 `json:"ID"`

	// 描述
	Description *string `json:"Description,omitempty"`

	// 生效类型，1：默认生效
	EffectType *int32 `json:"EffectType,omitempty"`

	// 代理模式，0：固定模式，1：解析模式
	Mode *int32 `json:"Mode,omitempty"`

	// 代理名称
	Name *string `json:"Name,omitempty"`

	// 代理列表
	ProxyConfigList []*UpdateProxyConfigBodyProxyConfigListItem `json:"ProxyConfigList,omitempty"`
}

type UpdateProxyConfigBodyProxyConfigListItem struct {

	// REQUIRED; 集群
	Cluster string `json:"Cluster"`

	// REQUIRED; 机房
	IDC string `json:"IDC"`

	// REQUIRED; 运营商
	ISP string `json:"ISP"`

	// REQUIRED; 地址列表
	ProxyList UpdateProxyConfigBodyProxyConfigListItemProxyList `json:"ProxyList"`
}

// UpdateProxyConfigBodyProxyConfigListItemProxyList - 地址列表
type UpdateProxyConfigBodyProxyConfigListItemProxyList struct {

	// REQUIRED; 地址
	URL string `json:"URL"`

	// REQUIRED; 权重
	Weight int32 `json:"Weight"`
}

type UpdateProxyConfigRes struct {

	// REQUIRED
	ResponseMetadata UpdateProxyConfigResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateProxyConfigResResponseMetadata struct {

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

	// REQUIRED; PushPriority：设置点播视频转推至第三方推流域名时是否使用推流优先级参数，缺省情况下表示不使用此参数，支持的取值及含义如下。
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

	// 点播视频文件循环播放模式，当拉流来源类型为点播视频（Type 为 1）时为必选参数，参数取值及含义如下所示。
	// * -1：无限循环，至任务结束；
	// * 0：有限次循环，循环次数为 PlayTimes 取值为准。
	CycleMode *int32 `json:"CycleMode,omitempty"`

	// 推流域名，推流地址（DstAddr）为空时必传；反之，则该参数不生效
	Domain *string `json:"Domain,omitempty"`

	// 推流地址，即直播源或点播视频转推的目标地址。
	DstAddr *string `json:"DstAddr,omitempty"`

	// 点播文件启播时间偏移值，单位为秒，仅当点播视频播放地址列表（SrcAddrS）只有一个地址，且未配置 Offsets 时生效，缺省情况下表示不进行偏移。
	Offset *float32 `json:"Offset,omitempty"`

	// 点播文件启播时间偏移值，单位为秒，数量与拉流地址列表中地址数量相等，缺省情况下表示不进行偏移。 拉流来源类型为点播视频（Type 为 1）时，参数生效。
	OffsetS []*float32 `json:"OffsetS,omitempty"`

	// 点播视频文件循环播放次数，当循环播放模式为有限次循环（CycleMode为0）时为必选参数。
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

	// 拉流转推任务的名称，由 1 到 20 位中文、大小写字母和数字组成，默认为空，表示不修改任务名称。
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

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 模版名称，由 1 到 50 位数字、字母、下划线及"-"和"."组成。可调用 ListVhostRecordPresetV2 [https://www.volcengine.com/docs/6469/1126858]
	// 接口，查询模版名称。
	Preset string `json:"Preset"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 录制模板详细配置。
	RecordPresetConfig *UpdateRecordPresetV2BodyRecordPresetConfig `json:"RecordPresetConfig,omitempty"`

	// 流名称。
	Stream *string `json:"Stream,omitempty"`
}

// UpdateRecordPresetV2BodyRecordPresetConfig - 录制模板详细配置。
type UpdateRecordPresetV2BodyRecordPresetConfig struct {

	// FLV 录制参数，开启 FLV 录制时设置。 :::tipFlvParam、HlsParam、Mp4Param至少开启一个。 :::
	FlvParam *UpdateRecordPresetV2BodyRecordPresetConfigFlvParam `json:"FlvParam,omitempty"`

	// HLS 录制参数，开启 HLS 录制时设置。 :::tipFlvParam、HlsParam、Mp4Param至少开启一个。 :::
	HlsParam *UpdateRecordPresetV2BodyRecordPresetConfigHlsParam `json:"HlsParam,omitempty"`

	// MP4 录制参数，开启 MP4 录制时设置。 :::tipFlvParam、HlsParam、Mp4Param至少开启一个。 :::
	Mp4Param *UpdateRecordPresetV2BodyRecordPresetConfigMp4Param `json:"Mp4Param,omitempty"`

	// 源流录制，默认值为 0。支持的取值如下所示。
	// * 0：不录制；
	// * 1：录制。
	// :::tipTranscodeRecord和OriginRecord的取值至少一个不为 0。 :::
	OriginRecord *int32 `json:"OriginRecord,omitempty"`

	// 录制 HLS 格式时，单个 TS 切片时长，单位为 s，默认值为 5，取值范围为 [5,30]。
	SliceDuration *int32 `json:"SliceDuration,omitempty"`

	// 转码流录制，默认值为 0。支持的取值如下所示。
	// * 0：不录制；
	// * 1：录制。
	// * 2：全部录制，如果录制转码流后缀列表（TranscodeSuffixList）为空则全部录制，不为空则录制 TranscodeSuffixList 命中的转码后缀。
	// :::tipTranscodeRecord 和 OriginRecord 的取值至少一个不为 0。 :::
	TranscodeRecord *int32 `json:"TranscodeRecord,omitempty"`

	// 录制转码流后缀列表，转码流录制配置为全部录制时（TranscodeRecord 配置等于 2）生效。
	TranscodeSuffixList []*string `json:"TranscodeSuffixList,omitempty"`
}

// UpdateRecordPresetV2BodyRecordPresetConfigFlvParam - FLV 录制参数，开启 FLV 录制时设置。 :::tipFlvParam、HlsParam、Mp4Param至少开启一个。 :::
type UpdateRecordPresetV2BodyRecordPresetConfigFlvParam struct {

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
	TOSParam *UpdateRecordPresetV2BodyRecordPresetConfigFlvParamTOSParam `json:"TOSParam,omitempty"`

	// VOD 存储相关配置。 :::tipTOSParam和VODParam配置且配置其中一个。 :::
	VODParam *UpdateRecordPresetV2BodyRecordPresetConfigFlvParamVODParam `json:"VODParam,omitempty"`
}

// UpdateRecordPresetV2BodyRecordPresetConfigFlvParamTOSParam - TOS 存储相关配置。 :::tipTOSParam和VODParam配置且配置其中一个。 :::
type UpdateRecordPresetV2BodyRecordPresetConfigFlvParamTOSParam struct {

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

// UpdateRecordPresetV2BodyRecordPresetConfigFlvParamVODParam - VOD 存储相关配置。 :::tipTOSParam和VODParam配置且配置其中一个。 :::
type UpdateRecordPresetV2BodyRecordPresetConfigFlvParamVODParam struct {

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

// UpdateRecordPresetV2BodyRecordPresetConfigHlsParam - HLS 录制参数，开启 HLS 录制时设置。 :::tipFlvParam、HlsParam、Mp4Param至少开启一个。 :::
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

// UpdateRecordPresetV2BodyRecordPresetConfigMp4Param - MP4 录制参数，开启 MP4 录制时设置。 :::tipFlvParam、HlsParam、Mp4Param至少开启一个。 :::
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

	// REQUIRED; Referer 防盗链信息列表。
	RefererInfoList []UpdateRefererBodyRefererInfoListItem `json:"RefererInfoList"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 应用名称，默认为所有应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。 :::tip 参数 Domain 和 App 传且仅传一个。 :::
	App *string `json:"App,omitempty"`

	// 拉流域名。 :::tip 参数 Domain 和 App 传且仅传一个。 :::
	Domain *string `json:"Domain,omitempty"`
}

type UpdateRefererBodyRefererInfoListItem struct {

	// REQUIRED; 用于标识 referer 防盗链的关键词默认取值为 referer。
	Key string `json:"Key"`

	// REQUIRED; 防盗链类型，支持如下取值。
	// * deny：黑名单；
	// * allow：白名单。
	Type string `json:"Type"`

	// 指定域名的优先级。默认值为 0，取值范围为 [0,100]，数值越大，优先级越高。如果优先级相同，则越早加入列表的域名优先级越高。
	Priority *int32 `json:"Priority,omitempty"`

	// 防盗链规则，即设置的黑名单或白名单的域名，最大长度限制 300 个字符。
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

type UpdateRelaySinkBody struct {

	// REQUIRED
	Group string `json:"Group"`

	// REQUIRED
	IsThunderRelaySink bool `json:"IsThunderRelaySink"`

	// REQUIRED
	IsTranscodeRelaySink bool `json:"IsTranscodeRelaySink"`

	// REQUIRED
	PassRequestParam bool `json:"PassRequestParam"`

	// REQUIRED
	RelaySinkDetailList []UpdateRelaySinkBodyRelaySinkDetailListItem `json:"RelaySinkDetailList"`

	// REQUIRED
	Vhost string  `json:"Vhost"`
	App   *string `json:"App,omitempty"`
}

type UpdateRelaySinkBodyRelaySinkDetailListItem struct {

	// REQUIRED
	CDN string `json:"CDN"`

	// REQUIRED
	PullDomainList []UpdateRelaySinkBodyRelaySinkDetailListPropertiesItemsItem `json:"PullDomainList"`

	// REQUIRED
	RelaySinkDomain string  `json:"RelaySinkDomain"`
	AK              *string `json:"AK,omitempty"`
	PushAuth        *bool   `json:"PushAuth,omitempty"`
	RelaySinkApp    *string `json:"RelaySinkApp,omitempty"`

	// Anything
	RelaySinkParams interface{} `json:"RelaySinkParams,omitempty"`
	SK              *string     `json:"SK,omitempty"`
	Status          *int32      `json:"Status,omitempty"`
	ValidDuration   *int32      `json:"ValidDuration,omitempty"`
	Weight          *int32      `json:"Weight,omitempty"`
}

type UpdateRelaySinkBodyRelaySinkDetailListPropertiesItemsItem struct {

	// REQUIRED
	Protocol string `json:"Protocol"`

	// REQUIRED
	PullDomain string `json:"PullDomain"`
}

type UpdateRelaySinkRes struct {

	// REQUIRED
	ResponseMetadata UpdateRelaySinkResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type UpdateRelaySinkResResponseMetadata struct {

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

	// REQUIRED; The domain name space
	Vhost string `json:"Vhost"`

	// The domain name
	Domain *string `json:"Domain,omitempty"`

	// The origin address rewrite configurations
	RewriteRule *UpdateRelaySourceRewriteBodyRewriteRule `json:"RewriteRule,omitempty"`
}

// UpdateRelaySourceRewriteBodyRewriteRule - The origin address rewrite configurations
type UpdateRelaySourceRewriteBodyRewriteRule struct {

	// REQUIRED; Whether to enable origin address rewrite
	// * true: Enable
	// * false: Disable
	Enable bool `json:"Enable"`

	// REQUIRED; A list of rewrite rules
	RewriteRuleList []UpdateRelaySourceRewriteBodyRewriteRuleListItem `json:"RewriteRuleList"`
}

type UpdateRelaySourceRewriteBodyRewriteRuleListItem struct {

	// Whether to include the query parameters of the original path in the new path.
	// * true: Include
	// * false: Exclude
	IncludeParams *bool `json:"IncludeParams,omitempty"`

	// The existing path. The path may include a maximum of five wildcards. However, do not use wildcards to represent query parameters.
	// The path must satisfy the following requirements:
	// * The path must not contain the protocol and domain name, such as http://example.domain.com.
	// * The path must start with /.
	// * The path can contain a maximum of 1,024 characters.
	// * The path cannot contain two consecutive asterisks, i.e. **.
	// * The path cannot contain two consecutive instances of ${, i.e. ${${.
	OriginPath *string `json:"OriginPath,omitempty"`

	// The new path. You can use ${1}, ${2}, ${3}, ${4}, and ${5} to represent the content matched by the first to the fifth wildcards
	// respectively. The path must satisfy the following requirements:
	// * The path must not contain the protocol and domain name, such as http://example.domain.com.
	// * The path must start with /.
	// * The path can contain a maximum of 1,024 characters.
	// * The path cannot contain two consecutive asterisks, i.e. **.
	// * The path cannot contain two consecutive instances of ${, i.e. ${${.
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

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console-stable.volcanicengine.com/live/main/domain/list]页面，查看直播流使用的域名所属的域名空间。
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

	// REQUIRED; 应用名称，即直播流地址的AppName字段取值，支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 直播流使用的域名。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console-stable.volcanicengine.com/live/main/domain/list]页面，查看直播流使用的域名。
	Domain string `json:"Domain"`

	// REQUIRED; 回源地址列表，支持 RTMP、FLV、HLS 和 SRT 回源协议。
	// :::tip
	// * 当源站使用了非默认端口时，支持在回源地址中以域名:端口或IP:端口的形式配置端口。
	// * 最多支持添加 10 个回源地址，回源失败时，将按照您添加的地址顺序轮循尝试。 :::
	SrcAddrS []string `json:"SrcAddrS"`

	// REQUIRED; 流名称，即直播流地址的StreamName字段取值，支持由大小写字母（A - Z、a - z）、数字（0 - 9）、字母、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
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

type UpdateSDKBody struct {

	// REQUIRED; sdk记录ID
	ID int32 `json:"ID"`

	// App名称
	AppName *string `json:"AppName,omitempty"`

	// App英文名称
	AppNameEN *string `json:"AppNameEN,omitempty"`

	// SDK版本，精简版：1、互动版：2，已经弃用
	SDKVersion *int32 `json:"SDKVersion,omitempty"`

	// 状态，0：未激活，1：激活，2：审批，3：过期，4：删除，5：试用过期，6：正式过期，7：试用激活，8：彻底删除
	Status *int32 `json:"Status,omitempty"`

	// 要迁移的目标账号
	TargetAccountID *string `json:"TargetAccountID,omitempty"`

	// 操作类型，1：激活，2：恢复，3：彻底删除
	Type *int32 `json:"Type,omitempty"`
}

type UpdateSDKLicenseBody struct {

	// REQUIRED; SDK记录ID
	ID int32 `json:"ID"`

	// REQUIRED; license类型，1:基础版，2：高级版，3：试用版
	LicenseType int32 `json:"LicenseType"`

	// REQUIRED; 操作类型，1：续签，2：转正（sdk记录只能为试用版），3：延期（sdk的记录只能为试用版）
	Type int32 `json:"Type"`

	// licenseID,续签和转正必填
	LicenseID *int32 `json:"LicenseID,omitempty"`

	// 流量包ID
	PackageID *string `json:"PackageID,omitempty"`
}

type UpdateSDKLicenseRes struct {

	// REQUIRED
	ResponseMetadata UpdateSDKLicenseResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateSDKLicenseResResponseMetadata struct {

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

type UpdateSDKRes struct {

	// REQUIRED
	ResponseMetadata UpdateSDKResResponseMetadata `json:"ResponseMetadata"`
}

type UpdateSDKResResponseMetadata struct {
	Action    *string                            `json:"Action,omitempty"`
	Error     *UpdateSDKResResponseMetadataError `json:"Error,omitempty"`
	Region    *string                            `json:"Region,omitempty"`
	RequestID *string                            `json:"RequestId,omitempty"`
	Service   *string                            `json:"Service,omitempty"`
	Version   *string                            `json:"Version,omitempty"`
}

type UpdateSDKResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type UpdateServiceBody struct {

	// 用户账号
	AccountID *string `json:"AccountID,omitempty"`

	// 公司名称
	CompanyName *string                       `json:"CompanyName,omitempty"`
	LimitConfig *UpdateServiceBodyLimitConfig `json:"LimitConfig,omitempty"`

	// 需要隐藏的面板
	PresetConfigHide []*UpdateServiceBodyPresetConfigHideItem `json:"PresetConfigHide,omitempty"`
}

type UpdateServiceBodyLimitConfig struct {

	// 账号下app限制数目
	AppLimit *int32 `json:"AppLimit,omitempty"`

	// 账号vhost下的domain限制数目
	DomainLimit *int32 `json:"DomainLimit,omitempty"`

	// 账号下的vhost限制数目，-1表示无限制，0表示使用默认配置数目
	VhostLimit *int32 `json:"VhostLimit,omitempty"`
}

type UpdateServiceBodyPresetConfigHideItem struct {

	// 1: 录制是否隐藏TOS 2: 截图是否隐藏TOS 3: 时移是否隐藏VOD 4: 云导播是否隐藏 5：海外加速计费是否隐藏 6：RTM单独加速计费是否隐藏 7：基础版License申请是否隐藏 8：高级版License申请是否隐藏 9：固定回源是否隐藏
	// 10: 月结欠费关停是否处理，1表示处理 11: IP限频是否隐藏 12：URL限频是否隐藏 13：URL参数限频是否隐藏
	// 14：IP访问相同URL限频是否隐藏 15: 活动带宽计费是否隐藏 16: 画质增强是否隐藏 17: Quic加速计费是否隐藏
	ConfigID *int32 `json:"ConfigID,omitempty"`

	// 是否隐藏
	IsHide *bool `json:"IsHide,omitempty"`
}

type UpdateServiceRes struct {

	// REQUIRED
	ResponseMetadata UpdateServiceResResponseMetadata `json:"ResponseMetadata"`
}

type UpdateServiceResResponseMetadata struct {
	Action    *string                                `json:"Action,omitempty"`
	Error     *UpdateServiceResResponseMetadataError `json:"Error,omitempty"`
	Region    *string                                `json:"Region,omitempty"`
	RequestID *string                                `json:"RequestId,omitempty"`
	Service   *string                                `json:"Service,omitempty"`
	Version   *string                                `json:"Version,omitempty"`
}

type UpdateServiceResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type UpdateSnapshotAuditPresetBody struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 截图审核配置的名称，您可以通过调用查询截图审核配置列表 [https://www.volcengine.com/docs/6469/1126870]接口获取。
	PresetName string  `json:"PresetName"`
	AuditType  *string `json:"AuditType,omitempty"`

	// ToS 存储对应的 Bucket。 :::tip 参数 Bucket 和 ServiceID 传且仅传一个。 :::
	Bucket *string `json:"Bucket,omitempty"`

	// 截图审核结果回调地址配置。
	CallbackDetailList []*UpdateSnapshotAuditPresetBodyCallbackDetailListItem `json:"CallbackDetailList,omitempty"`

	// 截图审核配置的描述。
	Description *string `json:"Description,omitempty"`

	// 推流域名。 :::tip 参数 Domain 和 Vhost 传且仅传一个。 :::
	Domain  *string `json:"Domain,omitempty"`
	Field37 *string `json:"Field37,omitempty"`

	// 截图间隔时间，单位秒，取值范围为[0.1,10]，支持保留两位小数。
	Interval *float32 `json:"Interval,omitempty"`

	// 审核标签，缺省情况下取值为 301、302、302、305 和 306，支持的取值及含义如下。
	// * 301：涉黄；
	// * 302：涉敏1；
	// * 303：涉敏2；
	// * 304：广告；
	// * 305：引人不适；
	// * 306：违禁；
	// * 307：二维码；
	// * 308：诈骗；
	// * 309：不良画面；
	// * 310：未成年相关；
	// * 320：文字违规。
	Label []*string `json:"Label,omitempty"`

	// veimageX 的服务 ID。 :::tip 参数 Bucket 和 ServiceID 传且仅传一个。 :::
	ServiceID *string `json:"ServiceID,omitempty"`

	// 截图存储规则，支持以 {Domain}/{App}/{Stream}/{UnixTimestamp} 样式设置存储规则，支持输入字母、数字、-、!、_、.、* 及占位符，最大长度为 180 个字符。
	SnapshotObject *string `json:"SnapshotObject,omitempty"`
	Status         *int32  `json:"Status,omitempty"`

	// ToS 存储对应 Bucket 下的存储目录，默认为空。 例如，存储位置为 live-test-tos-example/live/liveapp 时，StorageDir 取值为 live/liveapp。
	StorageDir *string `json:"StorageDir,omitempty"`

	// 存储策略。支持以下取值。
	// * 0：触发存储，只存储有风险图片；
	// * 1：全部存储，存储全部图片。
	StorageStrategy *int32 `json:"StorageStrategy,omitempty"`

	// 域名空间名称。 :::tip 参数 Domain 和 Vhost 传且仅传一个。 :::
	Vhost *string `json:"Vhost,omitempty"`
}

type UpdateSnapshotAuditPresetBodyCallbackDetailListItem struct {

	// REQUIRED; 回调地址的类型，当前仅支持 http。
	CallbackType string `json:"CallbackType"`

	// REQUIRED; 回调地址。
	URL string `json:"URL"`
}

type UpdateSnapshotAuditPresetRes struct {

	// REQUIRED
	ResponseMetadata UpdateSnapshotAuditPresetResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateSnapshotAuditPresetResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                             `json:"Version"`
	Error   *UpdateSnapshotAuditPresetResResponseMetadataError `json:"Error,omitempty"`
}

type UpdateSnapshotAuditPresetResResponseMetadataError struct {

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

	// REQUIRED
	VodNamespace string  `json:"VodNamespace"`
	ACL          *string `json:"ACL,omitempty"`
	AccessKey    *string `json:"AccessKey,omitempty"`
	AccountID    *string `json:"AccountID,omitempty"`
	AsShort      *int32  `json:"AsShort,omitempty"`

	// ToS 的存储 Bucket。 :::tipBucket 与 ServiceID 传且仅传一个。 :::
	Bucket   *string `json:"Bucket,omitempty"`
	Callback *string `json:"Callback,omitempty"`

	// 回调详情。
	CallbackDetailList []*UpdateSnapshotPresetBodyCallbackDetailListItem `json:"CallbackDetailList,omitempty"`
	Description        *string                                           `json:"Description,omitempty"`
	Duration           *int32                                            `json:"Duration,omitempty"`
	Format             []*string                                         `json:"Format,omitempty"`
	Height             *int32                                            `json:"Height,omitempty"`

	// 截图间隔时间，单位为 s，默认值为 10，取值范围为正整数
	Interval  *int32  `json:"Interval,omitempty"`
	NssConfig *string `json:"NssConfig,omitempty"`

	// 存储方式为覆盖截图时的存储规则，支持以 {Domain}/{App}/{Stream} 样式设置存储规则，支持输入字母、数字、"-"、"!"、"_"、"."、"*"及占位符。
	OverwriteObject  *string                                  `json:"OverwriteObject,omitempty"`
	PlatformTypeList []*string                                `json:"PlatformTypeList,omitempty"`
	PullDomain       *string                                  `json:"PullDomain,omitempty"`
	Quality          *int32                                   `json:"Quality,omitempty"`
	RecordConfig     *string                                  `json:"RecordConfig,omitempty"`
	RecordObject     *string                                  `json:"RecordObject,omitempty"`
	RecordTob        []*UpdateSnapshotPresetBodyRecordTobItem `json:"RecordTob,omitempty"`
	RegionConfig     *string                                  `json:"RegionConfig,omitempty"`
	ReserveDays      *int32                                   `json:"ReserveDays,omitempty"`

	// veImageX 的服务 ID。 :::tipBucket 与 ServiceID 传且仅传一个。 :::
	ServiceID      *string `json:"ServiceID,omitempty"`
	SliceDuration  *int32  `json:"SliceDuration,omitempty"`
	SnapshotConfig *string `json:"SnapshotConfig,omitempty"`

	// 截图格式。支持如下取值。- jpeg - jpg
	SnapshotFormat *string `json:"SnapshotFormat,omitempty"`

	// 存储方式为实时存储时的存储规则，支持以 {Domain}/{App}/{Stream}/{UnixTimestamp} 样式设置存储规则，支持输入字母、数字、"-"、"!"、"_"、"."、"*"及占位符。
	SnapshotObject *string `json:"SnapshotObject,omitempty"`
	Splice         *int32  `json:"Splice,omitempty"`

	// 截图模版状态。
	// * 1：开启
	// * 0：关闭
	Status *int32 `json:"Status,omitempty"`

	// ToS 的存储目录，不传为空。
	StorageDir *string `json:"StorageDir,omitempty"`
	TosCluster *string `json:"TosCluster,omitempty"`
	TosDC      *string `json:"TosDC,omitempty"`
	TosPSM     *string `json:"TosPSM,omitempty"`
	Width      *int32  `json:"Width,omitempty"`
	WorkflowID *string `json:"WorkflowID,omitempty"`
}

type UpdateSnapshotPresetBodyCallbackDetailListItem struct {

	// 回调类型，默认值为 http。
	CallbackType *string `json:"CallbackType,omitempty"`

	// 回调地址。
	URL *string `json:"URL,omitempty"`
}

type UpdateSnapshotPresetBodyRecordTobItem struct {
	Duration     *int32  `json:"Duration,omitempty"`
	Format       *string `json:"Format,omitempty"`
	RecordObject *string `json:"RecordObject,omitempty"`
	Splice       *int32  `json:"Splice,omitempty"`
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

type UpdateSnapshotPresetV2Body struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 截图配置的名称。
	Preset string `json:"Preset"`

	// REQUIRED; 截图配置的详细参数配置。
	SnapshotPresetConfig UpdateSnapshotPresetV2BodySnapshotPresetConfig `json:"SnapshotPresetConfig"`

	// REQUIRED; 域名空间名称。
	Vhost string `json:"Vhost"`

	// 截图配置生效状态，默认为生效。
	// * 1：生效；
	// * 0：不生效。
	Status *int32 `json:"Status,omitempty"`
}

// UpdateSnapshotPresetV2BodySnapshotPresetConfig - 截图配置的详细参数配置。
type UpdateSnapshotPresetV2BodySnapshotPresetConfig struct {

	// 截图间隔时间，单位为秒，默认值为 10，取值范围为正整数。
	Interval *int32 `json:"Interval,omitempty"`

	// 图片格式为 JPEG 时的截图参数，开启 JPEG 截图时设置。 :::tip JPEG 截图和 JPG 截图必须开启且只能开启一个。 :::
	JPEGParam *UpdateSnapshotPresetV2BodySnapshotPresetConfigJPEGParam `json:"JpegParam,omitempty"`

	// 截图格式为 JPG 时的截图参数，开启 JPG 截图时设置。 :::tip JPEG 截图和 JPG 截图必须开启且只能开启一个。 :::
	JpgParam *UpdateSnapshotPresetV2BodySnapshotPresetConfigJpgParam `json:"JpgParam,omitempty"`
}

// UpdateSnapshotPresetV2BodySnapshotPresetConfigJPEGParam - 图片格式为 JPEG 时的截图参数，开启 JPEG 截图时设置。 :::tip JPEG 截图和 JPG 截图必须开启且只能开启一个。
// :::
type UpdateSnapshotPresetV2BodySnapshotPresetConfigJPEGParam struct {

	// 当前格式的截图配置是否开启，默认为 false，取值及含义如下所示。
	// * false：不开启；
	// * true：开启。
	Enable *bool `json:"Enable,omitempty"`

	// 截图存储到 veImageX 时的配置。 :::tip TOSParam 和 ImageXParam 配置且配置其中一个。 :::
	ImageXParam *UpdateSnapshotPresetV2BodySnapshotPresetConfigJPEGParamImageXParam `json:"ImageXParam,omitempty"`

	// 截图存储到 TOS 时的配置。 :::tip TOSParam 和 ImageXParam 配置且配置其中一个。 :::
	TOSParam *UpdateSnapshotPresetV2BodySnapshotPresetConfigJPEGParamTOSParam `json:"TOSParam,omitempty"`
}

// UpdateSnapshotPresetV2BodySnapshotPresetConfigJPEGParamImageXParam - 截图存储到 veImageX 时的配置。 :::tip TOSParam 和 ImageXParam
// 配置且配置其中一个。 :::
type UpdateSnapshotPresetV2BodySnapshotPresetConfigJPEGParamImageXParam struct {

	// 截图是否使用 veImageX 存储，默认为 false，取值及含义如下所示。
	// * false：不使用；
	// * true：使用。
	Enable *bool `json:"Enable,omitempty"`

	// 存储方式为实时存储时的存储规则，支持以 {Domain}/{App}/{Stream}/{UnixTimestamp} 样式设置存储规则，支持输入字母、数字、-、!、_、.、*" 及占位符。 :::tip 参数 ExactObject 和
	// OverwriteObject 传且仅传一个。 :::
	ExactObject *string `json:"ExactObject,omitempty"`

	// 存储方式为覆盖截图时的存储规则，支持以 {Domain}/{App}/{Stream} 样式设置存储规则，支持输入字母、数字、-、!、_、.、* 及占位符。 :::tip 参数 ExactObject 和 OverwriteObject
	// 传且仅传一个。 :::
	OverwriteObject *string `json:"OverwriteObject,omitempty"`

	// 使用 veImageX 存储截图时，对应的 veImageX 的服务 ID。 :::tip 使用 veImageX 存储时 ServiceID 为必填项。 :::
	ServiceID *string `json:"ServiceID,omitempty"`
}

// UpdateSnapshotPresetV2BodySnapshotPresetConfigJPEGParamTOSParam - 截图存储到 TOS 时的配置。 :::tip TOSParam 和 ImageXParam 配置且配置其中一个。
// :::
type UpdateSnapshotPresetV2BodySnapshotPresetConfigJPEGParamTOSParam struct {

	// TOS 存储对应的 Bucket。 例如，存储路径为 live-test-tos-example/live/liveapp 时，Bucket 取值为 live-test-tos-example。 :::tip 使用 TOS 存储时 Bucket
	// 为必填项。 :::
	Bucket *string `json:"Bucket,omitempty"`

	// 截图是否使用 TOS 存储，默认为 false，取值及含义如下所示。
	// * false：不使用；
	// * true：使用。
	Enable *bool `json:"Enable,omitempty"`

	// 存储方式为实时存储时的存储规则，支持以 {Domain}/{App}/{Stream}/{UnixTimestamp} 样式设置存储规则，支持输入字母、数字、-、!、_、.、* 及占位符。 :::tip 参数 ExactObject 和
	// OverwriteObject 传且仅传一个。 :::
	ExactObject *string `json:"ExactObject,omitempty"`

	// 存储方式为覆盖截图时的存储规则，支持以 {Domain}/{App}/{Stream} 样式设置存储规则，支持输入字母、数字、-、!、_、.、* 及占位符。 :::tip 参数 ExactObject 和 OverwriteObject
	// 传且仅传一个。 :::
	OverwriteObject *string `json:"OverwriteObject,omitempty"`

	// Bucket 目录，默认为空。 例如，存储路径为 live-test-tos-example/live/liveapp 时，StorageDir 取值为 live/liveapp。
	StorageDir *string `json:"StorageDir,omitempty"`
}

// UpdateSnapshotPresetV2BodySnapshotPresetConfigJpgParam - 截图格式为 JPG 时的截图参数，开启 JPG 截图时设置。 :::tip JPEG 截图和 JPG 截图必须开启且只能开启一个。
// :::
type UpdateSnapshotPresetV2BodySnapshotPresetConfigJpgParam struct {
	Enable      *bool                                                              `json:"Enable,omitempty"`
	ImageXParam *UpdateSnapshotPresetV2BodySnapshotPresetConfigJpgParamImageXParam `json:"ImageXParam,omitempty"`
	TOSParam    *UpdateSnapshotPresetV2BodySnapshotPresetConfigJpgParamTOSParam    `json:"TOSParam,omitempty"`
}

type UpdateSnapshotPresetV2BodySnapshotPresetConfigJpgParamImageXParam struct {
	Enable          *bool   `json:"Enable,omitempty"`
	ExactObject     *string `json:"ExactObject,omitempty"`
	OverwriteObject *string `json:"OverwriteObject,omitempty"`
	ServiceID       *string `json:"ServiceID,omitempty"`
}

type UpdateSnapshotPresetV2BodySnapshotPresetConfigJpgParamTOSParam struct {
	ACL             *string `json:"ACL,omitempty"`
	AccessKey       *string `json:"AccessKey,omitempty"`
	Bucket          *string `json:"Bucket,omitempty"`
	Enable          *bool   `json:"Enable,omitempty"`
	ExactObject     *string `json:"ExactObject,omitempty"`
	OverwriteObject *string `json:"OverwriteObject,omitempty"`
	Region          *string `json:"Region,omitempty"`
	S3NetworkType   *string `json:"S3NetworkType,omitempty"`
	StorageDir      *string `json:"StorageDir,omitempty"`
	TosCluster      *string `json:"TosCluster,omitempty"`
	TosDC           *string `json:"TosDC,omitempty"`
	TosPSM          *string `json:"TosPSM,omitempty"`
}

type UpdateSnapshotPresetV2Res struct {

	// REQUIRED
	ResponseMetadata UpdateSnapshotPresetV2ResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateSnapshotPresetV2ResResponseMetadata struct {

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

type UpdateStreamQuotaConfigBody struct {

	// REQUIRED; 需要配置限额的推流域名或拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 限额配置详情。
	QuotaDetailList []UpdateStreamQuotaConfigBodyQuotaDetailListItem `json:"QuotaDetailList"`
}

type UpdateStreamQuotaConfigBodyQuotaDetailListItem struct {

	// 拉流域名的带宽限额配置。 :::tipDomain 为拉流域名时，本参数为必选参数。 :::
	BandwidthConfig *UpdateStreamQuotaConfigBodyQuotaDetailListItemBandwidthConfig `json:"BandwidthConfig,omitempty"`

	// 超过限额时返回的错误码，默认值为403。
	ErrCode *int32 `json:"ErrCode,omitempty"`

	// 超过限额时返回的错误信息，默认值为forbid。
	ErrMsg *string `json:"ErrMsg,omitempty"`

	// 推流域名的推流路数限额配置。 :::tipDomain 为推流域名时，本参数为必选参数。 :::
	StreamConfig *UpdateStreamQuotaConfigBodyQuotaDetailListItemStreamConfig `json:"StreamConfig,omitempty"`
}

// UpdateStreamQuotaConfigBodyQuotaDetailListItemBandwidthConfig - 拉流域名的带宽限额配置。 :::tipDomain 为拉流域名时，本参数为必选参数。 :::
type UpdateStreamQuotaConfigBodyQuotaDetailListItemBandwidthConfig struct {

	// REQUIRED; 带宽限额，取值[1~1000]。
	Quota int32 `json:"Quota"`

	// REQUIRED; 拉流带宽限额的计量单位，支持的取值如下所示。
	// * Mbps
	// * Gbps
	// * Tbps
	QuotaUnit string `json:"QuotaUnit"`

	// 拉流带宽限额告警阈值，取值范围为 [1,1000]，缺省情况表示不设置告警。 :::tip 该参数的取值需要小于等于拉流带宽限额Quota，否则会报错。 :::
	AlarmThreshold *int32 `json:"AlarmThreshold,omitempty"`

	// 拉流带宽限额告警的计量单位，缺省情况表示不设置告警。支持的取值如下所示。
	// * Mbps
	// * Gbps
	// * Tbps
	AlarmThresholdUnit *string `json:"AlarmThresholdUnit,omitempty"`
}

// UpdateStreamQuotaConfigBodyQuotaDetailListItemStreamConfig - 推流域名的推流路数限额配置。 :::tipDomain 为推流域名时，本参数为必选参数。 :::
type UpdateStreamQuotaConfigBodyQuotaDetailListItemStreamConfig struct {

	// REQUIRED; 推流路数限额，取值[10~200000]。
	Quota int32 `json:"Quota"`

	// 推流路数限额告警阈值，缺省情况表示不设置告警。取值范围为 [10,200000]。 :::tip 该参数的取值需要小于等于推流路数限额Quota，否则会报错。 :::
	AlarmThreshold *int32 `json:"AlarmThreshold,omitempty"`
}

type UpdateStreamQuotaConfigPatchBody struct {

	// REQUIRED; 批量添加限制，最多30个
	ConfigList []UpdateStreamQuotaConfigPatchBodyConfigListItem `json:"ConfigList"`
}

type UpdateStreamQuotaConfigPatchBodyConfigListItem struct {

	// REQUIRED; 需要配置限额的域名空间
	Vhost string `json:"Vhost"`

	// 限额配置详情。
	QuotaDetailList []*UpdateStreamQuotaConfigPatchBodyConfigListPropertiesItemsItem `json:"QuotaDetailList,omitempty"`
}

// UpdateStreamQuotaConfigPatchBodyConfigListItemQuotaDetailListItemStreamConfig - 目前域名空间只支持推流路数配置
type UpdateStreamQuotaConfigPatchBodyConfigListItemQuotaDetailListItemStreamConfig struct {

	// 推流路数限额告警阈值，缺省情况表示不设置告警。取值范围为 [10,200000]。 :::tip 该参数的取值需要小于等于推流路数限额Quota，否则会报错。 :::
	AlarmThreshold *int32 `json:"AlarmThreshold,omitempty"`

	// 限制的推流的qps，>= -1 0:使用调度默认限制 -1：不限制 xx>0: 具体数据，则限制为xx qps
	QPSLimit *int32 `json:"QPSLimit,omitempty"`

	// 推流路数限额，取值[10~200000]。
	Quota *int32 `json:"Quota,omitempty"`
}

type UpdateStreamQuotaConfigPatchBodyConfigListPropertiesItemsItem struct {

	// 目前域名空间只支持推流路数配置
	StreamConfig *UpdateStreamQuotaConfigPatchBodyConfigListItemQuotaDetailListItemStreamConfig `json:"StreamConfig,omitempty"`
}

type UpdateStreamQuotaConfigPatchRes struct {

	// REQUIRED
	ResponseMetadata UpdateStreamQuotaConfigPatchResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type UpdateStreamQuotaConfigPatchResResponseMetadata struct {

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

type UpdateStreamQuotaConfigRes struct {

	// REQUIRED
	ResponseMetadata UpdateStreamQuotaConfigResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateStreamQuotaConfigResResponseMetadata struct {

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
	Error   *UpdateStreamQuotaConfigResResponseMetadataError `json:"Error,omitempty"`
}

type UpdateStreamQuotaConfigResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateTimeShiftPresetV2Body struct {

	// REQUIRED
	MaxShiftTime int32 `json:"MaxShiftTime"`

	// REQUIRED
	Preset string `json:"Preset"`

	// REQUIRED
	Vhost        string  `json:"Vhost"`
	App          *string `json:"App,omitempty"`
	MasterFormat *string `json:"MasterFormat,omitempty"`
	Status       *int32  `json:"Status,omitempty"`
	Type         *string `json:"Type,omitempty"`
}

type UpdateTimeShiftPresetV2Res struct {

	// REQUIRED
	ResponseMetadata UpdateTimeShiftPresetV2ResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type UpdateTimeShiftPresetV2ResResponseMetadata struct {

	// REQUIRED
	Action string `json:"Action"`

	// REQUIRED
	Region string `json:"Region"`

	// REQUIRED
	Service string `json:"Service"`

	// REQUIRED
	Version   string                                           `json:"Version"`
	Error     *UpdateTimeShiftPresetV2ResResponseMetadataError `json:"Error,omitempty"`
	RequestID *string                                          `json:"RequestID,omitempty"`
}

type UpdateTimeShiftPresetV2ResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateTimeShiftPresetV3Body struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 最大时移时长，即观看时移的最长时间，单位为 s。支持的取值如下所示。
	// * 86400
	// * 259200
	// * 604800
	// * 1296000
	MaxShiftTime int32 `json:"MaxShiftTime"`

	// REQUIRED; 域名空间名称。
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

	// 开启时移的流名称，同一个 App 最多可指定 20 路。
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

	// REQUIRED; 转码配置的名称。
	Preset string `json:"Preset"`

	// REQUIRED; 域名空间名称。
	Vhost     string  `json:"Vhost"`
	ALayout   *string `json:"ALayout,omitempty"`
	AProfile  *string `json:"AProfile,omitempty"`
	AR        *int32  `json:"AR,omitempty"`
	AbrMode   *int32  `json:"AbrMode,omitempty"`
	AccountID *string `json:"AccountID,omitempty"`

	// 音频编码格式，支持以下 3 种类型。
	// * aac：使用 AAC 编码格式；
	// * copy：不进行转码，所有音频编码参数不生效；
	// * opus：使用 Opus 编码格式。
	Acodec         *string `json:"Acodec,omitempty"`
	AdvancedParam  *string `json:"AdvancedParam,omitempty"`
	AllowAudioCopy *int32  `json:"AllowAudioCopy,omitempty"`
	AllowVideoCopy *int32  `json:"AllowVideoCopy,omitempty"`
	An             *int32  `json:"An,omitempty"`

	// 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App *string `json:"App,omitempty"`

	// 宽高自适应模式开关，默认值为 0。支持的取值及含义如下。
	// * 0：关闭宽高自适应
	// * 1：开启宽高自适应 :::tip
	// * 关闭宽高自适应时，转码配置分辨率取 Width 和 Height 的值对转码视频进行拉伸；
	// * 开启宽高自适应时，转码配置分辨率按照 ShortSide 、 LongSide 、Width 、Height 的优先级取值，另一边等比缩放。
	// * 修改 As 为 0 时，请确认 Width 和 Height 的取值是否超出阈值；
	// * 修改 As 为 1 时，请确认 ShortSide 和 LongSide 的取值是否超出阈值。 :::
	As *string `json:"As,omitempty"`

	// 音频码率，单位为 kbps。
	AudioBitrate *int32 `json:"AudioBitrate,omitempty"`
	AutoTransAb  *int32 `json:"AutoTransAb,omitempty"`
	AutoTransAl  *int32 `json:"AutoTransAl,omitempty"`
	AutoTransAr  *int32 `json:"AutoTransAr,omitempty"`

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
	BCM         *int32 `json:"BCM,omitempty"`

	// 2 个参考帧之间的最大 B 帧数，默认值为 3。配置不同的视频编码格式时，最大 B 帧数的取值存在如下差异。
	// * H.264：取值范围为 [0,7]；
	// * H.265：取值范围为 [0,3]、7、15；
	// * H.266：取值范围为 [0,3]、7、15。
	// 取值为 0 时，表示去除 B 帧。
	BFrames  *int32  `json:"BFrames,omitempty"`
	Describe *string `json:"Describe,omitempty"`

	// 视频帧率，单位为 fps，默认值为 25，帧率越大，画面越流畅。 配置不同视频编码格式时，视频帧率的取值存在如下差异。
	// * H.264：取值范围为 [0,60]；
	// * H.265：取值范围为 [0,60]；
	// * H.266：取值范围为 [0,35]。
	FPS *int32 `json:"FPS,omitempty"`

	// IDR 帧之间的最大间隔，单位为秒，默认值为 4，取值范围为 [1,20]。
	GOP    *int32 `json:"GOP,omitempty"`
	GopMin *int32 `json:"GopMin,omitempty"`
	HVSPre *bool  `json:"HVSPre,omitempty"`

	// 视频高度，默认值为 0。配置不同视频编码格式时，视频高度的取值存在如下差异。
	// * H.264：取值范围为 [150,1920]；
	// * H.265：取值范围为 [150,1920]。
	// :::tip
	// * 当 As 的取值为 0 时，参数生效；反之则不生效；
	// * 编码格式为 H.266 时，不支持设置 Width 和 Height，请使用自适应配置。 :::
	Height *int32 `json:"Height,omitempty"`

	// 长边长度，配置不同的视频编码方式和转码类型时，长边长度的取值范围存在如下差异。
	// * Roi 取 false 时： * H.264：取值范围为 0 和 [150,4096]；
	// * H.265：取值范围为 0 和 [150,7680]；
	// * H.266：取值范围为 0 和 [150,1280]。
	//
	//
	// * Roi 取 true 时： * H.264：取值范围为 0 和 [150,1920]；
	// * H.265：取值范围为 0 和 [150,1920]。 :::tip
	//
	//
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
	PresetKind   *int32  `json:"PresetKind,omitempty"`
	PresetType   *int32  `json:"PresetType,omitempty"`
	Qp           *int32  `json:"Qp,omitempty"`
	RegionConfig *string `json:"RegionConfig,omitempty"`
	Revision     *string `json:"Revision,omitempty"`

	// 是否极智超清转码，取值及含义如下。
	// * true：极智超清转码；
	// * false：标准转码。 :::tip
	// * 修改 Roi 为 true，且 As 为 1 时，请确认 ShortSide 和 LongSide 的取值是否超出阈值。
	// * 视频编码格式为 H.266 时，转码类型不支持极智超清转码。 :::
	Roi  *bool `json:"Roi,omitempty"`
	SITI *bool `json:"SITI,omitempty"`

	// 短边长度，配置不同的视频编码方式和转码类型时，短边长度的取值范围存在如下差异。
	// * Roi 取 false 时： * H.264：取值范围为 0 和 [150,2160]；
	// * H.265：取值范围为 0 和 [150,4096]；
	// * H.266：取值范围为 0 和 [150,720]。
	//
	//
	// * Roi 取 true 时： * H.264：取值范围为 0 和 [150,1920]；
	// * H.265：取值范围为 0 和 [150,1920]。 :::tip
	//
	//
	// * 当 As 的取值为 1 即开启宽高自适应时，参数生效，反之则不生效。
	// * 当 As 的取值为 1 时，如果 LongSide 、 ShortSide 、Width 、Height 同时取 0，表示保持源流尺寸。 :::
	ShortSide *int32 `json:"ShortSide,omitempty"`
	Status    *int32 `json:"Status,omitempty"`

	// 转码停止时长，支持触发方式为拉流转码时设置，表示断开拉流后转码停止的时长，单位为秒，取值范围为 -1 和 [0,300]，-1 表示不停止转码，默认值为 60。
	StopInterval *int32 `json:"StopInterval,omitempty"`

	// 转码流后缀名。支持 10 个字符以内的大小写字母、下划线与中划线，常见后缀包括：sd、hd、uhd。 例如，配置的转���流后缀名为hd，则拉转码流时转码的流名为 stream-123456789_hd。
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
	// * h264：使用 H.264 编码格式；
	// * h265：使用 H.265 编码格式；
	// * h266：使用 H.266 编码格式；
	// * copy：不进行转码，所有视频编码参数不生效。
	Vcodec *string `json:"Vcodec,omitempty"`

	// 视频码率，单位为 bps，默认值为 1000000；取 0 时，表示使用源流码率。 配置不同的视频编码格式时，视频码率的取值范围存在如下差异。
	// * H.264：取值范围为 [0,30000000]；
	// * H.265：取值范围为 [0,30000000]；
	// * H.266：取值范围为 [0,6000000]。
	VideoBitrate *int32  `json:"VideoBitrate,omitempty"`
	Vn           *int32  `json:"Vn,omitempty"`
	Watermark    *string `json:"Watermark,omitempty"`

	// 视频宽度，默认值为 0。配置不同视频编码格式时，视频宽度的取值存在如下差异。
	// * H.264：取值范围为 [150,1920]；
	// * H.265：取值范围为 [150,1920]。
	// :::tip
	// * 当 As 的取值为 0 即关闭宽高自适应时，转码分辨率将取 Width 和 Height 的值对转码视频进行拉伸；
	// * Width 和 Height 任一配置为 0 时，转码视频将保持源流尺寸；
	// * 编码格式为 H.266 时，不支持设置 Width 和 Height，请使用自适应配置。 :::
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

type UpdateVhostTagsBody struct {

	// REQUIRED; 域名空间
	Vhost string `json:"Vhost"`

	// 标签列表，不填就更新为空
	Tags []*UpdateVhostTagsBodyTagsItem `json:"Tags,omitempty"`
}

type UpdateVhostTagsBodyTagsItem struct {

	// REQUIRED; 标签的类型，System, Custom
	Category string `json:"Category"`

	// REQUIRED; 标签key
	Key string `json:"Key"`

	// REQUIRED; 标签value
	Value string `json:"Value"`
}

type UpdateVhostTagsRes struct {

	// REQUIRED
	ResponseMetadata UpdateVhostTagsResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateVhostTagsResResponseMetadata struct {

	// REQUIRED; 请求的接口名，属于请求的公共参数。
	Action string `json:"Action"`

	// REQUIRED; 请求的Region，例如：cn-north-1
	Region string `json:"Region"`

	// REQUIRED; RequestID为每次API请求的唯一标识。
	RequestID string `json:"RequestId"`

	// REQUIRED; 请求的服务，属于请求的公共参数。
	Service string `json:"Service"`

	// REQUIRED; 请求的版本号，属于请求的公共参数。
	Version string                                   `json:"Version"`
	Error   *UpdateVhostTagsResResponseMetadataError `json:"Error,omitempty"`
}

type UpdateVhostTagsResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type UpdateWatermarkPresetBody struct {

	// REQUIRED; 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `json:"App"`

	// REQUIRED; 域名空间名称，由 1 到 60 位数字、字母、下划线及"-"和"."组成。
	Vhost string `json:"Vhost"`

	// 直播画面方向，支持 2 种取值。
	// * vertical：竖屏；
	// * horizontal：横屏。 :::tip 该参数属于历史版本参数，预计将于未来移除。建议使用预览背景高度（PreviewHeight）、预览背景宽度（PreviewWidth）参数代替。 :::
	Orientation *string `json:"Orientation,omitempty"`

	// 水印图片链接，图片最大 2MB，最小 100Bytes，最大分辨率为 1080×1080。图片使用 data URI 协议，格式为：data:[<mediatype>];[base64],<data>。
	// * mediatype：图片类型，支持 png、jpg、jpeg 格式；
	// * data：base64 编码的图片字符串。
	// :::warning 如果水印图片不更新，请勿在更新配置时传入该参数，否则会造成水印无法显示。 :::
	Picture *string `json:"Picture,omitempty"`

	// 水印图片对应的 HTTP 地址。与水印图片字符串字段二选一传入。同时传入时，以水印图片字符串参数为准。 :::warning 如果水印图片不更新，请勿在更新配置时传入该参数，否则会造成水印无法显示。 :::
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

type UpdateWatermarkPresetV2Body struct {

	// 应用名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App *string `json:"App,omitempty"`

	// 需要添加水印的直播画面方向，支持 2 种取值。
	// * vertical：竖屏；
	// * horizontal：横屏。
	Orientation *string `json:"Orientation,omitempty"`

	// 水印图片字符串，图片最大 2MB，最小 100Bytes，最大分辨率为 1080×1080。图片 Data URL 格式为：data:[<mediatype>];[base64],<data>。
	// * mediatype：图片类型，支持 png、jpg、jpeg 格式；
	// * data：base64 编码的图片字符串。
	// 例如，data:image/png;base64,iVBORw0KGg****mCC
	Picture    *string `json:"Picture,omitempty"`
	PictureURL *string `json:"PictureUrl,omitempty"`

	// 水平偏移，表示水印左侧边与转码流画面左侧边之间的距离，使用相对比率，取值范围为 [0,1]。
	PosX *float32 `json:"PosX,omitempty"`

	// 垂直偏移，表示水印顶部边与转码流画面顶部边之间的距离，使用相对比率，取值范围为 [0,1]。
	PosY           *float32 `json:"PosY,omitempty"`
	PresetName     *string  `json:"PresetName,omitempty"`
	PreviewHeight  *float32 `json:"PreviewHeight,omitempty"`
	PreviewWidth   *float32 `json:"PreviewWidth,omitempty"`
	RelativeHeight *float32 `json:"RelativeHeight,omitempty"`

	// 水印相对宽度，水印宽度占直播转码流画面宽度的比例，取值范围为 [0,1]，水印高度会随宽度等比缩放。
	RelativeWidth *float32 `json:"RelativeWidth,omitempty"`
	Scale         *float32 `json:"Scale,omitempty"`
	Stream        *string  `json:"Stream,omitempty"`

	// 域名空间名称。由 1 到 60 位数字、字母、下划线及"-"和"."组成。
	Vhost *string `json:"Vhost,omitempty"`
}

type UpdateWatermarkPresetV2Res struct {

	// REQUIRED
	ResponseMetadata UpdateWatermarkPresetV2ResResponseMetadata `json:"ResponseMetadata"`

	// Anything
	Result interface{} `json:"Result,omitempty"`
}

type UpdateWatermarkPresetV2ResResponseMetadata struct {

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

type ValidateCertBody struct {

	// REQUIRED; 证书链ID
	ChainID string `json:"ChainID"`

	// 账户id
	AccountID *string `json:"AccountID,omitempty"`

	// 域名
	Domain *string `json:"Domain,omitempty"`

	// 证书信息
	Rsa *ValidateCertBodyRsa `json:"Rsa,omitempty"`
}

// ValidateCertBodyRsa - 证书信息
type ValidateCertBodyRsa struct {

	// 证书信息
	PubKey *string `json:"PubKey,omitempty"`
}

type ValidateCertRes struct {

	// REQUIRED
	ResponseMetadata ValidateCertResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *ValidateCertResResult `json:"Result,omitempty"`
}

type ValidateCertResResponseMetadata struct {

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
	Error   *ValidateCertResResponseMetadataError `json:"Error,omitempty"`
}

type ValidateCertResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

// ValidateCertResResult - 视请求的接口而定
type ValidateCertResResult struct {

	// 证书内容合法
	CertValid *bool `json:"CertValid,omitempty"`

	// 证书与域名是否匹配
	DomainValid *bool `json:"DomainValid,omitempty"`

	// 检查失败原因
	Reason *string `json:"Reason,omitempty"`
}

type VerifyDomainOwnerBody struct {

	// REQUIRED; 推拉流域名
	Domain string `json:"Domain"`

	// REQUIRED; 校验方式，取值：
	// * dnsCheck：DNS验证。
	// * fileCheck：文件验证。
	VerifyType string `json:"VerifyType"`
}

type VerifyDomainOwnerRes struct {

	// REQUIRED
	ResponseMetadata VerifyDomainOwnerResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *VerifyDomainOwnerResResult `json:"Result,omitempty"`
}

type VerifyDomainOwnerResResponseMetadata struct {

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

// VerifyDomainOwnerResResult - 视请求的接口而定
type VerifyDomainOwnerResResult struct {

	// REQUIRED; 检查结果
	CheckResult bool `json:"CheckResult"`
}
type AddCommonTransPreset struct{}
type AddCommonTransPresetQuery struct{}
type AssociatePreset struct{}
type AssociatePresetQuery struct{}
type AssociateRefConfig struct{}
type AssociateRefConfigQuery struct{}
type BindCert struct{}
type BindCertQuery struct{}
type BindEncryptDRM struct{}
type BindEncryptDRMQuery struct{}
type CheckCustomLogConfig struct{}
type CheckCustomLogConfigQuery struct{}
type CreateApp struct{}
type CreateAppQuery struct{}
type CreateAvSlicePreset struct{}
type CreateAvSlicePresetQuery struct{}
type CreateCert struct{}
type CreateCertQuery struct{}
type CreateCustomLogConfig struct{}
type CreateCustomLogConfigQuery struct{}
type CreateDenseSnapshotPreset struct{}
type CreateDenseSnapshotPresetQuery struct{}
type CreateDomain struct{}
type CreateDomainQuery struct{}
type CreateDomainV2 struct{}
type CreateDomainV2Query struct{}
type CreateLiveAccountFeeConfig struct{}
type CreateLiveAccountFeeConfigQuery struct{}
type CreateProxyConfig struct{}
type CreateProxyConfigQuery struct{}
type CreatePullCDNSnapshotTask struct{}
type CreatePullCDNSnapshotTaskQuery struct{}
type CreatePullRecordTask struct{}
type CreatePullRecordTaskQuery struct{}
type CreatePullToPushTask struct{}
type CreatePullToPushTaskQuery struct{}
type CreateRecordPresetV2 struct{}
type CreateRecordPresetV2Query struct{}
type CreateRelaySourceV4 struct{}
type CreateRelaySourceV4Query struct{}
type CreateSDK struct{}
type CreateSDKQuery struct{}
type CreateSnapshotAuditPreset struct{}
type CreateSnapshotAuditPresetQuery struct{}
type CreateSnapshotPreset struct{}
type CreateSnapshotPresetQuery struct{}
type CreateSnapshotPresetV2 struct{}
type CreateSnapshotPresetV2Query struct{}
type CreateTicket struct{}
type CreateTicketQuery struct{}
type CreateTimeShiftPresetV2 struct{}
type CreateTimeShiftPresetV2Query struct{}
type CreateTimeShiftPresetV3 struct{}
type CreateTimeShiftPresetV3Query struct{}
type CreateTranscodePreset struct{}
type CreateTranscodePresetBatch struct{}
type CreateTranscodePresetBatchQuery struct{}
type CreateTranscodePresetPatchByAdmin struct{}
type CreateTranscodePresetPatchByAdminQuery struct{}
type CreateTranscodePresetQuery struct{}
type CreateVQScoreTask struct{}
type CreateVQScoreTaskQuery struct{}
type CreateVerifyContent struct{}
type CreateVerifyContentQuery struct{}
type CreateWatermarkPreset struct{}
type CreateWatermarkPresetQuery struct{}
type CreateWatermarkPresetV2 struct{}
type CreateWatermarkPresetV2Query struct{}
type DeleteAuth struct{}
type DeleteAuthQuery struct{}
type DeleteAvSlicePreset struct{}
type DeleteAvSlicePresetQuery struct{}
type DeleteCMAFConfig struct{}
type DeleteCMAFConfigQuery struct{}
type DeleteCallback struct{}
type DeleteCallbackQuery struct{}
type DeleteCert struct{}
type DeleteCertQuery struct{}
type DeleteCommonTransPreset struct{}
type DeleteCommonTransPresetQuery struct{}
type DeleteCustomLogConfig struct{}
type DeleteCustomLogConfigQuery struct{}
type DeleteDenseSnapshotPreset struct{}
type DeleteDenseSnapshotPresetQuery struct{}
type DeleteDenyConfigV2 struct{}
type DeleteDenyConfigV2Query struct{}
type DeleteDomain struct{}
type DeleteDomainQuery struct{}
type DeleteDomainV2 struct{}
type DeleteDomainV2Query struct{}
type DeleteHLSConfig struct{}
type DeleteHLSConfigQuery struct{}
type DeleteHTTPHeaderConfig struct{}
type DeleteHTTPHeaderConfigQuery struct{}
type DeleteHeaderConfig struct{}
type DeleteHeaderConfigQuery struct{}
type DeleteIPAccessRule struct{}
type DeleteIPAccessRuleQuery struct{}
type DeleteLatencyConfig struct{}
type DeleteLatencyConfigQuery struct{}
type DeleteLiveAccountFeeConfig struct{}
type DeleteLiveAccountFeeConfigQuery struct{}
type DeleteNSSRewriteConfig struct{}
type DeleteNSSRewriteConfigQuery struct{}
type DeleteProxyConfig struct{}
type DeleteProxyConfigAssociation struct{}
type DeleteProxyConfigAssociationQuery struct{}
type DeleteProxyConfigQuery struct{}
type DeletePullToPushTask struct{}
type DeletePullToPushTaskQuery struct{}
type DeleteRecordHistory struct{}
type DeleteRecordHistoryQuery struct{}
type DeleteRecordPreset struct{}
type DeleteRecordPresetQuery struct{}
type DeleteReferer struct{}
type DeleteRefererQuery struct{}
type DeleteRelaySink struct{}
type DeleteRelaySinkQuery struct{}
type DeleteRelaySourceRewrite struct{}
type DeleteRelaySourceRewriteQuery struct{}
type DeleteRelaySourceV3 struct{}
type DeleteRelaySourceV3Query struct{}
type DeleteRelaySourceV4 struct{}
type DeleteRelaySourceV4Query struct{}
type DeleteSDK struct{}
type DeleteSDKQuery struct{}
type DeleteSnapshotAuditPreset struct{}
type DeleteSnapshotAuditPresetQuery struct{}
type DeleteSnapshotPreset struct{}
type DeleteSnapshotPresetQuery struct{}
type DeleteStreamQuotaConfig struct{}
type DeleteStreamQuotaConfigQuery struct{}
type DeleteTimeShiftPresetV2 struct{}
type DeleteTimeShiftPresetV2Query struct{}
type DeleteTimeShiftPresetV3 struct{}
type DeleteTimeShiftPresetV3Query struct{}
type DeleteTranscodePreset struct{}
type DeleteTranscodePresetBatch struct{}
type DeleteTranscodePresetBatchQuery struct{}
type DeleteTranscodePresetPatchByAdmin struct{}
type DeleteTranscodePresetPatchByAdminQuery struct{}
type DeleteTranscodePresetQuery struct{}
type DeleteWatermarkPreset struct{}
type DeleteWatermarkPresetQuery struct{}
type DeleteWatermarkPresetV2 struct{}
type DeleteWatermarkPresetV2Query struct{}
type DescDenseSnapshotPresetDetail struct{}
type DescDenseSnapshotPresetDetailQuery struct{}
type DescribeActionHistory struct{}
type DescribeActionHistoryQuery struct{}
type DescribeAppIDParamsAvailable struct{}
type DescribeAppIDParamsAvailableQuery struct{}
type DescribeAuth struct{}
type DescribeAuthQuery struct{}
type DescribeBilling struct{}
type DescribeBillingBody struct{}
type DescribeBillingForAdmin struct{}
type DescribeBillingForAdminQuery struct{}
type DescribeBillingMonthAvailable struct{}
type DescribeBillingMonthAvailableBody struct{}
type DescribeBillingMonthAvailableQuery struct{}
type DescribeBillingQuery struct{}
type DescribeCDNSnapshotHistory struct{}
type DescribeCDNSnapshotHistoryQuery struct{}
type DescribeCMAFConfig struct{}
type DescribeCMAFConfigQuery struct{}
type DescribeCallback struct{}
type DescribeCallbackQuery struct{}
type DescribeCertDRM struct{}
type DescribeCertDRMBody struct{}
type DescribeCertDetailSecret struct{}
type DescribeCertDetailSecretQuery struct{}
type DescribeCertDetailSecretV2 struct{}
type DescribeCertDetailSecretV2Query struct{}
type DescribeCertDetailV2 struct{}
type DescribeCertDetailV2Query struct{}
type DescribeClosedStreamInfoByPage struct{}
type DescribeClosedStreamInfoByPageBody struct{}
type DescribeContentKey struct{}
type DescribeContentKeyBody struct{}
type DescribeContentKeyQuery struct{}
type DescribeCustomLogConfig struct{}
type DescribeCustomLogConfigBody struct{}
type DescribeCustomLogConfigQuery struct{}
type DescribeDenyConfig struct{}
type DescribeDenyConfigQuery struct{}
type DescribeDenyConfigV2 struct{}
type DescribeDenyConfigV2Query struct{}
type DescribeDomain struct{}
type DescribeDomainQuery struct{}
type DescribeDomainVerify struct{}
type DescribeDomainVerifyQuery struct{}
type DescribeEncryptDRM struct{}
type DescribeEncryptDRMBody struct{}
type DescribeEncryptDRMQuery struct{}
type DescribeForbiddenStreamInfoByPage struct{}
type DescribeForbiddenStreamInfoByPageBody struct{}
type DescribeHLSConfig struct{}
type DescribeHLSConfigQuery struct{}
type DescribeHTTPHeaderConfig struct{}
type DescribeHTTPHeaderConfigQuery struct{}
type DescribeHeaderConfig struct{}
type DescribeHeaderConfigQuery struct{}
type DescribeIPAccessRule struct{}
type DescribeIPAccessRuleQuery struct{}
type DescribeIPInfo struct{}
type DescribeIPInfoQuery struct{}
type DescribeLatencyConfig struct{}
type DescribeLatencyConfigQuery struct{}
type DescribeLicenseDRM struct{}
type DescribeLicenseDRMBody struct{}
type DescribeLiveAccountFeeConfig struct{}
type DescribeLiveAccountFeeConfigBody struct{}
type DescribeLiveAccountFeeConfigQuery struct{}
type DescribeLiveAccountFeeType struct{}
type DescribeLiveAccountFeeTypeQuery struct{}
type DescribeLiveActivityBandwidthData struct{}
type DescribeLiveActivityBandwidthDataQuery struct{}
type DescribeLiveAuditData struct{}
type DescribeLiveAuditDataQuery struct{}
type DescribeLiveBandwidthData struct{}
type DescribeLiveBandwidthDataQuery struct{}
type DescribeLiveBatchOnlineStreamMetrics struct{}
type DescribeLiveBatchOnlineStreamMetricsQuery struct{}
type DescribeLiveBatchPushStreamAvgMetrics struct{}
type DescribeLiveBatchPushStreamAvgMetricsQuery struct{}
type DescribeLiveBatchPushStreamMetrics struct{}
type DescribeLiveBatchPushStreamMetricsQuery struct{}
type DescribeLiveBatchSourceStreamAvgMetrics struct{}
type DescribeLiveBatchSourceStreamAvgMetricsQuery struct{}
type DescribeLiveBatchSourceStreamMetrics struct{}
type DescribeLiveBatchSourceStreamMetricsQuery struct{}
type DescribeLiveBatchStreamTrafficData struct{}
type DescribeLiveBatchStreamTrafficDataQuery struct{}
type DescribeLiveBatchStreamTranscodeData struct{}
type DescribeLiveBatchStreamTranscodeDataQuery struct{}
type DescribeLiveCustomizedLogData struct{}
type DescribeLiveCustomizedLogDataQuery struct{}
type DescribeLiveFeeConfig struct{}
type DescribeLiveFeeConfigBody struct{}
type DescribeLiveFeeConfigQuery struct{}
type DescribeLiveFreeTimeInterval struct{}
type DescribeLiveFreeTimeIntervalBody struct{}
type DescribeLiveFreeTimeIntervalQuery struct{}
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
type DescribeLivePushStreamCountData struct{}
type DescribeLivePushStreamCountDataQuery struct{}
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
type DescribeLiveStreamCountData struct{}
type DescribeLiveStreamCountDataQuery struct{}
type DescribeLiveStreamInfoByPage struct{}
type DescribeLiveStreamInfoByPageBody struct{}
type DescribeLiveStreamSessionData struct{}
type DescribeLiveStreamSessionDataQuery struct{}
type DescribeLiveStreamState struct{}
type DescribeLiveStreamStateBody struct{}
type DescribeLiveStreamUsageData struct{}
type DescribeLiveStreamUsageDataQuery struct{}
type DescribeLiveTimeShiftData struct{}
type DescribeLiveTimeShiftDataQuery struct{}
type DescribeLiveTrafficData struct{}
type DescribeLiveTrafficDataQuery struct{}
type DescribeLiveTranscodeData struct{}
type DescribeLiveTranscodeDataQuery struct{}
type DescribeNSSRewriteConfig struct{}
type DescribeNSSRewriteConfigQuery struct{}
type DescribePresetAssociation struct{}
type DescribePresetAssociationQuery struct{}
type DescribeProxyConfigAssociation struct{}
type DescribeProxyConfigAssociationQuery struct{}
type DescribeRecordTaskFileHistory struct{}
type DescribeRecordTaskFileHistoryQuery struct{}
type DescribeRefConfig struct{}
type DescribeRefConfigQuery struct{}
type DescribeReferer struct{}
type DescribeRefererQuery struct{}
type DescribeRelaySink struct{}
type DescribeRelaySinkQuery struct{}
type DescribeRelaySourceRewrite struct{}
type DescribeRelaySourceRewriteQuery struct{}
type DescribeRelaySourceV3 struct{}
type DescribeRelaySourceV3Query struct{}
type DescribeSDKDetail struct{}
type DescribeSDKDetailQuery struct{}
type DescribeSDKParamsAvailable struct{}
type DescribeSDKParamsAvailableQuery struct{}
type DescribeService struct{}
type DescribeServiceBody struct{}
type DescribeServiceQuery struct{}
type DescribeSnapshotAuditPresetDetail struct{}
type DescribeSnapshotAuditPresetDetailQuery struct{}
type DescribeStreamQuotaConfig struct{}
type DescribeStreamQuotaConfigQuery struct{}
type DescribeTimeShiftPresetDetail struct{}
type DescribeTimeShiftPresetDetailQuery struct{}
type DescribeTranscodePresetDetail struct{}
type DescribeTranscodePresetDetailQuery struct{}
type DescribeVQScoreTask struct{}
type DescribeVQScoreTaskQuery struct{}
type DescribeVhost struct{}
type DescribeVhostQuery struct{}
type DescribeWatermarkPresetDetail struct{}
type DescribeWatermarkPresetDetailQuery struct{}
type DisAssociatePreset struct{}
type DisAssociatePresetQuery struct{}
type DisableAuth struct{}
type DisableAuthQuery struct{}
type DisableDomain struct{}
type DisableDomainQuery struct{}
type DisassociateRefConfig struct{}
type DisassociateRefConfigQuery struct{}
type EnableAuth struct{}
type EnableAuthQuery struct{}
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
type GenerateTimeShiftPlayURL struct{}
type GenerateTimeShiftPlayURLQuery struct{}
type GetApps struct{}
type GetAppsBody struct{}
type GetAppsQuery struct{}
type GetPullCDNSnapshotTask struct{}
type GetPullCDNSnapshotTaskQuery struct{}
type GetPullRecordTask struct{}
type GetPullRecordTaskQuery struct{}
type GetTags struct{}
type GetTagsBody struct{}
type GetTagsQuery struct{}
type GetVqosRawData struct{}
type KillStream struct{}
type KillStreamQuery struct{}
type ListActionHistory struct{}
type ListActionHistoryQuery struct{}
type ListBindEncryptDRM struct{}
type ListBindEncryptDRMQuery struct{}
type ListCert struct{}
type ListCertBindInfo struct{}
type ListCertBindInfoQuery struct{}
type ListCertQuery struct{}
type ListCertV2 struct{}
type ListCertV2Query struct{}
type ListCommonTransPresetDetail struct{}
type ListCommonTransPresetDetailQuery struct{}
type ListDomainDetail struct{}
type ListDomainDetailQuery struct{}
type ListHeaderEnum struct{}
type ListHeaderEnumQuery struct{}
type ListInstance struct{}
type ListInstanceQuery struct{}
type ListObject struct{}
type ListObjectQuery struct{}
type ListProjects struct{}
type ListProjectsQuery struct{}
type ListProxyConfig struct{}
type ListProxyConfigQuery struct{}
type ListPullCDNSnapshotTask struct{}
type ListPullCDNSnapshotTaskQuery struct{}
type ListPullRecordTask struct{}
type ListPullRecordTaskQuery struct{}
type ListPullToPushTask struct{}
type ListPullToPushTaskBody struct{}
type ListReferenceInfo struct{}
type ListReferenceInfoQuery struct{}
type ListReferenceNames struct{}
type ListReferenceNamesQuery struct{}
type ListReferenceTypes struct{}
type ListReferenceTypesBody struct{}
type ListReferenceTypesQuery struct{}
type ListRelaySourceV4 struct{}
type ListRelaySourceV4Query struct{}
type ListResourcePackage struct{}
type ListResourcePackageQuery struct{}
type ListSDK struct{}
type ListSDKAdmin struct{}
type ListSDKAdminQuery struct{}
type ListSDKQuery struct{}
type ListServices struct{}
type ListServicesQuery struct{}
type ListTimeShiftPresetV2 struct{}
type ListTimeShiftPresetV2Query struct{}
type ListVQScoreTask struct{}
type ListVQScoreTaskQuery struct{}
type ListVhostDenseSnapshotPreset struct{}
type ListVhostDenseSnapshotPresetQuery struct{}
type ListVhostDetail struct{}
type ListVhostDetailByAdmin struct{}
type ListVhostDetailByAdminQuery struct{}
type ListVhostDetailQuery struct{}
type ListVhostDomainDetailByUserID struct{}
type ListVhostDomainDetailByUserIDQuery struct{}
type ListVhostRecordPresetV2 struct{}
type ListVhostRecordPresetV2Query struct{}
type ListVhostSnapshotAuditPreset struct{}
type ListVhostSnapshotAuditPresetQuery struct{}
type ListVhostSnapshotPreset struct{}
type ListVhostSnapshotPresetQuery struct{}
type ListVhostSnapshotPresetV2 struct{}
type ListVhostSnapshotPresetV2Query struct{}
type ListVhostTransCodePreset struct{}
type ListVhostTransCodePresetQuery struct{}
type ListVhostWatermarkPreset struct{}
type ListVhostWatermarkPresetQuery struct{}
type ListVideoClassifications struct{}
type ListVideoClassificationsQuery struct{}
type ListVqosDimensionValues struct{}
type ListVqosMetricsDimensions struct{}
type ListVqosMetricsDimensionsBody struct{}
type ListWatermarkPreset struct{}
type ListWatermarkPresetQuery struct{}
type ManagerPullPushDomainBind struct{}
type ManagerPullPushDomainBindQuery struct{}
type RejectDomain struct{}
type RejectDomainQuery struct{}
type RestartPullToPushTask struct{}
type RestartPullToPushTaskQuery struct{}
type ResumeStream struct{}
type ResumeStreamQuery struct{}
type StopPullCDNSnapshotTask struct{}
type StopPullCDNSnapshotTaskQuery struct{}
type StopPullRecordTask struct{}
type StopPullRecordTaskQuery struct{}
type StopPullToPushTask struct{}
type StopPullToPushTaskQuery struct{}
type TerminateInstance struct{}
type TerminateInstanceQuery struct{}
type UnBindEncryptDRM struct{}
type UnBindEncryptDRMQuery struct{}
type UnbindCert struct{}
type UnbindCertQuery struct{}
type UpdateActivityBilling struct{}
type UpdateActivityBillingQuery struct{}
type UpdateApp struct{}
type UpdateAppQuery struct{}
type UpdateAuthKey struct{}
type UpdateAuthKeyQuery struct{}
type UpdateAvSlicePreset struct{}
type UpdateAvSlicePresetQuery struct{}
type UpdateBilling struct{}
type UpdateBillingQuery struct{}
type UpdateCMAFConfig struct{}
type UpdateCMAFConfigQuery struct{}
type UpdateCallback struct{}
type UpdateCallbackQuery struct{}
type UpdateCert struct{}
type UpdateCertQuery struct{}
type UpdateDenseSnapshotPreset struct{}
type UpdateDenseSnapshotPresetQuery struct{}
type UpdateDenyConfig struct{}
type UpdateDenyConfigQuery struct{}
type UpdateDenyConfigV2 struct{}
type UpdateDenyConfigV2Query struct{}
type UpdateDomain struct{}
type UpdateDomainQuery struct{}
type UpdateDomainVhost struct{}
type UpdateDomainVhostQuery struct{}
type UpdateEncryptDRM struct{}
type UpdateEncryptDRMQuery struct{}
type UpdateHLSConfig struct{}
type UpdateHLSConfigQuery struct{}
type UpdateHTTPHeaderConfig struct{}
type UpdateHTTPHeaderConfigQuery struct{}
type UpdateHeaderConfig struct{}
type UpdateHeaderConfigQuery struct{}
type UpdateIPAccessRule struct{}
type UpdateIPAccessRuleQuery struct{}
type UpdateLatencyConfig struct{}
type UpdateLatencyConfigQuery struct{}
type UpdateNSSRewriteConfig struct{}
type UpdateNSSRewriteConfigQuery struct{}
type UpdatePresetAssociation struct{}
type UpdatePresetAssociationQuery struct{}
type UpdateProxyConfig struct{}
type UpdateProxyConfigAssociation struct{}
type UpdateProxyConfigAssociationQuery struct{}
type UpdateProxyConfigQuery struct{}
type UpdatePullToPushTask struct{}
type UpdatePullToPushTaskQuery struct{}
type UpdateRecordPresetV2 struct{}
type UpdateRecordPresetV2Query struct{}
type UpdateReferer struct{}
type UpdateRefererQuery struct{}
type UpdateRelaySink struct{}
type UpdateRelaySinkQuery struct{}
type UpdateRelaySourceRewrite struct{}
type UpdateRelaySourceRewriteQuery struct{}
type UpdateRelaySourceV3 struct{}
type UpdateRelaySourceV3Query struct{}
type UpdateRelaySourceV4 struct{}
type UpdateRelaySourceV4Query struct{}
type UpdateSDK struct{}
type UpdateSDKLicense struct{}
type UpdateSDKLicenseQuery struct{}
type UpdateSDKQuery struct{}
type UpdateService struct{}
type UpdateServiceQuery struct{}
type UpdateSnapshotAuditPreset struct{}
type UpdateSnapshotAuditPresetQuery struct{}
type UpdateSnapshotPreset struct{}
type UpdateSnapshotPresetQuery struct{}
type UpdateSnapshotPresetV2 struct{}
type UpdateSnapshotPresetV2Query struct{}
type UpdateStreamQuotaConfig struct{}
type UpdateStreamQuotaConfigPatch struct{}
type UpdateStreamQuotaConfigPatchQuery struct{}
type UpdateStreamQuotaConfigQuery struct{}
type UpdateTimeShiftPresetV2 struct{}
type UpdateTimeShiftPresetV2Query struct{}
type UpdateTimeShiftPresetV3 struct{}
type UpdateTimeShiftPresetV3Query struct{}
type UpdateTranscodePreset struct{}
type UpdateTranscodePresetQuery struct{}
type UpdateVhostTags struct{}
type UpdateVhostTagsQuery struct{}
type UpdateWatermarkPreset struct{}
type UpdateWatermarkPresetQuery struct{}
type UpdateWatermarkPresetV2 struct{}
type UpdateWatermarkPresetV2Query struct{}
type ValidateCert struct{}
type ValidateCertQuery struct{}
type VerifyDomainOwner struct{}
type VerifyDomainOwnerQuery struct{}
type AddCommonTransPresetReq struct {
	*AddCommonTransPresetQuery
	*AddCommonTransPresetBody
}
type AssociatePresetReq struct {
	*AssociatePresetQuery
	*AssociatePresetBody
}
type AssociateRefConfigReq struct {
	*AssociateRefConfigQuery
	*AssociateRefConfigBody
}
type BindCertReq struct {
	*BindCertQuery
	*BindCertBody
}
type BindEncryptDRMReq struct {
	*BindEncryptDRMQuery
	*BindEncryptDRMBody
}
type CheckCustomLogConfigReq struct {
	*CheckCustomLogConfigQuery
	*CheckCustomLogConfigBody
}
type CreateAppReq struct {
	*CreateAppQuery
	*CreateAppBody
}
type CreateAvSlicePresetReq struct {
	*CreateAvSlicePresetQuery
	*CreateAvSlicePresetBody
}
type CreateCertReq struct {
	*CreateCertQuery
	*CreateCertBody
}
type CreateCustomLogConfigReq struct {
	*CreateCustomLogConfigQuery
	*CreateCustomLogConfigBody
}
type CreateDenseSnapshotPresetReq struct {
	*CreateDenseSnapshotPresetQuery
	*CreateDenseSnapshotPresetBody
}
type CreateDomainReq struct {
	*CreateDomainQuery
	*CreateDomainBody
}
type CreateDomainV2Req struct {
	*CreateDomainV2Query
	*CreateDomainV2Body
}
type CreateLiveAccountFeeConfigReq struct {
	*CreateLiveAccountFeeConfigQuery
	*CreateLiveAccountFeeConfigBody
}
type CreateProxyConfigReq struct {
	*CreateProxyConfigQuery
	*CreateProxyConfigBody
}
type CreatePullCDNSnapshotTaskReq struct {
	*CreatePullCDNSnapshotTaskQuery
	*CreatePullCDNSnapshotTaskBody
}
type CreatePullRecordTaskReq struct {
	*CreatePullRecordTaskQuery
	*CreatePullRecordTaskBody
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
type CreateSDKReq struct {
	*CreateSDKQuery
	*CreateSDKBody
}
type CreateSnapshotAuditPresetReq struct {
	*CreateSnapshotAuditPresetQuery
	*CreateSnapshotAuditPresetBody
}
type CreateSnapshotPresetReq struct {
	*CreateSnapshotPresetQuery
	*CreateSnapshotPresetBody
}
type CreateSnapshotPresetV2Req struct {
	*CreateSnapshotPresetV2Query
	*CreateSnapshotPresetV2Body
}
type CreateTicketReq struct {
	*CreateTicketQuery
	*CreateTicketBody
}
type CreateTimeShiftPresetV2Req struct {
	*CreateTimeShiftPresetV2Query
	*CreateTimeShiftPresetV2Body
}
type CreateTimeShiftPresetV3Req struct {
	*CreateTimeShiftPresetV3Query
	*CreateTimeShiftPresetV3Body
}
type CreateTranscodePresetReq struct {
	*CreateTranscodePresetQuery
	*CreateTranscodePresetBody
}
type CreateTranscodePresetBatchReq struct {
	*CreateTranscodePresetBatchQuery
	*CreateTranscodePresetBatchBody
}
type CreateTranscodePresetPatchByAdminReq struct {
	*CreateTranscodePresetPatchByAdminQuery
	*CreateTranscodePresetPatchByAdminBody
}
type CreateVQScoreTaskReq struct {
	*CreateVQScoreTaskQuery
	*CreateVQScoreTaskBody
}
type CreateVerifyContentReq struct {
	*CreateVerifyContentQuery
	*CreateVerifyContentBody
}
type CreateWatermarkPresetReq struct {
	*CreateWatermarkPresetQuery
	*CreateWatermarkPresetBody
}
type CreateWatermarkPresetV2Req struct {
	*CreateWatermarkPresetV2Query
	*CreateWatermarkPresetV2Body
}
type DeleteAuthReq struct {
	*DeleteAuthQuery
	*DeleteAuthBody
}
type DeleteAvSlicePresetReq struct {
	*DeleteAvSlicePresetQuery
	*DeleteAvSlicePresetBody
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
type DeleteCommonTransPresetReq struct {
	*DeleteCommonTransPresetQuery
	*DeleteCommonTransPresetBody
}
type DeleteCustomLogConfigReq struct {
	*DeleteCustomLogConfigQuery
	*DeleteCustomLogConfigBody
}
type DeleteDenseSnapshotPresetReq struct {
	*DeleteDenseSnapshotPresetQuery
	*DeleteDenseSnapshotPresetBody
}
type DeleteDenyConfigV2Req struct {
	*DeleteDenyConfigV2Query
	*DeleteDenyConfigV2Body
}
type DeleteDomainReq struct {
	*DeleteDomainQuery
	*DeleteDomainBody
}
type DeleteDomainV2Req struct {
	*DeleteDomainV2Query
	*DeleteDomainV2Body
}
type DeleteHLSConfigReq struct {
	*DeleteHLSConfigQuery
	*DeleteHLSConfigBody
}
type DeleteHTTPHeaderConfigReq struct {
	*DeleteHTTPHeaderConfigQuery
	*DeleteHTTPHeaderConfigBody
}
type DeleteHeaderConfigReq struct {
	*DeleteHeaderConfigQuery
	*DeleteHeaderConfigBody
}
type DeleteIPAccessRuleReq struct {
	*DeleteIPAccessRuleQuery
	*DeleteIPAccessRuleBody
}
type DeleteLatencyConfigReq struct {
	*DeleteLatencyConfigQuery
	*DeleteLatencyConfigBody
}
type DeleteLiveAccountFeeConfigReq struct {
	*DeleteLiveAccountFeeConfigQuery
	*DeleteLiveAccountFeeConfigBody
}
type DeleteNSSRewriteConfigReq struct {
	*DeleteNSSRewriteConfigQuery
	*DeleteNSSRewriteConfigBody
}
type DeleteProxyConfigReq struct {
	*DeleteProxyConfigQuery
	*DeleteProxyConfigBody
}
type DeleteProxyConfigAssociationReq struct {
	*DeleteProxyConfigAssociationQuery
	*DeleteProxyConfigAssociationBody
}
type DeletePullToPushTaskReq struct {
	*DeletePullToPushTaskQuery
	*DeletePullToPushTaskBody
}
type DeleteRecordHistoryReq struct {
	*DeleteRecordHistoryQuery
	*DeleteRecordHistoryBody
}
type DeleteRecordPresetReq struct {
	*DeleteRecordPresetQuery
	*DeleteRecordPresetBody
}
type DeleteRefererReq struct {
	*DeleteRefererQuery
	*DeleteRefererBody
}
type DeleteRelaySinkReq struct {
	*DeleteRelaySinkQuery
	*DeleteRelaySinkBody
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
type DeleteSDKReq struct {
	*DeleteSDKQuery
	*DeleteSDKBody
}
type DeleteSnapshotAuditPresetReq struct {
	*DeleteSnapshotAuditPresetQuery
	*DeleteSnapshotAuditPresetBody
}
type DeleteSnapshotPresetReq struct {
	*DeleteSnapshotPresetQuery
	*DeleteSnapshotPresetBody
}
type DeleteStreamQuotaConfigReq struct {
	*DeleteStreamQuotaConfigQuery
	*DeleteStreamQuotaConfigBody
}
type DeleteTimeShiftPresetV2Req struct {
	*DeleteTimeShiftPresetV2Query
	*DeleteTimeShiftPresetV2Body
}
type DeleteTimeShiftPresetV3Req struct {
	*DeleteTimeShiftPresetV3Query
	*DeleteTimeShiftPresetV3Body
}
type DeleteTranscodePresetReq struct {
	*DeleteTranscodePresetQuery
	*DeleteTranscodePresetBody
}
type DeleteTranscodePresetBatchReq struct {
	*DeleteTranscodePresetBatchQuery
	*DeleteTranscodePresetBatchBody
}
type DeleteTranscodePresetPatchByAdminReq struct {
	*DeleteTranscodePresetPatchByAdminQuery
	*DeleteTranscodePresetPatchByAdminBody
}
type DeleteWatermarkPresetReq struct {
	*DeleteWatermarkPresetQuery
	*DeleteWatermarkPresetBody
}
type DeleteWatermarkPresetV2Req struct {
	*DeleteWatermarkPresetV2Query
	*DeleteWatermarkPresetV2Body
}
type DescDenseSnapshotPresetDetailReq struct {
	*DescDenseSnapshotPresetDetailQuery
	*DescDenseSnapshotPresetDetailBody
}
type DescribeActionHistoryReq struct {
	*DescribeActionHistoryQuery
	*DescribeActionHistoryBody
}
type DescribeAppIDParamsAvailableReq struct {
	*DescribeAppIDParamsAvailableQuery
	*DescribeAppIDParamsAvailableBody
}
type DescribeAuthReq struct {
	*DescribeAuthQuery
	*DescribeAuthBody
}
type DescribeBillingReq struct {
	*DescribeBillingQuery
	*DescribeBillingBody
}
type DescribeBillingForAdminReq struct {
	*DescribeBillingForAdminQuery
	*DescribeBillingForAdminBody
}
type DescribeBillingMonthAvailableReq struct {
	*DescribeBillingMonthAvailableQuery
	*DescribeBillingMonthAvailableBody
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
type DescribeCertDetailSecretReq struct {
	*DescribeCertDetailSecretQuery
	*DescribeCertDetailSecretBody
}
type DescribeCertDetailSecretV2Req struct {
	*DescribeCertDetailSecretV2Query
	*DescribeCertDetailSecretV2Body
}
type DescribeCertDetailV2Req struct {
	*DescribeCertDetailV2Query
	*DescribeCertDetailV2Body
}
type DescribeClosedStreamInfoByPageReq struct {
	*DescribeClosedStreamInfoByPageQuery
	*DescribeClosedStreamInfoByPageBody
}
type DescribeContentKeyReq struct {
	*DescribeContentKeyQuery
	*DescribeContentKeyBody
}
type DescribeCustomLogConfigReq struct {
	*DescribeCustomLogConfigQuery
	*DescribeCustomLogConfigBody
}
type DescribeDenyConfigReq struct {
	*DescribeDenyConfigQuery
	*DescribeDenyConfigBody
}
type DescribeDenyConfigV2Req struct {
	*DescribeDenyConfigV2Query
	*DescribeDenyConfigV2Body
}
type DescribeDomainReq struct {
	*DescribeDomainQuery
	*DescribeDomainBody
}
type DescribeDomainVerifyReq struct {
	*DescribeDomainVerifyQuery
	*DescribeDomainVerifyBody
}
type DescribeEncryptDRMReq struct {
	*DescribeEncryptDRMQuery
	*DescribeEncryptDRMBody
}
type DescribeForbiddenStreamInfoByPageReq struct {
	*DescribeForbiddenStreamInfoByPageQuery
	*DescribeForbiddenStreamInfoByPageBody
}
type DescribeHLSConfigReq struct {
	*DescribeHLSConfigQuery
	*DescribeHLSConfigBody
}
type DescribeHTTPHeaderConfigReq struct {
	*DescribeHTTPHeaderConfigQuery
	*DescribeHTTPHeaderConfigBody
}
type DescribeHeaderConfigReq struct {
	*DescribeHeaderConfigQuery
	*DescribeHeaderConfigBody
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
type DescribeLiveAccountFeeConfigReq struct {
	*DescribeLiveAccountFeeConfigQuery
	*DescribeLiveAccountFeeConfigBody
}
type DescribeLiveAccountFeeTypeReq struct {
	*DescribeLiveAccountFeeTypeQuery
	*DescribeLiveAccountFeeTypeBody
}
type DescribeLiveActivityBandwidthDataReq struct {
	*DescribeLiveActivityBandwidthDataQuery
	*DescribeLiveActivityBandwidthDataBody
}
type DescribeLiveAuditDataReq struct {
	*DescribeLiveAuditDataQuery
	*DescribeLiveAuditDataBody
}
type DescribeLiveBandwidthDataReq struct {
	*DescribeLiveBandwidthDataQuery
	*DescribeLiveBandwidthDataBody
}
type DescribeLiveBatchOnlineStreamMetricsReq struct {
	*DescribeLiveBatchOnlineStreamMetricsQuery
	*DescribeLiveBatchOnlineStreamMetricsBody
}
type DescribeLiveBatchPushStreamAvgMetricsReq struct {
	*DescribeLiveBatchPushStreamAvgMetricsQuery
	*DescribeLiveBatchPushStreamAvgMetricsBody
}
type DescribeLiveBatchPushStreamMetricsReq struct {
	*DescribeLiveBatchPushStreamMetricsQuery
	*DescribeLiveBatchPushStreamMetricsBody
}
type DescribeLiveBatchSourceStreamAvgMetricsReq struct {
	*DescribeLiveBatchSourceStreamAvgMetricsQuery
	*DescribeLiveBatchSourceStreamAvgMetricsBody
}
type DescribeLiveBatchSourceStreamMetricsReq struct {
	*DescribeLiveBatchSourceStreamMetricsQuery
	*DescribeLiveBatchSourceStreamMetricsBody
}
type DescribeLiveBatchStreamTrafficDataReq struct {
	*DescribeLiveBatchStreamTrafficDataQuery
	*DescribeLiveBatchStreamTrafficDataBody
}
type DescribeLiveBatchStreamTranscodeDataReq struct {
	*DescribeLiveBatchStreamTranscodeDataQuery
	*DescribeLiveBatchStreamTranscodeDataBody
}
type DescribeLiveCustomizedLogDataReq struct {
	*DescribeLiveCustomizedLogDataQuery
	*DescribeLiveCustomizedLogDataBody
}
type DescribeLiveFeeConfigReq struct {
	*DescribeLiveFeeConfigQuery
	*DescribeLiveFeeConfigBody
}
type DescribeLiveFreeTimeIntervalReq struct {
	*DescribeLiveFreeTimeIntervalQuery
	*DescribeLiveFreeTimeIntervalBody
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
type DescribeLivePushStreamCountDataReq struct {
	*DescribeLivePushStreamCountDataQuery
	*DescribeLivePushStreamCountDataBody
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
type DescribeLiveStreamCountDataReq struct {
	*DescribeLiveStreamCountDataQuery
	*DescribeLiveStreamCountDataBody
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
type DescribeLiveStreamUsageDataReq struct {
	*DescribeLiveStreamUsageDataQuery
	*DescribeLiveStreamUsageDataBody
}
type DescribeLiveTimeShiftDataReq struct {
	*DescribeLiveTimeShiftDataQuery
	*DescribeLiveTimeShiftDataBody
}
type DescribeLiveTrafficDataReq struct {
	*DescribeLiveTrafficDataQuery
	*DescribeLiveTrafficDataBody
}
type DescribeLiveTranscodeDataReq struct {
	*DescribeLiveTranscodeDataQuery
	*DescribeLiveTranscodeDataBody
}
type DescribeNSSRewriteConfigReq struct {
	*DescribeNSSRewriteConfigQuery
	*DescribeNSSRewriteConfigBody
}
type DescribePresetAssociationReq struct {
	*DescribePresetAssociationQuery
	*DescribePresetAssociationBody
}
type DescribeProxyConfigAssociationReq struct {
	*DescribeProxyConfigAssociationQuery
	*DescribeProxyConfigAssociationBody
}
type DescribeRecordTaskFileHistoryReq struct {
	*DescribeRecordTaskFileHistoryQuery
	*DescribeRecordTaskFileHistoryBody
}
type DescribeRefConfigReq struct {
	*DescribeRefConfigQuery
	*DescribeRefConfigBody
}
type DescribeRefererReq struct {
	*DescribeRefererQuery
	*DescribeRefererBody
}
type DescribeRelaySinkReq struct {
	*DescribeRelaySinkQuery
	*DescribeRelaySinkBody
}
type DescribeRelaySourceRewriteReq struct {
	*DescribeRelaySourceRewriteQuery
	*DescribeRelaySourceRewriteBody
}
type DescribeRelaySourceV3Req struct {
	*DescribeRelaySourceV3Query
	*DescribeRelaySourceV3Body
}
type DescribeSDKDetailReq struct {
	*DescribeSDKDetailQuery
	*DescribeSDKDetailBody
}
type DescribeSDKParamsAvailableReq struct {
	*DescribeSDKParamsAvailableQuery
	*DescribeSDKParamsAvailableBody
}
type DescribeServiceReq struct {
	*DescribeServiceQuery
	*DescribeServiceBody
}
type DescribeSnapshotAuditPresetDetailReq struct {
	*DescribeSnapshotAuditPresetDetailQuery
	*DescribeSnapshotAuditPresetDetailBody
}
type DescribeStreamQuotaConfigReq struct {
	*DescribeStreamQuotaConfigQuery
	*DescribeStreamQuotaConfigBody
}
type DescribeTimeShiftPresetDetailReq struct {
	*DescribeTimeShiftPresetDetailQuery
	*DescribeTimeShiftPresetDetailBody
}
type DescribeTranscodePresetDetailReq struct {
	*DescribeTranscodePresetDetailQuery
	*DescribeTranscodePresetDetailBody
}
type DescribeVQScoreTaskReq struct {
	*DescribeVQScoreTaskQuery
	*DescribeVQScoreTaskBody
}
type DescribeVhostReq struct {
	*DescribeVhostQuery
	*DescribeVhostBody
}
type DescribeWatermarkPresetDetailReq struct {
	*DescribeWatermarkPresetDetailQuery
	*DescribeWatermarkPresetDetailBody
}
type DisAssociatePresetReq struct {
	*DisAssociatePresetQuery
	*DisAssociatePresetBody
}
type DisableAuthReq struct {
	*DisableAuthQuery
	*DisableAuthBody
}
type DisableDomainReq struct {
	*DisableDomainQuery
	*DisableDomainBody
}
type DisassociateRefConfigReq struct {
	*DisassociateRefConfigQuery
	*DisassociateRefConfigBody
}
type EnableAuthReq struct {
	*EnableAuthQuery
	*EnableAuthBody
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
type GenerateTimeShiftPlayURLReq struct {
	*GenerateTimeShiftPlayURLQuery
	*GenerateTimeShiftPlayURLBody
}
type GetAppsReq struct {
	*GetAppsQuery
	*GetAppsBody
}
type GetPullCDNSnapshotTaskReq struct {
	*GetPullCDNSnapshotTaskQuery
	*GetPullCDNSnapshotTaskBody
}
type GetPullRecordTaskReq struct {
	*GetPullRecordTaskQuery
	*GetPullRecordTaskBody
}
type GetTagsReq struct {
	*GetTagsQuery
	*GetTagsBody
}
type GetVqosRawDataReq struct {
	*GetVqosRawDataQuery
	*GetVqosRawDataBody
}
type KillStreamReq struct {
	*KillStreamQuery
	*KillStreamBody
}
type ListActionHistoryReq struct {
	*ListActionHistoryQuery
	*ListActionHistoryBody
}
type ListBindEncryptDRMReq struct {
	*ListBindEncryptDRMQuery
	*ListBindEncryptDRMBody
}
type ListCertReq struct {
	*ListCertQuery
	*ListCertBody
}
type ListCertBindInfoReq struct {
	*ListCertBindInfoQuery
	*ListCertBindInfoBody
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
type ListHeaderEnumReq struct {
	*ListHeaderEnumQuery
	*ListHeaderEnumBody
}
type ListInstanceReq struct {
	*ListInstanceQuery
	*ListInstanceBody
}
type ListObjectReq struct {
	*ListObjectQuery
	*ListObjectBody
}
type ListProjectsReq struct {
	*ListProjectsQuery
	*ListProjectsBody
}
type ListProxyConfigReq struct {
	*ListProxyConfigQuery
	*ListProxyConfigBody
}
type ListPullCDNSnapshotTaskReq struct {
	*ListPullCDNSnapshotTaskQuery
	*ListPullCDNSnapshotTaskBody
}
type ListPullRecordTaskReq struct {
	*ListPullRecordTaskQuery
	*ListPullRecordTaskBody
}
type ListPullToPushTaskReq struct {
	*ListPullToPushTaskQuery
	*ListPullToPushTaskBody
}
type ListReferenceInfoReq struct {
	*ListReferenceInfoQuery
	*ListReferenceInfoBody
}
type ListReferenceNamesReq struct {
	*ListReferenceNamesQuery
	*ListReferenceNamesBody
}
type ListReferenceTypesReq struct {
	*ListReferenceTypesQuery
	*ListReferenceTypesBody
}
type ListRelaySourceV4Req struct {
	*ListRelaySourceV4Query
	*ListRelaySourceV4Body
}
type ListResourcePackageReq struct {
	*ListResourcePackageQuery
	*ListResourcePackageBody
}
type ListSDKReq struct {
	*ListSDKQuery
	*ListSDKBody
}
type ListSDKAdminReq struct {
	*ListSDKAdminQuery
	*ListSDKAdminBody
}
type ListServicesReq struct {
	*ListServicesQuery
	*ListServicesBody
}
type ListTimeShiftPresetV2Req struct {
	*ListTimeShiftPresetV2Query
	*ListTimeShiftPresetV2Body
}
type ListVQScoreTaskReq struct {
	*ListVQScoreTaskQuery
	*ListVQScoreTaskBody
}
type ListVhostDenseSnapshotPresetReq struct {
	*ListVhostDenseSnapshotPresetQuery
	*ListVhostDenseSnapshotPresetBody
}
type ListVhostDetailReq struct {
	*ListVhostDetailQuery
	*ListVhostDetailBody
}
type ListVhostDetailByAdminReq struct {
	*ListVhostDetailByAdminQuery
	*ListVhostDetailByAdminBody
}
type ListVhostDomainDetailByUserIDReq struct {
	*ListVhostDomainDetailByUserIDQuery
	*ListVhostDomainDetailByUserIDBody
}
type ListVhostRecordPresetV2Req struct {
	*ListVhostRecordPresetV2Query
	*ListVhostRecordPresetV2Body
}
type ListVhostSnapshotAuditPresetReq struct {
	*ListVhostSnapshotAuditPresetQuery
	*ListVhostSnapshotAuditPresetBody
}
type ListVhostSnapshotPresetReq struct {
	*ListVhostSnapshotPresetQuery
	*ListVhostSnapshotPresetBody
}
type ListVhostSnapshotPresetV2Req struct {
	*ListVhostSnapshotPresetV2Query
	*ListVhostSnapshotPresetV2Body
}
type ListVhostTransCodePresetReq struct {
	*ListVhostTransCodePresetQuery
	*ListVhostTransCodePresetBody
}
type ListVhostWatermarkPresetReq struct {
	*ListVhostWatermarkPresetQuery
	*ListVhostWatermarkPresetBody
}
type ListVideoClassificationsReq struct {
	*ListVideoClassificationsQuery
	*ListVideoClassificationsBody
}
type ListVqosDimensionValuesReq struct {
	*ListVqosDimensionValuesQuery
	*ListVqosDimensionValuesBody
}
type ListVqosMetricsDimensionsReq struct {
	*ListVqosMetricsDimensionsQuery
	*ListVqosMetricsDimensionsBody
}
type ListWatermarkPresetReq struct {
	*ListWatermarkPresetQuery
	*ListWatermarkPresetBody
}
type ManagerPullPushDomainBindReq struct {
	*ManagerPullPushDomainBindQuery
	*ManagerPullPushDomainBindBody
}
type RejectDomainReq struct {
	*RejectDomainQuery
	*RejectDomainBody
}
type RestartPullToPushTaskReq struct {
	*RestartPullToPushTaskQuery
	*RestartPullToPushTaskBody
}
type ResumeStreamReq struct {
	*ResumeStreamQuery
	*ResumeStreamBody
}
type StopPullCDNSnapshotTaskReq struct {
	*StopPullCDNSnapshotTaskQuery
	*StopPullCDNSnapshotTaskBody
}
type StopPullRecordTaskReq struct {
	*StopPullRecordTaskQuery
	*StopPullRecordTaskBody
}
type StopPullToPushTaskReq struct {
	*StopPullToPushTaskQuery
	*StopPullToPushTaskBody
}
type TerminateInstanceReq struct {
	*TerminateInstanceQuery
	*TerminateInstanceBody
}
type UnBindEncryptDRMReq struct {
	*UnBindEncryptDRMQuery
	*UnBindEncryptDRMBody
}
type UnbindCertReq struct {
	*UnbindCertQuery
	*UnbindCertBody
}
type UpdateActivityBillingReq struct {
	*UpdateActivityBillingQuery
	*UpdateActivityBillingBody
}
type UpdateAppReq struct {
	*UpdateAppQuery
	*UpdateAppBody
}
type UpdateAuthKeyReq struct {
	*UpdateAuthKeyQuery
	*UpdateAuthKeyBody
}
type UpdateAvSlicePresetReq struct {
	*UpdateAvSlicePresetQuery
	*UpdateAvSlicePresetBody
}
type UpdateBillingReq struct {
	*UpdateBillingQuery
	*UpdateBillingBody
}
type UpdateCMAFConfigReq struct {
	*UpdateCMAFConfigQuery
	*UpdateCMAFConfigBody
}
type UpdateCallbackReq struct {
	*UpdateCallbackQuery
	*UpdateCallbackBody
}
type UpdateCertReq struct {
	*UpdateCertQuery
	*UpdateCertBody
}
type UpdateDenseSnapshotPresetReq struct {
	*UpdateDenseSnapshotPresetQuery
	*UpdateDenseSnapshotPresetBody
}
type UpdateDenyConfigReq struct {
	*UpdateDenyConfigQuery
	*UpdateDenyConfigBody
}
type UpdateDenyConfigV2Req struct {
	*UpdateDenyConfigV2Query
	*UpdateDenyConfigV2Body
}
type UpdateDomainReq struct {
	*UpdateDomainQuery
	*UpdateDomainBody
}
type UpdateDomainVhostReq struct {
	*UpdateDomainVhostQuery
	*UpdateDomainVhostBody
}
type UpdateEncryptDRMReq struct {
	*UpdateEncryptDRMQuery
	*UpdateEncryptDRMBody
}
type UpdateHLSConfigReq struct {
	*UpdateHLSConfigQuery
	*UpdateHLSConfigBody
}
type UpdateHTTPHeaderConfigReq struct {
	*UpdateHTTPHeaderConfigQuery
	*UpdateHTTPHeaderConfigBody
}
type UpdateHeaderConfigReq struct {
	*UpdateHeaderConfigQuery
	*UpdateHeaderConfigBody
}
type UpdateIPAccessRuleReq struct {
	*UpdateIPAccessRuleQuery
	*UpdateIPAccessRuleBody
}
type UpdateLatencyConfigReq struct {
	*UpdateLatencyConfigQuery
	*UpdateLatencyConfigBody
}
type UpdateNSSRewriteConfigReq struct {
	*UpdateNSSRewriteConfigQuery
	*UpdateNSSRewriteConfigBody
}
type UpdatePresetAssociationReq struct {
	*UpdatePresetAssociationQuery
	*UpdatePresetAssociationBody
}
type UpdateProxyConfigReq struct {
	*UpdateProxyConfigQuery
	*UpdateProxyConfigBody
}
type UpdateProxyConfigAssociationReq struct {
	*UpdateProxyConfigAssociationQuery
	*UpdateProxyConfigAssociationBody
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
type UpdateRelaySinkReq struct {
	*UpdateRelaySinkQuery
	*UpdateRelaySinkBody
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
type UpdateSDKReq struct {
	*UpdateSDKQuery
	*UpdateSDKBody
}
type UpdateSDKLicenseReq struct {
	*UpdateSDKLicenseQuery
	*UpdateSDKLicenseBody
}
type UpdateServiceReq struct {
	*UpdateServiceQuery
	*UpdateServiceBody
}
type UpdateSnapshotAuditPresetReq struct {
	*UpdateSnapshotAuditPresetQuery
	*UpdateSnapshotAuditPresetBody
}
type UpdateSnapshotPresetReq struct {
	*UpdateSnapshotPresetQuery
	*UpdateSnapshotPresetBody
}
type UpdateSnapshotPresetV2Req struct {
	*UpdateSnapshotPresetV2Query
	*UpdateSnapshotPresetV2Body
}
type UpdateStreamQuotaConfigReq struct {
	*UpdateStreamQuotaConfigQuery
	*UpdateStreamQuotaConfigBody
}
type UpdateStreamQuotaConfigPatchReq struct {
	*UpdateStreamQuotaConfigPatchQuery
	*UpdateStreamQuotaConfigPatchBody
}
type UpdateTimeShiftPresetV2Req struct {
	*UpdateTimeShiftPresetV2Query
	*UpdateTimeShiftPresetV2Body
}
type UpdateTimeShiftPresetV3Req struct {
	*UpdateTimeShiftPresetV3Query
	*UpdateTimeShiftPresetV3Body
}
type UpdateTranscodePresetReq struct {
	*UpdateTranscodePresetQuery
	*UpdateTranscodePresetBody
}
type UpdateVhostTagsReq struct {
	*UpdateVhostTagsQuery
	*UpdateVhostTagsBody
}
type UpdateWatermarkPresetReq struct {
	*UpdateWatermarkPresetQuery
	*UpdateWatermarkPresetBody
}
type UpdateWatermarkPresetV2Req struct {
	*UpdateWatermarkPresetV2Query
	*UpdateWatermarkPresetV2Body
}
type ValidateCertReq struct {
	*ValidateCertQuery
	*ValidateCertBody
}
type VerifyDomainOwnerReq struct {
	*VerifyDomainOwnerQuery
	*VerifyDomainOwnerBody
}
