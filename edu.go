/**
 * @brief Encryption Decryption Unit 加密解密
 * @file edu
 * @author zhangpeng
 * @version 1.0
 * @date
 */

package box_lib

import (
	"crypto/md5"
	"fmt"
)

func EncryptByMD5(target string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(target)))
}
