#!/usr/bin/env bash
#chmod ugo+x config.ssh
#ssh ubuntu@95.163.209.195
pwd
openssl aes-256-cbc -K $encrypted_123456789dfs_key -iv $encrypted_123456789dfs_iv -in ./.travis/deploy_key.enc -out ${TRAVIS_BUILD_DIR}/.travis/deploy_key -d
chmod 600 ${TRAVIS_BUILD_DIR}/.travis/deploy_key
mv ${TRAVIS_BUILD_DIR}/.travis/deploy_key ~/.ssh/id_rsa
cat ${TRAVIS_BUILD_DIR}/.travis/ssh_config >> ~/.ssh/config
