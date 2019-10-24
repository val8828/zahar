package main

	

import (
   	"fmt"
	"github.com/PuerkitoBio/goquery"
//	"os"
	"sync"
    "flag"   
    "strings"
  
)

var (
	searchCombination string // 
	m map[string]string
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
			if strings.HasSuffix(val, "mp3") {
				if _,ok :=  m[val]; !ok {
					m[val] = s.Text()					
				}
				
			}
        }
	})		
	
}

func main() {
	m = make(map[string]string)
	initCommandLineProp()
	flag.Parse()	
	
	var wg sync.WaitGroup

	url := "https://downloads.khinsider.com/game-soundtracks/album/adk-world"

		// каждый выполним параллельно
		wg.Add(1)

		// закрываем в анонимной функции переменную из цикла,
		// что бы предотвартить её потерю во время обработки
		go func(url string) {
			defer wg.Done()
			parseUrl(url)
		}(url)
	
	// ждем завершения процессов
	wg.Wait()
	for key, value := range m {
    	fmt.Println("Key:", key, "Value:", value)
	}
}

func initCommandLineProp(){
	flag.StringVar(&searchCombination,"sc", "zelda", "name of the game to search")	
}
