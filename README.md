# go-couchdb-query-engine
CouchDB query implementation in Golang for in-memory object querying

**Build status**: [![Build Status](https://travis-ci.org/wearetheledger/go-couchdb-query-engine.svg?branch=master)](https://travis-ci.org/wearetheledger/go-couchdb-query-engine)

## Features
- skip
- limit
- selector

## Selector
The selector query accepts following couchdb operators:
- all
- and
- elematch
- eq
- exists
- gt
- gte
- in
- lt
- lte
- mod
- ne
- nin
- nor
- not
- or
- regex
- size
- type

## Example code
### Data

```golang
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
```

### Method
```golang
res, err := query.ParseCouchDBQuery(TestData, map[string]interface{}{
			"selector": map[string]interface{}{
				"size": map[string]interface{}{
					"$eq": 3,
				},
			},
		})
```
### Response
```golang
StateCouchDBQueryResult{
	StateCouchDBQueryObject{
		Key: "MARBLE0",
		Value: map[string]interface{}{
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
	},
}
```
