package main_test

import (
	"net/http"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("main", func() {
	It("Default server port to 8080", func() {
		result := runSigner()

		_, err := http.Get("http://127.0.0.1:8080")
		Expect(err).ToNot(HaveOccurred())

		result.Kill()
	})
})

func runSigner(args ...string) *Session {
	path, err := Build("github.com/cloudfoundry/blobstore_url_signer/")
	Expect(err).NotTo(HaveOccurred())

	session, err := Start(exec.Command(path, args...), GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	return session
}
