<!--{
	"Title": "Effective Go",
	"Template": true,
	"Breadcrumb": true
}-->

# Introduction

* goal
  * how to write Go code
    * clear
    * idiomatic 

* recommendations
  * read PREVIOUSLY
    * [Go language specification](https://go.dev/ref/spec)
    * [Tour of Go](../tour)
    * [How to Write Go Code?](HowToWriteGoCode)

* Go
  * NEW language
  * 👀's some ideas borrowed -- from -- EXISTING languages👀
  * recommendations
    * 👀understand its 
      * properties
      * idioms
      * conventions (naming, formatting, program construction, ...)👀

* | January, 2022
  * this document was written | Go's release, in 2009
  * NOT updated significantly
  * this guide
    * NOT useful for
      * libraries
      * Go ecosystem 
    * useful for
      * how to use the language 
  * see [28782](https://github.com/golang/go/issues/28782)

## Examples

* [Go package sources](https://go.dev/src/)
  * == Go core library
  * uses
    * how to use the language
      * _Example:_ [map](https://pkg.go.dev/strings#example-Map)

# Formatting

* goal
  * 1! style

* formatting approach
  * 👀machine take care of MOST formatting issues👀

* `gofmt` OR `go fmt`
  * == program / operates | package level
    * steps
      * reads a Go program
      * emits the source / 
        * standard 
          * style of indentation
            * -- based on -- tabs
            * ONLY if you must -> use spaces
          * vertical alignment
        * 👀if it's necessary -> reformat comments👀
        * ❌NO line length limits❌
          * if it's too long -> wrap it & indent -- with an -- extra tab
        * ❌NO need `()` ❌
          * _Example:_ control structures (`if`, `for`, `switch`)
          * Reason: 🧠use operator precendence🧠
          ```
          x<<8 + y<<16
          # == (x<<8) + (y<<16)
          ```

* _Example:_
    ```go
    type T struct {
        name string // name of the object
        value int // its value
    }
    ```
  `gofmt` line up the columns

    ```go
    type T struct {
        name    string // name of the object
        value   int    // its value
    }
    ```

# Commentary

<p>
Go provides C-style <code>/* */</code> block comments
and C++-style <code>//</code> line comments.
Line comments are the norm;
block comments appear mostly as package comments, but
are useful within an expression or to disable large swaths of code.
</p>

<p>
Comments that appear before top-level declarations, with no intervening newlines,
are considered to document the declaration itself.
These “doc comments” are the primary documentation for a given Go package or command.
For more about doc comments, see “<a href="/doc/comment">Go Doc Comments</a>”.
</p>

# Names

<p>
Names are as important in Go as in any other language.
They even have semantic effect:
the visibility of a name outside a package is determined by whether its
first character is upper case.
It's therefore worth spending a little time talking about naming conventions
in Go programs.
</p>


<h3 id="package-names">Package names</h3>

<p>
When a package is imported, the package name becomes an accessor for the
contents.  After
</p>

<pre>
import "bytes"
</pre>

<p>
the importing package can talk about <code>bytes.Buffer</code>.  It's
helpful if everyone using the package can use the same name to refer to
its contents, which implies that the package name should be good:
short, concise, evocative.  By convention, packages are given
lower case, single-word names; there should be no need for underscores
or mixedCaps.
Err on the side of brevity, since everyone using your
package will be typing that name.
And don't worry about collisions <i>a priori</i>.
The package name is only the default name for imports; it need not be unique
across all source code, and in the rare case of a collision the
importing package can choose a different name to use locally.
In any case, confusion is rare because the file name in the import
determines just which package is being used.
</p>

<p>
Another convention is that the package name is the base name of
its source directory;
the package in <code>src/encoding/base64</code>
is imported as <code>"encoding/base64"</code> but has name <code>base64</code>,
not <code>encoding_base64</code> and not <code>encodingBase64</code>.
</p>

<p>
The importer of a package will use the name to refer to its contents,
so exported names in the package can use that fact
to avoid repetition.
(Don't use the <code>import .</code> notation, which can simplify
tests that must run outside the package they are testing, but should otherwise be avoided.)
For instance, the buffered reader type in the <code>bufio</code> package is called <code>Reader</code>,
not <code>BufReader</code>, because users see it as <code>bufio.Reader</code>,
which is a clear, concise name.
Moreover,
because imported entities are always addressed with their package name, <code>bufio.Reader</code>
does not conflict with <code>io.Reader</code>.
Similarly, the function to make new instances of <code>ring.Ring</code>&mdash;which
is the definition of a <em>constructor</em> in Go&mdash;would
normally be called <code>NewRing</code>, but since
<code>Ring</code> is the only type exported by the package, and since the
package is called <code>ring</code>, it's called just <code>New</code>,
which clients of the package see as <code>ring.New</code>.
Use the package structure to help you choose good names.
</p>

<p>
Another short example is <code>once.Do</code>;
<code>once.Do(setup)</code> reads well and would not be improved by
writing <code>once.DoOrWaitUntilDone(setup)</code>.
Long names don't automatically make things more readable.
A helpful doc comment can often be more valuable than an extra long name.
</p>

<h3 id="Getters">Getters</h3>

<p>
Go doesn't provide automatic support for getters and setters.
There's nothing wrong with providing getters and setters yourself,
and it's often appropriate to do so, but it's neither idiomatic nor necessary
to put <code>Get</code> into the getter's name.  If you have a field called
<code>owner</code> (lower case, unexported), the getter method should be
called <code>Owner</code> (upper case, exported), not <code>GetOwner</code>.
The use of upper-case names for export provides the hook to discriminate
the field from the method.
A setter function, if needed, will likely be called <code>SetOwner</code>.
Both names read well in practice:
</p>
<pre>
owner := obj.Owner()
if owner != user {
    obj.SetOwner(user)
}
</pre>

## Interface names

* 1-method interfaces' naming
  * == "methodName" + "-er"
    * Reason: 🧠construct an agent noun🧠
    * _Example:_ Reader, Writer

* "Read", "Write", "Close", "Flush", "String", ... 
  * canonical signatures & meanings
  * use properly

* if your type implements a method / meaning == well-known type's method -> give SAME name & signature
  * _Example:_ string-converter method would be "String", NOT "ToString"

## MixedCaps

<p>
Finally, the convention in Go is to use <code>MixedCaps</code>
or <code>mixedCaps</code> rather than underscores to write
multiword names.
</p>

# Semicolons

<p>
Like C, Go's formal grammar uses semicolons to terminate statements,
but unlike in C, those semicolons do not appear in the source.
Instead the lexer uses a simple rule to insert semicolons automatically
as it scans, so the input text is mostly free of them.
</p>

<p>
The rule is this. If the last token before a newline is an identifier
(which includes words like <code>int</code> and <code>float64</code>),
a basic literal such as a number or string constant, or one of the
tokens
</p>
<pre>
break continue fallthrough return ++ -- ) }
</pre>
<p>
the lexer always inserts a semicolon after the token.
This could be summarized as, &ldquo;if the newline comes
after a token that could end a statement, insert a semicolon&rdquo;.
</p>

<p>
A semicolon can also be omitted immediately before a closing brace,
so a statement such as
</p>
<pre>
    go func() { for { dst &lt;- &lt;-src } }()
</pre>
<p>
needs no semicolons.
Idiomatic Go programs have semicolons only in places such as
<code>for</code> loop clauses, to separate the initializer, condition, and
continuation elements.  They are also necessary to separate multiple
statements on a line, should you write code that way.
</p>

<p>
One consequence of the semicolon insertion rules
is that you cannot put the opening brace of a
control structure (<code>if</code>, <code>for</code>, <code>switch</code>,
or <code>select</code>) on the next line.  If you do, a semicolon
will be inserted before the brace, which could cause unwanted
effects.  Write them like this
</p>

<pre>
if i &lt; f() {
    g()
}
</pre>
<p>
not like this
</p>
<pre>
if i &lt; f()  // wrong!
{           // wrong!
    g()
}
</pre>


# Control structures

<p>
The control structures of Go are related to those of C but differ
in important ways.
There is no <code>do</code> or <code>while</code> loop, only a
slightly generalized
<code>for</code>;
<code>switch</code> is more flexible;
<code>if</code> and <code>switch</code> accept an optional
initialization statement like that of <code>for</code>;
<code>break</code> and <code>continue</code> statements
take an optional label to identify what to break or continue;
and there are new control structures including a type switch and a
multiway communications multiplexer, <code>select</code>.
The syntax is also slightly different:
there are no parentheses
and the bodies must always be brace-delimited.
</p>

## `if`

<p>
In Go a simple <code>if</code> looks like this:
</p>
<pre>
if x &gt; 0 {
    return y
}
</pre>

<p>
Mandatory braces encourage writing simple <code>if</code> statements
on multiple lines.  It's good style to do so anyway,
especially when the body contains a control statement such as a
<code>return</code> or <code>break</code>.
</p>

<p>
Since <code>if</code> and <code>switch</code> accept an initialization
statement, it's common to see one used to set up a local variable.
</p>

<pre>
if err := file.Chmod(0664); err != nil {
    log.Print(err)
    return err
}
</pre>

<p id="else">
In the Go libraries, you'll find that
when an <code>if</code> statement doesn't flow into the next statement—that is,
the body ends in <code>break</code>, <code>continue</code>,
<code>goto</code>, or <code>return</code>—the unnecessary
<code>else</code> is omitted.
</p>

<pre>
f, err := os.Open(name)
if err != nil {
    return err
}
codeUsing(f)
</pre>

<p>
This is an example of a common situation where code must guard against a
sequence of error conditions.  The code reads well if the
successful flow of control runs down the page, eliminating error cases
as they arise.  Since error cases tend to end in <code>return</code>
statements, the resulting code needs no <code>else</code> statements.
</p>

<pre>
f, err := os.Open(name)
if err != nil {
    return err
}
d, err := f.Stat()
if err != nil {
    f.Close()
    return err
}
codeUsing(f, d)
</pre>


## Redeclaration and reassignment

<p>
An aside: The last example in the previous section demonstrates a detail of how the
<code>:=</code> short declaration form works.
The declaration that calls <code>os.Open</code> reads,
</p>

<pre>
f, err := os.Open(name)
</pre>

<p>
This statement declares two variables, <code>f</code> and <code>err</code>.
A few lines later, the call to <code>f.Stat</code> reads,
</p>

<pre>
d, err := f.Stat()
</pre>

<p>
which looks as if it declares <code>d</code> and <code>err</code>.
Notice, though, that <code>err</code> appears in both statements.
This duplication is legal: <code>err</code> is declared by the first statement,
but only <em>re-assigned</em> in the second.
This means that the call to <code>f.Stat</code> uses the existing
<code>err</code> variable declared above, and just gives it a new value.
</p>

<p>
In a <code>:=</code> declaration a variable <code>v</code> may appear even
if it has already been declared, provided:
</p>

<ul>
<li>this declaration is in the same scope as the existing declaration of <code>v</code>
(if <code>v</code> is already declared in an outer scope, the declaration will create a new variable §),</li>
<li>the corresponding value in the initialization is assignable to <code>v</code>, and</li>
<li>there is at least one other variable that is created by the declaration.</li>
</ul>

<p>
This unusual property is pure pragmatism,
making it easy to use a single <code>err</code> value, for example,
in a long <code>if-else</code> chain.
You'll see it used often.
</p>

<p>
§ It's worth noting here that in Go the scope of function parameters and return values
is the same as the function body, even though they appear lexically outside the braces
that enclose the body.
</p>

## `for`

<p>
The Go <code>for</code> loop is similar to&mdash;but not the same as&mdash;C's.
It unifies <code>for</code>
and <code>while</code> and there is no <code>do-while</code>.
There are three forms, only one of which has semicolons.
</p>
<pre>
// Like a C for
for init; condition; post { }

// Like a C while
for condition { }

// Like a C for(;;)
for { }
</pre>

<p>
Short declarations make it easy to declare the index variable right in the loop.
</p>
<pre>
sum := 0
for i := 0; i &lt; 10; i++ {
    sum += i
}
</pre>

<p>
If you're looping over an array, slice, string, or map,
or reading from a channel, a <code>range</code> clause can
manage the loop.
</p>
<pre>
for key, value := range oldMap {
    newMap[key] = value
}
</pre>

<p>
If you only need the first item in the range (the key or index), drop the second:
</p>
<pre>
for key := range m {
    if key.expired() {
        delete(m, key)
    }
}
</pre>

<p>
If you only need the second item in the range (the value), use the <em>blank identifier</em>, an underscore, to discard the first:
</p>
<pre>
sum := 0
for _, value := range array {
    sum += value
}
</pre>

<p>
The blank identifier has many uses, as described in <a href="#blank">a later section</a>.
</p>

<p>
For strings, the <code>range</code> does more work for you, breaking out individual
Unicode code points by parsing the UTF-8.
Erroneous encodings consume one byte and produce the
replacement rune U+FFFD.
(The name (with associated builtin type) <code>rune</code> is Go terminology for a
single Unicode code point.
See <a href="/ref/spec#Rune_literals">the language specification</a>
for details.)
The loop
</p>
<pre>
for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
    fmt.Printf("character %#U starts at byte position %d\n", char, pos)
}
</pre>
<p>
prints
</p>
<pre>
character U+65E5 '日' starts at byte position 0
character U+672C '本' starts at byte position 3
character U+FFFD '�' starts at byte position 6
character U+8A9E '語' starts at byte position 7
</pre>

<p>
Finally, Go has no comma operator and <code>++</code> and <code>--</code>
are statements not expressions.
Thus if you want to run multiple variables in a <code>for</code>
you should use parallel assignment (although that precludes <code>++</code> and <code>--</code>).
</p>
<pre>
// Reverse a
for i, j := 0, len(a)-1; i &lt; j; i, j = i+1, j-1 {
    a[i], a[j] = a[j], a[i]
}
</pre>

## `switch`

<p>
Go's <code>switch</code> is more general than C's.
The expressions need not be constants or even integers,
the cases are evaluated top to bottom until a match is found,
and if the <code>switch</code> has no expression it switches on
<code>true</code>.
It's therefore possible&mdash;and idiomatic&mdash;to write an
<code>if</code>-<code>else</code>-<code>if</code>-<code>else</code>
chain as a <code>switch</code>.
</p>

<pre>
func unhex(c byte) byte {
    switch {
    case '0' &lt;= c &amp;&amp; c &lt;= '9':
        return c - '0'
    case 'a' &lt;= c &amp;&amp; c &lt;= 'f':
        return c - 'a' + 10
    case 'A' &lt;= c &amp;&amp; c &lt;= 'F':
        return c - 'A' + 10
    }
    return 0
}
</pre>

<p>
There is no automatic fall through, but cases can be presented
in comma-separated lists.
</p>
<pre>
func shouldEscape(c byte) bool {
    switch c {
    case ' ', '?', '&amp;', '=', '#', '+', '%':
        return true
    }
    return false
}
</pre>

<p>
Although they are not nearly as common in Go as some other C-like
languages, <code>break</code> statements can be used to terminate
a <code>switch</code> early.
Sometimes, though, it's necessary to break out of a surrounding loop,
not the switch, and in Go that can be accomplished by putting a label
on the loop and "breaking" to that label.
This example shows both uses.
</p>

<pre>
Loop:
    for n := 0; n &lt; len(src); n += size {
        switch {
        case src[n] &lt; sizeOne:
            if validateOnly {
                break
            }
            size = 1
            update(src[n])

        case src[n] &lt; sizeTwo:
            if n+1 &gt;= len(src) {
                err = errShortInput
                break Loop
            }
            if validateOnly {
                break
            }
            size = 2
            update(src[n] + src[n+1]&lt;&lt;shift)
        }
    }
</pre>

<p>
Of course, the <code>continue</code> statement also accepts an optional label
but it applies only to loops.
</p>

<p>
To close this section, here's a comparison routine for byte slices that uses two
<code>switch</code> statements:
</p>
<pre>
// Compare returns an integer comparing the two byte slices,
// lexicographically.
// The result will be 0 if a == b, -1 if a &lt; b, and +1 if a &gt; b
func Compare(a, b []byte) int {
    for i := 0; i &lt; len(a) &amp;&amp; i &lt; len(b); i++ {
        switch {
        case a[i] &gt; b[i]:
            return 1
        case a[i] &lt; b[i]:
            return -1
        }
    }
    switch {
    case len(a) &gt; len(b):
        return 1
    case len(a) &lt; len(b):
        return -1
    }
    return 0
}
</pre>

## Type switch

<p>
A switch can also be used to discover the dynamic type of an interface
variable.  Such a <em>type switch</em> uses the syntax of a type
assertion with the keyword <code>type</code> inside the parentheses.
If the switch declares a variable in the expression, the variable will
have the corresponding type in each clause.
It's also idiomatic to reuse the name in such cases, in effect declaring
a new variable with the same name but a different type in each case.
</p>
<pre>
var t interface{}
t = functionOfSomeType()
switch t := t.(type) {
default:
    fmt.Printf("unexpected type %T\n", t)     // %T prints whatever type t has
case bool:
    fmt.Printf("boolean %t\n", t)             // t has type bool
case int:
    fmt.Printf("integer %d\n", t)             // t has type int
case *bool:
    fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
case *int:
    fmt.Printf("pointer to integer %d\n", *t) // t has type *int
}
</pre>

# Functions

## Multiple return values

<p>
One of Go's unusual features is that functions and methods
can return multiple values.  This form can be used to
improve on a couple of clumsy idioms in C programs: in-band
error returns such as <code>-1</code> for <code>EOF</code>
and modifying an argument passed by address.
</p>

<p>
In C, a write error is signaled by a negative count with the
error code secreted away in a volatile location.
In Go, <code>Write</code>
can return a count <i>and</i> an error: &ldquo;Yes, you wrote some
bytes but not all of them because you filled the device&rdquo;.
The signature of the <code>Write</code> method on files from
package <code>os</code> is:
</p>

<pre>
func (file *File) Write(b []byte) (n int, err error)
</pre>

<p>
and as the documentation says, it returns the number of bytes
written and a non-nil <code>error</code> when <code>n</code>
<code>!=</code> <code>len(b)</code>.
This is a common style; see the section on error handling for more examples.
</p>

<p>
A similar approach obviates the need to pass a pointer to a return
value to simulate a reference parameter.
Here's a simple-minded function to
grab a number from a position in a byte slice, returning the number
and the next position.
</p>

<pre>
func nextInt(b []byte, i int) (int, int) {
    for ; i &lt; len(b) &amp;&amp; !isDigit(b[i]); i++ {
    }
    x := 0
    for ; i &lt; len(b) &amp;&amp; isDigit(b[i]); i++ {
        x = x*10 + int(b[i]) - '0'
    }
    return x, i
}
</pre>

<p>
You could use it to scan the numbers in an input slice <code>b</code> like this:
</p>

<pre>
    for i := 0; i &lt; len(b); {
        x, i = nextInt(b, i)
        fmt.Println(x)
    }
</pre>

## Named result parameters

<p>
The return or result "parameters" of a Go function can be given names and
used as regular variables, just like the incoming parameters.
When named, they are initialized to the zero values for their types when
the function begins; if the function executes a <code>return</code> statement
with no arguments, the current values of the result parameters are
used as the returned values.
</p>

<p>
The names are not mandatory but they can make code shorter and clearer:
they're documentation.
If we name the results of <code>nextInt</code> it becomes
obvious which returned <code>int</code>
is which.
</p>

<pre>
func nextInt(b []byte, pos int) (value, nextPos int) {
</pre>

<p>
Because named results are initialized and tied to an unadorned return, they can simplify
as well as clarify.  Here's a version
of <code>io.ReadFull</code> that uses them well:
</p>

<pre>
func ReadFull(r Reader, buf []byte) (n int, err error) {
    for len(buf) &gt; 0 &amp;&amp; err == nil {
        var nr int
        nr, err = r.Read(buf)
        n += nr
        buf = buf[nr:]
    }
    return
}
</pre>

## Defer

<p>
Go's <code>defer</code> statement schedules a function call (the
<i>deferred</i> function) to be run immediately before the function
executing the <code>defer</code> returns.  It's an unusual but
effective way to deal with situations such as resources that must be
released regardless of which path a function takes to return.  The
canonical examples are unlocking a mutex or closing a file.
</p>

<pre>
// Contents returns the file's contents as a string.
func Contents(filename string) (string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer f.Close()  // f.Close will run when we're finished.

    var result []byte
    buf := make([]byte, 100)
    for {
        n, err := f.Read(buf[0:])
        result = append(result, buf[0:n]...) // append is discussed later.
        if err != nil {
            if err == io.EOF {
                break
            }
            return "", err  // f will be closed if we return here.
        }
    }
    return string(result), nil // f will be closed if we return here.
}
</pre>

<p>
Deferring a call to a function such as <code>Close</code> has two advantages.  First, it
guarantees that you will never forget to close the file, a mistake
that's easy to make if you later edit the function to add a new return
path.  Second, it means that the close sits near the open,
which is much clearer than placing it at the end of the function.
</p>

<p>
The arguments to the deferred function (which include the receiver if
the function is a method) are evaluated when the <i>defer</i>
executes, not when the <i>call</i> executes.  Besides avoiding worries
about variables changing values as the function executes, this means
that a single deferred call site can defer multiple function
executions.  Here's a silly example.
</p>

<pre>
for i := 0; i &lt; 5; i++ {
    defer fmt.Printf("%d ", i)
}
</pre>

<p>
Deferred functions are executed in LIFO order, so this code will cause
<code>4 3 2 1 0</code> to be printed when the function returns.  A
more plausible example is a simple way to trace function execution
through the program.  We could write a couple of simple tracing
routines like this:
</p>

<pre>
func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

// Use them like this:
func a() {
    trace("a")
    defer untrace("a")
    // do something....
}
</pre>

<p>
We can do better by exploiting the fact that arguments to deferred
functions are evaluated when the <code>defer</code> executes.  The
tracing routine can set up the argument to the untracing routine.
This example:
</p>

<pre>
func trace(s string) string {
    fmt.Println("entering:", s)
    return s
}

func un(s string) {
    fmt.Println("leaving:", s)
}

func a() {
    defer un(trace("a"))
    fmt.Println("in a")
}

func b() {
    defer un(trace("b"))
    fmt.Println("in b")
    a()
}

func main() {
    b()
}
</pre>

<p>
prints
</p>

<pre>
entering: b
in b
entering: a
in a
leaving: a
leaving: b
</pre>

<p>
For programmers accustomed to block-level resource management from
other languages, <code>defer</code> may seem peculiar, but its most
interesting and powerful applications come precisely from the fact
that it's not block-based but function-based.  In the section on
<code>panic</code> and <code>recover</code> we'll see another
example of its possibilities.
</p>

# Data

## Allocation with `new`

<p>
Go has two allocation primitives, the built-in functions
<code>new</code> and <code>make</code>.
They do different things and apply to different types, which can be confusing,
but the rules are simple.
Let's talk about <code>new</code> first.
It's a built-in function that allocates memory, but unlike its namesakes
in some other languages it does not <em>initialize</em> the memory,
it only <em>zeros</em> it.
That is,
<code>new(T)</code> allocates zeroed storage for a new item of type
<code>T</code> and returns its address, a value of type <code>*T</code>.
In Go terminology, it returns a pointer to a newly allocated zero value of type
<code>T</code>.
</p>

<p>
Since the memory returned by <code>new</code> is zeroed, it's helpful to arrange
when designing your data structures that the
zero value of each type can be used without further initialization.  This means a user of
the data structure can create one with <code>new</code> and get right to
work.
For example, the documentation for <code>bytes.Buffer</code> states that
"the zero value for <code>Buffer</code> is an empty buffer ready to use."
Similarly, <code>sync.Mutex</code> does not
have an explicit constructor or <code>Init</code> method.
Instead, the zero value for a <code>sync.Mutex</code>
is defined to be an unlocked mutex.
</p>

<p>
The zero-value-is-useful property works transitively. Consider this type declaration.
</p>

<pre>
type SyncedBuffer struct {
    lock    sync.Mutex
    buffer  bytes.Buffer
}
</pre>

<p>
Values of type <code>SyncedBuffer</code> are also ready to use immediately upon allocation
or just declaration.  In the next snippet, both <code>p</code> and <code>v</code> will work
correctly without further arrangement.
</p>

<pre>
p := new(SyncedBuffer)  // type *SyncedBuffer
var v SyncedBuffer      // type  SyncedBuffer
</pre>

## Constructors and composite literals

<p>
Sometimes the zero value isn't good enough and an initializing
constructor is necessary, as in this example derived from
package <code>os</code>.
</p>

<pre>
func NewFile(fd int, name string) *File {
    if fd &lt; 0 {
        return nil
    }
    f := new(File)
    f.fd = fd
    f.name = name
    f.dirinfo = nil
    f.nepipe = 0
    return f
}
</pre>

<p>
There's a lot of boilerplate in there.  We can simplify it
using a <i>composite literal</i>, which is
an expression that creates a
new instance each time it is evaluated.
</p>

<pre>
func NewFile(fd int, name string) *File {
    if fd &lt; 0 {
        return nil
    }
    f := File{fd, name, nil, 0}
    return &amp;f
}
</pre>

<p>
Note that, unlike in C, it's perfectly OK to return the address of a local variable;
the storage associated with the variable survives after the function
returns.
In fact, taking the address of a composite literal
allocates a fresh instance each time it is evaluated,
so we can combine these last two lines.
</p>

<pre>
    return &amp;File{fd, name, nil, 0}
</pre>

<p>
The fields of a composite literal are laid out in order and must all be present.
However, by labeling the elements explicitly as <i>field</i><code>:</code><i>value</i>
pairs, the initializers can appear in any
order, with the missing ones left as their respective zero values.  Thus we could say
</p>

<pre>
    return &amp;File{fd: fd, name: name}
</pre>

<p>
As a limiting case, if a composite literal contains no fields at all, it creates
a zero value for the type.  The expressions <code>new(File)</code> and <code>&amp;File{}</code> are equivalent.
</p>

<p>
Composite literals can also be created for arrays, slices, and maps,
with the field labels being indices or map keys as appropriate.
In these examples, the initializations work regardless of the values of <code>Enone</code>,
<code>Eio</code>, and <code>Einval</code>, as long as they are distinct.
</p>

<pre>
a := [...]string   {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
s := []string      {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
</pre>

## Allocation with `make`

<p>
Back to allocation.
The built-in function <code>make(T, </code><i>args</i><code>)</code> serves
a purpose different from <code>new(T)</code>.
It creates slices, maps, and channels only, and it returns an <em>initialized</em>
(not <em>zeroed</em>)
value of type <code>T</code> (not <code>*T</code>).
The reason for the distinction
is that these three types represent, under the covers, references to data structures that
must be initialized before use.
A slice, for example, is a three-item descriptor
containing a pointer to the data (inside an array), the length, and the
capacity, and until those items are initialized, the slice is <code>nil</code>.
For slices, maps, and channels,
<code>make</code> initializes the internal data structure and prepares
the value for use.
For instance,
</p>

<pre>
make([]int, 10, 100)
</pre>

<p>
allocates an array of 100 ints and then creates a slice
structure with length 10 and a capacity of 100 pointing at the first
10 elements of the array.
(When making a slice, the capacity can be omitted; see the section on slices
for more information.)
In contrast, <code>new([]int)</code> returns a pointer to a newly allocated, zeroed slice
structure, that is, a pointer to a <code>nil</code> slice value.
</p>

<p>
These examples illustrate the difference between <code>new</code> and
<code>make</code>.
</p>

<pre>
var p *[]int = new([]int)       // allocates slice structure; *p == nil; rarely useful
var v  []int = make([]int, 100) // the slice v now refers to a new array of 100 ints

// Unnecessarily complex:
var p *[]int = new([]int)
*p = make([]int, 100, 100)

// Idiomatic:
v := make([]int, 100)
</pre>

<p>
Remember that <code>make</code> applies only to maps, slices and channels
and does not return a pointer.
To obtain an explicit pointer allocate with <code>new</code> or take the address
of a variable explicitly.
</p>

## Arrays

<p>
Arrays are useful when planning the detailed layout of memory and sometimes
can help avoid allocation, but primarily
they are a building block for slices, the subject of the next section.
To lay the foundation for that topic, here are a few words about arrays.
</p>

<p>
There are major differences between the ways arrays work in Go and C.
In Go,
</p>
<ul>
<li>
Arrays are values. Assigning one array to another copies all the elements.
</li>
<li>
In particular, if you pass an array to a function, it
will receive a <i>copy</i> of the array, not a pointer to it.
<li>
The size of an array is part of its type.  The types <code>[10]int</code>
and <code>[20]int</code> are distinct.
</li>
</ul>

<p>
The value property can be useful but also expensive; if you want C-like behavior and efficiency,
you can pass a pointer to the array.
</p>

<pre>
func Sum(a *[3]float64) (sum float64) {
    for _, v := range *a {
        sum += v
    }
    return
}

array := [...]float64{7.0, 8.5, 9.1}
x := Sum(&amp;array)  // Note the explicit address-of operator
</pre>

<p>
But even this style isn't idiomatic Go.
Use slices instead.
</p>

## Slices

<p>
Slices wrap arrays to give a more general, powerful, and convenient
interface to sequences of data.  Except for items with explicit
dimension such as transformation matrices, most array programming in
Go is done with slices rather than simple arrays.
</p>
<p>
Slices hold references to an underlying array, and if you assign one
slice to another, both refer to the same array.
If a function takes a slice argument, changes it makes to
the elements of the slice will be visible to the caller, analogous to
passing a pointer to the underlying array.  A <code>Read</code>
function can therefore accept a slice argument rather than a pointer
and a count; the length within the slice sets an upper
limit of how much data to read.  Here is the signature of the
<code>Read</code> method of the <code>File</code> type in package
<code>os</code>:
</p>
<pre>
func (f *File) Read(buf []byte) (n int, err error)
</pre>
<p>
The method returns the number of bytes read and an error value, if
any.
To read into the first 32 bytes of a larger buffer
<code>buf</code>, <i>slice</i> (here used as a verb) the buffer.
</p>
<pre>
    n, err := f.Read(buf[0:32])
</pre>
<p>
Such slicing is common and efficient.  In fact, leaving efficiency aside for
the moment, the following snippet would also read the first 32 bytes of the buffer.
</p>
<pre>
    var n int
    var err error
    for i := 0; i &lt; 32; i++ {
        nbytes, e := f.Read(buf[i:i+1])  // Read one byte.
        n += nbytes
        if nbytes == 0 || e != nil {
            err = e
            break
        }
    }
</pre>
<p>
The length of a slice may be changed as long as it still fits within
the limits of the underlying array; just assign it to a slice of
itself.  The <i>capacity</i> of a slice, accessible by the built-in
function <code>cap</code>, reports the maximum length the slice may
assume.  Here is a function to append data to a slice.  If the data
exceeds the capacity, the slice is reallocated.  The
resulting slice is returned.  The function uses the fact that
<code>len</code> and <code>cap</code> are legal when applied to the
<code>nil</code> slice, and return 0.
</p>
<pre>
func Append(slice, data []byte) []byte {
    l := len(slice)
    if l + len(data) &gt; cap(slice) {  // reallocate
        // Allocate double what's needed, for future growth.
        newSlice := make([]byte, (l+len(data))*2)
        // The copy function is predeclared and works for any slice type.
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0:l+len(data)]
    copy(slice[l:], data)
    return slice
}
</pre>
<p>
We must return the slice afterwards because, although <code>Append</code>
can modify the elements of <code>slice</code>, the slice itself (the run-time data
structure holding the pointer, length, and capacity) is passed by value.
</p>

<p>
The idea of appending to a slice is so useful it's captured by the
<code>append</code> built-in function.  To understand that function's
design, though, we need a little more information, so we'll return
to it later.
</p>

## Two-dimensional slices

<p>
Go's arrays and slices are one-dimensional.
To create the equivalent of a 2D array or slice, it is necessary to define an array-of-arrays
or slice-of-slices, like this:
</p>

<pre>
type Transform [3][3]float64  // A 3x3 array, really an array of arrays.
type LinesOfText [][]byte     // A slice of byte slices.
</pre>

<p>
Because slices are variable-length, it is possible to have each inner
slice be a different length.
That can be a common situation, as in our <code>LinesOfText</code>
example: each line has an independent length.
</p>

<pre>
text := LinesOfText{
    []byte("Now is the time"),
    []byte("for all good gophers"),
    []byte("to bring some fun to the party."),
}
</pre>

<p>
Sometimes it's necessary to allocate a 2D slice, a situation that can arise when
processing scan lines of pixels, for instance.
There are two ways to achieve this.
One is to allocate each slice independently; the other
is to allocate a single array and point the individual slices into it.
Which to use depends on your application.
If the slices might grow or shrink, they should be allocated independently
to avoid overwriting the next line; if not, it can be more efficient to construct
the object with a single allocation.
For reference, here are sketches of the two methods.
First, a line at a time:
</p>

<pre>
// Allocate the top-level slice.
picture := make([][]uint8, YSize) // One row per unit of y.
// Loop over the rows, allocating the slice for each row.
for i := range picture {
    picture[i] = make([]uint8, XSize)
}
</pre>

<p>
And now as one allocation, sliced into lines:
</p>

<pre>
// Allocate the top-level slice, the same as before.
picture := make([][]uint8, YSize) // One row per unit of y.
// Allocate one large slice to hold all the pixels.
pixels := make([]uint8, XSize*YSize) // Has type []uint8 even though picture is [][]uint8.
// Loop over the rows, slicing each row from the front of the remaining pixels slice.
for i := range picture {
    picture[i], pixels = pixels[:XSize], pixels[XSize:]
}
</pre>

## Maps

* Maps
  * == built-in data structure /
    * convenient
    * powerful
    * values of one type (`key`) are -- associated with -- values of another type (`element` OR `value`)
    
    The key can be of any type for which the equality operator is defined,
    such as integers,
    floating point and complex numbers,
    strings, pointers, interfaces (as long as the dynamic type
    supports equality), structs and arrays.
    Slices cannot be used as map keys,
    because equality is not defined on them.
    Like slices, maps hold references to an underlying data structure.
    If you pass a map to a function
    that changes the contents of the map, the changes will be visible
    in the caller.

<p>
Maps can be constructed using the usual composite literal syntax
with colon-separated key-value pairs,
so it's easy to build them during initialization.
</p>
<pre>
var timeZone = map[string]int{
    "UTC":  0*60*60,
    "EST": -5*60*60,
    "CST": -6*60*60,
    "MST": -7*60*60,
    "PST": -8*60*60,
}
</pre>
<p>
Assigning and fetching map values looks syntactically just like
doing the same for arrays and slices except that the index doesn't
need to be an integer.
</p>
<pre>
offset := timeZone["EST"]
</pre>
<p>
An attempt to fetch a map value with a key that
is not present in the map will return the zero value for the type
of the entries
in the map.  For instance, if the map contains integers, looking
up a non-existent key will return <code>0</code>.
A set can be implemented as a map with value type <code>bool</code>.
Set the map entry to <code>true</code> to put the value in the set, and then
test it by simple indexing.
</p>
<pre>
attended := map[string]bool{
    "Ann": true,
    "Joe": true,
    ...
}

if attended[person] { // will be false if person is not in the map
    fmt.Println(person, "was at the meeting")
}
</pre>
<p>
Sometimes you need to distinguish a missing entry from
a zero value.  Is there an entry for <code>"UTC"</code>
or is that 0 because it's not in the map at all?
You can discriminate with a form of multiple assignment.
</p>
<pre>
var seconds int
var ok bool
seconds, ok = timeZone[tz]
</pre>
<p>
For obvious reasons this is called the &ldquo;comma ok&rdquo; idiom.
In this example, if <code>tz</code> is present, <code>seconds</code>
will be set appropriately and <code>ok</code> will be true; if not,
<code>seconds</code> will be set to zero and <code>ok</code> will
be false.
Here's a function that puts it together with a nice error report:
</p>
<pre>
func offset(tz string) int {
    if seconds, ok := timeZone[tz]; ok {
        return seconds
    }
    log.Println("unknown time zone:", tz)
    return 0
}
</pre>
<p>
To test for presence in the map without worrying about the actual value,
you can use the <a href="#blank">blank identifier</a> (<code>_</code>)
in place of the usual variable for the value.
</p>
<pre>
_, present := timeZone[tz]
</pre>
<p>
To delete a map entry, use the <code>delete</code>
built-in function, whose arguments are the map and the key to be deleted.
It's safe to do this even if the key is already absent
from the map.
</p>
<pre>
delete(timeZone, "PDT")  // Now on Standard Time
</pre>

## Printing

* Go's printing
  * vs C's `printf` family
    * similar
    * richer
    * MORE general
  * `fmt` package
  * capitalized names
    * `fmt.Printf`
    * `fmt.Fprintf`
    * `fmt.Sprintf`
    * ...

* string functions
  * _Example:_ `Sprintf`
  * return a string
    * != C == fill | provided buffer
    ```C
    char buffer[100];  // you need to provide a buffer
    sprintf(buffer, "Hello %s", name);  // fill in the provided buffer
    ```
    ```go
    result := fmt.Sprintf("Hello %s", name)  // return directly NEW string
    ```
  * ❌NOT required to provide -- a -- format string❌
  * / EACH formatted function (_Example:_ `Printf`, `Fprintf` & `Sprintf`), exist
    * no formatted functions (_Example:_ `Print`, `Fprint` & `Sprint`)
      * ❌NOT accept a format string❌
      * generate a default format / EACH argument
    * breaking line functions (_Example:_ `Println`, `Fprintln` & `Sprintln`)
      * ❌NOT accept a format string❌
      * generate a default format / EACH argument

* TODO: The <code>Println</code> versions also insert a blank
between arguments and append a newline to the output while
the <code>Print</code> versions add blanks only if the operand on neither side is a string.
In this example each line produces the same output.


<pre>
fmt.Printf("Hello %d\n", 23)
fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
fmt.Println("Hello", 23)
fmt.Println(fmt.Sprint("Hello ", 23))
</pre>
<p>
The formatted print functions <code>fmt.Fprint</code>
and friends take as a first argument any object
that implements the <code>io.Writer</code> interface; the variables <code>os.Stdout</code>
and <code>os.Stderr</code> are familiar instances.
</p>
<p>
Here things start to diverge from C.  First, the numeric formats such as <code>%d</code>
do not take flags for signedness or size; instead, the printing routines use the
type of the argument to decide these properties.
</p>
<pre>
var x uint64 = 1&lt;&lt;64 - 1
fmt.Printf("%d %x; %d %x\n", x, x, int64(x), int64(x))
</pre>
<p>
prints
</p>
<pre>
18446744073709551615 ffffffffffffffff; -1 -1
</pre>
<p>
If you just want the default conversion, such as decimal for integers, you can use
the catchall format <code>%v</code> (for &ldquo;value&rdquo;); the result is exactly
what <code>Print</code> and <code>Println</code> would produce.
Moreover, that format can print <em>any</em> value, even arrays, slices, structs, and
maps.  Here is a print statement for the time zone map defined in the previous section.
</p>
<pre>
fmt.Printf("%v\n", timeZone)  // or just fmt.Println(timeZone)
</pre>
<p>
which gives output:
</p>
<pre>
map[CST:-21600 EST:-18000 MST:-25200 PST:-28800 UTC:0]
</pre>
<p>
For maps, <code>Printf</code> and friends sort the output lexicographically by key.
</p>
<p>
When printing a struct, the modified format <code>%+v</code> annotates the
fields of the structure with their names, and for any value the alternate
format <code>%#v</code> prints the value in full Go syntax.
</p>
<pre>
type T struct {
    a int
    b float64
    c string
}
t := &amp;T{ 7, -2.35, "abc\tdef" }
fmt.Printf("%v\n", t)
fmt.Printf("%+v\n", t)
fmt.Printf("%#v\n", t)
fmt.Printf("%#v\n", timeZone)
</pre>
<p>
prints
</p>
<pre>
&amp;{7 -2.35 abc   def}
&amp;{a:7 b:-2.35 c:abc     def}
&amp;main.T{a:7, b:-2.35, c:"abc\tdef"}
map[string]int{"CST":-21600, "EST":-18000, "MST":-25200, "PST":-28800, "UTC":0}
</pre>
<p>
(Note the ampersands.)
That quoted string format is also available through <code>%q</code> when
applied to a value of type <code>string</code> or <code>[]byte</code>.
The alternate format <code>%#q</code> will use backquotes instead if possible.
(The <code>%q</code> format also applies to integers and runes, producing a
single-quoted rune constant.)
Also, <code>%x</code> works on strings, byte arrays and byte slices as well as
on integers, generating a long hexadecimal string, and with
a space in the format (<code>%&nbsp;x</code>) it puts spaces between the bytes.
</p>
<p>
Another handy format is <code>%T</code>, which prints the <em>type</em> of a value.
</p>
<pre>
fmt.Printf(&quot;%T\n&quot;, timeZone)
</pre>
<p>
prints
</p>
<pre>
map[string]int
</pre>
<p>
If you want to control the default format for a custom type, all that's required is to define
a method with the signature <code>String() string</code> on the type.
For our simple type <code>T</code>, that might look like this.
</p>
<pre>
func (t *T) String() string {
    return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
}
fmt.Printf("%v\n", t)
</pre>
<p>
to print in the format
</p>
<pre>
7/-2.35/"abc\tdef"
</pre>
<p>
(If you need to print <em>values</em> of type <code>T</code> as well as pointers to <code>T</code>,
the receiver for <code>String</code> must be of value type; this example used a pointer because
that's more efficient and idiomatic for struct types.
See the section below on <a href="#pointers_vs_values">pointers vs. value receivers</a> for more information.)
</p>

<p>
Our <code>String</code> method is able to call <code>Sprintf</code> because the
print routines are fully reentrant and can be wrapped this way.
There is one important detail to understand about this approach,
however: don't construct a <code>String</code> method by calling
<code>Sprintf</code> in a way that will recur into your <code>String</code>
method indefinitely.  This can happen if the <code>Sprintf</code>
call attempts to print the receiver directly as a string, which in
turn will invoke the method again.  It's a common and easy mistake
to make, as this example shows.
</p>

<pre>
type MyString string

func (m MyString) String() string {
    return fmt.Sprintf("MyString=%s", m) // Error: will recur forever.
}
</pre>

<p>
It's also easy to fix: convert the argument to the basic string type, which does not have the
method.
</p>

<pre>
type MyString string
func (m MyString) String() string {
    return fmt.Sprintf("MyString=%s", string(m)) // OK: note conversion.
}
</pre>

<p>
In the <a href="#initialization">initialization section</a> we'll see another technique that avoids this recursion.
</p>

<p>
Another printing technique is to pass a print routine's arguments directly to another such routine.
The signature of <code>Printf</code> uses the type <code>...interface{}</code>
for its final argument to specify that an arbitrary number of parameters (of arbitrary type)
can appear after the format.
</p>
<pre>
func Printf(format string, v ...interface{}) (n int, err error) {
</pre>
<p>
Within the function <code>Printf</code>, <code>v</code> acts like a variable of type
<code>[]interface{}</code> but if it is passed to another variadic function, it acts like
a regular list of arguments.
Here is the implementation of the
function <code>log.Println</code> we used above. It passes its arguments directly to
<code>fmt.Sprintln</code> for the actual formatting.
</p>
<pre>
// Println prints to the standard logger in the manner of fmt.Println.
func Println(v ...interface{}) {
    std.Output(2, fmt.Sprintln(v...))  // Output takes parameters (int, string)
}
</pre>
<p>
We write <code>...</code> after <code>v</code> in the nested call to <code>Sprintln</code> to tell the
compiler to treat <code>v</code> as a list of arguments; otherwise it would just pass
<code>v</code> as a single slice argument.
</p>
<p>
There's even more to printing than we've covered here.  See the <code>godoc</code> documentation
for package <code>fmt</code> for the details.
</p>
<p>
By the way, a <code>...</code> parameter can be of a specific type, for instance <code>...int</code>
for a min function that chooses the least of a list of integers:
</p>
<pre>
func Min(a ...int) int {
    min := int(^uint(0) &gt;&gt; 1)  // largest int
    for _, i := range a {
        if i &lt; min {
            min = i
        }
    }
    return min
}
</pre>

## Append
<p>
Now we have the missing piece we needed to explain the design of
the <code>append</code> built-in function.  The signature of <code>append</code>
is different from our custom <code>Append</code> function above.
Schematically, it's like this:
</p>
<pre>
func append(slice []<i>T</i>, elements ...<i>T</i>) []<i>T</i>
</pre>
<p>
where <i>T</i> is a placeholder for any given type.  You can't
actually write a function in Go where the type <code>T</code>
is determined by the caller.
That's why <code>append</code> is built in: it needs support from the
compiler.
</p>
<p>
What <code>append</code> does is append the elements to the end of
the slice and return the result.  The result needs to be returned
because, as with our hand-written <code>Append</code>, the underlying
array may change.  This simple example
</p>
<pre>
x := []int{1,2,3}
x = append(x, 4, 5, 6)
fmt.Println(x)
</pre>
<p>
prints <code>[1 2 3 4 5 6]</code>.  So <code>append</code> works a
little like <code>Printf</code>, collecting an arbitrary number of
arguments.
</p>
<p>
But what if we wanted to do what our <code>Append</code> does and
append a slice to a slice?  Easy: use <code>...</code> at the call
site, just as we did in the call to <code>Output</code> above.  This
snippet produces identical output to the one above.
</p>
<pre>
x := []int{1,2,3}
y := []int{4,5,6}
x = append(x, y...)
fmt.Println(x)
</pre>
<p>
Without that <code>...</code>, it wouldn't compile because the types
would be wrong; <code>y</code> is not of type <code>int</code>.
</p>

# Initialization

* Go initialization vs C or C++ initialization
  * 👀MORE powerful👀
    * Reason:🧠
      * ALLOWED building complex structures
      * | order initialized objects, NOT problem🧠

## Constants

* ⚠️created | compile time,⚠️
  * even defined -- as -- locals | functions
  * -> expressions == constant expressions / evaluable -- by the -- compiler
    ```go
     1<<3                   # == constant expression
    math.Sin(math.Pi/4)     # !=  constant expression -- Reason: math.Sin happens | run-time
    ```

* ⚠️ONLY ALLOWED types⚠️
  * numbers,
  * characters (runes),
  * strings 
  * booleans

* `iota` enumerator
  * allows
    * creating enumerated constants
      * Reason:🧠
        * `iota` can be -- part of an -- expression
        * expressions can be IMPLICITLY repeated🧠

* _Example:_ [here](progs/eff_bytesize.go)

* | user-defined types or scalar types
  * 👀if you define `String()` -> Go used it AUTOMATICALLY | print that value👀
  * _Example:_ [`type ByteSize`](progs/eff_bytesize.go)

## Variables

<p>
Variables can be initialized just like constants but the
initializer can be a general expression computed at run time.
</p>
<pre>
var (
    home   = os.Getenv("HOME")
    user   = os.Getenv("USER")
    gopath = os.Getenv("GOPATH")
)
</pre>

## The init function

<p>
Finally, each source file can define its own niladic <code>init</code> function to
set up whatever state is required.  (Actually each file can have multiple
<code>init</code> functions.)
And finally means finally: <code>init</code> is called after all the
variable declarations in the package have evaluated their initializers,
and those are evaluated only after all the imported packages have been
initialized.
</p>
<p>
Besides initializations that cannot be expressed as declarations,
a common use of <code>init</code> functions is to verify or repair
correctness of the program state before real execution begins.
</p>

<pre>
func init() {
    if user == "" {
        log.Fatal("$USER not set")
    }
    if home == "" {
        home = "/home/" + user
    }
    if gopath == "" {
        gopath = home + "/go"
    }
    // gopath may be overridden by --gopath flag on command line.
    flag.StringVar(&amp;gopath, "gopath", gopath, "override default GOPATH")
}
</pre>

# Methods

## Pointers vs. Values

* methods
  * can be defined | 
    * ANY named type
      * ⚠️EXCEPT to, pointer OR interface⚠️
      * _Example:_ [here](progs/eff_bytesize.go)
      * 's receiver
        * ❌NOT have to be a struct❌
          * _Example:_ [here](progs/eff_bytesize.go)

* TODO: 
<p>
In the discussion of slices above, we wrote an <code>Append</code>
function.  We can define it as a method on slices instead.  To do
this, we first declare a named type to which we can bind the method, and
then make the receiver for the method a value of that type.
</p>
<pre>
type ByteSlice []byte

func (slice ByteSlice) Append(data []byte) []byte {
    // Body exactly the same as the Append function defined above.
}
</pre>

<p>
This still requires the method to return the updated slice.  We can
eliminate that clumsiness by redefining the method to take a
<i>pointer</i> to a <code>ByteSlice</code> as its receiver, so the
method can overwrite the caller's slice.
</p>

<pre>
func (p *ByteSlice) Append(data []byte) {
    slice := *p
    // Body as above, without the return.
    *p = slice
}
</pre>
<p>
In fact, we can do even better.  If we modify our function so it looks
like a standard <code>Write</code> method, like this,
</p>
<pre>
func (p *ByteSlice) Write(data []byte) (n int, err error) {
    slice := *p
    // Again as above.
    *p = slice
    return len(data), nil
}
</pre>
<p>
then the type <code>*ByteSlice</code> satisfies the standard interface
<code>io.Writer</code>, which is handy.  For instance, we can
print into one.
</p>
<pre>
    var b ByteSlice
    fmt.Fprintf(&amp;b, "This hour has %d days\n", 7)
</pre>
<p>
We pass the address of a <code>ByteSlice</code>
because only <code>*ByteSlice</code> satisfies <code>io.Writer</code>.
The rule about pointers vs. values for receivers is that value methods
can be invoked on pointers and values, but pointer methods can only be
invoked on pointers.
</p>

<p>
This rule arises because pointer methods can modify the receiver; invoking
them on a value would cause the method to receive a copy of the value, so
any modifications would be discarded.
The language therefore disallows this mistake.
There is a handy exception, though. When the value is addressable, the
language takes care of the common case of invoking a pointer method on a
value by inserting the address operator automatically.
In our example, the variable <code>b</code> is addressable, so we can call
its <code>Write</code> method with just <code>b.Write</code>. The compiler
will rewrite that to <code>(&amp;b).Write</code> for us.
</p>

<p>
By the way, the idea of using <code>Write</code> on a slice of bytes
is central to the implementation of <code>bytes.Buffer</code>.
</p>

# Interfaces & OTHER types

## Interfaces
<p>
Interfaces in Go provide a way to specify the behavior of an
object: if something can do <em>this</em>, then it can be used
<em>here</em>.  We've seen a couple of simple examples already;
custom printers can be implemented by a <code>String</code> method
while <code>Fprintf</code> can generate output to anything
with a <code>Write</code> method.
Interfaces with only one or two methods are common in Go code, and are
usually given a name derived from the method, such as <code>io.Writer</code>
for something that implements <code>Write</code>.
</p>
<p>
A type can implement multiple interfaces.
For instance, a collection can be sorted
by the routines in package <code>sort</code> if it implements
<code>sort.Interface</code>, which contains <code>Len()</code>,
<code>Less(i, j int) bool</code>, and <code>Swap(i, j int)</code>,
and it could also have a custom formatter.
In this contrived example <code>Sequence</code> satisfies both.
</p>
{{code "/doc/progs/eff_sequence.go" `/^type/` "$"}}

## Conversions

<p>
The <code>String</code> method of <code>Sequence</code> is recreating the
work that <code>Sprint</code> already does for slices.
(It also has complexity O(N²), which is poor.) We can share the
effort (and also speed it up) if we convert the <code>Sequence</code> to a plain
<code>[]int</code> before calling <code>Sprint</code>.
</p>
<pre>
func (s Sequence) String() string {
    s = s.Copy()
    sort.Sort(s)
    return fmt.Sprint([]int(s))
}
</pre>
<p>
This method is another example of the conversion technique for calling
<code>Sprintf</code> safely from a <code>String</code> method.
Because the two types (<code>Sequence</code> and <code>[]int</code>)
are the same if we ignore the type name, it's legal to convert between them.
The conversion doesn't create a new value, it just temporarily acts
as though the existing value has a new type.
(There are other legal conversions, such as from integer to floating point, that
do create a new value.)
</p>
<p>
It's an idiom in Go programs to convert the
type of an expression to access a different
set of methods. As an example, we could use the existing
type <code>sort.IntSlice</code> to reduce the entire example
to this:
</p>
<pre>
type Sequence []int

// Method for printing - sorts the elements before printing
func (s Sequence) String() string {
    s = s.Copy()
    sort.IntSlice(s).Sort()
    return fmt.Sprint([]int(s))
}
</pre>
<p>
Now, instead of having <code>Sequence</code> implement multiple
interfaces (sorting and printing), we're using the ability of a data item to be
converted to multiple types (<code>Sequence</code>, <code>sort.IntSlice</code>
and <code>[]int</code>), each of which does some part of the job.
That's more unusual in practice but can be effective.
</p>

<h3 id="interface_conversions">Interface conversions and type assertions</h3>

<p>
<a href="#type_switch">Type switches</a> are a form of conversion: they take an interface and, for
each case in the switch, in a sense convert it to the type of that case.
Here's a simplified version of how the code under <code>fmt.Printf</code> turns a value into
a string using a type switch.
If it's already a string, we want the actual string value held by the interface, while if it has a
<code>String</code> method we want the result of calling the method.
</p>

<pre>
type Stringer interface {
    String() string
}

var value interface{} // Value provided by caller.
switch str := value.(type) {
case string:
    return str
case Stringer:
    return str.String()
}
</pre>

<p>
The first case finds a concrete value; the second converts the interface into another interface.
It's perfectly fine to mix types this way.
</p>

<p>
What if there's only one type we care about? If we know the value holds a <code>string</code>
and we just want to extract it?
A one-case type switch would do, but so would a <em>type assertion</em>.
A type assertion takes an interface value and extracts from it a value of the specified explicit type.
The syntax borrows from the clause opening a type switch, but with an explicit
type rather than the <code>type</code> keyword:
</p>

<pre>
value.(typeName)
</pre>

<p>
and the result is a new value with the static type <code>typeName</code>.
That type must either be the concrete type held by the interface, or a second interface
type that the value can be converted to.
To extract the string we know is in the value, we could write:
</p>

<pre>
str := value.(string)
</pre>

<p>
But if it turns out that the value does not contain a string, the program will crash with a run-time error.
To guard against that, use the "comma, ok" idiom to test, safely, whether the value is a string:
</p>

<pre>
str, ok := value.(string)
if ok {
    fmt.Printf("string value is: %q\n", str)
} else {
    fmt.Printf("value is not a string\n")
}
</pre>

<p>
If the type assertion fails, <code>str</code> will still exist and be of type string, but it will have
the zero value, an empty string.
</p>

<p>
As an illustration of the capability, here's an <code>if</code>-<code>else</code>
statement that's equivalent to the type switch that opened this section.
</p>

<pre>
if str, ok := value.(string); ok {
    return str
} else if str, ok := value.(Stringer); ok {
    return str.String()
}
</pre>

<h3 id="generality">Generality</h3>
<p>
If a type exists only to implement an interface and will
never have exported methods beyond that interface, there is
no need to export the type itself.
Exporting just the interface makes it clear the value has no
interesting behavior beyond what is described in the
interface.
It also avoids the need to repeat the documentation
on every instance of a common method.
</p>
<p>
In such cases, the constructor should return an interface value
rather than the implementing type.
As an example, in the hash libraries
both <code>crc32.NewIEEE</code> and <code>adler32.New</code>
return the interface type <code>hash.Hash32</code>.
Substituting the CRC-32 algorithm for Adler-32 in a Go program
requires only changing the constructor call;
the rest of the code is unaffected by the change of algorithm.
</p>
<p>
A similar approach allows the streaming cipher algorithms
in the various <code>crypto</code> packages to be
separated from the block ciphers they chain together.
The <code>Block</code> interface
in the <code>crypto/cipher</code> package specifies the
behavior of a block cipher, which provides encryption
of a single block of data.
Then, by analogy with the <code>bufio</code> package,
cipher packages that implement this interface
can be used to construct streaming ciphers, represented
by the <code>Stream</code> interface, without
knowing the details of the block encryption.
</p>
<p>
The  <code>crypto/cipher</code> interfaces look like this:
</p>
<pre>
type Block interface {
    BlockSize() int
    Encrypt(dst, src []byte)
    Decrypt(dst, src []byte)
}

type Stream interface {
    XORKeyStream(dst, src []byte)
}
</pre>

<p>
Here's the definition of the counter mode (CTR) stream,
which turns a block cipher into a streaming cipher; notice
that the block cipher's details are abstracted away:
</p>

<pre>
// NewCTR returns a Stream that encrypts/decrypts using the given Block in
// counter mode. The length of iv must be the same as the Block's block size.
func NewCTR(block Block, iv []byte) Stream
</pre>
<p>
<code>NewCTR</code> applies not
just to one specific encryption algorithm and data source but to any
implementation of the <code>Block</code> interface and any
<code>Stream</code>.  Because they return
interface values, replacing CTR
encryption with other encryption modes is a localized change.  The constructor
calls must be edited, but because the surrounding code must treat the result only
as a <code>Stream</code>, it won't notice the difference.
</p>

<h3 id="interface_methods">Interfaces and methods</h3>
<p>
Since almost anything can have methods attached, almost anything can
satisfy an interface.  One illustrative example is in the <code>http</code>
package, which defines the <code>Handler</code> interface.  Any object
that implements <code>Handler</code> can serve HTTP requests.
</p>
<pre>
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
</pre>
<p>
<code>ResponseWriter</code> is itself an interface that provides access
to the methods needed to return the response to the client.
Those methods include the standard <code>Write</code> method, so an
<code>http.ResponseWriter</code> can be used wherever an <code>io.Writer</code>
can be used.
<code>Request</code> is a struct containing a parsed representation
of the request from the client.
</p>
<p>
For brevity, let's ignore POSTs and assume HTTP requests are always
GETs; that simplification does not affect the way the handlers are set up.
Here's a trivial implementation of a handler to count the number of times
the page is visited.
</p>
<pre>
// Simple counter server.
type Counter struct {
    n int
}

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    ctr.n++
    fmt.Fprintf(w, "counter = %d\n", ctr.n)
}
</pre>
<p>
(Keeping with our theme, note how <code>Fprintf</code> can print to an
<code>http.ResponseWriter</code>.)
In a real server, access to <code>ctr.n</code> would need protection from
concurrent access.
See the <code>sync</code> and <code>atomic</code> packages for suggestions.
</p>
<p>
For reference, here's how to attach such a server to a node on the URL tree.
</p>
<pre>
import "net/http"
...
ctr := new(Counter)
http.Handle("/counter", ctr)
</pre>
<p>
But why make <code>Counter</code> a struct?  An integer is all that's needed.
(The receiver needs to be a pointer so the increment is visible to the caller.)
</p>
<pre>
// Simpler counter server.
type Counter int

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    *ctr++
    fmt.Fprintf(w, "counter = %d\n", *ctr)
}
</pre>
<p>
What if your program has some internal state that needs to be notified that a page
has been visited?  Tie a channel to the web page.
</p>
<pre>
// A channel that sends a notification on each visit.
// (Probably want the channel to be buffered.)
type Chan chan *http.Request

func (ch Chan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    ch &lt;- req
    fmt.Fprint(w, "notification sent")
}
</pre>
<p>
Finally, let's say we wanted to present on <code>/args</code> the arguments
used when invoking the server binary.
It's easy to write a function to print the arguments.
</p>
<pre>
func ArgServer() {
    fmt.Println(os.Args)
}
</pre>
<p>
How do we turn that into an HTTP server?  We could make <code>ArgServer</code>
a method of some type whose value we ignore, but there's a cleaner way.
Since we can define a method for any type except pointers and interfaces,
we can write a method for a function.
The <code>http</code> package contains this code:
</p>
<pre>
// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers.  If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler object that calls f.
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, req).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, req *Request) {
    f(w, req)
}
</pre>
<p>
<code>HandlerFunc</code> is a type with a method, <code>ServeHTTP</code>,
so values of that type can serve HTTP requests.  Look at the implementation
of the method: the receiver is a function, <code>f</code>, and the method
calls <code>f</code>.  That may seem odd but it's not that different from, say,
the receiver being a channel and the method sending on the channel.
</p>
<p>
To make <code>ArgServer</code> into an HTTP server, we first modify it
to have the right signature.
</p>
<pre>
// Argument server.
func ArgServer(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(w, os.Args)
}
</pre>
<p>
<code>ArgServer</code> now has the same signature as <code>HandlerFunc</code>,
so it can be converted to that type to access its methods,
just as we converted <code>Sequence</code> to <code>IntSlice</code>
to access <code>IntSlice.Sort</code>.
The code to set it up is concise:
</p>
<pre>
http.Handle("/args", http.HandlerFunc(ArgServer))
</pre>
<p>
When someone visits the page <code>/args</code>,
the handler installed at that page has value <code>ArgServer</code>
and type <code>HandlerFunc</code>.
The HTTP server will invoke the method <code>ServeHTTP</code>
of that type, with <code>ArgServer</code> as the receiver, which will in turn call
<code>ArgServer</code> (via the invocation <code>f(w, req)</code>
inside <code>HandlerFunc.ServeHTTP</code>).
The arguments will then be displayed.
</p>
<p>
In this section we have made an HTTP server from a struct, an integer,
a channel, and a function, all because interfaces are just sets of
methods, which can be defined for (almost) any type.
</p>

