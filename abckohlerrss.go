package abcrss

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// RSS defines the structure of the RSS feed.
type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

// Channel represents the RSS channel.
type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

// Item represents an RSS feed item.
type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	GUID        string `xml:"guid"`
	Thumbnail   string `xml:"thumbnail"`
}

// ABCJSON represents the JSON structure extracted from the page.
type ABCJSON struct {
	Props struct {
		PageProps struct {
			Gtm struct {
				NoGtm bool `json:"noGtm"`
			} `json:"gtm"`
			SnowplowPrepared struct {
				SnowplowCollectorURL string `json:"snowplowCollectorURL"`
				AppVersionID         string `json:"appVersionId"`
				AppID                string `json:"appId"`
				Environment          string `json:"environment"`
			} `json:"snowplowPrepared"`
			Environment struct {
				APIKey struct {
					Mapbox string `json:"mapbox"`
				} `json:"apiKey"`
				Application    string `json:"application"`
				AppEnvironment string `json:"appEnvironment"`
				Layout         struct {
					IsNewsApp bool `json:"isNewsApp"`
					IsFuture  bool `json:"isFuture"`
				} `json:"layout"`
			} `json:"environment"`
			Experimentsconfig struct {
				Experiments struct {
					StateEditions struct {
						Enabled  bool `json:"enabled"`
						Variants struct {
							A int `json:"a"`
						} `json:"variants"`
					} `json:"state-editions"`
					BetaOptin struct {
						Enabled  bool `json:"enabled"`
						Variants struct {
							A int `json:"a"`
						} `json:"variants"`
					} `json:"beta-optin"`
					BetaForced struct {
						Enabled  bool `json:"enabled"`
						Variants struct {
							A int `json:"a"`
						} `json:"variants"`
					} `json:"beta-forced"`
					ForYou struct {
						Enabled  bool `json:"enabled"`
						Variants struct {
							A int `json:"a"`
						} `json:"variants"`
					} `json:"for-you"`
					AmNewsNewsletterArticleEmbed struct {
						Enabled  bool `json:"enabled"`
						Variants struct {
							A int `json:"a"`
						} `json:"variants"`
					} `json:"am-news-newsletter-article-embed"`
					HomepageTrendingView struct {
						Enabled  bool `json:"enabled"`
						Variants struct {
							A int `json:"a"`
							B int `json:"b"`
						} `json:"variants"`
					} `json:"homepage-trending-view"`
				} `json:"experiments"`
			} `json:"experimentsconfig"`
			SiteNavigation []struct {
				LinkTo   string `json:"linkTo"`
				Children string `json:"children"`
				Active   bool   `json:"active"`
			} `json:"siteNavigation"`
			FooterNavigation []struct {
				LinkHref  string `json:"linkHref"`
				LinkTitle string `json:"linkTitle"`
			} `json:"footerNavigation"`
			Ticker   []any `json:"ticker"`
			Profiles struct {
				Environment string `json:"environment"`
				APIKey      string `json:"APIKey"`
				AnalyticsID string `json:"analyticsID"`
			} `json:"profiles"`
			DeveloperFlags []any `json:"developerFlags"`
			Toaster        struct {
				Show            bool   `json:"show"`
				LocalStorageKey string `json:"localStorageKey"`
				BackgroundColor string `json:"backgroundColor"`
				Title           string `json:"title"`
				Description     string `json:"description"`
				ImgSrc          string `json:"imgSrc"`
				CtaPrimary      struct {
					LinkTo          string `json:"linkTo"`
					LinkText        string `json:"linkText"`
					BackgroundColor string `json:"backgroundColor"`
				} `json:"ctaPrimary"`
				ModuleURI   string `json:"moduleURI"`
				Dismissable bool   `json:"dismissable"`
			} `json:"toaster"`
			UnstructuredTracking bool `json:"unstructuredTracking"`
			LoginBanner          bool `json:"loginBanner"`
			Location             struct {
				Pathname string `json:"pathname"`
				Search   string `json:"search"`
			} `json:"location"`
			Channelpage struct {
				HeadTagsPagePrepared struct {
					BasePath     string `json:"basePath"`
					FaviconPath  string `json:"faviconPath"`
					CanonicalURL string `json:"canonicalURL"`
					Description  string `json:"description"`
					Keywords     []any  `json:"keywords"`
					Lang         string `json:"lang"`
					Robots       any    `json:"robots"`
					Title        string `json:"title"`
				} `json:"headTagsPagePrepared"`
				HeadTagsCMSPrepared struct {
					ContentSource string `json:"contentSource"`
					DocType       string `json:"docType"`
					Generator     string `json:"generator"`
					ID            string `json:"id"`
					Version       int    `json:"version"`
				} `json:"headTagsCMSPrepared"`
				HeadTagsSocialPrepared struct {
					CanonicalURL string `json:"canonicalURL"`
					Description  string `json:"description"`
					Image        string `json:"image"`
					OgType       string `json:"ogType"`
					Title        string `json:"title"`
					Twitter      string `json:"twitter"`
				} `json:"headTagsSocialPrepared"`
				FixedHeaderPrepared struct {
					Title      string `json:"title"`
					HasShare   bool   `json:"hasShare"`
					BrandProps struct {
					} `json:"brandProps"`
					HideContext     bool `json:"hideContext"`
					NavigationItems []struct {
						LinkTo   string `json:"linkTo"`
						Active   bool   `json:"active"`
						Children string `json:"children"`
					} `json:"navigationItems"`
				} `json:"fixedHeaderPrepared"`
				DataLayer struct {
					Document struct {
						Language      string `json:"language"`
						CanonicalURL  string `json:"canonicalUrl"`
						ContentType   string `json:"contentType"`
						URI           string `json:"uri"`
						ContentSource string `json:"contentSource"`
						ID            string `json:"id"`
						Title         struct {
							Title string `json:"title"`
						} `json:"title"`
						PageTitle string `json:"pageTitle"`
						SiteRoot  struct {
							Segment string `json:"segment"`
							Title   string `json:"title"`
						} `json:"siteRoot"`
					} `json:"document"`
				} `json:"dataLayer"`
				Subbanner struct {
					DataLayerPrepared struct {
						URI           string `json:"uri"`
						ModuleURI     string `json:"moduleUri"`
						ContentSource string `json:"contentSource"`
						ContentType   string `json:"contentType"`
						Items         []struct {
							URI string `json:"uri"`
						} `json:"items"`
						ID string `json:"id"`
					} `json:"dataLayerPrepared"`
					Heading              string `json:"heading"`
					NavItemPreparedArray []struct {
						LinkTo     string `json:"linkTo"`
						Children   string `json:"children"`
						Active     bool   `json:"active"`
						ContentURI string `json:"contentUri"`
					} `json:"navItemPreparedArray"`
				} `json:"subbanner"`
				PageProps struct {
					Webview bool `json:"webview"`
				} `json:"pageProps"`
				SubNavigation     []any `json:"subNavigation"`
				BannerComponents  []any `json:"bannerComponents"`
				HeaderComponents  []any `json:"headerComponents"`
				SidebarComponents []any `json:"sidebarComponents"`
				Components        []struct {
					DocumentID string `json:"documentId"`
					Background string `json:"background,omitempty"`
					Theme      string `json:"theme,omitempty"`
					Component  struct {
						Name  string `json:"name"`
						Props struct {
							DataLayerPrepared struct {
								URI           string `json:"uri"`
								ModuleURI     string `json:"moduleUri"`
								ContentSource string `json:"contentSource"`
								ContentType   string `json:"contentType"`
								ID            string `json:"id"`
								Title         struct {
									Title string `json:"title"`
								} `json:"title"`
								Items []struct {
									URI string `json:"uri"`
								} `json:"items"`
							} `json:"dataLayerPrepared"`
							Heading string `json:"heading"`
							Layout  string `json:"layout"`
							Video   struct {
								ID     string `json:"id"`
								Config struct {
									Autostart string `json:"autostart"`
									Mute      bool   `json:"mute"`
									Sources   []struct {
										File     string `json:"file"`
										Type     string `json:"type"`
										FileSize int    `json:"fileSize"`
										Bitrate  int    `json:"bitrate"`
										Label    string `json:"label"`
									} `json:"sources"`
									Image      string `json:"image"`
									StreamType string `json:"streamType"`
								} `json:"config"`
								AnalyticsData struct {
									ContentID     string `json:"contentID"`
									ContentSource string `json:"contentSource"`
									ContentType   string `json:"contentType"`
									Duration      int    `json:"duration"`
									StreamType    string `json:"streamType"`
									URI           string `json:"uri"`
								} `json:"analyticsData"`
								Skin struct {
									Name string `json:"name"`
									URL  string `json:"url"`
								} `json:"skin"`
								Title string `json:"title"`
								Share struct {
									HasPopover bool `json:"hasPopover"`
									SendBy     struct {
										Title string `json:"title"`
										Items []struct {
											ShareBy string `json:"shareBy"`
										} `json:"items"`
									} `json:"sendBy"`
									ShareLink string `json:"shareLink"`
									ShareOn   struct {
										Title string `json:"title"`
										Items []struct {
											ShareBy string `json:"shareBy"`
											Via     string `json:"via,omitempty"`
										} `json:"items"`
									} `json:"shareOn"`
									Synopsis             string `json:"synopsis"`
									Title                string `json:"title"`
									UtilityBarShareLinks []struct {
										ShareBy string `json:"shareBy"`
									} `json:"utilityBarShareLinks"`
									UtmSource string `json:"utmSource"`
								} `json:"share"`
								PostDate struct {
									PublishedDate time.Time `json:"publishedDate"`
									UpdatedDate   any       `json:"updatedDate"`
								} `json:"postDate"`
							} `json:"video"`
							List []struct {
								Card struct {
									ContentURI string `json:"contentUri"`
									ID         string `json:"id"`
									Duration   int    `json:"duration"`
									Image      struct {
										ImgSrc   string `json:"imgSrc"`
										ImgRatio string `json:"imgRatio"`
									} `json:"image"`
									Title struct {
										Children string `json:"children"`
									} `json:"title"`
								} `json:"card"`
								Player struct {
									ID     string `json:"id"`
									Config struct {
										Autostart bool `json:"autostart"`
										Mute      bool `json:"mute"`
										Sources   []struct {
											File     string `json:"file"`
											Type     string `json:"type"`
											FileSize int    `json:"fileSize"`
											Bitrate  int    `json:"bitrate"`
											Label    string `json:"label"`
										} `json:"sources"`
										Image      string `json:"image"`
										StreamType string `json:"streamType"`
									} `json:"config"`
									AnalyticsData struct {
										ContentID     string `json:"contentID"`
										ContentSource string `json:"contentSource"`
										ContentType   string `json:"contentType"`
										Duration      int    `json:"duration"`
										StreamType    string `json:"streamType"`
										URI           string `json:"uri"`
									} `json:"analyticsData"`
									Skin struct {
										Name string `json:"name"`
										URL  string `json:"url"`
									} `json:"skin"`
									Title string `json:"title"`
									Share struct {
										HasPopover bool `json:"hasPopover"`
										SendBy     struct {
											Title string `json:"title"`
											Items []struct {
												ShareBy string `json:"shareBy"`
											} `json:"items"`
										} `json:"sendBy"`
										ShareLink string `json:"shareLink"`
										ShareOn   struct {
											Title string `json:"title"`
											Items []struct {
												ShareBy string `json:"shareBy"`
												Via     string `json:"via,omitempty"`
											} `json:"items"`
										} `json:"shareOn"`
										Synopsis             string `json:"synopsis"`
										Title                string `json:"title"`
										UtilityBarShareLinks []struct {
											ShareBy string `json:"shareBy"`
										} `json:"utilityBarShareLinks"`
										UtmSource string `json:"utmSource"`
									} `json:"share"`
									PostDate struct {
										PublishedDate time.Time `json:"publishedDate"`
										UpdatedDate   any       `json:"updatedDate"`
									} `json:"postDate"`
								} `json:"player"`
							} `json:"list"`
							MoreLinks []any `json:"moreLinks"`
						} `json:"props"`
						IsRecursive bool `json:"isRecursive"`
						IsSidebar   bool `json:"isSidebar"`
					} `json:"component"`
				} `json:"components"`
				Scripts       []any `json:"scripts"`
				DynamicTicker any   `json:"dynamicTicker"`
			} `json:"channelpage"`
		} `json:"pageProps"`
		NSsp bool `json:"__N_SSP"`
	} `json:"props"`
	Page  string `json:"page"`
	Query struct {
		Channel []string `json:"channel"`
	} `json:"query"`
	BuildID               string `json:"buildId"`
	AssetPrefix           string `json:"assetPrefix"`
	IsFallback            bool   `json:"isFallback"`
	IsExperimentalCompile bool   `json:"isExperimentalCompile"`
	DynamicIds            []int  `json:"dynamicIds"`
	Gssp                  bool   `json:"gssp"`
	AppGip                bool   `json:"appGip"`
	ScriptLoader          []any  `json:"scriptLoader"`
}

