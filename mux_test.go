package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewMux(t *testing.T) {
	// httptest.NewRecorder 関数と httptest.NewRequest 関数を利用することで
	// GoではHTTPサーバーを起動しなくても簡単にHTTPハンドラーに対するテストコードを作成できます
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)

	sut := NewMux()
	sut.ServeHTTP(w, r)
	resp := w.Result()
	t.Cleanup(func() { _ = resp.Body.Close() })

	if resp.StatusCode != http.StatusOK {
		t.Error("want status code 200, but", resp.StatusCode)
	}
	got, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}

	want := `{"status": " ok"}`
	if string(got) != want {
		t.Errorf("want %q, but got %q", want, got)
	}
}
