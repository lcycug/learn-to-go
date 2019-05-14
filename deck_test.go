package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()
	if len(cards) != 9 {
		t.Errorf("Expected result is 9, but got %v.", len(cards))
	}
}

func TestSaveToLocalAndLoadingFromLocal(t *testing.T) {
	os.Remove("_testFile")

	cards := newDeck()

	if len(cards) != 9 {
		t.Errorf("Expected Result is 9, but got %v", len(cards))
	}

	cards.saveToLocal()

	newCards := loadFromLocal()

	if len(newCards) != 9 {
		t.Errorf("Expected Result is 9, but got %v", len(newCards))
	}
}
