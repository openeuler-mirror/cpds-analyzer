package config

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
