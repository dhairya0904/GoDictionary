package dict

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
)

type Dictionary struct {
}

func (d *Dictionary) GetMeaning(word string) Orange {
	return _getMeaning(word)
}

func (d *Dictionary) GetMeanings(words []string) []Orange {

	maxGoRoutines := 10
	guard := make(chan struct{}, maxGoRoutines)
	result := make(chan Orange, len(words))

	var wg sync.WaitGroup
	wg.Add(len(words))

	for i := 0; i < len(words); i++ {
		guard <- struct{}{}
		go func(n int) {
			result <- _getMeaning(words[n])
			wg.Done()
			<-guard
		}(i)
	}


	wg.Wait()
	close(result)
	close(guard)

	var oranges []Orange

	for x := range result {
		oranges = append(oranges, x)
	}

	return oranges
}

func _getMeaning(word string) Orange {
	url := "https://mashape-community-urban-dictionary.p.rapidapi.com/define?term=" + word
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "mashape-community-urban-dictionary.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "<Your Key>")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	temp := string(body)

	var result map[string]interface{}
	json.Unmarshal([]byte(temp), &result)

	birds := result["list"].([]interface{})

	var apples []Apple

	for _, value := range birds {
		// Each value is an interface{} type, that is type asserted as a string
		oneApple := Apple{
			definition: value.(map[string]interface{})["definition"].(string),
			example:    value.(map[string]interface{})["example"].(string),
		}
		apples = append(apples, oneApple)
	}

	return Orange{
		Word: word,
		info: apples,
	}
}

type Orange struct {
	Word string
	info []Apple
}

type Apple struct {
	definition string
	example    string
}
