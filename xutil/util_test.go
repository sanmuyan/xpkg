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

func TestFillObj(t *testing.T) {
	type User struct {
		Name string `json:"name" copy:"force"`
		Age  int
		Is   bool
	}
	sv := User{Name: "a", Age: 1, Is: false}
	tv := User{Name: "b", Age: 2, Is: true}
	err := FillObj(&sv, &tv)
	if err != nil {
		t.Errorf("FillObj error: %s", err)
	}
	if tv.Age != 1 || tv.Is {
		t.Errorf("FillObj error: %d", tv.Age)
	}
}
