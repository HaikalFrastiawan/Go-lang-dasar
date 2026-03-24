package mysqldatabase

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('Haikal','HAIKAL')"
	_, err := db.ExecContext(ctx,script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert New Customer")
}


func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx,script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id: ", id)
		fmt.Println("Name:	", name)
	}

}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating , birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx,script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
    var id, name, email string
    var balance int32
    var rating float64
    var birthDate, createdAt time.Time
    var married bool

    err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
    if err != nil {
        panic(err)
    }
		fmt.Println("================")
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		fmt.Println("Email:", email)
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		fmt.Println("Birth Date:", birthDate)
		fmt.Println("Married:", married)
		fmt.Println("Created At:", createdAt)
}
defer rows.Close()

}

func TestSqlInjection(t *testing.T) {

	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"


	script := "SELECT username FROM user WHERE username = '"+username+"' AND password = '"+password+"' LIMIT 1"

	rows, err := db.QueryContext(ctx,script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	}else {
		fmt.Println("Gagal Login")
	}
}


func TestSqlInjectionSave(t *testing.T) {

	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "admin"


	script := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	fmt.Println(script)

	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	}else {
		fmt.Println("Gagal Login")
	}
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	username := "eko"
	password := "eko"

	script := "INSERT INTO user(username, password) VALUES(?,?)"
	_, err := db.ExecContext(ctx,script, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert New User")
}

//auto increment
func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	email := "haikal@gmail.com"
	coment := "Test komen"

	script := "INSERT INTO comments(email, comment) VALUES(?,?)"
	result, err := db.ExecContext(ctx,script, email, coment)
	if err != nil {
		panic(err)
	}

	insertid, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Success Insert New Cooment with id", insertid)
}

func TestPreperStetment(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO comments(email, comment) VALUES(?,?)"
	statement, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "haikal" + strconv.Itoa(i) + "@gmail.com"
		comment := "komentar ke " +strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("comment Id", id)

	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	
	script := "INSERT INTO comments(email, comment) VALUES(?,?)" 
	//do transaction
	for i := 0; i < 10; i++ {
		email := "haikal" + strconv.Itoa(i) + "@gmail.com"
		comment := "komentar ke " +strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("comment Id", id)

	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}

}