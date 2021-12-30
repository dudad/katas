package bookseller

import "testing"

func TestSingleBook(t *testing.T) {
	grouped := StockList([]string{"ABC 10"}, []string{"A"})
	if grouped != "(A - 20)" {
		t.Error()
	}
}
