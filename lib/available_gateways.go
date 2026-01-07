package ipfs_gateway

type GatewayNamesType string

const (
	Gateway_LightHouse   GatewayNamesType = "lighthouse"
	Gateway_SentX        GatewayNamesType = "sentx-bcdn"
	Gateway_HashpackBcdn GatewayNamesType = "hashpack-bcdn"
	Gateway_NftStorage   GatewayNamesType = "nftstorage"
	Gateway_IpfsIo       GatewayNamesType = "ipfs.io"
	Gateway_Filebase     GatewayNamesType = "filebase"
)

var GatewayNamesArray = [5]GatewayNamesType{Gateway_HashpackBcdn, Gateway_IpfsIo, Gateway_LightHouse, Gateway_NftStorage, Gateway_SentX}

/*
Available gateways
  - lighthouse
  - sentx-bcdn
  - nftstorage
  - ipfs.io
  - lighthouse
  - filebase
*/
func GetGatewayByName(gateway_name GatewayNamesType) IPFS_Gateway {
	switch gateway_name {
	case Gateway_LightHouse:
		return NewLightHousGateway()
	case Gateway_SentX:
		return NewSentxBCdnGateway()
	case Gateway_NftStorage:
		return NewNftStorageGateway()
	case Gateway_IpfsIo:
		return NewIpfsIoGateway()
	case Gateway_Filebase:
		return NewFilebaseGateway()
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
