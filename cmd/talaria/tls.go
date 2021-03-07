package main

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"time"
)

func TLSFromFiles(certpath, keypath string) (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(certpath, keypath)
	if err != nil {
		return nil, err
	}
	return &tls.Config{Certificates: []tls.Certificate{cert}}, nil
}

func TLSFromScratch(domain string) (*tls.Config, error) {
	_, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	var serialNumber *big.Int
	{
		var err error
		serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
		serialNumber, err = rand.Int(rand.Reader, serialNumberLimit)
		if err != nil {
			return nil, err
		}
	}

	cert := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{`talaria`},
		},
		DNSNames: []string{domain},

		NotBefore: time.Now(),
		NotAfter:  time.Now().AddDate(10, 0, 0),

		IsCA:                  true,
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &cert, &cert, priv.Public().(ed25519.PublicKey), priv)
	if err != nil {
		return nil, err
	}

	var certPem, keyPem bytes.Buffer
	{
		err := pem.Encode(&certPem, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
		if err != nil {
			return nil, err
		}
	}
	{
		privBytes, err := x509.MarshalPKCS8PrivateKey(priv)
		if err != nil {
			return nil, err
		}

		err = pem.Encode(&keyPem, &pem.Block{Type: "PRIVATE KEY", Bytes: privBytes})
		if err != nil {
			return nil, err
		}
	}

	tlsCert, err := tls.X509KeyPair(certPem.Bytes(), keyPem.Bytes())
	if err != nil {
		return nil, err
	}
	return &tls.Config{Certificates: []tls.Certificate{tlsCert}}, nil
}
