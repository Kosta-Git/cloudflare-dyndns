package pkg

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func CreateServiceFile() {
	executableLocation, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	realPath, err := filepath.EvalSymlinks(executableLocation)
	if err != nil {
		panic(err)
	}
	serviceContent := fmt.Sprintf(`[Unit]
Description=Cloudflare Dynamic DNS Updater
After=network.target

[Service]
ExecStart=%s start
Restart=on-failure
User=root
Group=root

[Install]
WantedBy=default.target`, realPath)
	err = os.WriteFile("/etc/systemd/system/cf-dyndns.service", []byte(serviceContent), 0644)
	if err != nil {
		log.Fatalf("Error writing service file: %v\n", err)
		return
	}
}
