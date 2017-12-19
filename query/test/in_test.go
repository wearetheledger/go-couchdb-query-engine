package test

import (
	"github.com/wearetheledger/go-couchdb-query-engine/query"
	"testing"
)

func TestIn(t *testing.T) {

	t.Run("size should equal one element in array", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"size": map[string]interface{}{
					"$in": []int{
						1, 3,
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

	t.Run("should contain all elements", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"owner": map[string]interface{}{
					"$in": []string{
						"alice", "bob", "arnold",
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
