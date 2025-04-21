package routes

import (
	// "fmt"
	"time"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func SystemResources(e *echo.Group) error {
	e.GET("/resources", func(c echo.Context) error {
		// Get CPU usage over a 1-second interval for all CPUs
		cpuPercentages, err := cpu.Percent(time.Second, true)
		if err != nil {
			log.Fatalf("Error getting CPU info: %v", err)
		}

		// Get memory usage
		vm, err := mem.VirtualMemory()
		if err != nil {
			log.Fatalf("Error getting memory info: %v", err)
		}

		// Return detailed CPU info and memory in JSON
		return c.JSON(200, map[string]interface{}{
			"cpu": cpuPercentages,
			"mem": map[string]interface{}{
				"usedPercent": vm.UsedPercent,
				"total":       vm.Total,
			},
		})
	})
	return nil
}
