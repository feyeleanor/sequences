include $(GOROOT)/src/Make.inc

TARG=gosequence

GOFILES=\
	gosequence.go\
	length.go\
	each.go\
	while.go\
	reduce.go\

include $(GOROOT)/src/Make.pkg