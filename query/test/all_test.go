package test

import (
	"test/couchdbMockQuery/query"
	"testing"
)

func TestAll(t *testing.T) {

	t.Run("should contain all elements", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"previousOwners": map[string]interface{}{
					"$all": []string{
						"alice",
						"donald",
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

	t.Run("should contain all elements", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"previousOwners": map[string]interface{}{
					"$all": []string{
						"alice",
						"donald",
						"tom",
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
