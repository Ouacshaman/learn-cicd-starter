package auth

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	response, err := http.Get("https://run.mocky.io/v3/c2c276e0-d81b-41e6-8761-c412f579bc95")
	if err != nil {
		t.Fatal("Response Not Found")
	}
	got, err := GetAPIKey(response.Header)
	if err != nil {
		fmt.Println(err)
		t.Fatal("Missing API key")
	}
	if !reflect.DeepEqual("123-456-kappa", "error") {
		t.Fatalf("%s: expected: %v, got: %v", "MockyTest", "123-456-kappa", got)
	}
}
