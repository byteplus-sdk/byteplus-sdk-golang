package live_v20230101

type CreateSubtitleTranscodePresetBodyPositionRelative string

type Enum0 string

type BindCertBody struct {

	// REQUIRED; 需要绑定的 HTTPS 证书的证书链 ID，可以通过查询证书列表 [https://www.volcengine.com/docs/6469/1126822]接口获取。
	ChainID string `json:"ChainID"`

	// REQUIRED; 填写需要配置 HTTPS 证书的域名。 您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看需要绑定证书的域名。
	Domain string `json:"Domain"`

	// 证书域名。
	CertDomain *string `json:"CertDomain,omitempty"`

	// 是否开启 HTTP/2 协议，默认为 false。取值如下：
	// * false: 关闭 HTTP/2 协议。
	// * true: 开启 HTTP/2 协议。
	HTTP2 *bool `json:"HTTP2,omitempty"`

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

	// REQUIRED; 加密类型，支持的取值及含义如下所示。
	// * FairPlay：使用 FairPlay 技术的商业 DRM 加密；
	// * Widevine：使用 Widevine 技术的商业 DRM 加密；
	// * PlayReady：使用 PlayReady 技术的商业 DRM 加密；
	// * ClearKey：HLS 标准加密。
	// :::tip DRM 加密与 HLS 标准加密不可同时配置。 :::
	DRMSystems []string `json:"DRMSystems"`

	// REQUIRED; 是否开启源流加密，取值及含义如下所示。
	// * true：开启；
	// * fasle：不开启。 :::tip 源流和转码流至少有一个需要开启录制。 :::
	EncryptOriginStream bool `json:"EncryptOriginStream"`

	// REQUIRED; 是否开启转码流加密，取值及含义如下所示。
	// * true：开启；
	// * fasle：不开启。 :::tip 源流和转码流至少有一个需要开启录制。 :::
	EncryptTranscodeStream bool `json:"EncryptTranscodeStream"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 是否开启当前 DRM 加密配置，取值及含义如下所示。
	// * true：（默认值）开启；
	// * false：关闭。
	Enable *bool `json:"Enable,omitempty"`

	// 开启转码流加密时待加密的转码流对应的转码流后缀配置。您可以调用查询转码配置列表 [https://www.volcengine.com/docs/6469/1126853]接口或在视频直播控制台的转码配置 [https://console.volcengine.com/live/main/application/transcode]页面，查看转码配置的转码流后缀。
	EncryptTranscodeSuffix []*string `json:"EncryptTranscodeSuffix,omitempty"`
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

// Components1523StvSchemasListvhostsubtitletranscodepresetresPropertiesResultPropertiesPresetlistItemsPropertiesTranscodepresetPropertiesSourcelanguage
// - 原文字幕展示参数配置。
type Components1523StvSchemasListvhostsubtitletranscodepresetresPropertiesResultPropertiesPresetlistItemsPropertiesTranscodepresetPropertiesSourcelanguage struct {
	Border Components1O8E0AlSchemasListvhostsubtitletranscodepresetresPropertiesResultPropertiesPresetlistItemsPropertiesTranscodepresetPropertiesSourcelanguagePropertiesBorder `json:"Border"`

	Display bool `json:"Display"`

	Font string `json:"Font"`

	FontColor string `json:"FontColor"`

	FontSize int32 `json:"FontSize"`

	Language string `json:"Language"`
}

type Components17Ohct5SchemasDescribeliveasrdurationdataresPropertiesResultPropertiesAsrdurationdetaildataItemsPropertiesAsrdurationdataItems struct {
	Duration float32 `json:"Duration"`

	TimeStamp string `json:"TimeStamp"`
}

type Components1C398ShSchemasListvhostsubtitletranscodepresetresPropertiesResultPropertiesPresetlistItemsPropertiesTranscodepresetPropertiesTargetlanguageItems struct {
	Border ListVhostSubtitleTranscodePresetResResultPresetListItemTranscodePresetTargetLanguageItemBorder `json:"Border"`

	Font string `json:"Font"`

	FontColor string `json:"FontColor"`

	FontSize int32 `json:"FontSize"`

	Language string `json:"Language"`
}

// Components1Hkcrc4SchemasListvhostsnapshotpresetresPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetPropertiesCallbackdetail
// - 回调信息。
type Components1Hkcrc4SchemasListvhostsnapshotpresetresPropertiesResultPropertiesPresetlistItemsPropertiesSlicepresetPropertiesCallbackdetail struct {
	URL string `json:"URL"`

	CallbackType *string `json:"CallbackType,omitempty"`
}

type Components1Nf1A8CSchemasListpulltopushtaskv2ResPropertiesResultPropertiesListItemsPropertiesVodsrcaddrsItems struct {
	SrcAddr string `json:"SrcAddr"`

	EndOffset *float32 `json:"EndOffset,omitempty"`

	StartOffset *float32 `json:"StartOffset,omitempty"`
}

// Components1O8E0AlSchemasListvhostsubtitletranscodepresetresPropertiesResultPropertiesPresetlistItemsPropertiesTranscodepresetPropertiesSourcelanguagePropertiesBorder
// - 原文字幕的字体描边配置。
type Components1O8E0AlSchemasListvhostsubtitletranscodepresetresPropertiesResultPropertiesPresetlistItemsPropertiesTranscodepresetPropertiesSourcelanguagePropertiesBorder struct {
	Color string `json:"Color"`

	Width int32 `json:"Width"`
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

type Components6Aa5KwSchemasListvhostremoteauthresPropertiesResultPropertiesRemoteauthconfiglistItemsPropertiesHeaderparamconfigPropertiesParamsItems struct {
	Type      string  `json:"Type"`
	ParamName *string `json:"ParamName,omitempty"`
	ToName    *string `json:"ToName,omitempty"`
	Value     *string `json:"Value,omitempty"`
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

type ComponentsGg7M1TSchemasListpulltopushtaskresPropertiesResultPropertiesListItemsPropertiesVodsrcaddrsItems struct {
	SrcAddr string `json:"SrcAddr"`

	EndOffset *float32 `json:"EndOffset,omitempty"`

	StartOffset *float32 `json:"StartOffset,omitempty"`
}

// ComponentsJ1MbxoSchemasListvhostsubtitletranscodepresetresPropertiesResultPropertiesPresetlistItemsPropertiesTranscodepresetPropertiesPosition
// - 字幕位置设置，通过设置字幕距离画面左右边距和底部边距来指定字幕位置。
// :::tip
// * 使用预设配置时，字幕位置设置不生效。
// * 不使用预设配置时，字幕位置设置必填。 :::
type ComponentsJ1MbxoSchemasListvhostsubtitletranscodepresetresPropertiesResultPropertiesPresetlistItemsPropertiesTranscodepresetPropertiesPosition struct {
	MarginHorizontal float32 `json:"MarginHorizontal"`

	MarginVertical float32 `json:"MarginVertical"`

	Relative string `json:"Relative"`
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

// ComponentsNopwcvSchemasListvhostremoteauthresPropertiesResultPropertiesRemoteauthconfiglistItemsPropertiesCacheconfigPropertiesDenykeys
// - 生成鉴权失败结果缓存 Key 使用的参数配置，最多支持配置 50 个参数。
// :::tip 鉴权成功与鉴权失败不使用相同配置，即 UseSameCache 配置 false 时生效。 :::
type ComponentsNopwcvSchemasListvhostremoteauthresPropertiesResultPropertiesRemoteauthconfiglistItemsPropertiesCacheconfigPropertiesDenykeys struct {
	Type      string  `json:"Type"`
	ParamName *string `json:"ParamName,omitempty"`
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

// ComponentsTmguxbSchemasListvhostremoteauthresPropertiesResultPropertiesRemoteauthconfiglistItemsPropertiesCacheconfigPropertiesCachekeys
// - 生成鉴权结果缓存 Key 使用的参数配置，最多支持配置 50 个参数。 :::tip
// * 鉴权成功与鉴权失败使用相同配置时，表示鉴权结果缓存 Key 配置；
// * 鉴权成功与鉴权失败不使用相同配置时，表示鉴权成功缓存的过期时间。 :::
type ComponentsTmguxbSchemasListvhostremoteauthresPropertiesResultPropertiesRemoteauthconfiglistItemsPropertiesCacheconfigPropertiesCachekeys struct {
	Type string `json:"Type"`

	ParamName *string `json:"ParamName,omitempty"`
}

type ContinuePullToPushTaskBody struct {

	// REQUIRED; 任务 ID，任务的唯一标识，您可以通过获取拉流转推任务列表 [https://www.volcengine.com/docs/6469/1126896]接口获取状态为停用的任务 ID。
	TaskID string `json:"TaskId"`

	// 任务所属的群组名称，您可以通过获取拉流转推任务列表 [https://www.volcengine.com/docs/6469/1126896]接口获取。 :::tip
	// * 使用主账号调用时，为非必填。
	// * 使用子账号调用时，为必填。 :::
	GroupName *string `json:"GroupName,omitempty"`
}

type ContinuePullToPushTaskRes struct {

	// REQUIRED
	ResponseMetadata ContinuePullToPushTaskResResponseMetadata `json:"ResponseMetadata"`
}

type ContinuePullToPushTaskResResponseMetadata struct {

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
	Error   *ContinuePullToPushTaskResResponseMetadataError `json:"Error,omitempty"`
}

type ContinuePullToPushTaskResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type CreateCarouselTaskBody struct {

	// REQUIRED; 轮播任务名称，运行中的任务，名称不能重复
	Name string `json:"Name"`

	// REQUIRED; 轮播规则，用于指定轮播播放的素材和行为等。
	Rule CreateCarouselTaskBodyRule `json:"Rule"`
}

// CreateCarouselTaskBodyRule - 轮播规则，用于指定轮播播放的素材和行为等。
type CreateCarouselTaskBodyRule struct {

	// REQUIRED; -1代表无限循环，最终停止由StopTime字段控制，或者系统默认的停止时间（3天）
	Loop int32 `json:"Loop"`

	// REQUIRED; 0为普通模式，此模式下系统会根据前后两个点播素材的头信息来判断是否能不断流拼接，如果不满足拼接条件，在进行素材切换时会断流。 2为转码模式，此模式下系统会将所有素材格式化为固定参数，用户可以配置这个音视频参数，如果不配置默认参数跟随第一个素材，在进行素材切换时不会断流
	Mode int32 `json:"Mode"`

	// REQUIRED; 轮播任务的推流参数，包括视频、音频、推流地址及回调信息。
	Output CreateCarouselTaskBodyRuleOutput `json:"Output"`

	// REQUIRED; 轮播素材列表，用于指定在轮播过程中播放的素材资源。
	Source []CreateCarouselTaskBodyRuleSourceItem `json:"Source"`

	// 播放时间，选填，默认会等待第一个视频缓存完毕，如果系统时间大于此值，则开始播放
	PlayTime *int32 `json:"PlayTime,omitempty"`

	// 停止时间，选填，当此字段被设置时，系统会遵循此时间设置关闭任务
	StopTime *int32 `json:"StopTime,omitempty"`
}

// CreateCarouselTaskBodyRuleOutput - 轮播任务的推流参数，包括视频、音频、推流地址及回调信息。
type CreateCarouselTaskBodyRuleOutput struct {

	// REQUIRED; 推流rtmp地址或者rtmps地址，支持多推，最多填写8条地址，最少1条地址
	URL []string `json:"Url"`

	// 转码模式下有效，可选配置推流的音频参数
	Audio *CreateCarouselTaskBodyRuleOutputAudio `json:"Audio,omitempty"`

	// 回调函数。
	Callback *CreateCarouselTaskBodyRuleOutputCallback `json:"Callback,omitempty"`

	// 转码模式下有效，可选配置推流的视频参数
	Video *CreateCarouselTaskBodyRuleOutputVideo `json:"Video,omitempty"`
}

// CreateCarouselTaskBodyRuleOutputAudio - 转码模式下有效，可选配置推流的音频参数
type CreateCarouselTaskBodyRuleOutputAudio struct {

	// 音频码率设置
	BitRate *int32 `json:"BitRate,omitempty"`

	// mono：单声道；stereo：双声道
	ChannelLayout *string `json:"ChannelLayout,omitempty"`

	// 采样率，可选：22000、32000、44100、48000
	SampleRate *int32 `json:"SampleRate,omitempty"`
}

// CreateCarouselTaskBodyRuleOutputCallback - 回调函数。
type CreateCarouselTaskBodyRuleOutputCallback struct {

	// REQUIRED; 回调地址，系统会将部分信息回调出去
	URL string `json:"Url"`
}

// CreateCarouselTaskBodyRuleOutputVideo - 转码模式下有效，可选配置推流的视频参数
type CreateCarouselTaskBodyRuleOutputVideo struct {

	// 视频码率，单位是bit
	BitRate *int32 `json:"BitRate,omitempty"`

	// 视频帧率取值[10,60]
	FrameRate *int32 `json:"FrameRate,omitempty"`

	// 取值[1,10]
	GOP *int32 `json:"GOP,omitempty"`

	// 取值范围[10,2160]
	Height *int32 `json:"Height,omitempty"`

	// 编码档位，可选：faster、medium、veryfast
	Preset *string `json:"Preset,omitempty"`

	// 取值范围[10,2160]
	Width *int32 `json:"Width,omitempty"`
}

type CreateCarouselTaskBodyRuleSourceItem struct {

	// REQUIRED; 播放列表内不允许重复
	ID string `json:"ID"`

	// REQUIRED; vod：点播文件；m3u8：m3u8文件
	Type string `json:"Type"`

	// REQUIRED; 轮播素材的公网可访问地址。确保提供的地址能够被公网正常访问，以便正确加载轮播素材内容。
	URL string `json:"Url"`

	// 此素材连续播放几次，字段必须大于等于0
	Loop *int32 `json:"Loop,omitempty"`

	// 可以控制当前素材跳过开头进行播放，单位是秒，注意此字段如果小于等于0或者大于视频长度不生效，只在素材type为vod时生效
	Seek *int32 `json:"Seek,omitempty"`
}

type CreateCarouselTaskRes struct {

	// REQUIRED
	ResponseMetadata CreateCarouselTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result CreateCarouselTaskResResult `json:"Result"`
}

type CreateCarouselTaskResResponseMetadata struct {

	// REQUIRED
	RequestID string `json:"RequestID"`
}

type CreateCarouselTaskResResult struct {

	// REQUIRED; 轮播任务数据对象。
	Data CreateCarouselTaskResResultData `json:"Data"`
}

// CreateCarouselTaskResResultData - 轮播任务数据对象。
type CreateCarouselTaskResResultData struct {

	// REQUIRED; 任务唯一标识
	TaskID string `json:"TaskID"`
}

type CreateCertBody struct {

	// REQUIRED; 证书信息。
	Rsa CreateCertBodyRsa `json:"Rsa"`

	// REQUIRED; 证书用途，当前仅支持设置为 https，表示用于 HTTPS 加密；
	UseWay string `json:"UseWay"`

	// 证书名称。
	CertName *string `json:"CertName,omitempty"`

	// 证书链 ID，用于标识整个证书链，包括叶子证书（服务器证书）、中间证书（中间 CA 证书）以及根证书（根 CA 证书）。 :::tip 使用当前接口更新证书时， ChainID 为必选参数。 :::
	ChainID *string `json:"ChainID,omitempty"`

	// 项目名称，默认值为 default，您可以登录访问控制 [https://console.volcengine.com/iam/resourcemanage/project]获取项目名称。
	ProjectName *string `json:"ProjectName,omitempty"`
}

// CreateCertBodyRsa - 证书信息。
type CreateCertBodyRsa struct {

	// REQUIRED; 证书私钥的内容，你需要在计算机上使用文本编辑器打开证书私钥，并将所有内容复制粘贴作为参数。 :::tip 请确保证书私钥没有密码保护。 :::
	Prikey string `json:"Prikey"`

	// REQUIRED; 证书内容，你需要在计算机上使用文本编辑器打开证书，并将所有内容复制粘贴作为参数。 :::tip
	// * 视频直播支持证书链校验。你只需要上传为你的域名颁发的证书，系统将自动检索完整的证书链。
	// * 如果你选择上传证书链，请务必包含服务器证书、中间证书和根证书，并按正确的顺序排列：首先是服务器证书，其次是中间证书，然后是根证书。错误的顺序将使证书链无效。 :::
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

type CreateCloudMixTaskBody struct {

	// REQUIRED; 混流任务详细配置。
	MixedRules CreateCloudMixTaskBodyMixedRules `json:"MixedRules"`

	// REQUIRED; 混流任务名称，与正在进行中的任务名称不能重复。
	Name string `json:"Name"`
}

// CreateCloudMixTaskBodyMixedRules - 混流任务详细配置。
type CreateCloudMixTaskBodyMixedRules struct {

	// REQUIRED; 混流输出布局配置。
	InputLayout CreateCloudMixTaskBodyMixedRulesInputLayout `json:"InputLayout"`

	// REQUIRED; 混流素材列表，最多支持配置 8 路素材。
	InputSource []CreateCloudMixTaskBodyMixedRulesInputSourceItem `json:"InputSource"`

	// REQUIRED; 混流输出视频质量参数配置。
	Output CreateCloudMixTaskBodyMixedRulesOutput `json:"Output"`
}

// CreateCloudMixTaskBodyMixedRulesInputLayout - 混流输出布局配置。
type CreateCloudMixTaskBodyMixedRulesInputLayout struct {

	// REQUIRED; 混流输出画布配置及素材布局配置。
	Scene CreateCloudMixTaskBodyMixedRulesInputLayoutScene `json:"Scene"`

	// 混流输出视频中 Logo 布局配置。
	// :::tip 支持最多配置 4 个 Logo，展示层级以添加顺序为准。 :::
	Logo []*CreateCloudMixTaskBodyMixedRulesInputLayoutLogoItem `json:"Logo,omitempty"`

	// 预设混流模板 ID。 模板配置还有变动，后续稳定后对外，且后续支持用户自定义模板。
	TemplateID *string `json:"TemplateID,omitempty"`
}

type CreateCloudMixTaskBodyMixedRulesInputLayoutLogoItem struct {

	// REQUIRED; Logo 图片在混流输出整体画面中的布局配置。
	Layout CreateCloudMixTaskBodyMixedRulesInputLayoutLogoItemLayout `json:"Layout"`

	// REQUIRED; Logo 图片访问地址。
	URL string `json:"Url"`
}

// CreateCloudMixTaskBodyMixedRulesInputLayoutLogoItemLayout - Logo 图片在混流输出整体画面中的布局配置。
type CreateCloudMixTaskBodyMixedRulesInputLayoutLogoItemLayout struct {

	// REQUIRED
	H int32 `json:"H"`

	// REQUIRED
	W int32 `json:"W"`

	// REQUIRED
	X int32 `json:"X"`

	// REQUIRED
	Y int32 `json:"Y"`
}

// CreateCloudMixTaskBodyMixedRulesInputLayoutScene - 混流输出画布配置及素材布局配置。
type CreateCloudMixTaskBodyMixedRulesInputLayoutScene struct {

	// REQUIRED; 混流输出整体画布高度，单位为 px，取值范围为 [10,2160]。
	Height int32 `json:"Height"`

	// REQUIRED; 混流素材在混流输出整体画面中的布局配置。 :::tip 混流素材布局中需包含所有素材的配置，且需与通过 Layer 参数与混流素材一一匹配。 :::
	Layout []CreateCloudMixTaskBodyMixedRulesInputLayoutSceneLayoutItem `json:"Layout"`

	// REQUIRED; 混流输出画布整体宽度，单位为 px，取值范围为 [10,2160]。
	Width int32 `json:"Width"`
}

type CreateCloudMixTaskBodyMixedRulesInputLayoutSceneLayoutItem struct {

	// REQUIRED; 当前素材或 Logo 图片在混流输出画面中的限制高度，单位为 px，取值范围为 [10,2160]。
	// :::tip 限制宽度和限制高度指定了素材展示的限制范围，当素材尺寸和限制值不一致时，素材将在限制范围内根据长边进行等比缩放，其余区域透明展示。 :::
	H int32 `json:"H"`

	// REQUIRED; 当配置素材布局时需要通过 Layer 参数与素材进行一一对应。 :::tip 配置 Logo 图片的布局时此参数不生效。 :::
	Layer int32 `json:"Layer"`

	// REQUIRED; 当前素材或 Logo 图片在混流输出画面中的限制宽度，单位为 px，取值范围为 [10,2160]。
	W int32 `json:"W"`

	// REQUIRED; 当前素材或 Logo 图片在输出画面中相对画面左上角的 X 偏移位置，单位为 px，取值范围为 0 到设置的画面宽度。
	X int32 `json:"X"`

	// REQUIRED; 当前素材或 Logo 图片在输出画面中相对画面左上角的 Y 偏移位置，单位为 px，取值范围为 0 到设置的画面高度。
	Y int32 `json:"Y"`
}

type CreateCloudMixTaskBodyMixedRulesInputSourceItem struct {

	// REQUIRED; 混流素材 ID，一个任务中素材 ID 不能重复，此 ID 用于任务状态回调消息中标识素材。
	ID string `json:"ID"`

	// REQUIRED; 混流素材的叠放顺序，1 为最底层，2 层在 1 层之上，以此类推，取值范围为[1,9999]。 :::tip 当前准备使用某个素材作为布局背景时，其叠放顺序应设置为所有素材中的最小值。 :::
	Layer int32 `json:"Layer"`

	// REQUIRED; 混流素材类型，支持的取值及含义如下所示。
	// * vod：视频点播中的素材，支持 MP4、FLV 格式素材；
	// * live：直播源素材，支持 RTMP、FLV 协议拉流地址；
	// * pic：图片素材，支持 png、jpg 格式图片。
	Type string `json:"Type"`

	// REQUIRED; 混流素材的访问地址。 :::tip 混流素材的访问地址需与混流素材的类型保持对应关系。 :::
	URL string `json:"Url"`
}

// CreateCloudMixTaskBodyMixedRulesOutput - 混流输出视频质量参数配置。
type CreateCloudMixTaskBodyMixedRulesOutput struct {

	// REQUIRED; 混流音频参数设置。
	Audio CreateCloudMixTaskBodyMixedRulesOutputAudio `json:"Audio"`

	// REQUIRED; 混流视频的推流地址，支持最多配置 8 个推流地址。
	URL []string `json:"Url"`

	// REQUIRED; 混流视频参数设置。
	Video CreateCloudMixTaskBodyMixedRulesOutputVideo `json:"Video"`

	// 任务状态回调地址配置。
	Callback *CreateCloudMixTaskBodyMixedRulesOutputCallback `json:"Callback,omitempty"`
}

// CreateCloudMixTaskBodyMixedRulesOutputAudio - 混流音频参数设置。
type CreateCloudMixTaskBodyMixedRulesOutputAudio struct {

	// REQUIRED; 混流输出流的音频码率，单位为 bps，取值范围为 [128000,320000]，常见取值及含义如下所示。
	// * 128000：128 kbps；
	// * 144000：144 kbps；
	// * 256000：256 kbps；
	// * 320000：320 kbps。
	BitRate int32 `json:"BitRate"`

	// REQUIRED; 混流输出流的音频声道设置，取值及含义如下所示。
	// * mono：单声道；
	// * stereo：立体声。
	ChannelLayout string `json:"ChannelLayout"`

	// REQUIRED; 混流输出流的音频采样率，单位为 HZ，常见取值及含义如下所示。
	// * 32000：32 kHZ；
	// * 44100：44.1 kHZ；
	// * 48000：48 kHZ。
	SampleRate int32 `json:"SampleRate"`
}

// CreateCloudMixTaskBodyMixedRulesOutputCallback - 任务状态回调地址配置。
type CreateCloudMixTaskBodyMixedRulesOutputCallback struct {

	// REQUIRED; 接收云端混流任务状态回调的 HTTP 地址。
	URL string `json:"Url"`
}

// CreateCloudMixTaskBodyMixedRulesOutputVideo - 混流视频参数设置。
type CreateCloudMixTaskBodyMixedRulesOutputVideo struct {

	// REQUIRED; 混流输出视频码率，单位为 bps，取值范围为 [1000000,20000000]，输入值小于或大于取值范围时会进行自动修正至最小值和最大值。
	BitRate int32 `json:"BitRate"`

	// REQUIRED; 混流输出视频编码格式，支持的取值及含义如下所示。
	// * h264：使用 H.264 编码格式；
	// * h265：使用 H.265 编码格式。
	Codec string `json:"Codec"`

	// REQUIRED; 混流输出视频帧率，单位为 fps，取值范围为 [10,60]，输入值小于或大于取值范围时会进行自动修正至最小值和最大值。
	FrameRate int32 `json:"FrameRate"`

	// REQUIRED; IDR 帧之间的最大间隔时间，单位为秒，取值范围为 [1,10]。
	GOP int32 `json:"GOP"`
}

type CreateCloudMixTaskRes struct {

	// REQUIRED
	ResponseMetadata CreateCloudMixTaskResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *CreateCloudMixTaskResResult `json:"Result,omitempty"`
}

type CreateCloudMixTaskResResponseMetadata struct {

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

// CreateCloudMixTaskResResult - 视请求的接口而定
type CreateCloudMixTaskResResult struct {

	// REQUIRED; 请求响应码，取值及含义如下。
	// * 0：请求成功；
	// * 500：内部处理错误；
	// * 400：请求异常。
	Code int32 `json:"Code"`

	// REQUIRED; 返回的数据。
	Data CreateCloudMixTaskResResultData `json:"Data"`

	// REQUIRED; 请求响应码对应的信息。
	Message string `json:"Message"`
}

// CreateCloudMixTaskResResultData - 返回的数据。
type CreateCloudMixTaskResResultData struct {

	// REQUIRED; 混流任务 ID。
	TaskID string `json:"TaskID"`
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

type CreateLiveVideoQualityAnalysisTaskBody struct {

	// REQUIRED; 源流地址，支持flv、hls、rtmp地址
	SrcURL string `json:"SrcURL"`

	// 任务运行时常，支持60-300，单位s，默认180
	Duration *int32 `json:"Duration,omitempty"`

	// 截图间隔，支持3-10s，单位s，不填默认为10s
	Interval *int32 `json:"Interval,omitempty"`

	// 任务名称，用来查询使用。
	Name *string `json:"Name,omitempty"`
}

type CreateLiveVideoQualityAnalysisTaskRes struct {

	// REQUIRED
	ResponseMetadata CreateLiveVideoQualityAnalysisTaskResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *CreateLiveVideoQualityAnalysisTaskResResult `json:"Result,omitempty"`
}

type CreateLiveVideoQualityAnalysisTaskResResponseMetadata struct {

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

// CreateLiveVideoQualityAnalysisTaskResResult - 视请求的接口而定
type CreateLiveVideoQualityAnalysisTaskResResult struct {

	// REQUIRED; 任务ID
	ID int64 `json:"ID"`

	// REQUIRED; 任务名称
	Name string `json:"Name"`
}

type CreatePullToPushGroupBody struct {

	// REQUIRED; 群组名称，支持有中文、大小写字母和数字组成，最大长度为 20 个字符。
	Name string `json:"Name"`

	// REQUIRED; 为任务群组设置所属项目，您可以在访问控制-项目列表 [https://console.volcengine.com/iam/resourcemanage/project]查看已有项目并对项目进行管理。 项目是火山引擎提供的一种资源管理方式，您可以对不同业务或项目使用的云资源进行分组管理，以实现基于项目的账单查看、子账号授权等功能。
	ProjectName string `json:"ProjectName"`

	// 为任务群组设置资源标签。您可以通过资源标签从不同维度对云资源进行分类和聚合管理，如资源分账等场景。 :::tip 如需使用标签进行资源分账，可以在可以在账单管理-费用标签 [https://console.volcengine.com/finance/bill/tag/]处管理启用标签，将对应标签运用到账单明细等数据中。
	// :::
	Tags []*CreatePullToPushGroupBodyTagsItem `json:"Tags,omitempty"`
}

type CreatePullToPushGroupBodyTagsItem struct {

	// REQUIRED; 标签 Key 值。
	Key string `json:"Key"`

	// REQUIRED; 标签 Value 值。
	Value string `json:"Value"`
}

type CreatePullToPushGroupRes struct {

	// REQUIRED
	ResponseMetadata CreatePullToPushGroupResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type CreatePullToPushGroupResResponseMetadata struct {

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

type CreatePullToPushTaskBody struct {

	// REQUIRED; 任务的结束时间，Unix 时间戳，单位为秒。 :::tip 拉流转推任务持续时间最长为 7 天。 :::
	EndTime int32 `json:"EndTime"`

	// REQUIRED; 任务的开始时间，Unix 时间戳，单位为秒。 :::tip 拉流转推任务持续时间最长为 7 天。 :::
	StartTime int32 `json:"StartTime"`

	// REQUIRED; 拉流来源类型，支持的取值及含义如下。
	// * 0：直播源；
	// * 1：点播视频。
	Type int32 `json:"Type"`

	// 推流应用名称，推流地址（DstAddr）为空时必传；反之，则该参数不生效。
	App *string `json:"App,omitempty"`

	// 接收拉流转推任务状态回调的地址，最大长度为 512 个字符，默认为空。
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

	// 群组所属名称，您可以调用 ListPullToPushGroup [https://www.volcengine.com/docs/6469/1327382] 获取可用的群组。 :::tip
	// * 使用主账号调用时，为非必填，默认加入 default 群组，default 群组不存在时会默认创建，并绑定 default 项目。
	// * 使用子账号调用时，为必填。 :::
	GroupName *string `json:"GroupName,omitempty"`

	// 点播文件启播时间偏移值，单位为秒，仅当点播视频播放地址列表（SrcAddrS）只有一个地址，且未配置 Offsets 时生效，缺省情况下为空表示不进行偏移。 :::tip 此字段为旧版本配置，请使用 VodSrcAddrs 配置点播视频地址和播放偏移时间。
	// :::
	Offset *float32 `json:"Offset,omitempty"`

	// 点播文件启播时间偏移值，单位为秒，数量与拉流地址列表中地址数量相等，缺省情况下为空表示不进行偏移。拉流来源类型为点播视频时，参数生效。 :::tip 此字段为旧版本配置，请使用 VodSrcAddrs 配置点播视频地址和播放偏移时间。 :::
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
	// :::tip 此字段为旧版本配置，请使用 VodSrcAddrs 配置点播视频地址和播放偏移时间。 :::
	SrcAddrS []*string `json:"SrcAddrS,omitempty"`

	// 推流的流名称，推流地址（DstAddr）为空时必传；反之，则该参数不生效。
	Stream *string `json:"Stream,omitempty"`

	// 拉流转推任务的名称，默认为空表示不配置任务名称。支持由中文、大小写字母（A - Z、a - z）和数字（0 - 9）组成，长度为 1 到 20 各字符。
	Title *string `json:"Title,omitempty"`

	// 点播文件地址和开始播放、结束播放的时间设置。 :::tip
	// * 当 Type 为点播类型时配置生效。
	// * 与 SrcAddrS 和 OffsetS 字段不可同时填写。 :::
	VodSrcAddrs []*CreatePullToPushTaskBodyVodSrcAddrsItem `json:"VodSrcAddrs,omitempty"`

	// 为拉流转推视频添加的水印配置信息。
	Watermark *CreatePullToPushTaskBodyWatermark `json:"Watermark,omitempty"`
}

type CreatePullToPushTaskBodyVodSrcAddrsItem struct {

	// REQUIRED; 点播文件地址。
	SrcAddr string `json:"SrcAddr"`

	// 当前点播文件结束播放的时间偏移值，单位为秒，默认为空时表示结束播放时间不进行偏移。
	EndOffset *float32 `json:"EndOffset,omitempty"`

	// 当前点播文件开始播放的时间偏移值，单位为秒。默认为空时表示开始播放时间不进行偏移。
	StartOffset *float32 `json:"StartOffset,omitempty"`
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

	// 应用名称，取值与直播流地址的 AppName 字段取值相同，由 1 到 30 位数字（0 - 9）、大写小字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，默认为空。 :::tip
	// * App 取值为空时，Stream 取值也需为空，表示录制配置为 Vhost 级别的全局配置。
	// * App 取值不为空时，Stream 取值含义请参见 Stream 参数说明。 :::
	App *string `json:"App,omitempty"`

	// 流名称，取值与直播流地址的 StreamName 字段取值相同，由 1 到 100 位数字（0 - 9）、大写小字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成。
	// :::tip
	// * App 取值不为空、Stream 取值为空时，表示录制配置为 Vhost + App 级别的配置。
	// * App 取值不为空、Stream 取值不为空时，表示录制为 Vhost + App + Stream 的配置。 :::
	Stream *string `json:"Stream,omitempty"`
}

// CreateRecordPresetV2BodyRecordPresetConfig - 直播流录制配置的详细配置。
type CreateRecordPresetV2BodyRecordPresetConfig struct {

	// 录制为 FLV 格式时的录制参数。 :::tip 您需至少配置一种录制格式，即 FlvParam、HlsParam、Mp4Param 至少开启一个。 :::
	FlvParam *CreateRecordPresetV2BodyRecordPresetConfigFlvParam `json:"FlvParam,omitempty"`

	// 录制为 HLS 格式时的录制参数。 :::tip 您需至少配置一种录制格式，即 FlvParam、HlsParam、Mp4Param 至少开启一个。 :::
	HlsParam *CreateRecordPresetV2BodyRecordPresetConfigHlsParam `json:"HlsParam,omitempty"`

	// 录制为 MP4 格式时的录制参数。 :::tip 您需至少配置一种录制格式，即 FlvParam、HlsParam、Mp4Param 至少开启一个。 :::
	Mp4Param *CreateRecordPresetV2BodyRecordPresetConfigMp4Param `json:"Mp4Param,omitempty"`

	// 是否录制源流，默认值为 0，支持的取值及含义如下所示。
	// * 0：不录制；
	// * 1：录制。
	// :::tip 转码流和源流需至少选一个进行录制，即 TranscodeRecord 和 OriginRecord 的取值不能同时为 0。 :::
	OriginRecord *int32 `json:"OriginRecord,omitempty"`

	// 录制为 HLS 格式时，单个 TS 切片时长，单位为秒，默认值为 10，取值范围为 [2,100]。
	SliceDuration *int32 `json:"SliceDuration,omitempty"`

	// 是否录制转码流，默认值为 0，支持的取值及含义如下所示。
	// * 0：不录制；
	// * 1：录制全部转码流；
	// * 2：录制指定转码流，根据转码后缀列表 TranscodeSuffixList决定录制哪些转码流。如果这个列表为空，则效果和设置为 1 一样，即录制所有转码流。
	// :::tip 转码流和源流需至少选一个进行录制，即 TranscodeRecord 和 OriginRecord 的取值不能同时为 0。 :::
	TranscodeRecord *int32 `json:"TranscodeRecord,omitempty"`

	// 转码后缀列表。
	TranscodeSuffixList []*string `json:"TranscodeSuffixList,omitempty"`
}

// CreateRecordPresetV2BodyRecordPresetConfigFlvParam - 录制为 FLV 格式时的录制参数。 :::tip 您需至少配置一种录制格式，即 FlvParam、HlsParam、Mp4Param
// 至少开启一个。 :::
type CreateRecordPresetV2BodyRecordPresetConfigFlvParam struct {

	// 实时录制场景下，断流等待时长，单位为秒，默认值为 180，取值范围为 [0,3600]。如果实际断流时间小于断流等待时长，录制任务不会停止；如果实际断流时间大于断流等待时长，录制任务会停止，断流恢复后重新开始一个新的录制任务。
	ContinueDuration *int32 `json:"ContinueDuration,omitempty"`

	// 断流录制场景下，单文件录制时长，单位为秒，默认值为 7200，取值范围为 -1 和 [300,86400]。
	// * 取值为 -1 时，表示不限制录制时长，录制结束后生成一个完整的录制文件。
	// * 取值为 [300,86400] 之间的值时，表示根据设置的录制文件时长，到达时长立即生成录制文件，完成录制后一起上传。
	// :::tip
	// * 断流录制场景仅在录制格式为 HLS 时生效，且断流录制和实时录制为二选一配置。
	// * 如录制过程中出现断流，对应生成的录制文件时长也会相应缩短。
	// :::
	Duration *int32 `json:"Duration,omitempty"`

	// 当前格式的录制是否开启，默认值为 false，支持的取值及含义如下所示。
	// * false：不开启；
	// * true：开启。
	Enable *bool `json:"Enable,omitempty"`

	// 实时录制场景下，单文件录制时长，单位为秒，默认值为 1800，取值范围为 [300,21600]。录制时间到达设置的单文件录制时长时，会立即生成录制文件实时上传存储。 :::tip 如录制过程中出现断流，对应生成的录制文件时长也会相应缩短。
	// :::
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

	// 录制文件存储到 TOS 时的存储路径和文件名规则。支持输入字母（A - Z、a - z）、数字（0 - 9）、短横线（-）、叹号（!）、下划线（_）、句点（.）、星号（*）及占位符。最大长度为 200 个字符，
	// 支持以下字段作为占位符：
	// * record：自定义字段，可遵照支持字符进行自定义。
	// * {PubDomain}：当前配置中的 vhost 值。
	// * {App}：当前配置中的 AppName 值。
	// * {Stream}：当前配置中的 StreamName 值。
	// * {StartTime}：录制开始的 Unix 时间戳，精度为 s。
	// * {EndTime}：录制结束的 Unix 时间戳，精度为 s。
	// 存储路径必须至少包含两级目录。例如：live/{App}/{Stream}
	// 合法示例：
	// record/{PubDomain}/{App}/{Stream}/{StartTime}-{EndTime}
	// {App}/archive/{Stream}/recording_{StartTime}
	// vod/{Stream}/!highlight_{EndTime}
	// a/b/custom_record
	// 错误示例：
	// single_level # 错误：路径层级不足两级
	// invalid_/{S@ream}/file # 错误：含非法字符@
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

// CreateRecordPresetV2BodyRecordPresetConfigHlsParam - 录制为 HLS 格式时的录制参数。 :::tip 您需至少配置一种录制格式，即 FlvParam、HlsParam、Mp4Param
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
	Result *CreateRecordPresetV2ResResult `json:"Result,omitempty"`
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

// CreateRecordPresetV2ResResult - 视请求的接口而定
type CreateRecordPresetV2ResResult struct {

	// REQUIRED; 录制配置 ID。
	PresetName string `json:"PresetName"`
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
	Result           *CreateSnapshotPresetResResult          `json:"Result,omitempty"`
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

type CreateSnapshotPresetResResult struct {

	// REQUIRED
	PresetName string `json:"PresetName"`
}

type CreateSubtitleTranscodePresetBody struct {

	// REQUIRED; 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 源语言参数
	SourceLanguage CreateSubtitleTranscodePresetBodySourceLanguage `json:"SourceLanguage"`

	// REQUIRED; 关联的转码配置后缀，一个字幕配置支持关联多个转码配置后缀。
	Suffixes []string `json:"Suffixes"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 字幕配置的描述信息。
	Description *string `json:"Description,omitempty"`

	// 预设配置，使用预设配置是系统将自动对字体大小、字幕行数、每行最大字符数和边距参数（MarginVertical 和 MarginHorizontal）进行智能化适配。默认为空，表示不使用预设配置，支持的预设配置如下所示。
	// * small ：小字幕。
	// * medium：中字幕。
	// * large：大字幕。 :::tip 使用预设配置时，字幕行数、每行最大字符数、左右边距和底部边距参数不生效，系统将使用预设配置自动进行计算。 :::
	DisplayPreset *string `json:"DisplayPreset,omitempty"`

	// 原文翻译成译文时使用的热词词库，总长度不超过 10000 个字符，默认为空。
	GlossaryWordList []*string `json:"GlossaryWordList,omitempty"`

	// 原文字幕识别时使用的热词词库，总长度不超过为 10000 个字符，默认为空。
	HotWordList []*string `json:"HotWordList,omitempty"`

	// 设置在 16:9 分辨率场景下，每行字幕展示的最大字符数。 :::tip
	// * 使用预设配置时，字幕每行最大字符数设置不生效。
	// * 不使用预设配置时，字幕每行最大字符数必填。
	// * 每个文字、字母、符号或数字均为一个字符。
	// * 当屏幕分辨率改变时，屏幕上显示的每行文字数量会相应调整，以适应新的分辨率，确保文字的显示效果和阅读体验。 :::
	MaxCharNumber *int32 `json:"MaxCharNumber,omitempty"`

	// 字幕展示的行数，同时适用于原文字幕和译文字幕，支持的取值及含义如下所示。
	// * 0：（默认值）根据字幕字数自动进行分行展示；
	// * 1：每种字幕展示一行；
	// * 2：每种字幕展示两行。 :::tip
	// * 使用预设配置时，字幕行数为自动分行展示。
	// * 超出行内字数限制时表示字幕将超过显示范围，此时字幕内容将被截断。 :::
	MaxRowNumber *int32 `json:"MaxRowNumber,omitempty"`

	// 字幕位置设置，通过设置字幕距离画面左右边距和底部边距来指定字幕位置。
	// :::tip
	// * 使用预设配置时，字幕位置设置不生效。
	// * 不使用预设配置时，字幕位置设置必填。 :::
	Position *CreateSubtitleTranscodePresetBodyPosition `json:"Position,omitempty"`

	// 字幕配置的名称，不可以与其他已有的配置名称重复。默认为空，表示由系统将自动分配配置名称。
	PresetName *string `json:"PresetName,omitempty"`

	// 译文字幕展示参数配置列表，当前最多支持配置一种译文。
	TargetLanguage []*CreateSubtitleTranscodePresetBodyTargetLanguageItem `json:"TargetLanguage,omitempty"`
}

