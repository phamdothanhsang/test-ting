package Utils

import (
	"bytes"
	"regexp"
	"strings"
	"unicode/utf8"

	"golang.org/x/text/unicode/norm"
)

func GenerateSearchKeyword(s string) string {
	if s != "" {
		s = RemoveAccent(s)
		s = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(s, "")
		return strings.ToLower(strings.Replace(norm.NFD.String(s), " ", "", -1))
	}
	return ""
}

func UpperFirstChar(s string) string {
	//ex: channelConfigID -> ChannelConfigID
	return strings.ToUpper(s[0:1]) + s[1:]
}

func GetSortedFieldQuery(sortedField string, sortedDirection bool, table string) string {
	//ex: channelConfigID -> "Newss"."ChannelConfigID" | "Newss"."ChannelConfigID" desc
	var sortedFieldQuery string = table + ".\"" + UpperFirstChar(sortedField) + "\""
	if sortedDirection == false {
		sortedFieldQuery = sortedFieldQuery + " desc"
	}
	return sortedFieldQuery
}

// Mang cac ky tu goc co dau var
var SOURCE_CHARACTERS, LL_LENGTH = stringToRune("ÀÁÂÃÈÉÊÌÍÒÓÔÕÙÚÝàáâãèéêìíòóôõùúýĂăĐđĨĩŨũƠơƯưẠạẢảẤấẦầẨẩẪẫẬậẮắẰằẲẳẴẵẶặẸẹẺẻẼẽẾếỀềỂểỄễỆệỈỉỊịỌọỎỏỐốỒồỔổỖỗỘộỚớỜờỞởỠỡỢợỤụỦủỨứỪừỬửỮữỰự")

// Mang cac ky tu thay the khong dau var
var DESTINATION_CHARACTERS, _ = stringToRune(`AAAAEEEIIOOOOUUYaaaaeeeiioooouuyAaDdIiUuOoUuAaAaAaAaAaAaAaAaAaAaAaAaEeEeEeEeEeEeEeEeIiIiOoOoOoOoOoOoOoOoOoOoOoOoUuUuUuUuUuUuUu`)

func stringToRune(s string) ([]string, int) {
	ll := utf8.RuneCountInString(s)
	var texts = make([]string, ll+1)
	var index = 0
	for _, runeValue := range s {
		texts[index] = string(runeValue)
		index++
	}
	return texts, ll
}
func binarySearch(sortedArray []string, key string, low int, high int) int {
	var middle int = (low + high) / 2
	if high < low {
		return -1
	}
	if key == sortedArray[middle] {
		return middle
	} else if key < sortedArray[middle] {
		return binarySearch(sortedArray, key, low, middle-1)
	} else {
		return binarySearch(sortedArray, key, middle+1, high)
	}
}
func removeAccentChar(ch string) string {
	var index int = binarySearch(SOURCE_CHARACTERS, ch, 0, LL_LENGTH)
	if index >= 0 {

		ch = DESTINATION_CHARACTERS[index]
	}
	return ch
}
func RemoveAccent(s string) string {
	var buffer bytes.Buffer
	for _, runeValue := range s {
		buffer.WriteString(removeAccentChar(string(runeValue)))
	}
	return buffer.String()

}
