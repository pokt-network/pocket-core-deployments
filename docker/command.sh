#!/usr/bin/expect

set command $argv; # Grab the first command line parameter
set timeout -1

spawn sh -c "cp /tmp/*.json /home/app/.pocket/config/"
sleep 2
if { $env(POCKET_CORE_KEY) eq "" }  {
    log_user 0
    spawn sh -c "$command"
    send -- "$env(POCKET_CORE_PASSPHRASE)\n"
    log_user 1

} else {
    log_user 0
    spawn pocket accounts import-raw $env(POCKET_CORE_KEY)
    sleep 1
    send -- "$env(POCKET_CORE_PASSPHRASE)\n"
    expect eof
    spawn sh -c "pocket accounts set-validator `pocket accounts list | cut -d' ' -f2- `"
    sleep 1
    send -- "$env(POCKET_CORE_PASSPHRASE)\n"
    expect eof
    log_user 1
    spawn sh -c "$command"
}

expect eof

exit
