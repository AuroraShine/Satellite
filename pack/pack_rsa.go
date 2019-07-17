package pack

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	. "satellite/utils"
)

func RSAEncrypt(src []byte) (dest []byte, err error) {
	block, _ := pem.Decode(RSAPublicKey)
	if block == nil {
		err = errors.New("RSA Public Key Error")
		return dest, err
	}
	pi, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return dest, err
	}
	pub := pi.(*rsa.PublicKey)
	dest, err = rsa.EncryptPKCS1v15(rand.Reader, pub, src)
	if err != nil {
		return dest, err
	}
	return dest, err
}