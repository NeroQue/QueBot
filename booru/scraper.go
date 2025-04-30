package booru

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Provider type for type safety
type Provider string

// Define our supported providers
const (
	Safebooru Provider = "safebooru"
	Danbooru  Provider = "danbooru"
	Gelbooru  Provider = "gelbooru"
)

// Result struct to hold image data
type Result struct {
	ImageURL string
	Source   string
	Tags     []string
}

// Error definitions
var (
	ErrNoResults       = errors.New("no results found")
	ErrInvalidProvider = errors.New("invalid provider")
	ErrScrapingFailed  = errors.New("failed to scrape website")
)

func URLBuilder(provider Provider, tag string) (string, error) {
	encodedTag := url.QueryEscape(tag)
	switch provider {
	case Safebooru:
		return "https://safebooru.org/index.php?page=post&s=list&tags=" + encodedTag, nil
	case Danbooru:
		return "https://danbooru.donmai.us/posts?tags=is%3Asfw+" + encodedTag, nil
	case Gelbooru:
		return "https://gelbooru.com/index.php?page=post&s=list&tags=rating%3Ageneral+" + encodedTag, nil
	}
	return "", ErrInvalidProvider
}

func getFullImage(provider Provider, postURL string) (string, error) {
	resp, err := http.Get(postURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", ErrScrapingFailed
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	switch provider {
	case Safebooru:
		// Find the full-size image
		imageURL, exists := doc.Find("#image").Attr("src")
		if !exists {
			return "", ErrNoResults
		}
		return imageURL, nil
	case Danbooru:
		imageURL, exists := doc.Find("#image").Attr("src")
		if !exists {
			return "", ErrNoResults
		}
		return imageURL, nil
	case Gelbooru:
		imageURL, exists := doc.Find("#image").Attr("src")
		if !exists {
			return "", ErrNoResults
		}
		return imageURL, nil
	default:
		return "", ErrInvalidProvider
	}
}

func Scrape(provider Provider, tag string) ([]Result, error) {
	url, err := URLBuilder(provider, tag)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, ErrScrapingFailed
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	results := []Result{}

	switch provider {
	case Safebooru:
		// Find all thumbnail spans
		doc.Find("span.thumb").Each(func(i int, s *goquery.Selection) {
			anchor := s.Find("a").First()
			if href, exists := anchor.Attr("href"); exists {
				postURL := "https://safebooru.org/" + href
				if imageURL, err := getFullImage(provider, postURL); err == nil {
					tags := []string{}
					if alt, exists := anchor.Find("img.preview").Attr("alt"); exists {
						tags = strings.Fields(alt)
					}
					results = append(results, Result{
						ImageURL: imageURL,
						Source:   postURL,
						Tags:     tags,
					})
				}
			}
		})
	case Danbooru:
		// Find all post previews
		doc.Find("article.post-preview").Each(func(i int, s *goquery.Selection) {
			// Get tags from data-tags attribute
			tags := strings.Fields(s.AttrOr("data-tags", ""))

			// Get the post URL and fetch the full image
			if href, exists := s.Find("a.post-preview-link").Attr("href"); exists {
				postURL := "https://danbooru.donmai.us" + href
				if imageURL, err := getFullImage(provider, postURL); err == nil {
					results = append(results, Result{
						ImageURL: imageURL,
						Source:   postURL,
						Tags:     tags,
					})
				}
			}
		})
	case Gelbooru:
		// Find all thumbnail previews
		doc.Find("article.thumbnail-preview").Each(func(i int, s *goquery.Selection) {
			anchor := s.Find("a").First()
			if href, exists := anchor.Attr("href"); exists {
				// Gelbooru provides full URLs
				postURL := href
				if imageURL, err := getFullImage(provider, postURL); err == nil {
					// Get tags from the img title attribute
					tags := []string{}
					if title, exists := s.Find("img").Attr("title"); exists {
						tags = strings.Fields(title)
					}
					results = append(results, Result{
						ImageURL: imageURL,
						Source:   postURL,
						Tags:     tags,
					})
				}
			}
		})
	}

	return results, nil
}

func GetRandomImage(provider Provider, tag string) (*Result, error) {
	fmt.Println("Scraping for tag:", tag, "from provider:", provider)
	results, err := Scrape(provider, tag)
	if err != nil {
		return nil, err
	}

	fmt.Println("Found", len(results), "results")

	if len(results) == 0 {
		return nil, ErrNoResults
	}

	return &results[rand.Intn(len(results))], nil
}
