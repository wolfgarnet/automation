package main

import "github.com/wolfgarnet/automation/html2"

func main() {
	c := html2.NewCrawler()
	c.Crawl("http://www.ejbyurterne.dk/index.php")
}
