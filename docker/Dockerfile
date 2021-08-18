FROM debian:buster

LABEL maintainer "Clement DAI <dai.clement.pro@gmail.com>"

ENV GO_ANONYMIZE=go-anonymize-mysqldump_linux_amd64

RUN apt update
RUN apt install -y \
    git \
    vim \
    make \
    curl 

RUN curl https://dl.google.com/go/go1.16.7.linux-amd64.tar.gz --output go1.16.7.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.16.7.linux-amd64.tar.gz
RUN echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc

RUN git clone https://github.com/humanmade/go-anonymize-mysqldump.git

WORKDIR /cdai-test/go-anonymize-mysqldump

RUN apt install -y gox

RUN curl -OL https://github.com/humanmade/go-anonymize-mysqldump/releases/download/0.3.0/$GO_ANONYMIZE.gz
RUN gunzip $GO_ANONYMIZE.gz
RUN mv $GO_ANONYMIZE /usr/local/bin/anonymize-mysqldump
RUN chmod +x /usr/local/bin/anonymize-mysqldump



ENTRYPOINT ["tail", "-f", "/dev/null"]
