package main

import "testing"

func TestAppName(t *testing.T) {
	expected := "Zoo Application"
	if AppName() != expected {
		t.Errorf("AppName() = %v; want %v", AppName(), expected)
	}
}