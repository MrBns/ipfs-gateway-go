package ipfs_gateway

import (
	"fmt"
	"net/http"
)

/*
This Method will must resolve the ipfs data. via fallback.

strategy.
  - It will looks data to Lighthouse.storage
  - then it will look at hashpack-bcdn
  - then it will look at sentx-bcdn*
  - then it will look at  nftstorage.link
  - at last it will look at ipfs.io
*/
func GetResponse_Must(url string) (*http.Response, error) {

	cid, err := IsValidIpfsUrlAndParse(url)

	if err != nil {
		return nil, err
	}

	// 1st: filebase
	filebaseIpfsResponse, err := GetResponseByGatewayName(cid, Gateway_Filebase)
	if err == nil {
		return filebaseIpfsResponse, nil
	}

	// 2nd: lighthouse;
	lightHouseResponse, err := GetResponseByGatewayName(cid, Gateway_LightHouse)
	if err == nil {
		return lightHouseResponse, nil
	}

	// 3rd: hashpack-bcdn
	hashpackBcdnResponse, err := GetResponseByGatewayName(cid, Gateway_HashpackBcdn)
	if err == nil {
		return hashpackBcdnResponse, nil
	}

	// 4th: Sentx-Bcdn
	sentxBcdnReponse, err := GetResponseByGatewayName(cid, Gateway_SentX)
	if err == nil {
		return sentxBcdnReponse, nil
	}

	// 5th: IPFS.io gateway
	ipfsIoResponse, err := GetResponseByGatewayName(cid, Gateway_IpfsIo)
	if err == nil {
		return ipfsIoResponse, err
	}

	return nil, fmt.Errorf("no ipfs gateway could resolve this cid %v. make sure this exist", cid)
}

func GetResponseByGatewayName(cid string, gateway_name GatewayNamesType) (*http.Response, error) {
	gateway := GetGatewayByName(gateway_name)

	res, err := gateway.Get(cid)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode >= 400 {
		return nil, fmt.Errorf("failed to fetch ipfs data for %v", cid)
	}

	res.Header.Set("F-Gateway", string(gateway_name))

	return res, nil

}
