#############################################################################

export PS1='\e[95m[DOCKER] \[\033[01;31m\]\u\[\033[01;33m\]@\[\033[01;36m\]\h \[\033[01;33m\]\w \[\033[01;35m\]\$ \[\033[00m\]'
umask 022

alias ls='ls $LS_OPTIONS'
alias ll='ls $LS_OPTIONS -l'
alias l='ls $LS_OPTIONS -lA'
alias ..='cd ..'
alias ...='cd ../..'
alias s='ssh -l root'
alias sshfs='sshfs -o allow_other,uid=1000,gid=1000'

#############################################################################

export EDITOR="vim"

export HISTFILESIZE=99999999
export HISTSIZE=99999999
export HISTCONTROL="ignoreboth"

export LS_OPTIONS='--color=auto -h'

#############################################################################

# ~/emsdk-portable/emsdk activate latest
# source ~/emsdk-portable/emsdk_env.sh