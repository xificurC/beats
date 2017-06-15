package memory

import (
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/metricbeat/module/docker"
)

type MemoryData struct {
	Time      common.Time
	Container *docker.Container
	Failcnt   uint64
	Limit     uint64
	MaxUsage  uint64
	TotalRss  uint64
	TotalRssP float64
	Usage     uint64
	UsageP    float64
}

type MemoryService struct{}

func (s *MemoryService) getMemoryStatsList(rawStats []docker.Stat) []MemoryData {
	formattedStats := []MemoryData{}
	for _, myRawStats := range rawStats {
		formattedStats = append(formattedStats, s.GetMemoryStats(myRawStats))
	}

	return formattedStats
}

func (s *MemoryService) GetMemoryStats(myRawStat docker.Stat) MemoryData {
	memLimit := float64(myRawStat.Stats.MemoryStats.Limit)
	totalRss := float64(myRawStat.Stats.MemoryStats.Stats.TotalRss)
	var totalRssP float64 = 0
	if memLimit != 0 {
		totalRssP = totalRss / memLimit
	}
	return MemoryData{
		Time:      common.Time(myRawStat.Stats.Read),
		Container: docker.NewContainer(&myRawStat.Container),
		Failcnt:   myRawStat.Stats.MemoryStats.Failcnt,
		Limit:     myRawStat.Stats.MemoryStats.Limit,
		MaxUsage:  myRawStat.Stats.MemoryStats.MaxUsage,
		TotalRss:  myRawStat.Stats.MemoryStats.Stats.TotalRss,
		TotalRssP: totalRssP,
		Usage:     myRawStat.Stats.MemoryStats.Usage,
		UsageP:    float64(myRawStat.Stats.MemoryStats.Usage) / float64(myRawStat.Stats.MemoryStats.Limit),
	}
}
