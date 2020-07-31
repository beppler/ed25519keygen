package main

import (
	"crypto/rand"
	"encoding/pem"
	"fmt"

	"github.com/mikesmitty/edkey"
	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/ssh"
)

func main() {
	// Generate a new private/public keypair for OpenSSH
	pubKey, privKey, _ := ed25519.GenerateKey(rand.Reader)
	publicKey, _ := ssh.NewPublicKey(pubKey)

	pemKey := &pem.Block{
		Type:  "OPENSSH PRIVATE KEY",
		Bytes: edkey.MarshalED25519PrivateKey(privKey),
	}
	privateKey := pem.EncodeToMemory(pemKey)
	authorizedKey := ssh.MarshalAuthorizedKey(publicKey)

	fmt.Printf("%s\n", privateKey)
	fmt.Printf("%s\n", authorizedKey)
}
