package commom

import (
	"fmt"
	"strings"
	"testing"
)

func TestConvertTaggedContentImg(t *testing.T) {
	result := `<img src="{STEAM_CLAN_IMAGE}/3703047/e5a57f7b21063e63c40a4073903fe02560ebe95c.png">`
	simpleTag := "[img]{STEAM_CLAN_IMAGE}/3703047/e5a57f7b21063e63c40a4073903fe02560ebe95c.png[/img]"

	newsItems := []NewsItem{}

	item := NewsItem{}
	item.Contents = simpleTag

	item2 := NewsItem{}
	item2.Contents = "aaaaaaaaaaaaaaaaaa" + simpleTag

	item3 := NewsItem{}
	item3.Contents = simpleTag + "aaaaaaaaaaaaaaaaaa"

	item4 := NewsItem{}
	item4.Contents = "aaaaaaaaaaaaaaaaaa" + simpleTag + "aaaaaaaaaaaaaaaaaa"

	item5 := NewsItem{}
	item5.Contents = `[img]{STEAM_CLAN_IMAGE}/3703047/e5a57f7b21063e63c40a4073903fe02560ebe95c.png[/img] Today Marci powers her way from [url=https://www.netflix.com/title/80994336]DOTA: Dragon's Blood[/url] into the battle of the Ancients as the newest hero to join the fight, proving that undying loyalty yields unrivaled power. Marci marches into battle ready to raise fists in defense of her companions. Head over to the [url=https://www.dota2.com/marci]Marci update page[/url] to learn about her abilities and more. [h3]Gameplay Update 7.30e[/h3] Also included in todays update is gameplay update 7.30e. Check out all the details on the [url=https://www.dota2.com/patches/7.30e]update page[/url]. [h3]Treasure of the Wordless Trek[/h3] The Treasure of the Wordless Trek is now for sale for $2.49. This all-new treasure features sets for Zeus, Sven, Puck, Lina, Brewmaster, Clockwerk, Lich, Pangolier, Templar Assassin, and Ancient Apparation, as well as a chance at a Very Rare Spectre set, and Extremely Rare Razor set. [h3]The International Grand Champions[/h3] From Eastern Europe Qualifiers to The International Champions, we'd like to officially congratulate Team Spirit on their monumental accomplishment. In one of the greatest Finals ever witnessed, Team Spirit triumphed in a winner-take-all game five against arguably the most feared Dota team in the world, PSG.LGD. Fittingly, the lineups for the deciding game set two dominant strategies of the tournament against each other, and in the end Team Spirit's Magnus-first draft was too much for PSG.LGD's vaunted Tiny/Lycan pair. They claimed the final game on the main stage — and with it the Aegis — just as they had won the first game of the Main Event, with an aggressive and exciting style of play that improved each day and thrilled Dota fans worldwide along the way. As bearers of the ultimate symbol of victory, these names shall forever be inscribed upon the Aegis of Champions: [b]2021 - Team Spirit[/b] [list] [*] Illya "Yatoro" Mulyarchuk [*] Alexander "TORONTOTOKYO" Khertek [*] Magomed "Collapse" Khalilov [*] Miroslaw "Mira" Kolpakov [*] Yaroslav "Miposhka" Naidenov [/list] [h3]Looking Back[/h3] If you missed any of the tournament, or just want to rewatch some of the incredible moments and dazzling plays, head over to [url=https://www.dota2.com/esports/ti10/]The International website[/url], where you can find replays for every match. There are photos and videos from the event over on our [url=https://www.youtube.com/playlist?list=PLxkyNsoBqOdDhwxZ9jT2gEzzmS5FSoGSz]Youtube Main Event[/url] and [url=https://www.youtube.com/playlist?list=PLxkyNsoBqOdA5JA7aZpOhQ8XSGjR769Sx]Event Entertainment[/url] playlists, [url=https://www.instagram.com/dota2]Dota 2 Instagram[/url], and the [url=https://www.flickr.com/photos/dota2ti/]Dota 2 Flickr[/url] — which also features content stretching back to the earliest incarnations of TI. We would like to thank all of the players, talent, and everyone in the Dota community for once again helping bring this celebration to life. The road to The International was not straightforward this time around, but we hope you enjoyed the world-class Dota and all of the broadcast content from this year's tournament as much as we did.`

	newsItems = append(newsItems, item)
	newsItems = append(newsItems, item2)
	newsItems = append(newsItems, item3)
	newsItems = append(newsItems, item4)

	news := News{}
	news.NewsItems = newsItems

	appNews := AppNews{}
	appNews.News = news

	appNews.ConvertTaggedContent()

	for _, item := range appNews.News.NewsItems {
		fmt.Println(item.Contents)
		if !strings.Contains(item.Contents, result) {
			t.Fatalf("TestConvertTaggedContentImg - %s - %s", item.Contents, result)
		}
	}
}

