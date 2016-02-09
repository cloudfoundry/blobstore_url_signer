package server_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/cloudfoundry/blobstore_url_signer/server"
	"github.com/cloudfoundry/blobstore_url_signer/signer/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("handlers", func() {

	var (
		fakeSigner    *fakes.FakeSigner
		serverHandler server.ServerHandlers
		resp          *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		fakeSigner = &fakes.FakeSigner{}
		serverHandler = server.NewServerHandlers(fakeSigner)
		resp = httptest.NewRecorder()
	})

	Describe("SignUrl()", func() {
		It("calls the signer to sign the url", func() {
			serverHandler.SignUrl(resp, &http.Request{})
			Expect(fakeSigner.SignCallCount()).To(Equal(1))
		})

		It("writes the signed URL back to requester", func() {
			fakeSigner.SignReturns("/link/?md5=signedurl")
			serverHandler.SignUrl(resp, &http.Request{})
			Expect(resp.Body.String()).To(Equal("/link/?md5=signedurl"))
		})
	})

})
