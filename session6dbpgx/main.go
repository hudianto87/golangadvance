package main

import (
	"belajargolangpart2/session6dbpgx/entity"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	//namadrive,user,password,ip,port,nama db
	dsn := "postgresql://postgres:P4ssw0rd@192.168.26.50:5432/traininggolang"
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connetted to db")

	//ambil data row
	var usr entity.User

	err = pool.QueryRow(ctx, "select id,name from users order by id desc limit 1").Scan(&usr.ID, &usr.Name)

	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println("user retrieve", usr)

	//exec untuk insert data
	_, err = pool.Exec(ctx, "insert into users(name,email,password,created_at,updated_at) values('budi','test@gmail.com','budi',NOW(),NOW())")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("data successfully inserted")

	//query untuk ambil semua data
	var users []entity.User
	rows, err := pool.Query(ctx, "select id,name from users order by id desc")
	if err != nil {
		log.Panicln(err)
	}

	for rows.Next() {
		var user entity.User
		rows.Scan(&user.ID, &user.Name)
		users = append(users, user)
	}

	fmt.Println("all user retrieve", users)
}
