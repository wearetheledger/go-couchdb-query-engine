package rules

import (
	"github.com/wearetheledger/go-couchdb-query-engine/query"
	"testing"
)

func TestLte(t *testing.T) {

	t.Run("Size should be lower than 3", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"size": map[string]interface{}{
					"$lte": 3,
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
}
