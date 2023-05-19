package main

import (
	"testing"
	"time"
)

type MinMaxData struct {
	x, y, result int
}

// To run all tests and have one output (PASS or FAIL)
// go test

// To run all tests and see output of each test (PASS OR FAIL FOR EACH TEST)
// go test -v

// To see test coverage
// go test -cover

func TestMax(t *testing.T){

testData := []MinMaxData{
	{1, 2, 2},
	{4, 8, 8},
	{4, 6, 6},
} 

for _, item := range testData{
	result := max(item.x, item.y)
	if (result == item.x){
		t.Errorf("max() FAILED, Expected: %v Received: %v", item.result, item.x)
	} else {
		t.Logf("max() PASSED, Expected: %v Received: %v", item.result, item.y)
	}
}

}

func TestMin(t *testing.T){

	testData := []MinMaxData{
		{2,4,2},
		{3,4,3},
		{4,9,4},
	}

	for _, item := range testData{

		result := min(item.x, item.y)
		if (result == item.y){
			t.Errorf("min() FAILED, Expected: %v Received: %v", item.result, item.y)
		} else {
			t.Logf("min() PASSED, Expected: %v Received: %v", item.result, item.y)
		}
	}

}

func TestWait(t *testing.T) {

	start := time.Now()

	Wait(1)

	elapsed := time.Since(start)

	tolerance := 100 * time.Millisecond
	expectedDuration := 1 * time.Second

	if elapsed < expectedDuration-tolerance || elapsed > expectedDuration+tolerance {
		t.Errorf("Wait() FAILED,Expected duration %s, got %s", expectedDuration, elapsed)
	} else {
		t.Logf("Wait() PASSED, Expected duration %s, got %s", expectedDuration, elapsed)
	}
}

func TestFillGrid(t *testing.T){

	result, result2 := FillGrid()

	if(gameGrid[result][result2] != question){
		t.Errorf("FillGrid() FAILED, Expected: %v Received: %v", true, false)
	} else {
		t.Logf("FillGrid() PASSED, Expected: %v Received: %v", true, true)
	}

}

func TestAdjacentBombsNoDisplayOnInit(t *testing.T){

	FillGrid()
	PlaceBombs(5)
	DisplayGrid()

	result := AdjacentBombs(0,0)

	if(result != 0){
		t.Errorf("FAILED, Expected: %v Received: %v", result, result)
	} else {
		t.Logf("PASSED, Expected: %v Received: %v", result, result)
	}


}