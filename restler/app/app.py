import uuid
from subprocess import run
import json
import boto3
import os
import shutil
import logging
from urllib.parse import urlparse

logger = logging.getLogger()
logger.setLevel(logging.INFO)


def handler(event, context):
    # Restler consistently tries to save files to working directory
    # So to avoid read only errors in lambda change wdir to /tmp
    os.chdir("/tmp")
    logger.info("logging started")
    local_api_spec_path = "/tmp/swagger.json"
    restler_compile_cmd = f"dotnet /RESTler/restler/Restler.dll --workingDirPath /tmp  compile --api_spec {local_api_spec_path} "
    token = ""
    repo = ""
    owner = ""
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
    # SQS check
    if "Records" in event:
        if len(event["Records"]) != 1:
            raise ValueError("This lambda only supports one record at a time")
        event = json.loads(event["Records"][0]["body"])
    s3 = boto3.client("s3")
    if "s3_location" in event:
        swagger_url = urlparse(event["s3_location"])
        logger.info(f"Trying to get: s3://{swagger_url.netloc}{swagger_url.path}")
        s3.download_file(
            swagger_url.netloc, swagger_url.path.lstrip("/"), local_api_spec_path
        )
    elif "swagger_json" in event:
        with open(local_api_spec_path, "w") as f:
            json.dump(event["swagger_json"], f)
    else:
        raise KeyError(
            "No swagger file provided: Input needs to have either s3_location or swagger_json key"
        )
    logger.info(f"swagger file saved at {local_api_spec_path}")
    run(restler_compile_cmd, shell=True)
    logger.info("swagger file complied")
    # force no Garbage collection, not supported in lambda
    with open("/tmp/Compile/engine_settings.json", "r") as f:
        engine_settings = json.load(f)
        engine_settings["garbage_collection_interval"] = 0
    with open("/tmp/Compile/engine_settings.json", "w") as f:
        json.dump(engine_settings, f)
    run(restler_fuzz_cmd, shell=True)
    logger.info("fuzzing-lean complete")
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
        logger.info(f"S3 uploaded response for {prefix}: {response}")
        logger.info(f"S3 {prefix} uploaded to s3://{bucket_name}/{key}")
    key = f"{random_prefix}/results.zip"
    snsMessage = {
        "location": f"{bucket_name}/{key}",
        "uuid": f"{random_prefix}",
        "token": f"{token}",
        "owner": f"{owner}",
        "repo": f"{repo}",
    }
    if os.environ.get("issues_sns_topic_arn", False):
        logger.info(f'Publish to SNS {os.environ["issues_sns_topic_arn"]}')
        sns = boto3.client("sns")
        response = sns.publish(
            TopicArn=os.environ["issues_sns_topic_arn"],
            Message=json.dumps({"default": json.dumps(snsMessage)}),
            MessageStructure="json",
        )
    with open("/tmp/FuzzLean/ResponseBuckets/runSummary.json", "r") as f:
        results = json.load(f)
    return results
