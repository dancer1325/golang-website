- Threading models
    - thread1 ← communicate sharing memory with → thread2
    - *example*
        - Java
        - C++
        - Python
- Shared data structures — are protected by — locks
    - access the data
        - — via — threads over the locks
        - thread-safe -- Example: Python's queue --
- Go concurrency
    - primitives
        - goroutines
        - channels
    - origins  -- Check '../blog/Bell Labs and CSP Threads' --
    - 👁️NOT locks 👁️
    - 1! goroutine has got access to the data / given time
- Check '../doc/codewalk'