package ipfs_gateway

import (
	"net/http"
)

type IPFS_Gateway interface {
	Get(cid string) (*http.Response, error)
	GetAsBytes(cid string) ([]byte, error)
	GetAndParse(cid string, v any) error
	ToHttp(cid string) (string, error)
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

/*
IPFS Exclusive

baseurl: https://ipfs.filebase.io/ipfs
*/
func NewFilebaseGateway() *ipfsIOGateway {
	return &ipfsIOGateway{
		baseGateway{
			isSubDomain: false,
			gatewayUrl:  "ipfs.filebase.io/ipfs",
		},
	}
}
