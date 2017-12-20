package test

import (
	"github.com/wearetheledger/go-couchdb-query-engine/query"
	"testing"
)

func TestNin(t *testing.T) {

	t.Run("size should not equal any elements", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"size": map[string]interface{}{
					"$nin": []int{
						1000, 99,
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

	t.Run("owner should not equal any elements", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"owner": map[string]interface{}{
					"$nin": []string{
						"maria", "aaron",
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
