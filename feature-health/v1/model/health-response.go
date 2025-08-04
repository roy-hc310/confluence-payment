package model

type HealthResponse struct {
	Name        string   `json:"name"`
	Uptime      string   `json:"up_time"`
	TotalMemory string   `json:"total_memory"`
	FreeMemory  string   `json:"free_memory"`
	UsedPercent string   `json:"used_memory"`
	Cpus        []string `json:"cpus"`
	HostOS      string   `json:"host_os"`
	HostId      string   `json:"host_id"`
}
