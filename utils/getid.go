package utils

func Getid(s *string) string {
	if len(*s) != 22 {
		return ""
	}
	return (*s)[3 : len(*s)-1]
}
