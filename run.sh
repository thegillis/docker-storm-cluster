#!/bin/bash

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

