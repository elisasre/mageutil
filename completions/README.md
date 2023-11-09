# Completions
Currently upstream mage doesn't support auto completions but since those are implemented separately from the binary itself it's not a problem.

### ZSH
Copy autocomplete script into your `$fpath` where applicable, e.g. on Linux:

```sh
cp mage.zsh /usr/share/zsh/site-functions/_mage
```

#### Faster autocomplete

The autocompletion utilizes `mage -l` for target definitions and descriptions. The binary is cached but hitting tab still has a second or so
delay for each refresh. To speed up things, you can add following to your `.zshrc` file:

```
zstyle ':completions:*:mage:*' hash-fast true
```

This enables the Mage's fast hashing for files for the autocompletion.

### Next

Start a new shell and auto completions should work properly.

## TODO:
 - [ ] bash completion
 - [x] completion for flags
 - [x] completion for default commands
 - [x] description for each command
 - [ ] handle non-mage directories without error
 - [ ] upstream pr
