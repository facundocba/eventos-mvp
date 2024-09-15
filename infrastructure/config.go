package infrastructure

import "os"

func GetDBConfig() string {
	return os.Getenv("DATABASE_URL")
}
