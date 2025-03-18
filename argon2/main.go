package main

import (
	"fmt"

	"golang.org/x/crypto/argon2"
)

func hashPassword(password, salt string) string {
	timeCost := uint32(3)           // 计算时间成本
	memoryCost := uint32(64 * 1024) // 内存成本（KB）
	parallelism := uint8(2)         // 并行度
	keyLen := uint32(32)            // 生成哈希的长度

	hash := argon2.IDKey([]byte(password), []byte(salt), timeCost, memoryCost, parallelism, keyLen)
	return fmt.Sprintf("%x", hash)
}

func main() {
	password := "mypassword"
	salt := "random_salt"

	hash1 := hashPassword(password, salt)
	hash2 := hashPassword(password, salt)

	fmt.Println("Hash 1:", hash1)
	fmt.Println("Hash 2:", hash2)

	if hash1 == hash2 {
		fmt.Println("哈希值相同")
	} else {
		fmt.Println("哈希值不同")
	}
}
