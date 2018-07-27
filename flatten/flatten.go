package flatten

func Flatten(data interface{}) []interface{} {

	var res []interface{}

	switch d := data.(type) {
	case []interface{}:
		for _, v := range d {
			res = append(res, Flatten(v)...)
		}
	case interface{}:
		res = append(res, d)

	}

	return res
}
