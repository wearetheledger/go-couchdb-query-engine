package test

import (
	"testing"

	"github.com/wearetheledger/go-couchdb-query-engine/query"
)

func TestSkip(t *testing.T) {
	t.Run("Element should be returned when not equal owner bob & skip 1", func(t *testing.T) {

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
			"skip": 1,
		})

		if err != nil {
			t.Error(err)
		}

		if len(res) != 1 {
			t.Log(len(res))
			t.Error("Query should have returned 2 results")
		}
	})

	t.Run("Element should be returned when not equal owner bob & skip 5", func(t *testing.T) {

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
			"skip": 5,
		})

		if err != nil {
			t.Error(err)
		}

		if len(res) != 0 {
			t.Error("Query should have returned 2 results")
		}
	})
}
