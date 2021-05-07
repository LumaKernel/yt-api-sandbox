package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	youtubeApiKey := os.Getenv("YOUTUBE_API_KEY")

	fmt.Println("hi")
	httpClient := &http.Client{
		Transport: &transport.APIKey{
			Key: youtubeApiKey,
		},
	}
	svc, err := youtube.New(httpClient)
	if err != nil {
		log.Fatal(err)
	}

	var categoryMap = make(map[string]string)
	{
		call := svc.
			VideoCategories.
			List([]string{}).
			RegionCode("us")
		res, err := call.Do()
		if err != nil {
			log.Fatal(err)
		}

		for _, item := range res.Items {
			categoryMap[item.Id] = item.Snippet.Title
			fmt.Println(item.Snippet.Title)
		}
	}

	// for _, item := range res.Items {
	// 	fmt.Println("")
	// 	fmt.Println(item.Id)
	// }

	call := svc.Videos.List([]string{
		"snippet",
	})
	call.MaxResults(2)
	call.Id("dnO0_ZGOJJY")
	// call.Id("mYfJxlgR2jw")
	res, err := call.Do()
	if err != nil {
		log.Fatal(err)
	}
	snippet := res.Items[0].Snippet
	fmt.Println(snippet.Title)
	fmt.Println(snippet.CategoryId)
	fmt.Println(categoryMap[snippet.CategoryId])
	fmt.Println(snippet.ChannelId)
	cid := snippet.ChannelId

	{
		call := svc.
			Channels.
			List([]string{
				"snippet",
			}).
			Id(cid)
		res, err := call.Do()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res.Items[0].Snippet.Title)
	}
}
