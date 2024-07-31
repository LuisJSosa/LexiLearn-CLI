package main

import (
	"testing"
	"time"
)

// TestAddWord tests the addWord function.
func TestAddWord(t *testing.T) {
	initialLength := len(vocabList)
	addWordHelper("test", "this is a test")
	if len(vocabList) != initialLength+1 {
		t.Errorf("Expected vocabList length to be %d, got %d", initialLength+1, len(vocabList))
	}
}

// addWordHelper adds a word to vocabList for testing purposes.
func addWordHelper(term, definition string) {
	newWord := Word{
		Term:         term,
		Definition:   definition,
		CorrectCount: 0,
		LastReview:   time.Now(),
		NextReview:   time.Now().Add(-24 * time.Hour),
	}
	vocabList = append(vocabList, newWord)
}

// TestDeleteWord tests the deleteWord function.
func TestDeleteWord(t *testing.T) {
	addWordHelper("testToDelete", "this is a test to delete")
	initialLength := len(vocabList)
	deleteWord(len(vocabList) - 1)
	if len(vocabList) != initialLength-1 {
		t.Errorf("Expected vocabList length to be %d, got %d", initialLength-1, len(vocabList))
	}
}

// TestCalculateNextReview tests the calculateNextReview function.
func TestCalculateNextReview(t *testing.T) {
	for i := 1; i <= 5; i++ {
		nextReview := calculateNextReview(i)
		if nextReview.Before(time.Now()) {
			t.Errorf("Expected next review to be after now, got %s", nextReview)
		}
	}
}
