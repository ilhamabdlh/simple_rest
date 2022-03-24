package main

import (
	"testing"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func TestCheckAccount(t *testing.T){
	r := httptest.NewRequest(http.MethodGet, "/account/55502", nil)
    w := httptest.NewRecorder()
	CheckAccount(w, r)
	res := w.Result()
	defer res.Body.Close()
	data, err:= ioutil.ReadAll(res.Body)

	if err != nil{
		t.Errorf("no value %v", err)
	}

	var check =  `{"account_number":"55502","custumer_name":"Linus Torvalds","balance":12200}`

	if string(data) == check{
		t.Errorf("account not found %v", string(data))
	}
}