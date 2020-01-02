package hander

import (
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

// func Test_deletedata_error(t *testing.T) {
// 	text := deletedata("2")
// 	assert.Equal(t, "Error", text, "hello")
// }
// func Test_deletedata_success(t *testing.T) {
// 	text := deletedata("1")
// 	assert.Equal(t, "OK", text, "hello")
// }

func Test_toisbn(t *testing.T) {
	text, err := toisbn("1234567890123")
	assert.NoError(t, err, "text")
	assert.Equal(t, "123-4-56-789012-3", text, "hello") //000-0-00-00667-6
}
func Test_toisbn2(t *testing.T) {
	text, err := toisbn("3")
	assert.NoError(t, err, "text")
	assert.Equal(t, "000-0-00-000000-3", text, "hello") //000-0-00-00667-6
}
func Test_toisbn3(t *testing.T) {
	text, err := toisbn("77734666666666334272123123")
	assert.Error(t, err, "1")
	assert.EqualError(t, err, "text out of range", "check error")
	assert.Equal(t, "", text, "hello") //000-0-00-00667-6
}

func Test_setisbnformat1(t *testing.T) {
	text, err := setisbnformat("123-45-1-678901-1")
	assert.NoError(t, err, "check error")
	assert.Equal(t, "1234516789011", text, "")
}

func Test_setisbnformat2(t *testing.T) {
	_, err := setisbnformat("123-45-1-67891-1")
	assert.Error(t, err, "check error")
	assert.EqualError(t, err, "cant convert isbn format", "")
}

func Test_setisbnformat3(t *testing.T) {
	_, err := setisbnformat("123-4567891-1")
	assert.Error(t, err, "check error")
	assert.EqualError(t, err, "string lenght not match", "")
}
