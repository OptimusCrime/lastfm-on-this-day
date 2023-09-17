FROM php:7.4.11-fpm

# Last.fm API is in UTC, so this makes date stuff a lot easier
ENV TZ=Etc/UTC

ARG ENV=prod

RUN cd /usr/src \
    && curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer \
    && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo $TZ > /etc/timezone \
    && echo "date.timezone = $TZ" > /usr/local/etc/php/conf.d/timezone.ini \
    && apt-get update \
    && apt-get install -y --fix-missing apt-utils gnupg \
    && echo "deb http://packages.dotdeb.org jessie all" >> /etc/apt/sources.list \
    && echo "deb-src http://packages.dotdeb.org jessie all" >> /etc/apt/sources.list \
    && curl -sS --insecure https://www.dotdeb.org/dotdeb.gpg | apt-key add - \
    && apt-get update \
    && apt-get install -y zlib1g-dev libzip-dev zip \
    && docker-php-ext-install zip

RUN if [ $ENV = "prod" ] ; then \
    mv "$PHP_INI_DIR/php.ini-production" "$PHP_INI_DIR/php.ini" ; \
fi ;

COPY ./backend /code/site
