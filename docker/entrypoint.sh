#!/usr/bin/expect
# parse arguments to get --passphrase flag
set passphrase_idx [lsearch -nocase -exact $argv "--passphrase"]
if { $passphrase_idx > 0} {
    set passphrase [lindex $argv [expr {$passphrase_idx +1}]]
    set user_command [lrange $argv 0 [expr {$passphrase_idx -1}]]
} elseif {$passphrase_idx <= 0} {
    set user_command $argv 
}

# Checks if POCKET_CORE_KEY environment variable is set or empty
if { [info exists env(POCKET_CORE_KEY2)] && $env(POCKET_CORE_KEY2) ne "" }  {
    log_user 0
    spawn sh -c "pocket-core accounts import-raw $env(POCKET_CORE_KEY)"
    log_user 1
    sleep 1

# Checks if the passphrase was passed
    if { [info exists passphrase] } { 
        send -- "$passphrase\n"
    } else { 
        send_user "Please use the --passphrase flag to set the passphrase\n"
        exit
    }

    expect eof
    
# Checks if a command was passed or if it is empty
    if { [info exist user_command] && $user_command ne ""} { 
        log_user 0
        spawn sh -c $user_command;
        log_user 1
    } else { 
        send_user "\nPlease enter a command\n"
        exit
    }

} else {
    if { [info exist user_command] && $user_command ne ""} { 
        log_user 0
        spawn sh -c $user_command;
        log_user 1
    } else { 
        send_user "\nPlease enter a command\n"
        exit
    }
}

sleep 1

if {[info exist passphrase]} { 
    send -- "$passphrase\n"
}

expect eof

exit