#!/bin/bash

# see if we have some basic configuration files available
if [ ! -f /opt/apache-storm/conf/storm.yaml ]
then
	echo "Starting with an empty storm configuration file"
	touch /opt/apache-storm/conf/storm.yaml
fi
if [ ! -f /opt/apache-storm/logback/cluster.xml ]
then
	echo "No cluster.xml found, creating one."
	cp /opt/apache-storm/logback-dist/cluster-console.xml /opt/apache-storm/logback/cluster.xml
fi

# check if we want to autoconfigure zookeeper

if [ "x$CONFIGURE_ZOOKEEPER" != "x" ]
then
	echo "Trying to autoconfigure zookeeper servers..."

	/configure-zookeeper-servers /opt/apache-storm/conf/storm.yaml
fi

if [ "x$STORM_CMD" != "x" ]
then

	echo "Running storm command ${STORM_CMD}"

	bin/storm ${STORM_CMD}

else

	echo "Nothing to run. Just waiting..."
	while true; do sleep 3; done

fi

