package main

import "testing"

func TestRot13(t *testing.T) {
	tt := struct {
		name  string
		input string
		want  string
	}{
		name:  "Rot13",
		input: "Hello There",
		want:  "Uryyb Gurer",
	}

	if got := Rot13(tt.input); got != tt.want {
		t.Errorf("Test-1 failed :- Rot13(%s) = %s, want %s", tt.input, got, tt.want)
	}
}

func TestRot132(t *testing.T) {
	tt := struct {
		name  string
		input string
		want  string
	}{
		name:  "Rot13",
		input: "Hello There 13",
		want:  "Uryyb Gurer 13",
	}

	if got := Rot13(tt.input); got != tt.want {
		t.Errorf("Test-2 failed :- Rot13(%s) = %s, want %s", tt.input, got, tt.want)
	}
}

func TestRot133(t *testing.T) {
	tt := struct {
		name  string
		input string
		want  string
	}{
		name:  "Rot13",
		input: "My Name Is Godwin Ouma.",
		want:  "Zl Anzr Vf Tbqjva Bhzn.",
	}

	if got := Rot13(tt.input); got != tt.want {
		t.Errorf("Test-3 failed :- Rot13(%s) = %s, want %s", tt.input, got, tt.want)
	}
}
