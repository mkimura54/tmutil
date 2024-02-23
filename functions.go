package tmutil

import "time"

// Contains は時間帯に指定時間が含まれているかどうかを判定する
func Contains(tt TimeTerm, t time.Time) bool {
	return ((tt.start.just && tt.start.value.Equal(t)) || tt.start.value.Before(t)) &&
		((tt.end.just && tt.end.value.Equal(t)) || tt.end.value.After(t))
}

// AnyOneContains は複数の時間帯のいずれかに指定時間が含まれているかどうかを判定する
func AnyOneContains(tts []TimeTerm, t time.Time) bool {
	for _, tt := range tts {
		if Contains(tt, t) {
			return true
		}
	}
	return false
}

// AllContains は複数の時間帯のすべてに指定時間が含まれているかどうかを判定する
func AllContains(tts []TimeTerm, t time.Time) bool {
	for _, tt := range tts {
		if !Contains(tt, t) {
			return false
		}
	}
	return true
}

// DuplicateTerm は2つの時間帯が重なる部分の時間帯を取得する
func DuplicateTerm(tt1, tt2 TimeTerm) (TimeTerm, bool) {
	if !tt1.valid || !tt2.valid {
		return TimeTerm{}, false
	}

	//         (1)     (2)      (3)           (4)                (5) (6)   (7)
	// tt1 : |-----|  |-----|  |-----|   |-----------|         |--|  |--|  |--|
	// tt2 :    |-------|       |--|   |---------------|          |--|     |--|
	if tt1.start.value.After(tt2.start.value) && tt1.end.value.After(tt2.start.value) &&
		tt1.start.value.Before(tt2.end.value) && tt1.end.value.Before(tt2.end.value) {
		return tt1, true // (4)
	} else if tt1.start.value.Before(tt2.start.value) && tt1.end.value.After(tt2.start.value) &&
		tt1.start.value.Before(tt2.end.value) && tt1.end.value.After(tt2.end.value) {
		return tt2, true // (3)
	} else if tt1.start.value.Before(tt2.end.value) && tt1.end.value.After(tt2.end.value) {
		return createNewTimeTerm(tt1.start, tt2.end), true // (2)
	} else if tt1.start.value.Before(tt2.start.value) && tt1.end.value.After(tt2.start.value) {
		return createNewTimeTerm(tt2.start, tt1.end), true // (1)
	} else {
		if tt1.start.just && tt1.end.just && tt2.start.just && tt2.end.just &&
			tt1.start.value.Equal(tt2.start.value) && tt1.end.value.Equal(tt2.end.value) {
			return createNewTimeTerm(tt1.start, tt2.end), true // (7)
		} else if tt1.start.just && tt2.end.just &&
			tt1.start.value.Equal(tt2.end.value) {
			return createNewTimeTerm(tt1.start, tt2.end), true // (6)
		} else if tt1.end.just && tt2.start.just &&
			tt1.end.value.Equal(tt2.start.value) {
			return createNewTimeTerm(tt1.end, tt2.start), true // (5)
		}
		return TimeTerm{}, false // 重なりなし
	}
}

// createNewTimeTerm は2つの日時の詳細な状態から新しい時間帯を生成する
func createNewTimeTerm(start, end timeState) TimeTerm {
	ntt, _ := NewTimeTerm(start.value, start.just, end.value, end.just)
	return ntt
}
