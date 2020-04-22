package runtime_test

import (
	"encoding/json"
	"testing"

	"errors"

	"github.com/gopub/gox"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCopyPlainTypes(t *testing.T) {
	type Item struct {
		Int     int
		Int8    int8
		Int16   int16
		Int32   int32
		Int64   int64
		Uint    uint
		Uint8   uint8
		Uint16  uint16
		Uint32  uint32
		Uint64  uint64
		Float32 float32
		Float64 float64
		String  string
		Bytes   []byte
	}

	t.Run("StructToPtrStruct", func(t *testing.T) {
		i1 := Item{
			Int:     1,
			Int8:    2,
			Int16:   3,
			Int32:   4,
			Int64:   5,
			Uint:    6,
			Uint8:   7,
			Uint16:  8,
			Uint32:  9,
			Uint64:  10,
			Float32: 11.1,
			Float64: 12.2,
			String:  "This is a string",
			Bytes:   []byte("This is a slice of bytes"),
		}

		i2 := Item{}
		err := gox.Copy(&i2, i1)
		assert.NoError(t, err)
		assert.Equal(t, i2, i1)
	})

	t.Run("PtrStructToPtrStruct", func(t *testing.T) {
		i1 := &Item{
			Int:     1,
			Uint:    6,
			Float32: 11.1,
			Float64: 12.2,
			String:  "This is a string",
			Bytes:   []byte("This is a slice of bytes"),
		}

		i2 := &Item{}
		err := gox.Copy(i2, i1)
		assert.NoError(t, err)
		assert.Equal(t, i2, i1)
	})

	t.Run("StructToStructError", func(t *testing.T) {
		i1 := &Item{
			Int:     1,
			Uint:    6,
			Float32: 11.1,
			Float64: 12.2,
			String:  "This is a string",
			Bytes:   []byte("This is a slice of bytes"),
		}

		i2 := Item{}
		err := gox.Copy(i2, i1)
		assert.Error(t, err)
	})

	t.Run("MapToStruct", func(t *testing.T) {
		m := map[string]interface{}{
			"Int": 1, "Uint": 2, "Float32": 3.3, "String": "s", "Bytes": []byte("bytes"),
		}
		i := &Item{}
		err := gox.Copy(i, m)
		assert.NoError(t, err)
	})
}

func TestCopyEmbeddedStruct(t *testing.T) {
	type SubItem struct {
		Int     int
		Uint    uint
		Float64 float64
		String  string
		Bytes   []byte
	}

	type Item struct {
		Int     int
		Uint    uint
		Float64 float64
		String  string
		Bytes   []byte
		SubItem SubItem
	}

	t.Run("StructToPtrStruct", func(t *testing.T) {
		i1 := &Item{
			Int:     1,
			Uint:    2,
			Float64: 3.3,
			String:  "This is a string",
			Bytes:   []byte("abc"),
			SubItem: SubItem{
				Int:     4,
				Uint:    5,
				Float64: 6.6,
				String:  "This is another string",
				Bytes:   []byte("def"),
			},
		}

		i2 := &Item{}
		err := gox.Copy(i2, i1)
		assert.NoError(t, err)
		assert.Equal(t, i2, i1)
	})
}

func TestCopyEmbeddedPtrStruct(t *testing.T) {
	type Item struct {
		Int     int
		Uint    uint
		Float64 float64
		String  string
		Bytes   []byte
		SubItem *Item
	}

	t.Run("StructToPtrStruct", func(t *testing.T) {
		i1 := &Item{
			Int:     1,
			Uint:    2,
			Float64: 3.3,
			String:  "This is a string",
			Bytes:   []byte("abc"),
			SubItem: &Item{
				Int:     4,
				Uint:    5,
				Float64: 6.6,
				String:  "This is another string",
				Bytes:   []byte("def"),
			},
		}

		i2 := &Item{}
		err := gox.Copy(i2, i1)
		assert.NoError(t, err)
		assert.Equal(t, i2, i1)
	})

	t.Run("MapToPtrStruct", func(t *testing.T) {
		m := map[string]interface{}{
			"Int":     1,
			"Uint":    2,
			"Float64": 3.3,
			"String":  "This is a string",
			"Bytes":   []byte("abc"),
			"SubItem": &Item{
				Int:     4,
				Uint:    5,
				Float64: 6.6,
				String:  "This is another string",
				Bytes:   []byte("def"),
			},
		}

		i := &Item{}
		err := gox.Copy(i, m)
		assert.NoError(t, err)
		assert.Equal(t, i.SubItem, m["SubItem"])
		jm, err := json.Marshal(m)
		require.NoError(t, err)
		ji, err := json.Marshal(i)
		require.NoError(t, err)
		assert.JSONEq(t, string(ji), string(jm))
	})
}

func TestCopyer(t *testing.T) {
	type Contact struct {
		Name        string
		PhoneNumber *gox.PhoneNumber
	}

	t.Run("Success", func(t *testing.T) {
		c := &Contact{}
		err := gox.Copy(c, map[string]interface{}{
			"Name":        "Tom",
			"PhoneNumber": "+8613800000001",
		})
		require.NoError(t, err)
		assert.Equal(t, &Contact{
			Name: "Tom",
			PhoneNumber: &gox.PhoneNumber{
				Code:   86,
				Number: 13800000001,
			},
		}, c)
	})

	t.Run("CopyString", func(t *testing.T) {
		pn := &gox.PhoneNumber{}
		err := gox.Copy(pn, "+8618600000001")
		assert.NoError(t, err)
		assert.Equal(t, pn, &gox.PhoneNumber{
			Code:   86,
			Number: 18600000001,
		})
	})
}

type Item struct {
	ID    int64
	Score int
}

func (i *Item) Validate() error {
	if i.ID < 0 {
		return errors.New("invalid id")
	}

	if i.Score < 0 {
		return errors.New("invalid score")
	}

	return nil
}

func TestValidator(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		i := &Item{}
		m := map[string]interface{}{"id": 123, "score": 100}
		err := gox.Copy(i, m)
		assert.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		i := &Item{}
		m := map[string]interface{}{"id": 123, "score": -10}
		err := gox.CopyWithNamer(i, m, gox.SnakeToCamelNamer)
		assert.Error(t, err)
	})
}
