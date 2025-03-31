package password

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// 安全参数配置（根据服务器性能调整）
const (
	time    = 1         // 迭代次数
	memory  = 64 * 1024 // 64MB 内存占用
	threads = 2         // 并行线程数
	keyLen  = 32        // 生成密钥长度
	saltLen = 16        // 盐值长度
)

type Argon2Hasher struct{}

func NewArgon2Hasher() *Argon2Hasher {
	return &Argon2Hasher{}
}

// 生成哈希（在注册/修改密码时调用）
func (h *Argon2Hasher) Generate(plain string) (string, error) {
	// 生成随机盐
	salt := make([]byte, saltLen)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("生成盐失败: %w", err)
	}

	// 生成Argon2id哈希
	hash := argon2.IDKey([]byte(plain), salt, time, memory, threads, keyLen)

	// 编码存储字符串格式：$算法$参数$盐$哈希
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	return fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		memory,
		time,
		threads,
		b64Salt,
		b64Hash,
	), nil
}

// 验证密码（在登录时调用）
func (h *Argon2Hasher) Compare(hashed, plain string) bool {
	// 解析存储的哈希字符串
	parts := strings.Split(hashed, "$")
	if len(parts) != 7 || parts[1] != "argon2id" {
		return false
	}
	var memory, time, threads uint32
	// 解码参数
	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &time, &threads)
	if err != nil {
		return false
	}

	// 解码盐和原始哈希
	salt, _ := base64.RawStdEncoding.DecodeString(parts[4])
	originalHash, _ := base64.RawStdEncoding.DecodeString(parts[5])

	// 重新计算哈希
	newHash := argon2.IDKey([]byte(plain), salt, time, memory, uint8(threads), keyLen)
	return subtle.ConstantTimeCompare(originalHash, newHash) == 1
}
