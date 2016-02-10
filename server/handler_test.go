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
		request       *http.Request
		err           error
	)

	BeforeEach(func() {
		fakeSigner = &fakes.FakeSigner{}
		serverHandler = server.NewServerHandlers(fakeSigner, "user", "pass")
		resp = httptest.NewRecorder()
		request, err = http.NewRequest("GET", "http://127.0.0.1:8080/sign?expire=123123&secret=topSecret&prefix=blobstore&path=1c/9a/3234-sdfs", nil)
		Expect(err).ToNot(HaveOccurred())
		request.SetBasicAuth("user", "pass")
	})

	Describe("SignUrl()", func() {
		It("calls the signer to sign the url", func() {
			serverHandler.SignUrl(resp, request)
			Expect(fakeSigner.SignCallCount()).To(Equal(1))
		})

		It("sends the signer the correct params", func() {
			serverHandler.SignUrl(resp, request)
			expire, prefix, path := fakeSigner.SignArgsForCall(0)
			Expect(expire).To(Equal("123123"))
			Expect(prefix).To(Equal("blobstore"))
			Expect(path).To(Equal("1c/9a/3234-sdfs"))
		})

		It("writes the signed URL back to requester", func() {
			fakeSigner.SignReturns("/link/?md5=signedurl")
			serverHandler.SignUrl(resp, request)
			Expect(resp.Body.String()).To(ContainSubstring("/link/?md5=signedurl"))
		})

		Context("Authentication", func() {
			Context("When succeed", func() {
				It("It returns 302", func() {
					serverHandler.SignUrl(resp, request)
					Expect(err).ToNot(HaveOccurred())

					Expect(resp.Code).To(Equal(302))
				})
			})

			Context("When user name does not match", func() {
				It("It returns 403", func() {
					request.SetBasicAuth("bad-user", "pass")

					serverHandler.SignUrl(resp, request)
					Expect(err).ToNot(HaveOccurred())

					Expect(resp.Code).To(Equal(403))
				})
			})

			Context("When password does not match", func() {
				It("It returns 403", func() {
					request.SetBasicAuth("user", "bad-pass")

					serverHandler.SignUrl(resp, request)
					Expect(err).ToNot(HaveOccurred())

					Expect(resp.Code).To(Equal(403))
				})
			})
		})
	})
})
