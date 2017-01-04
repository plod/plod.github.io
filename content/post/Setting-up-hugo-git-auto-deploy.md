+++
Description = ""
Tags = ["go", "golang", "hugo", "digital ocean", "ubuntu", "git"]
date = "2016-02-18T12:29:13Z"
title = "Setting up a hugo based website on ubuntu then auto deploy using gogs' webhooks"
image = "img/bgs/deploy-bg.jpg"

+++
## What is Hugo?

>Hugo is a static HTML and CSS website generator written in Go. It is optimized for speed, easy use and configurability. Hugo takes a directory with content and templates and renders them into a full HTML website.

>Hugo relies on Markdown files with front matter for meta data. And you can run Hugo from any directory. This works well for shared hosts and other systems where you don’t have a privileged account.

>Hugo renders a typical website of moderate size in a fraction of a second. A good rule of thumb is that each piece of content renders in around 1 millisecond.

>Hugo is designed to work well for any kind of website including blogs, tumbles and docs.

https://gohugo.io/

## What is Gogs?

>The goal of this project is to make the easiest, fastest, and most painless way of setting up a self-hosted Git service. With Go, this can be done with an independent binary distribution across ALL platforms that Go supports, including Linux, Mac OS X, Windows and ARM.

https://gogs.io/

It is a great self hosting github/gitlabs sort of project but coded in go.

---

I started by following [this digital ocean tutorial](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-hugo-a-static-site-generator-on-ubuntu-14-04)

I didn't want to use a package, my preference was to compile from [source](https://gohugo.io/overview/installing/). However I got the following error while trying that:
{{< highlight console >}}
go: missing Mercurial command. See https://golang.org/s/gogetcmd
package bitbucket.org/pkg/inflect: exec: "hg": executable file not found in $PATH
{{< / highlight >}}
Thanks to [this](http://kaiq.me/2015/12/23/go/go-get-tools/) and google translate I realised Mercurial was a dependancy.

so
{{< highlight console >}}
sudo apt-get install mercurial
{{< / highlight >}}
and then

{{< highlight console >}}
go get -u -v github.com/spf13/hugo
{{< / highlight >}}

got me cooking!


In the DO tutorial it says to run 

{{< highlight console >}}
sudo hugo genautocomplete
{{< / highlight >}}

this seems to be factored out to 

{{< highlight console >}}
sudo hugo gen autocomplete
{{< / highlight >}}

but because I compiled from source hugo isn't in sudo environment path so 

{{< highlight console >}}
sudo ~/go-workspace/bin/hugo gen autocomplete
{{< / highlight >}}

[~/go-workspace is my $GOPATH so change to yours]


While cloning the themes I decided hugo-themes was a better name for the directory in my home dir (themes could be anything while browsing later), of course this meant later, my symlink would need to look like:

{{< highlight console >}}
 ln -s ../hugo-themes/ themes
{{< / highlight >}}

It seems that some other switches are different from the DO tutorial (-themes= in git cloned version is -t )

Auto deploy
=========

This is based on the digital ocean tutorial [here](https://www.digitalocean.com/community/tutorials/how-to-deploy-a-hugo-site-to-production-with-git-hooks-on-ubuntu-14-04) the difference however was, is I use gogs, the server itself has no firewall access to my gogs server. I decided to solve this problem using the tutorial/gogs' web hooks/and a little php (that I'm sure I'll refactor as go at some stage)

First of all lets set up the gogs web hook:

Inside the repository in question click on the settings link on the right hand side

![Gogs Settings]
(/img/post-img/gogs-settings.png)

Then click on the Webhooks link on the left hand side

![Gogs Webhooks]
(/img/post-img/gogs-webhooks.png)

Finally enter a url the gogs server can connect to and that you can write php on (make a note of the Secret that you use)

![Gogs Webhook]
(/img/post-img/gogs-webhook.png)

Now the php that runs at the URL you entered above should look a little something like this.

{{< highlight php >}}
<?php
//gogs secret for web hook
$secret   = "";

//keyword you are looking for in commit message to decide if to desploy
$deployCommitKeyword = "[deploy] ";
$json = (array) json_decode(file_get_contents('php://input'));

if(array_key_exists('secret', $json)&&($json['secret']==$secret)){
    echo 'secret matched';
    if(array_key_exists('commits', $json)){
        for($i=0, $j=count($json['commits']); $i<$j; $i++){
            $json['commits'][$i] = (array) $json['commits'][$i];
            if(array_key_exists('message', $json['commits'][$i])&&(strstr($json['commits'][$i]['message'], $deployCommitKeyword))){
                do_deploy();
                break;
            }
    }
}

function do_deploy(){
    $GIT_REPO           = "$HOME/my-website.git";
    $WORKING_DIRECTORY  = "$HOME/my-website-working";
    $REMOTE_BACKUP_HTML = "backup_html/";
    $REMOTE_PUBLIC_HTML = "public_html/";
    $MY_SERVER_IP       = "server_domain_or_IP
1";
 
    if(!is_dir($WORKING_DIRECTORY)){
        $command = "git clone $GIT_REPO $WORKING_DIRECTORY";
        echo `$command`;
    }else{
        $command = "cd $WORKING_DIRECTORY; git pull";
        echo `$command`;
    }
    $command = "cd $WORKING_DIRECTORY; /path/to/hugo";
    echo `$command`;
    //lets rsync a copy of the working directory to backup
    $command = "ssh $MY_SERVER_IP rsync -r $REMOTE_PUBLIC_HTML $REMOTE_BACKUP_HTML";
    echo `$command`;
    $command = "rsync -r $WORKING_DIRECTORY/public/ $MY_SERVER_IP:$REMOTE_PUBLIC_HTML";
    echo `$command`;
}
{{< / highlight >}}

_*update_ I noticed that syntax hilighting was going missing on this post (turned out that pygmatize was not in path of the php process) (php runs as my user account) so I fixed this by changing

{{< highlight console >}}
$command = "cd $WORKING_DIRECTORY; /path/to/hugo";
echo `$command`;
{{< / highlight >}}

to

{{< highlight console >}}
$command = "cd $WORKING_DIRECTORY; export PATH=\"What my $PATH is at tty\" hugo";
echo `$command`;
{{< / highlight >}}

This felt a little dirty but after many attempts at a more elegent solution, I have stayed with this for the time being.
