FROM scratch

WORKDIR /data/blog
COPY ./blog-go-server /data/blog

EXPOSE 8000
CMD ["./blog-go-server"]