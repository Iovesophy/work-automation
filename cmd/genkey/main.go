package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"work-automation/pkg/auth"

	"golang.org/x/crypto/ssh/terminal"
)

const (
	PublicKeyPem  = "config/automationPublicKey.pem"
	PrivateKeyPem = "/.ssh/work/automationPrivateKey.pem"
)

func CreateKey() {
	pri, pub := auth.GenerateKeyPair(2024)

	fmt.Print("Input password :")
	plain, _ := ReadPassword()

	// Encode
	passWithPub := auth.EncryptWithPublicKey([]byte(plain), pub)
	file, err := os.Create(PublicKeyPem)
	if err != nil {
		log.Fatal(err)
	}
	publickey := gob.NewEncoder(file)
	publickey.Encode(passWithPub)
	file.Close()

	var home string
	home, err = os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	file, err = os.Create(home + PrivateKeyPem)
	if err != nil {
		log.Fatal(err)
	}
	privatekey := gob.NewEncoder(file)
	privatekey.Encode(pri)
	file.Close()

	// Decode
	file, err = os.Open(home + PrivateKeyPem)
	if err != nil {
		log.Fatal(err)
	}
	priDec := gob.NewDecoder(file)
	err = priDec.Decode(&pri)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	var passwdDec []byte
	file, err = os.Open(PublicKeyPem)
	if err != nil {
		log.Fatal(err)
	}
	pubDec := gob.NewDecoder(file)
	err = pubDec.Decode(&passwdDec)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	auth.DecryptWithPrivateKey(passwdDec, pri)
	fmt.Println()
	fmt.Print("done")
}

func ReadPassword() ([]byte, error) {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	defer signal.Stop(signalChan)

	currentStates, err := terminal.GetState(int(syscall.Stdin))
	if err != nil {
		return nil, err
	}

	go func() {
		<-signalChan
		terminal.Restore(int(syscall.Stdin), currentStates)
		os.Exit(1)
	}()

	return terminal.ReadPassword(syscall.Stdin)
}

func main() {
	CreateKey()
}
