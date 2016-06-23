import requests
import json

class UserAuthTestHelperMixin:
    def authRequest(self, apiPath,jsonArgs):
        baseURL = 'http://localhost:8080/auth/'
        fullURL = baseURL + apiPath
        print "UserAuthTestHelperMixin: Request: ",apiPath,": args=",json.dumps(jsonArgs)
        resp = requests.post(fullURL,json=jsonArgs)
        if resp.status_code != 200:
            print "UserAuthTestHelperMixin: Error response: ",resp.text
        self.assertEqual(resp.status_code,200,"expecting success return code from server")
        print "UserAuthTestHelperMixin: Response: ",json.dumps(resp.json())
        return resp.json()


    