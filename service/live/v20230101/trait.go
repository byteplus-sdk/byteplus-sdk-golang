package live_v20230101

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"

	common "github.com/byteplus-sdk/byteplus-sdk-golang/base"
)

type queryMarshalFilter func(key string, value []string) (accept bool)

func skipEmptyValue() queryMarshalFilter {
	return func(_ string, value []string) (accept bool) {
		if len(value) == 0 {
			return false
		}

		for _, item := range value {
			if item == "" {
				return false
			}
		}

		return true
	}
}

func marshalToQuery(model interface{}, filters ...queryMarshalFilter) (url.Values, error) {
	ret := url.Values{}
	if model == nil {
		return ret, nil
	}

	modelType := reflect.TypeOf(model)
	modelValue := reflect.ValueOf(model)
	if modelValue.IsNil() {
		return ret, nil
	}

	if modelType.Kind() == reflect.Ptr {
		modelValue = modelValue.Elem()
		modelType = modelValue.Type()
	} else {
		return ret, fmt.Errorf("not struct")
	}
	fieldCount := modelType.NumField()

	for i := 0; i < fieldCount; i++ {
		field := modelType.Field(i)
		queryKey := field.Tag.Get("query")
		if queryKey == "" {
			continue
		}

		queryRawValue := modelValue.FieldByName(field.Name)
		queryValues := make([]string, 0)
		if field.Type.Kind() != reflect.Array && field.Type.Kind() != reflect.Slice {
			value := resolveQueryValue(queryRawValue)
			if value == nil {
				continue
			}
			queryValues = append(queryValues, fmt.Sprintf("%v", value))
		} else {
			length := queryRawValue.Len()
			for idx := 0; idx < length; idx++ {
				value := resolveQueryValue(queryRawValue.Index(idx))
				if value == nil {
					continue
				}
				queryValues = append(queryValues, fmt.Sprintf("%v", value))
			}
		}

		for _, fun := range filters {
			if !fun(queryKey, queryValues) {
				goto afterAddQuery
			}
		}

		for _, t := range queryValues {
			ret.Add(queryKey, t)
		}
	afterAddQuery:
	}

	return ret, nil
}

func resolveQueryValue(queryRawValue reflect.Value) interface{} {
	if queryRawValue.Kind() == reflect.Ptr {
		if queryRawValue.IsNil() {
			return nil
		}
		decayedQueryRawValue := queryRawValue.Elem()
		decayedReflectValue := decayedQueryRawValue.Interface()
		return fmt.Sprintf("%v", decayedReflectValue)
	} else {
		queryReflectValue := queryRawValue.Interface()
		return fmt.Sprintf("%v", queryReflectValue)
	}
}

func marshalToJson(model interface{}) ([]byte, error) {
	if model == nil {
		return make([]byte, 0), nil
	}

	result, err := json.Marshal(model)
	if err != nil {
		return []byte{}, fmt.Errorf("can not marshal model to json, %v", err)
	}
	return result, nil
}

func unmarshalResultInto(data []byte, result interface{}) error {
	resp := new(common.CommonResponse)
	if err := json.Unmarshal(data, resp); err != nil {
		return fmt.Errorf("fail to unmarshal response, %v", err)
	}
	errObj := resp.ResponseMetadata.Error
	if errObj != nil && errObj.CodeN != 0 {
		return fmt.Errorf("request %s error %s", resp.ResponseMetadata.RequestId, errObj.Message)
	}

	if err := json.Unmarshal(data, result); err != nil {
		return fmt.Errorf("fail to unmarshal result, %v", err)
	}
	return nil
}

