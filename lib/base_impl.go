package ipfs_gateway

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	bns "github.com/mrbns/gokit/utility"
)

type baseGateway struct {
	isSubDomain bool
	_           [6]byte // Padding to align the next field (optional for clarity)
	gatewayUrl  string
}

func (v baseGateway) ToHttp(cid string) (string, error) {

	id, path, err := SplitIpfsURL(cid)

	if err != nil {
		return "", err
	}

	if v.isSubDomain {
		return "https://" + id + "." + v.gatewayUrl + "/" + path, nil
	} else {
		return "https://" + v.gatewayUrl + "/" + id + bns.Ternary(path != "", "/"+path, ""), nil
	}
}

func (v baseGateway) Get(cid string) (*http.Response, error) {

	url, err := v.ToHttp(cid)
	if err != nil {
		return nil, err
	}

	return http.Get(url)

}

func (v baseGateway) GetAsBytes(cid string) ([]byte, error) {
	res, err := v.Get(cid)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 199 || res.StatusCode >= 400 {
		return nil, fmt.Errorf("failed to fetch ipfs data for %v", cid)
	}

	result, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (st baseGateway) GetAndParse(cid string, v any) error {
	result, err := st.Get(cid)
	if err != nil {
		return err
	}

	err = json.NewDecoder(result.Body).Decode(v)
	return err
}
