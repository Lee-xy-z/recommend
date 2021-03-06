/**
* @Author: Zhangxinyu
* @Date: 2020/6/28 13:57
 */
package version

var (
	// commitFromGit is a constant representing the source version that
	// generated this build. It should be set during build via -ldflags.
	commitSHA string

	// versionFromGit is a constant representing the version tag that
	// generated this build.It should be set during build via -ldflags.
	latestVersion string

	//build date in ISO8601 format,output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
	date string
)

type Info struct {
	GitCommit  string `json:"GitCommit"`
	GitVersion string `json:"GitVersion"`
	BuildDate  string `json:"BuildDate"`
}

// Get creates and initialized Info object
func Get() Info {
	return Info{
		GitCommit:  commitSHA,
		GitVersion: latestVersion,
		BuildDate:  date,
	}
}
