package tmutil

import (
	"testing"
	"time"
)

func TestNewTimeTermJust(t *testing.T) {
	st := time.Date(2024, 2, 22, 22, 22, 22, 0, time.Local)
	et := time.Date(2024, 2, 22, 22, 33, 33, 0, time.Local)

	tt, err := NewTimeTermJust(st, et)
	if err != nil {
		t.Error("NewTimeTerm Error")
	}
	if !tt.valid ||
		tt.start.value.Format(time.DateTime) != st.Format(time.DateTime) ||
		tt.end.value.Format(time.DateTime) != et.Format(time.DateTime) {
		t.Error("NewTimeTerm Error")
	}
}

func TestNewTimeTerm(t *testing.T) {
	st := time.Date(2024, 2, 22, 22, 22, 22, 0, time.Local)
	et := time.Date(2024, 2, 22, 22, 33, 33, 0, time.Local)

	tt1, err := NewTimeTerm(st, true, et, true)
	if err != nil {
		t.Error("NewTimeTerm Error")
	}
	if !tt1.valid ||
		tt1.start.value.Format(time.DateTime) != st.Format(time.DateTime) ||
		tt1.end.value.Format(time.DateTime) != et.Format(time.DateTime) {
		t.Error("NewTimeTerm Error")
	}

	tt2, err := NewTimeTerm(et, true, st, true)
	if err == nil {
		t.Error("NewTimeTerm Error")
	}
	if tt2.valid {
		t.Error("NewTimeTerm Error")
	}
}

func TestContains1(t *testing.T) {
	tt, _ := NewTimeTerm(
		time.Date(2024, 2, 22, 22, 22, 22, 0, time.Local), true,
		time.Date(2024, 2, 22, 22, 33, 33, 0, time.Local), true,
	)

	if !Contains(tt, time.Date(2024, 2, 22, 22, 33, 22, 0, time.Local)) {
		t.Error("Contains Error")
	}

	if !Contains(tt, time.Date(2024, 2, 22, 22, 22, 22, 0, time.Local)) {
		t.Error("Contains Error")
	}

	if !Contains(tt, time.Date(2024, 2, 22, 22, 33, 33, 0, time.Local)) {
		t.Error("Contains Error")
	}

	if Contains(tt, time.Date(2024, 2, 22, 11, 11, 11, 0, time.Local)) {
		t.Error("Contains Error")
	}

	if Contains(tt, time.Date(2024, 2, 22, 44, 44, 44, 0, time.Local)) {
		t.Error("Contains Error")
	}
}

func TestContains2(t *testing.T) {
	tt, _ := NewTimeTerm(
		time.Date(2024, 2, 22, 22, 22, 22, 0, time.Local), true,
		time.Date(2024, 2, 22, 22, 33, 33, 0, time.Local), false,
	)

	if !Contains(tt, time.Date(2024, 2, 22, 22, 33, 22, 0, time.Local)) {
		t.Error("Contains Error")
	}

	if !Contains(tt, time.Date(2024, 2, 22, 22, 22, 22, 0, time.Local)) {
		t.Error("Contains Error")
	}

	if Contains(tt, time.Date(2024, 2, 22, 22, 33, 33, 0, time.Local)) {
		t.Error("Contains Error")
	}
}

func TestContains3(t *testing.T) {
	tt, _ := NewTimeTerm(
		time.Date(2024, 2, 22, 22, 22, 22, 0, time.Local), false,
		time.Date(2024, 2, 22, 22, 33, 33, 0, time.Local), true,
	)

	if !Contains(tt, time.Date(2024, 2, 22, 22, 33, 22, 0, time.Local)) {
		t.Error("Contains Error")
	}

	if Contains(tt, time.Date(2024, 2, 22, 22, 22, 22, 0, time.Local)) {
		t.Error("Contains Error")
	}

	if !Contains(tt, time.Date(2024, 2, 22, 22, 33, 33, 0, time.Local)) {
		t.Error("Contains Error")
	}
}

func TestAnyOneContains(t *testing.T) {
	var tt []TimeTerm
	tt1, _ := NewTimeTermJust(
		time.Date(2024, 2, 22, 22, 22, 22, 0, time.Local),
		time.Date(2024, 2, 22, 22, 33, 33, 0, time.Local),
	)
	tt = append(tt, tt1)
	tt2, _ := NewTimeTermJust(
		time.Date(2024, 2, 22, 22, 44, 44, 0, time.Local),
		time.Date(2024, 2, 22, 22, 55, 55, 0, time.Local),
	)
	tt = append(tt, tt2)

	if !AnyOneContains(tt, time.Date(2024, 2, 22, 22, 33, 22, 0, time.Local)) {
		t.Error("Contains Error")
	}
	if !AnyOneContains(tt, time.Date(2024, 2, 22, 22, 44, 55, 0, time.Local)) {
		t.Error("Contains Error")
	}
	if AnyOneContains(tt, time.Date(2024, 2, 22, 22, 11, 11, 0, time.Local)) {
		t.Error("Contains Error")
	}
}

func TestAllContains(t *testing.T) {
	var tt []TimeTerm
	tt1, _ := NewTimeTermJust(
		time.Date(2024, 2, 22, 22, 22, 22, 0, time.Local),
		time.Date(2024, 2, 22, 22, 33, 33, 0, time.Local),
	)
	tt = append(tt, tt1)
	tt2, _ := NewTimeTermJust(
		time.Date(2024, 2, 22, 22, 30, 30, 0, time.Local),
		time.Date(2024, 2, 22, 22, 55, 55, 0, time.Local),
	)
	tt = append(tt, tt2)

	if !AllContains(tt, time.Date(2024, 2, 22, 22, 30, 50, 0, time.Local)) {
		t.Error("Contains Error")
	}
	if AllContains(tt, time.Date(2024, 2, 22, 22, 33, 44, 0, time.Local)) {
		t.Error("Contains Error")
	}
	if AllContains(tt, time.Date(2024, 2, 22, 22, 44, 55, 0, time.Local)) {
		t.Error("Contains Error")
	}
	if AllContains(tt, time.Date(2024, 2, 22, 22, 11, 11, 0, time.Local)) {
		t.Error("Contains Error")
	}
}
