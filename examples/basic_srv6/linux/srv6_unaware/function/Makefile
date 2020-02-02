
CXXFLAGS += -std=c++11 -I$(HOME)/git/libslankdev
LDFLAGS +=

SRC = main.cc
OBJ = $(SRC:.cc=.o)
TARGET = a.out

def: install_docker

all: $(OBJ)
	$(CXX) $(CXXFLAGS) -o $(TARGET) $(OBJ) $(LDFLAGS)

clean:
	rm -f $(OBJ) $(TARGET)

install_docker: all
	docker cp a.out F1:/
