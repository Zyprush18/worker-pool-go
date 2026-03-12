package main

import "fmt"

func main() {
	urls:=  []string{
	// === BERITA INTERNASIONAL ===
	"https://www.reuters.com",
	"https://www.bbc.com/news",
	"https://www.theguardian.com",
	"https://apnews.com",
	"https://www.aljazeera.com",
	"https://www.npr.org/sections/news",
	"https://www.cbsnews.com",
	"https://www.nbcnews.com",
	"https://abcnews.go.com",
	"https://www.usatoday.com",

	// === BERITA TEKNOLOGI ===
	"https://techcrunch.com",
	"https://www.theverge.com",
	"https://www.wired.com",
	"https://arstechnica.com",
	"https://www.zdnet.com",
	"https://www.cnet.com",
	"https://venturebeat.com",
	"https://www.engadget.com",
	"https://thenextweb.com",
	"https://www.tomshardware.com",

	// === BERITA INDONESIA ===
	"https://www.kompas.com",
	"https://www.detik.com",
	"https://www.liputan6.com",
	"https://www.tribunnews.com",
	"https://www.cnnindonesia.com",
	"https://www.tempo.co",
	"https://www.merdeka.com",
	"https://www.okezone.com",
	"https://www.sindonews.com",
	"https://www.antaranews.com",

	// === BERITA BISNIS & KEUANGAN ===
	"https://www.bloomberg.com",
	"https://www.ft.com",
	"https://www.wsj.com",
	"https://fortune.com",
	"https://www.forbes.com",
	"https://www.businessinsider.com",
	"https://www.cnbc.com",
	"https://www.marketwatch.com",
	"https://finance.yahoo.com",
	"https://www.economist.com",

	// === BLOG & ARTIKEL UMUM ===
	"https://medium.com",
	"https://dev.to",
	"https://hashnode.com",
	"https://www.smashingmagazine.com",
	"https://css-tricks.com",
	"https://www.sitepoint.com",
	"https://www.digitalocean.com/community/tutorials",
	"https://scotch.io",
	"https://www.freecodecamp.org/news",
	"https://www.hongkiat.com",

	// === BERITA SAINS & RISET ===
	"https://www.sciencedaily.com",
	"https://phys.org",
	"https://www.newscientist.com",
	"https://www.livescience.com",
	"https://www.scientificamerican.com",
	"https://www.nature.com/news",
	"https://www.space.com",
	"https://futurism.com",
	"https://www.discovermagazine.com",
	"https://www.popularmechanics.com",

	// === FORUM & KOMUNITAS ===
	"https://news.ycombinator.com",
	"https://lobste.rs",
	"https://slashdot.org",
	"https://www.digg.com",
	"https://tildes.net",
	"https://www.kaskus.co.id",

	// === HIBURAN & GAYA HIDUP ===
	"https://www.buzzfeed.com",
	"https://mashable.com",
	"https://www.huffpost.com",
	"https://www.vice.com",
	"https://www.rollingstone.com",
	"https://www.esquire.com",
	"https://www.gq.com",
	"https://www.vogue.com",
	"https://www.menshealth.com",
	"https://www.healthline.com",

	// === OLAHRAGA ===
	"https://www.espn.com",
	"https://sports.yahoo.com",
	"https://www.skysports.com",
	"https://www.goal.com",
	"https://www.bola.net",
	"https://www.fourfourtwo.com",
	"https://bleacherreport.com",
	"https://www.fifa.com/news",
	"https://www.sportingnews.com",

	// === OPEN DATA & PUBLIC API ===
	"https://quotes.toscrape.com",
	"https://books.toscrape.com",
	"https://news.ycombinator.com/newest",
	"https://www.goodreads.com/list",
	"https://www.imdb.com/news",
	"https://openweathermap.org",
	"https://www.worldometers.info",
	"https://www.numbeo.com",
	"https://en.wikipedia.org/wiki/Main_Page",
}



	url := make(chan string, len(urls))
	for _, v := range urls {
		url <- v
	}
	close(url)

	

	worker := NewJob(20, url)
	worker.WorkerFetchPage()
	worker.ParseResultsPage()

	for v := range worker.resultParse {
		fmt.Println(v)
	}

	fmt.Printf("banyak url: %d\n", len(urls))
	fmt.Printf("total yg di kerjakan worker Fetch: %d\n", berapaFetch)
	fmt.Printf("total yg di kerjakan worker Parse: %d\n", berapaParse)
	fmt.Println("Success Scraping")
}