package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestSendHTTPMessage(t *testing.T) {
	testData := map[string]string{
		"chat_id": "345182391",
		"message": "hello, http!",
	}
	b, _ := json.Marshal(testData)
	req := httptest.NewRequest("POST", "/send/telegram", bytes.NewBuffer(b))
	w := httptest.NewRecorder()
	sendMessage(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != "ok\n" {
		t.Errorf("expected ok got %v", string(data))
	}
}
