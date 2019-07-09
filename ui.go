package up

import "os"

type itemType int

const (
	fileType itemType = iota
	dirType
)

type Item struct {
	Name string
	Type itemType
	URL  string
}

func GetItemType(fi os.FileInfo) itemType {
	if fi.IsDir() {
		return dirType
	}
	return fileType
}

const Template = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>serve</title>
</head>
<body>
    <h1>serve</h1>
    <ul>
        {{range .Items}}
            <li>
                <a href="{{.URL}}">
                    {{.Name}}
                </a>
            </li>
        {{end}}
    </ul>
</body>
</html>
`
