package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = ""
	password = ""
	dname    = ""
)

var db *sql.DB

func init() {
	var err error
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dname)
	db, err = sql.Open("postgres", connString)
	//db.SetMaxIdleConns(5)//Max 5 bağlantı
	//db.SetMaxOpenConns(10)//Anlık eş zamanlı db connection var
	//db.SetConnMaxIdleTime(1*time.Second)// Bir bağlantı kaç dk boşta kalabilir
	//db.SetConnMaxLifetime(30*time.Second)//Bağlantı maximum açık kalacak süresi
	if err != nil {
		panic(err)
	}

}

type Product struct {
	ID                 int
	Title, Description string
	Price              float32
}

func InsertProduct(data Product) {
	result, err := db.Exec("INSERT INTO products(title,description,price) VALUES($1,$2,$3)", data.Title, data.Description, data.Price)
	if err != nil {
		panic(err)
	}
	rowsAffected, err := result.RowsAffected()
	fmt.Printf("Eklenen Kayıt sayısı(%d)", rowsAffected)
}

func UpdateProcut(data Product) {
	result, err := db.Exec("UPDATE products SET title=$1,description=$2,price=$3 WHERE id=$4", data.Title, data.Description, data.Price, data.ID)
	if err != nil {
		panic(err)
	}
	rowsAffected, err := result.RowsAffected()
	fmt.Printf("Kayıt güncellendi(%d)", rowsAffected)
}

func GetProducts() {
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Kayıt bulunamadı")
			return
		}
		panic(err)

	}
	defer rows.Close()
	var products []*Product
	for rows.Next() {//Sonraki satıra geçme.
		prd := &Product{}
		err := rows.Scan(&prd.ID, &prd.Title, &prd.Description, &prd.Price)//Select alınan veriyi bunun içerisine doldur sonra dolduracağım
		if err != nil {
			panic(err)
		}
		products = append(products, prd)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	for _, pr := range products {
		fmt.Printf("%d %s %s %.2f\n", pr.ID, pr.Title, pr.Description, pr.Price)
	}
}

func GetProductsByID(id int) {
	var product string
	err := db.QueryRow("SELECT title FROM products WHERE id=$1", id).Scan(&product)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("Kayıt bulunamadı")
	case err != nil:
		panic(err)
	default:
		fmt.Println("Product :", product)
	}
}
