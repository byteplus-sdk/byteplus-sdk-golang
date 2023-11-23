package livesaas

import (
	"github.com/byteplus-sdk/byteplus-sdk-golang/base"
)

type GetTemporaryLoginTokenParam struct {
	ActivityId int64 `thrift:"ActivityId,1" frugal:"1,default,i64" json:"ActivityId"`
}

type GetTemporaryLoginTokenResponseBody struct {
	SecretKey  string `thrift:"SecretKey,1" frugal:"1,default,string" json:"SecretKey"`
	ActivityId int64  `thrift:"ActivityId,2" frugal:"2,default,i64" json:"ActivityId"`
}

type GetTemporaryLoginTokenResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *GetTemporaryLoginTokenResponseBody `json:"Result,omitempty"`
}

type UpdateActivityParam struct {
	AccountId                int64              `thrift:"AccountId,1" frugal:"1,default,i64" json:"AccountId"`
	AccountUserName          string             `thrift:"AccountUserName,2" frugal:"2,default,string" json:"AccountUserName"`
	Id                       int64              `thrift:"Id,3" frugal:"3,default,i64" json:"Id"`
	Name                     string             `thrift:"Name,4" frugal:"4,default,string" json:"Name"`
	LiveTime                 int64              `thrift:"LiveTime,5" frugal:"5,default,i64" json:"LiveTime"`
	ViewUrlPath              string             `thrift:"ViewUrlPath,6" frugal:"6,default,string" json:"ViewUrlPath"`
	CoverImage               string             `thrift:"CoverImage,7" frugal:"7,default,string" json:"CoverImage"`
	OldId                    int64              `thrift:"OldId,8" frugal:"8,default,i64" json:"OldId"`
	CopyStream               bool               `thrift:"CopyStream,9" frugal:"9,default,bool" json:"CopyStream"`
	ProtectModeCode          string             `thrift:"ProtectModeCode,10" frugal:"10,default,string" json:"ProtectModeCode"`
	LiveType                 int32              `thrift:"LiveType,11" frugal:"11,default,i32" json:"LiveType"`
	LiveSource               string             `thrift:"LiveSource,12" frugal:"12,default,string" json:"LiveSource"`
	MediaName                string             `thrift:"MediaName,13" frugal:"13,default,string" json:"MediaName"`
	MediaDuration            int32              `thrift:"MediaDuration,14" frugal:"14,default,i32" json:"MediaDuration"`
	OnLineStatus             int32              `thrift:"OnLineStatus,15" frugal:"15,default,i32" json:"OnLineStatus"`
	SiteTags                 []*SiteActivityTag `thrift:"SiteTags,16" frugal:"16,default,list<SiteActivityTag>" json:"SiteTags"`
	IsReplayAutoOnlineEnable int32              `thrift:"IsReplayAutoOnlineEnable,17" frugal:"17,default,i32" json:"IsReplayAutoOnlineEnable"`
	TemplateId               int64              `thrift:"TemplateId,18" frugal:"18,default,i64" json:"TemplateId"`
	TenantID                 string             `thrift:"TenantID,19" frugal:"19,default,string" json:"TenantID"`
	ColorThemeIndex          string             `thrift:"ColorThemeIndex,20" frugal:"20,default,string" json:"ColorThemeIndex"`
	ConfigVersion            int32              `thrift:"ConfigVersion,21" frugal:"21,default,i32" json:"ConfigVersion"`
	LiveZone                 int32              `thrift:"LiveZone,22" frugal:"22,default,i32" json:"LiveZone"`
	Extra                    string             `thrift:"Extra,23" frugal:"23,default,string" json:"Extra"`
}

type SiteActivityTag struct {
	Index   int32  `thrift:"Index,1" frugal:"1,default,i32" json:"Index"`
	Value   string `thrift:"Value,2" frugal:"2,default,string" json:"Value"`
	DbIndex int32  `thrift:"DbIndex,3" frugal:"3,default,i32" json:"DbIndex"`
	Show    int32  `thrift:"Show,4" frugal:"4,default,i32" json:"Show"`
	Name    string `thrift:"Name,5" frugal:"5,default,string" json:"Name"`
}

type CommonResp struct {
	ActivityId int64 `thrift:"ActivityId,1" frugal:"1,default,i64" json:"ActivityId"`
}

type CommonResponseResp struct {
	ResponseMetadata base.ResponseMetadata
	Result           *CommonResp `json:"Result,omitempty"`
}

type GetSDKTokenRequestParam struct {
	ActivityId int64  `thrift:"ActivityId,1" frugal:"1,default,i64" json:"ActivityId"`
	Mode       int32  `thrift:"Mode,2" frugal:"2,default,i32" json:"Mode"`
	UserIdStr  string `thrift:"UserIdStr,3" frugal:"3,default,string" json:"UserIdStr"`
	Nickname   string `thrift:"Nickname,4" frugal:"4,default,string" json:"Nickname"`
}

type GetSDKTokenResponseBody struct {
	Token string `thrift:"Token,2" frugal:"2,default,string" json:"Token"`
}

type GetSDKTokenResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *GetSDKTokenResponseBody `json:"Result,omitempty"`
}

// ListActivityAPI 房间列表

