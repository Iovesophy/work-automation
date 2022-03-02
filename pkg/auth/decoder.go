package auth

import (
	"crypto/rsa"
	"encoding/gob"
	"log"
	"os"
)

type Enc struct {
	privateKey *rsa.PrivateKey
	code       []byte
}

func (e Enc) decode(path string, pri string) []byte {
	// Decode code

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	// open privatekey
	file, err := os.Open(home + pri)
	if err != nil {
		log.Fatal(err)
	}
	privateKeyPemFile := gob.NewDecoder(file)
	err = privateKeyPemFile.Decode(&e.privateKey)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// open publicKey
	file, err = os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	publicKeyPemFile := gob.NewDecoder(file)
	err = publicKeyPemFile.Decode(&e.code)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	code := DecryptWithPrivateKey(e.code, e.privateKey)
	return code
}

func Decoder(publidkeyPath string) []byte {
	e := Enc{}
	return e.decode(publidkeyPath, "/.ssh/work/automationPrivateKey.pem")
}
