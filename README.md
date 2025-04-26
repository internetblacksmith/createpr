# CreatePR

[![Release](https://img.shields.io/github/release/internetblacksmith/createpr.svg?style=for-the-badge)](https://github.com/internetblacksmith/createpr/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=for-the-badge)](/LICENSE.md)
[![Build status](https://img.shields.io/github/actions/workflow/status/internetblacksmith/createpr/go.yml?style=for-the-badge&branch=main)](https://github.com/goreleaser/goreleaser/actions?workflow=build)
[![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=for-the-badge)](https://github.com/goreleaser)
[![GoReportCard](https://goreportcard.com/badge/github.com/internetblacksmith/createpr?style=for-the-badge)](https://goreportcard.com/report/github.com/internetblacksmith/createpr)

A simple command-line utility that opens GitHub's "New Pull Request" page in your browser directly from your terminal. Skip the manual navigation and instantly create PRs for your current branch.

## ✨ Features

- 🔄 **One Command**: Run `createpr` and immediately open the correct GitHub PR creation page
- 🔍 **Auto-detection**: Automatically detects GitHub repository URL from Git remotes
- 🌿 **Branch Aware**: Uses your current branch as the source branch for the PR
- 🔒 **Secure**: Works with both HTTPS and SSH remote URLs
- 🚀 **Fast**: No dependencies other than Git itself

## 📦 Installation

### Using Homebrew (macOS)

```bash
brew tap internetblacksmith/internetblacksmith
brew install createpr
```

### Using Scoop (Windows)

```bash
scoop bucket add internetblacksmith https://github.com/internetblacksmith/scoop-bucket
scoop install internetblacksmith/createpr
```

### Using Pre-built Binaries

1. Go to the [Releases page](https://github.com/internetblacksmith/createpr/releases)
2. Download the appropriate binary for your system:
   - `createpr_linux_amd64.tar.gz` for Linux (64-bit)
   - `createpr_darwin_amd64.tar.gz` for macOS (Intel)
   - `createpr_darwin_arm64.tar.gz` for macOS (Apple Silicon)
   - `createpr_windows_amd64.zip` for Windows (64-bit)
3. Extract the archive
4. Move the `createpr` executable to a directory in your PATH

### Using Go Install

If you have Go installed:

```bash
go install github.com/internetblacksmith/createpr@latest
```

## 🚀 Usage

Navigate to any Git repository in your terminal and run:

```bash
createpr
```

This will:
1. Detect the GitHub repository URL from your Git remote
2. Identify your current branch
3. Open your default web browser to GitHub's "New Pull Request" page with your branch pre-selected

### Examples

```bash
# Basic usage
cd ~/projects/my-repo
createpr
```

## 🧪 Development

### Prerequisites

- Go 1.18 or higher
- Git

### Building from Source

```bash
# Clone the repository
git clone https://github.com/internetblacksmith/createpr.git
cd createpr

# Build
go build

# Run tests
go test
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🤔 FAQ

### ☝️🤓 "Uhm, actually... GitHub CLI can already do this with `gh pr create --web`"

Yes, and pizza delivery exists too, but sometimes you just want to grab a slice without filling out forms! 

CreatePR is the "no authentication, zero API calls, instant gratification" solution:

- **Zero login needed** - No "please authenticate" dance required
- **Lightning fast** - Makes absolutely no network requests (the gh CLI makes several)
- **Featherweight champion** - Tiny binary that does one thing perfectly
- **Works offline** - Your spotty coffee shop WiFi can't stop you from getting to that PR page
- **No relationship baggage** - Doesn't need to know your GitHub username, tokens, or life story

If you're already happily married to the GitHub CLI, that's cool! But for a commitment-free PR creation experience, CreatePR is your speed-dating alternative.

### 🔥 Why did you build this when other tools exist?

I got tired of:
1. Click repository
2. Click Pull requests
3. Click New pull request
4. Click branch dropdown
5. Scroll... scroll... scroll...
6. Click my branch
7. Question my life choices

Now it's just `createpr` → *browser opens* → *PR ready to complete*

### 🛠️ Will you add feature X?

Probably not. The beauty of CreatePR is its single-minded focus - it opens PR pages and nothing else. It's not trying to boil the ocean or become your GitHub Swiss Army knife.

That said, if you have an idea that keeps it simple while improving the core functionality, submit an issue! Just remember: if your feature request includes the phrase "and then it could also..." it's probably a no.

### 👻 Does this actually create the PR for me?

Nope! CreatePR just teleports you to the PR creation page with your current branch pre-selected. You still get to write that thoughtful PR description your colleagues will definitely read thoroughly.

Think of it as an express elevator to the PR floor, not an automated PR robot.

## 🙏 Acknowledgments

- Inspired by a rant from [Theo](https://x.com/theo) in one of his streams 
- Built with [Go](https://golang.org/)
- Released with [GoReleaser](https://goreleaser.com/)

---

<p align="center">
  Made with ❤️ by <a href="https://github.com/internetblacksmith">Internet Blacksmith</a>
</p>
```

## Planned improvements

 - [ ] Add option to open the extended PR view
 - [ ] Handle other services BitBucket/Gitlab etc...
