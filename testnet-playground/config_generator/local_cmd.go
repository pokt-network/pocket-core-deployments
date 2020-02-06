package main

import "io/ioutil"

const (
	fileContents = `#!/usr/bin/expect

set timeout -1

if { $env(POCKET_CORE_KEY) eq "" }  {
    spawn pocket-core start
} else {
    spawn pocket-core accounts import-raw $env(POCKET_CORE_KEY)
    sleep 1
    send -- "yo\n"
    expect eof
    spawn pocket-core start --seeds $env(POCKET_CORE_SEEDS)
}

sleep 1

send -- "yo\n"

expect eof

exit`
)

func writeLocalCmd(homeDir string) {
	er := ioutil.WriteFile(homeDir+fs+"local_command.sh", []byte(fileContents), 0644)
	if er != nil {
		panic(er)
	}
}
