package main

import "fmt"

func avgNum(numbers ...float32) {
	length := len(numbers)

	var sum float32
	for i := 0; i < length; i++ {
		sum += numbers[i]
	}
	avg := sum / float32(length)

	fmt.Printf("Jumlah nilai siswa dalam kelas:%d\n", length)
	fmt.Printf("Nilai rata-rata siswa dalam kelas:%f\n", avg)
}

func main() {
	nilaiSiswa := []float32{85.5, 92.0, 78.5, 90.0, 88.5}
	avgNum(nilaiSiswa...)
}