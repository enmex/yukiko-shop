package minio

type Config struct {
	Endpoint    string
	AccessKey   string
	SecretKey   string
	SslRequired bool
	BucketName  string
}
