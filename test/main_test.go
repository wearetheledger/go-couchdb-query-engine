package test

var TestData = map[string]interface{}{
	"MARBLE1": map[string]interface{}{
		"objectType": "MARBLE",
		"owner":      "bob",
		"size":       3,
		"previousOwners": []string{
			"alice",
			"donald",
		},
		"family": map[string]interface{}{
			"name":   "colored",
			"origin": "spain",
		},
	},
	"MARBLE2": map[string]interface{}{
		"objectType": "MARBLE",
		"owner":      "alice",
		"size":       1,
		"previousOwners": []string{
			"donald",
		},
		"family": map[string]interface{}{
			"name":   "white",
			"origin": "france",
		},
		"verification": []interface{}{
			map[string]interface{}{
				"organizationId": "marbles inspectors inc.",
				"checkedAt":      "2017-12-18T12:11:34.171Z",
				"score":          5,
			},
		},
	},
	"MARBLE3": map[string]interface{}{
		"objectType": "MARBLE",
		"owner":      "arnold",
		"image":      "https://example.com/image.png",
		"size":       5,
		"previousOwners": []string{
			"alice",
		},
		"family": map[string]interface{}{
			"name":   "striped",
			"origin": "america",
		},
		"verification": []interface{}{
			map[string]interface{}{
				"organizationId": "marbles ID inc.",
				"checkedAt":      "2017-12-18T12:11:34.171Z",
				"score":          3,
			},
			map[string]interface{}{
				"organizationId": "marbles ID inc.",
				"checkedAt":      "2016-12-18T12:11:34.171Z",
				"score":          7,
			},
		},
	},
}

// TODO test
/*
res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
		"selector": map[string]interface{}{
			"objectType": map[string]interface{}{
				"ne": "MARBLE",
			},
		},
	})
*/
