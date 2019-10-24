package main

	

import (
   	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
//	"sync"
    "flag"   
    "strings"
  	"io"
  	"net/http"

)

var (
	searchCombination string // 
	pagesWithFiles map[string]string	
	albumName string
	WORKERS int = 2 //кол-во "потоков"
)


func _check(err error) {
	if err != nil {
		panic(err)
	}
}

// основная функция обработки
func parseUrl(url string) {

	//fmt.Println("request: " + url)
	
	// заворачиваем источник в goquery документ
	doc, err := goquery.NewDocument(url)

	_check(err)

	// в манере jquery, css селектором получаем все ссылки
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		if val, ok := s.Attr("href"); ok {
			//Check if url ends with mp3
			if strings.HasSuffix(val, "mp3") {
				val := "https://downloads.khinsider.com" + val
				// put all url in map (only first in)
				if _,ok :=  pagesWithFiles[val]; !ok {
					pagesWithFiles[val] = s.Text()					
				}
			}
	    }
	})
	//get album nane from first <h2> tag
	doc.Find("h2").EachWithBreak(func(i int, s *goquery.Selection) bool {		
		albumName =  s.Text()
		return false
	}) 	
}

func downloadFilesCNTRL() {
	//var countOfWorkers int
	//var alreadyDoneFiles int
	
	//keys := make([]string, 0, len(pagesWithFiles))
    /*for k := range pagesWithFiles {


        keys = append(keys, k)    
    }


*/
	for k, v := range pagesWithFiles {
		//go func() { 
			fmt.Println("key", k , "value", v )
		//}()
	}
	//(countOfWorkers <= WORKERS) && (alreadyDoneFiles < len(pagesWithFiles)){
//		currentLink := 
	
//	}
	
}

func DownloadFile(filepath string, url string) error {

    // Get the data
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // Create the file
    out, err := os.Create(filepath)
    if err != nil {
        return err
    }
    defer out.Close()

    // Write the body to file
    _, err = io.Copy(out, resp.Body)
    return err
}

func initial() {
	flag.StringVar(&searchCombination,"sc", "zelda", "страница с музакальным альбомом")	
	flag.IntVar(&WORKERS, "w", WORKERS, "количество потоков")

	pagesWithFiles = make(map[string]string)
}

func main() {
	
	initial()

	flag.Parse()		

	url := "https://downloads.khinsider.com/game-soundtracks/album/adk-world"

	parseUrl(url)		
	
	downloadFilesCNTRL()
}

