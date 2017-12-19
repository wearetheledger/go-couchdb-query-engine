package test

import (
	"test/couchdbMockQuery/query"
	"testing"
)

func TestNe(t *testing.T) {

	t.Run("Size should equal 3", func(t *testing.T) {
		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"size": map[string]interface{}{
					"$ne": 3,
				},
			},
		})

		if err != nil {
			t.Error(err)
		}

		if len(res) != 2 {
			t.Error("TestNe should have returned 2 result")
		}
	})

	t.Run("ObjectType should equal marble", func(t *testing.T) {
		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"objectType": map[string]interface{}{
					"$ne": "MARBLE",
				},
			},
		})

		if err != nil {
			t.Error(err)
		}

		if len(res) != 0 {
			t.Error("TestNeMultiple should have returned 0 results")
		}
	})
}