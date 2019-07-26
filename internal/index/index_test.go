package index_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/logingood/gooz/internal/index"
)

var ht *HashTable

var _ = Describe("Index", func() {
	Describe("Test hash table", func() {
		BeforeEach(func() {
			ht = &HashTable{}
		})
		Context("Test putting new element in hash table", func() {
			It("It should handle uninitialized hash table", func() {
				err := ht.Insert("test_key1", "test_value1")
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError("Hash table is not initialied, call New"))
			})

			It("It should put a Key-Value in the hash table", func() {
				ht = New()
				err := ht.Insert("test_key1", "test_value1")
				Expect(err).ToNot(HaveOccurred())

				length := ht.Length()
				Expect(length).To(Equal(1))

				err = ht.Insert("test_key2", "test_value2")

				length = ht.Length()
				Expect(length).To(Equal(2))
			})

			It("It should handle insert of empty keys", func() {
				ht = New()
				err := ht.Insert("", "test")
				Expect(err).ToNot(HaveOccurred())

				length := ht.Length()
				Expect(length).To(Equal(1))

				value := ht.Search("")
				Expect(value[0]).To(Equal("test"))

				err = ht.Insert("123", "test2")
				Expect(err).ToNot(HaveOccurred())
				value = ht.Search("123")
				Expect(value[0]).To(Equal("test2"))
			})

			It("Should lookup multiple keys on insert, use int key", func() {
				ht = New()
				err := ht.Insert(123, "test")
				Expect(err).ToNot(HaveOccurred())

				value := ht.Search(123)
				Expect(value[0]).To(Equal("test"))

				err = ht.Insert(123, "another_test")
				Expect(err).ToNot(HaveOccurred())

				err = ht.Insert(123, "extra_test")
				Expect(err).ToNot(HaveOccurred())

				value = ht.Search(123)
				Expect(len(value)).To(Equal(3))
				Expect(value[0]).To(Equal("extra_test"))
				Expect(value[1]).To(Equal("another_test"))
				Expect(value[2]).To(Equal("test"))
			})

			It("Should deal with integers and strings", func() {
				ht = New()
				// 123 int
				err := ht.Insert(123, "test")
				Expect(err).ToNot(HaveOccurred())

				// 123 is a string
				value := ht.Search("123")
				Expect(len(value)).To(Equal(0))
			})
		})

		Context("Test searching elements in hash table", func() {
			It("It should search keys in hash table", func() {
				ht = New()
				err := ht.Insert("test_key1", "test_value1")
				err = ht.Insert("test_key2", "test_value2")

				Expect(err).ToNot(HaveOccurred())

				value := ht.Search("test_key1")
				Expect(value[0]).To(Equal("test_value1"))

				value = ht.Search("test_key2")
				Expect(value[0]).To(Equal("test_value2"))
			})

			It("It should handle empty values", func() {
				ht = New()
				err := ht.Insert("test_key1", "")
				err = ht.Insert("test_key2", "")
				Expect(err).ToNot(HaveOccurred())

				value := ht.Search("test_key1")
				Expect(value[0]).To(Equal(""))

				value = ht.Search("test_key2")
				Expect(value[0]).To(Equal(""))
			})

			It("It should handle non existing keys", func() {
				ht = New()
				err := ht.Insert("test_key1", "")
				err = ht.Insert("test_key2", "")
				Expect(err).ToNot(HaveOccurred())

				value := ht.Search("non_existent_key")
				Expect(len(value)).To(Equal(0))
			})
		})

		Context("Test arbitarary structs as values", func() {
			It("It should handle non existing keys", func() {
				ht = New()
				type SomeType struct {
					A string
					B int64
					C string
				}

				err := ht.Insert("test_key1", &SomeType{A: "testA", B: 123, C: "testC"})
				err = ht.Insert("test_key2", &SomeType{A: "test2A", B: 345, C: "test2C"})

				Expect(err).ToNot(HaveOccurred())
				values := ht.Search("test_key1")
				Expect(values[0].(*SomeType).A).To(Equal("testA"))
				Expect(values[0].(*SomeType).B).To(Equal(int64(123)))
				Expect(values[0].(*SomeType).C).To(Equal("testC"))
			})
		})
	})
})
