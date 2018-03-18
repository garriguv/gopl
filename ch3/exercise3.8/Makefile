GO_FILES:=$(wildcard *.go)
NAME:=exercise3.8

X:=0.36
Y:=0.6
X_RANGE:=0.005
Y_RANGE:=0.005
WIDTH:=200
HEIGHT:=200

FORMATS:=complex128 complex64 big.Float big.Rat
PNG_FILES:=$(FORMATS:=.png)

.PHONY: images
images: $(PNG_FILES) $(NAME)

$(NAME): $(GO_FILES)
	go build

%.png:
	./$(NAME) --precision $(basename $@) \
		--width $(WIDTH) --height $(HEIGHT) \
		--x $(X) --y $(Y) \
		--xRange $(X_RANGE) --yRange $(Y_RANGE) \
		> $@
