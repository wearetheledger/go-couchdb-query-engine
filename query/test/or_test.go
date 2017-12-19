package test

import (
	"test/couchdbMockQuery/query"
	"testing"
)

func TestOr(t *testing.T) {

	t.Run("Element should be returned when owner bob or size 5", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
					"$or": []interface{}{
						map[string]interface{}{
							"owner": "bob",
						},
						map[string]interface{}{
							"size":  5,
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

	t.Run("Element should be returned when not equal owner bob & size 3", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
					"$or": []interface{}{
						map[string]interface{}{
							"owner": map[string]interface{}{
								"$ne": "alicia",
							},
						},
						map[string]interface{}{
							"size":  map[string]interface{}{
								"$ne": 20,
							},
						},
					},

			},
		})

		if err != nil {
			t.Error(err)
		}

		if len(res) != 3 {
			t.Error("Query should have returned 3 results")
		}
	})
}
