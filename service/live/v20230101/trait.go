package live_v20230101

import (
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

func (c *Live) DeleteTranscodePreset(arg *DeleteTranscodePresetBody) (*DeleteTranscodePresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DeleteTranscodePreset", url.Values{}, string(body))
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

func (c *Live) UpdateTranscodePreset(arg *UpdateTranscodePresetBody) (*UpdateTranscodePresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("UpdateTranscodePreset", url.Values{}, string(body))
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

func (c *Live) ListCommonTransPresetDetail(arg *ListCommonTransPresetDetailBody) (*ListCommonTransPresetDetailRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("ListCommonTransPresetDetail", url.Values{}, string(body))
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

func (c *Live) ListVhostTransCodePreset(arg *ListVhostTransCodePresetBody) (*ListVhostTransCodePresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("ListVhostTransCodePreset", url.Values{}, string(body))
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

func (c *Live) CreateTranscodePreset(arg *CreateTranscodePresetBody) (*CreateTranscodePresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("CreateTranscodePreset", url.Values{}, string(body))
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

func (c *Live) CreateWatermarkPreset(arg *CreateWatermarkPresetBody) (*CreateWatermarkPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("CreateWatermarkPreset", url.Values{}, string(body))
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

func (c *Live) UpdateWatermarkPreset(arg *UpdateWatermarkPresetBody) (*UpdateWatermarkPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("UpdateWatermarkPreset", url.Values{}, string(body))
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

func (c *Live) DeleteWatermarkPreset(arg *DeleteWatermarkPresetBody) (*DeleteWatermarkPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DeleteWatermarkPreset", url.Values{}, string(body))
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

func (c *Live) ListWatermarkPreset(arg *ListWatermarkPresetBody) (*ListWatermarkPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("ListWatermarkPreset", url.Values{}, string(body))
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

func (c *Live) ListVhostWatermarkPreset(arg *ListVhostWatermarkPresetBody) (*ListVhostWatermarkPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("ListVhostWatermarkPreset", url.Values{}, string(body))
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

func (c *Live) DeleteRecordPreset(arg *DeleteRecordPresetBody) (*DeleteRecordPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DeleteRecordPreset", url.Values{}, string(body))
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

func (c *Live) UpdateRecordPresetV2(arg *UpdateRecordPresetV2Body) (*UpdateRecordPresetV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("UpdateRecordPresetV2", url.Values{}, string(body))
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

func (c *Live) DescribeRecordTaskFileHistory(arg *DescribeRecordTaskFileHistoryBody) (*DescribeRecordTaskFileHistoryRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeRecordTaskFileHistory", url.Values{}, string(body))
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

func (c *Live) ListVhostRecordPresetV2(arg *ListVhostRecordPresetV2Body) (*ListVhostRecordPresetV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("ListVhostRecordPresetV2", url.Values{}, string(body))
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

func (c *Live) CreateRecordPresetV2(arg *CreateRecordPresetV2Body) (*CreateRecordPresetV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("CreateRecordPresetV2", url.Values{}, string(body))
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

func (c *Live) DeleteSnapshotPreset(arg *DeleteSnapshotPresetBody) (*DeleteSnapshotPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DeleteSnapshotPreset", url.Values{}, string(body))
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

func (c *Live) UpdateSnapshotPreset(arg *UpdateSnapshotPresetBody) (*UpdateSnapshotPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("UpdateSnapshotPreset", url.Values{}, string(body))
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

func (c *Live) DescribeCDNSnapshotHistory(arg *DescribeCDNSnapshotHistoryBody) (*DescribeCDNSnapshotHistoryRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeCDNSnapshotHistory", url.Values{}, string(body))
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

func (c *Live) ListVhostSnapshotPreset(arg *ListVhostSnapshotPresetBody) (*ListVhostSnapshotPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("ListVhostSnapshotPreset", url.Values{}, string(body))
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

func (c *Live) CreateSnapshotPreset(arg *CreateSnapshotPresetBody) (*CreateSnapshotPresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("CreateSnapshotPreset", url.Values{}, string(body))
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

func (c *Live) DeleteTimeShiftPresetV2(arg *DeleteTimeShiftPresetV2Body) (*DeleteTimeShiftPresetV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DeleteTimeShiftPresetV2", url.Values{}, string(body))
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

func (c *Live) CreateTimeShiftPresetV2(arg *CreateTimeShiftPresetV2Body) (*CreateTimeShiftPresetV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("CreateTimeShiftPresetV2", url.Values{}, string(body))
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

func (c *Live) UpdateTimeShiftPresetV2(arg *UpdateTimeShiftPresetV2Body) (*UpdateTimeShiftPresetV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("UpdateTimeShiftPresetV2", url.Values{}, string(body))
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

func (c *Live) ListTimeShiftPresetV2(arg *ListTimeShiftPresetV2Body) (*ListTimeShiftPresetV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("ListTimeShiftPresetV2", url.Values{}, string(body))
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

func (c *Live) DeleteCallback(arg *DeleteCallbackBody) (*DeleteCallbackRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DeleteCallback", url.Values{}, string(body))
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

func (c *Live) DescribeCallback(arg *DescribeCallbackBody) (*DescribeCallbackRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeCallback", url.Values{}, string(body))
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

func (c *Live) UpdateCallback(arg *UpdateCallbackBody) (*UpdateCallbackRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("UpdateCallback", url.Values{}, string(body))
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

func (c *Live) DeleteCert(arg *DeleteCertBody) (*DeleteCertRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DeleteCert", url.Values{}, string(body))
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

func (c *Live) DescribeCertDetailSecretV2(arg *DescribeCertDetailSecretV2Body) (*DescribeCertDetailSecretV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeCertDetailSecretV2", url.Values{}, string(body))
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

func (c *Live) ListCertV2(arg *ListCertV2Body) (*ListCertV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("ListCertV2", url.Values{}, string(body))
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

func (c *Live) CreateCert(arg *CreateCertBody) (*CreateCertRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("CreateCert", url.Values{}, string(body))
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

func (c *Live) BindCert(arg *BindCertBody) (*BindCertRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("BindCert", url.Values{}, string(body))
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

func (c *Live) UnbindCert(arg *UnbindCertBody) (*UnbindCertRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("UnbindCert", url.Values{}, string(body))
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

func (c *Live) DeleteDomain(arg *DeleteDomainBody) (*DeleteDomainRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DeleteDomain", url.Values{}, string(body))
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

func (c *Live) EnableDomain(arg *EnableDomainBody) (*EnableDomainRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("EnableDomain", url.Values{}, string(body))
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

func (c *Live) CreateDomainV2(arg *CreateDomainV2Body) (*CreateDomainV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("CreateDomainV2", url.Values{}, string(body))
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

func (c *Live) UpdateDomainVhost(arg *UpdateDomainVhostBody) (*UpdateDomainVhostRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("UpdateDomainVhost", url.Values{}, string(body))
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

func (c *Live) DescribeDomain(arg *DescribeDomainBody) (*DescribeDomainRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeDomain", url.Values{}, string(body))
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

func (c *Live) ListDomainDetail(arg *ListDomainDetailBody) (*ListDomainDetailRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("ListDomainDetail", url.Values{}, string(body))
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

func (c *Live) DisableDomain(arg *DisableDomainBody) (*DisableDomainRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DisableDomain", url.Values{}, string(body))
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

func (c *Live) CreateLiveVideoQualityAnalysisTask(arg *CreateLiveVideoQualityAnalysisTaskBody) (*CreateLiveVideoQualityAnalysisTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("CreateLiveVideoQualityAnalysisTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateLiveVideoQualityAnalysisTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteLiveVideoQualityAnalysisTask(arg *DeleteLiveVideoQualityAnalysisTaskBody) (*DeleteLiveVideoQualityAnalysisTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DeleteLiveVideoQualityAnalysisTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteLiveVideoQualityAnalysisTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) GetLiveVideoQualityAnalysisTaskDetail(arg *GetLiveVideoQualityAnalysisTaskDetailBody) (*GetLiveVideoQualityAnalysisTaskDetailRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("GetLiveVideoQualityAnalysisTaskDetail", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetLiveVideoQualityAnalysisTaskDetailRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListLiveVideoQualityAnalysisTasks(arg *ListLiveVideoQualityAnalysisTasksBody) (*ListLiveVideoQualityAnalysisTasksRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("ListLiveVideoQualityAnalysisTasks", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListLiveVideoQualityAnalysisTasksRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) StopPullToPushTask(arg *StopPullToPushTaskBody) (*StopPullToPushTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("StopPullToPushTask", url.Values{}, string(body))
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

func (c *Live) CreatePullToPushTask(arg *CreatePullToPushTaskBody) (*CreatePullToPushTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("CreatePullToPushTask", url.Values{}, string(body))
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

func (c *Live) CreatePullToPushGroup(arg *CreatePullToPushGroupBody) (*CreatePullToPushGroupRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("CreatePullToPushGroup", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreatePullToPushGroupRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeletePullToPushTask(arg *DeletePullToPushTaskBody) (*DeletePullToPushTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DeletePullToPushTask", url.Values{}, string(body))
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

func (c *Live) DeletePullToPushGroup(arg *DeletePullToPushGroupBody) (*DeletePullToPushGroupRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DeletePullToPushGroup", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeletePullToPushGroupRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) RestartPullToPushTask(arg *RestartPullToPushTaskBody) (*RestartPullToPushTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("RestartPullToPushTask", url.Values{}, string(body))
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

func (c *Live) UpdatePullToPushTask(arg *UpdatePullToPushTaskBody) (*UpdatePullToPushTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("UpdatePullToPushTask", url.Values{}, string(body))
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

func (c *Live) ListPullToPushGroup(arg *ListPullToPushGroupBody) (*ListPullToPushGroupRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("ListPullToPushGroup", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListPullToPushGroupRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListPullToPushTask(arg *ListPullToPushTaskQuery) (*ListPullToPushTaskRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Query("ListPullToPushTask", query)
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

func (c *Live) ListPullToPushTaskV2(arg *ListPullToPushTaskV2Body) (*ListPullToPushTaskV2Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("ListPullToPushTaskV2", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListPullToPushTaskV2Res)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteRelaySourceRewrite(arg *DeleteRelaySourceRewriteBody) (*DeleteRelaySourceRewriteRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DeleteRelaySourceRewrite", url.Values{}, string(body))
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

func (c *Live) DeleteRelaySourceV3(arg *DeleteRelaySourceV3Body) (*DeleteRelaySourceV3Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DeleteRelaySourceV3", url.Values{}, string(body))
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

func (c *Live) UpdateRelaySourceRewrite(arg *UpdateRelaySourceRewriteBody) (*UpdateRelaySourceRewriteRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("UpdateRelaySourceRewrite", url.Values{}, string(body))
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

func (c *Live) DescribeRelaySourceRewrite(arg *DescribeRelaySourceRewriteBody) (*DescribeRelaySourceRewriteRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeRelaySourceRewrite", url.Values{}, string(body))
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

func (c *Live) DescribeRelaySourceV3(arg *DescribeRelaySourceV3Body) (*DescribeRelaySourceV3Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeRelaySourceV3", url.Values{}, string(body))
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

func (c *Live) UpdateRelaySourceV3(arg *UpdateRelaySourceV3Body) (*UpdateRelaySourceV3Res, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("UpdateRelaySourceV3", url.Values{}, string(body))
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

func (c *Live) KillStream(arg *KillStreamBody) (*KillStreamRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("KillStream", url.Values{}, string(body))
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

func (c *Live) DescribeClosedStreamInfoByPage(arg *DescribeClosedStreamInfoByPageQuery) (*DescribeClosedStreamInfoByPageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Query("DescribeClosedStreamInfoByPage", query)
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

func (c *Live) DescribeLiveStreamInfoByPage(arg *DescribeLiveStreamInfoByPageQuery) (*DescribeLiveStreamInfoByPageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Query("DescribeLiveStreamInfoByPage", query)
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

func (c *Live) DescribeLiveStreamState(arg *DescribeLiveStreamStateQuery) (*DescribeLiveStreamStateRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Query("DescribeLiveStreamState", query)
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

func (c *Live) DescribeForbiddenStreamInfoByPage(arg *DescribeForbiddenStreamInfoByPageQuery) (*DescribeForbiddenStreamInfoByPageRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Query("DescribeForbiddenStreamInfoByPage", query)
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

func (c *Live) ForbidStream(arg *ForbidStreamBody) (*ForbidStreamRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("ForbidStream", url.Values{}, string(body))
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

func (c *Live) ResumeStream(arg *ResumeStreamBody) (*ResumeStreamRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("ResumeStream", url.Values{}, string(body))
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

func (c *Live) GeneratePlayURL(arg *GeneratePlayURLBody) (*GeneratePlayURLRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("GeneratePlayURL", url.Values{}, string(body))
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

func (c *Live) GeneratePushURL(arg *GeneratePushURLBody) (*GeneratePushURLRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("GeneratePushURL", url.Values{}, string(body))
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

func (c *Live) DescribeLiveRegionData() (*DescribeLiveRegionDataRes, error) {

	data, _, err := c.Client.Json("DescribeLiveRegionData", url.Values{}, "")
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

func (c *Live) DescribeLiveSourceStreamMetrics(arg *DescribeLiveSourceStreamMetricsBody) (*DescribeLiveSourceStreamMetricsRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeLiveSourceStreamMetrics", url.Values{}, string(body))
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

func (c *Live) DescribeLivePushStreamMetrics(arg *DescribeLivePushStreamMetricsBody) (*DescribeLivePushStreamMetricsRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeLivePushStreamMetrics", url.Values{}, string(body))
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

func (c *Live) DescribeLiveStreamSessionData(arg *DescribeLiveStreamSessionDataBody) (*DescribeLiveStreamSessionDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeLiveStreamSessionData", url.Values{}, string(body))
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

func (c *Live) DescribeLivePlayStatusCodeData(arg *DescribeLivePlayStatusCodeDataBody) (*DescribeLivePlayStatusCodeDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeLivePlayStatusCodeData", url.Values{}, string(body))
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

func (c *Live) DescribeLiveBatchPushStreamMetrics(arg *DescribeLiveBatchPushStreamMetricsBody) (*DescribeLiveBatchPushStreamMetricsRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeLiveBatchPushStreamMetrics", url.Values{}, string(body))
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

func (c *Live) DescribeLivePushStreamInfoData(arg *DescribeLivePushStreamInfoDataBody) (*DescribeLivePushStreamInfoDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeLivePushStreamInfoData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLivePushStreamInfoDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveTranscodeInfoData(arg *DescribeLiveTranscodeInfoDataBody) (*DescribeLiveTranscodeInfoDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeLiveTranscodeInfoData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveTranscodeInfoDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveSourceBandwidthData(arg *DescribeLiveSourceBandwidthDataBody) (*DescribeLiveSourceBandwidthDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeLiveSourceBandwidthData", url.Values{}, string(body))
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

func (c *Live) DescribeLiveSourceTrafficData(arg *DescribeLiveSourceTrafficDataBody) (*DescribeLiveSourceTrafficDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeLiveSourceTrafficData", url.Values{}, string(body))
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

func (c *Live) DescribeLiveMetricBandwidthData(arg *DescribeLiveMetricBandwidthDataBody) (*DescribeLiveMetricBandwidthDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeLiveMetricBandwidthData", url.Values{}, string(body))
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

func (c *Live) DescribeLiveMetricTrafficData(arg *DescribeLiveMetricTrafficDataBody) (*DescribeLiveMetricTrafficDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeLiveMetricTrafficData", url.Values{}, string(body))
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

func (c *Live) DescribeLiveISPData() (*DescribeLiveISPDataRes, error) {

	data, _, err := c.Client.Json("DescribeLiveISPData", url.Values{}, "")
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

func (c *Live) DescribeLiveP95PeakBandwidthData(arg *DescribeLiveP95PeakBandwidthDataBody) (*DescribeLiveP95PeakBandwidthDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeLiveP95PeakBandwidthData", url.Values{}, string(body))
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

func (c *Live) DescribeLiveASRDurationData(arg *DescribeLiveASRDurationDataBody) (*DescribeLiveASRDurationDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeLiveASRDurationData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveASRDurationDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLivePullToPushBandwidthData(arg *DescribeLivePullToPushBandwidthDataBody) (*DescribeLivePullToPushBandwidthDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeLivePullToPushBandwidthData", url.Values{}, string(body))
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

func (c *Live) DescribeLivePullToPushData(arg *DescribeLivePullToPushDataBody) (*DescribeLivePullToPushDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeLivePullToPushData", url.Values{}, string(body))
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

func (c *Live) DescribeLiveBandwidthData(arg *DescribeLiveBandwidthDataBody) (*DescribeLiveBandwidthDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeLiveBandwidthData", url.Values{}, string(body))
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

func (c *Live) DescribeLiveRecordData(arg *DescribeLiveRecordDataBody) (*DescribeLiveRecordDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeLiveRecordData", url.Values{}, string(body))
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

func (c *Live) DescribeLiveSnapshotData(arg *DescribeLiveSnapshotDataBody) (*DescribeLiveSnapshotDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeLiveSnapshotData", url.Values{}, string(body))
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

func (c *Live) DescribeLiveTrafficData(arg *DescribeLiveTrafficDataBody) (*DescribeLiveTrafficDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeLiveTrafficData", url.Values{}, string(body))
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

func (c *Live) DescribeLiveTranscodeData(arg *DescribeLiveTranscodeDataBody) (*DescribeLiveTranscodeDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeLiveTranscodeData", url.Values{}, string(body))
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

func (c *Live) DescribeLiveStorageSpaceData(arg *DescribeLiveStorageSpaceDataBody) (*DescribeLiveStorageSpaceDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeLiveStorageSpaceData", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeLiveStorageSpaceDataRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLiveLogData(arg *DescribeLiveLogDataBody) (*DescribeLiveLogDataRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeLiveLogData", url.Values{}, string(body))
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

func (c *Live) DeleteReferer(arg *DeleteRefererBody) (*DeleteRefererRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DeleteReferer", url.Values{}, string(body))
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

func (c *Live) DescribeReferer(arg *DescribeRefererBody) (*DescribeRefererRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeReferer", url.Values{}, string(body))
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

func (c *Live) DescribeAuth(arg *DescribeAuthBody) (*DescribeAuthRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeAuth", url.Values{}, string(body))
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

func (c *Live) UpdateReferer(arg *UpdateRefererBody) (*UpdateRefererRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("UpdateReferer", url.Values{}, string(body))
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

func (c *Live) UpdateAuthKey(arg *UpdateAuthKeyBody) (*UpdateAuthKeyRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("UpdateAuthKey", url.Values{}, string(body))
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

func (c *Live) DeleteHLSConfig(arg *DeleteHLSConfigBody) (*DeleteHLSConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DeleteHLSConfig", url.Values{}, string(body))
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

func (c *Live) UpdateHLSConfig(arg *UpdateHLSConfigBody) (*UpdateHLSConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("UpdateHLSConfig", url.Values{}, string(body))
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

func (c *Live) DescribeHLSConfig(arg *DescribeHLSConfigBody) (*DescribeHLSConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeHLSConfig", url.Values{}, string(body))
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

func (c *Live) DeleteHTTPHeaderConfig(arg *DeleteHTTPHeaderConfigBody) (*DeleteHTTPHeaderConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DeleteHTTPHeaderConfig", url.Values{}, string(body))
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

func (c *Live) EnableHTTPHeaderConfig(arg *EnableHTTPHeaderConfigBody) (*EnableHTTPHeaderConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("EnableHTTPHeaderConfig", url.Values{}, string(body))
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

func (c *Live) DescribeHTTPHeaderConfig(arg *DescribeHTTPHeaderConfigBody) (*DescribeHTTPHeaderConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeHTTPHeaderConfig", url.Values{}, string(body))
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

func (c *Live) UpdateHTTPHeaderConfig(arg *UpdateHTTPHeaderConfigBody) (*UpdateHTTPHeaderConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("UpdateHTTPHeaderConfig", url.Values{}, string(body))
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

func (c *Live) UpdateEncryptDRM(arg *UpdateEncryptDRMBody) (*UpdateEncryptDRMRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("UpdateEncryptDRM", url.Values{}, string(body))
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

func (c *Live) UpdateEncryptHLS(arg *UpdateEncryptHLSBody) (*UpdateEncryptHLSRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("UpdateEncryptHLS", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateEncryptHLSRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) GetHLSEncryptDataKey(arg *GetHLSEncryptDataKeyQuery) (*GetHLSEncryptDataKeyRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Query("GetHLSEncryptDataKey", query)
	if err != nil {
		return nil, err
	}

	result := new(GetHLSEncryptDataKeyRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeEncryptHLS() (*DescribeEncryptHLSRes, error) {

	data, _, err := c.Client.Json("DescribeEncryptHLS", url.Values{}, "")
	if err != nil {
		return nil, err
	}

	result := new(DescribeEncryptHLSRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeLicenseDRM(arg *DescribeLicenseDRMQuery) (*DescribeLicenseDRMRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeLicenseDRM", query, "")
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

func (c *Live) DescribeCertDRM(arg *DescribeCertDRMQuery) (*DescribeCertDRMRes, error) {
	query, err := marshalToQuery(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Query("DescribeCertDRM", query)
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

func (c *Live) DescribeEncryptDRM() (*DescribeEncryptDRMRes, error) {

	data, _, err := c.Client.Json("DescribeEncryptDRM", url.Values{}, "")
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

func (c *Live) BindEncryptDRM(arg *BindEncryptDRMBody) (*BindEncryptDRMRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("BindEncryptDRM", url.Values{}, string(body))
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

func (c *Live) UnBindEncryptDRM(arg *UnBindEncryptDRMBody) (*UnBindEncryptDRMRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("UnBindEncryptDRM", url.Values{}, string(body))
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

func (c *Live) ListBindEncryptDRM(arg *ListBindEncryptDRMBody) (*ListBindEncryptDRMRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("ListBindEncryptDRM", url.Values{}, string(body))
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

func (c *Live) DeleteIPAccessRule(arg *DeleteIPAccessRuleBody) (*DeleteIPAccessRuleRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DeleteIPAccessRule", url.Values{}, string(body))
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

func (c *Live) DeleteRegionAccessRule(arg *DeleteRegionAccessRuleBody) (*DeleteRegionAccessRuleRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DeleteRegionAccessRule", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteRegionAccessRuleRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateIPAccessRule(arg *UpdateIPAccessRuleBody) (*UpdateIPAccessRuleRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("UpdateIPAccessRule", url.Values{}, string(body))
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

func (c *Live) UpdateRegionAccessRule(arg *UpdateRegionAccessRuleBody) (*UpdateRegionAccessRuleRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("UpdateRegionAccessRule", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateRegionAccessRuleRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeIPAccessRule(arg *DescribeIPAccessRuleBody) (*DescribeIPAccessRuleRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeIPAccessRule", url.Values{}, string(body))
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

func (c *Live) DescribeRegionAccessRule(arg *DescribeRegionAccessRuleBody) (*DescribeRegionAccessRuleRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeRegionAccessRule", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeRegionAccessRuleRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteCMAFConfig(arg *DeleteCMAFConfigBody) (*DeleteCMAFConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DeleteCMAFConfig", url.Values{}, string(body))
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

func (c *Live) UpdateCMAFConfig(arg *UpdateCMAFConfigBody) (*UpdateCMAFConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("UpdateCMAFConfig", url.Values{}, string(body))
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

func (c *Live) DescribeCMAFConfig(arg *DescribeCMAFConfigBody) (*DescribeCMAFConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeCMAFConfig", url.Values{}, string(body))
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

func (c *Live) DeleteLatencyConfig(arg *DeleteLatencyConfigBody) (*DeleteLatencyConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DeleteLatencyConfig", url.Values{}, string(body))
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

func (c *Live) DescribeLatencyConfig(arg *DescribeLatencyConfigBody) (*DescribeLatencyConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeLatencyConfig", url.Values{}, string(body))
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

func (c *Live) UpdateLatencyConfig(arg *UpdateLatencyConfigBody) (*UpdateLatencyConfigRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("UpdateLatencyConfig", url.Values{}, string(body))
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

func (c *Live) CreateCloudMixTask(arg *CreateCloudMixTaskBody) (*CreateCloudMixTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("CreateCloudMixTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateCloudMixTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateCloudMixTask(arg *UpdateCloudMixTaskBody) (*UpdateCloudMixTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("UpdateCloudMixTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateCloudMixTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) GetCloudMixTaskDetail(arg *GetCloudMixTaskDetailBody) (*GetCloudMixTaskDetailRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("GetCloudMixTaskDetail", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(GetCloudMixTaskDetailRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListCloudMixTask(arg *ListCloudMixTaskBody) (*ListCloudMixTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("ListCloudMixTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListCloudMixTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteCloudMixTask(arg *DeleteCloudMixTaskBody) (*DeleteCloudMixTaskRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DeleteCloudMixTask", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteCloudMixTaskRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteClusterRateLimit(arg *DeleteClusterRateLimitBody) (*DeleteClusterRateLimitRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DeleteClusterRateLimit", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteClusterRateLimitRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeClusterRateLimit(arg *DescribeClusterRateLimitBody) (*DescribeClusterRateLimitRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeClusterRateLimit", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeClusterRateLimitRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateClusterRateLimit(arg *UpdateClusterRateLimitBody) (*UpdateClusterRateLimitRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("UpdateClusterRateLimit", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateClusterRateLimitRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteUserAgentAccessRule(arg *DeleteUserAgentAccessRuleBody) (*DeleteUserAgentAccessRuleRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DeleteUserAgentAccessRule", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteUserAgentAccessRuleRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeUserAgentAccessRule(arg *DescribeUserAgentAccessRuleBody) (*DescribeUserAgentAccessRuleRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeUserAgentAccessRule", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeUserAgentAccessRuleRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateUserAgentAccessRule(arg *UpdateUserAgentAccessRuleBody) (*UpdateUserAgentAccessRuleRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("UpdateUserAgentAccessRule", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateUserAgentAccessRuleRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteFormatAccessRule(arg *DeleteFormatAccessRuleBody) (*DeleteFormatAccessRuleRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DeleteFormatAccessRule", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteFormatAccessRuleRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DescribeFormatAccessRule(arg *DescribeFormatAccessRuleBody) (*DescribeFormatAccessRuleRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DescribeFormatAccessRule", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DescribeFormatAccessRuleRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateFormatAccessRule(arg *UpdateFormatAccessRuleBody) (*UpdateFormatAccessRuleRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("UpdateFormatAccessRule", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateFormatAccessRuleRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) DeleteSubtitleTranscodePreset(arg *DeleteSubtitleTranscodePresetBody) (*DeleteSubtitleTranscodePresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("DeleteSubtitleTranscodePreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(DeleteSubtitleTranscodePresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) UpdateSubtitleTranscodePreset(arg *UpdateSubtitleTranscodePresetBody) (*UpdateSubtitleTranscodePresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("UpdateSubtitleTranscodePreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(UpdateSubtitleTranscodePresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) ListVhostSubtitleTranscodePreset(arg *ListVhostSubtitleTranscodePresetBody) (*ListVhostSubtitleTranscodePresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("ListVhostSubtitleTranscodePreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(ListVhostSubtitleTranscodePresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Live) CreateSubtitleTranscodePreset(arg *CreateSubtitleTranscodePresetBody) (*CreateSubtitleTranscodePresetRes, error) {
	body, err := marshalToJson(arg)
	if err != nil {
		return nil, err
	}

	data, _, err := c.Client.Json("CreateSubtitleTranscodePreset", url.Values{}, string(body))
	if err != nil {
		return nil, err
	}

	result := new(CreateSubtitleTranscodePresetRes)
	err = unmarshalResultInto(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
