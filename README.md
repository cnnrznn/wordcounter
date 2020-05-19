# Wordcounter

This is an app that converts a URL into a wordcount. The backend performs a
GET request on the url and passes the page data through a wordcount
computation. The frontend is a simple form that passes the provided URL to
the backend and renders the HTML response

## Deployment
1. Have kubectl connected to a kubernetes cluster
2. `skaffold run` in root dir

## Based on
This code is based on the vscode example
[here](https://github.com/GoogleCloudPlatform/cloud-code-samples/tree/master/golang/go-guestbook).