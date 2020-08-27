package main

import (
	"dictionary/dictionary"
	"flag"
	"fmt"
	"os"
)

func main() {
	action := flag.String("action", "list", "Action to perform on the dictionary")
	d, err := dictionary.New("./badger")
	handlerErr(err)
	defer d.Close()

	flag.Parse()
	switch *action {
	case "list":
		actionList(d)
	case "add":
		actionAdd(d, flag.Args())
	case "define":
		actionGet(d, flag.Args())
	case "remove":
		actionRemove(d, flag.Args())
	default:
		fmt.Printf("Unknown action: %v\n", *action)
	}
}

func actionList(d *dictionary.Dictionary)  {
	words, entries, err := d.List()
	handlerErr(err)
	fmt.Println("Dictionary content")

	for _, word := range words {
		fmt.Println(entries[word])
	}
}

func actionAdd(d *dictionary.Dictionary, args []string)  {
	err := d.Add(args[0], args[1])
	handlerErr(err)
	fmt.Printf("%v added to the dictionary\n", args[0])
}

func actionGet(d *dictionary.Dictionary, args []string)  {
	entry ,err := d.Get(args[0])
	handlerErr(err)
	fmt.Println(entry)
}

func actionRemove(d *dictionary.Dictionary, args []string)  {
	err := d.Remove(args[0])
	handlerErr(err)
	fmt.Printf("%v was successfully removed\n", args[0])
}

func handlerErr(err error)  {
	if err != nil {
		fmt.Printf("Dictionary error: %v\n", err)
		os.Exit(1)
	}
}