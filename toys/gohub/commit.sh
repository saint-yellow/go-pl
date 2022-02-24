#!/usr/bin/env bash
# -*- coding: utf-8 -*-

git add .
git commit -m "$1"
git push -u origin-gitee master
git push -u origin-github master
