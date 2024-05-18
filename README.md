⚠️ Highly experimental plugin system for the Aurora server software

<h1 align="center">AuroraAPI 🌟</h1>

> A plugin API for Aurora<br />

## Support

| Platform     | External | Internal |
| :---         | :---:    | :---:    |
| **Windows**  | ❌       | ✅       |
| **Linux**    | ✅       | ✅       |

**Internal vs External:**
External plugins are loaded from a file. E.g. plugins/example.so. While internal plugins are directly built into the server by registering them in the plugin loader and building from source.

A break down of the pros and cons will be added later on.

**What does this mean?**
Plugins are supported on both platforms, but due to limitations with Go's plugin system, Windows is not fully supported. I am currently looking for workarounds and other solutions, even if they are a bit janky.

## Usage

See the [example plugin](https://github.com/MinimixMC/AuroraExamplePlugin)