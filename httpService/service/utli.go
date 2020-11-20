package service

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

//ReversString 卡片ID處理
func ReversString(input string) string {
	var arr []string
	s := strings.Split(input, "")
	ln := len(s)
	if ln%2 != 0 {
		s = append(s, "0")
	}
	for i := 0; i < len(s); i += 2 {
		tmp := s[i] + s[i+1]
		arr = append(arr, tmp)
	}
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return strings.Join(arr, "")
}

//HexToDecimal 16to10 補0
func HexToDecimal(src string) (string, error) {
	dec, err := strconv.ParseInt(src, 16, 64)
	if err != nil {
		log.Printf("Error in conversion: %s\n", err)
		return "", err
	}

	log.Printf("HexToDecimal Befor : %d", dec)
	ln := len(fmt.Sprintf("%d", dec))
	// format
	var result string
	if ln > 10 {
		result = fmt.Sprintf("%0*d", 17, dec)
	} else if ln < 10 {
		result = fmt.Sprintf("%0*d", 10, dec)
	} else {
		result = fmt.Sprintf("%d", dec)
	}
	log.Println("HexToDecimal after : " + result)
	return result, nil
}