// CreateSubtitleTranscodePresetBodyPosition - 字幕位置设置，通过设置字幕距离画面左右边距和底部边距来指定字幕位置。
// :::tip
// * 使用预设配置时，字幕位置设置不生效。
// * 不使用预设配置时，字幕位置设置必填。 :::
type CreateSubtitleTranscodePresetBodyPosition struct {

	// 左右边距，[0,0.2]
	MarginHorizontal *float32 `json:"MarginHorizontal,omitempty"`

	// 上下边距
	MarginVertical *float32 `json:"MarginVertical,omitempty"`

	// MarginVertical是相对顶部或底部,默认底部
	Relative *CreateSubtitleTranscodePresetBodyPositionRelative `json:"Relative,omitempty"`
}

// CreateSubtitleTranscodePresetBodySourceLanguage - 源语言参数
type CreateSubtitleTranscodePresetBodySourceLanguage struct {

	// REQUIRED; 是否展示源语言
	Display bool `json:"Display"`

	// REQUIRED; 字体
	Font string `json:"Font"`

	// REQUIRED; 字体颜色
	FontColor string `json:"FontColor"`

	// REQUIRED; 原文字幕的语言，取值及含义如下所示。
	// * zh：中英混合；
	// * en：英语；
	// * ko：韩语；
	// * ja：日语。
	Language string `json:"Language"`

	// 字幕阴影配置
	Border *CreateSubtitleTranscodePresetBodySourceLanguageBorder `json:"Border,omitempty"`
}

// CreateSubtitleTranscodePresetBodySourceLanguageBorder - 字幕阴影配置
type CreateSubtitleTranscodePresetBodySourceLanguageBorder struct {

	// REQUIRED; 填0的时候后端根据字体大小进行计算，字体大小/32*1.25
	Width float32 `json:"Width"`

	// 阴影颜色
	Color *string `json:"Color,omitempty"`
}

type CreateSubtitleTranscodePresetBodyTargetLanguageItem struct {

	// REQUIRED; 字体，覆盖全局参数
	Font string `json:"Font"`

	// REQUIRED; 字体颜色，覆盖全局参数
	FontColor string `json:"FontColor"`

	// REQUIRED; 译文字幕的语言，取值及含义如下所示。
	// * zh：中英混合；
	// * zh-Hant：繁体中文；
	// * en：英语；
	// * ko：韩语；
	// * ja：日语；
	// * ar：阿拉伯语；
	// * de：德语；
	// * es：西班牙语；
	// * fr：法语；
	// * hi：印地语；
	// * pt：葡萄牙语；
	// * ru：俄语；
	// * vi：越南语；
	// * th：泰语。
	Language string `json:"Language"`

	// 字幕阴影配置
	Border *CreateSubtitleTranscodePresetBodyTargetLanguageItemBorder `json:"Border,omitempty"`
}

// CreateSubtitleTranscodePresetBodyTargetLanguageItemBorder - 字幕阴影配置
type CreateSubtitleTranscodePresetBodyTargetLanguageItemBorder struct {

	// REQUIRED; 填0的时候后端根据字体大小进行计算，字体大小/32*1.25
	Width float32 `json:"Width"`

	// 阴影颜色
	Color *string `json:"Color,omitempty"`
}

