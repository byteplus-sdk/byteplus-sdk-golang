package cdn

import (
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
	"testing"
)

func SubmitUnblockTask(t *testing.T) {
	resp, err := DefaultInstance.SubmitUnblockTask(&cdn.SubmitUnblockTaskRequest{
		Urls: exampleDomain,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Result.TaskID)
}
