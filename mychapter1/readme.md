# 运行单元测试

```bash
$ go test -v
```

# 运行性能测试

```bash
# 运行全部性能测试
$ go test -v -bench=".*"
 
# 运行以_GetStudentIdsUsing为前缀的性能测试
$ go test -v -bench="_GetStudentIdsUsing.*" 

=== RUN   Test_CreateStudentList
--- PASS: Test_CreateStudentList (0.00s)
=== RUN   Test_GetStudentIdsUsingApply
--- PASS: Test_GetStudentIdsUsingApply (0.00s)
=== RUN   Test_GetStudentIdsUsingLoop
--- PASS: Test_GetStudentIdsUsingLoop (0.00s)
PASS
Benchmark_GetStudentIdsUsingApply-8      3000000               433 ns/op
Benchmark_GetStudentIdsUsingLoop-8      100000000               97.2 ns/op
ok      gobook-src/mychapter1   81.996s
```

# 得到性能测试报告

```bash
# 生成性能测试报告
$ go test -v -bench=".*"  -cpuprofile=cpu.prof

# 使用go tool pprof工具查看性能测试报告
$ go tool pprof mychapter1.test cpu.prof
Entering interactive mode (type "help" for commands)
(pprof) web
```

# 相关知识

- [Golang package testing](https://golang.org/pkg/testing/)
