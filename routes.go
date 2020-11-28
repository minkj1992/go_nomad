package main

func initializeRoutes() {
	// Handle the index route
	router.GET("/", showSearchPage)
	router.GET("/scrape", scrapeJobs)
}
