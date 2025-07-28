package parser

// GPSelement groups all GPS data in one record
type GPSelement struct {
	Longitude  float64 `json:"Longitude"`
	Latitude   float64 `json:"Latitude"`
	Altitude   int16   `json:"Altitude"`
	Angle      uint16  `json:"Angle"`
	Satellites uint8   `json:"Satellites"`
	Speed      uint16  `json:"Speed"`
}

// IOelement holds IO data for each record
type IOelement struct {
	EventID      int            `json:"EventID"`
	ElementCount int            `json:"ElementCount"`
	Elements     map[string]any `json:"Elements"` // use any for flexibility (could be int or string)
}

// AvlData represents a single AVL data record
type AvlData struct {
	Timestamp  string     `json:"Timestamp"` // ISO 8601 format (e.g. 2025-06-04T09:51:58.010Z)
	Priority   uint8      `json:"Priority"`
	GPSelement GPSelement `json:"GPSelement"`
	IOelement  IOelement  `json:"IOelement"`
}

// Content wraps the records (for output parity)
type Content struct {
	AVL_Datas []AvlData `json:"AVL_Datas"`
}

// TeltonikaAvlPacket is the root structure for a decoded Codec8 packet
type TeltonikaAvlPacket struct {
	Packet     string  `json:"Packet"`
	Preamble   uint32  `json:"Preamble"`
	DataLength uint32  `json:"Data_Length"`
	CodecID    uint8   `json:"CodecID"`
	CodecType  string  `json:"CodecType"`
	Quantity1  uint8   `json:"Quantity1"`
	Content    Content `json:"Content"`
	Quantity2  uint8   `json:"Quantity2"`
	CRC        uint32  `json:"CRC"`
	// meta: Skip for now (not needed for core parsing/output)
}
