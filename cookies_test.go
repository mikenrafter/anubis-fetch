package main

import "testing"

func TestParseCookieHeader(t *testing.T) {
	cookies := parseCookieHeader("auth=abc==; theme=dark; malformed; empty=")
	if len(cookies) != 3 {
		t.Fatalf("got %d cookies, want 3", len(cookies))
	}
	if cookies[0].Name != "auth" || cookies[0].Value != "abc==" {
		t.Fatalf("auth cookie = %#v", cookies[0])
	}
	if cookies[1].Name != "theme" || cookies[1].Value != "dark" {
		t.Fatalf("theme cookie = %#v", cookies[1])
	}
	if cookies[2].Name != "empty" || cookies[2].Value != "" {
		t.Fatalf("empty cookie = %#v", cookies[2])
	}
}
