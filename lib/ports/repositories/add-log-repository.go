package repositories

type AddLogRepository interface {
	Add(data AddLogRequest) AddLogResponse
}

type AddLogRequest struct {
	Type        string      `json:"type"`
	Description interface{} `json:"description"`
	Service     string      `json:"service"`
	Version     float64     `json:"version"`
}

type AddLogResponse struct {
	string
}