# The blank identifier

<p>
We've mentioned the blank identifier a couple of times now, in the context of
<a href="#for"><code>for</code> <code>range</code> loops</a>
and <a href="#maps">maps</a>.
The blank identifier can be assigned or declared with any value of any type, with the
value discarded harmlessly.
It's a bit like writing to the Unix <code>/dev/null</code> file:
it represents a write-only value
to be used as a place-holder
where a variable is needed but the actual value is irrelevant.
It has uses beyond those we've seen already.
</p>

<h3 id="blank_assign">The blank identifier in multiple assignment</h3>

<p>
The use of a blank identifier in a <code>for</code> <code>range</code> loop is a
special case of a general situation: multiple assignment.
</p>

<p>
If an assignment requires multiple values on the left side,
but one of the values will not be used by the program,
a blank identifier on the left-hand-side of
the assignment avoids the need
to create a dummy variable and makes it clear that the
value is to be discarded.
For instance, when calling a function that returns
a value and an error, but only the error is important,
use the blank identifier to discard the irrelevant value.
</p>

<pre>
if _, err := os.Stat(path); os.IsNotExist(err) {
    fmt.Printf("%s does not exist\n", path)
}
</pre>

<p>
Occasionally you'll see code that discards the error value in order
to ignore the error; this is terrible practice. Always check error returns;
they're provided for a reason.
</p>

