package main

import "testing"

var tests = []struct {
	testName    string
	divident    float32
	divisor     float32
	result      float32
	errExpected bool
}{
	{"passing_test", 100.0, 10.0, 10.0, false},
	{"failing_test", 100.0, 0.0, 0.0, true},
}

func TestMyDivision(t *testing.T) {
	for _, tt := range tests {
		res, err := divide(tt.divident, tt.divisor)
		if tt.errExpected {
			if err == nil {
				t.Error("Expected error, got None")
			}
		} else {
			if err != nil {
				t.Error("Should not throw Error: ", err.Error())
			}
		}

		if res != tt.result {
			t.Error("Results are not matching, Got:", res, " Expected:", tt.result)
		}
	}
}

func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)
	if err != nil {
		t.Error("Got an error for Divide")
	}
}

func TestBadDivide(t *testing.T) {
	_, err := divide(10.0, 0.0)
	if err == nil {
		t.Error("We should get an error, but did not get")
	}
}
