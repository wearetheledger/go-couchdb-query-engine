package test

import (
	"testing"

	"github.com/wearetheledger/go-couchdb-query-engine/query"
)

func TestImplicitAnd(t *testing.T) {

	t.Run("Element should be returned when owner bob & size 3 with implicit and and explicit eq", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"owner": map[string]interface{}{
					"$eq": "bob",
				},
				"size": map[string]interface{}{
					"$eq": 3,
				},
			},
		})

		if err != nil {
			t.Error(err)
		}

		if len(res) != 1 {
			t.Error("Query should have returned 1 results")
		}

	})

	t.Run("Element should be returned when owner bob & size 3 with implicit and and implicit eq", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"owner": "bob",
				"size":  3,
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
