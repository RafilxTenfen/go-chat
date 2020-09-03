
bot_SOURCES=$(shell find ./cmd/bot -name "*.go")

bot: build/$(OS)-$(ARCH)/bot
	cp $< $@

build/linux-amd64/bot: $(bot_SOURCES) $(SOURCES)
	env GOTRACEBACK=none CGO_ENABLED=0 GOARCH=amd64 GOOS=linux $(GOBUILD) $(FLAGS) $(LDFLAGS) -o $@ ./cmd/bot

build/darwin-amd64/bot: $(bot_SOURCES) $(SOURCES)
	env GOTRACEBACK=none CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin $(GOBUILD) $(FLAGS) $(LDFLAGS) -o $@ ./cmd/bot

build/linux-arm7/bot: $(bot_SOURCES) $(SOURCES)
	env GOTRACEBACK=none CGO_ENABLED=0 GOARM=7 GOARCH=arm GOOS=linux $(GOBUILD) $(FLAGS) $(LDFLAGS) -o $@ ./cmd/bot
user_SOURCES=$(shell find ./cmd/user -name "*.go")

user: build/$(OS)-$(ARCH)/user
	cp $< $@

build/linux-amd64/user: $(user_SOURCES) $(SOURCES)
	env GOTRACEBACK=none CGO_ENABLED=0 GOARCH=amd64 GOOS=linux $(GOBUILD) $(FLAGS) $(LDFLAGS) -o $@ ./cmd/user

build/darwin-amd64/user: $(user_SOURCES) $(SOURCES)
	env GOTRACEBACK=none CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin $(GOBUILD) $(FLAGS) $(LDFLAGS) -o $@ ./cmd/user

build/linux-arm7/user: $(user_SOURCES) $(SOURCES)
	env GOTRACEBACK=none CGO_ENABLED=0 GOARM=7 GOARCH=arm GOOS=linux $(GOBUILD) $(FLAGS) $(LDFLAGS) -o $@ ./cmd/user
chatapi_SOURCES=$(shell find ./cmd/chatapi -name "*.go")

chatapi: build/$(OS)-$(ARCH)/chatapi
	cp $< $@

build/linux-amd64/chatapi: $(chatapi_SOURCES) $(SOURCES)
	env GOTRACEBACK=none CGO_ENABLED=0 GOARCH=amd64 GOOS=linux $(GOBUILD) $(FLAGS) $(LDFLAGS) -o $@ ./cmd/chatapi

build/darwin-amd64/chatapi: $(chatapi_SOURCES) $(SOURCES)
	env GOTRACEBACK=none CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin $(GOBUILD) $(FLAGS) $(LDFLAGS) -o $@ ./cmd/chatapi

build/linux-arm7/chatapi: $(chatapi_SOURCES) $(SOURCES)
	env GOTRACEBACK=none CGO_ENABLED=0 GOARM=7 GOARCH=arm GOOS=linux $(GOBUILD) $(FLAGS) $(LDFLAGS) -o $@ ./cmd/chatapi
dist/linux-amd64.tar.gz: $(addprefix build/linux-amd64/,$(TARGETS))
	mkdir -p dist
	tar -czf $@ -C build/linux-amd64 .
dist/darwin-amd64.tar.gz: $(addprefix build/darwin-amd64/,$(TARGETS))
	mkdir -p dist
	tar -czf $@ -C build/darwin-amd64 .
dist/linux-arm7.tar.gz: $(addprefix build/linux-arm7/,$(TARGETS))
	mkdir -p dist
	tar -czf $@ -C build/linux-arm7 .
