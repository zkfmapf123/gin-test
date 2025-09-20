package secrets

import (
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func Test_env(t *testing.T) {
	err := SetValue(filepath.Join("..", "..", ".env"))

	assert.Equal(t, err, nil)

	assert.Equal(t, GetStringOrDefault("TEST_STR", "default"), "aa")
	assert.Equal(t, GetIntOrDefault("TEST_INT", 0), 10)
	assert.Equal(t, GetBoolOrDefault("TEST_BOOL", false), true)

	assert.Equal(t, viper.GetString("TEST_STR"), "aa")
	assert.Equal(t, viper.GetInt("TEST_INT"), 10)
	assert.Equal(t, viper.GetBool("TEST_BOOL"), true)
}
