package server

import (
	"net/http"

	"github.com/cloudfoundry/blobstore_url_signer/signer"
)

type ServerHandlers interface {
	SignUrl(w http.ResponseWriter, r *http.Request)
}

type handlers struct {
	signer signer.Signer
}

func NewServerHandlers(signer signer.Signer) ServerHandlers {
	return &handlers{
		signer: signer,
	}
}

func (h *handlers) SignUrl(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(h.signer.Sign()))
}
