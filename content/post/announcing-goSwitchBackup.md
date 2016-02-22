+++
categories = ["Development", "golang"]
date = "2016-02-22T16:19:11Z"
description = ""
draft = false
image = "/img/bgs/bg-switch.jpg"
tags = ["go", "golang", "development", "switches", "ssh", "cisco"]
title = "Announcing goSwitchBackup"

+++

goSwitchBackup is Go program I wrote, to automate the backing up of cisco switches over SSH for the network boys in work. It essentially SSHs (can you use it like a verb like that?) into an IP (Which should be the address of a cisco switch) using the password, once it receives the prompt character it trys to esculate privilidges using: {{< highlight console >}}enable{{< /highlight >}} Onnce it sees the "#" character it then trys to run the backup command: {{< highlight console >}}copy running-config tftp{{< /highlight >}}It will then give the tftp address used when the program was invoked, and call the file switchBackup-IP-OF-SWITCH-RFC3339timestamp (_I may make this configurable later_)

You can check out the code on github here:

https://github.com/plod/goSwitchBackup

Once you run or compile, or just go run the code, pass the -help command for helps with flags.

Any comments or requests for more information please post in the comments below (I am still quite new to Go so any helpful pointers would also be greatfully received).

### todo/wishlist:

* Add potential tftp server inside the program itself
* compare previous backup with current backup and send alert of changes (smtp)
* Add a verbose switch to echo some of the commented stuff that helps debug issues