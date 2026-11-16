package main

// from colly readme on github: https://github.com/gocolly/colly
import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"

	"github.com/gocolly/colly/v2"
)

type imagesUrl []string

func main() {
	const postImagesRegex string = `https://images[1-9].vinted.net/t/(\w*/)*?\w*.webp(\?s=\w*)?`

	// ask user for vinted url or share link
	var url string

	fmt.Print("Enter vinted url or share link: ")
	fmt.Scanln(&url)

	// create scrapper instance
	c := colly.NewCollector()

	// find all images from the website and look if they match the patter
	seen := make(map[string]bool) // avoid duplicates
	var images imagesUrl

	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		found_url := e.Attr("src")
		match, _ := regexp.MatchString(postImagesRegex, found_url)

		if match && !seen[found_url] {
			seen[found_url] = true
			images = append(images, found_url)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Printf("Found %d images on the page.\n", len(images))

		for i := 0; i < len(images); i++ {
			downloadImages(images[i], i)
		}

		fmt.Printf("Done scrapping.\nClosing the program...")
		os.Exit(0)
	})

	// navigate to the user url
	c.Visit(url)
}

// example took from https://siongui.github.io/2018/10/10/go-download-file-from-url/
func downloadImages(url string, index int) {
	filename := path.Base("images/image_" + strconv.Itoa(index) + ".webp")
	fmt.Println("Downloading ", url, " to ", filename)

	response, err := http.Get(url)
	if err != nil {
		return
	}

	defer response.Body.Close()

	f, err := os.Create(filename)
	if err != nil {
		return
	}
	defer f.Close()

	_, err = io.Copy(f, response.Body)
}
