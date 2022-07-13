Cli-tool that can be integrated with github so clients can send their requests directly to dyno via github actions.

Code for the commands is in my_tool.py.

Commands:
-send : with flag -d, --dest that sends the request to dyno with the path of the swagger.json file

e.g command

To test, best to create a virtual environment, then change to the directory with code and run:
> python setup.py develop
> dyno send -d C:/something/somewhere/swagger.json