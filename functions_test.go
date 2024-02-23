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
		t.Error("TestNewTimeTermJust Error")
	}
	if !tt.valid ||
		tt.start.value.Format(time.DateTime) != st.Format(time.DateTime) ||
		tt.end.value.Format(time.DateTime) != et.Format(time.DateTime) {
		t.Error("TestNewTimeTermJust Error")
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

func TestSeconds(t *testing.T) {
	st := time.Date(2024, 2, 22, 22, 22, 22, 0, time.Local)
	et := time.Date(2024, 2, 22, 22, 33, 33, 0, time.Local)

	tt1, _ := NewTimeTermJust(st, et)
	if tt1.Seconds() != 11*60+11 {
		t.Error("TestSeconds Error")
	}

	tt2, _ := NewTimeTermJust(et, st)
	if tt2.Seconds() != 0 {
		t.Error("TestSeconds Error")
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

func TestDuplicateTerm1(t *testing.T) {
	tt1a, _ := NewTimeTermJust(
		time.Date(2024, 2, 22, 10, 0, 0, 0, time.Local),
		time.Date(2024, 2, 22, 11, 0, 0, 0, time.Local),
	)
	tt1b, _ := NewTimeTermJust(
		time.Date(2024, 2, 22, 13, 0, 0, 0, time.Local),
		time.Date(2024, 2, 22, 14, 0, 0, 0, time.Local),
	)
	tt1c, _ := NewTimeTermJust(
		time.Date(2024, 2, 22, 15, 0, 0, 0, time.Local),
		time.Date(2024, 2, 22, 16, 0, 0, 0, time.Local),
	)
	tt1d, _ := NewTimeTermJust(
		time.Date(2024, 2, 22, 17, 0, 0, 0, time.Local),
		time.Date(2024, 2, 22, 18, 0, 0, 0, time.Local),
	)

	tt2a, _ := NewTimeTermJust(
		time.Date(2024, 2, 22, 10, 30, 0, 0, time.Local),
		time.Date(2024, 2, 22, 13, 30, 0, 0, time.Local),
	)
	tt2b, _ := NewTimeTermJust(
		time.Date(2024, 2, 22, 15, 20, 0, 0, time.Local),
		time.Date(2024, 2, 22, 15, 40, 0, 0, time.Local),
	)
	tt2c, _ := NewTimeTermJust(
		time.Date(2024, 2, 22, 16, 40, 0, 0, time.Local),
		time.Date(2024, 2, 22, 18, 20, 0, 0, time.Local),
	)
	tt2d, _ := NewTimeTermJust(
		time.Date(2024, 2, 22, 11, 0, 0, 0, time.Local),
		time.Date(2024, 2, 22, 13, 0, 0, 0, time.Local),
	)
	tt2e, _ := NewTimeTermJust(
		time.Date(2024, 2, 22, 10, 0, 0, 0, time.Local),
		time.Date(2024, 2, 22, 11, 0, 0, 0, time.Local),
	)

	d1, d1r := DuplicateTerm(tt1a, tt2a)
	d2, d2r := DuplicateTerm(tt1b, tt2a)
	d3, d3r := DuplicateTerm(tt1c, tt2b)
	d4, d4r := DuplicateTerm(tt1d, tt2c)
	_, d5r := DuplicateTerm(tt1a, tt2b)
	d6, d6r := DuplicateTerm(tt1a, tt2d)
	d7, d7r := DuplicateTerm(tt1b, tt2d)
	d8, d8r := DuplicateTerm(tt1a, tt2e)
	if !(isSameDateTime(d1.Start(), tt2a.Start()) && isSameDateTime(d1.End(), tt1a.End())) {
		t.Error("TestDuplicateTerm Error")
	}
	if !(isSameDateTime(d2.Start(), tt1b.Start()) && isSameDateTime(d2.End(), tt2a.End())) {
		t.Error("TestDuplicateTerm Error")
	}
	if !(isSameDateTime(d3.Start(), tt2b.Start()) && isSameDateTime(d3.End(), tt2b.End())) {
		t.Error("TestDuplicateTerm Error")
	}
	if !(isSameDateTime(d4.Start(), tt1d.Start()) && isSameDateTime(d4.End(), tt1d.End())) {
		t.Error("TestDuplicateTerm Error")
	}
	if !d1r || !d2r || !d3r || !d4r || d5r || !d6r || !d7r || !d8r {
		t.Error("TestDuplicateTerm Error")
	}
	if !(isSameDateTime(d6.Start(), tt1a.End()) && isSameDateTime(d6.End(), tt1a.End())) {
		t.Error("TestDuplicateTerm Error")
	}
	if !(isSameDateTime(d7.Start(), tt1b.Start()) && isSameDateTime(d7.End(), tt1b.Start())) {
		t.Error("TestDuplicateTerm Error")
	}
	if !(isSameDateTime(d8.Start(), tt1a.Start()) && isSameDateTime(d8.End(), tt1a.End())) {
		t.Error("TestDuplicateTerm Error")
	}
}

func TestDuplicateTerm2(t *testing.T) {
	tt1a, _ := NewTimeTerm(
		time.Date(2024, 2, 22, 10, 0, 0, 0, time.Local), false,
		time.Date(2024, 2, 22, 11, 0, 0, 0, time.Local), false,
	)
	tt1b, _ := NewTimeTerm(
		time.Date(2024, 2, 22, 13, 0, 0, 0, time.Local), false,
		time.Date(2024, 2, 22, 14, 0, 0, 0, time.Local), false,
	)
	tt1c, _ := NewTimeTerm(
		time.Date(2024, 2, 22, 15, 0, 0, 0, time.Local), false,
		time.Date(2024, 2, 22, 16, 0, 0, 0, time.Local), false,
	)
	tt1d, _ := NewTimeTerm(
		time.Date(2024, 2, 22, 17, 0, 0, 0, time.Local), false,
		time.Date(2024, 2, 22, 18, 0, 0, 0, time.Local), false,
	)

	tt2a, _ := NewTimeTerm(
		time.Date(2024, 2, 22, 10, 30, 0, 0, time.Local), false,
		time.Date(2024, 2, 22, 13, 30, 0, 0, time.Local), false,
	)
	tt2b, _ := NewTimeTerm(
		time.Date(2024, 2, 22, 15, 20, 0, 0, time.Local), false,
		time.Date(2024, 2, 22, 15, 40, 0, 0, time.Local), false,
	)
	tt2c, _ := NewTimeTerm(
		time.Date(2024, 2, 22, 16, 40, 0, 0, time.Local), false,
		time.Date(2024, 2, 22, 18, 20, 0, 0, time.Local), false,
	)
	tt2d, _ := NewTimeTerm(
		time.Date(2024, 2, 22, 11, 0, 0, 0, time.Local), false,
		time.Date(2024, 2, 22, 13, 0, 0, 0, time.Local), false,
	)
	tt2e, _ := NewTimeTerm(
		time.Date(2024, 2, 22, 10, 0, 0, 0, time.Local), false,
		time.Date(2024, 2, 22, 11, 0, 0, 0, time.Local), false,
	)

	d1, d1r := DuplicateTerm(tt1a, tt2a)
	d2, d2r := DuplicateTerm(tt1b, tt2a)
	d3, d3r := DuplicateTerm(tt1c, tt2b)
	d4, d4r := DuplicateTerm(tt1d, tt2c)
	_, d5r := DuplicateTerm(tt1a, tt2b)
	_, d6r := DuplicateTerm(tt1a, tt2d)
	_, d7r := DuplicateTerm(tt1b, tt2d)
	_, d8r := DuplicateTerm(tt1a, tt2e)
	if !(isSameDateTime(d1.Start(), tt2a.Start()) && isSameDateTime(d1.End(), tt1a.End())) {
		t.Error("TestDuplicateTerm Error")
	}
	if !(isSameDateTime(d2.Start(), tt1b.Start()) && isSameDateTime(d2.End(), tt2a.End())) {
		t.Error("TestDuplicateTerm Error")
	}
	if !(isSameDateTime(d3.Start(), tt2b.Start()) && isSameDateTime(d3.End(), tt2b.End())) {
		t.Error("TestDuplicateTerm Error")
	}
	if !(isSameDateTime(d4.Start(), tt1d.Start()) && isSameDateTime(d4.End(), tt1d.End())) {
		t.Error("TestDuplicateTerm Error")
	}
	if !d1r || !d2r || !d3r || !d4r || d5r || d6r || d7r || d8r {
		t.Error("TestDuplicateTerm Error")
	}
}

func isSameDateTime(t1, t2 time.Time) bool {
	ts1 := t1.Format(time.DateTime)
	ts2 := t2.Format(time.DateTime)
	return ts1 == ts2
}

func TestIsDuplicateTerm1(t *testing.T) {
	tt1a, _ := NewTimeTermJust(
		time.Date(2024, 2, 22, 10, 0, 0, 0, time.Local),
		time.Date(2024, 2, 22, 11, 0, 0, 0, time.Local),
	)
	tt1b, _ := NewTimeTermJust(
		time.Date(2024, 2, 22, 13, 0, 0, 0, time.Local),
		time.Date(2024, 2, 22, 14, 0, 0, 0, time.Local),
	)
	tt1c, _ := NewTimeTermJust(
		time.Date(2024, 2, 22, 15, 0, 0, 0, time.Local),
		time.Date(2024, 2, 22, 16, 0, 0, 0, time.Local),
	)
	tt1d, _ := NewTimeTermJust(
		time.Date(2024, 2, 22, 17, 0, 0, 0, time.Local),
		time.Date(2024, 2, 22, 18, 0, 0, 0, time.Local),
	)

	tt2a, _ := NewTimeTermJust(
		time.Date(2024, 2, 22, 10, 30, 0, 0, time.Local),
		time.Date(2024, 2, 22, 13, 30, 0, 0, time.Local),
	)
	tt2b, _ := NewTimeTermJust(
		time.Date(2024, 2, 22, 15, 20, 0, 0, time.Local),
		time.Date(2024, 2, 22, 15, 40, 0, 0, time.Local),
	)
	tt2c, _ := NewTimeTermJust(
		time.Date(2024, 2, 22, 16, 40, 0, 0, time.Local),
		time.Date(2024, 2, 22, 18, 20, 0, 0, time.Local),
	)
	tt2d, _ := NewTimeTermJust(
		time.Date(2024, 2, 22, 11, 0, 0, 0, time.Local),
		time.Date(2024, 2, 22, 13, 0, 0, 0, time.Local),
	)
	tt2e, _ := NewTimeTermJust(
		time.Date(2024, 2, 22, 10, 0, 0, 0, time.Local),
		time.Date(2024, 2, 22, 11, 0, 0, 0, time.Local),
	)

	d1 := IsDuplicateTerm(tt1a, tt2a)
	d2 := IsDuplicateTerm(tt1b, tt2a)
	d3 := IsDuplicateTerm(tt1c, tt2b)
	d4 := IsDuplicateTerm(tt1d, tt2c)
	d5 := IsDuplicateTerm(tt1a, tt2b)
	d6 := IsDuplicateTerm(tt1a, tt2d)
	d7 := IsDuplicateTerm(tt1b, tt2d)
	d8 := IsDuplicateTerm(tt1a, tt2e)
	if !d1 || !d2 || !d3 || !d4 || d5 || !d6 || !d7 || !d8 {
		t.Error("TestIsDuplicateTerm Error")
	}
}

func TestIsDuplicateTerm2(t *testing.T) {
	tt1a, _ := NewTimeTerm(
		time.Date(2024, 2, 22, 10, 0, 0, 0, time.Local), false,
		time.Date(2024, 2, 22, 11, 0, 0, 0, time.Local), false,
	)
	tt1b, _ := NewTimeTerm(
		time.Date(2024, 2, 22, 13, 0, 0, 0, time.Local), false,
		time.Date(2024, 2, 22, 14, 0, 0, 0, time.Local), false,
	)
	tt1c, _ := NewTimeTerm(
		time.Date(2024, 2, 22, 15, 0, 0, 0, time.Local), false,
		time.Date(2024, 2, 22, 16, 0, 0, 0, time.Local), false,
	)
	tt1d, _ := NewTimeTerm(
		time.Date(2024, 2, 22, 17, 0, 0, 0, time.Local), false,
		time.Date(2024, 2, 22, 18, 0, 0, 0, time.Local), false,
	)

	tt2a, _ := NewTimeTerm(
		time.Date(2024, 2, 22, 10, 30, 0, 0, time.Local), false,
		time.Date(2024, 2, 22, 13, 30, 0, 0, time.Local), false,
	)
	tt2b, _ := NewTimeTerm(
		time.Date(2024, 2, 22, 15, 20, 0, 0, time.Local), false,
		time.Date(2024, 2, 22, 15, 40, 0, 0, time.Local), false,
	)
	tt2c, _ := NewTimeTerm(
		time.Date(2024, 2, 22, 16, 40, 0, 0, time.Local), false,
		time.Date(2024, 2, 22, 18, 20, 0, 0, time.Local), false,
	)
	tt2d, _ := NewTimeTerm(
		time.Date(2024, 2, 22, 11, 0, 0, 0, time.Local), false,
		time.Date(2024, 2, 22, 13, 0, 0, 0, time.Local), false,
	)
	tt2e, _ := NewTimeTerm(
		time.Date(2024, 2, 22, 10, 0, 0, 0, time.Local), false,
		time.Date(2024, 2, 22, 11, 0, 0, 0, time.Local), false,
	)

	d1 := IsDuplicateTerm(tt1a, tt2a)
	d2 := IsDuplicateTerm(tt1b, tt2a)
	d3 := IsDuplicateTerm(tt1c, tt2b)
	d4 := IsDuplicateTerm(tt1d, tt2c)
	d5 := IsDuplicateTerm(tt1a, tt2b)
	d6 := IsDuplicateTerm(tt1a, tt2d)
	d7 := IsDuplicateTerm(tt1b, tt2d)
	d8 := IsDuplicateTerm(tt1a, tt2e)
	if !d1 || !d2 || !d3 || !d4 || d5 || d6 || d7 || d8 {
		t.Error("TestIsDuplicateTerm Error")
	}
}
