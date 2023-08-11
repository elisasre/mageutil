#compdef mage

local curcontext="$curcontext" state line _opts ret=1

_arguments -C '*: :->targets' && ret=0

case $state in
  targets)
    _values "mage target" \
      $(mage | awk 'FNR > 1 {print $1}') && \
      ret=0
    ;;
esac

return ret

# Local Variables:
# mode: Shell-Script
# sh-indentation: 2
# indent-tabs-mode: nil
# sh-basic-offset: 2
# End:
# vim: ft=zsh sw=2 ts=2 et
