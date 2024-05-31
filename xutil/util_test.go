package xutil

import (
	"testing"
)

func TestIsSubPath(t *testing.T) {
	if !IsSubPath("/api", "/api/user") {
		t.Errorf("IsSubPath error")
	}
	if !IsSubPath("/api/user", "/api/users") {
		t.Errorf("IsSubPath error")
	}
	if IsSubPath("/api/use", "/api/user") {
		t.Errorf("IsSubPath error")
	}
}
