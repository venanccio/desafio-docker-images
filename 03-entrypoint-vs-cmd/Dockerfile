FROM ubuntu:20.04
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt-get install -y ffmpeg && apt-get clean
COPY convert.sh /usr/local/bin/convert.sh
RUN chmod +x /usr/local/bin/convert.sh
ENTRYPOINT ["convert.sh"]
CMD ["--help"]
