# About CLI

Dyno's CLI tool allows clients to make fuzzing requests to the Dyno Platform. In additon, it can be integrated within CI/CD pipelines.
Currently, the CLI supports Open API files in both JSON and YAML.

# Usage

To use Dyno please look at the following:

$: dyno -h
    Usage: dyno [--verbose] [--debug] <command> [<args>]

    Options:
      --verbose, -v          enable verbose logging
      --debug, -d            enable debug logging
      --help, -h             display this help and exit

    Commands:
      send                   can also use -d to provide the path to file

# To send Open API File

$: dyno send {path to open api file} 

e.g:
   dyno send /file/is/here/file.json

Output: No output if successfully sent otherwise appropriate error will be thrown informing you whats wrong

# How to integrate it with Github Actions

To use Dyno within your Github Actions workflow, you will need to add the dyno send command in the workflow.yml. You can integrate this at any stage that suits you.
e.g
 steps:
      - uses: actions/dyno-cli@v1.2.1
      - name: Fuzz Api
        run: |
	        dyno send /path/to/file.json
 