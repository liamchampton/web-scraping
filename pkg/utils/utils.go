package utils

import (
	"strings"
)

func FormatPrice(price *string) {
	//r := regexp.MustCompile(`\$(\d+(\.\d+)?).*$`)
	////r := regexp.MustCompile(`(\d+\.\d{1,2})`) - not used
	//newPrices := r.FindStringSubmatch(*price)
	//fmt.Println("newPrices:", newPrices)
	//
	//if len(newPrices) > 1 {
	//	*price = newPrices[1]
	//} else {
	//	*price = "Unknown"
	//}

	r := strings.Count(*price, "£")

	//fmt.Println(r) // 1 or 2 £'s in a string

	if r > 1 {
		splitStr := strings.Split(*price, "£")
		*price = splitStr[1]
	}

}

func FormatStars(stars *string) {
	if len(*stars) >= 3 {
		*stars = (*stars)[0:3]
	} else {
		*stars = "Unknown"
	}
}