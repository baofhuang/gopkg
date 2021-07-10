package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"strings"
)

var inpath = flag.String("srcpath", "./testdata/in.txt", "please source file path")
var outpath = flag.String("dstpath", "./testdata/out.txt", "please destination file path")

func main() {
	flag.Parse()
	//读取文件
	source, err := os.Open(*inpath)
	if err != nil {
		log.Fatal(err)
	}
	defer source.Close()
	buffReader:=bufio.NewReader(source)
	//处理数据
	var rids []string
	var tmpdata []string
	for {
		tmp,_,err:=buffReader.ReadLine()
		if err != nil || err == io.EOF {
			break
		}
		tmpdata=append(tmpdata, string(tmp))
	}
	for i:=0;i<len(tmpdata);i++{
		start:=strings.LastIndex(tmpdata[i],":\"")+2
		end := strings.LastIndex(tmpdata[i],"\"}")
		rid:=tmpdata[i][start:end]
		rids=append(rids,rid)
	}
	//将处理后的数据写入目标文件
	out,e3:=os.OpenFile(*outpath,os.O_RDWR,771)
	if e3!=nil{
		log.Fatal(e3)
	}
	defer out.Close()
	bufWriter:=bufio.NewWriter(out)
	for i:=0;i<len(rids);i++{
		bufWriter.WriteString(rids[i])
		bufWriter.WriteString("\n")
	}
	bufWriter.Flush()
}
