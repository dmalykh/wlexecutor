
This is dead simple demonstration how [wlinterpreter](https://github.com/dmalykh/wlinterpreter) could be used. 

This code [Go 1.18.1 needed](https://github.com/golang/go/issues/51847).

```bash
Usage of /private/var/folders/_c/7f6xftgs7qzbtcmhkstkfg680000gq/T/GoLand/___go_build_main_go:
  -cell int
        Size of cell: 8, 32 or 64
  -help
        Displays this help
  -stack int
        Size of stack (default 30000)

```
For example:
```bash
$ echo "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++." | wlexecutor -cell 32
Hello World!
```