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
			Range: []int{3, 5},
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
		"length2": {
			Data:  0x00,
			Range: []int{},
			Size:  2,
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
	order := []string{"ee", "ee2", "length", "data1", "data2", "data3", "length2", "data4", "data5"}

	// 存储处理结果
	processedData := make(map[string][]byte)
	sumKey := ""
	sumStart := false
	sumTotal := 0

	// 第一步：处理所有数据并转换
	for _, key := range order {
		item := params[key]

		if len(item.Range) > 1 {
			sumKey = key
			sumStart = true
			sumTotal = 0
			processedData[key] = make([]byte, item.Size)
			continue
		}

		var dataBytes []byte

		switch v := item.Data.(type) {
		case int:
			dataBytes = byteToMuiByte(v, item.Size, "big")
		case []byte:
			// 根据 Size 调整字节切片长度
			if item.Size > 0 {
				if len(v) >= item.Size {
					dataBytes = v[:item.Size]
				} else {
					// 不足部分用0填充
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
			// 其他类型暂时不处理
			dataBytes = []byte{}
		}
		processedData[key] = dataBytes
		// 如果 sum 已开始，累加长度
		if sumStart && sumKey != "" {
			sumTotal += len(dataBytes)
		}
	}

	// 如果有 sum 字段，填充总和
	if sumKey != "" && sumTotal > 0 {
		item := params[sumKey]
		processedData[sumKey] = byteToMuiByte(sumTotal, item.Size, "big")
	}

	fmt.Println("Processed Data:")
	for _, key := range order {
		fmt.Printf("%s: %v (length: %d)\n", key, processedData[key], len(processedData[key]))
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
