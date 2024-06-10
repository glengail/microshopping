package csv_helper

import (
	"encoding/csv"
	"log"
	"mime/multipart"
)

func ReadCsv(fileHeader *multipart.FileHeader) ([][]string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	//关闭流
	defer func(f multipart.File) {
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}(file)
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var result [][]string
	//第一行是标题行
	for _, line := range lines[1:] {
		data := []string{line[0], line[1]}
		result = append(result, data)
	}
	return result, nil
}
