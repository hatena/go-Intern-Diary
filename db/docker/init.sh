#!/bin/bash
set -xe

mysqladmin -uroot create intern_diary
mysqladmin -uroot create intern_diary_test

mysql -uroot intern_diary < /app/db/schema.sql
mysql -uroot intern_diary_test < /app/db/schema.sql