func TestConvertTaggedContentUrl(t *testing.T) {
	result := `<a href="https://www.netflix.com/title/80994336">DOTA: Dragon's Blood</a>`
	simpleTag := "[url=https://www.netflix.com/title/80994336]DOTA: Dragon's Blood[/url]"

	newsItems := []NewsItem{}

	item := NewsItem{}
	item.Contents = simpleTag

	item2 := NewsItem{}
	item2.Contents = "aaaaaaaaaaaaaaaaaa" + simpleTag

	item3 := NewsItem{}
	item3.Contents = simpleTag + "aaaaaaaaaaaaaaaaaa"

	item4 := NewsItem{}
	item4.Contents = "aaaaaaaaaaaaaaaaaa" + simpleTag + "aaaaaaaaaaaaaaaaaa"

	item5 := NewsItem{}
	item5.Contents = `[img]{STEAM_CLAN_IMAGE}/3703047/e5a57f7b21063e63c40a4073903fe02560ebe95c.png[/img] Today Marci powers her way from [url=https://www.netflix.com/title/80994336]DOTA: Dragon's Blood[/url] into the battle of the Ancients as the newest hero to join the fight, proving that undying loyalty yields unrivaled power. Marci marches into battle ready to raise fists in defense of her companions. Head over to the [url=https://www.dota2.com/marci]Marci update page[/url] to learn about her abilities and more. [h3]Gameplay Update 7.30e[/h3] Also included in todays update is gameplay update 7.30e. Check out all the details on the [url=https://www.dota2.com/patches/7.30e]update page[/url]. [h3]Treasure of the Wordless Trek[/h3] The Treasure of the Wordless Trek is now for sale for $2.49. This all-new treasure features sets for Zeus, Sven, Puck, Lina, Brewmaster, Clockwerk, Lich, Pangolier, Templar Assassin, and Ancient Apparation, as well as a chance at a Very Rare Spectre set, and Extremely Rare Razor set. [h3]The International Grand Champions[/h3] From Eastern Europe Qualifiers to The International Champions, we'd like to officially congratulate Team Spirit on their monumental accomplishment. In one of the greatest Finals ever witnessed, Team Spirit triumphed in a winner-take-all game five against arguably the most feared Dota team in the world, PSG.LGD. Fittingly, the lineups for the deciding game set two dominant strategies of the tournament against each other, and in the end Team Spirit's Magnus-first draft was too much for PSG.LGD's vaunted Tiny/Lycan pair. They claimed the final game on the main stage — and with it the Aegis — just as they had won the first game of the Main Event, with an aggressive and exciting style of play that improved each day and thrilled Dota fans worldwide along the way. As bearers of the ultimate symbol of victory, these names shall forever be inscribed upon the Aegis of Champions: [b]2021 - Team Spirit[/b] [list] [*] Illya "Yatoro" Mulyarchuk [*] Alexander "TORONTOTOKYO" Khertek [*] Magomed "Collapse" Khalilov [*] Miroslaw "Mira" Kolpakov [*] Yaroslav "Miposhka" Naidenov [/list] [h3]Looking Back[/h3] If you missed any of the tournament, or just want to rewatch some of the incredible moments and dazzling plays, head over to [url=https://www.dota2.com/esports/ti10/]The International website[/url], where you can find replays for every match. There are photos and videos from the event over on our [url=https://www.youtube.com/playlist?list=PLxkyNsoBqOdDhwxZ9jT2gEzzmS5FSoGSz]Youtube Main Event[/url] and [url=https://www.youtube.com/playlist?list=PLxkyNsoBqOdA5JA7aZpOhQ8XSGjR769Sx]Event Entertainment[/url] playlists, [url=https://www.instagram.com/dota2]Dota 2 Instagram[/url], and the [url=https://www.flickr.com/photos/dota2ti/]Dota 2 Flickr[/url] — which also features content stretching back to the earliest incarnations of TI. We would like to thank all of the players, talent, and everyone in the Dota community for once again helping bring this celebration to life. The road to The International was not straightforward this time around, but we hope you enjoyed the world-class Dota and all of the broadcast content from this year's tournament as much as we did.`

	newsItems = append(newsItems, item)
	newsItems = append(newsItems, item2)
	newsItems = append(newsItems, item3)
	newsItems = append(newsItems, item4)

	news := News{}
	news.NewsItems = newsItems

	appNews := AppNews{}
	appNews.News = news

	appNews.ConvertTaggedContent()

	for _, item := range appNews.News.NewsItems {
		if !strings.Contains(item.Contents, result) {
			t.Fatalf("TestConvertTaggedContentImg - %s - %s", item.Contents, result)
		}
	}
}

