package hander

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_createstring_selectone1(t *testing.T) {
	var dum Article
	var err error
	var text string
	expected := `{"Id":"12","Title": "twelve","Desc": "test","Content": "Content","ISBN": "000000006676","Time": "2019-12-23T08:30:24Z","Recentupdate": "2019-12-23T08:30:24Z"}`

	json.Unmarshal([]byte(expected), &dum)
	text, err = createstring_selectone(dum)
	assert.NoError(t, err, "")
	assert.Equal(t, DB12, text, "")

}
func Test_createstring_selectone2(t *testing.T) {
	var dum Article
	var err error

	expected := `{"Id":"12","Title": "twelve","Desc": "test","Content": "Content","ISBN": "00100000a06676","Time": "2019-12-23T08:30:24Z","Recentupdate": "2019-12-23T08:30:24Z"}`

	json.Unmarshal([]byte(expected), &dum)
	_, err = createstring_selectone(dum)
	assert.EqualError(t, err, "cannot convert isbn")

}

func Test_createstring_selectone2(t *testing.T) {
	var dum Article
	var err error

	expected := `{"Id":"11","Title": "twelve","Desc": "test","Content": "Content","ISBN": "00100000a06676","Time": "2019-12-23T08:30:24Z","Recentupdate": "2019-12-23T08:30:24Z"}`

	json.Unmarshal([]byte(expected), &dum)
	_, err = createstring_selectone(dum)
	assert.EqualError(t, err, "cannot convert isbn")

}


