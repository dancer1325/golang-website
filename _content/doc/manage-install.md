* goal  
  * how to
    * install MULTIPLE versions of Go | SAME machine
    * uninstall Go

# Installing MULTIPLE Go versions | SAME machine

* use cases
  * test your code | MULTIPLE Go versions

* requirements
  * [git](https://git-scm.com/)

* steps
  * [install a Go version](install.md)
  * `go install [build flags] [packages]` & `goSpecifiedVersion download`
    * install ADDITIONAL Go versions
    * run the <a href="/cmd/go/#hdr-Compile_and_install_packages_and_dependencies"><code>go install</code> command</a>, specifying the download location of the version you want to install
    ```
    $ go install golang.org/dl/go1.24.4@latest    # NOT require to specify architecture == FINE -- via -- this command
    $ go1.24.4 download
    ```

* if you want to
  * run commands / ADDITIONAL version -> `goSpecifiedVersion commandDesired`
    ```
    $ go1.24.4 version
    ```
  * check where EACH version is installed -> `goSpecifiedVersion env GOROOT`
    ```
    $ go1.24.4 env GOROOT
    ```

# Uninstalling Go

## Linux / macOS / FreeBSD

* TODO:
<ol>

<li>Delete the go directory.

<p>
This is usually /usr/local/go.
</p>

</li>

<li>Remove the Go bin directory from your <code>PATH</code> environment variable.

<p>
Under Linux and FreeBSD, edit /etc/profile or $HOME/.profile
* If you installed Go with the macOS package, remove the /etc/paths.d/go file.
</p>

</li>

</ol>

## Windows

<p>
The simplest way to remove Go is via Add/Remove Programs in the Windows control panel:
</p>

<ol>

<li>In Control Panel, double-click <strong>Add/Remove Programs</strong>.</li>
<li>In <strong>Add/Remove Programs</strong>, select <strong>Go Programming Language,</strong> click Uninstall, then follow the prompts.</li>

</ol>

<p>
For removing Go with tools, you can also use the command line:
</p>

<ul>

<li>Uninstall using the command line by running the following command:

<pre>
msiexec /x go{{version}}.windows-{{cpu-arch}}.msi /q
</pre>

<p>

<strong>Note:</strong> Using this uninstall process for Windows will automatically remove Windows environment variables created by the original installation.
</p>

</li>

</ul>

