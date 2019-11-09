package htmltpl

const SubjectDetailsPage = `
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
<p><a href="{{ .SubjectAddress }}">Subject {{.Subject}} page</a></p>
</div>
<div class="w3-container w3-teal">
	<h1>Subject details</h1>
</div>
<div class="w3-container">
	<p>Name: {{.Subject}}</p>
	<p>Version: {{.Version}}<p>
	<p>Id: {{.Id}}<p>
	<p>Schema: {{.Schema}}<p>
</div>
</body>
</html>
`
