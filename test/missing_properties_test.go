package test

import (
	"testing"

	"github.com/wearetheledger/go-couchdb-query-engine/query"
)

func TestQueryWithSomeMissingProperties(t *testing.T) {

	t.Run("Element should be returned when properties are missing in other records", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"image": "https://example.com/image.png",
			},
		})

		if err != nil {
			t.Error(err)
		}

		if len(res) != 1 {
			t.Error("Query should have returned 1 results")
		}
	})
}
