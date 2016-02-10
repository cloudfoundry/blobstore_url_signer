package main

import (
	"flag"

	"github.com/cloudfoundry/blobstore_url_signer/server"
	"github.com/cloudfoundry/blobstore_url_signer/signer"
)

var (
	flagBlobstoreSecret   string
	flagBlobstoreUser     string
	flagBlobstorePassword string
)

func main() {
	flag.StringVar(&flagBlobstoreSecret, "secret", "", "The secret for signing webdav url")
	flag.StringVar(&flagBlobstoreUser, "user", "", "The username for identifying internal client")
	flag.StringVar(&flagBlobstorePassword, "password", "", "The password for identifying internal client")
	flag.Parse()

	urlSigner := signer.NewSigner(flagBlobstoreSecret)
	serverHandlers := server.NewServerHandlers(urlSigner, flagBlobstoreUser, flagBlobstorePassword)
	s := server.NewServer(8080, "127.0.0.1", serverHandlers)
	s.Start()
}
