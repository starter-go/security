package rbac

import "testing"

func TestParseFullPhoneNumber(t *testing.T) {

	list := make([]string, 0)
	list = append(list, "+86-123456789")
	list = append(list, " + 86 - 1234 56789 ")
	list = append(list, " + 86 - 1234 - 56789 ")
	list = append(list, " + 86 - 12 + 34 - 56789 ")
	list = append(list, " 66 + 86 - 12 - 34 - 56789  ")
	list = append(list, " 66 + 86  12  34  56789  ")

	for _, str := range list {
		full, err := ParseFullPhoneNumber(str)
		if err == nil {
			t.Log("", full.String(), " <<< [", str, "]")
		} else {
			t.Log("parse:[", str, "] error:", err.Error())
		}
	}
}

func TestFullPhoneNumberParts(t *testing.T) {

	list := make([]string, 0)
	list = append(list, "+86-123456789")
	list = append(list, " + 86 - 1234 56789 ")
	list = append(list, " + 86 - 1234 - 56789 ")
	list = append(list, " + 86 - 12 + 34 - 56789 ")
	list = append(list, " 66 + 86 - 12 - 34 - 56789  ")
	list = append(list, " 66 + 86  12  34  56789  ")

	for _, str := range list {

		// full, err := ParseFullPhoneNumber(str)
		// if err != nil {
		// 	continue
		// }

		full := FullPhoneNumber(str)
		region, phone, err := full.Parts()
		if err != nil {
			t.Log("Error: ", err)
		} else {
			t.Log("region:", region, " phone:", phone, " full:", full)
		}
	}
}
