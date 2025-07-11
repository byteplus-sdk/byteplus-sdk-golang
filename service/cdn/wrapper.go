package cdn

func (s *CDN) AddCdnDomain(dto *AddCdnDomainRequest) (responseBody *AddCdnDomainResponse, err error) {
	responseBody = new(AddCdnDomainResponse)
	if err = s.post("AddCdnDomain", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) StartCdnDomain(dto *StartCdnDomainRequest) (responseBody *StartCdnDomainResponse, err error) {
	responseBody = new(StartCdnDomainResponse)
	if err = s.post("StartCdnDomain", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) StopCdnDomain(dto *StopCdnDomainRequest) (responseBody *StopCdnDomainResponse, err error) {
	responseBody = new(StopCdnDomainResponse)
	if err = s.post("StopCdnDomain", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DeleteCdnDomain(dto *DeleteCdnDomainRequest) (responseBody *DeleteCdnDomainResponse, err error) {
	responseBody = new(DeleteCdnDomainResponse)
	if err = s.post("DeleteCdnDomain", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) ListCdnDomains(dto *ListCdnDomainsRequest) (responseBody *ListCdnDomainsResponse, err error) {
	responseBody = new(ListCdnDomainsResponse)
	if err = s.post("ListCdnDomains", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnConfig(dto *DescribeCdnConfigRequest) (responseBody *DescribeCdnConfigResponse, err error) {
	responseBody = new(DescribeCdnConfigResponse)
	if err = s.post("DescribeCdnConfig", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) UpdateCdnConfig(dto *UpdateCdnConfigRequest) (responseBody *UpdateCdnConfigResponse, err error) {
	responseBody = new(UpdateCdnConfigResponse)
	if err = s.post("UpdateCdnConfig", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnData(dto *DescribeCdnDataRequest) (responseBody *DescribeCdnDataResponse, err error) {
	responseBody = new(DescribeCdnDataResponse)
	if err = s.post("DescribeCdnData", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeEdgeNrtDataSummary(dto *DescribeEdgeNrtDataSummaryRequest) (responseBody *DescribeEdgeNrtDataSummaryResponse, err error) {
	responseBody = new(DescribeEdgeNrtDataSummaryResponse)
	if err = s.post("DescribeEdgeNrtDataSummary", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnOriginData(dto *DescribeCdnOriginDataRequest) (responseBody *DescribeCdnOriginDataResponse, err error) {
	responseBody = new(DescribeCdnOriginDataResponse)
	if err = s.post("DescribeCdnOriginData", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeOriginNrtDataSummary(dto *DescribeOriginNrtDataSummaryRequest) (responseBody *DescribeOriginNrtDataSummaryResponse, err error) {
	responseBody = new(DescribeOriginNrtDataSummaryResponse)
	if err = s.post("DescribeOriginNrtDataSummary", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnDataDetail(dto *DescribeCdnDataDetailRequest) (responseBody *DescribeCdnDataDetailResponse, err error) {
	responseBody = new(DescribeCdnDataDetailResponse)
	if err = s.post("DescribeCdnDataDetail", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeDistrictIspData(dto *DescribeDistrictIspDataRequest) (responseBody *DescribeDistrictIspDataResponse, err error) {
	responseBody = new(DescribeDistrictIspDataResponse)
	if err = s.post("DescribeDistrictIspData", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeEdgeStatisticalData(dto *DescribeEdgeStatisticalDataRequest) (responseBody *DescribeEdgeStatisticalDataResponse, err error) {
	responseBody = new(DescribeEdgeStatisticalDataResponse)
	if err = s.post("DescribeEdgeStatisticalData", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeEdgeTopNrtData(dto *DescribeEdgeTopNrtDataRequest) (responseBody *DescribeEdgeTopNrtDataResponse, err error) {
	responseBody = new(DescribeEdgeTopNrtDataResponse)
	if err = s.post("DescribeEdgeTopNrtData", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeOriginTopNrtData(dto *DescribeOriginTopNrtDataRequest) (responseBody *DescribeOriginTopNrtDataResponse, err error) {
	responseBody = new(DescribeOriginTopNrtDataResponse)
	if err = s.post("DescribeOriginTopNrtData", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeEdgeTopStatusCode(dto *DescribeEdgeTopStatusCodeRequest) (responseBody *DescribeEdgeTopStatusCodeResponse, err error) {
	responseBody = new(DescribeEdgeTopStatusCodeResponse)
	if err = s.post("DescribeEdgeTopStatusCode", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeOriginTopStatusCode(dto *DescribeOriginTopStatusCodeRequest) (responseBody *DescribeOriginTopStatusCodeResponse, err error) {
	responseBody = new(DescribeOriginTopStatusCodeResponse)
	if err = s.post("DescribeOriginTopStatusCode", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeEdgeTopStatisticalData(dto *DescribeEdgeTopStatisticalDataRequest) (responseBody *DescribeEdgeTopStatisticalDataResponse, err error) {
	responseBody = new(DescribeEdgeTopStatisticalDataResponse)
	if err = s.post("DescribeEdgeTopStatisticalData", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnRegionAndIsp(dto *DescribeCdnRegionAndIspRequest) (responseBody *DescribeCdnRegionAndIspResponse, err error) {
	responseBody = new(DescribeCdnRegionAndIspResponse)
	if err = s.post("DescribeCdnRegionAndIsp", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnService() (responseBody *DescribeCdnServiceResponse, err error) {
	responseBody = new(DescribeCdnServiceResponse)
	if err = s.post("DescribeCdnService", nil, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeAccountingData(dto *DescribeAccountingDataRequest) (responseBody *DescribeAccountingDataResponse, err error) {
	responseBody = new(DescribeAccountingDataResponse)
	if err = s.post("DescribeAccountingData", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) SubmitRefreshTask(dto *SubmitRefreshTaskRequest) (responseBody *SubmitRefreshTaskResponse, err error) {
	responseBody = new(SubmitRefreshTaskResponse)
	if err = s.post("SubmitRefreshTask", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) SubmitPreloadTask(dto *SubmitPreloadTaskRequest) (responseBody *SubmitPreloadTaskResponse, err error) {
	responseBody = new(SubmitPreloadTaskResponse)
	if err = s.post("SubmitPreloadTask", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeContentTasks(dto *DescribeContentTasksRequest) (responseBody *DescribeContentTasksResponse, err error) {
	responseBody = new(DescribeContentTasksResponse)
	if err = s.post("DescribeContentTasks", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeContentQuota() (responseBody *DescribeContentQuotaResponse, err error) {
	responseBody = new(DescribeContentQuotaResponse)
	if err = s.post("DescribeContentQuota", nil, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) SubmitBlockTask(dto *SubmitBlockTaskRequest) (responseBody *SubmitBlockTaskResponse, err error) {
	responseBody = new(SubmitBlockTaskResponse)
	if err = s.post("SubmitBlockTask", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) SubmitUnblockTask(dto *SubmitUnblockTaskRequest) (responseBody *SubmitUnblockTaskResponse, err error) {
	responseBody = new(SubmitUnblockTaskResponse)
	if err = s.post("SubmitUnblockTask", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeContentBlockTasks(dto *DescribeContentBlockTasksRequest) (responseBody *DescribeContentBlockTasksResponse, err error) {
	responseBody = new(DescribeContentBlockTasksResponse)
	if err = s.post("DescribeContentBlockTasks", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnAccessLog(dto *DescribeCdnAccessLogRequest) (responseBody *DescribeCdnAccessLogResponse, err error) {
	responseBody = new(DescribeCdnAccessLogResponse)
	if err = s.post("DescribeCdnAccessLog", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeIPInfo(dto *DescribeIPInfoRequest) (responseBody *DescribeIPInfoResponse, err error) {
	responseBody = new(DescribeIPInfoResponse)
	if err = s.post("DescribeIPInfo", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeIPListInfo(dto *DescribeIPListInfoRequest) (responseBody *DescribeIPListInfoResponse, err error) {
	responseBody = new(DescribeIPListInfoResponse)
	if err = s.post("DescribeIPListInfo", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCdnUpperIp(dto *DescribeCdnUpperIpRequest) (responseBody *DescribeCdnUpperIpResponse, err error) {
	responseBody = new(DescribeCdnUpperIpResponse)
	if err = s.post("DescribeCdnUpperIp", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) AddCdnCertificate(dto *AddCdnCertificateRequest) (responseBody *AddCdnCertificateResponse, err error) {
	responseBody = new(AddCdnCertificateResponse)
	if err = s.post("AddCdnCertificate", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) ListCertInfo(dto *ListCertInfoRequest) (responseBody *ListCertInfoResponse, err error) {
	responseBody = new(ListCertInfoResponse)
	if err = s.post("ListCertInfo", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) ListCdnCertInfo(dto *ListCdnCertInfoRequest) (responseBody *ListCdnCertInfoResponse, err error) {
	responseBody = new(ListCdnCertInfoResponse)
	if err = s.post("ListCdnCertInfo", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCertConfig(dto *DescribeCertConfigRequest) (responseBody *DescribeCertConfigResponse, err error) {
	responseBody = new(DescribeCertConfigResponse)
	if err = s.post("DescribeCertConfig", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) BatchDeployCert(dto *BatchDeployCertRequest) (responseBody *BatchDeployCertResponse, err error) {
	responseBody = new(BatchDeployCertResponse)
	if err = s.post("BatchDeployCert", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DeleteCdnCertificate(dto *DeleteCdnCertificateRequest) (responseBody *DeleteCdnCertificateResponse, err error) {
	responseBody = new(DeleteCdnCertificateResponse)
	if err = s.post("DeleteCdnCertificate", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeAccountingSummary(dto *DescribeAccountingSummaryRequest) (responseBody *DescribeAccountingSummaryResponse, err error) {
	responseBody = new(DescribeAccountingSummaryResponse)
	if err = s.post("DescribeAccountingSummary", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeTemplates(dto *DescribeTemplatesRequest) (responseBody *DescribeTemplatesResponse, err error) {
	responseBody = new(DescribeTemplatesResponse)
	if err = s.post("DescribeTemplates", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeServiceTemplate(dto *DescribeServiceTemplateRequest) (responseBody *DescribeServiceTemplateResponse, err error) {
	responseBody = new(DescribeServiceTemplateResponse)
	if err = s.post("DescribeServiceTemplate", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeCipherTemplate(dto *DescribeCipherTemplateRequest) (responseBody *DescribeCipherTemplateResponse, err error) {
	responseBody = new(DescribeCipherTemplateResponse)
	if err = s.post("DescribeCipherTemplate", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) CreateCipherTemplate(dto *CreateCipherTemplateRequest) (responseBody *CreateCipherTemplateResponse, err error) {
	responseBody = new(CreateCipherTemplateResponse)
	if err = s.post("CreateCipherTemplate", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) UpdateServiceTemplate(dto *UpdateServiceTemplateRequest) (responseBody *UpdateServiceTemplateResponse, err error) {
	responseBody = new(UpdateServiceTemplateResponse)
	if err = s.post("UpdateServiceTemplate", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) UpdateCipherTemplate(dto *UpdateCipherTemplateRequest) (responseBody *UpdateCipherTemplateResponse, err error) {
	responseBody = new(UpdateCipherTemplateResponse)
	if err = s.post("UpdateCipherTemplate", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DuplicateTemplate(dto *DuplicateTemplateRequest) (responseBody *DuplicateTemplateResponse, err error) {
	responseBody = new(DuplicateTemplateResponse)
	if err = s.post("DuplicateTemplate", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) LockTemplate(dto *LockTemplateRequest) (responseBody *LockTemplateResponse, err error) {
	responseBody = new(LockTemplateResponse)
	if err = s.post("LockTemplate", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DeleteTemplate(dto *DeleteTemplateRequest) (responseBody *DeleteTemplateResponse, err error) {
	responseBody = new(DeleteTemplateResponse)
	if err = s.post("DeleteTemplate", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeTemplateDomains(dto *DescribeTemplateDomainsRequest) (responseBody *DescribeTemplateDomainsResponse, err error) {
	responseBody = new(DescribeTemplateDomainsResponse)
	if err = s.post("DescribeTemplateDomains", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) AddTemplateDomain(dto *AddTemplateDomainRequest) (responseBody *AddTemplateDomainResponse, err error) {
	responseBody = new(AddTemplateDomainResponse)
	if err = s.post("AddTemplateDomain", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) UpdateTemplateDomain(dto *UpdateTemplateDomainRequest) (responseBody *UpdateTemplateDomainResponse, err error) {
	responseBody = new(UpdateTemplateDomainResponse)
	if err = s.post("UpdateTemplateDomain", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) CreateServiceTemplate(dto *CreateServiceTemplateRequest) (responseBody *CreateServiceTemplateResponse, err error) {
	responseBody = new(CreateServiceTemplateResponse)
	if err = s.post("CreateServiceTemplate", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeDistrictData(dto *DescribeDistrictDataRequest) (responseBody *DescribeDistrictDataResponse, err error) {
	responseBody = new(DescribeDistrictDataResponse)
	if err = s.post("DescribeDistrictData", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeEdgeData(dto *DescribeEdgeDataRequest) (responseBody *DescribeEdgeDataResponse, err error) {
	responseBody = new(DescribeEdgeDataResponse)
	if err = s.post("DescribeEdgeData", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeDistrictSummary(dto *DescribeDistrictSummaryRequest) (responseBody *DescribeDistrictSummaryResponse, err error) {
	responseBody = new(DescribeDistrictSummaryResponse)
	if err = s.post("DescribeDistrictSummary", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeEdgeSummary(dto *DescribeEdgeSummaryRequest) (responseBody *DescribeEdgeSummaryResponse, err error) {
	responseBody = new(DescribeEdgeSummaryResponse)
	if err = s.post("DescribeEdgeSummary", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeOriginData(dto *DescribeOriginDataRequest) (responseBody *DescribeOriginDataResponse, err error) {
	responseBody = new(DescribeOriginDataResponse)
	if err = s.post("DescribeOriginData", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeOriginSummary(dto *DescribeOriginSummaryRequest) (responseBody *DescribeOriginSummaryResponse, err error) {
	responseBody = new(DescribeOriginSummaryResponse)
	if err = s.post("DescribeOriginSummary", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeUserData(dto *DescribeUserDataRequest) (responseBody *DescribeUserDataResponse, err error) {
	responseBody = new(DescribeUserDataResponse)
	if err = s.post("DescribeUserData", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeDistrictRanking(dto *DescribeDistrictRankingRequest) (responseBody *DescribeDistrictRankingResponse, err error) {
	responseBody = new(DescribeDistrictRankingResponse)
	if err = s.post("DescribeDistrictRanking", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeEdgeRanking(dto *DescribeEdgeRankingRequest) (responseBody *DescribeEdgeRankingResponse, err error) {
	responseBody = new(DescribeEdgeRankingResponse)
	if err = s.post("DescribeEdgeRanking", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeOriginRanking(dto *DescribeOriginRankingRequest) (responseBody *DescribeOriginRankingResponse, err error) {
	responseBody = new(DescribeOriginRankingResponse)
	if err = s.post("DescribeOriginRanking", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeEdgeStatusCodeRanking(dto *DescribeEdgeStatusCodeRankingRequest) (responseBody *DescribeEdgeStatusCodeRankingResponse, err error) {
	responseBody = new(DescribeEdgeStatusCodeRankingResponse)
	if err = s.post("DescribeEdgeStatusCodeRanking", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeOriginStatusCodeRanking(dto *DescribeOriginStatusCodeRankingRequest) (responseBody *DescribeOriginStatusCodeRankingResponse, err error) {
	responseBody = new(DescribeOriginStatusCodeRankingResponse)
	if err = s.post("DescribeOriginStatusCodeRanking", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeStatisticalRanking(dto *DescribeStatisticalRankingRequest) (responseBody *DescribeStatisticalRankingResponse, err error) {
	responseBody = new(DescribeStatisticalRankingResponse)
	if err = s.post("DescribeStatisticalRanking", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) BatchUpdateCdnConfig(dto *BatchUpdateCdnConfigRequest) (responseBody *BatchUpdateCdnConfigResponse, err error) {
	responseBody = new(BatchUpdateCdnConfigResponse)
	if err = s.post("BatchUpdateCdnConfig", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) AddCertificate(dto *AddCertificateRequest) (responseBody *AddCertificateResponse, err error) {
	responseBody = new(AddCertificateResponse)
	if err = s.post("AddCertificate", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DeleteUsageReport(dto *DeleteUsageReportRequest) (responseBody *DeleteUsageReportResponse, err error) {
	responseBody = new(DeleteUsageReportResponse)
	if err = s.post("DeleteUsageReport", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) CreateUsageReport(dto *CreateUsageReportRequest) (responseBody *CreateUsageReportResponse, err error) {
	responseBody = new(CreateUsageReportResponse)
	if err = s.post("CreateUsageReport", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) ListUsageReports(dto *ListUsageReportsRequest) (responseBody *ListUsageReportsResponse, err error) {
	responseBody = new(ListUsageReportsResponse)
	if err = s.post("ListUsageReports", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeSharedConfig(dto *DescribeSharedConfigRequest) (responseBody *DescribeSharedConfigResponse, err error) {
	responseBody = new(DescribeSharedConfigResponse)
	if err = s.post("DescribeSharedConfig", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) ListSharedConfig(dto *ListSharedConfigRequest) (responseBody *ListSharedConfigResponse, err error) {
	responseBody = new(ListSharedConfigResponse)
	if err = s.post("ListSharedConfig", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DeleteSharedConfig(dto *DeleteSharedConfigRequest) (responseBody *DeleteSharedConfigResponse, err error) {
	responseBody = new(DeleteSharedConfigResponse)
	if err = s.post("DeleteSharedConfig", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) UpdateSharedConfig(dto *UpdateSharedConfigRequest) (responseBody *UpdateSharedConfigResponse, err error) {
	responseBody = new(UpdateSharedConfigResponse)
	if err = s.post("UpdateSharedConfig", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) TagResources(dto *TagResourcesRequest) (responseBody *TagResourcesResponse, err error) {
	responseBody = new(TagResourcesResponse)
	if err = s.post("TagResources", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) UntagResources(dto *UntagResourcesRequest) (responseBody *UntagResourcesResponse, err error) {
	responseBody = new(UntagResourcesResponse)
	if err = s.post("UntagResources", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) ReleaseTemplate(dto *ReleaseTemplateRequest) (responseBody *ReleaseTemplateResponse, err error) {
	responseBody = new(ReleaseTemplateResponse)
	if err = s.post("ReleaseTemplate", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) CreateRuleEngineTemplate(dto *CreateRuleEngineTemplateRequest) (responseBody *CreateRuleEngineTemplateResponse, err error) {
	responseBody = new(CreateRuleEngineTemplateResponse)
	if err = s.post("CreateRuleEngineTemplate", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) UpdateRuleEngineTemplate(dto *UpdateRuleEngineTemplateRequest) (responseBody *UpdateRuleEngineTemplateResponse, err error) {
	responseBody = new(UpdateRuleEngineTemplateResponse)
	if err = s.post("UpdateRuleEngineTemplate", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}

func (s *CDN) DescribeRuleEngineTemplate(dto *DescribeRuleEngineTemplateRequest) (responseBody *DescribeRuleEngineTemplateResponse, err error) {
	responseBody = new(DescribeRuleEngineTemplateResponse)
	if err = s.post("DescribeRuleEngineTemplate", dto, responseBody); err != nil {
		return
	}
	if err = validateResponse(responseBody.ResponseMetadata); err != nil {
		return
	}
	return
}
