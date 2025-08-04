package service

import (
	"confluence-payment/feature-health/v1/model"
	"fmt"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type HealthService struct {
	StartAt time.Time
}

func NewHealthService() *HealthService {
	return &HealthService{
		StartAt: time.Now(),
	}
}

func (h *HealthService) Health() (res model.HealthResponse, err error, statusCode int) {
	stat, err := mem.VirtualMemory()
	if err != nil {
		return res, err, http.StatusInternalServerError
	}
	totalMemory := stat.Total / 1024 / 1024
	freeMemory := stat.Free / 1024 / 1024

	CpuPercentages, err := cpu.Percent(0, true)
	if err != nil {
		return res, err, http.StatusInternalServerError
	}
	CPUs := []string{}

	for i, percentage := range CpuPercentages {
		cpuInfo := fmt.Sprintf("CPU[%d]: %f%%", i, percentage)
		CPUs = append(CPUs, cpuInfo)
	}

	hostStat, err := host.Info()
	if err != nil {
		return res, err, http.StatusInternalServerError
	}

	res = model.HealthResponse{
		Name:        "promotion",
		Uptime:      time.Now().Sub(h.StartAt).String(),
		TotalMemory: fmt.Sprintf("%d MB", totalMemory),
		FreeMemory:  fmt.Sprintf("%d MB", freeMemory),
		UsedPercent: fmt.Sprintf("%f%%", stat.UsedPercent),
		Cpus:        CPUs,
		HostOS:      hostStat.OS,
		HostId:      hostStat.HostID,
	}

	return res, nil, http.StatusOK
}
