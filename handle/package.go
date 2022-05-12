package handle


func Run(get map[string]interface{}, callback func(string, int)) {
	code := get["code"].(int)
	src  := get["src"].(string)

	callback(src, code)
}