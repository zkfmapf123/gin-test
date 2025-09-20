package secrets

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func SetValue(fileIn string) error {
	return godotenv.Load(fileIn)
}

func GetStringOrDefault(key, defaultValue string) string {

	v := os.Getenv(key)
	viper.Set(key, v)

	return v
}

func GetIntOrDefault(key string, defaultValue int) int {
	v := os.Getenv(key)

	intV, _ := strconv.Atoi(v)
	viper.Set(key, intV)

	return intV
}

func GetBoolOrDefault(key string, defaultValue bool) bool {
	v := os.Getenv(key)

	boolV, _ := strconv.ParseBool(v)
	viper.Set(key, boolV)

	return boolV
}
