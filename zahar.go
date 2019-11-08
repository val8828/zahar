package main

/*
zahar will download some mp3 files from https://downloads.khinsider.com

just give zahar link to page with list of tracks 

for example "zahar.exe -url=https://downloads.khinsider.com/game-soundtracks/album/death-brade"

to change default count of download thread add the option for example "-w=123" - 123 goroutine will be use
*/	

import (
   	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	"time"
    "flag"   
    "strings"
  	"io"
  	"net/http"
	"runtime"
	"path/filepath"
)

var (
	searchCombination string // 
	pagesWithFiles map[string] string
	albumName string
	WORKERS int = 20 //кол-во "потоков"
	url string
)


func _check(doc *goquery.Document, err error) int {
	if err != nil {
		panic(err)
	}
	if strings.Contains(doc.Find("title").Text(), "Error") {		
		return 404		
	}
	return 0
}

// основная функция обработки
func parseUrl(url string, level int )  {	
	// заворачиваем источник в goquery документ
	doc, err := goquery.NewDocument(url)

	switch level{
	case 1:
		if _check(doc, err)	== 0 {
			//get album name from first <h2> tag
			if(albumName == ""){
				doc.Find("h2").EachWithBreak(func(i int, s *goquery.Selection) bool {		
					albumName =  s.Text()
					return false
				}) 
			}			
				doc.Find("a").Each(func(i int, s *goquery.Selection) {
					if val, ok := s.Attr("href"); ok {
						//Check if url ends with mp3
						if strings.HasSuffix(val, "mp3") {		
							val := "https://downloads.khinsider.com" + val						
								parseUrl(val, 2)
						}
		    		}
				})			
		}

	case 2:				
		// в манере jquery, css селектором получаем все ссылки
		doc.Find("a").Each(func(i int, s *goquery.Selection) {			
			if val, ok := s.Attr("href"); ok {				
				//Check if url ends with mp3
				if strings.HasSuffix(val, "mp3") {								
					// put all url in map (only first in)
					if _,ok :=  pagesWithFiles[val]; !ok {						
						pagesWithFiles[val] = findSongName(doc)					
					}
				}
    		}
		})
	}	
}

func findSongName(doc *goquery.Document) string {
	var songName string
	doc.Find("b").Each(func(i int, s *goquery.Selection) {			
			if i == 3 {			
				songName = s.Text()
			}
		})
	return songName
}

func downloadFilesCNTRL() {	
	os.MkdirAll("downloads\\"+albumName,0777)
	for k, v := range pagesWithFiles {
		
		if runtime.NumGoroutine() >= WORKERS {
			time.Sleep(1 * time.Millisecond)
		}
		
		go func(i string, j string) { 			
			if err := DownloadFile(filepath.FromSlash("downloads"+"/"+albumName + "/" + j + ".mp3"), i); err != nil {
        		panic(err)
        		fmt.Println(err)
    		}
		}(k , v)

	}
		time.Sleep(20 *time.Second)		    	
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
	flag.StringVar(&url,"url", "https://downloads.khinsider.com/game-soundtracks/album/death-brade", "страница с музакальным альбомом")	
	flag.IntVar(&WORKERS, "w", 20, "количество потоков")
	flag.StringVar(&albumName,"an", "", "путь папки")	

	pagesWithFiles = make(map[string]string)
}

func main() {
	
	initial()

	flag.Parse()			

	parseUrl(url,1)		
	
	downloadFilesCNTRL()
}

