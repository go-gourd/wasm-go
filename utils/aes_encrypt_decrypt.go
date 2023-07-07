package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// EncryptAES 加密
func EncryptAES(key []byte, plaintext string) (string, error) {

	// 创建一个AES加密块
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// 创建一个加密模式
	mode := cipher.NewCBCEncrypter(block, key[:aes.BlockSize])

	plaintextBytes := []byte(plaintext)

	// 对明文进行填充
	plaintextBytes = PKCS7Padding(plaintextBytes, aes.BlockSize)

	// 创建一个字节数组用于存储加密后的密文
	ciphertext := make([]byte, len(plaintextBytes))

	// 加密明文
	mode.CryptBlocks(ciphertext, plaintextBytes)

	// 将密文转换为十六进制字符串
	//ciphertextHex := hex.EncodeToString(ciphertext)

	ciphertextHex := base64.StdEncoding.EncodeToString(ciphertext)

	return ciphertextHex, nil
}

// DecryptAES 解密
func DecryptAES(key []byte, ciphertextHex string) (string, error) {

	// 将密文转换为字节数组
	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextHex)
	if err != nil {
		panic(err)
	}

	// 创建一个AES解密块
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// 创建一个解密模式
	mode := cipher.NewCBCDecrypter(block, key[:aes.BlockSize])

	// 创建一个字节数组用于存储解密后的明文
	plaintext := make([]byte, len(ciphertext))

	// 解密密文
	mode.CryptBlocks(plaintext, ciphertext)

	// 去除填充
	plaintext = PKCS7UnPadding(plaintext)

	return string(plaintext), nil
}

// PKCS7Padding 对明文进行填充
func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// PKCS7UnPadding 去除填充
func PKCS7UnPadding(data []byte) []byte {
	padding := int(data[len(data)-1])
	return data[:len(data)-padding]
}
