package gameStats

import (
	"testing"
)

// To run all tests and have one output (PASS or FAIL)
// go test

// To run all tests and see output of each test (PASS OR FAIL FOR EACH TEST)
// go test -v

// To see test coverage
// go test -cover

func TestIncrementWins(t *testing.T){

	
	IncrementWins()

	if (Wins == 0){
		t.Errorf("IncrementWins() FAILED. Expected: %v Received: %v", Wins, 0)
	} else {
		t.Logf("IncrementWins() PASSED. Expected: %v Received: %v", 1, Wins)
	}

	// Clean up
	Wins = 0

}

func TestIncrementLosses(t *testing.T){

	IncrementLosses()

	if (Losses == 0){
		t.Errorf("IncrementLosses() FAILED. Expected: %v Received: %v", Losses, 0)
	} else {
		t.Logf("IncrementLosses() PASSED. Expected: %v Received: %v", 1, Losses)
	}

	// Clean up
	Losses = 0

}