func TestConvertTaggedContentBold(t *testing.T) {
	result := `<strong>2021 - Team Spirit</strong>`
	simpleTag := "[b]2021 - Team Spirit[/b]"

	newsItems := []NewsItem{}

	item := NewsItem{}
	item.Contents = simpleTag

	item2 := NewsItem{}
	item2.Contents = "aaaaaaaaaaaaaaaaaa" + simpleTag

	item3 := NewsItem{}
	item3.Contents = simpleTag + "aaaaaaaaaaaaaaaaaa"

	item4 := NewsItem{}
	item4.Contents = "aaaaaaaaaaaaaaaaaa" + simpleTag + "aaaaaaaaaaaaaaaaaa"

	item5 := NewsItem{}
	item5.Contents = `[img]{STEAM_CLAN_IMAGE}/3703047/e5a57f7b21063e63c40a4073903fe02560ebe95c.png[/img] Today Marci powers her way from [url=https://www.netflix.com/title/80994336]DOTA: Dragon's Blood[/url] into the battle of the Ancients as the newest hero to join the fight, proving that undying loyalty yields unrivaled power. Marci marches into battle ready to raise fists in defense of her companions. Head over to the [url=https://www.dota2.com/marci]Marci update page[/url] to learn about her abilities and more. [h3]Gameplay Update 7.30e[/h3] Also included in todays update is gameplay update 7.30e. Check out all the details on the [url=https://www.dota2.com/patches/7.30e]update page[/url]. [h3]Treasure of the Wordless Trek[/h3] The Treasure of the Wordless Trek is now for sale for $2.49. This all-new treasure features sets for Zeus, Sven, Puck, Lina, Brewmaster, Clockwerk, Lich, Pangolier, Templar Assassin, and Ancient Apparation, as well as a chance at a Very Rare Spectre set, and Extremely Rare Razor set. [h3]The International Grand Champions[/h3] From Eastern Europe Qualifiers to The International Champions, we'd like to officially congratulate Team Spirit on their monumental accomplishment. In one of the greatest Finals ever witnessed, Team Spirit triumphed in a winner-take-all game five against arguably the most feared Dota team in the world, PSG.LGD. Fittingly, the lineups for the deciding game set two dominant strategies of the tournament against each other, and in the end Team Spirit's Magnus-first draft was too much for PSG.LGD's vaunted Tiny/Lycan pair. They claimed the final game on the main stage — and with it the Aegis — just as they had won the first game of the Main Event, with an aggressive and exciting style of play that improved each day and thrilled Dota fans worldwide along the way. As bearers of the ultimate symbol of victory, these names shall forever be inscribed upon the Aegis of Champions: [b]2021 - Team Spirit[/b] [list] [*] Illya "Yatoro" Mulyarchuk [*] Alexander "TORONTOTOKYO" Khertek [*] Magomed "Collapse" Khalilov [*] Miroslaw "Mira" Kolpakov [*] Yaroslav "Miposhka" Naidenov [/list] [h3]Looking Back[/h3] If you missed any of the tournament, or just want to rewatch some of the incredible moments and dazzling plays, head over to [url=https://www.dota2.com/esports/ti10/]The International website[/url], where you can find replays for every match. There are photos and videos from the event over on our [url=https://www.youtube.com/playlist?list=PLxkyNsoBqOdDhwxZ9jT2gEzzmS5FSoGSz]Youtube Main Event[/url] and [url=https://www.youtube.com/playlist?list=PLxkyNsoBqOdA5JA7aZpOhQ8XSGjR769Sx]Event Entertainment[/url] playlists, [url=https://www.instagram.com/dota2]Dota 2 Instagram[/url], and the [url=https://www.flickr.com/photos/dota2ti/]Dota 2 Flickr[/url] — which also features content stretching back to the earliest incarnations of TI. We would like to thank all of the players, talent, and everyone in the Dota community for once again helping bring this celebration to life. The road to The International was not straightforward this time around, but we hope you enjoyed the world-class Dota and all of the broadcast content from this year's tournament as much as we did.`

	newsItems = append(newsItems, item)
	newsItems = append(newsItems, item2)
	newsItems = append(newsItems, item3)
	newsItems = append(newsItems, item4)

	news := News{}
	news.NewsItems = newsItems

	appNews := AppNews{}
	appNews.News = news

	appNews.ConvertTaggedContent()

	for _, item := range appNews.News.NewsItems {
		if !strings.Contains(item.Contents, result) {
			t.Fatalf("TestConvertTaggedContentImg - %s - %s", item.Contents, result)
		}
	}
}

