/*
 * @Descripttion:
 * @version:
 * @Author: 1314mylove
 * @Date: 2022-11-10 16:23:17
 * @LastEditors: 1314mylove
 * @LastEditTime: 2022-11-14 17:12:08
 */
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/thinkeridea/go-extend/exnet"
)

func main() {

	// 定义几个变量，用于接收命令行的参数值
	var ipstr string
	var num uint
	var filei string
	var filen string
	// &user 就是接收命令行中输入 -u 后面的参数值，其他同理
	flag.StringVar(&ipstr, "i", "", "需要转换的IP")
	flag.UintVar(&num, "n", 8888, "需要转换的十进制")
	flag.StringVar(&filei, "fi", "", "从文件读取需要转换的IP")
	flag.StringVar(&filen, "fn", "", "从文件读取需要转换的十进制")
	// 解析命令行参数写入注册的flag里
	flag.Parse()
	// 输出结果
	if flag.Lookup("i") != nil {
		if ipstr != "" {
			n, _ := exnet.IPString2Long(ipstr)
			fmt.Println(n)
		}
	}
	if flag.Lookup("n") != nil {
		if num != 8888 {
			s, _ := exnet.Long2IPString(num)
			fmt.Println(s)
		}
	}
	// 从文件读取需要转换的IP
	if flag.Lookup("fi") != nil {
		if filei != "" {
			f, err := os.Open(filei)
			if err != nil {
				fmt.Printf("error : %s", err)
			}

			fread := bufio.NewReader(f)
			for {
				line, _, err := fread.ReadLine()
				if err == io.EOF {
					break
				}
				fmt.Println(string(line))
				// 删除字符串前后空格
				n, _ := exnet.IPString2Long(strings.TrimSpace(string(line)))
				fmt.Println(n)
			}
		}
	}
	// 从文件读取需要转换的十进制
	if flag.Lookup("fn") != nil {
		if filen != "" {
			f, err := os.Open(filen)
			if err != nil {
				fmt.Printf("error : %s", err)
			}

			fread := bufio.NewReader(f)
			for {
				line, _, err := fread.ReadLine()
				if err == io.EOF {
					break
				}
				// fmt.Println("原始：", string(line))
				// 删除字符串前后空格
				intNum, _ := strconv.Atoi(strings.TrimSpace(string(line)))
				int64Num := uint(intNum)
				s, _ := exnet.Long2IPString(int64Num)
				// fmt.Println("转换后：", s)
				fmt.Println(s)
			}

		}
	}

}
