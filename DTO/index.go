package meta

import "example.com/t/enum"

type WS_META_File struct {
	EVENT enum.WS_EVENT `json:"event"`
	TYPE  enum.WS_TYPE  `json:"type"`
	DATA  any           `json:"data"`
}
