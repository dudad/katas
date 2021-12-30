package bookseller

import "testing"

func TestSingleBook(t *testing.T) {
	grouped := StockList([]string{"ABC 20"}, []string{"A"})
	if grouped != "(A : 20)" {
		t.Error("Retrurned:", grouped)
	}
}

func TestTwoBooksInTheSameCategory(t *testing.T) {
	grouped := StockList([]string{"ABC 10", "AUY 20"}, []string{"A"})
	if grouped != "(A : 30)" {
		t.Error("Retrurned:", grouped)
	}
}

func TestTwoBooksInDiffrentCategories(t *testing.T) {
	grouped := StockList([]string{"ABC 10", "BUY 20"}, []string{"A", "B"})
	if grouped != "(A : 10) - (B : 20)" {
		t.Error("Retrurned:", grouped)
	}
}

func TestOnlySpecifiedCategoriesAreListed(t *testing.T) {
	grouped := StockList([]string{"ABC 10", "BUY 20", "CAT 33"}, []string{"B", "A"})
	if grouped != "(B : 20) - (A : 10)" {
		t.Error("Retrurned:", grouped)
	}
}

func TestEmptyCategory(t *testing.T) {
	grouped := StockList([]string{"ABC 10", "BUY 20", "CAT 33"}, []string{"B", "A", "E"})
	if grouped != "(B : 20) - (A : 10) - (E : 0)" {
		t.Error("Retrurned:", grouped)
	}
}

func TestEmptyBookList(t *testing.T) {
	grouped := StockList([]string{}, []string{"B", "A", "E"})
	if grouped != "" {
		t.Error("Retrurned:", grouped)
	}
}

func TestEmptyCatList(t *testing.T) {
	grouped := StockList([]string{"ABC 10", "BUY 20", "CAT 33"}, []string{})
	if grouped != "" {
		t.Error("Retrurned:", grouped)
	}
}



