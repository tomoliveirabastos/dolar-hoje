package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/energye/systray"
	"github.com/energye/systray/icon"
)

type Cotacao struct {
	USDBRL USDBRL `json:"USDBRL"`
}

type USDBRL struct {
	Code       string    `json:"code"`
	Codein     string    `json:"codein"`
	Name       string    `json:"name"`
	High       string    `json:"high"`
	Low        string    `json:"low"`
	VarBid     string    `json:"varBid"`
	PctChange  string    `json:"pctChange"`
	Bid        string    `json:"bid"`
	Ask        string    `json:"ask"`
	Timestamp  string    `json:"timestamp"`
	CreateDate time.Time `json:"create_date"`
}

func main() {
	systray.Run(onReady, onExit)
}

func CheckDolar() Cotacao {
	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	req, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	res, err := io.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}

	var usd Cotacao

	json.Unmarshal(res, &usd)

	return usd
}

func onReady() {
	dolar := CheckDolar()
	title := fmt.Sprintf("Dolar - Alta: %s, Baixa %s", dolar.USDBRL.High, dolar.USDBRL.Low)
	systray.SetIcon(icon.Data)
	systray.SetTitle(title)
	systray.SetTooltip("")

	systray.SetOnClick(func(menu systray.IMenu) {
		menu.ShowMenu()
	})

	systray.SetOnDClick(func(menu systray.IMenu) {
		menu.ShowMenu()
	})

	systray.SetOnRClick(func(menu systray.IMenu) {
		menu.ShowMenu()
	})

	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	mQuit.SetIcon(icon.Data)
	mQuit.Click(systray.Quit)
}

func onExit() {
	fmt.Println("Saiu")
}
