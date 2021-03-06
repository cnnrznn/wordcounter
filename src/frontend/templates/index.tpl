{{ define "index" }}
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>Wordcounter</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
    <link rel="stylesheet" href="/static/style.css" type="text/css">
    <link rel="icon" href="/static/favicon.png" type="image/x-icon"/>
    <link href="https://fonts.googleapis.com/css?family=Roboto" rel="stylesheet">
</head>

<body>
    <div class="header">
        <div class="container">
            <h1>
                <a href="/">
                    A Web Wordcounter
                </a>
            </h1>
            <a href="#" class="text-muted">View on GitHub</a>
        </div>
    </div>

    <div class="container posts mt-0">
        <form class="form-inline" method="POST" action="/post">
            <label class="sr-only" for="name">Name</label>
            <div class="input-group mb-2 mr-sm-2">
                <div class="input-group-prepend">
                    <div class="input-group-text">URL</div>
                </div>
                <input type="text" class="form-control" id="url" name="url" required>
            </div>
            <button type="submit" class="btn btn-primary mb-2">Compute wordcount</button>
        </form>

        {{ if .Wordcount }}
            <div class="card my-3 col-12">
                <div class="card-body">
                    <h5 class="card-title">{{ .URL }}</h5>
                    <h6 class="card-subtitle mb-2 text-muted">Computed in {{ .Time }} seconds</h6>
                    <br>
                    {{ range .Wordcount }}
                    <p class="card-text">
                        {{ .Key }}, {{ .Val }}
                    </p>
                    {{ end }}
                </div>
            </div>
        {{ else }}
            <div class="alert alert-info" role="alert">
                Enter a URL to perform wordcount
            </div>
        {{ end }}
    </div>
</body>
</html>
{{ end }}
