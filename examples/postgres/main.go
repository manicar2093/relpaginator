package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-rel/postgres"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	_ "github.com/lib/pq"
	"github.com/manicar2093/relpaginator"
)

const (
	UserTable = "users"
)

type User struct {
	ID        int
	Name      string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c User) String() string {
	return fmt.Sprintf("User{Name: %s, Age: %d}", c.Name, c.Age)
}

func (c User) Table() string {
	return UserTable
}

func main() {
	ctx := context.TODO()
	adapter, err := postgres.Open("postgres://development:development@localhost:3456/relpaginatortest?sslmode=disable")

	if err != nil {
		log.Fatalln(err)
	}

	db := rel.New(adapter)
	if err := createTable(ctx, db); err != nil {
		log.Fatalln(err)
	}
	if err := populateDB(ctx, db); err != nil {
		log.Fatalln(err)
	}

	// You can create a
	paginator := relpaginator.NewPaginator(db, &relpaginator.RelPaginatorConfig{
		PageSize: 4,
	})

	usersFoundPage := []User{}
	page, err := paginator.CreatePagination(ctx, UserTable, &usersFoundPage, &relpaginator.PageSort{
		Page: 1,
	})
	if err != nil {
		log.Fatalln(err)
	}

	data := page.Data.(*[]User)

	log.Println("Found quantity: ", len(*data))
	for _, u := range *data {
		log.Println(u)
	}

}

func createTable(ctx context.Context, db rel.Repository) error {
	if _, _, err := db.Exec(ctx, `create table if not exists public.users (
		id integer,
		name varchar(255),
		age integer,
		created_at timestamp(3),
		updated_at timestamp(3)
	);`); err != nil {
		return err
	}

	return nil
}

func populateDB(ctx context.Context, db rel.Repository) error {
	if count := db.MustCount(ctx, UserTable, where.Eq("name", "name1")); count == 0 {
		users := []User{
			{Name: "name1", Age: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Name: "name2", Age: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Name: "name3", Age: 65, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Name: "name4", Age: 33, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Name: "name5", Age: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Name: "name6", Age: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Name: "name7", Age: 40, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Name: "name8", Age: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Name: "name9", Age: 22, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Name: "name10", Age: 75, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Name: "name11", Age: 25, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Name: "name12", Age: 56, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		}
		return db.InsertAll(ctx, &users)
	}
	return nil
}
