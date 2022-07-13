import click
import requests

@click.group()
def cli():
    pass


@cli.command()
@click.option('-d', '--dest', type=str, help='Path to Swagger File')
def send(dest):
    #click.echo(f'Shit {dest}')
    
    API_ENDPOINT = "https://o8cnchwjji.execute-api.ap-southeast-2.amazonaws.com/v1/post_json"
  
    path = dest
    data = ""
    with open(path, "r") as f:
        data = f.readlines()
    #print(data)
    #data = "yo"
    # data to be sent to api
    data = {'pathToFile':path}
    params = {'pathToFile':path}
    files = {'file': open(path, 'rb')}


    # sending post request and saving response as response object
    r = requests.post(url = API_ENDPOINT, data = data,params=params,files=files)

    # extracting response text 
    pastebin_url = r.text
    print("The pastebin URL is:%s"%pastebin_url)