func (c *Live) DeleteTranscodePreset(ctx context.Context, arg *DeleteTranscodePresetBody) (*DeleteTranscodePresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteTranscodePreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteTranscodePresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteCommonTransPreset(ctx context.Context, arg *DeleteCommonTransPresetBody) (*DeleteCommonTransPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteCommonTransPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteCommonTransPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteTranscodePresetPatchByAdmin(ctx context.Context, arg *DeleteTranscodePresetPatchByAdminBody) (*DeleteTranscodePresetPatchByAdminRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteTranscodePresetPatchByAdmin", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteTranscodePresetPatchByAdminRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) AddCommonTransPreset(ctx context.Context, arg *AddCommonTransPresetBody) (*AddCommonTransPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "AddCommonTransPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(AddCommonTransPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateTranscodePreset(ctx context.Context, arg *UpdateTranscodePresetBody) (*UpdateTranscodePresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateTranscodePreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateTranscodePresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListCommonTransPresetDetail(ctx context.Context, arg *ListCommonTransPresetDetailBody) (*ListCommonTransPresetDetailRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListCommonTransPresetDetail", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListCommonTransPresetDetailRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeTranscodePresetDetail(ctx context.Context, arg *DescribeTranscodePresetDetailBody) (*DescribeTranscodePresetDetailRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeTranscodePresetDetail", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeTranscodePresetDetailRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListVhostTransCodePreset(ctx context.Context, arg *ListVhostTransCodePresetBody) (*ListVhostTransCodePresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListVhostTransCodePreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListVhostTransCodePresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateTranscodePreset(ctx context.Context, arg *CreateTranscodePresetBody) (*CreateTranscodePresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateTranscodePreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateTranscodePresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateTranscodePresetPatchByAdmin(ctx context.Context, arg *CreateTranscodePresetPatchByAdminBody) (*CreateTranscodePresetPatchByAdminRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateTranscodePresetPatchByAdmin", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateTranscodePresetPatchByAdminRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteWatermarkPresetV2(ctx context.Context, arg *DeleteWatermarkPresetV2Body) (*DeleteWatermarkPresetV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteWatermarkPresetV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteWatermarkPresetV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateWatermarkPresetV2(ctx context.Context, arg *UpdateWatermarkPresetV2Body) (*UpdateWatermarkPresetV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateWatermarkPresetV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateWatermarkPresetV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeWatermarkPresetDetail(ctx context.Context, arg *DescribeWatermarkPresetDetailBody) (*DescribeWatermarkPresetDetailRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeWatermarkPresetDetail", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeWatermarkPresetDetailRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateWatermarkPreset(ctx context.Context, arg *CreateWatermarkPresetBody) (*CreateWatermarkPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateWatermarkPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateWatermarkPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateWatermarkPresetV2(ctx context.Context, arg *CreateWatermarkPresetV2Body) (*CreateWatermarkPresetV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateWatermarkPresetV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateWatermarkPresetV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateWatermarkPreset(ctx context.Context, arg *UpdateWatermarkPresetBody) (*UpdateWatermarkPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateWatermarkPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateWatermarkPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteWatermarkPreset(ctx context.Context, arg *DeleteWatermarkPresetBody) (*DeleteWatermarkPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteWatermarkPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteWatermarkPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListWatermarkPreset(ctx context.Context, arg *ListWatermarkPresetBody) (*ListWatermarkPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListWatermarkPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListWatermarkPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListVhostWatermarkPreset(ctx context.Context, arg *ListVhostWatermarkPresetBody) (*ListVhostWatermarkPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListVhostWatermarkPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListVhostWatermarkPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) StopPullRecordTask(ctx context.Context, arg *StopPullRecordTaskBody) (*StopPullRecordTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "StopPullRecordTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(StopPullRecordTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreatePullRecordTask(ctx context.Context, arg *CreatePullRecordTaskBody) (*CreatePullRecordTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreatePullRecordTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreatePullRecordTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteRecordHistory(ctx context.Context, arg *DeleteRecordHistoryBody) (*DeleteRecordHistoryRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteRecordHistory", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteRecordHistoryRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteRecordPreset(ctx context.Context, arg *DeleteRecordPresetBody) (*DeleteRecordPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteRecordPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteRecordPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateRecordPresetV2(ctx context.Context, arg *UpdateRecordPresetV2Body) (*UpdateRecordPresetV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateRecordPresetV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateRecordPresetV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeRecordTaskFileHistory(ctx context.Context, arg *DescribeRecordTaskFileHistoryBody) (*DescribeRecordTaskFileHistoryRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeRecordTaskFileHistory", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeRecordTaskFileHistoryRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListVhostRecordPresetV2(ctx context.Context, arg *ListVhostRecordPresetV2Body) (*ListVhostRecordPresetV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListVhostRecordPresetV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListVhostRecordPresetV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListPullRecordTask(ctx context.Context, arg *ListPullRecordTaskBody) (*ListPullRecordTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListPullRecordTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListPullRecordTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateRecordPresetV2(ctx context.Context, arg *CreateRecordPresetV2Body) (*CreateRecordPresetV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateRecordPresetV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateRecordPresetV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListVideoClassifications(ctx context.Context, arg *ListVideoClassificationsBody) (*ListVideoClassificationsRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListVideoClassifications", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListVideoClassificationsRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteSnapshotPreset(ctx context.Context, arg *DeleteSnapshotPresetBody) (*DeleteSnapshotPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteSnapshotPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteSnapshotPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateSnapshotPreset(ctx context.Context, arg *UpdateSnapshotPresetBody) (*UpdateSnapshotPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateSnapshotPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateSnapshotPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateSnapshotPresetV2(ctx context.Context, arg *UpdateSnapshotPresetV2Body) (*UpdateSnapshotPresetV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateSnapshotPresetV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateSnapshotPresetV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeCDNSnapshotHistory(ctx context.Context, arg *DescribeCDNSnapshotHistoryBody) (*DescribeCDNSnapshotHistoryRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeCDNSnapshotHistory", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeCDNSnapshotHistoryRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListVhostSnapshotPreset(ctx context.Context, arg *ListVhostSnapshotPresetBody) (*ListVhostSnapshotPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListVhostSnapshotPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListVhostSnapshotPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListVhostSnapshotPresetV2(ctx context.Context, arg *ListVhostSnapshotPresetV2Body) (*ListVhostSnapshotPresetV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListVhostSnapshotPresetV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListVhostSnapshotPresetV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateSnapshotPreset(ctx context.Context, arg *CreateSnapshotPresetBody) (*CreateSnapshotPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateSnapshotPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateSnapshotPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateSnapshotPresetV2(ctx context.Context, arg *CreateSnapshotPresetV2Body) (*CreateSnapshotPresetV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateSnapshotPresetV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateSnapshotPresetV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteTimeShiftPresetV2(ctx context.Context, arg *DeleteTimeShiftPresetV2Body) (*DeleteTimeShiftPresetV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteTimeShiftPresetV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteTimeShiftPresetV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteTimeShiftPresetV3(ctx context.Context, arg *DeleteTimeShiftPresetV3Body) (*DeleteTimeShiftPresetV3Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteTimeShiftPresetV3", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteTimeShiftPresetV3Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateTimeShiftPresetV2(ctx context.Context, arg *CreateTimeShiftPresetV2Body) (*CreateTimeShiftPresetV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateTimeShiftPresetV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateTimeShiftPresetV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateTimeShiftPresetV2(ctx context.Context, arg *UpdateTimeShiftPresetV2Body) (*UpdateTimeShiftPresetV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateTimeShiftPresetV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateTimeShiftPresetV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateTimeShiftPresetV3(ctx context.Context, arg *UpdateTimeShiftPresetV3Body) (*UpdateTimeShiftPresetV3Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateTimeShiftPresetV3", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateTimeShiftPresetV3Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeTimeShiftPresetDetail(ctx context.Context, arg *DescribeTimeShiftPresetDetailBody) (*DescribeTimeShiftPresetDetailRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeTimeShiftPresetDetail", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeTimeShiftPresetDetailRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListTimeShiftPresetV2(ctx context.Context, arg *ListTimeShiftPresetV2Body) (*ListTimeShiftPresetV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListTimeShiftPresetV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListTimeShiftPresetV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateTimeShiftPresetV3(ctx context.Context, arg *CreateTimeShiftPresetV3Body) (*CreateTimeShiftPresetV3Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateTimeShiftPresetV3", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateTimeShiftPresetV3Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) GenerateTimeShiftPlayURL(ctx context.Context, arg *GenerateTimeShiftPlayURLBody) (*GenerateTimeShiftPlayURLRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GenerateTimeShiftPlayURL", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GenerateTimeShiftPlayURLRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListVhostDomainDetailByUserID(ctx context.Context, arg *ListVhostDomainDetailByUserIDBody) (*ListVhostDomainDetailByUserIDRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListVhostDomainDetailByUserID", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListVhostDomainDetailByUserIDRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateVhostTags(ctx context.Context, arg *UpdateVhostTagsBody) (*UpdateVhostTagsRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateVhostTags", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateVhostTagsRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListVhostDetail(ctx context.Context, arg *ListVhostDetailBody) (*ListVhostDetailRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListVhostDetail", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListVhostDetailRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListVhostDetailByAdmin(ctx context.Context, arg *ListVhostDetailByAdminBody) (*ListVhostDetailByAdminRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListVhostDetailByAdmin", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListVhostDetailByAdminRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeVhost(ctx context.Context, arg *DescribeVhostBody) (*DescribeVhostRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeVhost", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeVhostRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) GetTags(ctx context.Context) (*GetTagsRes, error) {

	data, _, err := c.CtxQuery(ctx, "GetTags", url.Values{})
	if err != nil {
		return nil, err
	}

	result := new(GetTagsRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListProjects(ctx context.Context, arg *ListProjectsBody) (*ListProjectsRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListProjects", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListProjectsRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteCallback(ctx context.Context, arg *DeleteCallbackBody) (*DeleteCallbackRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteCallback", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteCallbackRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeCallback(ctx context.Context, arg *DescribeCallbackBody) (*DescribeCallbackRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeCallback", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeCallbackRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateCallback(ctx context.Context, arg *UpdateCallbackBody) (*UpdateCallbackRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateCallback", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateCallbackRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteAuth(ctx context.Context, arg *DeleteAuthBody) (*DeleteAuthRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteAuth", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteAuthRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) EnableAuth(ctx context.Context, arg *EnableAuthBody) (*EnableAuthRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "EnableAuth", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(EnableAuthRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeAuth(ctx context.Context, arg *DescribeAuthBody) (*DescribeAuthRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeAuth", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeAuthRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DisableAuth(ctx context.Context, arg *DisableAuthBody) (*DisableAuthRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DisableAuth", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DisableAuthRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteCert(ctx context.Context, arg *DeleteCertBody) (*DeleteCertRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteCert", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteCertRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateCert(ctx context.Context, arg *UpdateCertBody) (*UpdateCertRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateCert", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateCertRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeCertDetailSecretV2(ctx context.Context, arg *DescribeCertDetailSecretV2Body) (*DescribeCertDetailSecretV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeCertDetailSecretV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeCertDetailSecretV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListCertV2(ctx context.Context, arg *ListCertV2Body) (*ListCertV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListCertV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListCertV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeCertDetailV2(ctx context.Context, arg *DescribeCertDetailV2Body) (*DescribeCertDetailV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeCertDetailV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeCertDetailV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateCert(ctx context.Context, arg *CreateCertBody) (*CreateCertRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateCert", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateCertRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) BindCert(ctx context.Context, arg *BindCertBody) (*BindCertRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "BindCert", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(BindCertRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UnbindCert(ctx context.Context, arg *UnbindCertBody) (*UnbindCertRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UnbindCert", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UnbindCertRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListObject(ctx context.Context, arg *ListObjectBody) (*ListObjectRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListObject", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListObjectRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ManagerPullPushDomainBind(ctx context.Context, arg *ManagerPullPushDomainBindBody) (*ManagerPullPushDomainBindRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ManagerPullPushDomainBind", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ManagerPullPushDomainBindRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeCertDetailSecret(ctx context.Context, arg *DescribeCertDetailSecretBody) (*DescribeCertDetailSecretRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeCertDetailSecret", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeCertDetailSecretRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeDomainVerify(ctx context.Context, arg *DescribeDomainVerifyBody) (*DescribeDomainVerifyRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeDomainVerify", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeDomainVerifyRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateVerifyContent(ctx context.Context, arg *CreateVerifyContentBody) (*CreateVerifyContentRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateVerifyContent", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateVerifyContentRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListVqosDimensionValues(ctx context.Context, arg *ListVqosDimensionValuesReq) (*ListVqosDimensionValuesRes, error) {
	query, err := marshalToQuery(arg.ListVqosDimensionValuesQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.ListVqosDimensionValuesBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListVqosDimensionValues", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListVqosDimensionValuesRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListCert(ctx context.Context, arg *ListCertBody) (*ListCertRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListCert", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListCertRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListCertBindInfo(ctx context.Context, arg *ListCertBindInfoBody) (*ListCertBindInfoRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListCertBindInfo", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListCertBindInfoRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveFreeTimeInterval(ctx context.Context) (*DescribeLiveFreeTimeIntervalRes, error) {

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveFreeTimeInterval", url.Values{}, "")
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveFreeTimeIntervalRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) VerifyDomainOwner(ctx context.Context, arg *VerifyDomainOwnerBody) (*VerifyDomainOwnerRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "VerifyDomainOwner", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(VerifyDomainOwnerRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ValidateCert(ctx context.Context, arg *ValidateCertBody) (*ValidateCertRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ValidateCert", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ValidateCertRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteDomain(ctx context.Context, arg *DeleteDomainBody) (*DeleteDomainRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteDomain", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteDomainRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteDomainV2(ctx context.Context, arg *DeleteDomainV2Body) (*DeleteDomainV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteDomainV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteDomainV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) EnableDomain(ctx context.Context, arg *EnableDomainBody) (*EnableDomainRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "EnableDomain", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(EnableDomainRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateDomainV2(ctx context.Context, arg *CreateDomainV2Body) (*CreateDomainV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateDomainV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateDomainV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) RejectDomain(ctx context.Context, arg *RejectDomainBody) (*RejectDomainRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "RejectDomain", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(RejectDomainRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateDomainVhost(ctx context.Context, arg *UpdateDomainVhostBody) (*UpdateDomainVhostRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateDomainVhost", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateDomainVhostRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateDomain(ctx context.Context, arg *UpdateDomainBody) (*UpdateDomainRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateDomain", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateDomainRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeDomain(ctx context.Context, arg *DescribeDomainBody) (*DescribeDomainRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeDomain", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeDomainRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListDomainDetail(ctx context.Context, arg *ListDomainDetailBody) (*ListDomainDetailRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListDomainDetail", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListDomainDetailRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateDomain(ctx context.Context, arg *CreateDomainBody) (*CreateDomainRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateDomain", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateDomainRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DisableDomain(ctx context.Context, arg *DisableDomainBody) (*DisableDomainRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DisableDomain", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DisableDomainRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateVQScoreTask(ctx context.Context, arg *CreateVQScoreTaskBody) (*CreateVQScoreTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateVQScoreTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateVQScoreTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeVQScoreTask(ctx context.Context, arg *DescribeVQScoreTaskBody) (*DescribeVQScoreTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeVQScoreTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeVQScoreTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListVQScoreTask(ctx context.Context, arg *ListVQScoreTaskBody) (*ListVQScoreTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListVQScoreTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListVQScoreTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) StopPullToPushTask(ctx context.Context, arg *StopPullToPushTaskBody) (*StopPullToPushTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "StopPullToPushTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(StopPullToPushTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreatePullToPushTask(ctx context.Context, arg *CreatePullToPushTaskBody) (*CreatePullToPushTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreatePullToPushTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreatePullToPushTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeletePullToPushTask(ctx context.Context, arg *DeletePullToPushTaskBody) (*DeletePullToPushTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeletePullToPushTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeletePullToPushTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) RestartPullToPushTask(ctx context.Context, arg *RestartPullToPushTaskBody) (*RestartPullToPushTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "RestartPullToPushTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(RestartPullToPushTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdatePullToPushTask(ctx context.Context, arg *UpdatePullToPushTaskBody) (*UpdatePullToPushTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdatePullToPushTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdatePullToPushTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListPullToPushTask(ctx context.Context, arg *ListPullToPushTaskQuery) (*ListPullToPushTaskRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ListPullToPushTask", query)
	if err != nil {
		return nil, err
	}

	result := new(ListPullToPushTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteDenyConfigV2(ctx context.Context, arg *DeleteDenyConfigV2Body) (*DeleteDenyConfigV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteDenyConfigV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteDenyConfigV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeDenyConfigV2(ctx context.Context, arg *DescribeDenyConfigV2Body) (*DescribeDenyConfigV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeDenyConfigV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeDenyConfigV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteRelaySourceRewrite(ctx context.Context, arg *DeleteRelaySourceRewriteBody) (*DeleteRelaySourceRewriteRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteRelaySourceRewrite", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteRelaySourceRewriteRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteRelaySourceV4(ctx context.Context, arg *DeleteRelaySourceV4Body) (*DeleteRelaySourceV4Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteRelaySourceV4", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteRelaySourceV4Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteRelaySourceV3(ctx context.Context, arg *DeleteRelaySourceV3Body) (*DeleteRelaySourceV3Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteRelaySourceV3", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteRelaySourceV3Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateRelaySourceRewrite(ctx context.Context, arg *UpdateRelaySourceRewriteBody) (*UpdateRelaySourceRewriteRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateRelaySourceRewrite", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateRelaySourceRewriteRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateRelaySourceV4(ctx context.Context, arg *UpdateRelaySourceV4Body) (*UpdateRelaySourceV4Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateRelaySourceV4", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateRelaySourceV4Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeRelaySourceRewrite(ctx context.Context, arg *DescribeRelaySourceRewriteBody) (*DescribeRelaySourceRewriteRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeRelaySourceRewrite", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeRelaySourceRewriteRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListRelaySourceV4(ctx context.Context, arg *ListRelaySourceV4Body) (*ListRelaySourceV4Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListRelaySourceV4", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListRelaySourceV4Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeRelaySourceV3(ctx context.Context, arg *DescribeRelaySourceV3Body) (*DescribeRelaySourceV3Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeRelaySourceV3", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeRelaySourceV3Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateRelaySourceV4(ctx context.Context, arg *CreateRelaySourceV4Body) (*CreateRelaySourceV4Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateRelaySourceV4", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateRelaySourceV4Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateRelaySourceV3(ctx context.Context, arg *UpdateRelaySourceV3Body) (*UpdateRelaySourceV3Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateRelaySourceV3", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateRelaySourceV3Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) KillStream(ctx context.Context, arg *KillStreamBody) (*KillStreamRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "KillStream", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(KillStreamRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeClosedStreamInfoByPage(ctx context.Context, arg *DescribeClosedStreamInfoByPageQuery) (*DescribeClosedStreamInfoByPageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeClosedStreamInfoByPage", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeClosedStreamInfoByPageRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveStreamInfoByPage(ctx context.Context, arg *DescribeLiveStreamInfoByPageQuery) (*DescribeLiveStreamInfoByPageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeLiveStreamInfoByPage", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveStreamInfoByPageRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveStreamState(ctx context.Context, arg *DescribeLiveStreamStateQuery) (*DescribeLiveStreamStateRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeLiveStreamState", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveStreamStateRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeForbiddenStreamInfoByPage(ctx context.Context, arg *DescribeForbiddenStreamInfoByPageQuery) (*DescribeForbiddenStreamInfoByPageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeForbiddenStreamInfoByPage", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeForbiddenStreamInfoByPageRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ForbidStream(ctx context.Context, arg *ForbidStreamBody) (*ForbidStreamRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ForbidStream", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ForbidStreamRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ResumeStream(ctx context.Context, arg *ResumeStreamBody) (*ResumeStreamRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ResumeStream", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ResumeStreamRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) GeneratePlayURL(ctx context.Context, arg *GeneratePlayURLBody) (*GeneratePlayURLRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GeneratePlayURL", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GeneratePlayURLRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) GeneratePushURL(ctx context.Context, arg *GeneratePushURLBody) (*GeneratePushURLRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GeneratePushURL", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GeneratePushURLRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateSDKLicense(ctx context.Context, arg *UpdateSDKLicenseBody) (*UpdateSDKLicenseRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateSDKLicense", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateSDKLicenseRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateApp(ctx context.Context, arg *CreateAppBody) (*CreateAppRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateApp", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateAppRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteSDK(ctx context.Context, arg *DeleteSDKBody) (*DeleteSDKRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteSDK", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteSDKRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateApp(ctx context.Context, arg *UpdateAppBody) (*UpdateAppRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateApp", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateAppRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateSDK(ctx context.Context, arg *UpdateSDKBody) (*UpdateSDKRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateSDK", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateSDKRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeSDKDetail(ctx context.Context, arg *DescribeSDKDetailBody) (*DescribeSDKDetailRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeSDKDetail", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeSDKDetailRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeSDKParamsAvailable(ctx context.Context, arg *DescribeSDKParamsAvailableBody) (*DescribeSDKParamsAvailableRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeSDKParamsAvailable", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeSDKParamsAvailableRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeAppIDParamsAvailable(ctx context.Context, arg *DescribeAppIDParamsAvailableBody) (*DescribeAppIDParamsAvailableRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeAppIDParamsAvailable", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeAppIDParamsAvailableRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateSDK(ctx context.Context, arg *CreateSDKBody) (*CreateSDKRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateSDK", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateSDKRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListSDK(ctx context.Context, arg *ListSDKBody) (*ListSDKRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListSDK", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListSDKRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListSDKAdmin(ctx context.Context, arg *ListSDKAdminBody) (*ListSDKAdminRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListSDKAdmin", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListSDKAdminRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) GetApps(ctx context.Context) (*GetAppsRes, error) {

	data, _, err := c.CtxQuery(ctx, "GetApps", url.Values{})
	if err != nil {
		return nil, err
	}

	result := new(GetAppsRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateService(ctx context.Context, arg *UpdateServiceBody) (*UpdateServiceRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateService", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateServiceRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListServices(ctx context.Context, arg *ListServicesBody) (*ListServicesRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListServices", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListServicesRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeService(ctx context.Context) (*DescribeServiceRes, error) {

	data, _, err := c.Client.CtxJson(ctx, "DescribeService", url.Values{}, "")
	if err != nil {
		return nil, err
	}

	result := new(DescribeServiceRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateActivityBilling(ctx context.Context, arg *UpdateActivityBillingBody) (*UpdateActivityBillingRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateActivityBilling", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateActivityBillingRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateBilling(ctx context.Context, arg *UpdateBillingBody) (*UpdateBillingRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateBilling", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateBillingRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListInstance(ctx context.Context, arg *ListInstanceBody) (*ListInstanceRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListInstance", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListInstanceRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeBillingForAdmin(ctx context.Context, arg *DescribeBillingForAdminBody) (*DescribeBillingForAdminRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeBillingForAdmin", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeBillingForAdminRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeBilling(ctx context.Context) (*DescribeBillingRes, error) {

	data, _, err := c.Client.CtxJson(ctx, "DescribeBilling", url.Values{}, "")
	if err != nil {
		return nil, err
	}

	result := new(DescribeBillingRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeBillingMonthAvailable(ctx context.Context) (*DescribeBillingMonthAvailableRes, error) {

	data, _, err := c.Client.CtxJson(ctx, "DescribeBillingMonthAvailable", url.Values{}, "")
	if err != nil {
		return nil, err
	}

	result := new(DescribeBillingMonthAvailableRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListResourcePackage(ctx context.Context, arg *ListResourcePackageBody) (*ListResourcePackageRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListResourcePackage", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListResourcePackageRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) TerminateInstance(ctx context.Context, arg *TerminateInstanceBody) (*TerminateInstanceRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "TerminateInstance", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(TerminateInstanceRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteStreamQuotaConfig(ctx context.Context, arg *DeleteStreamQuotaConfigBody) (*DeleteStreamQuotaConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteStreamQuotaConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteStreamQuotaConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateStreamQuotaConfigPatch(ctx context.Context, arg *UpdateStreamQuotaConfigPatchBody) (*UpdateStreamQuotaConfigPatchRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateStreamQuotaConfigPatch", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateStreamQuotaConfigPatchRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeStreamQuotaConfig(ctx context.Context, arg *DescribeStreamQuotaConfigBody) (*DescribeStreamQuotaConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeStreamQuotaConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeStreamQuotaConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateStreamQuotaConfig(ctx context.Context, arg *UpdateStreamQuotaConfigBody) (*UpdateStreamQuotaConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateStreamQuotaConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateStreamQuotaConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListVqosMetricsDimensions(ctx context.Context, arg *ListVqosMetricsDimensionsQuery) (*ListVqosMetricsDimensionsRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "ListVqosMetricsDimensions", query)
	if err != nil {
		return nil, err
	}

	result := new(ListVqosMetricsDimensionsRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) GetVqosRawData(ctx context.Context, arg *GetVqosRawDataReq) (*GetVqosRawDataRes, error) {
	query, err := marshalToQuery(arg.GetVqosRawDataQuery)
	if err != nil {
		return nil, err
	}

	body, err := marshalToJson(arg.GetVqosRawDataBody)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetVqosRawData", query, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetVqosRawDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) StopPullCDNSnapshotTask(ctx context.Context, arg *StopPullCDNSnapshotTaskBody) (*StopPullCDNSnapshotTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "StopPullCDNSnapshotTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(StopPullCDNSnapshotTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreatePullCDNSnapshotTask(ctx context.Context, arg *CreatePullCDNSnapshotTaskBody) (*CreatePullCDNSnapshotTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreatePullCDNSnapshotTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreatePullCDNSnapshotTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) GetPullCDNSnapshotTask(ctx context.Context, arg *GetPullCDNSnapshotTaskBody) (*GetPullCDNSnapshotTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetPullCDNSnapshotTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetPullCDNSnapshotTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListPullCDNSnapshotTask(ctx context.Context, arg *ListPullCDNSnapshotTaskBody) (*ListPullCDNSnapshotTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListPullCDNSnapshotTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListPullCDNSnapshotTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) GetPullRecordTask(ctx context.Context, arg *GetPullRecordTaskBody) (*GetPullRecordTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "GetPullRecordTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetPullRecordTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteSnapshotAuditPreset(ctx context.Context, arg *DeleteSnapshotAuditPresetBody) (*DeleteSnapshotAuditPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteSnapshotAuditPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteSnapshotAuditPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateSnapshotAuditPreset(ctx context.Context, arg *UpdateSnapshotAuditPresetBody) (*UpdateSnapshotAuditPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateSnapshotAuditPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateSnapshotAuditPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeSnapshotAuditPresetDetail(ctx context.Context, arg *DescribeSnapshotAuditPresetDetailBody) (*DescribeSnapshotAuditPresetDetailRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeSnapshotAuditPresetDetail", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeSnapshotAuditPresetDetailRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListVhostSnapshotAuditPreset(ctx context.Context, arg *ListVhostSnapshotAuditPresetBody) (*ListVhostSnapshotAuditPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListVhostSnapshotAuditPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListVhostSnapshotAuditPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateSnapshotAuditPreset(ctx context.Context, arg *CreateSnapshotAuditPresetBody) (*CreateSnapshotAuditPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateSnapshotAuditPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateSnapshotAuditPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeIPInfo(ctx context.Context, arg *DescribeIPInfoBody) (*DescribeIPInfoRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeIpInfo", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeIPInfoRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveRegionData(ctx context.Context) (*DescribeLiveRegionDataRes, error) {

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveRegionData", url.Values{}, "")
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveRegionDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveSourceStreamMetrics(ctx context.Context, arg *DescribeLiveSourceStreamMetricsBody) (*DescribeLiveSourceStreamMetricsRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveSourceStreamMetrics", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveSourceStreamMetricsRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLivePushStreamMetrics(ctx context.Context, arg *DescribeLivePushStreamMetricsBody) (*DescribeLivePushStreamMetricsRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLivePushStreamMetrics", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLivePushStreamMetricsRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLivePlayStatusCodeData(ctx context.Context, arg *DescribeLivePlayStatusCodeDataBody) (*DescribeLivePlayStatusCodeDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLivePlayStatusCodeData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLivePlayStatusCodeDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveBatchSourceStreamMetrics(ctx context.Context, arg *DescribeLiveBatchSourceStreamMetricsBody) (*DescribeLiveBatchSourceStreamMetricsRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveBatchSourceStreamMetrics", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveBatchSourceStreamMetricsRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveBatchSourceStreamAvgMetrics(ctx context.Context, arg *DescribeLiveBatchSourceStreamAvgMetricsBody) (*DescribeLiveBatchSourceStreamAvgMetricsRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveBatchSourceStreamAvgMetrics", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveBatchSourceStreamAvgMetricsRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveBatchOnlineStreamMetrics(ctx context.Context, arg *DescribeLiveBatchOnlineStreamMetricsBody) (*DescribeLiveBatchOnlineStreamMetricsRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveBatchOnlineStreamMetrics", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveBatchOnlineStreamMetricsRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveBatchPushStreamMetrics(ctx context.Context, arg *DescribeLiveBatchPushStreamMetricsBody) (*DescribeLiveBatchPushStreamMetricsRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveBatchPushStreamMetrics", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveBatchPushStreamMetricsRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveBatchPushStreamAvgMetrics(ctx context.Context, arg *DescribeLiveBatchPushStreamAvgMetricsBody) (*DescribeLiveBatchPushStreamAvgMetricsRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveBatchPushStreamAvgMetrics", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveBatchPushStreamAvgMetricsRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveBatchStreamTranscodeData(ctx context.Context, arg *DescribeLiveBatchStreamTranscodeDataBody) (*DescribeLiveBatchStreamTranscodeDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveBatchStreamTranscodeData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveBatchStreamTranscodeDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveStreamCountData(ctx context.Context, arg *DescribeLiveStreamCountDataBody) (*DescribeLiveStreamCountDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveStreamCountData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveStreamCountDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLivePushStreamCountData(ctx context.Context, arg *DescribeLivePushStreamCountDataBody) (*DescribeLivePushStreamCountDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLivePushStreamCountData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLivePushStreamCountDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveSourceBandwidthData(ctx context.Context, arg *DescribeLiveSourceBandwidthDataBody) (*DescribeLiveSourceBandwidthDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveSourceBandwidthData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveSourceBandwidthDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveSourceTrafficData(ctx context.Context, arg *DescribeLiveSourceTrafficDataBody) (*DescribeLiveSourceTrafficDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveSourceTrafficData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveSourceTrafficDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveMetricBandwidthData(ctx context.Context, arg *DescribeLiveMetricBandwidthDataBody) (*DescribeLiveMetricBandwidthDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveMetricBandwidthData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveMetricBandwidthDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveMetricTrafficData(ctx context.Context, arg *DescribeLiveMetricTrafficDataBody) (*DescribeLiveMetricTrafficDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveMetricTrafficData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveMetricTrafficDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveBatchStreamTrafficData(ctx context.Context, arg *DescribeLiveBatchStreamTrafficDataBody) (*DescribeLiveBatchStreamTrafficDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveBatchStreamTrafficData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveBatchStreamTrafficDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveStreamSessionData(ctx context.Context, arg *DescribeLiveStreamSessionDataBody) (*DescribeLiveStreamSessionDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveStreamSessionData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveStreamSessionDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveISPData(ctx context.Context) (*DescribeLiveISPDataRes, error) {

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveISPData", url.Values{}, "")
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveISPDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveP95PeakBandwidthData(ctx context.Context, arg *DescribeLiveP95PeakBandwidthDataBody) (*DescribeLiveP95PeakBandwidthDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveP95PeakBandwidthData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveP95PeakBandwidthDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveAuditData(ctx context.Context, arg *DescribeLiveAuditDataBody) (*DescribeLiveAuditDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveAuditData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveAuditDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLivePullToPushBandwidthData(ctx context.Context, arg *DescribeLivePullToPushBandwidthDataBody) (*DescribeLivePullToPushBandwidthDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLivePullToPushBandwidthData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLivePullToPushBandwidthDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLivePullToPushData(ctx context.Context, arg *DescribeLivePullToPushDataBody) (*DescribeLivePullToPushDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLivePullToPushData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLivePullToPushDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveBandwidthData(ctx context.Context, arg *DescribeLiveBandwidthDataBody) (*DescribeLiveBandwidthDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveBandwidthData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveBandwidthDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveRecordData(ctx context.Context, arg *DescribeLiveRecordDataBody) (*DescribeLiveRecordDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveRecordData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveRecordDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveSnapshotData(ctx context.Context, arg *DescribeLiveSnapshotDataBody) (*DescribeLiveSnapshotDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveSnapshotData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveSnapshotDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveTrafficData(ctx context.Context, arg *DescribeLiveTrafficDataBody) (*DescribeLiveTrafficDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveTrafficData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveTrafficDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveTranscodeData(ctx context.Context, arg *DescribeLiveTranscodeDataBody) (*DescribeLiveTranscodeDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveTranscodeData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveTranscodeDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveTimeShiftData(ctx context.Context, arg *DescribeLiveTimeShiftDataBody) (*DescribeLiveTimeShiftDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveTimeShiftData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveTimeShiftDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeActionHistory(ctx context.Context, arg *DescribeActionHistoryBody) (*DescribeActionHistoryRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeActionHistory", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeActionHistoryRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListActionHistory(ctx context.Context, arg *ListActionHistoryBody) (*ListActionHistoryRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListActionHistory", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListActionHistoryRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteDenseSnapshotPreset(ctx context.Context, arg *DeleteDenseSnapshotPresetBody) (*DeleteDenseSnapshotPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteDenseSnapshotPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteDenseSnapshotPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateDenseSnapshotPreset(ctx context.Context, arg *UpdateDenseSnapshotPresetBody) (*UpdateDenseSnapshotPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateDenseSnapshotPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateDenseSnapshotPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListVhostDenseSnapshotPreset(ctx context.Context, arg *ListVhostDenseSnapshotPresetBody) (*ListVhostDenseSnapshotPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListVhostDenseSnapshotPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListVhostDenseSnapshotPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescDenseSnapshotPresetDetail(ctx context.Context, arg *DescDenseSnapshotPresetDetailBody) (*DescDenseSnapshotPresetDetailRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescDenseSnapshotPresetDetail", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescDenseSnapshotPresetDetailRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateDenseSnapshotPreset(ctx context.Context, arg *CreateDenseSnapshotPresetBody) (*CreateDenseSnapshotPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateDenseSnapshotPreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateDenseSnapshotPresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveCustomizedLogData(ctx context.Context, arg *DescribeLiveCustomizedLogDataBody) (*DescribeLiveCustomizedLogDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveCustomizedLogData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveCustomizedLogDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveLogData(ctx context.Context, arg *DescribeLiveLogDataBody) (*DescribeLiveLogDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveLogData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveLogDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) AssociatePreset(ctx context.Context, arg *AssociatePresetBody) (*AssociatePresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "AssociatePreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(AssociatePresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DisAssociatePreset(ctx context.Context, arg *DisAssociatePresetBody) (*DisAssociatePresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DisAssociatePreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DisAssociatePresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdatePresetAssociation(ctx context.Context, arg *UpdatePresetAssociationBody) (*UpdatePresetAssociationRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdatePresetAssociation", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdatePresetAssociationRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribePresetAssociation(ctx context.Context, arg *DescribePresetAssociationBody) (*DescribePresetAssociationRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribePresetAssociation", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribePresetAssociationRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateTicket(ctx context.Context, arg *CreateTicketBody) (*CreateTicketRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateTicket", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateTicketRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteReferer(ctx context.Context, arg *DeleteRefererBody) (*DeleteRefererRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteReferer", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteRefererRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeDenyConfig(ctx context.Context, arg *DescribeDenyConfigBody) (*DescribeDenyConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeDenyConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeDenyConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeReferer(ctx context.Context, arg *DescribeRefererBody) (*DescribeRefererRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeReferer", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeRefererRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateDenyConfig(ctx context.Context, arg *UpdateDenyConfigBody) (*UpdateDenyConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateDenyConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateDenyConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateDenyConfigV2(ctx context.Context, arg *UpdateDenyConfigV2Body) (*UpdateDenyConfigV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateDenyConfigV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateDenyConfigV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateReferer(ctx context.Context, arg *UpdateRefererBody) (*UpdateRefererRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateReferer", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateRefererRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateAuthKey(ctx context.Context, arg *UpdateAuthKeyBody) (*UpdateAuthKeyRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateAuthKey", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateAuthKeyRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteRelaySink(ctx context.Context, arg *DeleteRelaySinkBody) (*DeleteRelaySinkRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteRelaySink", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteRelaySinkRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateRelaySink(ctx context.Context, arg *UpdateRelaySinkBody) (*UpdateRelaySinkRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateRelaySink", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateRelaySinkRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeRelaySink(ctx context.Context, arg *DescribeRelaySinkBody) (*DescribeRelaySinkRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeRelaySink", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeRelaySinkRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteHLSConfig(ctx context.Context, arg *DeleteHLSConfigBody) (*DeleteHLSConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteHLSConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteHLSConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateHLSConfig(ctx context.Context, arg *UpdateHLSConfigBody) (*UpdateHLSConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateHLSConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateHLSConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeHLSConfig(ctx context.Context, arg *DescribeHLSConfigBody) (*DescribeHLSConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeHLSConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeHLSConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteHTTPHeaderConfig(ctx context.Context, arg *DeleteHTTPHeaderConfigBody) (*DeleteHTTPHeaderConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteHTTPHeaderConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteHTTPHeaderConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteHeaderConfig(ctx context.Context, arg *DeleteHeaderConfigBody) (*DeleteHeaderConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteHeaderConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteHeaderConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) EnableHTTPHeaderConfig(ctx context.Context, arg *EnableHTTPHeaderConfigBody) (*EnableHTTPHeaderConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "EnableHTTPHeaderConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(EnableHTTPHeaderConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateHTTPHeaderConfig(ctx context.Context, arg *UpdateHTTPHeaderConfigBody) (*UpdateHTTPHeaderConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateHTTPHeaderConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateHTTPHeaderConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateHeaderConfig(ctx context.Context, arg *UpdateHeaderConfigBody) (*UpdateHeaderConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateHeaderConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateHeaderConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeHTTPHeaderConfig(ctx context.Context, arg *DescribeHTTPHeaderConfigBody) (*DescribeHTTPHeaderConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeHTTPHeaderConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeHTTPHeaderConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeHeaderConfig(ctx context.Context, arg *DescribeHeaderConfigBody) (*DescribeHeaderConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeHeaderConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeHeaderConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListHeaderEnum(ctx context.Context, arg *ListHeaderEnumBody) (*ListHeaderEnumRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListHeaderEnum", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListHeaderEnumRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteNSSRewriteConfig(ctx context.Context, arg *DeleteNSSRewriteConfigBody) (*DeleteNSSRewriteConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteNSSRewriteConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteNSSRewriteConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateNSSRewriteConfig(ctx context.Context, arg *UpdateNSSRewriteConfigBody) (*UpdateNSSRewriteConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateNSSRewriteConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateNSSRewriteConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeNSSRewriteConfig(ctx context.Context, arg *DescribeNSSRewriteConfigBody) (*DescribeNSSRewriteConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeNSSRewriteConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeNSSRewriteConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveActivityBandwidthData(ctx context.Context, arg *DescribeLiveActivityBandwidthDataBody) (*DescribeLiveActivityBandwidthDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveActivityBandwidthData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveActivityBandwidthDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateLiveAccountFeeConfig(ctx context.Context, arg *CreateLiveAccountFeeConfigBody) (*CreateLiveAccountFeeConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateLiveAccountFeeConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateLiveAccountFeeConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteLiveAccountFeeConfig(ctx context.Context, arg *DeleteLiveAccountFeeConfigBody) (*DeleteLiveAccountFeeConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteLiveAccountFeeConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteLiveAccountFeeConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveAccountFeeConfig(ctx context.Context) (*DescribeLiveAccountFeeConfigRes, error) {

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveAccountFeeConfig", url.Values{}, "")
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveAccountFeeConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveStreamUsageData(ctx context.Context, arg *DescribeLiveStreamUsageDataBody) (*DescribeLiveStreamUsageDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveStreamUsageData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveStreamUsageDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveFeeConfig(ctx context.Context) (*DescribeLiveFeeConfigRes, error) {

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveFeeConfig", url.Values{}, "")
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveFeeConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveAccountFeeType(ctx context.Context, arg *DescribeLiveAccountFeeTypeBody) (*DescribeLiveAccountFeeTypeRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLiveAccountFeeType", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveAccountFeeTypeRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateEncryptDRM(ctx context.Context, arg *UpdateEncryptDRMBody) (*UpdateEncryptDRMRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateEncryptDRM", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateEncryptDRMRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeContentKey(ctx context.Context) (*DescribeContentKeyRes, error) {

	data, _, err := c.Client.CtxJson(ctx, "DescribeContentKey", url.Values{}, "")
	if err != nil {
		return nil, err
	}

	result := new(DescribeContentKeyRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeCertDRM(ctx context.Context, arg *DescribeCertDRMQuery) (*DescribeCertDRMRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.CtxQuery(ctx, "DescribeCertDRM", query)
	if err != nil {
		return nil, err
	}

	result := new(DescribeCertDRMRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLicenseDRM(ctx context.Context, arg *DescribeLicenseDRMQuery) (*DescribeLicenseDRMRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLicenseDRM", query, "")
	if err != nil {
		return nil, err
	}

	result := new(DescribeLicenseDRMRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeEncryptDRM(ctx context.Context) (*DescribeEncryptDRMRes, error) {

	data, _, err := c.Client.CtxJson(ctx, "DescribeEncryptDRM", url.Values{}, "")
	if err != nil {
		return nil, err
	}

	result := new(DescribeEncryptDRMRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) BindEncryptDRM(ctx context.Context, arg *BindEncryptDRMBody) (*BindEncryptDRMRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "BindEncryptDRM", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(BindEncryptDRMRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UnBindEncryptDRM(ctx context.Context, arg *UnBindEncryptDRMBody) (*UnBindEncryptDRMRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UnBindEncryptDRM", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UnBindEncryptDRMRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListBindEncryptDRM(ctx context.Context, arg *ListBindEncryptDRMBody) (*ListBindEncryptDRMRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListBindEncryptDRM", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListBindEncryptDRMRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateCustomLogConfig(ctx context.Context, arg *CreateCustomLogConfigBody) (*CreateCustomLogConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateCustomLogConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateCustomLogConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteCustomLogConfig(ctx context.Context, arg *DeleteCustomLogConfigBody) (*DeleteCustomLogConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteCustomLogConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteCustomLogConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeCustomLogConfig(ctx context.Context) (*DescribeCustomLogConfigRes, error) {

	data, _, err := c.Client.CtxJson(ctx, "DescribeCustomLogConfig", url.Values{}, "")
	if err != nil {
		return nil, err
	}

	result := new(DescribeCustomLogConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CheckCustomLogConfig(ctx context.Context, arg *CheckCustomLogConfigBody) (*CheckCustomLogConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CheckCustomLogConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CheckCustomLogConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateTranscodePresetBatch(ctx context.Context, arg *CreateTranscodePresetBatchBody) (*CreateTranscodePresetBatchRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateTranscodePresetBatch", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateTranscodePresetBatchRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteTranscodePresetBatch(ctx context.Context, arg *DeleteTranscodePresetBatchBody) (*DeleteTranscodePresetBatchRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteTranscodePresetBatch", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteTranscodePresetBatchRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) AssociateRefConfig(ctx context.Context, arg *AssociateRefConfigBody) (*AssociateRefConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "AssociateRefConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(AssociateRefConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DisassociateRefConfig(ctx context.Context, arg *DisassociateRefConfigBody) (*DisassociateRefConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DisassociateRefConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DisassociateRefConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeRefConfig(ctx context.Context, arg *DescribeRefConfigBody) (*DescribeRefConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeRefConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeRefConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListReferenceNames(ctx context.Context, arg *ListReferenceNamesBody) (*ListReferenceNamesRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListReferenceNames", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListReferenceNamesRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListReferenceTypes(ctx context.Context) (*ListReferenceTypesRes, error) {

	data, _, err := c.CtxQuery(ctx, "ListReferenceTypes", url.Values{})
	if err != nil {
		return nil, err
	}

	result := new(ListReferenceTypesRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListReferenceInfo(ctx context.Context, arg *ListReferenceInfoBody) (*ListReferenceInfoRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListReferenceInfo", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListReferenceInfoRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateAvSlicePreset(ctx context.Context, arg *CreateAvSlicePresetBody) (*CreateAvSlicePresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateAvSlicePreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateAvSlicePresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteAvSlicePreset(ctx context.Context, arg *DeleteAvSlicePresetBody) (*DeleteAvSlicePresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteAvSlicePreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteAvSlicePresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateAvSlicePreset(ctx context.Context, arg *UpdateAvSlicePresetBody) (*UpdateAvSlicePresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateAvSlicePreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateAvSlicePresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteIPAccessRule(ctx context.Context, arg *DeleteIPAccessRuleBody) (*DeleteIPAccessRuleRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteIPAccessRule", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteIPAccessRuleRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateIPAccessRule(ctx context.Context, arg *UpdateIPAccessRuleBody) (*UpdateIPAccessRuleRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateIPAccessRule", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateIPAccessRuleRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeIPAccessRule(ctx context.Context, arg *DescribeIPAccessRuleBody) (*DescribeIPAccessRuleRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeIPAccessRule", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeIPAccessRuleRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateProxyConfig(ctx context.Context, arg *CreateProxyConfigBody) (*CreateProxyConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "CreateProxyConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateProxyConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteProxyConfigAssociation(ctx context.Context, arg *DeleteProxyConfigAssociationBody) (*DeleteProxyConfigAssociationRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteProxyConfigAssociation", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteProxyConfigAssociationRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteProxyConfig(ctx context.Context, arg *DeleteProxyConfigBody) (*DeleteProxyConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteProxyConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteProxyConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateProxyConfigAssociation(ctx context.Context, arg *UpdateProxyConfigAssociationBody) (*UpdateProxyConfigAssociationRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateProxyConfigAssociation", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateProxyConfigAssociationRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateProxyConfig(ctx context.Context, arg *UpdateProxyConfigBody) (*UpdateProxyConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateProxyConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateProxyConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeProxyConfigAssociation(ctx context.Context, arg *DescribeProxyConfigAssociationBody) (*DescribeProxyConfigAssociationRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeProxyConfigAssociation", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeProxyConfigAssociationRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListProxyConfig(ctx context.Context, arg *ListProxyConfigBody) (*ListProxyConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "ListProxyConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListProxyConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteCMAFConfig(ctx context.Context, arg *DeleteCMAFConfigBody) (*DeleteCMAFConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteCMAFConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteCMAFConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateCMAFConfig(ctx context.Context, arg *UpdateCMAFConfigBody) (*UpdateCMAFConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateCMAFConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateCMAFConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeCMAFConfig(ctx context.Context, arg *DescribeCMAFConfigBody) (*DescribeCMAFConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeCMAFConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeCMAFConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteLatencyConfig(ctx context.Context, arg *DeleteLatencyConfigBody) (*DeleteLatencyConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DeleteLatencyConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteLatencyConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateLatencyConfig(ctx context.Context, arg *UpdateLatencyConfigBody) (*UpdateLatencyConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "UpdateLatencyConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateLatencyConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLatencyConfig(ctx context.Context, arg *DescribeLatencyConfigBody) (*DescribeLatencyConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.CtxJson(ctx, "DescribeLatencyConfig", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLatencyConfigRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
