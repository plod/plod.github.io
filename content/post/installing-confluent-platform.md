---
title: "Installing Confluent Platform"
date: 2018-03-03T13:57:02Z
draft: false
image: "img/bgs/franz-kafka.jpg"
tags: ["kafka", "big data", "java", "scala", "confluent", "script"]
description: "with java 9 on ubuntu 16.04"
---

I recently installed the newest version of confluent platform (4.0 at time of writing) on ubuntu 16.04 leaving these steps here for anyone else who needs to automate it.

{{< highlight console >}}
sudo add-apt-repository ppa:webupd8team/java -y
sudo apt-get update
echo oracle-java9-installer shared/accepted-oracle-license-v1-1 select true | sudo /usr/bin/debconf-set-selections
sudo apt install oracle-java9-installer oracle-java9-set-default -y
wget -qO - https://packages.confluent.io/deb/4.0/archive.key | sudo apt-key add -
sudo add-apt-repository "deb [arch=amd64] https://packages.confluent.io/deb/4.0 stable main"
sudo apt-get update
sudo apt-get install confluent-platform-oss-2.11 -y
{{< / highlight >}}

Points of note

- This uses java9, there are different things needed for java8
- you can use different apt repos for earlier (or post) version of confluent platform
- I'm using the scala 2.11 version you could change this to 2.10 in the package name
- above is for the oss package check confluent website for enterprise package name