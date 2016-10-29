package kr

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"
)

var analytics_user_agent = fmt.Sprintf("Mozilla/5.0 (Macintosh; Intel Mac OS X) AppleWebKit/602.2.14 (KHTML, like Gecko) Version/%s kr/%s", CURRENT_VERSION, CURRENT_VERSION)

var analytics_os = "OS X"

var cachedAnalyticsOSVersion *string
var osVersionMutex sync.Mutex

func getAnalyticsOSVersion() *string {
	osVersionMutex.Lock()
	defer osVersionMutex.Unlock()
	if cachedAnalyticsOSVersion != nil {
		return cachedAnalyticsOSVersion
	}

	analytics_os_version_bytes, err := exec.Command("sw_vers", "-productVersion").Output()
	if err != nil {
		return nil
	}
	stripped := strings.TrimSpace(string(analytics_os_version_bytes))
	cachedAnalyticsOSVersion = &stripped
	return cachedAnalyticsOSVersion
}