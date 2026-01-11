package cr

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	client := NewClient("test-token")
	if client.token != "test-token" {
		t.Errorf("Token mismatch")
	}
	if client.http.Timeout != 5*time.Second {
		t.Error("Timeout not set")
	}
}

func TestGetPlayer_InvalidResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
	}))
	defer server.Close()

	client := NewClient("test")
	client.baseURL = server.URL
	client.http = server.Client()

	_, err := client.GetPlayer("#TEST")
	if err == nil {
		t.Error("Expected error")
	}
}

func TestGetClan_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"name":"Test Clan","members":50}`))
	}))
	defer server.Close()

	client := NewClient("test")
	client.baseURL = server.URL
	client.http = server.Client()

	clan, err := client.GetClan("#CLAN")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if clan.Name != "Test Clan" {
		t.Errorf("Expected clan name 'Test Clan', got '%s'", clan.Name)
	}
}

func TestGetBattleLog_Empty(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`[]`))
	}))
	defer server.Close()

	client := NewClient("test")
	client.baseURL = server.URL
	client.http = server.Client()

	battleLog, err := client.GetBattleLog("#PLAYER")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(*battleLog) != 0 {
		t.Errorf("Expected empty battle log")
	}
}
