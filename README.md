# Hammer GameConfig.txt auto-initializer for NT;RE

Source code for building the InitHammerGameConfig tool, which initializes the NT;RE Hammer GameConfig with appropriate paths for the user.

The built executable file (`InitHammerGameConfig.exe`) is intended to be included at `steamapps\common\NEOTOKYOREBUILD\bin\x64` for the Steam build.
When the executable runs, it assumes to find a `steam_api64.dll` (or `steam_api.dll` for any hypothetical 32-bit build - in which case you'd also need a 32-bit Golang toolchain; probably not worth it)
Steamworks API library at that same path.

## Building
### Requirements
* A version of [Golang](https://go.dev/) specified at [go.mod](go.mod#L3)
* Assuming Windows platform (although this could well be made cross-platform, assuming our tools become Linux compatible at some point)
* Assuming `steam_api64.dll` located at the same path as the executable. For testing, you can [download the Steamworks SDK](https://partner.steamgames.com/downloads/steamworks_sdk.zip)
  and copy-paste the file from the `redistributable_bin` folder.
  * For the NT;RE Steam build, this dll file should already be in place for the Windows build.

### Build instructions
This assumes you have Git and Go installed and invokable from your command line as `git` and `go`.
```cmd
git clone <this_repo_url>
CD <this_repo_path>
go mod tidy
go build
```

## Usage
* Build the executable
* Make sure an up-to-date Steamworks SDK `steam_api64.dll` library exists in the same path
  * More specifically, the Steamworks SDK dll version used should be the same as the one our `go-steamworks` [pinned module version](https://github.com/Rainyan/GoHam/blob/c29f1f3060ec0e354ab82448dfbd7ba38417b26a/go.mod#L6) is targeting, which can be found at the [project README](https://github.com/hajimehoshi/go-steamworks/blob/ea9c0844b066/README.md#steamworks-sdk-version). Assuming commit `ea9c0844b066` in this example but check our go.mod to be sure of the up-to-date version.
* Run the executable
* Observe how `GameConfig.txt` was generated, or updated, in the executable folder.

## Acknowledgements
This tool uses the following open-source software:
* [go-steamworks](https://github.com/hajimehoshi/go-steamworks) - A Steamworks SDK binding for Go ([Apache-2.0 license](https://github.com/hajimehoshi/go-steamworks?tab=Apache-2.0-1-ov-file#readme))
* [ventil](https://github.com/noxer/ventil) - Valve Key-Value file parser in Go ([MIT license](https://github.com/noxer/ventil?tab=MIT-1-ov-file#readme))
