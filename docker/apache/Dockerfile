FROM php:7.2-apache

RUN apt-get update \
  && apt-get install -y \
    libapache2-mod-fcgid  \
  && apt-get clean

RUN  a2enmod rewrite
