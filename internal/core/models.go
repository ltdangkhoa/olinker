package core

type VendorConfig struct {
	Vendor  string `json:"vendor"`
	DLLPath string `json:"dll_path"`
	Port    int    `json:"port"`
}

type EncodeRequest struct {
	HotelPwd string `json:"HotelPwd"`
	CardID   string `json:"card_id"`
	RoomID   string `json:"room_id"`
	RoomName string `json:"room_name"`
	BeginTime string `json:"BeginTime"`
	EndTime   string `json:"EndTime"`
}

type EncodeResult struct {
	Status   int    `json:"status"`
	CardNo   string `json:"card_no"`
	RoomNo   string `json:"room_no,omitempty"`
	RoomName string `json:"room_name"`
	Stime    string `json:"stime,omitempty"`
}

type CancelRequest struct {
	CardID string `json:"card_id"`
}

type ExtendRequest struct {
	CardID  string `json:"card_id"`
	EndTime string `json:"EndTime"`
}

type CardInfo struct {
	Status   int    `json:"status"`
	CardNo   string `json:"card_no"`
	RoomName string `json:"room_name"`
}