type ListActivityParam struct {
	AccountId        int64                 `thrift:"AccountId,1" frugal:"1,default,i64" json:"AccountId"`
	AccountUserName  string                `thrift:"AccountUserName,2" frugal:"2,default,string" json:"AccountUserName"`
	FollowerUserName string                `thrift:"FollowerUserName,3" frugal:"3,default,string" json:"FollowerUserName"`
	Name             string                `thrift:"Name,4" frugal:"4,default,string" json:"Name"`
	Status           int32                 `thrift:"Status,5" frugal:"5,default,i32" json:"Status"`
	PageNo           int32                 `thrift:"PageNo,6" frugal:"6,default,i32" json:"PageNo"`
	PageItemCount    int32                 `thrift:"PageItemCount,7" frugal:"7,default,i32" json:"PageItemCount"`
	IsNeedTotalCount bool                  `thrift:"IsNeedTotalCount,8" frugal:"8,default,bool" json:"IsNeedTotalCount"`
	ProtectModeCode  string                `thrift:"ProtectModeCode,9" frugal:"9,default,string" json:"ProtectModeCode"`
	LiveTime         int64                 `thrift:"LiveTime,10" frugal:"10,default,i64" json:"LiveTime"`
	SiteTags         []*SiteActivityTag    `thrift:"SiteTags,11" frugal:"11,default,list<SiteActivityTag>" json:"SiteTags"`
	OnlineStatus     int32                 `thrift:"OnlineStatus,12" frugal:"12,default,i32" json:"OnlineStatus"`
	IsLockPreview    int32                 `thrift:"IsLockPreview,13" frugal:"13,default,i32" json:"IsLockPreview"`
	NeedTag          *bool                 `thrift:"NeedTag,14,optional" frugal:"14,optional,bool" json:"NeedTag,omitempty"`
	SiteTagNews      []*SiteActivityTagNew `thrift:"SiteTagNews,15" frugal:"15,default,list<SiteActivityTagNew>" json:"SiteTagNews"`
	AccountUserId    string                `thrift:"AccountUserId,16" frugal:"16,default,string" json:"AccountUserId"`
}

type SiteActivityTagNew struct {
	Index   int32    `thrift:"Index,1" frugal:"1,default,i32" json:"Index"`
	Value   []string `thrift:"Value,2" frugal:"2,default,list<string>" json:"Value"`
	DbIndex int32    `thrift:"DbIndex,3" frugal:"3,default,i32" json:"DbIndex"`
	Show    int32    `thrift:"Show,4" frugal:"4,default,i32" json:"Show"`
	Name    string   `thrift:"Name,5" frugal:"5,default,string" json:"Name"`
	ID      int64    `thrift:"ID,6" frugal:"6,default,i64" json:"ID"`
}

type ListActivityAPIForm struct {
	Id              int64              `thrift:"Id,1" frugal:"1,default,i64" json:"Id"`
	Name            string             `thrift:"Name,2" frugal:"2,default,string" json:"Name"`
	ViewUrl         string             `thrift:"ViewUrl,3" frugal:"3,default,string" json:"ViewUrl"`
	Status          int32              `thrift:"Status,4" frugal:"4,default,i32" json:"Status"`
	CoverImage      string             `thrift:"CoverImage,5" frugal:"5,default,string" json:"CoverImage"`
	CreateTime      int64              `thrift:"CreateTime,6" frugal:"6,default,i64" json:"CreateTime"`
	LiveTime        int64              `thrift:"LiveTime,7" frugal:"7,default,i64" json:"LiveTime"`
	OnlineStatus    int32              `thrift:"OnlineStatus,8" frugal:"8,default,i32" json:"OnlineStatus"`
	SiteTags        []*SiteActivityTag `thrift:"SiteTags,9" frugal:"9,default,list<SiteActivityTag>" json:"SiteTags"`
	StreamStartTime int64              `thrift:"StreamStartTime,10" frugal:"10,default,i64" json:"StreamStartTime"`
	IsLockPreview   int32              `thrift:"IsLockPreview,11" frugal:"11,default,i32" json:"IsLockPreview"`
}

type ListActivityAPIResp struct {
	PageNo         int32                  `thrift:"PageNo,1" frugal:"1,default,i32" json:"PageNo"`
	PageItemCount  int32                  `thrift:"PageItemCount,2" frugal:"2,default,i32" json:"PageItemCount"`
	TotalItemCount int32                  `thrift:"TotalItemCount,3" frugal:"3,default,i32" json:"TotalItemCount"`
	Activities     []*ListActivityAPIForm `thrift:"Activities,4" frugal:"4,default,list<ListActivityAPIForm>" json:"Activities"`
}

type ListActivityAPIResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *ListActivityAPIResp `json:"Result,omitempty"`
}

// UpdateActivityBasicConfigAPI

type RiskWarningSetting struct {
	IsRiskWarningEnable   int32  `thrift:"IsRiskWarningEnable,1" frugal:"1,default,i32" json:"IsRiskWarningEnable"`
	RiskWarningTitle      string `thrift:"RiskWarningTitle,2" frugal:"2,default,string" json:"RiskWarningTitle"`
	RiskWarningContent    string `thrift:"RiskWarningContent,3" frugal:"3,default,string" json:"RiskWarningContent"`
	RiskWarningButtonText string `thrift:"RiskWarningButtonText,4" frugal:"4,default,string" json:"RiskWarningButtonText"`
}

