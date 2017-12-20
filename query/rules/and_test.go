package rules

import (
	"github.com/wearetheledger/go-couchdb-query-engine/query"
	"testing"
)

func TestAnd(t *testing.T) {

	t.Run("Element should be returned when owner bob & size 3", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
					"$and": []interface{}{
						map[string]interface{}{
							"owner": "bob",
						},
						map[string]interface{}{
							"size":  3,
						},
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

	t.Run("Element should be returned when not equal owner bob & size 3", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
					"$and": []interface{}{
						map[string]interface{}{
							"owner": map[string]interface{}{
								"$ne": "bob",
							},
						},
						map[string]interface{}{
							"size":  map[string]interface{}{
								"$ne": 3,
							},
						},
					},

			},
		})

		if err != nil {
			t.Error(err)
		}

		if len(res) != 2 {
			t.Error("Query should have returned 2 results")
		}
	})
}
