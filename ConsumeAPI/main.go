package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Response struct{
	Name string `json:"name"`
	PokemonType [] PokemonType `json:"pokemon_list"`
}

type PokemonType struct{
	Species PokemonSpecies `json:"species"`
	EntryNum int `json:"entry_num"`
}
type PokemonSpecies struct {
	Name string `json:"name"`
}
func main(){
	url:= "http://pokeapi.co/api/v2/pokedex/kanto/"
	response,err:= http.Get(url)
	if err!=nil{
		fmt.Printf("Unable to get the pokemon type :%s\n",err.Error())	
	}
	data,err:= ioutil.ReadAll(response.Body)
	if err!=nil{
		fmt.Println(err.Error())	
	}
response.Body.Close()
	//fmt.Println(string(data))

var UserData Response

json.Unmarshal(data,&UserData)
fmt.Println(UserData.Name)

	fmt.Println(len(UserData.PokemonType))
	for i>=0;i<len(UserData.PokemonType);i++{
		fmt.Println("%s-%s\n",UserData.PokemonType[i].EntryNum,UserData.PokemonType[i].Species.Name)
	}
}

