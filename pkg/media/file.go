package media

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path"
)

func DownloadResource(url, targetDirectory string) (error, string) {
	r, err := http.Get(url)
	if err != nil {
		return err, ""
	}

	fileName := path.Base(r.Request.URL.Path)
	targetPath := path.Join(targetDirectory, fileName)
	err = os.MkdirAll(targetDirectory, 0700)
	if err != nil {
		return err, ""
	}

	if r.StatusCode != 200 {
		return errors.New("response error"), ""
	}
	f, err := os.Create(targetPath)
	if err != nil {
		return err, ""
	}

	_, err = io.Copy(f, r.Body)
	return err, targetPath
}
