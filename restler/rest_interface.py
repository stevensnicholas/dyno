"""
This is not expected to by a publicly exposed API.
It is to be used via the backend interface

TODO make the docker ephemeral not permanently running after each reqest
"""
from flask import Flask, request
from subprocess import run
import json

app = Flask(__name__)

@app.post("/fuzz_api")
def fuzz_api():
    restler_compile_cmd = "dotnet /RESTler/restler/Restler.dll compile --api_spec /swagger.json"
    restler_fuzz_cmd = r"""
    dotnet /RESTler/restler/Restler.dll fuzz-lean \
        --grammar_file /Compile/grammar.py \
        --dictionary_file /Compile/dict.json \
        --settings /Compile/engine_settings.json \
        --no_ssl
    """
    if not request.is_json:
        return {"error": "Request must be JSON"}, 415

    input_data = request.get_json()
    with open('/swagger.json', "w") as f:
        json.dump(input_data['swagger_file'],f)
    app.logger.warning(run("ls /RESTler", shell=True))
    run(restler_compile_cmd, shell=True)
    run(restler_fuzz_cmd, shell=True)
    with open("/FuzzLean/ResponseBuckets/runSummary.json", "r") as f:
        results = json.load(f)        
    return {"results": results}, 201


if __name__ == '__main__':
    from waitress import serve
    serve(app, host="0.0.0.0", port=5000)