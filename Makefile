include $(GOROOT)/src/Make.inc

TARG=sequences

GOFILES=\
	sequences.go\
	length.go\
	indexable.go\
	each.go\
	while.go\
	predicates.go\
	reduce.go\

include $(GOROOT)/src/Make.pkg