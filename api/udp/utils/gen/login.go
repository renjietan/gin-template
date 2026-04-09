package udp_utils_gen

import (
	"example.com/t/api/udp/utils/parse"
	"example.com/t/api/udp/utils/stuct"
	"github.com/elliotchance/orderedmap/v3"
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
