package test

import (
	"github.com/masputrawae/go-todo/utils"
	"testing"
)

func TestIsLetterSuccess(t *testing.T) {
	if !utils.IsLetter("A") {
		t.Errorf("Faild test")
		return
	}

	if !utils.IsLetter("a") {
		t.Errorf("Faild test")
		return
	}
}

func TestIsLetterFaild(t *testing.T) {
	if utils.IsLetter("World") {
		t.Errorf("Faild test")
		return
	}

	if utils.IsLetter("1") {
		t.Errorf("Faild test")
		return
	}
}
