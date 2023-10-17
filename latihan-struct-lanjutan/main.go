package main

import "fmt"

type Siswa struct {
	Nama 		string
	Umur		int
	Kelas		string
	Parent	OrangTua
}

type OrangTua struct {
	Nama 		string
	Umur		int
}

type SiswaSlice []Siswa

func main() {
	siswa1 := Siswa{
		Nama: "Ali",
		Umur: 15,
		Kelas: "9A",
		Parent: OrangTua{
			Nama: "Budi",
			Umur: 40,
		},
	}

	siswa2 := Siswa{
		Nama: "David",
		Umur: 14,
		Kelas: "8B",
		Parent: OrangTua{
			Nama: "Citra",
			Umur: 39,
		},
	}

	siswa3 := struct{
		Nama	string
		Umur	int
		Kelas	string
		Parent struct{
			Nama	string
			Umur	int
		}
	}{
		Nama: "Fina",
		Umur: 16,
		Kelas: "10C",
		Parent: OrangTua{
			Nama: "Eva",
			Umur: 35,
		},
	}

	daftarSiswa := SiswaSlice{
		{Nama: "Eva", Umur: 12, Kelas: "6B", Parent: OrangTua{Nama: "Rudi", Umur: 40}},
		{Nama: "Faisal", Umur: 11, Kelas: "5A", Parent: OrangTua{Nama: "Dewi", Umur: 38}},
		{Nama: "Grace", Umur: 10, Kelas: "4C", Parent: OrangTua{Nama: "Hendro", Umur: 42}},
	}



fmt.Println("Informasi Siswa 1:")
fmt.Printf("Nama Siswa: %s, Umur: %d, Kelas: %s\n", siswa1.Nama, siswa1.Umur, siswa1.Kelas)
fmt.Printf("Orang Tua: %s, Umur: %d\n", siswa1.Parent.Nama, siswa1.Parent.Umur)

fmt.Println("Informasi Siswa 2:")
fmt.Printf("Nama Siswa: %s, Umur: %d, Kelas: %s\n", siswa2.Nama, siswa2.Umur, siswa2.Kelas)
fmt.Printf("Orang Tua: %s, Umur: %d\n", siswa2.Parent.Nama, siswa2.Parent.Umur)

fmt.Println("Informasi Siswa 3:")
fmt.Printf("Nama Siswa: %s, Umur: %d, Kelas: %s\n", siswa3.Nama, siswa3.Umur, siswa3.Kelas)
fmt.Printf("Orang Tua: %s, Umur: %d\n", siswa3.Parent.Nama, siswa3.Parent.Umur)



fmt.Println("Daftar Siswa:")
for _, siswa := range daftarSiswa {
	fmt.Printf("Nama Siswa: %s, Umur: %d, Kelas: %s\n", siswa.Nama, siswa.Umur, siswa.Kelas)
	fmt.Printf("Orang Tua: %s, Umur: %d\n", siswa.Parent.Nama, siswa.Parent.Umur)
}
}