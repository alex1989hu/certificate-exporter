package version

//nolint:gochecknoglobals
var (
	// GitBranch overridden during build time.
	GitBranch = "dev"

	// GitCommit overridden during build time.
	GitCommit = "dirty"

	// Time overridden during build time.
	Time = ""
)
