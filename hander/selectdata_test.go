package hander

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type selectMock struct {
	Respond      string
	RespondA     Article
	RespondArray []Article
	Err          error
}

func (db selectMock) dbget1(string) (Article, error) {
	return db.RespondA, db.Err
}
func (db selectMock) Inserter(string) (string, error) {
	return db.Respond, db.Err
}
func (db selectMock) updatedata(string, string, string, string, string) (string, error) {
	return db.Respond, db.Err
}
func (db selectMock) deletedata(string) (string, error) {
	return db.Respond, db.Err
}
func (db selectMock) dbgetall() ([]Article, error) {
	return db.RespondArray, db.Err
}
func Test_select_one1(t *testing.T) {
	key := "20"
	mmock := selectMock{
		Err: errors.New("not found"),
	}
	db := SqlHandler{
		Command: mmock,
	}
	_, err := db.selectone(string(key))
	assert.EqualError(t, err, "not found") // test , act , exc
}
func Test_select_one2(t *testing.T) {
	key := "20"
	ans := ` {"Id": "20","Title": "TWENTY","Desc": "20","Content": "c20","ISBN": "4562312345641","Time": "2020-01-07T13:50:27+07:00","Recentupdate": "2020-01-07T13:50:27+07:00"}`

	var a Article
	json.Unmarshal([]byte(ans), &a)
	mmock := selectMock{
		RespondA: a,
	}
	db := SqlHandler{
		Command: mmock,
	}
	var ans1 string
	ans1, _ = db.selectone(string(key))
	assert.Equal(t, DB20, ans1, "Hello") // test , act , exc
}
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

func Test_createstring_selectone3(t *testing.T) {
	var dum Article
	var err error

	expected := `{"Id":"11","Title": "twelve","Desc": "test","Content": "Content","ISBN": "00100000a06676","Time": "2019-12-23T08:30:24Z","Recentupdate": "2019-12-23T08:30:24Z"}`

	json.Unmarshal([]byte(expected), &dum)
	_, err = createstring_selectone(dum)
	assert.EqualError(t, err, "cannot convert isbn")

}

func TestSqlHandler_selectall2(t *testing.T) {
	ans := ` {"Id": "20","Title": "TWENTY","Desc": "20","Content": "c20","ISBN": "456-2-31-234564-1","Time": "2020-01-07T13:50:27+07:00","Recentupdate": "2020-01-07T13:50:27+07:00"}`

	var a Article
	var set []Article

	json.Unmarshal([]byte(ans), &a)
	set = append(set, a)
	mmock := selectMock{
		RespondArray: set,
	}
	db := SqlHandler{
		Command: mmock,
	}
	ans1, _ := db.selectall2()
	//assert.EqualError(t, err, "not found")
	assert.Equal(t, "["+DB20+"]", ans1, "Hello")

}

// key := "20"
// mmock := selectMock{
// 	Err: errors.New("not found"),
// }
// db := SqlHandler{
// 	Command: mmock,
// }
// _, err := db.selectone(string(key))
// assert.EqualError(t, err, "not found") // test , act , exc
