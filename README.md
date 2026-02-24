# SimpleApps
![GitHub Repo stars](https://img.shields.io/github/stars/Luis-Harz/SimpleApps)
![Go](https://img.shields.io/badge/Go-1.18+-00ADD8?style=flat&logo=go&logoColor=white)
![Platform](https://img.shields.io/badge/platform-Linux%20%7C%20Windows-informational)
![License](https://img.shields.io/badge/license-MIT-green?style=flat)
![Repo Size](https://img.shields.io/github/repo-size/Luis-Harz/SimpleApps?style=flat)
![Last Commit](https://img.shields.io/github/last-commit/Luis-Harz/SimpleApps?style=flat)

A collection of simple, fun terminal tools written in Go — all bundled into one interactive CLI.

---

## Features

SimpleApps ships with 24 built-in tools accessible from a single menu:

| # | Tool | Description |
|---|------|-------------|
| 0 | **NumberChecker** | Check if a number is odd or even |
| 1 | **GradeChecker** | Calculate letter grades from scores |
| 2 | **UnitConverter** | Convert between length, time, mass, and volume units |
| 3 | **Number2Bar** | Visualize a number as a terminal progress bar |
| 4 | **CoinFlip** | Flip a coin with ASCII animation |
| 5 | **Countdown** | Countdown timer in HH:MM:SS format |
| 6 | **Timer** | Stopwatch with start, stop, reset, and show commands |
| 7 | **Clock** | Live ASCII clock (style 1) |
| 8 | **Magic 8-Ball** | Ask a question, get an answer |
| 9 | **800+ Lines Special** | Animated noise screen easter egg |
| 10 | **Calculator** | Expression evaluator supporting `+`, `-`, `*`, `/`, and `()` |
| 11 | **ToDo List** | Persistent to-do list saved to `todos.json` |
| 12 | **Map Gen** | Generate and save/load random ASCII maps |
| 13 | **Matrix** | Scrolling binary matrix effect |
| 14 | **FakeLogGen** | Animated fake hacker log generator |
| 15 | **SysMonitor** | Live CPU and RAM usage with bar graphs |
| 16 | **ClockV2** | Live block-character ASCII clock (style 2) |
| 17 | **ASCII Animations** | Auto-cycling collection of terminal animations |
| 18 | **SimpleChat** | WebSocket-based chat client |
| 19 | **PasswordGen** | Configurable random password generator |
| 20 | **SimpleFiles** | Terminal based File Explorer |
| 21 | **Game** | Simple Arrow Key movement(not exactly a game) |
| 22 | **SimpleClick** | A global counter |
| 23 | **Wget** | A Wget Client |
| 24 | **SimplePic** | Simple "Pictures" with 3 colors |

You can also launch any tool directly by passing its number as a CLI argument (e.g. `./SimpleApps 10`).

---

## Installation

### Prerequisites

- [Go](https://go.dev/dl/) 1.18 or later

### Build from source

```bash
git clone https://github.com/Luis-Harz/SimpleApps.git
cd SimpleApps
go build -o SimpleApps .
```

On Windows:
```bash
go build -o SimpleApps.exe .
```

### Or download a release

Pre-built binaries are available on the [Releases](https://github.com/Luis-Harz/SimpleApps/releases) page. Download the binary for your platform, make it executable (Linux/macOS), and run it. Make sure you also download the latest version.txt!

```bash
chmod +x SimpleApps
./SimpleApps
```

> **Note:** SimpleApps requires a terminal of at least **70 columns × 23 rows**. You'll get a warning on startup if your terminal is too small.

### Or Permanently install(Linux)
Download the Zip extract it and run

```bash
./Permanent.sh
```
---
## SimpleChat Server
You can host your own SimpleChat Server by download the files <a href="https://github.com/Luis-Harz/SimpleApps/tree/main/Chat">here</a>

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/my-new-tool`)
3. Commit your changes (`git commit -m 'Add my new tool'`)
4. Push to the branch (`git push origin feature/my-new-tool`)
5. Open a Pull Request

Please keep new tools consistent with the existing style — interactive, terminal-friendly, and self-contained within a single function.

---

## Credits
New Icon made by <a href="https://github.com/Gustaf2580">Gustaf2580</a> <br>

## License

This project is licensed under the [MIT License](LICENSE).
