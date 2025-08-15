package main

import "os"

func main() {
	// 创建文件夹
	_ = os.Mkdir("ql", 0777)

	// 创建多层级目录
	_ = os.MkdirAll("ql\\test1\\test2", 0700)

	// 删除目录，以及目录里面的内容
	_ = os.RemoveAll("ql")

	// 创建文件，go 语言创建文件后会打开文件，必须要关闭，否则其余人不可以用改文件
	file := createName("ql.txt")

	// 写入内容
	writeFile(file, "ql\r\nlll\r\n")

	// 删除文件
	//_ = os.RemoveAll("ql.txt")
}

func writeFile(f *os.File, s string) {
	writeString, _ := f.WriteString(s)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			err.Error()
		}
	}(f)
	println(writeString)
}

func createName(name string) *os.File {
	file, err := os.Create(name)
	if err != nil {
		err.Error()
	} else {
		return file
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			err.Error()
		}
	}(file)
	return file
}
