# Zog

重新设计一版 logger

## 概念

有两种 log 经常被混淆：一种是给人读，多用于开发调试，一种是给机器读，一般输出成 JSON 给其他脚本去分析。我想做的是前者。

log 又分 level，如常见的会分成 `DEBUG`/`INFO`/`ERROR`/`WARN`/`FATAL`，但是一个项目其实是很多功能分别打 log，我希望是能容易开关其中的不同功能而不是不同层级的 log。  
如一个需求是，A 和 B 两个功能的 `INFO` 都输出到 `info.log`，我并不是要开关 `INFO` 而是可能直接关掉 A 的所有 log。以及另外 C 功能我希望输出到 `info-c.log` 而不跟 A B 在一个文件。以及不同 level 我希望输出到不同的文件，可以通过 `tail -f` 不同的文件组合来跟踪。

现确认下常见的 SetLogger

* https://pkg.go.dev/google.golang.org/grpc/grpclog#SetLoggerV2
* https://pkg.go.dev/github.com/go-sql-driver/mysql#SetLogger

输出方式

* stdout
* file
* file rotation
* any io.Writer
