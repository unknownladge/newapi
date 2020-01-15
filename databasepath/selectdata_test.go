package databasepath

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type selectMock struct {
	Respond      string
	RespondA     Article
	RespondArray []Article
	Err          error
}

func (db selectMock) Dbget1(string) (Article, error) {
	return db.RespondA, db.Err
}
func (db selectMock) Inserter(string) (string, error) {
	return db.Respond, db.Err
}
func (db selectMock) Updatedata(string, string, string, string, string) (string, error) {
	return db.Respond, db.Err
}
func (db selectMock) Deletedata(string) (string, error) {
	return db.Respond, db.Err
}
func (db selectMock) Dbgetall() ([]Article, error) {
	return db.RespondArray, db.Err
}
func (db selectMock) Selectone(string) (string, error) {
	return db.Respond, db.Err
}
func (db selectMock) Selectall2() (string, error) {
	return db.Respond, db.Err
}

////////////////////////////////////////////////////////////////
type Ordermock struct {
	DbOrderMock  selectMock
	Respond      string
	RespondA     Article
	RespondArray []Article
	Err          error
}

// func (db Ordermock) selectone(string) (string, error) {
// 	return db.Respond, db.Err
// }
// func (db Ordermock) selectall2() (string, error) {
// 	return db.Respond, db.Err
// }

///////////////////////////////////////////////////////////////
func Test_select_one1(t *testing.T) {
	key := "20"
	mmock := selectMock{
		Err: errors.New("not found"),
	}
	db := SqlHandler{
		Command: mmock,
	}

	_, err := db.Selectone(string(key))
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
	ans1, _ = db.Selectone(string(key))  //db.selectone(string(key))
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
	ans1, _ := db.Selectall2()
	//assert.EqualError(t, err, "not found")
	assert.Equal(t, "["+DB20+"]", ans1, "Hello")

}

// func TestReturnAllArticles(t *testing.T) {
// 	//key := "20"
// 	ans := ` {"Id": "20","Title": "TWENTY","Desc": "20","Content": "c20","ISBN": "4562312345641","Time": "2020-01-07T13:50:27+07:00","Recentupdate": "2020-01-07T13:50:27+07:00"}`
// 	var set []Article
// 	var a Article
// 	json.Unmarshal([]byte(ans), &a)
// 	set = append(set, a)
// 	mmock := selectMock{
// 		Respond:      ans,
// 		RespondA:     a,
// 		RespondArray: set,
// 		Err:          errors.New("not found"),
// 	}
// 	mock := Ordermock{
// 		DbOrderMock:  mmock,
// 		Respond:      ans,
// 		RespondA:     a,
// 		RespondArray: set,
// 		Err:          errors.New("not found"),
// 	}

// }

func TestReturnRespond1(t *testing.T) {
	ans := ` {"Id": "20","Title": "TWENTY","Desc": "20","Content": "c20","ISBN": "4562312345641","Time": "2020-01-07T13:50:27+07:00","Recentupdate": "2020-01-07T13:50:27+07:00"}`
	var set []Article
	var a Article
	json.Unmarshal([]byte(ans), &a)
	set = append(set, a)
	mmock := selectMock{
		Respond:      ans,
		RespondA:     a,
		RespondArray: set,
		Err:          errors.New("not found"),
	}
	db := SqlHandler{
		Command: mmock,
	}

	re := Order{
		Dborder: db,
	}
	//ans1 := re.ReturnAllArticles

	req, err := http.NewRequest(http.MethodGet, "/article", nil)
	if err != nil {
		t.Error(err)
	}
	resp := httptest.NewRecorder()

	re.ReturnAllArticles(resp, req)
	if status := resp.Code; status != http.StatusOK {
		t.Errorf("wrong code: got %v want %v", status, http.StatusOK)
	}
}
func TestReturnRespond2(t *testing.T) {
	mmock := selectMock{
		Err: errors.New("Time Format error"),
	}
	db11 := SqlHandler{
		Command: mmock,
	}

	re := Order{
		Dborder: db11,
	}
	//ans1 := re.ReturnAllArticles

	req, err := http.NewRequest(http.MethodGet, "/article", nil)
	if err != nil {
		t.Error(err)
	}
	resp := httptest.NewRecorder()

	re.ReturnAllArticles(resp, req)

	if status := resp.Code; status == http.StatusOK {
		t.Errorf("wrong code: got %v want %v", status, http.StatusOK)
	}

}
func TestReturnRespond3(t *testing.T) {
	mmock := selectMock{
		Err: errors.New("Marshal error"),
	}
	db11 := SqlHandler{
		Command: mmock,
	}

	re := Order{
		Dborder: db11,
	}
	//ans1 := re.ReturnAllArticles

	req, err := http.NewRequest(http.MethodGet, "/article", nil)
	if err != nil {
		t.Error(err)
	}
	resp := httptest.NewRecorder()

	re.ReturnAllArticles(resp, req)

	if status := resp.Code; status == http.StatusOK {
		t.Errorf("wrong code: got %v want %v", status, http.StatusOK)
	}

}

