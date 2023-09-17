FROM nginx:1.25.2-alpine

ARG ENV=prod

COPY ./_docker/site.conf /etc/nginx/conf.d/site.conf

RUN chmod 644 /etc/nginx/conf.d/site.conf \
    && rm /etc/nginx/conf.d/default.conf

RUN if [ $ENV = "prod" ] ; then \
    sed -i 's/lastfm-php/lastfm-prod-php/g' /etc/nginx/conf.d/site.conf ; \
fi ;
