package types

import "sync"

/*
Implement a ConfigManager Singleton that:

Has fields appName string and version string
Is created once via GetInstance() using sync.Once
Has a GetConfig() string method that returns something like "AppName: MyApp, Version: 1.0"
In main, call GetInstance() three times and show all three point to the same object
*/

type ConfigManager struct {
	appName string
	version string
}

var instance *ConfigManager
var once sync.Once

func GetInstance() *ConfigManager {
	once.Do(func() {
		instance = &ConfigManager{appName: "Singleton", version: "1234567"}
	})
	return instance
}
