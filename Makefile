watch:
	@air -c .air.conf

migrate-up:
	@sql-migrate up

migrate-down:
	@sql-migrate down