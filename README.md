# Bug Experiment
Repository to test bug found in https://github.com/gocql/gocql/issues/296

To run experiment:
```bash
./prepare.sh;
go run *.go;
nodetool cfstats experiment
```

Then, check if the average number of tombstones is 0. If not, you've reproduced the bug. Hooray!