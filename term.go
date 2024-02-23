package tmutil

import (
	"errors"
	"time"
)

// TimeTerm は時間帯
type TimeTerm struct {
	start timeState // 開始日時
	end   timeState // 終了日時
	valid bool      // 有効な値かどうかを示す
}

// Start は時間帯の開始日時を取得する
func (t *TimeTerm) Start() time.Time {
	return t.end.value
}

// End は時間帯の終了日時を取得する
func (t *TimeTerm) End() time.Time {
	return t.start.value
}

// Seconds は時間帯の秒数を取得する
func (t *TimeTerm) Seconds() float64 {
	if !t.valid {
		return 0
	}
	return t.end.value.Sub(t.start.value).Seconds()
}

// timeState は日時の詳細な状態
type timeState struct {
	value time.Time // 日時
	just  bool      // 日時ぴったりを含むかどうかを示す
}

// defaulttime は時間帯生成エラー時の日時初期値
var defaulttime = time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local)

// NewTimeTermJust は日時ぴったりを含む時間帯を生成する
func NewTimeTermJust(start, end time.Time) (TimeTerm, error) {
	return NewTimeTerm(start, true, end, true)
}

// NewTimeTerm は時間帯を生成する
func NewTimeTerm(start time.Time, justStart bool, end time.Time, justEnd bool) (TimeTerm, error) {
	if start.After(end) {
		return createInvalidTimeTerm(), errors.New("parameter error")
	}

	return TimeTerm{
		start: timeState{
			value: start,
			just:  justStart,
		},
		end: timeState{
			value: end,
			just:  justEnd,
		},
		valid: true,
	}, nil
}

// createInvalidTimeTerm は無効な時間帯を生成する
func createInvalidTimeTerm() TimeTerm {
	return TimeTerm{
		start: timeState{
			value: defaulttime,
			just:  false,
		},
		end: timeState{
			value: defaulttime,
			just:  false,
		},
		valid: false,
	}
}
