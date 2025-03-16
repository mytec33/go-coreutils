package platform

import (
	"os/exec"
	"strings"
	"testing"
)

const ID = "id"
const WHOAMI = "whoami"

func nativeId(t *testing.T, args string) string {
	native, err := exec.Command(ID, args).Output()
	if err != nil {
		t.Fatalf("Failed to run whoami: %v", err)
	}

	return string(native)
}

func nativeWhoami(t *testing.T) string {
	native, err := exec.Command(WHOAMI).Output()
	if err != nil {
		t.Fatalf("Failed to run whoami: %v", err)
	}

	return string(native)
}

func TestGetUserIdExact(t *testing.T) {
	expected := strings.TrimSpace(nativeWhoami(t))

	got, err := GetUserId()
	if err != nil {
		t.Fatalf("failed to get user id %v", err)
	}

	if expected != got {
		t.Fatalf("expected %s, got %s\n", expected, got)
	}
}

func TestGetUserIdNormal(t *testing.T) {
	expected := nativeWhoami(t)

	got, err := GetUserId()
	if err != nil {
		t.Fatalf("failed to get user id %v", err)
	}
	got += "\n"

	if expected != got {
		t.Fatalf("expected %s, got %s\n", expected, got)
	}
}

func TestGetUserIdNormalVsId(t *testing.T) {
	expected := nativeId(t, "-un")

	got, err := GetUserId()
	if err != nil {
		t.Fatalf("failed to get user id %v", err)
	}
	got += "\n"

	if expected != got {
		t.Fatalf("expected %s, got %s\n", expected, got)
	}
}
