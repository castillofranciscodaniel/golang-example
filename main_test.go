package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

var ContainerTest = InitializeServer()

func TestUpperCaseHandler(t *testing.T) {

	byte := []byte(`{
    "id": 2,
    "name": "Attila's Blog",
    "price": 6
  }`)

	r := bytes.NewReader(byte)
	req := httptest.NewRequest(http.MethodGet, "/modifyProductById", r)

	var writer http.ResponseWriter
	ContainerTest.ProductHandler.HandlerProductByID(writer, req)


}