type BasicConfig struct {
	ActivityId                int64               `thrift:"ActivityId,1" frugal:"1,default,i64" json:"ActivityId"`
	Name                      string              `thrift:"Name,2" frugal:"2,default,string" json:"Name"`
	LiveTime                  int64               `thrift:"LiveTime,3" frugal:"3,default,i64" json:"LiveTime"`
	IsCoverImageEnable        int32               `thrift:"IsCoverImageEnable,4" frugal:"4,default,i32" json:"IsCoverImageEnable"`
	CoverImageUrl             string              `thrift:"CoverImageUrl,5" frugal:"5,default,string" json:"CoverImageUrl"`
	CoverImageUrlDefault      string              `thrift:"CoverImageUrlDefault,6" frugal:"6,default,string" json:"CoverImageUrlDefault"`
	IsPcBackImageEnable       int32               `thrift:"IsPcBackImageEnable,7" frugal:"7,default,i32" json:"IsPcBackImageEnable"`
	PcBackImageUrl            string              `thrift:"PcBackImageUrl,8" frugal:"8,default,string" json:"PcBackImageUrl"`
	PcBackImageUrlDefault     string              `thrift:"PcBackImageUrlDefault,9" frugal:"9,default,string" json:"PcBackImageUrlDefault"`
	IsMobileBackImageEnable   int32               `thrift:"IsMobileBackImageEnable,10" frugal:"10,default,i32" json:"IsMobileBackImageEnable"`
	MobileBackImageUrl        string              `thrift:"MobileBackImageUrl,11" frugal:"11,default,string" json:"MobileBackImageUrl"`
	MobileBackImageUrlDefault string              `thrift:"MobileBackImageUrlDefault,12" frugal:"12,default,string" json:"MobileBackImageUrlDefault"`
	IsPreviewVideoEnable      int32               `thrift:"IsPreviewVideoEnable,13" frugal:"13,default,i32" json:"IsPreviewVideoEnable"`
	PreviewVideoUrl           string              `thrift:"PreviewVideoUrl,14" frugal:"14,default,string" json:"PreviewVideoUrl"`
	PreviewVideoVid           string              `thrift:"PreviewVideoVid,15" frugal:"15,default,string" json:"PreviewVideoVid"`
	PreviewVideoVidDefault    string              `thrift:"PreviewVideoVidDefault,16" frugal:"16,default,string" json:"PreviewVideoVidDefault"`
	IsPeopleCountEnable       int32               `thrift:"IsPeopleCountEnable,17" frugal:"17,default,i32" json:"IsPeopleCountEnable"`
	IsHeaderImageEnable       int32               `thrift:"IsHeaderImageEnable,18" frugal:"18,default,i32" json:"IsHeaderImageEnable"`
	HeaderImageUrl            string              `thrift:"HeaderImageUrl,19" frugal:"19,default,string" json:"HeaderImageUrl"`
	IsWatermarkImageEnable    int32               `thrift:"IsWatermarkImageEnable,20" frugal:"20,default,i32" json:"IsWatermarkImageEnable"`
	WatermarkImageUrl         string              `thrift:"WatermarkImageUrl,21" frugal:"21,default,string" json:"WatermarkImageUrl"`
	IsThumbUpEnable           int32               `thrift:"IsThumbUpEnable,22" frugal:"22,default,i32" json:"IsThumbUpEnable"`
	ThumbUpUrl                string              `thrift:"ThumbUpUrl,23" frugal:"23,default,string" json:"ThumbUpUrl"`
	ThumbUpUrlDefault         string              `thrift:"ThumbUpUrlDefault,24" frugal:"24,default,string" json:"ThumbUpUrlDefault"`
	IsShareIconEnable         int32               `thrift:"IsShareIconEnable,25" frugal:"25,default,i32" json:"IsShareIconEnable"`
	ShareIconUrl              string              `thrift:"ShareIconUrl,26" frugal:"26,default,string" json:"ShareIconUrl"`
	ShareIconUrlDefault       string              `thrift:"ShareIconUrlDefault,27" frugal:"27,default,string" json:"ShareIconUrlDefault"`
	IsCommentTranslateEnable  int32               `thrift:"IsCommentTranslateEnable,28" frugal:"28,default,i32" json:"IsCommentTranslateEnable"`
	Announcement              string              `thrift:"Announcement,29" frugal:"29,default,string" json:"Announcement"`
	IsAnnouncementEnable      int32               `thrift:"IsAnnouncementEnable,30" frugal:"30,default,i32" json:"IsAnnouncementEnable"`
	BackgroundColor           string              `thrift:"BackgroundColor,31" frugal:"31,default,string" json:"BackgroundColor"`
	InteractionColor          string              `thrift:"InteractionColor,32" frugal:"32,default,string" json:"InteractionColor"`
	FontColor                 string              `thrift:"FontColor,33" frugal:"33,default,string" json:"FontColor"`
	ColorThemeIndex           string              `thrift:"ColorThemeIndex,34" frugal:"34,default,string" json:"ColorThemeIndex"`
	IsPCHeaderImageEnable     int32               `thrift:"IsPCHeaderImageEnable,35" frugal:"35,default,i32" json:"IsPCHeaderImageEnable"`
	PCHeaderImageUrl          string              `thrift:"PCHeaderImageUrl,36" frugal:"36,default,string" json:"PCHeaderImageUrl"`
	IsCountdownEnable         int32               `thrift:"IsCountdownEnable,37" frugal:"37,default,i32" json:"IsCountdownEnable"`
	IsAutoStartEnable         int32               `thrift:"IsAutoStartEnable,38" frugal:"38,default,i32" json:"IsAutoStartEnable"`
	IsPageLimitEnable         int32               `thrift:"IsPageLimitEnable,39" frugal:"39,default,i32" json:"IsPageLimitEnable"`
	PageLimitType             string              `thrift:"PageLimitType,40" frugal:"40,default,string" json:"PageLimitType"`
	IsLanguageEnable          int32               `thrift:"IsLanguageEnable,41" frugal:"41,default,i32" json:"IsLanguageEnable"`
	LanguageType              []int32             `thrift:"LanguageType,42" frugal:"42,default,list<i32>" json:"LanguageType"`
	SiteTags                  []*SiteActivityTag  `thrift:"SiteTags,43" frugal:"43,default,list<SiteActivityTag>" json:"SiteTags"`
	AutoStartType             int32               `thrift:"AutoStartType,44" frugal:"44,default,i32" json:"AutoStartType"`
	IsPlayerTopEnable         int32               `thrift:"IsPlayerTopEnable,45" frugal:"45,default,i32" json:"IsPlayerTopEnable"`
	PlayerTopType             []int64             `thrift:"PlayerTopType,46" frugal:"46,default,list<i64>" json:"PlayerTopType"`
	IsReplayAutoOnlineEnable  int32               `thrift:"IsReplayAutoOnlineEnable,47" frugal:"47,default,i32" json:"IsReplayAutoOnlineEnable"`
	PreviewVideoId            int64               `thrift:"PreviewVideoId,48" frugal:"48,default,i64" json:"PreviewVideoId"`
	AccountId                 int64               `thrift:"AccountId,49" frugal:"49,default,i64" json:"AccountId"`
	PreviewVideoReviewStatus  int32               `thrift:"PreviewVideoReviewStatus,50" frugal:"50,default,i32" json:"PreviewVideoReviewStatus"`
	DefaultSubtitleLanguage   string              `thrift:"DefaultSubtitleLanguage,51" frugal:"51,default,string" json:"DefaultSubtitleLanguage"`
	SourceSubtitleLanguage    string              `thrift:"SourceSubtitleLanguage,52" frugal:"52,default,string" json:"SourceSubtitleLanguage"`
	OpenLiveAvextractorTask   int32               `thrift:"OpenLiveAvextractorTask,53" frugal:"53,default,i32" json:"OpenLiveAvextractorTask"`
	IsTimeShift               int32               `thrift:"IsTimeShift,54" frugal:"54,default,i32" json:"IsTimeShift"`
	PreviewVideoCoverImage    string              `thrift:"PreviewVideoCoverImage,55" frugal:"55,default,string" json:"PreviewVideoCoverImage"`
	PreviewVideoMediaName     string              `thrift:"PreviewVideoMediaName,56" frugal:"56,default,string" json:"PreviewVideoMediaName"`
	IsPreviewPromptEnable     int32               `thrift:"IsPreviewPromptEnable,57" frugal:"57,default,i32" json:"IsPreviewPromptEnable"`
	PreviewPrompt             string              `thrift:"PreviewPrompt,58" frugal:"58,default,string" json:"PreviewPrompt"`
	IsReservationEnable       int32               `thrift:"IsReservationEnable,59" frugal:"59,default,i32" json:"IsReservationEnable"`
	ReservationTime           int64               `thrift:"ReservationTime,60" frugal:"60,default,i64" json:"ReservationTime"`
	ReservationText           string              `thrift:"ReservationText,61" frugal:"61,default,string" json:"ReservationText"`
	WatermarkPosition         int32               `thrift:"WatermarkPosition,62" frugal:"62,default,i32" json:"WatermarkPosition"`
	IsReplayBulletChat        int32               `thrift:"IsReplayBulletChat,63" frugal:"63,default,i32" json:"IsReplayBulletChat"`
	PresenterChatColor        string              `thrift:"PresenterChatColor,64" frugal:"64,default,string" json:"PresenterChatColor"`
	IsLiveBulletChat          int32               `thrift:"IsLiveBulletChat,65" frugal:"65,default,i32" json:"IsLiveBulletChat"`
	IsBackgroundBlur          int32               `thrift:"IsBackgroundBlur,66" frugal:"66,default,i32" json:"IsBackgroundBlur"`
	FeedbackMessage           string              `thrift:"FeedbackMessage,67" frugal:"67,default,string" json:"FeedbackMessage"`
	IsFeedbackEnable          int32               `thrift:"IsFeedbackEnable,68" frugal:"68,default,i32" json:"IsFeedbackEnable"`
	IsThumbUpNumberEnable     int32               `thrift:"IsThumbUpNumberEnable,69" frugal:"69,default,i32" json:"IsThumbUpNumberEnable"`
	SmsLanguage               int32               `thrift:"SmsLanguage,70" frugal:"70,default,i32" json:"SmsLanguage"`
	MobileChatBackgroundColor string              `thrift:"MobileChatBackgroundColor,71" frugal:"71,default,string" json:"MobileChatBackgroundColor"`
	MobileBackgroundColor     string              `thrift:"MobileBackgroundColor,72" frugal:"72,default,string" json:"MobileBackgroundColor"`
	ConfigVersion             int32               `thrift:"ConfigVersion,73" frugal:"73,default,i32" json:"ConfigVersion"`
	LiveZone                  int32               `thrift:"LiveZone,74" frugal:"74,default,i32" json:"LiveZone"`
	RiskWarningSetting        *RiskWarningSetting `thrift:"RiskWarningSetting,75" frugal:"75,default,RiskWarningSetting" json:"RiskWarningSetting"`
	IsSDKPlayEnable           int32               `thrift:"IsSDKPlayEnable,76" frugal:"76,default,i32" json:"IsSDKPlayEnable"`
	SDKPlayStatus             int32               `thrift:"SDKPlayStatus,77" frugal:"77,default,i32" json:"SDKPlayStatus"`
}

type ListActivityDetailInfoAPIResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           *ListActivityDetailInfoAPIResponseBody `json:"Result,omitempty"`
}

type ListActivityDetailInfoAPIResponseBody struct {
	PageNumber int32                 `thrift:"PageNumber,1" json:"PageNumber"`
	PageSize   int32                 `thrift:"PageSize,2" json:"PageSize"`
	TotalCount int32                 `thrift:"TotalCount,3" json:"TotalCount"`
	List       []*ActivityDetailInfo `thrift:"List,4" json:"List"`
}

type ActivityDetailInfo struct {
	ActivityId         int64               `thrift:"ActivityId,1" json:"ActivityId"`
	Basic              *Basic              `thrift:"Basic,2" json:"Basic"`
	Menus              *Menus              `thrift:"Menus,3" json:"Menus"`
	CommentConfig      *CommentConfig      `thrift:"CommentConfig,4" json:"CommentConfig"`
	Questionnaire      *Questionnaire      `thrift:"Questionnaire,5" json:"Questionnaire"`
	TTL                int32               `thrift:"TTL,6" json:"TTL"`
	InteractiveSession *InteractiveSession `thrift:"InteractiveSession,7" json:"InteractiveSession"`
	ImageTextConfig    *ImageTextConfig    `thrift:"ImageTextConfig,8" json:"ImageTextConfig"`
	ThumbUpConfig      *ThumbUpConfig      `thrift:"ThumbUpConfig,9" json:"ThumbUpConfig"`
	TaskAwardConfig    *TaskAwardConfig    `thrift:"TaskAwardConfig,10" json:"TaskAwardConfig"`
	VirtualPeopleCount int64               `thrift:"VirtualPeopleCount,11" json:"VirtualPeopleCount"`
	OnlinePeopleCount  int64               `thrift:"OnlinePeopleCount,12" json:"OnlinePeopleCount"`
	SDKToken           string              `thrift:"SDKToken,13" json:"SDKToken"`
}

