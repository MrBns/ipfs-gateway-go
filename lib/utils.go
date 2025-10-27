package ipfs_gateway

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

// detect if []byte type is a valid text type. using http.DetectContentType()
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
It check either input is ipfs url or base64.
if it is base64 then it will decode base64 and return decoded version.
*/
func CheckIpfsUrlAndParse(ipfsUrlOrBase64 string) (string, error) {

	var decodedStr string = ipfsUrlOrBase64

	if !strings.HasPrefix(ipfsUrlOrBase64, "ipfs://") {

		dcoded, err := base64.StdEncoding.DecodeString(ipfsUrlOrBase64)
		if err != nil {
			return "", err
		}

		decodedStr = string(dcoded)

		if !strings.HasPrefix(decodedStr, "ipfs://") {
			return "", fmt.Errorf("only valid ipfs url is supported or base64 encoded string of ipfs url")
		}

	}
	return decodedStr, nil
}

/*
this function take ipfs:// url as params. and

return  (id string, path string, error)
*/
func SplitIpfsURL(url string) (string, string, error) {
	purifiedCid := url

	if val, ok := strings.CutPrefix(url, "ipfs://"); ok {
		purifiedCid = val
	}

	splitCid := strings.Split(purifiedCid, "/")

	if len(purifiedCid) < 1 {
		return "", "", fmt.Errorf("%v is not a valid ipfs cid", purifiedCid)
	}
	id := splitCid[0]

	path := ""
	if len(splitCid) >= 2 {
		path = strings.Join(splitCid[1:], "/")
	}

	return id, path, nil
}
