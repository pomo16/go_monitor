package utils

import (
	"crypto/md5"
	"fmt"
)

//Md5AddSalt 添加盐值
func Md5AddSalt(base string, salt string, limit bool) string {
	tmp := fmt.Sprintf("%s|%s", salt, base)
	res := fmt.Sprintf("%x", md5.Sum([]byte(tmp)))
	if limit {
		return res[0:8]
	}
	return res
}
