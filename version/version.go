package version

import (
	"fmt"
)

const (
	// ServiceName 服务名称
	ServiceName = "keyauth"

	// Description 服务描述
	Description = "微服务权限中心"
)

var (
	GIT_COMMIT string
	GIT_BRANCH string
	BUILD_TIME string
	GO_VERSION string
)

// FullVersion show the version info
func FullVersion() string {
	version := fmt.Sprintf("Build Time: %s\nGit Branch: %s\nGit Commit: %s\nGo Version: %s\n", BUILD_TIME, GIT_BRANCH, GIT_COMMIT, GO_VERSION)
	return version
}

// Short 版本缩写
func Short() string {
	commit := ""
	if len(GIT_COMMIT) > 8 {
		commit = GIT_COMMIT[:8]
	}
	return fmt.Sprintf("%s[%s]", GIT_BRANCH, commit)
}
