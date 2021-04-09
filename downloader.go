package main

import (
	"github.com/oschwald/geoip2-golang"
	"io"
	"log"
	"net/http"
	"os"
)

const FILE_NAME = "geodb.file"

func DownloadFile(url string, fileName string) error {
	_, e := os.Stat(fileName)
	if e != nil {
		log.Printf("DB file is absent, downloading...")
		f, err := os.Create(fileName)
		if err != nil {
			return err
		}

		// file doesn't exist - download it
		resp, err := http.Get(url)
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		size, err := io.Copy(f, resp.Body)
		if err != nil {
			_ = f.Close()
			_ = os.Remove(fileName)
			return err
		}

		f.Close()

		// now, try to open file
		db, err := geoip2.Open(fileName)
		if err != nil {
			log.Printf("Looks like something is off with the file...")
			// if file has bad format - report it, and delete, so it can be re-downloaded later again
			_ = os.Remove(fileName)
			return err
		}
		db.Close()

		log.Printf("DB file successfully downloaded. Got %v bytes.", size)
	}

	return nil
}
