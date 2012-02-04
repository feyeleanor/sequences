include $(GOROOT)/src/Make.inc

TARG=sequences

GOFILES=\
	sequences.go\
	length.go\
	indexable.go\
	step.go\
	each.go\
	count.go\
	while.go\
	predicates.go\
	reduce.go\

include $(GOROOT)/src/Make.pkg