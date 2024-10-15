package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

// PKCS7Padding 填充
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS7UnPadding 移除填充
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// AesEncrypt 加密
func AesEncrypt(data []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	origData := PKCS7Padding(data, blockSize)
	encrypted := make([]byte, len(origData))
	iv := key[:blockSize] // 使用密钥的前16个字节作为IV
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(encrypted, origData)
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// AesDecrypt 解密
func AesDecrypt(encrypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	if len(encrypted) < blockSize {
		return nil, fmt.Errorf("密文长度太短")
	}
	// 假设IV是密文的前几个字节
	iv := encrypted[:blockSize]
	encrypted = encrypted[blockSize:]
	if len(encrypted)%blockSize != 0 {
		return nil, fmt.Errorf("密文不是块大小的倍数")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(encrypted))
	mode.CryptBlocks(decrypted, encrypted)
	decrypted = PKCS7UnPadding(decrypted)
	return decrypted, nil
}
