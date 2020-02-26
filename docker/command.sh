#!/usr/bin/expect

set command $argv; # Grab the first command line parameter
set timeout -1

if { $env(POCKET_CORE_KEY) eq "" }  {
    spawn sh -c "$command"

} else {
    spawn pocket-core accounts import-raw $env(POCKET_CORE_KEY)
    sleep 1
    send -- "$env(POCKET_CORE_PASSPHRASE)\n"
    expect eof
    spawn sh -c "$command"
}

sleep 1

send -- "$env(POCKET_CORE_PASSPHRASE)\n"

expect eof

exit
