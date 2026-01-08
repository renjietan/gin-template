package main

import (
	"encoding/binary"
	"fmt"

	"example.com/t/utility"
)

type RadioDataJson struct {
	Data  any   `json:"data"`
	Range []int `json:"rg"`
	Size  int   `json:"length"`
}

func main() {
	params := map[string]RadioDataJson{
		"ee": {
			Data:  0xee,
			Range: []int{},
			Size:  1,
		},
		"ee2": {
			Data:  0xee,
			Range: []int{},
			Size:  1,
		},
		"length": {
			Data:  0x00,
			Range: []int{3, 6},
			Size:  2,
		},
		"data1": {
			Data:  "hello world1",
			Range: []int{},
			Size:  0,
		},
		"data2": {
			Data:  "hello world2",
			Range: []int{},
			Size:  0,
		},
		"data3": {
			Data:  "hello",
			Range: []int{},
			Size:  0,
		},
		"data33": {
			Data:  "hello-1",
			Range: []int{},
			Size:  0,
		},
		"length2": {
			Data:  0x00,
			Range: []int{7, 8},
			Size:  4,
		},
		"data4": {
			Data:  "hello world1",
			Range: []int{},
			Size:  0,
		},
		"data5": {
			Data:  "hello world2",
			Range: []int{},
			Size:  0,
		},
	}

	// 定义处理顺序
	keys := []string{"ee", "ee2", "length", "data1", "data2", "data3", "data33", "length2", "data4", "data5"}

	// 存储处理结果
	processedData := make(map[string][]byte)
	res := []byte{}
	sumKey := ""
	sumStart := false
	sumTotal := 0
	sumRange := []int{}
	// 第一步：处理所有数据并转换
	for index, key := range keys {
		item := params[key]

		if len(item.Range) > 1 {
			sumKey = key
			sumStart = true
			sumTotal = 0
			sumRange = []int{}
			sumRange = append(sumRange, item.Range...)
			processedData[key] = make([]byte, item.Size)
			// res = append(res, processedData[key]...)
			continue
		}

		var dataBytes []byte

		switch v := item.Data.(type) {
		case int:
			dataBytes = byteToMuiByte(v, item.Size, "big")
		case []byte:
			if item.Size > 0 {
				if len(v) >= item.Size {
					dataBytes = v[:item.Size]
				} else {
					dataBytes = make([]byte, item.Size)
					copy(dataBytes, v)
				}
			} else {
				dataBytes = v
			}
		case string:
			// 直接转换为 []byte
			dataBytes = []byte(v)
			item.Size = len(dataBytes)
		default:
			dataBytes = []byte{}
		}
		processedData[key] = dataBytes
		// res = append(res, processedData[key]...)
		if sumStart && sumKey != "" && index >= sumRange[0] && index <= sumRange[1] {
			sumTotal += len(dataBytes)
			if index == sumRange[1] {
				item := params[sumKey]
				processedData[sumKey] = byteToMuiByte(sumTotal, item.Size, "big")
			}
		}
	}
	for _, key := range keys {
		res = append(res, processedData[key]...)
		fmt.Printf("%s: %v %#v(length: %d)\n", key, processedData[key], processedData[key], len(processedData[key]))
	}
}

func byteToMuiByte(i int, length int, endian string) []byte {
	buf := make([]byte, length)
	switch length {
	case 1:
		buf[0] = byte(i)
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
