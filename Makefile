# Source files
SOURCE = demo.go

# Target excutables
TARGET = demo.out

# Compiler to use
GC = go build

all: $(TARGET)

$(TARGET):
	$(GC) -o $(TARGET) $(SOURCE)
