package blocks

import (
	"log"
	"strconv"
	"strings"
)

// getKeyValues returns values for all paths, including arrays
// {"foo":"bar"} returns [bar] for string "foo"
// {"foo":["bar","bar","bar"]} returns [bar, bar, bar] for string "foo.[]"
// {"foo":[{"type":"bar"},{"type":"baz"}]} returns [bar, baz] for string "foo.[].type"
func getKeyValues(d interface{}, p string) []interface{} {
	var values []interface{}
	var key string
	var rest string

	keyIdx := strings.Index(p, ".")

	if keyIdx != -1 {
		key = p[:keyIdx]
		rest = p[keyIdx+1:]
	} else {
		key = p
	}

	bStart := strings.Index(key, "[")
	bEnd := strings.Index(key, "]")
	var id int64
	id = -1
	if bStart == 0 && bEnd != 1 {
		id, _ = strconv.ParseInt(key[bStart+1:bEnd], 10, 64)
	}

	switch d := d.(type) {
	case map[string]interface{}:
		if len(rest) > 0 {
			x := getKeyValues(d[key], rest)
			for _, z := range x {
				values = append(values, z)
			}
		} else {
			values = append(values, d[p])
		}

	case []interface{}:
		var ids []int64
		if id != -1 {
			ids = append(ids, int64(id))
		} else {
			for i := range d {
				ids = append(ids, int64(i))
			}
		}

		for _, id := range ids {
			if len(rest) > 0 {
				x := getKeyValues(d[id], rest)
				for _, z := range x {
					values = append(values, z)
				}
			} else {
				values = append(values, d[id])
			}
		}
	default:
	}

	return values
}

func equals(value interface{}, comparator interface{}) bool {
	switch value := value.(type) {
	case int:
		c := comparator.(float64)
		return value == int(c)
	case string:
		return value == comparator
	case bool:
		return value == comparator
	default:
		log.Println("cannot perform an equals operation on this type")
		return false
	}
}

func greaterthan(value interface{}, comparator interface{}) bool {
	switch value := value.(type) {
	case int:
		return value > int(comparator.(float64))
	case float64:
		return value > comparator.(float64)
	default:
		log.Println("cannot perform a greaterthan operation on this type")
		return false
	}
}

func lessthan(value interface{}, comparator interface{}) bool {
	switch value := value.(type) {
	case int:
		return value < int(comparator.(float64))
	case float64:
		return value < comparator.(float64)
	default:
		log.Println("cannot perform a lessthan operation on this type")
		return false
	}
}

func subsetof(value interface{}, comparator interface{}) bool {
	log.Println("HELLO")
	switch value := value.(type) {
	case string:
		log.Println("VALUE", value)
		log.Println("COMPARATOR", comparator.(string))
		return strings.Contains(value, comparator.(string))
	default:
		log.Println(value)
		log.Println("cannot perform a subsetof operation on this type")
		return false
	}
}
