package rss

import (
	"fmt"
	"net/url"
	"time"
)

type RssItem struct {
	Title       string
	Source      string
	SourceURL   string
	Link        string
	PublishDate time.Time
	Description string
}

func Parse(urls []string) []RssItem {
	var items []RssItem

	for _, stringURL := range urls {
		//fmt.Println(url)

		//res, err := http.Get(url)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//
		//data, err := io.ReadAll(res.Body)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//defer res.Body.Close()
		//
		////fmt.Println(data)
		//
		//var items Items
		//if err := xml.Unmarshal([]byte(data), &items); err != nil {
		//	panic(err)
		//}
		//
		//fmt.Println("items")
		//fmt.Println(items)
		//
		//for _, item := range items.Items {
		//	fmt.Println(item.Title)
		//	break
		//}
		u, err := url.Parse(stringURL)
		if err != nil {
			panic(err)
		}

		fmt.Println("\n" + u.Host)
		resp, err := Read(stringURL)
		if err != nil {
			fmt.Println(err)
		}

		channel, err := Regular(resp)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(channel.Title)
		for _, item := range channel.Item {
			//time, err := item.PubDate.Parse()
			//if err != nil {
			//	fmt.Println(err)
			//}
			fmt.Println(item.Title + " " + item.Link)
		}
	}

	//var w8 sync.WaitGroup
	//numberOfWorkers := len(urls)
	//w8.Add(numberOfWorkers)
	//
	//inputCh := make(chan string, 10)
	////outputCh := make(chan RssItem, len(dayStats))
	//
	//for k := 0; k < numberOfWorkers; k++ {
	//	go worker(inputCh, k, &w8)
	//}
	//
	//for _, url := range urls {
	//	inputCh <- url
	//}
	//close(inputCh)
	//
	//w8.Wait()

	return items
}

//func worker(in chan string, workerId int, w8 *sync.WaitGroup) {
//	for url := range in {
//		fmt.Println(url)
//	}
//}
