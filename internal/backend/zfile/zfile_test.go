package zfile_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/logingood/gooz/internal/backend/zfile"
	"github.com/spf13/afero"
)

var _ = Describe("Zfile", func() {
	appFS := afero.NewMemMapFs()

	Describe("Filesystem tests", func() {
		BeforeEach(func() {
			afero.WriteFile(appFS, "/tmp/orgs.json", []byte("orgs"), 0644)
			afero.WriteFile(appFS, "/tmp/users.json", []byte("users"), 0644)
			afero.WriteFile(appFS, "/tmp/tickets.json", []byte("tickets"), 0644)
		})

		Context("Backend should initialize", func() {
			It("Should initialize backend with filepath and filesystem", func() {
				backend := New(appFS, "/tmp/test")

				Expect(backend.FilePath).To(Equal("/tmp/test"))
				Expect(backend.GoozFS).ToNot(BeNil())
			})
		})

		Context("Files do not exist", func() {
			It("It should handle non-existing org file", func() {
				backend := New(appFS, "/tmp1231/nonexistent.json")
				err := backend.Open()

				Expect(err).ToNot(BeNil())
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError("open /tmp1231/nonexistent.json: file does not exist"))
				Expect(backend.FD).To(BeNil())
			})
		})

		Context("Handle Close of File", func() {
			It("It should handle non-existing org file", func() {
				backend := New(appFS, "/tmp/orgs.json")
				err := backend.Open()
				Expect(err).ToNot(HaveOccurred())
				Expect(backend.FD).ToNot(BeNil())

				// Shouldn't panic
				err = backend.Close()
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("handle close of un-opened file", func() {
			It("it should handle non-existing org file", func() {
				backend := New(appFS, "/tmp/orgs.json")
				// Shouldn't panic
				err := backend.Close()
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError("File descriptor is nil"))
			})
		})
	})

	Describe("Reading json", func() {
		Context("Read json from file", func() {
			It("it should handle closed FD when reads json", func() {
				backend := New(appFS, "/tmp/nonexistent.json")
				err := backend.Open()

				stringMap, err := backend.Read()
				Expect(err).To(MatchError("File descriptor is nil"))
				Expect(err).To(HaveOccurred())
				Expect(stringMap).To(BeNil())
			})

			It("it should handle malformed json", func() {
				afero.WriteFile(appFS, "/tmp/orgs.json", []byte("malformed"), 0644)
				backend := New(appFS, "/tmp/orgs.json")
				err := backend.Open()
				Expect(err).ToNot(HaveOccurred())
				Expect(backend.FD).ToNot(BeNil())

				stringMap, err := backend.Read()
				Expect(err).To(HaveOccurred())
				Expect(stringMap).To(BeNil())
				Expect(err).To(MatchError("invalid character 'm' looking for beginning of value"))
			})

			It("it should handle valid json but not marshable into []map", func() {
				afero.WriteFile(appFS, "/tmp/orgs.json", []byte("{\"hello\": \"i am json\"}"), 0644)
				backend := New(appFS, "/tmp/orgs.json")
				err := backend.Open()
				Expect(err).ToNot(HaveOccurred())
				Expect(backend.FD).ToNot(BeNil())

				stringMap, err := backend.Read()
				Expect(err).To(HaveOccurred())
				Expect(stringMap).To(BeNil())
				Expect(err).To(MatchError("json: cannot unmarshal object into Go value of type []map[string]interface {}"))
			})

			It("it should read []json", func() {
				afero.WriteFile(appFS, "/tmp/orgs.json", []byte("[{\"hello\": \"i am json\"}, {\"hello\": \"and me json\"}]"), 0644)
				backend := New(appFS, "/tmp/orgs.json")
				err := backend.Open()
				Expect(err).ToNot(HaveOccurred())
				Expect(backend.FD).ToNot(BeNil())

				stringMap, err := backend.Read()
				Expect(err).ToNot(HaveOccurred())
				Expect(stringMap).ToNot(BeNil())
				Expect(stringMap[0]["hello"]).To(Equal("i am json"))
				Expect(stringMap[1]["hello"]).To(Equal("and me json"))
			})
		})
	})
})
