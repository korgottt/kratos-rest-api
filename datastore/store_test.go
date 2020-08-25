package datastore

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDBConnection(t *testing.T) {
	asserts := assert.New(t)
	dbStore := IdentitiesDBStore{}

	asserts.NoError(dbStore.Init(), "init error")
	asserts.NoError(dbStore.db.Ping(), "something get wrong")
	asserts.NoError(dbStore.db.Close(), "closing error")

}

// func TestDBArticleStore(t *testing.T) {

// 	dbStore := initDB(t)
// 	testTitle := "My test article description"
// 	testSlug := utils.CreateSlug(testTitle)

// 	defer closeDB(t, dbStore)
// 	defer deleteTestData(dbStore, "article", fmt.Sprintf("title LIKE '%s'", testTitle))

// 	t.Run("test get if not exists", func(t *testing.T) {
// 		a, _ := dbStore.GetArticle(testSlug)
// 		assert.Equal(t, model.Article{}, a)
// 	})

// 	t.Run("test create and get if exists", func(t *testing.T) {
// 		inputArticle, _ := dbStore.CreateArticle(model.SingleArticleWrap{model.Article{Title: testTitle}})
// 		outputArticle, _ := dbStore.GetIdentityByID(testSlug)

// 		assert.Equal(t, inputArticle.Slug, outputArticle.Slug)
// 	})
// }

func initDB(t *testing.T) *IdentitiesDBStore {
	t.Helper()
	db := IdentitiesDBStore{}
	if err := db.Init(); err != nil {
		assert.FailNow(t, fmt.Sprintf("init error: %s", err))
	}
	return &db
}

func closeDB(t *testing.T, db *IdentitiesDBStore) {
	t.Helper()

	if err := db.Close(); err != nil {
		assert.FailNow(t, "db connection was not closed")
	}
}

func deleteTestData(db *IdentitiesDBStore, table, filter string) {
	db.db.Exec(fmt.Sprintf("DELETE FROM %s WHERE %s", table, filter))
}
