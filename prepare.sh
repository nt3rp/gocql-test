#!/bin/sh
echo "Delete Test Keyspace"
ccm node1 cqlsh -x "DROP KEYSPACE IF EXISTS \"experiment\""

echo "Create Test Keyspace"
ccm node1 cqlsh -x "CREATE KEYSPACE IF NOT EXISTS \"experiment\" WITH REPLICATION = {'class': 'SimpleStrategy', 'replication_factor' : 2}"

echo "Create Simple Table"
ccm node1 cqlsh -x "CREATE TABLE experiment.simple (k int PRIMARY KEY, s text, i int)"

echo "Create Complex Table"
ccm node1 cqlsh -x "CREATE TABLE experiment.complex (k int PRIMARY KEY, s text, l list<text>)"