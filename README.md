# pdebug
prettier debug.PrintStack() for golang


- Filters stack traces to list function calls only in YOUR package.
- Hide extra information (pointer addreses and machine readable arguments).
- Easier to Click through line by line in your IDE to see the flow of the code.

### Usage

```go
import "github.com/medyagh/pdebug"
pdebug.PrintStack()
```

### Example Output with pdbeug
```
        /Users/medya/workspace/minikube/cmd/minikube/main.go:95 +0x250
        /Users/medya/workspace/minikube/pkg/minikube/out/out.go:487
        /Users/medya/workspace/minikube/pkg/minikube/out/out_reason.go:99
        /Users/medya/workspace/minikube/pkg/minikube/out/out_reason.go:40
        /Users/medya/workspace/minikube/pkg/minikube/exit/exit.go:64
        /Users/medya/workspace/minikube/pkg/minikube/exit/exit.go:92
        /Users/medya/workspace/minikube/cmd/minikube/cmd/start.go:288
        /Users/medya/workspace/minikube/cmd/minikube/cmd/root.go:174

```

### Compare to debug.PrintStack()
```
runtime/debug.Stack()
        /usr/local/go/src/runtime/debug/stack.go:26 +0x64
runtime/debug.PrintStack()
        /usr/local/go/src/runtime/debug/stack.go:18 +0x1c
k8s.io/minikube/pkg/minikube/out.displayGitHubIssueMessage()
        /Users/medya/workspace/minikube/pkg/minikube/out/out.go:486 +0xc8
k8s.io/minikube/pkg/minikube/out.displayText({{0x1062af36d, 0xb}, 0x50, 0x17, {0x0, 0x0}, {0x0, 0x0}, {0x0, 0x0, ...}, ...}, ...)
        /Users/medya/workspace/minikube/pkg/minikube/out/out_reason.go:99 +0x344
k8s.io/minikube/pkg/minikube/out.Error({{0x1062af36d, 0xb}, 0x50, 0x17, {0x0, 0x0}, {0x0, 0x0}, {0x0, 0x0, ...}, ...}, ...)
        /Users/medya/workspace/minikube/pkg/minikube/out/out_reason.go:40 +0x2dc
k8s.io/minikube/pkg/minikube/exit.Message({{0x1062af36d, 0xb}, 0x50, 0x17, {0x0, 0x0}, {0x0, 0x0}, {0x0, 0x0, ...}, ...}, ...)
        /Users/medya/workspace/minikube/pkg/minikube/exit/exit.go:64 +0x354
k8s.io/minikube/pkg/minikube/exit.Error({{0x1062af36d, 0xb}, 0x50, 0x0, {0x0, 0x0}, {0x0, 0x0}, {0x0, 0x0, ...}, ...}, ...)
        /Users/medya/workspace/minikube/pkg/minikube/exit/exit.go:92 +0x240
k8s.io/minikube/cmd/minikube/cmd.runStart(0x109902200, {0x10629f68b?, 0x4?, 0x10629f6df?})
        /Users/medya/workspace/minikube/cmd/minikube/cmd/start.go:288 +0xd78
github.com/spf13/cobra.(*Command).execute(0x109902200, {0x10995cea0, 0x0, 0x0})
        /Users/medya/go/pkg/mod/github.com/spf13/cobra@v1.9.1/command.go:1019 +0x82c
github.com/spf13/cobra.(*Command).ExecuteC(0x109900100)
        /Users/medya/go/pkg/mod/github.com/spf13/cobra@v1.9.1/command.go:1148 +0x384
github.com/spf13/cobra.(*Command).Execute(...)
        /Users/medya/go/pkg/mod/github.com/spf13/cobra@v1.9.1/command.go:1071
k8s.io/minikube/cmd/minikube/cmd.Execute()
        /Users/medya/workspace/minikube/cmd/minikube/cmd/root.go:174 +0x550
main.main()
```


