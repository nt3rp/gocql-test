# Bug Reproduction
Repository to test bug found in https://github.com/gocql/gocql/issues/296

To run attempt to reproduce the bug:

- Install [`ccm`](https://github.com/pcmanus/ccm)
- Create a two node cassandra setup

```bash
ccm create bug-reproduction -v 2.0.5
ccm populate -n 2
ccm start
```

- Setup the test, and run it

```bash
./prepare.sh
go run *.go
nodetool cfstats experiment
```

The test sets up the schema, creates a prepare statement for a number of scenarios, and runs the test cases one thousand times each.
If the average number of tombstones is greater than 0, you may have reproduce the issue.

Still figuring this out.