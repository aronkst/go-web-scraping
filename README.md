# Go Web Scraping

This is a simple web scraping developed using the Go programming language that has the possibility to fetch information from websites rendered via JavaScript.

Its execution is done using Docker, where Chromium it's installed inside the container, which can be used to render the page, or not, and use the traditional way that takes the HTML of the page with a simple request.

Searching for site information can be passed easily via `body`, in a POST request.

There is also a way to return all the HTML from a page by making a request to a specific path.

To know more details, there is the Makefile file with the commands necessary for its execution and example of how to use, including `curl` commands.

# JavaScript Render

If the site you are fetching the data from has JavaScript rendering, use the `javascript` parameter of the `body` with the value of `true`.

# How to use

Builds the Docker application, must be the first command to be executed:

```
make build
```

Run the application:

```
make run
```

Start the application:

```
make start
```

Stop the application:

```
make stop
```

Example of how to fetch information within a web page using a URL as input.

```
make test-find
```

Example of how to fetch for information within a web page using HTML as input.

```
make test-find-html
```

Example of how to get all the HTML from a web page using a URL as input.

```
make test-html
```