type Basic struct {
	Name                      string              `thrift:"Name,1" json:"Name"`
	LiveTime                  int64               `thrift:"LiveTime,2" json:"LiveTime"`
	CoverImageUrl             string              `thrift:"CoverImageUrl,3" json:"CoverImageUrl"`
	PcBackImageUrl            string              `thrift:"PcBackImageUrl,4" json:"PcBackImageUrl"`
	MobileBackImage           string              `thrift:"MobileBackImage,5" json:"MobileBackImage"`
	Status                    int32               `thrift:"Status,6" json:"Status"`
	IsPcBackImageEnable       int32               `thrift:"IsPcBackImageEnable,7" json:"IsPcBackImageEnable"`
	IsPeopleCountEnable       int32               `thrift:"IsPeopleCountEnable,8" json:"IsPeopleCountEnable"`
	PreviewVideoUrl           string              `thrift:"PreviewVideoUrl,9" json:"PreviewVideoUrl"`
	Vid                       string              `thrift:"Vid,10" json:"Vid"`
	HeaderImageUrl            string              `thrift:"HeaderImageUrl,11" json:"HeaderImageUrl"`
	WatermarkImageUrl         string              `thrift:"WatermarkImageUrl,12" json:"WatermarkImageUrl"`
	IsThumbUpEnable           int32               `thrift:"IsThumbUpEnable,13" json:"IsThumbUpEnable"`
	ThumbUpUrl                string              `thrift:"ThumbUpUrl,14" json:"ThumbUpUrl"`
	ByteAppRoomId             string              `thrift:"ByteAppRoomId,15" json:"ByteAppRoomId"`
	ShareIconUrl              string              `thrift:"ShareIconUrl,16" json:"ShareIconUrl"`
	IsCommentTranslateEnable  int32               `thrift:"IsCommentTranslateEnable,17" json:"IsCommentTranslateEnable"`
	Announcement              string              `thrift:"Announcement,18" json:"Announcement"`
	IsAnnouncementEnable      int32               `thrift:"IsAnnouncementEnable,19" json:"IsAnnouncementEnable"`
	BackgroundColor           string              `thrift:"BackgroundColor,20" json:"BackgroundColor"`
	InteractionColor          string              `thrift:"InteractionColor,21" json:"InteractionColor"`
	IsMobileBackImageEnable   int32               `thrift:"IsMobileBackImageEnable,22" json:"IsMobileBackImageEnable"`
	FontColor                 string              `thrift:"FontColor,23" json:"FontColor"`
	ColorThemeIndex           string              `thrift:"ColorThemeIndex,24" json:"ColorThemeIndex"`
	IsPCHeaderImageEnable     int32               `thrift:"IsPCHeaderImageEnable,25" json:"IsPCHeaderImageEnable"`
	PCHeaderImageUrl          string              `thrift:"PCHeaderImageUrl,26" json:"PCHeaderImageUrl"`
	Countdown                 int64               `thrift:"Countdown,27" json:"Countdown"`
	IsCountdownEnable         int32               `thrift:"IsCountdownEnable,28" json:"IsCountdownEnable"`
	IsPageLimitEnable         int32               `thrift:"IsPageLimitEnable,29" json:"IsPageLimitEnable"`
	PageLimitType             string              `thrift:"PageLimitType,30" json:"PageLimitType"`
	IsLanguageEnable          int32               `thrift:"IsLanguageEnable,31" json:"IsLanguageEnable"`
	LanguageType              []int32             `thrift:"LanguageType,32" json:"LanguageType"`
	ReplayUrl                 string              `thrift:"ReplayUrl,33" json:"ReplayUrl"`
	IsFeedbackEnable          int32               `thrift:"IsFeedbackEnable,34" json:"IsFeedbackEnable"`
	FeedbackMessage           string              `thrift:"FeedbackMessage,35" json:"FeedbackMessage"`
	Replays                   []*ActMediaForm     `thrift:"Replays,36" json:"Replays"`
	IsHeaderImageEnable       int32               `thrift:"IsHeaderImageEnable,37" json:"IsHeaderImageEnable"`
	IsPlayerTopEnable         int32               `thrift:"IsPlayerTopEnable,38" json:"IsPlayerTopEnable"`
	PlayerTopType             []int64             `thrift:"PlayerTopType,39" json:"PlayerTopType"`
	CurrentTime               int64               `thrift:"CurrentTime,40" json:"CurrentTime"`
	LivingStartTime           int64               `thrift:"LivingStartTime,41" json:"LivingStartTime"`
	IsGenerateReplay          bool                `thrift:"IsGenerateReplay,42" json:"IsGenerateReplay"`
	IsPreviewPromptEnable     int32               `thrift:"IsPreviewPromptEnable,43" json:"IsPreviewPromptEnable"`
	PreviewPrompt             string              `thrift:"PreviewPrompt,44" json:"PreviewPrompt"`
	IsFrontierPollingEnable   bool                `thrift:"IsFrontierPollingEnable,45" json:"IsFrontierPollingEnable"`
	ChatPermission            []int32             `thrift:"ChatPermission,46" json:"ChatPermission"`
	WatchPermission           int32               `thrift:"WatchPermission,47" json:"WatchPermission"`
	AccountPermission         int32               `thrift:"AccountPermission,48" json:"AccountPermission"`
	IsReservationEnable       int32               `thrift:"IsReservationEnable,49" json:"IsReservationEnable"`
	WatermarkPosition         int32               `thrift:"WatermarkPosition,50" json:"WatermarkPosition"`
	CanShare                  int32               `thrift:"CanShare,51" json:"CanShare"`
	IsReplayBulletChat        int32               `thrift:"IsReplayBulletChat,52" json:"IsReplayBulletChat"`
	PresenterChatColor        string              `thrift:"PresenterChatColor,53" json:"PresenterChatColor"`
	IsLiveBulletChat          int32               `thrift:"IsLiveBulletChat,54" json:"IsLiveBulletChat"`
	IsBackgroundBlur          int32               `thrift:"IsBackgroundBlur,55" json:"IsBackgroundBlur"`
	IsGiftRewardEnable        int32               `thrift:"IsGiftRewardEnable,56" json:"IsGiftRewardEnable"`
	IsRankListEnable          int32               `thrift:"IsRankListEnable,57" json:"IsRankListEnable"`
	RankListName              string              `thrift:"RankListName,58" json:"RankListName"`
	RankListCount             int32               `thrift:"RankListCount,59" json:"RankListCount"`
	PropUnit                  string              `thrift:"PropUnit,60" json:"PropUnit"`
	SmsLanguage               int32               `thrift:"SmsLanguage,61" json:"SmsLanguage"`
	MobileChatBackgroundColor string              `thrift:"MobileChatBackgroundColor,62" json:"MobileChatBackgroundColor"`
	PCBackgroundColor         string              `thrift:"PCBackgroundColor,63" json:"PCBackgroundColor"`
	MobileBackgroundColor     string              `thrift:"MobileBackgroundColor,64" json:"MobileBackgroundColor"`
	ConfigVersion             int32               `thrift:"ConfigVersion,65" json:"ConfigVersion"`
	CanInvite                 int32               `thrift:"CanInvite,66" json:"CanInvite"`
	RiskWarningSetting        *RiskWarningSetting `thrift:"RiskWarningSetting,67" json:"RiskWarningSetting"`
	IsSDKPlayEnable           int32               `thrift:"IsSDKPlayEnable,68" json:"IsSDKPlayEnable"`
	SDKPlayStatus             int32               `thrift:"SDKPlayStatus,69" json:"SDKPlayStatus"`
	IsColorSync               int32               `thrift:"IsColorSync,70" json:"IsColorSync"`
}

