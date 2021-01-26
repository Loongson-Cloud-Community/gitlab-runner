package helpers

import (
	"os"

	"gitlab.com/gitlab-org/gitlab-runner/commands/helpers/archive"
	"gitlab.com/gitlab-org/gitlab-runner/commands/helpers/archive/fastzip"
	"gitlab.com/gitlab-org/gitlab-runner/helpers/featureflags"

	// auto-register default archivers/extractors
	_ "gitlab.com/gitlab-org/gitlab-runner/commands/helpers/archive/gziplegacy"
	_ "gitlab.com/gitlab-org/gitlab-runner/commands/helpers/archive/raw"
	_ "gitlab.com/gitlab-org/gitlab-runner/commands/helpers/archive/ziplegacy"

	"github.com/sirupsen/logrus"
)

func init() {
	// enable fastzip archiver/extractor
	if on, _ := featureflags.IsOn(os.Getenv(featureflags.UseFastzip)); on {
		archive.Register(archive.Zip, fastzip.NewArchiver, fastzip.NewExtractor)
	}
}

func getCompressionLevel(name string) archive.CompressionLevel {
	switch name {
	case "fastest":
		return archive.FastestCompression
	case "fast":
		return archive.FastCompression
	case "slow":
		return archive.SlowCompression
	case "slowest":
		return archive.SlowestCompression
	case "default", "":
		return archive.DefaultCompression
	}

	logrus.Warningf("compression level %q is invalid, falling back to default", name)

	return archive.DefaultCompression
}
