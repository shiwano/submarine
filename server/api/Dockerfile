FROM ruby:2.3.1-alpine

RUN apk update && \
    apk upgrade && \
    apk add --no-cache bash git openssh build-base libxml2-dev libxslt-dev mysql-dev && \
    gem install bundler && \
    mkdir -p /src
WORKDIR /src

COPY Gemfile /src
COPY Gemfile.lock /src
RUN bundle install --deployment

COPY config.json /src
COPY . /src

ARG env=development
ENV RAILS_ENV $env
EXPOSE 3000
ENTRYPOINT ["bundle", "exec", "rails"]
CMD ["server", "-b", "0.0.0.0"]
