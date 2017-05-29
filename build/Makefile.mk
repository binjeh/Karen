.EXPORT_ALL_VARIABLES:

.NOTPARALLEL:

.PHONY: __CATCHALL__
__CATCHALL__:
	@echo ""
	@echo "+--------------------------------------------------------------------------+"
	@echo "| Please execute explicit job names instead of typing 'make' without args. |"
	@echo "|            Use 'make help' to view a list of available tasks.            |"
	@echo "+--------------------------------------------------------------------------+"
	@echo ""
	exit 1
