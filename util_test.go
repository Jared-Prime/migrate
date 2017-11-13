package migrate

import (
	nurl "net/url"
	"testing"
	"time"
)

func TestSuintPanicsWithNegativeInput(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected suint to panic for -1")
		}
	}()
	suint(-1)
}

func TestSuint(t *testing.T) {
	if u := suint(0); u != 0 {
		t.Fatalf("expected 0, got %v", u)
	}
}

func TestFilterCustomQuery(t *testing.T) {
	n, err := nurl.Parse("foo://host?a=b&x-custom=foo&c=d")
	if err != nil {
		t.Fatal(err)
	}
	nx := FilterCustomQuery(n).Query()
	if nx.Get("x-custom") != "" {
		t.Fatalf("didn't expect x-custom")
	}
}

func TestCreateTimestamp(t *testing.T) {
	var result int64
	var err error

	startTime := time.Date(2012, time.May, 13, 0, 0, 0, 0, time.UTC)

	result, err = createTimestamp("unix", startTime)
	if err != nil {
		t.Fatal("unexpected error: %s", err.Error())
	}
	if result != 1336867200 {
		t.Fatal("expected timestamp to equal 1336867200", "got", result)
	}

	result, err = createTimestamp("epoch", startTime)
	if err != nil {
		t.Fatal("unexpected error: %s", err.Error())
	}
	if result != 1336867200 {
		t.Fatal("expected timestamp to equal 1336867200", "got", result)
	}

	result, err = createTimestamp("rails", startTime)
	if err != nil {
		t.Fatal("unexpected error: %s", err.Error())
	}
	if result != 20120513000000 {
		t.Fatal("expected timestamp to equal 20120513000000", "got", result)
	}

	result, err = createTimestamp("other", startTime)
	if result != 0 {
		t.Fatal("unexpected timestamp value:", result, "expected: other")
	}
	if err.Error() != "unsupported timestamp format: other" {
		t.Fatal("unexpected error message:", err.Error())
	}
}
