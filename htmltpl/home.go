package htmltpl

const HomePage = `
<!DOCTYPE html>
<html lang="en">
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="stylesheet" href="https://www.w3schools.com/w3css/4/w3.css">
<head>
    <meta charset="UTF-8">
    <title>Home</title>
</head>
<body>
<div class="w3-container w3-teal">
    <h1>Schema Registry UI</h1>
</div>
<div class="w3-container">
    <p>0.1.0</p>
    <h2>Schema registry server informations</h2>
    <p>Address : {{.Address}}</p>
    <p>Global compatibility : {{.GlobalCompat}}</p>
    <h2>Subjects</h2>
    <p>{{range .ListSubjects}}<a href="/subjects/{{ . }}/">{{ . }}</p>{{else}} empty subjects list{{end}}
</div>
</body>
</html>
`
