include $(GOROOT)/src/Make.inc

TARG=sequences

GOFILES=\
	sequences.go\
	length.go\
	each.go\
	while.go\
	reduce.go\

include $(GOROOT)/src/Make.pkg