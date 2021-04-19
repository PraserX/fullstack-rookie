package user

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

type User struct {
	Username  string
	FirstName string
	LastName  string
	password  Password
}

type Password struct {
	plaintext string
	sha256    string
}

func New(opts ...Option) *User {
	var options = &Options{}
	options.Username = ""
	options.FirstName = ""
	options.LastName = ""

	for _, opt := range opts {
		opt(options)
	}

	return &User{
		Username:  options.Username,
		FirstName: options.FirstName,
		LastName:  options.LastName,
		password:  Password{},
	}
}

func (u *User) Fullname() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

func (u *User) CreateRandomPassword(length int) string {
	const symbolTable = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789$!._"

	var plaintextPassword = ""

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		plaintextPassword += string(symbolTable[rand.Intn(len(symbolTable))])
		// Why we have to convert to string? Because string is compound of bytes.
		// Thats the reason why we have to convert one byte to string. The is no
		// operator overloading in Go. So it is possible concatenate only two
		// strings, not string and byte.
	}

	u.password.plaintext = plaintextPassword

	hash := sha256.New()
	hash.Write([]byte(plaintextPassword))
	u.password.sha256 = hex.EncodeToString(hash.Sum(nil))

	return u.password.sha256
}

// GetPasswordHash is non pointer receiver.
func (u User) GetPasswordHash() string {
	return u.password.sha256
}

// GetBadPasswordHash is non pointer receiver which trying to update our password.
// But it cannot, it is non pointer receiver (it works with copy of object/struct).
func (u User) GetBadPasswordHash() string {
	u.password = Password{plaintext: "bad", sha256: "bad"}
	return u.password.sha256
}
