package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"./weatherType"
)

func main() {
	var cityName string
	var response weatherType.Response
	flag.StringVar(&cityName, "c", "", "")
	flag.Parse()

	const key string = "4f3de5836fc8dcd56f17ecc1d0c09218"
	var body, err = request("http://api.openweathermap.org/data/2.5/weather?q=" + cityName + "&units=metric&appid=" + key)

	if err != nil {
		print(err)
	}

	err = json.Unmarshal([]byte(body), &response)

	if err != nil {
		fmt.Println(err)
	}
	//cod, err := strconv.Atoi(response.Cod)
	if response.Cod != 200 {
		fmt.Println("API error :", response.Cod)

	}

	affiche(response)

}

func request(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Print(err)
		return "connection error", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	return string(body), err
}

func affiche(response weatherType.Response) {
	fmt.Println("Ville : ", response.Name)
	fmt.Println("Température(°C) : ", response.Main.Temp)
	fmt.Println("Température min(°C) : ", response.Main.Temp_min)
	fmt.Println("Température max(°C) : ", response.Main.Temp_max)
	fmt.Println("Pression : ", response.Main.Pressure)
	fmt.Println("Humidité : ", response.Main.Humidity)
	fmt.Println("Vitesse du vent(m/s) : ", response.Wind.Speed)
}
