package main

import (
		"fmt"
		"net/http"
		"github.com/google/google-api-go-client/googleapi/transport"
		"google.golang.org/api/youtube/v3"
)

type Video struct {
		ID, Title, File string
}

type SearchResponse struct {
		query     string
		nextToken string
		videos    []Video
}

var client = &http.Client{
		Transport: &transport.APIKey{Key: "<your api key>"},
}

func main() {

		const query = "cat"
		const maxResults = 10 //MAX 50

		sr, err := search(query, maxResults)
		if err != nil {
				fmt.Println(err);
		} else {
				convert(sr.videos)
		}
}

func convert(videos[] Video) {
		for _, v := range videos {
				fmt.Print(v.ID);
				fmt.Print(v.Title);
				fmt.Print(v.File);
		}
}

func search(term string, num int64) (SearchResponse, error) {

		videos := []Video{}
		searchResponse := SearchResponse{}

		service, err := youtube.New(client)
		if err != nil {
				return searchResponse, err
		}

		call := service.Search.List("id,snippet").
		Q(term).
		Type("video").
		MaxResults(num)
		response, err := call.Do()

		if err != nil {
				return searchResponse, err
		}

		for _, item := range response.Items {
				switch item.Id.Kind {
				case "youtube#video":
						videos = append(videos, Video{item.Id.VideoId, item.Snippet.Title, ""})
				}
		}

		searchResponse.query = term;
		searchResponse.nextToken = response.NextPageToken;
		searchResponse.videos = videos;
		return searchResponse, nil
}
