package server

import (
	"net/http"
	"net/url"

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
	// userName, password, _ := r.BasicAuth()

	u, _ := url.Parse(r.URL.String())
	queries, _ := url.ParseQuery(u.RawQuery)
	expirationDate := queries["expire"][0]
	path := queries["path"][0]
	prefix := queries["prefix"][0]

	redirectUrl := h.signer.Sign(expirationDate, prefix, path)
	http.Redirect(w, r, redirectUrl, http.StatusFound)
}
