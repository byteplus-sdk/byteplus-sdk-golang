package cdn

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ListResourceTags(t *testing.T) {
	resp, err := DefaultInstance.ListResourceTags()
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", resp)
}
