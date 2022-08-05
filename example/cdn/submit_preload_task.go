package cdn

import (
	"fmt"
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
	"testing"
)

func SubmitPreloadTask(t *testing.T) {
	resp, err := DefaultInstance.SubmitPreloadTask(&cdn.SubmitPreloadTaskRequest{
		Urls: fmt.Sprintf("http://%s/1.txt", exampleDomain),
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.TaskID)
}
