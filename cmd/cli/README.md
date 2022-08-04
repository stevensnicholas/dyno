# About CLI

Dyno's custom cli tool allows the clients to integrate Dyno within their CI/CD pipeline. It takes the OpenApi file given and sends it over to the Dyno backened for fuzzing.
Currently, the cli only takes in either JSON or YAML as the format types for the specified OpenApi File.

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

# To send OpenApi File

$: dyno send {Path To Open Api File} 

e.g:
   dyno send C://file/is/here/file.json

Output: No output if successfully sent otherwise appropriate error will be thrown informing you whats wrong

# How to integrate it with Github Actions

To use Dyno within your Github Actions workflow, you will need to add the dyno send command in the workflow.yml. You can integrate this at any stage that suits you.
e.g
 steps:
      - uses: actions/dyno-cli@v1.2.1
      - name: Fuzz Api
        run: |
	        dyno send /path/to/file.json
 