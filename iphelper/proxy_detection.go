package iphelper

import (
	"context"
	"strconv"
	"time"
)

// ProxyDetector 检测 IP 是否是代理，是代理返回 true.
func (ipHelper *ipipIPHelper) ProxyDetector(ip string, threshold int64) (detected bool) {
	if threshold == 0 || ipHelper.ipProxyKVDB == nil {
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