const BaseURL = "https://www.abc.net.au"

func FetchAndParseToRSS() (RSS, error) {
	resp, err := http.Get(BaseURL + "/news/programs/kohler-report")
	if err != nil {
		return RSS{}, fmt.Errorf("fetching news to rss: %v", err)
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			log.Printf("Failed to close body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return RSS{}, fmt.Errorf("status code: %v", resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return RSS{}, fmt.Errorf("parsing news to rss: %v", err)
	}

	rss := RSS{
		Version: "2.0",
		Channel: Channel{},
	}

	// Extract JSON data from a <script> tag.
	jsonData := ""
	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		if strings.Contains(s.Text(), "\"components\"") {
			jsonData = s.Text()
		}
	})

	if jsonData == "" {
		return RSS{}, fmt.Errorf("no JSON data found")
	}

	var abcData ABCJSON
	if err := json.Unmarshal([]byte(jsonData), &abcData); err != nil {
		return RSS{}, fmt.Errorf("parsing JSON data: %v", err)
	}

	// Extract feed header information
	rss.Channel.Title = abcData.Props.PageProps.Channelpage.HeadTagsPagePrepared.Title
	rss.Channel.Link = abcData.Props.PageProps.Channelpage.HeadTagsPagePrepared.CanonicalURL
	rss.Channel.Description = abcData.Props.PageProps.Channelpage.HeadTagsPagePrepared.Description

	seenGUIDs := map[string]bool{}

	// Process components for items
	for _, component := range abcData.Props.PageProps.Channelpage.Components {
		if component.Component.Name == "VideoPlayer" {
			video := component.Component.Props.Video
			//pubDate, err := time.Parse(time.RFC3339, video.PostDate.PublishedDate)
			//if err != nil {
			//	log.Printf("Failed to parse publication date: %v", err)
			//	pubDate = time.Now()
			//}

			{
				guid := video.Share.ShareLink
				if !seenGUIDs[guid] {
					seenGUIDs[guid] = true

					rss.Channel.Items = append(rss.Channel.Items, Item{
						Title:       video.Title,
						Link:        video.Share.ShareLink,
						Description: video.Share.Synopsis,
						PubDate:     video.PostDate.PublishedDate.Format(time.RFC1123),
						GUID:        guid,
						Thumbnail:   video.Config.Image,
					})
				}
			}

			for _, listItem := range component.Component.Props.List {
				guid := listItem.Player.Share.ShareLink
				if seenGUIDs[guid] {
					continue
				}
				seenGUIDs[guid] = true

				rss.Channel.Items = append(rss.Channel.Items, Item{
					Title:       listItem.Card.Title.Children,
					Link:        guid,
					Description: listItem.Player.Share.Synopsis,
					PubDate:     listItem.Player.PostDate.PublishedDate.Format(time.RFC1123),
					GUID:        guid,
					Thumbnail:   video.Config.Image,
				})
			}
		}
	}

	return rss, nil
}
