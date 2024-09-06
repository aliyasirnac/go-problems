package defanging_an_ip_address

import "strings"

func defangIPaddr(address string) string {
	var sb strings.Builder
	for i := 0; i < len(address); i++ {
		ch := address[i]
		if ch == '.' {
			sb.WriteString("[.]")
			continue
		}
		sb.WriteByte(ch)
	}

	return sb.String()
}
