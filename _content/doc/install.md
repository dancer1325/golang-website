* goal
  * download & install Go

* ALTERNATIVES
  * [Managing Go installations](manage-install.md)
  * [Installing Go -- from -- source](install/source.md)

* Go distribution
  * := package / you
    * download
    * install | your system

# Go installation>
## Linux
* TODO:
<ol>
      <li>
        <strong>Remove any previous Go installation</strong> by deleting the /usr/local/go folder
		(if it exists), then extract the archive you just downloaded into /usr/local, creating a fresh
		Go tree in /usr/local/go:
		<pre class="CopyPaste">
      <span>$ rm -rf /usr/local/go &amp;&amp; tar -C /usr/local -xzf <span id="linux-filename">go1.14.3.linux-amd64.tar.gz</span></span>
      <button aria-label="Copy and paste the code.">
        <img class="CopyPaste-icon" src="/images/icons/copy-paste.svg" />
        <img class="CopyPaste-icon CopyPaste-icon-dark" src="/images/icons/copy-paste-dark.svg" />
      </button>
    </pre>
		<p>
          (You may need to run the command as root or through <code>sudo</code>).
        </p>
		<p>
		  <strong>Do not</strong> untar the archive into an existing /usr/local/go tree. This is known to
		  produce broken Go installations.
		 </p>
      </li>
      <li>
        Add /usr/local/go/bin to the <code>PATH</code> environment variable.
        <p>
          You can do this by adding the following line to your $HOME/.profile or
          /etc/profile (for a system-wide installation):
        </p>
        <pre class="CopyPaste">
          <span>export PATH=$PATH:/usr/local/go/bin</span>
          <button aria-label="Copy and paste the code.">
            <img class="CopyPaste-icon" src="/images/icons/copy-paste.svg" />
            <img class="CopyPaste-icon CopyPaste-icon-dark" src="/images/icons/copy-paste-dark.svg" />
          </button>
        </pre>
        <p>
          <strong>Note:</strong> Changes made to a profile file may not apply
          until the next time you log into your computer. To apply the changes
          immediately, just run the shell commands directly or execute them from
          the profile using a command such as
          <code>source $HOME/.profile</code>.
        </p>
      </li>
      <li>
        Verify that you've installed Go by opening a command prompt and typing
        the following command:
        <pre class="CopyPaste">
          <span>$ go version</span>
          <button aria-label="Copy and paste the code.">
            <img class="CopyPaste-icon" src="/images/icons/copy-paste.svg" />
            <img class="CopyPaste-icon CopyPaste-icon-dark" src="/images/icons/copy-paste-dark.svg" />
          </button>
        </pre>
      </li>
      <li>Confirm that the command prints the installed version of Go.</li>
    </ol>
  </div>

## Mac

* download the corresponding [here](https://go.dev/dl/)
* follow the prompts to install
  * Go distribution is installed | 👀"/usr/local/go"👀
  * "/usr/local/go/bin" added | `$PATH` environment variable
* `go version`
  * verify it's installed

## Windows

* TODO:

  <div
    role="tabpanel"
    id="windows-tab"
    class="TabSection-tabPanel"
    aria-labelledby="windows"
    hidden
  >
    <ol>
      <li>
        Open the MSI file you downloaded and follow the prompts to install Go.
        <p>
          By default, the installer will install Go to <code>Program Files</code>
          or <code>Program Files (x86)</code>. You can change the
          location as needed. After installing, you will need to close and
          reopen any open command prompts so that changes to the environment
          made by the installer are reflected at the command prompt.
        </p>
      </li>
      <li>
        Verify that you've installed Go.
        <ol>
          <li>
            In <strong>Windows</strong>, click the <strong>Start</strong> menu.
          </li>
          <li>
            In the menu's search box, type <code>cmd</code>, then press the
            <strong>Enter</strong> key.
          </li>
          <li>
            In the Command Prompt window that appears, type the following
            command:
            <pre class="CopyPaste">
              <span>$ go version</span>
              <button aria-label="Copy and paste the code.">
                <img class="CopyPaste-icon" src="/images/icons/copy-paste.svg" />
                <img class="CopyPaste-icon CopyPaste-icon-dark" src="/images/icons/copy-paste-dark.svg" />
              </button>
            </pre>
          </li>
          <li>Confirm that the command prints the installed version of Go.</li>
        </ol>
      </li>
    </ol>
  </div>
</div>

