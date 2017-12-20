package test

import (
	"github.com/wearetheledger/go-couchdb-query-engine/query"
	"testing"
)

func TestEq(t *testing.T) {
	t.Run("Size should equal 3", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"size": map[string]interface{}{
					"$eq": 3,
				},
			},
		})

		if err != nil {
			t.Error(err)
		}

		if len(res) != 1 {
			t.Error("Query should have returned 1 result")
		}
	})

	t.Run("ObjectType should equal marble", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"objectType": map[string]interface{}{
					"$eq": "MARBLE",
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

func TestImplicitEq(t *testing.T) {
	t.Run("Size should equal 3", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"size": 3,
			},
		})

		if err != nil {
			t.Error(err)
		}

		if len(res) != 1 {
			t.Error("Query should have returned 1 result")
		}
	})

	t.Run("previousOwners should equal array with alice & donald", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"previousOwners": []string{
					"alice",
					"donald",
				},
			},
		})

		if err != nil {
			t.Error(err)
		}

		if len(res) != 1 {
			t.Error("Query should have returned 1 result")
		}
	})

	t.Run("previousOwners should not equal boolean false", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"previousOwners": false,
			},
		})

		if err != nil {
			t.Error(err)
		}

		if len(res) != 0 {
			t.Error("Query should have returned 0 result")
		}
	})
}

func TestEQSubdocument(t *testing.T) {
	t.Run("Family name should equal colored implicit", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"family.name": "colored",
			},
		})

		if err != nil {
			t.Error(err)
		}

		if len(res) != 1 {
			t.Error("Query should have returned 1 result")
		}
	})
	t.Run("Family name should be colored", func(t *testing.T) {

		res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"family.name": map[string]interface{}{
					"$eq": "colored",
				},
			},
		})

		if err != nil {
			t.Error(err)
		}

		if len(res) != 1 {
			t.Error("Query should have returned 1 result")
		}
	})

}
