package prometheus

import (
	"testing"
	"time"
)

func TestLastOfLastMonth(t *testing.T) {
	testCases := []struct {
		input time.Time
		want  string
	}{
		{time.Date(2023, 3, 31, 0, 0, 0, 0, time.UTC), "2023-02-28"},
		{time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC), "2023-12-31"}, // 跨年
	}
	for _, tc := range testCases {
		got := time.Date(tc.input.Year(), tc.input.Month()-1, 1, 0, 0, 0, 0, tc.input.Location()).AddDate(0, 1, -1)
		if got.Format("2006-01-02") != tc.want {
			t.Errorf("输入 %v 期望 %s 实际 %s", tc.input, tc.want, got.Format("2006-01-02"))
		}
	}
}
