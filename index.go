package main

import (
	"bufio"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/speps/go-hashids/v2"
	"log"
	"os"
	"strconv"
	"strings"
)

var selectedLineForEncrypting int
var path string
var splittedPath string

func encryption(param int) string {
	hd := hashids.NewData()
	hd.Salt = os.Getenv("HASHID_SALT")
	hd.Alphabet = os.Getenv("HASHID_ALPHABET")
	hd.MinLength, _ = strconv.Atoi(os.Getenv("HASHID_MINLENGTH"))
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{param})
	return e
}

func readFile() {
	path = "/Users/talhasalt/Desktop/csv-splitter/example.csv"
	f, err := os.Open(path)
	fmt.Println("Dosya açıldı")
	f2, err := os.Create(path + "_manipulated" + ".csv")
	fmt.Println("Yeni dosyamız oluşturuldu")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	count := 0
	for scanner.Scan() {
		if count == 0 {
			selectedLineForEncrypting = getWhichLineEncrypted(scanner.Text())
			f2.WriteString(scanner.Text() + "\n")
		} else {
			row := strings.Split(scanner.Text(), ",")
			intTemp, _ := strconv.Atoi(strings.Trim(row[selectedLineForEncrypting], "\""))
			strconv.Atoi(row[selectedLineForEncrypting])
			//fmt.Println(intTemp)
			encrypted := "U-" + encryption(intTemp)
			row[2] = encrypted
			f2.WriteString(strings.Join(row, ",") + "\n")
			fmt.Println(count, ". Satır okundu, işlendi, yazıldı.")
		}
		count++
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}



func main() {
	var image = "\n ██████ ███████ ██    ██     ███████ ██████  ██      ██ ████████ ████████ ███████ ██████  \n██      ██      ██    ██     ██      ██   ██ ██      ██    ██       ██    ██      ██   ██ \n██      ███████ ██    ██     ███████ ██████  ██      ██    ██       ██    █████   ██████  \n██           ██  ██  ██           ██ ██      ██      ██    ██       ██    ██      ██   ██ \n ██████ ███████   ████       ███████ ██      ███████ ██    ██       ██    ███████ ██   ██ \n                                                                                          \n                                                                                          \n"
	println(image)
	if isEncrypted() == "y" {
		readFile()
	} else {
	}
}

func isEncrypted() string {
	prompt := promptui.Prompt{
		Label:     "Dosyayı parçalamadan önce şifrelemek istediğin bir sütun var mı?",
		IsConfirm: true,
	}
	result, _ := prompt.Run()
	if result != "y" && result != "n" {
		isEncrypted()
	}
	return result
}

func getWhichLineEncrypted(line string) int {
	row := strings.Split(line, ",")
	prompt := promptui.Select{
		Label: "Şifrelenmek istenen sütunu seçiniz",
		Items: row,
	}
	_, result, _ := prompt.Run()
	return indexOf(result, row)
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}
