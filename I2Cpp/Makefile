CC = g++
LDFLAGS = -lwiringPi
CPPFLAGS = -O -std=c++14 -pthread -Wall

TARGETS = CheckHDSP SimpleClockHDSP ResetHDSP DisplayHDSP DateHDSP TwoDisplayHDSP CheckInputs ThreadTest
TARGETSRC = $(addsuffix .c,$(TARGETS))
LIBS = SetupHDSP WriteCharacter ScrollText ClockModule DateModule SetupInputs WriteText
LIBSRCS = $(addsuffix .c,$(LIBS))
LIBOBJS = $(addsuffix .o,$(LIBS))
BASE = ../$(notdir $(shell pwd))


all: $(LIBOBJS) $(TARGETS)


%.o : %.c HDSP.h Inputs.h
	$(CC) $(CPPFLAGS) -c $*.c

% : %.c $(LIBOBJS)
	$(CC) $(CPPFLAGS) -o $@ $< $(LIBOBJS) $(LDFLAGS)


.PHONEY: clean package

clean:
	rm -f $(TARGETS) *.o

package:
	tar cvzf ./HDSPTools.tgz $(BASE)/*.c $(BASE)/*.h $(BASE)/*.rules $(BASE)/*.conf $(BASE)/README.txt $(BASE)/Makefile
