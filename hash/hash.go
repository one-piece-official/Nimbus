package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

// NOTE: 优化摘要函数的写法和性能, 参考 https://wangbjun.site/2020/coding/golang/md5.html
func MD5(s string) string {
	sum := md5.Sum([]byte(s))

	return hex.EncodeToString(sum[:])
}

func SHA1(s string) string {
	sum := sha1.Sum([]byte(s))

	return hex.EncodeToString(sum[:])
}
