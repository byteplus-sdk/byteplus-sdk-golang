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
