include $(GOROOT)/src/Make.inc

TARG=gosequence

GOFILES=\
	gosequence.go\
	each.go\
	while.go\
	reduce.go\

include $(GOROOT)/src/Make.pkg