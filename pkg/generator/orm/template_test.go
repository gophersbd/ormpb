package orm

import "testing"

func TestPrintInterface(t *testing.T) {
	t1 := printInterface(5)
	if t1 != `"5"` {
		t.Errorf("invalid string of interface %v; got %s, want %s", "5", t1, "5")
	}

	t2 := printInterface(5.2)
	if t2 != `"5.2"` {
		t.Errorf("invalid string of interface %v; got %s, want %s", "5.2", t1, "5.2")
	}

	t3 := printInterface(true)
	if t3 != `"true"` {
		t.Errorf("invalid string of interface %v; got %s, want %s", "true", t1, "true")
	}

	t4 := printInterface("check")
	if t4 != `"check"` {
		t.Errorf("invalid string of interface %v; got %s, want %s", "check", t1, `"check"`)
	}
}
