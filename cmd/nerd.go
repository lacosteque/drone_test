package main
import "fmt"
func main() {
    fmt.Println(`
NAME:
   nerd - versatile tool to support smart deployment process

USAGE:
   nerd [global options] command [command options] [arguments...]

COMMANDS:
   apply, a
   login, l
   config, c
   edit, e
   down
   reload, r
   up
   update
   distpacker
   push
   ps          List containers in all found compositions
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)""")
`)
}
