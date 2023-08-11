# Completions
Currently upstream mage doesn't support auto completions but since those are implemented separately from the binary itself it's not a problem.

### ZSH
Copy autocomplete script into $fpath:
```sh
cp mage.zsh /usr/share/zsh/site-functions/_mage
```
Start a new shell and auto completions should work properly.

# TODO:
    - [ ] bash completion
    - [ ] completion for flags
    - [ ] completion for default commands
    - [ ] description for each command
    - [ ] handle non-mage directories without error
    - [ ] upstream pr
