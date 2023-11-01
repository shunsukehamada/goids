package cmd

import (
	"fmt"
	"goids/gui"
	"os"

	"github.com/go-vgo/robotgo"
	"github.com/spf13/cobra"
)

var width, height int
var n int
var speed float64
var force float64
var fullScreen bool

var rootCmd = &cobra.Command{
	Use:   "goids",
	Short: "gopher boids flocking algorithm animation",
	Long:  `gopher boids flocking algorithm animation.`,
	Run: func(cmd *cobra.Command, args []string) {
		if width <= 0 || height <= 0 {
			fmt.Println("width and height must be positive")
			os.Exit(1)
		}
		if n <= 0 {
			fmt.Println("number of gopher must be positive")
			os.Exit(1)
		}
		if speed < 0 {
			fmt.Println("max speed must be non negative")
			os.Exit(1)
		}
		if force < 0 {
			fmt.Println("max force must be non negative")
			os.Exit(1)
		}
		if fullScreen {
			width, height = robotgo.GetScreenSize()
		}
		gui.Run(width, height, n, speed, force)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().IntVarP(&width, "width", "w", 640, "width of the window")
	rootCmd.Flags().IntVar(&height, "height", 480, "height of the window")
	rootCmd.Flags().IntVarP(&n, "number", "n", 30, "number of gopher")
	rootCmd.Flags().Float64VarP(&speed, "speed", "s", 3, "max speed of the gopher")
	rootCmd.Flags().Float64VarP(&force, "force", "f", 2, "max force of the gopher")
	rootCmd.Flags().BoolVar(&fullScreen, "full", false, "full screen mode")
}