<pre>
// Bad! This code will crash if path does not exist.
fi, _ := os.Stat(path)
if fi.IsDir() {
    fmt.Printf("%s is a directory\n", path)
}
</pre>

<h3 id="blank_unused">Unused imports and variables</h3>

<p>
It is an error to import a package or to declare a variable without using it.
Unused imports bloat the program and slow compilation,
while a variable that is initialized but not used is at least
a wasted computation and perhaps indicative of a
larger bug.
When a program is under active development, however,
unused imports and variables often arise and it can
be annoying to delete them just to have the compilation proceed,
only to have them be needed again later.
The blank identifier provides a workaround.
</p>
<p>
This half-written program has two unused imports
(<code>fmt</code> and <code>io</code>)
and an unused variable (<code>fd</code>),
so it will not compile, but it would be nice to see if the
code so far is correct.
</p>
{{code "/doc/progs/eff_unused1.go" `/package/` `$`}}
<p>
To silence complaints about the unused imports, use a
blank identifier to refer to a symbol from the imported package.
Similarly, assigning the unused variable <code>fd</code>
to the blank identifier will silence the unused variable error.
This version of the program does compile.
</p>
{{code "/doc/progs/eff_unused2.go" `/package/` `$`}}

<p>
By convention, the global declarations to silence import errors
should come right after the imports and be commented,
both to make them easy to find and as a reminder to clean things up later.
</p>