type ActMediaForm struct {
	MediaId             int64  `thrift:"MediaId,1" json:"MediaId"`
	MediaType           int32  `thrift:"MediaType,2" json:"MediaType"`
	Name                string `thrift:"Name,3" json:"Name"`
	Url                 string `thrift:"Url,4" json:"Url"`
	SourceType          int32  `thrift:"SourceType,5" json:"SourceType"`
	SourceId            int64  `thrift:"SourceId,6" json:"SourceId"`
	Duration            int32  `thrift:"Duration,7" json:"Duration"`
	StartTime           int64  `thrift:"StartTime,8" json:"StartTime"`
	EndTime             int64  `thrift:"EndTime,9" json:"EndTime"`
	AccountId           int64  `thrift:"AccountId,10" json:"AccountId"`
	Vid                 string `thrift:"Vid,11" json:"Vid"`
	ActivityMediaType   int32  `thrift:"ActivityMediaType,12" json:"ActivityMediaType"`
	OnlineStatus        int32  `thrift:"OnlineStatus,13" json:"OnlineStatus"`
	CreateTime          int64  `thrift:"CreateTime,14" json:"CreateTime"`
	CoverImage          string `thrift:"CoverImage,15" json:"CoverImage"`
	MediaLibraryStatus  int32  `thrift:"MediaLibraryStatus,16" json:"MediaLibraryStatus"`
	ReviewStatus        int32  `thrift:"ReviewStatus,17" json:"ReviewStatus"`
	MediaSize           int64  `thrift:"MediaSize,18" json:"MediaSize"`
	MediaLibraryVideoId int64  `thrift:"MediaLibraryVideoId,19" json:"MediaLibraryVideoId"`
	PointNum            int64  `thrift:"PointNum,20" json:"PointNum"`
}

type Menus struct {
	Menus                    []*MenuMsg                        `thrift:"Menus,1" json:"Menus"`
	Band                     *GetActivityBandResp              `thrift:"Band,2" json:"Band"`
	Product                  map[string][]*ProductMsg          `thrift:"Product,3" json:"Product"`
	Market                   *UpdateActivityMarketParam        `thrift:"Market,4" json:"Market"`
	AggregationActivities    []*AggregationActivity            `thrift:"AggregationActivities,5" json:"AggregationActivities"`
	AggregationActivitiesNew map[string][]*AggregationActivity `thrift:"AggregationActivitiesNew,6" json:"AggregationActivitiesNew"`
	ProductV2                map[string]*Products              `thrift:"ProductV2,7" json:"ProductV2"`
}

type MenuMsg struct {
	Name   string `thrift:"Name,1" json:"Name"`
	TypeA1 int32  `thrift:"Type,2" json:"Type"`
	Index  int32  `thrift:"Index,3" json:"Index"`
}

type GetActivityBandResp struct {
	RichText map[string]string `thrift:"RichText,1" json:"RichText"`
}

type ProductMsg struct {
	Title                 string `thrift:"Title,1" json:"Title"`
	Highlight             string `thrift:"Highlight,2" json:"Highlight"`
	IntroduceImage        string `thrift:"IntroduceImage,3" json:"IntroduceImage"`
	RedirectImage         string `thrift:"RedirectImage,4" json:"RedirectImage"`
	RedirectUrl           string `thrift:"RedirectUrl,5" json:"RedirectUrl"`
	Index                 int32  `thrift:"Index,6" json:"Index"`
	Id                    int64  `thrift:"Id,7" json:"Id"`
	PageAdvertisementType int32  `thrift:"PageAdvertisementType,8" json:"PageAdvertisementType"`
	ExplainStatus         int32  `thrift:"ExplainStatus,9" json:"ExplainStatus"`
	ExplainTime           int64  `thrift:"ExplainTime,10" json:"ExplainTime"`
	EnableStatus          int32  `thrift:"EnableStatus,11" json:"EnableStatus"`
	FloatingStatus        int32  `thrift:"FloatingStatus,12" json:"FloatingStatus"`
	Strikethrough         string `thrift:"Strikethrough,13" json:"Strikethrough"`
	DirectUrl             string `thrift:"DirectUrl,14" json:"DirectUrl"`
}

