
CXXFLAGS += -std=c++11 -I$(HOME)/git/libslankdev
LDFLAGS +=

SRC = main.cc
OBJ = $(SRC:.cc=.o)
TARGET = a.out

def:
	$(CXX) $(CXXFLAGS) -o edenman_chikuwa.out edenman_chikuwa.cc $(LDFLAGS)
	$(CXX) $(CXXFLAGS) -o ntt_ipa.out ntt_ipa.cc $(LDFLAGS)
	docker cp ntt_ipa.out R8:/usr/bin/ntt_ipa
	docker cp edenman_chikuwa.out R9:/usr/bin/edenman_chikuwa


all: $(OBJ)
	$(CXX) $(CXXFLAGS) -o $(TARGET) $(OBJ) $(LDFLAGS)

clean:
	rm -f $(OBJ) $(TARGET)

install_docker: all
	for c in `docker ps --format "{{.Names}}"`; do\
		docker cp a.out $$c:/usr/bin/filter ; done
