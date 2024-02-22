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
