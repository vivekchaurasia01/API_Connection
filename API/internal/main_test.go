package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleRoot(t *testing.T){
	w := httptest.NewRecorder()

	handleRoot(w,nil)

	desireCode := http.StatusOK
	if w.Code != desireCode{
		t.Errorf("bad response code,expected :%v\n but got :%v\nbody:%s\n",desireCode,w.Code,w.Body.String())
	}
	expectedMessage := []byte("Welcome_User")
	if !bytes.Equal(expectedMessage,w.Body.Bytes()){
		t.Errorf("bad response code,Got :%q,Expected %q",w.Body.String(),expectedMessage)
	}


}