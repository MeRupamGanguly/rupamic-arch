
`option go_package = "/gogen";` Ea proto file ka line dockerfile ke `RUN protoc --proto_path=/ticker/domain --go_out=/ticker/domain --go-grpc_out=/ticker/domain /ticker/domain/ticker.proto` `_out` path se append ho jata hai and `/ticker/domain/gogen` pai apne generated code ko rakhta hai.

```bash
docker build -t grpcgen -f dockerfile.ticker .
docker run -it grpcgen
docker cp 5d7352636391:/ticker/domain/gogen ./ticker/domain
```

Jenkins:

agent: Defines where the pipeline runs
Stages Divides the pipeline into multiple steps (such as Build, Test, Deploy).
post: Defines actions that are executed after the pipeline or a stage completes (e.g., cleanup, notifications).
Parallel Stages: To run stages in parallel and speed up pipeline execution.
&> This redirects both standard output (stdout) and standard error (stderr) to the specified location (in this case, /dev/null).


Steps:- 
Pipeleine -> Agent {label->Name}
env variables set -> environment{Name=""}
Stages{stage{steps{}}}


find . -path ./integration -prune -o -name "*_test.go" -exec go test -v {} \;
find: This is the command used to search for files and directories in a given location.

The dot (.) specifies the current directory as the starting point for the search. So it means "start looking from here (the current directory) and go down through all the subdirectories".

-path: This option matches files or directories with the given path pattern.

./integration: This is the path you want to exclude. It specifically matches the integration folder in the current directory.

-prune: When this option is used, find will exclude the matching directory or file from further processing (meaning it won’t search inside or return results from this path).

-o: This is a logical OR operator. It tells find to continue processing the next condition if the current condition is false.

-name "*_test.go": This tells find to look for files whose names end with _test.go. These are typically the test files in Go projects.

-exec: This option allows you to execute a command on the files found by find.


go test -v: This is the command that will be run on each file found by find. The go test -v command runs the Go tests with verbose output.

{}: This is a placeholder that represents the current file or directory found by find.

\;: This ends the -exec command and tells find to execute the command for each file it finds.

go test -v $(go list ./... | grep -v '/integration')

$(...): This is a way to run a command inside another command and replace it with the output of the inner command. It’s called command substitution.

go list ./...: This command lists all Go packages in the current directory (./) and its subdirectories recursively.

. refers to the current directory.
... is a special wildcard in Go that means all subdirectories. So, ./... means "start from the current directory and go through all subdirectories recursively."

|: This is the pipe operator, which takes the output of the command on the left and passes it as input to the command on the right.

grep -v '/integration': This command filters out (excludes) any line that contains /integration.
grep is used to search through the output of the previous command (go list ./...).
-v means invert match — it tells grep to exclude lines that match the given pattern (in this case, /integration).



