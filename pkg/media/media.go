package media

import (
	"io"
	"path"
	"strings"
)

type Media struct {
	Title    string `json:"title"`
	ImageUrl string `json:"imageUrl"`
}

type MediaSource interface {
	io.Closer

	Name() string
	Next(n int) (error, []Media)
}

type MediaCollection []MediaSource

func (c MediaCollection) Collect(messageChannel chan<- Media, n int) {
	for _, mediaSource := range c {
		go func(mediaSource MediaSource) {
			err, messages := mediaSource.Next(n)
			if err != nil {
				return
			}
			for _, m := range messages {
				go downloadAndBroadcast(messageChannel, m, mediaSource)
			}
		}(mediaSource)
	}
}

func downloadAndBroadcast(c chan<- Media, m Media, ms MediaSource) {
	mediaSaveDirectory := path.Join("public", "images", ms.Name())
	err, mediaPath := DownloadResource(m.ImageUrl, mediaSaveDirectory)
	if err != nil {
		return
	}
	m.ImageUrl = "/" + strings.Replace(mediaPath, "public/", "", 1)
	c <- m
}
