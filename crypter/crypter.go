package crypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

type Encrypter struct {
	Key string
}

func NewEncrypter() *Encrypter {
	key := os.Getenv("KEY")
	if key == "" {
		panic("Не пердан параметр Key в переменные окружения")
	}
	return &Encrypter{
		Key: key,
	}
}

func (enc *Encrypter) Encrypt(plainStr []byte) []byte {
	//func aes.NewCipher(key []byte) (cipher.Block, error)
	// обьект блок - симетричный блочный шифр
	blok, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		panic(err.Error())
	}
	//func cipher.NewGCM(cipher cipher.Block) (cipher.AEAD, error)
	// aesGSM - NewGCM возвращает заданный 128-битный блочный шифр
	aesGSM, err := cipher.NewGCM(blok)
	if err != nil {
		panic(err.Error())
	}
	// func (cipher.AEAD) NonceSize() int
	// Случайное уникальное значение которое испольдля шифравания
	nonce := make([]byte, aesGSM.NonceSize())
	// func io.ReadFull(r io.Reader, buf []byte) (n int, err error)
	// Reader — это глобальный, общий экземпляр криптографически безопасного генератора случайных чисел.
	// Он возвращает количество скопированных байтов и
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		panic(err.Error())
	}
	/*
		func (cipher.AEAD) Seal(dst []byte, nonce []byte,
		plaintext []byte, additionalData []byte) []byte
		Запечатываем шифр
	*/
	return aesGSM.Seal(nonce, nonce, plainStr, nil)
}

func (enc *Encrypter) Decrypt(encryptStr []byte) []byte {

	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	//-----
	nonceSize := aesgcm.NonceSize()
	nonce, cipherText := encryptStr[:nonceSize], encryptStr[nonceSize:]

	plaintext, err := aesgcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}
