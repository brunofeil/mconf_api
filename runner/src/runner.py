import requests
import sys
import json

if __name__ == '__main__':
    
    if len (sys.argv) != 2 :
        print(" -> arg missing <-")
        sys.exit (1)

    args = sys.argv[1:]
    args[0] = args[0].replace(' ', '+')

    r = requests.get('http://host.docker.internal:8080/'+args[0])
    response = r.text

    configured_json = json.loads(response)
    books = configured_json['docs']

    print("\nLivros encontrados:")
    for book in books:
        print("\t- "+book['title'])