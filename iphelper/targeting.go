package iphelper

import (
	"fmt"
)

// GetAddrFromIPDB 根据 IP 获取地理位置信息.
func (ipHelper *IPHelper) GetAddrFromIPDB(ip string) (mp map[string]string, err error) {
	if ipHelper.db == nil {
		return mp, fmt.Errorf("no ip db, IP: %s, err: %w", ip, err)
	}

	if mp, err = ipHelper.db.FindMap(ip, "CN"); err != nil {
		return mp, fmt.Errorf("run FindMap Error, IP: %s, err: %w", ip, err)
	}

	return mp, nil
}

// CheckIPAddressExistsInRegions - 检查 IP 归属省份和城市是否可以被投放广告，direction 可选值 include 白名单，exclude 黑名单.
func (ipHelper *IPHelper) CheckIPAddressExistsInRegions(regions []string, ip, direction string) (bool, error) {
	return ipHelper.checkIPAddressExistsInTargets(regions, ip, direction)
}

func (ipHelper *IPHelper) checkIPAddressExistsInTargets(targets []string, ip, direction string) (bool, error) {
	mp, err := ipHelper.GetAddrFromIPDB(ip)
	// NOTE: IPv6 的数据无法定位城市，不投放.
	if err != nil {
		return false, fmt.Errorf("get addr from ipdb failed, ip: %s, err: %w", ip, err)
	}

	for _, target := range targets {
		if target == mp[KeyRegion] || target == mp[KeyCity] {
			return direction == "include", nil // 在白名单中，可以投放
		}
	}

	return direction == "exclude", nil // 不在黑名单中，可以投放
}
