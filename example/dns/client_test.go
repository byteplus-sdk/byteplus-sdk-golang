package dns_test

import (
	"context"
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/service/dns"
	"github.com/stretchr/testify/assert"
)

const (
	ak = "ak"
	sk = "sk"
)

func TestTopZone(t *testing.T) {
	ctx := context.Background()
	c := dns.InitDNSBytePlusClient()
	c.SetAccessKey(ak)
	c.SetSecretKey(sk)
	respListZone, err := c.ListZones(ctx, &dns.ListZonesRequest{})
	assert.Nil(t, err)
	assert.NotNil(t, respListZone)
}
