package bookseller

import "testing"

func TestSingleBook(t *testing.T) {
	grouped := StockList([]string{"ABC 20"}, []string{"A"})
	if grouped != "(A - 20)" {
		t.Error("Retrurned:", grouped)
	}
}

func TestTwoBooksInTheSameCategory(t *testing.T) {
	grouped := StockList([]string{"ABC 10", "AUY 20"}, []string{"A"})
	if grouped != "(A - 30)" {
		t.Error("Retrurned:", grouped)
	}
}
