# Cake Store RESTful API

RESTful API ini dibuat dengan menggunakan bahasa pemrograman Golang. RESTful API ini memiliki fitur :
- Mengambil semua data cakes dan cake satuan dengan ID.
- Membuat data cake dengan ID.
- Menghapus data cake dengan ID

## Panduan Awal

Hal pertama yang perlu dilakukan adalah install Golang. Jika belum, [Klik Disini](https://go.dev/doc/install) untuk mendownload.

Selanjutnya, jalankan kode ```git clone https://github.com/fbebemoreno/cake-store-restful-api.git```

Perlu diperhatikan pada file `main.go` , masukkan port yang sedang tidak digunakan untuk menjalankan server API. 
Secara default port yang digunakan adalah `:8080` , dan di file ini diganti dengan `:5000`.

```golang
func main() {
  ...
  router.Run(":5000") //ganti ":5000" dengan port yang tidak digunakan, biarkan kosong untuk ":8080"
}
```

Lalu buat **Database** dengan nama `cake_store_restful_api`.
Nama file dan port server database dapat diganti dengan mengganti kode pada `./models/setup.go` :

```golang
...
func ConnectDB() {
  database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/cake_store_restful_api")) //ganti nama dan port disini
  ...
}
```

Lakukan `cd cake-store-restful-api` dan jalankan kode `go run main.go`.

## Notes

Bila mengirim data menggunakan **Postman**, terkadang data `updated_at`, `created_at`, dan `deleted_at` menampilkan data *date* yang sedikit berbeda. Namun pada **Database**, bentuk dan type data sudah sesuai.
