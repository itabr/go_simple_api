package models

import (
    "regexp"
    "errors"
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "encoding/json"
    "fmt"
    "io"
)

func AddUser(decoder *json.Decoder) (user, error) {
	var U user
	err := decoder.Decode(&U);if err != nil {
		return U, err
	}
	err = U.validateType();if err != nil {
		return U, err
	}
	U.Token = randToken()
	U.Password = U.encryptPassword()
	return U, nil
 }

type user struct {
	Fname    string
	Lname    string
	Email    string
	Password string
	Status   string
	Token    []byte
}

func (u user) validateType() error {
    // validate Fname
    if nameRegexp:= regexp.MustCompile("[a-zA-Z]{2,}"); !nameRegexp.MatchString(u.Fname) { return errors.New("Fname INVALID FORMAT") }
    // validate Lname
    if nameRegexp:= regexp.MustCompile("[a-zA-Z]{2,}"); !nameRegexp.MatchString(u.Lname) { return errors.New("Lname INVALID FORMAT") }
    // validate email
    if emailRegexp := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"); !emailRegexp.MatchString(u.Email) { return errors.New("Email INVALID FORMAT") }
    // validate password
    if passwordRegexp := regexp.MustCompile("[a-zA-Z0-9@#$%^&+=]{8,24}"); !passwordRegexp.MatchString(u.Password) { return errors.New("Password INVALID FORMAT") }
    if u.Token != nil { return errors.New("INVALID FORMAT") }
    if u.Status != "" { return errors.New("INVALID FORMAT") }
    return nil
}

func (u user) encryptPassword() string {
	cryp := encrypt( u.Token, u.Password)
	return cryp
}

func randToken() []byte {
    b := make([]byte, 24)
    rand.Read(b)
    return b
}

// encrypt string to base64 crypto using AES
func encrypt(key []byte, text string) string {

	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		errors.New("Server Error")
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		errors.New("Server Error")
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext)
}

// decrypt from base64 to decrypted string
func decrypt(key []byte, cryptoText string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		errors.New("Server Error")
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}
