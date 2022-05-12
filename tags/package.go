package tags

import (
	"strconv"
	"regexp"
	"fmt"
	"strings"
	"github.com/onecthree/nhendl/http"
)

type Defaults struct {
	Source		string
	parsed		string
}

func Default(source string) *Defaults {
	
	return new(Defaults)
}

func (this *Defaults) Clear() {
	re1, _ := regexp.Compile(`<section\sid\=\"tags\"\>((.|\n)*)\<\/section\>`)
	raw_parsed := re1.FindString(this.Source)

	re2 := regexp.MustCompile(`<[^>]*>`)
	html_parsed := re2.ReplaceAllString(raw_parsed, " ")

	re3 := regexp.MustCompile(`\s+`)
	mid_parsed := re3.ReplaceAllString(html_parsed, " ")

	re4 := regexp.MustCompile(`Favorite.*download`)
	this.parsed = re4.ReplaceAllString(mid_parsed, "")
}

func (this *Defaults) title() string {
	var regex, _ = regexp.Compile(`<h1\sclass\=\"title\".*\<\/h1\>`)
	var str = regex.FindString(this.Source)
	re := regexp.MustCompile(`<[^>]*>`)
	title := re.ReplaceAllString(str, "")

	return title
}

func (this *Defaults) sub(tag string) string {
	parodies_re := regexp.MustCompile(tag + `:.*?(\w+\:)`)
	parodies_filter := parodies_re.FindString(this.parsed)
	parodies_re = regexp.MustCompile(tag + `\:|\w+\:$`)
	parodies_filter = parodies_re.ReplaceAllString(parodies_filter, "")

	return parodies_filter
}

func (this *Defaults) Preview(args map[string]string) {
	if _, ok := args["--fulldesc"]; ok {
		fmt.Println("Title       :", this.title())
		fmt.Println("Parodies    :", this.sub("Parodies"))
		fmt.Println("Characters  :", this.sub("Characters"))
		fmt.Println("Tags        :", this.sub("Tags"))
		fmt.Println("Artists     :", this.sub("Artists"))
		fmt.Println("Groups      :", this.sub("Groups"))
		fmt.Println("Languages   :", this.sub("Languages"))
		fmt.Println("Categories  :", this.sub("Categories"))
		fmt.Println("Pages       :", this.sub("Pages"))
		fmt.Println("\n")
	} else {
		fmt.Println("nhendl: found a title for -- ", this.title())
	}
}

func (this *Defaults) Pages() int {
	var regex, _ = regexp.Compile(`[0-9]`)
	var page_str = regex.FindAllString(this.sub("Pages"), 3)
	pages, _ := strconv.Atoi(strings.Join(page_str, ""))
	return pages
}

func (this *Defaults) Host(code string) string {
	get_pages := http.Get("https://nhentai.net/g/" + code + "/1")
	
	host_re := regexp.MustCompile(`fit-both.*?1\.jpg`)
	host_filter := host_re.FindString(get_pages["src"].(string))
	
	host_re = regexp.MustCompile(`^.*src\=\"|1\.jpg`)
	host_filter = host_re.ReplaceAllString(host_filter, "")

	return host_filter
}