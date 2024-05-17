# Build:

Linux:
```sh
go build -buildmode=plugin -o build/example.dll .
```

Windows:

⚠️ Requires CGO

```sh
go build -buildmode=c-shared -o build/example.dll .
```