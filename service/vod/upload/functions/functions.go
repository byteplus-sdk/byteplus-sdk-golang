package functions

import (
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/vod/models/business"
)

func GetMetaFunc() business.VodUploadFunction {
	return business.VodUploadFunction{
		Name: "GetMeta",
	}
}

func SnapshotFunc(snapshot business.VodUploadFunctionInput) business.VodUploadFunction {
	return business.VodUploadFunction{
		Name: "Snapshot",
		Input: &business.VodUploadFunctionInput{
			SnapshotTime: snapshot.SnapshotTime,
		},
	}
}

func AddOptionInfoFunc(info business.VodUploadFunctionInput) business.VodUploadFunction {
	return business.VodUploadFunction{
		Name: "AddOptionInfo",
		Input: &business.VodUploadFunctionInput{
			Title:            info.Title,
			Tags:             info.Tags,
			Description:      info.Description,
			Category:         info.Category,
			RecordType:       info.RecordType,
			Format:           info.Format,
			ClassificationId: info.ClassificationId,
		},
	}
}

func StartWorkflowFunc(workflowInput business.VodUploadFunctionInput) business.VodUploadFunction {
	return business.VodUploadFunction{
		Name:  "StartWorkflow",
		Input: &business.VodUploadFunctionInput{TemplateId: workflowInput.TemplateId, TemplateIds: workflowInput.TemplateIds},
	}
}

func CaptionUploadFunc(captionUploadInput business.VodUploadFunctionInput) business.VodUploadFunction {
	return business.VodUploadFunction{
		Name: "CaptionUpload",
		Input: &business.VodUploadFunctionInput{
			Title:       captionUploadInput.Title,
			Format:      captionUploadInput.Format,
			Vid:         captionUploadInput.Fid,
			Fid:         captionUploadInput.Fid,
			Language:    captionUploadInput.Language,
			StoreUri:    captionUploadInput.StoreUri,
			Source:      captionUploadInput.Source,
			Tag:         captionUploadInput.Tag,
			AutoPublish: captionUploadInput.AutoPublish,
			ActionType:  captionUploadInput.ActionType,
		},
	}
}
