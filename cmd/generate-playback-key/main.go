package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"io/ioutil"
	"log"
)

func main() {
	keysize := flag.Int("keysize", 4096, "the key size (2048, 4096)")
	private := flag.String("privout", "key.pem", "output private key encoded as pkcs1 into file to this location")
	public := flag.String("pubout", "key.pub", "output public key in PKCS1 format to this location")

	flag.Parse()

	priv, err := rsa.GenerateKey(rand.Reader, *keysize)
	if err != nil {
		log.Fatalf("failed to generate key: %v", err)
	}

	pemPriv := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(priv),
		},
	)

	pubasn1, err := x509.MarshalPKIXPublicKey(&priv.PublicKey)
	if err != nil {
		log.Fatalf("failed to generate public key: %v", err)
	}

	pemPub := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubasn1,
		},
	)

	if err := ioutil.WriteFile(*private, pemPriv, 0644); err != nil {
		log.Fatalf("failed to write private key: %v", err)
	}
	if err := ioutil.WriteFile(*public, pemPub, 0644); err != nil {
		log.Fatalf("failed to write pubkey: %v", err)
	}
}
