import uuid
from subprocess import run
import json
import boto3
import os
import shutil

def handler(event, context):
    #Restler consistently tries to save files to working directory
    #So to avoid read only errors in lambda change wdir to /tmp
    os.chdir("/tmp")
    print("logging started")
    restler_compile_cmd = "dotnet /RESTler/restler/Restler.dll --workingDirPath /tmp  compile --api_spec /tmp/swagger.json "
    restler_fuzz_cmd = r"""
    dotnet /RESTler/restler/Restler.dll \
        --workingDirPath /tmp \
        --logsUploadRootDirPath /tmp \
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
    print("fuzzing-lean complete")
    s3 = boto3.client("s3")
    bucket_name = os.environ['results_upload_s3_bucket']
    random_prefix = uuid.uuid4()
    results_key = f"{random_prefix}/results.zip"
    logs_key = f"{random_prefix}/logs.zip"
    shutil.make_archive("/tmp/results", 'zip', "/tmp/FuzzLean")
    shutil.make_archive("/tmp/logs", 'zip', "/tmp/RestlerLogs")
    response_results = s3.upload_file("/tmp/results.zip", bucket_name, logs_key)
    response_logs = s3.upload_file("/tmp/logs.zip", bucket_name, results_key)
    print(f"S3 upload response for results: {response_results}")
    print(f"S3 upload response for logs: {response_logs}")
    with open("/tmp/FuzzLean/ResponseBuckets/runSummary.json", "r") as f:
        results = json.load(f)        
    return results