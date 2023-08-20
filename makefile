# Ime izlaznog binarnog fajla
BINARY_NAME=myapp

# Putanja do Go izvornih fajlova
SRC_PATH=./src

# Flags za kompajliranje (opciono)
BUILD_FLAGS=-ldflags "-s -w"

# Pravilo za build
build:
	go build $(BUILD_FLAGS) -o $(BINARY_NAME) $(SRC_PATH)

# Pravilo za clean
clean:
	if [ -f $(BINARY_NAME) ] ; then rm $(BINARY_NAME) ; fi

# Pravilo za run
run:
	go run $(SRC_PATH)

# Pravilo za testiranje
test:
	go test $(SRC_PATH)/...

# Pravilo za formatiranje koda
fmt:
	go fmt $(SRC_PATH)/...

# Pravilo za pretragu i zamenu
replace:
	find $(SRC_PATH) -type f -name '*.go' -exec sed -i 's/old_string/new_string/g' {} +

# Pravilo za generisanje dokumentacije
docs:
	godoc -http=:6060 -index

# Podrazumevano pravilo (build)
default: build