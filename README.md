# rEnv
Tool for Linux distro configuration

# Installation
To install run:
```bash
curl ... | sh
```
or
```bash
wget -O - ... | sh
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