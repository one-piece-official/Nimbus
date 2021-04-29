package iphelper

import (
	"github.com/ipipdotnet/ipdb-go"
	"github.com/one-piece-official/Nimbus/repository"
)

const (
	KeyCity    = "city_name"
	KeyCountry = "country_name"
	KeyRegion  = "region_name"
)

type IPHelper struct {
	db          *ipdb.City
	ipProxyKVDB repository.KVIface // 存储 IP 代理信息
}

func NewIPHelper(dbURL string, ipProxyKVDB repository.KVIface) *IPHelper {
	db, err := ipdb.NewCity(dbURL) // 初始化 ipdb 连接
	if err != nil {
		return nil
	}

	return &IPHelper{
		db:          db,
		ipProxyKVDB: ipProxyKVDB,
	}
}

func NewIPHelperWithIPDB(db *ipdb.City, ipProxyKVDB repository.KVIface) *IPHelper {
	return &IPHelper{
		db:          db,
		ipProxyKVDB: ipProxyKVDB,
	}
}
