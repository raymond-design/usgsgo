package shared

import "fmt"

func GetIvsUrl(c string) string {
	return fmt.Sprintf(IvsUrl, c)
}

func GetIvsUrlFromAuthFunc(f AuthFunc) string {
	params := f("")
	return GetIvsUrl(params.Get("clientString"))
}