type UpdateActivityMarketParam struct {
	ActivityId                          int64                `thrift:"ActivityId,1" json:"ActivityId"`
	IsBusinessAccountEnable             int32                `thrift:"IsBusinessAccountEnable,2" json:"IsBusinessAccountEnable"`
	BusinessAccountInfo                 *BusinessAccountInfo `thrift:"BusinessAccountInfo,3" json:"BusinessAccountInfo"`
	IsPageAdvertisementEnable           int32                `thrift:"IsPageAdvertisementEnable,4" json:"IsPageAdvertisementEnable"`
	IsPageHeaderAdvertisementEnable     int32                `thrift:"IsPageHeaderAdvertisementEnable,5" json:"IsPageHeaderAdvertisementEnable"`
	IsPageUpperAdvertisementEnable      int32                `thrift:"IsPageUpperAdvertisementEnable,6" json:"IsPageUpperAdvertisementEnable"`
	AdvertisementInfo                   []*AdvertisementInfo `thrift:"AdvertisementInfo,7" json:"AdvertisementInfo"`
	IsPageAdFloatingEnable              int32                `thrift:"IsPageAdFloatingEnable,8" json:"IsPageAdFloatingEnable"`
	IsPageHeaderAdFloatingEnable        int32                `thrift:"IsPageHeaderAdFloatingEnable,9" json:"IsPageHeaderAdFloatingEnable"`
	IsPageUpperAdFloatingEnable         int32                `thrift:"IsPageUpperAdFloatingEnable,10" json:"IsPageUpperAdFloatingEnable"`
	IsGuidePageEnable                   int32                `thrift:"IsGuidePageEnable,11" json:"IsGuidePageEnable"`
	GuidePageInfo                       *GuidePageInfo       `thrift:"GuidePageInfo,12" json:"GuidePageInfo"`
	IsPageUpperAdvertisementCloseEnable int32                `thrift:"IsPageUpperAdvertisementCloseEnable,13" json:"IsPageUpperAdvertisementCloseEnable"`
}

type BusinessAccountInfo struct {
	AccountName            string `thrift:"AccountName,1" json:"AccountName"`
	AccountHeadImage       string `thrift:"AccountHeadImage,2" json:"AccountHeadImage"`
	AccountHeadRedirectUrl string `thrift:"AccountHeadRedirectUrl,3" json:"AccountHeadRedirectUrl"`
}

type AdvertisementInfo struct {
	AdvertisementType        int32  `thrift:"AdvertisementType,1" json:"AdvertisementType"`
	AdvertisementContent     string `thrift:"AdvertisementContent,2" json:"AdvertisementContent"`
	AdvertisementRedirectUrl string `thrift:"AdvertisementRedirectUrl,3" json:"AdvertisementRedirectUrl"`
	Index                    int32  `thrift:"Index,4" json:"Index"`
	Id                       int64  `thrift:"Id,5" json:"Id"`
	PageAdvertisementType    int32  `thrift:"PageAdvertisementType,6" json:"PageAdvertisementType"`
}

type GuidePageInfo struct {
	GuidePageUrl string `thrift:"GuidePageUrl,1" json:"GuidePageUrl"`
}

type CommentConfig struct {
	IsHotListEnable          *int32  `thrift:"IsHotListEnable,1" json:"IsHotListEnable"`
	HotListName              *string `thrift:"HotListName,2" json:"HotListName"`
	HotListCount             *int32  `thrift:"HotListCount,3" json:"HotListCount"`
	VoiceInterval            *int32  `thrift:"VoiceInterval,4" json:"VoiceInterval"`
	InputBoxPrompt           *string `thrift:"InputBoxPrompt,5" json:"InputBoxPrompt"`
	PresenterName            *string `thrift:"PresenterName,6" json:"PresenterName"`
	IsBulletScreenEnable     *int32  `thrift:"IsBulletScreenEnable,7" json:"IsBulletScreenEnable"`
	IsCommentTranslateEnable *int32  `thrift:"IsCommentTranslateEnable,8" json:"IsCommentTranslateEnable"`
	IsSendCommentEnable      *int32  `thrift:"IsSendCommentEnable,9" json:"IsSendCommentEnable"`
	IsManualHotListEnable    *int32  `thrift:"IsManualHotListEnable,10" json:"IsManualHotListEnable"`
	IsLikeNumberShowEnable   *int32  `thrift:"IsLikeNumberShowEnable,11" json:"IsLikeNumberShowEnable"`
	IsWelcomeMessageEnable   *int32  `thrift:"IsWelcomeMessageEnable,12" json:"IsWelcomeMessageEnable"`
	WelcomeMessageTitle      *string `thrift:"WelcomeMessageTitle,13" json:"WelcomeMessageTitle"`
	WelcomeMessageContent    *string `thrift:"WelcomeMessageContent,14" json:"WelcomeMessageContent"`
}

type Questionnaire struct {
	QuestionnaireId             int64       `thrift:"QuestionnaireId,1" json:"QuestionnaireId"`
	Title                       string      `thrift:"Title,2" json:"Title"`
	Description                 string      `thrift:"Description,3" json:"Description"`
	TriggerTime                 string      `thrift:"TriggerTime,4" json:"TriggerTime"`
	PublishStatus               int32       `thrift:"PublishStatus,5" json:"PublishStatus"`
	CreateTime                  int64       `thrift:"CreateTime,6" json:"CreateTime"`
	Questions                   []*Question `thrift:"Questions,7" json:"Questions"`
	IsHeaderImageEnable         int32       `thrift:"IsHeaderImageEnable,8" json:"IsHeaderImageEnable"`
	HeaderImageUrl              string      `thrift:"HeaderImageUrl,9" json:"HeaderImageUrl"`
	PostCommitShowText          string      `thrift:"PostCommitShowText,10" json:"PostCommitShowText"`
	IsPostCommitShowImageEnable int32       `thrift:"IsPostCommitShowImageEnable,11" json:"IsPostCommitShowImageEnable"`
	PostCommitShowImageUrl      string      `thrift:"PostCommitShowImageUrl,12" json:"PostCommitShowImageUrl"`
	IsLarkGroupEnable           int32       `thrift:"IsLarkGroupEnable,13" json:"IsLarkGroupEnable"`
	LarkGroupType               int32       `thrift:"LarkGroupType,14" json:"LarkGroupType"`
	LarkGroupId                 string      `thrift:"LarkGroupId,15" json:"LarkGroupId"`
	LarkGroupName               string      `thrift:"LarkGroupName,16" json:"LarkGroupName"`
	LarkGroupOwner              string      `thrift:"LarkGroupOwner,17" json:"LarkGroupOwner"`
	ModifyTime                  int64       `thrift:"ModifyTime,18" json:"ModifyTime"`
	IsNotEdit                   int32       `thrift:"IsNotEdit,19" json:"IsNotEdit"`
}

