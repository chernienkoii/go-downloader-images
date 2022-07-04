package main

import (
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"unicode"
)

func main() {

	fileNamexlsx := "./images"

	xlsx, err := excelize.OpenFile(fileNamexlsx + ".xlsx")
	if err != nil {
		//		fmt.Println(err)
		fmt.Println(" error open file xlsx")
	} else {
		fmt.Println("open file xlsx OK ")
	}

	sheet := "images"
	n := 9999999
	for i := 1; i < n; i += 1 {
		name := 0
		var titleprinted bool = false
		for r := 'b'; r <= 'k'; r++ {
			R := unicode.ToUpper(r)
			cell := xlsx.GetCellValue(sheet, fmt.Sprintf("%c%d", R, i))
			title := xlsx.GetCellValue(sheet, fmt.Sprintf("A%d", i))
			// fmt.Println(title + " ---- title")
			// fmt.Println(cell + " ---- cell")
			name += 1
			if len(cell) > 0 {
				if titleprinted == false {
					titleprinted = true

					if _, err := os.Stat(title + "_frgroup"); errors.Is(err, os.ErrNotExist) {
						err := os.Mkdir(title+"_frgroup", os.ModePerm)
						if err != nil {
							log.Println(err)
						}
					}

				}
				//name := len(cell)
				//fmt.Println(name)
				fileName := title + "_frgroup/" + strconv.Itoa(name) + ".jpg"
				// fmt.Println(fileName + "----fileName")
				URL := cell
				//				fmt.Println(cell)
				err := downloadFile(URL, fileName)
				if err != nil {
					log.Fatal(err)
					continue
				}
				fmt.Println("File downlaod in current working directory", fileName)

			}
		}

	}
	fmt.Println("Script OK")
}
func downloadFile(URL, fileName string) error {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Received non 200 response code")
	}
	//Create a empty file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
