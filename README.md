> [!IMPORTANT]
> As it turns out, it appears most if not all tools that read GameConfig files support relative paths.
> This realisation turns this entire project mostly obsolete, if all the paths your require are resolvable with relative paths.
>
> As such, this project will be archived and is no longer maintained.
> If you need the features of this app for your own projects, feel free to fork of course.
>
> For an example GameConfig.txt to use as a base for the relative paths, see below:

```kv
"Configs"
{
    "Games"
    {
        "Neotokyo; Rebuild"
        {
            "GameDir"        "..\..\neo"
            "Hammer"
            {
                "GameData0"        "..\..\neo\rebuild.fgd"
                "TextureFormat"        "5"
                "MapFormat"        "4"
                "DefaultTextureScale"        "0.250000"
                "DefaultLightmapScale"        "16"
                "GameExe"        "..\..\ntre64.exe"
                "DefaultSolidEntity"        "func_detail"
                "DefaultPointEntity"        "info_player_start"
                "BSP"        "vbsp.exe"
                "Vis"        "vvis.exe"
                "Light"        "vrad.exe"
                "GameExeDir"        "..\.."
                "MapDir"        "..\..\neo\mapsrc"
                "BSPDir"        "..\..\neo\maps"
                "CordonTexture"        "tools\toolsskybox"
                "MaterialExcludeCount"        "0"
            }
        }
    }
    "SDKVersion"        "5"
}
```

# Hammer GameConfig.txt auto-initializer for NT;RE

Source code for building the `InitHammerGameConfig` tool, which initializes the NT;RE Hammer GameConfig.txt file with correct paths for the user.
The `InitHammerGameConfig.exe` executable is intended to be included at `steamapps\common\NEOTOKYOREBUILD\bin\x64` for the Steam build of NT;RE.

When `InitHammerGameConfig` runs, it expects to find a `steam_api64.dll` Steamworks API library file at that same path (for a hypothetical x86 32-bit build, the file would be `steam_api.dll`).

The executable can be run standalone, or as part of a Steam app [InstallScript](https://partner.steamgames.com/doc/sdk/installscripts), to automate Hammer setup for the user upon app installation.
It is mostly intended to be used as an InstallScript step.

## Installation
Pre-built releases are available to download at the [Releases page](https://github.com/Rainyan/GoHam/releases).

## Usage
* Download the release build, or build the app yourself.
* Make sure an up-to-date Steamworks SDK `steam_api64.dll` library exists in the same path as the `InitHammerGameConfig.exe` executable.
  * The easiest way to do this is by copying the files `InitHammerGameConfig.exe` and `GameConfig.txt.pre` from the release to your `steamapps/common/NEOTOKYOREBUILD/bin/x64` folder.
  * More specifically, the Steamworks SDK dll version used should be the same as the one our `go-steamworks` [pinned module version](https://github.com/Rainyan/GoHam/blob/c29f1f3060ec0e354ab82448dfbd7ba38417b26a/go.mod#L6) is targeting, which can be found at the [project README](https://github.com/hajimehoshi/go-steamworks/blob/ea9c0844b066/README.md#steamworks-sdk-version). Assuming commit `ea9c0844b066` in this example, but check our go.mod to be sure of the up-to-date version.
* Make sure the `base.fgd` and `halflife2.fgd` files exist in the folder `steamapps/common/NEOTOKYOREBUILD/bin/x64`.
  * At the time of writing this, you have to copy-paste them from `steamapps/common/NEOTOKYOREBUILD/bin` to the `x64` folder, but this should be fixed in the Steam release itself.
* Make sure Steam is running, and you are logged in with an account that has NT;RE installed on the machine.
* Run the `InitHammerGameConfig.exe` executable.
* Observe how `GameConfig.txt` was generated from the `GameConfig.txt.pre` file (or its contents updated if it already existed).
* Run the `hammer.exe` executable in `steamapps/common/NEOTOKYOREBUILD/bin/x64`, and it should now be ready to use with the generated `GameConfig.txt` for NT;RE mapping.

## Building
The following is only needed if you wish to build this app from source code.
### Requirements
* A modern version of Git
* A version of [Golang](https://go.dev/) specified at [go.mod](go.mod#L3)
* Assuming Windows platform (although this could well be made cross-platform, assuming our tools become Linux compatible at some point)
* Assuming `steam_api64.dll` located at the same path as the executable. For testing, you can [download the Steamworks SDK](https://partner.steamgames.com/downloads/steamworks_sdk.zip)
  and copy-paste the file from the `redistributable_bin` folder.
  * For the NT;RE Steam build, this dll file should already be in place for the Windows build.

### Build instructions
This assumes you have [Git and Go installed and invokable](#requirements) from your command line as `git` and `go`.
```cmd
git clone <this_repo_url>
CD <this_repo_path>
go mod tidy
go build
```

## Acknowledgements
This tool uses the following open-source software:
* [go-steamworks](https://github.com/hajimehoshi/go-steamworks) - A Steamworks SDK binding for Go ([Apache-2.0 license](LICENSES/LICENSE-go-steamworks.txt))
* [ventil](https://github.com/noxer/ventil) - Valve Key-Value file parser in Go ([MIT license](LICENSES/LICENSE-ventil.txt))
