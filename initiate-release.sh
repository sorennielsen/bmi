#!/usr/bin/env bash

set -eu

Echo "Starting release of BMI Calculator"

echo "${yellow}Version for new release:${normal}"
read -p "${bold}>${normal} " version

echo ""

echo "${yellow}Ready to continue with ${version}?${normal}"
echo "Press ${bold}Ctrl-C${normal} to abort, ${bold}ENTER${normal} to continue"

read

git tag -a "${version}" -m "Automated release initiated by $(whoami)"
git push origin refs/tags/${version}

echo ""
echo "${green}New version pushed to Git.${normal}"
echo ""
echo "Follow release process with below links:"
echo "GitHub releases: https://github.com/sorennielsen/bmi/releases"
echo "Drone builds: https://cloud.drone.io/sorennielsen/bmi/"
echo "Docker Hub: https://hub.docker.com/repository/docker/sorennielsen/bmi"


#read -p "Edit deployment for SLAYER Cloud? [${bold}yes${normal} to confirm] " confirm \
#	&& [[ $confirm == [yY] || $confirm == [yY][eE][sS] ]] || exit 1
#
#kubectl edit deployments.apps bmi
