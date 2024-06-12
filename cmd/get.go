/*
Copyright © 2024 Guilherme Lira
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("---")
		fmt.Println("Memory Usage:")
		PrintMemUsage()

		fmt.Println("")
		fmt.Println("---")
		fmt.Println("CPU Usage:")
		PrintCpuUsage()

		fmt.Println("")
		fmt.Println("---")
		fmt.Println("Disk Usage:")
		PrintDisk()
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func PrintMemUsage() {
	v, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Print memory information
	fmt.Printf("Total: %v MB\n", v.Total/1024/1024)
	fmt.Printf("Free: %v MB\n", v.Free/1024/1024)
	fmt.Printf("Used: %v MB\n", v.Used/1024/1024)
	fmt.Printf("Used Percent: %.2f%%\n", v.UsedPercent)
}

func PrintCpuUsage() {
	// Get info about single CPU core
	percentages, err := cpu.Percent(0, true)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Print CPU utilizations of all cores
	for i, percent := range percentages {
		fmt.Printf("CPU %d: %.2f%%\n", i, percent)
	}

	// Get total CPU utilization
	totalPercent, err := cpu.Percent(0, false)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Print total CPU utulization
	fmt.Printf("Total CPU Usage: %.2f%%\n", totalPercent[0])
}

func PrintDisk() {
	usage, _ := disk.Usage("/")
	fmt.Println("Mount Point: /")
	fmt.Printf("Used disk Percent: %.2f%%\n", usage.UsedPercent)
}
