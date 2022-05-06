package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_getAlbums(t *testing.T) {

	ts := httptest.NewServer(setupServer())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/albums", ts.URL))

	if err != nil {
		t.Fatalf("Got error %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Got error code %v", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

}
