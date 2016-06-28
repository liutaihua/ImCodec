CXXFLAGS=-g

bin = imcodec
objs = main.o encoder.o

$(bin): $(objs)
	g++ -o $(bin) $(objs) -lpthread

.phony: clean

clean:
	rm $(objs) $(bin)
