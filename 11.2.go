package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	pembaca := bufio.NewReader(os.Stdin)
	masukan, _ := pembaca.ReadString('\n')
	masukan = strings.TrimSpace(masukan)

	angka := strings.Split(masukan, " ")
	suara := make(map[int]int)
	totalSuara := -1 // -1 untuk tidak menghitung 0 di akhir
	suaraSah := 0

	// Hitung suara
	for _, num := range angka {
		totalSuara++
		nilai, err := strconv.Atoi(num)
		if err != nil || nilai == 0 {
			continue
		}

		if nilai >= 1 && nilai <= 20 {
			suara[nilai]++
			suaraSah++
		}
	}

	// Cari ketua dan wakil
	var suaraTertinggi, suaraTertinggiKedua int
	var ketua, wakil int

	// Tahap pertama: cari suara tertinggi
	for i := 1; i <= 20; i++ {
		if suara[i] > suaraTertinggi {
			suaraTertinggi = suara[i]
		}
	}

	// Cari ketua (nomor terkecil dengan suara tertinggi)
	for i := 1; i <= 20; i++ {
		if suara[i] == suaraTertinggi {
			ketua = i
			break
		}
	}

	// Cari suara tertinggi kedua
	for i := 1; i <= 20; i++ {
		if suara[i] > suaraTertinggiKedua && suara[i] <= suaraTertinggi && i != ketua {
			suaraTertinggiKedua = suara[i]
		}
	}

	// Cari wakil (nomor terkecil dengan suara tertinggi kedua)
	for i := 1; i <= 20; i++ {
		if suara[i] == suaraTertinggiKedua && i != ketua {
			wakil = i
			break
		}
	}

	fmt.Printf("Suara masuk: %d\n", totalSuara)
	fmt.Printf("Suara sah: %d\n", suaraSah)
	fmt.Printf("Ketua RT: %d\n", ketua)
	fmt.Printf("Wakil ketua: %d\n", wakil)
}
