package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Contoh inputan "02:00:00 PM"
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan waktu dalam format 12 jam (HH:MM:SS AM/PM): ")
	timeString, _ := reader.ReadString('\n')
	timeString = strings.TrimSpace(timeString)

	if !strings.Contains(timeString, "AM") && !strings.Contains(timeString, "PM") {
		timeString += " AM"
	}

	parsedTime, err := time.Parse("03:04:05 PM", timeString)
	if err != nil {
		fmt.Println("Gagal parsing waktu:", err)
		return
	}

	// Konversi format 24 jam
	time24Format := parsedTime.Format("15:04:05")

	fmt.Println("Waktu dalam format 24 jam:", time24Format)
}
