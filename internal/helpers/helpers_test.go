package helpers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/logingood/gooz/internal/helpers"
)

var _ = Describe("Helpers", func() {
	Describe("Test type detect and string", func() {
		Context("Stringfy different strings", func() {
			It("should detect string and stringify it", func() {
				var val interface{}
				val = "test"

				str, err := DetectTypeAndStringfy(val)
				Expect(err).ToNot(HaveOccurred())
				Expect(str).To(Equal("test"))
			})

			It("should detect float64 and stringify it", func() {
				var val interface{}
				val = float64(1)

				str, err := DetectTypeAndStringfy(val)
				Expect(err).ToNot(HaveOccurred())
				Expect(str).To(Equal("1"))
			})

			It("should detect int64 and stringify it", func() {
				var val interface{}
				val = int64(1)

				str, err := DetectTypeAndStringfy(val)
				Expect(err).ToNot(HaveOccurred())
				Expect(str).To(Equal("1"))
			})

			It("should detect int and stringify it", func() {
				var val interface{}
				val = int(1)

				str, err := DetectTypeAndStringfy(val)
				Expect(err).ToNot(HaveOccurred())
				Expect(str).To(Equal("1"))
			})

			It("should detect bool and stringify it", func() {
				var val interface{}
				val = true

				str, err := DetectTypeAndStringfy(val)
				Expect(err).ToNot(HaveOccurred())
				Expect(str).To(Equal("true"))
			})

			It("should detect unknown type and return err", func() {
				var val interface{}
				type TestType struct {
					string
				}

				val = &TestType{"test"}

				str, err := DetectTypeAndStringfy(val)
				Expect(err).To(HaveOccurred())
				Expect(str).To(Equal("failed to convert"))
				Expect(err).To(MatchError("Failed to detect type"))
			})

			It("should detect unknown type if value is nil and return err", func() {
				var val interface{}
				val = nil

				str, err := DetectTypeAndStringfy(val)
				Expect(err).To(HaveOccurred())
				Expect(str).To(Equal("failed to convert"))
				Expect(err).To(MatchError("Failed to detect type"))
			})

			It("should detect []interface{} and stringify it", func() {
				var (
					val   interface{}
					slice []interface{}
				)

				slice = []interface{}{"test", "test2"}
				val = slice

				str, err := DetectTypeAndStringfy(val)
				Expect(err).ToNot(HaveOccurred())
				Expect(str).To(Equal("test\ntest2"))
			})
		})
	})
})
