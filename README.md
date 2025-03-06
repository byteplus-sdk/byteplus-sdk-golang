# BytePlus developer SDK for Go

## Installing SDK

Suggest to use go version >= 1.14

```go get -u github.com/byteplus-sdk/byteplus-sdk-golang```

### AK/SK registration procedure
Main account and privileged sub-account may create AK/SK pair:

1. Log in to byteplus console
2. Choose "IAM" -> "Key Management"
3. You can find AK/SK pair details in the page, each account can have maximum of 2 pairs.
4. Create new token or click to view key detail.

### AK/SK Setting
- Explicitly call SetAccessKey/SetSecretKey method in Client

- Set env variable BYTEPLUS_ACCESSKEY="your ak"  BYTEPLUS_SECRETKEY = "your sk"

- json config at ～/.byteplus/config，format：
```json
{
	"ak":"Your-AK",
	"sk":"Your-SK"
}
```
Above options will take precedence in sequence, Suggest use option 1 for better trouble shooting

