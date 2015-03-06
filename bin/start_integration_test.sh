#!/bin/bash

# Need to first update the local repo before installing anything
apt-get -y update

# Install Git (TODO(CD): Remove this, for testing only)
apt-get -y install git

# Install the latest Marathon
apt-get -y install marathon=0.8.0-1.1.97.ubuntu1404

# List installed versions of external systems
dpkg -l marathon mesos zookeeper | grep '^ii'

# Start zookeeper
/usr/share/zookeeper/bin/zkServer.sh start

# Start Mesos master
/etc/init.d/mesos-master start

# Start Mesos slave
/etc/init.d/mesos-slave start

# Start Marathon
/etc/init.d/marathon start

# Give all of the processes above some time.
sleep 2

# Clean and recreate environment
cd /dcos-cli
make clean env

# Activate the virtual environment so that we can run make
source env/bin/activate

# Run the default target: E.g. test and package
make
