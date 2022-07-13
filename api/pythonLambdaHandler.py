import json
from cgi import parse_header, parse_multipart
from io import BytesIO

def lambda_handler(event, context):
    # TODO implement
    #data = json.loads(event['body'])
    #encoded_file = data['file']
    #decoded_file = base64.decodestring(event['body'])
    
    #print event['queryStringParameters']['filename']
    #c_type, c_data = parse_header(event['headers']['content-type'])
    #c_data['boundary'] = bytes(c_data['boundary']).encode("utf-8")

    #body_file = BytesIO(bytes(event['body']).encode("utf-8"))
    #form_data = parse_multipart(body_file, c_data)

    #s3 = boto3.resource('s3')
    #object = s3.Object('test-store-swagger', "swagger")
    #object.put(Body=form_data['upload'][0])
    
    pathToFile = event['queryStringParameters']['pathToFile']
    #content = event['queryStringParameters']['swaggerContent']
    pathToFile = pathToFile+'done'
    #file = event['body']
    
    #with open('swagger.txt', 'w') as f:
    #    f.write(pathToFile)
    
    print("path to swagger file is = "+pathToFile)
    
    fuzzResponse = {}
    fuzzResponse['pathToFile']=pathToFile
    #fuzzResponse['swaggerContent']=content
    fuzzResponse['file']=event['headers']['content-type']
    fuzzResponse['message']='Here are your results'
    
    return {
        'statusCode': 200,
        'body': json.dumps(fuzzResponse)
    }