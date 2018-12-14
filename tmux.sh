#!/bin/sh
tmux new-session './log.sh' \; \
    split-window -h './monitor-kr.sh' \; \
    split-window -v './run.sh' \; \
    select-pane -t 0 \; \
    split-window -v  \; \
    split-window -h './ip-monitor.sh' \;


