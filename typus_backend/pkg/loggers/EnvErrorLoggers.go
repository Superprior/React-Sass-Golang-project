package loggers

import "log"

func LogEnvError(varName string) {
	log.Printf("ERROR: COULD NOT LOAD '%s' VARIABLE FROM .env FILE\n", varName)
}
