# HSelect [![Build Status](https://travis-ci.com/gptankit/hselect.svg?branch=master)](https://travis-ci.com/gptankit/hselect) [![PkgGoDev](https://pkg.go.dev/badge/github.com/gptankit/hselect?tab=overview)](https://pkg.go.dev/github.com/gptankit/hselect?tab=overview)

**hselect** is a command line utility for selecting a reliable service among a group of errored and non-errored services (mirrors). It uses *harmonic* dispatch algorithm (https://github.com/gptankit/harmonic) for service selection and also updates *harmonic* with service errors if *dial* fails for the selected service. As *harmonic* can quickly adapt to varying error values, *hselect* can thus select a service with maximum probability of connection success.

Here are the steps to use *hselect* - </br>

**Install**

<pre>
$ mkdir hselect && git clone https://github.com/gptankit/hselect hselect/
$ cd hselect
$ make
$ make install
</pre>

This will install *hselect* binary in */usr/local/hselect/bin* folder.

**Run**

First, add *hselect* command to PATH -

<pre>
$ export PATH=/usr/local/hselect/bin:$PATH
</pre>

Now, run *hselect* with **-e** flag to pass a comma separated list of services - 

<pre>
// Selecting from go module mirrors
$ sudo hselect -e https://goproxy.io,https://proxy.golang.org,https://gocenter.io
https://proxy.golang.org
</pre>

You can optionally run *hselect* with **-v** flag for verbose output - 

<pre>
// Verbose output
$ sudo hselect -v -e https://goproxy.io,https://proxy.golang.org,https://gocenter.io
Parsing endpoints...Done.
Initializing cluster state...Done.
Populating errors...Done.
Selecting service...Done -> https://goproxy.io
Dialing service https://goproxy.io...Success.
Connection successful to https://goproxy.io
https://goproxy.io
</pre>

Output of *hselect* (when used without **-v**) can also be piped to *curl* or *wget* for further downloads - 

<pre>
// Downloading zip file from selected go module mirror
$ sudo hselect -e https://proxy.golang.org/github.com/spf13/viper/@v/v1.7.1.zip,https://goproxy.io/github.com/spf13/viper/@v/v1.7.1.zip,https://gocenter.io/github.com/spf13/viper/@v/v1.7.1.zip | xargs wget -O viper.zip
--2020-08-16 23:50:35--  https://goproxy.io/github.com/spf13/viper/@v/v1.7.1.zip
Resolving goproxy.io (goproxy.io)... 119.28.201.50
Connecting to goproxy.io (goproxy.io)|119.28.201.50|:443... connected.
HTTP request sent, awaiting response... 302 Found
Location: https://goproxy.onetool.net/github.com/spf13/viper/@v/v1.7.1.zip [following]
--2020-08-16 23:50:36--  https://goproxy.onetool.net/github.com/spf13/viper/@v/v1.7.1.zip
Resolving goproxy.onetool.net (goproxy.onetool.net)... 122.10.255.108, 124.156.41.23, 122.10.255.106, ...
Connecting to goproxy.onetool.net (goproxy.onetool.net)|122.10.255.108|:443... connected.
HTTP request sent, awaiting response... 200 OK
Length: 85905 (84K) [application/octet-stream]
Saving to: ‘viper.zip’

viper.zip                                    100%[==========================================================>]  83.89K  --.-KB/s    in 0.03s

2020-08-16 23:50:37 (2.51 MB/s) - ‘viper.zip’ saved [85905/85905]
</pre>

Feel free to play around and post feedbacks
