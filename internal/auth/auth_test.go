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
			input:          "https://run.mocky.io/v3/95b20dc8-6073-4598-858c-8309faf28ef8",
			expectedAPIKey: "123-456-kappa",
		},
		{
			name:           "test 2",
			input:          "",
			expectedAPIKey: "",
		},
	}
	for _, v := range tests {
		if v.input == "" {
			fmt.Printf("for %s, you enter %v", v.name, v.input)
			continue
		}
		response, err := http.Get(v.input)
		if err != nil {
			t.Fatal("Response Not Found")
		}
		got, err := GetAPIKey(response.Header)
		if err != nil {
			t.Fatal("APIKey Not Found")
		}
		if !reflect.DeepEqual(v.expectedAPIKey, got) {
			t.Fatalf("%s: expected: %v, got: %v", v.name, v.expectedAPIKey, got)
		}
	}
}
