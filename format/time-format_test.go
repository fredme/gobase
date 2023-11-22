package format

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeFormat(t *testing.T) {
	timeStr := "2019-01-01 12:00:00"
	ti, err := time.Parse(TimeFormat, timeStr)
	if err != nil {
		t.Fatalf("time.Parse error")
	}
	fmt.Printf("ti.Location(): %v\n", ti.Location())                 // UTC
	fmt.Printf("time.Now().Location(): %v\n", time.Now().Location()) // Local

	timeInBeijing, err := time.ParseInLocation(TimeFormat, timeStr, beijing)
	if err != nil {
		t.Fatalf("time.Parse error")
	}
	fmt.Printf("timeInBeijing.Location(): %v\n", timeInBeijing.Location())                   // Beijing Time
	fmt.Printf("time.Now().In(beijing).Location(): %v\n", time.Now().In(beijing).Location()) // Beijing Time
}
