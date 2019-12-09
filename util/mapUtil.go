package util

import "fmt"

func ConvertMap(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = ConvertMap(v)
		}
		return m2
	case map[interface{}]string:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = ConvertMap(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = ConvertMap(v)
		}
	}
	return i
}

func FlatMap(any map[interface{}]interface{}) map[interface{}]string {

	m := map[interface{}]string{}
	for k, v := range any {
		flatten(k.(string), v, m)
	}

	return m
}

func flatten(prefix string, value interface{}, flatMap map[interface{}]string) {
	subMap, ok := value.(map[interface{}]interface{})
	if ok {
		for k, v := range subMap {
			flatten(prefix+"."+k.(string), v, flatMap)
		}
		return
	}
	stringList, ok := value.([]interface{})
	if ok {
		flatten(fmt.Sprintf("%s.size", prefix), len(stringList), flatMap)
		for i, v := range stringList {
			flatten(fmt.Sprintf("%s.%d", prefix, i), v, flatMap)
		}
		return
	}
	flatMap[prefix] = fmt.Sprintf("%v", value)
}


