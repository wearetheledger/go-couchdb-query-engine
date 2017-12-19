package test

import (
	"test/couchdbMockQuery/query"
	"testing"
)

func TestElematch(t *testing.T) {

	t.Run("Array should contains element with organizationId", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"verification": map[string]interface{}{
					"$elemMatch": map[string]interface{}{
						"organizationId": "marbles ID inc.",
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

	t.Run("Array should contains element with score greater than 6", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"verification": map[string]interface{}{
					"$elemMatch": map[string]interface{}{
						"score": map[string]interface{}{
							"$gt": 6,
						},
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

	t.Run("Array should contains element with score greater than 9", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"verification": map[string]interface{}{
					"$elemMatch": map[string]interface{}{
						"score": map[string]interface{}{
							"$gt": 9,
						},
					},
				},
			},
		})

		if err != nil {
			t.Error(err)
		}

		if len(res) != 0 {
			t.Error("Query should have returned 0 results")
		}
	})
}
