package app

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/CSY54/go-url-shortener/src/url"
)

func performRequst(r http.Handler, method, path, json string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer([]byte(json)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestAPI(t *testing.T) {
	testcases := []struct{
		body string
		status int
		message string
	}{
		{
			`{"url":"http://localhost","expireAt":"2069-12-31T00:00:00Z"}`,
			200,
			`{"id":"AQ","shortUrl":"http://localhost/AQ"}`,
		},
		{
			`{"url":"http://localhost"}`,
			400,
			`{"error":"Key missing or value invalid"}`,
		},
		{
			`{"expireAt":"2069-12-31T00:00:00Z"}`,
			400,
			`{"error":"Key missing or value invalid"}`,
		},
		{
			`{"url":"ftp://localhost","expireAt":"2069-12-31T00:00:00Z"}`,
			400,
			`{"error":"Not a valid URL"}`,
		},
		{
			`{"url":"http://localhost","expireAt":"2000-12-31T00:00:00Z"}`,
			400,
			`{"error":"Time already expired"}`,
		},
	}

	r := Init(true)

	for _, tc := range testcases {
		w := performRequst(r, "POST", "/api/v1/urls", tc.body)
		if w.Code != tc.status || w.Body.String() != tc.message {
			t.Errorf("Expected: %d, %s, got: %d, %s", tc.status, tc.message, w.Code, w.Body)
		}
	}
}

func TestRedirect(t *testing.T) {
	now := time.Now()
	seeds := []url.Url{
		{Url: "http://localhost/1", ExpireAt: now.Add(time.Hour * 1)},
		{Url: "http://localhost/2", ExpireAt: now},
	}

	testcases := []struct{
		shortUrl string
		status int
		redirectUrl string
	}{
		{
			"AQ",
			302,
			"http://localhost/1",
		},
		{
			"Ag",
			404,
			"",
		},
		{
			"meow",
			404,
			"",
		},
	}

	db := setupDatabase(true)
	repo := url.ProvideUrlRepository(db)
	for _, seed := range seeds {
		repo.Create(seed)
	}

	r := Init(false)

	for _, tc := range testcases {
		w := performRequst(r, "GET", "/" + tc.shortUrl, "")
		if w.Code != tc.status || w.HeaderMap.Get("Location") != tc.redirectUrl {
			t.Errorf("Expected: %d, %s, got: %d, %s", tc.status, tc.redirectUrl, w.Code, w.HeaderMap.Get("Location"))
		}
	}
}
