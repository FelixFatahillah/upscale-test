package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("soal_1_input.txt")
	if err != nil {
		fmt.Println("Gagal membaca file:", err)
		return
	}

	// Mengubah isi file menjadi slice string
	lines := strings.Split(string(data), "\n")

	angka := make([]int, 0)
	jumlah := 0

	for _, line := range lines {
		// Delete char \r
		line = strings.TrimRight(line, "\r")

		if line == "" {
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Gagal mengkonversi angka:", err)
			return
		}
		jumlah += num
		angka = append(angka, num)
	}

	fmt.Println("Total angka pada file:", len(angka))
	fmt.Println("Jumlah semua angka:", jumlah)
}
