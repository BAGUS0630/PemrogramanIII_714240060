package model

type Response struct {
	Message string      `json:"message" example:"detail pesan"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty" example:"detail error"`
}

type Response401 struct {
	Message string `json:"message" example:"token tidak valid atau sudah expired"`
}

type Response403 struct {
	Message string `json:"message" example:"user tidak memiliki akses untuk fitur ini"`
}

type Response200 struct {
	Message string      `json:"message" example:"Berhasil"`
	Data    interface{} `json:"data,omitempty"`
}

type Response201 struct {
	Message string      `json:"message" example:"Berhasil dibuat"`
	Data    interface{} `json:"data,omitempty"`
}

