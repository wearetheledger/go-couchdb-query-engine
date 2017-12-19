package test

import (
	"test/couchdbMockQuery/query"
	"testing"
)

func TestExists(t *testing.T) {

	t.Run("Owner property should exist", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"owner": map[string]interface{}{
					"$exists": true,
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

	t.Run("ownership property should not exist", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"ownership": map[string]interface{}{
					"$exists": false,
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

	t.Run("verification property should exist on some", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"verification": map[string]interface{}{
					"$exists": true,
				},
			},
		})

		if err != nil {
			t.Error(err)
		}

		if len(res) != 2 {
			t.Errorf("Query should have returned 2 results, returned %d", len(res))
		}
	})
}
