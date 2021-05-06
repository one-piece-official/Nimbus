package iphelper

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/one-piece-official/Nimbus/repository"
)

type mapIPHelper struct {
	db          map[string]map[string]string
	ipProxyKVDB repository.KVIface // 存储 IP 代理信息
}

func (ipHelper *mapIPHelper) GetAddrFromIPDB(ip string) (mp map[string]string, err error) {
	if ipHelper.db == nil {
		return mp, fmt.Errorf("no ip db, IP: %s, err: %w", ip, err)
	}

	return ipHelper.db[ip], nil
}

func (ipHelper *mapIPHelper) CheckIPAddressExistsInRegions(regions []string, ip, direction string) (bool, error) {
	mp, err := ipHelper.GetAddrFromIPDB(ip)
	// NOTE: IPv6 的数据无法定位城市，不投放.
	if err != nil || mp == nil {
		return false, fmt.Errorf("get addr from ipdb failed, ip: %s, err: %w", ip, err)
	}

	for _, target := range regions {
		if target == mp[KeyRegion] || target == mp[KeyCity] {
			return direction == "include", nil // 在白名单中，可以投放
		}
	}

	return direction == "exclude", nil // 不在黑名单中，可以投放
}

func NewIPHelperWithMap(db map[string]map[string]string, ipProxyKVDB repository.KVIface) IPHelper {
	return &mapIPHelper{
		db:          db,
		ipProxyKVDB: ipProxyKVDB,
	}
}

// ProxyDetector 检测 IP 是否是代理，是代理返回 true.
func (ipHelper *mapIPHelper) ProxyDetector(ip string, threshold int64) (detected bool) {
	if threshold == 0 || ipHelper.db == nil {
		return false
	}

	proxyTimestamp, err := ipHelper.ipProxyKVDB.Get(context.Background(), ip)
	if err != nil {
		return false
	}

	now := time.Now()
	sec := now.Unix()

	timestamp, err := strconv.Atoi(proxyTimestamp)
	if err != nil {
		return false
	}

	detected = sec-int64(timestamp) <= threshold

	return detected
}
