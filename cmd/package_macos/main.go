package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

var plist = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CFBundleName</key>
	<string>Assurcast Reader</string>
	<key>CFBundleDisplayName</key>
	<string>Assurcast Reader</string>
	<key>CFBundleIdentifier</key>
	<string>com.audit-brands.assurcast-reader</string>
	<key>CFBundleVersion</key>
	<string>VERSION</string>
	<key>CFBundlePackageType</key>
	<string>APPL</string>
	<key>CFBundleExecutable</key>
	<string>assurcast-reader</string>

	<key>CFBundleIconFile</key>
	<string>icon</string>
	<key>LSApplicationCategoryType</key>
	<string>public.app-category.news</string>

	<key>NSHighResolutionCapable</key>
	<string>True</string>

	<key>LSMinimumSystemVersion</key>
	<string>10.13</string>
	<key>LSUIElement</key>
	<true/>
	<key>NSHumanReadableCopyright</key>
	<string>Copyright © 2026 Audit Risk Media. Based on narr by fiatjaf and yarr by Nazar Kanaev. MIT licensed.</string>
</dict>
</plist>
`

func run(cmd ...string) {
	fmt.Println(cmd)
	err := exec.Command(cmd[0], cmd[1:]...).Run()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var version, outdir string
	flag.StringVar(&version, "version", "0.0", "")
	flag.StringVar(&outdir, "outdir", "", "")
	flag.Parse()

	outfile := "assurcast-reader"
	appBundle := "Assurcast Reader.app"

	binDir := path.Join(outdir, appBundle, "Contents/MacOS")
	resDir := path.Join(outdir, appBundle, "Contents/Resources")

	plistFile := path.Join(outdir, appBundle, "Contents/Info.plist")
	pkginfoFile := path.Join(outdir, appBundle, "Contents/PkgInfo")

	os.MkdirAll(binDir, 0700)
	os.MkdirAll(resDir, 0700)

	f, _ := os.ReadFile(path.Join(outdir, outfile))
	os.WriteFile(path.Join(binDir, outfile), f, 0755)

	os.WriteFile(plistFile, []byte(strings.Replace(plist, "VERSION", version, 1)), 0644)
	os.WriteFile(pkginfoFile, []byte("APPL????"), 0644)

	iconFile := path.Join(outdir, "icon.png")
	iconsetDir := path.Join(outdir, "icon.iconset")
	os.Mkdir(iconsetDir, 0700)

	for _, res := range []int{1024, 512, 256, 128, 64, 32, 16} {
		outfile := fmt.Sprintf("icon_%dx%d.png", res, res)
		if res == 1024 || res == 64 {
			outfile = fmt.Sprintf("icon_%dx%d@2x.png", res/2, res/2)
		}
		cmd := []string{
			"sips", "-s", "format", "png", "--resampleWidth", strconv.Itoa(res),
			iconFile, "--out", path.Join(iconsetDir, outfile),
		}
		run(cmd...)
	}

	icnsFile := path.Join(resDir, "icon.icns")
	run("iconutil", "-c", "icns", iconsetDir, "-o", icnsFile)
}
