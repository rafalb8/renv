# rEnv
 - better version of dotfiles python script
 - in $home/.renv on root or split into multiple dirs -> every dir new env config
 - env dir must contain renv.json or sth

# example files
```json
// renv.json
{
    "include":["ohmyzsh.json"],
    "bin": ["btop", "git", "zsh"],
    "files": {
        ".zshrc": "$HOME/.zshrc",
        ".gitconfig": "$HOME/.gitconfig"
    }
}
// ohmyzsh.json
{
    "bin": ["curl","git", "zsh"],
    "cmd": [
        "curl ... | sh"
    ]
}
```
 
# cli
 - installed in /bin or other specified location with curl/wget ... | sh script
 - compiled for x86_64 and aarch64
 - db in $HOME/.config/renv
 - should work if run from root or from other user

## run steps
 - run `include`s
 - install bins from `bin`
 - run `cmd`
 - copy `files` (no patching)

## examples
 - renv apply home - apply config from home dir
 - renv apply work - apply config from work dir
 - renv apply - apply last config
 - renv edit - open vscode in $HOME/.renv
 - renv install btop zsh ... - install packages using local package manager 


