package main

import (
	"bufio"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/speps/go-hashids/v2"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)
import splitCsv "github.com/tolik505/split-csv"

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
	path = getFilePath()
	f, err := os.Open(path + ".csv")
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
	splitManipulatedFile()
}

func splitManipulatedFile() {
	fmt.Println("Şifreleme bitti.")
	splitter := splitCsv.New()
	splitter.FileChunkSize, _ = strconv.Atoi(setFileByte()) //in bytes (100MB)
	splittedPath = setFilePathForSplittedCSVS()
	_, _ = splitter.Split(path+"_manipulated"+".csv", splittedPath)
	fmt.Println("Dosya parçalanıyor.")
	fmt.Println("Dosya parçalandı")
	cmd := exec.Command("/bin/sh", "/Users/talhasalt/GolandProjects/awesomeProject/script.sh", splittedPath, setPasswordForSplittedFiles(), setNameForSecuredFiles())

	pipe, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {

	}
	reader := bufio.NewReader(pipe)
	line, err := reader.ReadString('\n')
	for err == nil {
		fmt.Println(line)
		line, err = reader.ReadString('\n')
	}
}

func splitNonEncryptedFiles() {
	path = getFilePath()
	splitter := splitCsv.New()
	splitter.FileChunkSize, _ = strconv.Atoi(setFileByte()) //in bytes (100MB)
	splittedPath = setFilePathForSplittedCSVS()
	_, _ = splitter.Split(path+".csv", splittedPath)
	fmt.Println("Dosya parçalanıyor.")
	fmt.Println("Dosya parçalandı")
	cmd := exec.Command("/bin/sh", "/Users/talhasalt/GolandProjects/awesomeProject/script.sh", splittedPath, setPasswordForSplittedFiles(), setNameForSecuredFiles())

	pipe, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {

	}
	reader := bufio.NewReader(pipe)
	line, err := reader.ReadString('\n')
	for err == nil {
		fmt.Println(line)
		line, err = reader.ReadString('\n')
	}
}

func main() {
	var image = "\n ██████ ███████ ██    ██     ███████ ██████  ██      ██ ████████ ████████ ███████ ██████  \n██      ██      ██    ██     ██      ██   ██ ██      ██    ██       ██    ██      ██   ██ \n██      ███████ ██    ██     ███████ ██████  ██      ██    ██       ██    █████   ██████  \n██           ██  ██  ██           ██ ██      ██      ██    ██       ██    ██      ██   ██ \n ██████ ███████   ████       ███████ ██      ███████ ██    ██       ██    ███████ ██   ██ \n                                                                                          \n                                                                                          \n"
	println(image)
	if isEncrypted() == "y" {
		readFile()
	} else {
		splitNonEncryptedFiles()
	}

}

func isEncrypted() string {
	prompt := promptui.Prompt{
		Label:     "Sütunlaradan herhangi birini şifrelemek istiyor musun?",
		IsConfirm: true,
	}
	result, _ := prompt.Run()
	if result != "y" && result != "n" {
		isEncrypted()
	}
	return result
}

func getFilePath() string {
	enterFilePath := promptui.Prompt{
		Label: "Parçalanacak dosyanın dosya konumunu gir (Sonuna '.csv' yazma)",
	}
	enteredFilePath, _ := enterFilePath.Run()
	return enteredFilePath
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

func setFilePathForSplittedCSVS() string {
	enterFilePath := promptui.Prompt{
		Label: "Parçalanmış dosyaların dosya konumunu giriniz",
	}
	enteredFilePath, _ := enterFilePath.Run()
	return enteredFilePath
}

func setPasswordForSplittedFiles() string {
	enterFilePath := promptui.Prompt{
		Label: "Parçalanmış dosyalar için şifreyi giriniz",
	}
	enteredFilePath, _ := enterFilePath.Run()
	return enteredFilePath
}

func setFileByte() string {
	setFileByte := promptui.Prompt{
		Label: "Parça boyutunu giriniz",
	}
	settedFileByte, _ := setFileByte.Run()
	return settedFileByte
}

func setNameForSecuredFiles() string {
	enterFilePath := promptui.Prompt{
		Label: "Şifrelenmiş dosyalar için isim gir",
	}
	enteredFilePath, _ := enterFilePath.Run()
	return enteredFilePath
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}
