package crypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
)

type Crypter struct {
	File     string
	Message  []byte
	Password string
}

func New(opts ...Option) *Crypter {
	var options = &Options{}
	options.File = ""

	for _, opt := range opts {
		opt(options)
	}

	return &Crypter{
		File: options.File,
	}
}

func (c *Crypter) OpenFile() error {
	var err error
	var data []byte

	if data, err = ioutil.ReadFile(c.File); err != nil {
		return fmt.Errorf("cannot open file: %v", err)
	}

	c.Message = data

	return nil
}

func (c *Crypter) SetPassword(password string) {
	c.Password = password
}

func (c *Crypter) Encrypt() (string, error) {
	var err error
	var gcm cipher.AEAD

	if gcm, err = c.getCipher(); err != nil {
		return "", fmt.Errorf("cannot get cipher")
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("cannot get nonce")
	}

	cipherText := gcm.Seal(nonce, nonce, c.Message, nil)
	ctBase64 := base64.StdEncoding.EncodeToString(cipherText)

	return ctBase64, nil
}

func (c *Crypter) Decrypt() (string, error) {
	var err error
	var gcm cipher.AEAD
	var message, encodedMessage []byte

	encodedMessage = c.Message
	if gcm, err = c.getCipher(); err != nil {
		return "", fmt.Errorf("cannot get cipher")
	}

	if message, err = base64.StdEncoding.DecodeString(string(encodedMessage)); err != nil {
		return "", fmt.Errorf("cannot decode message")
	}

	nonceSize := gcm.NonceSize()
	if len(message) < nonceSize {
		return "", fmt.Errorf("bad nonce")
	}

	nonce, ciphertext := message[:nonceSize], message[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("cannot decipher message")
	}

	return string(plaintext), nil
}

func (c *Crypter) derivateKey() ([]byte, error) {
	hash := sha256.New()
	hash.Write([]byte(c.Password))

	key := hash.Sum(nil)

	if len(key) != 32 {
		return nil, fmt.Errorf("bad key length")
	}

	return key, nil
}

func (c *Crypter) getCipher() (cipher.AEAD, error) {
	var err error
	var key []byte
	var aesCipher cipher.Block
	var gcm cipher.AEAD

	if key, err = c.derivateKey(); err != nil {
		return nil, fmt.Errorf("cannot derivate key")
	}

	if aesCipher, err = aes.NewCipher(key); err != nil {
		return nil, fmt.Errorf("cannot create cipher")
	}

	if gcm, err = cipher.NewGCM(aesCipher); err != nil {
		return nil, fmt.Errorf("cannot create AES GCM")
	}

	return gcm, nil
}
