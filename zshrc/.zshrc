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

source .credentials.sh

export GOENV_ROOT=$HOME/.goenv
export PATH=$GOENV_ROOT/bin:$PATH
eval "$(goenv init -)"

# remotehostのkuberneteに接続するため
if ! ps aux | grep "6443:kubernetes.docker.internal:6443" | grep -v grep > /dev/null
then
  ssh -fN -L 6443:kubernetes.docker.internal:6443 remotehost
fi
ssh remotehost cat .kube/config > .kube/config
