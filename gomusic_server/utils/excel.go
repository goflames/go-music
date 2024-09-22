package utils

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"reflect"
)

// excel工具类

// createExcel 接受一个文件名和通用的 slice 数组，生成 Excel 文件
func CreateExcel(fileName string, data interface{}) error {
	f := excelize.NewFile()
	sheetName := "歌单列表"
	f.NewSheet(sheetName)

	// 获取数据类型，并确保它是一个 slice
	value := reflect.ValueOf(data)
	if value.Kind() != reflect.Slice {
		return fmt.Errorf("data should be a slice")
	}

	// 获取 slice 中的第一个元素类型并写入表头
	if value.Len() > 0 {
		firstElem := value.Index(0)
		WriteHeaders(f, sheetName, firstElem)
	}

	// 写入数据
	for i := 0; i < value.Len(); i++ {
		row := i + 2 // 从第二行开始写入数据
		elem := value.Index(i)
		WriteRow(f, sheetName, elem, row)
	}

	// 保存 Excel 文件
	return f.SaveAs(fileName)
}

// writeHeaders 写入表头
func WriteHeaders(f *excelize.File, sheetName string, elem reflect.Value) {
	// 获取字段的数量
	elemType := elem.Type()
	for i := 0; i < elemType.NumField(); i++ {
		header := elemType.Field(i).Name
		col := string(rune('A' + i)) // 从 A 列开始
		f.SetCellValue(sheetName, fmt.Sprintf("%s1", col), header)
	}
}

// writeRow 写入一行数据
func WriteRow(f *excelize.File, sheetName string, elem reflect.Value, row int) {
	for i := 0; i < elem.NumField(); i++ {
		value := elem.Field(i).Interface()
		col := string(rune('A' + i)) // 从 A 列开始
		f.SetCellValue(sheetName, fmt.Sprintf("%s%d", col, row), value)
	}
}

// 获取文件大小
func GetFileSize(file *os.File) int64 {
	info, err := file.Stat()
	if err != nil {
		return 0
	}
	return info.Size()
}
