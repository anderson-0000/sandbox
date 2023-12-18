disable r
export PATH="$PATH:/usr/local/user-add/flutter/bin"

function mkdir-date () {
  mkdir $(date "+%Y%m%d")
}

function cd-date () {
  cd $(date "+%Y%m%d")
}

# gitコマンド補完
fpath=(
  ${HOME}/.zsh/completions
  ${fpath}
)
autoload -Uz compinit
compinit

alias ls="ls -G"
alias python='/usr/local/bin/python3.11'
alias pip='/usr/local/bin/pip3.11'

source /Users/sho/.credentials.sh

export GOENV_ROOT=$HOME/.goenv
export PATH=$GOENV_ROOT/bin:$PATH
eval "$(goenv init -)"

# remotehostのkuberneteに接続するため
if ping -c 1 -W 1 remotehost > /dev/null
then
  if ! ps aux | grep "6443:kubernetes.docker.internal:6443" | grep -v grep > /dev/null
  then
    ssh -fN -L 6443:kubernetes.docker.internal:6443 remotehost
    ssh remotehost cat .kube/config > /Users/sho/.kube/config
  fi
fi

source <(kubectl completion zsh)

### MANAGED BY RANCHER DESKTOP START (DO NOT EDIT)
export PATH="/Users/sho/.rd/bin:$PATH"
### MANAGED BY RANCHER DESKTOP END (DO NOT EDIT)
