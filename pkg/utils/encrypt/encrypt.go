package encrypt

import (
	"crypto/sha256"
	"encoding/hex"
)

func Encrypt(s string) string {
	// 创建一个新的SHA-256哈希函数
	hash := sha256.New()
	// 将密码转化为字节数组并写入哈希函数
	hash.Write([]byte(s + ".iam.encrypt"))
	// 计算哈希值
	hashed := hash.Sum(nil)
	// 将哈希值转化为十六进制字符串
	encryptedPassword := hex.EncodeToString(hashed)
	return encryptedPassword

}
