package main

import (
	"os"
	"path/filepath"
	"time"

	"github.com/joyme123/cola-device-plugin/pkg/server"
	log "github.com/sirupsen/logrus"
	"gopkg.in/fsnotify.v1"
)

func main() {
	log.Info("cola device plugin starting")
	colaSrv := server.NewColaServer()
	go colaSrv.Run()

	// 向 kubelet 注册
	if err := colaSrv.RegisterToKubelet(); err != nil {
		log.Fatalf("register to kubelet error: %v", err)
	} else {
		log.Infoln("register to kubelet successfully")
	}

	// 监听 kubelet.sock，一旦创建则重启
	devicePluginSocket := filepath.Join(server.DevicePluginPath, server.KubeletSocket)
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println("Failed to created FS watcher.")
		os.Exit(1)
	}
	defer watcher.Close()

	for {
		select {
		case event := <-watcher.Events:
			if event.Name == devicePluginSocket && event.Op&fsnotify.Create == fsnotify.Create {
				time.Sleep(time.Second)
				log.Fatalf("inotify: %s created, restarting.", devicePluginSocket)
			}
		case err := <-watcher.Errors:
			log.Fatalf("inotify: %s", err)
		}
	}
}
