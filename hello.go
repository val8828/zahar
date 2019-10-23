package main

	

import (
   	"fmt"
	"github.com/PuerkitoBio/goquery"
//	"os"
	"sync"
    "flag"    
  //  "net/http"    
)

var (
	searchCombination string // 
)


func _check(err error) {
	if err != nil {
		panic(err)
	}
}

// основная функция обработки
func parseUrl(url string) {
	fmt.Println("request: " + url)

	// заворачиваем источник в goquery документ
	doc, err := goquery.NewDocument(url)
	// получаем объект со всеми тегами "a"
	
	_check(err)

	// в манере jquery, css селектором получаем все ссылки
	doc.Find(".EchoTopic").Each(func(i int, s *goquery.Selection) {
		if val, ok := s.Attr("href"); ok {
                fmt.Println(val)
        }
	})
	
}

func main() {
	initCommandLineProp()
	flag.Parse()	
	
	var wg sync.WaitGroup

	url := "https://downloads.khinsider.com/search?search=zelda"

	// получаем список url из входных параметров
	//for _, url := range os.Args[1:] {
		// каждый выполним параллельно
		wg.Add(1)

		// закрываем в анонимной функции переменную из цикла,
		// что бы предотвартить её потерю во время обработки
		go func(url string) {
			defer wg.Done()
			parseUrl(url)
		}(url)
	//}

	// ждем завершения процессов
	wg.Wait()
}

func initCommandLineProp(){
	flag.StringVar(&searchCombination,"sc", "zelda", "name of the game to search")
	//flag.Parse()			
}
