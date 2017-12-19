package test

import (
	"test/couchdbMockQuery/query"
	"testing"
)

func TestSize(t *testing.T) {

	t.Run("PreviousOwners size should equal 2", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"previousOwners": map[string]interface{}{
					"$size": 2,
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
}
