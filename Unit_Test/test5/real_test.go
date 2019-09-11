package main

import (
	"reflect"
	"testing"
)

func Test_Parameter(t *testing.T) {

	t.Run("get userclient returns err if response body invalid", func(t *testing.T) {
		userclient.requester = &FakeClient{}

		_, err := userclient.Get()
		if err == nil {
			t.Error("Expected to get an error")
		}
	})

	t.Run("get userclient returns unmarshalled response body on success", func(t *testing.T) {

		userclient.requester = &FakeClient{
			Body: []byte(`{"status": "ok"}`),
		}
		got, _ := userclient.Get()
		want := map[string]interface{}{
			"status": "ok",
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
