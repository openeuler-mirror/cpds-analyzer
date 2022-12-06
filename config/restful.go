package config

import (
	"crypto/tls"

	"github.com/emicklei/go-restful"
)

var (
	certPath = "/etc/cpds/ca/cert.pem"
	keyPath  = "/etc/cpds/ca/key.pem"
)

func GetCertPath() string {
	return certPath
}

func GetKeyPath() string {
	return keyPath
}

func GetTlsConf() *tls.Config {
	tlsconf := &tls.Config{
		PreferServerCipherSuites: true,
		// Specifies that the minimum TLS version is 1.2
		MinVersion: tls.VersionTLS12,
	}

	// Use secure encryption
	tlsconf.CipherSuites = []uint16{
		tls.TLS_AES_128_GCM_SHA256,
		tls.TLS_CHACHA20_POLY1305_SHA256,
		tls.TLS_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
		tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
		tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
	}

	return tlsconf
}

func GetCors(c *restful.Container) restful.CrossOriginResourceSharing {
	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{"X-My-Header"},
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT"},
		CookiesAllowed: false,
		Container:      c,
	}

	return cors
}
