package reddit

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mmt/example/pkg/media"
	"net/http"
	"strings"
	"time"
)

type apiPosting struct {
	Data struct {
		Subreddit           string `json:"subreddit_name_prefixed"`
		IsRedditMediaDomain bool   `json:"is_reddit_media_domain"`
		Title               string `json:"title"`
		Url                 string `json:"url"`
	}
}

func (p apiPosting) Message() media.Media {
	return media.Media{
		Title:    p.Data.Subreddit + " " + p.Data.Title,
		ImageUrl: p.Data.Url,
	}
}

type apiResponse struct {
	Data struct {
		Children []apiPosting
		After    string
	}
}

type Source struct {
	sub      string
	lastPost string
	interval time.Duration
}

func (s *Source) Close() error {
	return nil
}

func (s *Source) Next(n int) (error, []media.Media) {
	url := fmt.Sprintf("https://www.reddit.com/r/%s/top.json?after=%s&limit=%d", s.sub, s.lastPost, n)
	c := http.Client{
		Timeout: 5 * time.Second,
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err, nil
	}
	request.Header.Set("User-Agent", "fhs mmt meme-bot:1.0.0 (TEST)")
	r, err := c.Do(request)
	if err != nil {
		return err, nil
	}
	if r.StatusCode != 200 {
		return errors.New("api error"), nil
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err, nil
	}

	var ar apiResponse
	err = json.Unmarshal(b, &ar)
	if err != nil {
		return err, nil
	}

	var messages []media.Media
	for _, c := range ar.Data.Children {
		containsImageExtension := strings.Index(c.Data.Url, ".jpg") != -1 || strings.Index(c.Data.Url, ".png") != -1
		if c.Data.IsRedditMediaDomain && containsImageExtension {
			messages = append(messages, c.Message())
		}
	}
	s.lastPost = ar.Data.After
	return nil, messages
}

func (s *Source) Name() string {
	return "reddit_" + strings.ToLower(s.sub)
}

func New(sub string) *Source {
	return &Source{
		sub: sub,
	}
}

var _ = media.MediaSource(&Source{})
