package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Structura map[string]Hero

const jname = "heroes.json"

type Hero struct {
	ManaType   string
	MeleeRange string
	Lane       string
}

func main() {
	var s []Structura
	s = Unm(jname)
	s = append(s, Scam())
	byted, _ := json.Marshal(s)
	os.WriteFile(jname, byted, 0666)
}
func Scam() Structura {
	bruh := make(Structura)
	var (
		name    string
		element Hero
	)
	fmt.Println("enter Name, ManaType (mana, energy, none), range type (melee, ranged), lane (bot, top, mid, jungle)")
	fmt.Scanf("%s %s %s %s", &name, &element.ManaType, &element.MeleeRange, &element.Lane)
	bruh[name] = element
	return bruh
}
func Unm(filename string) (umed []Structura) {
	byted, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(byted, &umed)
	return umed
}
