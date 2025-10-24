package ipfs_gateway

type GatewayNamesType string

const (
	Gateway_LightHouse   GatewayNamesType = "lighthouse"
	Gateway_SentX        GatewayNamesType = "sentx-bcdn"
	Gateway_HashpackBcdn GatewayNamesType = "hashpack-bcdn"
	Gateway_NftStorage   GatewayNamesType = "nftstorage"
	Gateway_IpfsIo       GatewayNamesType = "ipfs.io"
)

var GatewayNamesArray = [5]GatewayNamesType{Gateway_HashpackBcdn, Gateway_IpfsIo, Gateway_LightHouse, Gateway_NftStorage, Gateway_SentX}

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

func IsValidGateway(gateway GatewayNamesType) bool {
	for _, val := range GatewayNamesArray {
		if val == gateway {
			return true
		}
	}

	return false
}
