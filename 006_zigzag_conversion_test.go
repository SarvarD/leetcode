package leetcode

import "testing"

func TestZigZagConvert_3(t *testing.T) {
	res := ZigZagConvert("PAYPALISHIRING", 3)
	expected := "PAHNAPLSIIGYIR"

	if res != expected {
		t.Errorf("Expected %s, got %s", expected, res)
	}
}

func TestZigZagConvert_4(t *testing.T) {
	res := ZigZagConvert("PAYPALISHIRING", 4)
	expected := "PINALSIGYAHRPI"

	if res != expected {
		t.Errorf("Expected %s, got %s", expected, res)
	}
}

func TestZigZagConvert_6(t *testing.T) {
	res := ZigZagConvert("0123456789", 6)
	expected := "0192837465"

	if res != expected {
		t.Errorf("Expected %s, got %s", expected, res)
	}
}

func TestZigZagConvert_1(t *testing.T) {
	res := ZigZagConvert("sutlej", 1)
	expected := "sutlej"

	if res != expected {
		t.Errorf("Expected %s, got %s", expected, res)
	}
}

func TestZigZagConvert_2(t *testing.T) {
	res := ZigZagConvert("sutlej", 2)
	expected := "steulj"

	if res != expected {
		t.Errorf("Expected %s, got %s", expected, res)
	}
}
