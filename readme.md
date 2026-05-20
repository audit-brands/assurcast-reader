# Assurcast Reader

**Assurcast Reader** is a desktop and self-hosted RSS + Nostr long-form feed reader,
purpose-built for internal audit and assurance professionals who want a clean way to
follow [Assurcast](https://assurcast.com) and other industry sources.

The app is a single binary with an embedded SQLite database. It is built on top of
[narr](https://github.com/fiatjaf/narr) by fiatjaf, which is itself a fork of
[yarr](https://github.com/nkanaev/yarr) by Nazar Kanaev.

![screenshot](etc/promo.png)

## install

The latest prebuilt binaries for Linux/MacOS/Windows AMD64 are available
[here](https://github.com/audit-brands/assurcast-reader/releases/latest).

* MacOS

  Download `assurcast-reader-*-macos64.zip`, unzip it, place `Assurcast Reader.app` in `/Applications`, [open the app][macos-open], click the menu bar icon, select "Open".

* Windows

  Download `assurcast-reader-*-windows64.zip`, unzip it, open `assurcast-reader.exe`, click the system tray icon, select "Open".

* Linux

  Download `assurcast-reader-*-linux64.zip`, unzip it, place `assurcast-reader` in `$HOME/.local/bin`
  and run [the install script](etc/install-linux.sh).

[macos-open]: https://support.apple.com/en-gb/guide/mac-help/mh40616/mac

For self-hosting, see `assurcast-reader -h` for auth, tls & server configuration flags.

## see more

* [Building from source](doc/build.md)
* [Fever API support](doc/fever.md)

## credits

* [narr](https://github.com/fiatjaf/narr) by fiatjaf — Nostr long-form support
* [yarr](https://github.com/nkanaev/yarr) by Nazar Kanaev — original RSS reader
* [Feather](http://feathericons.com/) — icons

## license

MIT — see [license](license).
