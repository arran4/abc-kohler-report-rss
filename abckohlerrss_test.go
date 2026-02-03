package abcrss

import (
	"encoding/xml"
	"testing"
)

func TestStructs(t *testing.T) {
	// Test RSS struct instantiation
	rss := RSS{
		Version: "2.0",
		Channel: Channel{
			Title:       "Test Feed",
			Link:        "http://example.com",
			Description: "A test feed",
			Items: []Item{
				{
					Title:       "Test Item",
					Link:        "http://example.com/item",
					Description: "A test item",
					PubDate:     "Mon, 02 Jan 2006 15:04:05 MST",
					GUID:        "http://example.com/item",
					Thumbnail:   "http://example.com/image.jpg",
				},
			},
		},
	}

	if rss.Version != "2.0" {
		t.Errorf("Expected version 2.0, got %s", rss.Version)
	}

	// Test XML marshalling
	output, err := xml.Marshal(rss)
	if err != nil {
		t.Fatalf("Failed to marshal RSS: %v", err)
	}

	expected := `<rss version="2.0"><channel><title>Test Feed</title><link>http://example.com</link><description>A test feed</description><item><title>Test Item</title><link>http://example.com/item</link><description>A test item</description><pubDate>Mon, 02 Jan 2006 15:04:05 MST</pubDate><guid>http://example.com/item</guid><thumbnail>http://example.com/image.jpg</thumbnail></item></channel></rss>`
	if string(output) != expected {
		t.Errorf("XML output mismatch.\nExpected: %s\nGot:      %s", expected, string(output))
	}
}
