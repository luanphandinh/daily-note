install:
	go build -o ./build/note -v .
	chmod +x ./build/note
	grep -q "alias note=$(shell pwd)/build/note" ~/.bash_profile 2>/dev/null || echo "alias note=$(shell pwd)/build/note" >> ~/.bash_profile
	grep -q "alias note=$(shell pwd)/build/note" ~/.bashrc 2>/dev/null || echo "alias note=$(shell pwd)/build/note" >> ~/.bashrc
	grep -q "alias note=$(shell pwd)/build/note" ~/.zshrc 2>/dev/null || echo "alias note=$(shell pwd)/build/note" >> ~/.zshrc
