package main

import (
	"os"
	//"fmt"
	//"io/ioutil"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/red-two/lsbeat/beater"
)

func main() {
		err := beat.Run("lsbeat", "", beater.New)
		if err != nil {
			os.Exit(1)
		}
		// //apply run path "." without argument.
    // if len(os.Args) == 1 {
    //     listDir(".")
    // } else {
    //     listDir(os.Args[1])
    // }
}

// func listDir(dirFile string) {
//     files, _ := ioutil.ReadDir(dirFile)
//     for _, f := range files {
//         t := f.ModTime()
//         fmt.Println(f.Name(), dirFile+"/"+f.Name(), f.IsDir(), t, f.Size())
//         if f.IsDir() {
//             listDir(dirFile + "/" + f.Name())
//         }
//     }
// }
