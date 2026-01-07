package main

import (
	"encoding/binary"
	"fmt"
	"sync"

	"example.com/t/utility"
)

type RadioDataJson struct {
	Data any  `json:"data"`
	Sum  bool `json:"sum"`
	Size int  `json:"length"`
}

var wg sync.WaitGroup

func main() {
	var params = map[string]RadioDataJson{
		"ee": {
			Data: 0xee,
			Sum:  false,
			Size: 1,
		},
		"ee2": {
			Data: 0xee,
			Sum:  false,
			Size: 1,
		},
		"length": {
			Data: 0x00,
			Sum:  true,
			Size: 2,
		},
		"data1": {
			Data: "hello world1",
			Sum:  false,
			Size: 0,
		},
		"data2": {
			Data: "hello world2",
			Sum:  false,
			Size: 0,
		},
		"data3": {
			Data: "hello world3",
			Sum:  false,
			Size: 0,
		},
	}
	wg.Add(1)
	go RadioJsonToBuffer(params)
	wg.Wait()
	// fmt.Println("need_sum_field:", need_sum_field)
}

func RadioJsonToBuffer(params map[string]RadioDataJson) {
	var start_sum bool
	var need_sum_field string
	var total_size int
	for k, v := range params {
		if v.Sum {
			start_sum = true
			if need_sum_field != k {
				total_size = 0
			}
			need_sum_field = k
		} else {
			if start_sum && !v.Sum {
				total_size = total_size + v.Size
				sumVal := params[need_sum_field]
				sumVal.Data = byteToMuiByte(total_size, v.Size, "little")
				params[need_sum_field] = sumVal
			} else {
				var data []byte
				var size int
				fmt.Printf("int old --------%s=====================%v\n", k, v.Data)
				switch v.Data.(type) {
				case int:
					data = byteToMuiByte(v.Data.(int), v.Size, "little")
					size = len(v.Data.([]byte))
					v.Data = data
					v.Size = size
				case string:
					data = []byte(v.Data.(string))
					fmt.Println("========", k, v.Data)
					size = len(v.Data.([]byte))
					v.Data = data
					v.Size = size
				default:
					fmt.Printf("default-SIZE: %T:%v\n", v.Size, v.Size)
					fmt.Printf("default-SUM: %T:%v\n", v.Sum, v.Sum)
					fmt.Printf("default-DATA: %T:%v\n", v.Data, v.Data)
					fmt.Println("default-LENGTH:", params)
				}
			}
		}
	}
	fmt.Println("params:", params)

	wg.Done()
}

func byteToMuiByte(i int, length int, endian string) []byte {
	buf := make([]byte, length)
	switch length {
	case 2:
		utility.Tern(endian == "big",
			func() { binary.BigEndian.PutUint16(buf, uint16(i)) },
			func() { binary.LittleEndian.PutUint16(buf, uint16(i)) })
	case 4:
		utility.Tern(endian == "big",
			func() { binary.BigEndian.PutUint32(buf, uint32(i)) },
			func() { binary.LittleEndian.PutUint32(buf, uint32(i)) })
	case 8:
		utility.Tern(endian == "big",
			func() { binary.BigEndian.PutUint64(buf, uint64(i)) },
			func() { binary.LittleEndian.PutUint64(buf, uint64(i)) })
	}
	return buf
}
