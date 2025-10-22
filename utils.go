package ipfs_gateway

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

// detect []byte type
func IsValidTextType(data []byte) (bool, string) {

	mimetype := http.DetectContentType(data)

	// Check if it's text-based
	if strings.HasPrefix(mimetype, "text/") {
		return true, mimetype
	}

	// Some text formats might be detected as application/*
	textMimeTypes := []string{
		"application/json",
		"application/xml",
		"application/yaml",
		"application/toml",
	}

	for _, mimeType := range textMimeTypes {
		if strings.HasPrefix(mimetype, mimeType) {
			return true, mimetype
		}
	}

	return false, mimetype
}

/*
Available gateways
  - lighthouse
  - sentx-bcdn
  - nftstorage
  - ipfs.io
  - lighthouse
*/
func GetGatewayByName(gateway_name GatewayNamesType) IPFS_Gateway {
	switch gateway_name {
	case "lighthouse":
		return NewLightHousGateway()
	case "sentx-bcdn":
		return NewSentxBCdnGateway()
	case "nftstorage":
		return NewNftStorageGateway()
	case "ipfs.io":
		return NewIpfsIoGateway()
	default:
		return NewHashpackBcdnGateway()
	}
}

/*
It check either input is ipfs url or base64.
if it is base64 then it will decode base64 and return decoded version.
*/
func CheckIpfsUrlAndParse(ipfsUrlOrBase64 string) (string, error) {

	if !strings.HasPrefix(ipfsUrlOrBase64, "ipfs://") {
		data, err := base64.StdEncoding.DecodeString(ipfsUrlOrBase64)
		if err != nil {
			return "", err
		} else if ipfsUrlOrBase64 = string(data); !strings.HasPrefix(ipfsUrlOrBase64, "ipfs://") {
			return "", fmt.Errorf("only valid ipfs url is supported or base64 encoded string of ipfs url")
		}
	}
	return ipfsUrlOrBase64, nil
}
