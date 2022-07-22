from subprocess import run
import json

def handler(event, context):
    print("logging started")
    restler_compile_cmd = "dotnet /RESTler/restler/Restler.dll compile --api_spec /swagger.json"
    restler_fuzz_cmd = r"""
    dotnet /RESTler/restler/Restler.dll fuzz-lean \
        --grammar_file Compile/grammar.py \
        --dictionary_file Compile/dict.json \
        --settings Compile/engine_settings.json \
        --no_ssl
    """
    with open('/swagger.json', "w") as f:
        json.dump(event['swagger_file'],f)
    print("swagger file saved")
    run(restler_compile_cmd, shell=True)
    print("swagger file complied")
    run(restler_fuzz_cmd, shell=True)
    print("fuzzy complete")
    with open("FuzzLean/ResponseBuckets/runSummary.json", "r") as f:
        results = json.load(f)        
    return results