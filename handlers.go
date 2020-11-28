package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/minkj1992/go_nomad/scrapper"
)

// Account is for user (private fields)
type file struct {
	term   string
	format string
	name   string
}

// NewAccount is factory to make Account
func newFile(term, fmt string) *file {
	fileName := term + "." + fmt
	f := file{term: term, format: fmt, name: fileName}
	return &f
}

func showSearchPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{"title": "Job Scrapper"}, "index.html")
}

func scrapeJobs(c *gin.Context) {
	// https://github.com/gin-gonic/gin#querystring-parameters
	term := strings.ToLower(scrapper.CleanString(c.Query("term")))
	fileFmt := strings.ToLower(scrapper.CleanString(c.Query("choices-file-fmt")))

	f := newFile(term, fileFmt)
	downloadPath := "downloads/" + f.name
	scrapper.Scrape(f.term, f.format)
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	// c.Header("Content-Disposition", "attachment; filename="+f.name)
	c.Header("Content-Type", "application/octet-stream")
	c.FileAttachment(downloadPath, f.name)
	// defer os.Remove(f.name)
}
