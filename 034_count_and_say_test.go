package leetcode

import "testing"

func TestCountAndSayRepr(t *testing.T) {
	res := CountAndSayRepr("11122311")
	expected := "31221321"

	if res != expected {
		t.Errorf("Expected %s, got %s", expected, res)
	}

}

func TestCountAndSayReprSingleNumber(t *testing.T) {
	res := CountAndSayRepr("11111111")
	expected := "81"

	if res != expected {
		t.Errorf("Expected %s, got %s", expected, res)
	}

}

func TestCountAndSayReprSingleCharacter(t *testing.T) {
	res := CountAndSayRepr("3")
	expected := "13"

	if res != expected {
		t.Errorf("Expected %s, got %s", expected, res)
	}
}

func TestCountAndSayBaseCase(t *testing.T) {
	res := CountAndSay(1)
	expected := "1"

	if res != expected {
		t.Errorf("Expected %s, got %s", expected, res)
	}
}

func TestCountAndSay4(t *testing.T) {
	res := CountAndSay(4)
	expected := "1211"

	if res != expected {
		t.Errorf("Expected %s, got %s", expected, res)
	}
}
