package utils

import (
	"fmt"
	"log"
	"os"
)

func SetGithubEnvOutput(key, value string) {
	outputFilename := os.Getenv("GITHUB_OUTPUT")
	f, err := os.OpenFile(outputFilename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)
	if _, err := f.WriteString(fmt.Sprintf("%s=%d\n", key, value)); err != nil {
		log.Println(err)
	}
}
