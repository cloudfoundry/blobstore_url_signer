package main_test

import (
	"net/http"
	"os/exec"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("main", func() {
	var result *Session

	BeforeEach(func() {
		result = runSigner()
		time.Sleep(10 * time.Millisecond)
	})

	AfterEach(func() {
		result.Kill()
	})

	It("Default server port to 8080", func() {
		_, err := http.Get("http://127.0.0.1:8080")
		Expect(err).ToNot(HaveOccurred())
	})
})

func runSigner(args ...string) *Session {
	path, err := Build("github.com/cloudfoundry/blobstore_url_signer/")
	Expect(err).NotTo(HaveOccurred())

	session, err := Start(exec.Command(path, args...), GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	return session
}
