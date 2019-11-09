package htmltpl

const SubjectsPage = `
<!DOCTYPE html>
<html lang="en">
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="stylesheet" href="https://www.w3schools.com/w3css/4/w3.css">
<head>
    <meta charset="UTF-8">
    <title>Home</title>
</head>
<body>
<div class="w3-container">
	<p><a href="{{ .HomeAddresse }}">Home</a></p>
	</div>
<div class="w3-container w3-teal">
    <h1>Subject</h1>
</div>
<div class="w3-container">
    <p>Name : {{.Name}}</p>
	<h2>Versions</h2>
	<p>{{range .Versions}}<a href="{{ .Link }}">{{ .Version }}</p>{{end}}
</div>
</body>
</html>
`
