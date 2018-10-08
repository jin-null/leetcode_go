package main

import (
	"fmt"
)

func longestCommonPrefix(strs [] string) string {
	s:=stortest(strs)
	for i,r:=range s{
		for j:=0;j<len(strs);j++{
			if strs[j][i] != byte(r){
				return strs[j][:i]
			}
		}
	}
	return  ""
}
func stortest(str[] string)string{
	if len(str) == 0 {
		return ""
	}
	res := str[0]
	for _,i:=range str{

		if len(res)>len(i){

			res=i
		}
	}
	return res
}
func main() {
	strs := [] string{"","b"}
	output := longestCommonPrefix(strs)
	fmt.Println(output)
}
