package test

import (
	"github.com/wearetheledger/go-couchdb-query-engine/query"
	"testing"
)

func TestRegex(t *testing.T) {

	t.Run("Owner property should include B", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"owner": map[string]interface{}{
					"$regex": "^b",
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

	t.Run("Owner property should include letters", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"owner": map[string]interface{}{
					"$regex": "[a-z]",
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