type CreateSubtitleTranscodePresetRes struct {

	// REQUIRED
	ResponseMetadata CreateSubtitleTranscodePresetResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type CreateSubtitleTranscodePresetResResponseMetadata struct {

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
	App string `json:"App"`

	// REQUIRED
	MaxShiftTime int32 `json:"MaxShiftTime"`

	// REQUIRED
	PullDomain string `json:"PullDomain"`

	// REQUIRED
	Type string `json:"Type"`

	// REQUIRED
	Vhost        string  `json:"Vhost"`
	Bucket       *string `json:"Bucket,omitempty"`
	MasterFormat *string `json:"MasterFormat,omitempty"`
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

type CreateTranscodePresetBatchBody struct {

	// REQUIRED; 子流转码配置的详细参数。
	PresetList []CreateTranscodePresetBatchBodyPresetListItem `json:"PresetList"`

	// REQUIRED; 枚举值 create associate hls-abr dash-abr abr
	Type string `json:"Type"`

	// 子m3u8是否复制主m3u8参数
	CopyParams *bool `json:"CopyParams,omitempty"`
}

type CreateTranscodePresetBatchBodyPresetListItem struct {

	// REQUIRED; 直播流地址的 AppName 字段，支持由大小写字母（A - Z、a - z）、数字（0 - 9）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 配置参数。
	GOP int32 `json:"GOP"`

	// REQUIRED; 高度，取值范围为 [ ]，单位为，默认值为``。
	Height int32 `json:"Height"`

	// REQUIRED; 转码配置的后缀标识，支持由大小写字母（A - Z、a - z）、数字（0 - 9）和短横线（-）组成，长度为 1 到 10 个字符。 转码后缀通常以流名称后缀的形式来使用，常见的标识有sd、hd、uhd，例如，当转码配置的标识为hd时，拉取转码流时转码流的流名名称为源流的流名称_hd。
	SuffixName string `json:"SuffixName"`

	// REQUIRED; 视频编码格式。
	// * h264：使用 H.264 编码格式；
	// * h265：使用 H.265 编码格式；
	// :::tip 同一条流所有的 ABR 子流的视频编码格式需配置相同。 :::
	Vcodec string `json:"Vcodec"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看需要转码的直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// REQUIRED; 视频比特率。取值范围为 [ ]，单位为，默认值为``。
	VideoBitrate int32 `json:"VideoBitrate"`

	// REQUIRED; 宽度，取值范围为 [ ]，单位为，默认值为``。
	Width int32 `json:"Width"`

	// 音频编码格式。默认格式为 acc，支持以下 3 种类型。
	// * aac：使用 aac 编码格式；
	// * copy：不进行转码，所有音频编码参数不生效；
	// * opus：使用 opus 编码格式。
	Acodec *string `json:"Acodec,omitempty"`
	As     *string `json:"As,omitempty"`

	// 音频码率，单位为 kbps。
	AudioBitrate *int32 `json:"AudioBitrate,omitempty"`
	BFrames      *int32 `json:"BFrames,omitempty"`

	// 视频帧率，单位为 fps，取值范围为 [0,60]，默认为 25fps。帧率越大，画面越流畅。
	FPS      *int32 `json:"FPS,omitempty"`
	LongSide *int32 `json:"LongSide,omitempty"`

	// 转码类型是否为极智超清转码，默认值为 false，取值及含义如下所示。
	// * true：极智超清；
	// * false：标准转码。
	// :::tip 同一条流所有的 ABR 子流的转码方式需配置相同。 :::
	Roi       *bool  `json:"Roi,omitempty"`
	ShortSide *int32 `json:"ShortSide,omitempty"`
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

	// REQUIRED; 应用名称，取值与直播流地址的 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 转码后缀，支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）和短横线（-）组成，长度为 1 到 10 个字符。
	// 转码后缀通常以流名称后缀的形式来使用，常见的标识有 _sd、_hd、_uhd，例如，当转码配置的标识为 _hd 时，拉取转码流时转码流的流名名称为 源流的流名称_hd。
	SuffixName string `json:"SuffixName"`

	// REQUIRED; 视频编码格式，支持的取值及含义如下所示。
	// * h264：使用 H.264 的视频编码格式；
	// * h265：使用 H.265 的视频编码格式；
	// * h266：使用 H.266 的视频编码格式；
	// * copy：不进行视频转码，所有视频编码参数不生效，视频编码参数包括视频帧率（FPS）、视频码率（VideoBitrate）、分辨率设置（As、Width、Height、ShortSide、LongSide）、GOP 和 BFrames
	// 等。
	Vcodec string `json:"Vcodec"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，查看需要转码的直播流使用的域名所属的域名空间。
	Vhost   string `json:"Vhost"`
	AbrMode *int32 `json:"AbrMode,omitempty"`

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

	// 音频码率，单位为 kbps，默认值为 128，取值范围为 [0,1000]；取值为 0 时，表示与源流的音频码率相同。
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

	// 转码输出视频中 2 个参考帧之间的最大 B 帧数量，默认值为 3，取值为 0 时表示去除 B 帧。
	// 最大 B 帧数量的取值范围根据视频编码格式（Vcodec）的不同有所差异，取值范围如下所示。
	// * 视频编码格式为 H.264 （Vcodec 取值为 h264）时取值范围为 [0,7]；
	// * 视频编码格式为 H.265 或 H.266 （Vcodec 取值为 h265 或 h266）时取值范围为 [0,3]、7、15。
	BFrames *int32 `json:"BFrames,omitempty"`

	// 视频帧率，单位为 fps，默认值为 25，取值为 0 时表示与源流视频帧率相同。
	// 视频帧率的取值范围根据视频编码格式（Vcodec）的不同有所差异，视频码率的取值范围如下所示。
	// * 视频编码格式为 H.264 或 H.265 （Vcodec 取值为 h264 或 h265）时，视频帧率取值范围为 [0,60]；
	// * 视频编码格式为 H.266 （Vcodec 取值为 h266）时，视频帧率取值范围为 [0,35]。
	FPS *int32 `json:"FPS,omitempty"`

	// IDR 帧之间的最大间隔时间，单位为秒，取值范围为 [1,30]。
	GOP *int32 `json:"GOP,omitempty"`

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
	LongSide *int32 `json:"LongSide,omitempty"`

	// 转码类型是否为极智超清转码，默认值为 false，取值及含义如下。
	// * true：极智超清转码；
	// * false：标准转码。
	// :::tip 视频编码格式为 H.266 （Vcodec取值为h266）时，转码类型不支持极智超清转码。 :::
	Roi *bool `json:"Roi,omitempty"`

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

	// 转码停止时长，支持触发方式为拉流转码（TransType 取值为 Pull）时设置，表示断开拉流后转码停止的时长，单位为秒，取值范围为 -1 和 [0,300]，-1 表示不停止转码，默认值为 60。
	StopInterval *int32 `json:"StopInterval,omitempty"`

	// 转码触发方式，默认值为 Pull，支持的取值及含义如下。
	// * Push：推流转码，直播推流后会自动启动转码任务，生成转码流；
	// * Pull：拉流转码，直播推流后，需要主动播放转码流才会启动转码任务，生成转码流。
	TransType *string `json:"TransType,omitempty"`

	// 视频码率，单位为 bps，默认值为 1000000；取值为 0 时，表示与源流的视频码率相同。
	// 视频码率的取值范围根据视频编码格式（Vcodec）的不同有所差异，视频码率的取值范围如下所示。
	// * 视频编码格式为 H.264 或 H.265 （Vcodec 取值为 h264 或 h265）时，视频码率取值范围为 [0,30000000]；
	// * 视频编码格式为 H.266 （Vcodec 取值为 h266）时，视频码率取值范围为 [0,6000000]。
	VideoBitrate *int32 `json:"VideoBitrate,omitempty"`

	// 视频宽度，单位为 px，默认值为 0。
	// 视频宽度的取值范围根据视频编码格式（Vcodec）的不同所有差异，视频宽度取值如下所示。
	// * 视频编码格式为 H.264 或 H.265 （Vcodec 取值为 h264 或 h265）时，取值范围为 [150,1920]；
	// * 视频编码格式为 H.266 （Vcodec 取值为 h266）时，不支持设置 Width 和 Height。
	// :::tip
	// * 当关闭视频分辨率自适应（As 取值为 0）时，转码分辨率将取 Width 和 Height 的值对转码视频进行拉伸；
	// * 当关闭视频分辨率自适应（As 取值为 0）时，Width 和 Height 任一取值为 0 时，转码视频将保持源流尺寸。 :::
	Width *int32 `json:"Width,omitempty"`
}

type CreateTranscodePresetRes struct {

	// REQUIRED
	ResponseMetadata CreateTranscodePresetResResponseMetadata `json:"ResponseMetadata"`
	Result           *CreateTranscodePresetResResult          `json:"Result,omitempty"`
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

type CreateTranscodePresetResResult struct {

	// REQUIRED; 模板名称
	PresetName string `json:"PresetName"`
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

	// 流名称，取值与直播流地址中 StreamName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	// :::tip
	// * 默认为空，表示对指定的 AppName 下所有转码流均使用当前水印配置。
	// * 指定流名称时，表示仅对 AppName 下指定流名称的转码流使用当前水印配置。 :::
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

type DeleteCarouselTaskBody struct {

	// REQUIRED; 待删除的轮播任务 ID，任务的唯一标识。调用 CreateCarouselTask 接口创建轮播任务时返回。
	TaskID string `json:"TaskID"`
}

type DeleteCarouselTaskRes struct {

	// REQUIRED
	ResponseMetadata DeleteCarouselTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DeleteCarouselTaskResResult `json:"Result"`
}

type DeleteCarouselTaskResResponseMetadata struct {

	// REQUIRED
	RequestID string `json:"RequestID"`
}

type DeleteCarouselTaskResResult struct {

	// REQUIRED; 删除任务的操作结果信息，返回任务是否成功删除以及相关的 Mesos ID 和操作影响记录数。
	Data string `json:"Data"`
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

type DeleteCloudMixTaskBody struct {

	// REQUIRED; 混流任务 ID，您可以通过 ListCloudMixTask [https://www.volcengine.com/docs/6469/1271157] 接口获取待结束的混流任务 ID。
	TaskID string `json:"TaskID"`
}

type DeleteCloudMixTaskRes struct {

	// REQUIRED
	ResponseMetadata DeleteCloudMixTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DeleteCloudMixTaskResResult `json:"Result"`
}

type DeleteCloudMixTaskResResponseMetadata struct {

	// REQUIRED
	RequestID string `json:"RequestID"`
}

type DeleteCloudMixTaskResResult struct {

	// REQUIRED; 请求响应码，取值及含义如下。
	// * 0：请求成功；
	// * 500：内部处理错误；
	// * 400：请求异常。
	Code int32 `json:"Code"`

	// REQUIRED; 返回的数据。
	Data string `json:"Data"`

	// REQUIRED; 请求响应码对应的信息。
	Message string `json:"Message"`
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

type DeleteLiveVideoQualityAnalysisTaskBody struct {

	// 任务ID，和任务名二选一
	ID *int64 `json:"ID,omitempty"`

	// 任务名，和任务ID二选一
	Name *string `json:"Name,omitempty"`
}

type DeleteLiveVideoQualityAnalysisTaskRes struct {

	// REQUIRED
	ResponseMetadata DeleteLiveVideoQualityAnalysisTaskResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteLiveVideoQualityAnalysisTaskResResponseMetadata struct {

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

type DeletePullToPushGroupBody struct {

	// REQUIRED; 拉流转推群组名称，您可以调用 ListPullToPushGroup [https://www.volcengine.com/docs/6469/1327382] 接口获取群组名称。
	Name string `json:"Name"`
}

type DeletePullToPushGroupRes struct {

	// REQUIRED
	ResponseMetadata DeletePullToPushGroupResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeletePullToPushGroupResResponseMetadata struct {

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

	// 任务所属的群组名称，您可以通过获取拉流转推任务列表 [https://www.volcengine.com/docs/6469/1126896]接口获取。 :::tip
	// * 使用主账号调用时，为非必填。
	// * 使用子账号调用时，为必填。 :::
	GroupName *string `json:"GroupName,omitempty"`
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

type DeleteRemoteAuthBody struct {

	// REQUIRED; 拉流域名，您可以调用ListVhostRemoteAuth [https://www.volcengine.com/docs/6469/1250148]接口查看远程鉴权配置的 Domain 取值。
	Domain string `json:"Domain"`

	// REQUIRED; 域名空间，您可以调用ListVhostRemoteAuth [https://www.volcengine.com/docs/6469/1250148]接口查看远程鉴权配置的 Vhost 取值。
	Vhost string  `json:"Vhost"`
	App   *string `json:"App,omitempty"`
}

type DeleteRemoteAuthRes struct {

	// REQUIRED
	ResponseMetadata DeleteRemoteAuthResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *string `json:"Result,omitempty"`
}

type DeleteRemoteAuthResResponseMetadata struct {

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

type DeleteSubtitleTranscodePresetBody struct {

	// REQUIRED; 火山必填
	App string `json:"App"`

	// REQUIRED; 截图配置的名称，您可以调用 ListVhostSubtitleTranscodePreset [https://www.volcengine.com/docs/6469/1288712] 接口，获取待删除字幕配置的 PresetName
	// 取值。
	PresetName string `json:"PresetName"`

	// REQUIRED; 火山必填
	Vhost string `json:"Vhost"`
}

type DeleteSubtitleTranscodePresetRes struct {

	// REQUIRED
	ResponseMetadata DeleteSubtitleTranscodePresetResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type DeleteSubtitleTranscodePresetResResponseMetadata struct {

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

type DeleteTimeShiftPresetV2Body struct {

	// REQUIRED
	App string `json:"App"`

	// REQUIRED
	Preset string `json:"Preset"`

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名所属的域名空间。
	Vhost string  `json:"Vhost"`
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

type DeleteTranscodePresetBatchBody struct {

	// REQUIRED; 删除模版的信息
	PresetList []DeleteTranscodePresetBatchBodyPresetListItem `json:"PresetList"`

	// REQUIRED; associate create hls-abr
	Type string `json:"Type"`
}

type DeleteTranscodePresetBatchBodyPresetListItem struct {

	// REQUIRED; 解绑的app
	App string `json:"App"`

	// REQUIRED; 模版名
	Preset string `json:"Preset"`

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

	// 流名称，您可以调用ListVhostWatermarkPreset [https://www.volcengine.com/docs/6469/1126889]接口，查看待删除水印配置的 Stream 取值。
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

	// REQUIRED; 时间戳进制。取值如下：
	// * 2：二进制
	// * 8：八进制
	// * 10：十进制
	// * 16：十六进制
	// :::tipSceneType 取值为 push 时，该参数取值固定为 10。 :::
	TimeStampBase int32 `json:"TimeStampBase"`

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

	// 鉴权模版类型
	AuthType *string `json:"AuthType,omitempty"`

	// 生成加密字符串使用的加密字段。
	EncryptField []*string `json:"EncryptField,omitempty"`

	// 对称加密算法。
	EncryptionAlgorithm *string `json:"EncryptionAlgorithm,omitempty"`

	// 自定义鉴权密钥。
	SecretKey *string `json:"SecretKey,omitempty"`
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

	// REQUIRED
	CreateTime string `json:"CreateTime"`

	// REQUIRED
	UpdateTime        string   `json:"UpdateTime"`
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

	// 证书 ID，您可以通过ListCertV2 [https://www.volcengine.com/docs/6469/1126823]接口获取证书 ID。 :::tip 参数ChainID与CertID传且仅传一个。 :::
	CertID *string `json:"CertID,omitempty"`

	// 证书链 ID，您可以通过ListcCertV2 [https://www.volcengine.com/docs/6469/1126823]接口获取 证书链 ID。 :::tip 参数ChainID与CertID传且仅传一个。 :::
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

	// REQUIRED; 是否开启 HTTP/2 协议。取值如下：
	// * false: 关闭 HTTP/2 协议。
	// * true: 开启 HTTP/2 协议。
	HTTP2 bool `json:"HTTP2"`

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
	// * 5：欠费关停；
	// * 6：域名未备案被封禁。
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

type DescribeEncryptHLSRes struct {

	// REQUIRED
	ResponseMetadata DescribeEncryptHLSResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeEncryptHLSResResult `json:"Result,omitempty"`
}

type DescribeEncryptHLSResResponseMetadata struct {

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

// DescribeEncryptHLSResResult - 视请求的接口而定
type DescribeEncryptHLSResResult struct {

	// REQUIRED; 视频直播服务端生成密钥的更新周期，单位为秒，取值范围为 [60,604800]。
	CycleTime float32 `json:"CycleTime"`

	// REQUIRED; 客户自建密钥管理服务后，客户端向密钥管理服务请求获取密钥的地址。
	URL string `json:"URL"`
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
	InsertPDT bool `json:"InsertPDT"`

	// REQUIRED
	Interval float32 `json:"Interval"`

	// REQUIRED
	PDTInterval string `json:"PDTInterval"`

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

	// REQUIRED; 0: response 1: request
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
	CreateTime string `json:"CreateTime"`

	// REQUIRED
	Domain string `json:"Domain"`

	// REQUIRED; 单位ms
	GopCacheSize string `json:"GopCacheSize"`

	// REQUIRED
	UpdateTime string `json:"UpdateTime"`

	// REQUIRED
	Vhost string `json:"Vhost"`
}

type DescribeLicenseDRMQuery struct {

	// REQUIRED; 应用名称，取值与直播流地址的 AppName 字段相同，由大写小字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App" query:"App"`

	// REQUIRED; DRM 加密的类型，取值及含义如下所示。
	// * fp：FairPlay 加密；
	// * wv：Widevine 加密；
	// * pr：PlayReady 加密。
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

	// REQUIRED; 查询的结束时间。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的起始时间。
	StartTime string `json:"StartTime"`

	// AI字幕原始语言，枚举值：ZH(中文)，EN(英文), JA(日文), KO(韩文)
	ASRSourceType *string `json:"ASRSourceType,omitempty"`

	// AI字幕转换的语言类型，ASRTargetTypeList和ASRSourceType必须同时指定， 枚举值：ZH(中文)，EN(英文), JA(日文), KO(韩文)
	ASRTargetTypeList []*string `json:"ASRTargetTypeList,omitempty"`

	// 聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 93 天；
	// * 86400：1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询时间范围为 93 天。
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

	// REQUIRED; 聚合的时间粒度
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; AI字幕总时长总量，unit is minute
	Duration float32 `json:"Duration"`

	// REQUIRED; 查询的结束时间。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的起始时间。
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

	// REQUIRED; 时间点
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询最大时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询最大时间范围为 366 天；
	// * 86400：1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询最大时间范围为 366 天。
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
	// * huashu：华数。
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// byteplus比火山多了CMAF协议
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询时间范围内的下行峰值带宽，单位为 Mbps。
	PeakDownBandwidth float32 `json:"PeakDownBandwidth"`

	// REQUIRED; 查询时间范围内的上行峰值带宽，单位为 Mbps。
	PeakUpBandwidth float32 `json:"PeakUpBandwidth"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
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
	// * huashu：华数。
	ISPList []*string `json:"ISPList,omitempty"`

	// byteplus比火山多了CMAF协议
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域列表。
	RegionList []*DescribeLiveBandwidthDataResResultRegionListItem `json:"RegionList,omitempty"`

	// 客户端 IP 所属区域列表。
	UserRegionList []*DescribeLiveBandwidthDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveBandwidthDataResResultBandwidthDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的下行峰值带宽，单位为 Mbps。
	DownBandwidth float32 `json:"DownBandwidth"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
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

type DescribeLiveBatchPushStreamMetricsBody struct {

	// REQUIRED; 推流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的推流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 请控制查询的数据量，如果查询速度较慢请缩短查询时间范围
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	// :::tip 单次查询最大时间跨度为 1 天，历史查询最大时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 指标聚合算法，支持max:峰值聚合，avg：平均值，默认max
	AggType *string `json:"AggType,omitempty"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 5：5 秒；
	// * 30：30 秒；
	// * 60：（默认值）1 分钟。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用名称，取值与直播流地址中的 AppName 字段取值相同，查询流粒度数据时必传，且需同时传入 Stream。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30
	// 个字符。 :::tip 查询流粒度的监控数据时，需同时指定 App 和 Stream 来指定直播流。 :::
	App *string `json:"App,omitempty"`

	// 流名称，预置与直播流地址中的 StreamName 字段取值相同，查询流粒度数据时必传，且需同时传入 Stream。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到
	// 100 个字符。 :::tip 查询流粒度的监控数据时，需同时指定 App 和 Stream 来指定直播流。 :::
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 直推流的信息，包含域名、应用名称、流名称和监控数据。
	StreamMetricList []DescribeLiveBatchPushStreamMetricsResResultStreamMetricListItem `json:"StreamMetricList"`

	// 数据聚合时间粒度内，动态指标的聚合算法，取值及含义如下所示。
	// * max：（默认值）计算聚合时间粒度内的最大值；
	// * avg：计算聚合时间粒度内的平均值。
	AggType *string `json:"AggType,omitempty"`

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

	// REQUIRED; 标记一路推流的唯一id
	SessionID string `json:"SessionID"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`
}

type DescribeLiveBatchPushStreamMetricsResResultStreamMetricListPropertiesItemsItem struct {

	// REQUIRED; 音频编码格式。 Eg. ACC
	ACodec string `json:"ACodec"`

	// REQUIRED; 数据聚合时间粒度内，按聚合算法得出的音频码率，单位为 kbps。
	AudioBitrate float32 `json:"AudioBitrate"`

	// REQUIRED; 数据聚合时间粒度内，按聚合算法得出的相邻音频帧显示时间戳差值，单位为毫秒。
	AudioFrameGap float32 `json:"AudioFrameGap"`

	// REQUIRED; 数据聚合时间粒度内，按聚合算法得出的音频帧率（每秒传输的音频数据包个数）。
	AudioFramerate float32 `json:"AudioFramerate"`

	// REQUIRED; 数据聚合时间粒度内，最后一个音频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	AudioPts float64 `json:"AudioPts"`

	// REQUIRED; 数据聚合时间粒度内，按聚合算法得出的视频码率，单位为 kbps。
	Bitrate float32 `json:"Bitrate"`

	// REQUIRED; 客户端ip
	ClientIP string `json:"ClientIp"`

	// REQUIRED; 收到首帧的时间，，单位毫秒
	FirstFrameTime int32 `json:"FirstFrameTime"`

	// REQUIRED; 数据聚合时间粒度内，按聚合算法得出的视频帧率，单位为 fps。
	Framerate float32 `json:"Framerate"`

	// REQUIRED; 数据聚合时间粒度内，按聚合算法得出的音视频帧显示时间戳差值，即所有 AudioPts 与 VideoPts 差值的最大值，单位为毫秒。
	PtsDelta float32 `json:"PtsDelta"`

	// REQUIRED; 分辨率
	Resolution string `json:"Resolution"`

	// REQUIRED; 服务端ip
	ServerIP string `json:"ServerIp"`

	// REQUIRED; 推流开始时间，单位毫秒
	StreamBeginTime int64 `json:"StreamBeginTime"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 视频编码格式。 Eg. H264
	VCodec string `json:"VCodec"`

	// REQUIRED; 数据聚合时间粒度内，按聚合算法得出的相邻视频帧显示时间戳差值，单位为毫秒。
	VideoFrameGap float32 `json:"VideoFrameGap"`

	// REQUIRED; 数据聚合时间粒度内，最后一个视频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	VideoPts float64 `json:"VideoPts"`
}

type DescribeLiveBatchStreamTranscodeDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。 :::tip 查询历史数据的时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 域名列表，默认为空表示全部域名。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取直播流使用的域名信息。
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 数据分页的信息。
	Pagination DescribeLiveBatchStreamTranscodeDataResResultPagination `json:"Pagination"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
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

	// REQUIRED; 当前流的转码码率，单位为 kbps。
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

type DescribeLiveCallbackDataBody struct {

	// REQUIRED; 结束时间，单次查询31天，历史366天
	EndTime string `json:"EndTime"`

	// REQUIRED; 开始时间
	StartTime string `json:"StartTime"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 回调事件类型 推流开始-push_start 推流结束-push_end 截图回调-snapshot_event 录制回调-record_event
	CallbackEventType []*string `json:"CallbackEventType,omitempty"`

	// 回调状态 成功-success 失败-fail
	CallbackStatus []*string `json:"CallbackStatus,omitempty"`

	// 需查询的域名列表，缺省情况下表示当前账号下的所有域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 分页页码，默认是1，取值范围[1，10000]
	PageNum *int32 `json:"PageNum,omitempty"`

	// 每页的大小，默认100，取值范围[1, 1000]
	PageSize *int32 `json:"PageSize,omitempty"`

	// 流名称，用于精确定位某一路直播流。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveCallbackDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveCallbackDataResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *DescribeLiveCallbackDataResResult `json:"Result,omitempty"`
}

type DescribeLiveCallbackDataResResponseMetadata struct {

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

// DescribeLiveCallbackDataResResult - 视请求的接口而定
type DescribeLiveCallbackDataResResult struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的 UTC 时间，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的 UTC 时间，精度为秒。
	StartTime string `json:"StartTime"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 回调事件类型
	CallbackEventType []*string `json:"CallbackEventType,omitempty"`

	// 回调事件详情
	CallbackInfoList []*DescribeLiveCallbackDataResResultCallbackInfoListItem `json:"CallbackInfoList,omitempty"`

	// 回调状态
	CallbackStatus []*string `json:"CallbackStatus,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 查询结果的分页信息。
	Pagination *DescribeLiveCallbackDataResResultPagination `json:"Pagination,omitempty"`

	// 流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveCallbackDataResResultCallbackInfoListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 回调地址。
	CallbackAddress string `json:"CallbackAddress"`

	// REQUIRED; 回调请求体。
	CallbackBody string `json:"CallbackBody"`

	// REQUIRED; 回调错误说明。
	CallbackErrorCode string `json:"CallbackErrorCode"`

	// REQUIRED; 错误信息，当回调失败时返回。
	CallbackErrorMessage string `json:"CallbackErrorMessage"`

	// REQUIRED; 回调事件类型。
	CallbackEventType string `json:"CallbackEventType"`

	// REQUIRED; 回调请求方式。
	CallbackMethod string `json:"CallbackMethod"`

	// REQUIRED; 回调响应体。
	CallbackResponseBody string `json:"CallbackResponseBody"`

	// REQUIRED; 回调响应码。
	CallbackResponseCode string `json:"CallbackResponseCode"`

	// REQUIRED; 回调响应头信息。
	CallbackResponseHeader string `json:"CallbackResponseHeader"`

	// REQUIRED; 回调响应时间。
	CallbackResponseTime string `json:"CallbackResponseTime"`

	// REQUIRED; 回调状态
	CallbackStatus string `json:"CallbackStatus"`

	// REQUIRED; 回调发生时间。
	CallbackTime string `json:"CallbackTime"`

	// REQUIRED; 域名。
	Domain string `json:"Domain"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`
}

// DescribeLiveCallbackDataResResultPagination - 查询结果的分页信息。
type DescribeLiveCallbackDataResResultPagination struct {

	// REQUIRED; 当前所在分页的页码。
	PageNum int32 `json:"PageNum"`

	// REQUIRED; 每页显示的数据条数。
	PageSize int32 `json:"PageSize"`

	// REQUIRED; 查询结果的数据总条数。
	TotalCount int32 `json:"TotalCount"`
}

type DescribeLiveEdgeStatDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。 :::tip 历史查询最大时间范围为 366 天，单次查询最大时间跨度与数据拆分维度和数据聚合时间粒度有关，详细如下。
	// * 当不进行维度拆分或只使用一个维度拆分数据时： * 数据以 60 秒聚合时，单次查询最大时间跨度为 24 小时；
	// * 数据以 300 秒聚合时，单次查询最大时间跨度为 31 天；
	// * 数据以 3600 秒聚合时，单次查询最大时间跨度为 31 天。
	//
	//
	// * 当使用两个或两个以上维度拆分数据时： * 数据以 60 秒聚合时，单次查询最大时间跨度为 3 小时；
	// * 数据以 300 秒聚合时，单次查询最大时间跨度为 24 小时；
	// * 数据以 3600 秒聚合时，单次查询最大时间跨度为 7 天。 :::
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持以下取值。
	// * 60：1 分钟。
	// * 300：（默认值）5 分钟。
	// * 3600：1 小时。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用程序列表。
	AppList []*string `json:"AppList,omitempty"`

	// 数据拆分的维度，缺省情况下不进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * Protocol：推拉流协议；
	// * ISP：运营商；
	// * Area：CDN 节点 IP 所属大区。
	// * Country：CDN 节点 IP 所属国家。
	// * Province：CDN 节点 IP 所属省份。
	// * UserArea：客户端 IP 所属大区。
	// * UserCountry：客户端 IP 所属国家。
	// * UserProvince：客户端 IP 所属省份。
	// :::tip 中国（Country 或 UserCountry 为 CN）以外区域无 Province 字段，如果按 Province 或 UserProvince 字段拆分数据时，默认只返回 Country 为 CN 时的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，默认为空，表示您添加到视频直播中的所有域名。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的域名。
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
	// * huashu：华数。
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// byteplus比火山多了CMAF协议
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数RegionList和UserRegionList不支持同时传入。 :::
	RegionList []*DescribeLiveEdgeStatDataBodyRegionListItem `json:"RegionList,omitempty"`

	// 流列表。
	StreamList []*string `json:"StreamList,omitempty"`

	// 客户端 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数RegionList和UserRegionList不支持同时传入。 :::
	UserRegionList []*DescribeLiveEdgeStatDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveEdgeStatDataBodyRegionListItem struct {

	// 区域信息中的大区标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按国家筛选，需要同时传入Area和Country。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符，国外暂不支持该参数，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按省筛选，需要同时传入Area、Country和Province。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveEdgeStatDataBodyUserRegionListItem struct {
	Area     *string `json:"Area,omitempty"`
	Country  *string `json:"Country,omitempty"`
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveEdgeStatDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveEdgeStatDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeLiveEdgeStatDataResResult `json:"Result"`
}

type DescribeLiveEdgeStatDataResResponseMetadata struct {

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

type DescribeLiveEdgeStatDataResResult struct {

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 边缘统计数据列表。
	EdgeStatDataList []DescribeLiveEdgeStatDataResResultEdgeStatDataListItem `json:"EdgeStatDataList"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 下行峰值带宽，单位为 Mbps。
	PeakDownBandwidth float32 `json:"PeakDownBandwidth"`

	// REQUIRED; 带宽峰值。取值范围为 [ ]，单位为，默认值为``。
	PeakUpBandwidth float32 `json:"PeakUpBandwidth"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 下行总流量，单位为 GB。
	TotalDownTraffic float32 `json:"TotalDownTraffic"`

	// REQUIRED; 请求数。
	TotalRequest float32 `json:"TotalRequest"`

	// REQUIRED; 上行总流量，单位为 GB。
	TotalUpTraffic float32 `json:"TotalUpTraffic"`

	// 应用列表。
	AppList []*string `json:"AppList,omitempty"`

	// 数据拆分的维度，维度说明如下所示。
	// * Domain：域名；
	// * Protocol：推拉流协议；
	// * ISP：运营商；
	// * Area：CDN 节点 IP 所属大区。
	// * Country：CDN 节点 IP 所属国家。
	// * Province：CDN 节点 IP 所属省份。
	// * UserArea：客户端 IP 所属大区。
	// * UserCountry：客户端 IP 所属国家。
	// * UserProvince：客户端 IP 所属省份。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 边缘统计详细数据列表。
	EdgeStatDetailDataList []*DescribeLiveEdgeStatDataResResultEdgeStatDetailDataListItem `json:"EdgeStatDetailDataList,omitempty"`

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
	// * huashu：华数。
	ISPList []*string `json:"ISPList,omitempty"`

	// byteplus比火山多了CMAF协议
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域列表。
	RegionList []*DescribeLiveEdgeStatDataResResultRegionListItem `json:"RegionList,omitempty"`

	// 流列表。
	StreamList []*string `json:"StreamList,omitempty"`

	// 客户端 IP 所属区域列表。
	UserRegionList []*DescribeLiveEdgeStatDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveEdgeStatDataResResultEdgeStatDataListItem struct {

	// REQUIRED; 下行带宽。
	DownBandwidth float32 `json:"DownBandwidth"`

	// REQUIRED; 当前数据聚合时间粒度内产生的总下行流量，单位 GB。
	DownTraffic float32 `json:"DownTraffic"`

	// REQUIRED; 请求参数。
	Request float32 `json:"Request"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的 UTC+8 时间，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 上行带宽。
	UpBandwidth float32 `json:"UpBandwidth"`

	// REQUIRED; 当前数据聚合时间粒度内产生的总上行流量，单位 GB。
	UpTraffic float32 `json:"UpTraffic"`
}

type DescribeLiveEdgeStatDataResResultEdgeStatDetailDataListItem struct {

	// REQUIRED; 边缘统计数据列表。
	EdgeStatDataList []DescribeLiveEdgeStatDataResResultEdgeStatDetailDataListPropertiesItemsItem `json:"EdgeStatDataList"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的下行峰值带宽，单位为 Mbps。
	PeakDownBandwidth float32 `json:"PeakDownBandwidth"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的上行峰值带宽，单位为 Mbps。
	PeakUpBandwidth float32 `json:"PeakUpBandwidth"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的下行总流量，单位为 GB。
	TotalDownTraffic float32 `json:"TotalDownTraffic"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的请求数。
	TotalRequest float32 `json:"TotalRequest"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的上行总流量，单位为 GB。
	TotalUpTraffic float32 `json:"TotalUpTraffic"`

	// 区域。
	Area *string `json:"Area,omitempty"`

	// 国家。
	Country *string `json:"Country,omitempty"`

	// 按域名维度进行数据拆分时的域名信息。
	Domain *string `json:"Domain,omitempty"`

	// 按运营商维度进行数据拆分时的运营商信息。
	ISP *string `json:"ISP,omitempty"`

	// 按推拉流协议维度进行数据拆分时的协议信息。
	Protocol *string `json:"Protocol,omitempty"`

	// 用户区域。
	UserArea *string `json:"UserArea,omitempty"`

	// 用户所在国家。
	UserCountry *string `json:"UserCountry,omitempty"`
}

type DescribeLiveEdgeStatDataResResultEdgeStatDetailDataListPropertiesItemsItem struct {

	// REQUIRED
	DownBandwidth float32 `json:"DownBandwidth"`

	// REQUIRED; 下行流量，单位 GB
	DownTraffic float32 `json:"DownTraffic"`

	// REQUIRED
	Request float32 `json:"Request"`

	// REQUIRED; 时间片起始时刻。RFC3339 格式的 UTC 时间，精度为 s，例如，2022-04-13T00:00:00+08:00
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED
	UpBandwidth float32 `json:"UpBandwidth"`

	// REQUIRED; 上行流量，单位 GB
	UpTraffic float32 `json:"UpTraffic"`
}

type DescribeLiveEdgeStatDataResResultRegionListItem struct {

	// 区域信息中的大区标识符。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveEdgeStatDataResResultUserRegionListItem struct {
	Area     *string `json:"Area,omitempty"`
	Country  *string `json:"Country,omitempty"`
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
	// * huashu：华数。
	ISPList []DescribeLiveISPDataResResultISPListItem `json:"ISPList"`
}

type DescribeLiveISPDataResResultISPListItem struct {

	// REQUIRED; 运营商标识符。
	Code string `json:"Code"`

	// REQUIRED; 运营商名称。
	Name string `json:"Name"`
}

type DescribeLiveLogDataBody struct {

	// REQUIRED; 仅支持查询最近31天的数据
	EndTime string `json:"EndTime"`

	// REQUIRED; 仅支持查询最近31天的数据
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 日志文件的信息列表。
	LogInfoList []DescribeLiveLogDataResResultLogInfoListItem `json:"LogInfoList"`

	// REQUIRED; 数据分页信息。
	Pagination DescribeLiveLogDataResResultPagination `json:"Pagination"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// 聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 60：1 分钟。时间粒度为 1 分钟时，单次查询最大时间跨度为 24 小时，历史查询时间范围为 366 天；
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用名称，取值与直播流地址中的 AppName 字段取值相同，查询流粒度数据时必传，且需同时传入 Stream。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30
	// 个字符。 :::tip 查询流粒度的带宽监控数据时，需同时指定 App 和 Stream 来指定直播流。 :::
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
	// * huashu：华数。
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// byteplus比火山多了CMAF协议
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	RegionList []*DescribeLiveMetricBandwidthDataBodyRegionListItem `json:"RegionList,omitempty"`

	// 流名称，预置与直播流地址中的 StreamName 字段取值相同，查询流粒度数据时必传，且需同时传入 Stream。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到
	// 100 个字符。 :::tip 查询流粒度的带宽监控数据时，需同时指定 App 和 Stream 来指定直播流。 :::
	Stream *string `json:"Stream,omitempty"`

	// 客户端 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	UserRegionList []*DescribeLiveMetricBandwidthDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveMetricBandwidthDataBodyRegionListItem struct {

	// 区域信息中的大区标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按国家筛选，需要同时传入Area和Country。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符，国外暂不支持该参数，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按省筛选，需要同时传入Area、Country和Province。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveMetricBandwidthDataBodyUserRegionListItem struct {

	// 大区，映射关系请参见区域映射
	Area *string `json:"Area,omitempty"`

	// 国家，映射关系请参见区域映射。如果按国家筛选，需要同时传入 Area 和 Country。
	Country *string `json:"Country,omitempty"`

	// 国内为省，国外暂不支持该参数，映射关系请参见区域映射。如果按省筛选，需要同时传入 Area、Country 和 Province。
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

	// REQUIRED; 聚合的时间粒度，单位为秒。
	// * 60：1 分钟；
	// * 300：5 分钟；
	// * 3600：1 小时。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 所有时间粒度的数据。
	BandwidthDataList []DescribeLiveMetricBandwidthDataResResultBandwidthDataListItem `json:"BandwidthDataList"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询时间范围内的下行峰值，单位为 Mbps。
	PeakDownBandwidth float32 `json:"PeakDownBandwidth"`

	// REQUIRED; 查询时间范围内的上行峰值，单位为 Mbps。
	PeakUpBandwidth float32 `json:"PeakUpBandwidth"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// 查询流粒度数据时的应用名称。
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
	// * huashu：华数。
	ISPList []*string `json:"ISPList,omitempty"`

	// byteplus比火山多了CMAF协议
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域列表。
	RegionList []*DescribeLiveMetricBandwidthDataResResultRegionListItem `json:"RegionList,omitempty"`

	// 查询流粒度数据时的流名称。
	Stream *string `json:"Stream,omitempty"`

	// 客户端 IP 所属区域列表。
	UserRegionList []*DescribeLiveMetricBandwidthDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveMetricBandwidthDataResResultBandwidthDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的下行峰值带宽，单位为 Mbps。
	DownBandwidth float32 `json:"DownBandwidth"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
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

	// REQUIRED; 时间片起始时刻。RFC3339 格式的 UTC 时间，精度为 s，例如，2022-04-13T00:00:00+08:00
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 上行带宽，单位为 Mbps
	UpBandwidth float32 `json:"UpBandwidth"`
}

type DescribeLiveMetricBandwidthDataResResultRegionListItem struct {

	// 区域信息中的大区标识符。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的身份标识符。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveMetricBandwidthDataResResultUserRegionListItem struct {
	Area     *string `json:"Area,omitempty"`
	Country  *string `json:"Country,omitempty"`
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveMetricTrafficDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// 聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 60：1 分钟。时间粒度为 1 分钟时，单次查询最大时间跨度为 24 小时，历史查询时间范围为 366 天；
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用名称，取值与直播流地址中的 AppName 字段取值相同，查询流粒度数据时必传，且需同时传入 Stream。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30
	// 个字符。 :::tip 查询流粒度的流量监控数据时，需同时指定 App 和 Stream 来指定直播流。 :::
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
	// * huashu：华数。
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// byteplus比火山多了CMAF协议
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	RegionList []*DescribeLiveMetricTrafficDataBodyRegionListItem `json:"RegionList,omitempty"`

	// 指定查询的流量数据为闲时或忙时，缺省情况下为查询全部数据，支持的取值如下。
	// * busy：忙时；
	// * free：闲时。
	Stage *string `json:"Stage,omitempty"`

	// 流名称，取值与直播流地址中的 StreamName 字段取值相同，查询流粒度数据时必传，且需同时传入 Stream。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到
	// 100 个字符。 :::tip 查询流粒度的流量监控数据时，需同时指定 App 和 Stream 来指定直播流。 :::
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 查询时间范围内的下行总流量，单位为 GB。
	TotalDownTraffic float32 `json:"TotalDownTraffic"`

	// REQUIRED; 查询时间范围内的上行总流量，单位为 GB。
	TotalUpTraffic float32 `json:"TotalUpTraffic"`

	// REQUIRED; 所有时间粒度的数据。
	TrafficDataList []DescribeLiveMetricTrafficDataResResultTrafficDataListItem `json:"TrafficDataList"`

	// 查询流粒度数据时的应用名称。
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
	// * huashu：华数。
	ISPList []*string `json:"ISPList,omitempty"`

	// byteplus比火山多了CMAF协议
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域列表。
	RegionList []*DescribeLiveMetricTrafficDataResResultRegionListItem `json:"RegionList,omitempty"`

	// 流量数据为闲时或忙时，取值说明如下。
	// * busy：忙时；
	// * free：闲时。
	Stage *string `json:"Stage,omitempty"`

	// 查询流粒度数据时的流名称。
	Stream *string `json:"Stream,omitempty"`

	// 按维度拆分后的数据。 :::tip 配置数据拆分维度时，对应的维度参数传入多个值时会返回按维度进行拆分的数据；对应的维度只传入一个值时不返回按此维度进行拆分的数据。 :::
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

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。 :::tip 单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，当前接口默认且仅支持按 300 秒进行数据拆分。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 域名列表，默认为空，表示查询您视频直播产品下所有域名的 95 峰值带宽用量数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// byteplus比火山多了CMAF协议
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 时间范围内的上下行 95 峰值带宽总和。 :::tip 如果请求时，Regionlist中传入多个 region，则返回这些 region 的上下行带宽 95 峰值总和。 :::
	P95PeakBandwidth float32 `json:"P95PeakBandwidth"`

	// REQUIRED; 95 峰值带宽的时间戳，RFC3339 格式的时间戳，精度为秒。
	P95PeakTimestamp string `json:"P95PeakTimestamp"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// byteplus比火山多了CMAF协议
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// 聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 60：（默认值）1 分钟。时间粒度为 1 分钟时，单次查询最大时间跨度为 24 小时，历史查询时间范围为 366 天；
	// * 300：5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 数据拆分的维度，默认为空表示不按维度进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * ISP：运营商。
	// :::tip 配置数据拆分的维度时，对应的维度参数传入多个值时才会返回按此维度拆分的数据。例如，配置按 Domain 进行数据拆分时， DomainList 传入多个 Domain 值时，才会返回按 Domain 拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，默认为空时表示查询所有域名下产生的请求状态码占比数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取待查询请求状态码占比数据的域名。
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
	// * huashu：华数。
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// CDN 节点 IP 所属区域的列表，缺省情况下表示所有区域。 :::tip 参数 RegionList和UserRegionList 不支持同时传入。 :::
	RegionList []*DescribeLivePlayStatusCodeDataBodyRegionListItem `json:"RegionList,omitempty"`

	// 请求类型，取值及含义如下所示。
	// * Access：（默认值）推流请求和拉流请求；
	// * Source：回源请求。
	Type *string `json:"Type,omitempty"`

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

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
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
	// * huashu：华数。
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

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天；
	// * 86400：1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 支持域名拆分
	DetailField []*string `json:"DetailField,omitempty"`

	// 拉流转推任务群组列表，默认为空，表示查询所有拉流转推任务群组的带宽用量。
	GroupList []*string `json:"GroupList,omitempty"`
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 当前查询条件下的拉流转推峰值带宽，单位为 Mbps。
	PeakUpBandwidth float32 `json:"PeakUpBandwidth"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// 按维度拆分后的数据。 :::tip 当配置了数据拆分的维度时，对应的维度参数传入多个值才会返回按维度拆分的数据。 :::
	BandwidthDetailDataList []*DescribeLivePullToPushBandwidthDataResResultBandwidthDetailDataListItem `json:"BandwidthDetailDataList,omitempty"`

	// 数据拆分的维度。
	DetailField []*string `json:"DetailField,omitempty"`

	// 拉流转推任务群组列表。
	GroupList []*string `json:"GroupList,omitempty"`
}

type DescribeLivePullToPushBandwidthDataResResultBandwidthDataListItem struct {

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内的拉流转推峰值带宽，单位为 Mbps。
	UpBandwidth float32 `json:"UpBandwidth"`
}

type DescribeLivePullToPushBandwidthDataResResultBandwidthDetailDataListItem struct {

	// REQUIRED; 按维度进行数据拆分后，当前维度下所有时间粒度的数据。
	BandwidthDataList []DescribeLivePullToPushBandwidthDataResResultBandwidthDetailDataListPropertiesItemsItem `json:"BandwidthDataList"`

	// REQUIRED; 查询时间范围内的维度下的拉流转推峰值带宽，单位为 Mbps。
	PeakUpBandwidth float32 `json:"PeakUpBandwidth"`

	// 按推流地址类型维度进行数据拆分时的地址类型信息。
	DstAddrType *string `json:"DstAddrType,omitempty"`

	// 按任务群组维度进行数据拆分时的群组信息。
	Group *string `json:"Group,omitempty"`
}

type DescribeLivePullToPushBandwidthDataResResultBandwidthDetailDataListPropertiesItemsItem struct {

	// REQUIRED; 时间片起始时刻。RFC3339 格式的 UTC 时间，精度为 s，例如，2022-04-13T00:00:00+08:00
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 转推带宽，单位为 Mbps
	UpBandwidth float32 `json:"UpBandwidth"`
}

type DescribeLivePullToPushDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 60：1 分钟。时间粒度为 1 分钟时，单次查询最大时间跨度为 1 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询时间跨度为 93 天，历史查询时间范围为 366 天；
	// * 86400：（默认值）1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	// :::tip 查询流粒度数据时，需同时传入 App 和 Stream。 :::
	App *string `json:"App,omitempty"`

	// 支持群组拆分
	DetailField []*string `json:"DetailField,omitempty"`

	// 群组
	GroupList []*string `json:"GroupList,omitempty"`

	// 流名称，取值与直播流地址中 StreamName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。 :::tip 查询流粒度数据时，需同时传入
	// App 和 Stream。 :::
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 所有时间粒度的数据。
	PullToPushDataList []DescribeLivePullToPushDataResResultPullToPushDataListItem `json:"PullToPushDataList"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 当前查询条件下的拉流转推总时长，单位为分钟。
	TotalDuration float32 `json:"TotalDuration"`

	// 查询流粒度数据时的应用名称。
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，当前接口仅支持按 Group 即拉流转推任务群组维度进行数据拆分。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 拉流转推任务群组。
	GroupList []*string `json:"GroupList,omitempty"`

	// 按维度拆分后的数据。
	PullToPushDetailDataList []*DescribeLivePullToPushDataResResultPullToPushDetailDataListItem `json:"PullToPushDetailDataList,omitempty"`

	// 查询流粒度数据时的流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLivePullToPushDataResResultPullToPushDataListItem struct {

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内的拉流转推总时长，单位为分钟。
	Value float32 `json:"Value"`
}

type DescribeLivePullToPushDataResResultPullToPushDetailDataListItem struct {

	// REQUIRED; 按维度进行数据拆分后，当前维度下所有时间粒度的数据。
	PullToPushDataList []DescribeLivePullToPushDataResResultPullToPushDetailDataListPropertiesItemsItem `json:"PullToPushDataList"`

	// REQUIRED; 按维度进行数据拆分后，当前维度的拉流转推总时长，单位分钟。
	TotalDuration float32 `json:"TotalDuration"`

	// 按任务群组维度进行数据拆分时的群组信息。
	Group *string `json:"Group,omitempty"`
}

type DescribeLivePullToPushDataResResultPullToPushDetailDataListPropertiesItemsItem struct {

	// REQUIRED; 时间片起始时刻。RFC3339 格式的 UTC 时间，精度为 s，例如，2022-04-13T00:00:00+08:00
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

	// REQUIRED; 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 推流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的推流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	// :::tip 单次查询最大时间跨度为 1 天，历史查询最大时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// REQUIRED; 流名称，取值与直播流地址中 StreamName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
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

	// 数据聚合的时间粒度，单位为秒。
	// * 5：5 秒；
	// * 30：30 秒。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 推流域名。
	Domain *string `json:"Domain,omitempty"`

	// 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime *string `json:"EndTime,omitempty"`

	// 所有时间粒度的数据。
	MetricList []*DescribeLivePushStreamMetricsResResultMetricListItem `json:"MetricList,omitempty"`

	// 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime *string `json:"StartTime,omitempty"`

	// 流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLivePushStreamMetricsResResultMetricListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的音频码率最大值，单位为 kbps。
	AudioBitrate float32 `json:"AudioBitrate"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻音频帧显示时间戳差值的最大值，单位为毫秒。
	AudioFrameGap float32 `json:"AudioFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内的音频帧率最大值，单位为 fps。
	AudioFramerate float32 `json:"AudioFramerate"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个音频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	AudioPts float64 `json:"AudioPts"`

	// REQUIRED; 当前数据聚合时间粒度内的视频码率最大值，单位为 kbps。
	Bitrate float32 `json:"Bitrate"`

	// REQUIRED; 当前数据聚合时间粒度内的视频帧率最大值，单位为 fps。
	Framerate float32 `json:"Framerate"`

	// REQUIRED; 当前数据聚合时间粒度内，所有音视频帧显示时间戳差值的最大值，即所有 AudioPts 与 VideoPts 差值的最大值，单位为毫秒。
	PtsDelta float32 `json:"PtsDelta"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻视频帧显示时间戳差值的最大值，单位为毫秒。
	VideoFrameGap float32 `json:"VideoFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个视频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	VideoPts float64 `json:"VideoPts"`
}

type DescribeLiveRecordDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询最大时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天；
	// * 86400：1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	// :::tip 查询流粒度数据时，需同时传入 App 和 Stream。 :::
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，默认为空表示不按维度进行数据拆分，当前接口仅支持填写 Domain 表示按查询的域名为维度进行数据拆分。 :::tip 配置数据拆分的维度时，对应的维度参数传入多个值时才会返回按此维度拆分的数据。例如，配置按 Domain
	// 进行数据拆分时， DomainList 传入多个 Domain 值时，才会返回按 Domain 拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，默认为空，表示查询您视频直播产品下所有域名的录制用量数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 流名称，取值与直播流地址中 StreamName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	// :::tip 查询流粒度数据时，需同时传入 App 和 Stream。 :::
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 所有时间粒度的数据。
	RecordDataList []DescribeLiveRecordDataResResultRecordDataListItem `json:"RecordDataList"`

	// REQUIRED; 当前查询条件下的录制并发路数最大值。
	RecordPeak int32 `json:"RecordPeak"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// 查询流粒度数据时的应用名称。
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，当前接口仅支持按 Domain 即域名维度进行数据拆分。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 按维度拆分后的数据。
	RecordDetailDataList []*DescribeLiveRecordDataResResultRecordDetailDataListItem `json:"RecordDetailDataList,omitempty"`

	// 查询流粒度数据时的流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveRecordDataResResultRecordDataListItem struct {

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 300：5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天；
	// * 86400：（默认值）1 天。时间粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	// :::tip 查询流粒度数据时，需同时传入 App 和 Stream。 :::
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，默认为空表示不按维度进行数据拆分，当前接口仅支持填写 Domain 表示按查询的域名为维度进行数据拆分。
	// :::tip 配置数据拆分的维度时，对应的维度参数传入多个值时才会返回按此维度拆分的数据。例如，配置按 Domain 进行数据拆分时， DomainList 传入多个 Domain 值时，才会返回按 Domain 拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表，默认为空，表示查询您视频直播产品下所有域名的截图用量数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 流名称，取值与直播流地址中 StreamName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。 :::tip 查询流粒度数据时，需同时传入
	// App 和 Stream。 :::
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 所有时间粒度的数据。
	SnapshotDataList []DescribeLiveSnapshotDataResResultSnapshotDataListItem `json:"SnapshotDataList"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 当前查询条件下的截图总张数。
	Total int32 `json:"Total"`

	// 查询流粒度数据时的应用名称。
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，当前接口仅支持按 Domain 即域名维度进行数据拆分。
	DetailField []*string `json:"DetailField,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 按维度拆分后的数据。
	SnapshotDetailData []*DescribeLiveSnapshotDataResResultSnapshotDetailDataItem `json:"SnapshotDetailData,omitempty"`

	// 查询流粒度数据时的流名称。
	Stream *string `json:"Stream,omitempty"`
}

type DescribeLiveSnapshotDataResResultSnapshotDataListItem struct {

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// 聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 60：1 分钟。时间粒度为 1 分钟时，单次查询最大时间跨度为 24 小时，历史查询时间范围为 366 天；
	// * 300：（默认值）5 分钟。时间粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询时间范围为 366 天；
	// * 3600：1 小时。时间粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询时间范围为 366 天。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 回源流的应用名称，查询流粒度数据时必传，且需同时传入 Domain 和 Stream。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。 :::tip
	// 查询流粒度的回源带宽监控数据时，需同时指定 Domain 、App 和 Stream 来指定回源流。 :::
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
	// * huashu：华数。
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// 回源流的流名称，查询流粒度数据时必传，且需同时传入 Domain 和 App。支持由大小写字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。 :::tip 查询流粒度的回源带宽监控数据时，需同时指定
	// Domain 、App 和 Stream 来指定回源流。 :::
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

	// REQUIRED; 所有时间粒度的数据。
	BandwidthDataList []DescribeLiveSourceBandwidthDataResResultBandwidthDataListItem `json:"BandwidthDataList"`

	// 聚合的时间粒度，单位为秒。
	// * 60：1 分钟；
	// * 300：5 分钟；
	// * 3600：1 小时。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 查询流粒度数据时的应用名称。
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

	// 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime *string `json:"EndTime,omitempty"`

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
	// * huashu：华数。
	ISPList []*string `json:"ISPList,omitempty"`

	// 查询时间范围内的回源峰值带宽，单位为 Mbps。
	PeakBandwidth *float32 `json:"PeakBandwidth,omitempty"`

	// 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime *string `json:"StartTime,omitempty"`

	// 查询流粒度数据时的流名称。
	Stream *string `json:"Stream,omitempty"`

	// 客户端 IP 所属区域列表。
	UserRegionList []*DescribeLiveSourceBandwidthDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveSourceBandwidthDataResResultBandwidthDataListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的回源峰值带宽，单位为 Mbps。
	Bandwidth float32 `json:"Bandwidth"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
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

	// 时间片起始时刻。RFC3339 格式的 UTC 时间，精度为 s，例如，2022-04-13T00:00:00+08:00
	TimeStamp *string `json:"TimeStamp,omitempty"`
}

type DescribeLiveSourceBandwidthDataResResultUserRegionListItem struct {

	// 区域信息中的大区标识符。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveSourceStreamMetricsBody struct {

	// REQUIRED; 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	App string `json:"App"`

	// REQUIRED; 拉流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看回源流使用的拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	// :::tip 单次查询最大时间跨度为 1 天，历史查询最大时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// REQUIRED; 流名称，取值与直播流地址中 StreamName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
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

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 拉流域名。
	Domain string `json:"Domain"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 所有时间粒度的数据。
	MetricList []DescribeLiveSourceStreamMetricsResResultMetricListItem `json:"MetricList"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 流名称。
	Stream string `json:"Stream"`
}

type DescribeLiveSourceStreamMetricsResResultMetricListItem struct {

	// REQUIRED; 当前数据聚合时间粒度内的音频码率最大值，单位为 kbps。
	AudioBitrate float32 `json:"AudioBitrate"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻音频帧显示时间戳差值的最大值，单位为毫秒。
	AudioFrameGap float32 `json:"AudioFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内的音频帧率最大值，单位为 fps。
	AudioFramerate float32 `json:"AudioFramerate"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个音频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	AudioPts float64 `json:"AudioPts"`

	// REQUIRED; 当前数据聚合时间粒度内的视频码率最大值，单位为 kbps。
	Bitrate float32 `json:"Bitrate"`

	// REQUIRED; 当前数据聚合时间粒度内的视频帧率最大值，单位为 fps
	Framerate float32 `json:"Framerate"`

	// REQUIRED; 当前数据聚合时间粒度内，所有音视频帧显示时间戳差值的最大值，即所有 AudioPts 与 VideoPts 差值的最大值，单位为毫秒。
	PtsDelta float32 `json:"PtsDelta"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内，相邻视频帧显示时间戳差值的最大值，单位为毫秒。
	VideoFrameGap float32 `json:"VideoFrameGap"`

	// REQUIRED; 当前数据聚合时间粒度内，最后一个视频帧的显示时间戳 PTS（Presentation Time Stamp），单位为毫秒。
	VideoPts float64 `json:"VideoPts"`
}

type DescribeLiveSourceTrafficDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。 :::tip 历史查询最大时间范围为 366 天，单次查询最大时间跨度与数据拆分维度和数据聚合时间粒度有关，详细如下。
	// * 当不进行维度拆分或只使用一个维度拆分数据时： * 数据以 60 秒聚合时，单次查询最大时间跨度为 24 小时；
	// * 数据以 300 秒聚合时，单次查询最大时间跨度为 31 天；
	// * 数据以 3600 秒聚合时，单次查询最大时间跨度为 31 天。
	//
	//
	// * 当使用两个或两个以上维度拆分数据时： * 数据以 60 秒聚合时，单次查询最大时间跨度为 3 小时；
	// * 数据以 300 秒聚合时，单次查询最大时间跨度为 24 小时；
	// * 数据以 3600 秒聚合时，单次查询最大时间跨度为 7 天。 :::
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 60：1 分钟。
	// * 300：（默认值）5 分钟。
	// * 3600：1 小时。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 回源流的应用名称，查询流粒度数据时必传，且需同时传入 Domain 和 Stream。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。 :::tip
	// 查询流粒度数据时，需同时指定 Domain 、App 和 Stream 来指定回源流。 :::
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，默认为空表示不按维度进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * ISP：运营商。
	// :::tip 配置数据拆分的维度时，对应的维度参数传入多个值时才会返回按此维度拆分的数据。例如，配置按 Domain 进行数据拆分时， DomainList 传入多个 Domain 值时，才会返回按 Domain 拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 拉流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的拉流域名。
	// :::tip 查询流粒度数据时，需同时指定 Domain 、App 和
	// Stream 来指定回源流。 :::
	Domain *string `json:"Domain,omitempty"`

	// 拉流域名列表，默认为空，表示查询所有域名的回源流量带宽监控数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的拉流域名。
	// :::tipDomainList 和 Domain 传且仅传一个。 :::
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
	// * huashu：华数。
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// 回源流的流名称，查询流粒度数据时必传，且需同时传入 Domain 和 App。支持由大小写字母（A - Z、a - z）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。 :::tip 查询流粒度数据时，需同时指定
	// Domain 、App 和 Stream 来指定回源流。 :::
	Stream *string `json:"Stream,omitempty"`

	// 客户端 IP 所属区域的列表，缺省情况下表示所有区域。
	UserRegionList []*DescribeLiveSourceTrafficDataBodyUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveSourceTrafficDataBodyUserRegionListItem struct {

	// 区域信息中的大区标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按国家筛选，需要同时传入Area和Country。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符，国外暂不支持该参数，如何获取请参见查询区域标识符 [https://www.volcengine.com/docs/6469/1133973]。如果按省筛选，需要同时传入Area、Country和Province。
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

	// REQUIRED; 峰值带宽。
	PeakBandwidth float32 `json:"PeakBandwidth"`

	// REQUIRED; 查询时间范围内的回源总流量，单位为 GB。
	TotalTraffic float32 `json:"TotalTraffic"`

	// REQUIRED; 所有时间粒度的数据。
	TrafficDataList []DescribeLiveSourceTrafficDataResResultTrafficDataListItem `json:"TrafficDataList"`

	// 聚合的时间粒度，单位为秒。
	// * 60：1 分钟；
	// * 300：5 分钟；
	// * 3600：1 小时。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 查询流粒度数据时的应用名称。
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，维度说明如下所示。
	// * Domain：域名；
	// * ISP：运营商。
	DetailField []*string `json:"DetailField,omitempty"`

	// 查询流粒度数据时的域名。
	Domain *string `json:"Domain,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime *string `json:"EndTime,omitempty"`

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
	// * huashu：华数。
	ISPList []*string `json:"ISPList,omitempty"`

	// 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime *string `json:"StartTime,omitempty"`

	// 查询流粒度数据时的流名称。
	Stream *string `json:"Stream,omitempty"`

	// 按维度拆分后的数据。
	TrafficDetailDataList []*DescribeLiveSourceTrafficDataResResultTrafficDetailDataListItem `json:"TrafficDetailDataList,omitempty"`

	// 客户端 IP 所属区域列表。
	UserRegionList []*DescribeLiveSourceTrafficDataResResultUserRegionListItem `json:"UserRegionList,omitempty"`
}

type DescribeLiveSourceTrafficDataResResultTrafficDataListItem struct {

	// REQUIRED; 带宽。取值范围为 [ ]，单位为，默认值为``。
	Bandwidth float32 `json:"Bandwidth"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 当前数据聚合时间粒度内产生的回源流量，单位 GB。
	Traffic float32 `json:"Traffic"`
}

type DescribeLiveSourceTrafficDataResResultTrafficDetailDataListItem struct {

	// REQUIRED; 峰值带宽。
	PeakBandwidth float32 `json:"PeakBandwidth"`

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

	// REQUIRED; 带宽。
	Bandwidth float32 `json:"Bandwidth"`

	// REQUIRED; 时间片起始时刻。RFC3339 格式的 UTC 时间，精度为 s，例如，2022-04-13T00:00:00+08:00
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 回源流量，单位 GB
	Traffic float32 `json:"Traffic"`
}

type DescribeLiveSourceTrafficDataResResultUserRegionListItem struct {

	// 区域信息中的大区标识符。
	Area *string `json:"Area,omitempty"`

	// 区域信息中的国家标识符。
	Country *string `json:"Country,omitempty"`

	// 区域信息中的省份标识符。
	Province *string `json:"Province,omitempty"`
}

type DescribeLiveStorageSpaceDataBody struct {

	// REQUIRED; 结束时间，RFC3339 格式的 UTC 时间，精度为 s。
	EndTime string `json:"EndTime"`

	// REQUIRED; 开始时间，RFC3339 格式的 UTC 时间，精度为 s。
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

	// REQUIRED; 聚合时间粒度。
	Aggregation bool `json:"Aggregation"`

	// REQUIRED; 结束时间。
	EndTime string `json:"EndTime"`

	// REQUIRED; 开始时间。
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。 :::tip 历史查询最大时间范围为 366 天，单次查询最大时间跨度与数据拆分维度和数据聚合时间粒度有关，详细如下。
	// * 当不进行维度拆分或只使用一个维度拆分数据时： * 数据以 60 秒聚合时，单次查询最大时间跨度为 24 小时；
	// * 数据以 300 秒聚合时，单次查询最大时间跨度为 31 天；
	// * 数据以 3600 秒聚合时，单次查询最大时间跨度为 31 天。
	//
	//
	// * 当使用两个或两个以上维度拆分数据时： * 数据以 60 秒聚合时，单次查询最大时间跨度为 3 小时；
	// * 数据以 300 秒聚合时，单次查询最大时间跨度为 24 小时；
	// * 数据以 3600 秒聚合时，单次查询最大时间跨度为 7 天。 :::
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 60：1 分钟。
	// * 300：（默认值）5 分钟。
	// * 3600：1 小时。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用名称，取值与直播流地址中的 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。 :::tip 查询流粒度的请求数和在线人数数据时，需同时指定
	// Domain 、App 和 Stream 来指定直播流。 :::
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，默认为空表示不按维度进行数据拆分，支持的维度如下所示。
	// * Domain：域名；
	// * ISP：运营商；
	// * Protocol：推拉流协议。
	// :::tip 配置数据拆分的维度时，对应的维度参数传入多个值时才会返回按此维度拆分的数据。例如，配置按 Domain 进行数据拆分时， DomainList 传入多个 Domain 值时，才会返回按 Domain 拆分的数据。 :::
	DetailField []*string `json:"DetailField,omitempty"`

	// 拉流域名，您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的拉流域名。
	// :::tip 查询流粒度的请求数和在线人数数据时，需同时指定 Domain 、
	// App 和 Stream 来指定直播流。 :::
	Domain *string `json:"Domain,omitempty"`

	// 拉流域名列表，默认为空，表示查询所有域名的请求数和在线人数。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的拉流域名。
	// :::tipDomainList 和 Domain 传且仅传一个。 :::
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
	// * huashu：华数。
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// byteplus比火山多了CMAF协议
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// 指定拉流请求的 Referer 信息，默认为空，表示不对拉流请求的 Referer 字段进行校验。
	RefererList []*string `json:"RefererList,omitempty"`

	// CDN 节点 IP 所属区域的列表，缺省情况下表示所有区域。
	RegionList []*DescribeLiveStreamSessionDataBodyRegionListItem `json:"RegionList,omitempty"`

	// 数据流名称。
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
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询时间范围内的在线人数峰值。
	PeakOnlineUser int32 `json:"PeakOnlineUser"`

	// REQUIRED; 所有时间粒度的数据。
	SessionDataList []DescribeLiveStreamSessionDataResResultSessionDataListItem `json:"SessionDataList"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 总体边缘请求命中率
	TotalHitRate float64 `json:"TotalHitRate"`

	// REQUIRED; 查询时间范围内的请求数。
	TotalRequest int32 `json:"TotalRequest"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 数据拆分的维度，维度说明如下所示。
	// * Domain：域名；
	// * ISP：运营商；
	// * Protocol：推拉流协议；
	DetailField []*string `json:"DetailField,omitempty"`

	// 拉流域名。
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
	// * huashu：华数。
	ISPList []*string `json:"ISPList,omitempty"`

	// byteplus比火山多了CMAF协议
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// 拉流请求的 Referer 信息。
	RefererList []*string `json:"RefererList,omitempty"`

	// CDN 节点 IP 所属的区域列表，缺省情况下表示所有区域。
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

	// REQUIRED; 边缘请求命中率
	HitRate float64 `json:"HitRate"`

	// REQUIRED; 当前数据聚合时间粒度内的在线人数最大值。
	OnlineUser int32 `json:"OnlineUser"`

	// REQUIRED; 当前数据聚合时间粒度内的请求数。
	Request int32 `json:"Request"`

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
	TimeStamp string `json:"TimeStamp"`
}

type DescribeLiveStreamSessionDataResResultSessionDetailDataListItem struct {

	// REQUIRED; 按维度进行数据拆分后，当前维度的在线人数峰值。
	PeakOnlineUser int32 `json:"PeakOnlineUser"`

	// REQUIRED; 按维度进行数据拆分后，当前维度下所有时间粒度的数据。
	SessionDataList []DescribeLiveStreamSessionDataResResultSessionDetailDataListPropertiesItemsItem `json:"SessionDataList"`

	// REQUIRED; 总体边缘请求命中率
	TotalHitRate float32 `json:"TotalHitRate"`

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

	// REQUIRED; 边缘请求命中率
	HitRate float64 `json:"HitRate"`

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

type DescribeLiveTopPlayDataBody struct {

	// REQUIRED; 结束时间，RFC3339 格式，例如：2021-04-14T00:00:00+08:00 单次最长跨度是31天 历史查询范围是366天
	EndTime string `json:"EndTime"`

	// REQUIRED; 起始时间，RFC3339 格式，例如：2021-04-13T00:00:00+08:00
	StartTime string `json:"StartTime"`

	// VhostList和DomainList 2选1，支持为空，都为空查询全部
	DomainList []*string `json:"DomainList,omitempty"`

	// 查询数据的页码，默认值为 1，表示查询第一页的数据。 最多展示排名前1000的数据。
	PageNum *string `json:"PageNum,omitempty"`

	// 每页显示的数据条数，默认值为 10，取值范围为 [1,1000]。
	PageSize *string `json:"PageSize,omitempty"`

	// 查询类型，枚举值：
	// * Domain
	// * Stream（默认值）
	QueryType *string `json:"QueryType,omitempty"`

	// 排序指标，枚举值：
	// * PeakBandwidth 带宽峰值（默认值）
	// * AvgBandwidth 平均带宽
	// * TotalTraffic 流量加和
	SortBy *string `json:"SortBy,omitempty"`

	// VhostList和DomainList 2选1，支持为空，都为空查询全部
	VhostList []*string `json:"VhostList,omitempty"`
}

type DescribeLiveTopPlayDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveTopPlayDataResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED; 视请求的接口而定
	Result DescribeLiveTopPlayDataResResult `json:"Result"`
}

type DescribeLiveTopPlayDataResResponseMetadata struct {

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

// DescribeLiveTopPlayDataResResult - 视请求的接口而定
type DescribeLiveTopPlayDataResResult struct {

	// REQUIRED; 带宽和流量详细数据。
	DataItemList []DescribeLiveTopPlayDataResResultDataItemListItem `json:"DataItemList"`

	// REQUIRED; 域名列表。
	DomainList []string `json:"DomainList"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询结果的分页信息。
	Pagination DescribeLiveTopPlayDataResResultPagination `json:"Pagination"`

	// REQUIRED; 查询类型，取值及含义如下所示。
	// * Domain ：查询 TOPN 域名的的流量带宽信息。
	// * Stream（默认值）查询 TOPN 直播流的流量带宽信息。
	QueryType string `json:"QueryType"`

	// REQUIRED; TOPN 结果的排序指标，取值及含义如下所示。
	// * PeakBandwidth（默认值）：以峰值带宽值降序展示查询结果。
	// * AvgBandwidth：以平均带宽值降序展示查询结果。
	// * TotalTraffic：以流量加值降序展示查询结果。
	SortBy string `json:"SortBy"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 域名空间列表。
	VhostList []string `json:"VhostList"`
}

type DescribeLiveTopPlayDataResResultDataItemListItem struct {

	// REQUIRED; 平均带宽，单位bps
	AvgBandwidth float32 `json:"AvgBandwidth"`

	// REQUIRED; 域名
	Domain string `json:"Domain"`

	// REQUIRED; 带宽峰值，单位bps
	PeakBandwidth float32 `json:"PeakBandwidth"`

	// REQUIRED; 流量加和，单位Byte
	TotalTraffic float32 `json:"TotalTraffic"`

	// AppName，根据Domain查询时为空
	App *string `json:"App,omitempty"`

	// 流名，根据Domain查询时为空
	Stream *string `json:"Stream,omitempty"`
}

// DescribeLiveTopPlayDataResResultPagination - 查询结果的分页信息。
type DescribeLiveTopPlayDataResResultPagination struct {

	// REQUIRED; 当前所在分页的页码。
	PageNum string `json:"PageNum"`

	// REQUIRED; 每页显示的数据条数。
	PageSize string `json:"PageSize"`

	// REQUIRED; 查询结果的数据总条数。
	TotalCount string `json:"TotalCount"`
}

type DescribeLiveTrafficDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，支持的时间粒度如下所示。
	// * 300：（默认值）5 分钟。聚合粒度为 5 分钟时，单次查询最大时间跨度为 31 天，历史查询最大时间范围为 366 天；
	// * 3600：1 小时。聚合粒度为 1 小时时，单次查询最大时间跨度为 93 天，历史查询最大时间范围为 366 天；
	// * 86400：1 天。聚合粒度为 1 天时，单次查询最大时间跨度为 93 天，历史查询最大时间范围为 366 天。
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
	// * huashu：华数。
	// 您也可以通过 DescribeLiveISPData [https://www.volcengine.com/docs/6469/1133974] 接口获取运营商对应的标识符。
	ISPList []*string `json:"ISPList,omitempty"`

	// byteplus比火山多了CMAF协议
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。
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
	// * huashu：华数。
	ISPList []*string `json:"ISPList,omitempty"`

	// byteplus比火山多了CMAF协议
	ProtocolList []*string `json:"ProtocolList,omitempty"`

	// CDN 节点 IP 所属区域列表。
	RegionList []*DescribeLiveTrafficDataResResultRegionListItem `json:"RegionList,omitempty"`

	// 流量数据为闲时或忙时，取值说明如下。
	// * busy：忙时；
	// * free：闲时。
	Stage *string `json:"Stage,omitempty"`

	// 按维度拆分后的数据。 :::tip 当配置了数据拆分的维度时，对应的维度参数传入多个值才会返回按维度拆分的数据。 :::
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

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
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

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。 :::tip 单次查询最大时间跨度为 93 天，历史查询最大时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 数据聚合的时间粒度，单位为秒，当前接口默认且仅支持按 86400 秒进行数据聚合。
	Aggregation *int32 `json:"Aggregation,omitempty"`

	// 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。
	// :::tip 查询流粒度数据时，需同时传入 App 和 Stream。 :::
	App *string `json:"App,omitempty"`

	// 域名列表，默认为空，表示查询您视频直播产品下所有域名的转码用量数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 分辨率。- 480P：640 × 480； - 720P：1280 × 720； - 1080P：1920 × 1088； - 2K：2560 × 1440； - 4K：4096 × 2160；- 8K：大于4K； - 0P：纯音频流；
	Resolution []*string `json:"Resolution,omitempty"`

	// 流名称，取值与直播流地址中 StreamName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。 :::tip 查询流粒度数据时，需同时传入
	// App 和 Stream。 :::
	Stream *string `json:"Stream,omitempty"`

	// 视频编码格式，默认为空表示不指定编码格式，支持的取值和含义如下所示。
	// * Normal_H264：H.264 标准转码；
	// * Normal_H265：H.265 标准转码；
	// * Normal_H266：H.266 标准转码；
	// * ByteHD_H264：H.264 极智超清；
	// * ByteHD_H265：H.265 极智超清；
	// * ByteHD_H266：H.266 极智超清；
	// * ByteQE：画质增强；
	// * Audio：纯音频流。
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

	// REQUIRED; 数据聚合的时间粒度，单位为秒。
	Aggregation int32 `json:"Aggregation"`

	// REQUIRED; 查询时间范围内的转码总时长，单位为分钟。
	Duration float32 `json:"Duration"`

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的起始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 所有时间粒度的数据。
	TranscodeDataList []DescribeLiveTranscodeDataResResultTranscodeDataListItem `json:"TranscodeDataList"`

	// 查询流粒度数据时的应用名称。
	App *string `json:"App,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 分辨率。- 480P：640 × 480； - 720P：1280 × 720； - 1080P：1920 × 1088； - 2K：2560 × 1440； - 4K：4096 × 2160；- 8K：大于4K； - 0P：纯音频流；
	Resolution []*string `json:"Resolution,omitempty"`

	// 查询流粒度数据时的流名称。
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

	// REQUIRED; 数据按时间粒度聚合时，每个时间粒度的开始时间，RFC3339 格式的时间戳，精度为秒。
	TimeStamp string `json:"TimeStamp"`

	// REQUIRED; 视频编码格式，支持的取值和含义如下所示。- NormalH264：H.264 标准转码； - NormalH265：H.265 标准转码； - NormalH266：H.266 标准转码； - ByteHDH264：H.264
	// 极智超清； - ByteHDH265：H.265 极智超清； - ByteHDH266：H.266 极智超清；- ByteQE：画质增强；- Audio：纯音频流；
	TranscodeType string `json:"TranscodeType"`
}

type DescribeLiveTranscodeInfoDataBody struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 查询的开始时间，RFC3339 格式的时间戳，精度为秒。 :::tip 历史查询最大时间范围为 366 天。 :::
	StartTime string `json:"StartTime"`

	// 应用名称，取值与直播流地址中 AppName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 30 个字符。 :::tip 查询流粒度数据时，需同时传入App和Stream。
	// :::
	App *string `json:"App,omitempty"`

	// 域名列表，默认为空，表示查询您视频直播产品下所有域名的转码用量数据。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，获取待查询的域名。
	DomainList []*string `json:"DomainList,omitempty"`

	// 分页页数，默认1
	PageNum *int32 `json:"PageNum,omitempty"`

	// 每页大小， 默认20，取值范围[1,100000]
	PageSize *int32 `json:"PageSize,omitempty"`

	// 分辨率。- 480P：640 × 480； - 720P：1280 × 720； - 1080P：1920 × 1088； - 2K：2560 × 1440； - 4K：4096 × 2160；- 8K：大于4K； - 0P：纯音频流；
	Resolution []*string `json:"Resolution,omitempty"`

	// 流名称，取值与直播流地址中 StreamName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。 :::tip 查询流粒度数据时，需同时传入App和Stream。
	// :::
	Stream *string `json:"Stream,omitempty"`

	// 视频编码格式，支持的取值和含义如下所示。- NormalH264：H.264 标准转码； - NormalH265：H.265 标准转码； - NormalH266：H.266 标准转码； - ByteHDH264：H.264 极智超清；
	// - ByteHDH265：H.265 极智超清； - ByteHDH266：H.266 极智超清；- ByteQE：画质增强；- Audio：纯音频流；
	TranscodeType []*string `json:"TranscodeType,omitempty"`
}

type DescribeLiveTranscodeInfoDataRes struct {

	// REQUIRED
	ResponseMetadata DescribeLiveTranscodeInfoDataResResponseMetadata `json:"ResponseMetadata"`
	Result           *DescribeLiveTranscodeInfoDataResResult          `json:"Result,omitempty"`
}

type DescribeLiveTranscodeInfoDataResResponseMetadata struct {

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
	Error   *DescribeLiveTranscodeInfoDataResResponseMetadataError `json:"Error,omitempty"`
}

type DescribeLiveTranscodeInfoDataResResponseMetadataError struct {

	// 错误码
	Code *string `json:"Code,omitempty"`

	// 错误信息
	Message *string `json:"Message,omitempty"`
}

type DescribeLiveTranscodeInfoDataResResult struct {

	// REQUIRED; 查询的结束时间，RFC3339 格式的时间戳，精度为秒。
	EndTime string `json:"EndTime"`

	// REQUIRED; 分页信息
	Pagination DescribeLiveTranscodeInfoDataResResultPagination `json:"Pagination"`

	// REQUIRED; 查询的起始时间，RFC3339 格式的时间戳，精度为秒。
	StartTime string `json:"StartTime"`

	// REQUIRED; 所有时间粒度的数据。
	TranscodeInfoDataList []DescribeLiveTranscodeInfoDataResResultTranscodeInfoDataListItem `json:"TranscodeInfoDataList"`

	// 查询流粒度数据时的应用名称。
	App *string `json:"App,omitempty"`

	// 域名列表。
	DomainList []*string `json:"DomainList,omitempty"`

	// 分辨率。- 480P：640 × 480； - 720P：1280 × 720； - 1080P：1920 × 1088； - 2K：2560 × 1440； - 4K：4096 × 2160；- 8K：大于4K； - 0P：纯音频流；
	Resolution []*string `json:"Resolution,omitempty"`

	// 查询流粒度数据时的流名称。
	Stream *string `json:"Stream,omitempty"`

	// 视频编码格式，支持的取值和含义如下所示。- NormalH264：H.264 标准转码； - NormalH265：H.265 标准转码； - NormalH266：H.266 标准转码； - ByteHDH264：H.264 极智超清；
	// - ByteHDH265：H.265 极智超清； - ByteHDH266：H.266 极智超清；- ByteQE：画质增强；- Audio：纯音频流；
	TranscodeType []*string `json:"TranscodeType,omitempty"`
}

// DescribeLiveTranscodeInfoDataResResultPagination - 分页信息
type DescribeLiveTranscodeInfoDataResResultPagination struct {

	// REQUIRED; 当前页数
	PageCur int32 `json:"PageCur"`

	// REQUIRED; 每页大小
	PageSize int32 `json:"PageSize"`

	// REQUIRED; 总个数
	TotalCount int32 `json:"TotalCount"`
}

type DescribeLiveTranscodeInfoDataResResultTranscodeInfoDataListItem struct {

	// REQUIRED; 转码时长，单位分钟
	Duration float32 `json:"Duration"`

	// REQUIRED; 结束转码时间
	EndTime string `json:"EndTime"`

	// REQUIRED; 分辨率。- 480P：640 × 480； - 720P：1280 × 720； - 1080P：1920 × 1088； - 2K：2560 × 1440； - 4K：4096 × 2160；- 8K：大于4K； -
	// 0P：纯音频流；
	Resolution string `json:"Resolution"`

	// REQUIRED; 开始转码时间
	StartTime string `json:"StartTime"`

	// REQUIRED; 流名
	Stream string `json:"Stream"`

	// REQUIRED; 转码任务ID
	TaskID string `json:"TaskID"`

	// REQUIRED; 转码后缀
	TranscodeSuffix string `json:"TranscodeSuffix"`

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

type DescribeRemoteAuthBody struct {

	// REQUIRED; 需要配置远程鉴权的拉流域名。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看拉流域名信息。
	Domain string `json:"Domain"`

	// REQUIRED; 需要配置远程鉴权的拉流域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看拉流域名的域名空间信息。
	Vhost string  `json:"Vhost"`
	App   *string `json:"App,omitempty"`
}

type DescribeRemoteAuthRes struct {

	// REQUIRED
	ResponseMetadata DescribeRemoteAuthResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result DescribeRemoteAuthResResult `json:"Result"`
}

type DescribeRemoteAuthResResponseMetadata struct {

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

type DescribeRemoteAuthResResult struct {

	// REQUIRED; 鉴权成功状态码，所有状态码范围在[200,499]间
	AllowStatus []int32 `json:"AllowStatus"`

	// REQUIRED; 远程鉴权地址
	AuthURL string `json:"AuthURL"`

	// REQUIRED; 鉴权缓存配置，配置均为空时表示不使用缓存来判断鉴权结果。
	CacheConfig DescribeRemoteAuthResResultCacheConfig `json:"CacheConfig"`

	// REQUIRED; 鉴权异常时是否返回鉴权失败
	DenyOnFailed bool `json:"DenyOnFailed"`

	// REQUIRED; 返回状态码不在成功失败中时是否返回失败
	DenyOtherStatus bool `json:"DenyOtherStatus"`

	// REQUIRED; 鉴权失败时返回的状态码，范围在[400,699]内
	DenyReturnCode float32 `json:"DenyReturnCode"`

	// REQUIRED; 鉴权失败时鉴权服务器返回的状态码，所有状态码范围应在 [200,499] 之间，且和 AllowStatus 不重复。
	DenyRule DescribeRemoteAuthResResultDenyRule `json:"DenyRule"`

	// REQUIRED; 远程鉴权配置是否打开
	Enable bool `json:"Enable"`

	// REQUIRED; 参数在50个以内
	HeaderParamConfig DescribeRemoteAuthResResultHeaderParamConfig `json:"HeaderParamConfig"`

	// REQUIRED; 鉴权请求方法，返回包括POST或GET
	Method string `json:"Method"`

	// REQUIRED; 参数在50个以内
	QueryParamConfig DescribeRemoteAuthResResultQueryParamConfig `json:"QueryParamConfig"`

	// 远程鉴权路径，UseUserRequest为true时忽略此字段
	AuthURLPath *string `json:"AuthURLPath,omitempty"`

	// 鉴权请求 Body 参数。 :::tip 鉴权请求方法为 POST 时 Body 参数配置生效。 :::
	BodyParams []*DescribeRemoteAuthResResultBodyParamsItem `json:"BodyParams,omitempty"`

	// 是否开启对 HLS 协议拉流请求的鉴权开关，取值及含义如下所示。
	// * true：开启。
	// * false：（默认值）关闭。
	EnableHLSAuth *bool `json:"EnableHLSAuth,omitempty"`

	// 是否开启对 HLS 协议流的 TS 分片进行远程鉴权，取值及含义如下所示。
	// * true：开启。
	// * false：（默认值）关闭。
	// :::tip 开启 HLS 协议拉流请求的远程鉴权开关时，必须开启对 TS 分片的远程鉴权。 :::
	EnableHLSTSAuth *bool `json:"EnableHLSTSAuth,omitempty"`

	// 请求头是否保留大小写信息，true为原样请求，false转换为mime格式
	HeaderCaseSensitivity *bool `json:"HeaderCaseSensitivity,omitempty"`

	// 长度不超过1024
	HeaderHost *string `json:"HeaderHost,omitempty"`

	// 鉴权失败重试间隔，单位为秒，范围[1,30]
	RetryInterval *float32 `json:"RetryInterval,omitempty"`

	// 鉴权失败重试次数，范围[0,10]
	RetryTimes *float32 `json:"RetryTimes,omitempty"`

	// 鉴权超时时间，单位为秒。范围[0, 600]
	Timeout *float32 `json:"Timeout,omitempty"`

	// 是否使用用户请求的路径拼接到鉴权地址上
	UseUserRequest *bool `json:"UseUserRequest,omitempty"`
}

type DescribeRemoteAuthResResultBodyParamsItem struct {

	// REQUIRED; 参数类型，取值及含义如下所示。
	// * const_string：常量；
	// * header：用户请求的 Header 参数；
	// * query：用户请求的 URL 参数；
	// * vhost：参数值为变量 vhost 的参数，表示拉流请求中拉流域名所属的域名空间；
	// * domain：参数值为变量 domain 的参数，表示拉流请求中使用的拉流域名；
	// * app：参数值为变量 app 的参数，表示拉流请求中使用的 AppName；
	// * stream：参数值为变量 stream 的参数，表示拉流请求中使用的 StreamName；
	// * client_ip：参数值为变量 client_ip 的参数，表示拉流客户端 IP 地址；
	// * server_ip：参数值为变量 server_ip 的参数，表示响应拉流请求的 CDN 节点IP地址；
	// * request_uri：参数值为变量 request_uri 的参数，拉流请求地址的 URI。
	Type string `json:"Type"`

	// 参数名，最大长度为 100 个字符，不支持输入空格。 :::tip
	// * 参数类型为常量时表示常量参数的参数名；
	// * 参数类型为用户请求的 Header 参数或用户请求的 URL 参数时，表示指定用户请求中对应的参数名作为此处的参数名；
	// * 参数类型为变量时不生效。 :::
	ParamName *string `json:"ParamName,omitempty"`

	// 参数名的映射参数名，最大长度为 100 个字符，不支持输入空格。 :::tip
	// * 参数类型为常量时不生效；
	// * 参数类型为用户请求的 Header 参数或用户请求的 URL 参数时，表示鉴权请求时使用 ToName 值代替用户请求中对应的参数名；
	// * 参数类型为变量时，表示使用 ToName 取值作为此变量的参数名。 :::
	ToName *string `json:"ToName,omitempty"`

	// 参数类型为常量时的参数值，最大长度为 100 个字符，不支持输入空格。
	Value *string `json:"Value,omitempty"`
}

// DescribeRemoteAuthResResultCacheConfig - 鉴权缓存配置，配置均为空时表示不使用缓存来判断鉴权结果。
type DescribeRemoteAuthResResultCacheConfig struct {

	// REQUIRED; 生成鉴权结果缓存 Key 使用的参数配置，最多支持配置 50 个参数。 :::tip
	// * 鉴权成功与鉴权失败使用相同配置时，表示鉴权结果缓存 Key 配置；
	// * 鉴权成功与鉴权失败不使用相同配置时，表示鉴权成功缓存的过期时间。 :::
	CacheKeys DescribeRemoteAuthResResultCacheConfigCacheKeys `json:"CacheKeys"`

	// REQUIRED; 鉴权成功和鉴权失败是否使用相同的配置（包括缓存的 key 值和缓存过期时间）。
	// * true：使用相同配置；
	// * false：不使用相同配置。
	UseSameCache bool `json:"UseSameCache"`

	// 缓存过期时间，范围为[0,3600]
	CacheExpireSecond *float32 `json:"CacheExpireSecond,omitempty"`

	// 鉴权失败缓存的过期时间，默认值为 0 时表示不缓存，单位为秒，取值范围为 [0,3600]。 :::tip 鉴权成功与鉴权失败不使用相同配置，即 UseSameCache 配置 false 时生效。 :::
	DenyExpireSecond *float32 `json:"DenyExpireSecond,omitempty"`

	// 生成鉴权失败结果缓存 Key 使用的参数配置，最多支持配置 50 个参数。
	// :::tip 鉴权成功与鉴权失败不使用相同配置，即 UseSameCache 配置 false 时生效。 :::
	DenyKeys *DescribeRemoteAuthResResultCacheConfigDenyKeys `json:"DenyKeys,omitempty"`
}

// DescribeRemoteAuthResResultCacheConfigCacheKeys - 生成鉴权结果缓存 Key 使用的参数配置，最多支持配置 50 个参数。 :::tip
// * 鉴权成功与鉴权失败使用相同配置时，表示鉴权结果缓存 Key 配置；
// * 鉴权成功与鉴权失败不使用相同配置时，表示鉴权成功缓存的过期时间。 :::
type DescribeRemoteAuthResResultCacheConfigCacheKeys struct {

	// REQUIRED; 生成缓存 Key 使用的参数的类型，取值及含义如下所示。
	// * query：用户拉流请求 URL 参数；
	// * header：用户拉流请求 Header 参数；
	// * vhost：使用用户拉流域名所属的 vhost 值作为参数；
	// * domain：使用用户拉流域名值作为参数；
	// * app：使用用户拉流请求中的 AppName 值作为参数；
	// * stream：使用用户拉流请求中的 StreamName 值作为参数；
	// * client_ip：使用拉流用户客户端的 IP 地址作为参数；
	// * request_uri：使用拉流请求地址的 URI 做为参数。
	Type string `json:"Type"`

	// 参数名，当参数类型为 query 或 Header 时生效且必填，表示指定拉流请求中的参数名对应的参数值作为生成缓存 Key 的参数。
	ParamName *string `json:"ParamName,omitempty"`
}

// DescribeRemoteAuthResResultCacheConfigDenyKeys - 生成鉴权失败结果缓存 Key 使用的参数配置，最多支持配置 50 个参数。
// :::tip 鉴权成功与鉴权失败不使用相同配置，即 UseSameCache 配置 false 时生效。 :::
type DescribeRemoteAuthResResultCacheConfigDenyKeys struct {

	// REQUIRED
	Type      string  `json:"Type"`
	ParamName *string `json:"ParamName,omitempty"`
}

// DescribeRemoteAuthResResultDenyRule - 鉴权失败时鉴权服务器返回的状态码，所有状态码范围应在 [200,499] 之间，且和 AllowStatus 不重复。
type DescribeRemoteAuthResResultDenyRule struct {

	// REQUIRED; 鉴权失败状态码，所有状态码范围在[200,499]内
	DenyStatus []float32 `json:"DenyStatus"`
}

// DescribeRemoteAuthResResultHeaderParamConfig - 参数在50个以内
type DescribeRemoteAuthResResultHeaderParamConfig struct {

	// REQUIRED
	UseUserParam bool                                                      `json:"UseUserParam"`
	Params       []*DescribeRemoteAuthResResultHeaderParamConfigParamsItem `json:"Params,omitempty"`
}

type DescribeRemoteAuthResResultHeaderParamConfigParamsItem struct {

	// REQUIRED
	Type      string  `json:"Type"`
	ParamName *string `json:"ParamName,omitempty"`
	ToName    *string `json:"ToName,omitempty"`
	Value     *string `json:"Value,omitempty"`
}

// DescribeRemoteAuthResResultQueryParamConfig - 参数在50个以内
type DescribeRemoteAuthResResultQueryParamConfig struct {

	// REQUIRED; 远程鉴权请求的 URL 或 Header 参数使用用户请求参数还是自定义参数，取值及含义如下所示。
	// * true：使用用户请求参数；
	// * false：自定义参数。
	UseUserParam bool `json:"UseUserParam"`

	// 自定义参数时的参数配置。
	Params []*DescribeRemoteAuthResResultQueryParamConfigParamsItem `json:"Params,omitempty"`
}

type DescribeRemoteAuthResResultQueryParamConfigParamsItem struct {

	// REQUIRED
	Type      string  `json:"Type"`
	ParamName *string `json:"ParamName,omitempty"`
	ToName    *string `json:"ToName,omitempty"`
	Value     *string `json:"Value,omitempty"`
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

	// REQUIRED; 0: response 1: request
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

	// REQUIRED; 生成地址为源流地址/转码流地址还是abr地址
	StreamType string `json:"StreamType"`

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

	// 拉流地址的有效时长，单位为秒，超过有效时长后需要重新生成。取值范围为正整数，缺省值为 604800，即 7 天。 :::tip 如果同时设置 ValidDuration 和 ExpiredTime，以 ExpiredTime 的时间为准。
	// :::
	ValidDuration *int32 `json:"ValidDuration,omitempty"`
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

	// REQUIRED; 生成地址对应匹配到的鉴权模版类型
	AuthType string `json:"AuthType"`

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

	// REQUIRED; 子流地址。仅当 StreamType 为 abr 时返回。
	SubStreamURL []GeneratePlayURLResResultURLListPropertiesItemsItem `json:"SubStreamURL"`

	// REQUIRED; 地址类型，取值及含义如下所示。
	// * pull：拉流地址；
	// * 3rd_play(relay_source)：第三方回源地址，当配置了回源且 CDN 类型为第三方 CDN 时返回；
	// * 3rd_play(relay_sink)：第三方转推地址，当配置了拉流转推且 CDN 类型为第三方 CDN 时返回。
	Type string `json:"Type"`

	// REQUIRED; 生成的拉流地址。
	URL string `json:"URL"`
}

type GeneratePlayURLResResultURLListPropertiesItemsItem struct {

	// REQUIRED; 子流转码后缀。
	Suffix string `json:"Suffix"`

	// REQUIRED; 地址标签。包括 drm、hls加密等。
	Tag string `json:"Tag"`
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

	// 推流地址的有效时长，单位为秒，超过有效时长后需要重新生成。取值范围为正整数，默认值为 604800，即 7 天。 :::tip 如果同时设置 ValidDuration 和 ExpiredTime，以 ExpiredTime 的时间为准。
	// :::
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

	// REQUIRED; 生成地址对应匹配到的鉴权模版类型
	AuthType string `json:"AuthType"`

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

type GetCarouselDetailBody struct {

	// REQUIRED; 待查询的轮播任务 ID，任务的唯一标识。调用 CreateCarouselTask 接口创建轮播任务时返回。
	TaskID string `json:"TaskID"`
}

type GetCarouselDetailRes struct {

	// REQUIRED
	ResponseMetadata GetCarouselDetailResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result GetCarouselDetailResResult `json:"Result"`
}

type GetCarouselDetailResResponseMetadata struct {

	// REQUIRED
	RequestID string `json:"RequestID"`
}

type GetCarouselDetailResResult struct {

	// REQUIRED; 包含轮播任务相关信息的数据对象。
	Data GetCarouselDetailResResultData `json:"Data"`
}

// GetCarouselDetailResResultData - 包含轮播任务相关信息的数据对象。
type GetCarouselDetailResResultData struct {

	// REQUIRED; 最新的播放列表序列号
	LastOperationIndex int32 `json:"LastOperationIndex"`

	// REQUIRED; 当前播放列表序列号
	LastSuccessOperationIndex int32 `json:"LastSuccessOperationIndex"`

	// REQUIRED; 当前的播放信息，json字符串
	PlayInfo string `json:"PlayInfo"`

	// REQUIRED; 当前的播单信息
	Rule string `json:"Rule"`

	// REQUIRED; 任务状态： pending：任务等待调度中 prepare：任务初始化中 running：任务运行中 prestop：任务停止中 done：任务已经停止
	Status string `json:"Status"`
}

type GetCloudMixTaskDetailBody struct {

	// REQUIRED; 混流任务 ID，您可以通过 ListCloudMixTask [https://www.volcengine.com/docs/6469/1271157] 接口获取混流任务 ID。
	TaskID string `json:"TaskID"`
}

type GetCloudMixTaskDetailRes struct {

	// REQUIRED
	ResponseMetadata GetCloudMixTaskDetailResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result GetCloudMixTaskDetailResResult `json:"Result"`
}

type GetCloudMixTaskDetailResResponseMetadata struct {

	// REQUIRED
	RequestID string `json:"RequestID"`
}

type GetCloudMixTaskDetailResResult struct {

	// REQUIRED; 请求响应码，取值及含义如下。
	// * 0：请求成功；
	// * 500：内部处理错误；
	// * 400：请求异常。
	Code int32 `json:"Code"`

	// REQUIRED; 返回的数据。
	Data GetCloudMixTaskDetailResResultData `json:"Data"`

	// REQUIRED; 请求响应码对应的信息。
	Message string `json:"Message"`
}

// GetCloudMixTaskDetailResResultData - 返回的数据。
type GetCloudMixTaskDetailResResultData struct {

	// REQUIRED
	AccountID string `json:"AccountID"`

	// REQUIRED
	LastOperationErrCode int32 `json:"LastOperationErrCode"`

	// REQUIRED
	LastOperationErrMsg string `json:"LastOperationErrMsg"`

	// REQUIRED; 任务最近一次更新的版本标识。
	LastOperationIndex int32 `json:"LastOperationIndex"`

	// REQUIRED; 任务最近一次成功更新的版本标识。
	LastSuccessOperationIndex int32 `json:"LastSuccessOperationIndex"`

	// REQUIRED; 混流任务详细配置的 Json 字符串。
	Rule string `json:"Rule"`

	// REQUIRED; 混流任务状态，取值及含义如下所示。
	// * pending：任务调度被阻塞。
	// * prepare：正在准备任务资源。
	// * runing：任务进行中。
	// * prestop：正在清理任务资源。
	// * done：任务已结束。
	Status string `json:"Status"`

	// REQUIRED
	TaskErrCode int32 `json:"TaskErrCode"`

	// REQUIRED
	TaskErrMsg string `json:"TaskErrMsg"`

	// REQUIRED
	TaskErrSrcIDs string `json:"TaskErrSrcIDs"`

	// REQUIRED; 混流任务 ID。
	TaskID string `json:"TaskID"`
}

type GetHLSEncryptDataKeyQuery struct {

	// REQUIRED; 视频直播服务端生成的 M3U8 文件中写入的每个 TS 分片的密钥 ID。
	KeyID string `json:"KeyID" query:"KeyID"`
}

type GetHLSEncryptDataKeyRes struct {

	// REQUIRED
	ResponseMetadata GetHLSEncryptDataKeyResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *GetHLSEncryptDataKeyResResult `json:"Result,omitempty"`
}

type GetHLSEncryptDataKeyResResponseMetadata struct {

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

// GetHLSEncryptDataKeyResResult - 视请求的接口而定
type GetHLSEncryptDataKeyResResult struct {

	// REQUIRED; 密钥。
	DataKey string `json:"DataKey"`
}

type GetLiveVideoQualityAnalysisTaskDetailBody struct {

	// 查询的任务 ID。 :::tipName 和 ID 二选一必填。 :::
	ID *int64 `json:"ID,omitempty"`

	// 查询的任务名称。 :::tipName 和 ID 二选一必填。 :::
	Name *string `json:"Name,omitempty"`
}

type GetLiveVideoQualityAnalysisTaskDetailRes struct {

	// REQUIRED
	ResponseMetadata GetLiveVideoQualityAnalysisTaskDetailResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result GetLiveVideoQualityAnalysisTaskDetailResResult `json:"Result"`
}

type GetLiveVideoQualityAnalysisTaskDetailResResponseMetadata struct {

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

type GetLiveVideoQualityAnalysisTaskDetailResResult struct {

	// REQUIRED; 测评任务详细信息。
	Task GetLiveVideoQualityAnalysisTaskDetailResResultTask `json:"Task"`
}

// GetLiveVideoQualityAnalysisTaskDetailResResultTask - 测评任务详细信息。
type GetLiveVideoQualityAnalysisTaskDetailResResultTask struct {

	// REQUIRED; 测试任务的持续时长，单位为秒。
	Duration int32 `json:"Duration"`

	// REQUIRED; 请提供具体的参数ID和类型string，以便我为您生成参数描述。
	ID int64 `json:"ID"`

	// REQUIRED; 画质测评的打点间隔。
	Interval int32 `json:"Interval"`

	// REQUIRED; 任务名称。
	Name string `json:"Name"`

	// REQUIRED; 画质测评结果。
	ScoringResult GetLiveVideoQualityAnalysisTaskDetailResResultTaskScoringResult `json:"ScoringResult"`

	// REQUIRED; 画质测评视频流的播放地址。
	SrcURL string `json:"SrcURL"`
}

// GetLiveVideoQualityAnalysisTaskDetailResResultTaskScoringResult - 画质测评结果。
type GetLiveVideoQualityAnalysisTaskDetailResResultTaskScoringResult struct {

	// REQUIRED; 画质测评结果详细信息。
	VQScoreLive []GetLiveVideoQualityAnalysisTaskDetailResResultTaskScoringResultVQScoreLiveItem `json:"VQScoreLive"`
}

type GetLiveVideoQualityAnalysisTaskDetailResResultTaskScoringResultVQScoreLiveItem struct {

	// REQUIRED; 测评打点的时间，Unix 时间戳，精度为秒。
	Timestamp int32 `json:"Timestamp"`

	// REQUIRED; 测评点的画质得分，画质评分范围为 0 到 100，评分越高表示视频画面色彩越好。不同的评分段对应不同的视频质量感受：
	// * 0～60：主观感受较差。
	// * 60～70：主观感受良好。
	// * 70～100：主观感受清晰。
	Value float32 `json:"Value"`
}

type KillStreamBody struct {

	// REQUIRED; 直播流使用的应用名称。
	App string `json:"App"`

	// REQUIRED; 直播流使用的流名称。
	Stream string `json:"Stream"`

	// REQUIRED; 域名空间，您可以调用 DescribeLiveStreamInfoByPage [https://www.volcengine.com/docs/6469/1126841] 接口，查看待断开的在线流的信息，包括 Vhost、APP
	// 和 Stream。
	Vhost string `json:"Vhost"`

	// 推流域名。 参数 Domain 和 Vhost传且仅传一个。
	Domain *string `json:"Domain,omitempty"`

	// 禁推的结束时间，禁推有效期最长为 90 天，默认为当前时间加 90 天
	EndTime *string `json:"EndTime,omitempty"`
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

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED; 加密类型，支持的取值及含义如下所示。
	// * FairPlay：使用 FairPlay 技术的商业 DRM 加密；
	// * Widevine：使用 Widevine 技术的商业 DRM 加密；
	// * PlayReady：使用 PlayReady 技术的商业 DRM 加密；
	// * ClearKey：HLS 标准加密。
	// :::tip DRM 加密与 HLS 标准加密不可同时配置。 :::
	DRMSystems []string `json:"DRMSystems"`

	// REQUIRED; 是否开启源流加密，取值及含义如下所示。
	// * true：开启；
	// * fasle：不开启。
	EncryptOriginStream bool `json:"EncryptOriginStream"`

	// REQUIRED; 是否开启转码流加密，取值及含义如下所示。
	// * true：开启；
	// * fasle：不开启。
	EncryptTranscodeStream bool `json:"EncryptTranscodeStream"`

	// REQUIRED; 进行 DRM 加密的转码流对应的转码流后缀配置。
	EncryptTranscodeSuffix []string `json:"EncryptTranscodeSuffix"`

	// REQUIRED; 域名空间。
	Vhost string `json:"Vhost"`
}

type ListCarouselTaskBody struct {

	// REQUIRED; 分页功能，展示第几页
	Page int32 `json:"Page"`

	// REQUIRED; 分页功能，页大小
	PageSize int32 `json:"PageSize"`
}

type ListCarouselTaskRes struct {

	// REQUIRED
	ResponseMetadata ListCarouselTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result ListCarouselTaskResResult `json:"Result"`
}

type ListCarouselTaskResResponseMetadata struct {

	// REQUIRED
	RequestID string `json:"RequestID"`
}

type ListCarouselTaskResResult struct {

	// REQUIRED; 轮播任务数据对象。
	Data ListCarouselTaskResResultData `json:"Data"`
}

// ListCarouselTaskResResultData - 轮播任务数据对象。
type ListCarouselTaskResResultData struct {

	// REQUIRED; 满足查询条件的轮播任务总数。
	Count int32 `json:"Count"`

	// REQUIRED; 轮播任务的数组，每个元素表示一个任务的详细信息。
	Result []ListCarouselTaskResResultDataResultItem `json:"Result"`
}

type ListCarouselTaskResResultDataResultItem struct {

	// REQUIRED; 任务的创建时间，RFC3339 格式的时间戳，精度为秒。
	CreatedAt ListCarouselTaskResResultDataResultItemCreatedAt `json:"CreatedAt"`

	// REQUIRED; 轮播任务名称。
	Name string `json:"Name"`

	// REQUIRED; 轮播任务的当前状态。取值和含义如下：
	// * pending：任务等待调度中；
	// * prepare：任务初始化中；
	// * running：任务运行中；
	// * prestop：任务停止中；
	// * done：任务已经停止。
	Status string `json:"Status"`

	// REQUIRED; 任务的结束时间，RFC3339 格式的时间戳，精度为秒。
	StoppedAt ListCarouselTaskResResultDataResultItemStoppedAt `json:"StoppedAt"`

	// REQUIRED; 轮播任务的唯一标识。
	TaskID string `json:"TaskID"`

	// REQUIRED; 任务的更新时间，RFC3339 格式的时间戳，精度为秒。
	UpdatedAt ListCarouselTaskResResultDataResultItemUpdatedAt `json:"UpdatedAt"`
}

// ListCarouselTaskResResultDataResultItemCreatedAt - 任务的创建时间，RFC3339 格式的时间戳，精度为秒。
type ListCarouselTaskResResultDataResultItemCreatedAt struct {

	// REQUIRED; 任务的创建时间，RFC3339 格式的时间戳，精度为秒。
	Time string `json:"Time"`
}

// ListCarouselTaskResResultDataResultItemStoppedAt - 任务的结束时间，RFC3339 格式的时间戳，精度为秒。
type ListCarouselTaskResResultDataResultItemStoppedAt struct {

	// REQUIRED; 任务的结束时间，RFC3339 格式的时间戳，精度为秒。
	Time string `json:"Time"`
}

// ListCarouselTaskResResultDataResultItemUpdatedAt - 任务的更新时间，RFC3339 格式的时间戳，精度为秒。
type ListCarouselTaskResResultDataResultItemUpdatedAt struct {

	// REQUIRED; 任务的更新时间，RFC3339 格式的时间戳，精度为秒。
	Time string `json:"Time"`
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

	// 页码。不填默认返回全部。
	PageNum *int32 `json:"PageNum,omitempty"`

	// 分页大小。不填默认返回所有。
	PageSize *int32 `json:"PageSize,omitempty"`

	// 项目名称。
	ProjectName *string `json:"ProjectName,omitempty"`
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

	// REQUIRED; 本次查询所有证书的过期信息。
	ExpirationInfo ListCertV2ResResultExpirationInfo `json:"ExpirationInfo"`

	// 证书列表。
	CertList []*ListCertV2ResResultCertListItem `json:"CertList,omitempty"`

	// 总数。
	Total *int32 `json:"Total,omitempty"`
}

type ListCertV2ResResultCertListItem struct {

	// REQUIRED; 与证书绑定的域名列表。
	CertDomainList []string `json:"CertDomainList"`

	// REQUIRED; 证书 ID。
	CertID string `json:"CertID"`

	// REQUIRED; 证书名称。
	CertName string `json:"CertName"`

	// REQUIRED; 证书链 ID。
	ChainID string `json:"ChainID"`

	// REQUIRED; 火山引擎证书中心证书链 ID。
	ChainIDVolc string `json:"ChainIDVolc"`

	// REQUIRED; 创建时间。
	CreateTime string `json:"CreateTime"`

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

// ListCertV2ResResultExpirationInfo - 本次查询所有证书的过期信息。
type ListCertV2ResResultExpirationInfo struct {

	// REQUIRED; 生效数量。
	ActiveNum int32 `json:"ActiveNum"`

	// REQUIRED; 快要过期数量，一个月之内
	ClosingExpireNum int32 `json:"ClosingExpireNum"`

	// REQUIRED; 过期数量。
	ExpireNum int32 `json:"ExpireNum"`
}

type ListCloudMixTaskBody struct {

	// REQUIRED; 查询数据的页码。
	Page int32 `json:"Page"`

	// REQUIRED; 每页显示的数据条数，最大值为 100。
	PageSize int32 `json:"PageSize"`
}

type ListCloudMixTaskRes struct {

	// REQUIRED
	ResponseMetadata ListCloudMixTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result ListCloudMixTaskResResult `json:"Result"`
}

type ListCloudMixTaskResResponseMetadata struct {

	// REQUIRED
	RequestID string `json:"RequestID"`
}

type ListCloudMixTaskResResult struct {

	// REQUIRED; 请求响应码，取值及含义如下。
	// * 0：请求成功；
	// * 500：内部处理错误；
	// * 400：请求异常。
	Code int32 `json:"Code"`

	// REQUIRED; 返回的数据。
	Data ListCloudMixTaskResResultData `json:"Data"`

	// REQUIRED; 请求响应码对应的信息。
	Message string `json:"Message"`
}

// ListCloudMixTaskResResultData - 返回的数据。
type ListCloudMixTaskResResultData struct {

	// REQUIRED; 查询结果的数据总条数。
	Count int32 `json:"Count"`

	// REQUIRED; 查询结果数据详细信息。
	Result []ListCloudMixTaskResResultDataResultItem `json:"Result"`
}

type ListCloudMixTaskResResultDataResultItem struct {

	// REQUIRED
	AccountID string `json:"AccountID"`

	// REQUIRED
	CloudcastID string `json:"CloudcastID"`

	// REQUIRED; 混流任务创建时间。
	CreatedAt ListCloudMixTaskResResultDataResultItemCreatedAt `json:"CreatedAt"`

	// REQUIRED
	MesosID string `json:"MesosID"`

	// REQUIRED; 混流任务名称。
	Name string `json:"Name"`

	// REQUIRED
	Provider string `json:"Provider"`

	// REQUIRED; 混流任务状态，取值及含义如下所示。
	// * pending：任务调度被阻塞。
	// * prepare：正在准备任务资源。
	// * runing：任务进行中。
	// * prestop：正在清理任务资源。
	// * done：任务已结束。
	Status string `json:"Status"`

	// REQUIRED; 混流任务结束时间。
	StoppedAt ListCloudMixTaskResResultDataResultItemStoppedAt `json:"StoppedAt"`

	// REQUIRED; 混流任务 ID。
	TaskID string `json:"TaskID"`

	// REQUIRED
	TaskType string `json:"TaskType"`

	// REQUIRED
	UID string `json:"UID"`

	// REQUIRED; 混流任务更新时间。
	UpdatedAt ListCloudMixTaskResResultDataResultItemUpdatedAt `json:"UpdatedAt"`
}

// ListCloudMixTaskResResultDataResultItemCreatedAt - 混流任务创建时间。
type ListCloudMixTaskResResultDataResultItemCreatedAt struct {

	// REQUIRED; 时间。
	Time string `json:"Time"`
}

// ListCloudMixTaskResResultDataResultItemStoppedAt - 混流任务结束时间。
type ListCloudMixTaskResResultDataResultItemStoppedAt struct {

	// REQUIRED
	Time string `json:"Time"`
}

// ListCloudMixTaskResResultDataResultItemUpdatedAt - 混流任务更新时间。
type ListCloudMixTaskResResultDataResultItemUpdatedAt struct {

	// REQUIRED
	Time string `json:"Time"`
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

	// 视频高度，单位为 px。 :::tip
	// * 当关闭视频分辨率自适应（As 取值为 0）时，转码分辨率将取 Width 和 Height 的值对转码视频进行拉伸；
	// * 当关闭视频分辨率自适应（As 取值为 0）时，Width 和 Height 任一取值为 0 时，转码视频将保持源流尺寸；
	// * 编码格式为 H.266 时，不支持设置 Width 和 Height，请使用自适应配置。 :::
	Height *int32 `json:"Height,omitempty"`

	// 长边长度，单位为 px。 :::tip
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

	// 短边长度，单位为 px。 :::tip
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

	// 视频宽度，单位为 px。 :::tip
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
	// * 5：欠费关停；
	// * 6：域名未备案被封禁。
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

	// REQUIRED; HTTP/2协议。
	HTTP2 bool `json:"HTTP2"`

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
	// * 5：欠费关停；
	// * 6：域名未备案被封禁。
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

type ListLiveVideoQualityAnalysisTasksBody struct {

	// REQUIRED; 分页参数
	PageNum int32 `json:"PageNum"`

	// REQUIRED; 分页参数
	PageSize int32 `json:"PageSize"`

	// 查询的任务ID列表， 和Names二选一
	IDs []*int64 `json:"IDs,omitempty"`

	// 查询的任务名称列表， 和TaskIDs二选一
	Names []*string `json:"Names,omitempty"`
}

type ListLiveVideoQualityAnalysisTasksRes struct {

	// REQUIRED
	ResponseMetadata ListLiveVideoQualityAnalysisTasksResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result ListLiveVideoQualityAnalysisTasksResResult `json:"Result"`
}

type ListLiveVideoQualityAnalysisTasksResResponseMetadata struct {

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

type ListLiveVideoQualityAnalysisTasksResResult struct {

	// REQUIRED; 查询的数据的页码。
	PageNum int32 `json:"PageNum"`

	// REQUIRED; 每页显示的数据条数。
	PageSize int32 `json:"PageSize"`

	// REQUIRED; 画质测评任务列表。
	Tasks []ListLiveVideoQualityAnalysisTasksResResultTasksItem `json:"Tasks"`
}

type ListLiveVideoQualityAnalysisTasksResResultTasksItem struct {

	// 测评任务持续时长。
	Duration *int32 `json:"Duration,omitempty"`

	// 任务 ID。
	ID *int64 `json:"ID,omitempty"`

	// 画质测评的打点间隔。
	Interval *int32 `json:"Interval,omitempty"`

	// 任务名称。
	Name *string `json:"Name,omitempty"`

	// 进行画质测评的直播流地址。
	SrcURL *string `json:"SrcURL,omitempty"`
}

type ListPullToPushGroupBody struct {

	// REQUIRED; 查询数据的页码，取值范围为 [1,1000]。
	PageNum int32 `json:"PageNum"`

	// REQUIRED; 每页现实的数据条数，取值范围为 [1,1000]。
	PageSize int32 `json:"PageSize"`

	// 群组的状态，取值及含义如下所示。
	// * 0: （默认值）可用;
	// * 1: 已删除，不可用。
	StatusList []*int32 `json:"StatusList,omitempty"`

	// 标签过滤参数。
	TagFilters []*ListPullToPushGroupBodyTagFiltersItem `json:"TagFilters,omitempty"`
}

type ListPullToPushGroupBodyTagFiltersItem struct {

	// REQUIRED; 标签Key。
	Key string `json:"Key"`

	// 标签Value。
	Values []*string `json:"Values,omitempty"`
}

type ListPullToPushGroupRes struct {

	// REQUIRED
	ResponseMetadata ListPullToPushGroupResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result *ListPullToPushGroupResResult `json:"Result,omitempty"`
}

type ListPullToPushGroupResResponseMetadata struct {

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

// ListPullToPushGroupResResult - 视请求的接口而定
type ListPullToPushGroupResResult struct {

	// REQUIRED; 拉流转推群组列表。
	List []ListPullToPushGroupResResultListItem `json:"List"`

	// REQUIRED; 查询结果的数据条数。
	Total int32 `json:"Total"`
}

type ListPullToPushGroupResResultListItem struct {

	// REQUIRED; 账号。
	AccountID string `json:"AccountID"`

	// REQUIRED; 群组名称。
	Name string `json:"Name"`

	// REQUIRED; 群组所属的项目名称。
	ProjectName string `json:"ProjectName"`

	// REQUIRED; 群组的状态，取值及含义如下所示。
	// * 0: 可用;
	// * 1: 已删除，不可用。
	Status int32 `json:"Status"`

	// REQUIRED; 群组的标签信息。
	Tags []ListPullToPushGroupResResultListPropertiesItemsItem `json:"Tags"`
}

type ListPullToPushGroupResResultListPropertiesItemsItem struct {

	// REQUIRED; 标签类型，支持以下取值。
	// * System：系统内置标签；
	// * Custom：自定义标签。
	Category string `json:"Category"`

	// REQUIRED; 标签 Key 值。
	Key string `json:"Key"`

	// REQUIRED; 标签 Value 值。
	Value string `json:"Value"`
}

type ListPullToPushTaskQuery struct {

	// 群组名称。
	// * 使用主账号调用时，为非必填，默认为空，表示查询所有群组的任务信息。
	// * 使用子账号调用时，非必填。
	GroupName *string `json:"GroupName,omitempty" query:"GroupName"`

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

	// 任务所属的群组名称，您可以调用 ListPullToPushGroup [https://www.volcengine.com/docs/6469/1327382] 获取可用的群组。 :::tip
	// * 使用主账号调用时，为非必填，默认为空表示查询所有群组的任务列表。
	// * 使用子账号调用时，为必填。 :::
	GroupName *string `json:"GroupName,omitempty"`

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

	// 点播文件地址和开始播放、结束播放的时间设置。 :::tip
	// * 当 Type 为点播类型时配置生效。
	// * 与 SrcAddrS 和 OffsetS 字段不可同时填写。 :::
	VodSrcAddrs []*ComponentsGg7M1TSchemasListpulltopushtaskresPropertiesResultPropertiesListItemsPropertiesVodsrcaddrsItems `json:"VodSrcAddrs,omitempty"`

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

type ListPullToPushTaskV2Body struct {

	// 群组名称。
	GroupNames []*string `json:"GroupNames,omitempty"`

	// 查询数据的页码，默认为 1，表示查询第一页的数据。
	Page *int32 `json:"Page,omitempty"`

	// 每页显示的数据条数，默认为 20，最大值为 500。
	Size *int32 `json:"Size,omitempty"`

	// 拉流转推任务的名称，不区分大小写，支持模糊查询。 例如，title取值为doc时，则返回任务名称为docspace、docs、DOC等 title 中包含doc关键词的所有任务列表。
	Title *string `json:"Title,omitempty"`
}

type ListPullToPushTaskV2Res struct {

	// REQUIRED
	ResponseMetadata ListPullToPushTaskV2ResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result ListPullToPushTaskV2ResResult `json:"Result"`
}

type ListPullToPushTaskV2ResResponseMetadata struct {
	Action    *string                                       `json:"Action,omitempty"`
	Error     *ListPullToPushTaskV2ResResponseMetadataError `json:"Error,omitempty"`
	Region    *string                                       `json:"Region,omitempty"`
	RequestID *string                                       `json:"RequestId,omitempty"`
	Service   *string                                       `json:"Service,omitempty"`
	Version   *string                                       `json:"Version,omitempty"`
}

type ListPullToPushTaskV2ResResponseMetadataError struct {
	Code    *string `json:"Code,omitempty"`
	Message *string `json:"Message,omitempty"`
}

type ListPullToPushTaskV2ResResult struct {

	// 任务列表。
	List []*ListPullToPushTaskV2ResResultListItem `json:"List,omitempty"`

	// 分页数量信息。
	Pagination *ListPullToPushTaskV2ResResultPagination `json:"Pagination,omitempty"`
}

type ListPullToPushTaskV2ResResultListItem struct {

	// 接收拉流转推任务状态回调的地址。
	CallbackURL *string `json:"CallbackURL,omitempty"`

	// 续播策略，续播策略指转推点播视频进行直播时出现断流并恢复后，如何继续播放的策略，拉流来源类型为点播视频时参数生效，支持的取值及含义如下。
	// * 0：从断流处续播（默认值）；
	// * 1：从断流处+自然流逝时长处续播。
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

	// 任务所属的群组名称，您可以调用 ListPullToPushGroup [https://www.volcengine.com/docs/6469/1327382] 获取可用的群组。 :::tip
	// * 使用主账号调用时，为非必填，默认为空表示查询所有群组的任务列表。
	// * 使用子账号调用时，为必填。 :::
	GroupName *string `json:"GroupName,omitempty"`

	// 点播文件启播时间偏移值，单位为秒，数量与拉流地址列表中地址数量相等，缺省情况下为空表示不进行偏移。拉流来源类型为点播视频时，参数生效。
	OffsetS []*float32 `json:"OffsetS,omitempty"`

	// 点播视频文件循环播放次数，当循环播放模式为有限次循环（CycleMode为0）时配置生效。
	PlayTimes *int32 `json:"PlayTimes,omitempty"`

	// 是否开启点播预热，开启点播预热后，系统会自动将点播视频文件缓存到 CDN 节点上，当用户请求直播时，可以直播从 CDN 节点获取视频，从而提高直播流畅度。拉流来源类型为点播视频时，参数生效。
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

	// 点播文件地址和开始播放、结束播放的时间设置。 :::tip
	// * 当 Type 为点播类型时配置生效。
	// * 与 SrcAddrS 和 OffsetS 字段不可同时填写。 :::
	VodSrcAddrs []*Components1Nf1A8CSchemasListpulltopushtaskv2ResPropertiesResultPropertiesListItemsPropertiesVodsrcaddrsItems `json:"VodSrcAddrs,omitempty"`

	// 为拉流转推视频添加的水印配置信息。
	Watermark *ListPullToPushTaskV2ResResultListItemWatermark `json:"Watermark,omitempty"`
}

// ListPullToPushTaskV2ResResultListItemWatermark - 为拉流转推视频添加的水印配置信息。
type ListPullToPushTaskV2ResResultListItemWatermark struct {

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

// ListPullToPushTaskV2ResResultPagination - 分页数量信息。
type ListPullToPushTaskV2ResResultPagination struct {

	// 当前任务所在分页。
	PageCur *int32 `json:"PageCur,omitempty"`

	// 每页显示的数据条数。
	PageSize *int32 `json:"PageSize,omitempty"`

	// 查询结果的数据总页数。
	PageTotal *int32 `json:"PageTotal,omitempty"`

	// 查询结果的数据总条数。
	TotalCount *int32 `json:"TotalCount,omitempty"`
}

type ListTimeShiftPresetV2Body struct {

	// REQUIRED; 时移类型，默认类型为 vod。
	// * vod：点播时移，表示查询时移录制存储在 VOD 中的时移配置；
	// * tos：直播时移，表示查询时移录制存储在 TOS 以及 fcdn-tos 中的时移配置。
	Type string `json:"Type"`

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

type ListVhostRemoteAuthBody struct {

	// REQUIRED; 需要配置远程鉴权的拉流域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看拉流域名的域名空间信息。
	Vhost string `json:"Vhost"`
}

type ListVhostRemoteAuthRes struct {

	// REQUIRED
	ResponseMetadata ListVhostRemoteAuthResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result ListVhostRemoteAuthResResult `json:"Result"`
}

type ListVhostRemoteAuthResResponseMetadata struct {

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

type ListVhostRemoteAuthResResult struct {

	// REQUIRED; Vhost 下所有拉流域名的远程鉴权配置信息。
	RemoteAuthConfigList []ListVhostRemoteAuthResResultRemoteAuthConfigListItem `json:"RemoteAuthConfigList"`
}

type ListVhostRemoteAuthResResultRemoteAuthConfigListItem struct {

	// REQUIRED; 鉴权成功状态码，所有状态码范围在[200,499]间
	AllowStatus []int32 `json:"AllowStatus"`

	// REQUIRED; 配置所属的app
	App string `json:"App"`

	// REQUIRED; 远程鉴权地址
	AuthURL string `json:"AuthURL"`

	// REQUIRED; 鉴权缓存配置，配置均为空时表示不使用缓存来判断鉴权结果。
	CacheConfig ListVhostRemoteAuthResResultRemoteAuthConfigListItemCacheConfig `json:"CacheConfig"`

	// REQUIRED; 鉴权异常时是否返回鉴权失败
	DenyOnFailed bool `json:"DenyOnFailed"`

	// REQUIRED; 返回状态码不在成功失败中时是否返回失败
	DenyOtherStatus bool `json:"DenyOtherStatus"`

	// REQUIRED; 鉴权失败时返回的状态码，范围在[400,699]内
	DenyReturnCode float32 `json:"DenyReturnCode"`

	// REQUIRED; 鉴权失败时鉴权服务器返回的状态码，所有状态码范围应在 [200,499] 之间，且和 AllowStatus 不重复。
	DenyRule ListVhostRemoteAuthResResultRemoteAuthConfigListItemDenyRule `json:"DenyRule"`

	// REQUIRED; 配置所属的domain
	Domain string `json:"Domain"`

	// REQUIRED; 远程鉴权配置是否打开
	Enable bool `json:"Enable"`

	// REQUIRED; 参数在50个以内
	HeaderParamConfig ListVhostRemoteAuthResResultRemoteAuthConfigListItemHeaderParamConfig `json:"HeaderParamConfig"`

	// REQUIRED; 鉴权请求方法，返回包括POST或GET
	Method string `json:"Method"`

	// REQUIRED; 参数在50个以内
	QueryParamConfig ListVhostRemoteAuthResResultRemoteAuthConfigListItemQueryParamConfig `json:"QueryParamConfig"`

	// REQUIRED; 配置所属的vhost
	Vhost string `json:"Vhost"`

	// 远程鉴权路径，UseUserRequest为true时忽略此字段
	AuthURLPath *string `json:"AuthURLPath,omitempty"`

	// 鉴权请求 Body 参数，最多配置 50 个 Body 参数。 :::tip 鉴权请求方法为 POST 时 Body 参数配置生效。 :::
	BodyParams []*ListVhostRemoteAuthResResultRemoteAuthConfigListPropertiesItemsItem `json:"BodyParams,omitempty"`

	// 是否开启对 HLS 协议拉流请求的鉴权开关，取值及含义如下所示。
	// * true：开启。
	// * false：（默认值）关闭。
	EnableHLSAuth *bool `json:"EnableHLSAuth,omitempty"`

	// 是否开启对 HLS 协议流的 TS 分片进行远程鉴权，取值及含义如下所示。
	// * true：开启。
	// * false：（默认值）关闭。
	// :::tip 开启 HLS 协议拉流请求的远程鉴权开关时，必须开启对 TS 分片的远程鉴权。 :::
	EnableHLSTSAuth *bool `json:"EnableHLSTSAuth,omitempty"`

	// 请求头是否保留大小写信息，true为原样请求，false转换为mime格式
	HeaderCaseSensitivity *bool `json:"HeaderCaseSensitivity,omitempty"`

	// 长度不超过1024
	HeaderHost *string `json:"HeaderHost,omitempty"`

	// 鉴权失败重试间隔，单位为秒，范围[1,30]
	RetryInterval *float32 `json:"RetryInterval,omitempty"`

	// 鉴权失败重试次数，范围[0,10]
	RetryTimes *float32 `json:"RetryTimes,omitempty"`

	// 鉴权超时时间，单位为秒。范围[0, 600]
	Timeout *float32 `json:"Timeout,omitempty"`

	// 是否使用用户请求的路径拼接到鉴权地址上
	UseUserRequest *bool `json:"UseUserRequest,omitempty"`
}

// ListVhostRemoteAuthResResultRemoteAuthConfigListItemCacheConfig - 鉴权缓存配置，配置均为空时表示不使用缓存来判断鉴权结果。
type ListVhostRemoteAuthResResultRemoteAuthConfigListItemCacheConfig struct {

	// REQUIRED; 生成鉴权结果缓存 Key 使用的参数配置，最多支持配置 50 个参数。 :::tip
	// * 鉴权成功与鉴权失败使用相同配置时，表示鉴权结果缓存 Key 配置；
	// * 鉴权成功与鉴权失败不使用相同配置时，表示鉴权成功缓存的过期时间。 :::
	CacheKeys ComponentsTmguxbSchemasListvhostremoteauthresPropertiesResultPropertiesRemoteauthconfiglistItemsPropertiesCacheconfigPropertiesCachekeys `json:"CacheKeys"`

	// REQUIRED; 鉴权成功和鉴权失败是否使用相同的配置（包括缓存的 key 值和缓存过期时间）。
	// * true：使用相同配置；
	// * false：不使用相同配置。
	UseSameCache bool `json:"UseSameCache"`

	// 缓存过期时间，范围为[0,3600]
	CacheExpireSecond *float32 `json:"CacheExpireSecond,omitempty"`

	// 鉴权失败缓存的过期时间，默认值为 0 时表示不缓存，单位为秒，取值范围为 [0,3600]。 :::tip 鉴权成功与鉴权失败不使用相同配置，即 UseSameCache 配置 false 时生效。 :::
	DenyExpireSecond *float32 `json:"DenyExpireSecond,omitempty"`

	// 生成鉴权失败结果缓存 Key 使用的参数配置，最多支持配置 50 个参数。
	// :::tip 鉴权成功与鉴权失败不使用相同配置，即 UseSameCache 配置 false 时生效。 :::
	DenyKeys *ComponentsNopwcvSchemasListvhostremoteauthresPropertiesResultPropertiesRemoteauthconfiglistItemsPropertiesCacheconfigPropertiesDenykeys `json:"DenyKeys,omitempty"`
}

// ListVhostRemoteAuthResResultRemoteAuthConfigListItemDenyRule - 鉴权失败时鉴权服务器返回的状态码，所有状态码范围应在 [200,499] 之间，且和 AllowStatus 不重复。
type ListVhostRemoteAuthResResultRemoteAuthConfigListItemDenyRule struct {

	// REQUIRED; 鉴权失败状态码，所有状态码范围在[200,499]内
	DenyStatus []float32 `json:"DenyStatus"`
}

// ListVhostRemoteAuthResResultRemoteAuthConfigListItemHeaderParamConfig - 参数在50个以内
type ListVhostRemoteAuthResResultRemoteAuthConfigListItemHeaderParamConfig struct {

	// REQUIRED
	UseUserParam bool                                                                                                                                                `json:"UseUserParam"`
	Params       []*Components6Aa5KwSchemasListvhostremoteauthresPropertiesResultPropertiesRemoteauthconfiglistItemsPropertiesHeaderparamconfigPropertiesParamsItems `json:"Params,omitempty"`
}

// ListVhostRemoteAuthResResultRemoteAuthConfigListItemQueryParamConfig - 参数在50个以内
type ListVhostRemoteAuthResResultRemoteAuthConfigListItemQueryParamConfig struct {

	// REQUIRED; 远程鉴权请求的 URL 或 Header 参数使用用户请求参数还是自定义参数，取值及含义如下所示。
	// * true：使用用户请求参数；
	// * false：自定义参数。
	UseUserParam bool `json:"UseUserParam"`

	// 自定义参数时的参数配置。
	Params []*ListVhostRemoteAuthResResultRemoteAuthConfigListPropertiesPropertiesItemsItem `json:"Params,omitempty"`
}

type ListVhostRemoteAuthResResultRemoteAuthConfigListPropertiesItemsItem struct {

	// REQUIRED; 参数类型，取值及含义如下所示。
	// * const_string：常量；
	// * header：用户请求的 Header 参数；
	// * query：用户请求的 URL 参数；
	// * vhost：参数值为变量 vhost 的参数，表示拉流请求中拉流域名所属的域名空间；
	// * domain：参数值为变量 domain 的参数，表示拉流请求中使用的拉流域名；
	// * app：参数值为变量 app 的参数，表示拉流请求中使用的 AppName；
	// * stream：参数值为变量 stream 的参数，表示拉流请求中使用的 StreamName；
	// * client_ip：参数值为变量 client_ip 的参数，表示拉流客户端 IP 地址；
	// * server_ip：参数值为变量 server_ip 的参数，表示响应拉流请求的 CDN 节点IP地址；
	// * request_uri：参数值为变量 request_uri 的参数，拉流请求地址的 URI。
	Type string `json:"Type"`

	// 参数名，最大长度为 100 个字符，不支持输入空格。 :::tip
	// * 参数类型为常量时表示常量参数的参数名；
	// * 参数类型为用户请求的 Header 参数或用户请求的 URL 参数时，表示指定用户请求中对应的参数名作为此处的参数名；
	// * 参数类型为变量时不生效。 :::
	ParamName *string `json:"ParamName,omitempty"`

	// 参数名的映射参数名，最大长度为 100 个字符，不支持输入空格。 :::tip
	// * 参数类型为常量时不生效；
	// * 参数类型为用户请求的 Header 参数或用户请求的 URL 参数时，表示鉴权请求时使用 ToName 值代替用户请求中对应的参数名；
	// * 参数类型为变量时，表示使用 ToName 取值作为此变量的参数名。 :::
	ToName *string `json:"ToName,omitempty"`

	// 参数类型为常量时的参数值，最大长度为 100 个字符，不支持输入空格。
	Value *string `json:"Value,omitempty"`
}

type ListVhostRemoteAuthResResultRemoteAuthConfigListPropertiesPropertiesItemsItem struct {

	// REQUIRED
	Type      string  `json:"Type"`
	ParamName *string `json:"ParamName,omitempty"`
	ToName    *string `json:"ToName,omitempty"`
	Value     *string `json:"Value,omitempty"`
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

type ListVhostSubtitleTranscodePresetBody struct {

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`
}

type ListVhostSubtitleTranscodePresetRes struct {

	// REQUIRED
	ResponseMetadata ListVhostSubtitleTranscodePresetResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result ListVhostSubtitleTranscodePresetResResult `json:"Result"`
}

type ListVhostSubtitleTranscodePresetResResponseMetadata struct {

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

type ListVhostSubtitleTranscodePresetResResult struct {

	// REQUIRED; 字幕配置列表。
	PresetList []ListVhostSubtitleTranscodePresetResResultPresetListItem `json:"PresetList"`
}

type ListVhostSubtitleTranscodePresetResResultPresetListItem struct {

	// REQUIRED; 应用名称。
	App string `json:"App"`

	// REQUIRED
	Stream string `json:"Stream"`

	// REQUIRED; 转码后缀标识。
	Suffixes []string `json:"Suffixes"`

	// REQUIRED; 字幕配置详细参数。
	TranscodePreset ListVhostSubtitleTranscodePresetResResultPresetListItemTranscodePreset `json:"TranscodePreset"`

	// REQUIRED; 域名空间。
	Vhost string `json:"Vhost"`
}

// ListVhostSubtitleTranscodePresetResResultPresetListItemTranscodePreset - 字幕配置详细参数。
type ListVhostSubtitleTranscodePresetResResultPresetListItemTranscodePreset struct {

	// REQUIRED
	CreatedAt int32 `json:"CreatedAt"`

	// REQUIRED
	Delay int32 `json:"Delay"`

	// REQUIRED; 字幕配置的描述信息。
	Description string `json:"Description"`

	// REQUIRED; 预设配置，使用预设配置是系统将自动对字体大小、字幕行数、每行最大字符数和边距参数（MarginVertical 和 MarginHorizontal）进行智能化适配。默认为空，表示不使用预设配置，支持的预设配置如下所示。
	// * small ：小字幕。
	// * medium：中字幕。
	// * large：大字幕。 :::tip 使用预设配置时，字幕行数、每行最大字符数、左右边距和底部边距参数不生效，系统将使用预设配置自动进行计算。 :::
	DisplayPreset string `json:"DisplayPreset"`

	// REQUIRED; 原文翻译成译文时使用的热词词库。
	GlossaryWordList []string `json:"GlossaryWordList"`

	// REQUIRED; 原文字幕识别时使用的热词词库。
	HotWordList []string `json:"HotWordList"`

	// REQUIRED; 设置在 16:9 分辨率场景下，每行字幕展示的最大字符数。 :::tip
	// * 使用预设配置时，字幕每行最大字符数设置不生效。
	// * 不使用预设配置时，字幕每行最大字符数必填。
	// * 每个文字、字母、符号或数字均为一个字符。
	// * 当屏幕分辨率改变时，屏幕上显示的每行文字数量会相应调整，以适应新的分辨率，确保文字的显示效果和阅读体验。 :::
	MaxCharNumber int32 `json:"MaxCharNumber"`

	// REQUIRED; 字幕展示的行数，同时适用于原文字幕和译文字幕，支持的取值及含义如下所示。
	// * 0：（默认值）根据字幕字数自动进行分行展示；
	// * 1：每种字幕展示一行；
	// * 2：每种字幕展示两行。 :::tip
	// * 使用预设配置时，字幕行数为自动分行展示。
	// * 超出行内字数限制时表示字幕将超过显示范围，此时字幕内容将被截断。 :::
	MaxRowNumber int32 `json:"MaxRowNumber"`

	// REQUIRED; 字幕位置设置，通过设置字幕距离画面左右边距和底部边距来指定字幕位置。
	// :::tip
	// * 使用预设配置时，字幕位置设置不生效。
	// * 不使用预设配置时，字幕位置设置必填。 :::
	Position ComponentsJ1MbxoSchemasListvhostsubtitletranscodepresetresPropertiesResultPropertiesPresetlistItemsPropertiesTranscodepresetPropertiesPosition `json:"Position"`

	// REQUIRED; 字幕配置的名称。
	PresetName string `json:"PresetName"`

	// REQUIRED; 原文字幕展示参数配置。
	SourceLanguage Components1523StvSchemasListvhostsubtitletranscodepresetresPropertiesResultPropertiesPresetlistItemsPropertiesTranscodepresetPropertiesSourcelanguage `json:"SourceLanguage"`

	// REQUIRED
	Status int32 `json:"Status"`

	// REQUIRED
	SuffixName string `json:"SuffixName"`

	// REQUIRED; 译文字幕展示参数配置列表。
	TargetLanguage []Components1C398ShSchemasListvhostsubtitletranscodepresetresPropertiesResultPropertiesPresetlistItemsPropertiesTranscodepresetPropertiesTargetlanguageItems `json:"TargetLanguage"`

	// REQUIRED
	UpdatedAt int32 `json:"UpdatedAt"`
}

// ListVhostSubtitleTranscodePresetResResultPresetListItemTranscodePresetTargetLanguageItemBorder - 译文字幕的字体描边配置。
type ListVhostSubtitleTranscodePresetResResultPresetListItemTranscodePresetTargetLanguageItemBorder struct {

	// REQUIRED
	Color string `json:"Color"`

	// REQUIRED
	Width int32 `json:"Width"`
}

type ListVhostTransCodePresetBody struct {

	// REQUIRED; 域名空间，即直播流地址的域名所属的域名空间。您可以调用 ListDomainDetail [https://www.volcengine.com/docs/6469/1126815] 接口或在视频直播控制台的域名管理
	// [https://console.volcengine.com/live/main/domain/list]页面，查看需要录制的直播流使用的域名所属的域名空间。
	Vhost string `json:"Vhost"`

	// 是否是dash abr 请求
	IsDashAbr *bool `json:"IsDashAbr,omitempty"`

	// 是否是hls abr 请求
	IsHlsAbr *bool `json:"IsHlsAbr,omitempty"`
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

	// REQUIRED; 直播流地址的 AppName 字段。
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

	// REQUIRED; 直播流地址的 AppName 字段。
	App string `json:"App"`

	// REQUIRED; 域名空间。
	Vhost string `json:"Vhost"`

	// 转码配置具体信息。
	TranscodePreset *ListVhostTransCodePresetResResultCommonPresetListItemTranscodePreset `json:"TranscodePreset,omitempty"`
}

// ListVhostTransCodePresetResResultCommonPresetListItemTranscodePreset - 转码配置具体信息。
type ListVhostTransCodePresetResResultCommonPresetListItemTranscodePreset struct {

	// 音频编码格式，默认值为aac，支持的取值及含义如下所示。
	// * aac：使用 AAC 音频编码格式；
	// * opus：使用 Opus 音频编码格式。
	// * copy：不进行音频转码，所有音频编码参数不生效，音频编码参数包括音频码率（AudioBitrate）等。
	Acodec *string `json:"Acodec,omitempty"`

	// 视频分辨率自适应模式开关，默认值为0。支持的取值及含义如下。
	// * 0：关闭视频分辨率自适应；
	// * 1：开启视频分辨率自适应。 :::tip
	// * 关闭视频分辨率自适应模式（As取值为0）时，转码配置的视频分辨率取视频宽度（Width）和视频高度（Height）的值对转码视频进行拉伸；
	// * 开启视频分辨率自适应模式（As取值为1）时，转码配置的视频分辨率按照短边长度（ShortSide）、长边长度（LongSide）、视频宽度（Width）、视频高度（Height）的优先级取值，另一边等比缩放。 :::
	As *string `json:"As,omitempty"`

	// 音频码率，单位为 kbps。
	AudioBitrate *int32 `json:"AudioBitrate,omitempty"`

	// 是否开启转码视频分辨率不超过源流分辨率，默认值为1表示开启。开启后，当源流分辨率低于转码配置分辨率时（即源流宽低于转码配置宽且源流高低于转码配置高时），将按源流视频分辨率进行转码。
	// * 0：关闭；
	// * 1：开启。
	AutoTransResolution *int32 `json:"AutoTransResolution,omitempty"`

	// 是否开启转码视频码率不超过源流码率，默认值为1表示开启。开启后，当源流码率低于转码配置码率时，将按照源流视频码率进行转码。
	// * 0：关闭；
	// * 1：开启。
	AutoTransVb *int32 `json:"AutoTransVb,omitempty"`

	// 是否开启转码视频帧率不超过源流帧率，默认值为1表示开启。开启后，当源流帧率低于转码配置帧率时，将按照源流视频帧率进行转码。
	// * 0：关闭；
	// * 1：开启。
	AutoTransVr *int32 `json:"AutoTransVr,omitempty"`

	// 转码输出视频中 2 个参考帧之间的最大 B 帧数量，取值为0时表示去除 B 帧。
	BFrames *int32 `json:"BFrames,omitempty"`

	// 视频帧率，单位为 fps，帧率越大，画面越流畅。
	FPS *int32 `json:"FPS,omitempty"`

	// IDR 帧之间的最大间隔，单位为秒。
	GOP *int32 `json:"GOP,omitempty"`

	// 视频高度。 :::tip
	// * 当关闭视频分辨率自适应（As取值为0）时，转码分辨率将取Width和Height的值对转码视频进行拉伸；
	// * 当关闭视频分辨率自适应（As取值为0）时，Width和Height任一取值为0时，转码视频将保持源流尺寸。 :::
	Height *int32 `json:"Height,omitempty"`

	// 长边长度。 :::tip
	// * 当开启视频分辨率自适应模式时（As取值为1）时，参数生效，反之则不生效。
	// * 当开启视频分辨率自适应模式时（As取值为1）时，如果LongSide、ShortSide、Width、Height同时取0，表示保持源流尺寸。 :::
	LongSide *int32 `json:"LongSide,omitempty"`

	// 转码配置名称。
	Preset *string `json:"Preset,omitempty"`

	// 转码类型是否为极智超清转码，默认值为false，取值及含义如下。
	// * true：极智超清转码；
	// * false：标准转码。
	// :::tip 视频编码格式为 H.266 （Vcodec 取值为 h266）时，转码类型不支持极智超清转码。 :::
	Roi *bool `json:"Roi,omitempty"`

	// 短边长度。 :::tip
	// * 当 As 的取值为 1 即开启宽高自适应时，参数生效，反之则不生效。
	// * 当 As 的取值为 1 时，如果 LongSide 、 ShortSide 、Width 、Height 同时取 0，表示保持源流尺寸。 :::
	ShortSide *int32 `json:"ShortSide,omitempty"`

	// 转码停止时长，支持触发方式为拉流转码（TransType取值为Pull）时设置，表示断开拉流后转码停止的时长，单位为秒，取值范围为-1和 [0,300]，-1表示不停止转码，默认值为60。
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
	// * copy：不进行视频转码，所有视频编码参数不生效，视频编码参数包括视频帧率（FPS）、视频码率（VideoBitrate）、分辨率设置（As、Width、Height、ShortSide、LongSide）、GOP和BFrames等。
	Vcodec *string `json:"Vcodec,omitempty"`

	// 视频码率，单位为 kbps。
	VideoBitrate *int32 `json:"VideoBitrate,omitempty"`

	// 视频宽度。 :::tip
	// * 当关闭视频分辨率自适应（As取值为0）时，转码分辨率将取Width和Height的值对转码视频进行拉伸；
	// * 当关闭视频分辨率自适应（As取值为0）时，Width和Height任一取值为0时，转码视频将保持源流尺寸。 :::
	Width *int32 `json:"Width,omitempty"`
}

type ListVhostTransCodePresetResResultCustomizePresetListItem struct {

	// REQUIRED; 直播流地址的 AppName 字段。
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

	// 统计消息，提供可用配置和不可用配置的数量。
	StaticsMsg *string `json:"StaticsMsg,omitempty"`

	// 不可正常使用的水印配置列表，如水印图片获取失败等原因导致的配置不可用。返回不可正常使用的水印配置信息及配置不可用的原因。
	WatermarkErrMsgList []*ListVhostWatermarkPresetResResultWatermarkErrMsgListItem `json:"WatermarkErrMsgList,omitempty"`

	// 可正常使用的水印配置列表。
	WatermarkPresetList []*ListVhostWatermarkPresetResResultWatermarkPresetListItem `json:"WatermarkPresetList,omitempty"`
}

type ListVhostWatermarkPresetResResultWatermarkErrMsgListItem struct {

	// 火山引擎账号 ID。
	AccountID *string `json:"AccountID,omitempty"`

	// 应用名称。
	App *string `json:"App,omitempty"`

	// 获取水印模板失败的具体错误信息。
	ErrMsg *string `json:"ErrMsg,omitempty"`

	// 域名空间。
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

	// 流名称，取值与直播流地址中 StreamName 字段取值相同。支持由大小写字母（A - Z、a - z）、数字（0 - 9）、下划线（_）、短横线（-）和句点（.）组成，长度为 1 到 100 个字符。
	// :::tip
	// * 默认为空，表示查询的 AppName 级别对所有转码流生效的配置。
	// * 指定流名称时，表示查询仅对 AppName 下指定流名称的转码流生效的配置。 :::
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

type RelaunchPullToPushTaskBody struct {

	// REQUIRED; 任务 ID，任务的唯一标识，您可以通过获取拉流转推任务列表 [https://www.volcengine.com/docs/6469/1126896]接口获取状态为停用的任务 ID。
	TaskID string `json:"TaskId"`

	// 任务所属的群组名称，您可以通过获取拉流转推任务列表 [https://www.volcengine.com/docs/6469/1126896]接口获取。 :::tip
	// * 使用主账号调用时，为非必填。
	// * 使用子账号调用时，为必填。 :::
	GroupName *string `json:"GroupName,omitempty"`
}

type RelaunchPullToPushTaskRes struct {

	// REQUIRED
	ResponseMetadata RelaunchPullToPushTaskResResponseMetadata `json:"ResponseMetadata"`
}

type RelaunchPullToPushTaskResResponseMetadata struct {

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
	Error   *RelaunchPullToPushTaskResResponseMetadataError `json:"Error,omitempty"`
}

type RelaunchPullToPushTaskResResponseMetadataError struct {

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

	// 任务所属的群组名称，您可以通过获取拉流转推任务列表 [https://www.volcengine.com/docs/6469/1126896]接口获取。 :::tip
	// * 使用主账号调用时，为非必填。
	// * 使用子账号调用时，为必填。 :::
	GroupName *string `json:"GroupName,omitempty"`
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

	// REQUIRED; 鉴权配置参数，包括鉴权密钥、鉴权参数、加密字符串生成算法等。
	AuthDetailList []UpdateAuthKeyBodyAuthDetailListItem `json:"AuthDetailList"`

	// REQUIRED; 直播流使用的域名。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看直播流使用的域名。
	Domain string `json:"Domain"`

	// REQUIRED; 鉴权场景类型，取值及含义如下所示。
	// * push：推流鉴权；
	// * pull：拉流鉴权。
	SceneType string `json:"SceneType"`

	// 是否开启 URL 地址鉴权，取值及含义如下所示。
	// * false：关闭（默认值）；
	// * true：开启。
	PushPullEnable *bool `json:"PushPullEnable,omitempty"`

	// 时间戳进制。默认值为 10，但当 AuthType 取值为 TypeC 时，默认值为 16。取值如下：
	// * 2：二进制
	// * 8：八进制
	// * 10：十进制
	// * 16：十六进制
	// :::tipSceneType 取值为 push 时，该参数取值固定为 10。 :::
	TimeStampBase *int32 `json:"TimeStampBase,omitempty"`

	// 鉴权生效时长，单位为秒，默认值为 0，取值范围为 [0,2592000]，超出生效时长后 URL 无法使用。
	ValidDuration *int32 `json:"ValidDuration,omitempty"`
}

type UpdateAuthKeyBodyAuthDetailListItem struct {

	// REQUIRED; 推/拉流鉴权时必选
	AuthType string `json:"AuthType"`

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

type UpdateCarouselTaskBody struct {

	// REQUIRED; 轮播规则，用于指定轮播播放的素材和行为等。
	Rule UpdateCarouselTaskBodyRule `json:"Rule"`

	// REQUIRED; 待更新的轮播任务 ID，任务的唯一标识。调用 CreateCarouselTask 接口创建轮播任务时返回。
	TaskID string `json:"TaskID"`
}

// UpdateCarouselTaskBodyRule - 轮播规则，用于指定轮播播放的素材和行为等。
type UpdateCarouselTaskBodyRule struct {

	// REQUIRED; 轮播素材列表，用于指定在轮播过程中播放的素材资源。
	Source []UpdateCarouselTaskBodyRuleSourceItem `json:"Source"`

	// 循环次数。取值范围为 [ ]，单位为，默认值为``。
	Loop *int32 `json:"Loop,omitempty"`

	// 对素材更新后的播放行为进行控制
	SeekInfo *UpdateCarouselTaskBodyRuleSeekInfo `json:"SeekInfo,omitempty"`
}

// UpdateCarouselTaskBodyRuleSeekInfo - 对素材更新后的播放行为进行控制
type UpdateCarouselTaskBodyRuleSeekInfo struct {

	// 0 表示推完当前播放的素材后再进行素材切换；1 表示立刻切换到指定的素材、指定的进度
	Immediate *int64 `json:"Immediate,omitempty"`

	// 更新后播放的素材ID，为空代表不指定。
	SourceID *string `json:"SourceID,omitempty"`

	// 切换素材后，素材播放的位置。
	SourceSeek *int64 `json:"SourceSeek,omitempty"`
}

type UpdateCarouselTaskBodyRuleSourceItem struct {

	// REQUIRED; 注意，如果ID相同，此结构的其余字段也需要保证相同
	ID string `json:"ID"`

	// REQUIRED; 轮播素材的文件类型，用于指定素材的文件来源类型。支持以下取值：
	// * vod：点播 MP4 或 FLV 文件；
	// * m3u8：点播 M3U8 文件。
	// :::tip 如果素材的 ID 没有变化（即更新的 ID 与原素材的 ID 相同），Type 取值要和元素材保持一致。 :::
	Type string `json:"Type"`

	// REQUIRED; 轮播素材的公网可访问地址。确保提供的地址能够被公网正常访问，以便正确加载轮播素材内容。 :::tip 如果素材的 ID 没有变化（即更新的 ID 与原素材的 ID 相同），Url 取值要和元素材保持一致。 :::
	URL string `json:"Url"`

	// 指定此素材连续播放的次数。该字段值必须大于等于 0，不传时，将保持原有轮播配置。支持的取值及含义如下：
	// * 0：不循环播放；
	// * 其他正整数：按照指定次数循环播放。
	Loop *int32 `json:"Loop,omitempty"`

	// 用于控制当前素材播放时跳过开头的一段时间，例如，跳过片头，单位为秒。该字段仅在素材类型为视频点播（type=vod）时有效。以下是该字段的使用规则：
	// * 如果 Seek 的取值小于等于 0 或大于视频的实际时长，则该字段不生效。
	// * 确保根据点播素材的实际长度设置合适的值，以实现跳过片头的效果。
	Seek *int32 `json:"Seek,omitempty"`
}

type UpdateCarouselTaskRes struct {

	// REQUIRED
	ResponseMetadata UpdateCarouselTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result UpdateCarouselTaskResResult `json:"Result"`
}

type UpdateCarouselTaskResResponseMetadata struct {

	// REQUIRED
	RequestID string `json:"RequestID"`
}

type UpdateCarouselTaskResResult struct {

	// REQUIRED; 包含任务更新相关信息的数据对象。
	Data UpdateCarouselTaskResResultData `json:"Data"`
}

// UpdateCarouselTaskResResultData - 包含任务更新相关信息的数据对象。
type UpdateCarouselTaskResResultData struct {

	// REQUIRED; 当前生效的序列号
	OptID int32 `json:"OptID"`
}

type UpdateCloudMixTaskBody struct {

	// REQUIRED; 混流任务详细配置。
	MixedRules UpdateCloudMixTaskBodyMixedRules `json:"MixedRules"`

	// REQUIRED; 混流任务 ID，您可以通过 ListCloudMixTask [https://www.volcengine.com/docs/6469/1271157] 接口获取运行中的混流任务 ID。
	TaskID string `json:"TaskID"`
}

// UpdateCloudMixTaskBodyMixedRules - 混流任务详细配置。
type UpdateCloudMixTaskBodyMixedRules struct {

	// REQUIRED; 混流输出布局配置。
	InputLayout UpdateCloudMixTaskBodyMixedRulesInputLayout `json:"InputLayout"`

	// REQUIRED; 混流素材列表，最多支持配置 8 路输入源。
	InputSource []UpdateCloudMixTaskBodyMixedRulesInputSourceItem `json:"InputSource"`

	// REQUIRED; 混流输出流参数配置。 :::warning 更新云端混流任务时，Output 参数不支持变更，且必须传入与原混流任务一致的配置。 :::
	Output UpdateCloudMixTaskBodyMixedRulesOutput `json:"Output"`
}

// UpdateCloudMixTaskBodyMixedRulesInputLayout - 混流输出布局配置。
type UpdateCloudMixTaskBodyMixedRulesInputLayout struct {

	// REQUIRED; 混流输出画布配置和素材布局配置。
	Scene UpdateCloudMixTaskBodyMixedRulesInputLayoutScene `json:"Scene"`

	// 混流输出视频中 Logo 布局配置。 :::tip 支持最多配置 4 个 Logo，展示层级以添加顺序为准。 :::
	Logo []*UpdateCloudMixTaskBodyMixedRulesInputLayoutLogoItem `json:"Logo,omitempty"`
}

type UpdateCloudMixTaskBodyMixedRulesInputLayoutLogoItem struct {

	// Logo 图片在混流输出整体画面中的布局配置。
	Layout *UpdateCloudMixTaskBodyMixedRulesInputLayoutLogoItemLayout `json:"Layout,omitempty"`

	// Logo 图片访问地址。
	URL *string `json:"Url,omitempty"`
}

// UpdateCloudMixTaskBodyMixedRulesInputLayoutLogoItemLayout - Logo 图片在混流输出整体画面中的布局配置。
type UpdateCloudMixTaskBodyMixedRulesInputLayoutLogoItemLayout struct {

	// REQUIRED
	H int32 `json:"H"`

	// REQUIRED
	W int32 `json:"W"`

	// REQUIRED
	X int32 `json:"X"`

	// REQUIRED
	Y int32 `json:"Y"`
}

// UpdateCloudMixTaskBodyMixedRulesInputLayoutScene - 混流输出画布配置和素材布局配置。
type UpdateCloudMixTaskBodyMixedRulesInputLayoutScene struct {

	// REQUIRED; 混流输出整体画布高度，单位为 px，取值范围为 [10,2160]。
	Height int32 `json:"Height"`

	// REQUIRED; 混流素材在混流输出整体画面中的布局配置。 :::tip 混流素材布局中需包含所有素材的配置，且需与通过 Layer 参数与混流素材一一匹配。 :::
	Layout []UpdateCloudMixTaskBodyMixedRulesInputLayoutSceneLayoutItem `json:"Layout"`

	// REQUIRED; 混流输出画布整体宽度，单位为 px，取值范围为 [10,2160]。
	Width int32 `json:"Width"`
}

type UpdateCloudMixTaskBodyMixedRulesInputLayoutSceneLayoutItem struct {

	// REQUIRED; 当前素材或 Logo 图片在混流输出画面中的限制高度，单位为 px，取值范围为 [10,2160]。
	// :::tip 限制宽度和限制高度指定了素材展示的限制范围，当素材尺寸和限制值不一致时，素材将在限制范围内根据长边进行等比缩放，其余区域透明展示。 :::
	H int32 `json:"H"`

	// REQUIRED; 当配置素材布局时需要通过 Layer 参数与素材进行一一对应。 :::tip 配置 Logo 图片的布局时此参数不生效。 :::
	Layer int32 `json:"Layer"`

	// REQUIRED; 当前素材或 Logo 图片在混流输出画面中的限制宽度，单位为 px，取值范围为 [10,2160]。
	W int32 `json:"W"`

	// REQUIRED; 当前素材或 Logo 图片在输出画面中相对画面左上角的 X 偏移位置，单位为 px，取值范围为 0 到设置的画面宽度。
	X int32 `json:"X"`

	// REQUIRED; 当前素材或 Logo 图片在输出画面中相对画面左上角的 Y 偏移位置，单位为 px，取值范围为 0 到设置的画面高度。
	Y int32 `json:"Y"`
}

type UpdateCloudMixTaskBodyMixedRulesInputSourceItem struct {

	// REQUIRED; 混流素材 ID，一个任务中素材 ID 不能重复，此 ID 用于任务状态回调消息中标识素材。
	ID string `json:"ID"`

	// REQUIRED; 混流素材的叠放顺序，1 为最底层，2 层在 1 层之上，以此类推，取值范围为[1,9999]。 :::tip 当前准备使用某个素材作为布局背景时，其叠放顺序应设置为所有素材中的最小值。 :::
	Layer int32 `json:"Layer"`

	// REQUIRED; 混流素材类型，支持的取值及含义如下所示。
	// * vod：视频点播中的素材，支持 MP4、FLV 格式素材；
	// * live：直播源素材，支持 RTMP、FLV 协议拉流地址；
	// * pic：图片素材，支持 png、jpg 格式图片。
	Type string `json:"Type"`

	// REQUIRED; 混流素材访问地址。
	URL string `json:"Url"`
}

// UpdateCloudMixTaskBodyMixedRulesOutput - 混流输出流参数配置。 :::warning 更新云端混流任务时，Output 参数不支持变更，且必须传入与原混流任务一致的配置。 :::
type UpdateCloudMixTaskBodyMixedRulesOutput struct {

	// REQUIRED; 混流声音参数设置。
	Audio UpdateCloudMixTaskBodyMixedRulesOutputAudio `json:"Audio"`

	// REQUIRED; 混流视频的推流地址。
	URL []string `json:"Url"`

	// REQUIRED; 混流画面参数设置。
	Video UpdateCloudMixTaskBodyMixedRulesOutputVideo `json:"Video"`
}

// UpdateCloudMixTaskBodyMixedRulesOutputAudio - 混流声音参数设置。
type UpdateCloudMixTaskBodyMixedRulesOutputAudio struct {

	// REQUIRED; 混流输出流的音频码率，单位为 bps，取值范围为 [128000,320000]，常见取值及含义如下所示。
	// * 128000：128 kbps；
	// * 144000：144 kbps；
	// * 256000：256 kbps；
	// * 320000：320 kbps。
	BitRate int32 `json:"BitRate"`

	// REQUIRED; 混流输出流的音频声道设置，取值及含义如下所示。
	// * mono：单声道；
	// * stereo：立体声。
	ChannelLayout string `json:"ChannelLayout"`

	// REQUIRED; 混流输出流的音频采样率，单位为 HZ，常见取值及含义如下所示。
	// * 32000：32 kHZ；
	// * 44100：44.1 kHZ；
	// * 48000：48 kHZ。
	SampleRate int32 `json:"SampleRate"`
}

// UpdateCloudMixTaskBodyMixedRulesOutputVideo - 混流画面参数设置。
type UpdateCloudMixTaskBodyMixedRulesOutputVideo struct {

	// REQUIRED; 混流输出视频码率，单位为 bps，取值范围为 [1000000,20000000]，输入值小于或大于取值范围时会进行自动修正至最小值和最大值。
	BitRate int32 `json:"BitRate"`

	// REQUIRED; 混流输出视频编码格式，支持的取值及含义如下所示。
	// * h264：使用 H.264 编码格式；
	// * h265：使用 H.265 编码格式。
	Codec string `json:"Codec"`

	// REQUIRED; 混流输出视频帧率，单位为 fps，取值范围为 [10,60]，输入值小于或大于取值范围时会进行自动修正至最小值和最大值。
	FrameRate int32 `json:"FrameRate"`

	// REQUIRED; IDR 帧之间的最大间隔时间，单位为秒，取值范围为 [1,10]。
	GOP int32 `json:"GOP"`
}

type UpdateCloudMixTaskRes struct {

	// REQUIRED
	ResponseMetadata UpdateCloudMixTaskResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result UpdateCloudMixTaskResResult `json:"Result"`
}

type UpdateCloudMixTaskResResponseMetadata struct {

	// REQUIRED
	RequestID string `json:"RequestID"`
}

type UpdateCloudMixTaskResResult struct {

	// REQUIRED; 请求响应码，取值及含义如下。
	// * 0：请求成功；
	// * 500：内部处理错误；
	// * 400：请求异常。
	Code int32 `json:"Code"`

	// REQUIRED; 返回的数据。
	Data UpdateCloudMixTaskResResultData `json:"Data"`

	// REQUIRED; 请求响应码对应的信息。
	Message string `json:"Message"`
}

// UpdateCloudMixTaskResResultData - 返回的数据。
type UpdateCloudMixTaskResResultData struct {

	// REQUIRED; 任务版本标识符，用来标识更新后的任务版本。
	OptID int32 `json:"OptID"`
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

type UpdateEncryptHLSBody struct {

	// REQUIRED; 视频直播服务端生成密钥的更新周期，单位为秒，取值范围为 [60,604800]。
	CycleTime string `json:"CycleTime"`

	// REQUIRED; 客户自建密钥管理服务后，客户端向密钥管理服务请求获取密钥的地址。
	URL string `json:"URL"`
}

type UpdateEncryptHLSRes struct {

	// REQUIRED
	ResponseMetadata UpdateEncryptHLSResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateEncryptHLSResResponseMetadata struct {

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
	MaxSize     *int32   `json:"MaxSize,omitempty"`
	PDTInterval *float32 `json:"PDTInterval,omitempty"`

	// 可选枚举值 "audio_only" "video_only "video_keyframe_only" "video_single_keyframe_only"
	PacketFilter       *string  `json:"PacketFilter,omitempty"`
	PartTargetDuration *float32 `json:"PartTargetDuration,omitempty"`

	// ts存储位置
	Path *string `json:"Path,omitempty"`

	// m3u8的ts个数，默认3个
	PlaylistLength *int32 `json:"PlaylistLength,omitempty"`

	// 预取ts个数
	PrefetchNum *int32 `json:"PrefetchNum,omitempty"`

	// 服务类型
	ServiceType *string `json:"ServiceType,omitempty"`

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

	// REQUIRED; 0: response 1: request
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

	// REQUIRED; 与Vhost二选一
	Domain string `json:"Domain"`

	// REQUIRED; 单位毫秒，大于等于0
	GopCacheSize int32   `json:"GopCacheSize"`
	Vhost        *string `json:"Vhost,omitempty"`
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

	// 接收拉流转推任务状态回调的地址，最大长度为 512 个字符。
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

	// 任务所属的群组名称，您可以通过获取拉流转推任务列表 [https://www.volcengine.com/docs/6469/1126896]接口获取。 :::tip
	// * 群组名称不支持更新，仅做校验使用。
	// * 使用主账号调用时，为非必填。
	// * 使用子账号调用时，为必填。 :::
	GroupName *string `json:"GroupName,omitempty"`

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

	// 点播文件地址和开始播放、结束播放的时间设置。 :::tip
	// * 当 Type 为点播类型时配置生效。
	// * 与 SrcAddrS 和 OffsetS 字段不可同时填写。 :::
	VodSrcAddrs []*UpdatePullToPushTaskBodyVodSrcAddrsItem `json:"VodSrcAddrs,omitempty"`

	// 为拉流转推视频添加的水印配置信息。
	Watermark *UpdatePullToPushTaskBodyWatermark `json:"Watermark,omitempty"`
}

type UpdatePullToPushTaskBodyVodSrcAddrsItem struct {

	// REQUIRED; 点播文件地址。
	SrcAddr string `json:"SrcAddr"`

	// 当前点播文件结束播放的时间偏移值，单位为秒，默认为空时表示结束播放时间不进行偏移。
	EndOffset *float32 `json:"EndOffset,omitempty"`

	// 当前点播文件开始播放的时间偏移值，单位为秒。默认为空时表示开始播放时间不进行偏移。
	StartOffset *float32 `json:"StartOffset,omitempty"`
}

// UpdatePullToPushTaskBodyWatermark - 为拉流转推视频添加的水印配置信息。
type UpdatePullToPushTaskBodyWatermark struct {

	// REQUIRED; 水印图片字符串，图片最大 2MB，最小 100Bytes，最大分辨率为 1080×1080。图片 Data URL 格式为：data:image/<mediatype>;base64,<data>。
	// * mediatype：图片类型，支持 png、jpg、jpeg 格式；
	// * data：base64 编码的图片字符串。
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

	// 转码流后缀列表，是否录制转码设置为根据转码流列表匹配（TranscodeRecord 取值为 2）时生效，TranscodeSuffixList 默认配置为空，效果等同于录制全部转码流。
	TranscodeSuffixList []*string `json:"TranscodeSuffixList,omitempty"`
}

// UpdateRecordPresetV2BodyRecordPresetConfigFlvParam - 录制为 FLV 格式时的录制参数。 :::tip 您需至少配置一个录制格式，即 FlvParam、HlsParam、Mp4Param
// 至少开启一个。 :::
type UpdateRecordPresetV2BodyRecordPresetConfigFlvParam struct {

	// 实时录制场景下，断流等待时长，单位为秒，默认值为 180，取值范围为 [0,3600]。如果实际断流时间小于断流等待时长，录制任务不会停止；如果实际断流时间大于断流等待时长，录制任务会停止，断流恢复后重新开始一个新的录制任务。
	ContinueDuration *int32 `json:"ContinueDuration,omitempty"`

	// 断流录制场景下，单文件录制时长，单位为秒，默认值为 7200，取值范围为 -1 和 [300,86400]。
	// * 取值为 -1 时，表示不限制录制时长，录制结束后生成一个完整的录制文件。
	// * 取值为 [300,86400] 之间的值时，表示根据设置的录制文件时长，到达时长立即生成录制文件，完成录制后一起上传。
	// :::tip
	// * 断流录制场景仅在录制格式为 HLS 时生效，且断流录制和实时录制为二选一配置。
	// * 如录制过程中出现断流，对应生成的录制文件时长也会相应缩短。 :::
	Duration *int32 `json:"Duration,omitempty"`

	// 当前格式的录制是否开启，默认 false，取值及含义如下所示。
	// * false：不开启；
	// * true：开启。
	Enable *bool `json:"Enable,omitempty"`

	// 实时录制场景下，单文件录制时长，单位为秒，默认值为 1800，取值范围为 [300,21600]。录制时间到达设置的单文件录制时长时，会立即生成录制文件实时上传存储。 :::tip 如录制过程中出现断流，对应生成的录制文件时长也会相应缩短。
	// :::
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

	// 录制文件存储到 TOS 时的存储路径和文件名规则。支持输入字母（A - Z、a - z）、数字（0 - 9）、短横线（-）、叹号（!）、下划线（_）、句点（.）、星号（*）及占位符。最大长度为 200 个字符，
	// 支持以下字段作为占位符：
	// * record：自定义字段，可遵照支持字符进行自定义。
	// * {PubDomain}：当前配置中的 vhost 值。
	// * {App}：当前配置中的 AppName 值。
	// * {Stream}：当前配置中的 StreamName 值。
	// * {StartTime}：录制开始的 Unix 时间戳，精度为 s。
	// * {EndTime}：录制结束的 Unix 时间戳，精度为 s。
	// 存储路径必须至少包含两级目录。例如：live/{App}/{Stream}
	// 合法示例：
	// record/{PubDomain}/{App}/{Stream}/{StartTime}-{EndTime}
	// {App}/archive/{Stream}/recording_{StartTime}
	// vod/{Stream}/!highlight_{EndTime}
	// a/b/custom_record
	// 错误示例：
	// single_level # 错误：路径层级不足两级
	// invalid_/{S@ream}/file # 错误：含非法字符@
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
	ExactObject     *string `json:"ExactObject,omitempty"`
	ExactObjectName *string `json:"ExactObjectName,omitempty"`

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
	Type string `json:"Type"`

	// 限制国家列表，传入使用国家代码，遵循iso-3166
	CountryList []*string `json:"CountryList,omitempty"`

	// 限制省份列表，目前仅支持中国香港地区、中国澳门地区和中国台湾地区。分别对应代码为HK, MO, TW
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

	// REQUIRED; volcengine可以传入rtmp/flv, byteplus可以传入rtmp\flv\dash\hls
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

type UpdateRemoteAuthBody struct {

	// REQUIRED; 鉴权成功状态码，所有状态码范围应在[200,499]间，和DenyCode不能有交集
	AllowStatus []int32 `json:"AllowStatus"`

	// REQUIRED; 远程鉴权地址
	AuthURL string `json:"AuthURL"`

	// REQUIRED; 鉴权异常时是否返回鉴权失败
	DenyOnFailed bool `json:"DenyOnFailed"`

	// REQUIRED; 返回状态码不在成功失败中时是否返回失败
	DenyOtherStatus bool `json:"DenyOtherStatus"`

	// REQUIRED; 鉴权失败时返回的状态码，范围应在[400,699]内
	DenyReturnCode float32 `json:"DenyReturnCode"`

	// REQUIRED; 鉴权失败时鉴权服务器返回的状态码，所有状态码范围应在 [200,499] 之间，且和 AllowStatus 不重复。
	DenyRule UpdateRemoteAuthBodyDenyRule `json:"DenyRule"`

	// REQUIRED; 需要配置远程鉴权的拉流域名。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看拉流域名信息。
	Domain string `json:"Domain"`

	// REQUIRED; 远程鉴权配置是否打开
	Enable bool `json:"Enable"`

	// REQUIRED; 鉴权请求头参数，参数应在50个以内
	HeaderParamConfig UpdateRemoteAuthBodyHeaderParamConfig `json:"HeaderParamConfig"`

	// REQUIRED; 鉴权请求方法，支持传入POST或GET
	Method string `json:"Method"`

	// REQUIRED; 鉴权请求参数配置，参数应在50个以内
	QueryParamConfig UpdateRemoteAuthBodyQueryParamConfig `json:"QueryParamConfig"`

	// REQUIRED; 是否使用用户请求的路径拼接到鉴权地址上
	UseUserRequest bool `json:"UseUserRequest"`

	// REQUIRED; 需要配置远程鉴权的拉流域名所属的域名空间。您可以调用ListDomainDetail [https://www.volcengine.com/docs/6469/1126815]接口或在视频直播控制台的域名管理 [https://console.volcengine.com/live/main/domain/list]页面，查看拉流域名的域名空间信息。
	Vhost string `json:"Vhost"`

	// 远程鉴权路径，UseUserRequest为true时忽略此字段
	AuthURLPath *string `json:"AuthURLPath,omitempty"`

	// 鉴权请求body参数，参数不能超过50个
	BodyParams []*UpdateRemoteAuthBodyParamsItem `json:"BodyParams,omitempty"`

	// 鉴权缓存配置
	CacheConfig *UpdateRemoteAuthBodyCacheConfig `json:"CacheConfig,omitempty"`

	// hls鉴权开关
	EnableHLSAuth *bool `json:"EnableHLSAuth,omitempty"`

	// hls的ts片鉴权开关
	EnableHLSTSAuth *bool `json:"EnableHLSTSAuth,omitempty"`

	// 请求头是否保留大小写信息，true为原样请求，false转换为mime格式，不传默认为false
	HeaderCaseSensitivity *bool `json:"HeaderCaseSensitivity,omitempty"`

	// 长度不超过1024
	HeaderHost *string `json:"HeaderHost,omitempty"`

	// 鉴权失败重试间隔，单位为秒，范围[1,30]，不传入默认为10
	RetryInterval *float32 `json:"RetryInterval,omitempty"`

	// 鉴权失败重试次数，范围[0,10]，不传入默认为3
	RetryTimes *float32 `json:"RetryTimes,omitempty"`

	// 鉴权超时时间，单位为秒。范围[0, 600]，不填默认为1
	Timeout *float32 `json:"Timeout,omitempty"`
}

// UpdateRemoteAuthBodyCacheConfig - 鉴权缓存配置
type UpdateRemoteAuthBodyCacheConfig struct {

	// REQUIRED; 生成鉴权结果缓存 Key 使用的参数配置，最多支持配置 50 个参数。 :::tip
	// * 鉴权成功与鉴权失败使用相同配置时，表示鉴权结果缓存 Key 配置；
	// * 鉴权成功与鉴权失败不使用相同配置时，表示鉴权成功缓存的过期时间。 :::
	CacheKeys UpdateRemoteAuthBodyCacheConfigCacheKeys `json:"CacheKeys"`

	// REQUIRED; 鉴权成功失败是否使用相同的配置
	UseSameCache bool `json:"UseSameCache"`

	// 缓存过期时间，范围为[0,3600]，不传入默认为0
	CacheExpireSecond *float32 `json:"CacheExpireSecond,omitempty"`

	// 鉴权失败缓存的过期时间，范围为[0,3600]，不传入默认为0
	DenyExpireSecond *float32 `json:"DenyExpireSecond,omitempty"`

	// 生成鉴权失败结果缓存 Key 使用的参数配置，最多支持配置 50 个参数。
	// :::tip 鉴权成功与鉴权失败不使用相同配置，即 UseSameCache 配置 false 时生效。 :::
	DenyKeys *UpdateRemoteAuthBodyCacheConfigDenyKeys `json:"DenyKeys,omitempty"`
}

// UpdateRemoteAuthBodyCacheConfigCacheKeys - 生成鉴权结果缓存 Key 使用的参数配置，最多支持配置 50 个参数。 :::tip
// * 鉴权成功与鉴权失败使用相同配置时，表示鉴权结果缓存 Key 配置；
// * 鉴权成功与鉴权失败不使用相同配置时，表示鉴权成功缓存的过期时间。 :::
type UpdateRemoteAuthBodyCacheConfigCacheKeys struct {

	// REQUIRED; 缓存key的类型，包括query, header, vhost, domain, app, stream, client_ip, request_uri
	Type string `json:"Type"`

	// 缓存具体的key
	ParamName *string `json:"ParamName,omitempty"`
}

// UpdateRemoteAuthBodyCacheConfigDenyKeys - 生成鉴权失败结果缓存 Key 使用的参数配置，最多支持配置 50 个参数。
// :::tip 鉴权成功与鉴权失败不使用相同配置，即 UseSameCache 配置 false 时生效。 :::
type UpdateRemoteAuthBodyCacheConfigDenyKeys struct {

	// REQUIRED; 缓存key的类型，包括query, header, vhost, domain, app, stream, client_ip, request_uri
	Type string `json:"Type"`

	// 缓存具体的key
	ParamName *string `json:"ParamName,omitempty"`
}

// UpdateRemoteAuthBodyDenyRule - 鉴权失败时鉴权服务器返回的状态码，所有状态码范围应在 [200,499] 之间，且和 AllowStatus 不重复。
type UpdateRemoteAuthBodyDenyRule struct {

	// REQUIRED; 鉴权失败状态码，所有状态码范围应在[200,499]内，和SuccessCode不能有交集
	DenyStatus []float32 `json:"DenyStatus"`
}

// UpdateRemoteAuthBodyHeaderParamConfig - 鉴权请求头参数，参数应在50个以内
type UpdateRemoteAuthBodyHeaderParamConfig struct {

	// REQUIRED; 是否直接使用用户的请求头
	UseUserParam bool `json:"UseUserParam"`

	// 自定义鉴权请求的 Header 参数时的参数配置。
	Params []*UpdateRemoteAuthBodyHeaderParamConfigParamsItem `json:"Params,omitempty"`
}

type UpdateRemoteAuthBodyHeaderParamConfigParamsItem struct {

	// REQUIRED
	Type      string  `json:"Type"`
	ParamName *string `json:"ParamName,omitempty"`
	ToName    *string `json:"ToName,omitempty"`
	Value     *string `json:"Value,omitempty"`
}

type UpdateRemoteAuthBodyParamsItem struct {

	// REQUIRED; 参数类型，取值及含义如下所示。
	// * const_string：常量；
	// * header：用户请求的 Header 参数；
	// * query：用户请求的 URL 参数；
	// * vhost：参数值为变量 vhost 的参数，表示拉流请求中拉流域名所属的域名空间；
	// * domain：参数值为变量 domain 的参数，表示拉流请求中使用的拉流域名；
	// * app：参数值为变量 app 的参数，表示拉流请求中使用的 AppName；
	// * stream：参数值为变量 stream 的参数，表示拉流请求中使用的 StreamName；
	// * client_ip：参数值为变量 client_ip 的参数，表示拉流客户端 IP 地址；
	// * server_ip：参数值为变量 server_ip 的参数，表示响应拉流请求的 CDN 节点IP地址；
	// * request_uri：参数值为变量 request_uri 的参数，拉流请求地址的 URI。
	Type string `json:"Type"`

	// 参数名，最大长度为 100 个字符，不支持输入空格。
	// :::tip
	// * 参数类型为常量时表示常量参数的参数名；
	// * 参数类型为用户请求的 Header 参数或用户请求的 URL 参数时，表示指定用户请求中对应的参数名作为此处的参数名；
	// * 参数类型为变量时不生效。 :::
	ParamName *string `json:"ParamName,omitempty"`

	// 参数名的映射参数名，最大长度为 100 个字符，不支持输入空格。 :::tip
	// * 参数类型为常量时不生效；
	// * 参数类型为用户请求的 Header 参数或用户请求的 URL 参数时，表示鉴权请求时使用 ToName 值代替用户请求中对应的参数名；
	// * 参数类型为变量时，表示使用 ToName 取值作为此变量的参数名。 :::
	ToName *string `json:"ToName,omitempty"`

	// 参数类型为常量时的参数值，最大长度为 100 个字符，不支持输入空格。
	Value *string `json:"Value,omitempty"`
}

// UpdateRemoteAuthBodyQueryParamConfig - 鉴权请求参数配置，参数应在50个以内
type UpdateRemoteAuthBodyQueryParamConfig struct {

	// REQUIRED; 是否直接使用用户请求参数
	UseUserParam bool `json:"UseUserParam"`

	// 自定义参数时的参数配置，支持配置常量参数和变量参数总和最多不超过 25 个，提取用户请求参数最多不超过 25 个。
	Params []*UpdateRemoteAuthBodyQueryParamConfigParamsItem `json:"Params,omitempty"`
}

type UpdateRemoteAuthBodyQueryParamConfigParamsItem struct {

	// REQUIRED; 参数类型，可以传入const_string,header,query,vhost,domain,app,stream,client_ip,server_ip,request_uri
	Type string `json:"Type"`

	// 原query/header/body/变量的参数名，长度在100字符内，type为header、query、const_string时必传
	ParamName *string `json:"ParamName,omitempty"`

	// 鉴权请求中的query/header/body名，长度在100字符内，当不传入ToName时，新的参数名默认为原参数名
	ToName *string `json:"ToName,omitempty"`

	// 仅type=const_string时传入，常量的值，长度在100字符内
	Value *string `json:"Value,omitempty"`
}

type UpdateRemoteAuthRes struct {

	// REQUIRED
	ResponseMetadata UpdateRemoteAuthResResponseMetadata `json:"ResponseMetadata"`
	Result           *string                             `json:"Result,omitempty"`
}

type UpdateRemoteAuthResResponseMetadata struct {

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

type UpdateSubtitleTranscodePresetBody struct {

	// REQUIRED; 应用名称，您可以调用ListVhostSubtitleTranscodePreset [https://www.volcengine.com/docs/6469/1288712]接口，获取待更新字幕配置的 App 取值。
	App string `json:"App"`

	// REQUIRED; 截图配置的名称，您可以调用ListVhostSubtitleTranscodePreset [https://www.volcengine.com/docs/6469/1288712]接口，获取待更新字幕配置的 PresetName
	// 取值。
	PresetName string `json:"PresetName"`

	// REQUIRED; 原文字幕展示参数配置。
	SourceLanguage UpdateSubtitleTranscodePresetBodySourceLanguage `json:"SourceLanguage"`

	// REQUIRED; 关联转码配置后缀，一个字幕配置支持关联多个转码配置后缀。
	Suffixes []string `json:"Suffixes"`

	// REQUIRED; 域名空间，您可以调用 ListVhostSubtitleTranscodePreset [https://www.volcengine.com/docs/6469/1288712] 接口，获取待更新字幕配置的 Vhost
	// 取值。
	Vhost string `json:"Vhost"`

	// 字幕配置的描述信息。
	Description *string `json:"Description,omitempty"`

	// 预设配置，使用预设配置是系统将自动对字体大小、字幕行数、每行最大字符数和边距参数（MarginVertical 和 MarginHorizontal）进行智能化适配。默认为空，表示不使用预设配置，支持的预设配置如下所示。
	// * small ：小字幕。
	// * medium：中字幕。
	// * large：大字幕。 :::tip 使用预设配置时，字幕行数、每行最大字符数、左右边距和底部边距参数不生效，系统将使用预设配置自动进行计算。 :::
	DisplayPreset *string `json:"DisplayPreset,omitempty"`

	// 原文翻译成译文时使用的热词词库，总长度不超过 10000 个字符，默认为空。
	GlossaryWordList []*string `json:"GlossaryWordList,omitempty"`

	// 原文字幕识别时使用的热词词库，总长度不超过为 10000 个字符，默认为空。
	HotWordList []*string `json:"HotWordList,omitempty"`

	// 设置在 16:9 分辨率场景下，每行字幕展示的最大字符数。 :::tip
	// * 使用预设配置时，字幕每行最大字符数设置不生效。
	// * 不使用预设配置时，字幕每行最大字符数必填。
	// * 每个文字、字母、符号或数字均为一个字符。
	// * 当屏幕分辨率改变时，屏幕上显示的每行文字数量会相应调整，以适应新的分辨率，确保文字的显示效果和阅读体验。 :::
	MaxCharNumber *int32 `json:"MaxCharNumber,omitempty"`

	// 字幕展示的行数，同时适用于原文字幕和译文字幕，支持的取值及含义如下所示。
	// * 0：（默认值）根据字幕字数自动进行分行展示；
	// * 1：每种字幕展示一行；
	// * 2：每种字幕展示两行。 :::tip
	// * 使用预设配置时，字幕行数为自动分行展示。
	// * 超出行内字数限制时表示字幕将超过显示范围，此时字幕内容将被截断。 :::
	MaxRowNumber *int32 `json:"MaxRowNumber,omitempty"`

	// 字幕位置设置，通过设置字幕距离画面左右边距和底部边距来指定字幕位置。
	// :::tip
	// * 使用预设配置时，字幕位置设置不生效。
	// * 不使用预设配置时，字幕位置设置必填。 :::
	Position *UpdateSubtitleTranscodePresetBodyPosition `json:"Position,omitempty"`

	// 译文字幕展示参数配置列表，当前最多支持配置一种译文。
	TargetLanguage []*UpdateSubtitleTranscodePresetBodyTargetLanguageItem `json:"TargetLanguage,omitempty"`
}

// UpdateSubtitleTranscodePresetBodyPosition - 字幕位置设置，通过设置字幕距离画面左右边距和底部边距来指定字幕位置。
// :::tip
// * 使用预设配置时，字幕位置设置不生效。
// * 不使用预设配置时，字幕位置设置必填。 :::
type UpdateSubtitleTranscodePresetBodyPosition struct {

	// 字幕距离画面两侧的边距与画面宽度的占比，使用归一化百分表示，取值范围为 [0,0.2]。
	MarginHorizontal *float32 `json:"MarginHorizontal,omitempty"`

	// 字幕距离画面底部的边距与画面高度的占比，使用归一化百分表示，取值范围为 [0,0.5]。
	MarginVertical *float32 `json:"MarginVertical,omitempty"`
	Relative       *string  `json:"Relative,omitempty"`
}

// UpdateSubtitleTranscodePresetBodySourceLanguage - 原文字幕展示参数配置。
type UpdateSubtitleTranscodePresetBodySourceLanguage struct {

	// REQUIRED; 是否展示原文字幕，取值及含义如下所示。
	// * true：展示，此时将展示原文和译文双语字幕
	// * false：不展示，此时将只展示译文字幕。
	// :::tip 原文字幕语言和译文字幕语言相同时，仅展示译文字幕。 :::
	Display bool `json:"Display"`

	// REQUIRED; 原文字幕的字体，原文字幕字体根据原文字幕语言取值不同而不同，取值及含义如下所示。
	// * 当原文字幕的语言是 zh 时，支持以下字体取值。 * siyuanheiti：思源黑体；
	// * songtixi：宋体细；
	// * songticu：宋体粗；
	// * heitifan：黑体繁；
	// * kaiti：楷体。
	//
	//
	// * 当原文字幕的语言是 en 时，支持以下字体取值。 * inter：Inter；
	// * roboto：Roboto；
	// * opposans：OPPOSans；
	// * siyuansongti：思源宋体；
	// * montserrat：Montserrat。
	//
	//
	// * 当原文字幕的语言是 ko 和 ja 时，支持 notosans(Noto Sans) 字体。
	Font string `json:"Font"`

	// REQUIRED; 原文字幕的字体颜色，支持以下几种方法进行定义。
	// * 支持以 0x 或 # 开头，后面跟着十六进制颜色 RGB 值，再跟着 @+十六进制/百分比来表示的透明度值，来定义字幕的字体颜色。例如，设置 RGB 值为 FF0000，透明度为 5%的颜色时，您可以传入 0xFF0000@0x80、0xFF0000@0.5、#FF0000@0x80
	// 或 #FF0000@0.5。
	// * 支持使用前端框架 FFmpeg 规定的颜色关键字，来定义字幕的字体颜色。例如，AliceBlue 表示 0xF0F8FF、AntiqueWhite 表示 0xFAEBD7、Black 表示 0x000000 等。 :::tip 查看详细颜色定义方法及更多颜色关键字，请参考
	// FFmpeg 的颜色定义语法
	// [https://ffmpeg.org/ffmpeg-utils.html#color-syntax]。 :::
	FontColor string `json:"FontColor"`

	// REQUIRED; 原文字幕的语言，取值及含义如下所示。
	// * zh：中英混合；
	// * en：英语；
	// * ko：韩语；
	// * ja：日语。
	Language string `json:"Language"`

	// 原文字幕的阴影配置。
	Border *UpdateSubtitleTranscodePresetBodySourceLanguageBorder `json:"Border,omitempty"`

	// 原文字幕的字体大小，单位为 px，默认为空。 :::tip
	// * 使用了预设配置时，字幕字体大小设置不生效。
	// * 不使用预设配置时，字幕字体大小为必选参数。 :::
	FontSize *int32 `json:"FontSize,omitempty"`
}

// UpdateSubtitleTranscodePresetBodySourceLanguageBorder - 原文字幕的阴影配置。
type UpdateSubtitleTranscodePresetBodySourceLanguageBorder struct {

	// REQUIRED; 描边的颜色，支持以下几种方法进行定义。
	// * 支持以 0x 或 # 开头，后面跟着十六进制颜色 RGB 值，再跟着 @+十六进制/百分比来表示的透明度值，来定义字幕的字体颜色。例如，设置 RGB 值为 FF0000，透明度为 5%的颜色时，您可以传入 0xFF0000@0x80、0xFF0000@0.5、#FF0000@0x80
	// 或 #FF0000@0.5。
	// * 支持使用前端框架 FFmpeg 规定的颜色关键字，来定义字幕的字体颜色。例如，AliceBlue 表示 0xF0F8FF、AntiqueWhite 表示 0xFAEBD7、Black 表示 0x000000 等。 :::tip 查看详细颜色定义方法及更多颜色关键字，请参考
	// FFmpeg 的颜色定义语法
	// [https://ffmpeg.org/ffmpeg-utils.html#color-syntax]。 :::
	Color string `json:"Color"`

	// 填0的时候后端根据字体大小进行计算，字体大小/32*1.25
	Width *float32 `json:"Width,omitempty"`
}

type UpdateSubtitleTranscodePresetBodyTargetLanguageItem struct {

	// REQUIRED; 译文字幕的字体，译文字幕字体根据译文字幕语言取值不同而不同，取值及含义如下所示。
	// * 当译文字幕的语言是 zh 时，支持以下字体取值。 * siyuanheiti：思源黑体；
	// * songtixi：宋体细；
	// * songticu：宋体粗；
	// * heitifan：黑体繁；
	// * kaiti：楷体。
	//
	//
	// * 当译文字幕的语言是 zh-Hant 时，支持 siyuanheiti （思源黑体）字体。
	// * 当译文字幕的语言是 en 时，支持以下字体取值。 * inter：Inter；
	// * roboto：Roboto；
	// * opposans：OPPOSans；
	// * siyuansongti：思源宋体；
	// * montserrat：Montserrat。
	//
	//
	// * 当译文字幕的语言是 ko、ja、ar、de、es、fr、hi、pt、 ru、 vi、 th 时，支持 notosans(Noto Sans) 字体。
	Font string `json:"Font"`

	// REQUIRED; 译文字幕的字体颜色，支持以下几种方法进行定义。
	// * 支持以 0x 或 # 开头，后面跟着十六进制颜色 RGB 值，再跟着 @+十六进制/百分比来表示的透明度值，来定义字幕的字体颜色。例如，设置 RGB 值为 FF0000，透明度为 5%的颜色时，您可以传入 0xFF0000@0x80、0xFF0000@0.5、#FF0000@0x80
	// 或 #FF0000@0.5。
	// * 支持使用前端框架 FFmpeg 规定的颜色关键字，来定义字幕的字体颜色。例如，AliceBlue 表示 0xF0F8FF、AntiqueWhite 表示 0xFAEBD7、Black 表示 0x000000 等。 :::tip 查看详细颜色定义方法及更多颜色关键字，请参考
	// FFmpeg 的颜色定义语法
	// [https://ffmpeg.org/ffmpeg-utils.html#color-syntax]。 :::
	FontColor string `json:"FontColor"`

	// REQUIRED; 译文字幕的语言，取值及含义如下所示。
	// * zh：中英混合；
	// * zh-Hant：繁体中文；
	// * en：英语；
	// * ko：韩语；
	// * ja：日语；
	// * ar：阿拉伯语；
	// * de：德语；
	// * es：西班牙语；
	// * fr：法语；
	// * hi：印地语；
	// * pt：葡萄牙语；
	// * ru：俄语；
	// * vi：越南语；
	// * th：泰语。
	Language string `json:"Language"`

	// 填0的时候后端根据字体大小进行计算，字体大小/32*1.25
	Border *UpdateSubtitleTranscodePresetBodyTargetLanguageItemBorder `json:"Border,omitempty"`

	// 译文字幕的字体大小，单位为 px，默认为空。 :::tip
	// * 使用预设配置时，字幕字体大小设置不生效。
	// * 不使用预设配置时，字幕字体大小为必选参数。 :::
	FontSize *int32 `json:"FontSize,omitempty"`
}

// UpdateSubtitleTranscodePresetBodyTargetLanguageItemBorder - 填0的时候后端根据字体大小进行计算，字体大小/32*1.25
type UpdateSubtitleTranscodePresetBodyTargetLanguageItemBorder struct {

	// REQUIRED
	Color string   `json:"Color"`
	Width *float32 `json:"Width,omitempty"`
}

type UpdateSubtitleTranscodePresetRes struct {

	// REQUIRED
	ResponseMetadata UpdateSubtitleTranscodePresetResResponseMetadata `json:"ResponseMetadata"`

	// 视请求的接口而定
	Result interface{} `json:"Result,omitempty"`
}

type UpdateSubtitleTranscodePresetResResponseMetadata struct {

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

type UpdateTimeShiftPresetV2Body struct {

	// REQUIRED
	App string `json:"App"`

	// REQUIRED
	MaxShiftTime int32 `json:"MaxShiftTime"`

	// REQUIRED
	Preset string `json:"Preset"`

	// REQUIRED
	Type string `json:"Type"`

	// REQUIRED
	Vhost        string  `json:"Vhost"`
	MasterFormat *string `json:"MasterFormat,omitempty"`
	Status       *int32  `json:"Status,omitempty"`
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

	// IDR 帧之间的最大间隔时间，单位为秒，取值范围为 [1,30]。
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
	// * h264：使用 H.264 的视频编码格式；
	// * h265：使用 H.265 的视频编码格式；
	// * h266：使用 H.266 的视频编码格式；
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
	IsAbr        *string                                                                                                                    `json:"IsAbr,omitempty"`
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

	// 流名称，您可以调用ListVhostWatermarkPreset [https://www.volcengine.com/docs/6469/1126889]接口，查看待更新水印配置的 Stream 取值。
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
type ContinuePullToPushTask struct{}
type ContinuePullToPushTaskQuery struct{}
type CreateCarouselTask struct{}
type CreateCarouselTaskQuery struct{}
type CreateCert struct{}
type CreateCertQuery struct{}
type CreateCloudMixTask struct{}
type CreateCloudMixTaskQuery struct{}
type CreateDomainV2 struct{}
type CreateDomainV2Query struct{}
type CreateLiveVideoQualityAnalysisTask struct{}
type CreateLiveVideoQualityAnalysisTaskQuery struct{}
type CreatePullToPushGroup struct{}
type CreatePullToPushGroupQuery struct{}
type CreatePullToPushTask struct{}
type CreatePullToPushTaskQuery struct{}
type CreateRecordPresetV2 struct{}
type CreateRecordPresetV2Query struct{}
type CreateSnapshotPreset struct{}
type CreateSnapshotPresetQuery struct{}
type CreateSubtitleTranscodePreset struct{}
type CreateSubtitleTranscodePresetQuery struct{}
type CreateTimeShiftPresetV2 struct{}
type CreateTimeShiftPresetV2Query struct{}
type CreateTranscodePreset struct{}
type CreateTranscodePresetBatch struct{}
type CreateTranscodePresetBatchQuery struct{}
type CreateTranscodePresetQuery struct{}
type CreateWatermarkPreset struct{}
type CreateWatermarkPresetQuery struct{}
type DeleteCMAFConfig struct{}
type DeleteCMAFConfigQuery struct{}
type DeleteCallback struct{}
type DeleteCallbackQuery struct{}
type DeleteCarouselTask struct{}
type DeleteCarouselTaskQuery struct{}
type DeleteCert struct{}
type DeleteCertQuery struct{}
type DeleteCloudMixTask struct{}
type DeleteCloudMixTaskQuery struct{}
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
type DeleteLiveVideoQualityAnalysisTask struct{}
type DeleteLiveVideoQualityAnalysisTaskQuery struct{}
type DeletePullToPushGroup struct{}
type DeletePullToPushGroupQuery struct{}
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
type DeleteRemoteAuth struct{}
type DeleteRemoteAuthQuery struct{}
type DeleteSnapshotPreset struct{}
type DeleteSnapshotPresetQuery struct{}
type DeleteSubtitleTranscodePreset struct{}
type DeleteSubtitleTranscodePresetQuery struct{}
type DeleteTimeShiftPresetV2 struct{}
type DeleteTimeShiftPresetV2Query struct{}
type DeleteTranscodePreset struct{}
type DeleteTranscodePresetBatch struct{}
type DeleteTranscodePresetBatchQuery struct{}
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
type DescribeEncryptHLS struct{}
type DescribeEncryptHLSBody struct{}
type DescribeEncryptHLSQuery struct{}
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
type DescribeLatencyConfig struct{}
type DescribeLatencyConfigQuery struct{}
type DescribeLicenseDRM struct{}
type DescribeLicenseDRMBody struct{}
type DescribeLiveASRDurationData struct{}
type DescribeLiveASRDurationDataQuery struct{}
type DescribeLiveBandwidthData struct{}
type DescribeLiveBandwidthDataQuery struct{}
type DescribeLiveBatchPushStreamMetrics struct{}
type DescribeLiveBatchPushStreamMetricsQuery struct{}
type DescribeLiveBatchStreamTranscodeData struct{}
type DescribeLiveBatchStreamTranscodeDataQuery struct{}
type DescribeLiveCallbackData struct{}
type DescribeLiveCallbackDataQuery struct{}
type DescribeLiveEdgeStatData struct{}
type DescribeLiveEdgeStatDataQuery struct{}
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
type DescribeLiveTopPlayData struct{}
type DescribeLiveTopPlayDataQuery struct{}
type DescribeLiveTrafficData struct{}
type DescribeLiveTrafficDataQuery struct{}
type DescribeLiveTranscodeData struct{}
type DescribeLiveTranscodeDataQuery struct{}
type DescribeLiveTranscodeInfoData struct{}
type DescribeLiveTranscodeInfoDataQuery struct{}
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
type DescribeRemoteAuth struct{}
type DescribeRemoteAuthQuery struct{}
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
type GetCarouselDetail struct{}
type GetCarouselDetailQuery struct{}
type GetCloudMixTaskDetail struct{}
type GetCloudMixTaskDetailQuery struct{}
type GetHLSEncryptDataKey struct{}
type GetHLSEncryptDataKeyBody struct{}
type GetLiveVideoQualityAnalysisTaskDetail struct{}
type GetLiveVideoQualityAnalysisTaskDetailQuery struct{}
type KillStream struct{}
type KillStreamQuery struct{}
type ListBindEncryptDRM struct{}
type ListBindEncryptDRMQuery struct{}
type ListCarouselTask struct{}
type ListCarouselTaskQuery struct{}
type ListCertV2 struct{}
type ListCertV2Query struct{}
type ListCloudMixTask struct{}
type ListCloudMixTaskQuery struct{}
type ListCommonTransPresetDetail struct{}
type ListCommonTransPresetDetailQuery struct{}
type ListDomainDetail struct{}
type ListDomainDetailQuery struct{}
type ListLiveVideoQualityAnalysisTasks struct{}
type ListLiveVideoQualityAnalysisTasksQuery struct{}
type ListPullToPushGroup struct{}
type ListPullToPushGroupQuery struct{}
type ListPullToPushTask struct{}
type ListPullToPushTaskBody struct{}
type ListPullToPushTaskV2 struct{}
type ListPullToPushTaskV2Query struct{}
type ListTimeShiftPresetV2 struct{}
type ListTimeShiftPresetV2Query struct{}
type ListVhostRecordPresetV2 struct{}
type ListVhostRecordPresetV2Query struct{}
type ListVhostRemoteAuth struct{}
type ListVhostRemoteAuthQuery struct{}
type ListVhostSnapshotPreset struct{}
type ListVhostSnapshotPresetQuery struct{}
type ListVhostSubtitleTranscodePreset struct{}
type ListVhostSubtitleTranscodePresetQuery struct{}
type ListVhostTransCodePreset struct{}
type ListVhostTransCodePresetQuery struct{}
type ListVhostWatermarkPreset struct{}
type ListVhostWatermarkPresetQuery struct{}
type ListWatermarkPreset struct{}
type ListWatermarkPresetQuery struct{}
type RelaunchPullToPushTask struct{}
type RelaunchPullToPushTaskQuery struct{}
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
type UpdateCarouselTask struct{}
type UpdateCarouselTaskQuery struct{}
type UpdateCloudMixTask struct{}
type UpdateCloudMixTaskQuery struct{}
type UpdateClusterRateLimit struct{}
type UpdateClusterRateLimitQuery struct{}
type UpdateDomainVhost struct{}
type UpdateDomainVhostQuery struct{}
type UpdateEncryptDRM struct{}
type UpdateEncryptDRMQuery struct{}
type UpdateEncryptHLS struct{}
type UpdateEncryptHLSQuery struct{}
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
type UpdateRemoteAuth struct{}
type UpdateRemoteAuthQuery struct{}
type UpdateSnapshotPreset struct{}
type UpdateSnapshotPresetQuery struct{}
type UpdateSubtitleTranscodePreset struct{}
type UpdateSubtitleTranscodePresetQuery struct{}
type UpdateTimeShiftPresetV2 struct{}
type UpdateTimeShiftPresetV2Query struct{}
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
type ContinuePullToPushTaskReq struct {
	*ContinuePullToPushTaskQuery
	*ContinuePullToPushTaskBody
}
type CreateCarouselTaskReq struct {
	*CreateCarouselTaskQuery
	*CreateCarouselTaskBody
}
type CreateCertReq struct {
	*CreateCertQuery
	*CreateCertBody
}
type CreateCloudMixTaskReq struct {
	*CreateCloudMixTaskQuery
	*CreateCloudMixTaskBody
}
type CreateDomainV2Req struct {
	*CreateDomainV2Query
	*CreateDomainV2Body
}
type CreateLiveVideoQualityAnalysisTaskReq struct {
	*CreateLiveVideoQualityAnalysisTaskQuery
	*CreateLiveVideoQualityAnalysisTaskBody
}
type CreatePullToPushGroupReq struct {
	*CreatePullToPushGroupQuery
	*CreatePullToPushGroupBody
}
type CreatePullToPushTaskReq struct {
	*CreatePullToPushTaskQuery
	*CreatePullToPushTaskBody
}
type CreateRecordPresetV2Req struct {
	*CreateRecordPresetV2Query
	*CreateRecordPresetV2Body
}
type CreateSnapshotPresetReq struct {
	*CreateSnapshotPresetQuery
	*CreateSnapshotPresetBody
}
type CreateSubtitleTranscodePresetReq struct {
	*CreateSubtitleTranscodePresetQuery
	*CreateSubtitleTranscodePresetBody
}
type CreateTimeShiftPresetV2Req struct {
	*CreateTimeShiftPresetV2Query
	*CreateTimeShiftPresetV2Body
}
type CreateTranscodePresetReq struct {
	*CreateTranscodePresetQuery
	*CreateTranscodePresetBody
}
type CreateTranscodePresetBatchReq struct {
	*CreateTranscodePresetBatchQuery
	*CreateTranscodePresetBatchBody
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
type DeleteCarouselTaskReq struct {
	*DeleteCarouselTaskQuery
	*DeleteCarouselTaskBody
}
type DeleteCertReq struct {
	*DeleteCertQuery
	*DeleteCertBody
}
type DeleteCloudMixTaskReq struct {
	*DeleteCloudMixTaskQuery
	*DeleteCloudMixTaskBody
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
type DeleteLiveVideoQualityAnalysisTaskReq struct {
	*DeleteLiveVideoQualityAnalysisTaskQuery
	*DeleteLiveVideoQualityAnalysisTaskBody
}
type DeletePullToPushGroupReq struct {
	*DeletePullToPushGroupQuery
	*DeletePullToPushGroupBody
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
type DeleteRemoteAuthReq struct {
	*DeleteRemoteAuthQuery
	*DeleteRemoteAuthBody
}
type DeleteSnapshotPresetReq struct {
	*DeleteSnapshotPresetQuery
	*DeleteSnapshotPresetBody
}
type DeleteSubtitleTranscodePresetReq struct {
	*DeleteSubtitleTranscodePresetQuery
	*DeleteSubtitleTranscodePresetBody
}
type DeleteTimeShiftPresetV2Req struct {
	*DeleteTimeShiftPresetV2Query
	*DeleteTimeShiftPresetV2Body
}
type DeleteTranscodePresetReq struct {
	*DeleteTranscodePresetQuery
	*DeleteTranscodePresetBody
}
type DeleteTranscodePresetBatchReq struct {
	*DeleteTranscodePresetBatchQuery
	*DeleteTranscodePresetBatchBody
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
type DescribeEncryptHLSReq struct {
	*DescribeEncryptHLSQuery
	*DescribeEncryptHLSBody
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
type DescribeLiveBatchPushStreamMetricsReq struct {
	*DescribeLiveBatchPushStreamMetricsQuery
	*DescribeLiveBatchPushStreamMetricsBody
}
type DescribeLiveBatchStreamTranscodeDataReq struct {
	*DescribeLiveBatchStreamTranscodeDataQuery
	*DescribeLiveBatchStreamTranscodeDataBody
}
type DescribeLiveCallbackDataReq struct {
	*DescribeLiveCallbackDataQuery
	*DescribeLiveCallbackDataBody
}
type DescribeLiveEdgeStatDataReq struct {
	*DescribeLiveEdgeStatDataQuery
	*DescribeLiveEdgeStatDataBody
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
type DescribeLiveTopPlayDataReq struct {
	*DescribeLiveTopPlayDataQuery
	*DescribeLiveTopPlayDataBody
}
type DescribeLiveTrafficDataReq struct {
	*DescribeLiveTrafficDataQuery
	*DescribeLiveTrafficDataBody
}
type DescribeLiveTranscodeDataReq struct {
	*DescribeLiveTranscodeDataQuery
	*DescribeLiveTranscodeDataBody
}
type DescribeLiveTranscodeInfoDataReq struct {
	*DescribeLiveTranscodeInfoDataQuery
	*DescribeLiveTranscodeInfoDataBody
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
type DescribeRemoteAuthReq struct {
	*DescribeRemoteAuthQuery
	*DescribeRemoteAuthBody
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
type GetCarouselDetailReq struct {
	*GetCarouselDetailQuery
	*GetCarouselDetailBody
}
type GetCloudMixTaskDetailReq struct {
	*GetCloudMixTaskDetailQuery
	*GetCloudMixTaskDetailBody
}
type GetHLSEncryptDataKeyReq struct {
	*GetHLSEncryptDataKeyQuery
	*GetHLSEncryptDataKeyBody
}
type GetLiveVideoQualityAnalysisTaskDetailReq struct {
	*GetLiveVideoQualityAnalysisTaskDetailQuery
	*GetLiveVideoQualityAnalysisTaskDetailBody
}
type KillStreamReq struct {
	*KillStreamQuery
	*KillStreamBody
}
type ListBindEncryptDRMReq struct {
	*ListBindEncryptDRMQuery
	*ListBindEncryptDRMBody
}
type ListCarouselTaskReq struct {
	*ListCarouselTaskQuery
	*ListCarouselTaskBody
}
type ListCertV2Req struct {
	*ListCertV2Query
	*ListCertV2Body
}
type ListCloudMixTaskReq struct {
	*ListCloudMixTaskQuery
	*ListCloudMixTaskBody
}
type ListCommonTransPresetDetailReq struct {
	*ListCommonTransPresetDetailQuery
	*ListCommonTransPresetDetailBody
}
type ListDomainDetailReq struct {
	*ListDomainDetailQuery
	*ListDomainDetailBody
}
type ListLiveVideoQualityAnalysisTasksReq struct {
	*ListLiveVideoQualityAnalysisTasksQuery
	*ListLiveVideoQualityAnalysisTasksBody
}
type ListPullToPushGroupReq struct {
	*ListPullToPushGroupQuery
	*ListPullToPushGroupBody
}
type ListPullToPushTaskReq struct {
	*ListPullToPushTaskQuery
	*ListPullToPushTaskBody
}
type ListPullToPushTaskV2Req struct {
	*ListPullToPushTaskV2Query
	*ListPullToPushTaskV2Body
}
type ListTimeShiftPresetV2Req struct {
	*ListTimeShiftPresetV2Query
	*ListTimeShiftPresetV2Body
}
type ListVhostRecordPresetV2Req struct {
	*ListVhostRecordPresetV2Query
	*ListVhostRecordPresetV2Body
}
type ListVhostRemoteAuthReq struct {
	*ListVhostRemoteAuthQuery
	*ListVhostRemoteAuthBody
}
type ListVhostSnapshotPresetReq struct {
	*ListVhostSnapshotPresetQuery
	*ListVhostSnapshotPresetBody
}
type ListVhostSubtitleTranscodePresetReq struct {
	*ListVhostSubtitleTranscodePresetQuery
	*ListVhostSubtitleTranscodePresetBody
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
type RelaunchPullToPushTaskReq struct {
	*RelaunchPullToPushTaskQuery
	*RelaunchPullToPushTaskBody
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
type UpdateCarouselTaskReq struct {
	*UpdateCarouselTaskQuery
	*UpdateCarouselTaskBody
}
type UpdateCloudMixTaskReq struct {
	*UpdateCloudMixTaskQuery
	*UpdateCloudMixTaskBody
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
type UpdateEncryptHLSReq struct {
	*UpdateEncryptHLSQuery
	*UpdateEncryptHLSBody
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
type UpdateRemoteAuthReq struct {
	*UpdateRemoteAuthQuery
	*UpdateRemoteAuthBody
}
type UpdateSnapshotPresetReq struct {
	*UpdateSnapshotPresetQuery
	*UpdateSnapshotPresetBody
}
type UpdateSubtitleTranscodePresetReq struct {
	*UpdateSubtitleTranscodePresetQuery
	*UpdateSubtitleTranscodePresetBody
}
type UpdateTimeShiftPresetV2Req struct {
	*UpdateTimeShiftPresetV2Query
	*UpdateTimeShiftPresetV2Body
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
