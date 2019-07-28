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
	Describe("Find the longest map and return keys slice", func() {
		Context("Working with maps", func() {
			It("should detect longest from maps and return rows and header", func() {
				inputData := make([]map[string]interface{}, 4)
				map1 := make(map[string]interface{}, 0)
				map1["k1"] = "v11"
				map1["k2"] = "v12"
				map1["k3"] = "v13"

				map2 := make(map[string]interface{}, 0)
				map2["k1"] = "v21"
				map2["k2"] = "v22"

				inputData[0] = map1
				inputData[1] = map2

				header, rows := FindTheLongestMapAndSliceKeys(inputData)
				Expect(header).To(Equal([]string{"k1", "k2", "k3"}))
				Expect(rows[0]).To(Equal([]string{"v11", "v12", "v13"}))
				Expect(rows[1]).To(Equal([]string{"v21", "v22"}))
			})
			It("should detect longest from maps and return rows and header for 1 row map", func() {
				inputData := make([]map[string]interface{}, 4)
				map1 := make(map[string]interface{}, 0)
				map1["k1"] = "v11"
				map1["k2"] = "v12"
				map1["k3"] = "v13"

				inputData[0] = map1

				header, rows := FindTheLongestMapAndSliceKeys(inputData)
				Expect(header).To(Equal([]string{"k1", "k2", "k3"}))
				Expect(rows[0]).To(Equal([]string{"v11", "v12", "v13"}))
			})
			It("should handle empty map", func() {
				inputData := make([]map[string]interface{}, 4)

				header, rows := FindTheLongestMapAndSliceKeys(inputData)
				Expect(header).To(BeNil())
				Expect(len(rows[0])).To(Equal(0))
			})
		})
	})
})
