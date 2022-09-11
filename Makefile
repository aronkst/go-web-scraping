run:
	docker compose up

start:
	docker compose up -d

stop:
	docker compose stop

logs:
	docker compose logs -f

test-find:
	curl --request POST --url 'http://localhost:3000/find' --header 'Content-Type: application/json' --data '{"url": "https://en.wikipedia.org/wiki/Main_Page","javascript": false,"find": [{"name": "text","class": "div#mp-welcome h1 span","attribute": ""},{"name": "href","class": "div#mp-welcome h1 span a","attribute": "href"},{"name": "list","class": "ul#footer-places li","find": [{"name": "text","class": "a","attribute": ""},{"name": "href","class": "a","attribute": "href"}]}]}'

test-find-html:
	curl --request POST --url 'http://localhost:3000/find' --header 'Content-Type: application/json' --data '{"html": "<h1 title=\"title-test\">H1 Value</h1>","find": [{"name": "text","class": "h1","attribute": ""},{"name": "title","class": "h1","attribute": "title"}]}'

test-html:
	curl --request POST --url 'http://localhost:3000/html' --header 'Content-Type: application/json' --data '{"url": "https://en.wikipedia.org/wiki/Main_Page","javascript": false}'
