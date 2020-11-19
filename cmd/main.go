package main

import (
	"log"
	"mmt/example/pkg/media"
	"mmt/example/pkg/media/reddit"
	"mmt/example/pkg/web"
)

func catCollection() []media.MediaSource {
	return []media.MediaSource{
		reddit.New("cat"),
		reddit.New("cats"),
		reddit.New("awww"),
	}
}

func memeCollection() []media.MediaSource {
	return []media.MediaSource{
		reddit.New("ProgrammerHumor"),
		reddit.New("PrequelMemes"),
		reddit.New("OTMemes"),
		reddit.New("dankmemes"),
		reddit.New("memes"),
	}
}

func main() {
	/*mediaSources := []web.MediaSource{
		web.New("PrequelMemes"),
		web.New("dankmemes"),
		web.New("shitposting"),
		web.New("memes"),
	}

	ticker := time.NewTicker(5 * time.Second)
	for {
		<-ticker.C
		CrawlSubs(mediaSources)
	}*/

	const loadNPosts = 5
	mediaSources := catCollection()

	log.Fatal(web.Serve(8080, mediaSources, loadNPosts))
}

/*
func CrawlSubs(mediaSources []web.MediaSource) {
	for _, mediaSource := range mediaSources {
		err, messages := mediaSource.Next(5)
		if err != nil {
			continue
		}
		for _, m := range messages {
			err, mediaPath := SaveMedia(m, mediaSource)
			if err == nil {
				fmt.Printf("Downloaded file %s", mediaPath)
			}
		}
	}
}

func SaveMedia(m web.Media, source web.MediaSource) (error, string) {
	mediaSaveDirectory := path.Join("public", "images", source.Name())
	err, mediaPath := web.DownloadResource(m.ImageUrl, mediaSaveDirectory)
	if err != nil {
		return err, ""
	}
	return nil, mediaPath
}
*/
