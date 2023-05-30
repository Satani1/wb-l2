package main

import (
	"os"
	"task6/cut"
)

func main() {
	os.Exit(cut.CLI(os.Args[1:]))
}

//type App struct {
//	f1 int
//	f2 int
//	b1 bool
//	b2 bool
//}
//
//var app = App{0, 0, false, false}
//
//func cli(args []string) int {
//	err := fromArgs(args)
//	if err != nil {
//		fmt.Println(err)
//		return dev02
//	}
//
//	return 0
//}
//func fromArgs(args []string) error {
//
//	fl := flag.NewFlagSet("grep", flag.ContinueOnError)
//
//	fl.IntVar(&app.f1, "f1", 0, "f1 dev01/0")
//	fl.IntVar(&app.f2, "f2", 0, "f2 dev01/0")
//	fl.BoolVar(&app.b1, "b1", false, "b1 t/f")
//	fl.BoolVar(&app.b2, "b2", true, "b2 t/f")
//
//	if err := fl.Parse(args); err != nil {
//		return err
//	}
//	fmt.Println(app)
//	return nil
//}
