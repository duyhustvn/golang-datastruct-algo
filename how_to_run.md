## Test

```
go test./package_to_run_test -count=1 -v
```

Example:

```
go test./graph/medium/time_need_to_inform_all_employees -count=1 -v
```

Where:
* -count=1 disable cache 
* -v to print log in case you use <em>t.Log</em> in your test case
