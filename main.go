package main

import (
	"flag"
	"os"
	"fmt"
	"encoding/csv"
	"strings"
)

type problem struct {
	question string
	answer string
}
func main(){
	csvFilename := flag.String("csv","problems.csv","a csv file in the format of problem , solution")
	flag.Parse()
 	file , err := os.Open(*csvFilename)
 	if err!=nil{
 	exit(fmt.Sprintf("Failed to open csv file :%s",*csvFilename))
	}
	r:=csv.NewReader(file)
	lines , err := r.ReadAll()
	if err != nil{
		exit("failed to parse the provided csv file")
	}
	counter := 0
	problems := parseLines(lines)
	for i , p:=range problems{
		fmt.Printf("Problem #%d : %s =\n",i+1,p.question)
		var ans string
		fmt.Scanf("%s\n",&ans)
		if ans == p.answer{
			counter++
		}

	}
	fmt.Printf(" you scored %d out of %d .\n",counter,len(problems))

}
func  parseLines(lines [][]string) []problem  {
	ret := make([]problem,len(lines))
	for i , lines :=range lines{
		ret[i] = problem{
		question: lines[0],
		answer: strings.TrimSpace(lines[1]),
		}}
		return ret
	}


func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}