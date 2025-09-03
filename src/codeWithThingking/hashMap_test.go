package codewiththingking_test

import (
	codewiththingking "go_alg/src/codeWithThingking"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	s1 := "anagram"
	t1 := "nagaram"
	expected := true
	actual := codewiththingking.IsAnagram(s1, t1)
	if expected != actual {
		t.Errorf("s1 should be equal to t1")
	}
	s2 := "rat"
	t2 := "car"
	expected = false
	actual = codewiththingking.IsAnagram(s2, t2)
	if expected != actual {
		t.Errorf("s1 should be not equal to t2")
	}
}

func TestCanConstruct(t *testing.T) {
	s1 := "a"
	t1 := "b"
	expected := false
	actual := codewiththingking.CanConstruct(s1, t1)
	if expected != actual {
		t.Errorf("s1 : %s , t1 : %s", s1, t1)
	}
	s2 := "aa"
	t2 := "ab"
	expected = false
	actual = codewiththingking.CanConstruct(s2, t2)
	if expected != actual {
		t.Errorf("s2 : %s , t2 : %s", s2, t2)
	}
	s3 := "aa"
	t3 := "aab"
	expected = true
	actual = codewiththingking.CanConstruct(s3, t3)
	if expected != actual {
		t.Errorf("s3 : %s , t3 : %s", s3, t3)
	}
}

func TestIRemoveDuplicates(t *testing.T) {
	s := "abbaca"
	expected := "ca"
	actual := codewiththingking.IRemoveDuplicates(s)
	if expected != actual {
		t.Errorf("expected not equal to the actual ,%s, %s", expected, actual)
	}
}
