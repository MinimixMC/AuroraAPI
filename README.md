⚠️ Highly experimental

<h1 align="center">AuroraAPI 🔮</h1>

> An API for the Aurora server software<br />

## Plugin Support

| Platform     | External | Internal |
| :---         | :---:    | :---:    |
| **Windows**  | ❌       | ✅       |
| **Linux**    | ✅       | ✅       |

**Internal vs External:**</br>
External plugins are loaded from a file. E.g. plugins/example.so. While internal plugins are directly built into the server by registering them in the plugin loader and building from source.

A break down of the pros and cons will be added later on.

**What does this mean?**</br>
Plugins are supported on both platforms, but due to limitations with Go's plugin system, Windows is not fully supported. I am currently looking for workarounds and other solutions, even if they are a bit janky.

## Usage
Add the package:
```sh
go get -u https://github.com/MinimixMC/AuroraAPI
```
See the [example plugin](https://github.com/MinimixMC/AuroraExamplePlugin)
