package udp_utils_gen

import (
	"github.com/elliotchance/orderedmap/v3"

	udp_utils_parse "example.com/t/udp/utils/parse"
	udp_utils_struct "example.com/t/udp/utils/stuct"
)

func Login() (res []byte) {
	m := orderedmap.NewOrderedMap[string, udp_utils_struct.Base]()
	m.Set("buf", udp_utils_struct.Base{
		Data:  "test_login",
		Range: []int{},
		Size:  0,
	})
	res = udp_utils_parse.ParseSendBuf(m)
	return res
}