func TestReturnRespond4(t *testing.T) {
	mmock := selectMock{
		Err: errors.New("Scan row error"),
	}
	db11 := SqlHandler{
		Command: mmock,
	}

	re := Order{
		Dborder: db11,
	}
	//ans1 := re.ReturnAllArticles

	req, err := http.NewRequest(http.MethodGet, "/article", nil)
	if err != nil {
		t.Error(err)
	}
	resp := httptest.NewRecorder()

	re.ReturnAllArticles(resp, req)

	if status := resp.Code; status == http.StatusOK {
		t.Errorf("wrong code: got %v want %v", status, http.StatusOK)
	}

}
func TestReturnSingleRespond1(t *testing.T) {
	ans := ` {"Id": "20","Title": "TWENTY","Desc": "20","Content": "c20","ISBN": "4562312345641","Time": "2020-01-07T13:50:27+07:00","Recentupdate": "2020-01-07T13:50:27+07:00"}`
	var set []Article
	var a Article
	json.Unmarshal([]byte(ans), &a)
	set = append(set, a)
	mmock := selectMock{
		Respond:      ans,
		RespondA:     a,
		RespondArray: set,
		//	Err:          errors.New("not found"),
	}
	db := SqlHandler{
		Command: mmock,
	}

	re := Order{
		Dborder: db,
	}
	//ans1 := re.ReturnAllArticles

	req, err := http.NewRequest(http.MethodGet, "/article", nil)
	if err != nil {
		t.Error(err)
	}
	resp := httptest.NewRecorder()

	re.ReturnSingleArticle(resp, req)
	if status := resp.Code; status != http.StatusOK {
		t.Errorf("wrong code: got %v want %v", status, http.StatusOK)
	}
}
func TestReturnSingleRespond2(t *testing.T) {
	mmock := selectMock{
		Err: errors.New("Article not found"),
	}
	db := SqlHandler{
		Command: mmock,
	}

	re := Order{
		Dborder: db,
	}
	//ans1 := re.ReturnAllArticles

	req, err := http.NewRequest(http.MethodGet, "/article", nil)
	if err != nil {
		t.Error(err)
	}
	resp := httptest.NewRecorder()

	re.ReturnSingleArticle(resp, req)
	if status := resp.Code; status == http.StatusOK {
		t.Errorf("wrong code: got %v want %v", status, http.StatusOK)
	}
}
func TestReturnSingleRespond3(t *testing.T) {
	mmock := selectMock{
		Err: errors.New("id not found"),
	}
	db := SqlHandler{
		Command: mmock,
	}

	re := Order{
		Dborder: db,
	}
	//ans1 := re.ReturnAllArticles

	req, err := http.NewRequest(http.MethodGet, "/article", nil)
	if err != nil {
		t.Error(err)
	}
	resp := httptest.NewRecorder()

	re.ReturnSingleArticle(resp, req)
	if status := resp.Code; status == http.StatusOK {
		t.Errorf("wrong code: got %v want %v", status, http.StatusOK)
	}
}

func TestReturnSingleRespond4(t *testing.T) {
	mmock := selectMock{
		Err: errors.New("cannot convert isbn"),
	}
	db := SqlHandler{
		Command: mmock,
	}

	re := Order{
		Dborder: db,
	}
	//ans1 := re.ReturnAllArticles

	req, err := http.NewRequest(http.MethodGet, "/article", nil)
	if err != nil {
		t.Error(err)
	}
	resp := httptest.NewRecorder()

	re.ReturnSingleArticle(resp, req)
	if status := resp.Code; status == http.StatusOK {
		t.Errorf("wrong code: got %v want %v", status, http.StatusOK)
	}
}

func TestReturnSingleRespond5(t *testing.T) {
	mmock := selectMock{
		Err: errors.New("Marshal error"),
	}
	db := SqlHandler{
		Command: mmock,
	}

	re := Order{
		Dborder: db,
	}
	//ans1 := re.ReturnAllArticles

	req, err := http.NewRequest(http.MethodGet, "/article", nil)
	if err != nil {
		t.Error(err)
	}
	resp := httptest.NewRecorder()

	re.ReturnSingleArticle(resp, req)
	if status := resp.Code; status == http.StatusOK {
		t.Errorf("wrong code: got %v want %v", status, http.StatusOK)
	}
}

func TestReturnSingleRespond6(t *testing.T) {
	mmock := selectMock{
		Err: errors.New("Error"),
	}
	db := SqlHandler{
		Command: mmock,
	}

	re := Order{
		Dborder: db,
	}
	//ans1 := re.ReturnAllArticles

	req, err := http.NewRequest(http.MethodGet, "/article", nil)
	if err != nil {
		t.Error(err)
	}
	resp := httptest.NewRecorder()

	re.ReturnSingleArticle(resp, req)
	if status := resp.Code; status == http.StatusOK {
		t.Errorf("wrong code: got %v want %v", status, http.StatusOK)
	}
}
