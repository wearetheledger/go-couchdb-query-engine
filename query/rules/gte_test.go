package rules

import (
	"github.com/wearetheledger/go-couchdb-query-engine/query"
	"testing"
)

func TestGte(t *testing.T) {

	t.Run("Size should be greater than or equal 3", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"size": map[string]interface{}{
					"$gte": 3,
				},
			},
		})

		if err != nil {
			t.Error(err)
		}

		if len(res) != 2 {
			t.Error("should have returned 2 result")
		}
	})

	t.Run("Size should be greater than or equal 1", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"size": map[string]interface{}{
					"$gte": 1,
				},
			},
		})

		if err != nil {
			t.Error(err)
		}

		if len(res) != 3 {
			t.Error("should have returned 3 result")
		}
	})
}
