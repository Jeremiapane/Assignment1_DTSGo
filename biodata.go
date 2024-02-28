package main

import (
	"fmt"
	"os"
	"strconv"
)

type Biodata struct {
    nomorUrut        int
    namaPeserta      string
    alamatPeserta    string
    pekerjaanPeserta string
    alasanPeserta    string
}

func getBiodata(absen int) (Biodata, bool) {
    dataPeserta := map[int]Biodata{
        1: {1, "Agus", "Bandung", "Mahasiswa", "Karena bahasa go sedang populer"},
        2: {2, "Tina", "Jakarta", "Lulusan Baru", "Pengen kerja tapi harus bisa bahasa go"},
        3: {3, "Irwan", "Bekasi", "Mahasiswa", "Kawan-kawan udah pada jago ngoding c++"},
        4: {4, "Freya", "Sulawesi", "Mahasiswi", "Di kampus lagi belajar go"},
        5: {5, "Praroroo", "Subang", "Pekerja Swasta", "Ingin menguasai bahasa go "},
    }

    if absen > len(dataPeserta) {
        return Biodata{}, false
    }

    return dataPeserta[absen], true
}

func main() {
    args := os.Args
    if len(args) < 2 {
        fmt.Println("Usage: go run biodata.go <no_urut>")
        return
    }

    absen := os.Args[1]
    dataAbsen, err := strconv.Atoi(absen)
    if err != nil {
        fmt.Println("Mohon masukan no urut dengan bilangan bulat")
        return
    }


    data, exists := getBiodata(dataAbsen)
    if !exists {
        fmt.Println("Data belum tersedia")
        return
    }

    fmt.Println("Data Peserta Golang")
    fmt.Println("No Urut	: ", data.nomorUrut)
    fmt.Println("Nama		: ", data.namaPeserta)
    fmt.Println("Alamat		: ", data.alamatPeserta)
    fmt.Println("Pekerjaan	: ", data.pekerjaanPeserta)
    fmt.Println("Alasan		: ", data.alasanPeserta)
}
