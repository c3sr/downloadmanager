package downloadmanager

import (
	"os"
	"testing"

	"github.com/c3sr/config"
	"github.com/stretchr/testify/assert"
)

func TestDownloadJSON(t *testing.T) {
	const url = "http://data.dmlc.ml/models/imagenet/inception-bn/Inception-BN-symbol.json"
	target, _, err := DownloadFile(url, "/tmp/inception.json", Cache(false), MD5Sum("9f0f51b5e686c000b68629e4300decdd"))

	assert.NoError(t, err)
	assert.NotEmpty(t, target)
}

func TestMain(m *testing.M) {
	config.Init(
		config.AppName("carml"),
		config.VerboseMode(true),
		config.DebugMode(true),
	)
	os.Exit(m.Run())
}
