package signer_test

import (
	"github.com/cloudfoundry/blobstore_url_signer/signer"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Signer", func() {

	Context("Sign", func() {

		It("removes all =", func() {
		})

		It("returns a signed URL", func() {
			expires := "2147483647"
			secret := "secret"
			path := "/s/link"
			clientIp := "127.0.0.1"

			signer := signer.NewSigner(expires, secret, path, clientIp)
			signedUrl := signer.Sign()
			Expect(signedUrl).To(Equal("/s/link?md5=_e4Nc3iduzkWRm01TBBNYw&expires=2147483647"))
		})
	})

	Context("SanitizeString", func() {
		It("replaces '/' with '_'", func() {
			sanitizedString := signer.SanitizeString("i am /a /string")
			Expect(sanitizedString).To(Equal("i am _a _string"))
		})

		It("replaces '+' with '-'", func() {
			sanitizedString := signer.SanitizeString("i am +a +string")
			Expect(sanitizedString).To(Equal("i am -a -string"))
		})

		It("removes '='", func() {
			sanitizedString := signer.SanitizeString("i am =a =string")
			Expect(sanitizedString).To(Equal("i am a string"))
		})
	})
})
