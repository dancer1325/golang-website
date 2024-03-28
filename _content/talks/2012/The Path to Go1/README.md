# Go
* := Concurrent Open Source language
  * == native compilation + static types + lightweight dynamic feel (js or python)
* goal
  * 👁️write software at Google for the big machines in the back office 👁️
    * **Note:** Originally all was C++ & Java
* features
  * native code generation
    * == compiler — generates — native code
      * **Note:** 👁️ there are 2 compiler suites 👁️
  * statically typed
  * composition via interfaces
    * — to join — things
    * contrary to standard Object Oriented Model (❓)
  * memory safe
  * garbage collected
    * — opposed to — manage your own memory
  * standard library
    * based on the principle 👁️design before code 👁️

# Go1
* 1@ release
* development process
  * possible version control systems in those days
    * mercurial — chosen one —
    * subversion
  * linear history without branches
    * **Reason:** 🧠 all committers are updated to the current state 🧠
* contributions
  * originally launched on Linux & Mac
* development cycle
  * bit chaotic
  * approaches used
    * weekly snapshots
      * steps
        * pick an earlier stable version confirmed by the builders
        * tag in mercurial
      * has also problems
    * formal release process
* TODO: Keep on 'What is Go 1'

# Gofix
* := tool which
  * code — is mechanically — updated to new
    * language &
    * library
* how does it work?
  * Go code — is parsed to → tree
  * tree — based on various plug-in modules is → rewritten
* NO panacea
  * == certain versioning issues persist
    * 👁→ some companies perceived as unstable language 👁️
