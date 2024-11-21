FROM debian:12.8

# установка пакетов
RUN echo 'debconf debconf/frontend select Noninteractive' | debconf-set-selections
RUN  apt-get update -y && \
  apt-get install -yq --no-install-recommends software-properties-common apt-utils bzip2 procps wget gnupg vim curl bash zip unzip && \
  apt-get update && \
  apt-get install -yq --no-install-recommends locales locales-all git && \
  rm -rf /var/lib/apt/lists/* && \
  rm -rf /var/cache/apt/archives/*
    
ENV LC_ALL ru_RU.UTF-8
ENV LANG ru_RU.UTF-8
ENV LANGUAGE ru_RU.UTF-8
ENV TZ Europe/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

#создание пользователя
RUN adduser --disabled-password --gecos '' sdl
#папка для программ
RUN chown sdl:sdl /opt

#переключение пользователя
USER sdl
# рабочая папки для установки компонент
WORKDIR /opt
RUN wget https://go.dev/dl/go1.23.3.linux-amd64.tar.gz && \
    tar -C /opt -xzf go1.23.3.linux-amd64.tar.gz && \
    rm -rf go1.23.3.linux-amd64.tar.gz

ENV PATH=$PATH:/opt/go/bin
ENV GOROOT=/opt/go

WORKDIR /home/sdl

CMD ["bash", "-c", "tail -f /dev/null"]
