/**
 * @brief Coder-DECoder 编码解码
 * @file codec
 * @author zhangpeng
 * @version 1.0
 * @date
 */

package box_lib

import (
	"encoding/base64"
)

func EncodeByBase64(target string) string {
	strBytes := []byte(target)
	return base64.StdEncoding.EncodeToString(strBytes)
}

func DecodeByBase64(target string) (string, error) {
	result, err := base64.StdEncoding.DecodeString(target)
	if err != nil {
		return string(result), err
	}
	return string(result), nil
}
