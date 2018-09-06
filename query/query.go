package query

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/wearetheledger/go-couchdb-query-engine/query/rules"
)

var Registry []rules.IGenericRule

func init() {

	Registry = []rules.IGenericRule{
		// ===
		// Condition Operators == http://docs.couchdb.org/en/2.1.0/api/database/find.html?highlight=find#condition-operators
		// ===

		/// Operator type - (In)equality
		rules.IGenericRule(rules.Lt),
		rules.IGenericRule(rules.Lte),

		rules.IGenericRule(rules.Eq),
		rules.IGenericRule(rules.Ne),

		rules.IGenericRule(rules.Gte),
		rules.IGenericRule(rules.Gt),

		/// Operator type - Object
		rules.IGenericRule(rules.Exists),
		rules.IGenericRule(rules.Type),

		/// Operator type - Array
		rules.IGenericRule(rules.In),
		rules.IGenericRule(rules.Nin),
		rules.IGenericRule(rules.Size),

		/// Operator type - Miscellaneous
		rules.IGenericRule(rules.Mod),
		rules.IGenericRule(rules.Regex),

		// ===
		// Combination Operators == http://docs.couchdb.org/en/2.1.0/api/database/find.html?highlight=find#combination-operators
		// ===

		rules.IGenericRule(rules.And(test)),
		rules.IGenericRule(rules.Or(test)),
		rules.IGenericRule(rules.Not(test)),
		rules.IGenericRule(rules.Nor(test)),
		rules.IGenericRule(rules.All),
		rules.IGenericRule(rules.ElemMatch(test)),
	}
}

type StateCouchDBQueryObject struct {
	Key   string
	Value interface{}
}

type StateCouchDBQueryResult []StateCouchDBQueryObject

// Query engine which controls and accept the whole query. It makes sure every field is filtered.
func ParseCouchDBQuery(data map[string]interface{}, userQueryMap map[string]interface{}) (StateCouchDBQueryResult, error) {

	userQueryAsBytes, err := json.Marshal(userQueryMap)

	if err != nil {
		return nil, err
	}

	return ParseCouchDBQueryString(data, string(userQueryAsBytes))

}

func ParseCouchDBQueryString(data map[string]interface{}, userQueryString string) (StateCouchDBQueryResult, error) {

	var userQuery map[string]interface{}
	parseErr := json.Unmarshal([]byte(userQueryString), &userQuery)

	if parseErr != nil {
		return StateCouchDBQueryResult{}, errors.New("Error parsing query string")
	}

	if userQuery["selector"] == nil {
		return StateCouchDBQueryResult{}, errors.New("Invalid query, selector required")
	}

	var response StateCouchDBQueryResult

	for k, v := range data {
		result, err := test(v, userQuery["selector"])

		if err != nil {
			return []StateCouchDBQueryObject{}, err

		} else if result {
			response = append(response, StateCouchDBQueryObject{
				k,
				v,
			})

		}
	}

	// Take slice between skip and limit

	var skip = 0
	var limit = len(response)

	if userQuery["skip"] != nil {
		skipF64, ok := userQuery["skip"].(float64)

		if ok {
			skip = int(skipF64)
		} else {
			return StateCouchDBQueryResult{}, errors.New("Skip must be an integer")
		}
	}

	if userQuery["limit"] != nil {
		limitF64, ok := userQuery["limit"].(float64)

		if ok {
			limit = int(limitF64)
		} else {
			return StateCouchDBQueryResult{}, errors.New("Limit must be an integer")
		}

	}

	if skip < len(response) {
		response = response[skip:]
	} else {
		return StateCouchDBQueryResult{}, nil
	}

	if limit > len(response) {
		limit = len(response)
	}

	return response[:limit], nil

}

// Matcher function which delegates the work to the different rules based on the key and compares values to
// implicit match values.

func test(data interface{}, userQuery interface{}) (bool, error) {

	last := true

	if canDecend(userQuery) && reflect.TypeOf(userQuery).Kind() == reflect.Map {

		userQuery, ok := userQuery.(map[string]interface{})

		if !ok {
			return false, errors.New("Data is not a map")
		}

		for k, v := range userQuery {

			if string(k[0]) == "$" {
				rule, err := getRule(k)

				if err != nil {
					return false, err
				}

				matched, err1 := rule.Match(data, v)

				if err1 != nil {
					return false, err1
				}

				last = last && matched

			} else {

				dvp, dk := resolveSubdocumentQuery(data, k)

				var result bool
				var err1 error

				if dvp != nil && len(dk) == 1 {
					result, err1 = test(dvp[dk[0]], v)

				} else {
					result, err1 = test(nil, v)
				}

				if err1 != nil {
					return false, err1
				}

				last = last && result

			}
		}

		return last, nil

	} else {

		if userQuery == nil || data == nil {
			return false, nil
		}

		bool, err1 := rules.Eq.Match(data, userQuery)

		if err1 != nil {
			return false, err1
		}

		last = last && bool

		return last, nil
	}

	return false, nil
}

// getRule function is used to find rule based on it's name
func getRule(name string) (rules.IGenericRule, error) {

	for _, rule := range Registry {
		if rule.GetName() == name {
			return rule, nil
		}
	}

	return nil, fmt.Errorf("Rule %s not found!", name)

}

// resolveSubdocumentQuery function is used to match names like "foo.bar" to get the "bar" value in a subdocument
func resolveSubdocumentQuery(data interface{}, key string) (map[string]interface{}, []string) {

	var last []string
	var deepObject map[string]interface{}

	stack := strings.Split(key, ".")

	if len(stack) > 0 {
		item := stack[len(stack)-1]

		stack = stack[:len(stack)-1]

		last = append([]string{item}, last...)
	}

	_, ok := data.(map[string]interface{})

	if ok {
		deepObject, _ = data.(map[string]interface{})

		for len(stack) > 0 {
			x := stack[0]
			stack = stack[1:]

			t := interface{}(deepObject[x])

			newObj, ok := t.(map[string]interface{})

			if ok {
				deepObject = newObj
			} else {
				stack = stack[:len(stack)-1]
			}

		}
	}

	last = append(stack, last...)

	if deepObject == nil {
		deepObject = data.(map[string]interface{})
	}

	return deepObject, last

}

func canDecend(a interface{}) bool {

	rt := reflect.TypeOf(a)

	if (rt.Kind() == reflect.Slice || rt.Kind() == reflect.Slice) && reflect.ValueOf(a).Len() > 0 {
		return true
	}

	if rt.Kind() == reflect.Map && len(reflect.ValueOf(a).MapKeys()) > 0 {
		return true
	}

	return false
}
