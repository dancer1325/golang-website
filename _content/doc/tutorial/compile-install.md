* goal
  * "Create a Go module" tutorial LAST part
  * 2 OTHER `go` commands / build code
    * [`go build`](https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies)
    * [`go install`](/golang-website/_content/ref/mod.md#go-install-go-install)

* `go run` command
  * uses
    * | program / you make frequent changes,
      * compile it
      * run it 
  * ❌NOT generate a binary executable❌

* `go build`
  * compiles the packages + their dependencies
  * ❌NOT install the compiled results❌

* `go install`
  * compiles & installs the packages

* TODO: 
<ol>
  <li>
    From the command line in the hello directory, run the <code>go build</code>
    command to compile the code into an executable.

    <pre>$ go build</pre>
  </li>

  <li>
    From the command line in the hello directory, run the new <code>hello</code>
    executable to confirm that the code works.

    <p>
      Note that your result might differ depending on whether you changed
      your greetings.go code after testing it.
    </p>

    <ul>
      <li>
        On Linux or Mac:

        <pre>
$ ./hello
map[Darrin:Great to see you, Darrin! Gladys:Hail, Gladys! Well met! Samantha:Hail, Samantha! Well met!]
</pre>
      </li>

      <li>
        On Windows:

        <pre>
$ hello.exe
map[Darrin:Great to see you, Darrin! Gladys:Hail, Gladys! Well met! Samantha:Hail, Samantha! Well met!]
</pre>
      </li>
    </ul>
    <p>
      You've compiled the application into an executable so you can run it.
      But to run it currently, your prompt needs either to be in the executable's
      directory, or to specify the executable's path.
    </p>
    <p>
      Next, you'll install the executable so you can run it without specifying
      its path.
    </p>
  </li>

  <li>
    Discover the Go install path, where the <code>go</code> command will install
    the current package.

    <p>
      You can discover the install path by running the
      <a href="/cmd/go/#hdr-List_packages_or_modules">
        <code>go list</code> command</a>, as in the following example:
    </p>

    <pre>
$ go list -f '{{.Target}}'
</pre>

    <p>
      For example, the command's output might say <code>/home/gopher/bin/hello</code>,
      meaning that binaries are installed to /home/gopher/bin. You'll need this
      install directory in the next step.
    </p>
  </li>

  <li>
    Add the Go install directory to your system's shell path.

    <p>
      That way, you'll be able to run your program's executable without
      specifying where the executable is.
    </p>

    <ul>
      <li>
        On Linux or Mac, run the following command:

        <pre>
$ export PATH=$PATH:/path/to/your/install/directory
</pre
        >
      </li>

      <li>
        On Windows, run the following command:

        <pre>
$ set PATH=%PATH%;C:\path\to\your\install\directory
</pre
        >
      </li>
    </ul>

    <p>
      As an alternative, if you already have a directory like
      <code>$HOME/bin</code> in your shell path and you'd like to install your
      Go programs there, you can change the install target by setting the
      <code>GOBIN</code> variable using the
      <a href="/cmd/go/#hdr-Print_Go_environment_information">
        <code>go env</code> command</a>:
    </p>

    <pre>
$ go env -w GOBIN=/path/to/your/bin
</pre>

    <p>
      or
    </p>

    <pre>
$ go env -w GOBIN=C:\path\to\your\bin
</pre
    >
  </li>

  <li>
    Once you've updated the shell path, run the <code>go install</code> command
    to compile and install the package.

    <pre>$ go install</pre>
  </li>

  <li>
    Run your application by simply typing its name. To make this interesting,
    open a new command prompt and run the <code>hello</code> executable name
    in some other directory.

    <pre>
$ hello
map[Darrin:Hail, Darrin! Well met! Gladys:Great to see you, Gladys! Samantha:Hail, Samantha! Well met!]
</pre
    >
  </li>
</ol>

<p>
  That wraps up this Go tutorial!
</p>

<p class="Navigation">
  <a class="Navigation-prev" href="add-a-test.html">&lt; Add a test</a>
  <a class="Navigation-next" href="module-conclusion.html">Conclusion and links
    to more information &gt;</a>
</p>
