from subprocess import run
import json

def handler(event, context):
    print("logging started")
    restler_compile_cmd = "dotnet /RESTler/restler/Restler.dll --workingDirPath /tmp  compile --api_spec /tmp/swagger.json "
    restler_fuzz_cmd = r"""
    dotnet /RESTler/restler/Restler.dll --workingDirPath /tmp \
        fuzz-lean \
        --grammar_file /tmp/Compile/grammar.py \
        --dictionary_file /tmp/Compile/dict.json \
        --settings /tmp/Compile/engine_settings.json \
        --no_ssl
    """
    with open('/tmp/swagger.json', "w") as f:
        json.dump(event['swagger_file'],f)
    print("swagger file saved")
    run(restler_compile_cmd, shell=True)
    print("swagger file complied")
    run(restler_fuzz_cmd, shell=True)
    print("fuzzy complete")
    with open("/tmp/FuzzLean/ResponseBuckets/runSummary.json", "r") as f:
        results = json.load(f)        
    return results