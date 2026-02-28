package main

import(
	"fmt"
	"time"
	"flag"
	"strings"
    "strconv"
	"text/tabwriter"
	"os"
)

func main(){

	layout := "Mon,02 Jan 2006 15:04:05"

	listPtr := flag.Bool("list",false,"show the tasks")
	addPtr := flag.String("add","default","add new task")
	completePtr := flag.Int("complete",0,"compelete task")
	editPtr := flag.String("edit","default","edit task")

	flag.Parse()

	if *listPtr {
		tasks, err := Read()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		fmt.Fprintln(w, "ID\tTITLE\tSTATUS\tCreated AT\t\t\t\t\t\tCOMPLETED AT")
		fmt.Fprintln(w, "--\t-----\t------\t----------\t\t\t\t\t\t------------")

		for _, t := range tasks {
			status := "❌"
			if t.Status == "done" {
				status = "✅"
			}
			if !t.CompletedAt.IsZero(){
				fmt.Fprintf(w,"%d\t%s\t%s\t%s\t%s\n", t.Number,t.Title, status, t.CreatedAt.Format(layout), t.CompletedAt.Format(layout))
			}else {
				fmt.Fprintf(w,"%d\t%s\t%s\t%s\n", t.Number,t.Title, status, t.CreatedAt.Format(layout))
			}
		}
		w.Flush()
	}

	if *addPtr !="default" {
		tk := Todo{
		Number: 0,
		Title: *addPtr,
		Status: "Pending",
		CreatedAt: time.Now(),
		}

		err := Write(tk)
		if err != nil {
			fmt.Printf("there is error: %v",err)
		}

		fmt.Println("Task Added")
	}

	if *completePtr != 0 {
		err := Complete(*completePtr)

		if err != nil {
			fmt.Printf("there is error: %v",err)
			return
		}
		fmt.Println("Complete done")
	}

	if *editPtr != "default" {
		strValue := *editPtr
		parts := strings.Split(strValue, ":") // Returns ["2", "ddssd"]

		num, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: First part is not a number")
		}

		editerror := Edit(num,parts[1])
		if editerror != nil {
			fmt.Printf("there is error: %v",editerror)
			return
		}

		fmt.Println("Edit Done")
	}

}