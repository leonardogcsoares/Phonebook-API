package config

import (
	"crypto/rsa"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
)

// Options represents the configuration options
type Options struct {
	JWTExpirationDelta int64
	PublicKeyPath      string
	PrivateKeyPath     string
	// NOTE or substitute for db_string
	VerifyKey *rsa.PublicKey
	SignKey   *rsa.PrivateKey
}

var options *Options

// New returns the instance for the configuration options
func New() (*Options, error) {
	if options != nil {
		return options, nil
	}

	options := &Options{
		JWTExpirationDelta: 2, // 2 hours
		PublicKeyPath:      "keys/app.rsa.pub",
		PrivateKeyPath:     "keys/app.rsa",
	}
	err := options.readKeyFiles()

	return options, err
}

func (o *Options) readKeyFiles() error {
	var err error
	var signBytes, verifyBytes []byte

	if signBytes, err = ioutil.ReadFile(o.PrivateKeyPath); err != nil {
		return err
	}
	if o.SignKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes); err != nil {
		return err
	}
	if verifyBytes, err = ioutil.ReadFile(o.PublicKeyPath); err != nil {
		return err
	}
	if o.VerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes); err != nil {
		return err
	}

	return nil

}