type Question struct {
	QuestionId     int64     `thrift:"QuestionId,1" json:"QuestionId"`
	Question       string    `thrift:"Question,2" json:"Question"`
	IsRequire      int32     `thrift:"IsRequire,3" json:"IsRequire"`
	QuestionTag    string    `thrift:"QuestionTag,4" json:"QuestionTag"`
	QuestionSubTag string    `thrift:"QuestionSubTag,5" json:"QuestionSubTag"`
	Options        []*Option `thrift:"Options,6" json:"Options"`
}

type Option struct {
	OptionId   int64  `thrift:"OptionId,1" json:"OptionId"`
	OptionName string `thrift:"OptionName,2" json:"OptionName"`
}

type InteractiveSession struct {
	ActivityId            int64  `thrift:"ActivityId,1" json:"ActivityId"`
	VoiceInterval         int32  `thrift:"VoiceInterval,2" json:"VoiceInterval"`
	InputBoxPrompt        string `thrift:"InputBoxPrompt,3" json:"InputBoxPrompt"`
	PresenterName         string `thrift:"PresenterName,4" json:"PresenterName"`
	IsReviewSessionEnable int32  `thrift:"IsReviewSessionEnable,5" json:"IsReviewSessionEnable"`
	IsSendSessionEnable   int32  `thrift:"IsSendSessionEnable,6" json:"IsSendSessionEnable"`
	ShowType              int32  `thrift:"ShowType,7" json:"ShowType"`
	Name                  string `thrift:"Name,8" json:"Name"`
	SendNumber            int64  `thrift:"SendNumber,9" json:"SendNumber"`
}

type ImageTextConfig struct {
	ShowType      *int32 `thrift:"ShowType,1" json:"ShowType"`
	DisplayEnable *int32 `thrift:"DisplayEnable,2" json:"DisplayEnable"`
}

type ThumbUpConfig struct {
	IsThumbUpEnable       int32  `thrift:"IsThumbUpEnable,1" json:"IsThumbUpEnable"`
	IsThumbUpNumberEnable int32  `thrift:"IsThumbUpNumberEnable,2" json:"IsThumbUpNumberEnable"`
	ThumbUpUrl            string `thrift:"ThumbUpUrl,3" json:"ThumbUpUrl"`
	ShowThumbUpNumber     int64  `thrift:"ShowThumbUpNumber,4" json:"ShowThumbUpNumber"`
	ThumbUpVersion        int64  `thrift:"ThumbUpVersion,5" json:"ThumbUpVersion"`
}

type TaskAwardConfig struct {
	TaskAwardId           string                  `thrift:"TaskAwardId,1" json:"TaskAwardId"`
	TaskAwardStatus       int32                   `thrift:"TaskAwardStatus,2" json:"TaskAwardStatus"`
	TaskAwardInstructions string                  `thrift:"TaskAwardInstructions,3" json:"TaskAwardInstructions"`
	TaskAwardIcon         string                  `thrift:"TaskAwardIcon,4" json:"TaskAwardIcon"`
	TaskAwardRule         []*TaskAwardRuleContext `thrift:"TaskAwardRule,5" json:"TaskAwardRule"`
	TaskAwardStartTime    int64                   `thrift:"TaskAwardStartTime,6" json:"TaskAwardStartTime"`
	TaskAwardEndTime      int64                   `thrift:"TaskAwardEndTime,7" json:"TaskAwardEndTime"`
}

type TaskAwardRuleContext struct {
	WatchTime             int64  `thrift:"WatchTime,1" json:"WatchTime"`
	AwardItemName         string `thrift:"AwardItemName,2" json:"AwardItemName"`
	AwardItemIcon         string `thrift:"AwardItemIcon,3" json:"AwardItemIcon"`
	AwardItemType         int32  `thrift:"AwardItemType,4" json:"AwardItemType"`
	AwardCount            int32  `thrift:"AwardCount,5" json:"AwardCount"`
	TaskAwardItemId       string `thrift:"TaskAwardItemId,6" json:"TaskAwardItemId"`
	AwardParticipantLimit int32  `thrift:"AwardParticipantLimit,7" json:"AwardParticipantLimit"`
}

type AggregationActivity struct {
	ActivityId int64  `thrift:"ActivityId,1" json:"ActivityId"`
	Name       string `thrift:"Name,2" json:"Name"`
	PlayUrl    string `thrift:"PlayUrl,3" json:"PlayUrl"`
	CoverImage string `thrift:"CoverImage,4" json:"CoverImage"`
	IsDisplay  int32  `thrift:"IsDisplay,5" json:"IsDisplay"`
	Status     int32  `thrift:"Status,6" json:"Status"`
	LiveTime   int64  `thrift:"LiveTime,7" json:"LiveTime"`
	Index      int32  `thrift:"Index,8" json:"Index"`
}

type Products struct {
	EnableFloating bool          `thrift:"EnableFloating,1" json:"EnableFloating"`
	Product        []*ProductMsg `thrift:"Product,2" json:"Product"`
	EnableUA       bool          `thrift:"EnableUA,3" json:"EnableUA"`
	UAAddress      string        `thrift:"UAAddress,4" json:"UAAddress"`
}
