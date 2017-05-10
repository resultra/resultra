## Starting the Server

./bin/datasheetServer --config ~/datasheetConfig.json

## Profiling

Profiling is enabled with the --profile command line option. After stopping the server with Cntrl-C, a dump of the profile results can be obtained as follows:

	go tool pprof -text ./bin/datasheetServer PPROF_FILE

Where PPROF_FILE is the name of the pprof file generated for profiling.