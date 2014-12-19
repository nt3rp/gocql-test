#!/bin/sh

echo "Delete Test Keyspace"
cqlsh -e "DROP KEYSPACE IF EXISTS \"experiment\""

echo "Create Test Keyspace"
cqlsh -e "CREATE KEYSPACE IF NOT EXISTS \"experiment\" WITH REPLICATION = {'class': 'SimpleStrategy', 'replication_factor' : 2};"

echo "Create Simple Table"
cqlsh -e "CREATE TABLE experiment.simple (k int PRIMARY KEY, s text, i int)"

echo "Create Complex Table"
cqlsh -e "CREATE TABLE experiment.complex (k int PRIMARY KEY, s text, l list<text>)"
