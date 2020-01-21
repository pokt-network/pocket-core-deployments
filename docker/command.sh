#!/usr/bin/expect

set timeout -1

if { $env(POCKET_CORE_KEY) eq "" }  {
    spawn pocket-core start
} else {
    spawn pocket-core accounts import-raw $env(POCKET_CORE_KEY)
    sleep 1
    send -- "yo\n"
    expect eof
    spawn pocket-core start --persistent_peers $env(POCKET_CORE_PERSISTENT_PEERS)
}

sleep 1

send -- "yo\n"

expect eof

exit
