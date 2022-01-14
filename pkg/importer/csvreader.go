package importer

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/fusidic/trace-importer/pkg/microservices"
	"github.com/fusidic/trace-importer/pkg/nodes"
	"github.com/gocarina/gocsv"
)

type reader interface {
	readNode() (chan nodes.Node, error)
	readMS() (chan microservices.MicroService, error)
}

type csvReader struct {
	filePath string
}

func newCsvReader(filePath string) *csvReader {
	return &csvReader{
		filePath: filePath,
	}
}

func (c csvReader) readNode() (chan nodes.Node, error) {
	fileHanlde, err := os.OpenFile(c.filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer fileHanlde.Close()

	ch := make(chan nodes.Node)

	go func() {
		err = gocsv.UnmarshalToChan(fileHanlde, ch)
		if err != nil {
			log.Fatal(err)
		}
	}()

	return ch, nil
}

// TODO: see if could be simplified with generics.
func (c csvReader) readMS() (chan microservices.MicroService, error) {
	fileHanlde, err := os.OpenFile(c.filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer fileHanlde.Close()

	ch := make(chan microservices.MicroService)

	go func() {
		err = gocsv.UnmarshalToChan(fileHanlde, ch)
		if err != nil {
			log.Fatal(err)
		}
	}()

	return ch, nil
}

func (c csvReader) Read() [][]string {
	//打开文件(只读模式)，创建io.read接口实例
	f, err := os.Open(c.filePath)

	if err != nil {
		log.Println("csv文件打开失败!")
	}
	defer f.Close()

	//创建csv读取接口实例
	reader := csv.NewReader(f)

	//获取一行内容，一般为第一行内容
	// read, _ := reader.Read() //返回切片类型：[chen  hai wei]
	// log.Println(read)

	//读取所有内容
	readAll, err := reader.ReadAll() //返回切片类型：[[s s ds] [a a a]]
	if err != nil {
		log.Fatalln("Error in reading csv file", err)
	}
	return readAll
}
