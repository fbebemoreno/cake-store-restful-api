# Cake Store RESTful API

This RESTful API was developed using **Golang** programming language. The features consist of :
- GET all cakes data and singular cake using ID.
- CREATE cake data.
- UPDATE cake data using ID.
- DELETE cake data using ID.

## Getting Started

The first thing you need to do is install **Golang**. If you haven't downloaded it yet, [Click Here](https://go.dev/doc/install) to go to **Golang**'s website and download it.

After that, run this code on your terminal : ```git clone https://github.com/fbebemoreno/cake-store-restful-api.git```

Keep in mind on `main.go` file, you need to put a port that is currently not in use by any server so you can run API's server. By default, the port in use is `:8080`, and in this file was changed to `:5000`.

```golang
func main() {
  ...
  router.Run(":5000") //ganti ":5000" dengan port yang tidak digunakan, biarkan kosong untuk ":8080"
}
```

After that, make a **Database** with the name `cake_store_restful_api`.
You can always change the name and server port of the database by changing it in `./models/setup.go` :

```golang
...
func ConnectDB() {
  database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/cake_store_restful_api")) //ganti nama dan port disini
  ...
}
```

Lastly, run `cd cake-store-restful-api` and run `go run main.go` on your terminal.

## Notes

If you're using **Postman** to test this API, sometimes `updated_at`, `created_at`, and `deleted_at` doesn't show the correct date data. Although, in the **Database**, the data type and structure are already correct.