<h3 id="blank_import">Import for side effect</h3>

<p>
An unused import like <code>fmt</code> or <code>io</code> in the
previous example should eventually be used or removed:
blank assignments identify code as a work in progress.
But sometimes it is useful to import a package only for its
side effects, without any explicit use.
For example, during its <code>init</code> function,
the <code><a href="/pkg/net/http/pprof/">net/http/pprof</a></code>
package registers HTTP handlers that provide
debugging information. It has an exported API, but
most clients need only the handler registration and
access the data through a web page.
To import the package only for its side effects, rename the package
to the blank identifier:
</p>
<pre>
import _ "net/http/pprof"
</pre>
<p>
This form of import makes clear that the package is being
imported for its side effects, because there is no other possible
use of the package: in this file, it doesn't have a name.
(If it did, and we didn't use that name, the compiler would reject the program.)
</p>

<h3 id="blank_implements">Interface checks</h3>

<p>
As we saw in the discussion of <a href="#interfaces_and_types">interfaces</a> above,
a type need not declare explicitly that it implements an interface.
Instead, a type implements the interface just by implementing the interface's methods.
In practice, most interface conversions are static and therefore checked at compile time.
For example, passing an <code>*os.File</code> to a function
expecting an <code>io.Reader</code> will not compile unless
<code>*os.File</code> implements the <code>io.Reader</code> interface.
</p>

<p>
Some interface checks do happen at run-time, though.
One instance is in the <code><a href="/pkg/encoding/json/">encoding/json</a></code>
package, which defines a <code><a href="/pkg/encoding/json/#Marshaler">Marshaler</a></code>
interface. When the JSON encoder receives a value that implements that interface,
the encoder invokes the value's marshaling method to convert it to JSON
instead of doing the standard conversion.
The encoder checks this property at run time with a <a href="#interface_conversions">type assertion</a> like:
</p>

<pre>
m, ok := val.(json.Marshaler)
</pre>

<p>
If it's necessary only to ask whether a type implements an interface, without
actually using the interface itself, perhaps as part of an error check, use the blank
identifier to ignore the type-asserted value:
</p>

<pre>
if _, ok := val.(json.Marshaler); ok {
    fmt.Printf("value %v of type %T implements json.Marshaler\n", val, val)
}
</pre>

<p>
One place this situation arises is when it is necessary to guarantee within the package implementing the type that
it actually satisfies the interface.
If a type—for example,
<code><a href="/pkg/encoding/json/#RawMessage">json.RawMessage</a></code>—needs
a custom JSON representation, it should implement
<code>json.Marshaler</code>, but there are no static conversions that would
cause the compiler to verify this automatically.
If the type inadvertently fails to satisfy the interface, the JSON encoder will still work,
but will not use the custom implementation.
To guarantee that the implementation is correct,
a global declaration using the blank identifier can be used in the package:
</p>
<pre>
var _ json.Marshaler = (*RawMessage)(nil)
</pre>
<p>
In this declaration, the assignment involving a conversion of a
<code>*RawMessage</code> to a <code>Marshaler</code>
requires that <code>*RawMessage</code> implements <code>Marshaler</code>,
and that property will be checked at compile time.
Should the <code>json.Marshaler</code> interface change, this package
will no longer compile and we will be on notice that it needs to be updated.
</p>

<p>
The appearance of the blank identifier in this construct indicates that
the declaration exists only for the type checking,
not to create a variable.
Don't do this for every type that satisfies an interface, though.
By convention, such declarations are only used
when there are no static conversions already present in the code,
which is a rare event.
</p>


# Embedding

<p>
Go does not provide the typical, type-driven notion of subclassing,
but it does have the ability to &ldquo;borrow&rdquo; pieces of an
implementation by <em>embedding</em> types within a struct or
interface.
</p>
<p>
Interface embedding is very simple.
We've mentioned the <code>io.Reader</code> and <code>io.Writer</code> interfaces before;
here are their definitions.
</p>
<pre>
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
</pre>
<p>
The <code>io</code> package also exports several other interfaces
that specify objects that can implement several such methods.
For instance, there is <code>io.ReadWriter</code>, an interface
containing both <code>Read</code> and <code>Write</code>.
We could specify <code>io.ReadWriter</code> by listing the
two methods explicitly, but it's easier and more evocative
to embed the two interfaces to form the new one, like this:
</p>
<pre>
// ReadWriter is the interface that combines the Reader and Writer interfaces.
type ReadWriter interface {
    Reader
    Writer
}
</pre>
<p>
This says just what it looks like: A <code>ReadWriter</code> can do
what a <code>Reader</code> does <em>and</em> what a <code>Writer</code>
does; it is a union of the embedded interfaces.
Only interfaces can be embedded within interfaces.
</p>
<p>
The same basic idea applies to structs, but with more far-reaching
implications.  The <code>bufio</code> package has two struct types,
<code>bufio.Reader</code> and <code>bufio.Writer</code>, each of
which of course implements the analogous interfaces from package
<code>io</code>.
And <code>bufio</code> also implements a buffered reader/writer,
which it does by combining a reader and a writer into one struct
using embedding: it lists the types within the struct
but does not give them field names.
</p>
<pre>
// ReadWriter stores pointers to a Reader and a Writer.
// It implements io.ReadWriter.
type ReadWriter struct {
    *Reader  // *bufio.Reader
    *Writer  // *bufio.Writer
}
</pre>
<p>
The embedded elements are pointers to structs and of course
must be initialized to point to valid structs before they
can be used.
The <code>ReadWriter</code> struct could be written as
</p>
<pre>
type ReadWriter struct {
    reader *Reader
    writer *Writer
}
</pre>
<p>
but then to promote the methods of the fields and to
satisfy the <code>io</code> interfaces, we would also need
to provide forwarding methods, like this:
</p>
<pre>
func (rw *ReadWriter) Read(p []byte) (n int, err error) {
    return rw.reader.Read(p)
}
</pre>
<p>
By embedding the structs directly, we avoid this bookkeeping.
The methods of embedded types come along for free, which means that <code>bufio.ReadWriter</code>
not only has the methods of <code>bufio.Reader</code> and <code>bufio.Writer</code>,
it also satisfies all three interfaces:
<code>io.Reader</code>,
<code>io.Writer</code>, and
<code>io.ReadWriter</code>.
</p>
<p>
There's an important way in which embedding differs from subclassing.  When we embed a type,
the methods of that type become methods of the outer type,
but when they are invoked the receiver of the method is the inner type, not the outer one.
In our example, when the <code>Read</code> method of a <code>bufio.ReadWriter</code> is
invoked, it has exactly the same effect as the forwarding method written out above;
the receiver is the <code>reader</code> field of the <code>ReadWriter</code>, not the
<code>ReadWriter</code> itself.
</p>
<p>
Embedding can also be a simple convenience.
This example shows an embedded field alongside a regular, named field.
</p>
<pre>
type Job struct {
    Command string
    *log.Logger
}
</pre>
<p>
The <code>Job</code> type now has the <code>Print</code>, <code>Printf</code>, <code>Println</code>
and other
methods of <code>*log.Logger</code>.  We could have given the <code>Logger</code>
a field name, of course, but it's not necessary to do so.  And now, once
initialized, we can
log to the <code>Job</code>:
</p>
<pre>
job.Println("starting now...")
</pre>
<p>
The <code>Logger</code> is a regular field of the <code>Job</code> struct,
so we can initialize it in the usual way inside the constructor for <code>Job</code>, like this,
</p>
<pre>
func NewJob(command string, logger *log.Logger) *Job {
    return &amp;Job{command, logger}
}
</pre>
<p>
or with a composite literal,
</p>
<pre>
job := &amp;Job{command, log.New(os.Stderr, "Job: ", log.Ldate)}
</pre>
<p>
If we need to refer to an embedded field directly, the type name of the field,
ignoring the package qualifier, serves as a field name, as it did
in the <code>Read</code> method of our <code>ReadWriter</code> struct.
Here, if we needed to access the
<code>*log.Logger</code> of a <code>Job</code> variable <code>job</code>,
we would write <code>job.Logger</code>,
which would be useful if we wanted to refine the methods of <code>Logger</code>.
</p>
<pre>
func (job *Job) Printf(format string, args ...interface{}) {
    job.Logger.Printf("%q: %s", job.Command, fmt.Sprintf(format, args...))
}
</pre>
<p>
Embedding types introduces the problem of name conflicts but the rules to resolve
them are simple.
First, a field or method <code>X</code> hides any other item <code>X</code> in a more deeply
nested part of the type.
If <code>log.Logger</code> contained a field or method called <code>Command</code>, the <code>Command</code> field
of <code>Job</code> would dominate it.
</p>
<p>
Second, if the same name appears at the same nesting level, it is usually an error;
it would be erroneous to embed <code>log.Logger</code> if the <code>Job</code> struct
contained another field or method called <code>Logger</code>.
However, if the duplicate name is never mentioned in the program outside the type definition, it is OK.
This qualification provides some protection against changes made to types embedded from outside; there
is no problem if a field is added that conflicts with another field in another subtype if neither field
is ever used.
</p>


# Concurrency

## Share by communicating

* Concurrent programming
  * COMMON problem
    * implement correct access -- to -- shared variables
  * Go's approach
    * ⭐️shared values
      * are passed -- through -- channels⭐️
      * NEVER ACTIVELY shared -- by -- separate goroutine
        * == 1! goroutine has access -- to the -- value | ANY given time⭐️
    * slogan
      ```
      Do not communicate by sharing memory;
      instead, share memory by communicating.
      ```  

<p>
This approach can be taken too far.  Reference counts may be best done
by putting a mutex around an integer variable, for instance.  But as a
high-level approach, using channels to control access makes it easier
to write clear, correct programs.
</p>
<p>
One way to think about this model is to consider a typical single-threaded
program running on one CPU. It has no need for synchronization primitives.
Now run another such instance; it too needs no synchronization.  Now let those
two communicate; if the communication is the synchronizer, there's still no need
for other synchronization.  Unix pipelines, for example, fit this model
perfectly.  Although Go's approach to concurrency originates in Hoare's
Communicating Sequential Processes (CSP),
it can also be seen as a type-safe generalization of Unix pipes.
</p>

## Goroutines

* NOT reuse EXISTING terms (threads, coroutines, processes)
  * Reason: 🧠inaccurate connotations🧠

* Goroutine
  * 's model
    * | SAME address space, MULTIPLE goroutines executing concurrently 
    * lightweight
      * Reason: 🧠's costing (lightly)> allocation of stack space🧠
      * stacks
        * start small
        * grow -- by -- allocating (& freeing) heap storage
  * 👀are multiplexed | MULTIPLE OS threads👀
    * == if 1 is blocked (_Example:_ waiting for I/O) -> others continue to run

* `go functionName()` or `go methodName()`
  * run the call | NEW goroutine
    * == | Unix, shell's `&` notation
      * == run a command | background
  * | complete the call,
    * the goroutine exits SILENTLY
  ```
  go list.Sort()
  # list.Sort is run CONCURRENTLY == NOT wait for it 
  ```
  
* function literal
  * == closures
    * == 👀as long as goroutine is active -> variables / referred to by the function, survive👀 
  * uses
    * | goroutine invocation
      ```
      func Announce(message string, delay time.Duration) {
        
        // `delay` & `message` surive TILL this goroutine is active
        go func() {
          time.Sleep(delay)
          fmt.Println(message)
        }()  // Note the parentheses - must call the function.
      }
      ```
  * ❌NO way of signaling completion❌
    * -> 👀use channels👀

## Channels

* `make(chan TypesToSenViaTheChannel, channelBufferSize)`
  * create a channel /
    * == reference -- to an -- underlying data structure 
  * `make` ALSO valid | maps
  * `channelBufferSize`
    * OPTIONAL
    * integer parameter
      * by default, 0,
        * == unbuffered or synchronous channel
  ```
  ci := make(chan int)            // unbuffered channel of integers
  cj := make(chan int, 0)         // unbuffered channel of integers
  cs := make(chan *os.File, 100)  // buffered channel of pointers to Files
  ```

<p>
Unbuffered channels combine communication&mdash;the exchange of a value&mdash;with
synchronization&mdash;guaranteeing that two calculations (goroutines) are in
a known state.
</p>
<p>
There are lots of nice idioms using channels.  Here's one to get us started.
In the previous section we launched a sort in the background. A channel
can allow the launching goroutine to wait for the sort to complete.
</p>
<pre>
c := make(chan int)  // Allocate a channel.
// Start the sort in a goroutine; when it completes, signal on the channel.
go func() {
    list.Sort()
    c &lt;- 1  // Send a signal; value does not matter.
}()
doSomethingForAWhile()
&lt;-c   // Wait for sort to finish; discard sent value.
</pre>
<p>
Receivers always block until there is data to receive.
If the channel is unbuffered, the sender blocks until the receiver has
received the value.
If the channel has a buffer, the sender blocks only until the
value has been copied to the buffer; if the buffer is full, this
means waiting until some receiver has retrieved a value.
</p>
<p>
A buffered channel can be used like a semaphore, for instance to
limit throughput.  In this example, incoming requests are passed
to <code>handle</code>, which sends a value into the channel, processes
the request, and then receives a value from the channel
to ready the &ldquo;semaphore&rdquo; for the next consumer.
The capacity of the channel buffer limits the number of
simultaneous calls to <code>process</code>.
</p>
<pre>
var sem = make(chan int, MaxOutstanding)

func handle(r *Request) {
    sem &lt;- 1    // Wait for active queue to drain.
    process(r)  // May take a long time.
    &lt;-sem       // Done; enable next request to run.
}

func Serve(queue chan *Request) {
    for {
        req := &lt;-queue
        go handle(req)  // Don't wait for handle to finish.
    }
}
</pre>

<p>
Once <code>MaxOutstanding</code> handlers are executing <code>process</code>,
any more will block trying to send into the filled channel buffer,
until one of the existing handlers finishes and receives from the buffer.
</p>

<p>
This design has a problem, though: <code>Serve</code>
creates a new goroutine for
every incoming request, even though only <code>MaxOutstanding</code>
of them can run at any moment.
As a result, the program can consume unlimited resources if the requests come in too fast.
We can address that deficiency by changing <code>Serve</code> to
gate the creation of the goroutines:
</p>

<pre>
func Serve(queue chan *Request) {
    for req := range queue {
        sem &lt;- 1
        go func() {
            process(req)
            &lt;-sem
        }()
    }
}</pre>

<p>
(Note that in Go versions before 1.22 this code has a bug: the loop
variable is shared across all goroutines.
See the <a href="/wiki/LoopvarExperiment">Go wiki</a> for details.)
</p>

<p>
Another approach that manages resources well is to start a fixed
number of <code>handle</code> goroutines all reading from the request
channel.
The number of goroutines limits the number of simultaneous
calls to <code>process</code>.
This <code>Serve</code> function also accepts a channel on which
it will be told to exit; after launching the goroutines it blocks
receiving from that channel.
</p>

<pre>
func handle(queue chan *Request) {
    for r := range queue {
        process(r)
    }
}

func Serve(clientRequests chan *Request, quit chan bool) {
    // Start handlers
    for i := 0; i &lt; MaxOutstanding; i++ {
        go handle(clientRequests)
    }
    &lt;-quit  // Wait to be told to exit.
}
</pre>

## Channels of channels
<p>
One of the most important properties of Go is that
a channel is a first-class value that can be allocated and passed
around like any other.  A common use of this property is
to implement safe, parallel demultiplexing.
</p>
<p>
In the example in the previous section, <code>handle</code> was
an idealized handler for a request but we didn't define the
type it was handling.  If that type includes a channel on which
to reply, each client can provide its own path for the answer.
Here's a schematic definition of type <code>Request</code>.
</p>
<pre>
type Request struct {
    args        []int
    f           func([]int) int
    resultChan  chan int
}
</pre>
<p>
The client provides a function and its arguments, as well as
a channel inside the request object on which to receive the answer.
</p>
<pre>
func sum(a []int) (s int) {
    for _, v := range a {
        s += v
    }
    return
}

request := &amp;Request{[]int{3, 4, 5}, sum, make(chan int)}
// Send request
clientRequests &lt;- request
// Wait for response.
fmt.Printf("answer: %d\n", &lt;-request.resultChan)
</pre>
<p>
On the server side, the handler function is the only thing that changes.
</p>
<pre>
func handle(queue chan *Request) {
    for req := range queue {
        req.resultChan &lt;- req.f(req.args)
    }
}
</pre>
<p>
There's clearly a lot more to do to make it realistic, but this
code is a framework for a rate-limited, parallel, non-blocking RPC
system, and there's not a mutex in sight.
</p>

## Parallelization
<p>
Another application of these ideas is to parallelize a calculation
across multiple CPU cores.  If the calculation can be broken into
separate pieces that can execute independently, it can be parallelized,
with a channel to signal when each piece completes.
</p>
<p>
Let's say we have an expensive operation to perform on a vector of items,
and that the value of the operation on each item is independent,
as in this idealized example.
</p>
<pre>
type Vector []float64

// Apply the operation to v[i], v[i+1] ... up to v[n-1].
func (v Vector) DoSome(i, n int, u Vector, c chan int) {
    for ; i &lt; n; i++ {
        v[i] += u.Op(v[i])
    }
    c &lt;- 1    // signal that this piece is done
}
</pre>
<p>
We launch the pieces independently in a loop, one per CPU.
They can complete in any order but it doesn't matter; we just
count the completion signals by draining the channel after
launching all the goroutines.
</p>
<pre>
const numCPU = 4 // number of CPU cores

func (v Vector) DoAll(u Vector) {
    c := make(chan int, numCPU)  // Buffering optional but sensible.
    for i := 0; i &lt; numCPU; i++ {
        go v.DoSome(i*len(v)/numCPU, (i+1)*len(v)/numCPU, u, c)
    }
    // Drain the channel.
    for i := 0; i &lt; numCPU; i++ {
        &lt;-c    // wait for one task to complete
    }
    // All done.
}
</pre>
<p>
Rather than create a constant value for numCPU, we can ask the runtime what
value is appropriate.
The function <code><a href="/pkg/runtime#NumCPU">runtime.NumCPU</a></code>
returns the number of hardware CPU cores in the machine, so we could write
</p>
<pre>
var numCPU = runtime.NumCPU()
</pre>
<p>
There is also a function
<code><a href="/pkg/runtime#GOMAXPROCS">runtime.GOMAXPROCS</a></code>,
which reports (or sets)
the user-specified number of cores that a Go program can have running
simultaneously.
It defaults to the value of <code>runtime.NumCPU</code> but can be
overridden by setting the similarly named shell environment variable
or by calling the function with a positive number.  Calling it with
zero just queries the value.
Therefore if we want to honor the user's resource request, we should write
</p>
<pre>
var numCPU = runtime.GOMAXPROCS(0)
</pre>
<p>
Be sure not to confuse the ideas of concurrency—structuring a program
as independently executing components—and parallelism—executing
calculations in parallel for efficiency on multiple CPUs.
Although the concurrency features of Go can make some problems easy
to structure as parallel computations, Go is a concurrent language,
not a parallel one, and not all parallelization problems fit Go's model.
For a discussion of the distinction, see the talk cited in
<a href="/blog/concurrency-is-not-parallelism">this
blog post</a>.

## A leaky buffer

<p>
The tools of concurrent programming can even make non-concurrent
ideas easier to express.  Here's an example abstracted from an RPC
package.  The client goroutine loops receiving data from some source,
perhaps a network.  To avoid allocating and freeing buffers, it keeps
a free list, and uses a buffered channel to represent it.  If the
channel is empty, a new buffer gets allocated.
Once the message buffer is ready, it's sent to the server on
<code>serverChan</code>.
</p>
<pre>
var freeList = make(chan *Buffer, 100)
var serverChan = make(chan *Buffer)

func client() {
    for {
        var b *Buffer
        // Grab a buffer if available; allocate if not.
        select {
        case b = &lt;-freeList:
            // Got one; nothing more to do.
        default:
            // None free, so allocate a new one.
            b = new(Buffer)
        }
        load(b)              // Read next message from the net.
        serverChan &lt;- b      // Send to server.
    }
}
</pre>
<p>
The server loop receives each message from the client, processes it,
and returns the buffer to the free list.
</p>
<pre>
func server() {
    for {
        b := &lt;-serverChan    // Wait for work.
        process(b)
        // Reuse buffer if there's room.
        select {
        case freeList &lt;- b:
            // Buffer on free list; nothing more to do.
        default:
            // Free list full, just carry on.
        }
    }
}
</pre>
<p>
The client attempts to retrieve a buffer from <code>freeList</code>;
if none is available, it allocates a fresh one.
The server's send to <code>freeList</code> puts <code>b</code> back
on the free list unless the list is full, in which case the
buffer is dropped on the floor to be reclaimed by
the garbage collector.
(The <code>default</code> clauses in the <code>select</code>
statements execute when no other case is ready,
meaning that the <code>selects</code> never block.)
This implementation builds a leaky bucket free list
in just a few lines, relying on the buffered channel and
the garbage collector for bookkeeping.
</p>

# Errors

<p>
Library routines must often return some sort of error indication to
the caller.
As mentioned earlier, Go's multivalue return makes it
easy to return a detailed error description alongside the normal
return value.
It is good style to use this feature to provide detailed error information.
For example, as we'll see, <code>os.Open</code> doesn't
just return a <code>nil</code> pointer on failure, it also returns an
error value that describes what went wrong.
</p>

<p>
By convention, errors have type <code>error</code>,
a simple built-in interface.
</p>
<pre>
type error interface {
    Error() string
}
</pre>
<p>
A library writer is free to implement this interface with a
richer model under the covers, making it possible not only
to see the error but also to provide some context.
As mentioned, alongside the usual <code>*os.File</code>
return value, <code>os.Open</code> also returns an
error value.
If the file is opened successfully, the error will be <code>nil</code>,
but when there is a problem, it will hold an
<code>os.PathError</code>:
</p>
<pre>
// PathError records an error and the operation and
// file path that caused it.
type PathError struct {
    Op string    // "open", "unlink", etc.
    Path string  // The associated file.
    Err error    // Returned by the system call.
}

func (e *PathError) Error() string {
    return e.Op + " " + e.Path + ": " + e.Err.Error()
}
</pre>
<p>
<code>PathError</code>'s <code>Error</code> generates
a string like this:
</p>
<pre>
open /etc/passwx: no such file or directory
</pre>
<p>
Such an error, which includes the problematic file name, the
operation, and the operating system error it triggered, is useful even
if printed far from the call that caused it;
it is much more informative than the plain
"no such file or directory".
</p>

<p>
When feasible, error strings should identify their origin, such as by having
a prefix naming the operation or package that generated the error.  For example, in package
<code>image</code>, the string representation for a decoding error due to an
unknown format is "image: unknown format".
</p>

<p>
Callers that care about the precise error details can
use a type switch or a type assertion to look for specific
errors and extract details.  For <code>PathErrors</code>
this might include examining the internal <code>Err</code>
field for recoverable failures.
</p>

<pre>
for try := 0; try &lt; 2; try++ {
    file, err = os.Create(filename)
    if err == nil {
        return
    }
    if e, ok := err.(*os.PathError); ok &amp;&amp; e.Err == syscall.ENOSPC {
        deleteTempFiles()  // Recover some space.
        continue
    }
    return
}
</pre>

<p>
The second <code>if</code> statement here is another <a href="#interface_conversions">type assertion</a>.
If it fails, <code>ok</code> will be false, and <code>e</code>
will be <code>nil</code>.
If it succeeds,  <code>ok</code> will be true, which means the
error was of type <code>*os.PathError</code>, and then so is <code>e</code>,
which we can examine for more information about the error.
</p>

<h3 id="panic">Panic</h3>

<p>
The usual way to report an error to a caller is to return an
<code>error</code> as an extra return value.  The canonical
<code>Read</code> method is a well-known instance; it returns a byte
count and an <code>error</code>.  But what if the error is
unrecoverable?  Sometimes the program simply cannot continue.
</p>

<p>
For this purpose, there is a built-in function <code>panic</code>
that in effect creates a run-time error that will stop the program
(but see the next section).  The function takes a single argument
of arbitrary type&mdash;often a string&mdash;to be printed as the
program dies.  It's also a way to indicate that something impossible has
happened, such as exiting an infinite loop.
</p>


<pre>
// A toy implementation of cube root using Newton's method.
func CubeRoot(x float64) float64 {
    z := x/3   // Arbitrary initial value
    for i := 0; i &lt; 1e6; i++ {
        prevz := z
        z -= (z*z*z-x) / (3*z*z)
        if veryClose(z, prevz) {
            return z
        }
    }
    // A million iterations has not converged; something is wrong.
    panic(fmt.Sprintf("CubeRoot(%g) did not converge", x))
}
</pre>

<p>
This is only an example but real library functions should
avoid <code>panic</code>.  If the problem can be masked or worked
around, it's always better to let things continue to run rather
than taking down the whole program.  One possible counterexample
is during initialization: if the library truly cannot set itself up,
it might be reasonable to panic, so to speak.
</p>

<pre>
var user = os.Getenv("USER")

func init() {
    if user == "" {
        panic("no value for $USER")
    }
}
</pre>

<h3 id="recover">Recover</h3>

<p>
When <code>panic</code> is called, including implicitly for run-time
errors such as indexing a slice out of bounds or failing a type
assertion, it immediately stops execution of the current function
and begins unwinding the stack of the goroutine, running any deferred
functions along the way.  If that unwinding reaches the top of the
goroutine's stack, the program dies.  However, it is possible to
use the built-in function <code>recover</code> to regain control
of the goroutine and resume normal execution.
</p>

<p>
A call to <code>recover</code> stops the unwinding and returns the
argument passed to <code>panic</code>.  Because the only code that
runs while unwinding is inside deferred functions, <code>recover</code>
is only useful inside deferred functions.
</p>

<p>
One application of <code>recover</code> is to shut down a failing goroutine
inside a server without killing the other executing goroutines.
</p>

<pre>
func server(workChan &lt;-chan *Work) {
    for work := range workChan {
        go safelyDo(work)
    }
}

func safelyDo(work *Work) {
    defer func() {
        if err := recover(); err != nil {
            log.Println("work failed:", err)
        }
    }()
    do(work)
}
</pre>

<p>
In this example, if <code>do(work)</code> panics, the result will be
logged and the goroutine will exit cleanly without disturbing the
others.  There's no need to do anything else in the deferred closure;
calling <code>recover</code> handles the condition completely.
</p>

<p>
Because <code>recover</code> always returns <code>nil</code> unless called directly
from a deferred function, deferred code can call library routines that themselves
use <code>panic</code> and <code>recover</code> without failing.  As an example,
the deferred function in <code>safelyDo</code> might call a logging function before
calling <code>recover</code>, and that logging code would run unaffected
by the panicking state.
</p>

<p>
With our recovery pattern in place, the <code>do</code>
function (and anything it calls) can get out of any bad situation
cleanly by calling <code>panic</code>.  We can use that idea to
simplify error handling in complex software.  Let's look at an
idealized version of a <code>regexp</code> package, which reports
parsing errors by calling <code>panic</code> with a local
error type.  Here's the definition of <code>Error</code>,
an <code>error</code> method, and the <code>Compile</code> function.
</p>

<pre>
// Error is the type of a parse error; it satisfies the error interface.
type Error string
func (e Error) Error() string {
    return string(e)
}

// error is a method of *Regexp that reports parsing errors by
// panicking with an Error.
func (regexp *Regexp) error(err string) {
    panic(Error(err))
}

// Compile returns a parsed representation of the regular expression.
func Compile(str string) (regexp *Regexp, err error) {
    regexp = new(Regexp)
    // doParse will panic if there is a parse error.
    defer func() {
        if e := recover(); e != nil {
            regexp = nil    // Clear return value.
            err = e.(Error) // Will re-panic if not a parse error.
        }
    }()
    return regexp.doParse(str), nil
}
</pre>

<p>
If <code>doParse</code> panics, the recovery block will set the
return value to <code>nil</code>&mdash;deferred functions can modify
named return values.  It will then check, in the assignment
to <code>err</code>, that the problem was a parse error by asserting
that it has the local type <code>Error</code>.
If it does not, the type assertion will fail, causing a run-time error
that continues the stack unwinding as though nothing had interrupted
it.
This check means that if something unexpected happens, such
as an index out of bounds, the code will fail even though we
are using <code>panic</code> and <code>recover</code> to handle
parse errors.
</p>

<p>
With error handling in place, the <code>error</code> method (because it's a
method bound to a type, it's fine, even natural, for it to have the same name
as the builtin <code>error</code> type)
makes it easy to report parse errors without worrying about unwinding
the parse stack by hand:
</p>

<pre>
if pos == 0 {
    re.error("'*' illegal at start of expression")
}
</pre>

<p>
Useful though this pattern is, it should be used only within a package.
<code>Parse</code> turns its internal <code>panic</code> calls into
<code>error</code> values; it does not expose <code>panics</code>
to its client.  That is a good rule to follow.
</p>

<p>
By the way, this re-panic idiom changes the panic value if an actual
error occurs.  However, both the original and new failures will be
presented in the crash report, so the root cause of the problem will
still be visible.  Thus this simple re-panic approach is usually
sufficient&mdash;it's a crash after all&mdash;but if you want to
display only the original value, you can write a little more code to
filter unexpected problems and re-panic with the original error.
That's left as an exercise for the reader.
</p>


# A web server

<p>
Let's finish with a complete Go program, a web server.
This one is actually a kind of web re-server.
Google provides a service at <code>chart.apis.google.com</code>
that does automatic formatting of data into charts and graphs.
It's hard to use interactively, though,
because you need to put the data into the URL as a query.
The program here provides a nicer interface to one form of data: given a short piece of text,
it calls on the chart server to produce a QR code, a matrix of boxes that encode the
text.
That image can be grabbed with your cell phone's camera and interpreted as,
for instance, a URL, saving you typing the URL into the phone's tiny keyboard.
</p>
<p>
Here's the complete program.
An explanation follows.
</p>
{{code "/doc/progs/eff_qr.go" `/package/` `$`}}
<p>
The pieces up to <code>main</code> should be easy to follow.
The one flag sets a default HTTP port for our server.  The template
variable <code>templ</code> is where the fun happens. It builds an HTML template
that will be executed by the server to display the page; more about
that in a moment.
</p>
<p>
The <code>main</code> function parses the flags and, using the mechanism
we talked about above, binds the function <code>QR</code> to the root path
for the server.  Then <code>http.ListenAndServe</code> is called to start the
server; it blocks while the server runs.
</p>
<p>
<code>QR</code> just receives the request, which contains form data, and
executes the template on the data in the form value named <code>s</code>.
</p>
<p>
The template package <code>html/template</code> is powerful;
this program just touches on its capabilities.
In essence, it rewrites a piece of HTML text on the fly by substituting elements derived
from data items passed to <code>templ.Execute</code>, in this case the
form value.
Within the template text (<code>templateStr</code>),
double-brace-delimited pieces denote template actions.
The piece from <code>{{"{{if .}}"}}</code>
to <code>{{"{{end}}"}}</code> executes only if the value of the current data item, called <code>.</code> (dot),
is non-empty.
That is, when the string is empty, this piece of the template is suppressed.
</p>
<p>
The two snippets <code>{{"{{.}}"}}</code> say to show the data presented to
the template—the query string—on the web page.
The HTML template package automatically provides appropriate escaping so the
text is safe to display.
</p>
<p>
The rest of the template string is just the HTML to show when the page loads.
If this is too quick an explanation, see the <a href="/pkg/html/template/">documentation</a>
for the template package for a more thorough discussion.
</p>
<p>
And there you have it: a useful web server in a few lines of code plus some
data-driven HTML text.
Go is powerful enough to make a lot happen in a few lines.
</p>

<!--
TODO
<pre>
verifying implementation
type Color uint32

// Check that Color implements image.Color and image.Image
var _ image.Color = Black
var _ image.Image = Black
</pre>
-->
