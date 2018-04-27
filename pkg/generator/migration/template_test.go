package migration

import "testing"

func TestToSnake(t *testing.T) {
	t1 := ToSnake("Name")
	if t1 != "name" {
		t.Errorf("invalid snake format of %v; got %s, want %s", "Name", t1, "name")
	}
	t2 := ToSnake("NameCheck")
	if t2 != "name_check" {
		t.Errorf("invalid snake format of %v; got %s, want %s", "NameCheck", t2, "name_check")
	}
	t3 := ToSnake("aCheck")
	if t3 != "a_check" {
		t.Errorf("invalid snake format of %v; got %s, want %s", "aCheck", t3, "a_check")
	}
}
