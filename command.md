basic command for BYP
// The Core Flaw: System-Wide vs. User-Space
When an administrator restricts "external downloads" for a non-root user, they usually do so by locking down the system tools (apt, snap, or blocking wget/curl). However, standard Linux architectures inherently allow users complete control over their own home directory

** command to Test Network Fetching (User-Space Network Access)
See if your non-sudo user is blocked from making outbound web requests entirely, or if only the installation tools are blocked. Try to fetch a benign webpage header using standard utilities:
1.  curl -I https://www.google.com

The Vulnerability: Any application running under your user account can pull down files from the internet using raw network sockets, curl, wget, or Python scripts. The operating system is not filtering outbound web traffic.

** Purge the Broken HTML File
2. rm tor-browser-linux-x86_64-13.5.1.tar.xz

3. curl -L -O https://www.torproject.org/dist/torbrowser/15.0.15/tor-browser-linux-x86_64-15.0.15.tar.xz

**. Extract the Genuine Archive
4. tar -xvJf tor-browser-linux-x86_64-15.0.15.tar.xz

5. Run the Browser Out of User Space
Navigate into the newly extracted directory and launch the browser application binary directly:

Bash
cd tor-browser
./start-tor-browser.desktop


encription..................
annother way to achieve this 

1. wget https://www.torproject.org/dist/torbrowser/14.5.7/tor-browser-linux-x86_64-14.5.7.tar.xz
// to check the sytem cmpt version
2.uname -m
3. wget https://dist.torproject.org/torbrowser/15.0.15/tor-browser-linux-x86_64-15.0.15.tar.xz
// extract the file from the achieve
5. tar -xf tor-browser-linux-x86_64-15.0.15.tar.xz
    // if the extraction fail or If that gives a permission error:
6.chmod +x start-tor-browser.desktop
7. cd tor-browser
8. ./start-tor-browser.desktop
