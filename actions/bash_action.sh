#!/bin/sh

#logger zzzz will run sh -c '$ACTION_CMD'
sh -c "$INPUT_CMD"
#logger zzzz EXIT CODE $? sh -c '$ACTION_CMD'
