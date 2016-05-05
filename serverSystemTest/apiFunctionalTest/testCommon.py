import requests

class TestHelperMixin:
    def apiRequest(self, apiPath,jsonArgs):
        baseURL = 'http://localhost:8080/api/'
        fullURL = baseURL + apiPath
        resp = requests.post(fullURL,json=jsonArgs)
        self.assertEqual(resp.status_code,200,"expecting success return code from server")
        return resp.json()
