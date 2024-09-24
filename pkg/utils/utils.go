package utils

import (
	"strconv"
	"strings"
	"time"
)

// StringToTimeDuration - convert string to time.Duration
// example input: "1s", "1m", "1h", "1d"
func StringToTimeDuration(s string) time.Duration {

	defaultDuration := 30 * time.Second

	// Cek apakah input memiliki suffix "d" (hari)
	if strings.HasSuffix(s, "d") {
		// Ambil angka di depannya (tanpa "d")
		numberStr := strings.TrimSuffix(s, "d")
		days, err := strconv.Atoi(numberStr)
		if err != nil {
			return defaultDuration
		}

		duration := time.Duration(days) * 24 * time.Hour
		return duration
	}

	// Jika tidak ada "d", parse menggunakan time.ParseDuration
	duration, err := time.ParseDuration(s)
	if err != nil {
		return defaultDuration
	}

	// Tentukan unit berdasarkan akhiran
	var unit string
	if strings.HasSuffix(s, "s") {
		unit = "second"
	} else if strings.HasSuffix(s, "m") {
		unit = "minute"
	} else if strings.HasSuffix(s, "h") {
		unit = "hour"
	}

	// Ambil angka durasi dari hasil parsing
	value := int64(duration.Seconds())

	// Sesuaikan durasi berdasarkan unit
	switch unit {
	case "minute":
		value = int64(duration.Minutes())
	case "hour":
		value = int64(duration.Hours())
	}

	// Jika lebih dari 1, tambahkan plural "s"
	if value > 1 {
		unit += "s"
	}

	return duration

}
