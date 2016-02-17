package main

import (
	"flag"

	"github.com/cloudfoundry/blobstore_url_signer/server"
	"github.com/cloudfoundry/blobstore_url_signer/signer"
)

var (
	flagBlobstoreSecret string
)

func main() {
	flag.StringVar(&flagBlobstoreSecret, "secret", "", "The secret for signing webdav url")
	flag.Parse()

	urlSigner := signer.NewSigner(flagBlobstoreSecret)
	serverHandlers := server.NewServerHandlers(urlSigner)
	s := server.NewServer(8080, "127.0.0.1", serverHandlers)
	s.Start()
}