func TestConvertTaggedContentHeader(t *testing.T) {
	simpleTag := "[h3]Treasure of the Wordless Trek[/h3]"
	result := `<h3>Treasure of the Wordless Trek</h3>`

	newsItems := []NewsItem{}

	item := NewsItem{}
	item.Contents = simpleTag

	item2 := NewsItem{}
	item2.Contents = "aaaaaaaaaaaaaaaaaa" + simpleTag

	item3 := NewsItem{}
	item3.Contents = simpleTag + "aaaaaaaaaaaaaaaaaa"

	item4 := NewsItem{}
	item4.Contents = "aaaaaaaaaaaaaaaaaa" + simpleTag + "aaaaaaaaaaaaaaaaaa"

	item5 := NewsItem{}
	item5.Contents = `[img]{STEAM_CLAN_IMAGE}/3703047/e5a57f7b21063e63c40a4073903fe02560ebe95c.png[/img] Today Marci powers her way from [url=https://www.netflix.com/title/80994336]DOTA: Dragon's Blood[/url] into the battle of the Ancients as the newest hero to join the fight, proving that undying loyalty yields unrivaled power. Marci marches into battle ready to raise fists in defense of her companions. Head over to the [url=https://www.dota2.com/marci]Marci update page[/url] to learn about her abilities and more. [h3]Gameplay Update 7.30e[/h3] Also included in todays update is gameplay update 7.30e. Check out all the details on the [url=https://www.dota2.com/patches/7.30e]update page[/url]. [h3]Treasure of the Wordless Trek[/h3] The Treasure of the Wordless Trek is now for sale for $2.49. This all-new treasure features sets for Zeus, Sven, Puck, Lina, Brewmaster, Clockwerk, Lich, Pangolier, Templar Assassin, and Ancient Apparation, as well as a chance at a Very Rare Spectre set, and Extremely Rare Razor set. [h3]The International Grand Champions[/h3] From Eastern Europe Qualifiers to The International Champions, we'd like to officially congratulate Team Spirit on their monumental accomplishment. In one of the greatest Finals ever witnessed, Team Spirit triumphed in a winner-take-all game five against arguably the most feared Dota team in the world, PSG.LGD. Fittingly, the lineups for the deciding game set two dominant strategies of the tournament against each other, and in the end Team Spirit's Magnus-first draft was too much for PSG.LGD's vaunted Tiny/Lycan pair. They claimed the final game on the main stage — and with it the Aegis — just as they had won the first game of the Main Event, with an aggressive and exciting style of play that improved each day and thrilled Dota fans worldwide along the way. As bearers of the ultimate symbol of victory, these names shall forever be inscribed upon the Aegis of Champions: [b]2021 - Team Spirit[/b] [list] [*] Illya "Yatoro" Mulyarchuk [*] Alexander "TORONTOTOKYO" Khertek [*] Magomed "Collapse" Khalilov [*] Miroslaw "Mira" Kolpakov [*] Yaroslav "Miposhka" Naidenov [/list] [h3]Looking Back[/h3] If you missed any of the tournament, or just want to rewatch some of the incredible moments and dazzling plays, head over to [url=https://www.dota2.com/esports/ti10/]The International website[/url], where you can find replays for every match. There are photos and videos from the event over on our [url=https://www.youtube.com/playlist?list=PLxkyNsoBqOdDhwxZ9jT2gEzzmS5FSoGSz]Youtube Main Event[/url] and [url=https://www.youtube.com/playlist?list=PLxkyNsoBqOdA5JA7aZpOhQ8XSGjR769Sx]Event Entertainment[/url] playlists, [url=https://www.instagram.com/dota2]Dota 2 Instagram[/url], and the [url=https://www.flickr.com/photos/dota2ti/]Dota 2 Flickr[/url] — which also features content stretching back to the earliest incarnations of TI. We would like to thank all of the players, talent, and everyone in the Dota community for once again helping bring this celebration to life. The road to The International was not straightforward this time around, but we hope you enjoyed the world-class Dota and all of the broadcast content from this year's tournament as much as we did.`

	newsItems = append(newsItems, item)
	newsItems = append(newsItems, item2)
	newsItems = append(newsItems, item3)
	newsItems = append(newsItems, item4)

	news := News{}
	news.NewsItems = newsItems

	appNews := AppNews{}
	appNews.News = news

	appNews.ConvertTaggedContent()

	for _, item := range appNews.News.NewsItems {
		if !strings.Contains(item.Contents, result) {
			t.Fatalf("TestConvertTaggedContentImg - %s - %s", item.Contents, result)
		}
	}
}
