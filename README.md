
This is dead simple demonstration how [wlinterpreter](https://github.com/dmalykh/wlinterpreter) could be used. 

This code [Go 1.18.1 needed](https://github.com/golang/go/issues/51847).

```bash
Usage:
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