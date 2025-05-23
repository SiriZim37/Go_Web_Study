package migrations

import (
	"course-go/config"
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
)

func Migrate() {
	db := config.GetDB()
	m := gormigrate.New(
		db,
		gormigrate.DefaultOptions,
		[]*gormigrate.Migration{
			m1596813596CreateArticlesTable(),
			m1596889997CreateCategoriesTable(),
			m1596954993AddCategoryIDToArticles(),
			m1596977447CreateUsersTable(),
			m1597000245AddUserIDToArticles(),
		},
	)

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
}
