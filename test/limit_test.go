package test

import (
	"testing"

	"github.com/wearetheledger/go-couchdb-query-engine/query"
)

func TestLimit(t *testing.T) {

	t.Run("Element should be returned when not equal owner bob & limit 1", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"$and": []interface{}{
					map[string]interface{}{
						"owner": map[string]interface{}{
							"$ne": "bob",
						},
					},
					map[string]interface{}{
						"size": map[string]interface{}{
							"$ne": 3,
						},
					},
				},
			},
			"limit": 1,
		})

		if err != nil {
			t.Error(err)
		}

		if len(res) != 1 {
			t.Error("Query should have returned 2 results")
		}
	})

	t.Run("Element should be returned when not equal owner bob & limit 5", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"$and": []interface{}{
					map[string]interface{}{
						"owner": map[string]interface{}{
							"$ne": "bob",
						},
					},
					map[string]interface{}{
						"size": map[string]interface{}{
							"$ne": 3,
						},
					},
				},
			},
			"limit": 5,
		})

		if err != nil {
			t.Error(err)
		}

		if len(res) != 2 {
			t.Error("Query should have returned 2 results")
		}
	})
}
