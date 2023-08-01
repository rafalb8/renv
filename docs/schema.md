# JSON file schema

 * distro - distro check
 * test - shell boolean expression
 * include - include other renv files 
 * packages - packages to install
 * files - files to copy
 * cmd - shell commands to be executed

## Example
```json
{
    "include": [
        "ohmyzsh.json",
        "archlinux/renv.json"
    ],
    "packages": [
        "btop",
        "bat",
        "exa",
        "rsync",
        "xclip",
        "tldr",
        "vim",
        "jq"
    ],
    "files": {
        "dotfiles/.zshrc": "$HOME/.zshrc",
        "dotfiles/.gitconfig": "$HOME/.gitconfig",
        "dotfiles/MangoHud.conf": "$HOME/.config/MangoHud/MangoHud.conf"
    }
}
```