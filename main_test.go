package main

import "testing"

func TestToTest(t *testing.T) {

	if ToTest() != "mock" {
		t.Errorf("fail")
	}
}
