package live_v20200801

type ListStorageSpaceBody struct {

	// REQUIRED
	PageNum int32 `json:"PageNum"`

	// REQUIRED
	PageSize int32 `json:"PageSize"`
}

type ListStorageSpaceDetailBody struct {
	StorageSpaceList []*string `json:"StorageSpaceList,omitempty"`
}

type ListStorageSpaceDetailRes struct {

	// REQUIRED
	ResponseMetadata ListStorageSpaceDetailResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result ListStorageSpaceDetailResResult `json:"Result"`
}

type ListStorageSpaceDetailResResponseMetadata struct {

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

type ListStorageSpaceDetailResResult struct {

	// REQUIRED
	StorageSpaceListDetail []ListStorageSpaceDetailResResultStorageSpaceListDetailItem `json:"StorageSpaceListDetail"`
}

type ListStorageSpaceDetailResResultStorageSpaceListDetailItem struct {
	AccountID    *string `json:"AccountID,omitempty"`
	CnameDomain  *string `json:"CnameDomain,omitempty"`
	Name         *string `json:"Name,omitempty"`
	PublicDomain *string `json:"PublicDomain,omitempty"`
	TTL          *int32  `json:"TTL,omitempty"`
}

type ListStorageSpaceRes struct {

	// REQUIRED
	ResponseMetadata ListStorageSpaceResResponseMetadata `json:"ResponseMetadata"`

	// REQUIRED
	Result ListStorageSpaceResResult `json:"Result"`
}

type ListStorageSpaceResResponseMetadata struct {

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

type ListStorageSpaceResResult struct {

	// REQUIRED
	CurPage int32 `json:"CurPage"`

	// REQUIRED
	StorageSpace []string `json:"StorageSpace"`

	// REQUIRED
	Total int32 `json:"Total"`
}
type ListStorageSpace struct{}
type ListStorageSpaceDetail struct{}
type ListStorageSpaceDetailQuery struct{}
type ListStorageSpaceQuery struct{}
type ListStorageSpaceReq struct {
	*ListStorageSpaceQuery
	*ListStorageSpaceBody
}
type ListStorageSpaceDetailReq struct {
	*ListStorageSpaceDetailQuery
	*ListStorageSpaceDetailBody
}
