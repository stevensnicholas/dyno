import uuid
from subprocess import run
import json
import boto3
import os
import shutil


def handler(event, context):
    # Restler consistently tries to save files to working directory
    # So to avoid read only errors in lambda change wdir to /tmp
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
    with open("/tmp/swagger.json", "w") as f:
        json.dump(event["swagger_file"], f)
    print("swagger file saved")
    run(restler_compile_cmd, shell=True)
    with open("/tmp/Compile/engine_settings.json", "r") as f:
        engine_settings = json.load(f)
        engine_settings["garbage_collection_interval"] = 0
    with open("/tmp/Compile/engine_settings.json", "w") as f:
        json.dump("/tmp/Compile/engine_settings.json")
    print("swagger file complied")
    run(restler_fuzz_cmd, shell=True)
    print("fuzzing-lean complete")
    s3 = boto3.client("s3")
    bucket_name = os.environ["results_upload_s3_bucket"]
    random_prefix = uuid.uuid4()
    for folder, prefix in [
        ("RestlerLogs", "logs"),
        ("FuzzLean", "results"),
        ("Compile", "compile"),
    ]:
        key = f"{random_prefix}/{prefix}.zip"
        shutil.make_archive(f"/tmp/{prefix}", "zip", f"/tmp/{folder}")
        response = s3.upload_file(f"/tmp/{ prefix }.zip", bucket_name, key)
        print(f"S3 uploaded response for {prefix}: {response}")
        print(f"S3 {prefix} uploaded to s3://{bucket_name}/{key}")
    with open("/tmp/FuzzLean/ResponseBuckets/runSummary.json", "r") as f:
        results = json.load(f)
    return results
