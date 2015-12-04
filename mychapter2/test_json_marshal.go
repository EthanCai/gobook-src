package main

import (
	"fmt" // 我们需要fmt包中的Println函数
	"encoding/json"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Runtime error caught: %v", r)
		}
	}()

	original_text := `[{"text": ["棕木门口，内有大量铜和木布置，造型奇特的桌椅，书架还以花色雕刻的茶色玻璃点缀，不浮华，也没有故弄玄虚的高雅，却又散发着匠心独特低调时尚气质。\n\"\b\r\\\/\t\uFFFF"], "imageIds": ["7520"], "type": 1}]`

	pagesNet := []FabulaPageNet{}

	Unmarshal(original_text, &pagesNet)

	text := Marshal(pagesNet)

	fmt.Println(text)

}


type FabulaPageNet struct {
	Type     int         `json:"type,omitempty"`
	ImageIds []string    `json:"imageIds,omitempty"`
	Text     []string    `json:"text,omitempty"`
}

func Unmarshal(str string, data interface{}) {

	if len(str) == 0 {
		return
	}

	err := json.Unmarshal([]byte(str), data)

	if err != nil {
		panic(err)
	}
}

func Marshal(data interface{}) string {
	bytes, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	return string(bytes)
}
