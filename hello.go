package main

import "fmt"

func main() {
	fmt.Println("Hello, saya sedang belajar Go")
	fmt.Print("Ini tidak menambahkan baris baru, haru ditambahkan \\n. \n")
	fmt.Println("Hello, saya sedang belajar Go")

	// ini adalah  komentar inline
	/*
	* ini adalah komentar multiline
	* komentar ini tidak akan dieksekusi
	 */

	/*
		dibawah ini membahas tentang variable digolang
		keyword var => digunakan untuk deklarasi variable
		var <nama-variable> <tipe-data>
		contoh : var firstName string => string adalah tipe data
	*/
	fmt.Println()
	var firstName string = "Rahmatulah"

	var lastName string

	lastName = "Sidik"

	fmt.Printf("Hello nama saya %s %s!\n", firstName, lastName)
	fmt.Println("Hello", firstName, lastName+"!")

	/*
		deklarasi variable tanpa tipe data
		mendeklarasikan variable tanpa tipe data, tidak usah menuliskan keyword var
		<nama-variable> := <isi-variable>

		diperbolehkan menggunakan keyword var, tpi ganti menggunakan  =
		var <nama-variable> = <nilai-variable>
	*/

	i := 10
	j := 2
	fmt.Println("i x j = ", i*j)

	var charA = "A"
	var charB = "B"

	fmt.Printf("Character %s, Character %s.\n", charA, charB)

	/*
		deklarasi multi variable
	*/
	fmt.Println()
	var first, second, third string

	first, second, third = "satu", "dua", "tiga"
	fmt.Println(first, second, third)

	/*
		Variable underscore
		jika ingin mendeklarasikan varibale, tpi tidak ingin digunakan,
		gunakan Underscore( _ ) => predifined variable
		_ = isi variable
		Perlu diketahui, bahwa isi variabel underscore tidak dapat ditampilkan.
		Data yang sudah masuk variabel tersebut akan hilang. Ibarat blackhole, sekali masuk,
		tidak akan bisa keluar :-)
	*/
	_ = "predifined variable"

}
