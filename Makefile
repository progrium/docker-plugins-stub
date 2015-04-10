NAME=docker-plugins-stub

dev:
	@docker history $(NAME):dev &> /dev/null \
		|| docker build -f Dockerfile -t $(NAME):dev .
	@docker run --rm --plugin \
		-v $(PWD):/go/src/github.com/progrium/docker-plugins-stub \
		-p 8000:8000 \
		$(NAME):dev
