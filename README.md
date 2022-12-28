# rEnv
Tool for Linux distro configuration

# Installation
To install run:
```bash
curl -sfL https://raw.githubusercontent.com/rafalb8/renv/main/install.sh | sh
```

or
```bash
go install github.com/rafalb8/renv@latest
```

# Usage

## Apply enviroment
```bash
renv apply # to apply from $HOME/.renv or last applied configuration
renv apply path/to/config.json # to apply single file
renv apply path/to/dir # to apply renv.json in dir
```

## Edit enviroment
```bash
renv edit # run VS Code inside last applied config directory
```

## Install command
```bash
renv install vim # installs package using local package manager
```