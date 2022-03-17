package _map

import (
	"fmt"
	"time"
)

func GuessWhatDayTodayIs(s string) {

	days := make(map[string]int)

	days["Monday"] = 1
	days["Tuesday"] = 2
	days["Wednesday"] = 3
	days["Thursday"] = 4
	days["Friday"] = 5
	days["Saturday"] = 6
	days["Sunday"] = 7

	day := time.Now().Weekday()

	if days[s] == int(day) {
		fmt.Println(day, "correct")
	} else {
		fmt.Println(day, "wrong")
	}
}

func PrintDays() {

	days := make(map[string]int)

	days["Monday"] = 1
	days["Tuesday"] = 2
	days["Wednesday"] = 3
	days["Thursday"] = 4
	days["Friday"] = 5
	days["Saturday"] = 6
	days["Sunday"] = 7

	for key, value := range days {
		fmt.Println(key, value)
	}
}
