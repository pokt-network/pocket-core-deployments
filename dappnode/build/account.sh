#!/usr/bin/expect
    
if { "$env(ACCOUNT)" == "Shielded Key" } {
    log_user 0
    spawn pocket accounts import-armored /tmp/shielded_key.json
    sleep 1
    send -- "$env(SHIELDED_PASSPHRASE)\n"
    sleep 1
    send -- "$env(SHIELDED_PASSPHRASE)\n"
    expect eof
    spawn sh -c "pocket accounts set-validator `pocket accounts list | cut -d' ' -f2- `"
    sleep 1
    send -- "$env(SHIELDED_PASSPHRASE)\n"

} else {

    log_user 0
    spawn sh -c "cp /tmp/*.json /home/app/.pocket/config/"
    spawn pocket accounts import-raw $env(POCKET_CORE_KEY)
    sleep 1
    send -- "$env(POCKET_CORE_PASSPHRASE)\n"
    expect eof
    spawn sh -c "pocket accounts set-validator `pocket accounts list | cut -d' ' -f2- `"
    sleep 1
    send -- "$env(POCKET_CORE_PASSPHRASE)\n"
} 

expect eof

exit
