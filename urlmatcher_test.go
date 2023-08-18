package urlmatcher

import "testing"

func TestMatcherSimple(t *testing.T) {

	m := NewMatcher("/users/{id}")

	if _, err := m.Match("/users/a"); err != nil {
		t.Fail()
	}

	if _, err := m.Match("/users/1"); err != nil {
		t.Fail()
	}

}

func TestMatcherRegexpFailed(t *testing.T) {

	m := NewMatcher("/users/{id:[0-9]*}")

	if _, err := m.Match("/users/a"); err == nil {
		t.Fail()
	}
}

func TestMatcherRegexpSuccess(t *testing.T) {

	m := NewMatcher("/users/{id:[0-9]*}")
	match, err := m.Match("/users/1")

	if err != nil {
		t.Fatal("Expected no error")
	}

	if !match.Has("id") {
		t.Fatalf("Expected id")
	}

	if match.Var("id") != "1" {
		t.Fatalf("Expected id to be 1")
	}
}

func TestMatcherRegexpExtra(t *testing.T) {

	m := NewMatcher("/users/{id:(me|you)}")

	if _, err := m.Match("/users/me"); err != nil {
		t.Fatal("Expected me to pass")
	}

	if _, err := m.Match("/users/you"); err != nil {
		t.Fatal("Expected me to pass")
	}

	if _, err := m.Match("/users/them"); err == nil {
		t.Fatal("Expected them to not pass")
	}
}

func TestMatcherMultipleVars(t *testing.T) {

	m := NewMatcher("/users/{id:[0-9]*}/{action:(delete|update)}")

	if _, err := m.Match("/users/1/delete"); err != nil {
		t.Fail()
	}

	if _, err := m.Match("/users/1/update"); err != nil {
		t.Fail()
	}

	if m, err := m.Match("/users/1/delete"); err != nil || !m.Has("id") || !m.Has("action") {
		t.Fail()
	}
}
