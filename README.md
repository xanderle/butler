# Butler

Butler keeps your home in order

Uses a toml file to specify which file should be symlinked where

See example host.toml and dotfiles for example

## Usage
all files must be within the dotfiles directory like so
```
dotfiles/
        zsh/
            zshrc
            zprofile
        alacritty
            alacritty.yml
```
Commands are executed like so
```shell
butler # default will symlink everything into place
butler -config=./path/to/file.toml # overrides default config file
butler -clean # cleans up existing symlinks
butler zsh alacritty # deploys zsh and alacritty
butler zsh -clean # cleans up symlinks for zsh and alacritty
```
## Roadmap
1. symlink files into place
2. template files and allow them to be "built"

