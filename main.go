package ipfs_gateway

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	bns "github.com/mrbns/gokit/utility"
)

type IPFS_Gateway interface {
	Get(cid string) (*http.Response, error)
	GetAsBytes(cid string) ([]byte, error)
	GetAndParse(cid string, v any) error
	ToHttp(cid string) (string, error)
}

type GatewayNamesType string

const (
	Gateway_LightHouse   GatewayNamesType = "lighthouse"
	Gateway_SentX        GatewayNamesType = "sentx-bcdn"
	Gateway_HashpackBcdn GatewayNamesType = "hashpack-bcdn"
	Gateway_NftStorage   GatewayNamesType = "nftstorage"
	Gateway_IpfsIo       GatewayNamesType = "ipfs.io"
)

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

type baseGateway struct {
	isSubDomain bool
	_           [7]byte // Padding to align the next field (optional for clarity)
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

	if res.StatusCode < 200 || res.StatusCode >= 400 {
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

/*
 ===============
	LightHouse.storage NFT gateway
 =================
*/

/*
Lighthouse.Storage ipfs gateway;
baseurl = gateway.lighthouse.storage/ipfs
*/
type lightHouseGateway struct {
	baseGateway
}

func NewLightHousGateway() *lightHouseGateway {
	return &lightHouseGateway{
		baseGateway{
			gatewayUrl:  "gateway.lighthouse.storage/ipfs",
			isSubDomain: false,
		},
	}
}

/*
Hashpack B-CDN Ipfs Getway

url: hashpack.b-cdn.net
*/
type hashpackBcdnGateway struct {
	baseGateway
}

func NewHashpackBcdnGateway() *hashpackBcdnGateway {
	return &hashpackBcdnGateway{
		baseGateway{
			gatewayUrl:  "hashpack.b-cdn.net/ipfs",
			isSubDomain: false,
		},
	}
}

/*
Nftstorage.link IPFS gateway

baseurl: ipfs.nftstorage.link (subdomain)
*/
type nftStorageGateway struct {
	baseGateway
}

func NewNftStorageGateway() *nftStorageGateway {
	return &nftStorageGateway{
		baseGateway{
			isSubDomain: true,
			gatewayUrl:  "ipfs.nftstorage.link",
		},
	}
}

/*
Sentx.B-CDN IPFS gateway

baseurl: https://sentx.b-cdn.net
*/
type sentxBCdnGateway struct {
	baseGateway
}

func NewSentxBCdnGateway() *sentxBCdnGateway {
	return &sentxBCdnGateway{
		baseGateway{
			isSubDomain: false,
			gatewayUrl:  "sentx.b-cdn.net",
		},
	}
}

/*
Sentx.B-CDN IPFS gateway

baseurl: https://sentx.b-cdn.net
*/
type ipfsIOGateway struct {
	baseGateway
}

func NewIpfsIoGateway() *ipfsIOGateway {
	return &ipfsIOGateway{
		baseGateway{
			isSubDomain: false,
			gatewayUrl:  "ipfs.io/ipfs",
		},
	}
}
