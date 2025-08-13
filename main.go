package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/hajimehoshi/go-steamworks"
	"github.com/noxer/ventil"
)

// Initializes Steamworks API with the provided Steam AppID.
func steamInPlaceInit(appid steamworks.AppId_t) error {
	// Uses the file-based appid detection method.
	// https://partner.steamgames.com/doc/sdk/api#SteamAPI_Init
	const appidFile = "steam_appid.txt"
	f, err := os.Open(appidFile)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
		wf, err := os.Create(appidFile)
		if err != nil {
			return err
		}
		// If we made the file, we should also clean it up.
		// But if it existed originally, don't delete it.
		defer os.Remove(appidFile)
		defer wf.Close()
		n, err := wf.WriteString(strconv.Itoa(int(appid)))
		if err != nil {
			return err
		} else if n <= 0 {
			return fmt.Errorf("n (%d) <= 0", n)
		}
		err = wf.Sync()
		if err != nil {
			return err
		}
		// Re-enter once, because we now have the file set up.
		return steamInPlaceInit(appid)
	}
	defer f.Close()
	// Sanity check; verify file contents match the app ID
	contents, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	if string(contents) != strconv.Itoa(int(appid)) {
		return fmt.Errorf("steam_appid.txt contents differ from appid (%d). Contents:\n%s", appid, contents)
	}
	return steamworks.Init()
}

// Initializes Steamworks API with the provided Steam AppID,
// and then processes the GameConfig KeyValues to fixup paths for the user.
// The KeyValues are updated in-place, via the pointer.
func parseCfgInPlace(appid steamworks.AppId_t, kv *ventil.KV) error {
	err := steamInPlaceInit(appid)
	if err != nil {
		return err
	}

	replaces := map[string]string{
		// %INSTALLDIR% is a faux variable, mimicing the one defined at: https://partner.steamgames.com/doc/sdk/installscripts#environment_variables
		// So we gotta manually replace these occurrences with the full app path.
		"%INSTALLDIR%": steamworks.SteamApps().GetAppInstallDir(appid),
	}

	// Note that while we parse backslashes here too, they're not recommended for the pre-GameConfig,
	// because the KeyValues parser we're using to read the file will not parse backslashes correctly.
	// For example "\neo" would be interpreted as: '\n' + "eo"
	regexes := []*regexp.Regexp{
		regexp.MustCompile(`\\+`),
		regexp.MustCompile(`\/+`),
	}

	kv.Tree(func(key string, kv *ventil.KV) {
		if kv.Value == "" {
			return
		}
		for k, v := range replaces {
			kv.Value = strings.ReplaceAll(kv.Value, k, v)
		}
		for _, re := range regexes {
			kv.Value = re.ReplaceAllString(kv.Value, `\`)
		}
	})
	return nil
}

// Validates the GameConfig format, checking for incorrect contents.
func validateCfgFormat(cfgPath string) error {
	file, err := os.Open(cfgPath)
	if err != nil {
		return err
	}
	defer file.Close()

	formattedContents, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	// Expecting single backslashes for Windows-style path separators
	if strings.Contains(string(formattedContents), `\\`) {
		return fmt.Errorf("invalid formatting: multiple contiguous backslashes found in:\n%s", formattedContents)
	}
	// Currently not expecting to see Unix-style path separators
	if strings.Contains(string(formattedContents), "/") {
		return fmt.Errorf("invalid formatting: forwardslash but expected backslash found in:\n%s", formattedContents)
	}
	// Assuming we have handled all %COMMAND% style tokens, and no other % characters are expected either
	if strings.Contains(string(formattedContents), "%") {
		return fmt.Errorf("invalid formatting: stray percent sign(s) (%%) found in:\n%s", formattedContents)
	}

	return nil
}

// Entry point
func main() {
	const appid = 3172910 // Steam AppID for "Neotokyo; Rebuild"
	const cfgPathWrite = "GameConfig.txt"
	const cfgPathRead = cfgPathWrite + ".pre"

	cfg, err := ventil.ParseFile(cfgPathRead)
	if err != nil {
		log.Fatal(err)
	}

	err = parseCfgInPlace(appid, cfg)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(cfgPathWrite)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	cfg.WriteTo(file)

	validateCfgFormat(cfgPathWrite)
}
