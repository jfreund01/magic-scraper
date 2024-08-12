package main

import (
  "encoding/json";
  "fmt";
  "github.com/gocolly/colly";
)

const EDHREC_BASE_URL string = "https://edhrec.com/"
var EDHREC_URLS = []string{ 
  "top/", 
  "top/week", 
  "top/month",
  "/commanders",
  "/commanders/week",
  "/commanders/month", 
} 


type Card struct {
  name string
  set string
  foil bool
  avg_price float64
}

func main() {
  c := colly.NewCollector()
  
  
  
  c.OnHTML("script[type='application/json']", func(e *colly.HTMLElement) {
    var jsonData map[string]interface{}

    err := json.Unmarshal([]byte(e.Text), &jsonData)
    
    if err != nil {
      fmt.Println("Error parsing JSON")
      return
    }
    
  
    cardlists := jsonData["props"].
    (map[string]interface{})["pageProps"].
    (map[string]interface{})["data"].
    (map[string]interface{})["container"].
    (map[string]interface{})["json_dict"].
    (map[string]interface{})["cardlists"].
    ([]interface{})[0].
    (map[string]interface{})["cardviews"].
    ([]interface{})
    
    cards := make([]string, len(cardlists))

    for _, card := range cardlists {
      card := card.(map[string]interface{})
      cards = append(cards, card["sanitized_wo"].(string))
    }

    fmt.Println(cards)
    fmt.Println(len(cards))   
    
  })

  c.OnRequest(func(r *colly.Request) {
    fmt.Println("Visiting : ", r.URL.String())  
  })
 
  for _, url := range EDHREC_URLS {
    full_url := EDHREC_BASE_URL + url
    c.Visit(full_url)
  }
}

