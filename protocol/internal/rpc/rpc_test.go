package rpc

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func setupTestServer(handler http.HandlerFunc) *httptest.Server {
	return httptest.NewServer(handler)
}
func TestCall_SuccessAndStructure(t *testing.T) {
	server := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read body", http.StatusInternalServerError)
			return
		}
		var req Request
		if err := json.Unmarshal(body, &req); err != nil {
			http.Error(w, "failed to parse json", http.StatusBadRequest)
			return
		}
		if req.JSONRPC != "2.0" || req.Method != "test_method" || len(req.Params) != 1 {
			t.Errorf("Server received bad request structure: %s", string(body))
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"jsonrpc":"2.0", "id":1, "result":"0x12345"}`))
	})
	defer server.Close()

	client := NewClient(server.URL, 2*time.Second)
	result, err := client.Call(context.Background(), "test_method", "param1")
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if string(result) != `"0x12345"` {
		t.Errorf("Expected result '\"0x12345\"', got %s", string(result))
	}
}
func TestCall_RPCError(t *testing.T) {
	server := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"jsonrpc":"2.0", "id":1, "error":{"code":-32600, "message":"Invalid Request"}}`))
	})
	defer server.Close()

	client := NewClient(server.URL, 2*time.Second)
	_, err := client.Call(context.Background(), "bad_method")

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
	if !strings.Contains(err.Error(), "rpc error -32600") {
		t.Errorf("Expected RPC error message, got: %v", err)
	}
}

func TestCall_HTTPError(t *testing.T) {
	server := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})
	defer server.Close()

	client := NewClient(server.URL, 2*time.Second)
	_, err := client.Call(context.Background(), "test_method")

	if err == nil {
		t.Fatal("Expected HTTP error, got nil")
	}
	if !strings.Contains(err.Error(), "server returned http status") {
		t.Errorf("Expected status error, got: %v", err)
	}
}

func TestCall_Timeout(t *testing.T) {
	server := setupTestServer(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
	})
	defer server.Close()

	client := NewClient(server.URL, 50*time.Millisecond)
	_, err := client.Call(context.Background(), "test_method")

	if err == nil {
		t.Fatal("Expected timeout error, got nil")
	}

	// Better timeout check: Ask Go's error system if this is specifically a network timeout error
	var netErr net.Error
	if !errors.As(err, &netErr) || !netErr.Timeout() {
		t.Errorf("Expected a true network timeout error, got: %v", err)
	}
}
