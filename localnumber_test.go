package localnumber

import "testing"

func TestUKPhoneNumber(t *testing.T) {
	input := "+44 1932 567890"
	expected := "01932567890"
	output, err := Format(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if output != expected {
		t.Fatalf("output: %v, expected: %v", output, expected)
	}
}

func TestUKPhoneNumberLocal(t *testing.T) {
	input := "01932567890"
	expected := "01932567890"
	output, err := Format(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if output != expected {
		t.Fatalf("output: %v, expected: %v", output, expected)
	}
}

func TestUKPhoneNumberLocalSpace(t *testing.T) {
	input := "01932 567890"
	expected := "01932567890"
	output, err := Format(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if output != expected {
		t.Fatalf("output: %v, expected: %v", output, expected)
	}
}

func TestIndianPhoneNumber(t *testing.T) {
	input := "+91 6297 062979"
	expected := "6297062979"
	output, err := Format(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if output != expected {
		t.Fatalf("output: %v, expected: %v", output, expected)
	}
}
