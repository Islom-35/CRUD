package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Respons struct {
	Stores  []*Store
	Branch  []*Branch
	Vacancy []*Vacancy
}
type Store struct {
	ID       int
	Name     string
	Branches []*Branch
}
type Branch struct {
	ID          int
	Name        string
	PhoneNumber string
	StoreID     int
	Vacancies   []*Vacancy
}
type Vacancy struct {
	ID       int
	Position string
	Salary   int
	BranchId int
}

func connect() *sql.DB {
	connect := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "1234", "market")

	db, err := sql.Open("postgres", connect)

	if err != nil {
		panic(err)
	}
	return db
}
func insertStore(db *sql.DB, stores []Store) {
	for _, s := range stores {
		_, err := db.Exec("INSERT INTO store(id,name) Values($1, &2)", s.ID, s.Name)
		if err != nil {
			panic(err)
		}
	}
}
func insertBranch(db *sql.DB, stores []Store) {
	for _, s := range stores {
		for _, b := range s.Branches {
			_, err := db.Exec("INSERT INTO branch(id,name, phonenumber, store_id) VALUES"+
				"($1, $2, $3, $4)", b.ID, b.Name, b.PhoneNumber, b.StoreID)
			if err != nil {
				panic(err)
			}
		}
	}
}
func insertVacancy(db *sql.DB, stores []Store) {
	for _, s := range stores {
		for _, b := range s.Branches {
			for _, v := range b.Vacancies {
				_, err := db.Exec("INSERT INTO vacancy(id.position,salary,branch_id) VALUES"+
					"($1, $2, $3, $4)", v.ID, v.Position, v.Salary, v.BranchId)
				if err != nil {
					return
				}
			}
		}
	}
}

// func reloadDatabase(db *sql.DB) {
// 	_, err := db.Exec("Drop Table store")
// 		if err != nil {
// 			panic(err)
// 		}

//		_, err1 := db.Exec("CREATE Table store(id int,name varchar(64))")
//		if err1 != nil {
//			panic(err1)
//		}
//		_, err2 := db.Exec("Drop Table branch")
//		if err2 != nil {
//			panic(err)
//		}
//		_, err3 := db.Exec("CREATE Table branch(id int,name varchar(64),phonenumber text,store_id int)")
//		if err3 != nil {
//			panic(err1)
//		}
//		_, err4 := db.Exec("Drop Table vacancy")
//		if err != nil {
//			panic(err4)
//		}
//		_, err5 := db.Exec("CREATE Table vacancy(id int,name varchar(64)),salary int")
//		if err5 != nil {
//			panic(err5)
//		}
//	}
//
//	func deleteData(db *sql.DB) {
//		_, err1 := db.Exec("Delete From vacancy Where id= ($1)", 8)
//		if err1 != nil {
//			panic(err1)
//		}
//	}
func updateData(db *sql.DB) {
	branch := Branch{
		ID:   1,
		Name: "Abay",
	}
	_, err := db.Exec("Update branch SEt name=($1) Where id=($2)", branch.Name, branch.ID)
	if err != nil {
		panic(err)
	}
}
func getStoreData(db *sql.DB) {
	fmt.Println("    Table store")
	rows, err := db.Query("Select id,name From store")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			panic(err)
		}
		fmt.Printf("ID: %d,Name: %s\n", id, name)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
}

func getBranchData(db *sql.DB) {
	fmt.Println("  Table Branch")
	rows, err := db.Query("Select id, name, phonenumber,store_id from branch")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var phonenumber string
		var store_id int
		if err := rows.Scan(&id, &name, &phonenumber, &store_id); err != nil {
			panic(err)
		}
		fmt.Printf("ID: %d, Name: %s, Phone number: %v, Store_id: %d\n", id, name, phonenumber, store_id)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}

}
func getVacancyData(db *sql.DB) {
	fmt.Println("    Table Vacancy")
	rows, err := db.Query("Select id,position, salary, branch_id from vacancy")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var position string
		var salary int
		var branch_id int

		if err := rows.Scan(&id, &position, &branch_id); err != nil {
			panic(err)
		}
		fmt.Printf("ID: %d,Position: %s,Salary: %d, Branch_id:%d\n", id, position, salary, branch_id)
	}
}
func main() {
	db := connect()
	defer db.Close()
	stores := []Store{
		Store{
			ID:   1,
			Name: "Karzinka",
			Branches: []*Branch{
				&Branch{
					ID:          1,
					Name:        "Koxinur",
					PhoneNumber: "998919717124",
					StoreID:     1,
					Vacancies: []*Vacancy{
						&Vacancy{
							ID:       1,
							Position: "Driver",
							Salary:   300,
							BranchId: 1,
						},
						&Vacancy{
							ID:       2,
							Position: "Driver",
							Salary:   400,
							BranchId: 1,
						},
					},
				},
				&Branch{
					ID:          2,
					Name:        "Bunyodkor",
					PhoneNumber: "998919717124",
					StoreID:     1,
					Vacancies: []*Vacancy{
						&Vacancy{
							ID:       3,
							Position: "Driver",
							Salary:   300,
							BranchId: 2,
						},
						&Vacancy{
							ID:       4,
							Position: "Driver",
							Salary:   400,
							BranchId: 2,
						},
					},
				},
			},
		},
	}
	// reloadDatabase(db)
	connect()
	insertStore(db, stores)
	insertBranch(db, stores)
	insertVacancy(db, stores)
}
