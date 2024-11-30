# run

This package contains a universal mechanism to manage goroutine lifecycles. It 
implements an actor-runner with deterministic teardown. It is based on ideas 
from the https://github.com/oklog/run/ and enhances it with
configuration registration and validation as well as pre-run phase logic.

See godoc for information how to use  
[run.Group](https://pkg.go.dev/github.com/basvanbeek/run)
