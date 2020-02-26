#!/usr/bin/expect

proc parseargs {argc argv} {
    global OPTS
    foreach {key val} $argv {
        switch -exact -- $key {
            "-c"   { set OPTS(command)   $val }
            "-p"   { set OPTS(pass)   $val }
            "-h"   { set OPTS(help) $val}
        }
    }
}

parseargs $argc $argv

# Print to console script help
if { [info exists OPTS(help)] } {
    send_user " -c\t command\t (must be between quotes) \"<command here>\" \n -p\t passphrase \n -h\t this help \n"
    exit
}

# Checks if POCKET_CORE_KEY environment variable is set or empty
if { [info exists env(POCKET_CORE_KEY)] && $env(POCKET_CORE_KEY) ne "" }  {
    log_user 0
    spawn sh -c "pocket-core accounts import-raw $env(POCKET_CORE_KEY)"
    log_user 1
    sleep 1

# Checks if the passphrase was passed
    if { [info exists OPTS(pass)] } { 
        send -- "$OPTS(pass)\n"
    } else { 
        send_user "Please use the -p flag to set the passphrase\n"
        exit
    }

    expect eof
    
# Checks if a command was passed or if it is empty
    if { [info exist OPTS(command)] && $OPTS(command) ne ""} { 
        log_user 0
        spawn sh -c $OPTS(command);
        log_user 1
    } else { 
        send_user "\n\nPlease use the -c flag to enter your commands (between quotes) \"<command here>\"\n"
        exit
    }

} else {
    if { [info exist OPTS(command)] && $OPTS(command) ne ""} { 
        log_user 0
        spawn sh -c $OPTS(command);
        log_user 1
    } else { 
        send_user "\n\nPlease use the -c flag to enter your commands (between quotes) \"<command here>\"\n"
        exit
    }
}

sleep 1

if {[info exist OPTS(pass)]} { 
    send -- "$OPTS(pass)\n"
}

expect eof

exit