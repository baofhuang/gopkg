package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main()  {
	//准备读取文件
	fileName := "./testdata/acs_role_member2.csv"
	fs, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	defer fs.Close()
	r := csv.NewReader(fs)
	var data []string
	//针对大文件，一行一行的读取文件
	for {
		tmp, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		data=append(data,tmp[2])
	}
	//取出roleid
	var RoleIDs []string
	for _,v:=range data{
		RoleIDs = append(RoleIDs,v)
	}
	//role_id去重
	var RoleIds []string
	m := make(map[string]bool)
	for _, v := range RoleIDs {
		if _, ok := m[v]; !ok {
			RoleIds = append(RoleIds, v)
			m[v] = true
		}
	}
	fmt.Println("total num:%+v",len(RoleIds))
	//输出ROlE_ID
	out,e3:=os.Create("./testdata/role_id.txt")
	if e3!=nil{
		log.Fatal(e3)
	}
	defer out.Close()
	bufWriter:=bufio.NewWriter(out)
	for i:=0;i<len(RoleIds);i++{
		bufWriter.WriteString(RoleIds[i])
		bufWriter.WriteString("\n")
	}
	bufWriter.Flush()
}
