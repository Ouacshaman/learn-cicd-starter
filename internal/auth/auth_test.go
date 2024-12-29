package auth

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedAPIKey string
	}{
		{
			name:           "test 1",
			input:          "https://run.mocky.io/v3/c2c276e0-d81b-41e6-8761-c412f579bc95",
			expectedAPIKey: "123-456-kappa",
		},
	}
	for _, v := range tests {
		response, err := http.Get(v.input)
		if err != nil {
			fmt.Println("unable to get response from URL:", err)
			return
		}
		got, err := GetAPIKey(response.Header)
		if err != nil {
			fmt.Println(err)
			return
		}
		if !reflect.DeepEqual(v.expectedAPIKey, got) {
			t.Fatalf("%s: expected: %v, got: %v", v.name, v.expectedAPIKey, got)
		}
	}
}
