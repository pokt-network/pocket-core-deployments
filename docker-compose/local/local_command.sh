#!/usr/bin/expect

#set command $argv; # Grab the first command line parameter
set timeout -1

# Import the account into the keybase
spawn pocket accounts import-raw $env(POCKET_CORE_KEY)
sleep 1
send -- "$env(POCKET_CORE_PASSPHRASE)\n"
expect eof

# Set the validator account
spawn sh -c "pocket accounts set-validator $env(POCKET_CORE_ADDRESS)"
sleep 1
send -- "$env(POCKET_CORE_PASSPHRASE)\n"
expect eof

# Start pocket core
spawn sh -c "pocket start --seeds $env(POCKET_CORE_SEEDS)"

expect eof
exit