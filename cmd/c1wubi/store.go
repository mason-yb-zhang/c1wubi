package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const defaultDataDir = "/storage/c1/local-apps/apps/wubi/data"

func dataDir() string {
	if d := os.Getenv("C1WUBI_DATA"); d != "" {
		return d
	}
	return defaultDataDir
}

type stats struct {
	BestStreak int `json:"bestStreak"`
	TotalOK    int `json:"totalOk"`
	TotalTry   int `json:"totalTry"`
}

func statsPath() string { return filepath.Join(dataDir(), "stats.json") }

func loadStats() stats {
	data, err := os.ReadFile(statsPath())
	if err != nil {
		return stats{}
	}
	var s stats
	_ = json.Unmarshal(data, &s)
	return s
}

func saveStats(s stats) error {
	if err := os.MkdirAll(dataDir(), 0o755); err != nil {
		return err
	}
	data, _ := json.MarshalIndent(s, "", "  ")
	tmp := statsPath() + ".tmp"
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, statsPath())
}
