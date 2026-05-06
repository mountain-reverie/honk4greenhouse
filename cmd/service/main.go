package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/urfave/cli/v2"
	"github.com/mountain-reverie/honk4greenhouse/internal/server"
)

func main() {
	app := &cli.App{
		Name:        "honk4greenhouse",
		Usage:       "Your greenhouse design companion",
		Description: "A website for designing efficient and potentially passive greenhouses optimized for Canadian climate conditions",
		Version:     getVersion(),
		
		// Global flags
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Value:   "8080",
				Usage:   "HTTP server port",
				EnvVars: []string{"PORT"},
			},
		},
		
		// Default action (serve)
		Action: serveAction,
		
		Commands: []*cli.Command{
			{
				Name:   "serve",
				Usage:  "Start the HTTP server",
				Action: serveAction,
			},
			{
				Name:   "version",
				Usage:  "Show detailed version information",
				Action: versionAction,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func serveAction(c *cli.Context) error {
	port := c.String("port")
	fmt.Printf("🏠 Starting honk4greenhouse server on port %s\n", port)
	
	r := server.New()
	
	if err := server.Start(r, port); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	
	return nil
}

func versionAction(c *cli.Context) error {
	printDetailedBuildInfo()
	return nil
}

func getVersion() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		if info.Main.Version != "" && info.Main.Version != "(devel)" {
			return info.Main.Version
		}
		
		// Try to get VCS revision as fallback
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				if len(setting.Value) >= 7 {
					return setting.Value[:7] // Short commit hash
				}
				return setting.Value
			}
		}
	}
	return "dev"
}

func printDetailedBuildInfo() {
	fmt.Println("🏠 honk4greenhouse - Your greenhouse design companion")
	fmt.Println()
	
	if info, ok := debug.ReadBuildInfo(); ok {
		fmt.Printf("Module: %s\n", info.Main.Path)
		if info.Main.Version != "(devel)" && info.Main.Version != "" {
			fmt.Printf("Version: %s\n", info.Main.Version)
		}
		
		// Print VCS information if available
		for _, setting := range info.Settings {
			switch setting.Key {
			case "vcs.revision":
				fmt.Printf("Commit: %s\n", setting.Value)
			case "vcs.time":
				fmt.Printf("Build Time: %s\n", setting.Value)
			case "vcs.modified":
				if setting.Value == "true" {
					fmt.Println("Status: modified")
				} else {
					fmt.Println("Status: clean")
				}
			}
		}
		
		fmt.Printf("Go Version: %s\n", info.GoVersion)
	} else {
		fmt.Println("Build information not available")
	}
}
