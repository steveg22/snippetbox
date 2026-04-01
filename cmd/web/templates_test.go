package main

import (
	"testing"
	"time"

	"github.com/steveg22/snippetbox/internal/assert"
)

func TestHumanDate(t *testing.T) {
	dublin, _ := time.LoadLocation("Europe/Dublin")
	paris, _ := time.LoadLocation("Europe/Paris")
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2024, 3, 17, 10, 15, 0, 0, time.UTC),
			want: "17 Mar 2024 at 10:15",
		}, {
			name: "empty",
			tm:   time.Time{},
			want: "",
		}, {
			name: "Dublin (GMT)",
			//tm:   time.Date(2024, 3, 17, 10, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			tm:   time.Date(2024, 3, 17, 10, 15, 0, 0, dublin),
			want: "17 Mar 2024 at 10:15",
		}, {
			name: "Dublin (IST)",
			tm:   time.Date(2024, 4, 17, 10, 15, 0, 0, dublin),
			want: "17 Apr 2024 at 09:15",
		}, {
			name: "Paris (CET)",
			tm:   time.Date(2024, 3, 17, 10, 15, 0, 0, paris),
			want: "17 Mar 2024 at 09:15",
		},
	}
	for _, tt := range tests {
		// Use the t.Run() function to run a sub-test for each test case. The
		// first parameter to this is the name of the test (used to identify the
		// sub-test in any log output) and the second parameter is an anonymous
		// function containing the actual test for each case.
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			assert.Equal(t, hd, tt.want)
		})
	}
}
