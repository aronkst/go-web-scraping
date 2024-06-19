# Go Web Scraping

This open-source project offers a straightforward web scraping solution developed in the Go programming language. It is designed to fetch information from websites, including those rendered with JavaScript.

The project utilizes Docker for execution, with Chromium installed inside the container. This setup allows you to render web pages using Chromium or opt for the traditional method of fetching HTML through simple HTTP requests.

Information retrieval from websites is facilitated by sending a POST request with the desired parameters in the `body`. Additionally, you can obtain the complete HTML of a page by making a request to a specific endpoint.

For detailed instructions on execution and usage, refer to the Makefile included in the project. The Makefile contains all necessary commands, along with examples and `curl` commands for practical application.

# JavaScript Render

If the site you are fetching the data from has JavaScript rendering, use the `javascript` parameter of the `body` with the value of `true`.

# How to use

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

View logs for Docker containers:

```
make logs
```

Example of how to fetch information within a web page using a URL as input:

```
make test-find
```

Example of how to fetch for information within a web page using HTML as input:

```
make test-find-html
```

Example of how to get all the HTML from a web page using a URL as input:

```
make test-html
```
