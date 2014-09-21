package utils

func String(str string) *string {
	p := new(string)
	*p = str
	return p
}
