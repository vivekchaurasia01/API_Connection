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

func TestHandleGoodbye(t *testing.T){
	w := httptest.NewRecorder()

	handleGoodbye(w,nil)

	desireCode := http.StatusOK
	if w.Code != desireCode{
		t.Errorf("bad response code,expected :%v\n but got :%v\nbody:%s\n",desireCode,w.Code,w.Body.String())
	}
	expectedMessage := []byte("Goodbye_User(Thank You for using our services)")
	if !bytes.Equal(expectedMessage,w.Body.Bytes()){
		t.Errorf("bad response code,Got :%q,Expected %q",w.Body.String(),expectedMessage)
	}
}
func TestHandleHelloParameterize(t *testing.T){
	req := httptest.NewRequest(http.MethodGet,"/hello?user=TestMan",nil)

	w := httptest.NewRecorder()

	handleHelloParameterize(w,req)

		desireCode := http.StatusOK

	if w.Code != desireCode{
		t.Errorf("bad response code,expected :%v\n but got :%v\nbody:%s\n",desireCode,w.Code,w.Body.String())
	}
	expectedMessage := []byte("hello,TestMan!\n")
	if !bytes.Equal(expectedMessage,w.Body.Bytes()){
		t.Errorf("bad response code,Got:%q,Expected %q",w.Body.String(),expectedMessage)
	}
}

func TestHandleHelloParameterizeNoParam(t *testing.T){
	req := httptest.NewRequest(http.MethodGet,"/hello/",nil)

	w := httptest.NewRecorder()

	handleHelloParameterize(w,req)
		desireCode := http.StatusOK

	if w.Code != desireCode{
		t.Errorf("bad response code,expected :%v\n but got :%v\nbody:%s\n",desireCode,w.Code,w.Body.String())
	}
	expectedMessage := []byte("hello,User!\n")
	if !bytes.Equal(expectedMessage,w.Body.Bytes()){
		t.Errorf("bad response code,Got:%q,Expected %q",w.Body.String(),expectedMessage)
	}
}
func TestHandleHelloParameterizeWrongParam(t *testing.T){
	req := httptest.NewRequest(http.MethodGet,"/hello?xyz=pqr",nil)
	
	w := httptest.NewRecorder()

	handleHelloParameterize(w,req)
		desireCode := http.StatusOK

	if w.Code != desireCode{
		t.Errorf("bad response code,expected :%v\n but got :%v\nbody:%s\n",desireCode,w.Code,w.Body.String())
	}
	expectedMessage := []byte("hello,User!\n")
	if !bytes.Equal(expectedMessage,w.Body.Bytes()){
		t.Errorf("bad response code,Got:%q,Expected %q",w.Body.String(),expectedMessage)
	}
}